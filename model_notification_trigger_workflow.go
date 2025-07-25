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

// checks if the NotificationTriggerWorkflow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationTriggerWorkflow{}

// NotificationTriggerWorkflow TriggerWorkflow complexType holds the workflow task info.
type NotificationTriggerWorkflow struct {
	NotificationAction
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Target workflow definition moid to trigger.
	WorkflowDefinitionMoid *string `json:"WorkflowDefinitionMoid,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _NotificationTriggerWorkflow NotificationTriggerWorkflow

// NewNotificationTriggerWorkflow instantiates a new NotificationTriggerWorkflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationTriggerWorkflow(classId string, objectType string) *NotificationTriggerWorkflow {
	this := NotificationTriggerWorkflow{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNotificationTriggerWorkflowWithDefaults instantiates a new NotificationTriggerWorkflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationTriggerWorkflowWithDefaults() *NotificationTriggerWorkflow {
	this := NotificationTriggerWorkflow{}
	var classId string = "notification.TriggerWorkflow"
	this.ClassId = classId
	var objectType string = "notification.TriggerWorkflow"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationTriggerWorkflow) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationTriggerWorkflow) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationTriggerWorkflow) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "notification.TriggerWorkflow" of the ClassId field.
func (o *NotificationTriggerWorkflow) GetDefaultClassId() interface{} {
	return "notification.TriggerWorkflow"
}

// GetObjectType returns the ObjectType field value
func (o *NotificationTriggerWorkflow) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationTriggerWorkflow) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationTriggerWorkflow) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "notification.TriggerWorkflow" of the ObjectType field.
func (o *NotificationTriggerWorkflow) GetDefaultObjectType() interface{} {
	return "notification.TriggerWorkflow"
}

// GetWorkflowDefinitionMoid returns the WorkflowDefinitionMoid field value if set, zero value otherwise.
func (o *NotificationTriggerWorkflow) GetWorkflowDefinitionMoid() string {
	if o == nil || IsNil(o.WorkflowDefinitionMoid) {
		var ret string
		return ret
	}
	return *o.WorkflowDefinitionMoid
}

// GetWorkflowDefinitionMoidOk returns a tuple with the WorkflowDefinitionMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTriggerWorkflow) GetWorkflowDefinitionMoidOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowDefinitionMoid) {
		return nil, false
	}
	return o.WorkflowDefinitionMoid, true
}

// HasWorkflowDefinitionMoid returns a boolean if a field has been set.
func (o *NotificationTriggerWorkflow) HasWorkflowDefinitionMoid() bool {
	if o != nil && !IsNil(o.WorkflowDefinitionMoid) {
		return true
	}

	return false
}

// SetWorkflowDefinitionMoid gets a reference to the given string and assigns it to the WorkflowDefinitionMoid field.
func (o *NotificationTriggerWorkflow) SetWorkflowDefinitionMoid(v string) {
	o.WorkflowDefinitionMoid = &v
}

func (o NotificationTriggerWorkflow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationTriggerWorkflow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNotificationAction, errNotificationAction := json.Marshal(o.NotificationAction)
	if errNotificationAction != nil {
		return map[string]interface{}{}, errNotificationAction
	}
	errNotificationAction = json.Unmarshal([]byte(serializedNotificationAction), &toSerialize)
	if errNotificationAction != nil {
		return map[string]interface{}{}, errNotificationAction
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.WorkflowDefinitionMoid) {
		toSerialize["WorkflowDefinitionMoid"] = o.WorkflowDefinitionMoid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotificationTriggerWorkflow) UnmarshalJSON(data []byte) (err error) {
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
	type NotificationTriggerWorkflowWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Target workflow definition moid to trigger.
		WorkflowDefinitionMoid *string `json:"WorkflowDefinitionMoid,omitempty"`
	}

	varNotificationTriggerWorkflowWithoutEmbeddedStruct := NotificationTriggerWorkflowWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNotificationTriggerWorkflowWithoutEmbeddedStruct)
	if err == nil {
		varNotificationTriggerWorkflow := _NotificationTriggerWorkflow{}
		varNotificationTriggerWorkflow.ClassId = varNotificationTriggerWorkflowWithoutEmbeddedStruct.ClassId
		varNotificationTriggerWorkflow.ObjectType = varNotificationTriggerWorkflowWithoutEmbeddedStruct.ObjectType
		varNotificationTriggerWorkflow.WorkflowDefinitionMoid = varNotificationTriggerWorkflowWithoutEmbeddedStruct.WorkflowDefinitionMoid
		*o = NotificationTriggerWorkflow(varNotificationTriggerWorkflow)
	} else {
		return err
	}

	varNotificationTriggerWorkflow := _NotificationTriggerWorkflow{}

	err = json.Unmarshal(data, &varNotificationTriggerWorkflow)
	if err == nil {
		o.NotificationAction = varNotificationTriggerWorkflow.NotificationAction
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "WorkflowDefinitionMoid")

		// remove fields from embedded structs
		reflectNotificationAction := reflect.ValueOf(o.NotificationAction)
		for i := 0; i < reflectNotificationAction.Type().NumField(); i++ {
			t := reflectNotificationAction.Type().Field(i)

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

type NullableNotificationTriggerWorkflow struct {
	value *NotificationTriggerWorkflow
	isSet bool
}

func (v NullableNotificationTriggerWorkflow) Get() *NotificationTriggerWorkflow {
	return v.value
}

func (v *NullableNotificationTriggerWorkflow) Set(val *NotificationTriggerWorkflow) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTriggerWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTriggerWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTriggerWorkflow(val *NotificationTriggerWorkflow) *NullableNotificationTriggerWorkflow {
	return &NullableNotificationTriggerWorkflow{value: val, isSet: true}
}

func (v NullableNotificationTriggerWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTriggerWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
