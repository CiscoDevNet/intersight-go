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

// ComputeBaseClusterRelationship - A relationship to the 'compute.BaseCluster' resource, or the expanded 'compute.BaseCluster' resource, or the 'null' value.
type ComputeBaseClusterRelationship struct {
	ComputeBaseCluster *ComputeBaseCluster
	MoMoRef            *MoMoRef
}

// ComputeBaseClusterAsComputeBaseClusterRelationship is a convenience function that returns ComputeBaseCluster wrapped in ComputeBaseClusterRelationship
func ComputeBaseClusterAsComputeBaseClusterRelationship(v *ComputeBaseCluster) ComputeBaseClusterRelationship {
	return ComputeBaseClusterRelationship{
		ComputeBaseCluster: v,
	}
}

// MoMoRefAsComputeBaseClusterRelationship is a convenience function that returns MoMoRef wrapped in ComputeBaseClusterRelationship
func MoMoRefAsComputeBaseClusterRelationship(v *MoMoRef) ComputeBaseClusterRelationship {
	return ComputeBaseClusterRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ComputeBaseClusterRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'compute.BaseCluster'
	if jsonDict["ClassId"] == "compute.BaseCluster" {
		// try to unmarshal JSON data into ComputeBaseCluster
		err = json.Unmarshal(data, &dst.ComputeBaseCluster)
		if err == nil {
			return nil // data stored in dst.ComputeBaseCluster, return on the first match
		} else {
			dst.ComputeBaseCluster = nil
			return fmt.Errorf("failed to unmarshal ComputeBaseClusterRelationship as ComputeBaseCluster: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ComputeBaseClusterRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ComputeBaseClusterRelationship) MarshalJSON() ([]byte, error) {
	if src.ComputeBaseCluster != nil {
		return json.Marshal(&src.ComputeBaseCluster)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ComputeBaseClusterRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ComputeBaseCluster != nil {
		return obj.ComputeBaseCluster
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableComputeBaseClusterRelationship struct {
	value *ComputeBaseClusterRelationship
	isSet bool
}

func (v NullableComputeBaseClusterRelationship) Get() *ComputeBaseClusterRelationship {
	return v.value
}

func (v *NullableComputeBaseClusterRelationship) Set(val *ComputeBaseClusterRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeBaseClusterRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeBaseClusterRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeBaseClusterRelationship(val *ComputeBaseClusterRelationship) *NullableComputeBaseClusterRelationship {
	return &NullableComputeBaseClusterRelationship{value: val, isSet: true}
}

func (v NullableComputeBaseClusterRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeBaseClusterRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
