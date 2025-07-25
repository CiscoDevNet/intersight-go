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

// ApplianceClusterWorkerNodeResponse - The response body of a HTTP GET request for the 'appliance.ClusterWorkerNode' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'appliance.ClusterWorkerNode' resources.
type ApplianceClusterWorkerNodeResponse struct {
	ApplianceClusterWorkerNodeList *ApplianceClusterWorkerNodeList
	MoAggregateTransform           *MoAggregateTransform
	MoDocumentCount                *MoDocumentCount
	MoTagSummary                   *MoTagSummary
}

// ApplianceClusterWorkerNodeListAsApplianceClusterWorkerNodeResponse is a convenience function that returns ApplianceClusterWorkerNodeList wrapped in ApplianceClusterWorkerNodeResponse
func ApplianceClusterWorkerNodeListAsApplianceClusterWorkerNodeResponse(v *ApplianceClusterWorkerNodeList) ApplianceClusterWorkerNodeResponse {
	return ApplianceClusterWorkerNodeResponse{
		ApplianceClusterWorkerNodeList: v,
	}
}

// MoAggregateTransformAsApplianceClusterWorkerNodeResponse is a convenience function that returns MoAggregateTransform wrapped in ApplianceClusterWorkerNodeResponse
func MoAggregateTransformAsApplianceClusterWorkerNodeResponse(v *MoAggregateTransform) ApplianceClusterWorkerNodeResponse {
	return ApplianceClusterWorkerNodeResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsApplianceClusterWorkerNodeResponse is a convenience function that returns MoDocumentCount wrapped in ApplianceClusterWorkerNodeResponse
func MoDocumentCountAsApplianceClusterWorkerNodeResponse(v *MoDocumentCount) ApplianceClusterWorkerNodeResponse {
	return ApplianceClusterWorkerNodeResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsApplianceClusterWorkerNodeResponse is a convenience function that returns MoTagSummary wrapped in ApplianceClusterWorkerNodeResponse
func MoTagSummaryAsApplianceClusterWorkerNodeResponse(v *MoTagSummary) ApplianceClusterWorkerNodeResponse {
	return ApplianceClusterWorkerNodeResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApplianceClusterWorkerNodeResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'appliance.ClusterWorkerNode.List'
	if jsonDict["ObjectType"] == "appliance.ClusterWorkerNode.List" {
		// try to unmarshal JSON data into ApplianceClusterWorkerNodeList
		err = json.Unmarshal(data, &dst.ApplianceClusterWorkerNodeList)
		if err == nil {
			return nil // data stored in dst.ApplianceClusterWorkerNodeList, return on the first match
		} else {
			dst.ApplianceClusterWorkerNodeList = nil
			return fmt.Errorf("failed to unmarshal ApplianceClusterWorkerNodeResponse as ApplianceClusterWorkerNodeList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ApplianceClusterWorkerNodeResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ApplianceClusterWorkerNodeResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ApplianceClusterWorkerNodeResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApplianceClusterWorkerNodeResponse) MarshalJSON() ([]byte, error) {
	if src.ApplianceClusterWorkerNodeList != nil {
		return json.Marshal(&src.ApplianceClusterWorkerNodeList)
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
func (obj *ApplianceClusterWorkerNodeResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ApplianceClusterWorkerNodeList != nil {
		return obj.ApplianceClusterWorkerNodeList
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

type NullableApplianceClusterWorkerNodeResponse struct {
	value *ApplianceClusterWorkerNodeResponse
	isSet bool
}

func (v NullableApplianceClusterWorkerNodeResponse) Get() *ApplianceClusterWorkerNodeResponse {
	return v.value
}

func (v *NullableApplianceClusterWorkerNodeResponse) Set(val *ApplianceClusterWorkerNodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceClusterWorkerNodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceClusterWorkerNodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceClusterWorkerNodeResponse(val *ApplianceClusterWorkerNodeResponse) *NullableApplianceClusterWorkerNodeResponse {
	return &NullableApplianceClusterWorkerNodeResponse{value: val, isSet: true}
}

func (v NullableApplianceClusterWorkerNodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceClusterWorkerNodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
