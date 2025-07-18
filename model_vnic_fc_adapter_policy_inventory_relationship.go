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

// VnicFcAdapterPolicyInventoryRelationship - A relationship to the 'vnic.FcAdapterPolicyInventory' resource, or the expanded 'vnic.FcAdapterPolicyInventory' resource, or the 'null' value.
type VnicFcAdapterPolicyInventoryRelationship struct {
	MoMoRef                      *MoMoRef
	VnicFcAdapterPolicyInventory *VnicFcAdapterPolicyInventory
}

// MoMoRefAsVnicFcAdapterPolicyInventoryRelationship is a convenience function that returns MoMoRef wrapped in VnicFcAdapterPolicyInventoryRelationship
func MoMoRefAsVnicFcAdapterPolicyInventoryRelationship(v *MoMoRef) VnicFcAdapterPolicyInventoryRelationship {
	return VnicFcAdapterPolicyInventoryRelationship{
		MoMoRef: v,
	}
}

// VnicFcAdapterPolicyInventoryAsVnicFcAdapterPolicyInventoryRelationship is a convenience function that returns VnicFcAdapterPolicyInventory wrapped in VnicFcAdapterPolicyInventoryRelationship
func VnicFcAdapterPolicyInventoryAsVnicFcAdapterPolicyInventoryRelationship(v *VnicFcAdapterPolicyInventory) VnicFcAdapterPolicyInventoryRelationship {
	return VnicFcAdapterPolicyInventoryRelationship{
		VnicFcAdapterPolicyInventory: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VnicFcAdapterPolicyInventoryRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal VnicFcAdapterPolicyInventoryRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'vnic.FcAdapterPolicyInventory'
	if jsonDict["ClassId"] == "vnic.FcAdapterPolicyInventory" {
		// try to unmarshal JSON data into VnicFcAdapterPolicyInventory
		err = json.Unmarshal(data, &dst.VnicFcAdapterPolicyInventory)
		if err == nil {
			return nil // data stored in dst.VnicFcAdapterPolicyInventory, return on the first match
		} else {
			dst.VnicFcAdapterPolicyInventory = nil
			return fmt.Errorf("failed to unmarshal VnicFcAdapterPolicyInventoryRelationship as VnicFcAdapterPolicyInventory: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VnicFcAdapterPolicyInventoryRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.VnicFcAdapterPolicyInventory != nil {
		return json.Marshal(&src.VnicFcAdapterPolicyInventory)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VnicFcAdapterPolicyInventoryRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.VnicFcAdapterPolicyInventory != nil {
		return obj.VnicFcAdapterPolicyInventory
	}

	// all schemas are nil
	return nil
}

type NullableVnicFcAdapterPolicyInventoryRelationship struct {
	value *VnicFcAdapterPolicyInventoryRelationship
	isSet bool
}

func (v NullableVnicFcAdapterPolicyInventoryRelationship) Get() *VnicFcAdapterPolicyInventoryRelationship {
	return v.value
}

func (v *NullableVnicFcAdapterPolicyInventoryRelationship) Set(val *VnicFcAdapterPolicyInventoryRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcAdapterPolicyInventoryRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcAdapterPolicyInventoryRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcAdapterPolicyInventoryRelationship(val *VnicFcAdapterPolicyInventoryRelationship) *NullableVnicFcAdapterPolicyInventoryRelationship {
	return &NullableVnicFcAdapterPolicyInventoryRelationship{value: val, isSet: true}
}

func (v NullableVnicFcAdapterPolicyInventoryRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcAdapterPolicyInventoryRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
