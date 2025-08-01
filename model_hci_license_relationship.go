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

// HciLicenseRelationship - A relationship to the 'hci.License' resource, or the expanded 'hci.License' resource, or the 'null' value.
type HciLicenseRelationship struct {
	HciLicense *HciLicense
	MoMoRef    *MoMoRef
}

// HciLicenseAsHciLicenseRelationship is a convenience function that returns HciLicense wrapped in HciLicenseRelationship
func HciLicenseAsHciLicenseRelationship(v *HciLicense) HciLicenseRelationship {
	return HciLicenseRelationship{
		HciLicense: v,
	}
}

// MoMoRefAsHciLicenseRelationship is a convenience function that returns MoMoRef wrapped in HciLicenseRelationship
func MoMoRefAsHciLicenseRelationship(v *MoMoRef) HciLicenseRelationship {
	return HciLicenseRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HciLicenseRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hci.License'
	if jsonDict["ClassId"] == "hci.License" {
		// try to unmarshal JSON data into HciLicense
		err = json.Unmarshal(data, &dst.HciLicense)
		if err == nil {
			return nil // data stored in dst.HciLicense, return on the first match
		} else {
			dst.HciLicense = nil
			return fmt.Errorf("failed to unmarshal HciLicenseRelationship as HciLicense: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HciLicenseRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HciLicenseRelationship) MarshalJSON() ([]byte, error) {
	if src.HciLicense != nil {
		return json.Marshal(&src.HciLicense)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HciLicenseRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HciLicense != nil {
		return obj.HciLicense
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHciLicenseRelationship struct {
	value *HciLicenseRelationship
	isSet bool
}

func (v NullableHciLicenseRelationship) Get() *HciLicenseRelationship {
	return v.value
}

func (v *NullableHciLicenseRelationship) Set(val *HciLicenseRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHciLicenseRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHciLicenseRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHciLicenseRelationship(val *HciLicenseRelationship) *NullableHciLicenseRelationship {
	return &NullableHciLicenseRelationship{value: val, isSet: true}
}

func (v NullableHciLicenseRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHciLicenseRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
