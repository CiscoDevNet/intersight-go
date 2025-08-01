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

// BootPxeDeviceRelationship - A relationship to the 'boot.PxeDevice' resource, or the expanded 'boot.PxeDevice' resource, or the 'null' value.
type BootPxeDeviceRelationship struct {
	BootPxeDevice *BootPxeDevice
	MoMoRef       *MoMoRef
}

// BootPxeDeviceAsBootPxeDeviceRelationship is a convenience function that returns BootPxeDevice wrapped in BootPxeDeviceRelationship
func BootPxeDeviceAsBootPxeDeviceRelationship(v *BootPxeDevice) BootPxeDeviceRelationship {
	return BootPxeDeviceRelationship{
		BootPxeDevice: v,
	}
}

// MoMoRefAsBootPxeDeviceRelationship is a convenience function that returns MoMoRef wrapped in BootPxeDeviceRelationship
func MoMoRefAsBootPxeDeviceRelationship(v *MoMoRef) BootPxeDeviceRelationship {
	return BootPxeDeviceRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BootPxeDeviceRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'boot.PxeDevice'
	if jsonDict["ClassId"] == "boot.PxeDevice" {
		// try to unmarshal JSON data into BootPxeDevice
		err = json.Unmarshal(data, &dst.BootPxeDevice)
		if err == nil {
			return nil // data stored in dst.BootPxeDevice, return on the first match
		} else {
			dst.BootPxeDevice = nil
			return fmt.Errorf("failed to unmarshal BootPxeDeviceRelationship as BootPxeDevice: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal BootPxeDeviceRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BootPxeDeviceRelationship) MarshalJSON() ([]byte, error) {
	if src.BootPxeDevice != nil {
		return json.Marshal(&src.BootPxeDevice)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BootPxeDeviceRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BootPxeDevice != nil {
		return obj.BootPxeDevice
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableBootPxeDeviceRelationship struct {
	value *BootPxeDeviceRelationship
	isSet bool
}

func (v NullableBootPxeDeviceRelationship) Get() *BootPxeDeviceRelationship {
	return v.value
}

func (v *NullableBootPxeDeviceRelationship) Set(val *BootPxeDeviceRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableBootPxeDeviceRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableBootPxeDeviceRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootPxeDeviceRelationship(val *BootPxeDeviceRelationship) *NullableBootPxeDeviceRelationship {
	return &NullableBootPxeDeviceRelationship{value: val, isSet: true}
}

func (v NullableBootPxeDeviceRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootPxeDeviceRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
