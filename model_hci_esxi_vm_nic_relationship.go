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

// HciEsxiVmNicRelationship - A relationship to the 'hci.EsxiVmNic' resource, or the expanded 'hci.EsxiVmNic' resource, or the 'null' value.
type HciEsxiVmNicRelationship struct {
	HciEsxiVmNic *HciEsxiVmNic
	MoMoRef      *MoMoRef
}

// HciEsxiVmNicAsHciEsxiVmNicRelationship is a convenience function that returns HciEsxiVmNic wrapped in HciEsxiVmNicRelationship
func HciEsxiVmNicAsHciEsxiVmNicRelationship(v *HciEsxiVmNic) HciEsxiVmNicRelationship {
	return HciEsxiVmNicRelationship{
		HciEsxiVmNic: v,
	}
}

// MoMoRefAsHciEsxiVmNicRelationship is a convenience function that returns MoMoRef wrapped in HciEsxiVmNicRelationship
func MoMoRefAsHciEsxiVmNicRelationship(v *MoMoRef) HciEsxiVmNicRelationship {
	return HciEsxiVmNicRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HciEsxiVmNicRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'hci.EsxiVmNic'
	if jsonDict["ClassId"] == "hci.EsxiVmNic" {
		// try to unmarshal JSON data into HciEsxiVmNic
		err = json.Unmarshal(data, &dst.HciEsxiVmNic)
		if err == nil {
			return nil // data stored in dst.HciEsxiVmNic, return on the first match
		} else {
			dst.HciEsxiVmNic = nil
			return fmt.Errorf("failed to unmarshal HciEsxiVmNicRelationship as HciEsxiVmNic: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal HciEsxiVmNicRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HciEsxiVmNicRelationship) MarshalJSON() ([]byte, error) {
	if src.HciEsxiVmNic != nil {
		return json.Marshal(&src.HciEsxiVmNic)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HciEsxiVmNicRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HciEsxiVmNic != nil {
		return obj.HciEsxiVmNic
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHciEsxiVmNicRelationship struct {
	value *HciEsxiVmNicRelationship
	isSet bool
}

func (v NullableHciEsxiVmNicRelationship) Get() *HciEsxiVmNicRelationship {
	return v.value
}

func (v *NullableHciEsxiVmNicRelationship) Set(val *HciEsxiVmNicRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHciEsxiVmNicRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHciEsxiVmNicRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHciEsxiVmNicRelationship(val *HciEsxiVmNicRelationship) *NullableHciEsxiVmNicRelationship {
	return &NullableHciEsxiVmNicRelationship{value: val, isSet: true}
}

func (v NullableHciEsxiVmNicRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHciEsxiVmNicRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
