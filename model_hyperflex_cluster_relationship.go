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

// HyperflexClusterRelationship - A relationship to the 'hyperflex.Cluster' resource, or the expanded 'hyperflex.Cluster' resource, or the 'null' value.
type HyperflexClusterRelationship struct {
	HyperflexCluster *HyperflexCluster
	MoMoRef          *MoMoRef
}

// HyperflexClusterAsHyperflexClusterRelationship is a convenience function that returns HyperflexCluster wrapped in HyperflexClusterRelationship
func HyperflexClusterAsHyperflexClusterRelationship(v *HyperflexCluster) HyperflexClusterRelationship {
	return HyperflexClusterRelationship{
		HyperflexCluster: v,
	}
}

// MoMoRefAsHyperflexClusterRelationship is a convenience function that returns MoMoRef wrapped in HyperflexClusterRelationship
func MoMoRefAsHyperflexClusterRelationship(v *MoMoRef) HyperflexClusterRelationship {
	return HyperflexClusterRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexClusterRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hyperflex.Cluster'
	if jsonDict["ClassId"] == "hyperflex.Cluster" {
		// try to unmarshal JSON data into HyperflexCluster
		err = json.Unmarshal(data, &dst.HyperflexCluster)
		if err == nil {
			return nil // data stored in dst.HyperflexCluster, return on the first match
		} else {
			dst.HyperflexCluster = nil
			return fmt.Errorf("failed to unmarshal HyperflexClusterRelationship as HyperflexCluster: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexClusterRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexClusterRelationship) MarshalJSON() ([]byte, error) {
	if src.HyperflexCluster != nil {
		return json.Marshal(&src.HyperflexCluster)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HyperflexClusterRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HyperflexCluster != nil {
		return obj.HyperflexCluster
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHyperflexClusterRelationship struct {
	value *HyperflexClusterRelationship
	isSet bool
}

func (v NullableHyperflexClusterRelationship) Get() *HyperflexClusterRelationship {
	return v.value
}

func (v *NullableHyperflexClusterRelationship) Set(val *HyperflexClusterRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterRelationship(val *HyperflexClusterRelationship) *NullableHyperflexClusterRelationship {
	return &NullableHyperflexClusterRelationship{value: val, isSet: true}
}

func (v NullableHyperflexClusterRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
