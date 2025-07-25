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

// LicenseIncCustomerOpResponse - The response body of a HTTP GET request for the 'license.IncCustomerOp' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'license.IncCustomerOp' resources.
type LicenseIncCustomerOpResponse struct {
	LicenseIncCustomerOpList *LicenseIncCustomerOpList
	MoAggregateTransform     *MoAggregateTransform
	MoDocumentCount          *MoDocumentCount
	MoTagSummary             *MoTagSummary
}

// LicenseIncCustomerOpListAsLicenseIncCustomerOpResponse is a convenience function that returns LicenseIncCustomerOpList wrapped in LicenseIncCustomerOpResponse
func LicenseIncCustomerOpListAsLicenseIncCustomerOpResponse(v *LicenseIncCustomerOpList) LicenseIncCustomerOpResponse {
	return LicenseIncCustomerOpResponse{
		LicenseIncCustomerOpList: v,
	}
}

// MoAggregateTransformAsLicenseIncCustomerOpResponse is a convenience function that returns MoAggregateTransform wrapped in LicenseIncCustomerOpResponse
func MoAggregateTransformAsLicenseIncCustomerOpResponse(v *MoAggregateTransform) LicenseIncCustomerOpResponse {
	return LicenseIncCustomerOpResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsLicenseIncCustomerOpResponse is a convenience function that returns MoDocumentCount wrapped in LicenseIncCustomerOpResponse
func MoDocumentCountAsLicenseIncCustomerOpResponse(v *MoDocumentCount) LicenseIncCustomerOpResponse {
	return LicenseIncCustomerOpResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsLicenseIncCustomerOpResponse is a convenience function that returns MoTagSummary wrapped in LicenseIncCustomerOpResponse
func MoTagSummaryAsLicenseIncCustomerOpResponse(v *MoTagSummary) LicenseIncCustomerOpResponse {
	return LicenseIncCustomerOpResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LicenseIncCustomerOpResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'license.IncCustomerOp.List'
	if jsonDict["ObjectType"] == "license.IncCustomerOp.List" {
		// try to unmarshal JSON data into LicenseIncCustomerOpList
		err = json.Unmarshal(data, &dst.LicenseIncCustomerOpList)
		if err == nil {
			return nil // data stored in dst.LicenseIncCustomerOpList, return on the first match
		} else {
			dst.LicenseIncCustomerOpList = nil
			return fmt.Errorf("failed to unmarshal LicenseIncCustomerOpResponse as LicenseIncCustomerOpList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal LicenseIncCustomerOpResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal LicenseIncCustomerOpResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal LicenseIncCustomerOpResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LicenseIncCustomerOpResponse) MarshalJSON() ([]byte, error) {
	if src.LicenseIncCustomerOpList != nil {
		return json.Marshal(&src.LicenseIncCustomerOpList)
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
func (obj *LicenseIncCustomerOpResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.LicenseIncCustomerOpList != nil {
		return obj.LicenseIncCustomerOpList
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

type NullableLicenseIncCustomerOpResponse struct {
	value *LicenseIncCustomerOpResponse
	isSet bool
}

func (v NullableLicenseIncCustomerOpResponse) Get() *LicenseIncCustomerOpResponse {
	return v.value
}

func (v *NullableLicenseIncCustomerOpResponse) Set(val *LicenseIncCustomerOpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseIncCustomerOpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseIncCustomerOpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseIncCustomerOpResponse(val *LicenseIncCustomerOpResponse) *NullableLicenseIncCustomerOpResponse {
	return &NullableLicenseIncCustomerOpResponse{value: val, isSet: true}
}

func (v NullableLicenseIncCustomerOpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseIncCustomerOpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
