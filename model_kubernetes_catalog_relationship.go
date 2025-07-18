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

// KubernetesCatalogRelationship - A relationship to the 'kubernetes.Catalog' resource, or the expanded 'kubernetes.Catalog' resource, or the 'null' value.
type KubernetesCatalogRelationship struct {
	KubernetesCatalog *KubernetesCatalog
	MoMoRef           *MoMoRef
}

// KubernetesCatalogAsKubernetesCatalogRelationship is a convenience function that returns KubernetesCatalog wrapped in KubernetesCatalogRelationship
func KubernetesCatalogAsKubernetesCatalogRelationship(v *KubernetesCatalog) KubernetesCatalogRelationship {
	return KubernetesCatalogRelationship{
		KubernetesCatalog: v,
	}
}

// MoMoRefAsKubernetesCatalogRelationship is a convenience function that returns MoMoRef wrapped in KubernetesCatalogRelationship
func MoMoRefAsKubernetesCatalogRelationship(v *MoMoRef) KubernetesCatalogRelationship {
	return KubernetesCatalogRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KubernetesCatalogRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'kubernetes.Catalog'
	if jsonDict["ClassId"] == "kubernetes.Catalog" {
		// try to unmarshal JSON data into KubernetesCatalog
		err = json.Unmarshal(data, &dst.KubernetesCatalog)
		if err == nil {
			return nil // data stored in dst.KubernetesCatalog, return on the first match
		} else {
			dst.KubernetesCatalog = nil
			return fmt.Errorf("failed to unmarshal KubernetesCatalogRelationship as KubernetesCatalog: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal KubernetesCatalogRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KubernetesCatalogRelationship) MarshalJSON() ([]byte, error) {
	if src.KubernetesCatalog != nil {
		return json.Marshal(&src.KubernetesCatalog)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *KubernetesCatalogRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.KubernetesCatalog != nil {
		return obj.KubernetesCatalog
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableKubernetesCatalogRelationship struct {
	value *KubernetesCatalogRelationship
	isSet bool
}

func (v NullableKubernetesCatalogRelationship) Get() *KubernetesCatalogRelationship {
	return v.value
}

func (v *NullableKubernetesCatalogRelationship) Set(val *KubernetesCatalogRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCatalogRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCatalogRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCatalogRelationship(val *KubernetesCatalogRelationship) *NullableKubernetesCatalogRelationship {
	return &NullableKubernetesCatalogRelationship{value: val, isSet: true}
}

func (v NullableKubernetesCatalogRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCatalogRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
