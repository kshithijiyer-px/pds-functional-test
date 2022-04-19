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

// ModelsBackupSchedule struct for ModelsBackupSchedule
type ModelsBackupSchedule struct {
	// An ID of the backup schedule to help the UI identify validation failures. The type is UUID and it is generated on the UI side.
	Id *string `json:"id,omitempty"`
	// Number of retained backup jobs. Must be 1 or greater.
	RetentionCount *int32 `json:"retention_count,omitempty"`
	// CRON expression for the backup schedule.
	Schedule *string `json:"schedule,omitempty"`
	// Type of the backup schedule.
	Type *string `json:"type,omitempty"`
}

// NewModelsBackupSchedule instantiates a new ModelsBackupSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsBackupSchedule() *ModelsBackupSchedule {
	this := ModelsBackupSchedule{}
	return &this
}

// NewModelsBackupScheduleWithDefaults instantiates a new ModelsBackupSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsBackupScheduleWithDefaults() *ModelsBackupSchedule {
	this := ModelsBackupSchedule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsBackupSchedule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupSchedule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsBackupSchedule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsBackupSchedule) SetId(v string) {
	o.Id = &v
}

// GetRetentionCount returns the RetentionCount field value if set, zero value otherwise.
func (o *ModelsBackupSchedule) GetRetentionCount() int32 {
	if o == nil || o.RetentionCount == nil {
		var ret int32
		return ret
	}
	return *o.RetentionCount
}

// GetRetentionCountOk returns a tuple with the RetentionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupSchedule) GetRetentionCountOk() (*int32, bool) {
	if o == nil || o.RetentionCount == nil {
		return nil, false
	}
	return o.RetentionCount, true
}

// HasRetentionCount returns a boolean if a field has been set.
func (o *ModelsBackupSchedule) HasRetentionCount() bool {
	if o != nil && o.RetentionCount != nil {
		return true
	}

	return false
}

// SetRetentionCount gets a reference to the given int32 and assigns it to the RetentionCount field.
func (o *ModelsBackupSchedule) SetRetentionCount(v int32) {
	o.RetentionCount = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ModelsBackupSchedule) GetSchedule() string {
	if o == nil || o.Schedule == nil {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupSchedule) GetScheduleOk() (*string, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ModelsBackupSchedule) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *ModelsBackupSchedule) SetSchedule(v string) {
	o.Schedule = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelsBackupSchedule) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupSchedule) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelsBackupSchedule) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ModelsBackupSchedule) SetType(v string) {
	o.Type = &v
}

func (o ModelsBackupSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RetentionCount != nil {
		toSerialize["retention_count"] = o.RetentionCount
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableModelsBackupSchedule struct {
	value *ModelsBackupSchedule
	isSet bool
}

func (v NullableModelsBackupSchedule) Get() *ModelsBackupSchedule {
	return v.value
}

func (v *NullableModelsBackupSchedule) Set(val *ModelsBackupSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsBackupSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsBackupSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsBackupSchedule(val *ModelsBackupSchedule) *NullableModelsBackupSchedule {
	return &NullableModelsBackupSchedule{value: val, isSet: true}
}

func (v NullableModelsBackupSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsBackupSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

