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

// StorageDriveGroupRelationship - A relationship to the 'storage.DriveGroup' resource, or the expanded 'storage.DriveGroup' resource, or the 'null' value.
type StorageDriveGroupRelationship struct {
	MoMoRef           *MoMoRef
	StorageDriveGroup *StorageDriveGroup
}

// MoMoRefAsStorageDriveGroupRelationship is a convenience function that returns MoMoRef wrapped in StorageDriveGroupRelationship
func MoMoRefAsStorageDriveGroupRelationship(v *MoMoRef) StorageDriveGroupRelationship {
	return StorageDriveGroupRelationship{
		MoMoRef: v,
	}
}

// StorageDriveGroupAsStorageDriveGroupRelationship is a convenience function that returns StorageDriveGroup wrapped in StorageDriveGroupRelationship
func StorageDriveGroupAsStorageDriveGroupRelationship(v *StorageDriveGroup) StorageDriveGroupRelationship {
	return StorageDriveGroupRelationship{
		StorageDriveGroup: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageDriveGroupRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal StorageDriveGroupRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.DriveGroup'
	if jsonDict["ClassId"] == "storage.DriveGroup" {
		// try to unmarshal JSON data into StorageDriveGroup
		err = json.Unmarshal(data, &dst.StorageDriveGroup)
		if err == nil {
			return nil // data stored in dst.StorageDriveGroup, return on the first match
		} else {
			dst.StorageDriveGroup = nil
			return fmt.Errorf("failed to unmarshal StorageDriveGroupRelationship as StorageDriveGroup: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageDriveGroupRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.StorageDriveGroup != nil {
		return json.Marshal(&src.StorageDriveGroup)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageDriveGroupRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.StorageDriveGroup != nil {
		return obj.StorageDriveGroup
	}

	// all schemas are nil
	return nil
}

type NullableStorageDriveGroupRelationship struct {
	value *StorageDriveGroupRelationship
	isSet bool
}

func (v NullableStorageDriveGroupRelationship) Get() *StorageDriveGroupRelationship {
	return v.value
}

func (v *NullableStorageDriveGroupRelationship) Set(val *StorageDriveGroupRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDriveGroupRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDriveGroupRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDriveGroupRelationship(val *StorageDriveGroupRelationship) *NullableStorageDriveGroupRelationship {
	return &NullableStorageDriveGroupRelationship{value: val, isSet: true}
}

func (v NullableStorageDriveGroupRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDriveGroupRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
