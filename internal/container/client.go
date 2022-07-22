// package container implements a client for a container
// registry. It can be used to upload container images.
package container

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/containers/image/v5/docker/archive"
	_ "github.com/containers/image/v5/oci/archive"
	_ "github.com/containers/image/v5/oci/layout"
	"github.com/osbuild/osbuild-composer/internal/common"

	"github.com/containers/common/pkg/retry"
	"github.com/containers/image/v5/copy"
	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/docker/reference"
	"github.com/containers/image/v5/manifest"
	"github.com/containers/image/v5/signature"
	"github.com/containers/image/v5/transports"
	"github.com/containers/image/v5/types"
	"github.com/opencontainers/go-digest"

	imgspecv1 "github.com/opencontainers/image-spec/specs-go/v1"
)

const (
	DefaultUserAgent  = "osbuild-composer/1.0"
	DefaultPolicyPath = "/etc/containers/policy.json"
)

var (
	defaultAuthFile = GetDefaultAuthFile()
)

// GetDefaultAuthFile returns the authentication file to use for the
// current environment.
//
// This is basically a re-implementation of `getPathToAuthWithOS` from
// containers/image/pkg/docker/config/config.go[1], but we ensure that
// the returned path is either accessible or nonexistent. This is needed
// since any other error than os.ErrNotExist will lead to an overall
// failure and thus prevent any operation even with public resources.
//
// [1] https://github.com/containers/image/blob/55ea76c7db702ed1af60924a0b57c8da533d9e5a/pkg/docker/config/config.go#L506
func GetDefaultAuthFile() string {

	dirExistsOrEmpty := func(path string) bool {
		_, err := os.Stat(path)
		// document the error case
		return err == nil || os.IsNotExist(err)
	}

	if runtimeDir := os.Getenv("XDG_RUNTIME_DIR"); runtimeDir != "" {
		path := filepath.Join(runtimeDir, "containers", "auth.json")
		if dirExistsOrEmpty(path) {
			return path
		}
	}

	path := fmt.Sprintf("/run/containers/%d/auth.json", os.Getuid())
	if dirExistsOrEmpty(path) {
		return path
	}

	return filepath.Join("var", "empty", "containers-auth.json")
}

// A Client to interact with the given Target object at a
// container registry, like e.g. uploading an image to.
// All mentioned defaults are only set when using the
// NewClient constructor.
type Client struct {
	Target reference.Named // the target object to interact with

	ReportWriter io.Writer // used for writing status reports, defaults to os.Stdout

	PrecomputeDigests bool // precompute digest in order to avoid uploads
	MaxRetries        int  // how often to retry http requests

	UserAgent string // user agent string to use for requests, defaults to DefaultUserAgent

	// internal state
	policy *signature.Policy
	sysCtx *types.SystemContext
}

// NewClient constructs a new Client for target with default options.
// It will add the "latest" tag if target does not contain it.
func NewClient(target string) (*Client, error) {

	ref, err := reference.ParseNormalizedNamed(target)
	if err != nil {
		return nil, fmt.Errorf("failed to parse '%s': %w", target, err)
	}

	var policy *signature.Policy
	if _, err := os.Stat(DefaultPolicyPath); err == nil {
		policy, err = signature.NewPolicyFromFile(DefaultPolicyPath)
		if err != nil {
			return nil, err
		}
	} else {
		policy = &signature.Policy{
			Default: []signature.PolicyRequirement{
				signature.NewPRInsecureAcceptAnything(),
			},
		}
	}

	if err != nil {
		return nil, err
	}

	client := Client{
		Target: reference.TagNameOnly(ref),

		ReportWriter:      os.Stdout,
		PrecomputeDigests: true,

		UserAgent: DefaultUserAgent,

		sysCtx: &types.SystemContext{
			RegistriesDirPath:        "",
			SystemRegistriesConfPath: "",
			BigFilesTemporaryDir:     "/var/tmp",

			OSChoice: "linux",

			AuthFilePath: defaultAuthFile,
		},
		policy: policy,
	}

	return &client, nil
}

// SetAuthFilePath sets the location of the `containers-auth.json(5)` file.
func (cl *Client) SetAuthFilePath(path string) {
	cl.sysCtx.AuthFilePath = path
}

// GetAuthFilePath gets the location of the `containers-auth.json(5)` file.
func (cl *Client) GetAuthFilePath() string {
	return cl.sysCtx.AuthFilePath
}

func (cl *Client) SetArchitectureChoice(arch string) {
	// Translate some well-known Composer architecture strings
	// into the corresponding container ones

	variant := ""

	switch arch {
	case "x86_64":
		arch = "amd64"

	case "aarch64":
		arch = "arm64"
		if variant == "" {
			variant = "v8"
		}

	case "armhfp":
		arch = "arm"
		if variant == "" {
			variant = "v7"
		}

		//ppc64le and s390x are the same
	}

	cl.sysCtx.ArchitectureChoice = arch
	cl.sysCtx.VariantChoice = variant
}

func (cl *Client) SetVariantChoice(variant string) {
	cl.sysCtx.VariantChoice = variant
}

// SetCredentials will set username and password for Client
func (cl *Client) SetCredentials(username, password string) {

	if cl.sysCtx.DockerAuthConfig == nil {
		cl.sysCtx.DockerAuthConfig = &types.DockerAuthConfig{}
	}

	cl.sysCtx.DockerAuthConfig.Username = username
	cl.sysCtx.DockerAuthConfig.Password = password
}

// SetSkipTLSVerify controls if TLS verification happens when
// making requests. If nil is passed it falls back to the default.
func (cl *Client) SetTLSVerify(verify *bool) {
	if verify == nil {
		cl.sysCtx.DockerInsecureSkipTLSVerify = types.OptionalBoolUndefined
	} else {
		cl.sysCtx.DockerInsecureSkipTLSVerify = types.NewOptionalBool(!*verify)
	}
}

// GetSkipTLSVerify returns current TLS verification state.
func (cl *Client) GetTLSVerify() *bool {

	skip := cl.sysCtx.DockerInsecureSkipTLSVerify

	if skip == types.OptionalBoolUndefined {
		return nil
	}

	// NB: we invert the state, i.e. verify == (skip == false)
	return common.BoolToPtr(skip == types.OptionalBoolFalse)
}

// SkipTLSVerify is a convenience helper that internally calls
// SetTLSVerify with false
func (cl *Client) SkipTLSVerify() {
	cl.SetTLSVerify(common.BoolToPtr(false))
}

func parseImageName(name string) (types.ImageReference, error) {

	parts := strings.SplitN(name, ":", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid image name '%s'", name)
	}

	transport := transports.Get(parts[0])
	if transport == nil {
		return nil, fmt.Errorf("unknown transport '%s'", parts[0])
	}

	return transport.ParseReference(parts[1])
}

// UploadImage takes an container image located at from and uploads it
// to the Target of Client. If tag is set, i.e. not the empty string,
// it will replace any previously set tag or digest of the target.
// Returns the digest of the manifest that was written to the server.
func (cl *Client) UploadImage(ctx context.Context, from, tag string) (digest.Digest, error) {

	targetCtx := *cl.sysCtx
	targetCtx.DockerRegistryPushPrecomputeDigests = cl.PrecomputeDigests

	policyContext, err := signature.NewPolicyContext(cl.policy)

	if err != nil {
		return "", err
	}

	srcRef, err := parseImageName(from)
	if err != nil {
		return "", fmt.Errorf("invalid source name '%s': %w", from, err)
	}

	target := cl.Target

	if tag != "" {
		target = reference.TrimNamed(target)
		target, err = reference.WithTag(target, tag)
		if err != nil {
			return "", fmt.Errorf("error creating reference with tag '%s': %w", tag, err)
		}
	}

	destRef, err := docker.NewReference(target)
	if err != nil {
		return "", err
	}

	retryOpts := retry.RetryOptions{
		MaxRetry: cl.MaxRetries,
	}

	var manifestDigest digest.Digest

	err = retry.RetryIfNecessary(ctx, func() error {
		manifestBytes, err := copy.Image(ctx, policyContext, destRef, srcRef, &copy.Options{
			RemoveSignatures:      false,
			SignBy:                "",
			SignPassphrase:        "",
			ReportWriter:          cl.ReportWriter,
			SourceCtx:             cl.sysCtx,
			DestinationCtx:        &targetCtx,
			ForceManifestMIMEType: "",
			ImageListSelection:    copy.CopyAllImages,
			PreserveDigests:       false,
		})

		if err != nil {
			return err
		}

		manifestDigest, err = manifest.Digest(manifestBytes)

		return err

	}, &retryOpts)

	if err != nil {
		return "", err
	}

	return manifestDigest, nil
}

// A RawManifest contains the raw manifest Data and its MimeType
type RawManifest struct {
	Data     []byte
	MimeType string
}

// Digest computes the digest from the raw manifest data
func (m RawManifest) Digest() (digest.Digest, error) {
	return manifest.Digest(m.Data)
}

// GetManifest fetches the raw manifest data from the server. If digest is not empty
// it will override any given tag for the Client's Target.
func (cl *Client) GetManifest(ctx context.Context, digest digest.Digest) (r RawManifest, err error) {
	target := cl.Target

	if digest != "" {
		t := reference.TrimNamed(cl.Target)
		t, err = reference.WithDigest(t, digest)
		if err != nil {
			return
		}

		target = t
	}

	ref, err := docker.NewReference(target)
	if err != nil {
		return
	}

	src, err := ref.NewImageSource(ctx, cl.sysCtx)

	if err != nil {
		return
	}

	defer func() {
		if e := src.Close(); e != nil {
			err = fmt.Errorf("could not close image: %w", e)
		}
	}()

	retryOpts := retry.RetryOptions{
		MaxRetry: cl.MaxRetries,
	}

	if err = retry.RetryIfNecessary(ctx, func() error {
		r.Data, r.MimeType, err = src.GetManifest(ctx, nil)
		return err
	}, &retryOpts); err != nil {
		return
	}

	return
}

type manifestList interface {
	ChooseInstance(ctx *types.SystemContext) (digest.Digest, error)
}

type resolvedIds struct {
	Manifest digest.Digest
	Config   digest.Digest
}

func (cl *Client) resolveManifestList(ctx context.Context, list manifestList) (resolvedIds, error) {
	digest, err := list.ChooseInstance(cl.sysCtx)
	if err != nil {
		return resolvedIds{}, err
	}

	raw, err := cl.GetManifest(ctx, digest)

	if err != nil {
		return resolvedIds{}, fmt.Errorf("error getting manifest: %w", err)
	}

	if err != nil {
		return resolvedIds{}, nil
	}

	return cl.resolveRawManifest(ctx, raw)
}

func (cl *Client) resolveRawManifest(ctx context.Context, rm RawManifest) (resolvedIds, error) {

	var imageID digest.Digest

	switch rm.MimeType {
	case manifest.DockerV2ListMediaType:
		list, err := manifest.Schema2ListFromManifest(rm.Data)
		if err != nil {
			return resolvedIds{}, err
		}

		return cl.resolveManifestList(ctx, list)

	case imgspecv1.MediaTypeImageIndex:
		index, err := manifest.OCI1IndexFromManifest(rm.Data)
		if err != nil {
			return resolvedIds{}, err
		}

		return cl.resolveManifestList(ctx, index)

	case imgspecv1.MediaTypeImageManifest:
		m, err := manifest.OCI1FromManifest(rm.Data)
		if err != nil {
			return resolvedIds{}, nil
		}
		imageID = m.ConfigInfo().Digest

	case manifest.DockerV2Schema2MediaType:
		m, err := manifest.Schema2FromManifest(rm.Data)

		if err != nil {
			return resolvedIds{}, nil
		}

		imageID = m.ConfigInfo().Digest

	default:
		return resolvedIds{}, fmt.Errorf("unsupported manifest format '%s'", rm.MimeType)
	}

	dg, err := rm.Digest()

	if err != nil {
		return resolvedIds{}, err
	}

	return resolvedIds{
		Manifest: dg,
		Config:   imageID,
	}, nil
}

// Resolve the Client's Target to the manifest digest and the corresponding image id
// which is the digest of the configuration object. It uses the architecture and
// variant specified via SetArchitectureChoice or the corresponding defaults for
// the host.
func (cl *Client) Resolve(ctx context.Context, name string) (Spec, error) {

	raw, err := cl.GetManifest(ctx, "")

	if err != nil {
		return Spec{}, fmt.Errorf("error getting manifest: %w", err)
	}

	ids, err := cl.resolveRawManifest(ctx, raw)
	if err != nil {
		return Spec{}, err
	}

	spec := NewSpec(cl.Target, ids.Manifest, ids.Config)
	spec.TLSVerify = cl.GetTLSVerify()

	return spec, nil
}
