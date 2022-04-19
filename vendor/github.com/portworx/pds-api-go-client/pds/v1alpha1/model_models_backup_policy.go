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

// ModelsBackupPolicy struct for ModelsBackupPolicy
type ModelsBackupPolicy struct {
	AccountId *string `json:"account_id,omitempty"`
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	// Name of the backup policy. Must be unique for the given tenant.
	Name *string `json:"name,omitempty"`
	// An array of the backup schedules. Must be non-empty.
	Schedules []ModelsBackupSchedule `json:"schedules,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewModelsBackupPolicy instantiates a new ModelsBackupPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsBackupPolicy() *ModelsBackupPolicy {
	this := ModelsBackupPolicy{}
	return &this
}

// NewModelsBackupPolicyWithDefaults instantiates a new ModelsBackupPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsBackupPolicyWithDefaults() *ModelsBackupPolicy {
	this := ModelsBackupPolicy{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ModelsBackupPolicy) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupPolicy) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ModelsBackupPolicy) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ModelsBackupPolicy) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsBackupPolicy) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupPolicy) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsBackupPolicy) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsBackupPolicy) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsBackupPolicy) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupPolicy) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsBackupPolicy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsBackupPolicy) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsBackupPolicy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupPolicy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsBackupPolicy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsBackupPolicy) SetName(v string) {
	o.Name = &v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *ModelsBackupPolicy) GetSchedules() []ModelsBackupSchedule {
	if o == nil || o.Schedules == nil {
		var ret []ModelsBackupSchedule
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupPolicy) GetSchedulesOk() ([]ModelsBackupSchedule, bool) {
	if o == nil || o.Schedules == nil {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *ModelsBackupPolicy) HasSchedules() bool {
	if o != nil && o.Schedules != nil {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []ModelsBackupSchedule and assigns it to the Schedules field.
func (o *ModelsBackupPolicy) SetSchedules(v []ModelsBackupSchedule) {
	o.Schedules = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ModelsBackupPolicy) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupPolicy) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ModelsBackupPolicy) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ModelsBackupPolicy) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsBackupPolicy) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupPolicy) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsBackupPolicy) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsBackupPolicy) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsBackupPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Schedules != nil {
		toSerialize["schedules"] = o.Schedules
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableModelsBackupPolicy struct {
	value *ModelsBackupPolicy
	isSet bool
}

func (v NullableModelsBackupPolicy) Get() *ModelsBackupPolicy {
	return v.value
}

func (v *NullableModelsBackupPolicy) Set(val *ModelsBackupPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsBackupPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsBackupPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsBackupPolicy(val *ModelsBackupPolicy) *NullableModelsBackupPolicy {
	return &NullableModelsBackupPolicy{value: val, isSet: true}
}

func (v NullableModelsBackupPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsBackupPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

