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

// checks if the IaasCustomTaskInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IaasCustomTaskInfo{}

// IaasCustomTaskInfo List out the execution of the Custom Tasks with Names.
type IaasCustomTaskInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Task decription or Comment of the Custom task.
	TaskDescription *string `json:"TaskDescription,omitempty"`
	// Number of times this task has executed.
	TaskExecutionCount *int64 `json:"TaskExecutionCount,omitempty"`
	// Task Label in the Workflow.
	TaskLabel *string `json:"TaskLabel,omitempty"`
	// Name of the Custom Task in UCSD.
	TaskName             *string                          `json:"TaskName,omitempty"`
	Guid                 NullableIaasUcsdInfoRelationship `json:"Guid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasCustomTaskInfo IaasCustomTaskInfo

// NewIaasCustomTaskInfo instantiates a new IaasCustomTaskInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasCustomTaskInfo(classId string, objectType string) *IaasCustomTaskInfo {
	this := IaasCustomTaskInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasCustomTaskInfoWithDefaults instantiates a new IaasCustomTaskInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasCustomTaskInfoWithDefaults() *IaasCustomTaskInfo {
	this := IaasCustomTaskInfo{}
	var classId string = "iaas.CustomTaskInfo"
	this.ClassId = classId
	var objectType string = "iaas.CustomTaskInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasCustomTaskInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasCustomTaskInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iaas.CustomTaskInfo" of the ClassId field.
func (o *IaasCustomTaskInfo) GetDefaultClassId() interface{} {
	return "iaas.CustomTaskInfo"
}

// GetObjectType returns the ObjectType field value
func (o *IaasCustomTaskInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasCustomTaskInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iaas.CustomTaskInfo" of the ObjectType field.
func (o *IaasCustomTaskInfo) GetDefaultObjectType() interface{} {
	return "iaas.CustomTaskInfo"
}

// GetTaskDescription returns the TaskDescription field value if set, zero value otherwise.
func (o *IaasCustomTaskInfo) GetTaskDescription() string {
	if o == nil || IsNil(o.TaskDescription) {
		var ret string
		return ret
	}
	return *o.TaskDescription
}

// GetTaskDescriptionOk returns a tuple with the TaskDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetTaskDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.TaskDescription) {
		return nil, false
	}
	return o.TaskDescription, true
}

// HasTaskDescription returns a boolean if a field has been set.
func (o *IaasCustomTaskInfo) HasTaskDescription() bool {
	if o != nil && !IsNil(o.TaskDescription) {
		return true
	}

	return false
}

// SetTaskDescription gets a reference to the given string and assigns it to the TaskDescription field.
func (o *IaasCustomTaskInfo) SetTaskDescription(v string) {
	o.TaskDescription = &v
}

// GetTaskExecutionCount returns the TaskExecutionCount field value if set, zero value otherwise.
func (o *IaasCustomTaskInfo) GetTaskExecutionCount() int64 {
	if o == nil || IsNil(o.TaskExecutionCount) {
		var ret int64
		return ret
	}
	return *o.TaskExecutionCount
}

// GetTaskExecutionCountOk returns a tuple with the TaskExecutionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetTaskExecutionCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TaskExecutionCount) {
		return nil, false
	}
	return o.TaskExecutionCount, true
}

// HasTaskExecutionCount returns a boolean if a field has been set.
func (o *IaasCustomTaskInfo) HasTaskExecutionCount() bool {
	if o != nil && !IsNil(o.TaskExecutionCount) {
		return true
	}

	return false
}

// SetTaskExecutionCount gets a reference to the given int64 and assigns it to the TaskExecutionCount field.
func (o *IaasCustomTaskInfo) SetTaskExecutionCount(v int64) {
	o.TaskExecutionCount = &v
}

// GetTaskLabel returns the TaskLabel field value if set, zero value otherwise.
func (o *IaasCustomTaskInfo) GetTaskLabel() string {
	if o == nil || IsNil(o.TaskLabel) {
		var ret string
		return ret
	}
	return *o.TaskLabel
}

// GetTaskLabelOk returns a tuple with the TaskLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetTaskLabelOk() (*string, bool) {
	if o == nil || IsNil(o.TaskLabel) {
		return nil, false
	}
	return o.TaskLabel, true
}

// HasTaskLabel returns a boolean if a field has been set.
func (o *IaasCustomTaskInfo) HasTaskLabel() bool {
	if o != nil && !IsNil(o.TaskLabel) {
		return true
	}

	return false
}

// SetTaskLabel gets a reference to the given string and assigns it to the TaskLabel field.
func (o *IaasCustomTaskInfo) SetTaskLabel(v string) {
	o.TaskLabel = &v
}

// GetTaskName returns the TaskName field value if set, zero value otherwise.
func (o *IaasCustomTaskInfo) GetTaskName() string {
	if o == nil || IsNil(o.TaskName) {
		var ret string
		return ret
	}
	return *o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasCustomTaskInfo) GetTaskNameOk() (*string, bool) {
	if o == nil || IsNil(o.TaskName) {
		return nil, false
	}
	return o.TaskName, true
}

// HasTaskName returns a boolean if a field has been set.
func (o *IaasCustomTaskInfo) HasTaskName() bool {
	if o != nil && !IsNil(o.TaskName) {
		return true
	}

	return false
}

// SetTaskName gets a reference to the given string and assigns it to the TaskName field.
func (o *IaasCustomTaskInfo) SetTaskName(v string) {
	o.TaskName = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasCustomTaskInfo) GetGuid() IaasUcsdInfoRelationship {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasCustomTaskInfo) GetGuidOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasCustomTaskInfo) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableIaasUcsdInfoRelationship and assigns it to the Guid field.
func (o *IaasCustomTaskInfo) SetGuid(v IaasUcsdInfoRelationship) {
	o.Guid.Set(&v)
}

// SetGuidNil sets the value for Guid to be an explicit nil
func (o *IaasCustomTaskInfo) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *IaasCustomTaskInfo) UnsetGuid() {
	o.Guid.Unset()
}

func (o IaasCustomTaskInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IaasCustomTaskInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.TaskDescription) {
		toSerialize["TaskDescription"] = o.TaskDescription
	}
	if !IsNil(o.TaskExecutionCount) {
		toSerialize["TaskExecutionCount"] = o.TaskExecutionCount
	}
	if !IsNil(o.TaskLabel) {
		toSerialize["TaskLabel"] = o.TaskLabel
	}
	if !IsNil(o.TaskName) {
		toSerialize["TaskName"] = o.TaskName
	}
	if o.Guid.IsSet() {
		toSerialize["Guid"] = o.Guid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IaasCustomTaskInfo) UnmarshalJSON(data []byte) (err error) {
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
	type IaasCustomTaskInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Task decription or Comment of the Custom task.
		TaskDescription *string `json:"TaskDescription,omitempty"`
		// Number of times this task has executed.
		TaskExecutionCount *int64 `json:"TaskExecutionCount,omitempty"`
		// Task Label in the Workflow.
		TaskLabel *string `json:"TaskLabel,omitempty"`
		// Name of the Custom Task in UCSD.
		TaskName *string                          `json:"TaskName,omitempty"`
		Guid     NullableIaasUcsdInfoRelationship `json:"Guid,omitempty"`
	}

	varIaasCustomTaskInfoWithoutEmbeddedStruct := IaasCustomTaskInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIaasCustomTaskInfoWithoutEmbeddedStruct)
	if err == nil {
		varIaasCustomTaskInfo := _IaasCustomTaskInfo{}
		varIaasCustomTaskInfo.ClassId = varIaasCustomTaskInfoWithoutEmbeddedStruct.ClassId
		varIaasCustomTaskInfo.ObjectType = varIaasCustomTaskInfoWithoutEmbeddedStruct.ObjectType
		varIaasCustomTaskInfo.TaskDescription = varIaasCustomTaskInfoWithoutEmbeddedStruct.TaskDescription
		varIaasCustomTaskInfo.TaskExecutionCount = varIaasCustomTaskInfoWithoutEmbeddedStruct.TaskExecutionCount
		varIaasCustomTaskInfo.TaskLabel = varIaasCustomTaskInfoWithoutEmbeddedStruct.TaskLabel
		varIaasCustomTaskInfo.TaskName = varIaasCustomTaskInfoWithoutEmbeddedStruct.TaskName
		varIaasCustomTaskInfo.Guid = varIaasCustomTaskInfoWithoutEmbeddedStruct.Guid
		*o = IaasCustomTaskInfo(varIaasCustomTaskInfo)
	} else {
		return err
	}

	varIaasCustomTaskInfo := _IaasCustomTaskInfo{}

	err = json.Unmarshal(data, &varIaasCustomTaskInfo)
	if err == nil {
		o.MoBaseMo = varIaasCustomTaskInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TaskDescription")
		delete(additionalProperties, "TaskExecutionCount")
		delete(additionalProperties, "TaskLabel")
		delete(additionalProperties, "TaskName")
		delete(additionalProperties, "Guid")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIaasCustomTaskInfo struct {
	value *IaasCustomTaskInfo
	isSet bool
}

func (v NullableIaasCustomTaskInfo) Get() *IaasCustomTaskInfo {
	return v.value
}

func (v *NullableIaasCustomTaskInfo) Set(val *IaasCustomTaskInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasCustomTaskInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasCustomTaskInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasCustomTaskInfo(val *IaasCustomTaskInfo) *NullableIaasCustomTaskInfo {
	return &NullableIaasCustomTaskInfo{value: val, isSet: true}
}

func (v NullableIaasCustomTaskInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasCustomTaskInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
