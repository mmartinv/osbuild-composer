/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
	"reflect"
)


// ArtifactsAPIService ArtifactsAPI service
type ArtifactsAPIService service

type ArtifactsAPIArtifactsCreateRequest struct {
	ctx context.Context
	ApiService *ArtifactsAPIService
	file *os.File
	size *int64
	md5 *string
	sha1 *string
	sha224 *string
	sha256 *string
	sha384 *string
	sha512 *string
}

// The stored file.
func (r ArtifactsAPIArtifactsCreateRequest) File(file *os.File) ArtifactsAPIArtifactsCreateRequest {
	r.file = file
	return r
}

// The size of the file in bytes.
func (r ArtifactsAPIArtifactsCreateRequest) Size(size int64) ArtifactsAPIArtifactsCreateRequest {
	r.size = &size
	return r
}

// The MD5 checksum of the file if available.
func (r ArtifactsAPIArtifactsCreateRequest) Md5(md5 string) ArtifactsAPIArtifactsCreateRequest {
	r.md5 = &md5
	return r
}

// The SHA-1 checksum of the file if available.
func (r ArtifactsAPIArtifactsCreateRequest) Sha1(sha1 string) ArtifactsAPIArtifactsCreateRequest {
	r.sha1 = &sha1
	return r
}

// The SHA-224 checksum of the file if available.
func (r ArtifactsAPIArtifactsCreateRequest) Sha224(sha224 string) ArtifactsAPIArtifactsCreateRequest {
	r.sha224 = &sha224
	return r
}

// The SHA-256 checksum of the file if available.
func (r ArtifactsAPIArtifactsCreateRequest) Sha256(sha256 string) ArtifactsAPIArtifactsCreateRequest {
	r.sha256 = &sha256
	return r
}

// The SHA-384 checksum of the file if available.
func (r ArtifactsAPIArtifactsCreateRequest) Sha384(sha384 string) ArtifactsAPIArtifactsCreateRequest {
	r.sha384 = &sha384
	return r
}

// The SHA-512 checksum of the file if available.
func (r ArtifactsAPIArtifactsCreateRequest) Sha512(sha512 string) ArtifactsAPIArtifactsCreateRequest {
	r.sha512 = &sha512
	return r
}

func (r ArtifactsAPIArtifactsCreateRequest) Execute() (*ArtifactResponse, *http.Response, error) {
	return r.ApiService.ArtifactsCreateExecute(r)
}

/*
ArtifactsCreate Create an artifact

A customized named ModelViewSet that knows how to register itself with the Pulp API router.

This viewset is discoverable by its name.
"Normal" Django Models and Master/Detail models are supported by the ``register_with`` method.

Attributes:
    lookup_field (str): The name of the field by which an object should be looked up, in
        addition to any parent lookups if this ViewSet is nested. Defaults to 'pk'
    endpoint_name (str): The name of the final path segment that should identify the ViewSet's
        collection endpoint.
    nest_prefix (str): Optional prefix under which this ViewSet should be nested. This must
        correspond to the "parent_prefix" of a router with rest_framework_nested.NestedMixin.
        None indicates this ViewSet should not be nested.
    parent_lookup_kwargs (dict): Optional mapping of key names that would appear in self.kwargs
        to django model filter expressions that can be used with the corresponding value from
        self.kwargs, used only by a nested ViewSet to filter based on the parent object's
        identity.
    schema (DefaultSchema): The schema class to use by default in a viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ArtifactsAPIArtifactsCreateRequest
*/
func (a *ArtifactsAPIService) ArtifactsCreate(ctx context.Context) ArtifactsAPIArtifactsCreateRequest {
	return ArtifactsAPIArtifactsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ArtifactResponse
func (a *ArtifactsAPIService) ArtifactsCreateExecute(r ArtifactsAPIArtifactsCreateRequest) (*ArtifactResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsAPIService.ArtifactsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/artifacts/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "size", r.size, "")
	}
	if r.md5 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "md5", r.md5, "")
	}
	if r.sha1 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha1", r.sha1, "")
	}
	if r.sha224 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha224", r.sha224, "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha256", r.sha256, "")
	}
	if r.sha384 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha384", r.sha384, "")
	}
	if r.sha512 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha512", r.sha512, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ArtifactsAPIArtifactsDeleteRequest struct {
	ctx context.Context
	ApiService *ArtifactsAPIService
	artifactHref string
}

func (r ArtifactsAPIArtifactsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ArtifactsDeleteExecute(r)
}

/*
ArtifactsDelete Delete an artifact

Remove Artifact only if it is not associated with any Content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param artifactHref
 @return ArtifactsAPIArtifactsDeleteRequest
*/
func (a *ArtifactsAPIService) ArtifactsDelete(ctx context.Context, artifactHref string) ArtifactsAPIArtifactsDeleteRequest {
	return ArtifactsAPIArtifactsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		artifactHref: artifactHref,
	}
}

// Execute executes the request
func (a *ArtifactsAPIService) ArtifactsDeleteExecute(r ArtifactsAPIArtifactsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsAPIService.ArtifactsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{artifact_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"artifact_href"+"}", parameterValueToString(r.artifactHref, "artifactHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ArtifactsAPIArtifactsListRequest struct {
	ctx context.Context
	ApiService *ArtifactsAPIService
	limit *int32
	md5 *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	repositoryVersion *string
	sha1 *string
	sha224 *string
	sha256 *string
	sha384 *string
	sha512 *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ArtifactsAPIArtifactsListRequest) Limit(limit int32) ArtifactsAPIArtifactsListRequest {
	r.limit = &limit
	return r
}

// Filter results where md5 matches value
func (r ArtifactsAPIArtifactsListRequest) Md5(md5 string) ArtifactsAPIArtifactsListRequest {
	r.md5 = &md5
	return r
}

// The initial index from which to return the results.
func (r ArtifactsAPIArtifactsListRequest) Offset(offset int32) ArtifactsAPIArtifactsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;file&#x60; - File * &#x60;-file&#x60; - File (descending) * &#x60;size&#x60; - Size * &#x60;-size&#x60; - Size (descending) * &#x60;md5&#x60; - Md5 * &#x60;-md5&#x60; - Md5 (descending) * &#x60;sha1&#x60; - Sha1 * &#x60;-sha1&#x60; - Sha1 (descending) * &#x60;sha224&#x60; - Sha224 * &#x60;-sha224&#x60; - Sha224 (descending) * &#x60;sha256&#x60; - Sha256 * &#x60;-sha256&#x60; - Sha256 (descending) * &#x60;sha384&#x60; - Sha384 * &#x60;-sha384&#x60; - Sha384 (descending) * &#x60;sha512&#x60; - Sha512 * &#x60;-sha512&#x60; - Sha512 (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ArtifactsAPIArtifactsListRequest) Ordering(ordering []string) ArtifactsAPIArtifactsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ArtifactsAPIArtifactsListRequest) PulpHrefIn(pulpHrefIn []string) ArtifactsAPIArtifactsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ArtifactsAPIArtifactsListRequest) PulpIdIn(pulpIdIn []string) ArtifactsAPIArtifactsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ArtifactsAPIArtifactsListRequest) RepositoryVersion(repositoryVersion string) ArtifactsAPIArtifactsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Filter results where sha1 matches value
func (r ArtifactsAPIArtifactsListRequest) Sha1(sha1 string) ArtifactsAPIArtifactsListRequest {
	r.sha1 = &sha1
	return r
}

// Filter results where sha224 matches value
func (r ArtifactsAPIArtifactsListRequest) Sha224(sha224 string) ArtifactsAPIArtifactsListRequest {
	r.sha224 = &sha224
	return r
}

// Filter results where sha256 matches value
func (r ArtifactsAPIArtifactsListRequest) Sha256(sha256 string) ArtifactsAPIArtifactsListRequest {
	r.sha256 = &sha256
	return r
}

// Filter results where sha384 matches value
func (r ArtifactsAPIArtifactsListRequest) Sha384(sha384 string) ArtifactsAPIArtifactsListRequest {
	r.sha384 = &sha384
	return r
}

// Filter results where sha512 matches value
func (r ArtifactsAPIArtifactsListRequest) Sha512(sha512 string) ArtifactsAPIArtifactsListRequest {
	r.sha512 = &sha512
	return r
}

// A list of fields to include in the response.
func (r ArtifactsAPIArtifactsListRequest) Fields(fields []string) ArtifactsAPIArtifactsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ArtifactsAPIArtifactsListRequest) ExcludeFields(excludeFields []string) ArtifactsAPIArtifactsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ArtifactsAPIArtifactsListRequest) Execute() (*PaginatedArtifactResponseList, *http.Response, error) {
	return r.ApiService.ArtifactsListExecute(r)
}

/*
ArtifactsList List artifacts

A customized named ModelViewSet that knows how to register itself with the Pulp API router.

This viewset is discoverable by its name.
"Normal" Django Models and Master/Detail models are supported by the ``register_with`` method.

Attributes:
    lookup_field (str): The name of the field by which an object should be looked up, in
        addition to any parent lookups if this ViewSet is nested. Defaults to 'pk'
    endpoint_name (str): The name of the final path segment that should identify the ViewSet's
        collection endpoint.
    nest_prefix (str): Optional prefix under which this ViewSet should be nested. This must
        correspond to the "parent_prefix" of a router with rest_framework_nested.NestedMixin.
        None indicates this ViewSet should not be nested.
    parent_lookup_kwargs (dict): Optional mapping of key names that would appear in self.kwargs
        to django model filter expressions that can be used with the corresponding value from
        self.kwargs, used only by a nested ViewSet to filter based on the parent object's
        identity.
    schema (DefaultSchema): The schema class to use by default in a viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ArtifactsAPIArtifactsListRequest
*/
func (a *ArtifactsAPIService) ArtifactsList(ctx context.Context) ArtifactsAPIArtifactsListRequest {
	return ArtifactsAPIArtifactsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedArtifactResponseList
func (a *ArtifactsAPIService) ArtifactsListExecute(r ArtifactsAPIArtifactsListRequest) (*PaginatedArtifactResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedArtifactResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsAPIService.ArtifactsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/artifacts/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.md5 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "md5", r.md5, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.sha1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha1", r.sha1, "")
	}
	if r.sha224 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha224", r.sha224, "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "")
	}
	if r.sha384 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha384", r.sha384, "")
	}
	if r.sha512 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha512", r.sha512, "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ArtifactsAPIArtifactsReadRequest struct {
	ctx context.Context
	ApiService *ArtifactsAPIService
	artifactHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ArtifactsAPIArtifactsReadRequest) Fields(fields []string) ArtifactsAPIArtifactsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ArtifactsAPIArtifactsReadRequest) ExcludeFields(excludeFields []string) ArtifactsAPIArtifactsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ArtifactsAPIArtifactsReadRequest) Execute() (*ArtifactResponse, *http.Response, error) {
	return r.ApiService.ArtifactsReadExecute(r)
}

/*
ArtifactsRead Inspect an artifact

A customized named ModelViewSet that knows how to register itself with the Pulp API router.

This viewset is discoverable by its name.
"Normal" Django Models and Master/Detail models are supported by the ``register_with`` method.

Attributes:
    lookup_field (str): The name of the field by which an object should be looked up, in
        addition to any parent lookups if this ViewSet is nested. Defaults to 'pk'
    endpoint_name (str): The name of the final path segment that should identify the ViewSet's
        collection endpoint.
    nest_prefix (str): Optional prefix under which this ViewSet should be nested. This must
        correspond to the "parent_prefix" of a router with rest_framework_nested.NestedMixin.
        None indicates this ViewSet should not be nested.
    parent_lookup_kwargs (dict): Optional mapping of key names that would appear in self.kwargs
        to django model filter expressions that can be used with the corresponding value from
        self.kwargs, used only by a nested ViewSet to filter based on the parent object's
        identity.
    schema (DefaultSchema): The schema class to use by default in a viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param artifactHref
 @return ArtifactsAPIArtifactsReadRequest
*/
func (a *ArtifactsAPIService) ArtifactsRead(ctx context.Context, artifactHref string) ArtifactsAPIArtifactsReadRequest {
	return ArtifactsAPIArtifactsReadRequest{
		ApiService: a,
		ctx: ctx,
		artifactHref: artifactHref,
	}
}

// Execute executes the request
//  @return ArtifactResponse
func (a *ArtifactsAPIService) ArtifactsReadExecute(r ArtifactsAPIArtifactsReadRequest) (*ArtifactResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsAPIService.ArtifactsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{artifact_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"artifact_href"+"}", parameterValueToString(r.artifactHref, "artifactHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}