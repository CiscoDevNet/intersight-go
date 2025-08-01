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

// StorageSasPortResponse - The response body of a HTTP GET request for the 'storage.SasPort' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'storage.SasPort' resources.
type StorageSasPortResponse struct {
	MoAggregateTransform *MoAggregateTransform
	MoDocumentCount      *MoDocumentCount
	MoTagSummary         *MoTagSummary
	StorageSasPortList   *StorageSasPortList
}

// MoAggregateTransformAsStorageSasPortResponse is a convenience function that returns MoAggregateTransform wrapped in StorageSasPortResponse
func MoAggregateTransformAsStorageSasPortResponse(v *MoAggregateTransform) StorageSasPortResponse {
	return StorageSasPortResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsStorageSasPortResponse is a convenience function that returns MoDocumentCount wrapped in StorageSasPortResponse
func MoDocumentCountAsStorageSasPortResponse(v *MoDocumentCount) StorageSasPortResponse {
	return StorageSasPortResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsStorageSasPortResponse is a convenience function that returns MoTagSummary wrapped in StorageSasPortResponse
func MoTagSummaryAsStorageSasPortResponse(v *MoTagSummary) StorageSasPortResponse {
	return StorageSasPortResponse{
		MoTagSummary: v,
	}
}

// StorageSasPortListAsStorageSasPortResponse is a convenience function that returns StorageSasPortList wrapped in StorageSasPortResponse
func StorageSasPortListAsStorageSasPortResponse(v *StorageSasPortList) StorageSasPortResponse {
	return StorageSasPortResponse{
		StorageSasPortList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageSasPortResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal StorageSasPortResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal StorageSasPortResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal StorageSasPortResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.SasPort.List'
	if jsonDict["ObjectType"] == "storage.SasPort.List" {
		// try to unmarshal JSON data into StorageSasPortList
		err = json.Unmarshal(data, &dst.StorageSasPortList)
		if err == nil {
			return nil // data stored in dst.StorageSasPortList, return on the first match
		} else {
			dst.StorageSasPortList = nil
			return fmt.Errorf("failed to unmarshal StorageSasPortResponse as StorageSasPortList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageSasPortResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.StorageSasPortList != nil {
		return json.Marshal(&src.StorageSasPortList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageSasPortResponse) GetActualInstance() interface{} {
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

	if obj.StorageSasPortList != nil {
		return obj.StorageSasPortList
	}

	// all schemas are nil
	return nil
}

type NullableStorageSasPortResponse struct {
	value *StorageSasPortResponse
	isSet bool
}

func (v NullableStorageSasPortResponse) Get() *StorageSasPortResponse {
	return v.value
}

func (v *NullableStorageSasPortResponse) Set(val *StorageSasPortResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSasPortResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSasPortResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSasPortResponse(val *StorageSasPortResponse) *NullableStorageSasPortResponse {
	return &NullableStorageSasPortResponse{value: val, isSet: true}
}

func (v NullableStorageSasPortResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSasPortResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
