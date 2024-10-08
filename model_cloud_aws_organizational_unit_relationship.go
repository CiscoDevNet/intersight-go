/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-18775
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// CloudAwsOrganizationalUnitRelationship - A relationship to the 'cloud.AwsOrganizationalUnit' resource, or the expanded 'cloud.AwsOrganizationalUnit' resource, or the 'null' value.
type CloudAwsOrganizationalUnitRelationship struct {
	CloudAwsOrganizationalUnit *CloudAwsOrganizationalUnit
	MoMoRef                    *MoMoRef
}

// CloudAwsOrganizationalUnitAsCloudAwsOrganizationalUnitRelationship is a convenience function that returns CloudAwsOrganizationalUnit wrapped in CloudAwsOrganizationalUnitRelationship
func CloudAwsOrganizationalUnitAsCloudAwsOrganizationalUnitRelationship(v *CloudAwsOrganizationalUnit) CloudAwsOrganizationalUnitRelationship {
	return CloudAwsOrganizationalUnitRelationship{
		CloudAwsOrganizationalUnit: v,
	}
}

// MoMoRefAsCloudAwsOrganizationalUnitRelationship is a convenience function that returns MoMoRef wrapped in CloudAwsOrganizationalUnitRelationship
func MoMoRefAsCloudAwsOrganizationalUnitRelationship(v *MoMoRef) CloudAwsOrganizationalUnitRelationship {
	return CloudAwsOrganizationalUnitRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CloudAwsOrganizationalUnitRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'cloud.AwsOrganizationalUnit'
	if jsonDict["ClassId"] == "cloud.AwsOrganizationalUnit" {
		// try to unmarshal JSON data into CloudAwsOrganizationalUnit
		err = json.Unmarshal(data, &dst.CloudAwsOrganizationalUnit)
		if err == nil {
			return nil // data stored in dst.CloudAwsOrganizationalUnit, return on the first match
		} else {
			dst.CloudAwsOrganizationalUnit = nil
			return fmt.Errorf("failed to unmarshal CloudAwsOrganizationalUnitRelationship as CloudAwsOrganizationalUnit: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal CloudAwsOrganizationalUnitRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CloudAwsOrganizationalUnitRelationship) MarshalJSON() ([]byte, error) {
	if src.CloudAwsOrganizationalUnit != nil {
		return json.Marshal(&src.CloudAwsOrganizationalUnit)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CloudAwsOrganizationalUnitRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CloudAwsOrganizationalUnit != nil {
		return obj.CloudAwsOrganizationalUnit
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableCloudAwsOrganizationalUnitRelationship struct {
	value *CloudAwsOrganizationalUnitRelationship
	isSet bool
}

func (v NullableCloudAwsOrganizationalUnitRelationship) Get() *CloudAwsOrganizationalUnitRelationship {
	return v.value
}

func (v *NullableCloudAwsOrganizationalUnitRelationship) Set(val *CloudAwsOrganizationalUnitRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsOrganizationalUnitRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsOrganizationalUnitRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsOrganizationalUnitRelationship(val *CloudAwsOrganizationalUnitRelationship) *NullableCloudAwsOrganizationalUnitRelationship {
	return &NullableCloudAwsOrganizationalUnitRelationship{value: val, isSet: true}
}

func (v NullableCloudAwsOrganizationalUnitRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsOrganizationalUnitRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
