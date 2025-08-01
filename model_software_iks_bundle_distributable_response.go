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

// SoftwareIksBundleDistributableResponse - The response body of a HTTP GET request for the 'software.IksBundleDistributable' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'software.IksBundleDistributable' resources.
type SoftwareIksBundleDistributableResponse struct {
	MoAggregateTransform               *MoAggregateTransform
	MoDocumentCount                    *MoDocumentCount
	MoTagSummary                       *MoTagSummary
	SoftwareIksBundleDistributableList *SoftwareIksBundleDistributableList
}

// MoAggregateTransformAsSoftwareIksBundleDistributableResponse is a convenience function that returns MoAggregateTransform wrapped in SoftwareIksBundleDistributableResponse
func MoAggregateTransformAsSoftwareIksBundleDistributableResponse(v *MoAggregateTransform) SoftwareIksBundleDistributableResponse {
	return SoftwareIksBundleDistributableResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsSoftwareIksBundleDistributableResponse is a convenience function that returns MoDocumentCount wrapped in SoftwareIksBundleDistributableResponse
func MoDocumentCountAsSoftwareIksBundleDistributableResponse(v *MoDocumentCount) SoftwareIksBundleDistributableResponse {
	return SoftwareIksBundleDistributableResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsSoftwareIksBundleDistributableResponse is a convenience function that returns MoTagSummary wrapped in SoftwareIksBundleDistributableResponse
func MoTagSummaryAsSoftwareIksBundleDistributableResponse(v *MoTagSummary) SoftwareIksBundleDistributableResponse {
	return SoftwareIksBundleDistributableResponse{
		MoTagSummary: v,
	}
}

// SoftwareIksBundleDistributableListAsSoftwareIksBundleDistributableResponse is a convenience function that returns SoftwareIksBundleDistributableList wrapped in SoftwareIksBundleDistributableResponse
func SoftwareIksBundleDistributableListAsSoftwareIksBundleDistributableResponse(v *SoftwareIksBundleDistributableList) SoftwareIksBundleDistributableResponse {
	return SoftwareIksBundleDistributableResponse{
		SoftwareIksBundleDistributableList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SoftwareIksBundleDistributableResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("failed to unmarshal SoftwareIksBundleDistributableResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SoftwareIksBundleDistributableResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SoftwareIksBundleDistributableResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'software.IksBundleDistributable.List'
	if jsonDict["ObjectType"] == "software.IksBundleDistributable.List" {
		// try to unmarshal JSON data into SoftwareIksBundleDistributableList
		err = json.Unmarshal(data, &dst.SoftwareIksBundleDistributableList)
		if err == nil {
			return nil // data stored in dst.SoftwareIksBundleDistributableList, return on the first match
		} else {
			dst.SoftwareIksBundleDistributableList = nil
			return fmt.Errorf("failed to unmarshal SoftwareIksBundleDistributableResponse as SoftwareIksBundleDistributableList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SoftwareIksBundleDistributableResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.SoftwareIksBundleDistributableList != nil {
		return json.Marshal(&src.SoftwareIksBundleDistributableList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SoftwareIksBundleDistributableResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
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

	if obj.SoftwareIksBundleDistributableList != nil {
		return obj.SoftwareIksBundleDistributableList
	}

	// all schemas are nil
	return nil
}

type NullableSoftwareIksBundleDistributableResponse struct {
	value *SoftwareIksBundleDistributableResponse
	isSet bool
}

func (v NullableSoftwareIksBundleDistributableResponse) Get() *SoftwareIksBundleDistributableResponse {
	return v.value
}

func (v *NullableSoftwareIksBundleDistributableResponse) Set(val *SoftwareIksBundleDistributableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareIksBundleDistributableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareIksBundleDistributableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareIksBundleDistributableResponse(val *SoftwareIksBundleDistributableResponse) *NullableSoftwareIksBundleDistributableResponse {
	return &NullableSoftwareIksBundleDistributableResponse{value: val, isSet: true}
}

func (v NullableSoftwareIksBundleDistributableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareIksBundleDistributableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
