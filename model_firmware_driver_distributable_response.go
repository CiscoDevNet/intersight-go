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

// FirmwareDriverDistributableResponse - The response body of a HTTP GET request for the 'firmware.DriverDistributable' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'firmware.DriverDistributable' resources.
type FirmwareDriverDistributableResponse struct {
	FirmwareDriverDistributableList *FirmwareDriverDistributableList
	MoAggregateTransform            *MoAggregateTransform
	MoDocumentCount                 *MoDocumentCount
	MoTagSummary                    *MoTagSummary
}

// FirmwareDriverDistributableListAsFirmwareDriverDistributableResponse is a convenience function that returns FirmwareDriverDistributableList wrapped in FirmwareDriverDistributableResponse
func FirmwareDriverDistributableListAsFirmwareDriverDistributableResponse(v *FirmwareDriverDistributableList) FirmwareDriverDistributableResponse {
	return FirmwareDriverDistributableResponse{
		FirmwareDriverDistributableList: v,
	}
}

// MoAggregateTransformAsFirmwareDriverDistributableResponse is a convenience function that returns MoAggregateTransform wrapped in FirmwareDriverDistributableResponse
func MoAggregateTransformAsFirmwareDriverDistributableResponse(v *MoAggregateTransform) FirmwareDriverDistributableResponse {
	return FirmwareDriverDistributableResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsFirmwareDriverDistributableResponse is a convenience function that returns MoDocumentCount wrapped in FirmwareDriverDistributableResponse
func MoDocumentCountAsFirmwareDriverDistributableResponse(v *MoDocumentCount) FirmwareDriverDistributableResponse {
	return FirmwareDriverDistributableResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsFirmwareDriverDistributableResponse is a convenience function that returns MoTagSummary wrapped in FirmwareDriverDistributableResponse
func MoTagSummaryAsFirmwareDriverDistributableResponse(v *MoTagSummary) FirmwareDriverDistributableResponse {
	return FirmwareDriverDistributableResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FirmwareDriverDistributableResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'firmware.DriverDistributable.List'
	if jsonDict["ObjectType"] == "firmware.DriverDistributable.List" {
		// try to unmarshal JSON data into FirmwareDriverDistributableList
		err = json.Unmarshal(data, &dst.FirmwareDriverDistributableList)
		if err == nil {
			return nil // data stored in dst.FirmwareDriverDistributableList, return on the first match
		} else {
			dst.FirmwareDriverDistributableList = nil
			return fmt.Errorf("failed to unmarshal FirmwareDriverDistributableResponse as FirmwareDriverDistributableList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FirmwareDriverDistributableResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FirmwareDriverDistributableResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FirmwareDriverDistributableResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FirmwareDriverDistributableResponse) MarshalJSON() ([]byte, error) {
	if src.FirmwareDriverDistributableList != nil {
		return json.Marshal(&src.FirmwareDriverDistributableList)
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
func (obj *FirmwareDriverDistributableResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FirmwareDriverDistributableList != nil {
		return obj.FirmwareDriverDistributableList
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

type NullableFirmwareDriverDistributableResponse struct {
	value *FirmwareDriverDistributableResponse
	isSet bool
}

func (v NullableFirmwareDriverDistributableResponse) Get() *FirmwareDriverDistributableResponse {
	return v.value
}

func (v *NullableFirmwareDriverDistributableResponse) Set(val *FirmwareDriverDistributableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareDriverDistributableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareDriverDistributableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareDriverDistributableResponse(val *FirmwareDriverDistributableResponse) *NullableFirmwareDriverDistributableResponse {
	return &NullableFirmwareDriverDistributableResponse{value: val, isSet: true}
}

func (v NullableFirmwareDriverDistributableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareDriverDistributableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
