// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newPublicAdvertisedPrefixesClientHook clientHook

// PublicAdvertisedPrefixesCallOptions contains the retry settings for each method of PublicAdvertisedPrefixesClient.
type PublicAdvertisedPrefixesCallOptions struct {
	Announce []gax.CallOption
	Delete   []gax.CallOption
	Get      []gax.CallOption
	Insert   []gax.CallOption
	List     []gax.CallOption
	Patch    []gax.CallOption
	Withdraw []gax.CallOption
}

func defaultPublicAdvertisedPrefixesRESTCallOptions() *PublicAdvertisedPrefixesCallOptions {
	return &PublicAdvertisedPrefixesCallOptions{
		Announce: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		Delete: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		Get: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		Insert: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		List: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		Patch: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		Withdraw: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
	}
}

// internalPublicAdvertisedPrefixesClient is an interface that defines the methods available from Google Compute Engine API.
type internalPublicAdvertisedPrefixesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Announce(context.Context, *computepb.AnnouncePublicAdvertisedPrefixeRequest, ...gax.CallOption) (*Operation, error)
	Delete(context.Context, *computepb.DeletePublicAdvertisedPrefixeRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetPublicAdvertisedPrefixeRequest, ...gax.CallOption) (*computepb.PublicAdvertisedPrefix, error)
	Insert(context.Context, *computepb.InsertPublicAdvertisedPrefixeRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListPublicAdvertisedPrefixesRequest, ...gax.CallOption) *PublicAdvertisedPrefixIterator
	Patch(context.Context, *computepb.PatchPublicAdvertisedPrefixeRequest, ...gax.CallOption) (*Operation, error)
	Withdraw(context.Context, *computepb.WithdrawPublicAdvertisedPrefixeRequest, ...gax.CallOption) (*Operation, error)
}

// PublicAdvertisedPrefixesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The PublicAdvertisedPrefixes API.
type PublicAdvertisedPrefixesClient struct {
	// The internal transport-dependent client.
	internalClient internalPublicAdvertisedPrefixesClient

	// The call options for this service.
	CallOptions *PublicAdvertisedPrefixesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PublicAdvertisedPrefixesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PublicAdvertisedPrefixesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PublicAdvertisedPrefixesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Announce announces the specified PublicAdvertisedPrefix
func (c *PublicAdvertisedPrefixesClient) Announce(ctx context.Context, req *computepb.AnnouncePublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Announce(ctx, req, opts...)
}

// Delete deletes the specified PublicAdvertisedPrefix
func (c *PublicAdvertisedPrefixesClient) Delete(ctx context.Context, req *computepb.DeletePublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified PublicAdvertisedPrefix resource.
func (c *PublicAdvertisedPrefixesClient) Get(ctx context.Context, req *computepb.GetPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.PublicAdvertisedPrefix, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a PublicAdvertisedPrefix in the specified project using the parameters that are included in the request.
func (c *PublicAdvertisedPrefixesClient) Insert(ctx context.Context, req *computepb.InsertPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List lists the PublicAdvertisedPrefixes for a project.
func (c *PublicAdvertisedPrefixesClient) List(ctx context.Context, req *computepb.ListPublicAdvertisedPrefixesRequest, opts ...gax.CallOption) *PublicAdvertisedPrefixIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified Router resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *PublicAdvertisedPrefixesClient) Patch(ctx context.Context, req *computepb.PatchPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// Withdraw withdraws the specified PublicAdvertisedPrefix
func (c *PublicAdvertisedPrefixesClient) Withdraw(ctx context.Context, req *computepb.WithdrawPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Withdraw(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type publicAdvertisedPrefixesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// operationClient is used to call the operation-specific management service.
	operationClient *GlobalOperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing PublicAdvertisedPrefixesClient
	CallOptions **PublicAdvertisedPrefixesCallOptions
}

// NewPublicAdvertisedPrefixesRESTClient creates a new public advertised prefixes rest client.
//
// The PublicAdvertisedPrefixes API.
func NewPublicAdvertisedPrefixesRESTClient(ctx context.Context, opts ...option.ClientOption) (*PublicAdvertisedPrefixesClient, error) {
	clientOpts := append(defaultPublicAdvertisedPrefixesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultPublicAdvertisedPrefixesRESTCallOptions()
	c := &publicAdvertisedPrefixesRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	o := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opC, err := NewGlobalOperationsRESTClient(ctx, o...)
	if err != nil {
		return nil, err
	}
	c.operationClient = opC

	return &PublicAdvertisedPrefixesClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultPublicAdvertisedPrefixesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://compute.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *publicAdvertisedPrefixesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *publicAdvertisedPrefixesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	if err := c.operationClient.Close(); err != nil {
		return err
	}
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *publicAdvertisedPrefixesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// Announce announces the specified PublicAdvertisedPrefix
func (c *publicAdvertisedPrefixesRESTClient) Announce(ctx context.Context, req *computepb.AnnouncePublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/publicAdvertisedPrefixes/%v/announce", req.GetProject(), req.GetPublicAdvertisedPrefix())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_advertised_prefix", url.QueryEscape(req.GetPublicAdvertisedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Announce[0:len((*c.CallOptions).Announce):len((*c.CallOptions).Announce)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&globalOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
		},
	}
	return op, nil
}

// Delete deletes the specified PublicAdvertisedPrefix
func (c *publicAdvertisedPrefixesRESTClient) Delete(ctx context.Context, req *computepb.DeletePublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/publicAdvertisedPrefixes/%v", req.GetProject(), req.GetPublicAdvertisedPrefix())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_advertised_prefix", url.QueryEscape(req.GetPublicAdvertisedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&globalOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
		},
	}
	return op, nil
}

// Get returns the specified PublicAdvertisedPrefix resource.
func (c *publicAdvertisedPrefixesRESTClient) Get(ctx context.Context, req *computepb.GetPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.PublicAdvertisedPrefix, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/publicAdvertisedPrefixes/%v", req.GetProject(), req.GetPublicAdvertisedPrefix())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_advertised_prefix", url.QueryEscape(req.GetPublicAdvertisedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.PublicAdvertisedPrefix{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// Insert creates a PublicAdvertisedPrefix in the specified project using the parameters that are included in the request.
func (c *publicAdvertisedPrefixesRESTClient) Insert(ctx context.Context, req *computepb.InsertPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetPublicAdvertisedPrefixResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/publicAdvertisedPrefixes", req.GetProject())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&globalOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
		},
	}
	return op, nil
}

// List lists the PublicAdvertisedPrefixes for a project.
func (c *publicAdvertisedPrefixesRESTClient) List(ctx context.Context, req *computepb.ListPublicAdvertisedPrefixesRequest, opts ...gax.CallOption) *PublicAdvertisedPrefixIterator {
	it := &PublicAdvertisedPrefixIterator{}
	req = proto.Clone(req).(*computepb.ListPublicAdvertisedPrefixesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.PublicAdvertisedPrefix, string, error) {
		resp := &computepb.PublicAdvertisedPrefixList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(uint32(math.MaxInt32))
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/publicAdvertisedPrefixes", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := io.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Patch patches the specified Router resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *publicAdvertisedPrefixesRESTClient) Patch(ctx context.Context, req *computepb.PatchPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetPublicAdvertisedPrefixResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/publicAdvertisedPrefixes/%v", req.GetProject(), req.GetPublicAdvertisedPrefix())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_advertised_prefix", url.QueryEscape(req.GetPublicAdvertisedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&globalOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
		},
	}
	return op, nil
}

// Withdraw withdraws the specified PublicAdvertisedPrefix
func (c *publicAdvertisedPrefixesRESTClient) Withdraw(ctx context.Context, req *computepb.WithdrawPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/publicAdvertisedPrefixes/%v/withdraw", req.GetProject(), req.GetPublicAdvertisedPrefix())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_advertised_prefix", url.QueryEscape(req.GetPublicAdvertisedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Withdraw[0:len((*c.CallOptions).Withdraw):len((*c.CallOptions).Withdraw)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&globalOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
		},
	}
	return op, nil
}
