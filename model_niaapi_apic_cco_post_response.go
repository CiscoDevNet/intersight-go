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

// NiaapiApicCcoPostResponse - The response body of a HTTP GET request for the 'niaapi.ApicCcoPost' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'niaapi.ApicCcoPost' resources.
type NiaapiApicCcoPostResponse struct {
	MoAggregateTransform  *MoAggregateTransform
	MoDocumentCount       *MoDocumentCount
	MoTagSummary          *MoTagSummary
	NiaapiApicCcoPostList *NiaapiApicCcoPostList
}

// MoAggregateTransformAsNiaapiApicCcoPostResponse is a convenience function that returns MoAggregateTransform wrapped in NiaapiApicCcoPostResponse
func MoAggregateTransformAsNiaapiApicCcoPostResponse(v *MoAggregateTransform) NiaapiApicCcoPostResponse {
	return NiaapiApicCcoPostResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsNiaapiApicCcoPostResponse is a convenience function that returns MoDocumentCount wrapped in NiaapiApicCcoPostResponse
func MoDocumentCountAsNiaapiApicCcoPostResponse(v *MoDocumentCount) NiaapiApicCcoPostResponse {
	return NiaapiApicCcoPostResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsNiaapiApicCcoPostResponse is a convenience function that returns MoTagSummary wrapped in NiaapiApicCcoPostResponse
func MoTagSummaryAsNiaapiApicCcoPostResponse(v *MoTagSummary) NiaapiApicCcoPostResponse {
	return NiaapiApicCcoPostResponse{
		MoTagSummary: v,
	}
}

// NiaapiApicCcoPostListAsNiaapiApicCcoPostResponse is a convenience function that returns NiaapiApicCcoPostList wrapped in NiaapiApicCcoPostResponse
func NiaapiApicCcoPostListAsNiaapiApicCcoPostResponse(v *NiaapiApicCcoPostList) NiaapiApicCcoPostResponse {
	return NiaapiApicCcoPostResponse{
		NiaapiApicCcoPostList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiaapiApicCcoPostResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal NiaapiApicCcoPostResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiaapiApicCcoPostResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiaapiApicCcoPostResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niaapi.ApicCcoPost.List'
	if jsonDict["ObjectType"] == "niaapi.ApicCcoPost.List" {
		// try to unmarshal JSON data into NiaapiApicCcoPostList
		err = json.Unmarshal(data, &dst.NiaapiApicCcoPostList)
		if err == nil {
			return nil // data stored in dst.NiaapiApicCcoPostList, return on the first match
		} else {
			dst.NiaapiApicCcoPostList = nil
			return fmt.Errorf("failed to unmarshal NiaapiApicCcoPostResponse as NiaapiApicCcoPostList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiaapiApicCcoPostResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NiaapiApicCcoPostList != nil {
		return json.Marshal(&src.NiaapiApicCcoPostList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiaapiApicCcoPostResponse) GetActualInstance() interface{} {
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

	if obj.NiaapiApicCcoPostList != nil {
		return obj.NiaapiApicCcoPostList
	}

	// all schemas are nil
	return nil
}

type NullableNiaapiApicCcoPostResponse struct {
	value *NiaapiApicCcoPostResponse
	isSet bool
}

func (v NullableNiaapiApicCcoPostResponse) Get() *NiaapiApicCcoPostResponse {
	return v.value
}

func (v *NullableNiaapiApicCcoPostResponse) Set(val *NiaapiApicCcoPostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiApicCcoPostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiApicCcoPostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiApicCcoPostResponse(val *NiaapiApicCcoPostResponse) *NullableNiaapiApicCcoPostResponse {
	return &NullableNiaapiApicCcoPostResponse{value: val, isSet: true}
}

func (v NullableNiaapiApicCcoPostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiApicCcoPostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
