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

// WorkflowServiceItemInstanceRelationship - A relationship to the 'workflow.ServiceItemInstance' resource, or the expanded 'workflow.ServiceItemInstance' resource, or the 'null' value.
type WorkflowServiceItemInstanceRelationship struct {
	MoMoRef                     *MoMoRef
	WorkflowServiceItemInstance *WorkflowServiceItemInstance
}

// MoMoRefAsWorkflowServiceItemInstanceRelationship is a convenience function that returns MoMoRef wrapped in WorkflowServiceItemInstanceRelationship
func MoMoRefAsWorkflowServiceItemInstanceRelationship(v *MoMoRef) WorkflowServiceItemInstanceRelationship {
	return WorkflowServiceItemInstanceRelationship{
		MoMoRef: v,
	}
}

// WorkflowServiceItemInstanceAsWorkflowServiceItemInstanceRelationship is a convenience function that returns WorkflowServiceItemInstance wrapped in WorkflowServiceItemInstanceRelationship
func WorkflowServiceItemInstanceAsWorkflowServiceItemInstanceRelationship(v *WorkflowServiceItemInstance) WorkflowServiceItemInstanceRelationship {
	return WorkflowServiceItemInstanceRelationship{
		WorkflowServiceItemInstance: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *WorkflowServiceItemInstanceRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal WorkflowServiceItemInstanceRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'workflow.ServiceItemInstance'
	if jsonDict["ClassId"] == "workflow.ServiceItemInstance" {
		// try to unmarshal JSON data into WorkflowServiceItemInstance
		err = json.Unmarshal(data, &dst.WorkflowServiceItemInstance)
		if err == nil {
			return nil // data stored in dst.WorkflowServiceItemInstance, return on the first match
		} else {
			dst.WorkflowServiceItemInstance = nil
			return fmt.Errorf("failed to unmarshal WorkflowServiceItemInstanceRelationship as WorkflowServiceItemInstance: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WorkflowServiceItemInstanceRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.WorkflowServiceItemInstance != nil {
		return json.Marshal(&src.WorkflowServiceItemInstance)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WorkflowServiceItemInstanceRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.WorkflowServiceItemInstance != nil {
		return obj.WorkflowServiceItemInstance
	}

	// all schemas are nil
	return nil
}

type NullableWorkflowServiceItemInstanceRelationship struct {
	value *WorkflowServiceItemInstanceRelationship
	isSet bool
}

func (v NullableWorkflowServiceItemInstanceRelationship) Get() *WorkflowServiceItemInstanceRelationship {
	return v.value
}

func (v *NullableWorkflowServiceItemInstanceRelationship) Set(val *WorkflowServiceItemInstanceRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemInstanceRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemInstanceRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemInstanceRelationship(val *WorkflowServiceItemInstanceRelationship) *NullableWorkflowServiceItemInstanceRelationship {
	return &NullableWorkflowServiceItemInstanceRelationship{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemInstanceRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemInstanceRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
