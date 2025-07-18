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

// PartnerintegrationMetricsRelationship - A relationship to the 'partnerintegration.Metrics' resource, or the expanded 'partnerintegration.Metrics' resource, or the 'null' value.
type PartnerintegrationMetricsRelationship struct {
	MoMoRef                   *MoMoRef
	PartnerintegrationMetrics *PartnerintegrationMetrics
}

// MoMoRefAsPartnerintegrationMetricsRelationship is a convenience function that returns MoMoRef wrapped in PartnerintegrationMetricsRelationship
func MoMoRefAsPartnerintegrationMetricsRelationship(v *MoMoRef) PartnerintegrationMetricsRelationship {
	return PartnerintegrationMetricsRelationship{
		MoMoRef: v,
	}
}

// PartnerintegrationMetricsAsPartnerintegrationMetricsRelationship is a convenience function that returns PartnerintegrationMetrics wrapped in PartnerintegrationMetricsRelationship
func PartnerintegrationMetricsAsPartnerintegrationMetricsRelationship(v *PartnerintegrationMetrics) PartnerintegrationMetricsRelationship {
	return PartnerintegrationMetricsRelationship{
		PartnerintegrationMetrics: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PartnerintegrationMetricsRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal PartnerintegrationMetricsRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'partnerintegration.Metrics'
	if jsonDict["ClassId"] == "partnerintegration.Metrics" {
		// try to unmarshal JSON data into PartnerintegrationMetrics
		err = json.Unmarshal(data, &dst.PartnerintegrationMetrics)
		if err == nil {
			return nil // data stored in dst.PartnerintegrationMetrics, return on the first match
		} else {
			dst.PartnerintegrationMetrics = nil
			return fmt.Errorf("failed to unmarshal PartnerintegrationMetricsRelationship as PartnerintegrationMetrics: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PartnerintegrationMetricsRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.PartnerintegrationMetrics != nil {
		return json.Marshal(&src.PartnerintegrationMetrics)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PartnerintegrationMetricsRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.PartnerintegrationMetrics != nil {
		return obj.PartnerintegrationMetrics
	}

	// all schemas are nil
	return nil
}

type NullablePartnerintegrationMetricsRelationship struct {
	value *PartnerintegrationMetricsRelationship
	isSet bool
}

func (v NullablePartnerintegrationMetricsRelationship) Get() *PartnerintegrationMetricsRelationship {
	return v.value
}

func (v *NullablePartnerintegrationMetricsRelationship) Set(val *PartnerintegrationMetricsRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerintegrationMetricsRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerintegrationMetricsRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerintegrationMetricsRelationship(val *PartnerintegrationMetricsRelationship) *NullablePartnerintegrationMetricsRelationship {
	return &NullablePartnerintegrationMetricsRelationship{value: val, isSet: true}
}

func (v NullablePartnerintegrationMetricsRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerintegrationMetricsRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
