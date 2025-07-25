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

// FcpoolFcBlockRelationship - A relationship to the 'fcpool.FcBlock' resource, or the expanded 'fcpool.FcBlock' resource, or the 'null' value.
type FcpoolFcBlockRelationship struct {
	FcpoolFcBlock *FcpoolFcBlock
	MoMoRef       *MoMoRef
}

// FcpoolFcBlockAsFcpoolFcBlockRelationship is a convenience function that returns FcpoolFcBlock wrapped in FcpoolFcBlockRelationship
func FcpoolFcBlockAsFcpoolFcBlockRelationship(v *FcpoolFcBlock) FcpoolFcBlockRelationship {
	return FcpoolFcBlockRelationship{
		FcpoolFcBlock: v,
	}
}

// MoMoRefAsFcpoolFcBlockRelationship is a convenience function that returns MoMoRef wrapped in FcpoolFcBlockRelationship
func MoMoRefAsFcpoolFcBlockRelationship(v *MoMoRef) FcpoolFcBlockRelationship {
	return FcpoolFcBlockRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FcpoolFcBlockRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'fcpool.FcBlock'
	if jsonDict["ClassId"] == "fcpool.FcBlock" {
		// try to unmarshal JSON data into FcpoolFcBlock
		err = json.Unmarshal(data, &dst.FcpoolFcBlock)
		if err == nil {
			return nil // data stored in dst.FcpoolFcBlock, return on the first match
		} else {
			dst.FcpoolFcBlock = nil
			return fmt.Errorf("failed to unmarshal FcpoolFcBlockRelationship as FcpoolFcBlock: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FcpoolFcBlockRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FcpoolFcBlockRelationship) MarshalJSON() ([]byte, error) {
	if src.FcpoolFcBlock != nil {
		return json.Marshal(&src.FcpoolFcBlock)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FcpoolFcBlockRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FcpoolFcBlock != nil {
		return obj.FcpoolFcBlock
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableFcpoolFcBlockRelationship struct {
	value *FcpoolFcBlockRelationship
	isSet bool
}

func (v NullableFcpoolFcBlockRelationship) Get() *FcpoolFcBlockRelationship {
	return v.value
}

func (v *NullableFcpoolFcBlockRelationship) Set(val *FcpoolFcBlockRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolFcBlockRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolFcBlockRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolFcBlockRelationship(val *FcpoolFcBlockRelationship) *NullableFcpoolFcBlockRelationship {
	return &NullableFcpoolFcBlockRelationship{value: val, isSet: true}
}

func (v NullableFcpoolFcBlockRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolFcBlockRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
