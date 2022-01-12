/*
Devtron Labs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// HibernateRequest struct for HibernateRequest
type HibernateRequest struct {
	// helm app id
	AppId     *string                  `json:"appId,omitempty"`
	Resources *[]HibernateTargetObject `json:"resources,omitempty"`
}

// NewHibernateRequest instantiates a new HibernateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHibernateRequest() *HibernateRequest {
	this := HibernateRequest{}
	return &this
}

// NewHibernateRequestWithDefaults instantiates a new HibernateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHibernateRequestWithDefaults() *HibernateRequest {
	this := HibernateRequest{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *HibernateRequest) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HibernateRequest) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *HibernateRequest) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *HibernateRequest) SetAppId(v string) {
	o.AppId = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *HibernateRequest) GetResources() []HibernateTargetObject {
	if o == nil || o.Resources == nil {
		var ret []HibernateTargetObject
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HibernateRequest) GetResourcesOk() (*[]HibernateTargetObject, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *HibernateRequest) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []HibernateTargetObject and assigns it to the Resources field.
func (o *HibernateRequest) SetResources(v []HibernateTargetObject) {
	o.Resources = &v
}

func (o HibernateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableHibernateRequest struct {
	value *HibernateRequest
	isSet bool
}

func (v NullableHibernateRequest) Get() *HibernateRequest {
	return v.value
}

func (v *NullableHibernateRequest) Set(val *HibernateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHibernateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHibernateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHibernateRequest(val *HibernateRequest) *NullableHibernateRequest {
	return &NullableHibernateRequest{value: val, isSet: true}
}

func (v NullableHibernateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHibernateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
