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

// NiatelemetryNiaInventoryRelationship - A relationship to the 'niatelemetry.NiaInventory' resource, or the expanded 'niatelemetry.NiaInventory' resource, or the 'null' value.
type NiatelemetryNiaInventoryRelationship struct {
	MoMoRef                  *MoMoRef
	NiatelemetryNiaInventory *NiatelemetryNiaInventory
}

// MoMoRefAsNiatelemetryNiaInventoryRelationship is a convenience function that returns MoMoRef wrapped in NiatelemetryNiaInventoryRelationship
func MoMoRefAsNiatelemetryNiaInventoryRelationship(v *MoMoRef) NiatelemetryNiaInventoryRelationship {
	return NiatelemetryNiaInventoryRelationship{
		MoMoRef: v,
	}
}

// NiatelemetryNiaInventoryAsNiatelemetryNiaInventoryRelationship is a convenience function that returns NiatelemetryNiaInventory wrapped in NiatelemetryNiaInventoryRelationship
func NiatelemetryNiaInventoryAsNiatelemetryNiaInventoryRelationship(v *NiatelemetryNiaInventory) NiatelemetryNiaInventoryRelationship {
	return NiatelemetryNiaInventoryRelationship{
		NiatelemetryNiaInventory: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiatelemetryNiaInventoryRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal NiatelemetryNiaInventoryRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niatelemetry.NiaInventory'
	if jsonDict["ClassId"] == "niatelemetry.NiaInventory" {
		// try to unmarshal JSON data into NiatelemetryNiaInventory
		err = json.Unmarshal(data, &dst.NiatelemetryNiaInventory)
		if err == nil {
			return nil // data stored in dst.NiatelemetryNiaInventory, return on the first match
		} else {
			dst.NiatelemetryNiaInventory = nil
			return fmt.Errorf("failed to unmarshal NiatelemetryNiaInventoryRelationship as NiatelemetryNiaInventory: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiatelemetryNiaInventoryRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.NiatelemetryNiaInventory != nil {
		return json.Marshal(&src.NiatelemetryNiaInventory)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiatelemetryNiaInventoryRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.NiatelemetryNiaInventory != nil {
		return obj.NiatelemetryNiaInventory
	}

	// all schemas are nil
	return nil
}

type NullableNiatelemetryNiaInventoryRelationship struct {
	value *NiatelemetryNiaInventoryRelationship
	isSet bool
}

func (v NullableNiatelemetryNiaInventoryRelationship) Get() *NiatelemetryNiaInventoryRelationship {
	return v.value
}

func (v *NullableNiatelemetryNiaInventoryRelationship) Set(val *NiatelemetryNiaInventoryRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNiaInventoryRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNiaInventoryRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNiaInventoryRelationship(val *NiatelemetryNiaInventoryRelationship) *NullableNiatelemetryNiaInventoryRelationship {
	return &NullableNiatelemetryNiaInventoryRelationship{value: val, isSet: true}
}

func (v NullableNiatelemetryNiaInventoryRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNiaInventoryRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
