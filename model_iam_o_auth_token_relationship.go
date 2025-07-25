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

// IamOAuthTokenRelationship - A relationship to the 'iam.OAuthToken' resource, or the expanded 'iam.OAuthToken' resource, or the 'null' value.
type IamOAuthTokenRelationship struct {
	IamOAuthToken *IamOAuthToken
	MoMoRef       *MoMoRef
}

// IamOAuthTokenAsIamOAuthTokenRelationship is a convenience function that returns IamOAuthToken wrapped in IamOAuthTokenRelationship
func IamOAuthTokenAsIamOAuthTokenRelationship(v *IamOAuthToken) IamOAuthTokenRelationship {
	return IamOAuthTokenRelationship{
		IamOAuthToken: v,
	}
}

// MoMoRefAsIamOAuthTokenRelationship is a convenience function that returns MoMoRef wrapped in IamOAuthTokenRelationship
func MoMoRefAsIamOAuthTokenRelationship(v *MoMoRef) IamOAuthTokenRelationship {
	return IamOAuthTokenRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamOAuthTokenRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'iam.OAuthToken'
	if jsonDict["ClassId"] == "iam.OAuthToken" {
		// try to unmarshal JSON data into IamOAuthToken
		err = json.Unmarshal(data, &dst.IamOAuthToken)
		if err == nil {
			return nil // data stored in dst.IamOAuthToken, return on the first match
		} else {
			dst.IamOAuthToken = nil
			return fmt.Errorf("failed to unmarshal IamOAuthTokenRelationship as IamOAuthToken: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamOAuthTokenRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamOAuthTokenRelationship) MarshalJSON() ([]byte, error) {
	if src.IamOAuthToken != nil {
		return json.Marshal(&src.IamOAuthToken)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IamOAuthTokenRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IamOAuthToken != nil {
		return obj.IamOAuthToken
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableIamOAuthTokenRelationship struct {
	value *IamOAuthTokenRelationship
	isSet bool
}

func (v NullableIamOAuthTokenRelationship) Get() *IamOAuthTokenRelationship {
	return v.value
}

func (v *NullableIamOAuthTokenRelationship) Set(val *IamOAuthTokenRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableIamOAuthTokenRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableIamOAuthTokenRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamOAuthTokenRelationship(val *IamOAuthTokenRelationship) *NullableIamOAuthTokenRelationship {
	return &NullableIamOAuthTokenRelationship{value: val, isSet: true}
}

func (v NullableIamOAuthTokenRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamOAuthTokenRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
