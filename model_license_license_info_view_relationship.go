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

// LicenseLicenseInfoViewRelationship - A relationship to the 'license.LicenseInfoView' resource, or the expanded 'license.LicenseInfoView' resource, or the 'null' value.
type LicenseLicenseInfoViewRelationship struct {
	LicenseLicenseInfoView *LicenseLicenseInfoView
	MoMoRef                *MoMoRef
}

// LicenseLicenseInfoViewAsLicenseLicenseInfoViewRelationship is a convenience function that returns LicenseLicenseInfoView wrapped in LicenseLicenseInfoViewRelationship
func LicenseLicenseInfoViewAsLicenseLicenseInfoViewRelationship(v *LicenseLicenseInfoView) LicenseLicenseInfoViewRelationship {
	return LicenseLicenseInfoViewRelationship{
		LicenseLicenseInfoView: v,
	}
}

// MoMoRefAsLicenseLicenseInfoViewRelationship is a convenience function that returns MoMoRef wrapped in LicenseLicenseInfoViewRelationship
func MoMoRefAsLicenseLicenseInfoViewRelationship(v *MoMoRef) LicenseLicenseInfoViewRelationship {
	return LicenseLicenseInfoViewRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LicenseLicenseInfoViewRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'license.LicenseInfoView'
	if jsonDict["ClassId"] == "license.LicenseInfoView" {
		// try to unmarshal JSON data into LicenseLicenseInfoView
		err = json.Unmarshal(data, &dst.LicenseLicenseInfoView)
		if err == nil {
			return nil // data stored in dst.LicenseLicenseInfoView, return on the first match
		} else {
			dst.LicenseLicenseInfoView = nil
			return fmt.Errorf("failed to unmarshal LicenseLicenseInfoViewRelationship as LicenseLicenseInfoView: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal LicenseLicenseInfoViewRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LicenseLicenseInfoViewRelationship) MarshalJSON() ([]byte, error) {
	if src.LicenseLicenseInfoView != nil {
		return json.Marshal(&src.LicenseLicenseInfoView)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LicenseLicenseInfoViewRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.LicenseLicenseInfoView != nil {
		return obj.LicenseLicenseInfoView
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableLicenseLicenseInfoViewRelationship struct {
	value *LicenseLicenseInfoViewRelationship
	isSet bool
}

func (v NullableLicenseLicenseInfoViewRelationship) Get() *LicenseLicenseInfoViewRelationship {
	return v.value
}

func (v *NullableLicenseLicenseInfoViewRelationship) Set(val *LicenseLicenseInfoViewRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseLicenseInfoViewRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseLicenseInfoViewRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseLicenseInfoViewRelationship(val *LicenseLicenseInfoViewRelationship) *NullableLicenseLicenseInfoViewRelationship {
	return &NullableLicenseLicenseInfoViewRelationship{value: val, isSet: true}
}

func (v NullableLicenseLicenseInfoViewRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseLicenseInfoViewRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
