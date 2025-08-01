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

// StorageNetAppAggregateRelationship - A relationship to the 'storage.NetAppAggregate' resource, or the expanded 'storage.NetAppAggregate' resource, or the 'null' value.
type StorageNetAppAggregateRelationship struct {
	MoMoRef                *MoMoRef
	StorageNetAppAggregate *StorageNetAppAggregate
}

// MoMoRefAsStorageNetAppAggregateRelationship is a convenience function that returns MoMoRef wrapped in StorageNetAppAggregateRelationship
func MoMoRefAsStorageNetAppAggregateRelationship(v *MoMoRef) StorageNetAppAggregateRelationship {
	return StorageNetAppAggregateRelationship{
		MoMoRef: v,
	}
}

// StorageNetAppAggregateAsStorageNetAppAggregateRelationship is a convenience function that returns StorageNetAppAggregate wrapped in StorageNetAppAggregateRelationship
func StorageNetAppAggregateAsStorageNetAppAggregateRelationship(v *StorageNetAppAggregate) StorageNetAppAggregateRelationship {
	return StorageNetAppAggregateRelationship{
		StorageNetAppAggregate: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageNetAppAggregateRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal StorageNetAppAggregateRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.NetAppAggregate'
	if jsonDict["ClassId"] == "storage.NetAppAggregate" {
		// try to unmarshal JSON data into StorageNetAppAggregate
		err = json.Unmarshal(data, &dst.StorageNetAppAggregate)
		if err == nil {
			return nil // data stored in dst.StorageNetAppAggregate, return on the first match
		} else {
			dst.StorageNetAppAggregate = nil
			return fmt.Errorf("failed to unmarshal StorageNetAppAggregateRelationship as StorageNetAppAggregate: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageNetAppAggregateRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.StorageNetAppAggregate != nil {
		return json.Marshal(&src.StorageNetAppAggregate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageNetAppAggregateRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.StorageNetAppAggregate != nil {
		return obj.StorageNetAppAggregate
	}

	// all schemas are nil
	return nil
}

type NullableStorageNetAppAggregateRelationship struct {
	value *StorageNetAppAggregateRelationship
	isSet bool
}

func (v NullableStorageNetAppAggregateRelationship) Get() *StorageNetAppAggregateRelationship {
	return v.value
}

func (v *NullableStorageNetAppAggregateRelationship) Set(val *StorageNetAppAggregateRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppAggregateRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppAggregateRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppAggregateRelationship(val *StorageNetAppAggregateRelationship) *NullableStorageNetAppAggregateRelationship {
	return &NullableStorageNetAppAggregateRelationship{value: val, isSet: true}
}

func (v NullableStorageNetAppAggregateRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppAggregateRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
