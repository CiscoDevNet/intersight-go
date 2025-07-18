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

// AssetDeploymentRelationship - A relationship to the 'asset.Deployment' resource, or the expanded 'asset.Deployment' resource, or the 'null' value.
type AssetDeploymentRelationship struct {
	AssetDeployment *AssetDeployment
	MoMoRef         *MoMoRef
}

// AssetDeploymentAsAssetDeploymentRelationship is a convenience function that returns AssetDeployment wrapped in AssetDeploymentRelationship
func AssetDeploymentAsAssetDeploymentRelationship(v *AssetDeployment) AssetDeploymentRelationship {
	return AssetDeploymentRelationship{
		AssetDeployment: v,
	}
}

// MoMoRefAsAssetDeploymentRelationship is a convenience function that returns MoMoRef wrapped in AssetDeploymentRelationship
func MoMoRefAsAssetDeploymentRelationship(v *MoMoRef) AssetDeploymentRelationship {
	return AssetDeploymentRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AssetDeploymentRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'asset.Deployment'
	if jsonDict["ClassId"] == "asset.Deployment" {
		// try to unmarshal JSON data into AssetDeployment
		err = json.Unmarshal(data, &dst.AssetDeployment)
		if err == nil {
			return nil // data stored in dst.AssetDeployment, return on the first match
		} else {
			dst.AssetDeployment = nil
			return fmt.Errorf("failed to unmarshal AssetDeploymentRelationship as AssetDeployment: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal AssetDeploymentRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AssetDeploymentRelationship) MarshalJSON() ([]byte, error) {
	if src.AssetDeployment != nil {
		return json.Marshal(&src.AssetDeployment)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AssetDeploymentRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AssetDeployment != nil {
		return obj.AssetDeployment
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableAssetDeploymentRelationship struct {
	value *AssetDeploymentRelationship
	isSet bool
}

func (v NullableAssetDeploymentRelationship) Get() *AssetDeploymentRelationship {
	return v.value
}

func (v *NullableAssetDeploymentRelationship) Set(val *AssetDeploymentRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeploymentRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeploymentRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeploymentRelationship(val *AssetDeploymentRelationship) *NullableAssetDeploymentRelationship {
	return &NullableAssetDeploymentRelationship{value: val, isSet: true}
}

func (v NullableAssetDeploymentRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeploymentRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
