/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-10371
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// EquipmentSwitchOperationResponse - The response body of a HTTP GET request for the 'equipment.SwitchOperation' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'equipment.SwitchOperation' resources.
type EquipmentSwitchOperationResponse struct {
	EquipmentSwitchOperationList *EquipmentSwitchOperationList
	MoAggregateTransform         *MoAggregateTransform
	MoDocumentCount              *MoDocumentCount
	MoTagSummary                 *MoTagSummary
}

// EquipmentSwitchOperationListAsEquipmentSwitchOperationResponse is a convenience function that returns EquipmentSwitchOperationList wrapped in EquipmentSwitchOperationResponse
func EquipmentSwitchOperationListAsEquipmentSwitchOperationResponse(v *EquipmentSwitchOperationList) EquipmentSwitchOperationResponse {
	return EquipmentSwitchOperationResponse{
		EquipmentSwitchOperationList: v,
	}
}

// MoAggregateTransformAsEquipmentSwitchOperationResponse is a convenience function that returns MoAggregateTransform wrapped in EquipmentSwitchOperationResponse
func MoAggregateTransformAsEquipmentSwitchOperationResponse(v *MoAggregateTransform) EquipmentSwitchOperationResponse {
	return EquipmentSwitchOperationResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsEquipmentSwitchOperationResponse is a convenience function that returns MoDocumentCount wrapped in EquipmentSwitchOperationResponse
func MoDocumentCountAsEquipmentSwitchOperationResponse(v *MoDocumentCount) EquipmentSwitchOperationResponse {
	return EquipmentSwitchOperationResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsEquipmentSwitchOperationResponse is a convenience function that returns MoTagSummary wrapped in EquipmentSwitchOperationResponse
func MoTagSummaryAsEquipmentSwitchOperationResponse(v *MoTagSummary) EquipmentSwitchOperationResponse {
	return EquipmentSwitchOperationResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EquipmentSwitchOperationResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'equipment.SwitchOperation.List'
	if jsonDict["ObjectType"] == "equipment.SwitchOperation.List" {
		// try to unmarshal JSON data into EquipmentSwitchOperationList
		err = json.Unmarshal(data, &dst.EquipmentSwitchOperationList)
		if err == nil {
			return nil // data stored in dst.EquipmentSwitchOperationList, return on the first match
		} else {
			dst.EquipmentSwitchOperationList = nil
			return fmt.Errorf("Failed to unmarshal EquipmentSwitchOperationResponse as EquipmentSwitchOperationList: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal EquipmentSwitchOperationResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal EquipmentSwitchOperationResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal EquipmentSwitchOperationResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EquipmentSwitchOperationResponse) MarshalJSON() ([]byte, error) {
	if src.EquipmentSwitchOperationList != nil {
		return json.Marshal(&src.EquipmentSwitchOperationList)
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
func (obj *EquipmentSwitchOperationResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EquipmentSwitchOperationList != nil {
		return obj.EquipmentSwitchOperationList
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

type NullableEquipmentSwitchOperationResponse struct {
	value *EquipmentSwitchOperationResponse
	isSet bool
}

func (v NullableEquipmentSwitchOperationResponse) Get() *EquipmentSwitchOperationResponse {
	return v.value
}

func (v *NullableEquipmentSwitchOperationResponse) Set(val *EquipmentSwitchOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSwitchOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSwitchOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSwitchOperationResponse(val *EquipmentSwitchOperationResponse) *NullableEquipmentSwitchOperationResponse {
	return &NullableEquipmentSwitchOperationResponse{value: val, isSet: true}
}

func (v NullableEquipmentSwitchOperationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSwitchOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}