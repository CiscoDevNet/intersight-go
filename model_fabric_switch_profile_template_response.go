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

// FabricSwitchProfileTemplateResponse - The response body of a HTTP GET request for the 'fabric.SwitchProfileTemplate' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'fabric.SwitchProfileTemplate' resources.
type FabricSwitchProfileTemplateResponse struct {
	FabricSwitchProfileTemplateList *FabricSwitchProfileTemplateList
	MoAggregateTransform            *MoAggregateTransform
	MoDocumentCount                 *MoDocumentCount
	MoTagSummary                    *MoTagSummary
}

// FabricSwitchProfileTemplateListAsFabricSwitchProfileTemplateResponse is a convenience function that returns FabricSwitchProfileTemplateList wrapped in FabricSwitchProfileTemplateResponse
func FabricSwitchProfileTemplateListAsFabricSwitchProfileTemplateResponse(v *FabricSwitchProfileTemplateList) FabricSwitchProfileTemplateResponse {
	return FabricSwitchProfileTemplateResponse{
		FabricSwitchProfileTemplateList: v,
	}
}

// MoAggregateTransformAsFabricSwitchProfileTemplateResponse is a convenience function that returns MoAggregateTransform wrapped in FabricSwitchProfileTemplateResponse
func MoAggregateTransformAsFabricSwitchProfileTemplateResponse(v *MoAggregateTransform) FabricSwitchProfileTemplateResponse {
	return FabricSwitchProfileTemplateResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsFabricSwitchProfileTemplateResponse is a convenience function that returns MoDocumentCount wrapped in FabricSwitchProfileTemplateResponse
func MoDocumentCountAsFabricSwitchProfileTemplateResponse(v *MoDocumentCount) FabricSwitchProfileTemplateResponse {
	return FabricSwitchProfileTemplateResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsFabricSwitchProfileTemplateResponse is a convenience function that returns MoTagSummary wrapped in FabricSwitchProfileTemplateResponse
func MoTagSummaryAsFabricSwitchProfileTemplateResponse(v *MoTagSummary) FabricSwitchProfileTemplateResponse {
	return FabricSwitchProfileTemplateResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FabricSwitchProfileTemplateResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'fabric.SwitchProfileTemplate.List'
	if jsonDict["ObjectType"] == "fabric.SwitchProfileTemplate.List" {
		// try to unmarshal JSON data into FabricSwitchProfileTemplateList
		err = json.Unmarshal(data, &dst.FabricSwitchProfileTemplateList)
		if err == nil {
			return nil // data stored in dst.FabricSwitchProfileTemplateList, return on the first match
		} else {
			dst.FabricSwitchProfileTemplateList = nil
			return fmt.Errorf("failed to unmarshal FabricSwitchProfileTemplateResponse as FabricSwitchProfileTemplateList: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FabricSwitchProfileTemplateResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FabricSwitchProfileTemplateResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FabricSwitchProfileTemplateResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FabricSwitchProfileTemplateResponse) MarshalJSON() ([]byte, error) {
	if src.FabricSwitchProfileTemplateList != nil {
		return json.Marshal(&src.FabricSwitchProfileTemplateList)
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
func (obj *FabricSwitchProfileTemplateResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FabricSwitchProfileTemplateList != nil {
		return obj.FabricSwitchProfileTemplateList
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

type NullableFabricSwitchProfileTemplateResponse struct {
	value *FabricSwitchProfileTemplateResponse
	isSet bool
}

func (v NullableFabricSwitchProfileTemplateResponse) Get() *FabricSwitchProfileTemplateResponse {
	return v.value
}

func (v *NullableFabricSwitchProfileTemplateResponse) Set(val *FabricSwitchProfileTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSwitchProfileTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSwitchProfileTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSwitchProfileTemplateResponse(val *FabricSwitchProfileTemplateResponse) *NullableFabricSwitchProfileTemplateResponse {
	return &NullableFabricSwitchProfileTemplateResponse{value: val, isSet: true}
}

func (v NullableFabricSwitchProfileTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSwitchProfileTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
