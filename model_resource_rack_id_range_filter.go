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

// checks if the ResourceRackIdRangeFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRackIdRangeFilter{}

// ResourceRackIdRangeFilter Consist of two boundaries that define the range, namely the minimum and the maximum RackId value.
type ResourceRackIdRangeFilter struct {
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

type _ResourceRackIdRangeFilter ResourceRackIdRangeFilter

// NewResourceRackIdRangeFilter instantiates a new ResourceRackIdRangeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRackIdRangeFilter(classId string, objectType string) *ResourceRackIdRangeFilter {
	this := ResourceRackIdRangeFilter{}
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

// NewResourceRackIdRangeFilterWithDefaults instantiates a new ResourceRackIdRangeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRackIdRangeFilterWithDefaults() *ResourceRackIdRangeFilter {
	this := ResourceRackIdRangeFilter{}
	var classId string = "resource.RackIdRangeFilter"
	this.ClassId = classId
	var objectType string = "resource.RackIdRangeFilter"
	this.ObjectType = objectType
	var maxValue int64 = 0
	this.MaxValue = &maxValue
	var minValue int64 = 0
	this.MinValue = &minValue
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceRackIdRangeFilter) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceRackIdRangeFilter) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceRackIdRangeFilter) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "resource.RackIdRangeFilter" of the ClassId field.
func (o *ResourceRackIdRangeFilter) GetDefaultClassId() interface{} {
	return "resource.RackIdRangeFilter"
}

// GetObjectType returns the ObjectType field value
func (o *ResourceRackIdRangeFilter) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceRackIdRangeFilter) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceRackIdRangeFilter) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "resource.RackIdRangeFilter" of the ObjectType field.
func (o *ResourceRackIdRangeFilter) GetDefaultObjectType() interface{} {
	return "resource.RackIdRangeFilter"
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *ResourceRackIdRangeFilter) GetMaxValue() int64 {
	if o == nil || IsNil(o.MaxValue) {
		var ret int64
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRackIdRangeFilter) GetMaxValueOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxValue) {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *ResourceRackIdRangeFilter) HasMaxValue() bool {
	if o != nil && !IsNil(o.MaxValue) {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given int64 and assigns it to the MaxValue field.
func (o *ResourceRackIdRangeFilter) SetMaxValue(v int64) {
	o.MaxValue = &v
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *ResourceRackIdRangeFilter) GetMinValue() int64 {
	if o == nil || IsNil(o.MinValue) {
		var ret int64
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRackIdRangeFilter) GetMinValueOk() (*int64, bool) {
	if o == nil || IsNil(o.MinValue) {
		return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *ResourceRackIdRangeFilter) HasMinValue() bool {
	if o != nil && !IsNil(o.MinValue) {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given int64 and assigns it to the MinValue field.
func (o *ResourceRackIdRangeFilter) SetMinValue(v int64) {
	o.MinValue = &v
}

func (o ResourceRackIdRangeFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRackIdRangeFilter) ToMap() (map[string]interface{}, error) {
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

func (o *ResourceRackIdRangeFilter) UnmarshalJSON(data []byte) (err error) {
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
	type ResourceRackIdRangeFilterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Maximum value for the range limit. It should be greater than the minimum value.
		MaxValue *int64 `json:"MaxValue,omitempty"`
		// Minimum value for the range limit.
		MinValue *int64 `json:"MinValue,omitempty"`
	}

	varResourceRackIdRangeFilterWithoutEmbeddedStruct := ResourceRackIdRangeFilterWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varResourceRackIdRangeFilterWithoutEmbeddedStruct)
	if err == nil {
		varResourceRackIdRangeFilter := _ResourceRackIdRangeFilter{}
		varResourceRackIdRangeFilter.ClassId = varResourceRackIdRangeFilterWithoutEmbeddedStruct.ClassId
		varResourceRackIdRangeFilter.ObjectType = varResourceRackIdRangeFilterWithoutEmbeddedStruct.ObjectType
		varResourceRackIdRangeFilter.MaxValue = varResourceRackIdRangeFilterWithoutEmbeddedStruct.MaxValue
		varResourceRackIdRangeFilter.MinValue = varResourceRackIdRangeFilterWithoutEmbeddedStruct.MinValue
		*o = ResourceRackIdRangeFilter(varResourceRackIdRangeFilter)
	} else {
		return err
	}

	varResourceRackIdRangeFilter := _ResourceRackIdRangeFilter{}

	err = json.Unmarshal(data, &varResourceRackIdRangeFilter)
	if err == nil {
		o.ResourceRangeFilter = varResourceRackIdRangeFilter.ResourceRangeFilter
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

type NullableResourceRackIdRangeFilter struct {
	value *ResourceRackIdRangeFilter
	isSet bool
}

func (v NullableResourceRackIdRangeFilter) Get() *ResourceRackIdRangeFilter {
	return v.value
}

func (v *NullableResourceRackIdRangeFilter) Set(val *ResourceRackIdRangeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRackIdRangeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRackIdRangeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRackIdRangeFilter(val *ResourceRackIdRangeFilter) *NullableResourceRackIdRangeFilter {
	return &NullableResourceRackIdRangeFilter{value: val, isSet: true}
}

func (v NullableResourceRackIdRangeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRackIdRangeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}