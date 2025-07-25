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

// checks if the ConsoleConsoleConfigList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsoleConsoleConfigList{}

// ConsoleConsoleConfigList This resource list is returned as a response to a HTTP GET request that does not include a specific resource identifier.
type ConsoleConsoleConfigList struct {
	MoBaseResponse
	// The total number of 'console.ConsoleConfig' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty"`
	// The array of 'console.ConsoleConfig' resources matching the request.
	Results              []ConsoleConsoleConfig `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConsoleConsoleConfigList ConsoleConsoleConfigList

// NewConsoleConsoleConfigList instantiates a new ConsoleConsoleConfigList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsoleConsoleConfigList(objectType string) *ConsoleConsoleConfigList {
	this := ConsoleConsoleConfigList{}
	this.ObjectType = objectType
	return &this
}

// NewConsoleConsoleConfigListWithDefaults instantiates a new ConsoleConsoleConfigList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsoleConsoleConfigListWithDefaults() *ConsoleConsoleConfigList {
	this := ConsoleConsoleConfigList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ConsoleConsoleConfigList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleConsoleConfigList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ConsoleConsoleConfigList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ConsoleConsoleConfigList) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsoleConsoleConfigList) GetResults() []ConsoleConsoleConfig {
	if o == nil {
		var ret []ConsoleConsoleConfig
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsoleConsoleConfigList) GetResultsOk() ([]ConsoleConsoleConfig, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ConsoleConsoleConfigList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ConsoleConsoleConfig and assigns it to the Results field.
func (o *ConsoleConsoleConfigList) SetResults(v []ConsoleConsoleConfig) {
	o.Results = v
}

func (o ConsoleConsoleConfigList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsoleConsoleConfigList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseResponse, errMoBaseResponse := json.Marshal(o.MoBaseResponse)
	if errMoBaseResponse != nil {
		return map[string]interface{}{}, errMoBaseResponse
	}
	errMoBaseResponse = json.Unmarshal([]byte(serializedMoBaseResponse), &toSerialize)
	if errMoBaseResponse != nil {
		return map[string]interface{}{}, errMoBaseResponse
	}
	if !IsNil(o.Count) {
		toSerialize["Count"] = o.Count
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConsoleConsoleConfigList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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
	type ConsoleConsoleConfigListWithoutEmbeddedStruct struct {
		// The total number of 'console.ConsoleConfig' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
		Count *int32 `json:"Count,omitempty"`
		// The array of 'console.ConsoleConfig' resources matching the request.
		Results []ConsoleConsoleConfig `json:"Results,omitempty"`
	}

	varConsoleConsoleConfigListWithoutEmbeddedStruct := ConsoleConsoleConfigListWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varConsoleConsoleConfigListWithoutEmbeddedStruct)
	if err == nil {
		varConsoleConsoleConfigList := _ConsoleConsoleConfigList{}
		varConsoleConsoleConfigList.Count = varConsoleConsoleConfigListWithoutEmbeddedStruct.Count
		varConsoleConsoleConfigList.Results = varConsoleConsoleConfigListWithoutEmbeddedStruct.Results
		*o = ConsoleConsoleConfigList(varConsoleConsoleConfigList)
	} else {
		return err
	}

	varConsoleConsoleConfigList := _ConsoleConsoleConfigList{}

	err = json.Unmarshal(data, &varConsoleConsoleConfigList)
	if err == nil {
		o.MoBaseResponse = varConsoleConsoleConfigList.MoBaseResponse
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "Count")
		delete(additionalProperties, "Results")

		// remove fields from embedded structs
		reflectMoBaseResponse := reflect.ValueOf(o.MoBaseResponse)
		for i := 0; i < reflectMoBaseResponse.Type().NumField(); i++ {
			t := reflectMoBaseResponse.Type().Field(i)

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

type NullableConsoleConsoleConfigList struct {
	value *ConsoleConsoleConfigList
	isSet bool
}

func (v NullableConsoleConsoleConfigList) Get() *ConsoleConsoleConfigList {
	return v.value
}

func (v *NullableConsoleConsoleConfigList) Set(val *ConsoleConsoleConfigList) {
	v.value = val
	v.isSet = true
}

func (v NullableConsoleConsoleConfigList) IsSet() bool {
	return v.isSet
}

func (v *NullableConsoleConsoleConfigList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsoleConsoleConfigList(val *ConsoleConsoleConfigList) *NullableConsoleConsoleConfigList {
	return &NullableConsoleConsoleConfigList{value: val, isSet: true}
}

func (v NullableConsoleConsoleConfigList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsoleConsoleConfigList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
