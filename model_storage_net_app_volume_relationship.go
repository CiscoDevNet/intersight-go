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

// StorageNetAppVolumeRelationship - A relationship to the 'storage.NetAppVolume' resource, or the expanded 'storage.NetAppVolume' resource, or the 'null' value.
type StorageNetAppVolumeRelationship struct {
	MoMoRef             *MoMoRef
	StorageNetAppVolume *StorageNetAppVolume
}

// MoMoRefAsStorageNetAppVolumeRelationship is a convenience function that returns MoMoRef wrapped in StorageNetAppVolumeRelationship
func MoMoRefAsStorageNetAppVolumeRelationship(v *MoMoRef) StorageNetAppVolumeRelationship {
	return StorageNetAppVolumeRelationship{
		MoMoRef: v,
	}
}

// StorageNetAppVolumeAsStorageNetAppVolumeRelationship is a convenience function that returns StorageNetAppVolume wrapped in StorageNetAppVolumeRelationship
func StorageNetAppVolumeAsStorageNetAppVolumeRelationship(v *StorageNetAppVolume) StorageNetAppVolumeRelationship {
	return StorageNetAppVolumeRelationship{
		StorageNetAppVolume: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageNetAppVolumeRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal StorageNetAppVolumeRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.NetAppVolume'
	if jsonDict["ClassId"] == "storage.NetAppVolume" {
		// try to unmarshal JSON data into StorageNetAppVolume
		err = json.Unmarshal(data, &dst.StorageNetAppVolume)
		if err == nil {
			return nil // data stored in dst.StorageNetAppVolume, return on the first match
		} else {
			dst.StorageNetAppVolume = nil
			return fmt.Errorf("failed to unmarshal StorageNetAppVolumeRelationship as StorageNetAppVolume: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageNetAppVolumeRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.StorageNetAppVolume != nil {
		return json.Marshal(&src.StorageNetAppVolume)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageNetAppVolumeRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.StorageNetAppVolume != nil {
		return obj.StorageNetAppVolume
	}

	// all schemas are nil
	return nil
}

type NullableStorageNetAppVolumeRelationship struct {
	value *StorageNetAppVolumeRelationship
	isSet bool
}

func (v NullableStorageNetAppVolumeRelationship) Get() *StorageNetAppVolumeRelationship {
	return v.value
}

func (v *NullableStorageNetAppVolumeRelationship) Set(val *StorageNetAppVolumeRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppVolumeRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppVolumeRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppVolumeRelationship(val *StorageNetAppVolumeRelationship) *NullableStorageNetAppVolumeRelationship {
	return &NullableStorageNetAppVolumeRelationship{value: val, isSet: true}
}

func (v NullableStorageNetAppVolumeRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppVolumeRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
