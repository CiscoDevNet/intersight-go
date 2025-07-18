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

// HciGpuRelationship - A relationship to the 'hci.Gpu' resource, or the expanded 'hci.Gpu' resource, or the 'null' value.
type HciGpuRelationship struct {
	HciGpu  *HciGpu
	MoMoRef *MoMoRef
}

// HciGpuAsHciGpuRelationship is a convenience function that returns HciGpu wrapped in HciGpuRelationship
func HciGpuAsHciGpuRelationship(v *HciGpu) HciGpuRelationship {
	return HciGpuRelationship{
		HciGpu: v,
	}
}

// MoMoRefAsHciGpuRelationship is a convenience function that returns MoMoRef wrapped in HciGpuRelationship
func MoMoRefAsHciGpuRelationship(v *MoMoRef) HciGpuRelationship {
	return HciGpuRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HciGpuRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hci.Gpu'
	if jsonDict["ClassId"] == "hci.Gpu" {
		// try to unmarshal JSON data into HciGpu
		err = json.Unmarshal(data, &dst.HciGpu)
		if err == nil {
			return nil // data stored in dst.HciGpu, return on the first match
		} else {
			dst.HciGpu = nil
			return fmt.Errorf("failed to unmarshal HciGpuRelationship as HciGpu: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HciGpuRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HciGpuRelationship) MarshalJSON() ([]byte, error) {
	if src.HciGpu != nil {
		return json.Marshal(&src.HciGpu)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HciGpuRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HciGpu != nil {
		return obj.HciGpu
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHciGpuRelationship struct {
	value *HciGpuRelationship
	isSet bool
}

func (v NullableHciGpuRelationship) Get() *HciGpuRelationship {
	return v.value
}

func (v *NullableHciGpuRelationship) Set(val *HciGpuRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHciGpuRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHciGpuRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHciGpuRelationship(val *HciGpuRelationship) *NullableHciGpuRelationship {
	return &NullableHciGpuRelationship{value: val, isSet: true}
}

func (v NullableHciGpuRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHciGpuRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
