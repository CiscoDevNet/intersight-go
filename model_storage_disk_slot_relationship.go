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

// StorageDiskSlotRelationship - A relationship to the 'storage.DiskSlot' resource, or the expanded 'storage.DiskSlot' resource, or the 'null' value.
type StorageDiskSlotRelationship struct {
	MoMoRef         *MoMoRef
	StorageDiskSlot *StorageDiskSlot
}

// MoMoRefAsStorageDiskSlotRelationship is a convenience function that returns MoMoRef wrapped in StorageDiskSlotRelationship
func MoMoRefAsStorageDiskSlotRelationship(v *MoMoRef) StorageDiskSlotRelationship {
	return StorageDiskSlotRelationship{
		MoMoRef: v,
	}
}

// StorageDiskSlotAsStorageDiskSlotRelationship is a convenience function that returns StorageDiskSlot wrapped in StorageDiskSlotRelationship
func StorageDiskSlotAsStorageDiskSlotRelationship(v *StorageDiskSlot) StorageDiskSlotRelationship {
	return StorageDiskSlotRelationship{
		StorageDiskSlot: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageDiskSlotRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal StorageDiskSlotRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.DiskSlot'
	if jsonDict["ClassId"] == "storage.DiskSlot" {
		// try to unmarshal JSON data into StorageDiskSlot
		err = json.Unmarshal(data, &dst.StorageDiskSlot)
		if err == nil {
			return nil // data stored in dst.StorageDiskSlot, return on the first match
		} else {
			dst.StorageDiskSlot = nil
			return fmt.Errorf("failed to unmarshal StorageDiskSlotRelationship as StorageDiskSlot: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageDiskSlotRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.StorageDiskSlot != nil {
		return json.Marshal(&src.StorageDiskSlot)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageDiskSlotRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.StorageDiskSlot != nil {
		return obj.StorageDiskSlot
	}

	// all schemas are nil
	return nil
}

type NullableStorageDiskSlotRelationship struct {
	value *StorageDiskSlotRelationship
	isSet bool
}

func (v NullableStorageDiskSlotRelationship) Get() *StorageDiskSlotRelationship {
	return v.value
}

func (v *NullableStorageDiskSlotRelationship) Set(val *StorageDiskSlotRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDiskSlotRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDiskSlotRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDiskSlotRelationship(val *StorageDiskSlotRelationship) *NullableStorageDiskSlotRelationship {
	return &NullableStorageDiskSlotRelationship{value: val, isSet: true}
}

func (v NullableStorageDiskSlotRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDiskSlotRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
