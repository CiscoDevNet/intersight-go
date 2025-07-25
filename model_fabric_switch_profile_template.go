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

// checks if the FabricSwitchProfileTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricSwitchProfileTemplate{}

// FabricSwitchProfileTemplate The template specifies the policy configuration for each of the managed network switch.
type FabricSwitchProfileTemplate struct {
	FabricBaseSwitchProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType                   string                                                 `json:"ObjectType"`
	SwitchClusterProfileTemplate NullableFabricSwitchClusterProfileTemplateRelationship `json:"SwitchClusterProfileTemplate,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _FabricSwitchProfileTemplate FabricSwitchProfileTemplate

// NewFabricSwitchProfileTemplate instantiates a new FabricSwitchProfileTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSwitchProfileTemplate(classId string, objectType string) *FabricSwitchProfileTemplate {
	this := FabricSwitchProfileTemplate{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	var action string = "No-op"
	this.Action = &action
	var switchId string = "None"
	this.SwitchId = &switchId
	return &this
}

// NewFabricSwitchProfileTemplateWithDefaults instantiates a new FabricSwitchProfileTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSwitchProfileTemplateWithDefaults() *FabricSwitchProfileTemplate {
	this := FabricSwitchProfileTemplate{}
	var classId string = "fabric.SwitchProfileTemplate"
	this.ClassId = classId
	var objectType string = "fabric.SwitchProfileTemplate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSwitchProfileTemplate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfileTemplate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSwitchProfileTemplate) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.SwitchProfileTemplate" of the ClassId field.
func (o *FabricSwitchProfileTemplate) GetDefaultClassId() interface{} {
	return "fabric.SwitchProfileTemplate"
}

// GetObjectType returns the ObjectType field value
func (o *FabricSwitchProfileTemplate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfileTemplate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSwitchProfileTemplate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.SwitchProfileTemplate" of the ObjectType field.
func (o *FabricSwitchProfileTemplate) GetDefaultObjectType() interface{} {
	return "fabric.SwitchProfileTemplate"
}

// GetSwitchClusterProfileTemplate returns the SwitchClusterProfileTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchProfileTemplate) GetSwitchClusterProfileTemplate() FabricSwitchClusterProfileTemplateRelationship {
	if o == nil || IsNil(o.SwitchClusterProfileTemplate.Get()) {
		var ret FabricSwitchClusterProfileTemplateRelationship
		return ret
	}
	return *o.SwitchClusterProfileTemplate.Get()
}

// GetSwitchClusterProfileTemplateOk returns a tuple with the SwitchClusterProfileTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchProfileTemplate) GetSwitchClusterProfileTemplateOk() (*FabricSwitchClusterProfileTemplateRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SwitchClusterProfileTemplate.Get(), o.SwitchClusterProfileTemplate.IsSet()
}

// HasSwitchClusterProfileTemplate returns a boolean if a field has been set.
func (o *FabricSwitchProfileTemplate) HasSwitchClusterProfileTemplate() bool {
	if o != nil && o.SwitchClusterProfileTemplate.IsSet() {
		return true
	}

	return false
}

// SetSwitchClusterProfileTemplate gets a reference to the given NullableFabricSwitchClusterProfileTemplateRelationship and assigns it to the SwitchClusterProfileTemplate field.
func (o *FabricSwitchProfileTemplate) SetSwitchClusterProfileTemplate(v FabricSwitchClusterProfileTemplateRelationship) {
	o.SwitchClusterProfileTemplate.Set(&v)
}

// SetSwitchClusterProfileTemplateNil sets the value for SwitchClusterProfileTemplate to be an explicit nil
func (o *FabricSwitchProfileTemplate) SetSwitchClusterProfileTemplateNil() {
	o.SwitchClusterProfileTemplate.Set(nil)
}

// UnsetSwitchClusterProfileTemplate ensures that no value is present for SwitchClusterProfileTemplate, not even an explicit nil
func (o *FabricSwitchProfileTemplate) UnsetSwitchClusterProfileTemplate() {
	o.SwitchClusterProfileTemplate.Unset()
}

func (o FabricSwitchProfileTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricSwitchProfileTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricBaseSwitchProfile, errFabricBaseSwitchProfile := json.Marshal(o.FabricBaseSwitchProfile)
	if errFabricBaseSwitchProfile != nil {
		return map[string]interface{}{}, errFabricBaseSwitchProfile
	}
	errFabricBaseSwitchProfile = json.Unmarshal([]byte(serializedFabricBaseSwitchProfile), &toSerialize)
	if errFabricBaseSwitchProfile != nil {
		return map[string]interface{}{}, errFabricBaseSwitchProfile
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.SwitchClusterProfileTemplate.IsSet() {
		toSerialize["SwitchClusterProfileTemplate"] = o.SwitchClusterProfileTemplate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricSwitchProfileTemplate) UnmarshalJSON(data []byte) (err error) {
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
	type FabricSwitchProfileTemplateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType                   string                                                 `json:"ObjectType"`
		SwitchClusterProfileTemplate NullableFabricSwitchClusterProfileTemplateRelationship `json:"SwitchClusterProfileTemplate,omitempty"`
	}

	varFabricSwitchProfileTemplateWithoutEmbeddedStruct := FabricSwitchProfileTemplateWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricSwitchProfileTemplateWithoutEmbeddedStruct)
	if err == nil {
		varFabricSwitchProfileTemplate := _FabricSwitchProfileTemplate{}
		varFabricSwitchProfileTemplate.ClassId = varFabricSwitchProfileTemplateWithoutEmbeddedStruct.ClassId
		varFabricSwitchProfileTemplate.ObjectType = varFabricSwitchProfileTemplateWithoutEmbeddedStruct.ObjectType
		varFabricSwitchProfileTemplate.SwitchClusterProfileTemplate = varFabricSwitchProfileTemplateWithoutEmbeddedStruct.SwitchClusterProfileTemplate
		*o = FabricSwitchProfileTemplate(varFabricSwitchProfileTemplate)
	} else {
		return err
	}

	varFabricSwitchProfileTemplate := _FabricSwitchProfileTemplate{}

	err = json.Unmarshal(data, &varFabricSwitchProfileTemplate)
	if err == nil {
		o.FabricBaseSwitchProfile = varFabricSwitchProfileTemplate.FabricBaseSwitchProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SwitchClusterProfileTemplate")

		// remove fields from embedded structs
		reflectFabricBaseSwitchProfile := reflect.ValueOf(o.FabricBaseSwitchProfile)
		for i := 0; i < reflectFabricBaseSwitchProfile.Type().NumField(); i++ {
			t := reflectFabricBaseSwitchProfile.Type().Field(i)

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

type NullableFabricSwitchProfileTemplate struct {
	value *FabricSwitchProfileTemplate
	isSet bool
}

func (v NullableFabricSwitchProfileTemplate) Get() *FabricSwitchProfileTemplate {
	return v.value
}

func (v *NullableFabricSwitchProfileTemplate) Set(val *FabricSwitchProfileTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSwitchProfileTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSwitchProfileTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSwitchProfileTemplate(val *FabricSwitchProfileTemplate) *NullableFabricSwitchProfileTemplate {
	return &NullableFabricSwitchProfileTemplate{value: val, isSet: true}
}

func (v NullableFabricSwitchProfileTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSwitchProfileTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
