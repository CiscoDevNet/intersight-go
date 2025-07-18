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

// IaasServiceRequestResponse - The response body of a HTTP GET request for the 'iaas.ServiceRequest' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'iaas.ServiceRequest' resources.
type IaasServiceRequestResponse struct {
	IaasServiceRequestList *IaasServiceRequestList
	MoAggregateTransform   *MoAggregateTransform
	MoDocumentCount        *MoDocumentCount
	MoTagSummary           *MoTagSummary
}

// IaasServiceRequestListAsIaasServiceRequestResponse is a convenience function that returns IaasServiceRequestList wrapped in IaasServiceRequestResponse
func IaasServiceRequestListAsIaasServiceRequestResponse(v *IaasServiceRequestList) IaasServiceRequestResponse {
	return IaasServiceRequestResponse{
		IaasServiceRequestList: v,
	}
}

// MoAggregateTransformAsIaasServiceRequestResponse is a convenience function that returns MoAggregateTransform wrapped in IaasServiceRequestResponse
func MoAggregateTransformAsIaasServiceRequestResponse(v *MoAggregateTransform) IaasServiceRequestResponse {
	return IaasServiceRequestResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsIaasServiceRequestResponse is a convenience function that returns MoDocumentCount wrapped in IaasServiceRequestResponse
func MoDocumentCountAsIaasServiceRequestResponse(v *MoDocumentCount) IaasServiceRequestResponse {
	return IaasServiceRequestResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsIaasServiceRequestResponse is a convenience function that returns MoTagSummary wrapped in IaasServiceRequestResponse
func MoTagSummaryAsIaasServiceRequestResponse(v *MoTagSummary) IaasServiceRequestResponse {
	return IaasServiceRequestResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IaasServiceRequestResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'iaas.ServiceRequest.List'
	if jsonDict["ObjectType"] == "iaas.ServiceRequest.List" {
		// try to unmarshal JSON data into IaasServiceRequestList
		err = json.Unmarshal(data, &dst.IaasServiceRequestList)
		if err == nil {
			return nil // data stored in dst.IaasServiceRequestList, return on the first match
		} else {
			dst.IaasServiceRequestList = nil
			return fmt.Errorf("failed to unmarshal IaasServiceRequestResponse as IaasServiceRequestList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IaasServiceRequestResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IaasServiceRequestResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IaasServiceRequestResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IaasServiceRequestResponse) MarshalJSON() ([]byte, error) {
	if src.IaasServiceRequestList != nil {
		return json.Marshal(&src.IaasServiceRequestList)
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
func (obj *IaasServiceRequestResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IaasServiceRequestList != nil {
		return obj.IaasServiceRequestList
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

type NullableIaasServiceRequestResponse struct {
	value *IaasServiceRequestResponse
	isSet bool
}

func (v NullableIaasServiceRequestResponse) Get() *IaasServiceRequestResponse {
	return v.value
}

func (v *NullableIaasServiceRequestResponse) Set(val *IaasServiceRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasServiceRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasServiceRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasServiceRequestResponse(val *IaasServiceRequestResponse) *NullableIaasServiceRequestResponse {
	return &NullableIaasServiceRequestResponse{value: val, isSet: true}
}

func (v NullableIaasServiceRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasServiceRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
