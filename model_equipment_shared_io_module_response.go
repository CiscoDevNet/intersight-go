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

// EquipmentSharedIoModuleResponse - The response body of a HTTP GET request for the 'equipment.SharedIoModule' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'equipment.SharedIoModule' resources.
type EquipmentSharedIoModuleResponse struct {
	EquipmentSharedIoModuleList *EquipmentSharedIoModuleList
	MoAggregateTransform        *MoAggregateTransform
	MoDocumentCount             *MoDocumentCount
	MoTagSummary                *MoTagSummary
}

// EquipmentSharedIoModuleListAsEquipmentSharedIoModuleResponse is a convenience function that returns EquipmentSharedIoModuleList wrapped in EquipmentSharedIoModuleResponse
func EquipmentSharedIoModuleListAsEquipmentSharedIoModuleResponse(v *EquipmentSharedIoModuleList) EquipmentSharedIoModuleResponse {
	return EquipmentSharedIoModuleResponse{
		EquipmentSharedIoModuleList: v,
	}
}

// MoAggregateTransformAsEquipmentSharedIoModuleResponse is a convenience function that returns MoAggregateTransform wrapped in EquipmentSharedIoModuleResponse
func MoAggregateTransformAsEquipmentSharedIoModuleResponse(v *MoAggregateTransform) EquipmentSharedIoModuleResponse {
	return EquipmentSharedIoModuleResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsEquipmentSharedIoModuleResponse is a convenience function that returns MoDocumentCount wrapped in EquipmentSharedIoModuleResponse
func MoDocumentCountAsEquipmentSharedIoModuleResponse(v *MoDocumentCount) EquipmentSharedIoModuleResponse {
	return EquipmentSharedIoModuleResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsEquipmentSharedIoModuleResponse is a convenience function that returns MoTagSummary wrapped in EquipmentSharedIoModuleResponse
func MoTagSummaryAsEquipmentSharedIoModuleResponse(v *MoTagSummary) EquipmentSharedIoModuleResponse {
	return EquipmentSharedIoModuleResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EquipmentSharedIoModuleResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'equipment.SharedIoModule.List'
	if jsonDict["ObjectType"] == "equipment.SharedIoModule.List" {
		// try to unmarshal JSON data into EquipmentSharedIoModuleList
		err = json.Unmarshal(data, &dst.EquipmentSharedIoModuleList)
		if err == nil {
			return nil // data stored in dst.EquipmentSharedIoModuleList, return on the first match
		} else {
			dst.EquipmentSharedIoModuleList = nil
			return fmt.Errorf("failed to unmarshal EquipmentSharedIoModuleResponse as EquipmentSharedIoModuleList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal EquipmentSharedIoModuleResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal EquipmentSharedIoModuleResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal EquipmentSharedIoModuleResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EquipmentSharedIoModuleResponse) MarshalJSON() ([]byte, error) {
	if src.EquipmentSharedIoModuleList != nil {
		return json.Marshal(&src.EquipmentSharedIoModuleList)
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
func (obj *EquipmentSharedIoModuleResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EquipmentSharedIoModuleList != nil {
		return obj.EquipmentSharedIoModuleList
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

type NullableEquipmentSharedIoModuleResponse struct {
	value *EquipmentSharedIoModuleResponse
	isSet bool
}

func (v NullableEquipmentSharedIoModuleResponse) Get() *EquipmentSharedIoModuleResponse {
	return v.value
}

func (v *NullableEquipmentSharedIoModuleResponse) Set(val *EquipmentSharedIoModuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSharedIoModuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSharedIoModuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSharedIoModuleResponse(val *EquipmentSharedIoModuleResponse) *NullableEquipmentSharedIoModuleResponse {
	return &NullableEquipmentSharedIoModuleResponse{value: val, isSet: true}
}

func (v NullableEquipmentSharedIoModuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSharedIoModuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
