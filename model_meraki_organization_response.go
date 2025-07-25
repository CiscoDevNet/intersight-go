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

// MerakiOrganizationResponse - The response body of a HTTP GET request for the 'meraki.Organization' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'meraki.Organization' resources.
type MerakiOrganizationResponse struct {
	MerakiOrganizationList *MerakiOrganizationList
	MoAggregateTransform   *MoAggregateTransform
	MoDocumentCount        *MoDocumentCount
	MoTagSummary           *MoTagSummary
}

// MerakiOrganizationListAsMerakiOrganizationResponse is a convenience function that returns MerakiOrganizationList wrapped in MerakiOrganizationResponse
func MerakiOrganizationListAsMerakiOrganizationResponse(v *MerakiOrganizationList) MerakiOrganizationResponse {
	return MerakiOrganizationResponse{
		MerakiOrganizationList: v,
	}
}

// MoAggregateTransformAsMerakiOrganizationResponse is a convenience function that returns MoAggregateTransform wrapped in MerakiOrganizationResponse
func MoAggregateTransformAsMerakiOrganizationResponse(v *MoAggregateTransform) MerakiOrganizationResponse {
	return MerakiOrganizationResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsMerakiOrganizationResponse is a convenience function that returns MoDocumentCount wrapped in MerakiOrganizationResponse
func MoDocumentCountAsMerakiOrganizationResponse(v *MoDocumentCount) MerakiOrganizationResponse {
	return MerakiOrganizationResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsMerakiOrganizationResponse is a convenience function that returns MoTagSummary wrapped in MerakiOrganizationResponse
func MoTagSummaryAsMerakiOrganizationResponse(v *MoTagSummary) MerakiOrganizationResponse {
	return MerakiOrganizationResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MerakiOrganizationResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'meraki.Organization.List'
	if jsonDict["ObjectType"] == "meraki.Organization.List" {
		// try to unmarshal JSON data into MerakiOrganizationList
		err = json.Unmarshal(data, &dst.MerakiOrganizationList)
		if err == nil {
			return nil // data stored in dst.MerakiOrganizationList, return on the first match
		} else {
			dst.MerakiOrganizationList = nil
			return fmt.Errorf("failed to unmarshal MerakiOrganizationResponse as MerakiOrganizationList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MerakiOrganizationResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MerakiOrganizationResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MerakiOrganizationResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MerakiOrganizationResponse) MarshalJSON() ([]byte, error) {
	if src.MerakiOrganizationList != nil {
		return json.Marshal(&src.MerakiOrganizationList)
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
func (obj *MerakiOrganizationResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MerakiOrganizationList != nil {
		return obj.MerakiOrganizationList
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

type NullableMerakiOrganizationResponse struct {
	value *MerakiOrganizationResponse
	isSet bool
}

func (v NullableMerakiOrganizationResponse) Get() *MerakiOrganizationResponse {
	return v.value
}

func (v *NullableMerakiOrganizationResponse) Set(val *MerakiOrganizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMerakiOrganizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMerakiOrganizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerakiOrganizationResponse(val *MerakiOrganizationResponse) *NullableMerakiOrganizationResponse {
	return &NullableMerakiOrganizationResponse{value: val, isSet: true}
}

func (v NullableMerakiOrganizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerakiOrganizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
