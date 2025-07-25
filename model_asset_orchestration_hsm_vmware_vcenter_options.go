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

// checks if the AssetOrchestrationHsmVmwareVcenterOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetOrchestrationHsmVmwareVcenterOptions{}

// AssetOrchestrationHsmVmwareVcenterOptions Captures configuration specific to the VMware Vcenter target for the Hardware Support Manager.
type AssetOrchestrationHsmVmwareVcenterOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// HsmEnabled controls whether Hardware Support Manager is enabled or not. vCenter Server version 7.0 or later.
	HsmEnabled *bool `json:"HsmEnabled,omitempty"`
	// Indicates whether the value of the 'clientSecret' property has been set.
	IsClientSecretSet    *bool `json:"IsClientSecretSet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetOrchestrationHsmVmwareVcenterOptions AssetOrchestrationHsmVmwareVcenterOptions

// NewAssetOrchestrationHsmVmwareVcenterOptions instantiates a new AssetOrchestrationHsmVmwareVcenterOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetOrchestrationHsmVmwareVcenterOptions(classId string, objectType string) *AssetOrchestrationHsmVmwareVcenterOptions {
	this := AssetOrchestrationHsmVmwareVcenterOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetOrchestrationHsmVmwareVcenterOptionsWithDefaults instantiates a new AssetOrchestrationHsmVmwareVcenterOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetOrchestrationHsmVmwareVcenterOptionsWithDefaults() *AssetOrchestrationHsmVmwareVcenterOptions {
	this := AssetOrchestrationHsmVmwareVcenterOptions{}
	var classId string = "asset.OrchestrationHsmVmwareVcenterOptions"
	this.ClassId = classId
	var objectType string = "asset.OrchestrationHsmVmwareVcenterOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetOrchestrationHsmVmwareVcenterOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "asset.OrchestrationHsmVmwareVcenterOptions" of the ClassId field.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetDefaultClassId() interface{} {
	return "asset.OrchestrationHsmVmwareVcenterOptions"
}

// GetObjectType returns the ObjectType field value
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetOrchestrationHsmVmwareVcenterOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "asset.OrchestrationHsmVmwareVcenterOptions" of the ObjectType field.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetDefaultObjectType() interface{} {
	return "asset.OrchestrationHsmVmwareVcenterOptions"
}

// GetHsmEnabled returns the HsmEnabled field value if set, zero value otherwise.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetHsmEnabled() bool {
	if o == nil || IsNil(o.HsmEnabled) {
		var ret bool
		return ret
	}
	return *o.HsmEnabled
}

// GetHsmEnabledOk returns a tuple with the HsmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetHsmEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.HsmEnabled) {
		return nil, false
	}
	return o.HsmEnabled, true
}

// HasHsmEnabled returns a boolean if a field has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) HasHsmEnabled() bool {
	if o != nil && !IsNil(o.HsmEnabled) {
		return true
	}

	return false
}

// SetHsmEnabled gets a reference to the given bool and assigns it to the HsmEnabled field.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) SetHsmEnabled(v bool) {
	o.HsmEnabled = &v
}

// GetIsClientSecretSet returns the IsClientSecretSet field value if set, zero value otherwise.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetIsClientSecretSet() bool {
	if o == nil || IsNil(o.IsClientSecretSet) {
		var ret bool
		return ret
	}
	return *o.IsClientSecretSet
}

// GetIsClientSecretSetOk returns a tuple with the IsClientSecretSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetIsClientSecretSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsClientSecretSet) {
		return nil, false
	}
	return o.IsClientSecretSet, true
}

// HasIsClientSecretSet returns a boolean if a field has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) HasIsClientSecretSet() bool {
	if o != nil && !IsNil(o.IsClientSecretSet) {
		return true
	}

	return false
}

// SetIsClientSecretSet gets a reference to the given bool and assigns it to the IsClientSecretSet field.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) SetIsClientSecretSet(v bool) {
	o.IsClientSecretSet = &v
}

func (o AssetOrchestrationHsmVmwareVcenterOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetOrchestrationHsmVmwareVcenterOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return map[string]interface{}{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return map[string]interface{}{}, errAssetServiceOptions
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.HsmEnabled) {
		toSerialize["HsmEnabled"] = o.HsmEnabled
	}
	if !IsNil(o.IsClientSecretSet) {
		toSerialize["IsClientSecretSet"] = o.IsClientSecretSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetOrchestrationHsmVmwareVcenterOptions) UnmarshalJSON(data []byte) (err error) {
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
	type AssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// HsmEnabled controls whether Hardware Support Manager is enabled or not. vCenter Server version 7.0 or later.
		HsmEnabled *bool `json:"HsmEnabled,omitempty"`
		// Indicates whether the value of the 'clientSecret' property has been set.
		IsClientSecretSet *bool `json:"IsClientSecretSet,omitempty"`
	}

	varAssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct := AssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetOrchestrationHsmVmwareVcenterOptions := _AssetOrchestrationHsmVmwareVcenterOptions{}
		varAssetOrchestrationHsmVmwareVcenterOptions.ClassId = varAssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct.ClassId
		varAssetOrchestrationHsmVmwareVcenterOptions.ObjectType = varAssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct.ObjectType
		varAssetOrchestrationHsmVmwareVcenterOptions.HsmEnabled = varAssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct.HsmEnabled
		varAssetOrchestrationHsmVmwareVcenterOptions.IsClientSecretSet = varAssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct.IsClientSecretSet
		*o = AssetOrchestrationHsmVmwareVcenterOptions(varAssetOrchestrationHsmVmwareVcenterOptions)
	} else {
		return err
	}

	varAssetOrchestrationHsmVmwareVcenterOptions := _AssetOrchestrationHsmVmwareVcenterOptions{}

	err = json.Unmarshal(data, &varAssetOrchestrationHsmVmwareVcenterOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetOrchestrationHsmVmwareVcenterOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HsmEnabled")
		delete(additionalProperties, "IsClientSecretSet")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

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

type NullableAssetOrchestrationHsmVmwareVcenterOptions struct {
	value *AssetOrchestrationHsmVmwareVcenterOptions
	isSet bool
}

func (v NullableAssetOrchestrationHsmVmwareVcenterOptions) Get() *AssetOrchestrationHsmVmwareVcenterOptions {
	return v.value
}

func (v *NullableAssetOrchestrationHsmVmwareVcenterOptions) Set(val *AssetOrchestrationHsmVmwareVcenterOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOrchestrationHsmVmwareVcenterOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOrchestrationHsmVmwareVcenterOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOrchestrationHsmVmwareVcenterOptions(val *AssetOrchestrationHsmVmwareVcenterOptions) *NullableAssetOrchestrationHsmVmwareVcenterOptions {
	return &NullableAssetOrchestrationHsmVmwareVcenterOptions{value: val, isSet: true}
}

func (v NullableAssetOrchestrationHsmVmwareVcenterOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOrchestrationHsmVmwareVcenterOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
