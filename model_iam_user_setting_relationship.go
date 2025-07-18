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

// IamUserSettingRelationship - A relationship to the 'iam.UserSetting' resource, or the expanded 'iam.UserSetting' resource, or the 'null' value.
type IamUserSettingRelationship struct {
	IamUserSetting *IamUserSetting
	MoMoRef        *MoMoRef
}

// IamUserSettingAsIamUserSettingRelationship is a convenience function that returns IamUserSetting wrapped in IamUserSettingRelationship
func IamUserSettingAsIamUserSettingRelationship(v *IamUserSetting) IamUserSettingRelationship {
	return IamUserSettingRelationship{
		IamUserSetting: v,
	}
}

// MoMoRefAsIamUserSettingRelationship is a convenience function that returns MoMoRef wrapped in IamUserSettingRelationship
func MoMoRefAsIamUserSettingRelationship(v *MoMoRef) IamUserSettingRelationship {
	return IamUserSettingRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamUserSettingRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'iam.UserSetting'
	if jsonDict["ClassId"] == "iam.UserSetting" {
		// try to unmarshal JSON data into IamUserSetting
		err = json.Unmarshal(data, &dst.IamUserSetting)
		if err == nil {
			return nil // data stored in dst.IamUserSetting, return on the first match
		} else {
			dst.IamUserSetting = nil
			return fmt.Errorf("failed to unmarshal IamUserSettingRelationship as IamUserSetting: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamUserSettingRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamUserSettingRelationship) MarshalJSON() ([]byte, error) {
	if src.IamUserSetting != nil {
		return json.Marshal(&src.IamUserSetting)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IamUserSettingRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IamUserSetting != nil {
		return obj.IamUserSetting
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableIamUserSettingRelationship struct {
	value *IamUserSettingRelationship
	isSet bool
}

func (v NullableIamUserSettingRelationship) Get() *IamUserSettingRelationship {
	return v.value
}

func (v *NullableIamUserSettingRelationship) Set(val *IamUserSettingRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableIamUserSettingRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableIamUserSettingRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamUserSettingRelationship(val *IamUserSettingRelationship) *NullableIamUserSettingRelationship {
	return &NullableIamUserSettingRelationship{value: val, isSet: true}
}

func (v NullableIamUserSettingRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamUserSettingRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
