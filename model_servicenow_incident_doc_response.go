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

// ServicenowIncidentDocResponse - The response body of a HTTP GET request for the 'servicenow.IncidentDoc' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'servicenow.IncidentDoc' resources.
type ServicenowIncidentDocResponse struct {
	MoAggregateTransform      *MoAggregateTransform
	MoDocumentCount           *MoDocumentCount
	MoTagSummary              *MoTagSummary
	ServicenowIncidentDocList *ServicenowIncidentDocList
}

// MoAggregateTransformAsServicenowIncidentDocResponse is a convenience function that returns MoAggregateTransform wrapped in ServicenowIncidentDocResponse
func MoAggregateTransformAsServicenowIncidentDocResponse(v *MoAggregateTransform) ServicenowIncidentDocResponse {
	return ServicenowIncidentDocResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsServicenowIncidentDocResponse is a convenience function that returns MoDocumentCount wrapped in ServicenowIncidentDocResponse
func MoDocumentCountAsServicenowIncidentDocResponse(v *MoDocumentCount) ServicenowIncidentDocResponse {
	return ServicenowIncidentDocResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsServicenowIncidentDocResponse is a convenience function that returns MoTagSummary wrapped in ServicenowIncidentDocResponse
func MoTagSummaryAsServicenowIncidentDocResponse(v *MoTagSummary) ServicenowIncidentDocResponse {
	return ServicenowIncidentDocResponse{
		MoTagSummary: v,
	}
}

// ServicenowIncidentDocListAsServicenowIncidentDocResponse is a convenience function that returns ServicenowIncidentDocList wrapped in ServicenowIncidentDocResponse
func ServicenowIncidentDocListAsServicenowIncidentDocResponse(v *ServicenowIncidentDocList) ServicenowIncidentDocResponse {
	return ServicenowIncidentDocResponse{
		ServicenowIncidentDocList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServicenowIncidentDocResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal ServicenowIncidentDocResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ServicenowIncidentDocResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ServicenowIncidentDocResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'servicenow.IncidentDoc.List'
	if jsonDict["ObjectType"] == "servicenow.IncidentDoc.List" {
		// try to unmarshal JSON data into ServicenowIncidentDocList
		err = json.Unmarshal(data, &dst.ServicenowIncidentDocList)
		if err == nil {
			return nil // data stored in dst.ServicenowIncidentDocList, return on the first match
		} else {
			dst.ServicenowIncidentDocList = nil
			return fmt.Errorf("failed to unmarshal ServicenowIncidentDocResponse as ServicenowIncidentDocList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServicenowIncidentDocResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.ServicenowIncidentDocList != nil {
		return json.Marshal(&src.ServicenowIncidentDocList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServicenowIncidentDocResponse) GetActualInstance() interface{} {
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

	if obj.ServicenowIncidentDocList != nil {
		return obj.ServicenowIncidentDocList
	}

	// all schemas are nil
	return nil
}

type NullableServicenowIncidentDocResponse struct {
	value *ServicenowIncidentDocResponse
	isSet bool
}

func (v NullableServicenowIncidentDocResponse) Get() *ServicenowIncidentDocResponse {
	return v.value
}

func (v *NullableServicenowIncidentDocResponse) Set(val *ServicenowIncidentDocResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServicenowIncidentDocResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServicenowIncidentDocResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicenowIncidentDocResponse(val *ServicenowIncidentDocResponse) *NullableServicenowIncidentDocResponse {
	return &NullableServicenowIncidentDocResponse{value: val, isSet: true}
}

func (v NullableServicenowIncidentDocResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicenowIncidentDocResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
