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

// NiatelemetryMsoSiteDetailsResponse - The response body of a HTTP GET request for the 'niatelemetry.MsoSiteDetails' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'niatelemetry.MsoSiteDetails' resources.
type NiatelemetryMsoSiteDetailsResponse struct {
	MoAggregateTransform           *MoAggregateTransform
	MoDocumentCount                *MoDocumentCount
	MoTagSummary                   *MoTagSummary
	NiatelemetryMsoSiteDetailsList *NiatelemetryMsoSiteDetailsList
}

// MoAggregateTransformAsNiatelemetryMsoSiteDetailsResponse is a convenience function that returns MoAggregateTransform wrapped in NiatelemetryMsoSiteDetailsResponse
func MoAggregateTransformAsNiatelemetryMsoSiteDetailsResponse(v *MoAggregateTransform) NiatelemetryMsoSiteDetailsResponse {
	return NiatelemetryMsoSiteDetailsResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsNiatelemetryMsoSiteDetailsResponse is a convenience function that returns MoDocumentCount wrapped in NiatelemetryMsoSiteDetailsResponse
func MoDocumentCountAsNiatelemetryMsoSiteDetailsResponse(v *MoDocumentCount) NiatelemetryMsoSiteDetailsResponse {
	return NiatelemetryMsoSiteDetailsResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsNiatelemetryMsoSiteDetailsResponse is a convenience function that returns MoTagSummary wrapped in NiatelemetryMsoSiteDetailsResponse
func MoTagSummaryAsNiatelemetryMsoSiteDetailsResponse(v *MoTagSummary) NiatelemetryMsoSiteDetailsResponse {
	return NiatelemetryMsoSiteDetailsResponse{
		MoTagSummary: v,
	}
}

// NiatelemetryMsoSiteDetailsListAsNiatelemetryMsoSiteDetailsResponse is a convenience function that returns NiatelemetryMsoSiteDetailsList wrapped in NiatelemetryMsoSiteDetailsResponse
func NiatelemetryMsoSiteDetailsListAsNiatelemetryMsoSiteDetailsResponse(v *NiatelemetryMsoSiteDetailsList) NiatelemetryMsoSiteDetailsResponse {
	return NiatelemetryMsoSiteDetailsResponse{
		NiatelemetryMsoSiteDetailsList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiatelemetryMsoSiteDetailsResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal NiatelemetryMsoSiteDetailsResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryMsoSiteDetailsResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryMsoSiteDetailsResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niatelemetry.MsoSiteDetails.List'
	if jsonDict["ObjectType"] == "niatelemetry.MsoSiteDetails.List" {
		// try to unmarshal JSON data into NiatelemetryMsoSiteDetailsList
		err = json.Unmarshal(data, &dst.NiatelemetryMsoSiteDetailsList)
		if err == nil {
			return nil // data stored in dst.NiatelemetryMsoSiteDetailsList, return on the first match
		} else {
			dst.NiatelemetryMsoSiteDetailsList = nil
			return fmt.Errorf("failed to unmarshal NiatelemetryMsoSiteDetailsResponse as NiatelemetryMsoSiteDetailsList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiatelemetryMsoSiteDetailsResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NiatelemetryMsoSiteDetailsList != nil {
		return json.Marshal(&src.NiatelemetryMsoSiteDetailsList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiatelemetryMsoSiteDetailsResponse) GetActualInstance() interface{} {
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

	if obj.NiatelemetryMsoSiteDetailsList != nil {
		return obj.NiatelemetryMsoSiteDetailsList
	}

	// all schemas are nil
	return nil
}

type NullableNiatelemetryMsoSiteDetailsResponse struct {
	value *NiatelemetryMsoSiteDetailsResponse
	isSet bool
}

func (v NullableNiatelemetryMsoSiteDetailsResponse) Get() *NiatelemetryMsoSiteDetailsResponse {
	return v.value
}

func (v *NullableNiatelemetryMsoSiteDetailsResponse) Set(val *NiatelemetryMsoSiteDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryMsoSiteDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryMsoSiteDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryMsoSiteDetailsResponse(val *NiatelemetryMsoSiteDetailsResponse) *NullableNiatelemetryMsoSiteDetailsResponse {
	return &NullableNiatelemetryMsoSiteDetailsResponse{value: val, isSet: true}
}

func (v NullableNiatelemetryMsoSiteDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryMsoSiteDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
