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

// checks if the VnicArfsSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicArfsSettings{}

// VnicArfsSettings Settings for Accelerated Receive Flow Steering to reduce the network latency and increase CPU cache efficiency.
type VnicArfsSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of Accelerated Receive Flow Steering on the virtual ethernet interface.
	Enabled              *bool `json:"Enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicArfsSettings VnicArfsSettings

// NewVnicArfsSettings instantiates a new VnicArfsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicArfsSettings(classId string, objectType string) *VnicArfsSettings {
	this := VnicArfsSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewVnicArfsSettingsWithDefaults instantiates a new VnicArfsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicArfsSettingsWithDefaults() *VnicArfsSettings {
	this := VnicArfsSettings{}
	var classId string = "vnic.ArfsSettings"
	this.ClassId = classId
	var objectType string = "vnic.ArfsSettings"
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicArfsSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicArfsSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicArfsSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.ArfsSettings" of the ClassId field.
func (o *VnicArfsSettings) GetDefaultClassId() interface{} {
	return "vnic.ArfsSettings"
}

// GetObjectType returns the ObjectType field value
func (o *VnicArfsSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicArfsSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicArfsSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.ArfsSettings" of the ObjectType field.
func (o *VnicArfsSettings) GetDefaultObjectType() interface{} {
	return "vnic.ArfsSettings"
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VnicArfsSettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicArfsSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VnicArfsSettings) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VnicArfsSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o VnicArfsSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicArfsSettings) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicArfsSettings) UnmarshalJSON(data []byte) (err error) {
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
	type VnicArfsSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Status of Accelerated Receive Flow Steering on the virtual ethernet interface.
		Enabled *bool `json:"Enabled,omitempty"`
	}

	varVnicArfsSettingsWithoutEmbeddedStruct := VnicArfsSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicArfsSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicArfsSettings := _VnicArfsSettings{}
		varVnicArfsSettings.ClassId = varVnicArfsSettingsWithoutEmbeddedStruct.ClassId
		varVnicArfsSettings.ObjectType = varVnicArfsSettingsWithoutEmbeddedStruct.ObjectType
		varVnicArfsSettings.Enabled = varVnicArfsSettingsWithoutEmbeddedStruct.Enabled
		*o = VnicArfsSettings(varVnicArfsSettings)
	} else {
		return err
	}

	varVnicArfsSettings := _VnicArfsSettings{}

	err = json.Unmarshal(data, &varVnicArfsSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicArfsSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")

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

type NullableVnicArfsSettings struct {
	value *VnicArfsSettings
	isSet bool
}

func (v NullableVnicArfsSettings) Get() *VnicArfsSettings {
	return v.value
}

func (v *NullableVnicArfsSettings) Set(val *VnicArfsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicArfsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicArfsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicArfsSettings(val *VnicArfsSettings) *NullableVnicArfsSettings {
	return &NullableVnicArfsSettings{value: val, isSet: true}
}

func (v NullableVnicArfsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicArfsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
