/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2025051220
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// FunctionsFunctionRelationship - A relationship to the 'functions.Function' resource, or the expanded 'functions.Function' resource, or the 'null' value.
type FunctionsFunctionRelationship struct {
	FunctionsFunction *FunctionsFunction
	MoMoRef           *MoMoRef
}

// FunctionsFunctionAsFunctionsFunctionRelationship is a convenience function that returns FunctionsFunction wrapped in FunctionsFunctionRelationship
func FunctionsFunctionAsFunctionsFunctionRelationship(v *FunctionsFunction) FunctionsFunctionRelationship {
	return FunctionsFunctionRelationship{
		FunctionsFunction: v,
	}
}

// MoMoRefAsFunctionsFunctionRelationship is a convenience function that returns MoMoRef wrapped in FunctionsFunctionRelationship
func MoMoRefAsFunctionsFunctionRelationship(v *MoMoRef) FunctionsFunctionRelationship {
	return FunctionsFunctionRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FunctionsFunctionRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'functions.Function'
	if jsonDict["ClassId"] == "functions.Function" {
		// try to unmarshal JSON data into FunctionsFunction
		err = json.Unmarshal(data, &dst.FunctionsFunction)
		if err == nil {
			return nil // data stored in dst.FunctionsFunction, return on the first match
		} else {
			dst.FunctionsFunction = nil
			return fmt.Errorf("failed to unmarshal FunctionsFunctionRelationship as FunctionsFunction: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal FunctionsFunctionRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FunctionsFunctionRelationship) MarshalJSON() ([]byte, error) {
	if src.FunctionsFunction != nil {
		return json.Marshal(&src.FunctionsFunction)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FunctionsFunctionRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FunctionsFunction != nil {
		return obj.FunctionsFunction
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableFunctionsFunctionRelationship struct {
	value *FunctionsFunctionRelationship
	isSet bool
}

func (v NullableFunctionsFunctionRelationship) Get() *FunctionsFunctionRelationship {
	return v.value
}

func (v *NullableFunctionsFunctionRelationship) Set(val *FunctionsFunctionRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionsFunctionRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionsFunctionRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionsFunctionRelationship(val *FunctionsFunctionRelationship) *NullableFunctionsFunctionRelationship {
	return &NullableFunctionsFunctionRelationship{value: val, isSet: true}
}

func (v NullableFunctionsFunctionRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionsFunctionRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
