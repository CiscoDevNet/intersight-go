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

// checks if the CondAlarmClassification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CondAlarmClassification{}

// CondAlarmClassification A group of related alarms that can be used in suppressions.
type CondAlarmClassification struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                               `json:"ObjectType"`
	AppliesTo  []CondAlarmClassificationEligibility `json:"AppliesTo,omitempty"`
	// A description about the alarms group that usually gives what kind of alarms are part of this classification.
	Description *string `json:"Description,omitempty"`
	// The name that uniquely identifies the alarm classificaton.
	Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.:-]{1,64}$"`
	// An array of relationships to condAlarmDefinition resources.
	AlarmDefinitions     []CondAlarmDefinitionRelationship `json:"AlarmDefinitions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CondAlarmClassification CondAlarmClassification

// NewCondAlarmClassification instantiates a new CondAlarmClassification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondAlarmClassification(classId string, objectType string) *CondAlarmClassification {
	this := CondAlarmClassification{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCondAlarmClassificationWithDefaults instantiates a new CondAlarmClassification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondAlarmClassificationWithDefaults() *CondAlarmClassification {
	this := CondAlarmClassification{}
	var classId string = "cond.AlarmClassification"
	this.ClassId = classId
	var objectType string = "cond.AlarmClassification"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CondAlarmClassification) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CondAlarmClassification) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CondAlarmClassification) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "cond.AlarmClassification" of the ClassId field.
func (o *CondAlarmClassification) GetDefaultClassId() interface{} {
	return "cond.AlarmClassification"
}

// GetObjectType returns the ObjectType field value
func (o *CondAlarmClassification) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CondAlarmClassification) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CondAlarmClassification) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "cond.AlarmClassification" of the ObjectType field.
func (o *CondAlarmClassification) GetDefaultObjectType() interface{} {
	return "cond.AlarmClassification"
}

// GetAppliesTo returns the AppliesTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CondAlarmClassification) GetAppliesTo() []CondAlarmClassificationEligibility {
	if o == nil {
		var ret []CondAlarmClassificationEligibility
		return ret
	}
	return o.AppliesTo
}

// GetAppliesToOk returns a tuple with the AppliesTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CondAlarmClassification) GetAppliesToOk() ([]CondAlarmClassificationEligibility, bool) {
	if o == nil || IsNil(o.AppliesTo) {
		return nil, false
	}
	return o.AppliesTo, true
}

// HasAppliesTo returns a boolean if a field has been set.
func (o *CondAlarmClassification) HasAppliesTo() bool {
	if o != nil && !IsNil(o.AppliesTo) {
		return true
	}

	return false
}

// SetAppliesTo gets a reference to the given []CondAlarmClassificationEligibility and assigns it to the AppliesTo field.
func (o *CondAlarmClassification) SetAppliesTo(v []CondAlarmClassificationEligibility) {
	o.AppliesTo = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CondAlarmClassification) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmClassification) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CondAlarmClassification) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CondAlarmClassification) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CondAlarmClassification) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmClassification) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CondAlarmClassification) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CondAlarmClassification) SetName(v string) {
	o.Name = &v
}

// GetAlarmDefinitions returns the AlarmDefinitions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CondAlarmClassification) GetAlarmDefinitions() []CondAlarmDefinitionRelationship {
	if o == nil {
		var ret []CondAlarmDefinitionRelationship
		return ret
	}
	return o.AlarmDefinitions
}

// GetAlarmDefinitionsOk returns a tuple with the AlarmDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CondAlarmClassification) GetAlarmDefinitionsOk() ([]CondAlarmDefinitionRelationship, bool) {
	if o == nil || IsNil(o.AlarmDefinitions) {
		return nil, false
	}
	return o.AlarmDefinitions, true
}

// HasAlarmDefinitions returns a boolean if a field has been set.
func (o *CondAlarmClassification) HasAlarmDefinitions() bool {
	if o != nil && !IsNil(o.AlarmDefinitions) {
		return true
	}

	return false
}

// SetAlarmDefinitions gets a reference to the given []CondAlarmDefinitionRelationship and assigns it to the AlarmDefinitions field.
func (o *CondAlarmClassification) SetAlarmDefinitions(v []CondAlarmDefinitionRelationship) {
	o.AlarmDefinitions = v
}

func (o CondAlarmClassification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CondAlarmClassification) ToMap() (map[string]interface{}, error) {
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
	if o.AppliesTo != nil {
		toSerialize["AppliesTo"] = o.AppliesTo
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.AlarmDefinitions != nil {
		toSerialize["AlarmDefinitions"] = o.AlarmDefinitions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CondAlarmClassification) UnmarshalJSON(data []byte) (err error) {
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
	type CondAlarmClassificationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                               `json:"ObjectType"`
		AppliesTo  []CondAlarmClassificationEligibility `json:"AppliesTo,omitempty"`
		// A description about the alarms group that usually gives what kind of alarms are part of this classification.
		Description *string `json:"Description,omitempty"`
		// The name that uniquely identifies the alarm classificaton.
		Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.:-]{1,64}$"`
		// An array of relationships to condAlarmDefinition resources.
		AlarmDefinitions []CondAlarmDefinitionRelationship `json:"AlarmDefinitions,omitempty"`
	}

	varCondAlarmClassificationWithoutEmbeddedStruct := CondAlarmClassificationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCondAlarmClassificationWithoutEmbeddedStruct)
	if err == nil {
		varCondAlarmClassification := _CondAlarmClassification{}
		varCondAlarmClassification.ClassId = varCondAlarmClassificationWithoutEmbeddedStruct.ClassId
		varCondAlarmClassification.ObjectType = varCondAlarmClassificationWithoutEmbeddedStruct.ObjectType
		varCondAlarmClassification.AppliesTo = varCondAlarmClassificationWithoutEmbeddedStruct.AppliesTo
		varCondAlarmClassification.Description = varCondAlarmClassificationWithoutEmbeddedStruct.Description
		varCondAlarmClassification.Name = varCondAlarmClassificationWithoutEmbeddedStruct.Name
		varCondAlarmClassification.AlarmDefinitions = varCondAlarmClassificationWithoutEmbeddedStruct.AlarmDefinitions
		*o = CondAlarmClassification(varCondAlarmClassification)
	} else {
		return err
	}

	varCondAlarmClassification := _CondAlarmClassification{}

	err = json.Unmarshal(data, &varCondAlarmClassification)
	if err == nil {
		o.MoBaseMo = varCondAlarmClassification.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AppliesTo")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "AlarmDefinitions")

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

type NullableCondAlarmClassification struct {
	value *CondAlarmClassification
	isSet bool
}

func (v NullableCondAlarmClassification) Get() *CondAlarmClassification {
	return v.value
}

func (v *NullableCondAlarmClassification) Set(val *CondAlarmClassification) {
	v.value = val
	v.isSet = true
}

func (v NullableCondAlarmClassification) IsSet() bool {
	return v.isSet
}

func (v *NullableCondAlarmClassification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondAlarmClassification(val *CondAlarmClassification) *NullableCondAlarmClassification {
	return &NullableCondAlarmClassification{value: val, isSet: true}
}

func (v NullableCondAlarmClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondAlarmClassification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
