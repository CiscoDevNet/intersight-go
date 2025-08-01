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

// IamSystemResponse - The response body of a HTTP GET request for the 'iam.System' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'iam.System' resources.
type IamSystemResponse struct {
	IamSystemList        *IamSystemList
	MoAggregateTransform *MoAggregateTransform
	MoDocumentCount      *MoDocumentCount
	MoTagSummary         *MoTagSummary
}

// IamSystemListAsIamSystemResponse is a convenience function that returns IamSystemList wrapped in IamSystemResponse
func IamSystemListAsIamSystemResponse(v *IamSystemList) IamSystemResponse {
	return IamSystemResponse{
		IamSystemList: v,
	}
}

// MoAggregateTransformAsIamSystemResponse is a convenience function that returns MoAggregateTransform wrapped in IamSystemResponse
func MoAggregateTransformAsIamSystemResponse(v *MoAggregateTransform) IamSystemResponse {
	return IamSystemResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsIamSystemResponse is a convenience function that returns MoDocumentCount wrapped in IamSystemResponse
func MoDocumentCountAsIamSystemResponse(v *MoDocumentCount) IamSystemResponse {
	return IamSystemResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsIamSystemResponse is a convenience function that returns MoTagSummary wrapped in IamSystemResponse
func MoTagSummaryAsIamSystemResponse(v *MoTagSummary) IamSystemResponse {
	return IamSystemResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamSystemResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'iam.System.List'
	if jsonDict["ObjectType"] == "iam.System.List" {
		// try to unmarshal JSON data into IamSystemList
		err = json.Unmarshal(data, &dst.IamSystemList)
		if err == nil {
			return nil // data stored in dst.IamSystemList, return on the first match
		} else {
			dst.IamSystemList = nil
			return fmt.Errorf("failed to unmarshal IamSystemResponse as IamSystemList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamSystemResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamSystemResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamSystemResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamSystemResponse) MarshalJSON() ([]byte, error) {
	if src.IamSystemList != nil {
		return json.Marshal(&src.IamSystemList)
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
func (obj *IamSystemResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IamSystemList != nil {
		return obj.IamSystemList
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

type NullableIamSystemResponse struct {
	value *IamSystemResponse
	isSet bool
}

func (v NullableIamSystemResponse) Get() *IamSystemResponse {
	return v.value
}

func (v *NullableIamSystemResponse) Set(val *IamSystemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSystemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSystemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSystemResponse(val *IamSystemResponse) *NullableIamSystemResponse {
	return &NullableIamSystemResponse{value: val, isSet: true}
}

func (v NullableIamSystemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSystemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
