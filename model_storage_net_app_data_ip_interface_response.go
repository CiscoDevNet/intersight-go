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

// StorageNetAppDataIpInterfaceResponse - The response body of a HTTP GET request for the 'storage.NetAppDataIpInterface' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'storage.NetAppDataIpInterface' resources.
type StorageNetAppDataIpInterfaceResponse struct {
	MoAggregateTransform             *MoAggregateTransform
	MoDocumentCount                  *MoDocumentCount
	MoTagSummary                     *MoTagSummary
	StorageNetAppDataIpInterfaceList *StorageNetAppDataIpInterfaceList
}

// MoAggregateTransformAsStorageNetAppDataIpInterfaceResponse is a convenience function that returns MoAggregateTransform wrapped in StorageNetAppDataIpInterfaceResponse
func MoAggregateTransformAsStorageNetAppDataIpInterfaceResponse(v *MoAggregateTransform) StorageNetAppDataIpInterfaceResponse {
	return StorageNetAppDataIpInterfaceResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsStorageNetAppDataIpInterfaceResponse is a convenience function that returns MoDocumentCount wrapped in StorageNetAppDataIpInterfaceResponse
func MoDocumentCountAsStorageNetAppDataIpInterfaceResponse(v *MoDocumentCount) StorageNetAppDataIpInterfaceResponse {
	return StorageNetAppDataIpInterfaceResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsStorageNetAppDataIpInterfaceResponse is a convenience function that returns MoTagSummary wrapped in StorageNetAppDataIpInterfaceResponse
func MoTagSummaryAsStorageNetAppDataIpInterfaceResponse(v *MoTagSummary) StorageNetAppDataIpInterfaceResponse {
	return StorageNetAppDataIpInterfaceResponse{
		MoTagSummary: v,
	}
}

// StorageNetAppDataIpInterfaceListAsStorageNetAppDataIpInterfaceResponse is a convenience function that returns StorageNetAppDataIpInterfaceList wrapped in StorageNetAppDataIpInterfaceResponse
func StorageNetAppDataIpInterfaceListAsStorageNetAppDataIpInterfaceResponse(v *StorageNetAppDataIpInterfaceList) StorageNetAppDataIpInterfaceResponse {
	return StorageNetAppDataIpInterfaceResponse{
		StorageNetAppDataIpInterfaceList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageNetAppDataIpInterfaceResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal StorageNetAppDataIpInterfaceResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal StorageNetAppDataIpInterfaceResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal StorageNetAppDataIpInterfaceResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.NetAppDataIpInterface.List'
	if jsonDict["ObjectType"] == "storage.NetAppDataIpInterface.List" {
		// try to unmarshal JSON data into StorageNetAppDataIpInterfaceList
		err = json.Unmarshal(data, &dst.StorageNetAppDataIpInterfaceList)
		if err == nil {
			return nil // data stored in dst.StorageNetAppDataIpInterfaceList, return on the first match
		} else {
			dst.StorageNetAppDataIpInterfaceList = nil
			return fmt.Errorf("failed to unmarshal StorageNetAppDataIpInterfaceResponse as StorageNetAppDataIpInterfaceList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageNetAppDataIpInterfaceResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.StorageNetAppDataIpInterfaceList != nil {
		return json.Marshal(&src.StorageNetAppDataIpInterfaceList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageNetAppDataIpInterfaceResponse) GetActualInstance() interface{} {
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

	if obj.StorageNetAppDataIpInterfaceList != nil {
		return obj.StorageNetAppDataIpInterfaceList
	}

	// all schemas are nil
	return nil
}

type NullableStorageNetAppDataIpInterfaceResponse struct {
	value *StorageNetAppDataIpInterfaceResponse
	isSet bool
}

func (v NullableStorageNetAppDataIpInterfaceResponse) Get() *StorageNetAppDataIpInterfaceResponse {
	return v.value
}

func (v *NullableStorageNetAppDataIpInterfaceResponse) Set(val *StorageNetAppDataIpInterfaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppDataIpInterfaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppDataIpInterfaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppDataIpInterfaceResponse(val *StorageNetAppDataIpInterfaceResponse) *NullableStorageNetAppDataIpInterfaceResponse {
	return &NullableStorageNetAppDataIpInterfaceResponse{value: val, isSet: true}
}

func (v NullableStorageNetAppDataIpInterfaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppDataIpInterfaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
