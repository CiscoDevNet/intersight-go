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

// checks if the TelemetryDruidStringAnyAggregator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidStringAnyAggregator{}

// TelemetryDruidStringAnyAggregator Returns any value including null. This aggregator can simplify and optimize the performance by returning the first encountered value (including null).
type TelemetryDruidStringAnyAggregator struct {
	// The aggregator type.
	Type string `json:"type"`
	// Output name for the 'any' value.
	Name string `json:"name"`
	// Name of the metric column.
	FieldName            string `json:"fieldName"`
	MaxStringBytes       *int32 `json:"maxStringBytes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidStringAnyAggregator TelemetryDruidStringAnyAggregator

// NewTelemetryDruidStringAnyAggregator instantiates a new TelemetryDruidStringAnyAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidStringAnyAggregator(type_ string, name string, fieldName string) *TelemetryDruidStringAnyAggregator {
	this := TelemetryDruidStringAnyAggregator{}
	this.Type = type_
	this.Name = name
	this.FieldName = fieldName
	var maxStringBytes int32 = 1024
	this.MaxStringBytes = &maxStringBytes
	return &this
}

// NewTelemetryDruidStringAnyAggregatorWithDefaults instantiates a new TelemetryDruidStringAnyAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidStringAnyAggregatorWithDefaults() *TelemetryDruidStringAnyAggregator {
	this := TelemetryDruidStringAnyAggregator{}
	var maxStringBytes int32 = 1024
	this.MaxStringBytes = &maxStringBytes
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidStringAnyAggregator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidStringAnyAggregator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidStringAnyAggregator) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *TelemetryDruidStringAnyAggregator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidStringAnyAggregator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TelemetryDruidStringAnyAggregator) SetName(v string) {
	o.Name = v
}

// GetFieldName returns the FieldName field value
func (o *TelemetryDruidStringAnyAggregator) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidStringAnyAggregator) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *TelemetryDruidStringAnyAggregator) SetFieldName(v string) {
	o.FieldName = v
}

// GetMaxStringBytes returns the MaxStringBytes field value if set, zero value otherwise.
func (o *TelemetryDruidStringAnyAggregator) GetMaxStringBytes() int32 {
	if o == nil || IsNil(o.MaxStringBytes) {
		var ret int32
		return ret
	}
	return *o.MaxStringBytes
}

// GetMaxStringBytesOk returns a tuple with the MaxStringBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidStringAnyAggregator) GetMaxStringBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxStringBytes) {
		return nil, false
	}
	return o.MaxStringBytes, true
}

// HasMaxStringBytes returns a boolean if a field has been set.
func (o *TelemetryDruidStringAnyAggregator) HasMaxStringBytes() bool {
	if o != nil && !IsNil(o.MaxStringBytes) {
		return true
	}

	return false
}

// SetMaxStringBytes gets a reference to the given int32 and assigns it to the MaxStringBytes field.
func (o *TelemetryDruidStringAnyAggregator) SetMaxStringBytes(v int32) {
	o.MaxStringBytes = &v
}

func (o TelemetryDruidStringAnyAggregator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidStringAnyAggregator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["fieldName"] = o.FieldName
	if !IsNil(o.MaxStringBytes) {
		toSerialize["maxStringBytes"] = o.MaxStringBytes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidStringAnyAggregator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"fieldName",
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
	varTelemetryDruidStringAnyAggregator := _TelemetryDruidStringAnyAggregator{}

	err = json.Unmarshal(data, &varTelemetryDruidStringAnyAggregator)

	if err != nil {
		return err
	}

	*o = TelemetryDruidStringAnyAggregator(varTelemetryDruidStringAnyAggregator)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "fieldName")
		delete(additionalProperties, "maxStringBytes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidStringAnyAggregator struct {
	value *TelemetryDruidStringAnyAggregator
	isSet bool
}

func (v NullableTelemetryDruidStringAnyAggregator) Get() *TelemetryDruidStringAnyAggregator {
	return v.value
}

func (v *NullableTelemetryDruidStringAnyAggregator) Set(val *TelemetryDruidStringAnyAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidStringAnyAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidStringAnyAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidStringAnyAggregator(val *TelemetryDruidStringAnyAggregator) *NullableTelemetryDruidStringAnyAggregator {
	return &NullableTelemetryDruidStringAnyAggregator{value: val, isSet: true}
}

func (v NullableTelemetryDruidStringAnyAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidStringAnyAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
