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

// BackupTargetsApiService BackupTargetsApi service
type BackupTargetsApiService service

type ApiApiBackupTargetsIdDeleteRequest struct {
	ctx context.Context
	ApiService *BackupTargetsApiService
	id string
	force *string
}

// Delete backup target even if the deletion job fails on any deployment targets
func (r ApiApiBackupTargetsIdDeleteRequest) Force(force string) ApiApiBackupTargetsIdDeleteRequest {
	r.force = &force
	return r
}

func (r ApiApiBackupTargetsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiBackupTargetsIdDeleteExecute(r)
}

/*
ApiBackupTargetsIdDelete Delete BackupTargets

Removes a single BackupTarget

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id BackupTarget ID (must be valid UUID)
 @return ApiApiBackupTargetsIdDeleteRequest
*/
func (a *BackupTargetsApiService) ApiBackupTargetsIdDelete(ctx context.Context, id string) ApiApiBackupTargetsIdDeleteRequest {
	return ApiApiBackupTargetsIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *BackupTargetsApiService) ApiBackupTargetsIdDeleteExecute(r ApiApiBackupTargetsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupTargetsApiService.ApiBackupTargetsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/backup-targets/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.force != nil {
		localVarQueryParams.Add("force", parameterToString(*r.force, ""))
	}
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiApiBackupTargetsIdGetRequest struct {
	ctx context.Context
	ApiService *BackupTargetsApiService
	id string
}


func (r ApiApiBackupTargetsIdGetRequest) Execute() (*ModelsBackupTarget, *http.Response, error) {
	return r.ApiService.ApiBackupTargetsIdGetExecute(r)
}

/*
ApiBackupTargetsIdGet Get BackupTarget

Fetches a single BackupTarget

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id BackupTarget ID (must be valid UUID)
 @return ApiApiBackupTargetsIdGetRequest
*/
func (a *BackupTargetsApiService) ApiBackupTargetsIdGet(ctx context.Context, id string) ApiApiBackupTargetsIdGetRequest {
	return ApiApiBackupTargetsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsBackupTarget
func (a *BackupTargetsApiService) ApiBackupTargetsIdGetExecute(r ApiApiBackupTargetsIdGetRequest) (*ModelsBackupTarget, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsBackupTarget
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupTargetsApiService.ApiBackupTargetsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/backup-targets/{id}"
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

type ApiApiBackupTargetsIdPutRequest struct {
	ctx context.Context
	ApiService *BackupTargetsApiService
	id string
	body *ControllersUpdateBackupTargetRequest
}

// Object with partial update fields
func (r ApiApiBackupTargetsIdPutRequest) Body(body ControllersUpdateBackupTargetRequest) ApiApiBackupTargetsIdPutRequest {
	r.body = &body
	return r
}

func (r ApiApiBackupTargetsIdPutRequest) Execute() (*ModelsBackupTarget, *http.Response, error) {
	return r.ApiService.ApiBackupTargetsIdPutExecute(r)
}

/*
ApiBackupTargetsIdPut Update BackupTarget

Updates existing BackupTarget

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id BackupTarget ID (must be valid UUID)
 @return ApiApiBackupTargetsIdPutRequest
*/
func (a *BackupTargetsApiService) ApiBackupTargetsIdPut(ctx context.Context, id string) ApiApiBackupTargetsIdPutRequest {
	return ApiApiBackupTargetsIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsBackupTarget
func (a *BackupTargetsApiService) ApiBackupTargetsIdPutExecute(r ApiApiBackupTargetsIdPutRequest) (*ModelsBackupTarget, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsBackupTarget
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupTargetsApiService.ApiBackupTargetsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/backup-targets/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
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

type ApiApiBackupTargetsIdRetryPostRequest struct {
	ctx context.Context
	ApiService *BackupTargetsApiService
	id string
}


func (r ApiApiBackupTargetsIdRetryPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiBackupTargetsIdRetryPostExecute(r)
}

/*
ApiBackupTargetsIdRetryPost Retry sync of a BackupTarget

Retries to sync failed BackupTargetState

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id BackupTargetState ID (must be valid UUID)
 @return ApiApiBackupTargetsIdRetryPostRequest
*/
func (a *BackupTargetsApiService) ApiBackupTargetsIdRetryPost(ctx context.Context, id string) ApiApiBackupTargetsIdRetryPostRequest {
	return ApiApiBackupTargetsIdRetryPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *BackupTargetsApiService) ApiBackupTargetsIdRetryPostExecute(r ApiApiBackupTargetsIdRetryPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupTargetsApiService.ApiBackupTargetsIdRetryPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/backup-targets/{id}/retry"
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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiApiBackupTargetsIdStatesGetRequest struct {
	ctx context.Context
	ApiService *BackupTargetsApiService
	id string
	limit *string
	continuation *string
	sortBy *string
	backupTargetId *string
	deploymentTargetId *string
	state *string
}

// Maximum number of rows to return (could be less)
func (r ApiApiBackupTargetsIdStatesGetRequest) Limit(limit string) ApiApiBackupTargetsIdStatesGetRequest {
	r.limit = &limit
	return r
}
// Use a token returned by a previous query to continue listing with the next batch of rows
func (r ApiApiBackupTargetsIdStatesGetRequest) Continuation(continuation string) ApiApiBackupTargetsIdStatesGetRequest {
	r.continuation = &continuation
	return r
}
// A given BackupTargetState attribute to sort results by (one of: state, deployment_target_id, backup_target_id)
func (r ApiApiBackupTargetsIdStatesGetRequest) SortBy(sortBy string) ApiApiBackupTargetsIdStatesGetRequest {
	r.sortBy = &sortBy
	return r
}
// Filter results by BackupTarget ID
func (r ApiApiBackupTargetsIdStatesGetRequest) BackupTargetId(backupTargetId string) ApiApiBackupTargetsIdStatesGetRequest {
	r.backupTargetId = &backupTargetId
	return r
}
// Filter results by DeploymentTarget ID
func (r ApiApiBackupTargetsIdStatesGetRequest) DeploymentTargetId(deploymentTargetId string) ApiApiBackupTargetsIdStatesGetRequest {
	r.deploymentTargetId = &deploymentTargetId
	return r
}
// Filter results by state
func (r ApiApiBackupTargetsIdStatesGetRequest) State(state string) ApiApiBackupTargetsIdStatesGetRequest {
	r.state = &state
	return r
}

func (r ApiApiBackupTargetsIdStatesGetRequest) Execute() (*ControllersPaginatedBackupTargetStates, *http.Response, error) {
	return r.ApiService.ApiBackupTargetsIdStatesGetExecute(r)
}

/*
ApiBackupTargetsIdStatesGet List BackupTarget's BackupTargetStates

Lists BackupTarget's BackupTargetStates

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Backup Target ID (must be valid UUID)
 @return ApiApiBackupTargetsIdStatesGetRequest
*/
func (a *BackupTargetsApiService) ApiBackupTargetsIdStatesGet(ctx context.Context, id string) ApiApiBackupTargetsIdStatesGetRequest {
	return ApiApiBackupTargetsIdStatesGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ControllersPaginatedBackupTargetStates
func (a *BackupTargetsApiService) ApiBackupTargetsIdStatesGetExecute(r ApiApiBackupTargetsIdStatesGetRequest) (*ControllersPaginatedBackupTargetStates, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ControllersPaginatedBackupTargetStates
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupTargetsApiService.ApiBackupTargetsIdStatesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/backup-targets/{id}/states"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.continuation != nil {
		localVarQueryParams.Add("continuation", parameterToString(*r.continuation, ""))
	}
	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.backupTargetId != nil {
		localVarQueryParams.Add("backup_target_id", parameterToString(*r.backupTargetId, ""))
	}
	if r.deploymentTargetId != nil {
		localVarQueryParams.Add("deployment_target_id", parameterToString(*r.deploymentTargetId, ""))
	}
	if r.state != nil {
		localVarQueryParams.Add("state", parameterToString(*r.state, ""))
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

type ApiApiTenantsIdBackupTargetsGetRequest struct {
	ctx context.Context
	ApiService *BackupTargetsApiService
	id string
	limit *string
	continuation *string
	sortBy *string
	id2 *string
	name *string
	type_ *string
	bucket *string
	region *string
	backupCredentialsId *string
}

// Maximum number of rows to return (could be less)
func (r ApiApiTenantsIdBackupTargetsGetRequest) Limit(limit string) ApiApiTenantsIdBackupTargetsGetRequest {
	r.limit = &limit
	return r
}
// Use a token returned by a previous query to continue listing with the next batch of rows
func (r ApiApiTenantsIdBackupTargetsGetRequest) Continuation(continuation string) ApiApiTenantsIdBackupTargetsGetRequest {
	r.continuation = &continuation
	return r
}
// A given BackupTarget attribute to sort results by (one of: id, name, type, created_at)
func (r ApiApiTenantsIdBackupTargetsGetRequest) SortBy(sortBy string) ApiApiTenantsIdBackupTargetsGetRequest {
	r.sortBy = &sortBy
	return r
}
// Filter results by BackupTarget ID
func (r ApiApiTenantsIdBackupTargetsGetRequest) Id2(id2 string) ApiApiTenantsIdBackupTargetsGetRequest {
	r.id2 = &id2
	return r
}
// Filter results by BackupTarget name
func (r ApiApiTenantsIdBackupTargetsGetRequest) Name(name string) ApiApiTenantsIdBackupTargetsGetRequest {
	r.name = &name
	return r
}
// Filter results by BackupTarget type
func (r ApiApiTenantsIdBackupTargetsGetRequest) Type_(type_ string) ApiApiTenantsIdBackupTargetsGetRequest {
	r.type_ = &type_
	return r
}
// Filter results by bucket
func (r ApiApiTenantsIdBackupTargetsGetRequest) Bucket(bucket string) ApiApiTenantsIdBackupTargetsGetRequest {
	r.bucket = &bucket
	return r
}
// Filter results by region
func (r ApiApiTenantsIdBackupTargetsGetRequest) Region(region string) ApiApiTenantsIdBackupTargetsGetRequest {
	r.region = &region
	return r
}
// Filter results by BackupCredentials ID
func (r ApiApiTenantsIdBackupTargetsGetRequest) BackupCredentialsId(backupCredentialsId string) ApiApiTenantsIdBackupTargetsGetRequest {
	r.backupCredentialsId = &backupCredentialsId
	return r
}

func (r ApiApiTenantsIdBackupTargetsGetRequest) Execute() (*ControllersPaginatedBackupTargets, *http.Response, error) {
	return r.ApiService.ApiTenantsIdBackupTargetsGetExecute(r)
}

/*
ApiTenantsIdBackupTargetsGet List Tenant's BackupTargets

Lists Tenant's BackupTargets

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Tenant ID (must be valid UUID)
 @return ApiApiTenantsIdBackupTargetsGetRequest
*/
func (a *BackupTargetsApiService) ApiTenantsIdBackupTargetsGet(ctx context.Context, id string) ApiApiTenantsIdBackupTargetsGetRequest {
	return ApiApiTenantsIdBackupTargetsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ControllersPaginatedBackupTargets
func (a *BackupTargetsApiService) ApiTenantsIdBackupTargetsGetExecute(r ApiApiTenantsIdBackupTargetsGetRequest) (*ControllersPaginatedBackupTargets, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ControllersPaginatedBackupTargets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupTargetsApiService.ApiTenantsIdBackupTargetsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenants/{id}/backup-targets"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.continuation != nil {
		localVarQueryParams.Add("continuation", parameterToString(*r.continuation, ""))
	}
	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.id2 != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id2, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	}
	if r.bucket != nil {
		localVarQueryParams.Add("bucket", parameterToString(*r.bucket, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.backupCredentialsId != nil {
		localVarQueryParams.Add("backup_credentials_id", parameterToString(*r.backupCredentialsId, ""))
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

type ApiApiTenantsIdBackupTargetsPostRequest struct {
	ctx context.Context
	ApiService *BackupTargetsApiService
	id string
	body *ControllersCreateTenantBackupTarget
}

// Request body containing the backup target config
func (r ApiApiTenantsIdBackupTargetsPostRequest) Body(body ControllersCreateTenantBackupTarget) ApiApiTenantsIdBackupTargetsPostRequest {
	r.body = &body
	return r
}

func (r ApiApiTenantsIdBackupTargetsPostRequest) Execute() (*ModelsBackupTarget, *http.Response, error) {
	return r.ApiService.ApiTenantsIdBackupTargetsPostExecute(r)
}

/*
ApiTenantsIdBackupTargetsPost Create BackupTarget

Creates a new BackupTarget

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Tenant ID (must be valid UUID)
 @return ApiApiTenantsIdBackupTargetsPostRequest
*/
func (a *BackupTargetsApiService) ApiTenantsIdBackupTargetsPost(ctx context.Context, id string) ApiApiTenantsIdBackupTargetsPostRequest {
	return ApiApiTenantsIdBackupTargetsPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsBackupTarget
func (a *BackupTargetsApiService) ApiTenantsIdBackupTargetsPostExecute(r ApiApiTenantsIdBackupTargetsPostRequest) (*ModelsBackupTarget, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsBackupTarget
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupTargetsApiService.ApiTenantsIdBackupTargetsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenants/{id}/backup-targets"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
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