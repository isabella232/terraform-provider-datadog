/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// AWSTagFilterListResponseFilters A list of tag filters.
type AWSTagFilterListResponseFilters struct {
	Namespace *AWSNamespace `json:"namespace,omitempty"`
	// The tag filter string.
	TagFilterStr *string `json:"tag_filter_str,omitempty"`
}

// NewAWSTagFilterListResponseFilters instantiates a new AWSTagFilterListResponseFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSTagFilterListResponseFilters() *AWSTagFilterListResponseFilters {
	this := AWSTagFilterListResponseFilters{}
	return &this
}

// NewAWSTagFilterListResponseFiltersWithDefaults instantiates a new AWSTagFilterListResponseFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSTagFilterListResponseFiltersWithDefaults() *AWSTagFilterListResponseFilters {
	this := AWSTagFilterListResponseFilters{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *AWSTagFilterListResponseFilters) GetNamespace() AWSNamespace {
	if o == nil || o.Namespace == nil {
		var ret AWSNamespace
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSTagFilterListResponseFilters) GetNamespaceOk() (*AWSNamespace, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *AWSTagFilterListResponseFilters) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given AWSNamespace and assigns it to the Namespace field.
func (o *AWSTagFilterListResponseFilters) SetNamespace(v AWSNamespace) {
	o.Namespace = &v
}

// GetTagFilterStr returns the TagFilterStr field value if set, zero value otherwise.
func (o *AWSTagFilterListResponseFilters) GetTagFilterStr() string {
	if o == nil || o.TagFilterStr == nil {
		var ret string
		return ret
	}
	return *o.TagFilterStr
}

// GetTagFilterStrOk returns a tuple with the TagFilterStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSTagFilterListResponseFilters) GetTagFilterStrOk() (*string, bool) {
	if o == nil || o.TagFilterStr == nil {
		return nil, false
	}
	return o.TagFilterStr, true
}

// HasTagFilterStr returns a boolean if a field has been set.
func (o *AWSTagFilterListResponseFilters) HasTagFilterStr() bool {
	if o != nil && o.TagFilterStr != nil {
		return true
	}

	return false
}

// SetTagFilterStr gets a reference to the given string and assigns it to the TagFilterStr field.
func (o *AWSTagFilterListResponseFilters) SetTagFilterStr(v string) {
	o.TagFilterStr = &v
}

func (o AWSTagFilterListResponseFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.TagFilterStr != nil {
		toSerialize["tag_filter_str"] = o.TagFilterStr
	}
	return json.Marshal(toSerialize)
}

type NullableAWSTagFilterListResponseFilters struct {
	value *AWSTagFilterListResponseFilters
	isSet bool
}

func (v NullableAWSTagFilterListResponseFilters) Get() *AWSTagFilterListResponseFilters {
	return v.value
}

func (v *NullableAWSTagFilterListResponseFilters) Set(val *AWSTagFilterListResponseFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSTagFilterListResponseFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSTagFilterListResponseFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSTagFilterListResponseFilters(val *AWSTagFilterListResponseFilters) *NullableAWSTagFilterListResponseFilters {
	return &NullableAWSTagFilterListResponseFilters{value: val, isSet: true}
}

func (v NullableAWSTagFilterListResponseFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSTagFilterListResponseFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
