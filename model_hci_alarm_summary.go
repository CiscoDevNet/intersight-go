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

// checks if the HciAlarmSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HciAlarmSummary{}

// HciAlarmSummary The summary of alarm counts based on alarm severity.
type HciAlarmSummary struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The count of alarms that have severity type Critical.
	Critical *int64 `json:"Critical,omitempty"`
	// The health of the managed endpoint. The health is computed from the highest-severity alarm raised on the endpoint. * `Healthy` - The Enum value represents that the entity is healthy. * `Warning` - The Enum value Warning represents that the entity has one or more active warnings on it. * `Critical` - The Enum value Critical represents that the entity is in a critical state.
	Health *string `json:"Health,omitempty"`
	// The count of alarms that have severity type Info.
	Info *int64 `json:"Info,omitempty"`
	// The count of alarms that have severity type Warning.
	Warning              *int64 `json:"Warning,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HciAlarmSummary HciAlarmSummary

// NewHciAlarmSummary instantiates a new HciAlarmSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHciAlarmSummary(classId string, objectType string) *HciAlarmSummary {
	this := HciAlarmSummary{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHciAlarmSummaryWithDefaults instantiates a new HciAlarmSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHciAlarmSummaryWithDefaults() *HciAlarmSummary {
	this := HciAlarmSummary{}
	var classId string = "hci.AlarmSummary"
	this.ClassId = classId
	var objectType string = "hci.AlarmSummary"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HciAlarmSummary) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HciAlarmSummary) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HciAlarmSummary) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hci.AlarmSummary" of the ClassId field.
func (o *HciAlarmSummary) GetDefaultClassId() interface{} {
	return "hci.AlarmSummary"
}

// GetObjectType returns the ObjectType field value
func (o *HciAlarmSummary) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HciAlarmSummary) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HciAlarmSummary) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hci.AlarmSummary" of the ObjectType field.
func (o *HciAlarmSummary) GetDefaultObjectType() interface{} {
	return "hci.AlarmSummary"
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *HciAlarmSummary) GetCritical() int64 {
	if o == nil || IsNil(o.Critical) {
		var ret int64
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciAlarmSummary) GetCriticalOk() (*int64, bool) {
	if o == nil || IsNil(o.Critical) {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *HciAlarmSummary) HasCritical() bool {
	if o != nil && !IsNil(o.Critical) {
		return true
	}

	return false
}

// SetCritical gets a reference to the given int64 and assigns it to the Critical field.
func (o *HciAlarmSummary) SetCritical(v int64) {
	o.Critical = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *HciAlarmSummary) GetHealth() string {
	if o == nil || IsNil(o.Health) {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciAlarmSummary) GetHealthOk() (*string, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *HciAlarmSummary) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *HciAlarmSummary) SetHealth(v string) {
	o.Health = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *HciAlarmSummary) GetInfo() int64 {
	if o == nil || IsNil(o.Info) {
		var ret int64
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciAlarmSummary) GetInfoOk() (*int64, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *HciAlarmSummary) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given int64 and assigns it to the Info field.
func (o *HciAlarmSummary) SetInfo(v int64) {
	o.Info = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *HciAlarmSummary) GetWarning() int64 {
	if o == nil || IsNil(o.Warning) {
		var ret int64
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciAlarmSummary) GetWarningOk() (*int64, bool) {
	if o == nil || IsNil(o.Warning) {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *HciAlarmSummary) HasWarning() bool {
	if o != nil && !IsNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given int64 and assigns it to the Warning field.
func (o *HciAlarmSummary) SetWarning(v int64) {
	o.Warning = &v
}

func (o HciAlarmSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HciAlarmSummary) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Critical) {
		toSerialize["Critical"] = o.Critical
	}
	if !IsNil(o.Health) {
		toSerialize["Health"] = o.Health
	}
	if !IsNil(o.Info) {
		toSerialize["Info"] = o.Info
	}
	if !IsNil(o.Warning) {
		toSerialize["Warning"] = o.Warning
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HciAlarmSummary) UnmarshalJSON(data []byte) (err error) {
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
	type HciAlarmSummaryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The count of alarms that have severity type Critical.
		Critical *int64 `json:"Critical,omitempty"`
		// The health of the managed endpoint. The health is computed from the highest-severity alarm raised on the endpoint. * `Healthy` - The Enum value represents that the entity is healthy. * `Warning` - The Enum value Warning represents that the entity has one or more active warnings on it. * `Critical` - The Enum value Critical represents that the entity is in a critical state.
		Health *string `json:"Health,omitempty"`
		// The count of alarms that have severity type Info.
		Info *int64 `json:"Info,omitempty"`
		// The count of alarms that have severity type Warning.
		Warning *int64 `json:"Warning,omitempty"`
	}

	varHciAlarmSummaryWithoutEmbeddedStruct := HciAlarmSummaryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHciAlarmSummaryWithoutEmbeddedStruct)
	if err == nil {
		varHciAlarmSummary := _HciAlarmSummary{}
		varHciAlarmSummary.ClassId = varHciAlarmSummaryWithoutEmbeddedStruct.ClassId
		varHciAlarmSummary.ObjectType = varHciAlarmSummaryWithoutEmbeddedStruct.ObjectType
		varHciAlarmSummary.Critical = varHciAlarmSummaryWithoutEmbeddedStruct.Critical
		varHciAlarmSummary.Health = varHciAlarmSummaryWithoutEmbeddedStruct.Health
		varHciAlarmSummary.Info = varHciAlarmSummaryWithoutEmbeddedStruct.Info
		varHciAlarmSummary.Warning = varHciAlarmSummaryWithoutEmbeddedStruct.Warning
		*o = HciAlarmSummary(varHciAlarmSummary)
	} else {
		return err
	}

	varHciAlarmSummary := _HciAlarmSummary{}

	err = json.Unmarshal(data, &varHciAlarmSummary)
	if err == nil {
		o.MoBaseComplexType = varHciAlarmSummary.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Critical")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "Info")
		delete(additionalProperties, "Warning")

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

type NullableHciAlarmSummary struct {
	value *HciAlarmSummary
	isSet bool
}

func (v NullableHciAlarmSummary) Get() *HciAlarmSummary {
	return v.value
}

func (v *NullableHciAlarmSummary) Set(val *HciAlarmSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableHciAlarmSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableHciAlarmSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHciAlarmSummary(val *HciAlarmSummary) *NullableHciAlarmSummary {
	return &NullableHciAlarmSummary{value: val, isSet: true}
}

func (v NullableHciAlarmSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHciAlarmSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
