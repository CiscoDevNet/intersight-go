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

// checks if the TelemetryDruidExtractionFunctionLowerCase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidExtractionFunctionLowerCase{}

// TelemetryDruidExtractionFunctionLowerCase Returns the dimension values as all lower case. Optionally user can specify the language to use in order to perform lower transformation.
type TelemetryDruidExtractionFunctionLowerCase struct {
	Type                 string  `json:"type"`
	Locale               *string `json:"locale,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidExtractionFunctionLowerCase TelemetryDruidExtractionFunctionLowerCase

// NewTelemetryDruidExtractionFunctionLowerCase instantiates a new TelemetryDruidExtractionFunctionLowerCase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidExtractionFunctionLowerCase(type_ string) *TelemetryDruidExtractionFunctionLowerCase {
	this := TelemetryDruidExtractionFunctionLowerCase{}
	this.Type = type_
	return &this
}

// NewTelemetryDruidExtractionFunctionLowerCaseWithDefaults instantiates a new TelemetryDruidExtractionFunctionLowerCase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidExtractionFunctionLowerCaseWithDefaults() *TelemetryDruidExtractionFunctionLowerCase {
	this := TelemetryDruidExtractionFunctionLowerCase{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidExtractionFunctionLowerCase) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionLowerCase) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidExtractionFunctionLowerCase) SetType(v string) {
	o.Type = v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionLowerCase) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionLowerCase) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionLowerCase) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *TelemetryDruidExtractionFunctionLowerCase) SetLocale(v string) {
	o.Locale = &v
}

func (o TelemetryDruidExtractionFunctionLowerCase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidExtractionFunctionLowerCase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidExtractionFunctionLowerCase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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
	varTelemetryDruidExtractionFunctionLowerCase := _TelemetryDruidExtractionFunctionLowerCase{}

	err = json.Unmarshal(data, &varTelemetryDruidExtractionFunctionLowerCase)

	if err != nil {
		return err
	}

	*o = TelemetryDruidExtractionFunctionLowerCase(varTelemetryDruidExtractionFunctionLowerCase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "locale")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidExtractionFunctionLowerCase struct {
	value *TelemetryDruidExtractionFunctionLowerCase
	isSet bool
}

func (v NullableTelemetryDruidExtractionFunctionLowerCase) Get() *TelemetryDruidExtractionFunctionLowerCase {
	return v.value
}

func (v *NullableTelemetryDruidExtractionFunctionLowerCase) Set(val *TelemetryDruidExtractionFunctionLowerCase) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidExtractionFunctionLowerCase) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidExtractionFunctionLowerCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidExtractionFunctionLowerCase(val *TelemetryDruidExtractionFunctionLowerCase) *NullableTelemetryDruidExtractionFunctionLowerCase {
	return &NullableTelemetryDruidExtractionFunctionLowerCase{value: val, isSet: true}
}

func (v NullableTelemetryDruidExtractionFunctionLowerCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidExtractionFunctionLowerCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
