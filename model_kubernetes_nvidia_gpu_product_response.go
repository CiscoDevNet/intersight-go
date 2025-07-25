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

// KubernetesNvidiaGpuProductResponse - The response body of a HTTP GET request for the 'kubernetes.NvidiaGpuProduct' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'kubernetes.NvidiaGpuProduct' resources.
type KubernetesNvidiaGpuProductResponse struct {
	KubernetesNvidiaGpuProductList *KubernetesNvidiaGpuProductList
	MoAggregateTransform           *MoAggregateTransform
	MoDocumentCount                *MoDocumentCount
	MoTagSummary                   *MoTagSummary
}

// KubernetesNvidiaGpuProductListAsKubernetesNvidiaGpuProductResponse is a convenience function that returns KubernetesNvidiaGpuProductList wrapped in KubernetesNvidiaGpuProductResponse
func KubernetesNvidiaGpuProductListAsKubernetesNvidiaGpuProductResponse(v *KubernetesNvidiaGpuProductList) KubernetesNvidiaGpuProductResponse {
	return KubernetesNvidiaGpuProductResponse{
		KubernetesNvidiaGpuProductList: v,
	}
}

// MoAggregateTransformAsKubernetesNvidiaGpuProductResponse is a convenience function that returns MoAggregateTransform wrapped in KubernetesNvidiaGpuProductResponse
func MoAggregateTransformAsKubernetesNvidiaGpuProductResponse(v *MoAggregateTransform) KubernetesNvidiaGpuProductResponse {
	return KubernetesNvidiaGpuProductResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsKubernetesNvidiaGpuProductResponse is a convenience function that returns MoDocumentCount wrapped in KubernetesNvidiaGpuProductResponse
func MoDocumentCountAsKubernetesNvidiaGpuProductResponse(v *MoDocumentCount) KubernetesNvidiaGpuProductResponse {
	return KubernetesNvidiaGpuProductResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsKubernetesNvidiaGpuProductResponse is a convenience function that returns MoTagSummary wrapped in KubernetesNvidiaGpuProductResponse
func MoTagSummaryAsKubernetesNvidiaGpuProductResponse(v *MoTagSummary) KubernetesNvidiaGpuProductResponse {
	return KubernetesNvidiaGpuProductResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KubernetesNvidiaGpuProductResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'kubernetes.NvidiaGpuProduct.List'
	if jsonDict["ObjectType"] == "kubernetes.NvidiaGpuProduct.List" {
		// try to unmarshal JSON data into KubernetesNvidiaGpuProductList
		err = json.Unmarshal(data, &dst.KubernetesNvidiaGpuProductList)
		if err == nil {
			return nil // data stored in dst.KubernetesNvidiaGpuProductList, return on the first match
		} else {
			dst.KubernetesNvidiaGpuProductList = nil
			return fmt.Errorf("failed to unmarshal KubernetesNvidiaGpuProductResponse as KubernetesNvidiaGpuProductList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal KubernetesNvidiaGpuProductResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal KubernetesNvidiaGpuProductResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal KubernetesNvidiaGpuProductResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KubernetesNvidiaGpuProductResponse) MarshalJSON() ([]byte, error) {
	if src.KubernetesNvidiaGpuProductList != nil {
		return json.Marshal(&src.KubernetesNvidiaGpuProductList)
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
func (obj *KubernetesNvidiaGpuProductResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.KubernetesNvidiaGpuProductList != nil {
		return obj.KubernetesNvidiaGpuProductList
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

type NullableKubernetesNvidiaGpuProductResponse struct {
	value *KubernetesNvidiaGpuProductResponse
	isSet bool
}

func (v NullableKubernetesNvidiaGpuProductResponse) Get() *KubernetesNvidiaGpuProductResponse {
	return v.value
}

func (v *NullableKubernetesNvidiaGpuProductResponse) Set(val *KubernetesNvidiaGpuProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNvidiaGpuProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNvidiaGpuProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNvidiaGpuProductResponse(val *KubernetesNvidiaGpuProductResponse) *NullableKubernetesNvidiaGpuProductResponse {
	return &NullableKubernetesNvidiaGpuProductResponse{value: val, isSet: true}
}

func (v NullableKubernetesNvidiaGpuProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNvidiaGpuProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
