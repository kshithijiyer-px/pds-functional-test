/*
PDS API

Portworx Data Services API Server

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// ModelsProjectRoleBinding struct for ModelsProjectRoleBinding
type ModelsProjectRoleBinding struct {
	ActorId *string `json:"actor_id,omitempty"`
	ActorType *string `json:"actor_type,omitempty"`
	ProjectId *string `json:"project_id,omitempty"`
	RoleName *string `json:"role_name,omitempty"`
}

// NewModelsProjectRoleBinding instantiates a new ModelsProjectRoleBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsProjectRoleBinding() *ModelsProjectRoleBinding {
	this := ModelsProjectRoleBinding{}
	return &this
}

// NewModelsProjectRoleBindingWithDefaults instantiates a new ModelsProjectRoleBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsProjectRoleBindingWithDefaults() *ModelsProjectRoleBinding {
	this := ModelsProjectRoleBinding{}
	return &this
}

// GetActorId returns the ActorId field value if set, zero value otherwise.
func (o *ModelsProjectRoleBinding) GetActorId() string {
	if o == nil || o.ActorId == nil {
		var ret string
		return ret
	}
	return *o.ActorId
}

// GetActorIdOk returns a tuple with the ActorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsProjectRoleBinding) GetActorIdOk() (*string, bool) {
	if o == nil || o.ActorId == nil {
		return nil, false
	}
	return o.ActorId, true
}

// HasActorId returns a boolean if a field has been set.
func (o *ModelsProjectRoleBinding) HasActorId() bool {
	if o != nil && o.ActorId != nil {
		return true
	}

	return false
}

// SetActorId gets a reference to the given string and assigns it to the ActorId field.
func (o *ModelsProjectRoleBinding) SetActorId(v string) {
	o.ActorId = &v
}

// GetActorType returns the ActorType field value if set, zero value otherwise.
func (o *ModelsProjectRoleBinding) GetActorType() string {
	if o == nil || o.ActorType == nil {
		var ret string
		return ret
	}
	return *o.ActorType
}

// GetActorTypeOk returns a tuple with the ActorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsProjectRoleBinding) GetActorTypeOk() (*string, bool) {
	if o == nil || o.ActorType == nil {
		return nil, false
	}
	return o.ActorType, true
}

// HasActorType returns a boolean if a field has been set.
func (o *ModelsProjectRoleBinding) HasActorType() bool {
	if o != nil && o.ActorType != nil {
		return true
	}

	return false
}

// SetActorType gets a reference to the given string and assigns it to the ActorType field.
func (o *ModelsProjectRoleBinding) SetActorType(v string) {
	o.ActorType = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ModelsProjectRoleBinding) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsProjectRoleBinding) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ModelsProjectRoleBinding) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ModelsProjectRoleBinding) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *ModelsProjectRoleBinding) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsProjectRoleBinding) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *ModelsProjectRoleBinding) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *ModelsProjectRoleBinding) SetRoleName(v string) {
	o.RoleName = &v
}

func (o ModelsProjectRoleBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActorId != nil {
		toSerialize["actor_id"] = o.ActorId
	}
	if o.ActorType != nil {
		toSerialize["actor_type"] = o.ActorType
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableModelsProjectRoleBinding struct {
	value *ModelsProjectRoleBinding
	isSet bool
}

func (v NullableModelsProjectRoleBinding) Get() *ModelsProjectRoleBinding {
	return v.value
}

func (v *NullableModelsProjectRoleBinding) Set(val *ModelsProjectRoleBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsProjectRoleBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsProjectRoleBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsProjectRoleBinding(val *ModelsProjectRoleBinding) *NullableModelsProjectRoleBinding {
	return &NullableModelsProjectRoleBinding{value: val, isSet: true}
}

func (v NullableModelsProjectRoleBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsProjectRoleBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

