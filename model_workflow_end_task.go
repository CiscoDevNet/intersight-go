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

// checks if the WorkflowEndTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowEndTask{}

// WorkflowEndTask An EndTask denotes the completion of a workflow.
type WorkflowEndTask struct {
	WorkflowControlTask
	AdditionalProperties map[string]interface{}
}

type _WorkflowEndTask WorkflowEndTask

// NewWorkflowEndTask instantiates a new WorkflowEndTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowEndTask(classId string, objectType string) *WorkflowEndTask {
	this := WorkflowEndTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowEndTaskWithDefaults instantiates a new WorkflowEndTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowEndTaskWithDefaults() *WorkflowEndTask {
	this := WorkflowEndTask{}
	return &this
}

func (o WorkflowEndTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowEndTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowControlTask, errWorkflowControlTask := json.Marshal(o.WorkflowControlTask)
	if errWorkflowControlTask != nil {
		return map[string]interface{}{}, errWorkflowControlTask
	}
	errWorkflowControlTask = json.Unmarshal([]byte(serializedWorkflowControlTask), &toSerialize)
	if errWorkflowControlTask != nil {
		return map[string]interface{}{}, errWorkflowControlTask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowEndTask) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowEndTaskWithoutEmbeddedStruct struct {
	}

	varWorkflowEndTaskWithoutEmbeddedStruct := WorkflowEndTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowEndTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowEndTask := _WorkflowEndTask{}
		*o = WorkflowEndTask(varWorkflowEndTask)
	} else {
		return err
	}

	varWorkflowEndTask := _WorkflowEndTask{}

	err = json.Unmarshal(data, &varWorkflowEndTask)
	if err == nil {
		o.WorkflowControlTask = varWorkflowEndTask.WorkflowControlTask
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectWorkflowControlTask := reflect.ValueOf(o.WorkflowControlTask)
		for i := 0; i < reflectWorkflowControlTask.Type().NumField(); i++ {
			t := reflectWorkflowControlTask.Type().Field(i)

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

type NullableWorkflowEndTask struct {
	value *WorkflowEndTask
	isSet bool
}

func (v NullableWorkflowEndTask) Get() *WorkflowEndTask {
	return v.value
}

func (v *NullableWorkflowEndTask) Set(val *WorkflowEndTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowEndTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowEndTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowEndTask(val *WorkflowEndTask) *NullableWorkflowEndTask {
	return &NullableWorkflowEndTask{value: val, isSet: true}
}

func (v NullableWorkflowEndTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowEndTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
