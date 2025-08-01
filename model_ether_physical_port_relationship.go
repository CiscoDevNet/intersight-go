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

// EtherPhysicalPortRelationship - A relationship to the 'ether.PhysicalPort' resource, or the expanded 'ether.PhysicalPort' resource, or the 'null' value.
type EtherPhysicalPortRelationship struct {
	EtherPhysicalPort *EtherPhysicalPort
	MoMoRef           *MoMoRef
}

// EtherPhysicalPortAsEtherPhysicalPortRelationship is a convenience function that returns EtherPhysicalPort wrapped in EtherPhysicalPortRelationship
func EtherPhysicalPortAsEtherPhysicalPortRelationship(v *EtherPhysicalPort) EtherPhysicalPortRelationship {
	return EtherPhysicalPortRelationship{
		EtherPhysicalPort: v,
	}
}

// MoMoRefAsEtherPhysicalPortRelationship is a convenience function that returns MoMoRef wrapped in EtherPhysicalPortRelationship
func MoMoRefAsEtherPhysicalPortRelationship(v *MoMoRef) EtherPhysicalPortRelationship {
	return EtherPhysicalPortRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EtherPhysicalPortRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'ether.PhysicalPort'
	if jsonDict["ClassId"] == "ether.PhysicalPort" {
		// try to unmarshal JSON data into EtherPhysicalPort
		err = json.Unmarshal(data, &dst.EtherPhysicalPort)
		if err == nil {
			return nil // data stored in dst.EtherPhysicalPort, return on the first match
		} else {
			dst.EtherPhysicalPort = nil
			return fmt.Errorf("failed to unmarshal EtherPhysicalPortRelationship as EtherPhysicalPort: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal EtherPhysicalPortRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EtherPhysicalPortRelationship) MarshalJSON() ([]byte, error) {
	if src.EtherPhysicalPort != nil {
		return json.Marshal(&src.EtherPhysicalPort)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EtherPhysicalPortRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EtherPhysicalPort != nil {
		return obj.EtherPhysicalPort
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableEtherPhysicalPortRelationship struct {
	value *EtherPhysicalPortRelationship
	isSet bool
}

func (v NullableEtherPhysicalPortRelationship) Get() *EtherPhysicalPortRelationship {
	return v.value
}

func (v *NullableEtherPhysicalPortRelationship) Set(val *EtherPhysicalPortRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableEtherPhysicalPortRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableEtherPhysicalPortRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtherPhysicalPortRelationship(val *EtherPhysicalPortRelationship) *NullableEtherPhysicalPortRelationship {
	return &NullableEtherPhysicalPortRelationship{value: val, isSet: true}
}

func (v NullableEtherPhysicalPortRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtherPhysicalPortRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
