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

// HyperflexServerFirmwareVersionEntryRelationship - A relationship to the 'hyperflex.ServerFirmwareVersionEntry' resource, or the expanded 'hyperflex.ServerFirmwareVersionEntry' resource, or the 'null' value.
type HyperflexServerFirmwareVersionEntryRelationship struct {
	HyperflexServerFirmwareVersionEntry *HyperflexServerFirmwareVersionEntry
	MoMoRef                             *MoMoRef
}

// HyperflexServerFirmwareVersionEntryAsHyperflexServerFirmwareVersionEntryRelationship is a convenience function that returns HyperflexServerFirmwareVersionEntry wrapped in HyperflexServerFirmwareVersionEntryRelationship
func HyperflexServerFirmwareVersionEntryAsHyperflexServerFirmwareVersionEntryRelationship(v *HyperflexServerFirmwareVersionEntry) HyperflexServerFirmwareVersionEntryRelationship {
	return HyperflexServerFirmwareVersionEntryRelationship{
		HyperflexServerFirmwareVersionEntry: v,
	}
}

// MoMoRefAsHyperflexServerFirmwareVersionEntryRelationship is a convenience function that returns MoMoRef wrapped in HyperflexServerFirmwareVersionEntryRelationship
func MoMoRefAsHyperflexServerFirmwareVersionEntryRelationship(v *MoMoRef) HyperflexServerFirmwareVersionEntryRelationship {
	return HyperflexServerFirmwareVersionEntryRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexServerFirmwareVersionEntryRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hyperflex.ServerFirmwareVersionEntry'
	if jsonDict["ClassId"] == "hyperflex.ServerFirmwareVersionEntry" {
		// try to unmarshal JSON data into HyperflexServerFirmwareVersionEntry
		err = json.Unmarshal(data, &dst.HyperflexServerFirmwareVersionEntry)
		if err == nil {
			return nil // data stored in dst.HyperflexServerFirmwareVersionEntry, return on the first match
		} else {
			dst.HyperflexServerFirmwareVersionEntry = nil
			return fmt.Errorf("failed to unmarshal HyperflexServerFirmwareVersionEntryRelationship as HyperflexServerFirmwareVersionEntry: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexServerFirmwareVersionEntryRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexServerFirmwareVersionEntryRelationship) MarshalJSON() ([]byte, error) {
	if src.HyperflexServerFirmwareVersionEntry != nil {
		return json.Marshal(&src.HyperflexServerFirmwareVersionEntry)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HyperflexServerFirmwareVersionEntryRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HyperflexServerFirmwareVersionEntry != nil {
		return obj.HyperflexServerFirmwareVersionEntry
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHyperflexServerFirmwareVersionEntryRelationship struct {
	value *HyperflexServerFirmwareVersionEntryRelationship
	isSet bool
}

func (v NullableHyperflexServerFirmwareVersionEntryRelationship) Get() *HyperflexServerFirmwareVersionEntryRelationship {
	return v.value
}

func (v *NullableHyperflexServerFirmwareVersionEntryRelationship) Set(val *HyperflexServerFirmwareVersionEntryRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServerFirmwareVersionEntryRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServerFirmwareVersionEntryRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServerFirmwareVersionEntryRelationship(val *HyperflexServerFirmwareVersionEntryRelationship) *NullableHyperflexServerFirmwareVersionEntryRelationship {
	return &NullableHyperflexServerFirmwareVersionEntryRelationship{value: val, isSet: true}
}

func (v NullableHyperflexServerFirmwareVersionEntryRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServerFirmwareVersionEntryRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
