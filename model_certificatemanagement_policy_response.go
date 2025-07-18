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

// CertificatemanagementPolicyResponse - The response body of a HTTP GET request for the 'certificatemanagement.Policy' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'certificatemanagement.Policy' resources.
type CertificatemanagementPolicyResponse struct {
	CertificatemanagementPolicyList *CertificatemanagementPolicyList
	MoAggregateTransform            *MoAggregateTransform
	MoDocumentCount                 *MoDocumentCount
	MoTagSummary                    *MoTagSummary
}

// CertificatemanagementPolicyListAsCertificatemanagementPolicyResponse is a convenience function that returns CertificatemanagementPolicyList wrapped in CertificatemanagementPolicyResponse
func CertificatemanagementPolicyListAsCertificatemanagementPolicyResponse(v *CertificatemanagementPolicyList) CertificatemanagementPolicyResponse {
	return CertificatemanagementPolicyResponse{
		CertificatemanagementPolicyList: v,
	}
}

// MoAggregateTransformAsCertificatemanagementPolicyResponse is a convenience function that returns MoAggregateTransform wrapped in CertificatemanagementPolicyResponse
func MoAggregateTransformAsCertificatemanagementPolicyResponse(v *MoAggregateTransform) CertificatemanagementPolicyResponse {
	return CertificatemanagementPolicyResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsCertificatemanagementPolicyResponse is a convenience function that returns MoDocumentCount wrapped in CertificatemanagementPolicyResponse
func MoDocumentCountAsCertificatemanagementPolicyResponse(v *MoDocumentCount) CertificatemanagementPolicyResponse {
	return CertificatemanagementPolicyResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsCertificatemanagementPolicyResponse is a convenience function that returns MoTagSummary wrapped in CertificatemanagementPolicyResponse
func MoTagSummaryAsCertificatemanagementPolicyResponse(v *MoTagSummary) CertificatemanagementPolicyResponse {
	return CertificatemanagementPolicyResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CertificatemanagementPolicyResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'certificatemanagement.Policy.List'
	if jsonDict["ObjectType"] == "certificatemanagement.Policy.List" {
		// try to unmarshal JSON data into CertificatemanagementPolicyList
		err = json.Unmarshal(data, &dst.CertificatemanagementPolicyList)
		if err == nil {
			return nil // data stored in dst.CertificatemanagementPolicyList, return on the first match
		} else {
			dst.CertificatemanagementPolicyList = nil
			return fmt.Errorf("failed to unmarshal CertificatemanagementPolicyResponse as CertificatemanagementPolicyList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal CertificatemanagementPolicyResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal CertificatemanagementPolicyResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal CertificatemanagementPolicyResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CertificatemanagementPolicyResponse) MarshalJSON() ([]byte, error) {
	if src.CertificatemanagementPolicyList != nil {
		return json.Marshal(&src.CertificatemanagementPolicyList)
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
func (obj *CertificatemanagementPolicyResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CertificatemanagementPolicyList != nil {
		return obj.CertificatemanagementPolicyList
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

type NullableCertificatemanagementPolicyResponse struct {
	value *CertificatemanagementPolicyResponse
	isSet bool
}

func (v NullableCertificatemanagementPolicyResponse) Get() *CertificatemanagementPolicyResponse {
	return v.value
}

func (v *NullableCertificatemanagementPolicyResponse) Set(val *CertificatemanagementPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificatemanagementPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificatemanagementPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificatemanagementPolicyResponse(val *CertificatemanagementPolicyResponse) *NullableCertificatemanagementPolicyResponse {
	return &NullableCertificatemanagementPolicyResponse{value: val, isSet: true}
}

func (v NullableCertificatemanagementPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificatemanagementPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
