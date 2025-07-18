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

// VirtualizationVmwarePhysicalNetworkInterfaceRelationship - A relationship to the 'virtualization.VmwarePhysicalNetworkInterface' resource, or the expanded 'virtualization.VmwarePhysicalNetworkInterface' resource, or the 'null' value.
type VirtualizationVmwarePhysicalNetworkInterfaceRelationship struct {
	MoMoRef                                      *MoMoRef
	VirtualizationVmwarePhysicalNetworkInterface *VirtualizationVmwarePhysicalNetworkInterface
}

// MoMoRefAsVirtualizationVmwarePhysicalNetworkInterfaceRelationship is a convenience function that returns MoMoRef wrapped in VirtualizationVmwarePhysicalNetworkInterfaceRelationship
func MoMoRefAsVirtualizationVmwarePhysicalNetworkInterfaceRelationship(v *MoMoRef) VirtualizationVmwarePhysicalNetworkInterfaceRelationship {
	return VirtualizationVmwarePhysicalNetworkInterfaceRelationship{
		MoMoRef: v,
	}
}

// VirtualizationVmwarePhysicalNetworkInterfaceAsVirtualizationVmwarePhysicalNetworkInterfaceRelationship is a convenience function that returns VirtualizationVmwarePhysicalNetworkInterface wrapped in VirtualizationVmwarePhysicalNetworkInterfaceRelationship
func VirtualizationVmwarePhysicalNetworkInterfaceAsVirtualizationVmwarePhysicalNetworkInterfaceRelationship(v *VirtualizationVmwarePhysicalNetworkInterface) VirtualizationVmwarePhysicalNetworkInterfaceRelationship {
	return VirtualizationVmwarePhysicalNetworkInterfaceRelationship{
		VirtualizationVmwarePhysicalNetworkInterface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VirtualizationVmwarePhysicalNetworkInterfaceRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal VirtualizationVmwarePhysicalNetworkInterfaceRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'virtualization.VmwarePhysicalNetworkInterface'
	if jsonDict["ClassId"] == "virtualization.VmwarePhysicalNetworkInterface" {
		// try to unmarshal JSON data into VirtualizationVmwarePhysicalNetworkInterface
		err = json.Unmarshal(data, &dst.VirtualizationVmwarePhysicalNetworkInterface)
		if err == nil {
			return nil // data stored in dst.VirtualizationVmwarePhysicalNetworkInterface, return on the first match
		} else {
			dst.VirtualizationVmwarePhysicalNetworkInterface = nil
			return fmt.Errorf("failed to unmarshal VirtualizationVmwarePhysicalNetworkInterfaceRelationship as VirtualizationVmwarePhysicalNetworkInterface: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VirtualizationVmwarePhysicalNetworkInterfaceRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.VirtualizationVmwarePhysicalNetworkInterface != nil {
		return json.Marshal(&src.VirtualizationVmwarePhysicalNetworkInterface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VirtualizationVmwarePhysicalNetworkInterfaceRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.VirtualizationVmwarePhysicalNetworkInterface != nil {
		return obj.VirtualizationVmwarePhysicalNetworkInterface
	}

	// all schemas are nil
	return nil
}

type NullableVirtualizationVmwarePhysicalNetworkInterfaceRelationship struct {
	value *VirtualizationVmwarePhysicalNetworkInterfaceRelationship
	isSet bool
}

func (v NullableVirtualizationVmwarePhysicalNetworkInterfaceRelationship) Get() *VirtualizationVmwarePhysicalNetworkInterfaceRelationship {
	return v.value
}

func (v *NullableVirtualizationVmwarePhysicalNetworkInterfaceRelationship) Set(val *VirtualizationVmwarePhysicalNetworkInterfaceRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwarePhysicalNetworkInterfaceRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwarePhysicalNetworkInterfaceRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwarePhysicalNetworkInterfaceRelationship(val *VirtualizationVmwarePhysicalNetworkInterfaceRelationship) *NullableVirtualizationVmwarePhysicalNetworkInterfaceRelationship {
	return &NullableVirtualizationVmwarePhysicalNetworkInterfaceRelationship{value: val, isSet: true}
}

func (v NullableVirtualizationVmwarePhysicalNetworkInterfaceRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwarePhysicalNetworkInterfaceRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
