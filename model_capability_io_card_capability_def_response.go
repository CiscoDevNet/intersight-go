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

// CapabilityIoCardCapabilityDefResponse - The response body of a HTTP GET request for the 'capability.IoCardCapabilityDef' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'capability.IoCardCapabilityDef' resources.
type CapabilityIoCardCapabilityDefResponse struct {
	CapabilityIoCardCapabilityDefList *CapabilityIoCardCapabilityDefList
	MoAggregateTransform              *MoAggregateTransform
	MoDocumentCount                   *MoDocumentCount
	MoTagSummary                      *MoTagSummary
}

// CapabilityIoCardCapabilityDefListAsCapabilityIoCardCapabilityDefResponse is a convenience function that returns CapabilityIoCardCapabilityDefList wrapped in CapabilityIoCardCapabilityDefResponse
func CapabilityIoCardCapabilityDefListAsCapabilityIoCardCapabilityDefResponse(v *CapabilityIoCardCapabilityDefList) CapabilityIoCardCapabilityDefResponse {
	return CapabilityIoCardCapabilityDefResponse{
		CapabilityIoCardCapabilityDefList: v,
	}
}

// MoAggregateTransformAsCapabilityIoCardCapabilityDefResponse is a convenience function that returns MoAggregateTransform wrapped in CapabilityIoCardCapabilityDefResponse
func MoAggregateTransformAsCapabilityIoCardCapabilityDefResponse(v *MoAggregateTransform) CapabilityIoCardCapabilityDefResponse {
	return CapabilityIoCardCapabilityDefResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsCapabilityIoCardCapabilityDefResponse is a convenience function that returns MoDocumentCount wrapped in CapabilityIoCardCapabilityDefResponse
func MoDocumentCountAsCapabilityIoCardCapabilityDefResponse(v *MoDocumentCount) CapabilityIoCardCapabilityDefResponse {
	return CapabilityIoCardCapabilityDefResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsCapabilityIoCardCapabilityDefResponse is a convenience function that returns MoTagSummary wrapped in CapabilityIoCardCapabilityDefResponse
func MoTagSummaryAsCapabilityIoCardCapabilityDefResponse(v *MoTagSummary) CapabilityIoCardCapabilityDefResponse {
	return CapabilityIoCardCapabilityDefResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CapabilityIoCardCapabilityDefResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'capability.IoCardCapabilityDef.List'
	if jsonDict["ObjectType"] == "capability.IoCardCapabilityDef.List" {
		// try to unmarshal JSON data into CapabilityIoCardCapabilityDefList
		err = json.Unmarshal(data, &dst.CapabilityIoCardCapabilityDefList)
		if err == nil {
			return nil // data stored in dst.CapabilityIoCardCapabilityDefList, return on the first match
		} else {
			dst.CapabilityIoCardCapabilityDefList = nil
			return fmt.Errorf("failed to unmarshal CapabilityIoCardCapabilityDefResponse as CapabilityIoCardCapabilityDefList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal CapabilityIoCardCapabilityDefResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal CapabilityIoCardCapabilityDefResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal CapabilityIoCardCapabilityDefResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CapabilityIoCardCapabilityDefResponse) MarshalJSON() ([]byte, error) {
	if src.CapabilityIoCardCapabilityDefList != nil {
		return json.Marshal(&src.CapabilityIoCardCapabilityDefList)
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
func (obj *CapabilityIoCardCapabilityDefResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CapabilityIoCardCapabilityDefList != nil {
		return obj.CapabilityIoCardCapabilityDefList
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

type NullableCapabilityIoCardCapabilityDefResponse struct {
	value *CapabilityIoCardCapabilityDefResponse
	isSet bool
}

func (v NullableCapabilityIoCardCapabilityDefResponse) Get() *CapabilityIoCardCapabilityDefResponse {
	return v.value
}

func (v *NullableCapabilityIoCardCapabilityDefResponse) Set(val *CapabilityIoCardCapabilityDefResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityIoCardCapabilityDefResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityIoCardCapabilityDefResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityIoCardCapabilityDefResponse(val *CapabilityIoCardCapabilityDefResponse) *NullableCapabilityIoCardCapabilityDefResponse {
	return &NullableCapabilityIoCardCapabilityDefResponse{value: val, isSet: true}
}

func (v NullableCapabilityIoCardCapabilityDefResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityIoCardCapabilityDefResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
