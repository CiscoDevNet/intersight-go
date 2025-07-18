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

// IamEndPointPrivilegeRelationship - A relationship to the 'iam.EndPointPrivilege' resource, or the expanded 'iam.EndPointPrivilege' resource, or the 'null' value.
type IamEndPointPrivilegeRelationship struct {
	IamEndPointPrivilege *IamEndPointPrivilege
	MoMoRef              *MoMoRef
}

// IamEndPointPrivilegeAsIamEndPointPrivilegeRelationship is a convenience function that returns IamEndPointPrivilege wrapped in IamEndPointPrivilegeRelationship
func IamEndPointPrivilegeAsIamEndPointPrivilegeRelationship(v *IamEndPointPrivilege) IamEndPointPrivilegeRelationship {
	return IamEndPointPrivilegeRelationship{
		IamEndPointPrivilege: v,
	}
}

// MoMoRefAsIamEndPointPrivilegeRelationship is a convenience function that returns MoMoRef wrapped in IamEndPointPrivilegeRelationship
func MoMoRefAsIamEndPointPrivilegeRelationship(v *MoMoRef) IamEndPointPrivilegeRelationship {
	return IamEndPointPrivilegeRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamEndPointPrivilegeRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'iam.EndPointPrivilege'
	if jsonDict["ClassId"] == "iam.EndPointPrivilege" {
		// try to unmarshal JSON data into IamEndPointPrivilege
		err = json.Unmarshal(data, &dst.IamEndPointPrivilege)
		if err == nil {
			return nil // data stored in dst.IamEndPointPrivilege, return on the first match
		} else {
			dst.IamEndPointPrivilege = nil
			return fmt.Errorf("failed to unmarshal IamEndPointPrivilegeRelationship as IamEndPointPrivilege: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamEndPointPrivilegeRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamEndPointPrivilegeRelationship) MarshalJSON() ([]byte, error) {
	if src.IamEndPointPrivilege != nil {
		return json.Marshal(&src.IamEndPointPrivilege)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IamEndPointPrivilegeRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IamEndPointPrivilege != nil {
		return obj.IamEndPointPrivilege
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableIamEndPointPrivilegeRelationship struct {
	value *IamEndPointPrivilegeRelationship
	isSet bool
}

func (v NullableIamEndPointPrivilegeRelationship) Get() *IamEndPointPrivilegeRelationship {
	return v.value
}

func (v *NullableIamEndPointPrivilegeRelationship) Set(val *IamEndPointPrivilegeRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointPrivilegeRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointPrivilegeRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointPrivilegeRelationship(val *IamEndPointPrivilegeRelationship) *NullableIamEndPointPrivilegeRelationship {
	return &NullableIamEndPointPrivilegeRelationship{value: val, isSet: true}
}

func (v NullableIamEndPointPrivilegeRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointPrivilegeRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
