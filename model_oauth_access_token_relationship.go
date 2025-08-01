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

// OauthAccessTokenRelationship - A relationship to the 'oauth.AccessToken' resource, or the expanded 'oauth.AccessToken' resource, or the 'null' value.
type OauthAccessTokenRelationship struct {
	MoMoRef          *MoMoRef
	OauthAccessToken *OauthAccessToken
}

// MoMoRefAsOauthAccessTokenRelationship is a convenience function that returns MoMoRef wrapped in OauthAccessTokenRelationship
func MoMoRefAsOauthAccessTokenRelationship(v *MoMoRef) OauthAccessTokenRelationship {
	return OauthAccessTokenRelationship{
		MoMoRef: v,
	}
}

// OauthAccessTokenAsOauthAccessTokenRelationship is a convenience function that returns OauthAccessToken wrapped in OauthAccessTokenRelationship
func OauthAccessTokenAsOauthAccessTokenRelationship(v *OauthAccessToken) OauthAccessTokenRelationship {
	return OauthAccessTokenRelationship{
		OauthAccessToken: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OauthAccessTokenRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal OauthAccessTokenRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'oauth.AccessToken'
	if jsonDict["ClassId"] == "oauth.AccessToken" {
		// try to unmarshal JSON data into OauthAccessToken
		err = json.Unmarshal(data, &dst.OauthAccessToken)
		if err == nil {
			return nil // data stored in dst.OauthAccessToken, return on the first match
		} else {
			dst.OauthAccessToken = nil
			return fmt.Errorf("failed to unmarshal OauthAccessTokenRelationship as OauthAccessToken: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OauthAccessTokenRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.OauthAccessToken != nil {
		return json.Marshal(&src.OauthAccessToken)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OauthAccessTokenRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.OauthAccessToken != nil {
		return obj.OauthAccessToken
	}

	// all schemas are nil
	return nil
}

type NullableOauthAccessTokenRelationship struct {
	value *OauthAccessTokenRelationship
	isSet bool
}

func (v NullableOauthAccessTokenRelationship) Get() *OauthAccessTokenRelationship {
	return v.value
}

func (v *NullableOauthAccessTokenRelationship) Set(val *OauthAccessTokenRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableOauthAccessTokenRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableOauthAccessTokenRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauthAccessTokenRelationship(val *OauthAccessTokenRelationship) *NullableOauthAccessTokenRelationship {
	return &NullableOauthAccessTokenRelationship{value: val, isSet: true}
}

func (v NullableOauthAccessTokenRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauthAccessTokenRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
