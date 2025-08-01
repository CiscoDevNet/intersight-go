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

// NiatelemetryApicSnmpCommunityDetailsResponse - The response body of a HTTP GET request for the 'niatelemetry.ApicSnmpCommunityDetails' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'niatelemetry.ApicSnmpCommunityDetails' resources.
type NiatelemetryApicSnmpCommunityDetailsResponse struct {
	MoAggregateTransform                     *MoAggregateTransform
	MoDocumentCount                          *MoDocumentCount
	MoTagSummary                             *MoTagSummary
	NiatelemetryApicSnmpCommunityDetailsList *NiatelemetryApicSnmpCommunityDetailsList
}

// MoAggregateTransformAsNiatelemetryApicSnmpCommunityDetailsResponse is a convenience function that returns MoAggregateTransform wrapped in NiatelemetryApicSnmpCommunityDetailsResponse
func MoAggregateTransformAsNiatelemetryApicSnmpCommunityDetailsResponse(v *MoAggregateTransform) NiatelemetryApicSnmpCommunityDetailsResponse {
	return NiatelemetryApicSnmpCommunityDetailsResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsNiatelemetryApicSnmpCommunityDetailsResponse is a convenience function that returns MoDocumentCount wrapped in NiatelemetryApicSnmpCommunityDetailsResponse
func MoDocumentCountAsNiatelemetryApicSnmpCommunityDetailsResponse(v *MoDocumentCount) NiatelemetryApicSnmpCommunityDetailsResponse {
	return NiatelemetryApicSnmpCommunityDetailsResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsNiatelemetryApicSnmpCommunityDetailsResponse is a convenience function that returns MoTagSummary wrapped in NiatelemetryApicSnmpCommunityDetailsResponse
func MoTagSummaryAsNiatelemetryApicSnmpCommunityDetailsResponse(v *MoTagSummary) NiatelemetryApicSnmpCommunityDetailsResponse {
	return NiatelemetryApicSnmpCommunityDetailsResponse{
		MoTagSummary: v,
	}
}

// NiatelemetryApicSnmpCommunityDetailsListAsNiatelemetryApicSnmpCommunityDetailsResponse is a convenience function that returns NiatelemetryApicSnmpCommunityDetailsList wrapped in NiatelemetryApicSnmpCommunityDetailsResponse
func NiatelemetryApicSnmpCommunityDetailsListAsNiatelemetryApicSnmpCommunityDetailsResponse(v *NiatelemetryApicSnmpCommunityDetailsList) NiatelemetryApicSnmpCommunityDetailsResponse {
	return NiatelemetryApicSnmpCommunityDetailsResponse{
		NiatelemetryApicSnmpCommunityDetailsList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiatelemetryApicSnmpCommunityDetailsResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal NiatelemetryApicSnmpCommunityDetailsResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryApicSnmpCommunityDetailsResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryApicSnmpCommunityDetailsResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niatelemetry.ApicSnmpCommunityDetails.List'
	if jsonDict["ObjectType"] == "niatelemetry.ApicSnmpCommunityDetails.List" {
		// try to unmarshal JSON data into NiatelemetryApicSnmpCommunityDetailsList
		err = json.Unmarshal(data, &dst.NiatelemetryApicSnmpCommunityDetailsList)
		if err == nil {
			return nil // data stored in dst.NiatelemetryApicSnmpCommunityDetailsList, return on the first match
		} else {
			dst.NiatelemetryApicSnmpCommunityDetailsList = nil
			return fmt.Errorf("failed to unmarshal NiatelemetryApicSnmpCommunityDetailsResponse as NiatelemetryApicSnmpCommunityDetailsList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiatelemetryApicSnmpCommunityDetailsResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NiatelemetryApicSnmpCommunityDetailsList != nil {
		return json.Marshal(&src.NiatelemetryApicSnmpCommunityDetailsList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiatelemetryApicSnmpCommunityDetailsResponse) GetActualInstance() interface{} {
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

	if obj.NiatelemetryApicSnmpCommunityDetailsList != nil {
		return obj.NiatelemetryApicSnmpCommunityDetailsList
	}

	// all schemas are nil
	return nil
}

type NullableNiatelemetryApicSnmpCommunityDetailsResponse struct {
	value *NiatelemetryApicSnmpCommunityDetailsResponse
	isSet bool
}

func (v NullableNiatelemetryApicSnmpCommunityDetailsResponse) Get() *NiatelemetryApicSnmpCommunityDetailsResponse {
	return v.value
}

func (v *NullableNiatelemetryApicSnmpCommunityDetailsResponse) Set(val *NiatelemetryApicSnmpCommunityDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicSnmpCommunityDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicSnmpCommunityDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicSnmpCommunityDetailsResponse(val *NiatelemetryApicSnmpCommunityDetailsResponse) *NullableNiatelemetryApicSnmpCommunityDetailsResponse {
	return &NullableNiatelemetryApicSnmpCommunityDetailsResponse{value: val, isSet: true}
}

func (v NullableNiatelemetryApicSnmpCommunityDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicSnmpCommunityDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
