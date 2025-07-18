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

// IamResourceLimitsRelationship - A relationship to the 'iam.ResourceLimits' resource, or the expanded 'iam.ResourceLimits' resource, or the 'null' value.
type IamResourceLimitsRelationship struct {
	IamResourceLimits *IamResourceLimits
	MoMoRef           *MoMoRef
}

// IamResourceLimitsAsIamResourceLimitsRelationship is a convenience function that returns IamResourceLimits wrapped in IamResourceLimitsRelationship
func IamResourceLimitsAsIamResourceLimitsRelationship(v *IamResourceLimits) IamResourceLimitsRelationship {
	return IamResourceLimitsRelationship{
		IamResourceLimits: v,
	}
}

// MoMoRefAsIamResourceLimitsRelationship is a convenience function that returns MoMoRef wrapped in IamResourceLimitsRelationship
func MoMoRefAsIamResourceLimitsRelationship(v *MoMoRef) IamResourceLimitsRelationship {
	return IamResourceLimitsRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamResourceLimitsRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'iam.ResourceLimits'
	if jsonDict["ClassId"] == "iam.ResourceLimits" {
		// try to unmarshal JSON data into IamResourceLimits
		err = json.Unmarshal(data, &dst.IamResourceLimits)
		if err == nil {
			return nil // data stored in dst.IamResourceLimits, return on the first match
		} else {
			dst.IamResourceLimits = nil
			return fmt.Errorf("failed to unmarshal IamResourceLimitsRelationship as IamResourceLimits: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamResourceLimitsRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamResourceLimitsRelationship) MarshalJSON() ([]byte, error) {
	if src.IamResourceLimits != nil {
		return json.Marshal(&src.IamResourceLimits)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IamResourceLimitsRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IamResourceLimits != nil {
		return obj.IamResourceLimits
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableIamResourceLimitsRelationship struct {
	value *IamResourceLimitsRelationship
	isSet bool
}

func (v NullableIamResourceLimitsRelationship) Get() *IamResourceLimitsRelationship {
	return v.value
}

func (v *NullableIamResourceLimitsRelationship) Set(val *IamResourceLimitsRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableIamResourceLimitsRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableIamResourceLimitsRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamResourceLimitsRelationship(val *IamResourceLimitsRelationship) *NullableIamResourceLimitsRelationship {
	return &NullableIamResourceLimitsRelationship{value: val, isSet: true}
}

func (v NullableIamResourceLimitsRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamResourceLimitsRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
