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

// checks if the SchedulerBaseMonthlyCadenceParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulerBaseMonthlyCadenceParams{}

// SchedulerBaseMonthlyCadenceParams The monthly cadence applied to monthly schedules.
type SchedulerBaseMonthlyCadenceParams struct {
	SchedulerBaseCadenceParams
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Significant business days, such as days when reports are generated for analysis. * `None` - Placeholder. One of the following two fields must be selected. * `FirstWeekDay` - First week day of the month. * `MonthLastDay` - The last day of the month.
	CustomDayOfMonth     *string `json:"CustomDayOfMonth,omitempty"`
	DayOfMonth           []int64 `json:"DayOfMonth,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchedulerBaseMonthlyCadenceParams SchedulerBaseMonthlyCadenceParams

// NewSchedulerBaseMonthlyCadenceParams instantiates a new SchedulerBaseMonthlyCadenceParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulerBaseMonthlyCadenceParams(classId string, objectType string) *SchedulerBaseMonthlyCadenceParams {
	this := SchedulerBaseMonthlyCadenceParams{}
	this.ClassId = classId
	this.ObjectType = objectType
	var customDayOfMonth string = "None"
	this.CustomDayOfMonth = &customDayOfMonth
	return &this
}

// NewSchedulerBaseMonthlyCadenceParamsWithDefaults instantiates a new SchedulerBaseMonthlyCadenceParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulerBaseMonthlyCadenceParamsWithDefaults() *SchedulerBaseMonthlyCadenceParams {
	this := SchedulerBaseMonthlyCadenceParams{}
	var classId string = "scheduler.MonthlyCadenceParams"
	this.ClassId = classId
	var objectType string = "scheduler.MonthlyCadenceParams"
	this.ObjectType = objectType
	var customDayOfMonth string = "None"
	this.CustomDayOfMonth = &customDayOfMonth
	return &this
}

// GetClassId returns the ClassId field value
func (o *SchedulerBaseMonthlyCadenceParams) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SchedulerBaseMonthlyCadenceParams) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SchedulerBaseMonthlyCadenceParams) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "scheduler.MonthlyCadenceParams" of the ClassId field.
func (o *SchedulerBaseMonthlyCadenceParams) GetDefaultClassId() interface{} {
	return "scheduler.MonthlyCadenceParams"
}

// GetObjectType returns the ObjectType field value
func (o *SchedulerBaseMonthlyCadenceParams) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SchedulerBaseMonthlyCadenceParams) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SchedulerBaseMonthlyCadenceParams) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "scheduler.MonthlyCadenceParams" of the ObjectType field.
func (o *SchedulerBaseMonthlyCadenceParams) GetDefaultObjectType() interface{} {
	return "scheduler.MonthlyCadenceParams"
}

// GetCustomDayOfMonth returns the CustomDayOfMonth field value if set, zero value otherwise.
func (o *SchedulerBaseMonthlyCadenceParams) GetCustomDayOfMonth() string {
	if o == nil || IsNil(o.CustomDayOfMonth) {
		var ret string
		return ret
	}
	return *o.CustomDayOfMonth
}

// GetCustomDayOfMonthOk returns a tuple with the CustomDayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerBaseMonthlyCadenceParams) GetCustomDayOfMonthOk() (*string, bool) {
	if o == nil || IsNil(o.CustomDayOfMonth) {
		return nil, false
	}
	return o.CustomDayOfMonth, true
}

// HasCustomDayOfMonth returns a boolean if a field has been set.
func (o *SchedulerBaseMonthlyCadenceParams) HasCustomDayOfMonth() bool {
	if o != nil && !IsNil(o.CustomDayOfMonth) {
		return true
	}

	return false
}

// SetCustomDayOfMonth gets a reference to the given string and assigns it to the CustomDayOfMonth field.
func (o *SchedulerBaseMonthlyCadenceParams) SetCustomDayOfMonth(v string) {
	o.CustomDayOfMonth = &v
}

// GetDayOfMonth returns the DayOfMonth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulerBaseMonthlyCadenceParams) GetDayOfMonth() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulerBaseMonthlyCadenceParams) GetDayOfMonthOk() ([]int64, bool) {
	if o == nil || IsNil(o.DayOfMonth) {
		return nil, false
	}
	return o.DayOfMonth, true
}

// HasDayOfMonth returns a boolean if a field has been set.
func (o *SchedulerBaseMonthlyCadenceParams) HasDayOfMonth() bool {
	if o != nil && !IsNil(o.DayOfMonth) {
		return true
	}

	return false
}

// SetDayOfMonth gets a reference to the given []int64 and assigns it to the DayOfMonth field.
func (o *SchedulerBaseMonthlyCadenceParams) SetDayOfMonth(v []int64) {
	o.DayOfMonth = v
}

func (o SchedulerBaseMonthlyCadenceParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulerBaseMonthlyCadenceParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSchedulerBaseCadenceParams, errSchedulerBaseCadenceParams := json.Marshal(o.SchedulerBaseCadenceParams)
	if errSchedulerBaseCadenceParams != nil {
		return map[string]interface{}{}, errSchedulerBaseCadenceParams
	}
	errSchedulerBaseCadenceParams = json.Unmarshal([]byte(serializedSchedulerBaseCadenceParams), &toSerialize)
	if errSchedulerBaseCadenceParams != nil {
		return map[string]interface{}{}, errSchedulerBaseCadenceParams
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.CustomDayOfMonth) {
		toSerialize["CustomDayOfMonth"] = o.CustomDayOfMonth
	}
	if o.DayOfMonth != nil {
		toSerialize["DayOfMonth"] = o.DayOfMonth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchedulerBaseMonthlyCadenceParams) UnmarshalJSON(data []byte) (err error) {
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
	type SchedulerBaseMonthlyCadenceParamsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Significant business days, such as days when reports are generated for analysis. * `None` - Placeholder. One of the following two fields must be selected. * `FirstWeekDay` - First week day of the month. * `MonthLastDay` - The last day of the month.
		CustomDayOfMonth *string `json:"CustomDayOfMonth,omitempty"`
		DayOfMonth       []int64 `json:"DayOfMonth,omitempty"`
	}

	varSchedulerBaseMonthlyCadenceParamsWithoutEmbeddedStruct := SchedulerBaseMonthlyCadenceParamsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSchedulerBaseMonthlyCadenceParamsWithoutEmbeddedStruct)
	if err == nil {
		varSchedulerBaseMonthlyCadenceParams := _SchedulerBaseMonthlyCadenceParams{}
		varSchedulerBaseMonthlyCadenceParams.ClassId = varSchedulerBaseMonthlyCadenceParamsWithoutEmbeddedStruct.ClassId
		varSchedulerBaseMonthlyCadenceParams.ObjectType = varSchedulerBaseMonthlyCadenceParamsWithoutEmbeddedStruct.ObjectType
		varSchedulerBaseMonthlyCadenceParams.CustomDayOfMonth = varSchedulerBaseMonthlyCadenceParamsWithoutEmbeddedStruct.CustomDayOfMonth
		varSchedulerBaseMonthlyCadenceParams.DayOfMonth = varSchedulerBaseMonthlyCadenceParamsWithoutEmbeddedStruct.DayOfMonth
		*o = SchedulerBaseMonthlyCadenceParams(varSchedulerBaseMonthlyCadenceParams)
	} else {
		return err
	}

	varSchedulerBaseMonthlyCadenceParams := _SchedulerBaseMonthlyCadenceParams{}

	err = json.Unmarshal(data, &varSchedulerBaseMonthlyCadenceParams)
	if err == nil {
		o.SchedulerBaseCadenceParams = varSchedulerBaseMonthlyCadenceParams.SchedulerBaseCadenceParams
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CustomDayOfMonth")
		delete(additionalProperties, "DayOfMonth")

		// remove fields from embedded structs
		reflectSchedulerBaseCadenceParams := reflect.ValueOf(o.SchedulerBaseCadenceParams)
		for i := 0; i < reflectSchedulerBaseCadenceParams.Type().NumField(); i++ {
			t := reflectSchedulerBaseCadenceParams.Type().Field(i)

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

type NullableSchedulerBaseMonthlyCadenceParams struct {
	value *SchedulerBaseMonthlyCadenceParams
	isSet bool
}

func (v NullableSchedulerBaseMonthlyCadenceParams) Get() *SchedulerBaseMonthlyCadenceParams {
	return v.value
}

func (v *NullableSchedulerBaseMonthlyCadenceParams) Set(val *SchedulerBaseMonthlyCadenceParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulerBaseMonthlyCadenceParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulerBaseMonthlyCadenceParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulerBaseMonthlyCadenceParams(val *SchedulerBaseMonthlyCadenceParams) *NullableSchedulerBaseMonthlyCadenceParams {
	return &NullableSchedulerBaseMonthlyCadenceParams{value: val, isSet: true}
}

func (v NullableSchedulerBaseMonthlyCadenceParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulerBaseMonthlyCadenceParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
