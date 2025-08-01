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

// NiatelemetryApicFanDetailsResponse - The response body of a HTTP GET request for the 'niatelemetry.ApicFanDetails' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'niatelemetry.ApicFanDetails' resources.
type NiatelemetryApicFanDetailsResponse struct {
	MoAggregateTransform           *MoAggregateTransform
	MoDocumentCount                *MoDocumentCount
	MoTagSummary                   *MoTagSummary
	NiatelemetryApicFanDetailsList *NiatelemetryApicFanDetailsList
}

// MoAggregateTransformAsNiatelemetryApicFanDetailsResponse is a convenience function that returns MoAggregateTransform wrapped in NiatelemetryApicFanDetailsResponse
func MoAggregateTransformAsNiatelemetryApicFanDetailsResponse(v *MoAggregateTransform) NiatelemetryApicFanDetailsResponse {
	return NiatelemetryApicFanDetailsResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsNiatelemetryApicFanDetailsResponse is a convenience function that returns MoDocumentCount wrapped in NiatelemetryApicFanDetailsResponse
func MoDocumentCountAsNiatelemetryApicFanDetailsResponse(v *MoDocumentCount) NiatelemetryApicFanDetailsResponse {
	return NiatelemetryApicFanDetailsResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsNiatelemetryApicFanDetailsResponse is a convenience function that returns MoTagSummary wrapped in NiatelemetryApicFanDetailsResponse
func MoTagSummaryAsNiatelemetryApicFanDetailsResponse(v *MoTagSummary) NiatelemetryApicFanDetailsResponse {
	return NiatelemetryApicFanDetailsResponse{
		MoTagSummary: v,
	}
}

// NiatelemetryApicFanDetailsListAsNiatelemetryApicFanDetailsResponse is a convenience function that returns NiatelemetryApicFanDetailsList wrapped in NiatelemetryApicFanDetailsResponse
func NiatelemetryApicFanDetailsListAsNiatelemetryApicFanDetailsResponse(v *NiatelemetryApicFanDetailsList) NiatelemetryApicFanDetailsResponse {
	return NiatelemetryApicFanDetailsResponse{
		NiatelemetryApicFanDetailsList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiatelemetryApicFanDetailsResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal NiatelemetryApicFanDetailsResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryApicFanDetailsResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryApicFanDetailsResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niatelemetry.ApicFanDetails.List'
	if jsonDict["ObjectType"] == "niatelemetry.ApicFanDetails.List" {
		// try to unmarshal JSON data into NiatelemetryApicFanDetailsList
		err = json.Unmarshal(data, &dst.NiatelemetryApicFanDetailsList)
		if err == nil {
			return nil // data stored in dst.NiatelemetryApicFanDetailsList, return on the first match
		} else {
			dst.NiatelemetryApicFanDetailsList = nil
			return fmt.Errorf("failed to unmarshal NiatelemetryApicFanDetailsResponse as NiatelemetryApicFanDetailsList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiatelemetryApicFanDetailsResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NiatelemetryApicFanDetailsList != nil {
		return json.Marshal(&src.NiatelemetryApicFanDetailsList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiatelemetryApicFanDetailsResponse) GetActualInstance() interface{} {
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

	if obj.NiatelemetryApicFanDetailsList != nil {
		return obj.NiatelemetryApicFanDetailsList
	}

	// all schemas are nil
	return nil
}

type NullableNiatelemetryApicFanDetailsResponse struct {
	value *NiatelemetryApicFanDetailsResponse
	isSet bool
}

func (v NullableNiatelemetryApicFanDetailsResponse) Get() *NiatelemetryApicFanDetailsResponse {
	return v.value
}

func (v *NullableNiatelemetryApicFanDetailsResponse) Set(val *NiatelemetryApicFanDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicFanDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicFanDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicFanDetailsResponse(val *NiatelemetryApicFanDetailsResponse) *NullableNiatelemetryApicFanDetailsResponse {
	return &NullableNiatelemetryApicFanDetailsResponse{value: val, isSet: true}
}

func (v NullableNiatelemetryApicFanDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicFanDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
