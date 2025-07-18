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

// IamQualifierRelationship - A relationship to the 'iam.Qualifier' resource, or the expanded 'iam.Qualifier' resource, or the 'null' value.
type IamQualifierRelationship struct {
	IamQualifier *IamQualifier
	MoMoRef      *MoMoRef
}

// IamQualifierAsIamQualifierRelationship is a convenience function that returns IamQualifier wrapped in IamQualifierRelationship
func IamQualifierAsIamQualifierRelationship(v *IamQualifier) IamQualifierRelationship {
	return IamQualifierRelationship{
		IamQualifier: v,
	}
}

// MoMoRefAsIamQualifierRelationship is a convenience function that returns MoMoRef wrapped in IamQualifierRelationship
func MoMoRefAsIamQualifierRelationship(v *MoMoRef) IamQualifierRelationship {
	return IamQualifierRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamQualifierRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'iam.Qualifier'
	if jsonDict["ClassId"] == "iam.Qualifier" {
		// try to unmarshal JSON data into IamQualifier
		err = json.Unmarshal(data, &dst.IamQualifier)
		if err == nil {
			return nil // data stored in dst.IamQualifier, return on the first match
		} else {
			dst.IamQualifier = nil
			return fmt.Errorf("failed to unmarshal IamQualifierRelationship as IamQualifier: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamQualifierRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamQualifierRelationship) MarshalJSON() ([]byte, error) {
	if src.IamQualifier != nil {
		return json.Marshal(&src.IamQualifier)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IamQualifierRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IamQualifier != nil {
		return obj.IamQualifier
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableIamQualifierRelationship struct {
	value *IamQualifierRelationship
	isSet bool
}

func (v NullableIamQualifierRelationship) Get() *IamQualifierRelationship {
	return v.value
}

func (v *NullableIamQualifierRelationship) Set(val *IamQualifierRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableIamQualifierRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableIamQualifierRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamQualifierRelationship(val *IamQualifierRelationship) *NullableIamQualifierRelationship {
	return &NullableIamQualifierRelationship{value: val, isSet: true}
}

func (v NullableIamQualifierRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamQualifierRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
