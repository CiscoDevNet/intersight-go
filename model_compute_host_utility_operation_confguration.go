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
	"reflect"
	"strings"
)

// checks if the ComputeHostUtilityOperationConfguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputeHostUtilityOperationConfguration{}

// ComputeHostUtilityOperationConfguration Host Utility operation related confgiruations to use.
type ComputeHostUtilityOperationConfguration struct {
	MoBaseComplexType
	AdditionalProperties map[string]interface{}
}

type _ComputeHostUtilityOperationConfguration ComputeHostUtilityOperationConfguration

// NewComputeHostUtilityOperationConfguration instantiates a new ComputeHostUtilityOperationConfguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeHostUtilityOperationConfguration(classId string, objectType string) *ComputeHostUtilityOperationConfguration {
	this := ComputeHostUtilityOperationConfguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputeHostUtilityOperationConfgurationWithDefaults instantiates a new ComputeHostUtilityOperationConfguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeHostUtilityOperationConfgurationWithDefaults() *ComputeHostUtilityOperationConfguration {
	this := ComputeHostUtilityOperationConfguration{}
	return &this
}

func (o ComputeHostUtilityOperationConfguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputeHostUtilityOperationConfguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ComputeHostUtilityOperationConfguration) UnmarshalJSON(data []byte) (err error) {
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
	type ComputeHostUtilityOperationConfgurationWithoutEmbeddedStruct struct {
	}

	varComputeHostUtilityOperationConfgurationWithoutEmbeddedStruct := ComputeHostUtilityOperationConfgurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varComputeHostUtilityOperationConfgurationWithoutEmbeddedStruct)
	if err == nil {
		varComputeHostUtilityOperationConfguration := _ComputeHostUtilityOperationConfguration{}
		*o = ComputeHostUtilityOperationConfguration(varComputeHostUtilityOperationConfguration)
	} else {
		return err
	}

	varComputeHostUtilityOperationConfguration := _ComputeHostUtilityOperationConfguration{}

	err = json.Unmarshal(data, &varComputeHostUtilityOperationConfguration)
	if err == nil {
		o.MoBaseComplexType = varComputeHostUtilityOperationConfguration.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableComputeHostUtilityOperationConfguration struct {
	value *ComputeHostUtilityOperationConfguration
	isSet bool
}

func (v NullableComputeHostUtilityOperationConfguration) Get() *ComputeHostUtilityOperationConfguration {
	return v.value
}

func (v *NullableComputeHostUtilityOperationConfguration) Set(val *ComputeHostUtilityOperationConfguration) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeHostUtilityOperationConfguration) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeHostUtilityOperationConfguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeHostUtilityOperationConfguration(val *ComputeHostUtilityOperationConfguration) *NullableComputeHostUtilityOperationConfguration {
	return &NullableComputeHostUtilityOperationConfguration{value: val, isSet: true}
}

func (v NullableComputeHostUtilityOperationConfguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeHostUtilityOperationConfguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
