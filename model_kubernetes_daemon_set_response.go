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

// KubernetesDaemonSetResponse - The response body of a HTTP GET request for the 'kubernetes.DaemonSet' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'kubernetes.DaemonSet' resources.
type KubernetesDaemonSetResponse struct {
	KubernetesDaemonSetList *KubernetesDaemonSetList
	MoAggregateTransform    *MoAggregateTransform
	MoDocumentCount         *MoDocumentCount
	MoTagSummary            *MoTagSummary
}

// KubernetesDaemonSetListAsKubernetesDaemonSetResponse is a convenience function that returns KubernetesDaemonSetList wrapped in KubernetesDaemonSetResponse
func KubernetesDaemonSetListAsKubernetesDaemonSetResponse(v *KubernetesDaemonSetList) KubernetesDaemonSetResponse {
	return KubernetesDaemonSetResponse{
		KubernetesDaemonSetList: v,
	}
}

// MoAggregateTransformAsKubernetesDaemonSetResponse is a convenience function that returns MoAggregateTransform wrapped in KubernetesDaemonSetResponse
func MoAggregateTransformAsKubernetesDaemonSetResponse(v *MoAggregateTransform) KubernetesDaemonSetResponse {
	return KubernetesDaemonSetResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsKubernetesDaemonSetResponse is a convenience function that returns MoDocumentCount wrapped in KubernetesDaemonSetResponse
func MoDocumentCountAsKubernetesDaemonSetResponse(v *MoDocumentCount) KubernetesDaemonSetResponse {
	return KubernetesDaemonSetResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsKubernetesDaemonSetResponse is a convenience function that returns MoTagSummary wrapped in KubernetesDaemonSetResponse
func MoTagSummaryAsKubernetesDaemonSetResponse(v *MoTagSummary) KubernetesDaemonSetResponse {
	return KubernetesDaemonSetResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KubernetesDaemonSetResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'kubernetes.DaemonSet.List'
	if jsonDict["ObjectType"] == "kubernetes.DaemonSet.List" {
		// try to unmarshal JSON data into KubernetesDaemonSetList
		err = json.Unmarshal(data, &dst.KubernetesDaemonSetList)
		if err == nil {
			return nil // data stored in dst.KubernetesDaemonSetList, return on the first match
		} else {
			dst.KubernetesDaemonSetList = nil
			return fmt.Errorf("failed to unmarshal KubernetesDaemonSetResponse as KubernetesDaemonSetList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal KubernetesDaemonSetResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal KubernetesDaemonSetResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal KubernetesDaemonSetResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KubernetesDaemonSetResponse) MarshalJSON() ([]byte, error) {
	if src.KubernetesDaemonSetList != nil {
		return json.Marshal(&src.KubernetesDaemonSetList)
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
func (obj *KubernetesDaemonSetResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.KubernetesDaemonSetList != nil {
		return obj.KubernetesDaemonSetList
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

type NullableKubernetesDaemonSetResponse struct {
	value *KubernetesDaemonSetResponse
	isSet bool
}

func (v NullableKubernetesDaemonSetResponse) Get() *KubernetesDaemonSetResponse {
	return v.value
}

func (v *NullableKubernetesDaemonSetResponse) Set(val *KubernetesDaemonSetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesDaemonSetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesDaemonSetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesDaemonSetResponse(val *KubernetesDaemonSetResponse) *NullableKubernetesDaemonSetResponse {
	return &NullableKubernetesDaemonSetResponse{value: val, isSet: true}
}

func (v NullableKubernetesDaemonSetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesDaemonSetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
