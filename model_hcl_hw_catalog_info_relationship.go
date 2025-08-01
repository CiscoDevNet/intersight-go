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

// HclHwCatalogInfoRelationship - A relationship to the 'hcl.HwCatalogInfo' resource, or the expanded 'hcl.HwCatalogInfo' resource, or the 'null' value.
type HclHwCatalogInfoRelationship struct {
	HclHwCatalogInfo *HclHwCatalogInfo
	MoMoRef          *MoMoRef
}

// HclHwCatalogInfoAsHclHwCatalogInfoRelationship is a convenience function that returns HclHwCatalogInfo wrapped in HclHwCatalogInfoRelationship
func HclHwCatalogInfoAsHclHwCatalogInfoRelationship(v *HclHwCatalogInfo) HclHwCatalogInfoRelationship {
	return HclHwCatalogInfoRelationship{
		HclHwCatalogInfo: v,
	}
}

// MoMoRefAsHclHwCatalogInfoRelationship is a convenience function that returns MoMoRef wrapped in HclHwCatalogInfoRelationship
func MoMoRefAsHclHwCatalogInfoRelationship(v *MoMoRef) HclHwCatalogInfoRelationship {
	return HclHwCatalogInfoRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HclHwCatalogInfoRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hcl.HwCatalogInfo'
	if jsonDict["ClassId"] == "hcl.HwCatalogInfo" {
		// try to unmarshal JSON data into HclHwCatalogInfo
		err = json.Unmarshal(data, &dst.HclHwCatalogInfo)
		if err == nil {
			return nil // data stored in dst.HclHwCatalogInfo, return on the first match
		} else {
			dst.HclHwCatalogInfo = nil
			return fmt.Errorf("failed to unmarshal HclHwCatalogInfoRelationship as HclHwCatalogInfo: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HclHwCatalogInfoRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HclHwCatalogInfoRelationship) MarshalJSON() ([]byte, error) {
	if src.HclHwCatalogInfo != nil {
		return json.Marshal(&src.HclHwCatalogInfo)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HclHwCatalogInfoRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HclHwCatalogInfo != nil {
		return obj.HclHwCatalogInfo
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHclHwCatalogInfoRelationship struct {
	value *HclHwCatalogInfoRelationship
	isSet bool
}

func (v NullableHclHwCatalogInfoRelationship) Get() *HclHwCatalogInfoRelationship {
	return v.value
}

func (v *NullableHclHwCatalogInfoRelationship) Set(val *HclHwCatalogInfoRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHclHwCatalogInfoRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHclHwCatalogInfoRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclHwCatalogInfoRelationship(val *HclHwCatalogInfoRelationship) *NullableHclHwCatalogInfoRelationship {
	return &NullableHclHwCatalogInfoRelationship{value: val, isSet: true}
}

func (v NullableHclHwCatalogInfoRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclHwCatalogInfoRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
