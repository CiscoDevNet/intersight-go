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

// HyperflexAppCatalogResponse - The response body of a HTTP GET request for the 'hyperflex.AppCatalog' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'hyperflex.AppCatalog' resources.
type HyperflexAppCatalogResponse struct {
	HyperflexAppCatalogList *HyperflexAppCatalogList
	MoAggregateTransform    *MoAggregateTransform
	MoDocumentCount         *MoDocumentCount
	MoTagSummary            *MoTagSummary
}

// HyperflexAppCatalogListAsHyperflexAppCatalogResponse is a convenience function that returns HyperflexAppCatalogList wrapped in HyperflexAppCatalogResponse
func HyperflexAppCatalogListAsHyperflexAppCatalogResponse(v *HyperflexAppCatalogList) HyperflexAppCatalogResponse {
	return HyperflexAppCatalogResponse{
		HyperflexAppCatalogList: v,
	}
}

// MoAggregateTransformAsHyperflexAppCatalogResponse is a convenience function that returns MoAggregateTransform wrapped in HyperflexAppCatalogResponse
func MoAggregateTransformAsHyperflexAppCatalogResponse(v *MoAggregateTransform) HyperflexAppCatalogResponse {
	return HyperflexAppCatalogResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsHyperflexAppCatalogResponse is a convenience function that returns MoDocumentCount wrapped in HyperflexAppCatalogResponse
func MoDocumentCountAsHyperflexAppCatalogResponse(v *MoDocumentCount) HyperflexAppCatalogResponse {
	return HyperflexAppCatalogResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsHyperflexAppCatalogResponse is a convenience function that returns MoTagSummary wrapped in HyperflexAppCatalogResponse
func MoTagSummaryAsHyperflexAppCatalogResponse(v *MoTagSummary) HyperflexAppCatalogResponse {
	return HyperflexAppCatalogResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexAppCatalogResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'hyperflex.AppCatalog.List'
	if jsonDict["ObjectType"] == "hyperflex.AppCatalog.List" {
		// try to unmarshal JSON data into HyperflexAppCatalogList
		err = json.Unmarshal(data, &dst.HyperflexAppCatalogList)
		if err == nil {
			return nil // data stored in dst.HyperflexAppCatalogList, return on the first match
		} else {
			dst.HyperflexAppCatalogList = nil
			return fmt.Errorf("failed to unmarshal HyperflexAppCatalogResponse as HyperflexAppCatalogList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexAppCatalogResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexAppCatalogResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexAppCatalogResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexAppCatalogResponse) MarshalJSON() ([]byte, error) {
	if src.HyperflexAppCatalogList != nil {
		return json.Marshal(&src.HyperflexAppCatalogList)
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
func (obj *HyperflexAppCatalogResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HyperflexAppCatalogList != nil {
		return obj.HyperflexAppCatalogList
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

type NullableHyperflexAppCatalogResponse struct {
	value *HyperflexAppCatalogResponse
	isSet bool
}

func (v NullableHyperflexAppCatalogResponse) Get() *HyperflexAppCatalogResponse {
	return v.value
}

func (v *NullableHyperflexAppCatalogResponse) Set(val *HyperflexAppCatalogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAppCatalogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAppCatalogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAppCatalogResponse(val *HyperflexAppCatalogResponse) *NullableHyperflexAppCatalogResponse {
	return &NullableHyperflexAppCatalogResponse{value: val, isSet: true}
}

func (v NullableHyperflexAppCatalogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAppCatalogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
