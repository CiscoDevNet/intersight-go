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

// ComputePersonalityResponse - The response body of a HTTP GET request for the 'compute.Personality' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'compute.Personality' resources.
type ComputePersonalityResponse struct {
	ComputePersonalityList *ComputePersonalityList
	MoAggregateTransform   *MoAggregateTransform
	MoDocumentCount        *MoDocumentCount
	MoTagSummary           *MoTagSummary
}

// ComputePersonalityListAsComputePersonalityResponse is a convenience function that returns ComputePersonalityList wrapped in ComputePersonalityResponse
func ComputePersonalityListAsComputePersonalityResponse(v *ComputePersonalityList) ComputePersonalityResponse {
	return ComputePersonalityResponse{
		ComputePersonalityList: v,
	}
}

// MoAggregateTransformAsComputePersonalityResponse is a convenience function that returns MoAggregateTransform wrapped in ComputePersonalityResponse
func MoAggregateTransformAsComputePersonalityResponse(v *MoAggregateTransform) ComputePersonalityResponse {
	return ComputePersonalityResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsComputePersonalityResponse is a convenience function that returns MoDocumentCount wrapped in ComputePersonalityResponse
func MoDocumentCountAsComputePersonalityResponse(v *MoDocumentCount) ComputePersonalityResponse {
	return ComputePersonalityResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsComputePersonalityResponse is a convenience function that returns MoTagSummary wrapped in ComputePersonalityResponse
func MoTagSummaryAsComputePersonalityResponse(v *MoTagSummary) ComputePersonalityResponse {
	return ComputePersonalityResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ComputePersonalityResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'compute.Personality.List'
	if jsonDict["ObjectType"] == "compute.Personality.List" {
		// try to unmarshal JSON data into ComputePersonalityList
		err = json.Unmarshal(data, &dst.ComputePersonalityList)
		if err == nil {
			return nil // data stored in dst.ComputePersonalityList, return on the first match
		} else {
			dst.ComputePersonalityList = nil
			return fmt.Errorf("failed to unmarshal ComputePersonalityResponse as ComputePersonalityList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ComputePersonalityResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ComputePersonalityResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ComputePersonalityResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ComputePersonalityResponse) MarshalJSON() ([]byte, error) {
	if src.ComputePersonalityList != nil {
		return json.Marshal(&src.ComputePersonalityList)
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
func (obj *ComputePersonalityResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ComputePersonalityList != nil {
		return obj.ComputePersonalityList
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

type NullableComputePersonalityResponse struct {
	value *ComputePersonalityResponse
	isSet bool
}

func (v NullableComputePersonalityResponse) Get() *ComputePersonalityResponse {
	return v.value
}

func (v *NullableComputePersonalityResponse) Set(val *ComputePersonalityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableComputePersonalityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableComputePersonalityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputePersonalityResponse(val *ComputePersonalityResponse) *NullableComputePersonalityResponse {
	return &NullableComputePersonalityResponse{value: val, isSet: true}
}

func (v NullableComputePersonalityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputePersonalityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
