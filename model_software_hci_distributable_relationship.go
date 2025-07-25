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

// SoftwareHciDistributableRelationship - A relationship to the 'software.HciDistributable' resource, or the expanded 'software.HciDistributable' resource, or the 'null' value.
type SoftwareHciDistributableRelationship struct {
	MoMoRef                  *MoMoRef
	SoftwareHciDistributable *SoftwareHciDistributable
}

// MoMoRefAsSoftwareHciDistributableRelationship is a convenience function that returns MoMoRef wrapped in SoftwareHciDistributableRelationship
func MoMoRefAsSoftwareHciDistributableRelationship(v *MoMoRef) SoftwareHciDistributableRelationship {
	return SoftwareHciDistributableRelationship{
		MoMoRef: v,
	}
}

// SoftwareHciDistributableAsSoftwareHciDistributableRelationship is a convenience function that returns SoftwareHciDistributable wrapped in SoftwareHciDistributableRelationship
func SoftwareHciDistributableAsSoftwareHciDistributableRelationship(v *SoftwareHciDistributable) SoftwareHciDistributableRelationship {
	return SoftwareHciDistributableRelationship{
		SoftwareHciDistributable: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SoftwareHciDistributableRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal SoftwareHciDistributableRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'software.HciDistributable'
	if jsonDict["ClassId"] == "software.HciDistributable" {
		// try to unmarshal JSON data into SoftwareHciDistributable
		err = json.Unmarshal(data, &dst.SoftwareHciDistributable)
		if err == nil {
			return nil // data stored in dst.SoftwareHciDistributable, return on the first match
		} else {
			dst.SoftwareHciDistributable = nil
			return fmt.Errorf("failed to unmarshal SoftwareHciDistributableRelationship as SoftwareHciDistributable: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SoftwareHciDistributableRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.SoftwareHciDistributable != nil {
		return json.Marshal(&src.SoftwareHciDistributable)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SoftwareHciDistributableRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.SoftwareHciDistributable != nil {
		return obj.SoftwareHciDistributable
	}

	// all schemas are nil
	return nil
}

type NullableSoftwareHciDistributableRelationship struct {
	value *SoftwareHciDistributableRelationship
	isSet bool
}

func (v NullableSoftwareHciDistributableRelationship) Get() *SoftwareHciDistributableRelationship {
	return v.value
}

func (v *NullableSoftwareHciDistributableRelationship) Set(val *SoftwareHciDistributableRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareHciDistributableRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareHciDistributableRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareHciDistributableRelationship(val *SoftwareHciDistributableRelationship) *NullableSoftwareHciDistributableRelationship {
	return &NullableSoftwareHciDistributableRelationship{value: val, isSet: true}
}

func (v NullableSoftwareHciDistributableRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareHciDistributableRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
