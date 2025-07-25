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

// ComputeBladeRelationship - A relationship to the 'compute.Blade' resource, or the expanded 'compute.Blade' resource, or the 'null' value.
type ComputeBladeRelationship struct {
	ComputeBlade *ComputeBlade
	MoMoRef      *MoMoRef
}

// ComputeBladeAsComputeBladeRelationship is a convenience function that returns ComputeBlade wrapped in ComputeBladeRelationship
func ComputeBladeAsComputeBladeRelationship(v *ComputeBlade) ComputeBladeRelationship {
	return ComputeBladeRelationship{
		ComputeBlade: v,
	}
}

// MoMoRefAsComputeBladeRelationship is a convenience function that returns MoMoRef wrapped in ComputeBladeRelationship
func MoMoRefAsComputeBladeRelationship(v *MoMoRef) ComputeBladeRelationship {
	return ComputeBladeRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ComputeBladeRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'compute.Blade'
	if jsonDict["ClassId"] == "compute.Blade" {
		// try to unmarshal JSON data into ComputeBlade
		err = json.Unmarshal(data, &dst.ComputeBlade)
		if err == nil {
			return nil // data stored in dst.ComputeBlade, return on the first match
		} else {
			dst.ComputeBlade = nil
			return fmt.Errorf("failed to unmarshal ComputeBladeRelationship as ComputeBlade: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ComputeBladeRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ComputeBladeRelationship) MarshalJSON() ([]byte, error) {
	if src.ComputeBlade != nil {
		return json.Marshal(&src.ComputeBlade)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ComputeBladeRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ComputeBlade != nil {
		return obj.ComputeBlade
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableComputeBladeRelationship struct {
	value *ComputeBladeRelationship
	isSet bool
}

func (v NullableComputeBladeRelationship) Get() *ComputeBladeRelationship {
	return v.value
}

func (v *NullableComputeBladeRelationship) Set(val *ComputeBladeRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeBladeRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeBladeRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeBladeRelationship(val *ComputeBladeRelationship) *NullableComputeBladeRelationship {
	return &NullableComputeBladeRelationship{value: val, isSet: true}
}

func (v NullableComputeBladeRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeBladeRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
