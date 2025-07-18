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

// HyperflexServerFirmwareVersionEntryResponse - The response body of a HTTP GET request for the 'hyperflex.ServerFirmwareVersionEntry' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'hyperflex.ServerFirmwareVersionEntry' resources.
type HyperflexServerFirmwareVersionEntryResponse struct {
	HyperflexServerFirmwareVersionEntryList *HyperflexServerFirmwareVersionEntryList
	MoAggregateTransform                    *MoAggregateTransform
	MoDocumentCount                         *MoDocumentCount
	MoTagSummary                            *MoTagSummary
}

// HyperflexServerFirmwareVersionEntryListAsHyperflexServerFirmwareVersionEntryResponse is a convenience function that returns HyperflexServerFirmwareVersionEntryList wrapped in HyperflexServerFirmwareVersionEntryResponse
func HyperflexServerFirmwareVersionEntryListAsHyperflexServerFirmwareVersionEntryResponse(v *HyperflexServerFirmwareVersionEntryList) HyperflexServerFirmwareVersionEntryResponse {
	return HyperflexServerFirmwareVersionEntryResponse{
		HyperflexServerFirmwareVersionEntryList: v,
	}
}

// MoAggregateTransformAsHyperflexServerFirmwareVersionEntryResponse is a convenience function that returns MoAggregateTransform wrapped in HyperflexServerFirmwareVersionEntryResponse
func MoAggregateTransformAsHyperflexServerFirmwareVersionEntryResponse(v *MoAggregateTransform) HyperflexServerFirmwareVersionEntryResponse {
	return HyperflexServerFirmwareVersionEntryResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsHyperflexServerFirmwareVersionEntryResponse is a convenience function that returns MoDocumentCount wrapped in HyperflexServerFirmwareVersionEntryResponse
func MoDocumentCountAsHyperflexServerFirmwareVersionEntryResponse(v *MoDocumentCount) HyperflexServerFirmwareVersionEntryResponse {
	return HyperflexServerFirmwareVersionEntryResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsHyperflexServerFirmwareVersionEntryResponse is a convenience function that returns MoTagSummary wrapped in HyperflexServerFirmwareVersionEntryResponse
func MoTagSummaryAsHyperflexServerFirmwareVersionEntryResponse(v *MoTagSummary) HyperflexServerFirmwareVersionEntryResponse {
	return HyperflexServerFirmwareVersionEntryResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexServerFirmwareVersionEntryResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'hyperflex.ServerFirmwareVersionEntry.List'
	if jsonDict["ObjectType"] == "hyperflex.ServerFirmwareVersionEntry.List" {
		// try to unmarshal JSON data into HyperflexServerFirmwareVersionEntryList
		err = json.Unmarshal(data, &dst.HyperflexServerFirmwareVersionEntryList)
		if err == nil {
			return nil // data stored in dst.HyperflexServerFirmwareVersionEntryList, return on the first match
		} else {
			dst.HyperflexServerFirmwareVersionEntryList = nil
			return fmt.Errorf("failed to unmarshal HyperflexServerFirmwareVersionEntryResponse as HyperflexServerFirmwareVersionEntryList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexServerFirmwareVersionEntryResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexServerFirmwareVersionEntryResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexServerFirmwareVersionEntryResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexServerFirmwareVersionEntryResponse) MarshalJSON() ([]byte, error) {
	if src.HyperflexServerFirmwareVersionEntryList != nil {
		return json.Marshal(&src.HyperflexServerFirmwareVersionEntryList)
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
func (obj *HyperflexServerFirmwareVersionEntryResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HyperflexServerFirmwareVersionEntryList != nil {
		return obj.HyperflexServerFirmwareVersionEntryList
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

type NullableHyperflexServerFirmwareVersionEntryResponse struct {
	value *HyperflexServerFirmwareVersionEntryResponse
	isSet bool
}

func (v NullableHyperflexServerFirmwareVersionEntryResponse) Get() *HyperflexServerFirmwareVersionEntryResponse {
	return v.value
}

func (v *NullableHyperflexServerFirmwareVersionEntryResponse) Set(val *HyperflexServerFirmwareVersionEntryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServerFirmwareVersionEntryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServerFirmwareVersionEntryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServerFirmwareVersionEntryResponse(val *HyperflexServerFirmwareVersionEntryResponse) *NullableHyperflexServerFirmwareVersionEntryResponse {
	return &NullableHyperflexServerFirmwareVersionEntryResponse{value: val, isSet: true}
}

func (v NullableHyperflexServerFirmwareVersionEntryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServerFirmwareVersionEntryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
