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

// checks if the VnicFcNetworkPolicyInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicFcNetworkPolicyInventory{}

// VnicFcNetworkPolicyInventory A Fibre Channel Network policy governs the vSAN configuration for the virtual interfaces.
type VnicFcNetworkPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                       `json:"ObjectType"`
	VsanSettings         NullableVnicVsanSettings     `json:"VsanSettings,omitempty"`
	TargetMo             NullableMoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcNetworkPolicyInventory VnicFcNetworkPolicyInventory

// NewVnicFcNetworkPolicyInventory instantiates a new VnicFcNetworkPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcNetworkPolicyInventory(classId string, objectType string) *VnicFcNetworkPolicyInventory {
	this := VnicFcNetworkPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicFcNetworkPolicyInventoryWithDefaults instantiates a new VnicFcNetworkPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcNetworkPolicyInventoryWithDefaults() *VnicFcNetworkPolicyInventory {
	this := VnicFcNetworkPolicyInventory{}
	var classId string = "vnic.FcNetworkPolicyInventory"
	this.ClassId = classId
	var objectType string = "vnic.FcNetworkPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcNetworkPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcNetworkPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcNetworkPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.FcNetworkPolicyInventory" of the ClassId field.
func (o *VnicFcNetworkPolicyInventory) GetDefaultClassId() interface{} {
	return "vnic.FcNetworkPolicyInventory"
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcNetworkPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcNetworkPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcNetworkPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.FcNetworkPolicyInventory" of the ObjectType field.
func (o *VnicFcNetworkPolicyInventory) GetDefaultObjectType() interface{} {
	return "vnic.FcNetworkPolicyInventory"
}

// GetVsanSettings returns the VsanSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcNetworkPolicyInventory) GetVsanSettings() VnicVsanSettings {
	if o == nil || IsNil(o.VsanSettings.Get()) {
		var ret VnicVsanSettings
		return ret
	}
	return *o.VsanSettings.Get()
}

// GetVsanSettingsOk returns a tuple with the VsanSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcNetworkPolicyInventory) GetVsanSettingsOk() (*VnicVsanSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.VsanSettings.Get(), o.VsanSettings.IsSet()
}

// HasVsanSettings returns a boolean if a field has been set.
func (o *VnicFcNetworkPolicyInventory) HasVsanSettings() bool {
	if o != nil && o.VsanSettings.IsSet() {
		return true
	}

	return false
}

// SetVsanSettings gets a reference to the given NullableVnicVsanSettings and assigns it to the VsanSettings field.
func (o *VnicFcNetworkPolicyInventory) SetVsanSettings(v VnicVsanSettings) {
	o.VsanSettings.Set(&v)
}

// SetVsanSettingsNil sets the value for VsanSettings to be an explicit nil
func (o *VnicFcNetworkPolicyInventory) SetVsanSettingsNil() {
	o.VsanSettings.Set(nil)
}

// UnsetVsanSettings ensures that no value is present for VsanSettings, not even an explicit nil
func (o *VnicFcNetworkPolicyInventory) UnsetVsanSettings() {
	o.VsanSettings.Unset()
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcNetworkPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || IsNil(o.TargetMo.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo.Get()
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcNetworkPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetMo.Get(), o.TargetMo.IsSet()
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicFcNetworkPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo.IsSet() {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given NullableMoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicFcNetworkPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo.Set(&v)
}

// SetTargetMoNil sets the value for TargetMo to be an explicit nil
func (o *VnicFcNetworkPolicyInventory) SetTargetMoNil() {
	o.TargetMo.Set(nil)
}

// UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil
func (o *VnicFcNetworkPolicyInventory) UnsetTargetMo() {
	o.TargetMo.Unset()
}

func (o VnicFcNetworkPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicFcNetworkPolicyInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.VsanSettings.IsSet() {
		toSerialize["VsanSettings"] = o.VsanSettings.Get()
	}
	if o.TargetMo.IsSet() {
		toSerialize["TargetMo"] = o.TargetMo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicFcNetworkPolicyInventory) UnmarshalJSON(data []byte) (err error) {
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
	type VnicFcNetworkPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                       `json:"ObjectType"`
		VsanSettings NullableVnicVsanSettings     `json:"VsanSettings,omitempty"`
		TargetMo     NullableMoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varVnicFcNetworkPolicyInventoryWithoutEmbeddedStruct := VnicFcNetworkPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicFcNetworkPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varVnicFcNetworkPolicyInventory := _VnicFcNetworkPolicyInventory{}
		varVnicFcNetworkPolicyInventory.ClassId = varVnicFcNetworkPolicyInventoryWithoutEmbeddedStruct.ClassId
		varVnicFcNetworkPolicyInventory.ObjectType = varVnicFcNetworkPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varVnicFcNetworkPolicyInventory.VsanSettings = varVnicFcNetworkPolicyInventoryWithoutEmbeddedStruct.VsanSettings
		varVnicFcNetworkPolicyInventory.TargetMo = varVnicFcNetworkPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = VnicFcNetworkPolicyInventory(varVnicFcNetworkPolicyInventory)
	} else {
		return err
	}

	varVnicFcNetworkPolicyInventory := _VnicFcNetworkPolicyInventory{}

	err = json.Unmarshal(data, &varVnicFcNetworkPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varVnicFcNetworkPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "VsanSettings")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableVnicFcNetworkPolicyInventory struct {
	value *VnicFcNetworkPolicyInventory
	isSet bool
}

func (v NullableVnicFcNetworkPolicyInventory) Get() *VnicFcNetworkPolicyInventory {
	return v.value
}

func (v *NullableVnicFcNetworkPolicyInventory) Set(val *VnicFcNetworkPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcNetworkPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcNetworkPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcNetworkPolicyInventory(val *VnicFcNetworkPolicyInventory) *NullableVnicFcNetworkPolicyInventory {
	return &NullableVnicFcNetworkPolicyInventory{value: val, isSet: true}
}

func (v NullableVnicFcNetworkPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcNetworkPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
