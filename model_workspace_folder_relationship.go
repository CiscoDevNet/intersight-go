/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-17769
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// WorkspaceFolderRelationship - A relationship to the 'workspace.Folder' resource, or the expanded 'workspace.Folder' resource, or the 'null' value.
type WorkspaceFolderRelationship struct {
	MoMoRef         *MoMoRef
	WorkspaceFolder *WorkspaceFolder
}

// MoMoRefAsWorkspaceFolderRelationship is a convenience function that returns MoMoRef wrapped in WorkspaceFolderRelationship
func MoMoRefAsWorkspaceFolderRelationship(v *MoMoRef) WorkspaceFolderRelationship {
	return WorkspaceFolderRelationship{
		MoMoRef: v,
	}
}

// WorkspaceFolderAsWorkspaceFolderRelationship is a convenience function that returns WorkspaceFolder wrapped in WorkspaceFolderRelationship
func WorkspaceFolderAsWorkspaceFolderRelationship(v *WorkspaceFolder) WorkspaceFolderRelationship {
	return WorkspaceFolderRelationship{
		WorkspaceFolder: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *WorkspaceFolderRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal WorkspaceFolderRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'workspace.Folder'
	if jsonDict["ClassId"] == "workspace.Folder" {
		// try to unmarshal JSON data into WorkspaceFolder
		err = json.Unmarshal(data, &dst.WorkspaceFolder)
		if err == nil {
			return nil // data stored in dst.WorkspaceFolder, return on the first match
		} else {
			dst.WorkspaceFolder = nil
			return fmt.Errorf("failed to unmarshal WorkspaceFolderRelationship as WorkspaceFolder: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WorkspaceFolderRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.WorkspaceFolder != nil {
		return json.Marshal(&src.WorkspaceFolder)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WorkspaceFolderRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.WorkspaceFolder != nil {
		return obj.WorkspaceFolder
	}

	// all schemas are nil
	return nil
}

type NullableWorkspaceFolderRelationship struct {
	value *WorkspaceFolderRelationship
	isSet bool
}

func (v NullableWorkspaceFolderRelationship) Get() *WorkspaceFolderRelationship {
	return v.value
}

func (v *NullableWorkspaceFolderRelationship) Set(val *WorkspaceFolderRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceFolderRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceFolderRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceFolderRelationship(val *WorkspaceFolderRelationship) *NullableWorkspaceFolderRelationship {
	return &NullableWorkspaceFolderRelationship{value: val, isSet: true}
}

func (v NullableWorkspaceFolderRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceFolderRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}