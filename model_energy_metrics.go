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
	"reflect"
	"strings"
)

// checks if the EnergyMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnergyMetrics{}

// EnergyMetrics Returns carbon intensity, power breakdown, and electricity rate power metrics for a set of locations and a time span.
type EnergyMetrics struct {
	MoBaseMo
	AdditionalProperties map[string]interface{}
}

type _EnergyMetrics EnergyMetrics

// NewEnergyMetrics instantiates a new EnergyMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnergyMetrics(classId string, objectType string) *EnergyMetrics {
	this := EnergyMetrics{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEnergyMetricsWithDefaults instantiates a new EnergyMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnergyMetricsWithDefaults() *EnergyMetrics {
	this := EnergyMetrics{}
	return &this
}

func (o EnergyMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnergyMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnergyMetrics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type EnergyMetricsWithoutEmbeddedStruct struct {
	}

	varEnergyMetricsWithoutEmbeddedStruct := EnergyMetricsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEnergyMetricsWithoutEmbeddedStruct)
	if err == nil {
		varEnergyMetrics := _EnergyMetrics{}
		*o = EnergyMetrics(varEnergyMetrics)
	} else {
		return err
	}

	varEnergyMetrics := _EnergyMetrics{}

	err = json.Unmarshal(data, &varEnergyMetrics)
	if err == nil {
		o.MoBaseMo = varEnergyMetrics.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableEnergyMetrics struct {
	value *EnergyMetrics
	isSet bool
}

func (v NullableEnergyMetrics) Get() *EnergyMetrics {
	return v.value
}

func (v *NullableEnergyMetrics) Set(val *EnergyMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableEnergyMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableEnergyMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnergyMetrics(val *EnergyMetrics) *NullableEnergyMetrics {
	return &NullableEnergyMetrics{value: val, isSet: true}
}

func (v NullableEnergyMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnergyMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}