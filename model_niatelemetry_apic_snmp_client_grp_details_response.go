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

// NiatelemetryApicSnmpClientGrpDetailsResponse - The response body of a HTTP GET request for the 'niatelemetry.ApicSnmpClientGrpDetails' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'niatelemetry.ApicSnmpClientGrpDetails' resources.
type NiatelemetryApicSnmpClientGrpDetailsResponse struct {
	MoAggregateTransform                     *MoAggregateTransform
	MoDocumentCount                          *MoDocumentCount
	MoTagSummary                             *MoTagSummary
	NiatelemetryApicSnmpClientGrpDetailsList *NiatelemetryApicSnmpClientGrpDetailsList
}

// MoAggregateTransformAsNiatelemetryApicSnmpClientGrpDetailsResponse is a convenience function that returns MoAggregateTransform wrapped in NiatelemetryApicSnmpClientGrpDetailsResponse
func MoAggregateTransformAsNiatelemetryApicSnmpClientGrpDetailsResponse(v *MoAggregateTransform) NiatelemetryApicSnmpClientGrpDetailsResponse {
	return NiatelemetryApicSnmpClientGrpDetailsResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsNiatelemetryApicSnmpClientGrpDetailsResponse is a convenience function that returns MoDocumentCount wrapped in NiatelemetryApicSnmpClientGrpDetailsResponse
func MoDocumentCountAsNiatelemetryApicSnmpClientGrpDetailsResponse(v *MoDocumentCount) NiatelemetryApicSnmpClientGrpDetailsResponse {
	return NiatelemetryApicSnmpClientGrpDetailsResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsNiatelemetryApicSnmpClientGrpDetailsResponse is a convenience function that returns MoTagSummary wrapped in NiatelemetryApicSnmpClientGrpDetailsResponse
func MoTagSummaryAsNiatelemetryApicSnmpClientGrpDetailsResponse(v *MoTagSummary) NiatelemetryApicSnmpClientGrpDetailsResponse {
	return NiatelemetryApicSnmpClientGrpDetailsResponse{
		MoTagSummary: v,
	}
}

// NiatelemetryApicSnmpClientGrpDetailsListAsNiatelemetryApicSnmpClientGrpDetailsResponse is a convenience function that returns NiatelemetryApicSnmpClientGrpDetailsList wrapped in NiatelemetryApicSnmpClientGrpDetailsResponse
func NiatelemetryApicSnmpClientGrpDetailsListAsNiatelemetryApicSnmpClientGrpDetailsResponse(v *NiatelemetryApicSnmpClientGrpDetailsList) NiatelemetryApicSnmpClientGrpDetailsResponse {
	return NiatelemetryApicSnmpClientGrpDetailsResponse{
		NiatelemetryApicSnmpClientGrpDetailsList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiatelemetryApicSnmpClientGrpDetailsResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal NiatelemetryApicSnmpClientGrpDetailsResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryApicSnmpClientGrpDetailsResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryApicSnmpClientGrpDetailsResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niatelemetry.ApicSnmpClientGrpDetails.List'
	if jsonDict["ObjectType"] == "niatelemetry.ApicSnmpClientGrpDetails.List" {
		// try to unmarshal JSON data into NiatelemetryApicSnmpClientGrpDetailsList
		err = json.Unmarshal(data, &dst.NiatelemetryApicSnmpClientGrpDetailsList)
		if err == nil {
			return nil // data stored in dst.NiatelemetryApicSnmpClientGrpDetailsList, return on the first match
		} else {
			dst.NiatelemetryApicSnmpClientGrpDetailsList = nil
			return fmt.Errorf("failed to unmarshal NiatelemetryApicSnmpClientGrpDetailsResponse as NiatelemetryApicSnmpClientGrpDetailsList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiatelemetryApicSnmpClientGrpDetailsResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NiatelemetryApicSnmpClientGrpDetailsList != nil {
		return json.Marshal(&src.NiatelemetryApicSnmpClientGrpDetailsList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiatelemetryApicSnmpClientGrpDetailsResponse) GetActualInstance() interface{} {
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

	if obj.NiatelemetryApicSnmpClientGrpDetailsList != nil {
		return obj.NiatelemetryApicSnmpClientGrpDetailsList
	}

	// all schemas are nil
	return nil
}

type NullableNiatelemetryApicSnmpClientGrpDetailsResponse struct {
	value *NiatelemetryApicSnmpClientGrpDetailsResponse
	isSet bool
}

func (v NullableNiatelemetryApicSnmpClientGrpDetailsResponse) Get() *NiatelemetryApicSnmpClientGrpDetailsResponse {
	return v.value
}

func (v *NullableNiatelemetryApicSnmpClientGrpDetailsResponse) Set(val *NiatelemetryApicSnmpClientGrpDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicSnmpClientGrpDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicSnmpClientGrpDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicSnmpClientGrpDetailsResponse(val *NiatelemetryApicSnmpClientGrpDetailsResponse) *NullableNiatelemetryApicSnmpClientGrpDetailsResponse {
	return &NullableNiatelemetryApicSnmpClientGrpDetailsResponse{value: val, isSet: true}
}

func (v NullableNiatelemetryApicSnmpClientGrpDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicSnmpClientGrpDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
