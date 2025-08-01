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

// IppoolReservationRelationship - A relationship to the 'ippool.Reservation' resource, or the expanded 'ippool.Reservation' resource, or the 'null' value.
type IppoolReservationRelationship struct {
	IppoolReservation *IppoolReservation
	MoMoRef           *MoMoRef
}

// IppoolReservationAsIppoolReservationRelationship is a convenience function that returns IppoolReservation wrapped in IppoolReservationRelationship
func IppoolReservationAsIppoolReservationRelationship(v *IppoolReservation) IppoolReservationRelationship {
	return IppoolReservationRelationship{
		IppoolReservation: v,
	}
}

// MoMoRefAsIppoolReservationRelationship is a convenience function that returns MoMoRef wrapped in IppoolReservationRelationship
func MoMoRefAsIppoolReservationRelationship(v *MoMoRef) IppoolReservationRelationship {
	return IppoolReservationRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IppoolReservationRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'ippool.Reservation'
	if jsonDict["ClassId"] == "ippool.Reservation" {
		// try to unmarshal JSON data into IppoolReservation
		err = json.Unmarshal(data, &dst.IppoolReservation)
		if err == nil {
			return nil // data stored in dst.IppoolReservation, return on the first match
		} else {
			dst.IppoolReservation = nil
			return fmt.Errorf("failed to unmarshal IppoolReservationRelationship as IppoolReservation: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IppoolReservationRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IppoolReservationRelationship) MarshalJSON() ([]byte, error) {
	if src.IppoolReservation != nil {
		return json.Marshal(&src.IppoolReservation)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IppoolReservationRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IppoolReservation != nil {
		return obj.IppoolReservation
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableIppoolReservationRelationship struct {
	value *IppoolReservationRelationship
	isSet bool
}

func (v NullableIppoolReservationRelationship) Get() *IppoolReservationRelationship {
	return v.value
}

func (v *NullableIppoolReservationRelationship) Set(val *IppoolReservationRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolReservationRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolReservationRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolReservationRelationship(val *IppoolReservationRelationship) *NullableIppoolReservationRelationship {
	return &NullableIppoolReservationRelationship{value: val, isSet: true}
}

func (v NullableIppoolReservationRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolReservationRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
