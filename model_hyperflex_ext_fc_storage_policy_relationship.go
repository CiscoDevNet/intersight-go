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

// HyperflexExtFcStoragePolicyRelationship - A relationship to the 'hyperflex.ExtFcStoragePolicy' resource, or the expanded 'hyperflex.ExtFcStoragePolicy' resource, or the 'null' value.
type HyperflexExtFcStoragePolicyRelationship struct {
	HyperflexExtFcStoragePolicy *HyperflexExtFcStoragePolicy
	MoMoRef                     *MoMoRef
}

// HyperflexExtFcStoragePolicyAsHyperflexExtFcStoragePolicyRelationship is a convenience function that returns HyperflexExtFcStoragePolicy wrapped in HyperflexExtFcStoragePolicyRelationship
func HyperflexExtFcStoragePolicyAsHyperflexExtFcStoragePolicyRelationship(v *HyperflexExtFcStoragePolicy) HyperflexExtFcStoragePolicyRelationship {
	return HyperflexExtFcStoragePolicyRelationship{
		HyperflexExtFcStoragePolicy: v,
	}
}

// MoMoRefAsHyperflexExtFcStoragePolicyRelationship is a convenience function that returns MoMoRef wrapped in HyperflexExtFcStoragePolicyRelationship
func MoMoRefAsHyperflexExtFcStoragePolicyRelationship(v *MoMoRef) HyperflexExtFcStoragePolicyRelationship {
	return HyperflexExtFcStoragePolicyRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexExtFcStoragePolicyRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hyperflex.ExtFcStoragePolicy'
	if jsonDict["ClassId"] == "hyperflex.ExtFcStoragePolicy" {
		// try to unmarshal JSON data into HyperflexExtFcStoragePolicy
		err = json.Unmarshal(data, &dst.HyperflexExtFcStoragePolicy)
		if err == nil {
			return nil // data stored in dst.HyperflexExtFcStoragePolicy, return on the first match
		} else {
			dst.HyperflexExtFcStoragePolicy = nil
			return fmt.Errorf("failed to unmarshal HyperflexExtFcStoragePolicyRelationship as HyperflexExtFcStoragePolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexExtFcStoragePolicyRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexExtFcStoragePolicyRelationship) MarshalJSON() ([]byte, error) {
	if src.HyperflexExtFcStoragePolicy != nil {
		return json.Marshal(&src.HyperflexExtFcStoragePolicy)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HyperflexExtFcStoragePolicyRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HyperflexExtFcStoragePolicy != nil {
		return obj.HyperflexExtFcStoragePolicy
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHyperflexExtFcStoragePolicyRelationship struct {
	value *HyperflexExtFcStoragePolicyRelationship
	isSet bool
}

func (v NullableHyperflexExtFcStoragePolicyRelationship) Get() *HyperflexExtFcStoragePolicyRelationship {
	return v.value
}

func (v *NullableHyperflexExtFcStoragePolicyRelationship) Set(val *HyperflexExtFcStoragePolicyRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexExtFcStoragePolicyRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexExtFcStoragePolicyRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexExtFcStoragePolicyRelationship(val *HyperflexExtFcStoragePolicyRelationship) *NullableHyperflexExtFcStoragePolicyRelationship {
	return &NullableHyperflexExtFcStoragePolicyRelationship{value: val, isSet: true}
}

func (v NullableHyperflexExtFcStoragePolicyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexExtFcStoragePolicyRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
