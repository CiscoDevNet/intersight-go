/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2024112619
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// HciComplianceRelationship - A relationship to the 'hci.Compliance' resource, or the expanded 'hci.Compliance' resource, or the 'null' value.
type HciComplianceRelationship struct {
	HciCompliance *HciCompliance
	MoMoRef       *MoMoRef
}

// HciComplianceAsHciComplianceRelationship is a convenience function that returns HciCompliance wrapped in HciComplianceRelationship
func HciComplianceAsHciComplianceRelationship(v *HciCompliance) HciComplianceRelationship {
	return HciComplianceRelationship{
		HciCompliance: v,
	}
}

// MoMoRefAsHciComplianceRelationship is a convenience function that returns MoMoRef wrapped in HciComplianceRelationship
func MoMoRefAsHciComplianceRelationship(v *MoMoRef) HciComplianceRelationship {
	return HciComplianceRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HciComplianceRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hci.Compliance'
	if jsonDict["ClassId"] == "hci.Compliance" {
		// try to unmarshal JSON data into HciCompliance
		err = json.Unmarshal(data, &dst.HciCompliance)
		if err == nil {
			return nil // data stored in dst.HciCompliance, return on the first match
		} else {
			dst.HciCompliance = nil
			return fmt.Errorf("failed to unmarshal HciComplianceRelationship as HciCompliance: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HciComplianceRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HciComplianceRelationship) MarshalJSON() ([]byte, error) {
	if src.HciCompliance != nil {
		return json.Marshal(&src.HciCompliance)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HciComplianceRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HciCompliance != nil {
		return obj.HciCompliance
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHciComplianceRelationship struct {
	value *HciComplianceRelationship
	isSet bool
}

func (v NullableHciComplianceRelationship) Get() *HciComplianceRelationship {
	return v.value
}

func (v *NullableHciComplianceRelationship) Set(val *HciComplianceRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHciComplianceRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHciComplianceRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHciComplianceRelationship(val *HciComplianceRelationship) *NullableHciComplianceRelationship {
	return &NullableHciComplianceRelationship{value: val, isSet: true}
}

func (v NullableHciComplianceRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHciComplianceRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}