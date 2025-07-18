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

// HyperflexServerFirmwareVersionRelationship - A relationship to the 'hyperflex.ServerFirmwareVersion' resource, or the expanded 'hyperflex.ServerFirmwareVersion' resource, or the 'null' value.
type HyperflexServerFirmwareVersionRelationship struct {
	HyperflexServerFirmwareVersion *HyperflexServerFirmwareVersion
	MoMoRef                        *MoMoRef
}

// HyperflexServerFirmwareVersionAsHyperflexServerFirmwareVersionRelationship is a convenience function that returns HyperflexServerFirmwareVersion wrapped in HyperflexServerFirmwareVersionRelationship
func HyperflexServerFirmwareVersionAsHyperflexServerFirmwareVersionRelationship(v *HyperflexServerFirmwareVersion) HyperflexServerFirmwareVersionRelationship {
	return HyperflexServerFirmwareVersionRelationship{
		HyperflexServerFirmwareVersion: v,
	}
}

// MoMoRefAsHyperflexServerFirmwareVersionRelationship is a convenience function that returns MoMoRef wrapped in HyperflexServerFirmwareVersionRelationship
func MoMoRefAsHyperflexServerFirmwareVersionRelationship(v *MoMoRef) HyperflexServerFirmwareVersionRelationship {
	return HyperflexServerFirmwareVersionRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexServerFirmwareVersionRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hyperflex.ServerFirmwareVersion'
	if jsonDict["ClassId"] == "hyperflex.ServerFirmwareVersion" {
		// try to unmarshal JSON data into HyperflexServerFirmwareVersion
		err = json.Unmarshal(data, &dst.HyperflexServerFirmwareVersion)
		if err == nil {
			return nil // data stored in dst.HyperflexServerFirmwareVersion, return on the first match
		} else {
			dst.HyperflexServerFirmwareVersion = nil
			return fmt.Errorf("failed to unmarshal HyperflexServerFirmwareVersionRelationship as HyperflexServerFirmwareVersion: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexServerFirmwareVersionRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexServerFirmwareVersionRelationship) MarshalJSON() ([]byte, error) {
	if src.HyperflexServerFirmwareVersion != nil {
		return json.Marshal(&src.HyperflexServerFirmwareVersion)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HyperflexServerFirmwareVersionRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HyperflexServerFirmwareVersion != nil {
		return obj.HyperflexServerFirmwareVersion
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHyperflexServerFirmwareVersionRelationship struct {
	value *HyperflexServerFirmwareVersionRelationship
	isSet bool
}

func (v NullableHyperflexServerFirmwareVersionRelationship) Get() *HyperflexServerFirmwareVersionRelationship {
	return v.value
}

func (v *NullableHyperflexServerFirmwareVersionRelationship) Set(val *HyperflexServerFirmwareVersionRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServerFirmwareVersionRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServerFirmwareVersionRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServerFirmwareVersionRelationship(val *HyperflexServerFirmwareVersionRelationship) *NullableHyperflexServerFirmwareVersionRelationship {
	return &NullableHyperflexServerFirmwareVersionRelationship{value: val, isSet: true}
}

func (v NullableHyperflexServerFirmwareVersionRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServerFirmwareVersionRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
