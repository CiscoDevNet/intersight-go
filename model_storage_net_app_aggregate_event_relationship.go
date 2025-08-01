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

// StorageNetAppAggregateEventRelationship - A relationship to the 'storage.NetAppAggregateEvent' resource, or the expanded 'storage.NetAppAggregateEvent' resource, or the 'null' value.
type StorageNetAppAggregateEventRelationship struct {
	MoMoRef                     *MoMoRef
	StorageNetAppAggregateEvent *StorageNetAppAggregateEvent
}

// MoMoRefAsStorageNetAppAggregateEventRelationship is a convenience function that returns MoMoRef wrapped in StorageNetAppAggregateEventRelationship
func MoMoRefAsStorageNetAppAggregateEventRelationship(v *MoMoRef) StorageNetAppAggregateEventRelationship {
	return StorageNetAppAggregateEventRelationship{
		MoMoRef: v,
	}
}

// StorageNetAppAggregateEventAsStorageNetAppAggregateEventRelationship is a convenience function that returns StorageNetAppAggregateEvent wrapped in StorageNetAppAggregateEventRelationship
func StorageNetAppAggregateEventAsStorageNetAppAggregateEventRelationship(v *StorageNetAppAggregateEvent) StorageNetAppAggregateEventRelationship {
	return StorageNetAppAggregateEventRelationship{
		StorageNetAppAggregateEvent: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageNetAppAggregateEventRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal StorageNetAppAggregateEventRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.NetAppAggregateEvent'
	if jsonDict["ClassId"] == "storage.NetAppAggregateEvent" {
		// try to unmarshal JSON data into StorageNetAppAggregateEvent
		err = json.Unmarshal(data, &dst.StorageNetAppAggregateEvent)
		if err == nil {
			return nil // data stored in dst.StorageNetAppAggregateEvent, return on the first match
		} else {
			dst.StorageNetAppAggregateEvent = nil
			return fmt.Errorf("failed to unmarshal StorageNetAppAggregateEventRelationship as StorageNetAppAggregateEvent: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageNetAppAggregateEventRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.StorageNetAppAggregateEvent != nil {
		return json.Marshal(&src.StorageNetAppAggregateEvent)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageNetAppAggregateEventRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.StorageNetAppAggregateEvent != nil {
		return obj.StorageNetAppAggregateEvent
	}

	// all schemas are nil
	return nil
}

type NullableStorageNetAppAggregateEventRelationship struct {
	value *StorageNetAppAggregateEventRelationship
	isSet bool
}

func (v NullableStorageNetAppAggregateEventRelationship) Get() *StorageNetAppAggregateEventRelationship {
	return v.value
}

func (v *NullableStorageNetAppAggregateEventRelationship) Set(val *StorageNetAppAggregateEventRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppAggregateEventRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppAggregateEventRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppAggregateEventRelationship(val *StorageNetAppAggregateEventRelationship) *NullableStorageNetAppAggregateEventRelationship {
	return &NullableStorageNetAppAggregateEventRelationship{value: val, isSet: true}
}

func (v NullableStorageNetAppAggregateEventRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppAggregateEventRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
