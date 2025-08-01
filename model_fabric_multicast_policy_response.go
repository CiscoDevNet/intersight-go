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

// FabricMulticastPolicyResponse - The response body of a HTTP GET request for the 'fabric.MulticastPolicy' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'fabric.MulticastPolicy' resources.
type FabricMulticastPolicyResponse struct {
	FabricMulticastPolicyList *FabricMulticastPolicyList
	MoAggregateTransform      *MoAggregateTransform
	MoDocumentCount           *MoDocumentCount
	MoTagSummary              *MoTagSummary
}

// FabricMulticastPolicyListAsFabricMulticastPolicyResponse is a convenience function that returns FabricMulticastPolicyList wrapped in FabricMulticastPolicyResponse
func FabricMulticastPolicyListAsFabricMulticastPolicyResponse(v *FabricMulticastPolicyList) FabricMulticastPolicyResponse {
	return FabricMulticastPolicyResponse{
		FabricMulticastPolicyList: v,
	}
}

// MoAggregateTransformAsFabricMulticastPolicyResponse is a convenience function that returns MoAggregateTransform wrapped in FabricMulticastPolicyResponse
func MoAggregateTransformAsFabricMulticastPolicyResponse(v *MoAggregateTransform) FabricMulticastPolicyResponse {
	return FabricMulticastPolicyResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsFabricMulticastPolicyResponse is a convenience function that returns MoDocumentCount wrapped in FabricMulticastPolicyResponse
func MoDocumentCountAsFabricMulticastPolicyResponse(v *MoDocumentCount) FabricMulticastPolicyResponse {
	return FabricMulticastPolicyResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsFabricMulticastPolicyResponse is a convenience function that returns MoTagSummary wrapped in FabricMulticastPolicyResponse
func MoTagSummaryAsFabricMulticastPolicyResponse(v *MoTagSummary) FabricMulticastPolicyResponse {
	return FabricMulticastPolicyResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FabricMulticastPolicyResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'fabric.MulticastPolicy.List'
	if jsonDict["ObjectType"] == "fabric.MulticastPolicy.List" {
		// try to unmarshal JSON data into FabricMulticastPolicyList
		err = json.Unmarshal(data, &dst.FabricMulticastPolicyList)
		if err == nil {
			return nil // data stored in dst.FabricMulticastPolicyList, return on the first match
		} else {
			dst.FabricMulticastPolicyList = nil
			return fmt.Errorf("failed to unmarshal FabricMulticastPolicyResponse as FabricMulticastPolicyList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FabricMulticastPolicyResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FabricMulticastPolicyResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FabricMulticastPolicyResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FabricMulticastPolicyResponse) MarshalJSON() ([]byte, error) {
	if src.FabricMulticastPolicyList != nil {
		return json.Marshal(&src.FabricMulticastPolicyList)
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
func (obj *FabricMulticastPolicyResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FabricMulticastPolicyList != nil {
		return obj.FabricMulticastPolicyList
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

type NullableFabricMulticastPolicyResponse struct {
	value *FabricMulticastPolicyResponse
	isSet bool
}

func (v NullableFabricMulticastPolicyResponse) Get() *FabricMulticastPolicyResponse {
	return v.value
}

func (v *NullableFabricMulticastPolicyResponse) Set(val *FabricMulticastPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricMulticastPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricMulticastPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricMulticastPolicyResponse(val *FabricMulticastPolicyResponse) *NullableFabricMulticastPolicyResponse {
	return &NullableFabricMulticastPolicyResponse{value: val, isSet: true}
}

func (v NullableFabricMulticastPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricMulticastPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
