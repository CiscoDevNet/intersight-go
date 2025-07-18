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

// KubernetesClusterProfileRelationship - A relationship to the 'kubernetes.ClusterProfile' resource, or the expanded 'kubernetes.ClusterProfile' resource, or the 'null' value.
type KubernetesClusterProfileRelationship struct {
	KubernetesClusterProfile *KubernetesClusterProfile
	MoMoRef                  *MoMoRef
}

// KubernetesClusterProfileAsKubernetesClusterProfileRelationship is a convenience function that returns KubernetesClusterProfile wrapped in KubernetesClusterProfileRelationship
func KubernetesClusterProfileAsKubernetesClusterProfileRelationship(v *KubernetesClusterProfile) KubernetesClusterProfileRelationship {
	return KubernetesClusterProfileRelationship{
		KubernetesClusterProfile: v,
	}
}

// MoMoRefAsKubernetesClusterProfileRelationship is a convenience function that returns MoMoRef wrapped in KubernetesClusterProfileRelationship
func MoMoRefAsKubernetesClusterProfileRelationship(v *MoMoRef) KubernetesClusterProfileRelationship {
	return KubernetesClusterProfileRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KubernetesClusterProfileRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'kubernetes.ClusterProfile'
	if jsonDict["ClassId"] == "kubernetes.ClusterProfile" {
		// try to unmarshal JSON data into KubernetesClusterProfile
		err = json.Unmarshal(data, &dst.KubernetesClusterProfile)
		if err == nil {
			return nil // data stored in dst.KubernetesClusterProfile, return on the first match
		} else {
			dst.KubernetesClusterProfile = nil
			return fmt.Errorf("failed to unmarshal KubernetesClusterProfileRelationship as KubernetesClusterProfile: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal KubernetesClusterProfileRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KubernetesClusterProfileRelationship) MarshalJSON() ([]byte, error) {
	if src.KubernetesClusterProfile != nil {
		return json.Marshal(&src.KubernetesClusterProfile)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *KubernetesClusterProfileRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.KubernetesClusterProfile != nil {
		return obj.KubernetesClusterProfile
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableKubernetesClusterProfileRelationship struct {
	value *KubernetesClusterProfileRelationship
	isSet bool
}

func (v NullableKubernetesClusterProfileRelationship) Get() *KubernetesClusterProfileRelationship {
	return v.value
}

func (v *NullableKubernetesClusterProfileRelationship) Set(val *KubernetesClusterProfileRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterProfileRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterProfileRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterProfileRelationship(val *KubernetesClusterProfileRelationship) *NullableKubernetesClusterProfileRelationship {
	return &NullableKubernetesClusterProfileRelationship{value: val, isSet: true}
}

func (v NullableKubernetesClusterProfileRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterProfileRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
