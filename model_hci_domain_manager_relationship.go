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

// HciDomainManagerRelationship - A relationship to the 'hci.DomainManager' resource, or the expanded 'hci.DomainManager' resource, or the 'null' value.
type HciDomainManagerRelationship struct {
	HciDomainManager *HciDomainManager
	MoMoRef          *MoMoRef
}

// HciDomainManagerAsHciDomainManagerRelationship is a convenience function that returns HciDomainManager wrapped in HciDomainManagerRelationship
func HciDomainManagerAsHciDomainManagerRelationship(v *HciDomainManager) HciDomainManagerRelationship {
	return HciDomainManagerRelationship{
		HciDomainManager: v,
	}
}

// MoMoRefAsHciDomainManagerRelationship is a convenience function that returns MoMoRef wrapped in HciDomainManagerRelationship
func MoMoRefAsHciDomainManagerRelationship(v *MoMoRef) HciDomainManagerRelationship {
	return HciDomainManagerRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HciDomainManagerRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hci.DomainManager'
	if jsonDict["ClassId"] == "hci.DomainManager" {
		// try to unmarshal JSON data into HciDomainManager
		err = json.Unmarshal(data, &dst.HciDomainManager)
		if err == nil {
			return nil // data stored in dst.HciDomainManager, return on the first match
		} else {
			dst.HciDomainManager = nil
			return fmt.Errorf("failed to unmarshal HciDomainManagerRelationship as HciDomainManager: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HciDomainManagerRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HciDomainManagerRelationship) MarshalJSON() ([]byte, error) {
	if src.HciDomainManager != nil {
		return json.Marshal(&src.HciDomainManager)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HciDomainManagerRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HciDomainManager != nil {
		return obj.HciDomainManager
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHciDomainManagerRelationship struct {
	value *HciDomainManagerRelationship
	isSet bool
}

func (v NullableHciDomainManagerRelationship) Get() *HciDomainManagerRelationship {
	return v.value
}

func (v *NullableHciDomainManagerRelationship) Set(val *HciDomainManagerRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHciDomainManagerRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHciDomainManagerRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHciDomainManagerRelationship(val *HciDomainManagerRelationship) *NullableHciDomainManagerRelationship {
	return &NullableHciDomainManagerRelationship{value: val, isSet: true}
}

func (v NullableHciDomainManagerRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHciDomainManagerRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}