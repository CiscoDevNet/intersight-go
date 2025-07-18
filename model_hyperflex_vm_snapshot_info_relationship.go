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

// HyperflexVmSnapshotInfoRelationship - A relationship to the 'hyperflex.VmSnapshotInfo' resource, or the expanded 'hyperflex.VmSnapshotInfo' resource, or the 'null' value.
type HyperflexVmSnapshotInfoRelationship struct {
	HyperflexVmSnapshotInfo *HyperflexVmSnapshotInfo
	MoMoRef                 *MoMoRef
}

// HyperflexVmSnapshotInfoAsHyperflexVmSnapshotInfoRelationship is a convenience function that returns HyperflexVmSnapshotInfo wrapped in HyperflexVmSnapshotInfoRelationship
func HyperflexVmSnapshotInfoAsHyperflexVmSnapshotInfoRelationship(v *HyperflexVmSnapshotInfo) HyperflexVmSnapshotInfoRelationship {
	return HyperflexVmSnapshotInfoRelationship{
		HyperflexVmSnapshotInfo: v,
	}
}

// MoMoRefAsHyperflexVmSnapshotInfoRelationship is a convenience function that returns MoMoRef wrapped in HyperflexVmSnapshotInfoRelationship
func MoMoRefAsHyperflexVmSnapshotInfoRelationship(v *MoMoRef) HyperflexVmSnapshotInfoRelationship {
	return HyperflexVmSnapshotInfoRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexVmSnapshotInfoRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hyperflex.VmSnapshotInfo'
	if jsonDict["ClassId"] == "hyperflex.VmSnapshotInfo" {
		// try to unmarshal JSON data into HyperflexVmSnapshotInfo
		err = json.Unmarshal(data, &dst.HyperflexVmSnapshotInfo)
		if err == nil {
			return nil // data stored in dst.HyperflexVmSnapshotInfo, return on the first match
		} else {
			dst.HyperflexVmSnapshotInfo = nil
			return fmt.Errorf("failed to unmarshal HyperflexVmSnapshotInfoRelationship as HyperflexVmSnapshotInfo: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexVmSnapshotInfoRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexVmSnapshotInfoRelationship) MarshalJSON() ([]byte, error) {
	if src.HyperflexVmSnapshotInfo != nil {
		return json.Marshal(&src.HyperflexVmSnapshotInfo)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HyperflexVmSnapshotInfoRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HyperflexVmSnapshotInfo != nil {
		return obj.HyperflexVmSnapshotInfo
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHyperflexVmSnapshotInfoRelationship struct {
	value *HyperflexVmSnapshotInfoRelationship
	isSet bool
}

func (v NullableHyperflexVmSnapshotInfoRelationship) Get() *HyperflexVmSnapshotInfoRelationship {
	return v.value
}

func (v *NullableHyperflexVmSnapshotInfoRelationship) Set(val *HyperflexVmSnapshotInfoRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVmSnapshotInfoRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVmSnapshotInfoRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVmSnapshotInfoRelationship(val *HyperflexVmSnapshotInfoRelationship) *NullableHyperflexVmSnapshotInfoRelationship {
	return &NullableHyperflexVmSnapshotInfoRelationship{value: val, isSet: true}
}

func (v NullableHyperflexVmSnapshotInfoRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVmSnapshotInfoRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
