/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2024112619
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

// checks if the ResourceAdaptorsRangeFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAdaptorsRangeFilter{}

// ResourceAdaptorsRangeFilter Consist of two boundaries that define the range, namely the minimum and the maximum adaptors value.
type ResourceAdaptorsRangeFilter struct {
	ResourceRangeFilter
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Maximum value for the range limit. It should be greater than the minimum value.
	MaxValue *int64 `json:"MaxValue,omitempty"`
	// Minimum value for the range limit.
	MinValue             *int64 `json:"MinValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceAdaptorsRangeFilter ResourceAdaptorsRangeFilter

// NewResourceAdaptorsRangeFilter instantiates a new ResourceAdaptorsRangeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAdaptorsRangeFilter(classId string, objectType string) *ResourceAdaptorsRangeFilter {
	this := ResourceAdaptorsRangeFilter{}
	this.ClassId = classId
	this.ObjectType = objectType
	var conditionType string = "RANGE"
	this.ConditionType = &conditionType
	var maxValue int64 = 0
	this.MaxValue = &maxValue
	var minValue int64 = 0
	this.MinValue = &minValue
	return &this
}

// NewResourceAdaptorsRangeFilterWithDefaults instantiates a new ResourceAdaptorsRangeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAdaptorsRangeFilterWithDefaults() *ResourceAdaptorsRangeFilter {
	this := ResourceAdaptorsRangeFilter{}
	var classId string = "resource.AdaptorsRangeFilter"
	this.ClassId = classId
	var objectType string = "resource.AdaptorsRangeFilter"
	this.ObjectType = objectType
	var maxValue int64 = 0
	this.MaxValue = &maxValue
	var minValue int64 = 0
	this.MinValue = &minValue
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceAdaptorsRangeFilter) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceAdaptorsRangeFilter) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceAdaptorsRangeFilter) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "resource.AdaptorsRangeFilter" of the ClassId field.
func (o *ResourceAdaptorsRangeFilter) GetDefaultClassId() interface{} {
	return "resource.AdaptorsRangeFilter"
}

// GetObjectType returns the ObjectType field value
func (o *ResourceAdaptorsRangeFilter) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceAdaptorsRangeFilter) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceAdaptorsRangeFilter) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "resource.AdaptorsRangeFilter" of the ObjectType field.
func (o *ResourceAdaptorsRangeFilter) GetDefaultObjectType() interface{} {
	return "resource.AdaptorsRangeFilter"
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *ResourceAdaptorsRangeFilter) GetMaxValue() int64 {
	if o == nil || IsNil(o.MaxValue) {
		var ret int64
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAdaptorsRangeFilter) GetMaxValueOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxValue) {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *ResourceAdaptorsRangeFilter) HasMaxValue() bool {
	if o != nil && !IsNil(o.MaxValue) {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given int64 and assigns it to the MaxValue field.
func (o *ResourceAdaptorsRangeFilter) SetMaxValue(v int64) {
	o.MaxValue = &v
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *ResourceAdaptorsRangeFilter) GetMinValue() int64 {
	if o == nil || IsNil(o.MinValue) {
		var ret int64
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAdaptorsRangeFilter) GetMinValueOk() (*int64, bool) {
	if o == nil || IsNil(o.MinValue) {
		return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *ResourceAdaptorsRangeFilter) HasMinValue() bool {
	if o != nil && !IsNil(o.MinValue) {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given int64 and assigns it to the MinValue field.
func (o *ResourceAdaptorsRangeFilter) SetMinValue(v int64) {
	o.MinValue = &v
}

func (o ResourceAdaptorsRangeFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAdaptorsRangeFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedResourceRangeFilter, errResourceRangeFilter := json.Marshal(o.ResourceRangeFilter)
	if errResourceRangeFilter != nil {
		return map[string]interface{}{}, errResourceRangeFilter
	}
	errResourceRangeFilter = json.Unmarshal([]byte(serializedResourceRangeFilter), &toSerialize)
	if errResourceRangeFilter != nil {
		return map[string]interface{}{}, errResourceRangeFilter
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.MaxValue) {
		toSerialize["MaxValue"] = o.MaxValue
	}
	if !IsNil(o.MinValue) {
		toSerialize["MinValue"] = o.MinValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceAdaptorsRangeFilter) UnmarshalJSON(data []byte) (err error) {
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
	type ResourceAdaptorsRangeFilterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Maximum value for the range limit. It should be greater than the minimum value.
		MaxValue *int64 `json:"MaxValue,omitempty"`
		// Minimum value for the range limit.
		MinValue *int64 `json:"MinValue,omitempty"`
	}

	varResourceAdaptorsRangeFilterWithoutEmbeddedStruct := ResourceAdaptorsRangeFilterWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varResourceAdaptorsRangeFilterWithoutEmbeddedStruct)
	if err == nil {
		varResourceAdaptorsRangeFilter := _ResourceAdaptorsRangeFilter{}
		varResourceAdaptorsRangeFilter.ClassId = varResourceAdaptorsRangeFilterWithoutEmbeddedStruct.ClassId
		varResourceAdaptorsRangeFilter.ObjectType = varResourceAdaptorsRangeFilterWithoutEmbeddedStruct.ObjectType
		varResourceAdaptorsRangeFilter.MaxValue = varResourceAdaptorsRangeFilterWithoutEmbeddedStruct.MaxValue
		varResourceAdaptorsRangeFilter.MinValue = varResourceAdaptorsRangeFilterWithoutEmbeddedStruct.MinValue
		*o = ResourceAdaptorsRangeFilter(varResourceAdaptorsRangeFilter)
	} else {
		return err
	}

	varResourceAdaptorsRangeFilter := _ResourceAdaptorsRangeFilter{}

	err = json.Unmarshal(data, &varResourceAdaptorsRangeFilter)
	if err == nil {
		o.ResourceRangeFilter = varResourceAdaptorsRangeFilter.ResourceRangeFilter
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MaxValue")
		delete(additionalProperties, "MinValue")

		// remove fields from embedded structs
		reflectResourceRangeFilter := reflect.ValueOf(o.ResourceRangeFilter)
		for i := 0; i < reflectResourceRangeFilter.Type().NumField(); i++ {
			t := reflectResourceRangeFilter.Type().Field(i)

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

type NullableResourceAdaptorsRangeFilter struct {
	value *ResourceAdaptorsRangeFilter
	isSet bool
}

func (v NullableResourceAdaptorsRangeFilter) Get() *ResourceAdaptorsRangeFilter {
	return v.value
}

func (v *NullableResourceAdaptorsRangeFilter) Set(val *ResourceAdaptorsRangeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAdaptorsRangeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAdaptorsRangeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAdaptorsRangeFilter(val *ResourceAdaptorsRangeFilter) *NullableResourceAdaptorsRangeFilter {
	return &NullableResourceAdaptorsRangeFilter{value: val, isSet: true}
}

func (v NullableResourceAdaptorsRangeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAdaptorsRangeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}