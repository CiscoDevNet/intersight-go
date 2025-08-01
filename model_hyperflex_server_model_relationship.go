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

// HyperflexServerModelRelationship - A relationship to the 'hyperflex.ServerModel' resource, or the expanded 'hyperflex.ServerModel' resource, or the 'null' value.
type HyperflexServerModelRelationship struct {
	HyperflexServerModel *HyperflexServerModel
	MoMoRef              *MoMoRef
}

// HyperflexServerModelAsHyperflexServerModelRelationship is a convenience function that returns HyperflexServerModel wrapped in HyperflexServerModelRelationship
func HyperflexServerModelAsHyperflexServerModelRelationship(v *HyperflexServerModel) HyperflexServerModelRelationship {
	return HyperflexServerModelRelationship{
		HyperflexServerModel: v,
	}
}

// MoMoRefAsHyperflexServerModelRelationship is a convenience function that returns MoMoRef wrapped in HyperflexServerModelRelationship
func MoMoRefAsHyperflexServerModelRelationship(v *MoMoRef) HyperflexServerModelRelationship {
	return HyperflexServerModelRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexServerModelRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hyperflex.ServerModel'
	if jsonDict["ClassId"] == "hyperflex.ServerModel" {
		// try to unmarshal JSON data into HyperflexServerModel
		err = json.Unmarshal(data, &dst.HyperflexServerModel)
		if err == nil {
			return nil // data stored in dst.HyperflexServerModel, return on the first match
		} else {
			dst.HyperflexServerModel = nil
			return fmt.Errorf("failed to unmarshal HyperflexServerModelRelationship as HyperflexServerModel: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HyperflexServerModelRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexServerModelRelationship) MarshalJSON() ([]byte, error) {
	if src.HyperflexServerModel != nil {
		return json.Marshal(&src.HyperflexServerModel)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HyperflexServerModelRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HyperflexServerModel != nil {
		return obj.HyperflexServerModel
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHyperflexServerModelRelationship struct {
	value *HyperflexServerModelRelationship
	isSet bool
}

func (v NullableHyperflexServerModelRelationship) Get() *HyperflexServerModelRelationship {
	return v.value
}

func (v *NullableHyperflexServerModelRelationship) Set(val *HyperflexServerModelRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServerModelRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServerModelRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServerModelRelationship(val *HyperflexServerModelRelationship) *NullableHyperflexServerModelRelationship {
	return &NullableHyperflexServerModelRelationship{value: val, isSet: true}
}

func (v NullableHyperflexServerModelRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServerModelRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
