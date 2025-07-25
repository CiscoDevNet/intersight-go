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

// StorageNetAppNodeCdpNeighborResponse - The response body of a HTTP GET request for the 'storage.NetAppNodeCdpNeighbor' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'storage.NetAppNodeCdpNeighbor' resources.
type StorageNetAppNodeCdpNeighborResponse struct {
	MoAggregateTransform             *MoAggregateTransform
	MoDocumentCount                  *MoDocumentCount
	MoTagSummary                     *MoTagSummary
	StorageNetAppNodeCdpNeighborList *StorageNetAppNodeCdpNeighborList
}

// MoAggregateTransformAsStorageNetAppNodeCdpNeighborResponse is a convenience function that returns MoAggregateTransform wrapped in StorageNetAppNodeCdpNeighborResponse
func MoAggregateTransformAsStorageNetAppNodeCdpNeighborResponse(v *MoAggregateTransform) StorageNetAppNodeCdpNeighborResponse {
	return StorageNetAppNodeCdpNeighborResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsStorageNetAppNodeCdpNeighborResponse is a convenience function that returns MoDocumentCount wrapped in StorageNetAppNodeCdpNeighborResponse
func MoDocumentCountAsStorageNetAppNodeCdpNeighborResponse(v *MoDocumentCount) StorageNetAppNodeCdpNeighborResponse {
	return StorageNetAppNodeCdpNeighborResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsStorageNetAppNodeCdpNeighborResponse is a convenience function that returns MoTagSummary wrapped in StorageNetAppNodeCdpNeighborResponse
func MoTagSummaryAsStorageNetAppNodeCdpNeighborResponse(v *MoTagSummary) StorageNetAppNodeCdpNeighborResponse {
	return StorageNetAppNodeCdpNeighborResponse{
		MoTagSummary: v,
	}
}

// StorageNetAppNodeCdpNeighborListAsStorageNetAppNodeCdpNeighborResponse is a convenience function that returns StorageNetAppNodeCdpNeighborList wrapped in StorageNetAppNodeCdpNeighborResponse
func StorageNetAppNodeCdpNeighborListAsStorageNetAppNodeCdpNeighborResponse(v *StorageNetAppNodeCdpNeighborList) StorageNetAppNodeCdpNeighborResponse {
	return StorageNetAppNodeCdpNeighborResponse{
		StorageNetAppNodeCdpNeighborList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageNetAppNodeCdpNeighborResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal StorageNetAppNodeCdpNeighborResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal StorageNetAppNodeCdpNeighborResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal StorageNetAppNodeCdpNeighborResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.NetAppNodeCdpNeighbor.List'
	if jsonDict["ObjectType"] == "storage.NetAppNodeCdpNeighbor.List" {
		// try to unmarshal JSON data into StorageNetAppNodeCdpNeighborList
		err = json.Unmarshal(data, &dst.StorageNetAppNodeCdpNeighborList)
		if err == nil {
			return nil // data stored in dst.StorageNetAppNodeCdpNeighborList, return on the first match
		} else {
			dst.StorageNetAppNodeCdpNeighborList = nil
			return fmt.Errorf("failed to unmarshal StorageNetAppNodeCdpNeighborResponse as StorageNetAppNodeCdpNeighborList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageNetAppNodeCdpNeighborResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.StorageNetAppNodeCdpNeighborList != nil {
		return json.Marshal(&src.StorageNetAppNodeCdpNeighborList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageNetAppNodeCdpNeighborResponse) GetActualInstance() interface{} {
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

	if obj.StorageNetAppNodeCdpNeighborList != nil {
		return obj.StorageNetAppNodeCdpNeighborList
	}

	// all schemas are nil
	return nil
}

type NullableStorageNetAppNodeCdpNeighborResponse struct {
	value *StorageNetAppNodeCdpNeighborResponse
	isSet bool
}

func (v NullableStorageNetAppNodeCdpNeighborResponse) Get() *StorageNetAppNodeCdpNeighborResponse {
	return v.value
}

func (v *NullableStorageNetAppNodeCdpNeighborResponse) Set(val *StorageNetAppNodeCdpNeighborResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNodeCdpNeighborResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNodeCdpNeighborResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNodeCdpNeighborResponse(val *StorageNetAppNodeCdpNeighborResponse) *NullableStorageNetAppNodeCdpNeighborResponse {
	return &NullableStorageNetAppNodeCdpNeighborResponse{value: val, isSet: true}
}

func (v NullableStorageNetAppNodeCdpNeighborResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNodeCdpNeighborResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
