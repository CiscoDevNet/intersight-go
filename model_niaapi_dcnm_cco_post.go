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

// checks if the NiaapiDcnmCcoPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiaapiDcnmCcoPost{}

// NiaapiDcnmCcoPost The post reporting a new release is available for DCNM.
type NiaapiDcnmCcoPost struct {
	NiaapiNewReleasePost
	AdditionalProperties map[string]interface{}
}

type _NiaapiDcnmCcoPost NiaapiDcnmCcoPost

// NewNiaapiDcnmCcoPost instantiates a new NiaapiDcnmCcoPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiDcnmCcoPost(classId string, objectType string) *NiaapiDcnmCcoPost {
	this := NiaapiDcnmCcoPost{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiDcnmCcoPostWithDefaults instantiates a new NiaapiDcnmCcoPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiDcnmCcoPostWithDefaults() *NiaapiDcnmCcoPost {
	this := NiaapiDcnmCcoPost{}
	return &this
}

func (o NiaapiDcnmCcoPost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiaapiDcnmCcoPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNiaapiNewReleasePost, errNiaapiNewReleasePost := json.Marshal(o.NiaapiNewReleasePost)
	if errNiaapiNewReleasePost != nil {
		return map[string]interface{}{}, errNiaapiNewReleasePost
	}
	errNiaapiNewReleasePost = json.Unmarshal([]byte(serializedNiaapiNewReleasePost), &toSerialize)
	if errNiaapiNewReleasePost != nil {
		return map[string]interface{}{}, errNiaapiNewReleasePost
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiaapiDcnmCcoPost) UnmarshalJSON(data []byte) (err error) {
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
	type NiaapiDcnmCcoPostWithoutEmbeddedStruct struct {
	}

	varNiaapiDcnmCcoPostWithoutEmbeddedStruct := NiaapiDcnmCcoPostWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiaapiDcnmCcoPostWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiDcnmCcoPost := _NiaapiDcnmCcoPost{}
		*o = NiaapiDcnmCcoPost(varNiaapiDcnmCcoPost)
	} else {
		return err
	}

	varNiaapiDcnmCcoPost := _NiaapiDcnmCcoPost{}

	err = json.Unmarshal(data, &varNiaapiDcnmCcoPost)
	if err == nil {
		o.NiaapiNewReleasePost = varNiaapiDcnmCcoPost.NiaapiNewReleasePost
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectNiaapiNewReleasePost := reflect.ValueOf(o.NiaapiNewReleasePost)
		for i := 0; i < reflectNiaapiNewReleasePost.Type().NumField(); i++ {
			t := reflectNiaapiNewReleasePost.Type().Field(i)

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

type NullableNiaapiDcnmCcoPost struct {
	value *NiaapiDcnmCcoPost
	isSet bool
}

func (v NullableNiaapiDcnmCcoPost) Get() *NiaapiDcnmCcoPost {
	return v.value
}

func (v *NullableNiaapiDcnmCcoPost) Set(val *NiaapiDcnmCcoPost) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiDcnmCcoPost) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiDcnmCcoPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiDcnmCcoPost(val *NiaapiDcnmCcoPost) *NullableNiaapiDcnmCcoPost {
	return &NullableNiaapiDcnmCcoPost{value: val, isSet: true}
}

func (v NullableNiaapiDcnmCcoPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiDcnmCcoPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
