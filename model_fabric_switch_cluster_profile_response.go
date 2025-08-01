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

// FabricSwitchClusterProfileResponse - The response body of a HTTP GET request for the 'fabric.SwitchClusterProfile' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'fabric.SwitchClusterProfile' resources.
type FabricSwitchClusterProfileResponse struct {
	FabricSwitchClusterProfileList *FabricSwitchClusterProfileList
	MoAggregateTransform           *MoAggregateTransform
	MoDocumentCount                *MoDocumentCount
	MoTagSummary                   *MoTagSummary
}

// FabricSwitchClusterProfileListAsFabricSwitchClusterProfileResponse is a convenience function that returns FabricSwitchClusterProfileList wrapped in FabricSwitchClusterProfileResponse
func FabricSwitchClusterProfileListAsFabricSwitchClusterProfileResponse(v *FabricSwitchClusterProfileList) FabricSwitchClusterProfileResponse {
	return FabricSwitchClusterProfileResponse{
		FabricSwitchClusterProfileList: v,
	}
}

// MoAggregateTransformAsFabricSwitchClusterProfileResponse is a convenience function that returns MoAggregateTransform wrapped in FabricSwitchClusterProfileResponse
func MoAggregateTransformAsFabricSwitchClusterProfileResponse(v *MoAggregateTransform) FabricSwitchClusterProfileResponse {
	return FabricSwitchClusterProfileResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsFabricSwitchClusterProfileResponse is a convenience function that returns MoDocumentCount wrapped in FabricSwitchClusterProfileResponse
func MoDocumentCountAsFabricSwitchClusterProfileResponse(v *MoDocumentCount) FabricSwitchClusterProfileResponse {
	return FabricSwitchClusterProfileResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsFabricSwitchClusterProfileResponse is a convenience function that returns MoTagSummary wrapped in FabricSwitchClusterProfileResponse
func MoTagSummaryAsFabricSwitchClusterProfileResponse(v *MoTagSummary) FabricSwitchClusterProfileResponse {
	return FabricSwitchClusterProfileResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FabricSwitchClusterProfileResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'fabric.SwitchClusterProfile.List'
	if jsonDict["ObjectType"] == "fabric.SwitchClusterProfile.List" {
		// try to unmarshal JSON data into FabricSwitchClusterProfileList
		err = json.Unmarshal(data, &dst.FabricSwitchClusterProfileList)
		if err == nil {
			return nil // data stored in dst.FabricSwitchClusterProfileList, return on the first match
		} else {
			dst.FabricSwitchClusterProfileList = nil
			return fmt.Errorf("failed to unmarshal FabricSwitchClusterProfileResponse as FabricSwitchClusterProfileList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FabricSwitchClusterProfileResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FabricSwitchClusterProfileResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FabricSwitchClusterProfileResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FabricSwitchClusterProfileResponse) MarshalJSON() ([]byte, error) {
	if src.FabricSwitchClusterProfileList != nil {
		return json.Marshal(&src.FabricSwitchClusterProfileList)
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
func (obj *FabricSwitchClusterProfileResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FabricSwitchClusterProfileList != nil {
		return obj.FabricSwitchClusterProfileList
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

type NullableFabricSwitchClusterProfileResponse struct {
	value *FabricSwitchClusterProfileResponse
	isSet bool
}

func (v NullableFabricSwitchClusterProfileResponse) Get() *FabricSwitchClusterProfileResponse {
	return v.value
}

func (v *NullableFabricSwitchClusterProfileResponse) Set(val *FabricSwitchClusterProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSwitchClusterProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSwitchClusterProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSwitchClusterProfileResponse(val *FabricSwitchClusterProfileResponse) *NullableFabricSwitchClusterProfileResponse {
	return &NullableFabricSwitchClusterProfileResponse{value: val, isSet: true}
}

func (v NullableFabricSwitchClusterProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSwitchClusterProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
