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

// EquipmentIoCardRelationship - A relationship to the 'equipment.IoCard' resource, or the expanded 'equipment.IoCard' resource, or the 'null' value.
type EquipmentIoCardRelationship struct {
	EquipmentIoCard *EquipmentIoCard
	MoMoRef         *MoMoRef
}

// EquipmentIoCardAsEquipmentIoCardRelationship is a convenience function that returns EquipmentIoCard wrapped in EquipmentIoCardRelationship
func EquipmentIoCardAsEquipmentIoCardRelationship(v *EquipmentIoCard) EquipmentIoCardRelationship {
	return EquipmentIoCardRelationship{
		EquipmentIoCard: v,
	}
}

// MoMoRefAsEquipmentIoCardRelationship is a convenience function that returns MoMoRef wrapped in EquipmentIoCardRelationship
func MoMoRefAsEquipmentIoCardRelationship(v *MoMoRef) EquipmentIoCardRelationship {
	return EquipmentIoCardRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EquipmentIoCardRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'equipment.IoCard'
	if jsonDict["ClassId"] == "equipment.IoCard" {
		// try to unmarshal JSON data into EquipmentIoCard
		err = json.Unmarshal(data, &dst.EquipmentIoCard)
		if err == nil {
			return nil // data stored in dst.EquipmentIoCard, return on the first match
		} else {
			dst.EquipmentIoCard = nil
			return fmt.Errorf("failed to unmarshal EquipmentIoCardRelationship as EquipmentIoCard: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal EquipmentIoCardRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EquipmentIoCardRelationship) MarshalJSON() ([]byte, error) {
	if src.EquipmentIoCard != nil {
		return json.Marshal(&src.EquipmentIoCard)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EquipmentIoCardRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EquipmentIoCard != nil {
		return obj.EquipmentIoCard
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableEquipmentIoCardRelationship struct {
	value *EquipmentIoCardRelationship
	isSet bool
}

func (v NullableEquipmentIoCardRelationship) Get() *EquipmentIoCardRelationship {
	return v.value
}

func (v *NullableEquipmentIoCardRelationship) Set(val *EquipmentIoCardRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCardRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCardRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCardRelationship(val *EquipmentIoCardRelationship) *NullableEquipmentIoCardRelationship {
	return &NullableEquipmentIoCardRelationship{value: val, isSet: true}
}

func (v NullableEquipmentIoCardRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCardRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
