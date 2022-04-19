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

// AccountRoleBindingsApiService AccountRoleBindingsApi service
type AccountRoleBindingsApiService service

type ApiApiAccountsIdInvitationsPostRequest struct {
	ctx context.Context
	ApiService *AccountRoleBindingsApiService
	id string
	body *ControllersInvitationRequest
}

// Request body containing the invitation details.
func (r ApiApiAccountsIdInvitationsPostRequest) Body(body ControllersInvitationRequest) ApiApiAccountsIdInvitationsPostRequest {
	r.body = &body
	return r
}

func (r ApiApiAccountsIdInvitationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiAccountsIdInvitationsPostExecute(r)
}

/*
ApiAccountsIdInvitationsPost Create Invitation

Adds role binding to existing user. The plan is to send invites to non-existing users in the future.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Account ID (must be valid UUID)
 @return ApiApiAccountsIdInvitationsPostRequest
*/
func (a *AccountRoleBindingsApiService) ApiAccountsIdInvitationsPost(ctx context.Context, id string) ApiApiAccountsIdInvitationsPostRequest {
	return ApiApiAccountsIdInvitationsPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *AccountRoleBindingsApiService) ApiAccountsIdInvitationsPostExecute(r ApiApiAccountsIdInvitationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountRoleBindingsApiService.ApiAccountsIdInvitationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/accounts/{id}/invitations"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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

type ApiApiAccountsIdRoleBindingsDeleteRequest struct {
	ctx context.Context
	ApiService *AccountRoleBindingsApiService
	id string
	actorType *string
}

// AccountRoleBinding actor type
func (r ApiApiAccountsIdRoleBindingsDeleteRequest) ActorType(actorType string) ApiApiAccountsIdRoleBindingsDeleteRequest {
	r.actorType = &actorType
	return r
}

func (r ApiApiAccountsIdRoleBindingsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiAccountsIdRoleBindingsDeleteExecute(r)
}

/*
ApiAccountsIdRoleBindingsDelete Delete AccountRoleBinding

Removes a single AccountRoleBinding

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Account ID (must be valid UUID)
 @return ApiApiAccountsIdRoleBindingsDeleteRequest
*/
func (a *AccountRoleBindingsApiService) ApiAccountsIdRoleBindingsDelete(ctx context.Context, id string) ApiApiAccountsIdRoleBindingsDeleteRequest {
	return ApiApiAccountsIdRoleBindingsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *AccountRoleBindingsApiService) ApiAccountsIdRoleBindingsDeleteExecute(r ApiApiAccountsIdRoleBindingsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountRoleBindingsApiService.ApiAccountsIdRoleBindingsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/accounts/{id}/role-bindings"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.actorType
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

type ApiApiAccountsIdRoleBindingsGetRequest struct {
	ctx context.Context
	ApiService *AccountRoleBindingsApiService
	id string
	sortBy *string
	roleName *string
	actorId *string
	actorType *string
}

// A given AccountRoleBinding attribute to sort results by (one of: role_name, actor_id)
func (r ApiApiAccountsIdRoleBindingsGetRequest) SortBy(sortBy string) ApiApiAccountsIdRoleBindingsGetRequest {
	r.sortBy = &sortBy
	return r
}
// Filter results by AccountRoleBinding assigned role name
func (r ApiApiAccountsIdRoleBindingsGetRequest) RoleName(roleName string) ApiApiAccountsIdRoleBindingsGetRequest {
	r.roleName = &roleName
	return r
}
// Filter results by AccountRoleBinding actor id
func (r ApiApiAccountsIdRoleBindingsGetRequest) ActorId(actorId string) ApiApiAccountsIdRoleBindingsGetRequest {
	r.actorId = &actorId
	return r
}
// Filter results by AccountRoleBinding actor type
func (r ApiApiAccountsIdRoleBindingsGetRequest) ActorType(actorType string) ApiApiAccountsIdRoleBindingsGetRequest {
	r.actorType = &actorType
	return r
}

func (r ApiApiAccountsIdRoleBindingsGetRequest) Execute() (*ControllersPaginatedAccountRoleBindings, *http.Response, error) {
	return r.ApiService.ApiAccountsIdRoleBindingsGetExecute(r)
}

/*
ApiAccountsIdRoleBindingsGet List AccountRoleBinding

Lists AccountRoleBinding

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Account ID (must be valid UUID)
 @return ApiApiAccountsIdRoleBindingsGetRequest
*/
func (a *AccountRoleBindingsApiService) ApiAccountsIdRoleBindingsGet(ctx context.Context, id string) ApiApiAccountsIdRoleBindingsGetRequest {
	return ApiApiAccountsIdRoleBindingsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ControllersPaginatedAccountRoleBindings
func (a *AccountRoleBindingsApiService) ApiAccountsIdRoleBindingsGetExecute(r ApiApiAccountsIdRoleBindingsGetRequest) (*ControllersPaginatedAccountRoleBindings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ControllersPaginatedAccountRoleBindings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountRoleBindingsApiService.ApiAccountsIdRoleBindingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/accounts/{id}/role-bindings"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.roleName != nil {
		localVarQueryParams.Add("role_name", parameterToString(*r.roleName, ""))
	}
	if r.actorId != nil {
		localVarQueryParams.Add("actor_id", parameterToString(*r.actorId, ""))
	}
	if r.actorType != nil {
		localVarQueryParams.Add("actor_type", parameterToString(*r.actorType, ""))
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

type ApiApiAccountsIdRoleBindingsPutRequest struct {
	ctx context.Context
	ApiService *AccountRoleBindingsApiService
	id string
	body *ControllersUpsertAccountRoleBindingRequest
}

// Request body containing the account role binding
func (r ApiApiAccountsIdRoleBindingsPutRequest) Body(body ControllersUpsertAccountRoleBindingRequest) ApiApiAccountsIdRoleBindingsPutRequest {
	r.body = &body
	return r
}

func (r ApiApiAccountsIdRoleBindingsPutRequest) Execute() (*ModelsAccountRoleBinding, *http.Response, error) {
	return r.ApiService.ApiAccountsIdRoleBindingsPutExecute(r)
}

/*
ApiAccountsIdRoleBindingsPut Create/Update AccountRoleBinding

Creates a new AccountRoleBinding, or updates role_name if binding for (actor_id, actor_type) already exists.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Account ID (must be valid UUID)
 @return ApiApiAccountsIdRoleBindingsPutRequest
*/
func (a *AccountRoleBindingsApiService) ApiAccountsIdRoleBindingsPut(ctx context.Context, id string) ApiApiAccountsIdRoleBindingsPutRequest {
	return ApiApiAccountsIdRoleBindingsPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsAccountRoleBinding
func (a *AccountRoleBindingsApiService) ApiAccountsIdRoleBindingsPutExecute(r ApiApiAccountsIdRoleBindingsPutRequest) (*ModelsAccountRoleBinding, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsAccountRoleBinding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountRoleBindingsApiService.ApiAccountsIdRoleBindingsPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/accounts/{id}/role-bindings"
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

type ApiApiUsersIdAccountRoleBindingsGetRequest struct {
	ctx context.Context
	ApiService *AccountRoleBindingsApiService
	id string
	sortBy *string
	roleName *string
}

// A given AccountRoleBinding attribute to sort results by (allowed: role_name)
func (r ApiApiUsersIdAccountRoleBindingsGetRequest) SortBy(sortBy string) ApiApiUsersIdAccountRoleBindingsGetRequest {
	r.sortBy = &sortBy
	return r
}
// Filter results by AccountRoleBinding assigned role_name
func (r ApiApiUsersIdAccountRoleBindingsGetRequest) RoleName(roleName string) ApiApiUsersIdAccountRoleBindingsGetRequest {
	r.roleName = &roleName
	return r
}

func (r ApiApiUsersIdAccountRoleBindingsGetRequest) Execute() (*ControllersPaginatedAccountRoleBindings, *http.Response, error) {
	return r.ApiService.ApiUsersIdAccountRoleBindingsGetExecute(r)
}

/*
ApiUsersIdAccountRoleBindingsGet List AccountRoleBindings of a given user

Every user can read its own bindings. Only pds-admin can read bindings of other users.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id User ID (must be valid UUID)
 @return ApiApiUsersIdAccountRoleBindingsGetRequest
*/
func (a *AccountRoleBindingsApiService) ApiUsersIdAccountRoleBindingsGet(ctx context.Context, id string) ApiApiUsersIdAccountRoleBindingsGetRequest {
	return ApiApiUsersIdAccountRoleBindingsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ControllersPaginatedAccountRoleBindings
func (a *AccountRoleBindingsApiService) ApiUsersIdAccountRoleBindingsGetExecute(r ApiApiUsersIdAccountRoleBindingsGetRequest) (*ControllersPaginatedAccountRoleBindings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ControllersPaginatedAccountRoleBindings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountRoleBindingsApiService.ApiUsersIdAccountRoleBindingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/users/{id}/account-role-bindings"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.roleName != nil {
		localVarQueryParams.Add("role_name", parameterToString(*r.roleName, ""))
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
