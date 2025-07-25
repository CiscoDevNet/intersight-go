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

// LicenseSmartlicenseTokenResponse - The response body of a HTTP GET request for the 'license.SmartlicenseToken' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'license.SmartlicenseToken' resources.
type LicenseSmartlicenseTokenResponse struct {
	LicenseSmartlicenseTokenList *LicenseSmartlicenseTokenList
	MoAggregateTransform         *MoAggregateTransform
	MoDocumentCount              *MoDocumentCount
	MoTagSummary                 *MoTagSummary
}

// LicenseSmartlicenseTokenListAsLicenseSmartlicenseTokenResponse is a convenience function that returns LicenseSmartlicenseTokenList wrapped in LicenseSmartlicenseTokenResponse
func LicenseSmartlicenseTokenListAsLicenseSmartlicenseTokenResponse(v *LicenseSmartlicenseTokenList) LicenseSmartlicenseTokenResponse {
	return LicenseSmartlicenseTokenResponse{
		LicenseSmartlicenseTokenList: v,
	}
}

// MoAggregateTransformAsLicenseSmartlicenseTokenResponse is a convenience function that returns MoAggregateTransform wrapped in LicenseSmartlicenseTokenResponse
func MoAggregateTransformAsLicenseSmartlicenseTokenResponse(v *MoAggregateTransform) LicenseSmartlicenseTokenResponse {
	return LicenseSmartlicenseTokenResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsLicenseSmartlicenseTokenResponse is a convenience function that returns MoDocumentCount wrapped in LicenseSmartlicenseTokenResponse
func MoDocumentCountAsLicenseSmartlicenseTokenResponse(v *MoDocumentCount) LicenseSmartlicenseTokenResponse {
	return LicenseSmartlicenseTokenResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsLicenseSmartlicenseTokenResponse is a convenience function that returns MoTagSummary wrapped in LicenseSmartlicenseTokenResponse
func MoTagSummaryAsLicenseSmartlicenseTokenResponse(v *MoTagSummary) LicenseSmartlicenseTokenResponse {
	return LicenseSmartlicenseTokenResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LicenseSmartlicenseTokenResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'license.SmartlicenseToken.List'
	if jsonDict["ObjectType"] == "license.SmartlicenseToken.List" {
		// try to unmarshal JSON data into LicenseSmartlicenseTokenList
		err = json.Unmarshal(data, &dst.LicenseSmartlicenseTokenList)
		if err == nil {
			return nil // data stored in dst.LicenseSmartlicenseTokenList, return on the first match
		} else {
			dst.LicenseSmartlicenseTokenList = nil
			return fmt.Errorf("failed to unmarshal LicenseSmartlicenseTokenResponse as LicenseSmartlicenseTokenList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal LicenseSmartlicenseTokenResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal LicenseSmartlicenseTokenResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal LicenseSmartlicenseTokenResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LicenseSmartlicenseTokenResponse) MarshalJSON() ([]byte, error) {
	if src.LicenseSmartlicenseTokenList != nil {
		return json.Marshal(&src.LicenseSmartlicenseTokenList)
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
func (obj *LicenseSmartlicenseTokenResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.LicenseSmartlicenseTokenList != nil {
		return obj.LicenseSmartlicenseTokenList
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

type NullableLicenseSmartlicenseTokenResponse struct {
	value *LicenseSmartlicenseTokenResponse
	isSet bool
}

func (v NullableLicenseSmartlicenseTokenResponse) Get() *LicenseSmartlicenseTokenResponse {
	return v.value
}

func (v *NullableLicenseSmartlicenseTokenResponse) Set(val *LicenseSmartlicenseTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseSmartlicenseTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseSmartlicenseTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseSmartlicenseTokenResponse(val *LicenseSmartlicenseTokenResponse) *NullableLicenseSmartlicenseTokenResponse {
	return &NullableLicenseSmartlicenseTokenResponse{value: val, isSet: true}
}

func (v NullableLicenseSmartlicenseTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseSmartlicenseTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
