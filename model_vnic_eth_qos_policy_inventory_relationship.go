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

// VnicEthQosPolicyInventoryRelationship - A relationship to the 'vnic.EthQosPolicyInventory' resource, or the expanded 'vnic.EthQosPolicyInventory' resource, or the 'null' value.
type VnicEthQosPolicyInventoryRelationship struct {
	MoMoRef                   *MoMoRef
	VnicEthQosPolicyInventory *VnicEthQosPolicyInventory
}

// MoMoRefAsVnicEthQosPolicyInventoryRelationship is a convenience function that returns MoMoRef wrapped in VnicEthQosPolicyInventoryRelationship
func MoMoRefAsVnicEthQosPolicyInventoryRelationship(v *MoMoRef) VnicEthQosPolicyInventoryRelationship {
	return VnicEthQosPolicyInventoryRelationship{
		MoMoRef: v,
	}
}

// VnicEthQosPolicyInventoryAsVnicEthQosPolicyInventoryRelationship is a convenience function that returns VnicEthQosPolicyInventory wrapped in VnicEthQosPolicyInventoryRelationship
func VnicEthQosPolicyInventoryAsVnicEthQosPolicyInventoryRelationship(v *VnicEthQosPolicyInventory) VnicEthQosPolicyInventoryRelationship {
	return VnicEthQosPolicyInventoryRelationship{
		VnicEthQosPolicyInventory: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VnicEthQosPolicyInventoryRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal VnicEthQosPolicyInventoryRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'vnic.EthQosPolicyInventory'
	if jsonDict["ClassId"] == "vnic.EthQosPolicyInventory" {
		// try to unmarshal JSON data into VnicEthQosPolicyInventory
		err = json.Unmarshal(data, &dst.VnicEthQosPolicyInventory)
		if err == nil {
			return nil // data stored in dst.VnicEthQosPolicyInventory, return on the first match
		} else {
			dst.VnicEthQosPolicyInventory = nil
			return fmt.Errorf("failed to unmarshal VnicEthQosPolicyInventoryRelationship as VnicEthQosPolicyInventory: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VnicEthQosPolicyInventoryRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.VnicEthQosPolicyInventory != nil {
		return json.Marshal(&src.VnicEthQosPolicyInventory)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VnicEthQosPolicyInventoryRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.VnicEthQosPolicyInventory != nil {
		return obj.VnicEthQosPolicyInventory
	}

	// all schemas are nil
	return nil
}

type NullableVnicEthQosPolicyInventoryRelationship struct {
	value *VnicEthQosPolicyInventoryRelationship
	isSet bool
}

func (v NullableVnicEthQosPolicyInventoryRelationship) Get() *VnicEthQosPolicyInventoryRelationship {
	return v.value
}

func (v *NullableVnicEthQosPolicyInventoryRelationship) Set(val *VnicEthQosPolicyInventoryRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthQosPolicyInventoryRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthQosPolicyInventoryRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthQosPolicyInventoryRelationship(val *VnicEthQosPolicyInventoryRelationship) *NullableVnicEthQosPolicyInventoryRelationship {
	return &NullableVnicEthQosPolicyInventoryRelationship{value: val, isSet: true}
}

func (v NullableVnicEthQosPolicyInventoryRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthQosPolicyInventoryRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
