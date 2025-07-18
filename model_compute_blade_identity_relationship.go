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

// ComputeBladeIdentityRelationship - A relationship to the 'compute.BladeIdentity' resource, or the expanded 'compute.BladeIdentity' resource, or the 'null' value.
type ComputeBladeIdentityRelationship struct {
	ComputeBladeIdentity *ComputeBladeIdentity
	MoMoRef              *MoMoRef
}

// ComputeBladeIdentityAsComputeBladeIdentityRelationship is a convenience function that returns ComputeBladeIdentity wrapped in ComputeBladeIdentityRelationship
func ComputeBladeIdentityAsComputeBladeIdentityRelationship(v *ComputeBladeIdentity) ComputeBladeIdentityRelationship {
	return ComputeBladeIdentityRelationship{
		ComputeBladeIdentity: v,
	}
}

// MoMoRefAsComputeBladeIdentityRelationship is a convenience function that returns MoMoRef wrapped in ComputeBladeIdentityRelationship
func MoMoRefAsComputeBladeIdentityRelationship(v *MoMoRef) ComputeBladeIdentityRelationship {
	return ComputeBladeIdentityRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ComputeBladeIdentityRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'compute.BladeIdentity'
	if jsonDict["ClassId"] == "compute.BladeIdentity" {
		// try to unmarshal JSON data into ComputeBladeIdentity
		err = json.Unmarshal(data, &dst.ComputeBladeIdentity)
		if err == nil {
			return nil // data stored in dst.ComputeBladeIdentity, return on the first match
		} else {
			dst.ComputeBladeIdentity = nil
			return fmt.Errorf("failed to unmarshal ComputeBladeIdentityRelationship as ComputeBladeIdentity: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ComputeBladeIdentityRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ComputeBladeIdentityRelationship) MarshalJSON() ([]byte, error) {
	if src.ComputeBladeIdentity != nil {
		return json.Marshal(&src.ComputeBladeIdentity)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ComputeBladeIdentityRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ComputeBladeIdentity != nil {
		return obj.ComputeBladeIdentity
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableComputeBladeIdentityRelationship struct {
	value *ComputeBladeIdentityRelationship
	isSet bool
}

func (v NullableComputeBladeIdentityRelationship) Get() *ComputeBladeIdentityRelationship {
	return v.value
}

func (v *NullableComputeBladeIdentityRelationship) Set(val *ComputeBladeIdentityRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeBladeIdentityRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeBladeIdentityRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeBladeIdentityRelationship(val *ComputeBladeIdentityRelationship) *NullableComputeBladeIdentityRelationship {
	return &NullableComputeBladeIdentityRelationship{value: val, isSet: true}
}

func (v NullableComputeBladeIdentityRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeBladeIdentityRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
