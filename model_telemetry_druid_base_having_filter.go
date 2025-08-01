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

// checks if the TelemetryDruidBaseHavingFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidBaseHavingFilter{}

// TelemetryDruidBaseHavingFilter A having clause is a JSON object identifying which rows from a groupBy query should be returned, by specifying conditions on aggregated values. It is essentially the equivalent of the HAVING clause in SQL.
type TelemetryDruidBaseHavingFilter struct {
	// The having filter type.
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidBaseHavingFilter TelemetryDruidBaseHavingFilter

// NewTelemetryDruidBaseHavingFilter instantiates a new TelemetryDruidBaseHavingFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidBaseHavingFilter(type_ string) *TelemetryDruidBaseHavingFilter {
	this := TelemetryDruidBaseHavingFilter{}
	this.Type = type_
	return &this
}

// NewTelemetryDruidBaseHavingFilterWithDefaults instantiates a new TelemetryDruidBaseHavingFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidBaseHavingFilterWithDefaults() *TelemetryDruidBaseHavingFilter {
	this := TelemetryDruidBaseHavingFilter{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidBaseHavingFilter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidBaseHavingFilter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidBaseHavingFilter) SetType(v string) {
	o.Type = v
}

func (o TelemetryDruidBaseHavingFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidBaseHavingFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidBaseHavingFilter) UnmarshalJSON(data []byte) (err error) {
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
	varTelemetryDruidBaseHavingFilter := _TelemetryDruidBaseHavingFilter{}

	err = json.Unmarshal(data, &varTelemetryDruidBaseHavingFilter)

	if err != nil {
		return err
	}

	*o = TelemetryDruidBaseHavingFilter(varTelemetryDruidBaseHavingFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidBaseHavingFilter struct {
	value *TelemetryDruidBaseHavingFilter
	isSet bool
}

func (v NullableTelemetryDruidBaseHavingFilter) Get() *TelemetryDruidBaseHavingFilter {
	return v.value
}

func (v *NullableTelemetryDruidBaseHavingFilter) Set(val *TelemetryDruidBaseHavingFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidBaseHavingFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidBaseHavingFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidBaseHavingFilter(val *TelemetryDruidBaseHavingFilter) *NullableTelemetryDruidBaseHavingFilter {
	return &NullableTelemetryDruidBaseHavingFilter{value: val, isSet: true}
}

func (v NullableTelemetryDruidBaseHavingFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidBaseHavingFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
