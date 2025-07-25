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

// SecurityUnitRelationship - A relationship to the 'security.Unit' resource, or the expanded 'security.Unit' resource, or the 'null' value.
type SecurityUnitRelationship struct {
	MoMoRef      *MoMoRef
	SecurityUnit *SecurityUnit
}

// MoMoRefAsSecurityUnitRelationship is a convenience function that returns MoMoRef wrapped in SecurityUnitRelationship
func MoMoRefAsSecurityUnitRelationship(v *MoMoRef) SecurityUnitRelationship {
	return SecurityUnitRelationship{
		MoMoRef: v,
	}
}

// SecurityUnitAsSecurityUnitRelationship is a convenience function that returns SecurityUnit wrapped in SecurityUnitRelationship
func SecurityUnitAsSecurityUnitRelationship(v *SecurityUnit) SecurityUnitRelationship {
	return SecurityUnitRelationship{
		SecurityUnit: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SecurityUnitRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal SecurityUnitRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'security.Unit'
	if jsonDict["ClassId"] == "security.Unit" {
		// try to unmarshal JSON data into SecurityUnit
		err = json.Unmarshal(data, &dst.SecurityUnit)
		if err == nil {
			return nil // data stored in dst.SecurityUnit, return on the first match
		} else {
			dst.SecurityUnit = nil
			return fmt.Errorf("failed to unmarshal SecurityUnitRelationship as SecurityUnit: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SecurityUnitRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.SecurityUnit != nil {
		return json.Marshal(&src.SecurityUnit)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SecurityUnitRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.SecurityUnit != nil {
		return obj.SecurityUnit
	}

	// all schemas are nil
	return nil
}

type NullableSecurityUnitRelationship struct {
	value *SecurityUnitRelationship
	isSet bool
}

func (v NullableSecurityUnitRelationship) Get() *SecurityUnitRelationship {
	return v.value
}

func (v *NullableSecurityUnitRelationship) Set(val *SecurityUnitRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityUnitRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityUnitRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityUnitRelationship(val *SecurityUnitRelationship) *NullableSecurityUnitRelationship {
	return &NullableSecurityUnitRelationship{value: val, isSet: true}
}

func (v NullableSecurityUnitRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityUnitRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
