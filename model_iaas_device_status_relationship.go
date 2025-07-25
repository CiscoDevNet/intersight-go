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

// IaasDeviceStatusRelationship - A relationship to the 'iaas.DeviceStatus' resource, or the expanded 'iaas.DeviceStatus' resource, or the 'null' value.
type IaasDeviceStatusRelationship struct {
	IaasDeviceStatus *IaasDeviceStatus
	MoMoRef          *MoMoRef
}

// IaasDeviceStatusAsIaasDeviceStatusRelationship is a convenience function that returns IaasDeviceStatus wrapped in IaasDeviceStatusRelationship
func IaasDeviceStatusAsIaasDeviceStatusRelationship(v *IaasDeviceStatus) IaasDeviceStatusRelationship {
	return IaasDeviceStatusRelationship{
		IaasDeviceStatus: v,
	}
}

// MoMoRefAsIaasDeviceStatusRelationship is a convenience function that returns MoMoRef wrapped in IaasDeviceStatusRelationship
func MoMoRefAsIaasDeviceStatusRelationship(v *MoMoRef) IaasDeviceStatusRelationship {
	return IaasDeviceStatusRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IaasDeviceStatusRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'iaas.DeviceStatus'
	if jsonDict["ClassId"] == "iaas.DeviceStatus" {
		// try to unmarshal JSON data into IaasDeviceStatus
		err = json.Unmarshal(data, &dst.IaasDeviceStatus)
		if err == nil {
			return nil // data stored in dst.IaasDeviceStatus, return on the first match
		} else {
			dst.IaasDeviceStatus = nil
			return fmt.Errorf("failed to unmarshal IaasDeviceStatusRelationship as IaasDeviceStatus: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IaasDeviceStatusRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IaasDeviceStatusRelationship) MarshalJSON() ([]byte, error) {
	if src.IaasDeviceStatus != nil {
		return json.Marshal(&src.IaasDeviceStatus)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IaasDeviceStatusRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IaasDeviceStatus != nil {
		return obj.IaasDeviceStatus
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableIaasDeviceStatusRelationship struct {
	value *IaasDeviceStatusRelationship
	isSet bool
}

func (v NullableIaasDeviceStatusRelationship) Get() *IaasDeviceStatusRelationship {
	return v.value
}

func (v *NullableIaasDeviceStatusRelationship) Set(val *IaasDeviceStatusRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasDeviceStatusRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasDeviceStatusRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasDeviceStatusRelationship(val *IaasDeviceStatusRelationship) *NullableIaasDeviceStatusRelationship {
	return &NullableIaasDeviceStatusRelationship{value: val, isSet: true}
}

func (v NullableIaasDeviceStatusRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasDeviceStatusRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
