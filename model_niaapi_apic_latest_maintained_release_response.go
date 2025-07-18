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

// NiaapiApicLatestMaintainedReleaseResponse - The response body of a HTTP GET request for the 'niaapi.ApicLatestMaintainedRelease' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'niaapi.ApicLatestMaintainedRelease' resources.
type NiaapiApicLatestMaintainedReleaseResponse struct {
	MoAggregateTransform                  *MoAggregateTransform
	MoDocumentCount                       *MoDocumentCount
	MoTagSummary                          *MoTagSummary
	NiaapiApicLatestMaintainedReleaseList *NiaapiApicLatestMaintainedReleaseList
}

// MoAggregateTransformAsNiaapiApicLatestMaintainedReleaseResponse is a convenience function that returns MoAggregateTransform wrapped in NiaapiApicLatestMaintainedReleaseResponse
func MoAggregateTransformAsNiaapiApicLatestMaintainedReleaseResponse(v *MoAggregateTransform) NiaapiApicLatestMaintainedReleaseResponse {
	return NiaapiApicLatestMaintainedReleaseResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsNiaapiApicLatestMaintainedReleaseResponse is a convenience function that returns MoDocumentCount wrapped in NiaapiApicLatestMaintainedReleaseResponse
func MoDocumentCountAsNiaapiApicLatestMaintainedReleaseResponse(v *MoDocumentCount) NiaapiApicLatestMaintainedReleaseResponse {
	return NiaapiApicLatestMaintainedReleaseResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsNiaapiApicLatestMaintainedReleaseResponse is a convenience function that returns MoTagSummary wrapped in NiaapiApicLatestMaintainedReleaseResponse
func MoTagSummaryAsNiaapiApicLatestMaintainedReleaseResponse(v *MoTagSummary) NiaapiApicLatestMaintainedReleaseResponse {
	return NiaapiApicLatestMaintainedReleaseResponse{
		MoTagSummary: v,
	}
}

// NiaapiApicLatestMaintainedReleaseListAsNiaapiApicLatestMaintainedReleaseResponse is a convenience function that returns NiaapiApicLatestMaintainedReleaseList wrapped in NiaapiApicLatestMaintainedReleaseResponse
func NiaapiApicLatestMaintainedReleaseListAsNiaapiApicLatestMaintainedReleaseResponse(v *NiaapiApicLatestMaintainedReleaseList) NiaapiApicLatestMaintainedReleaseResponse {
	return NiaapiApicLatestMaintainedReleaseResponse{
		NiaapiApicLatestMaintainedReleaseList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiaapiApicLatestMaintainedReleaseResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal NiaapiApicLatestMaintainedReleaseResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiaapiApicLatestMaintainedReleaseResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiaapiApicLatestMaintainedReleaseResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niaapi.ApicLatestMaintainedRelease.List'
	if jsonDict["ObjectType"] == "niaapi.ApicLatestMaintainedRelease.List" {
		// try to unmarshal JSON data into NiaapiApicLatestMaintainedReleaseList
		err = json.Unmarshal(data, &dst.NiaapiApicLatestMaintainedReleaseList)
		if err == nil {
			return nil // data stored in dst.NiaapiApicLatestMaintainedReleaseList, return on the first match
		} else {
			dst.NiaapiApicLatestMaintainedReleaseList = nil
			return fmt.Errorf("failed to unmarshal NiaapiApicLatestMaintainedReleaseResponse as NiaapiApicLatestMaintainedReleaseList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiaapiApicLatestMaintainedReleaseResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NiaapiApicLatestMaintainedReleaseList != nil {
		return json.Marshal(&src.NiaapiApicLatestMaintainedReleaseList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiaapiApicLatestMaintainedReleaseResponse) GetActualInstance() interface{} {
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

	if obj.NiaapiApicLatestMaintainedReleaseList != nil {
		return obj.NiaapiApicLatestMaintainedReleaseList
	}

	// all schemas are nil
	return nil
}

type NullableNiaapiApicLatestMaintainedReleaseResponse struct {
	value *NiaapiApicLatestMaintainedReleaseResponse
	isSet bool
}

func (v NullableNiaapiApicLatestMaintainedReleaseResponse) Get() *NiaapiApicLatestMaintainedReleaseResponse {
	return v.value
}

func (v *NullableNiaapiApicLatestMaintainedReleaseResponse) Set(val *NiaapiApicLatestMaintainedReleaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiApicLatestMaintainedReleaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiApicLatestMaintainedReleaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiApicLatestMaintainedReleaseResponse(val *NiaapiApicLatestMaintainedReleaseResponse) *NullableNiaapiApicLatestMaintainedReleaseResponse {
	return &NullableNiaapiApicLatestMaintainedReleaseResponse{value: val, isSet: true}
}

func (v NullableNiaapiApicLatestMaintainedReleaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiApicLatestMaintainedReleaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
