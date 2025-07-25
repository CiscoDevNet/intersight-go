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

// HyperflexClusterNetworkPolicyRelationship - A relationship to the 'hyperflex.ClusterNetworkPolicy' resource, or the expanded 'hyperflex.ClusterNetworkPolicy' resource, or the 'null' value.
type HyperflexClusterNetworkPolicyRelationship struct {
	HyperflexClusterNetworkPolicy *HyperflexClusterNetworkPolicy
	MoMoRef                       *MoMoRef
}

// HyperflexClusterNetworkPolicyAsHyperflexClusterNetworkPolicyRelationship is a convenience function that returns HyperflexClusterNetworkPolicy wrapped in HyperflexClusterNetworkPolicyRelationship
func HyperflexClusterNetworkPolicyAsHyperflexClusterNetworkPolicyRelationship(v *HyperflexClusterNetworkPolicy) HyperflexClusterNetworkPolicyRelationship {
	return HyperflexClusterNetworkPolicyRelationship{
		HyperflexClusterNetworkPolicy: v,
	}
}

// MoMoRefAsHyperflexClusterNetworkPolicyRelationship is a convenience function that returns MoMoRef wrapped in HyperflexClusterNetworkPolicyRelationship
func MoMoRefAsHyperflexClusterNetworkPolicyRelationship(v *MoMoRef) HyperflexClusterNetworkPolicyRelationship {
	return HyperflexClusterNetworkPolicyRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexClusterNetworkPolicyRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hyperflex.ClusterNetworkPolicy'
	if jsonDict["ClassId"] == "hyperflex.ClusterNetworkPolicy" {
		// try to unmarshal JSON data into HyperflexClusterNetworkPolicy
		err = json.Unmarshal(data, &dst.HyperflexClusterNetworkPolicy)
		if err == nil {
			return nil // data stored in dst.HyperflexClusterNetworkPolicy, return on the first match
		} else {
			dst.HyperflexClusterNetworkPolicy = nil
			return fmt.Errorf("failed to unmarshal HyperflexClusterNetworkPolicyRelationship as HyperflexClusterNetworkPolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexClusterNetworkPolicyRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexClusterNetworkPolicyRelationship) MarshalJSON() ([]byte, error) {
	if src.HyperflexClusterNetworkPolicy != nil {
		return json.Marshal(&src.HyperflexClusterNetworkPolicy)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HyperflexClusterNetworkPolicyRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HyperflexClusterNetworkPolicy != nil {
		return obj.HyperflexClusterNetworkPolicy
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHyperflexClusterNetworkPolicyRelationship struct {
	value *HyperflexClusterNetworkPolicyRelationship
	isSet bool
}

func (v NullableHyperflexClusterNetworkPolicyRelationship) Get() *HyperflexClusterNetworkPolicyRelationship {
	return v.value
}

func (v *NullableHyperflexClusterNetworkPolicyRelationship) Set(val *HyperflexClusterNetworkPolicyRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterNetworkPolicyRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterNetworkPolicyRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterNetworkPolicyRelationship(val *HyperflexClusterNetworkPolicyRelationship) *NullableHyperflexClusterNetworkPolicyRelationship {
	return &NullableHyperflexClusterNetworkPolicyRelationship{value: val, isSet: true}
}

func (v NullableHyperflexClusterNetworkPolicyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterNetworkPolicyRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
