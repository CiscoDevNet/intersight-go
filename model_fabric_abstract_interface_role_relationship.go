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

// FabricAbstractInterfaceRoleRelationship - A relationship to the 'fabric.AbstractInterfaceRole' resource, or the expanded 'fabric.AbstractInterfaceRole' resource, or the 'null' value.
type FabricAbstractInterfaceRoleRelationship struct {
	FabricAbstractInterfaceRole *FabricAbstractInterfaceRole
	MoMoRef                     *MoMoRef
}

// FabricAbstractInterfaceRoleAsFabricAbstractInterfaceRoleRelationship is a convenience function that returns FabricAbstractInterfaceRole wrapped in FabricAbstractInterfaceRoleRelationship
func FabricAbstractInterfaceRoleAsFabricAbstractInterfaceRoleRelationship(v *FabricAbstractInterfaceRole) FabricAbstractInterfaceRoleRelationship {
	return FabricAbstractInterfaceRoleRelationship{
		FabricAbstractInterfaceRole: v,
	}
}

// MoMoRefAsFabricAbstractInterfaceRoleRelationship is a convenience function that returns MoMoRef wrapped in FabricAbstractInterfaceRoleRelationship
func MoMoRefAsFabricAbstractInterfaceRoleRelationship(v *MoMoRef) FabricAbstractInterfaceRoleRelationship {
	return FabricAbstractInterfaceRoleRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FabricAbstractInterfaceRoleRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'fabric.AbstractInterfaceRole'
	if jsonDict["ClassId"] == "fabric.AbstractInterfaceRole" {
		// try to unmarshal JSON data into FabricAbstractInterfaceRole
		err = json.Unmarshal(data, &dst.FabricAbstractInterfaceRole)
		if err == nil {
			return nil // data stored in dst.FabricAbstractInterfaceRole, return on the first match
		} else {
			dst.FabricAbstractInterfaceRole = nil
			return fmt.Errorf("failed to unmarshal FabricAbstractInterfaceRoleRelationship as FabricAbstractInterfaceRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("failed to unmarshal FabricAbstractInterfaceRoleRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FabricAbstractInterfaceRoleRelationship) MarshalJSON() ([]byte, error) {
	if src.FabricAbstractInterfaceRole != nil {
		return json.Marshal(&src.FabricAbstractInterfaceRole)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FabricAbstractInterfaceRoleRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FabricAbstractInterfaceRole != nil {
		return obj.FabricAbstractInterfaceRole
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableFabricAbstractInterfaceRoleRelationship struct {
	value *FabricAbstractInterfaceRoleRelationship
	isSet bool
}

func (v NullableFabricAbstractInterfaceRoleRelationship) Get() *FabricAbstractInterfaceRoleRelationship {
	return v.value
}

func (v *NullableFabricAbstractInterfaceRoleRelationship) Set(val *FabricAbstractInterfaceRoleRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricAbstractInterfaceRoleRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricAbstractInterfaceRoleRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricAbstractInterfaceRoleRelationship(val *FabricAbstractInterfaceRoleRelationship) *NullableFabricAbstractInterfaceRoleRelationship {
	return &NullableFabricAbstractInterfaceRoleRelationship{value: val, isSet: true}
}

func (v NullableFabricAbstractInterfaceRoleRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricAbstractInterfaceRoleRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
