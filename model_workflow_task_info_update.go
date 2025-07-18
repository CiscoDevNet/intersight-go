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

// checks if the WorkflowTaskInfoUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowTaskInfoUpdate{}

// WorkflowTaskInfoUpdate TaskInfoUpdate is used to capture the updates that are required for the task instance inside an executing workflow. WaitTask, included inside the workflow, can be updated using this. The values provided for task updates will depend on the type of wait task. The status of the task must be provided for all updates. The input values must also be provided for a WaitType that includes an input definition. The input values must satisfy all the constraints specified in the associated input definitions.
type WorkflowTaskInfoUpdate struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Inputs for the specified TaskInfo. Inputs must only be provided for tasks which has included an input definition and the inputs must match the constraints specified in the input definition.
	Input interface{} `json:"Input,omitempty"`
	// Name of the task being updated and this name must match the task instance name included inside the workflow definition. This name is also captured in the RefName property of the TaskInfo object for the task.
	Name *string `json:"Name,omitempty"`
	// New status of the task being updated, only Failed and Completed statuses are supported, Completed is the default value in case no status is provided. * `Scheduled` - The enum represents the status when task is in scheduled state. * `InProgress` - The enum represents the status when task is in-progress state. * `NoOp` - The enum represents the status when task is a noop. * `Timeout` - The enum represents the status when task has timed out. * `Completed` - The enum represents the status when task has completed. * `Failed` - The enum represents the status when task has failed.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTaskInfoUpdate WorkflowTaskInfoUpdate

// NewWorkflowTaskInfoUpdate instantiates a new WorkflowTaskInfoUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskInfoUpdate(classId string, objectType string) *WorkflowTaskInfoUpdate {
	this := WorkflowTaskInfoUpdate{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Scheduled"
	this.Status = &status
	return &this
}

// NewWorkflowTaskInfoUpdateWithDefaults instantiates a new WorkflowTaskInfoUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskInfoUpdateWithDefaults() *WorkflowTaskInfoUpdate {
	this := WorkflowTaskInfoUpdate{}
	var classId string = "workflow.TaskInfoUpdate"
	this.ClassId = classId
	var objectType string = "workflow.TaskInfoUpdate"
	this.ObjectType = objectType
	var status string = "Scheduled"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTaskInfoUpdate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoUpdate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTaskInfoUpdate) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.TaskInfoUpdate" of the ClassId field.
func (o *WorkflowTaskInfoUpdate) GetDefaultClassId() interface{} {
	return "workflow.TaskInfoUpdate"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTaskInfoUpdate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoUpdate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTaskInfoUpdate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.TaskInfoUpdate" of the ObjectType field.
func (o *WorkflowTaskInfoUpdate) GetDefaultObjectType() interface{} {
	return "workflow.TaskInfoUpdate"
}

// GetInput returns the Input field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskInfoUpdate) GetInput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskInfoUpdate) GetInputOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return &o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *WorkflowTaskInfoUpdate) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given interface{} and assigns it to the Input field.
func (o *WorkflowTaskInfoUpdate) SetInput(v interface{}) {
	o.Input = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowTaskInfoUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowTaskInfoUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowTaskInfoUpdate) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowTaskInfoUpdate) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskInfoUpdate) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowTaskInfoUpdate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowTaskInfoUpdate) SetStatus(v string) {
	o.Status = &v
}

func (o WorkflowTaskInfoUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowTaskInfoUpdate) ToMap() (map[string]interface{}, error) {
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
	if o.Input != nil {
		toSerialize["Input"] = o.Input
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowTaskInfoUpdate) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowTaskInfoUpdateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Inputs for the specified TaskInfo. Inputs must only be provided for tasks which has included an input definition and the inputs must match the constraints specified in the input definition.
		Input interface{} `json:"Input,omitempty"`
		// Name of the task being updated and this name must match the task instance name included inside the workflow definition. This name is also captured in the RefName property of the TaskInfo object for the task.
		Name *string `json:"Name,omitempty"`
		// New status of the task being updated, only Failed and Completed statuses are supported, Completed is the default value in case no status is provided. * `Scheduled` - The enum represents the status when task is in scheduled state. * `InProgress` - The enum represents the status when task is in-progress state. * `NoOp` - The enum represents the status when task is a noop. * `Timeout` - The enum represents the status when task has timed out. * `Completed` - The enum represents the status when task has completed. * `Failed` - The enum represents the status when task has failed.
		Status *string `json:"Status,omitempty"`
	}

	varWorkflowTaskInfoUpdateWithoutEmbeddedStruct := WorkflowTaskInfoUpdateWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowTaskInfoUpdateWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowTaskInfoUpdate := _WorkflowTaskInfoUpdate{}
		varWorkflowTaskInfoUpdate.ClassId = varWorkflowTaskInfoUpdateWithoutEmbeddedStruct.ClassId
		varWorkflowTaskInfoUpdate.ObjectType = varWorkflowTaskInfoUpdateWithoutEmbeddedStruct.ObjectType
		varWorkflowTaskInfoUpdate.Input = varWorkflowTaskInfoUpdateWithoutEmbeddedStruct.Input
		varWorkflowTaskInfoUpdate.Name = varWorkflowTaskInfoUpdateWithoutEmbeddedStruct.Name
		varWorkflowTaskInfoUpdate.Status = varWorkflowTaskInfoUpdateWithoutEmbeddedStruct.Status
		*o = WorkflowTaskInfoUpdate(varWorkflowTaskInfoUpdate)
	} else {
		return err
	}

	varWorkflowTaskInfoUpdate := _WorkflowTaskInfoUpdate{}

	err = json.Unmarshal(data, &varWorkflowTaskInfoUpdate)
	if err == nil {
		o.MoBaseComplexType = varWorkflowTaskInfoUpdate.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Input")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Status")

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

type NullableWorkflowTaskInfoUpdate struct {
	value *WorkflowTaskInfoUpdate
	isSet bool
}

func (v NullableWorkflowTaskInfoUpdate) Get() *WorkflowTaskInfoUpdate {
	return v.value
}

func (v *NullableWorkflowTaskInfoUpdate) Set(val *WorkflowTaskInfoUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskInfoUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskInfoUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskInfoUpdate(val *WorkflowTaskInfoUpdate) *NullableWorkflowTaskInfoUpdate {
	return &NullableWorkflowTaskInfoUpdate{value: val, isSet: true}
}

func (v NullableWorkflowTaskInfoUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskInfoUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
