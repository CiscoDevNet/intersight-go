/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-15473
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the CapabilityProcessorUnitUpdateConstraintMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityProcessorUnitUpdateConstraintMeta{}

// CapabilityProcessorUnitUpdateConstraintMeta Internal meta-data to enable processor unit update related constraints.
type CapabilityProcessorUnitUpdateConstraintMeta struct {
	CapabilityUpdateConstraintMeta
	AdditionalProperties map[string]interface{}
}

type _CapabilityProcessorUnitUpdateConstraintMeta CapabilityProcessorUnitUpdateConstraintMeta

// NewCapabilityProcessorUnitUpdateConstraintMeta instantiates a new CapabilityProcessorUnitUpdateConstraintMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityProcessorUnitUpdateConstraintMeta(classId string, objectType string) *CapabilityProcessorUnitUpdateConstraintMeta {
	this := CapabilityProcessorUnitUpdateConstraintMeta{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityProcessorUnitUpdateConstraintMetaWithDefaults instantiates a new CapabilityProcessorUnitUpdateConstraintMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityProcessorUnitUpdateConstraintMetaWithDefaults() *CapabilityProcessorUnitUpdateConstraintMeta {
	this := CapabilityProcessorUnitUpdateConstraintMeta{}
	return &this
}

func (o CapabilityProcessorUnitUpdateConstraintMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityProcessorUnitUpdateConstraintMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityUpdateConstraintMeta, errCapabilityUpdateConstraintMeta := json.Marshal(o.CapabilityUpdateConstraintMeta)
	if errCapabilityUpdateConstraintMeta != nil {
		return map[string]interface{}{}, errCapabilityUpdateConstraintMeta
	}
	errCapabilityUpdateConstraintMeta = json.Unmarshal([]byte(serializedCapabilityUpdateConstraintMeta), &toSerialize)
	if errCapabilityUpdateConstraintMeta != nil {
		return map[string]interface{}{}, errCapabilityUpdateConstraintMeta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityProcessorUnitUpdateConstraintMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type CapabilityProcessorUnitUpdateConstraintMetaWithoutEmbeddedStruct struct {
	}

	varCapabilityProcessorUnitUpdateConstraintMetaWithoutEmbeddedStruct := CapabilityProcessorUnitUpdateConstraintMetaWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityProcessorUnitUpdateConstraintMetaWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityProcessorUnitUpdateConstraintMeta := _CapabilityProcessorUnitUpdateConstraintMeta{}
		*o = CapabilityProcessorUnitUpdateConstraintMeta(varCapabilityProcessorUnitUpdateConstraintMeta)
	} else {
		return err
	}

	varCapabilityProcessorUnitUpdateConstraintMeta := _CapabilityProcessorUnitUpdateConstraintMeta{}

	err = json.Unmarshal(data, &varCapabilityProcessorUnitUpdateConstraintMeta)
	if err == nil {
		o.CapabilityUpdateConstraintMeta = varCapabilityProcessorUnitUpdateConstraintMeta.CapabilityUpdateConstraintMeta
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectCapabilityUpdateConstraintMeta := reflect.ValueOf(o.CapabilityUpdateConstraintMeta)
		for i := 0; i < reflectCapabilityUpdateConstraintMeta.Type().NumField(); i++ {
			t := reflectCapabilityUpdateConstraintMeta.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityProcessorUnitUpdateConstraintMeta struct {
	value *CapabilityProcessorUnitUpdateConstraintMeta
	isSet bool
}

func (v NullableCapabilityProcessorUnitUpdateConstraintMeta) Get() *CapabilityProcessorUnitUpdateConstraintMeta {
	return v.value
}

func (v *NullableCapabilityProcessorUnitUpdateConstraintMeta) Set(val *CapabilityProcessorUnitUpdateConstraintMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityProcessorUnitUpdateConstraintMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityProcessorUnitUpdateConstraintMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityProcessorUnitUpdateConstraintMeta(val *CapabilityProcessorUnitUpdateConstraintMeta) *NullableCapabilityProcessorUnitUpdateConstraintMeta {
	return &NullableCapabilityProcessorUnitUpdateConstraintMeta{value: val, isSet: true}
}

func (v NullableCapabilityProcessorUnitUpdateConstraintMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityProcessorUnitUpdateConstraintMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}