/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-17769
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

// checks if the SchedulerMonthlyWeekDayFormat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulerMonthlyWeekDayFormat{}

// SchedulerMonthlyWeekDayFormat An alternative for monthly scheduled task. For e.g. third Tuesday of every month.
type SchedulerMonthlyWeekDayFormat struct {
	MoBaseComplexType
	DayOfWeek []string `json:"DayOfWeek,omitempty"`
	// The week of the month, 1 through 5.
	WeekOfMonth          *int64 `json:"WeekOfMonth,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchedulerMonthlyWeekDayFormat SchedulerMonthlyWeekDayFormat

// NewSchedulerMonthlyWeekDayFormat instantiates a new SchedulerMonthlyWeekDayFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulerMonthlyWeekDayFormat(classId string, objectType string) *SchedulerMonthlyWeekDayFormat {
	this := SchedulerMonthlyWeekDayFormat{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSchedulerMonthlyWeekDayFormatWithDefaults instantiates a new SchedulerMonthlyWeekDayFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulerMonthlyWeekDayFormatWithDefaults() *SchedulerMonthlyWeekDayFormat {
	this := SchedulerMonthlyWeekDayFormat{}
	return &this
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulerMonthlyWeekDayFormat) GetDayOfWeek() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulerMonthlyWeekDayFormat) GetDayOfWeekOk() ([]string, bool) {
	if o == nil || IsNil(o.DayOfWeek) {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *SchedulerMonthlyWeekDayFormat) HasDayOfWeek() bool {
	if o != nil && IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given []string and assigns it to the DayOfWeek field.
func (o *SchedulerMonthlyWeekDayFormat) SetDayOfWeek(v []string) {
	o.DayOfWeek = v
}

// GetWeekOfMonth returns the WeekOfMonth field value if set, zero value otherwise.
func (o *SchedulerMonthlyWeekDayFormat) GetWeekOfMonth() int64 {
	if o == nil || IsNil(o.WeekOfMonth) {
		var ret int64
		return ret
	}
	return *o.WeekOfMonth
}

// GetWeekOfMonthOk returns a tuple with the WeekOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerMonthlyWeekDayFormat) GetWeekOfMonthOk() (*int64, bool) {
	if o == nil || IsNil(o.WeekOfMonth) {
		return nil, false
	}
	return o.WeekOfMonth, true
}

// HasWeekOfMonth returns a boolean if a field has been set.
func (o *SchedulerMonthlyWeekDayFormat) HasWeekOfMonth() bool {
	if o != nil && !IsNil(o.WeekOfMonth) {
		return true
	}

	return false
}

// SetWeekOfMonth gets a reference to the given int64 and assigns it to the WeekOfMonth field.
func (o *SchedulerMonthlyWeekDayFormat) SetWeekOfMonth(v int64) {
	o.WeekOfMonth = &v
}

func (o SchedulerMonthlyWeekDayFormat) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulerMonthlyWeekDayFormat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if o.DayOfWeek != nil {
		toSerialize["DayOfWeek"] = o.DayOfWeek
	}
	if !IsNil(o.WeekOfMonth) {
		toSerialize["WeekOfMonth"] = o.WeekOfMonth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchedulerMonthlyWeekDayFormat) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type SchedulerMonthlyWeekDayFormatWithoutEmbeddedStruct struct {
		DayOfWeek []string `json:"DayOfWeek,omitempty"`
		// The week of the month, 1 through 5.
		WeekOfMonth *int64 `json:"WeekOfMonth,omitempty"`
	}

	varSchedulerMonthlyWeekDayFormatWithoutEmbeddedStruct := SchedulerMonthlyWeekDayFormatWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSchedulerMonthlyWeekDayFormatWithoutEmbeddedStruct)
	if err == nil {
		varSchedulerMonthlyWeekDayFormat := _SchedulerMonthlyWeekDayFormat{}
		varSchedulerMonthlyWeekDayFormat.DayOfWeek = varSchedulerMonthlyWeekDayFormatWithoutEmbeddedStruct.DayOfWeek
		varSchedulerMonthlyWeekDayFormat.WeekOfMonth = varSchedulerMonthlyWeekDayFormatWithoutEmbeddedStruct.WeekOfMonth
		*o = SchedulerMonthlyWeekDayFormat(varSchedulerMonthlyWeekDayFormat)
	} else {
		return err
	}

	varSchedulerMonthlyWeekDayFormat := _SchedulerMonthlyWeekDayFormat{}

	err = json.Unmarshal(data, &varSchedulerMonthlyWeekDayFormat)
	if err == nil {
		o.MoBaseComplexType = varSchedulerMonthlyWeekDayFormat.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "DayOfWeek")
		delete(additionalProperties, "WeekOfMonth")

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

type NullableSchedulerMonthlyWeekDayFormat struct {
	value *SchedulerMonthlyWeekDayFormat
	isSet bool
}

func (v NullableSchedulerMonthlyWeekDayFormat) Get() *SchedulerMonthlyWeekDayFormat {
	return v.value
}

func (v *NullableSchedulerMonthlyWeekDayFormat) Set(val *SchedulerMonthlyWeekDayFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulerMonthlyWeekDayFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulerMonthlyWeekDayFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulerMonthlyWeekDayFormat(val *SchedulerMonthlyWeekDayFormat) *NullableSchedulerMonthlyWeekDayFormat {
	return &NullableSchedulerMonthlyWeekDayFormat{value: val, isSet: true}
}

func (v NullableSchedulerMonthlyWeekDayFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulerMonthlyWeekDayFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}