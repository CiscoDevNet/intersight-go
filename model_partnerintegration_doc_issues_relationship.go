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

// PartnerintegrationDocIssuesRelationship - A relationship to the 'partnerintegration.DocIssues' resource, or the expanded 'partnerintegration.DocIssues' resource, or the 'null' value.
type PartnerintegrationDocIssuesRelationship struct {
	MoMoRef                     *MoMoRef
	PartnerintegrationDocIssues *PartnerintegrationDocIssues
}

// MoMoRefAsPartnerintegrationDocIssuesRelationship is a convenience function that returns MoMoRef wrapped in PartnerintegrationDocIssuesRelationship
func MoMoRefAsPartnerintegrationDocIssuesRelationship(v *MoMoRef) PartnerintegrationDocIssuesRelationship {
	return PartnerintegrationDocIssuesRelationship{
		MoMoRef: v,
	}
}

// PartnerintegrationDocIssuesAsPartnerintegrationDocIssuesRelationship is a convenience function that returns PartnerintegrationDocIssues wrapped in PartnerintegrationDocIssuesRelationship
func PartnerintegrationDocIssuesAsPartnerintegrationDocIssuesRelationship(v *PartnerintegrationDocIssues) PartnerintegrationDocIssuesRelationship {
	return PartnerintegrationDocIssuesRelationship{
		PartnerintegrationDocIssues: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PartnerintegrationDocIssuesRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal PartnerintegrationDocIssuesRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'partnerintegration.DocIssues'
	if jsonDict["ClassId"] == "partnerintegration.DocIssues" {
		// try to unmarshal JSON data into PartnerintegrationDocIssues
		err = json.Unmarshal(data, &dst.PartnerintegrationDocIssues)
		if err == nil {
			return nil // data stored in dst.PartnerintegrationDocIssues, return on the first match
		} else {
			dst.PartnerintegrationDocIssues = nil
			return fmt.Errorf("failed to unmarshal PartnerintegrationDocIssuesRelationship as PartnerintegrationDocIssues: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PartnerintegrationDocIssuesRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.PartnerintegrationDocIssues != nil {
		return json.Marshal(&src.PartnerintegrationDocIssues)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PartnerintegrationDocIssuesRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.PartnerintegrationDocIssues != nil {
		return obj.PartnerintegrationDocIssues
	}

	// all schemas are nil
	return nil
}

type NullablePartnerintegrationDocIssuesRelationship struct {
	value *PartnerintegrationDocIssuesRelationship
	isSet bool
}

func (v NullablePartnerintegrationDocIssuesRelationship) Get() *PartnerintegrationDocIssuesRelationship {
	return v.value
}

func (v *NullablePartnerintegrationDocIssuesRelationship) Set(val *PartnerintegrationDocIssuesRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerintegrationDocIssuesRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerintegrationDocIssuesRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerintegrationDocIssuesRelationship(val *PartnerintegrationDocIssuesRelationship) *NullablePartnerintegrationDocIssuesRelationship {
	return &NullablePartnerintegrationDocIssuesRelationship{value: val, isSet: true}
}

func (v NullablePartnerintegrationDocIssuesRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerintegrationDocIssuesRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
