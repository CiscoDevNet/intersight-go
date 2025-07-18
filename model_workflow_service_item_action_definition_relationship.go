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

// WorkflowServiceItemActionDefinitionRelationship - A relationship to the 'workflow.ServiceItemActionDefinition' resource, or the expanded 'workflow.ServiceItemActionDefinition' resource, or the 'null' value.
type WorkflowServiceItemActionDefinitionRelationship struct {
	MoMoRef                             *MoMoRef
	WorkflowServiceItemActionDefinition *WorkflowServiceItemActionDefinition
}

// MoMoRefAsWorkflowServiceItemActionDefinitionRelationship is a convenience function that returns MoMoRef wrapped in WorkflowServiceItemActionDefinitionRelationship
func MoMoRefAsWorkflowServiceItemActionDefinitionRelationship(v *MoMoRef) WorkflowServiceItemActionDefinitionRelationship {
	return WorkflowServiceItemActionDefinitionRelationship{
		MoMoRef: v,
	}
}

// WorkflowServiceItemActionDefinitionAsWorkflowServiceItemActionDefinitionRelationship is a convenience function that returns WorkflowServiceItemActionDefinition wrapped in WorkflowServiceItemActionDefinitionRelationship
func WorkflowServiceItemActionDefinitionAsWorkflowServiceItemActionDefinitionRelationship(v *WorkflowServiceItemActionDefinition) WorkflowServiceItemActionDefinitionRelationship {
	return WorkflowServiceItemActionDefinitionRelationship{
		WorkflowServiceItemActionDefinition: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *WorkflowServiceItemActionDefinitionRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal WorkflowServiceItemActionDefinitionRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'workflow.ServiceItemActionDefinition'
	if jsonDict["ClassId"] == "workflow.ServiceItemActionDefinition" {
		// try to unmarshal JSON data into WorkflowServiceItemActionDefinition
		err = json.Unmarshal(data, &dst.WorkflowServiceItemActionDefinition)
		if err == nil {
			return nil // data stored in dst.WorkflowServiceItemActionDefinition, return on the first match
		} else {
			dst.WorkflowServiceItemActionDefinition = nil
			return fmt.Errorf("failed to unmarshal WorkflowServiceItemActionDefinitionRelationship as WorkflowServiceItemActionDefinition: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WorkflowServiceItemActionDefinitionRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.WorkflowServiceItemActionDefinition != nil {
		return json.Marshal(&src.WorkflowServiceItemActionDefinition)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WorkflowServiceItemActionDefinitionRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.WorkflowServiceItemActionDefinition != nil {
		return obj.WorkflowServiceItemActionDefinition
	}

	// all schemas are nil
	return nil
}

type NullableWorkflowServiceItemActionDefinitionRelationship struct {
	value *WorkflowServiceItemActionDefinitionRelationship
	isSet bool
}

func (v NullableWorkflowServiceItemActionDefinitionRelationship) Get() *WorkflowServiceItemActionDefinitionRelationship {
	return v.value
}

func (v *NullableWorkflowServiceItemActionDefinitionRelationship) Set(val *WorkflowServiceItemActionDefinitionRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemActionDefinitionRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemActionDefinitionRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemActionDefinitionRelationship(val *WorkflowServiceItemActionDefinitionRelationship) *NullableWorkflowServiceItemActionDefinitionRelationship {
	return &NullableWorkflowServiceItemActionDefinitionRelationship{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemActionDefinitionRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemActionDefinitionRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
