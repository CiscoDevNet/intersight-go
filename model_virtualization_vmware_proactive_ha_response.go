/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-15830
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// VirtualizationVmwareProactiveHaResponse - The response body of a HTTP GET request for the 'virtualization.VmwareProactiveHa' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'virtualization.VmwareProactiveHa' resources.
type VirtualizationVmwareProactiveHaResponse struct {
	MoAggregateTransform                *MoAggregateTransform
	MoDocumentCount                     *MoDocumentCount
	MoTagSummary                        *MoTagSummary
	VirtualizationVmwareProactiveHaList *VirtualizationVmwareProactiveHaList
}

// MoAggregateTransformAsVirtualizationVmwareProactiveHaResponse is a convenience function that returns MoAggregateTransform wrapped in VirtualizationVmwareProactiveHaResponse
func MoAggregateTransformAsVirtualizationVmwareProactiveHaResponse(v *MoAggregateTransform) VirtualizationVmwareProactiveHaResponse {
	return VirtualizationVmwareProactiveHaResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsVirtualizationVmwareProactiveHaResponse is a convenience function that returns MoDocumentCount wrapped in VirtualizationVmwareProactiveHaResponse
func MoDocumentCountAsVirtualizationVmwareProactiveHaResponse(v *MoDocumentCount) VirtualizationVmwareProactiveHaResponse {
	return VirtualizationVmwareProactiveHaResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsVirtualizationVmwareProactiveHaResponse is a convenience function that returns MoTagSummary wrapped in VirtualizationVmwareProactiveHaResponse
func MoTagSummaryAsVirtualizationVmwareProactiveHaResponse(v *MoTagSummary) VirtualizationVmwareProactiveHaResponse {
	return VirtualizationVmwareProactiveHaResponse{
		MoTagSummary: v,
	}
}

// VirtualizationVmwareProactiveHaListAsVirtualizationVmwareProactiveHaResponse is a convenience function that returns VirtualizationVmwareProactiveHaList wrapped in VirtualizationVmwareProactiveHaResponse
func VirtualizationVmwareProactiveHaListAsVirtualizationVmwareProactiveHaResponse(v *VirtualizationVmwareProactiveHaList) VirtualizationVmwareProactiveHaResponse {
	return VirtualizationVmwareProactiveHaResponse{
		VirtualizationVmwareProactiveHaList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VirtualizationVmwareProactiveHaResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal VirtualizationVmwareProactiveHaResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal VirtualizationVmwareProactiveHaResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal VirtualizationVmwareProactiveHaResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'virtualization.VmwareProactiveHa.List'
	if jsonDict["ObjectType"] == "virtualization.VmwareProactiveHa.List" {
		// try to unmarshal JSON data into VirtualizationVmwareProactiveHaList
		err = json.Unmarshal(data, &dst.VirtualizationVmwareProactiveHaList)
		if err == nil {
			return nil // data stored in dst.VirtualizationVmwareProactiveHaList, return on the first match
		} else {
			dst.VirtualizationVmwareProactiveHaList = nil
			return fmt.Errorf("failed to unmarshal VirtualizationVmwareProactiveHaResponse as VirtualizationVmwareProactiveHaList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VirtualizationVmwareProactiveHaResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.VirtualizationVmwareProactiveHaList != nil {
		return json.Marshal(&src.VirtualizationVmwareProactiveHaList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VirtualizationVmwareProactiveHaResponse) GetActualInstance() interface{} {
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

	if obj.VirtualizationVmwareProactiveHaList != nil {
		return obj.VirtualizationVmwareProactiveHaList
	}

	// all schemas are nil
	return nil
}

type NullableVirtualizationVmwareProactiveHaResponse struct {
	value *VirtualizationVmwareProactiveHaResponse
	isSet bool
}

func (v NullableVirtualizationVmwareProactiveHaResponse) Get() *VirtualizationVmwareProactiveHaResponse {
	return v.value
}

func (v *NullableVirtualizationVmwareProactiveHaResponse) Set(val *VirtualizationVmwareProactiveHaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareProactiveHaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareProactiveHaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareProactiveHaResponse(val *VirtualizationVmwareProactiveHaResponse) *NullableVirtualizationVmwareProactiveHaResponse {
	return &NullableVirtualizationVmwareProactiveHaResponse{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareProactiveHaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareProactiveHaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}