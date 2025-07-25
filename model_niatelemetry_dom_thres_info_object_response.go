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

// NiatelemetryDomThresInfoObjectResponse - The response body of a HTTP GET request for the 'niatelemetry.DomThresInfoObject' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'niatelemetry.DomThresInfoObject' resources.
type NiatelemetryDomThresInfoObjectResponse struct {
	MoAggregateTransform               *MoAggregateTransform
	MoDocumentCount                    *MoDocumentCount
	MoTagSummary                       *MoTagSummary
	NiatelemetryDomThresInfoObjectList *NiatelemetryDomThresInfoObjectList
}

// MoAggregateTransformAsNiatelemetryDomThresInfoObjectResponse is a convenience function that returns MoAggregateTransform wrapped in NiatelemetryDomThresInfoObjectResponse
func MoAggregateTransformAsNiatelemetryDomThresInfoObjectResponse(v *MoAggregateTransform) NiatelemetryDomThresInfoObjectResponse {
	return NiatelemetryDomThresInfoObjectResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsNiatelemetryDomThresInfoObjectResponse is a convenience function that returns MoDocumentCount wrapped in NiatelemetryDomThresInfoObjectResponse
func MoDocumentCountAsNiatelemetryDomThresInfoObjectResponse(v *MoDocumentCount) NiatelemetryDomThresInfoObjectResponse {
	return NiatelemetryDomThresInfoObjectResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsNiatelemetryDomThresInfoObjectResponse is a convenience function that returns MoTagSummary wrapped in NiatelemetryDomThresInfoObjectResponse
func MoTagSummaryAsNiatelemetryDomThresInfoObjectResponse(v *MoTagSummary) NiatelemetryDomThresInfoObjectResponse {
	return NiatelemetryDomThresInfoObjectResponse{
		MoTagSummary: v,
	}
}

// NiatelemetryDomThresInfoObjectListAsNiatelemetryDomThresInfoObjectResponse is a convenience function that returns NiatelemetryDomThresInfoObjectList wrapped in NiatelemetryDomThresInfoObjectResponse
func NiatelemetryDomThresInfoObjectListAsNiatelemetryDomThresInfoObjectResponse(v *NiatelemetryDomThresInfoObjectList) NiatelemetryDomThresInfoObjectResponse {
	return NiatelemetryDomThresInfoObjectResponse{
		NiatelemetryDomThresInfoObjectList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiatelemetryDomThresInfoObjectResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal NiatelemetryDomThresInfoObjectResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryDomThresInfoObjectResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal NiatelemetryDomThresInfoObjectResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niatelemetry.DomThresInfoObject.List'
	if jsonDict["ObjectType"] == "niatelemetry.DomThresInfoObject.List" {
		// try to unmarshal JSON data into NiatelemetryDomThresInfoObjectList
		err = json.Unmarshal(data, &dst.NiatelemetryDomThresInfoObjectList)
		if err == nil {
			return nil // data stored in dst.NiatelemetryDomThresInfoObjectList, return on the first match
		} else {
			dst.NiatelemetryDomThresInfoObjectList = nil
			return fmt.Errorf("failed to unmarshal NiatelemetryDomThresInfoObjectResponse as NiatelemetryDomThresInfoObjectList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiatelemetryDomThresInfoObjectResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NiatelemetryDomThresInfoObjectList != nil {
		return json.Marshal(&src.NiatelemetryDomThresInfoObjectList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiatelemetryDomThresInfoObjectResponse) GetActualInstance() interface{} {
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

	if obj.NiatelemetryDomThresInfoObjectList != nil {
		return obj.NiatelemetryDomThresInfoObjectList
	}

	// all schemas are nil
	return nil
}

type NullableNiatelemetryDomThresInfoObjectResponse struct {
	value *NiatelemetryDomThresInfoObjectResponse
	isSet bool
}

func (v NullableNiatelemetryDomThresInfoObjectResponse) Get() *NiatelemetryDomThresInfoObjectResponse {
	return v.value
}

func (v *NullableNiatelemetryDomThresInfoObjectResponse) Set(val *NiatelemetryDomThresInfoObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryDomThresInfoObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryDomThresInfoObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryDomThresInfoObjectResponse(val *NiatelemetryDomThresInfoObjectResponse) *NullableNiatelemetryDomThresInfoObjectResponse {
	return &NullableNiatelemetryDomThresInfoObjectResponse{value: val, isSet: true}
}

func (v NullableNiatelemetryDomThresInfoObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryDomThresInfoObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
