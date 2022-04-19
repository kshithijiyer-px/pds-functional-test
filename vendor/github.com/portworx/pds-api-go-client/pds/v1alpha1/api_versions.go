/*
PDS API

Portworx Data Services API Server

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// VersionsApiService VersionsApi service
type VersionsApiService service

type ApiApiDataServicesIdVersionsGetRequest struct {
	ctx context.Context
	ApiService *VersionsApiService
	id string
	sortBy *string
	limit *string
	continuation *string
	id2 *string
	name *string
}

// A given Version attribute to sort results by (one of: id, name, created_at)
func (r ApiApiDataServicesIdVersionsGetRequest) SortBy(sortBy string) ApiApiDataServicesIdVersionsGetRequest {
	r.sortBy = &sortBy
	return r
}
// Maximum number of rows to return (could be less)
func (r ApiApiDataServicesIdVersionsGetRequest) Limit(limit string) ApiApiDataServicesIdVersionsGetRequest {
	r.limit = &limit
	return r
}
// Use a token returned by a previous query to continue listing with the next batch of rows
func (r ApiApiDataServicesIdVersionsGetRequest) Continuation(continuation string) ApiApiDataServicesIdVersionsGetRequest {
	r.continuation = &continuation
	return r
}
// Filter results by Version id
func (r ApiApiDataServicesIdVersionsGetRequest) Id2(id2 string) ApiApiDataServicesIdVersionsGetRequest {
	r.id2 = &id2
	return r
}
// Filter results by Version&#39;s name
func (r ApiApiDataServicesIdVersionsGetRequest) Name(name string) ApiApiDataServicesIdVersionsGetRequest {
	r.name = &name
	return r
}

func (r ApiApiDataServicesIdVersionsGetRequest) Execute() (*ControllersPaginatedVersions, *http.Response, error) {
	return r.ApiService.ApiDataServicesIdVersionsGetExecute(r)
}

/*
ApiDataServicesIdVersionsGet List Data Service's Versions

Lists Versions belonging to the Data Service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Data Service ID (must be valid UUID)
 @return ApiApiDataServicesIdVersionsGetRequest
*/
func (a *VersionsApiService) ApiDataServicesIdVersionsGet(ctx context.Context, id string) ApiApiDataServicesIdVersionsGetRequest {
	return ApiApiDataServicesIdVersionsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ControllersPaginatedVersions
func (a *VersionsApiService) ApiDataServicesIdVersionsGetExecute(r ApiApiDataServicesIdVersionsGetRequest) (*ControllersPaginatedVersions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ControllersPaginatedVersions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VersionsApiService.ApiDataServicesIdVersionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/data-services/{id}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.continuation != nil {
		localVarQueryParams.Add("continuation", parameterToString(*r.continuation, ""))
	}
	if r.id2 != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id2, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiApiVersionsIdGetRequest struct {
	ctx context.Context
	ApiService *VersionsApiService
	id string
}


func (r ApiApiVersionsIdGetRequest) Execute() (*ModelsVersion, *http.Response, error) {
	return r.ApiService.ApiVersionsIdGetExecute(r)
}

/*
ApiVersionsIdGet Get Version

Fetches a single Version

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Version ID (must be valid UUID)
 @return ApiApiVersionsIdGetRequest
*/
func (a *VersionsApiService) ApiVersionsIdGet(ctx context.Context, id string) ApiApiVersionsIdGetRequest {
	return ApiApiVersionsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsVersion
func (a *VersionsApiService) ApiVersionsIdGetExecute(r ApiApiVersionsIdGetRequest) (*ModelsVersion, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsVersion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VersionsApiService.ApiVersionsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/versions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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