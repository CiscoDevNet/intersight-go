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

// FcpoolPoolMemberRelationship - A relationship to the 'fcpool.PoolMember' resource, or the expanded 'fcpool.PoolMember' resource, or the 'null' value.
type FcpoolPoolMemberRelationship struct {
	FcpoolPoolMember *FcpoolPoolMember
	MoMoRef          *MoMoRef
}

// FcpoolPoolMemberAsFcpoolPoolMemberRelationship is a convenience function that returns FcpoolPoolMember wrapped in FcpoolPoolMemberRelationship
func FcpoolPoolMemberAsFcpoolPoolMemberRelationship(v *FcpoolPoolMember) FcpoolPoolMemberRelationship {
	return FcpoolPoolMemberRelationship{
		FcpoolPoolMember: v,
	}
}

// MoMoRefAsFcpoolPoolMemberRelationship is a convenience function that returns MoMoRef wrapped in FcpoolPoolMemberRelationship
func MoMoRefAsFcpoolPoolMemberRelationship(v *MoMoRef) FcpoolPoolMemberRelationship {
	return FcpoolPoolMemberRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FcpoolPoolMemberRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'fcpool.PoolMember'
	if jsonDict["ClassId"] == "fcpool.PoolMember" {
		// try to unmarshal JSON data into FcpoolPoolMember
		err = json.Unmarshal(data, &dst.FcpoolPoolMember)
		if err == nil {
			return nil // data stored in dst.FcpoolPoolMember, return on the first match
		} else {
			dst.FcpoolPoolMember = nil
			return fmt.Errorf("failed to unmarshal FcpoolPoolMemberRelationship as FcpoolPoolMember: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FcpoolPoolMemberRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FcpoolPoolMemberRelationship) MarshalJSON() ([]byte, error) {
	if src.FcpoolPoolMember != nil {
		return json.Marshal(&src.FcpoolPoolMember)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FcpoolPoolMemberRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FcpoolPoolMember != nil {
		return obj.FcpoolPoolMember
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableFcpoolPoolMemberRelationship struct {
	value *FcpoolPoolMemberRelationship
	isSet bool
}

func (v NullableFcpoolPoolMemberRelationship) Get() *FcpoolPoolMemberRelationship {
	return v.value
}

func (v *NullableFcpoolPoolMemberRelationship) Set(val *FcpoolPoolMemberRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolPoolMemberRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolPoolMemberRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolPoolMemberRelationship(val *FcpoolPoolMemberRelationship) *NullableFcpoolPoolMemberRelationship {
	return &NullableFcpoolPoolMemberRelationship{value: val, isSet: true}
}

func (v NullableFcpoolPoolMemberRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolPoolMemberRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
