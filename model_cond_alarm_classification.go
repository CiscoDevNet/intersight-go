/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14828
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

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
	Name *string `json:"Name,omitempty"`
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
	if o == nil || o.AppliesTo == nil {
		return nil, false
	}
	return o.AppliesTo, true
}

// HasAppliesTo returns a boolean if a field has been set.
func (o *CondAlarmClassification) HasAppliesTo() bool {
	if o != nil && o.AppliesTo != nil {
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
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmClassification) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CondAlarmClassification) HasDescription() bool {
	if o != nil && o.Description != nil {
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
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmClassification) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CondAlarmClassification) HasName() bool {
	if o != nil && o.Name != nil {
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
	if o == nil || o.AlarmDefinitions == nil {
		return nil, false
	}
	return o.AlarmDefinitions, true
}

// HasAlarmDefinitions returns a boolean if a field has been set.
func (o *CondAlarmClassification) HasAlarmDefinitions() bool {
	if o != nil && o.AlarmDefinitions != nil {
		return true
	}

	return false
}

// SetAlarmDefinitions gets a reference to the given []CondAlarmDefinitionRelationship and assigns it to the AlarmDefinitions field.
func (o *CondAlarmClassification) SetAlarmDefinitions(v []CondAlarmDefinitionRelationship) {
	o.AlarmDefinitions = v
}

func (o CondAlarmClassification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AppliesTo != nil {
		toSerialize["AppliesTo"] = o.AppliesTo
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.AlarmDefinitions != nil {
		toSerialize["AlarmDefinitions"] = o.AlarmDefinitions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CondAlarmClassification) UnmarshalJSON(bytes []byte) (err error) {
	type CondAlarmClassificationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                               `json:"ObjectType"`
		AppliesTo  []CondAlarmClassificationEligibility `json:"AppliesTo,omitempty"`
		// A description about the alarms group that usually gives what kind of alarms are part of this classification.
		Description *string `json:"Description,omitempty"`
		// The name that uniquely identifies the alarm classificaton.
		Name *string `json:"Name,omitempty"`
		// An array of relationships to condAlarmDefinition resources.
		AlarmDefinitions []CondAlarmDefinitionRelationship `json:"AlarmDefinitions,omitempty"`
	}

	varCondAlarmClassificationWithoutEmbeddedStruct := CondAlarmClassificationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCondAlarmClassificationWithoutEmbeddedStruct)
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

	err = json.Unmarshal(bytes, &varCondAlarmClassification)
	if err == nil {
		o.MoBaseMo = varCondAlarmClassification.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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