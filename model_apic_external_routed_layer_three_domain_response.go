/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-17769
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// ApicExternalRoutedLayerThreeDomainResponse - The response body of a HTTP GET request for the 'apic.ExternalRoutedLayerThreeDomain' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'apic.ExternalRoutedLayerThreeDomain' resources.
type ApicExternalRoutedLayerThreeDomainResponse struct {
	ApicExternalRoutedLayerThreeDomainList *ApicExternalRoutedLayerThreeDomainList
	MoAggregateTransform                   *MoAggregateTransform
	MoDocumentCount                        *MoDocumentCount
	MoTagSummary                           *MoTagSummary
}

// ApicExternalRoutedLayerThreeDomainListAsApicExternalRoutedLayerThreeDomainResponse is a convenience function that returns ApicExternalRoutedLayerThreeDomainList wrapped in ApicExternalRoutedLayerThreeDomainResponse
func ApicExternalRoutedLayerThreeDomainListAsApicExternalRoutedLayerThreeDomainResponse(v *ApicExternalRoutedLayerThreeDomainList) ApicExternalRoutedLayerThreeDomainResponse {
	return ApicExternalRoutedLayerThreeDomainResponse{
		ApicExternalRoutedLayerThreeDomainList: v,
	}
}

// MoAggregateTransformAsApicExternalRoutedLayerThreeDomainResponse is a convenience function that returns MoAggregateTransform wrapped in ApicExternalRoutedLayerThreeDomainResponse
func MoAggregateTransformAsApicExternalRoutedLayerThreeDomainResponse(v *MoAggregateTransform) ApicExternalRoutedLayerThreeDomainResponse {
	return ApicExternalRoutedLayerThreeDomainResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsApicExternalRoutedLayerThreeDomainResponse is a convenience function that returns MoDocumentCount wrapped in ApicExternalRoutedLayerThreeDomainResponse
func MoDocumentCountAsApicExternalRoutedLayerThreeDomainResponse(v *MoDocumentCount) ApicExternalRoutedLayerThreeDomainResponse {
	return ApicExternalRoutedLayerThreeDomainResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsApicExternalRoutedLayerThreeDomainResponse is a convenience function that returns MoTagSummary wrapped in ApicExternalRoutedLayerThreeDomainResponse
func MoTagSummaryAsApicExternalRoutedLayerThreeDomainResponse(v *MoTagSummary) ApicExternalRoutedLayerThreeDomainResponse {
	return ApicExternalRoutedLayerThreeDomainResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApicExternalRoutedLayerThreeDomainResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'apic.ExternalRoutedLayerThreeDomain.List'
	if jsonDict["ObjectType"] == "apic.ExternalRoutedLayerThreeDomain.List" {
		// try to unmarshal JSON data into ApicExternalRoutedLayerThreeDomainList
		err = json.Unmarshal(data, &dst.ApicExternalRoutedLayerThreeDomainList)
		if err == nil {
			return nil // data stored in dst.ApicExternalRoutedLayerThreeDomainList, return on the first match
		} else {
			dst.ApicExternalRoutedLayerThreeDomainList = nil
			return fmt.Errorf("failed to unmarshal ApicExternalRoutedLayerThreeDomainResponse as ApicExternalRoutedLayerThreeDomainList: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("failed to unmarshal ApicExternalRoutedLayerThreeDomainResponse as MoAggregateTransform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("failed to unmarshal ApicExternalRoutedLayerThreeDomainResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			return nil // data stored in dst.MoTagSummary, return on the first match
		} else {
			dst.MoTagSummary = nil
			return fmt.Errorf("failed to unmarshal ApicExternalRoutedLayerThreeDomainResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApicExternalRoutedLayerThreeDomainResponse) MarshalJSON() ([]byte, error) {
	if src.ApicExternalRoutedLayerThreeDomainList != nil {
		return json.Marshal(&src.ApicExternalRoutedLayerThreeDomainList)
	}

	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApicExternalRoutedLayerThreeDomainResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ApicExternalRoutedLayerThreeDomainList != nil {
		return obj.ApicExternalRoutedLayerThreeDomainList
	}

	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	// all schemas are nil
	return nil
}

type NullableApicExternalRoutedLayerThreeDomainResponse struct {
	value *ApicExternalRoutedLayerThreeDomainResponse
	isSet bool
}

func (v NullableApicExternalRoutedLayerThreeDomainResponse) Get() *ApicExternalRoutedLayerThreeDomainResponse {
	return v.value
}

func (v *NullableApicExternalRoutedLayerThreeDomainResponse) Set(val *ApicExternalRoutedLayerThreeDomainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApicExternalRoutedLayerThreeDomainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApicExternalRoutedLayerThreeDomainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApicExternalRoutedLayerThreeDomainResponse(val *ApicExternalRoutedLayerThreeDomainResponse) *NullableApicExternalRoutedLayerThreeDomainResponse {
	return &NullableApicExternalRoutedLayerThreeDomainResponse{value: val, isSet: true}
}

func (v NullableApicExternalRoutedLayerThreeDomainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApicExternalRoutedLayerThreeDomainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}