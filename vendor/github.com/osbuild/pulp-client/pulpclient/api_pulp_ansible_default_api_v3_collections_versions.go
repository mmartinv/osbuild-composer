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
	"reflect"
)


// PulpAnsibleDefaultApiV3CollectionsVersionsAPIService PulpAnsibleDefaultApiV3CollectionsVersionsAPI service
type PulpAnsibleDefaultApiV3CollectionsVersionsAPIService service

type PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDeleteRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleDefaultApiV3CollectionsVersionsAPIService
	name string
	namespace string
	version string
}

func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDeleteExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDelete Method for PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDelete

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param version
 @return PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDeleteRequest

Deprecated
*/
func (a *PulpAnsibleDefaultApiV3CollectionsVersionsAPIService) PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDelete(ctx context.Context, name string, namespace string, version string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDeleteRequest {
	return PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		version: version,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
// Deprecated
func (a *PulpAnsibleDefaultApiV3CollectionsVersionsAPIService) PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDeleteExecute(r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleDefaultApiV3CollectionsVersionsAPIService.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/collections/{namespace}/{name}/versions/{version}/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", parameterValueToString(r.name, "name"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", parameterValueToString(r.namespace, "namespace"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", parameterValueToString(r.version, "version"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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

type PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleDefaultApiV3CollectionsVersionsAPIService
	name string
	namespace string
	isHighest *bool
	limit *int32
	name2 *string
	namespace2 *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	tags *string
	version *string
	fields *[]string
	excludeFields *[]string
}

func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) IsHighest(isHighest bool) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.isHighest = &isHighest
	return r
}

// Number of results to return per page.
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) Limit(limit int32) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.limit = &limit
	return r
}

func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) Name2(name2 string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.name2 = &name2
	return r
}

func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) Namespace2(namespace2 string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.namespace2 = &namespace2
	return r
}

// The initial index from which to return the results.
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) Offset(offset int32) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;authors&#x60; - Authors * &#x60;-authors&#x60; - Authors (descending) * &#x60;contents&#x60; - Contents * &#x60;-contents&#x60; - Contents (descending) * &#x60;dependencies&#x60; - Dependencies * &#x60;-dependencies&#x60; - Dependencies (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;docs_blob&#x60; - Docs blob * &#x60;-docs_blob&#x60; - Docs blob (descending) * &#x60;manifest&#x60; - Manifest * &#x60;-manifest&#x60; - Manifest (descending) * &#x60;files&#x60; - Files * &#x60;-files&#x60; - Files (descending) * &#x60;documentation&#x60; - Documentation * &#x60;-documentation&#x60; - Documentation (descending) * &#x60;homepage&#x60; - Homepage * &#x60;-homepage&#x60; - Homepage (descending) * &#x60;issues&#x60; - Issues * &#x60;-issues&#x60; - Issues (descending) * &#x60;license&#x60; - License * &#x60;-license&#x60; - License (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;namespace&#x60; - Namespace * &#x60;-namespace&#x60; - Namespace (descending) * &#x60;repository&#x60; - Repository * &#x60;-repository&#x60; - Repository (descending) * &#x60;requires_ansible&#x60; - Requires ansible * &#x60;-requires_ansible&#x60; - Requires ansible (descending) * &#x60;version&#x60; - Version * &#x60;-version&#x60; - Version (descending) * &#x60;version_major&#x60; - Version major * &#x60;-version_major&#x60; - Version major (descending) * &#x60;version_minor&#x60; - Version minor * &#x60;-version_minor&#x60; - Version minor (descending) * &#x60;version_patch&#x60; - Version patch * &#x60;-version_patch&#x60; - Version patch (descending) * &#x60;version_prerelease&#x60; - Version prerelease * &#x60;-version_prerelease&#x60; - Version prerelease (descending) * &#x60;is_highest&#x60; - Is highest * &#x60;-is_highest&#x60; - Is highest (descending) * &#x60;search_vector&#x60; - Search vector * &#x60;-search_vector&#x60; - Search vector (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) Ordering(ordering []string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) PulpHrefIn(pulpHrefIn []string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) PulpIdIn(pulpIdIn []string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) Q(q string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) RepositoryVersion(repositoryVersion string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter by comma separate list of tags that must all be matched
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) Tags(tags string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.tags = &tags
	return r
}

// Filter results where version matches value
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) Version(version string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.version = &version
	return r
}

// A list of fields to include in the response.
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) Fields(fields []string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) ExcludeFields(excludeFields []string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) Execute() (*PaginatedCollectionVersionListResponseList, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsList Method for PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsList

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @return PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest

Deprecated
*/
func (a *PulpAnsibleDefaultApiV3CollectionsVersionsAPIService) PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsList(ctx context.Context, name string, namespace string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest {
	return PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
	}
}

// Execute executes the request
//  @return PaginatedCollectionVersionListResponseList
// Deprecated
func (a *PulpAnsibleDefaultApiV3CollectionsVersionsAPIService) PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListExecute(r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsListRequest) (*PaginatedCollectionVersionListResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCollectionVersionListResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleDefaultApiV3CollectionsVersionsAPIService.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/collections/{namespace}/{name}/versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", parameterValueToString(r.name, "name"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", parameterValueToString(r.namespace, "namespace"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isHighest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_highest", r.isHighest, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name2, "")
	}
	if r.namespace2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace2, "")
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
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
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

type PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleDefaultApiV3CollectionsVersionsAPIService
	name string
	namespace string
	version string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadRequest) Fields(fields []string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadRequest) ExcludeFields(excludeFields []string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadRequest) Execute() (*CollectionVersionResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsRead Method for PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsRead

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param version
 @return PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadRequest

Deprecated
*/
func (a *PulpAnsibleDefaultApiV3CollectionsVersionsAPIService) PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsRead(ctx context.Context, name string, namespace string, version string) PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadRequest {
	return PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		version: version,
	}
}

// Execute executes the request
//  @return CollectionVersionResponse
// Deprecated
func (a *PulpAnsibleDefaultApiV3CollectionsVersionsAPIService) PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadExecute(r PulpAnsibleDefaultApiV3CollectionsVersionsAPIPulpAnsibleGalaxyDefaultApiV3CollectionsVersionsReadRequest) (*CollectionVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleDefaultApiV3CollectionsVersionsAPIService.PulpAnsibleGalaxyDefaultApiV3CollectionsVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/collections/{namespace}/{name}/versions/{version}/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", parameterValueToString(r.name, "name"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", parameterValueToString(r.namespace, "namespace"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", parameterValueToString(r.version, "version"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

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
