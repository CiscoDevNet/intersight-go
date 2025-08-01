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

// BiosUnitRelationship - A relationship to the 'bios.Unit' resource, or the expanded 'bios.Unit' resource, or the 'null' value.
type BiosUnitRelationship struct {
	BiosUnit *BiosUnit
	MoMoRef  *MoMoRef
}

// BiosUnitAsBiosUnitRelationship is a convenience function that returns BiosUnit wrapped in BiosUnitRelationship
func BiosUnitAsBiosUnitRelationship(v *BiosUnit) BiosUnitRelationship {
	return BiosUnitRelationship{
		BiosUnit: v,
	}
}

// MoMoRefAsBiosUnitRelationship is a convenience function that returns MoMoRef wrapped in BiosUnitRelationship
func MoMoRefAsBiosUnitRelationship(v *MoMoRef) BiosUnitRelationship {
	return BiosUnitRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BiosUnitRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'bios.Unit'
	if jsonDict["ClassId"] == "bios.Unit" {
		// try to unmarshal JSON data into BiosUnit
		err = json.Unmarshal(data, &dst.BiosUnit)
		if err == nil {
			return nil // data stored in dst.BiosUnit, return on the first match
		} else {
			dst.BiosUnit = nil
			return fmt.Errorf("failed to unmarshal BiosUnitRelationship as BiosUnit: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal BiosUnitRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BiosUnitRelationship) MarshalJSON() ([]byte, error) {
	if src.BiosUnit != nil {
		return json.Marshal(&src.BiosUnit)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BiosUnitRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BiosUnit != nil {
		return obj.BiosUnit
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableBiosUnitRelationship struct {
	value *BiosUnitRelationship
	isSet bool
}

func (v NullableBiosUnitRelationship) Get() *BiosUnitRelationship {
	return v.value
}

func (v *NullableBiosUnitRelationship) Set(val *BiosUnitRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosUnitRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosUnitRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosUnitRelationship(val *BiosUnitRelationship) *NullableBiosUnitRelationship {
	return &NullableBiosUnitRelationship{value: val, isSet: true}
}

func (v NullableBiosUnitRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosUnitRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
