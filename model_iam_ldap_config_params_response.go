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

// IamLdapConfigParamsResponse - The response body of a HTTP GET request for the 'iam.LdapConfigParams' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'iam.LdapConfigParams' resources.
type IamLdapConfigParamsResponse struct {
	IamLdapConfigParamsList *IamLdapConfigParamsList
	MoAggregateTransform    *MoAggregateTransform
	MoDocumentCount         *MoDocumentCount
	MoTagSummary            *MoTagSummary
}

// IamLdapConfigParamsListAsIamLdapConfigParamsResponse is a convenience function that returns IamLdapConfigParamsList wrapped in IamLdapConfigParamsResponse
func IamLdapConfigParamsListAsIamLdapConfigParamsResponse(v *IamLdapConfigParamsList) IamLdapConfigParamsResponse {
	return IamLdapConfigParamsResponse{
		IamLdapConfigParamsList: v,
	}
}

// MoAggregateTransformAsIamLdapConfigParamsResponse is a convenience function that returns MoAggregateTransform wrapped in IamLdapConfigParamsResponse
func MoAggregateTransformAsIamLdapConfigParamsResponse(v *MoAggregateTransform) IamLdapConfigParamsResponse {
	return IamLdapConfigParamsResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsIamLdapConfigParamsResponse is a convenience function that returns MoDocumentCount wrapped in IamLdapConfigParamsResponse
func MoDocumentCountAsIamLdapConfigParamsResponse(v *MoDocumentCount) IamLdapConfigParamsResponse {
	return IamLdapConfigParamsResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsIamLdapConfigParamsResponse is a convenience function that returns MoTagSummary wrapped in IamLdapConfigParamsResponse
func MoTagSummaryAsIamLdapConfigParamsResponse(v *MoTagSummary) IamLdapConfigParamsResponse {
	return IamLdapConfigParamsResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamLdapConfigParamsResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'iam.LdapConfigParams.List'
	if jsonDict["ObjectType"] == "iam.LdapConfigParams.List" {
		// try to unmarshal JSON data into IamLdapConfigParamsList
		err = json.Unmarshal(data, &dst.IamLdapConfigParamsList)
		if err == nil {
			return nil // data stored in dst.IamLdapConfigParamsList, return on the first match
		} else {
			dst.IamLdapConfigParamsList = nil
			return fmt.Errorf("failed to unmarshal IamLdapConfigParamsResponse as IamLdapConfigParamsList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamLdapConfigParamsResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamLdapConfigParamsResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamLdapConfigParamsResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamLdapConfigParamsResponse) MarshalJSON() ([]byte, error) {
	if src.IamLdapConfigParamsList != nil {
		return json.Marshal(&src.IamLdapConfigParamsList)
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
func (obj *IamLdapConfigParamsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IamLdapConfigParamsList != nil {
		return obj.IamLdapConfigParamsList
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

type NullableIamLdapConfigParamsResponse struct {
	value *IamLdapConfigParamsResponse
	isSet bool
}

func (v NullableIamLdapConfigParamsResponse) Get() *IamLdapConfigParamsResponse {
	return v.value
}

func (v *NullableIamLdapConfigParamsResponse) Set(val *IamLdapConfigParamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIamLdapConfigParamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIamLdapConfigParamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamLdapConfigParamsResponse(val *IamLdapConfigParamsResponse) *NullableIamLdapConfigParamsResponse {
	return &NullableIamLdapConfigParamsResponse{value: val, isSet: true}
}

func (v NullableIamLdapConfigParamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamLdapConfigParamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
