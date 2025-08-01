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

// checks if the WorkflowSuccessEndTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowSuccessEndTask{}

// WorkflowSuccessEndTask A SuccessEndTask denotes the successful completion of a workflow.
type WorkflowSuccessEndTask struct {
	WorkflowEndTask
	AdditionalProperties map[string]interface{}
}

type _WorkflowSuccessEndTask WorkflowSuccessEndTask

// NewWorkflowSuccessEndTask instantiates a new WorkflowSuccessEndTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSuccessEndTask(classId string, objectType string) *WorkflowSuccessEndTask {
	this := WorkflowSuccessEndTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowSuccessEndTaskWithDefaults instantiates a new WorkflowSuccessEndTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSuccessEndTaskWithDefaults() *WorkflowSuccessEndTask {
	this := WorkflowSuccessEndTask{}
	return &this
}

func (o WorkflowSuccessEndTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowSuccessEndTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowEndTask, errWorkflowEndTask := json.Marshal(o.WorkflowEndTask)
	if errWorkflowEndTask != nil {
		return map[string]interface{}{}, errWorkflowEndTask
	}
	errWorkflowEndTask = json.Unmarshal([]byte(serializedWorkflowEndTask), &toSerialize)
	if errWorkflowEndTask != nil {
		return map[string]interface{}{}, errWorkflowEndTask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowSuccessEndTask) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowSuccessEndTaskWithoutEmbeddedStruct struct {
	}

	varWorkflowSuccessEndTaskWithoutEmbeddedStruct := WorkflowSuccessEndTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowSuccessEndTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowSuccessEndTask := _WorkflowSuccessEndTask{}
		*o = WorkflowSuccessEndTask(varWorkflowSuccessEndTask)
	} else {
		return err
	}

	varWorkflowSuccessEndTask := _WorkflowSuccessEndTask{}

	err = json.Unmarshal(data, &varWorkflowSuccessEndTask)
	if err == nil {
		o.WorkflowEndTask = varWorkflowSuccessEndTask.WorkflowEndTask
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectWorkflowEndTask := reflect.ValueOf(o.WorkflowEndTask)
		for i := 0; i < reflectWorkflowEndTask.Type().NumField(); i++ {
			t := reflectWorkflowEndTask.Type().Field(i)

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

type NullableWorkflowSuccessEndTask struct {
	value *WorkflowSuccessEndTask
	isSet bool
}

func (v NullableWorkflowSuccessEndTask) Get() *WorkflowSuccessEndTask {
	return v.value
}

func (v *NullableWorkflowSuccessEndTask) Set(val *WorkflowSuccessEndTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSuccessEndTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSuccessEndTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSuccessEndTask(val *WorkflowSuccessEndTask) *NullableWorkflowSuccessEndTask {
	return &NullableWorkflowSuccessEndTask{value: val, isSet: true}
}

func (v NullableWorkflowSuccessEndTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSuccessEndTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
