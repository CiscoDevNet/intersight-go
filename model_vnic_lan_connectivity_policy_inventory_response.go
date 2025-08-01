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

// VnicLanConnectivityPolicyInventoryResponse - The response body of a HTTP GET request for the 'vnic.LanConnectivityPolicyInventory' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'vnic.LanConnectivityPolicyInventory' resources.
type VnicLanConnectivityPolicyInventoryResponse struct {
	MoAggregateTransform                   *MoAggregateTransform
	MoDocumentCount                        *MoDocumentCount
	MoTagSummary                           *MoTagSummary
	VnicLanConnectivityPolicyInventoryList *VnicLanConnectivityPolicyInventoryList
}

// MoAggregateTransformAsVnicLanConnectivityPolicyInventoryResponse is a convenience function that returns MoAggregateTransform wrapped in VnicLanConnectivityPolicyInventoryResponse
func MoAggregateTransformAsVnicLanConnectivityPolicyInventoryResponse(v *MoAggregateTransform) VnicLanConnectivityPolicyInventoryResponse {
	return VnicLanConnectivityPolicyInventoryResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsVnicLanConnectivityPolicyInventoryResponse is a convenience function that returns MoDocumentCount wrapped in VnicLanConnectivityPolicyInventoryResponse
func MoDocumentCountAsVnicLanConnectivityPolicyInventoryResponse(v *MoDocumentCount) VnicLanConnectivityPolicyInventoryResponse {
	return VnicLanConnectivityPolicyInventoryResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsVnicLanConnectivityPolicyInventoryResponse is a convenience function that returns MoTagSummary wrapped in VnicLanConnectivityPolicyInventoryResponse
func MoTagSummaryAsVnicLanConnectivityPolicyInventoryResponse(v *MoTagSummary) VnicLanConnectivityPolicyInventoryResponse {
	return VnicLanConnectivityPolicyInventoryResponse{
		MoTagSummary: v,
	}
}

// VnicLanConnectivityPolicyInventoryListAsVnicLanConnectivityPolicyInventoryResponse is a convenience function that returns VnicLanConnectivityPolicyInventoryList wrapped in VnicLanConnectivityPolicyInventoryResponse
func VnicLanConnectivityPolicyInventoryListAsVnicLanConnectivityPolicyInventoryResponse(v *VnicLanConnectivityPolicyInventoryList) VnicLanConnectivityPolicyInventoryResponse {
	return VnicLanConnectivityPolicyInventoryResponse{
		VnicLanConnectivityPolicyInventoryList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VnicLanConnectivityPolicyInventoryResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal VnicLanConnectivityPolicyInventoryResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal VnicLanConnectivityPolicyInventoryResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal VnicLanConnectivityPolicyInventoryResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'vnic.LanConnectivityPolicyInventory.List'
	if jsonDict["ObjectType"] == "vnic.LanConnectivityPolicyInventory.List" {
		// try to unmarshal JSON data into VnicLanConnectivityPolicyInventoryList
		err = json.Unmarshal(data, &dst.VnicLanConnectivityPolicyInventoryList)
		if err == nil {
			return nil // data stored in dst.VnicLanConnectivityPolicyInventoryList, return on the first match
		} else {
			dst.VnicLanConnectivityPolicyInventoryList = nil
			return fmt.Errorf("failed to unmarshal VnicLanConnectivityPolicyInventoryResponse as VnicLanConnectivityPolicyInventoryList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VnicLanConnectivityPolicyInventoryResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.VnicLanConnectivityPolicyInventoryList != nil {
		return json.Marshal(&src.VnicLanConnectivityPolicyInventoryList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VnicLanConnectivityPolicyInventoryResponse) GetActualInstance() interface{} {
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

	if obj.VnicLanConnectivityPolicyInventoryList != nil {
		return obj.VnicLanConnectivityPolicyInventoryList
	}

	// all schemas are nil
	return nil
}

type NullableVnicLanConnectivityPolicyInventoryResponse struct {
	value *VnicLanConnectivityPolicyInventoryResponse
	isSet bool
}

func (v NullableVnicLanConnectivityPolicyInventoryResponse) Get() *VnicLanConnectivityPolicyInventoryResponse {
	return v.value
}

func (v *NullableVnicLanConnectivityPolicyInventoryResponse) Set(val *VnicLanConnectivityPolicyInventoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicLanConnectivityPolicyInventoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicLanConnectivityPolicyInventoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicLanConnectivityPolicyInventoryResponse(val *VnicLanConnectivityPolicyInventoryResponse) *NullableVnicLanConnectivityPolicyInventoryResponse {
	return &NullableVnicLanConnectivityPolicyInventoryResponse{value: val, isSet: true}
}

func (v NullableVnicLanConnectivityPolicyInventoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicLanConnectivityPolicyInventoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
