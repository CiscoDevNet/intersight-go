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

// VnicFcQosPolicyRelationship - A relationship to the 'vnic.FcQosPolicy' resource, or the expanded 'vnic.FcQosPolicy' resource, or the 'null' value.
type VnicFcQosPolicyRelationship struct {
	MoMoRef         *MoMoRef
	VnicFcQosPolicy *VnicFcQosPolicy
}

// MoMoRefAsVnicFcQosPolicyRelationship is a convenience function that returns MoMoRef wrapped in VnicFcQosPolicyRelationship
func MoMoRefAsVnicFcQosPolicyRelationship(v *MoMoRef) VnicFcQosPolicyRelationship {
	return VnicFcQosPolicyRelationship{
		MoMoRef: v,
	}
}

// VnicFcQosPolicyAsVnicFcQosPolicyRelationship is a convenience function that returns VnicFcQosPolicy wrapped in VnicFcQosPolicyRelationship
func VnicFcQosPolicyAsVnicFcQosPolicyRelationship(v *VnicFcQosPolicy) VnicFcQosPolicyRelationship {
	return VnicFcQosPolicyRelationship{
		VnicFcQosPolicy: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VnicFcQosPolicyRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal VnicFcQosPolicyRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'vnic.FcQosPolicy'
	if jsonDict["ClassId"] == "vnic.FcQosPolicy" {
		// try to unmarshal JSON data into VnicFcQosPolicy
		err = json.Unmarshal(data, &dst.VnicFcQosPolicy)
		if err == nil {
			return nil // data stored in dst.VnicFcQosPolicy, return on the first match
		} else {
			dst.VnicFcQosPolicy = nil
			return fmt.Errorf("failed to unmarshal VnicFcQosPolicyRelationship as VnicFcQosPolicy: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VnicFcQosPolicyRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.VnicFcQosPolicy != nil {
		return json.Marshal(&src.VnicFcQosPolicy)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VnicFcQosPolicyRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.VnicFcQosPolicy != nil {
		return obj.VnicFcQosPolicy
	}

	// all schemas are nil
	return nil
}

type NullableVnicFcQosPolicyRelationship struct {
	value *VnicFcQosPolicyRelationship
	isSet bool
}

func (v NullableVnicFcQosPolicyRelationship) Get() *VnicFcQosPolicyRelationship {
	return v.value
}

func (v *NullableVnicFcQosPolicyRelationship) Set(val *VnicFcQosPolicyRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcQosPolicyRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcQosPolicyRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcQosPolicyRelationship(val *VnicFcQosPolicyRelationship) *NullableVnicFcQosPolicyRelationship {
	return &NullableVnicFcQosPolicyRelationship{value: val, isSet: true}
}

func (v NullableVnicFcQosPolicyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcQosPolicyRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
