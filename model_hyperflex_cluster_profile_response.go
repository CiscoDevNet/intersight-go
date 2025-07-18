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

// HyperflexClusterProfileResponse - The response body of a HTTP GET request for the 'hyperflex.ClusterProfile' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'hyperflex.ClusterProfile' resources.
type HyperflexClusterProfileResponse struct {
	HyperflexClusterProfileList *HyperflexClusterProfileList
	MoAggregateTransform        *MoAggregateTransform
	MoDocumentCount             *MoDocumentCount
	MoTagSummary                *MoTagSummary
}

// HyperflexClusterProfileListAsHyperflexClusterProfileResponse is a convenience function that returns HyperflexClusterProfileList wrapped in HyperflexClusterProfileResponse
func HyperflexClusterProfileListAsHyperflexClusterProfileResponse(v *HyperflexClusterProfileList) HyperflexClusterProfileResponse {
	return HyperflexClusterProfileResponse{
		HyperflexClusterProfileList: v,
	}
}

// MoAggregateTransformAsHyperflexClusterProfileResponse is a convenience function that returns MoAggregateTransform wrapped in HyperflexClusterProfileResponse
func MoAggregateTransformAsHyperflexClusterProfileResponse(v *MoAggregateTransform) HyperflexClusterProfileResponse {
	return HyperflexClusterProfileResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsHyperflexClusterProfileResponse is a convenience function that returns MoDocumentCount wrapped in HyperflexClusterProfileResponse
func MoDocumentCountAsHyperflexClusterProfileResponse(v *MoDocumentCount) HyperflexClusterProfileResponse {
	return HyperflexClusterProfileResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsHyperflexClusterProfileResponse is a convenience function that returns MoTagSummary wrapped in HyperflexClusterProfileResponse
func MoTagSummaryAsHyperflexClusterProfileResponse(v *MoTagSummary) HyperflexClusterProfileResponse {
	return HyperflexClusterProfileResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexClusterProfileResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'hyperflex.ClusterProfile.List'
	if jsonDict["ObjectType"] == "hyperflex.ClusterProfile.List" {
		// try to unmarshal JSON data into HyperflexClusterProfileList
		err = json.Unmarshal(data, &dst.HyperflexClusterProfileList)
		if err == nil {
			return nil // data stored in dst.HyperflexClusterProfileList, return on the first match
		} else {
			dst.HyperflexClusterProfileList = nil
			return fmt.Errorf("failed to unmarshal HyperflexClusterProfileResponse as HyperflexClusterProfileList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexClusterProfileResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexClusterProfileResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexClusterProfileResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexClusterProfileResponse) MarshalJSON() ([]byte, error) {
	if src.HyperflexClusterProfileList != nil {
		return json.Marshal(&src.HyperflexClusterProfileList)
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
func (obj *HyperflexClusterProfileResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HyperflexClusterProfileList != nil {
		return obj.HyperflexClusterProfileList
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

type NullableHyperflexClusterProfileResponse struct {
	value *HyperflexClusterProfileResponse
	isSet bool
}

func (v NullableHyperflexClusterProfileResponse) Get() *HyperflexClusterProfileResponse {
	return v.value
}

func (v *NullableHyperflexClusterProfileResponse) Set(val *HyperflexClusterProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterProfileResponse(val *HyperflexClusterProfileResponse) *NullableHyperflexClusterProfileResponse {
	return &NullableHyperflexClusterProfileResponse{value: val, isSet: true}
}

func (v NullableHyperflexClusterProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
