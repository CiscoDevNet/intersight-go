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

// ServerConfigChangeDetailRelationship - A relationship to the 'server.ConfigChangeDetail' resource, or the expanded 'server.ConfigChangeDetail' resource, or the 'null' value.
type ServerConfigChangeDetailRelationship struct {
	MoMoRef                  *MoMoRef
	ServerConfigChangeDetail *ServerConfigChangeDetail
}

// MoMoRefAsServerConfigChangeDetailRelationship is a convenience function that returns MoMoRef wrapped in ServerConfigChangeDetailRelationship
func MoMoRefAsServerConfigChangeDetailRelationship(v *MoMoRef) ServerConfigChangeDetailRelationship {
	return ServerConfigChangeDetailRelationship{
		MoMoRef: v,
	}
}

// ServerConfigChangeDetailAsServerConfigChangeDetailRelationship is a convenience function that returns ServerConfigChangeDetail wrapped in ServerConfigChangeDetailRelationship
func ServerConfigChangeDetailAsServerConfigChangeDetailRelationship(v *ServerConfigChangeDetail) ServerConfigChangeDetailRelationship {
	return ServerConfigChangeDetailRelationship{
		ServerConfigChangeDetail: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServerConfigChangeDetailRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal ServerConfigChangeDetailRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'server.ConfigChangeDetail'
	if jsonDict["ClassId"] == "server.ConfigChangeDetail" {
		// try to unmarshal JSON data into ServerConfigChangeDetail
		err = json.Unmarshal(data, &dst.ServerConfigChangeDetail)
		if err == nil {
			return nil // data stored in dst.ServerConfigChangeDetail, return on the first match
		} else {
			dst.ServerConfigChangeDetail = nil
			return fmt.Errorf("failed to unmarshal ServerConfigChangeDetailRelationship as ServerConfigChangeDetail: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServerConfigChangeDetailRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.ServerConfigChangeDetail != nil {
		return json.Marshal(&src.ServerConfigChangeDetail)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServerConfigChangeDetailRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.ServerConfigChangeDetail != nil {
		return obj.ServerConfigChangeDetail
	}

	// all schemas are nil
	return nil
}

type NullableServerConfigChangeDetailRelationship struct {
	value *ServerConfigChangeDetailRelationship
	isSet bool
}

func (v NullableServerConfigChangeDetailRelationship) Get() *ServerConfigChangeDetailRelationship {
	return v.value
}

func (v *NullableServerConfigChangeDetailRelationship) Set(val *ServerConfigChangeDetailRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigChangeDetailRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigChangeDetailRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigChangeDetailRelationship(val *ServerConfigChangeDetailRelationship) *NullableServerConfigChangeDetailRelationship {
	return &NullableServerConfigChangeDetailRelationship{value: val, isSet: true}
}

func (v NullableServerConfigChangeDetailRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigChangeDetailRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
