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

// IamEndPointUserPolicyResponse - The response body of a HTTP GET request for the 'iam.EndPointUserPolicy' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'iam.EndPointUserPolicy' resources.
type IamEndPointUserPolicyResponse struct {
	IamEndPointUserPolicyList *IamEndPointUserPolicyList
	MoAggregateTransform      *MoAggregateTransform
	MoDocumentCount           *MoDocumentCount
	MoTagSummary              *MoTagSummary
}

// IamEndPointUserPolicyListAsIamEndPointUserPolicyResponse is a convenience function that returns IamEndPointUserPolicyList wrapped in IamEndPointUserPolicyResponse
func IamEndPointUserPolicyListAsIamEndPointUserPolicyResponse(v *IamEndPointUserPolicyList) IamEndPointUserPolicyResponse {
	return IamEndPointUserPolicyResponse{
		IamEndPointUserPolicyList: v,
	}
}

// MoAggregateTransformAsIamEndPointUserPolicyResponse is a convenience function that returns MoAggregateTransform wrapped in IamEndPointUserPolicyResponse
func MoAggregateTransformAsIamEndPointUserPolicyResponse(v *MoAggregateTransform) IamEndPointUserPolicyResponse {
	return IamEndPointUserPolicyResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsIamEndPointUserPolicyResponse is a convenience function that returns MoDocumentCount wrapped in IamEndPointUserPolicyResponse
func MoDocumentCountAsIamEndPointUserPolicyResponse(v *MoDocumentCount) IamEndPointUserPolicyResponse {
	return IamEndPointUserPolicyResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsIamEndPointUserPolicyResponse is a convenience function that returns MoTagSummary wrapped in IamEndPointUserPolicyResponse
func MoTagSummaryAsIamEndPointUserPolicyResponse(v *MoTagSummary) IamEndPointUserPolicyResponse {
	return IamEndPointUserPolicyResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamEndPointUserPolicyResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'iam.EndPointUserPolicy.List'
	if jsonDict["ObjectType"] == "iam.EndPointUserPolicy.List" {
		// try to unmarshal JSON data into IamEndPointUserPolicyList
		err = json.Unmarshal(data, &dst.IamEndPointUserPolicyList)
		if err == nil {
			return nil // data stored in dst.IamEndPointUserPolicyList, return on the first match
		} else {
			dst.IamEndPointUserPolicyList = nil
			return fmt.Errorf("failed to unmarshal IamEndPointUserPolicyResponse as IamEndPointUserPolicyList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamEndPointUserPolicyResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamEndPointUserPolicyResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamEndPointUserPolicyResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamEndPointUserPolicyResponse) MarshalJSON() ([]byte, error) {
	if src.IamEndPointUserPolicyList != nil {
		return json.Marshal(&src.IamEndPointUserPolicyList)
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
func (obj *IamEndPointUserPolicyResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IamEndPointUserPolicyList != nil {
		return obj.IamEndPointUserPolicyList
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

type NullableIamEndPointUserPolicyResponse struct {
	value *IamEndPointUserPolicyResponse
	isSet bool
}

func (v NullableIamEndPointUserPolicyResponse) Get() *IamEndPointUserPolicyResponse {
	return v.value
}

func (v *NullableIamEndPointUserPolicyResponse) Set(val *IamEndPointUserPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointUserPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointUserPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointUserPolicyResponse(val *IamEndPointUserPolicyResponse) *NullableIamEndPointUserPolicyResponse {
	return &NullableIamEndPointUserPolicyResponse{value: val, isSet: true}
}

func (v NullableIamEndPointUserPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointUserPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
