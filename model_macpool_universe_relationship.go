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

// MacpoolUniverseRelationship - A relationship to the 'macpool.Universe' resource, or the expanded 'macpool.Universe' resource, or the 'null' value.
type MacpoolUniverseRelationship struct {
	MacpoolUniverse *MacpoolUniverse
	MoMoRef         *MoMoRef
}

// MacpoolUniverseAsMacpoolUniverseRelationship is a convenience function that returns MacpoolUniverse wrapped in MacpoolUniverseRelationship
func MacpoolUniverseAsMacpoolUniverseRelationship(v *MacpoolUniverse) MacpoolUniverseRelationship {
	return MacpoolUniverseRelationship{
		MacpoolUniverse: v,
	}
}

// MoMoRefAsMacpoolUniverseRelationship is a convenience function that returns MoMoRef wrapped in MacpoolUniverseRelationship
func MoMoRefAsMacpoolUniverseRelationship(v *MoMoRef) MacpoolUniverseRelationship {
	return MacpoolUniverseRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MacpoolUniverseRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'macpool.Universe'
	if jsonDict["ClassId"] == "macpool.Universe" {
		// try to unmarshal JSON data into MacpoolUniverse
		err = json.Unmarshal(data, &dst.MacpoolUniverse)
		if err == nil {
			return nil // data stored in dst.MacpoolUniverse, return on the first match
		} else {
			dst.MacpoolUniverse = nil
			return fmt.Errorf("failed to unmarshal MacpoolUniverseRelationship as MacpoolUniverse: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MacpoolUniverseRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MacpoolUniverseRelationship) MarshalJSON() ([]byte, error) {
	if src.MacpoolUniverse != nil {
		return json.Marshal(&src.MacpoolUniverse)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MacpoolUniverseRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MacpoolUniverse != nil {
		return obj.MacpoolUniverse
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableMacpoolUniverseRelationship struct {
	value *MacpoolUniverseRelationship
	isSet bool
}

func (v NullableMacpoolUniverseRelationship) Get() *MacpoolUniverseRelationship {
	return v.value
}

func (v *NullableMacpoolUniverseRelationship) Set(val *MacpoolUniverseRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableMacpoolUniverseRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableMacpoolUniverseRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacpoolUniverseRelationship(val *MacpoolUniverseRelationship) *NullableMacpoolUniverseRelationship {
	return &NullableMacpoolUniverseRelationship{value: val, isSet: true}
}

func (v NullableMacpoolUniverseRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacpoolUniverseRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
