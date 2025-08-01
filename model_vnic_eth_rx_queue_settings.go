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

// checks if the VnicEthRxQueueSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicEthRxQueueSettings{}

// VnicEthRxQueueSettings Receive Queue resource settings.
type VnicEthRxQueueSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of queue resources to allocate.
	Count *int64 `json:"Count,omitempty"`
	// The number of descriptors in each queue.
	RingSize             *int64 `json:"RingSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicEthRxQueueSettings VnicEthRxQueueSettings

// NewVnicEthRxQueueSettings instantiates a new VnicEthRxQueueSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthRxQueueSettings(classId string, objectType string) *VnicEthRxQueueSettings {
	this := VnicEthRxQueueSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var count int64 = 4
	this.Count = &count
	var ringSize int64 = 512
	this.RingSize = &ringSize
	return &this
}

// NewVnicEthRxQueueSettingsWithDefaults instantiates a new VnicEthRxQueueSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthRxQueueSettingsWithDefaults() *VnicEthRxQueueSettings {
	this := VnicEthRxQueueSettings{}
	var classId string = "vnic.EthRxQueueSettings"
	this.ClassId = classId
	var objectType string = "vnic.EthRxQueueSettings"
	this.ObjectType = objectType
	var count int64 = 4
	this.Count = &count
	var ringSize int64 = 512
	this.RingSize = &ringSize
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicEthRxQueueSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicEthRxQueueSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicEthRxQueueSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.EthRxQueueSettings" of the ClassId field.
func (o *VnicEthRxQueueSettings) GetDefaultClassId() interface{} {
	return "vnic.EthRxQueueSettings"
}

// GetObjectType returns the ObjectType field value
func (o *VnicEthRxQueueSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicEthRxQueueSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicEthRxQueueSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.EthRxQueueSettings" of the ObjectType field.
func (o *VnicEthRxQueueSettings) GetDefaultObjectType() interface{} {
	return "vnic.EthRxQueueSettings"
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *VnicEthRxQueueSettings) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthRxQueueSettings) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *VnicEthRxQueueSettings) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *VnicEthRxQueueSettings) SetCount(v int64) {
	o.Count = &v
}

// GetRingSize returns the RingSize field value if set, zero value otherwise.
func (o *VnicEthRxQueueSettings) GetRingSize() int64 {
	if o == nil || IsNil(o.RingSize) {
		var ret int64
		return ret
	}
	return *o.RingSize
}

// GetRingSizeOk returns a tuple with the RingSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthRxQueueSettings) GetRingSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.RingSize) {
		return nil, false
	}
	return o.RingSize, true
}

// HasRingSize returns a boolean if a field has been set.
func (o *VnicEthRxQueueSettings) HasRingSize() bool {
	if o != nil && !IsNil(o.RingSize) {
		return true
	}

	return false
}

// SetRingSize gets a reference to the given int64 and assigns it to the RingSize field.
func (o *VnicEthRxQueueSettings) SetRingSize(v int64) {
	o.RingSize = &v
}

func (o VnicEthRxQueueSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicEthRxQueueSettings) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Count) {
		toSerialize["Count"] = o.Count
	}
	if !IsNil(o.RingSize) {
		toSerialize["RingSize"] = o.RingSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicEthRxQueueSettings) UnmarshalJSON(data []byte) (err error) {
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
	type VnicEthRxQueueSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of queue resources to allocate.
		Count *int64 `json:"Count,omitempty"`
		// The number of descriptors in each queue.
		RingSize *int64 `json:"RingSize,omitempty"`
	}

	varVnicEthRxQueueSettingsWithoutEmbeddedStruct := VnicEthRxQueueSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicEthRxQueueSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicEthRxQueueSettings := _VnicEthRxQueueSettings{}
		varVnicEthRxQueueSettings.ClassId = varVnicEthRxQueueSettingsWithoutEmbeddedStruct.ClassId
		varVnicEthRxQueueSettings.ObjectType = varVnicEthRxQueueSettingsWithoutEmbeddedStruct.ObjectType
		varVnicEthRxQueueSettings.Count = varVnicEthRxQueueSettingsWithoutEmbeddedStruct.Count
		varVnicEthRxQueueSettings.RingSize = varVnicEthRxQueueSettingsWithoutEmbeddedStruct.RingSize
		*o = VnicEthRxQueueSettings(varVnicEthRxQueueSettings)
	} else {
		return err
	}

	varVnicEthRxQueueSettings := _VnicEthRxQueueSettings{}

	err = json.Unmarshal(data, &varVnicEthRxQueueSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicEthRxQueueSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Count")
		delete(additionalProperties, "RingSize")

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

type NullableVnicEthRxQueueSettings struct {
	value *VnicEthRxQueueSettings
	isSet bool
}

func (v NullableVnicEthRxQueueSettings) Get() *VnicEthRxQueueSettings {
	return v.value
}

func (v *NullableVnicEthRxQueueSettings) Set(val *VnicEthRxQueueSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthRxQueueSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthRxQueueSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthRxQueueSettings(val *VnicEthRxQueueSettings) *NullableVnicEthRxQueueSettings {
	return &NullableVnicEthRxQueueSettings{value: val, isSet: true}
}

func (v NullableVnicEthRxQueueSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthRxQueueSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
