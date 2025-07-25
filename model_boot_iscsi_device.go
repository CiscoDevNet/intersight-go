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

// checks if the BootIscsiDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BootIscsiDevice{}

// BootIscsiDevice Iscsi Boot Device configured on the server.
type BootIscsiDevice struct {
	BootConfiguredDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                      `json:"ObjectType"`
	ComputePhysical      NullableComputePhysicalRelationship         `json:"ComputePhysical,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootIscsiDevice BootIscsiDevice

// NewBootIscsiDevice instantiates a new BootIscsiDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootIscsiDevice(classId string, objectType string) *BootIscsiDevice {
	this := BootIscsiDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBootIscsiDeviceWithDefaults instantiates a new BootIscsiDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootIscsiDeviceWithDefaults() *BootIscsiDevice {
	this := BootIscsiDevice{}
	var classId string = "boot.IscsiDevice"
	this.ClassId = classId
	var objectType string = "boot.IscsiDevice"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootIscsiDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootIscsiDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootIscsiDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "boot.IscsiDevice" of the ClassId field.
func (o *BootIscsiDevice) GetDefaultClassId() interface{} {
	return "boot.IscsiDevice"
}

// GetObjectType returns the ObjectType field value
func (o *BootIscsiDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootIscsiDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootIscsiDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "boot.IscsiDevice" of the ObjectType field.
func (o *BootIscsiDevice) GetDefaultObjectType() interface{} {
	return "boot.IscsiDevice"
}

// GetComputePhysical returns the ComputePhysical field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootIscsiDevice) GetComputePhysical() ComputePhysicalRelationship {
	if o == nil || IsNil(o.ComputePhysical.Get()) {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.ComputePhysical.Get()
}

// GetComputePhysicalOk returns a tuple with the ComputePhysical field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootIscsiDevice) GetComputePhysicalOk() (*ComputePhysicalRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputePhysical.Get(), o.ComputePhysical.IsSet()
}

// HasComputePhysical returns a boolean if a field has been set.
func (o *BootIscsiDevice) HasComputePhysical() bool {
	if o != nil && o.ComputePhysical.IsSet() {
		return true
	}

	return false
}

// SetComputePhysical gets a reference to the given NullableComputePhysicalRelationship and assigns it to the ComputePhysical field.
func (o *BootIscsiDevice) SetComputePhysical(v ComputePhysicalRelationship) {
	o.ComputePhysical.Set(&v)
}

// SetComputePhysicalNil sets the value for ComputePhysical to be an explicit nil
func (o *BootIscsiDevice) SetComputePhysicalNil() {
	o.ComputePhysical.Set(nil)
}

// UnsetComputePhysical ensures that no value is present for ComputePhysical, not even an explicit nil
func (o *BootIscsiDevice) UnsetComputePhysical() {
	o.ComputePhysical.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootIscsiDevice) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootIscsiDevice) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BootIscsiDevice) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BootIscsiDevice) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *BootIscsiDevice) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *BootIscsiDevice) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootIscsiDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootIscsiDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BootIscsiDevice) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BootIscsiDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *BootIscsiDevice) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *BootIscsiDevice) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o BootIscsiDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BootIscsiDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBootConfiguredDevice, errBootConfiguredDevice := json.Marshal(o.BootConfiguredDevice)
	if errBootConfiguredDevice != nil {
		return map[string]interface{}{}, errBootConfiguredDevice
	}
	errBootConfiguredDevice = json.Unmarshal([]byte(serializedBootConfiguredDevice), &toSerialize)
	if errBootConfiguredDevice != nil {
		return map[string]interface{}{}, errBootConfiguredDevice
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.ComputePhysical.IsSet() {
		toSerialize["ComputePhysical"] = o.ComputePhysical.Get()
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BootIscsiDevice) UnmarshalJSON(data []byte) (err error) {
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
	type BootIscsiDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType          string                                      `json:"ObjectType"`
		ComputePhysical     NullableComputePhysicalRelationship         `json:"ComputePhysical,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varBootIscsiDeviceWithoutEmbeddedStruct := BootIscsiDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBootIscsiDeviceWithoutEmbeddedStruct)
	if err == nil {
		varBootIscsiDevice := _BootIscsiDevice{}
		varBootIscsiDevice.ClassId = varBootIscsiDeviceWithoutEmbeddedStruct.ClassId
		varBootIscsiDevice.ObjectType = varBootIscsiDeviceWithoutEmbeddedStruct.ObjectType
		varBootIscsiDevice.ComputePhysical = varBootIscsiDeviceWithoutEmbeddedStruct.ComputePhysical
		varBootIscsiDevice.InventoryDeviceInfo = varBootIscsiDeviceWithoutEmbeddedStruct.InventoryDeviceInfo
		varBootIscsiDevice.RegisteredDevice = varBootIscsiDeviceWithoutEmbeddedStruct.RegisteredDevice
		*o = BootIscsiDevice(varBootIscsiDevice)
	} else {
		return err
	}

	varBootIscsiDevice := _BootIscsiDevice{}

	err = json.Unmarshal(data, &varBootIscsiDevice)
	if err == nil {
		o.BootConfiguredDevice = varBootIscsiDevice.BootConfiguredDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ComputePhysical")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectBootConfiguredDevice := reflect.ValueOf(o.BootConfiguredDevice)
		for i := 0; i < reflectBootConfiguredDevice.Type().NumField(); i++ {
			t := reflectBootConfiguredDevice.Type().Field(i)

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

type NullableBootIscsiDevice struct {
	value *BootIscsiDevice
	isSet bool
}

func (v NullableBootIscsiDevice) Get() *BootIscsiDevice {
	return v.value
}

func (v *NullableBootIscsiDevice) Set(val *BootIscsiDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableBootIscsiDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableBootIscsiDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootIscsiDevice(val *BootIscsiDevice) *NullableBootIscsiDevice {
	return &NullableBootIscsiDevice{value: val, isSet: true}
}

func (v NullableBootIscsiDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootIscsiDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
