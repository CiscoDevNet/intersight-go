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

// checks if the WorkflowControlTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowControlTask{}

// WorkflowControlTask Control tasks run within the workflow engine and are not executed by remote worker microservices. They are used to modify flow of control within a workflow.
type WorkflowControlTask struct {
	WorkflowWorkflowTask
	AdditionalProperties map[string]interface{}
}

type _WorkflowControlTask WorkflowControlTask

// NewWorkflowControlTask instantiates a new WorkflowControlTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowControlTask(classId string, objectType string) *WorkflowControlTask {
	this := WorkflowControlTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowControlTaskWithDefaults instantiates a new WorkflowControlTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowControlTaskWithDefaults() *WorkflowControlTask {
	this := WorkflowControlTask{}
	return &this
}

func (o WorkflowControlTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowControlTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowWorkflowTask, errWorkflowWorkflowTask := json.Marshal(o.WorkflowWorkflowTask)
	if errWorkflowWorkflowTask != nil {
		return map[string]interface{}{}, errWorkflowWorkflowTask
	}
	errWorkflowWorkflowTask = json.Unmarshal([]byte(serializedWorkflowWorkflowTask), &toSerialize)
	if errWorkflowWorkflowTask != nil {
		return map[string]interface{}{}, errWorkflowWorkflowTask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowControlTask) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
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
	type WorkflowControlTaskWithoutEmbeddedStruct struct {
	}

	varWorkflowControlTaskWithoutEmbeddedStruct := WorkflowControlTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowControlTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowControlTask := _WorkflowControlTask{}
		*o = WorkflowControlTask(varWorkflowControlTask)
	} else {
		return err
	}

	varWorkflowControlTask := _WorkflowControlTask{}

	err = json.Unmarshal(data, &varWorkflowControlTask)
	if err == nil {
		o.WorkflowWorkflowTask = varWorkflowControlTask.WorkflowWorkflowTask
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectWorkflowWorkflowTask := reflect.ValueOf(o.WorkflowWorkflowTask)
		for i := 0; i < reflectWorkflowWorkflowTask.Type().NumField(); i++ {
			t := reflectWorkflowWorkflowTask.Type().Field(i)

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

type NullableWorkflowControlTask struct {
	value *WorkflowControlTask
	isSet bool
}

func (v NullableWorkflowControlTask) Get() *WorkflowControlTask {
	return v.value
}

func (v *NullableWorkflowControlTask) Set(val *WorkflowControlTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowControlTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowControlTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowControlTask(val *WorkflowControlTask) *NullableWorkflowControlTask {
	return &NullableWorkflowControlTask{value: val, isSet: true}
}

func (v NullableWorkflowControlTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowControlTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
