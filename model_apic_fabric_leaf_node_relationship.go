/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-17769
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// ApicFabricLeafNodeRelationship - A relationship to the 'apic.FabricLeafNode' resource, or the expanded 'apic.FabricLeafNode' resource, or the 'null' value.
type ApicFabricLeafNodeRelationship struct {
	ApicFabricLeafNode *ApicFabricLeafNode
	MoMoRef            *MoMoRef
}

// ApicFabricLeafNodeAsApicFabricLeafNodeRelationship is a convenience function that returns ApicFabricLeafNode wrapped in ApicFabricLeafNodeRelationship
func ApicFabricLeafNodeAsApicFabricLeafNodeRelationship(v *ApicFabricLeafNode) ApicFabricLeafNodeRelationship {
	return ApicFabricLeafNodeRelationship{
		ApicFabricLeafNode: v,
	}
}

// MoMoRefAsApicFabricLeafNodeRelationship is a convenience function that returns MoMoRef wrapped in ApicFabricLeafNodeRelationship
func MoMoRefAsApicFabricLeafNodeRelationship(v *MoMoRef) ApicFabricLeafNodeRelationship {
	return ApicFabricLeafNodeRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApicFabricLeafNodeRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'apic.FabricLeafNode'
	if jsonDict["ClassId"] == "apic.FabricLeafNode" {
		// try to unmarshal JSON data into ApicFabricLeafNode
		err = json.Unmarshal(data, &dst.ApicFabricLeafNode)
		if err == nil {
			return nil // data stored in dst.ApicFabricLeafNode, return on the first match
		} else {
			dst.ApicFabricLeafNode = nil
			return fmt.Errorf("failed to unmarshal ApicFabricLeafNodeRelationship as ApicFabricLeafNode: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ApicFabricLeafNodeRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApicFabricLeafNodeRelationship) MarshalJSON() ([]byte, error) {
	if src.ApicFabricLeafNode != nil {
		return json.Marshal(&src.ApicFabricLeafNode)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApicFabricLeafNodeRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ApicFabricLeafNode != nil {
		return obj.ApicFabricLeafNode
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableApicFabricLeafNodeRelationship struct {
	value *ApicFabricLeafNodeRelationship
	isSet bool
}

func (v NullableApicFabricLeafNodeRelationship) Get() *ApicFabricLeafNodeRelationship {
	return v.value
}

func (v *NullableApicFabricLeafNodeRelationship) Set(val *ApicFabricLeafNodeRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableApicFabricLeafNodeRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableApicFabricLeafNodeRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApicFabricLeafNodeRelationship(val *ApicFabricLeafNodeRelationship) *NullableApicFabricLeafNodeRelationship {
	return &NullableApicFabricLeafNodeRelationship{value: val, isSet: true}
}

func (v NullableApicFabricLeafNodeRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApicFabricLeafNodeRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}