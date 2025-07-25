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

// StorageFlexUtilControllerRelationship - A relationship to the 'storage.FlexUtilController' resource, or the expanded 'storage.FlexUtilController' resource, or the 'null' value.
type StorageFlexUtilControllerRelationship struct {
	MoMoRef                   *MoMoRef
	StorageFlexUtilController *StorageFlexUtilController
}

// MoMoRefAsStorageFlexUtilControllerRelationship is a convenience function that returns MoMoRef wrapped in StorageFlexUtilControllerRelationship
func MoMoRefAsStorageFlexUtilControllerRelationship(v *MoMoRef) StorageFlexUtilControllerRelationship {
	return StorageFlexUtilControllerRelationship{
		MoMoRef: v,
	}
}

// StorageFlexUtilControllerAsStorageFlexUtilControllerRelationship is a convenience function that returns StorageFlexUtilController wrapped in StorageFlexUtilControllerRelationship
func StorageFlexUtilControllerAsStorageFlexUtilControllerRelationship(v *StorageFlexUtilController) StorageFlexUtilControllerRelationship {
	return StorageFlexUtilControllerRelationship{
		StorageFlexUtilController: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageFlexUtilControllerRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal StorageFlexUtilControllerRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.FlexUtilController'
	if jsonDict["ClassId"] == "storage.FlexUtilController" {
		// try to unmarshal JSON data into StorageFlexUtilController
		err = json.Unmarshal(data, &dst.StorageFlexUtilController)
		if err == nil {
			return nil // data stored in dst.StorageFlexUtilController, return on the first match
		} else {
			dst.StorageFlexUtilController = nil
			return fmt.Errorf("failed to unmarshal StorageFlexUtilControllerRelationship as StorageFlexUtilController: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageFlexUtilControllerRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.StorageFlexUtilController != nil {
		return json.Marshal(&src.StorageFlexUtilController)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageFlexUtilControllerRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.StorageFlexUtilController != nil {
		return obj.StorageFlexUtilController
	}

	// all schemas are nil
	return nil
}

type NullableStorageFlexUtilControllerRelationship struct {
	value *StorageFlexUtilControllerRelationship
	isSet bool
}

func (v NullableStorageFlexUtilControllerRelationship) Get() *StorageFlexUtilControllerRelationship {
	return v.value
}

func (v *NullableStorageFlexUtilControllerRelationship) Set(val *StorageFlexUtilControllerRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexUtilControllerRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexUtilControllerRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexUtilControllerRelationship(val *StorageFlexUtilControllerRelationship) *NullableStorageFlexUtilControllerRelationship {
	return &NullableStorageFlexUtilControllerRelationship{value: val, isSet: true}
}

func (v NullableStorageFlexUtilControllerRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexUtilControllerRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
