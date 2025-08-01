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

// EquipmentIoCardBaseRelationship - A relationship to the 'equipment.IoCardBase' resource, or the expanded 'equipment.IoCardBase' resource, or the 'null' value.
type EquipmentIoCardBaseRelationship struct {
	EquipmentIoCardBase *EquipmentIoCardBase
	MoMoRef             *MoMoRef
}

// EquipmentIoCardBaseAsEquipmentIoCardBaseRelationship is a convenience function that returns EquipmentIoCardBase wrapped in EquipmentIoCardBaseRelationship
func EquipmentIoCardBaseAsEquipmentIoCardBaseRelationship(v *EquipmentIoCardBase) EquipmentIoCardBaseRelationship {
	return EquipmentIoCardBaseRelationship{
		EquipmentIoCardBase: v,
	}
}

// MoMoRefAsEquipmentIoCardBaseRelationship is a convenience function that returns MoMoRef wrapped in EquipmentIoCardBaseRelationship
func MoMoRefAsEquipmentIoCardBaseRelationship(v *MoMoRef) EquipmentIoCardBaseRelationship {
	return EquipmentIoCardBaseRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EquipmentIoCardBaseRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'equipment.IoCardBase'
	if jsonDict["ClassId"] == "equipment.IoCardBase" {
		// try to unmarshal JSON data into EquipmentIoCardBase
		err = json.Unmarshal(data, &dst.EquipmentIoCardBase)
		if err == nil {
			return nil // data stored in dst.EquipmentIoCardBase, return on the first match
		} else {
			dst.EquipmentIoCardBase = nil
			return fmt.Errorf("failed to unmarshal EquipmentIoCardBaseRelationship as EquipmentIoCardBase: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal EquipmentIoCardBaseRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EquipmentIoCardBaseRelationship) MarshalJSON() ([]byte, error) {
	if src.EquipmentIoCardBase != nil {
		return json.Marshal(&src.EquipmentIoCardBase)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EquipmentIoCardBaseRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EquipmentIoCardBase != nil {
		return obj.EquipmentIoCardBase
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableEquipmentIoCardBaseRelationship struct {
	value *EquipmentIoCardBaseRelationship
	isSet bool
}

func (v NullableEquipmentIoCardBaseRelationship) Get() *EquipmentIoCardBaseRelationship {
	return v.value
}

func (v *NullableEquipmentIoCardBaseRelationship) Set(val *EquipmentIoCardBaseRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCardBaseRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCardBaseRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCardBaseRelationship(val *EquipmentIoCardBaseRelationship) *NullableEquipmentIoCardBaseRelationship {
	return &NullableEquipmentIoCardBaseRelationship{value: val, isSet: true}
}

func (v NullableEquipmentIoCardBaseRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCardBaseRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
