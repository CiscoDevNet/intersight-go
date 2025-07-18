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

// ProcessorUnitResponse - The response body of a HTTP GET request for the 'processor.Unit' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'processor.Unit' resources.
type ProcessorUnitResponse struct {
	MoAggregateTransform *MoAggregateTransform
	MoDocumentCount      *MoDocumentCount
	MoTagSummary         *MoTagSummary
	ProcessorUnitList    *ProcessorUnitList
}

// MoAggregateTransformAsProcessorUnitResponse is a convenience function that returns MoAggregateTransform wrapped in ProcessorUnitResponse
func MoAggregateTransformAsProcessorUnitResponse(v *MoAggregateTransform) ProcessorUnitResponse {
	return ProcessorUnitResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsProcessorUnitResponse is a convenience function that returns MoDocumentCount wrapped in ProcessorUnitResponse
func MoDocumentCountAsProcessorUnitResponse(v *MoDocumentCount) ProcessorUnitResponse {
	return ProcessorUnitResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsProcessorUnitResponse is a convenience function that returns MoTagSummary wrapped in ProcessorUnitResponse
func MoTagSummaryAsProcessorUnitResponse(v *MoTagSummary) ProcessorUnitResponse {
	return ProcessorUnitResponse{
		MoTagSummary: v,
	}
}

// ProcessorUnitListAsProcessorUnitResponse is a convenience function that returns ProcessorUnitList wrapped in ProcessorUnitResponse
func ProcessorUnitListAsProcessorUnitResponse(v *ProcessorUnitList) ProcessorUnitResponse {
	return ProcessorUnitResponse{
		ProcessorUnitList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ProcessorUnitResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("failed to unmarshal ProcessorUnitResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ProcessorUnitResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ProcessorUnitResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'processor.Unit.List'
	if jsonDict["ObjectType"] == "processor.Unit.List" {
		// try to unmarshal JSON data into ProcessorUnitList
		err = json.Unmarshal(data, &dst.ProcessorUnitList)
		if err == nil {
			return nil // data stored in dst.ProcessorUnitList, return on the first match
		} else {
			dst.ProcessorUnitList = nil
			return fmt.Errorf("failed to unmarshal ProcessorUnitResponse as ProcessorUnitList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ProcessorUnitResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.ProcessorUnitList != nil {
		return json.Marshal(&src.ProcessorUnitList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ProcessorUnitResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
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

	if obj.ProcessorUnitList != nil {
		return obj.ProcessorUnitList
	}

	// all schemas are nil
	return nil
}

type NullableProcessorUnitResponse struct {
	value *ProcessorUnitResponse
	isSet bool
}

func (v NullableProcessorUnitResponse) Get() *ProcessorUnitResponse {
	return v.value
}

func (v *NullableProcessorUnitResponse) Set(val *ProcessorUnitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorUnitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorUnitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorUnitResponse(val *ProcessorUnitResponse) *NullableProcessorUnitResponse {
	return &NullableProcessorUnitResponse{value: val, isSet: true}
}

func (v NullableProcessorUnitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorUnitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
