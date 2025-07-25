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

// checks if the KubernetesKeyValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesKeyValue{}

// KubernetesKeyValue Key and value pair used to represent generic data.
type KubernetesKeyValue struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Key or property name in a key/value pair.
	Key *string `json:"Key,omitempty"`
	// Property value in a key/value pair.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesKeyValue KubernetesKeyValue

// NewKubernetesKeyValue instantiates a new KubernetesKeyValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesKeyValue(classId string, objectType string) *KubernetesKeyValue {
	this := KubernetesKeyValue{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesKeyValueWithDefaults instantiates a new KubernetesKeyValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesKeyValueWithDefaults() *KubernetesKeyValue {
	this := KubernetesKeyValue{}
	var classId string = "kubernetes.KeyValue"
	this.ClassId = classId
	var objectType string = "kubernetes.KeyValue"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesKeyValue) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesKeyValue) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesKeyValue) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.KeyValue" of the ClassId field.
func (o *KubernetesKeyValue) GetDefaultClassId() interface{} {
	return "kubernetes.KeyValue"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesKeyValue) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesKeyValue) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesKeyValue) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.KeyValue" of the ObjectType field.
func (o *KubernetesKeyValue) GetDefaultObjectType() interface{} {
	return "kubernetes.KeyValue"
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *KubernetesKeyValue) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesKeyValue) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *KubernetesKeyValue) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *KubernetesKeyValue) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *KubernetesKeyValue) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesKeyValue) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *KubernetesKeyValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *KubernetesKeyValue) SetValue(v string) {
	o.Value = &v
}

func (o KubernetesKeyValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesKeyValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesKeyValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
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
	type KubernetesKeyValueWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Key or property name in a key/value pair.
		Key *string `json:"Key,omitempty"`
		// Property value in a key/value pair.
		Value *string `json:"Value,omitempty"`
	}

	varKubernetesKeyValueWithoutEmbeddedStruct := KubernetesKeyValueWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesKeyValueWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesKeyValue := _KubernetesKeyValue{}
		varKubernetesKeyValue.ClassId = varKubernetesKeyValueWithoutEmbeddedStruct.ClassId
		varKubernetesKeyValue.ObjectType = varKubernetesKeyValueWithoutEmbeddedStruct.ObjectType
		varKubernetesKeyValue.Key = varKubernetesKeyValueWithoutEmbeddedStruct.Key
		varKubernetesKeyValue.Value = varKubernetesKeyValueWithoutEmbeddedStruct.Value
		*o = KubernetesKeyValue(varKubernetesKeyValue)
	} else {
		return err
	}

	varKubernetesKeyValue := _KubernetesKeyValue{}

	err = json.Unmarshal(data, &varKubernetesKeyValue)
	if err == nil {
		o.MoBaseComplexType = varKubernetesKeyValue.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Value")

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

type NullableKubernetesKeyValue struct {
	value *KubernetesKeyValue
	isSet bool
}

func (v NullableKubernetesKeyValue) Get() *KubernetesKeyValue {
	return v.value
}

func (v *NullableKubernetesKeyValue) Set(val *KubernetesKeyValue) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesKeyValue) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesKeyValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesKeyValue(val *KubernetesKeyValue) *NullableKubernetesKeyValue {
	return &NullableKubernetesKeyValue{value: val, isSet: true}
}

func (v NullableKubernetesKeyValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesKeyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
