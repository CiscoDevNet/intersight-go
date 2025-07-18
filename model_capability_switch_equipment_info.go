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

// checks if the CapabilitySwitchEquipmentInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilitySwitchEquipmentInfo{}

// CapabilitySwitchEquipmentInfo Type to represent additional switch equipment information.
type CapabilitySwitchEquipmentInfo struct {
	CapabilitySwitchIdentityDef
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Maximum number of fan modules/trays present on this switch.
	MaxFanModules *int64 `json:"MaxFanModules,omitempty"`
	// Maximum number of fans present in a fan module/tray on this switch.
	MaxFansInFanModule   *int64 `json:"MaxFansInFanModule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySwitchEquipmentInfo CapabilitySwitchEquipmentInfo

// NewCapabilitySwitchEquipmentInfo instantiates a new CapabilitySwitchEquipmentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchEquipmentInfo(classId string, objectType string) *CapabilitySwitchEquipmentInfo {
	this := CapabilitySwitchEquipmentInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilitySwitchEquipmentInfoWithDefaults instantiates a new CapabilitySwitchEquipmentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchEquipmentInfoWithDefaults() *CapabilitySwitchEquipmentInfo {
	this := CapabilitySwitchEquipmentInfo{}
	var classId string = "capability.SwitchEquipmentInfo"
	this.ClassId = classId
	var objectType string = "capability.SwitchEquipmentInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySwitchEquipmentInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchEquipmentInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySwitchEquipmentInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.SwitchEquipmentInfo" of the ClassId field.
func (o *CapabilitySwitchEquipmentInfo) GetDefaultClassId() interface{} {
	return "capability.SwitchEquipmentInfo"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySwitchEquipmentInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchEquipmentInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySwitchEquipmentInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.SwitchEquipmentInfo" of the ObjectType field.
func (o *CapabilitySwitchEquipmentInfo) GetDefaultObjectType() interface{} {
	return "capability.SwitchEquipmentInfo"
}

// GetMaxFanModules returns the MaxFanModules field value if set, zero value otherwise.
func (o *CapabilitySwitchEquipmentInfo) GetMaxFanModules() int64 {
	if o == nil || IsNil(o.MaxFanModules) {
		var ret int64
		return ret
	}
	return *o.MaxFanModules
}

// GetMaxFanModulesOk returns a tuple with the MaxFanModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchEquipmentInfo) GetMaxFanModulesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxFanModules) {
		return nil, false
	}
	return o.MaxFanModules, true
}

// HasMaxFanModules returns a boolean if a field has been set.
func (o *CapabilitySwitchEquipmentInfo) HasMaxFanModules() bool {
	if o != nil && !IsNil(o.MaxFanModules) {
		return true
	}

	return false
}

// SetMaxFanModules gets a reference to the given int64 and assigns it to the MaxFanModules field.
func (o *CapabilitySwitchEquipmentInfo) SetMaxFanModules(v int64) {
	o.MaxFanModules = &v
}

// GetMaxFansInFanModule returns the MaxFansInFanModule field value if set, zero value otherwise.
func (o *CapabilitySwitchEquipmentInfo) GetMaxFansInFanModule() int64 {
	if o == nil || IsNil(o.MaxFansInFanModule) {
		var ret int64
		return ret
	}
	return *o.MaxFansInFanModule
}

// GetMaxFansInFanModuleOk returns a tuple with the MaxFansInFanModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchEquipmentInfo) GetMaxFansInFanModuleOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxFansInFanModule) {
		return nil, false
	}
	return o.MaxFansInFanModule, true
}

// HasMaxFansInFanModule returns a boolean if a field has been set.
func (o *CapabilitySwitchEquipmentInfo) HasMaxFansInFanModule() bool {
	if o != nil && !IsNil(o.MaxFansInFanModule) {
		return true
	}

	return false
}

// SetMaxFansInFanModule gets a reference to the given int64 and assigns it to the MaxFansInFanModule field.
func (o *CapabilitySwitchEquipmentInfo) SetMaxFansInFanModule(v int64) {
	o.MaxFansInFanModule = &v
}

func (o CapabilitySwitchEquipmentInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilitySwitchEquipmentInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilitySwitchIdentityDef, errCapabilitySwitchIdentityDef := json.Marshal(o.CapabilitySwitchIdentityDef)
	if errCapabilitySwitchIdentityDef != nil {
		return map[string]interface{}{}, errCapabilitySwitchIdentityDef
	}
	errCapabilitySwitchIdentityDef = json.Unmarshal([]byte(serializedCapabilitySwitchIdentityDef), &toSerialize)
	if errCapabilitySwitchIdentityDef != nil {
		return map[string]interface{}{}, errCapabilitySwitchIdentityDef
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.MaxFanModules) {
		toSerialize["MaxFanModules"] = o.MaxFanModules
	}
	if !IsNil(o.MaxFansInFanModule) {
		toSerialize["MaxFansInFanModule"] = o.MaxFansInFanModule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilitySwitchEquipmentInfo) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilitySwitchEquipmentInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Maximum number of fan modules/trays present on this switch.
		MaxFanModules *int64 `json:"MaxFanModules,omitempty"`
		// Maximum number of fans present in a fan module/tray on this switch.
		MaxFansInFanModule *int64 `json:"MaxFansInFanModule,omitempty"`
	}

	varCapabilitySwitchEquipmentInfoWithoutEmbeddedStruct := CapabilitySwitchEquipmentInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilitySwitchEquipmentInfoWithoutEmbeddedStruct)
	if err == nil {
		varCapabilitySwitchEquipmentInfo := _CapabilitySwitchEquipmentInfo{}
		varCapabilitySwitchEquipmentInfo.ClassId = varCapabilitySwitchEquipmentInfoWithoutEmbeddedStruct.ClassId
		varCapabilitySwitchEquipmentInfo.ObjectType = varCapabilitySwitchEquipmentInfoWithoutEmbeddedStruct.ObjectType
		varCapabilitySwitchEquipmentInfo.MaxFanModules = varCapabilitySwitchEquipmentInfoWithoutEmbeddedStruct.MaxFanModules
		varCapabilitySwitchEquipmentInfo.MaxFansInFanModule = varCapabilitySwitchEquipmentInfoWithoutEmbeddedStruct.MaxFansInFanModule
		*o = CapabilitySwitchEquipmentInfo(varCapabilitySwitchEquipmentInfo)
	} else {
		return err
	}

	varCapabilitySwitchEquipmentInfo := _CapabilitySwitchEquipmentInfo{}

	err = json.Unmarshal(data, &varCapabilitySwitchEquipmentInfo)
	if err == nil {
		o.CapabilitySwitchIdentityDef = varCapabilitySwitchEquipmentInfo.CapabilitySwitchIdentityDef
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MaxFanModules")
		delete(additionalProperties, "MaxFansInFanModule")

		// remove fields from embedded structs
		reflectCapabilitySwitchIdentityDef := reflect.ValueOf(o.CapabilitySwitchIdentityDef)
		for i := 0; i < reflectCapabilitySwitchIdentityDef.Type().NumField(); i++ {
			t := reflectCapabilitySwitchIdentityDef.Type().Field(i)

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

type NullableCapabilitySwitchEquipmentInfo struct {
	value *CapabilitySwitchEquipmentInfo
	isSet bool
}

func (v NullableCapabilitySwitchEquipmentInfo) Get() *CapabilitySwitchEquipmentInfo {
	return v.value
}

func (v *NullableCapabilitySwitchEquipmentInfo) Set(val *CapabilitySwitchEquipmentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchEquipmentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchEquipmentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchEquipmentInfo(val *CapabilitySwitchEquipmentInfo) *NullableCapabilitySwitchEquipmentInfo {
	return &NullableCapabilitySwitchEquipmentInfo{value: val, isSet: true}
}

func (v NullableCapabilitySwitchEquipmentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchEquipmentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
