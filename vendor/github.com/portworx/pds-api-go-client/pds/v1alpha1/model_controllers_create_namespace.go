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

// ControllersCreateNamespace struct for ControllersCreateNamespace
type ControllersCreateNamespace struct {
	// Name of the namespace
	Name *string `json:"name,omitempty"`
}

// NewControllersCreateNamespace instantiates a new ControllersCreateNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersCreateNamespace() *ControllersCreateNamespace {
	this := ControllersCreateNamespace{}
	return &this
}

// NewControllersCreateNamespaceWithDefaults instantiates a new ControllersCreateNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersCreateNamespaceWithDefaults() *ControllersCreateNamespace {
	this := ControllersCreateNamespace{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ControllersCreateNamespace) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateNamespace) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ControllersCreateNamespace) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ControllersCreateNamespace) SetName(v string) {
	o.Name = &v
}

func (o ControllersCreateNamespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableControllersCreateNamespace struct {
	value *ControllersCreateNamespace
	isSet bool
}

func (v NullableControllersCreateNamespace) Get() *ControllersCreateNamespace {
	return v.value
}

func (v *NullableControllersCreateNamespace) Set(val *ControllersCreateNamespace) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersCreateNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersCreateNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersCreateNamespace(val *ControllersCreateNamespace) *NullableControllersCreateNamespace {
	return &NullableControllersCreateNamespace{value: val, isSet: true}
}

func (v NullableControllersCreateNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersCreateNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


