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

// checks if the VnicFcInterruptSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicFcInterruptSettings{}

// VnicFcInterruptSettings Interrupt Settings for the virtual fibre channel interface.
type VnicFcInterruptSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option. * `MSIx` - Message Signaled Interrupt (MSI) mechanism with the optional extension (MSIx). MSIx is the recommended and default option. * `MSI` - Message Signaled Interrupt (MSI) mechanism that treats messages as interrupts. * `INTx` - Line-based interrupt (INTx) mechanism similar to the one used in Legacy systems.
	Mode                 *string `json:"Mode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcInterruptSettings VnicFcInterruptSettings

// NewVnicFcInterruptSettings instantiates a new VnicFcInterruptSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcInterruptSettings(classId string, objectType string) *VnicFcInterruptSettings {
	this := VnicFcInterruptSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var mode string = "MSIx"
	this.Mode = &mode
	return &this
}

// NewVnicFcInterruptSettingsWithDefaults instantiates a new VnicFcInterruptSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcInterruptSettingsWithDefaults() *VnicFcInterruptSettings {
	this := VnicFcInterruptSettings{}
	var classId string = "vnic.FcInterruptSettings"
	this.ClassId = classId
	var objectType string = "vnic.FcInterruptSettings"
	this.ObjectType = objectType
	var mode string = "MSIx"
	this.Mode = &mode
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcInterruptSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcInterruptSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcInterruptSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.FcInterruptSettings" of the ClassId field.
func (o *VnicFcInterruptSettings) GetDefaultClassId() interface{} {
	return "vnic.FcInterruptSettings"
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcInterruptSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcInterruptSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcInterruptSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.FcInterruptSettings" of the ObjectType field.
func (o *VnicFcInterruptSettings) GetDefaultObjectType() interface{} {
	return "vnic.FcInterruptSettings"
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VnicFcInterruptSettings) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcInterruptSettings) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VnicFcInterruptSettings) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *VnicFcInterruptSettings) SetMode(v string) {
	o.Mode = &v
}

func (o VnicFcInterruptSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicFcInterruptSettings) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Mode) {
		toSerialize["Mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicFcInterruptSettings) UnmarshalJSON(data []byte) (err error) {
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
	type VnicFcInterruptSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option. * `MSIx` - Message Signaled Interrupt (MSI) mechanism with the optional extension (MSIx). MSIx is the recommended and default option. * `MSI` - Message Signaled Interrupt (MSI) mechanism that treats messages as interrupts. * `INTx` - Line-based interrupt (INTx) mechanism similar to the one used in Legacy systems.
		Mode *string `json:"Mode,omitempty"`
	}

	varVnicFcInterruptSettingsWithoutEmbeddedStruct := VnicFcInterruptSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicFcInterruptSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicFcInterruptSettings := _VnicFcInterruptSettings{}
		varVnicFcInterruptSettings.ClassId = varVnicFcInterruptSettingsWithoutEmbeddedStruct.ClassId
		varVnicFcInterruptSettings.ObjectType = varVnicFcInterruptSettingsWithoutEmbeddedStruct.ObjectType
		varVnicFcInterruptSettings.Mode = varVnicFcInterruptSettingsWithoutEmbeddedStruct.Mode
		*o = VnicFcInterruptSettings(varVnicFcInterruptSettings)
	} else {
		return err
	}

	varVnicFcInterruptSettings := _VnicFcInterruptSettings{}

	err = json.Unmarshal(data, &varVnicFcInterruptSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicFcInterruptSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Mode")

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

type NullableVnicFcInterruptSettings struct {
	value *VnicFcInterruptSettings
	isSet bool
}

func (v NullableVnicFcInterruptSettings) Get() *VnicFcInterruptSettings {
	return v.value
}

func (v *NullableVnicFcInterruptSettings) Set(val *VnicFcInterruptSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcInterruptSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcInterruptSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcInterruptSettings(val *VnicFcInterruptSettings) *NullableVnicFcInterruptSettings {
	return &NullableVnicFcInterruptSettings{value: val, isSet: true}
}

func (v NullableVnicFcInterruptSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcInterruptSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
