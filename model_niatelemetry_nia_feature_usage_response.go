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

// NiatelemetryNiaFeatureUsageResponse - The response body of a HTTP GET request for the 'niatelemetry.NiaFeatureUsage' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'niatelemetry.NiaFeatureUsage' resources.
type NiatelemetryNiaFeatureUsageResponse struct {
	MoAggregateTransform            *MoAggregateTransform
	MoDocumentCount                 *MoDocumentCount
	MoTagSummary                    *MoTagSummary
	NiatelemetryNiaFeatureUsageList *NiatelemetryNiaFeatureUsageList
}

// MoAggregateTransformAsNiatelemetryNiaFeatureUsageResponse is a convenience function that returns MoAggregateTransform wrapped in NiatelemetryNiaFeatureUsageResponse
func MoAggregateTransformAsNiatelemetryNiaFeatureUsageResponse(v *MoAggregateTransform) NiatelemetryNiaFeatureUsageResponse {
	return NiatelemetryNiaFeatureUsageResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsNiatelemetryNiaFeatureUsageResponse is a convenience function that returns MoDocumentCount wrapped in NiatelemetryNiaFeatureUsageResponse
func MoDocumentCountAsNiatelemetryNiaFeatureUsageResponse(v *MoDocumentCount) NiatelemetryNiaFeatureUsageResponse {
	return NiatelemetryNiaFeatureUsageResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsNiatelemetryNiaFeatureUsageResponse is a convenience function that returns MoTagSummary wrapped in NiatelemetryNiaFeatureUsageResponse
func MoTagSummaryAsNiatelemetryNiaFeatureUsageResponse(v *MoTagSummary) NiatelemetryNiaFeatureUsageResponse {
	return NiatelemetryNiaFeatureUsageResponse{
		MoTagSummary: v,
	}
}

// NiatelemetryNiaFeatureUsageListAsNiatelemetryNiaFeatureUsageResponse is a convenience function that returns NiatelemetryNiaFeatureUsageList wrapped in NiatelemetryNiaFeatureUsageResponse
func NiatelemetryNiaFeatureUsageListAsNiatelemetryNiaFeatureUsageResponse(v *NiatelemetryNiaFeatureUsageList) NiatelemetryNiaFeatureUsageResponse {
	return NiatelemetryNiaFeatureUsageResponse{
		NiatelemetryNiaFeatureUsageList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiatelemetryNiaFeatureUsageResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal NiatelemetryNiaFeatureUsageResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryNiaFeatureUsageResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryNiaFeatureUsageResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niatelemetry.NiaFeatureUsage.List'
	if jsonDict["ObjectType"] == "niatelemetry.NiaFeatureUsage.List" {
		// try to unmarshal JSON data into NiatelemetryNiaFeatureUsageList
		err = json.Unmarshal(data, &dst.NiatelemetryNiaFeatureUsageList)
		if err == nil {
			return nil // data stored in dst.NiatelemetryNiaFeatureUsageList, return on the first match
		} else {
			dst.NiatelemetryNiaFeatureUsageList = nil
			return fmt.Errorf("failed to unmarshal NiatelemetryNiaFeatureUsageResponse as NiatelemetryNiaFeatureUsageList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiatelemetryNiaFeatureUsageResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NiatelemetryNiaFeatureUsageList != nil {
		return json.Marshal(&src.NiatelemetryNiaFeatureUsageList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiatelemetryNiaFeatureUsageResponse) GetActualInstance() interface{} {
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

	if obj.NiatelemetryNiaFeatureUsageList != nil {
		return obj.NiatelemetryNiaFeatureUsageList
	}

	// all schemas are nil
	return nil
}

type NullableNiatelemetryNiaFeatureUsageResponse struct {
	value *NiatelemetryNiaFeatureUsageResponse
	isSet bool
}

func (v NullableNiatelemetryNiaFeatureUsageResponse) Get() *NiatelemetryNiaFeatureUsageResponse {
	return v.value
}

func (v *NullableNiatelemetryNiaFeatureUsageResponse) Set(val *NiatelemetryNiaFeatureUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNiaFeatureUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNiaFeatureUsageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNiaFeatureUsageResponse(val *NiatelemetryNiaFeatureUsageResponse) *NullableNiatelemetryNiaFeatureUsageResponse {
	return &NullableNiatelemetryNiaFeatureUsageResponse{value: val, isSet: true}
}

func (v NullableNiatelemetryNiaFeatureUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNiaFeatureUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
