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

// MonitoringHealthStatusResponse - The response body of a HTTP GET request for the 'monitoring.HealthStatus' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'monitoring.HealthStatus' resources.
type MonitoringHealthStatusResponse struct {
	MoAggregateTransform       *MoAggregateTransform
	MoDocumentCount            *MoDocumentCount
	MoTagSummary               *MoTagSummary
	MonitoringHealthStatusList *MonitoringHealthStatusList
}

// MoAggregateTransformAsMonitoringHealthStatusResponse is a convenience function that returns MoAggregateTransform wrapped in MonitoringHealthStatusResponse
func MoAggregateTransformAsMonitoringHealthStatusResponse(v *MoAggregateTransform) MonitoringHealthStatusResponse {
	return MonitoringHealthStatusResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsMonitoringHealthStatusResponse is a convenience function that returns MoDocumentCount wrapped in MonitoringHealthStatusResponse
func MoDocumentCountAsMonitoringHealthStatusResponse(v *MoDocumentCount) MonitoringHealthStatusResponse {
	return MonitoringHealthStatusResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsMonitoringHealthStatusResponse is a convenience function that returns MoTagSummary wrapped in MonitoringHealthStatusResponse
func MoTagSummaryAsMonitoringHealthStatusResponse(v *MoTagSummary) MonitoringHealthStatusResponse {
	return MonitoringHealthStatusResponse{
		MoTagSummary: v,
	}
}

// MonitoringHealthStatusListAsMonitoringHealthStatusResponse is a convenience function that returns MonitoringHealthStatusList wrapped in MonitoringHealthStatusResponse
func MonitoringHealthStatusListAsMonitoringHealthStatusResponse(v *MonitoringHealthStatusList) MonitoringHealthStatusResponse {
	return MonitoringHealthStatusResponse{
		MonitoringHealthStatusList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MonitoringHealthStatusResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal MonitoringHealthStatusResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MonitoringHealthStatusResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MonitoringHealthStatusResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'monitoring.HealthStatus.List'
	if jsonDict["ObjectType"] == "monitoring.HealthStatus.List" {
		// try to unmarshal JSON data into MonitoringHealthStatusList
		err = json.Unmarshal(data, &dst.MonitoringHealthStatusList)
		if err == nil {
			return nil // data stored in dst.MonitoringHealthStatusList, return on the first match
		} else {
			dst.MonitoringHealthStatusList = nil
			return fmt.Errorf("failed to unmarshal MonitoringHealthStatusResponse as MonitoringHealthStatusList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MonitoringHealthStatusResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.MonitoringHealthStatusList != nil {
		return json.Marshal(&src.MonitoringHealthStatusList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MonitoringHealthStatusResponse) GetActualInstance() interface{} {
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

	if obj.MonitoringHealthStatusList != nil {
		return obj.MonitoringHealthStatusList
	}

	// all schemas are nil
	return nil
}

type NullableMonitoringHealthStatusResponse struct {
	value *MonitoringHealthStatusResponse
	isSet bool
}

func (v NullableMonitoringHealthStatusResponse) Get() *MonitoringHealthStatusResponse {
	return v.value
}

func (v *NullableMonitoringHealthStatusResponse) Set(val *MonitoringHealthStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringHealthStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringHealthStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringHealthStatusResponse(val *MonitoringHealthStatusResponse) *NullableMonitoringHealthStatusResponse {
	return &NullableMonitoringHealthStatusResponse{value: val, isSet: true}
}

func (v NullableMonitoringHealthStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringHealthStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
