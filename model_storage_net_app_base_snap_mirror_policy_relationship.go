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

// StorageNetAppBaseSnapMirrorPolicyRelationship - A relationship to the 'storage.NetAppBaseSnapMirrorPolicy' resource, or the expanded 'storage.NetAppBaseSnapMirrorPolicy' resource, or the 'null' value.
type StorageNetAppBaseSnapMirrorPolicyRelationship struct {
	MoMoRef                           *MoMoRef
	StorageNetAppBaseSnapMirrorPolicy *StorageNetAppBaseSnapMirrorPolicy
}

// MoMoRefAsStorageNetAppBaseSnapMirrorPolicyRelationship is a convenience function that returns MoMoRef wrapped in StorageNetAppBaseSnapMirrorPolicyRelationship
func MoMoRefAsStorageNetAppBaseSnapMirrorPolicyRelationship(v *MoMoRef) StorageNetAppBaseSnapMirrorPolicyRelationship {
	return StorageNetAppBaseSnapMirrorPolicyRelationship{
		MoMoRef: v,
	}
}

// StorageNetAppBaseSnapMirrorPolicyAsStorageNetAppBaseSnapMirrorPolicyRelationship is a convenience function that returns StorageNetAppBaseSnapMirrorPolicy wrapped in StorageNetAppBaseSnapMirrorPolicyRelationship
func StorageNetAppBaseSnapMirrorPolicyAsStorageNetAppBaseSnapMirrorPolicyRelationship(v *StorageNetAppBaseSnapMirrorPolicy) StorageNetAppBaseSnapMirrorPolicyRelationship {
	return StorageNetAppBaseSnapMirrorPolicyRelationship{
		StorageNetAppBaseSnapMirrorPolicy: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageNetAppBaseSnapMirrorPolicyRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("failed to unmarshal StorageNetAppBaseSnapMirrorPolicyRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.NetAppBaseSnapMirrorPolicy'
	if jsonDict["ClassId"] == "storage.NetAppBaseSnapMirrorPolicy" {
		// try to unmarshal JSON data into StorageNetAppBaseSnapMirrorPolicy
		err = json.Unmarshal(data, &dst.StorageNetAppBaseSnapMirrorPolicy)
		if err == nil {
			return nil // data stored in dst.StorageNetAppBaseSnapMirrorPolicy, return on the first match
		} else {
			dst.StorageNetAppBaseSnapMirrorPolicy = nil
			return fmt.Errorf("failed to unmarshal StorageNetAppBaseSnapMirrorPolicyRelationship as StorageNetAppBaseSnapMirrorPolicy: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageNetAppBaseSnapMirrorPolicyRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.StorageNetAppBaseSnapMirrorPolicy != nil {
		return json.Marshal(&src.StorageNetAppBaseSnapMirrorPolicy)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageNetAppBaseSnapMirrorPolicyRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.StorageNetAppBaseSnapMirrorPolicy != nil {
		return obj.StorageNetAppBaseSnapMirrorPolicy
	}

	// all schemas are nil
	return nil
}

type NullableStorageNetAppBaseSnapMirrorPolicyRelationship struct {
	value *StorageNetAppBaseSnapMirrorPolicyRelationship
	isSet bool
}

func (v NullableStorageNetAppBaseSnapMirrorPolicyRelationship) Get() *StorageNetAppBaseSnapMirrorPolicyRelationship {
	return v.value
}

func (v *NullableStorageNetAppBaseSnapMirrorPolicyRelationship) Set(val *StorageNetAppBaseSnapMirrorPolicyRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppBaseSnapMirrorPolicyRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppBaseSnapMirrorPolicyRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppBaseSnapMirrorPolicyRelationship(val *StorageNetAppBaseSnapMirrorPolicyRelationship) *NullableStorageNetAppBaseSnapMirrorPolicyRelationship {
	return &NullableStorageNetAppBaseSnapMirrorPolicyRelationship{value: val, isSet: true}
}

func (v NullableStorageNetAppBaseSnapMirrorPolicyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppBaseSnapMirrorPolicyRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
