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

// checks if the AccessConfigurationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessConfigurationType{}

// AccessConfigurationType The type of configuration, In-Band and/or Out-Of-Band to be configured on the CIMC.
type AccessConfigurationType struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This flag enables the use of In-Band configuration for end-point access.
	ConfigureInband *bool `json:"ConfigureInband,omitempty"`
	// This flag enables the use of Out-Of-Band configuration for end-point access.
	ConfigureOutOfBand   *bool `json:"ConfigureOutOfBand,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessConfigurationType AccessConfigurationType

// NewAccessConfigurationType instantiates a new AccessConfigurationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessConfigurationType(classId string, objectType string) *AccessConfigurationType {
	this := AccessConfigurationType{}
	this.ClassId = classId
	this.ObjectType = objectType
	var configureInband bool = true
	this.ConfigureInband = &configureInband
	var configureOutOfBand bool = false
	this.ConfigureOutOfBand = &configureOutOfBand
	return &this
}

// NewAccessConfigurationTypeWithDefaults instantiates a new AccessConfigurationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessConfigurationTypeWithDefaults() *AccessConfigurationType {
	this := AccessConfigurationType{}
	var classId string = "access.ConfigurationType"
	this.ClassId = classId
	var objectType string = "access.ConfigurationType"
	this.ObjectType = objectType
	var configureInband bool = true
	this.ConfigureInband = &configureInband
	var configureOutOfBand bool = false
	this.ConfigureOutOfBand = &configureOutOfBand
	return &this
}

// GetClassId returns the ClassId field value
func (o *AccessConfigurationType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AccessConfigurationType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AccessConfigurationType) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "access.ConfigurationType" of the ClassId field.
func (o *AccessConfigurationType) GetDefaultClassId() interface{} {
	return "access.ConfigurationType"
}

// GetObjectType returns the ObjectType field value
func (o *AccessConfigurationType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AccessConfigurationType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AccessConfigurationType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "access.ConfigurationType" of the ObjectType field.
func (o *AccessConfigurationType) GetDefaultObjectType() interface{} {
	return "access.ConfigurationType"
}

// GetConfigureInband returns the ConfigureInband field value if set, zero value otherwise.
func (o *AccessConfigurationType) GetConfigureInband() bool {
	if o == nil || IsNil(o.ConfigureInband) {
		var ret bool
		return ret
	}
	return *o.ConfigureInband
}

// GetConfigureInbandOk returns a tuple with the ConfigureInband field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessConfigurationType) GetConfigureInbandOk() (*bool, bool) {
	if o == nil || IsNil(o.ConfigureInband) {
		return nil, false
	}
	return o.ConfigureInband, true
}

// HasConfigureInband returns a boolean if a field has been set.
func (o *AccessConfigurationType) HasConfigureInband() bool {
	if o != nil && !IsNil(o.ConfigureInband) {
		return true
	}

	return false
}

// SetConfigureInband gets a reference to the given bool and assigns it to the ConfigureInband field.
func (o *AccessConfigurationType) SetConfigureInband(v bool) {
	o.ConfigureInband = &v
}

// GetConfigureOutOfBand returns the ConfigureOutOfBand field value if set, zero value otherwise.
func (o *AccessConfigurationType) GetConfigureOutOfBand() bool {
	if o == nil || IsNil(o.ConfigureOutOfBand) {
		var ret bool
		return ret
	}
	return *o.ConfigureOutOfBand
}

// GetConfigureOutOfBandOk returns a tuple with the ConfigureOutOfBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessConfigurationType) GetConfigureOutOfBandOk() (*bool, bool) {
	if o == nil || IsNil(o.ConfigureOutOfBand) {
		return nil, false
	}
	return o.ConfigureOutOfBand, true
}

// HasConfigureOutOfBand returns a boolean if a field has been set.
func (o *AccessConfigurationType) HasConfigureOutOfBand() bool {
	if o != nil && !IsNil(o.ConfigureOutOfBand) {
		return true
	}

	return false
}

// SetConfigureOutOfBand gets a reference to the given bool and assigns it to the ConfigureOutOfBand field.
func (o *AccessConfigurationType) SetConfigureOutOfBand(v bool) {
	o.ConfigureOutOfBand = &v
}

func (o AccessConfigurationType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessConfigurationType) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ConfigureInband) {
		toSerialize["ConfigureInband"] = o.ConfigureInband
	}
	if !IsNil(o.ConfigureOutOfBand) {
		toSerialize["ConfigureOutOfBand"] = o.ConfigureOutOfBand
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessConfigurationType) UnmarshalJSON(data []byte) (err error) {
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
	type AccessConfigurationTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This flag enables the use of In-Band configuration for end-point access.
		ConfigureInband *bool `json:"ConfigureInband,omitempty"`
		// This flag enables the use of Out-Of-Band configuration for end-point access.
		ConfigureOutOfBand *bool `json:"ConfigureOutOfBand,omitempty"`
	}

	varAccessConfigurationTypeWithoutEmbeddedStruct := AccessConfigurationTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAccessConfigurationTypeWithoutEmbeddedStruct)
	if err == nil {
		varAccessConfigurationType := _AccessConfigurationType{}
		varAccessConfigurationType.ClassId = varAccessConfigurationTypeWithoutEmbeddedStruct.ClassId
		varAccessConfigurationType.ObjectType = varAccessConfigurationTypeWithoutEmbeddedStruct.ObjectType
		varAccessConfigurationType.ConfigureInband = varAccessConfigurationTypeWithoutEmbeddedStruct.ConfigureInband
		varAccessConfigurationType.ConfigureOutOfBand = varAccessConfigurationTypeWithoutEmbeddedStruct.ConfigureOutOfBand
		*o = AccessConfigurationType(varAccessConfigurationType)
	} else {
		return err
	}

	varAccessConfigurationType := _AccessConfigurationType{}

	err = json.Unmarshal(data, &varAccessConfigurationType)
	if err == nil {
		o.MoBaseComplexType = varAccessConfigurationType.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigureInband")
		delete(additionalProperties, "ConfigureOutOfBand")

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

type NullableAccessConfigurationType struct {
	value *AccessConfigurationType
	isSet bool
}

func (v NullableAccessConfigurationType) Get() *AccessConfigurationType {
	return v.value
}

func (v *NullableAccessConfigurationType) Set(val *AccessConfigurationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessConfigurationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessConfigurationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessConfigurationType(val *AccessConfigurationType) *NullableAccessConfigurationType {
	return &NullableAccessConfigurationType{value: val, isSet: true}
}

func (v NullableAccessConfigurationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessConfigurationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
