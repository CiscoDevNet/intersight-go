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

// InventoryGenericInventoryHolderRelationship - A relationship to the 'inventory.GenericInventoryHolder' resource, or the expanded 'inventory.GenericInventoryHolder' resource, or the 'null' value.
type InventoryGenericInventoryHolderRelationship struct {
	InventoryGenericInventoryHolder *InventoryGenericInventoryHolder
	MoMoRef                         *MoMoRef
}

// InventoryGenericInventoryHolderAsInventoryGenericInventoryHolderRelationship is a convenience function that returns InventoryGenericInventoryHolder wrapped in InventoryGenericInventoryHolderRelationship
func InventoryGenericInventoryHolderAsInventoryGenericInventoryHolderRelationship(v *InventoryGenericInventoryHolder) InventoryGenericInventoryHolderRelationship {
	return InventoryGenericInventoryHolderRelationship{
		InventoryGenericInventoryHolder: v,
	}
}

// MoMoRefAsInventoryGenericInventoryHolderRelationship is a convenience function that returns MoMoRef wrapped in InventoryGenericInventoryHolderRelationship
func MoMoRefAsInventoryGenericInventoryHolderRelationship(v *MoMoRef) InventoryGenericInventoryHolderRelationship {
	return InventoryGenericInventoryHolderRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *InventoryGenericInventoryHolderRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'inventory.GenericInventoryHolder'
	if jsonDict["ClassId"] == "inventory.GenericInventoryHolder" {
		// try to unmarshal JSON data into InventoryGenericInventoryHolder
		err = json.Unmarshal(data, &dst.InventoryGenericInventoryHolder)
		if err == nil {
			return nil // data stored in dst.InventoryGenericInventoryHolder, return on the first match
		} else {
			dst.InventoryGenericInventoryHolder = nil
			return fmt.Errorf("failed to unmarshal InventoryGenericInventoryHolderRelationship as InventoryGenericInventoryHolder: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal InventoryGenericInventoryHolderRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InventoryGenericInventoryHolderRelationship) MarshalJSON() ([]byte, error) {
	if src.InventoryGenericInventoryHolder != nil {
		return json.Marshal(&src.InventoryGenericInventoryHolder)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InventoryGenericInventoryHolderRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.InventoryGenericInventoryHolder != nil {
		return obj.InventoryGenericInventoryHolder
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableInventoryGenericInventoryHolderRelationship struct {
	value *InventoryGenericInventoryHolderRelationship
	isSet bool
}

func (v NullableInventoryGenericInventoryHolderRelationship) Get() *InventoryGenericInventoryHolderRelationship {
	return v.value
}

func (v *NullableInventoryGenericInventoryHolderRelationship) Set(val *InventoryGenericInventoryHolderRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryGenericInventoryHolderRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryGenericInventoryHolderRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryGenericInventoryHolderRelationship(val *InventoryGenericInventoryHolderRelationship) *NullableInventoryGenericInventoryHolderRelationship {
	return &NullableInventoryGenericInventoryHolderRelationship{value: val, isSet: true}
}

func (v NullableInventoryGenericInventoryHolderRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryGenericInventoryHolderRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
