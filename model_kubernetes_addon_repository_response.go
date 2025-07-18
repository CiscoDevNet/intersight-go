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

// KubernetesAddonRepositoryResponse - The response body of a HTTP GET request for the 'kubernetes.AddonRepository' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'kubernetes.AddonRepository' resources.
type KubernetesAddonRepositoryResponse struct {
	KubernetesAddonRepositoryList *KubernetesAddonRepositoryList
	MoAggregateTransform          *MoAggregateTransform
	MoDocumentCount               *MoDocumentCount
	MoTagSummary                  *MoTagSummary
}

// KubernetesAddonRepositoryListAsKubernetesAddonRepositoryResponse is a convenience function that returns KubernetesAddonRepositoryList wrapped in KubernetesAddonRepositoryResponse
func KubernetesAddonRepositoryListAsKubernetesAddonRepositoryResponse(v *KubernetesAddonRepositoryList) KubernetesAddonRepositoryResponse {
	return KubernetesAddonRepositoryResponse{
		KubernetesAddonRepositoryList: v,
	}
}

// MoAggregateTransformAsKubernetesAddonRepositoryResponse is a convenience function that returns MoAggregateTransform wrapped in KubernetesAddonRepositoryResponse
func MoAggregateTransformAsKubernetesAddonRepositoryResponse(v *MoAggregateTransform) KubernetesAddonRepositoryResponse {
	return KubernetesAddonRepositoryResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsKubernetesAddonRepositoryResponse is a convenience function that returns MoDocumentCount wrapped in KubernetesAddonRepositoryResponse
func MoDocumentCountAsKubernetesAddonRepositoryResponse(v *MoDocumentCount) KubernetesAddonRepositoryResponse {
	return KubernetesAddonRepositoryResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsKubernetesAddonRepositoryResponse is a convenience function that returns MoTagSummary wrapped in KubernetesAddonRepositoryResponse
func MoTagSummaryAsKubernetesAddonRepositoryResponse(v *MoTagSummary) KubernetesAddonRepositoryResponse {
	return KubernetesAddonRepositoryResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KubernetesAddonRepositoryResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'kubernetes.AddonRepository.List'
	if jsonDict["ObjectType"] == "kubernetes.AddonRepository.List" {
		// try to unmarshal JSON data into KubernetesAddonRepositoryList
		err = json.Unmarshal(data, &dst.KubernetesAddonRepositoryList)
		if err == nil {
			return nil // data stored in dst.KubernetesAddonRepositoryList, return on the first match
		} else {
			dst.KubernetesAddonRepositoryList = nil
			return fmt.Errorf("failed to unmarshal KubernetesAddonRepositoryResponse as KubernetesAddonRepositoryList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal KubernetesAddonRepositoryResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal KubernetesAddonRepositoryResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal KubernetesAddonRepositoryResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KubernetesAddonRepositoryResponse) MarshalJSON() ([]byte, error) {
	if src.KubernetesAddonRepositoryList != nil {
		return json.Marshal(&src.KubernetesAddonRepositoryList)
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
func (obj *KubernetesAddonRepositoryResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.KubernetesAddonRepositoryList != nil {
		return obj.KubernetesAddonRepositoryList
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

type NullableKubernetesAddonRepositoryResponse struct {
	value *KubernetesAddonRepositoryResponse
	isSet bool
}

func (v NullableKubernetesAddonRepositoryResponse) Get() *KubernetesAddonRepositoryResponse {
	return v.value
}

func (v *NullableKubernetesAddonRepositoryResponse) Set(val *KubernetesAddonRepositoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAddonRepositoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAddonRepositoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAddonRepositoryResponse(val *KubernetesAddonRepositoryResponse) *NullableKubernetesAddonRepositoryResponse {
	return &NullableKubernetesAddonRepositoryResponse{value: val, isSet: true}
}

func (v NullableKubernetesAddonRepositoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAddonRepositoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
