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

// StorageNetAppNonDataIpInterfaceEventRelationship - A relationship to the 'storage.NetAppNonDataIpInterfaceEvent' resource, or the expanded 'storage.NetAppNonDataIpInterfaceEvent' resource, or the 'null' value.
type StorageNetAppNonDataIpInterfaceEventRelationship struct {
	MoMoRef                              *MoMoRef
	StorageNetAppNonDataIpInterfaceEvent *StorageNetAppNonDataIpInterfaceEvent
}

// MoMoRefAsStorageNetAppNonDataIpInterfaceEventRelationship is a convenience function that returns MoMoRef wrapped in StorageNetAppNonDataIpInterfaceEventRelationship
func MoMoRefAsStorageNetAppNonDataIpInterfaceEventRelationship(v *MoMoRef) StorageNetAppNonDataIpInterfaceEventRelationship {
	return StorageNetAppNonDataIpInterfaceEventRelationship{
		MoMoRef: v,
	}
}

// StorageNetAppNonDataIpInterfaceEventAsStorageNetAppNonDataIpInterfaceEventRelationship is a convenience function that returns StorageNetAppNonDataIpInterfaceEvent wrapped in StorageNetAppNonDataIpInterfaceEventRelationship
func StorageNetAppNonDataIpInterfaceEventAsStorageNetAppNonDataIpInterfaceEventRelationship(v *StorageNetAppNonDataIpInterfaceEvent) StorageNetAppNonDataIpInterfaceEventRelationship {
	return StorageNetAppNonDataIpInterfaceEventRelationship{
		StorageNetAppNonDataIpInterfaceEvent: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageNetAppNonDataIpInterfaceEventRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal StorageNetAppNonDataIpInterfaceEventRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.NetAppNonDataIpInterfaceEvent'
	if jsonDict["ClassId"] == "storage.NetAppNonDataIpInterfaceEvent" {
		// try to unmarshal JSON data into StorageNetAppNonDataIpInterfaceEvent
		err = json.Unmarshal(data, &dst.StorageNetAppNonDataIpInterfaceEvent)
		if err == nil {
			return nil // data stored in dst.StorageNetAppNonDataIpInterfaceEvent, return on the first match
		} else {
			dst.StorageNetAppNonDataIpInterfaceEvent = nil
			return fmt.Errorf("failed to unmarshal StorageNetAppNonDataIpInterfaceEventRelationship as StorageNetAppNonDataIpInterfaceEvent: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageNetAppNonDataIpInterfaceEventRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.StorageNetAppNonDataIpInterfaceEvent != nil {
		return json.Marshal(&src.StorageNetAppNonDataIpInterfaceEvent)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageNetAppNonDataIpInterfaceEventRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.StorageNetAppNonDataIpInterfaceEvent != nil {
		return obj.StorageNetAppNonDataIpInterfaceEvent
	}

	// all schemas are nil
	return nil
}

type NullableStorageNetAppNonDataIpInterfaceEventRelationship struct {
	value *StorageNetAppNonDataIpInterfaceEventRelationship
	isSet bool
}

func (v NullableStorageNetAppNonDataIpInterfaceEventRelationship) Get() *StorageNetAppNonDataIpInterfaceEventRelationship {
	return v.value
}

func (v *NullableStorageNetAppNonDataIpInterfaceEventRelationship) Set(val *StorageNetAppNonDataIpInterfaceEventRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNonDataIpInterfaceEventRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNonDataIpInterfaceEventRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNonDataIpInterfaceEventRelationship(val *StorageNetAppNonDataIpInterfaceEventRelationship) *NullableStorageNetAppNonDataIpInterfaceEventRelationship {
	return &NullableStorageNetAppNonDataIpInterfaceEventRelationship{value: val, isSet: true}
}

func (v NullableStorageNetAppNonDataIpInterfaceEventRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNonDataIpInterfaceEventRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
