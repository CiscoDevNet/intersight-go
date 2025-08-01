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

// WorkflowTaskInfoRelationship - A relationship to the 'workflow.TaskInfo' resource, or the expanded 'workflow.TaskInfo' resource, or the 'null' value.
type WorkflowTaskInfoRelationship struct {
	MoMoRef          *MoMoRef
	WorkflowTaskInfo *WorkflowTaskInfo
}

// MoMoRefAsWorkflowTaskInfoRelationship is a convenience function that returns MoMoRef wrapped in WorkflowTaskInfoRelationship
func MoMoRefAsWorkflowTaskInfoRelationship(v *MoMoRef) WorkflowTaskInfoRelationship {
	return WorkflowTaskInfoRelationship{
		MoMoRef: v,
	}
}

// WorkflowTaskInfoAsWorkflowTaskInfoRelationship is a convenience function that returns WorkflowTaskInfo wrapped in WorkflowTaskInfoRelationship
func WorkflowTaskInfoAsWorkflowTaskInfoRelationship(v *WorkflowTaskInfo) WorkflowTaskInfoRelationship {
	return WorkflowTaskInfoRelationship{
		WorkflowTaskInfo: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *WorkflowTaskInfoRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("failed to unmarshal WorkflowTaskInfoRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'workflow.TaskInfo'
	if jsonDict["ClassId"] == "workflow.TaskInfo" {
		// try to unmarshal JSON data into WorkflowTaskInfo
		err = json.Unmarshal(data, &dst.WorkflowTaskInfo)
		if err == nil {
			return nil // data stored in dst.WorkflowTaskInfo, return on the first match
		} else {
			dst.WorkflowTaskInfo = nil
			return fmt.Errorf("failed to unmarshal WorkflowTaskInfoRelationship as WorkflowTaskInfo: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WorkflowTaskInfoRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.WorkflowTaskInfo != nil {
		return json.Marshal(&src.WorkflowTaskInfo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WorkflowTaskInfoRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.WorkflowTaskInfo != nil {
		return obj.WorkflowTaskInfo
	}

	// all schemas are nil
	return nil
}

type NullableWorkflowTaskInfoRelationship struct {
	value *WorkflowTaskInfoRelationship
	isSet bool
}

func (v NullableWorkflowTaskInfoRelationship) Get() *WorkflowTaskInfoRelationship {
	return v.value
}

func (v *NullableWorkflowTaskInfoRelationship) Set(val *WorkflowTaskInfoRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskInfoRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskInfoRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskInfoRelationship(val *WorkflowTaskInfoRelationship) *NullableWorkflowTaskInfoRelationship {
	return &NullableWorkflowTaskInfoRelationship{value: val, isSet: true}
}

func (v NullableWorkflowTaskInfoRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskInfoRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
