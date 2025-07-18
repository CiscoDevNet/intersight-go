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

// checks if the TelemetryDruidDefaultLimitSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidDefaultLimitSpec{}

// TelemetryDruidDefaultLimitSpec The default limit spec takes a limit and the list of columns to do an orderBy operation over.
type TelemetryDruidDefaultLimitSpec struct {
	// The limit spec type.
	Type string `json:"type"`
	// How many rows to return. If not specified, all rows will be returned.
	Limit                int32                             `json:"limit"`
	Columns              []TelemetryDruidOrderByColumnSpec `json:"columns"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidDefaultLimitSpec TelemetryDruidDefaultLimitSpec

// NewTelemetryDruidDefaultLimitSpec instantiates a new TelemetryDruidDefaultLimitSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidDefaultLimitSpec(type_ string, limit int32, columns []TelemetryDruidOrderByColumnSpec) *TelemetryDruidDefaultLimitSpec {
	this := TelemetryDruidDefaultLimitSpec{}
	this.Type = type_
	this.Limit = limit
	this.Columns = columns
	return &this
}

// NewTelemetryDruidDefaultLimitSpecWithDefaults instantiates a new TelemetryDruidDefaultLimitSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidDefaultLimitSpecWithDefaults() *TelemetryDruidDefaultLimitSpec {
	this := TelemetryDruidDefaultLimitSpec{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidDefaultLimitSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDefaultLimitSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidDefaultLimitSpec) SetType(v string) {
	o.Type = v
}

// GetLimit returns the Limit field value
func (o *TelemetryDruidDefaultLimitSpec) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDefaultLimitSpec) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *TelemetryDruidDefaultLimitSpec) SetLimit(v int32) {
	o.Limit = v
}

// GetColumns returns the Columns field value
func (o *TelemetryDruidDefaultLimitSpec) GetColumns() []TelemetryDruidOrderByColumnSpec {
	if o == nil {
		var ret []TelemetryDruidOrderByColumnSpec
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDefaultLimitSpec) GetColumnsOk() ([]TelemetryDruidOrderByColumnSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Columns, true
}

// SetColumns sets field value
func (o *TelemetryDruidDefaultLimitSpec) SetColumns(v []TelemetryDruidOrderByColumnSpec) {
	o.Columns = v
}

func (o TelemetryDruidDefaultLimitSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidDefaultLimitSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["limit"] = o.Limit
	toSerialize["columns"] = o.Columns

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidDefaultLimitSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"limit",
		"columns",
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
	varTelemetryDruidDefaultLimitSpec := _TelemetryDruidDefaultLimitSpec{}

	err = json.Unmarshal(data, &varTelemetryDruidDefaultLimitSpec)

	if err != nil {
		return err
	}

	*o = TelemetryDruidDefaultLimitSpec(varTelemetryDruidDefaultLimitSpec)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "columns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidDefaultLimitSpec struct {
	value *TelemetryDruidDefaultLimitSpec
	isSet bool
}

func (v NullableTelemetryDruidDefaultLimitSpec) Get() *TelemetryDruidDefaultLimitSpec {
	return v.value
}

func (v *NullableTelemetryDruidDefaultLimitSpec) Set(val *TelemetryDruidDefaultLimitSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidDefaultLimitSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidDefaultLimitSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidDefaultLimitSpec(val *TelemetryDruidDefaultLimitSpec) *NullableTelemetryDruidDefaultLimitSpec {
	return &NullableTelemetryDruidDefaultLimitSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidDefaultLimitSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidDefaultLimitSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
