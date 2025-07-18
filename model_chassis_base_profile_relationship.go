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

// ChassisBaseProfileRelationship - A relationship to the 'chassis.BaseProfile' resource, or the expanded 'chassis.BaseProfile' resource, or the 'null' value.
type ChassisBaseProfileRelationship struct {
	ChassisBaseProfile *ChassisBaseProfile
	MoMoRef            *MoMoRef
}

// ChassisBaseProfileAsChassisBaseProfileRelationship is a convenience function that returns ChassisBaseProfile wrapped in ChassisBaseProfileRelationship
func ChassisBaseProfileAsChassisBaseProfileRelationship(v *ChassisBaseProfile) ChassisBaseProfileRelationship {
	return ChassisBaseProfileRelationship{
		ChassisBaseProfile: v,
	}
}

// MoMoRefAsChassisBaseProfileRelationship is a convenience function that returns MoMoRef wrapped in ChassisBaseProfileRelationship
func MoMoRefAsChassisBaseProfileRelationship(v *MoMoRef) ChassisBaseProfileRelationship {
	return ChassisBaseProfileRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChassisBaseProfileRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'chassis.BaseProfile'
	if jsonDict["ClassId"] == "chassis.BaseProfile" {
		// try to unmarshal JSON data into ChassisBaseProfile
		err = json.Unmarshal(data, &dst.ChassisBaseProfile)
		if err == nil {
			return nil // data stored in dst.ChassisBaseProfile, return on the first match
		} else {
			dst.ChassisBaseProfile = nil
			return fmt.Errorf("failed to unmarshal ChassisBaseProfileRelationship as ChassisBaseProfile: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ChassisBaseProfileRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChassisBaseProfileRelationship) MarshalJSON() ([]byte, error) {
	if src.ChassisBaseProfile != nil {
		return json.Marshal(&src.ChassisBaseProfile)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ChassisBaseProfileRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ChassisBaseProfile != nil {
		return obj.ChassisBaseProfile
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableChassisBaseProfileRelationship struct {
	value *ChassisBaseProfileRelationship
	isSet bool
}

func (v NullableChassisBaseProfileRelationship) Get() *ChassisBaseProfileRelationship {
	return v.value
}

func (v *NullableChassisBaseProfileRelationship) Set(val *ChassisBaseProfileRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisBaseProfileRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisBaseProfileRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisBaseProfileRelationship(val *ChassisBaseProfileRelationship) *NullableChassisBaseProfileRelationship {
	return &NullableChassisBaseProfileRelationship{value: val, isSet: true}
}

func (v NullableChassisBaseProfileRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisBaseProfileRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
