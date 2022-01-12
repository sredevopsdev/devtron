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

// EnvironmentDetail struct for EnvironmentDetail
type EnvironmentDetail struct {
	// name of the environemnt
	EnvironmentName *string `json:"environmentName,omitempty"`
	// id in which app is deployed
	EnvironmentId *int32 `json:"environmentId,omitempty"`
	// namespace corresponding to the environemnt
	Namespace *string `json:"namespace,omitempty"`
	// if given environemnt is marked as production or not, nullable
	IsPrduction *bool `json:"isPrduction,omitempty"`
}

// NewEnvironmentDetail instantiates a new EnvironmentDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentDetail() *EnvironmentDetail {
	this := EnvironmentDetail{}
	return &this
}

// NewEnvironmentDetailWithDefaults instantiates a new EnvironmentDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentDetailWithDefaults() *EnvironmentDetail {
	this := EnvironmentDetail{}
	return &this
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *EnvironmentDetail) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDetail) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *EnvironmentDetail) HasEnvironmentName() bool {
	if o != nil && o.EnvironmentName != nil {
		return true
	}

	return false
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *EnvironmentDetail) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *EnvironmentDetail) GetEnvironmentId() int32 {
	if o == nil || o.EnvironmentId == nil {
		var ret int32
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDetail) GetEnvironmentIdOk() (*int32, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *EnvironmentDetail) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId != nil {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given int32 and assigns it to the EnvironmentId field.
func (o *EnvironmentDetail) SetEnvironmentId(v int32) {
	o.EnvironmentId = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *EnvironmentDetail) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDetail) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *EnvironmentDetail) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *EnvironmentDetail) SetNamespace(v string) {
	o.Namespace = &v
}

// GetIsPrduction returns the IsPrduction field value if set, zero value otherwise.
func (o *EnvironmentDetail) GetIsPrduction() bool {
	if o == nil || o.IsPrduction == nil {
		var ret bool
		return ret
	}
	return *o.IsPrduction
}

// GetIsPrductionOk returns a tuple with the IsPrduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDetail) GetIsPrductionOk() (*bool, bool) {
	if o == nil || o.IsPrduction == nil {
		return nil, false
	}
	return o.IsPrduction, true
}

// HasIsPrduction returns a boolean if a field has been set.
func (o *EnvironmentDetail) HasIsPrduction() bool {
	if o != nil && o.IsPrduction != nil {
		return true
	}

	return false
}

// SetIsPrduction gets a reference to the given bool and assigns it to the IsPrduction field.
func (o *EnvironmentDetail) SetIsPrduction(v bool) {
	o.IsPrduction = &v
}

func (o EnvironmentDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnvironmentName != nil {
		toSerialize["environmentName"] = o.EnvironmentName
	}
	if o.EnvironmentId != nil {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.IsPrduction != nil {
		toSerialize["isPrduction"] = o.IsPrduction
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentDetail struct {
	value *EnvironmentDetail
	isSet bool
}

func (v NullableEnvironmentDetail) Get() *EnvironmentDetail {
	return v.value
}

func (v *NullableEnvironmentDetail) Set(val *EnvironmentDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentDetail(val *EnvironmentDetail) *NullableEnvironmentDetail {
	return &NullableEnvironmentDetail{value: val, isSet: true}
}

func (v NullableEnvironmentDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
