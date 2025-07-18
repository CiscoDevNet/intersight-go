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

// ConvergedinfraBaseComplianceDetailsRelationship - A relationship to the 'convergedinfra.BaseComplianceDetails' resource, or the expanded 'convergedinfra.BaseComplianceDetails' resource, or the 'null' value.
type ConvergedinfraBaseComplianceDetailsRelationship struct {
	ConvergedinfraBaseComplianceDetails *ConvergedinfraBaseComplianceDetails
	MoMoRef                             *MoMoRef
}

// ConvergedinfraBaseComplianceDetailsAsConvergedinfraBaseComplianceDetailsRelationship is a convenience function that returns ConvergedinfraBaseComplianceDetails wrapped in ConvergedinfraBaseComplianceDetailsRelationship
func ConvergedinfraBaseComplianceDetailsAsConvergedinfraBaseComplianceDetailsRelationship(v *ConvergedinfraBaseComplianceDetails) ConvergedinfraBaseComplianceDetailsRelationship {
	return ConvergedinfraBaseComplianceDetailsRelationship{
		ConvergedinfraBaseComplianceDetails: v,
	}
}

// MoMoRefAsConvergedinfraBaseComplianceDetailsRelationship is a convenience function that returns MoMoRef wrapped in ConvergedinfraBaseComplianceDetailsRelationship
func MoMoRefAsConvergedinfraBaseComplianceDetailsRelationship(v *MoMoRef) ConvergedinfraBaseComplianceDetailsRelationship {
	return ConvergedinfraBaseComplianceDetailsRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ConvergedinfraBaseComplianceDetailsRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'convergedinfra.BaseComplianceDetails'
	if jsonDict["ClassId"] == "convergedinfra.BaseComplianceDetails" {
		// try to unmarshal JSON data into ConvergedinfraBaseComplianceDetails
		err = json.Unmarshal(data, &dst.ConvergedinfraBaseComplianceDetails)
		if err == nil {
			return nil // data stored in dst.ConvergedinfraBaseComplianceDetails, return on the first match
		} else {
			dst.ConvergedinfraBaseComplianceDetails = nil
			return fmt.Errorf("failed to unmarshal ConvergedinfraBaseComplianceDetailsRelationship as ConvergedinfraBaseComplianceDetails: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ConvergedinfraBaseComplianceDetailsRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConvergedinfraBaseComplianceDetailsRelationship) MarshalJSON() ([]byte, error) {
	if src.ConvergedinfraBaseComplianceDetails != nil {
		return json.Marshal(&src.ConvergedinfraBaseComplianceDetails)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ConvergedinfraBaseComplianceDetailsRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ConvergedinfraBaseComplianceDetails != nil {
		return obj.ConvergedinfraBaseComplianceDetails
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableConvergedinfraBaseComplianceDetailsRelationship struct {
	value *ConvergedinfraBaseComplianceDetailsRelationship
	isSet bool
}

func (v NullableConvergedinfraBaseComplianceDetailsRelationship) Get() *ConvergedinfraBaseComplianceDetailsRelationship {
	return v.value
}

func (v *NullableConvergedinfraBaseComplianceDetailsRelationship) Set(val *ConvergedinfraBaseComplianceDetailsRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraBaseComplianceDetailsRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraBaseComplianceDetailsRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraBaseComplianceDetailsRelationship(val *ConvergedinfraBaseComplianceDetailsRelationship) *NullableConvergedinfraBaseComplianceDetailsRelationship {
	return &NullableConvergedinfraBaseComplianceDetailsRelationship{value: val, isSet: true}
}

func (v NullableConvergedinfraBaseComplianceDetailsRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraBaseComplianceDetailsRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
