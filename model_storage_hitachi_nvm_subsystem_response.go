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

// StorageHitachiNvmSubsystemResponse - The response body of a HTTP GET request for the 'storage.HitachiNvmSubsystem' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'storage.HitachiNvmSubsystem' resources.
type StorageHitachiNvmSubsystemResponse struct {
	MoAggregateTransform           *MoAggregateTransform
	MoDocumentCount                *MoDocumentCount
	MoTagSummary                   *MoTagSummary
	StorageHitachiNvmSubsystemList *StorageHitachiNvmSubsystemList
}

// MoAggregateTransformAsStorageHitachiNvmSubsystemResponse is a convenience function that returns MoAggregateTransform wrapped in StorageHitachiNvmSubsystemResponse
func MoAggregateTransformAsStorageHitachiNvmSubsystemResponse(v *MoAggregateTransform) StorageHitachiNvmSubsystemResponse {
	return StorageHitachiNvmSubsystemResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsStorageHitachiNvmSubsystemResponse is a convenience function that returns MoDocumentCount wrapped in StorageHitachiNvmSubsystemResponse
func MoDocumentCountAsStorageHitachiNvmSubsystemResponse(v *MoDocumentCount) StorageHitachiNvmSubsystemResponse {
	return StorageHitachiNvmSubsystemResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsStorageHitachiNvmSubsystemResponse is a convenience function that returns MoTagSummary wrapped in StorageHitachiNvmSubsystemResponse
func MoTagSummaryAsStorageHitachiNvmSubsystemResponse(v *MoTagSummary) StorageHitachiNvmSubsystemResponse {
	return StorageHitachiNvmSubsystemResponse{
		MoTagSummary: v,
	}
}

// StorageHitachiNvmSubsystemListAsStorageHitachiNvmSubsystemResponse is a convenience function that returns StorageHitachiNvmSubsystemList wrapped in StorageHitachiNvmSubsystemResponse
func StorageHitachiNvmSubsystemListAsStorageHitachiNvmSubsystemResponse(v *StorageHitachiNvmSubsystemList) StorageHitachiNvmSubsystemResponse {
	return StorageHitachiNvmSubsystemResponse{
		StorageHitachiNvmSubsystemList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageHitachiNvmSubsystemResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal StorageHitachiNvmSubsystemResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal StorageHitachiNvmSubsystemResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal StorageHitachiNvmSubsystemResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.HitachiNvmSubsystem.List'
	if jsonDict["ObjectType"] == "storage.HitachiNvmSubsystem.List" {
		// try to unmarshal JSON data into StorageHitachiNvmSubsystemList
		err = json.Unmarshal(data, &dst.StorageHitachiNvmSubsystemList)
		if err == nil {
			return nil // data stored in dst.StorageHitachiNvmSubsystemList, return on the first match
		} else {
			dst.StorageHitachiNvmSubsystemList = nil
			return fmt.Errorf("failed to unmarshal StorageHitachiNvmSubsystemResponse as StorageHitachiNvmSubsystemList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageHitachiNvmSubsystemResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.StorageHitachiNvmSubsystemList != nil {
		return json.Marshal(&src.StorageHitachiNvmSubsystemList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageHitachiNvmSubsystemResponse) GetActualInstance() interface{} {
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

	if obj.StorageHitachiNvmSubsystemList != nil {
		return obj.StorageHitachiNvmSubsystemList
	}

	// all schemas are nil
	return nil
}

type NullableStorageHitachiNvmSubsystemResponse struct {
	value *StorageHitachiNvmSubsystemResponse
	isSet bool
}

func (v NullableStorageHitachiNvmSubsystemResponse) Get() *StorageHitachiNvmSubsystemResponse {
	return v.value
}

func (v *NullableStorageHitachiNvmSubsystemResponse) Set(val *StorageHitachiNvmSubsystemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiNvmSubsystemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiNvmSubsystemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiNvmSubsystemResponse(val *StorageHitachiNvmSubsystemResponse) *NullableStorageHitachiNvmSubsystemResponse {
	return &NullableStorageHitachiNvmSubsystemResponse{value: val, isSet: true}
}

func (v NullableStorageHitachiNvmSubsystemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiNvmSubsystemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
