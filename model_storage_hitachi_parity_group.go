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

// checks if the StorageHitachiParityGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageHitachiParityGroup{}

// StorageHitachiParityGroup A parity group in Hitachi storage array.
type StorageHitachiParityGroup struct {
	StorageBaseRaidGroup
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Speed (rpm) of the disk belonging to the parity group.
	DiskSpeed *string `json:"DiskSpeed,omitempty"`
	// Type of the disk belonging to the parity group.
	DiskType *string `json:"DiskType,omitempty"`
	// Value of the accelerated compression of the parity group. true, Accelerated compression for the parity group is enabled. false, Accelerated compression for the parity group is disabled.
	IsAcceleratedCompressionEnabled *bool `json:"IsAcceleratedCompressionEnabled,omitempty"`
	// Value of the copy back mode setting of the parity group. true, Copy back mode is enabled. false, Copy back mode is disabled.
	IsCopyBackModeEnabled *bool `json:"IsCopyBackModeEnabled,omitempty"`
	// Value of the encryption setting of the parity group. true, Encryption is enabled. false, Encryption is disabled.
	IsEncryptionEnabled  *bool                                       `json:"IsEncryptionEnabled,omitempty"`
	Array                NullableStorageHitachiArrayRelationship     `json:"Array,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiParityGroup StorageHitachiParityGroup

// NewStorageHitachiParityGroup instantiates a new StorageHitachiParityGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiParityGroup(classId string, objectType string) *StorageHitachiParityGroup {
	this := StorageHitachiParityGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiParityGroupWithDefaults instantiates a new StorageHitachiParityGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiParityGroupWithDefaults() *StorageHitachiParityGroup {
	this := StorageHitachiParityGroup{}
	var classId string = "storage.HitachiParityGroup"
	this.ClassId = classId
	var objectType string = "storage.HitachiParityGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiParityGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiParityGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.HitachiParityGroup" of the ClassId field.
func (o *StorageHitachiParityGroup) GetDefaultClassId() interface{} {
	return "storage.HitachiParityGroup"
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiParityGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiParityGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.HitachiParityGroup" of the ObjectType field.
func (o *StorageHitachiParityGroup) GetDefaultObjectType() interface{} {
	return "storage.HitachiParityGroup"
}

// GetDiskSpeed returns the DiskSpeed field value if set, zero value otherwise.
func (o *StorageHitachiParityGroup) GetDiskSpeed() string {
	if o == nil || IsNil(o.DiskSpeed) {
		var ret string
		return ret
	}
	return *o.DiskSpeed
}

// GetDiskSpeedOk returns a tuple with the DiskSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetDiskSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.DiskSpeed) {
		return nil, false
	}
	return o.DiskSpeed, true
}

// HasDiskSpeed returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasDiskSpeed() bool {
	if o != nil && !IsNil(o.DiskSpeed) {
		return true
	}

	return false
}

// SetDiskSpeed gets a reference to the given string and assigns it to the DiskSpeed field.
func (o *StorageHitachiParityGroup) SetDiskSpeed(v string) {
	o.DiskSpeed = &v
}

// GetDiskType returns the DiskType field value if set, zero value otherwise.
func (o *StorageHitachiParityGroup) GetDiskType() string {
	if o == nil || IsNil(o.DiskType) {
		var ret string
		return ret
	}
	return *o.DiskType
}

// GetDiskTypeOk returns a tuple with the DiskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetDiskTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DiskType) {
		return nil, false
	}
	return o.DiskType, true
}

// HasDiskType returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasDiskType() bool {
	if o != nil && !IsNil(o.DiskType) {
		return true
	}

	return false
}

// SetDiskType gets a reference to the given string and assigns it to the DiskType field.
func (o *StorageHitachiParityGroup) SetDiskType(v string) {
	o.DiskType = &v
}

// GetIsAcceleratedCompressionEnabled returns the IsAcceleratedCompressionEnabled field value if set, zero value otherwise.
func (o *StorageHitachiParityGroup) GetIsAcceleratedCompressionEnabled() bool {
	if o == nil || IsNil(o.IsAcceleratedCompressionEnabled) {
		var ret bool
		return ret
	}
	return *o.IsAcceleratedCompressionEnabled
}

// GetIsAcceleratedCompressionEnabledOk returns a tuple with the IsAcceleratedCompressionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetIsAcceleratedCompressionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAcceleratedCompressionEnabled) {
		return nil, false
	}
	return o.IsAcceleratedCompressionEnabled, true
}

// HasIsAcceleratedCompressionEnabled returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasIsAcceleratedCompressionEnabled() bool {
	if o != nil && !IsNil(o.IsAcceleratedCompressionEnabled) {
		return true
	}

	return false
}

// SetIsAcceleratedCompressionEnabled gets a reference to the given bool and assigns it to the IsAcceleratedCompressionEnabled field.
func (o *StorageHitachiParityGroup) SetIsAcceleratedCompressionEnabled(v bool) {
	o.IsAcceleratedCompressionEnabled = &v
}

// GetIsCopyBackModeEnabled returns the IsCopyBackModeEnabled field value if set, zero value otherwise.
func (o *StorageHitachiParityGroup) GetIsCopyBackModeEnabled() bool {
	if o == nil || IsNil(o.IsCopyBackModeEnabled) {
		var ret bool
		return ret
	}
	return *o.IsCopyBackModeEnabled
}

// GetIsCopyBackModeEnabledOk returns a tuple with the IsCopyBackModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetIsCopyBackModeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCopyBackModeEnabled) {
		return nil, false
	}
	return o.IsCopyBackModeEnabled, true
}

// HasIsCopyBackModeEnabled returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasIsCopyBackModeEnabled() bool {
	if o != nil && !IsNil(o.IsCopyBackModeEnabled) {
		return true
	}

	return false
}

// SetIsCopyBackModeEnabled gets a reference to the given bool and assigns it to the IsCopyBackModeEnabled field.
func (o *StorageHitachiParityGroup) SetIsCopyBackModeEnabled(v bool) {
	o.IsCopyBackModeEnabled = &v
}

// GetIsEncryptionEnabled returns the IsEncryptionEnabled field value if set, zero value otherwise.
func (o *StorageHitachiParityGroup) GetIsEncryptionEnabled() bool {
	if o == nil || IsNil(o.IsEncryptionEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEncryptionEnabled
}

// GetIsEncryptionEnabledOk returns a tuple with the IsEncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetIsEncryptionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEncryptionEnabled) {
		return nil, false
	}
	return o.IsEncryptionEnabled, true
}

// HasIsEncryptionEnabled returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasIsEncryptionEnabled() bool {
	if o != nil && !IsNil(o.IsEncryptionEnabled) {
		return true
	}

	return false
}

// SetIsEncryptionEnabled gets a reference to the given bool and assigns it to the IsEncryptionEnabled field.
func (o *StorageHitachiParityGroup) SetIsEncryptionEnabled(v bool) {
	o.IsEncryptionEnabled = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiParityGroup) GetArray() StorageHitachiArrayRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiParityGroup) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiParityGroup) SetArray(v StorageHitachiArrayRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StorageHitachiParityGroup) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StorageHitachiParityGroup) UnsetArray() {
	o.Array.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiParityGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiParityGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiParityGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageHitachiParityGroup) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageHitachiParityGroup) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o StorageHitachiParityGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageHitachiParityGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseRaidGroup, errStorageBaseRaidGroup := json.Marshal(o.StorageBaseRaidGroup)
	if errStorageBaseRaidGroup != nil {
		return map[string]interface{}{}, errStorageBaseRaidGroup
	}
	errStorageBaseRaidGroup = json.Unmarshal([]byte(serializedStorageBaseRaidGroup), &toSerialize)
	if errStorageBaseRaidGroup != nil {
		return map[string]interface{}{}, errStorageBaseRaidGroup
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DiskSpeed) {
		toSerialize["DiskSpeed"] = o.DiskSpeed
	}
	if !IsNil(o.DiskType) {
		toSerialize["DiskType"] = o.DiskType
	}
	if !IsNil(o.IsAcceleratedCompressionEnabled) {
		toSerialize["IsAcceleratedCompressionEnabled"] = o.IsAcceleratedCompressionEnabled
	}
	if !IsNil(o.IsCopyBackModeEnabled) {
		toSerialize["IsCopyBackModeEnabled"] = o.IsCopyBackModeEnabled
	}
	if !IsNil(o.IsEncryptionEnabled) {
		toSerialize["IsEncryptionEnabled"] = o.IsEncryptionEnabled
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageHitachiParityGroup) UnmarshalJSON(data []byte) (err error) {
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
	type StorageHitachiParityGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Speed (rpm) of the disk belonging to the parity group.
		DiskSpeed *string `json:"DiskSpeed,omitempty"`
		// Type of the disk belonging to the parity group.
		DiskType *string `json:"DiskType,omitempty"`
		// Value of the accelerated compression of the parity group. true, Accelerated compression for the parity group is enabled. false, Accelerated compression for the parity group is disabled.
		IsAcceleratedCompressionEnabled *bool `json:"IsAcceleratedCompressionEnabled,omitempty"`
		// Value of the copy back mode setting of the parity group. true, Copy back mode is enabled. false, Copy back mode is disabled.
		IsCopyBackModeEnabled *bool `json:"IsCopyBackModeEnabled,omitempty"`
		// Value of the encryption setting of the parity group. true, Encryption is enabled. false, Encryption is disabled.
		IsEncryptionEnabled *bool                                       `json:"IsEncryptionEnabled,omitempty"`
		Array               NullableStorageHitachiArrayRelationship     `json:"Array,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageHitachiParityGroupWithoutEmbeddedStruct := StorageHitachiParityGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageHitachiParityGroupWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiParityGroup := _StorageHitachiParityGroup{}
		varStorageHitachiParityGroup.ClassId = varStorageHitachiParityGroupWithoutEmbeddedStruct.ClassId
		varStorageHitachiParityGroup.ObjectType = varStorageHitachiParityGroupWithoutEmbeddedStruct.ObjectType
		varStorageHitachiParityGroup.DiskSpeed = varStorageHitachiParityGroupWithoutEmbeddedStruct.DiskSpeed
		varStorageHitachiParityGroup.DiskType = varStorageHitachiParityGroupWithoutEmbeddedStruct.DiskType
		varStorageHitachiParityGroup.IsAcceleratedCompressionEnabled = varStorageHitachiParityGroupWithoutEmbeddedStruct.IsAcceleratedCompressionEnabled
		varStorageHitachiParityGroup.IsCopyBackModeEnabled = varStorageHitachiParityGroupWithoutEmbeddedStruct.IsCopyBackModeEnabled
		varStorageHitachiParityGroup.IsEncryptionEnabled = varStorageHitachiParityGroupWithoutEmbeddedStruct.IsEncryptionEnabled
		varStorageHitachiParityGroup.Array = varStorageHitachiParityGroupWithoutEmbeddedStruct.Array
		varStorageHitachiParityGroup.RegisteredDevice = varStorageHitachiParityGroupWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageHitachiParityGroup(varStorageHitachiParityGroup)
	} else {
		return err
	}

	varStorageHitachiParityGroup := _StorageHitachiParityGroup{}

	err = json.Unmarshal(data, &varStorageHitachiParityGroup)
	if err == nil {
		o.StorageBaseRaidGroup = varStorageHitachiParityGroup.StorageBaseRaidGroup
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DiskSpeed")
		delete(additionalProperties, "DiskType")
		delete(additionalProperties, "IsAcceleratedCompressionEnabled")
		delete(additionalProperties, "IsCopyBackModeEnabled")
		delete(additionalProperties, "IsEncryptionEnabled")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseRaidGroup := reflect.ValueOf(o.StorageBaseRaidGroup)
		for i := 0; i < reflectStorageBaseRaidGroup.Type().NumField(); i++ {
			t := reflectStorageBaseRaidGroup.Type().Field(i)

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

type NullableStorageHitachiParityGroup struct {
	value *StorageHitachiParityGroup
	isSet bool
}

func (v NullableStorageHitachiParityGroup) Get() *StorageHitachiParityGroup {
	return v.value
}

func (v *NullableStorageHitachiParityGroup) Set(val *StorageHitachiParityGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiParityGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiParityGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiParityGroup(val *StorageHitachiParityGroup) *NullableStorageHitachiParityGroup {
	return &NullableStorageHitachiParityGroup{value: val, isSet: true}
}

func (v NullableStorageHitachiParityGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiParityGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
