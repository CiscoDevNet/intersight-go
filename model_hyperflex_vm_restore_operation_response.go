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

// HyperflexVmRestoreOperationResponse - The response body of a HTTP GET request for the 'hyperflex.VmRestoreOperation' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'hyperflex.VmRestoreOperation' resources.
type HyperflexVmRestoreOperationResponse struct {
	HyperflexVmRestoreOperationList *HyperflexVmRestoreOperationList
	MoAggregateTransform            *MoAggregateTransform
	MoDocumentCount                 *MoDocumentCount
	MoTagSummary                    *MoTagSummary
}

// HyperflexVmRestoreOperationListAsHyperflexVmRestoreOperationResponse is a convenience function that returns HyperflexVmRestoreOperationList wrapped in HyperflexVmRestoreOperationResponse
func HyperflexVmRestoreOperationListAsHyperflexVmRestoreOperationResponse(v *HyperflexVmRestoreOperationList) HyperflexVmRestoreOperationResponse {
	return HyperflexVmRestoreOperationResponse{
		HyperflexVmRestoreOperationList: v,
	}
}

// MoAggregateTransformAsHyperflexVmRestoreOperationResponse is a convenience function that returns MoAggregateTransform wrapped in HyperflexVmRestoreOperationResponse
func MoAggregateTransformAsHyperflexVmRestoreOperationResponse(v *MoAggregateTransform) HyperflexVmRestoreOperationResponse {
	return HyperflexVmRestoreOperationResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsHyperflexVmRestoreOperationResponse is a convenience function that returns MoDocumentCount wrapped in HyperflexVmRestoreOperationResponse
func MoDocumentCountAsHyperflexVmRestoreOperationResponse(v *MoDocumentCount) HyperflexVmRestoreOperationResponse {
	return HyperflexVmRestoreOperationResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsHyperflexVmRestoreOperationResponse is a convenience function that returns MoTagSummary wrapped in HyperflexVmRestoreOperationResponse
func MoTagSummaryAsHyperflexVmRestoreOperationResponse(v *MoTagSummary) HyperflexVmRestoreOperationResponse {
	return HyperflexVmRestoreOperationResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexVmRestoreOperationResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'hyperflex.VmRestoreOperation.List'
	if jsonDict["ObjectType"] == "hyperflex.VmRestoreOperation.List" {
		// try to unmarshal JSON data into HyperflexVmRestoreOperationList
		err = json.Unmarshal(data, &dst.HyperflexVmRestoreOperationList)
		if err == nil {
			return nil // data stored in dst.HyperflexVmRestoreOperationList, return on the first match
		} else {
			dst.HyperflexVmRestoreOperationList = nil
			return fmt.Errorf("failed to unmarshal HyperflexVmRestoreOperationResponse as HyperflexVmRestoreOperationList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexVmRestoreOperationResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexVmRestoreOperationResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexVmRestoreOperationResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexVmRestoreOperationResponse) MarshalJSON() ([]byte, error) {
	if src.HyperflexVmRestoreOperationList != nil {
		return json.Marshal(&src.HyperflexVmRestoreOperationList)
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
func (obj *HyperflexVmRestoreOperationResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HyperflexVmRestoreOperationList != nil {
		return obj.HyperflexVmRestoreOperationList
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

type NullableHyperflexVmRestoreOperationResponse struct {
	value *HyperflexVmRestoreOperationResponse
	isSet bool
}

func (v NullableHyperflexVmRestoreOperationResponse) Get() *HyperflexVmRestoreOperationResponse {
	return v.value
}

func (v *NullableHyperflexVmRestoreOperationResponse) Set(val *HyperflexVmRestoreOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVmRestoreOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVmRestoreOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVmRestoreOperationResponse(val *HyperflexVmRestoreOperationResponse) *NullableHyperflexVmRestoreOperationResponse {
	return &NullableHyperflexVmRestoreOperationResponse{value: val, isSet: true}
}

func (v NullableHyperflexVmRestoreOperationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVmRestoreOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
