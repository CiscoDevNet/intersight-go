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

// checks if the IqnpoolReservationReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IqnpoolReservationReference{}

// IqnpoolReservationReference The reference to the reservation object.
type IqnpoolReservationReference struct {
	PoolReservationReference
	AdditionalProperties map[string]interface{}
}

type _IqnpoolReservationReference IqnpoolReservationReference

// NewIqnpoolReservationReference instantiates a new IqnpoolReservationReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIqnpoolReservationReference(classId string, objectType string) *IqnpoolReservationReference {
	this := IqnpoolReservationReference{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIqnpoolReservationReferenceWithDefaults instantiates a new IqnpoolReservationReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIqnpoolReservationReferenceWithDefaults() *IqnpoolReservationReference {
	this := IqnpoolReservationReference{}
	return &this
}

func (o IqnpoolReservationReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IqnpoolReservationReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolReservationReference, errPoolReservationReference := json.Marshal(o.PoolReservationReference)
	if errPoolReservationReference != nil {
		return map[string]interface{}{}, errPoolReservationReference
	}
	errPoolReservationReference = json.Unmarshal([]byte(serializedPoolReservationReference), &toSerialize)
	if errPoolReservationReference != nil {
		return map[string]interface{}{}, errPoolReservationReference
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IqnpoolReservationReference) UnmarshalJSON(data []byte) (err error) {
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
	type IqnpoolReservationReferenceWithoutEmbeddedStruct struct {
	}

	varIqnpoolReservationReferenceWithoutEmbeddedStruct := IqnpoolReservationReferenceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIqnpoolReservationReferenceWithoutEmbeddedStruct)
	if err == nil {
		varIqnpoolReservationReference := _IqnpoolReservationReference{}
		*o = IqnpoolReservationReference(varIqnpoolReservationReference)
	} else {
		return err
	}

	varIqnpoolReservationReference := _IqnpoolReservationReference{}

	err = json.Unmarshal(data, &varIqnpoolReservationReference)
	if err == nil {
		o.PoolReservationReference = varIqnpoolReservationReference.PoolReservationReference
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectPoolReservationReference := reflect.ValueOf(o.PoolReservationReference)
		for i := 0; i < reflectPoolReservationReference.Type().NumField(); i++ {
			t := reflectPoolReservationReference.Type().Field(i)

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

type NullableIqnpoolReservationReference struct {
	value *IqnpoolReservationReference
	isSet bool
}

func (v NullableIqnpoolReservationReference) Get() *IqnpoolReservationReference {
	return v.value
}

func (v *NullableIqnpoolReservationReference) Set(val *IqnpoolReservationReference) {
	v.value = val
	v.isSet = true
}

func (v NullableIqnpoolReservationReference) IsSet() bool {
	return v.isSet
}

func (v *NullableIqnpoolReservationReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIqnpoolReservationReference(val *IqnpoolReservationReference) *NullableIqnpoolReservationReference {
	return &NullableIqnpoolReservationReference{value: val, isSet: true}
}

func (v NullableIqnpoolReservationReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIqnpoolReservationReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
