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

// CapabilityStorageControllerUpdateConstraintMetaResponse - The response body of a HTTP GET request for the 'capability.StorageControllerUpdateConstraintMeta' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'capability.StorageControllerUpdateConstraintMeta' resources.
type CapabilityStorageControllerUpdateConstraintMetaResponse struct {
	CapabilityStorageControllerUpdateConstraintMetaList *CapabilityStorageControllerUpdateConstraintMetaList
	MoAggregateTransform                                *MoAggregateTransform
	MoDocumentCount                                     *MoDocumentCount
	MoTagSummary                                        *MoTagSummary
}

// CapabilityStorageControllerUpdateConstraintMetaListAsCapabilityStorageControllerUpdateConstraintMetaResponse is a convenience function that returns CapabilityStorageControllerUpdateConstraintMetaList wrapped in CapabilityStorageControllerUpdateConstraintMetaResponse
func CapabilityStorageControllerUpdateConstraintMetaListAsCapabilityStorageControllerUpdateConstraintMetaResponse(v *CapabilityStorageControllerUpdateConstraintMetaList) CapabilityStorageControllerUpdateConstraintMetaResponse {
	return CapabilityStorageControllerUpdateConstraintMetaResponse{
		CapabilityStorageControllerUpdateConstraintMetaList: v,
	}
}

// MoAggregateTransformAsCapabilityStorageControllerUpdateConstraintMetaResponse is a convenience function that returns MoAggregateTransform wrapped in CapabilityStorageControllerUpdateConstraintMetaResponse
func MoAggregateTransformAsCapabilityStorageControllerUpdateConstraintMetaResponse(v *MoAggregateTransform) CapabilityStorageControllerUpdateConstraintMetaResponse {
	return CapabilityStorageControllerUpdateConstraintMetaResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsCapabilityStorageControllerUpdateConstraintMetaResponse is a convenience function that returns MoDocumentCount wrapped in CapabilityStorageControllerUpdateConstraintMetaResponse
func MoDocumentCountAsCapabilityStorageControllerUpdateConstraintMetaResponse(v *MoDocumentCount) CapabilityStorageControllerUpdateConstraintMetaResponse {
	return CapabilityStorageControllerUpdateConstraintMetaResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsCapabilityStorageControllerUpdateConstraintMetaResponse is a convenience function that returns MoTagSummary wrapped in CapabilityStorageControllerUpdateConstraintMetaResponse
func MoTagSummaryAsCapabilityStorageControllerUpdateConstraintMetaResponse(v *MoTagSummary) CapabilityStorageControllerUpdateConstraintMetaResponse {
	return CapabilityStorageControllerUpdateConstraintMetaResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CapabilityStorageControllerUpdateConstraintMetaResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'capability.StorageControllerUpdateConstraintMeta.List'
	if jsonDict["ObjectType"] == "capability.StorageControllerUpdateConstraintMeta.List" {
		// try to unmarshal JSON data into CapabilityStorageControllerUpdateConstraintMetaList
		err = json.Unmarshal(data, &dst.CapabilityStorageControllerUpdateConstraintMetaList)
		if err == nil {
			return nil // data stored in dst.CapabilityStorageControllerUpdateConstraintMetaList, return on the first match
		} else {
			dst.CapabilityStorageControllerUpdateConstraintMetaList = nil
			return fmt.Errorf("failed to unmarshal CapabilityStorageControllerUpdateConstraintMetaResponse as CapabilityStorageControllerUpdateConstraintMetaList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal CapabilityStorageControllerUpdateConstraintMetaResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal CapabilityStorageControllerUpdateConstraintMetaResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal CapabilityStorageControllerUpdateConstraintMetaResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CapabilityStorageControllerUpdateConstraintMetaResponse) MarshalJSON() ([]byte, error) {
	if src.CapabilityStorageControllerUpdateConstraintMetaList != nil {
		return json.Marshal(&src.CapabilityStorageControllerUpdateConstraintMetaList)
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
func (obj *CapabilityStorageControllerUpdateConstraintMetaResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CapabilityStorageControllerUpdateConstraintMetaList != nil {
		return obj.CapabilityStorageControllerUpdateConstraintMetaList
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

type NullableCapabilityStorageControllerUpdateConstraintMetaResponse struct {
	value *CapabilityStorageControllerUpdateConstraintMetaResponse
	isSet bool
}

func (v NullableCapabilityStorageControllerUpdateConstraintMetaResponse) Get() *CapabilityStorageControllerUpdateConstraintMetaResponse {
	return v.value
}

func (v *NullableCapabilityStorageControllerUpdateConstraintMetaResponse) Set(val *CapabilityStorageControllerUpdateConstraintMetaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityStorageControllerUpdateConstraintMetaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityStorageControllerUpdateConstraintMetaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityStorageControllerUpdateConstraintMetaResponse(val *CapabilityStorageControllerUpdateConstraintMetaResponse) *NullableCapabilityStorageControllerUpdateConstraintMetaResponse {
	return &NullableCapabilityStorageControllerUpdateConstraintMetaResponse{value: val, isSet: true}
}

func (v NullableCapabilityStorageControllerUpdateConstraintMetaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityStorageControllerUpdateConstraintMetaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
