/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-17769
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// DnacTransitResponse - The response body of a HTTP GET request for the 'dnac.Transit' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'dnac.Transit' resources.
type DnacTransitResponse struct {
	DnacTransitList      *DnacTransitList
	MoAggregateTransform *MoAggregateTransform
	MoDocumentCount      *MoDocumentCount
	MoTagSummary         *MoTagSummary
}

// DnacTransitListAsDnacTransitResponse is a convenience function that returns DnacTransitList wrapped in DnacTransitResponse
func DnacTransitListAsDnacTransitResponse(v *DnacTransitList) DnacTransitResponse {
	return DnacTransitResponse{
		DnacTransitList: v,
	}
}

// MoAggregateTransformAsDnacTransitResponse is a convenience function that returns MoAggregateTransform wrapped in DnacTransitResponse
func MoAggregateTransformAsDnacTransitResponse(v *MoAggregateTransform) DnacTransitResponse {
	return DnacTransitResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsDnacTransitResponse is a convenience function that returns MoDocumentCount wrapped in DnacTransitResponse
func MoDocumentCountAsDnacTransitResponse(v *MoDocumentCount) DnacTransitResponse {
	return DnacTransitResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsDnacTransitResponse is a convenience function that returns MoTagSummary wrapped in DnacTransitResponse
func MoTagSummaryAsDnacTransitResponse(v *MoTagSummary) DnacTransitResponse {
	return DnacTransitResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DnacTransitResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'dnac.Transit.List'
	if jsonDict["ObjectType"] == "dnac.Transit.List" {
		// try to unmarshal JSON data into DnacTransitList
		err = json.Unmarshal(data, &dst.DnacTransitList)
		if err == nil {
			return nil // data stored in dst.DnacTransitList, return on the first match
		} else {
			dst.DnacTransitList = nil
			return fmt.Errorf("failed to unmarshal DnacTransitResponse as DnacTransitList: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("failed to unmarshal DnacTransitResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal DnacTransitResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal DnacTransitResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DnacTransitResponse) MarshalJSON() ([]byte, error) {
	if src.DnacTransitList != nil {
		return json.Marshal(&src.DnacTransitList)
	}

	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DnacTransitResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DnacTransitList != nil {
		return obj.DnacTransitList
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

	// all schemas are nil
	return nil
}

type NullableDnacTransitResponse struct {
	value *DnacTransitResponse
	isSet bool
}

func (v NullableDnacTransitResponse) Get() *DnacTransitResponse {
	return v.value
}

func (v *NullableDnacTransitResponse) Set(val *DnacTransitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDnacTransitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDnacTransitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnacTransitResponse(val *DnacTransitResponse) *NullableDnacTransitResponse {
	return &NullableDnacTransitResponse{value: val, isSet: true}
}

func (v NullableDnacTransitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnacTransitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}