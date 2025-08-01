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

// checks if the WorkflowMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowMessage{}

// WorkflowMessage Intermediate Task or Workflow message with severity.
type WorkflowMessage struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An i18n message that can be translated into multiple languages to support internationalization.
	Message *string `json:"Message,omitempty"`
	// The severity of the Task or Workflow message warning/error/info etc. * `Info` - The enum represents the log level to be used to convey info message. * `Warning` - The enum represents the log level to be used to convey warning message. * `Debug` - The enum represents the log level to be used to convey debug message. * `Error` - The enum represents the log level to be used to convey error message.
	Severity             *string `json:"Severity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowMessage WorkflowMessage

// NewWorkflowMessage instantiates a new WorkflowMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowMessage(classId string, objectType string) *WorkflowMessage {
	this := WorkflowMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowMessageWithDefaults instantiates a new WorkflowMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowMessageWithDefaults() *WorkflowMessage {
	this := WorkflowMessage{}
	var classId string = "workflow.Message"
	this.ClassId = classId
	var objectType string = "workflow.Message"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.Message" of the ClassId field.
func (o *WorkflowMessage) GetDefaultClassId() interface{} {
	return "workflow.Message"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.Message" of the ObjectType field.
func (o *WorkflowMessage) GetDefaultObjectType() interface{} {
	return "workflow.Message"
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *WorkflowMessage) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowMessage) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *WorkflowMessage) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *WorkflowMessage) SetMessage(v string) {
	o.Message = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *WorkflowMessage) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowMessage) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *WorkflowMessage) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *WorkflowMessage) SetSeverity(v string) {
	o.Severity = &v
}

func (o WorkflowMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowMessage) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if !IsNil(o.Severity) {
		toSerialize["Severity"] = o.Severity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowMessage) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// An i18n message that can be translated into multiple languages to support internationalization.
		Message *string `json:"Message,omitempty"`
		// The severity of the Task or Workflow message warning/error/info etc. * `Info` - The enum represents the log level to be used to convey info message. * `Warning` - The enum represents the log level to be used to convey warning message. * `Debug` - The enum represents the log level to be used to convey debug message. * `Error` - The enum represents the log level to be used to convey error message.
		Severity *string `json:"Severity,omitempty"`
	}

	varWorkflowMessageWithoutEmbeddedStruct := WorkflowMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowMessageWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowMessage := _WorkflowMessage{}
		varWorkflowMessage.ClassId = varWorkflowMessageWithoutEmbeddedStruct.ClassId
		varWorkflowMessage.ObjectType = varWorkflowMessageWithoutEmbeddedStruct.ObjectType
		varWorkflowMessage.Message = varWorkflowMessageWithoutEmbeddedStruct.Message
		varWorkflowMessage.Severity = varWorkflowMessageWithoutEmbeddedStruct.Severity
		*o = WorkflowMessage(varWorkflowMessage)
	} else {
		return err
	}

	varWorkflowMessage := _WorkflowMessage{}

	err = json.Unmarshal(data, &varWorkflowMessage)
	if err == nil {
		o.MoBaseComplexType = varWorkflowMessage.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Severity")

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

type NullableWorkflowMessage struct {
	value *WorkflowMessage
	isSet bool
}

func (v NullableWorkflowMessage) Get() *WorkflowMessage {
	return v.value
}

func (v *NullableWorkflowMessage) Set(val *WorkflowMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowMessage(val *WorkflowMessage) *NullableWorkflowMessage {
	return &NullableWorkflowMessage{value: val, isSet: true}
}

func (v NullableWorkflowMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
