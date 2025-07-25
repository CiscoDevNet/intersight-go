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

// IamCertificateRequestResponse - The response body of a HTTP GET request for the 'iam.CertificateRequest' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'iam.CertificateRequest' resources.
type IamCertificateRequestResponse struct {
	IamCertificateRequestList *IamCertificateRequestList
	MoAggregateTransform      *MoAggregateTransform
	MoDocumentCount           *MoDocumentCount
	MoTagSummary              *MoTagSummary
}

// IamCertificateRequestListAsIamCertificateRequestResponse is a convenience function that returns IamCertificateRequestList wrapped in IamCertificateRequestResponse
func IamCertificateRequestListAsIamCertificateRequestResponse(v *IamCertificateRequestList) IamCertificateRequestResponse {
	return IamCertificateRequestResponse{
		IamCertificateRequestList: v,
	}
}

// MoAggregateTransformAsIamCertificateRequestResponse is a convenience function that returns MoAggregateTransform wrapped in IamCertificateRequestResponse
func MoAggregateTransformAsIamCertificateRequestResponse(v *MoAggregateTransform) IamCertificateRequestResponse {
	return IamCertificateRequestResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsIamCertificateRequestResponse is a convenience function that returns MoDocumentCount wrapped in IamCertificateRequestResponse
func MoDocumentCountAsIamCertificateRequestResponse(v *MoDocumentCount) IamCertificateRequestResponse {
	return IamCertificateRequestResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsIamCertificateRequestResponse is a convenience function that returns MoTagSummary wrapped in IamCertificateRequestResponse
func MoTagSummaryAsIamCertificateRequestResponse(v *MoTagSummary) IamCertificateRequestResponse {
	return IamCertificateRequestResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamCertificateRequestResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'iam.CertificateRequest.List'
	if jsonDict["ObjectType"] == "iam.CertificateRequest.List" {
		// try to unmarshal JSON data into IamCertificateRequestList
		err = json.Unmarshal(data, &dst.IamCertificateRequestList)
		if err == nil {
			return nil // data stored in dst.IamCertificateRequestList, return on the first match
		} else {
			dst.IamCertificateRequestList = nil
			return fmt.Errorf("failed to unmarshal IamCertificateRequestResponse as IamCertificateRequestList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamCertificateRequestResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamCertificateRequestResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamCertificateRequestResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamCertificateRequestResponse) MarshalJSON() ([]byte, error) {
	if src.IamCertificateRequestList != nil {
		return json.Marshal(&src.IamCertificateRequestList)
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
func (obj *IamCertificateRequestResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IamCertificateRequestList != nil {
		return obj.IamCertificateRequestList
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

type NullableIamCertificateRequestResponse struct {
	value *IamCertificateRequestResponse
	isSet bool
}

func (v NullableIamCertificateRequestResponse) Get() *IamCertificateRequestResponse {
	return v.value
}

func (v *NullableIamCertificateRequestResponse) Set(val *IamCertificateRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIamCertificateRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIamCertificateRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamCertificateRequestResponse(val *IamCertificateRequestResponse) *NullableIamCertificateRequestResponse {
	return &NullableIamCertificateRequestResponse{value: val, isSet: true}
}

func (v NullableIamCertificateRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamCertificateRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
