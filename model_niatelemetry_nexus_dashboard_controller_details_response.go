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

// NiatelemetryNexusDashboardControllerDetailsResponse - The response body of a HTTP GET request for the 'niatelemetry.NexusDashboardControllerDetails' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'niatelemetry.NexusDashboardControllerDetails' resources.
type NiatelemetryNexusDashboardControllerDetailsResponse struct {
	MoAggregateTransform                            *MoAggregateTransform
	MoDocumentCount                                 *MoDocumentCount
	MoTagSummary                                    *MoTagSummary
	NiatelemetryNexusDashboardControllerDetailsList *NiatelemetryNexusDashboardControllerDetailsList
}

// MoAggregateTransformAsNiatelemetryNexusDashboardControllerDetailsResponse is a convenience function that returns MoAggregateTransform wrapped in NiatelemetryNexusDashboardControllerDetailsResponse
func MoAggregateTransformAsNiatelemetryNexusDashboardControllerDetailsResponse(v *MoAggregateTransform) NiatelemetryNexusDashboardControllerDetailsResponse {
	return NiatelemetryNexusDashboardControllerDetailsResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsNiatelemetryNexusDashboardControllerDetailsResponse is a convenience function that returns MoDocumentCount wrapped in NiatelemetryNexusDashboardControllerDetailsResponse
func MoDocumentCountAsNiatelemetryNexusDashboardControllerDetailsResponse(v *MoDocumentCount) NiatelemetryNexusDashboardControllerDetailsResponse {
	return NiatelemetryNexusDashboardControllerDetailsResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsNiatelemetryNexusDashboardControllerDetailsResponse is a convenience function that returns MoTagSummary wrapped in NiatelemetryNexusDashboardControllerDetailsResponse
func MoTagSummaryAsNiatelemetryNexusDashboardControllerDetailsResponse(v *MoTagSummary) NiatelemetryNexusDashboardControllerDetailsResponse {
	return NiatelemetryNexusDashboardControllerDetailsResponse{
		MoTagSummary: v,
	}
}

// NiatelemetryNexusDashboardControllerDetailsListAsNiatelemetryNexusDashboardControllerDetailsResponse is a convenience function that returns NiatelemetryNexusDashboardControllerDetailsList wrapped in NiatelemetryNexusDashboardControllerDetailsResponse
func NiatelemetryNexusDashboardControllerDetailsListAsNiatelemetryNexusDashboardControllerDetailsResponse(v *NiatelemetryNexusDashboardControllerDetailsList) NiatelemetryNexusDashboardControllerDetailsResponse {
	return NiatelemetryNexusDashboardControllerDetailsResponse{
		NiatelemetryNexusDashboardControllerDetailsList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiatelemetryNexusDashboardControllerDetailsResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal NiatelemetryNexusDashboardControllerDetailsResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryNexusDashboardControllerDetailsResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryNexusDashboardControllerDetailsResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niatelemetry.NexusDashboardControllerDetails.List'
	if jsonDict["ObjectType"] == "niatelemetry.NexusDashboardControllerDetails.List" {
		// try to unmarshal JSON data into NiatelemetryNexusDashboardControllerDetailsList
		err = json.Unmarshal(data, &dst.NiatelemetryNexusDashboardControllerDetailsList)
		if err == nil {
			return nil // data stored in dst.NiatelemetryNexusDashboardControllerDetailsList, return on the first match
		} else {
			dst.NiatelemetryNexusDashboardControllerDetailsList = nil
			return fmt.Errorf("failed to unmarshal NiatelemetryNexusDashboardControllerDetailsResponse as NiatelemetryNexusDashboardControllerDetailsList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiatelemetryNexusDashboardControllerDetailsResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NiatelemetryNexusDashboardControllerDetailsList != nil {
		return json.Marshal(&src.NiatelemetryNexusDashboardControllerDetailsList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiatelemetryNexusDashboardControllerDetailsResponse) GetActualInstance() interface{} {
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

	if obj.NiatelemetryNexusDashboardControllerDetailsList != nil {
		return obj.NiatelemetryNexusDashboardControllerDetailsList
	}

	// all schemas are nil
	return nil
}

type NullableNiatelemetryNexusDashboardControllerDetailsResponse struct {
	value *NiatelemetryNexusDashboardControllerDetailsResponse
	isSet bool
}

func (v NullableNiatelemetryNexusDashboardControllerDetailsResponse) Get() *NiatelemetryNexusDashboardControllerDetailsResponse {
	return v.value
}

func (v *NullableNiatelemetryNexusDashboardControllerDetailsResponse) Set(val *NiatelemetryNexusDashboardControllerDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNexusDashboardControllerDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNexusDashboardControllerDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNexusDashboardControllerDetailsResponse(val *NiatelemetryNexusDashboardControllerDetailsResponse) *NullableNiatelemetryNexusDashboardControllerDetailsResponse {
	return &NullableNiatelemetryNexusDashboardControllerDetailsResponse{value: val, isSet: true}
}

func (v NullableNiatelemetryNexusDashboardControllerDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNexusDashboardControllerDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
