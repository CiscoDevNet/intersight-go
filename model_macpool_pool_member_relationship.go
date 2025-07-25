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

// MacpoolPoolMemberRelationship - A relationship to the 'macpool.PoolMember' resource, or the expanded 'macpool.PoolMember' resource, or the 'null' value.
type MacpoolPoolMemberRelationship struct {
	MacpoolPoolMember *MacpoolPoolMember
	MoMoRef           *MoMoRef
}

// MacpoolPoolMemberAsMacpoolPoolMemberRelationship is a convenience function that returns MacpoolPoolMember wrapped in MacpoolPoolMemberRelationship
func MacpoolPoolMemberAsMacpoolPoolMemberRelationship(v *MacpoolPoolMember) MacpoolPoolMemberRelationship {
	return MacpoolPoolMemberRelationship{
		MacpoolPoolMember: v,
	}
}

// MoMoRefAsMacpoolPoolMemberRelationship is a convenience function that returns MoMoRef wrapped in MacpoolPoolMemberRelationship
func MoMoRefAsMacpoolPoolMemberRelationship(v *MoMoRef) MacpoolPoolMemberRelationship {
	return MacpoolPoolMemberRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MacpoolPoolMemberRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'macpool.PoolMember'
	if jsonDict["ClassId"] == "macpool.PoolMember" {
		// try to unmarshal JSON data into MacpoolPoolMember
		err = json.Unmarshal(data, &dst.MacpoolPoolMember)
		if err == nil {
			return nil // data stored in dst.MacpoolPoolMember, return on the first match
		} else {
			dst.MacpoolPoolMember = nil
			return fmt.Errorf("failed to unmarshal MacpoolPoolMemberRelationship as MacpoolPoolMember: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MacpoolPoolMemberRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MacpoolPoolMemberRelationship) MarshalJSON() ([]byte, error) {
	if src.MacpoolPoolMember != nil {
		return json.Marshal(&src.MacpoolPoolMember)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MacpoolPoolMemberRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MacpoolPoolMember != nil {
		return obj.MacpoolPoolMember
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableMacpoolPoolMemberRelationship struct {
	value *MacpoolPoolMemberRelationship
	isSet bool
}

func (v NullableMacpoolPoolMemberRelationship) Get() *MacpoolPoolMemberRelationship {
	return v.value
}

func (v *NullableMacpoolPoolMemberRelationship) Set(val *MacpoolPoolMemberRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableMacpoolPoolMemberRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableMacpoolPoolMemberRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacpoolPoolMemberRelationship(val *MacpoolPoolMemberRelationship) *NullableMacpoolPoolMemberRelationship {
	return &NullableMacpoolPoolMemberRelationship{value: val, isSet: true}
}

func (v NullableMacpoolPoolMemberRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacpoolPoolMemberRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
