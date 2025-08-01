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

// SoftwarerepositoryCategoryMapperResponse - The response body of a HTTP GET request for the 'softwarerepository.CategoryMapper' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'softwarerepository.CategoryMapper' resources.
type SoftwarerepositoryCategoryMapperResponse struct {
	MoAggregateTransform                 *MoAggregateTransform
	MoDocumentCount                      *MoDocumentCount
	MoTagSummary                         *MoTagSummary
	SoftwarerepositoryCategoryMapperList *SoftwarerepositoryCategoryMapperList
}

// MoAggregateTransformAsSoftwarerepositoryCategoryMapperResponse is a convenience function that returns MoAggregateTransform wrapped in SoftwarerepositoryCategoryMapperResponse
func MoAggregateTransformAsSoftwarerepositoryCategoryMapperResponse(v *MoAggregateTransform) SoftwarerepositoryCategoryMapperResponse {
	return SoftwarerepositoryCategoryMapperResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsSoftwarerepositoryCategoryMapperResponse is a convenience function that returns MoDocumentCount wrapped in SoftwarerepositoryCategoryMapperResponse
func MoDocumentCountAsSoftwarerepositoryCategoryMapperResponse(v *MoDocumentCount) SoftwarerepositoryCategoryMapperResponse {
	return SoftwarerepositoryCategoryMapperResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsSoftwarerepositoryCategoryMapperResponse is a convenience function that returns MoTagSummary wrapped in SoftwarerepositoryCategoryMapperResponse
func MoTagSummaryAsSoftwarerepositoryCategoryMapperResponse(v *MoTagSummary) SoftwarerepositoryCategoryMapperResponse {
	return SoftwarerepositoryCategoryMapperResponse{
		MoTagSummary: v,
	}
}

// SoftwarerepositoryCategoryMapperListAsSoftwarerepositoryCategoryMapperResponse is a convenience function that returns SoftwarerepositoryCategoryMapperList wrapped in SoftwarerepositoryCategoryMapperResponse
func SoftwarerepositoryCategoryMapperListAsSoftwarerepositoryCategoryMapperResponse(v *SoftwarerepositoryCategoryMapperList) SoftwarerepositoryCategoryMapperResponse {
	return SoftwarerepositoryCategoryMapperResponse{
		SoftwarerepositoryCategoryMapperList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SoftwarerepositoryCategoryMapperResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal SoftwarerepositoryCategoryMapperResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SoftwarerepositoryCategoryMapperResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SoftwarerepositoryCategoryMapperResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'softwarerepository.CategoryMapper.List'
	if jsonDict["ObjectType"] == "softwarerepository.CategoryMapper.List" {
		// try to unmarshal JSON data into SoftwarerepositoryCategoryMapperList
		err = json.Unmarshal(data, &dst.SoftwarerepositoryCategoryMapperList)
		if err == nil {
			return nil // data stored in dst.SoftwarerepositoryCategoryMapperList, return on the first match
		} else {
			dst.SoftwarerepositoryCategoryMapperList = nil
			return fmt.Errorf("failed to unmarshal SoftwarerepositoryCategoryMapperResponse as SoftwarerepositoryCategoryMapperList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SoftwarerepositoryCategoryMapperResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.SoftwarerepositoryCategoryMapperList != nil {
		return json.Marshal(&src.SoftwarerepositoryCategoryMapperList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SoftwarerepositoryCategoryMapperResponse) GetActualInstance() interface{} {
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

	if obj.SoftwarerepositoryCategoryMapperList != nil {
		return obj.SoftwarerepositoryCategoryMapperList
	}

	// all schemas are nil
	return nil
}

type NullableSoftwarerepositoryCategoryMapperResponse struct {
	value *SoftwarerepositoryCategoryMapperResponse
	isSet bool
}

func (v NullableSoftwarerepositoryCategoryMapperResponse) Get() *SoftwarerepositoryCategoryMapperResponse {
	return v.value
}

func (v *NullableSoftwarerepositoryCategoryMapperResponse) Set(val *SoftwarerepositoryCategoryMapperResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCategoryMapperResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCategoryMapperResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCategoryMapperResponse(val *SoftwarerepositoryCategoryMapperResponse) *NullableSoftwarerepositoryCategoryMapperResponse {
	return &NullableSoftwarerepositoryCategoryMapperResponse{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCategoryMapperResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCategoryMapperResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
