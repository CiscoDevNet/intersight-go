/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2025062323
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// ApicOutResponse - The response body of a HTTP GET request for the 'apic.Out' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'apic.Out' resources.
type ApicOutResponse struct {
	ApicOutList          *ApicOutList
	MoAggregateTransform *MoAggregateTransform
	MoDocumentCount      *MoDocumentCount
	MoTagSummary         *MoTagSummary
}

// ApicOutListAsApicOutResponse is a convenience function that returns ApicOutList wrapped in ApicOutResponse
func ApicOutListAsApicOutResponse(v *ApicOutList) ApicOutResponse {
	return ApicOutResponse{
		ApicOutList: v,
	}
}

// MoAggregateTransformAsApicOutResponse is a convenience function that returns MoAggregateTransform wrapped in ApicOutResponse
func MoAggregateTransformAsApicOutResponse(v *MoAggregateTransform) ApicOutResponse {
	return ApicOutResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsApicOutResponse is a convenience function that returns MoDocumentCount wrapped in ApicOutResponse
func MoDocumentCountAsApicOutResponse(v *MoDocumentCount) ApicOutResponse {
	return ApicOutResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsApicOutResponse is a convenience function that returns MoTagSummary wrapped in ApicOutResponse
func MoTagSummaryAsApicOutResponse(v *MoTagSummary) ApicOutResponse {
	return ApicOutResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApicOutResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'apic.Out.List'
	if jsonDict["ObjectType"] == "apic.Out.List" {
		// try to unmarshal JSON data into ApicOutList
		err = json.Unmarshal(data, &dst.ApicOutList)
		if err == nil {
			return nil // data stored in dst.ApicOutList, return on the first match
		} else {
			dst.ApicOutList = nil
			return fmt.Errorf("failed to unmarshal ApicOutResponse as ApicOutList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ApicOutResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ApicOutResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ApicOutResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApicOutResponse) MarshalJSON() ([]byte, error) {
	if src.ApicOutList != nil {
		return json.Marshal(&src.ApicOutList)
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
func (obj *ApicOutResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ApicOutList != nil {
		return obj.ApicOutList
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

type NullableApicOutResponse struct {
	value *ApicOutResponse
	isSet bool
}

func (v NullableApicOutResponse) Get() *ApicOutResponse {
	return v.value
}

func (v *NullableApicOutResponse) Set(val *ApicOutResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApicOutResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApicOutResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApicOutResponse(val *ApicOutResponse) *NullableApicOutResponse {
	return &NullableApicOutResponse{value: val, isSet: true}
}

func (v NullableApicOutResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApicOutResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
