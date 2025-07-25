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

// checks if the StoragePureProtectionGroupSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoragePureProtectionGroupSnapshot{}

// StoragePureProtectionGroupSnapshot Protection group snapshot entity in Pure protection group.
type StoragePureProtectionGroupSnapshot struct {
	StorageBaseProtectionGroupSnapshot
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                         `json:"ObjectType"`
	Array                NullableStoragePureArrayRelationship           `json:"Array,omitempty"`
	ProtectionGroup      NullableStoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureProtectionGroupSnapshot StoragePureProtectionGroupSnapshot

// NewStoragePureProtectionGroupSnapshot instantiates a new StoragePureProtectionGroupSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureProtectionGroupSnapshot(classId string, objectType string) *StoragePureProtectionGroupSnapshot {
	this := StoragePureProtectionGroupSnapshot{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureProtectionGroupSnapshotWithDefaults instantiates a new StoragePureProtectionGroupSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureProtectionGroupSnapshotWithDefaults() *StoragePureProtectionGroupSnapshot {
	this := StoragePureProtectionGroupSnapshot{}
	var classId string = "storage.PureProtectionGroupSnapshot"
	this.ClassId = classId
	var objectType string = "storage.PureProtectionGroupSnapshot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureProtectionGroupSnapshot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureProtectionGroupSnapshot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureProtectionGroupSnapshot) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.PureProtectionGroupSnapshot" of the ClassId field.
func (o *StoragePureProtectionGroupSnapshot) GetDefaultClassId() interface{} {
	return "storage.PureProtectionGroupSnapshot"
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureProtectionGroupSnapshot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureProtectionGroupSnapshot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureProtectionGroupSnapshot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.PureProtectionGroupSnapshot" of the ObjectType field.
func (o *StoragePureProtectionGroupSnapshot) GetDefaultObjectType() interface{} {
	return "storage.PureProtectionGroupSnapshot"
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureProtectionGroupSnapshot) GetArray() StoragePureArrayRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureProtectionGroupSnapshot) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureProtectionGroupSnapshot) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureProtectionGroupSnapshot) SetArray(v StoragePureArrayRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StoragePureProtectionGroupSnapshot) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StoragePureProtectionGroupSnapshot) UnsetArray() {
	o.Array.Unset()
}

// GetProtectionGroup returns the ProtectionGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureProtectionGroupSnapshot) GetProtectionGroup() StoragePureProtectionGroupRelationship {
	if o == nil || IsNil(o.ProtectionGroup.Get()) {
		var ret StoragePureProtectionGroupRelationship
		return ret
	}
	return *o.ProtectionGroup.Get()
}

// GetProtectionGroupOk returns a tuple with the ProtectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureProtectionGroupSnapshot) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionGroup.Get(), o.ProtectionGroup.IsSet()
}

// HasProtectionGroup returns a boolean if a field has been set.
func (o *StoragePureProtectionGroupSnapshot) HasProtectionGroup() bool {
	if o != nil && o.ProtectionGroup.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroup gets a reference to the given NullableStoragePureProtectionGroupRelationship and assigns it to the ProtectionGroup field.
func (o *StoragePureProtectionGroupSnapshot) SetProtectionGroup(v StoragePureProtectionGroupRelationship) {
	o.ProtectionGroup.Set(&v)
}

// SetProtectionGroupNil sets the value for ProtectionGroup to be an explicit nil
func (o *StoragePureProtectionGroupSnapshot) SetProtectionGroupNil() {
	o.ProtectionGroup.Set(nil)
}

// UnsetProtectionGroup ensures that no value is present for ProtectionGroup, not even an explicit nil
func (o *StoragePureProtectionGroupSnapshot) UnsetProtectionGroup() {
	o.ProtectionGroup.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureProtectionGroupSnapshot) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureProtectionGroupSnapshot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureProtectionGroupSnapshot) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureProtectionGroupSnapshot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StoragePureProtectionGroupSnapshot) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StoragePureProtectionGroupSnapshot) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o StoragePureProtectionGroupSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoragePureProtectionGroupSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseProtectionGroupSnapshot, errStorageBaseProtectionGroupSnapshot := json.Marshal(o.StorageBaseProtectionGroupSnapshot)
	if errStorageBaseProtectionGroupSnapshot != nil {
		return map[string]interface{}{}, errStorageBaseProtectionGroupSnapshot
	}
	errStorageBaseProtectionGroupSnapshot = json.Unmarshal([]byte(serializedStorageBaseProtectionGroupSnapshot), &toSerialize)
	if errStorageBaseProtectionGroupSnapshot != nil {
		return map[string]interface{}{}, errStorageBaseProtectionGroupSnapshot
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}
	if o.ProtectionGroup.IsSet() {
		toSerialize["ProtectionGroup"] = o.ProtectionGroup.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StoragePureProtectionGroupSnapshot) UnmarshalJSON(data []byte) (err error) {
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
	type StoragePureProtectionGroupSnapshotWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string                                         `json:"ObjectType"`
		Array            NullableStoragePureArrayRelationship           `json:"Array,omitempty"`
		ProtectionGroup  NullableStoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	}

	varStoragePureProtectionGroupSnapshotWithoutEmbeddedStruct := StoragePureProtectionGroupSnapshotWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStoragePureProtectionGroupSnapshotWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureProtectionGroupSnapshot := _StoragePureProtectionGroupSnapshot{}
		varStoragePureProtectionGroupSnapshot.ClassId = varStoragePureProtectionGroupSnapshotWithoutEmbeddedStruct.ClassId
		varStoragePureProtectionGroupSnapshot.ObjectType = varStoragePureProtectionGroupSnapshotWithoutEmbeddedStruct.ObjectType
		varStoragePureProtectionGroupSnapshot.Array = varStoragePureProtectionGroupSnapshotWithoutEmbeddedStruct.Array
		varStoragePureProtectionGroupSnapshot.ProtectionGroup = varStoragePureProtectionGroupSnapshotWithoutEmbeddedStruct.ProtectionGroup
		varStoragePureProtectionGroupSnapshot.RegisteredDevice = varStoragePureProtectionGroupSnapshotWithoutEmbeddedStruct.RegisteredDevice
		*o = StoragePureProtectionGroupSnapshot(varStoragePureProtectionGroupSnapshot)
	} else {
		return err
	}

	varStoragePureProtectionGroupSnapshot := _StoragePureProtectionGroupSnapshot{}

	err = json.Unmarshal(data, &varStoragePureProtectionGroupSnapshot)
	if err == nil {
		o.StorageBaseProtectionGroupSnapshot = varStoragePureProtectionGroupSnapshot.StorageBaseProtectionGroupSnapshot
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ProtectionGroup")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseProtectionGroupSnapshot := reflect.ValueOf(o.StorageBaseProtectionGroupSnapshot)
		for i := 0; i < reflectStorageBaseProtectionGroupSnapshot.Type().NumField(); i++ {
			t := reflectStorageBaseProtectionGroupSnapshot.Type().Field(i)

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

type NullableStoragePureProtectionGroupSnapshot struct {
	value *StoragePureProtectionGroupSnapshot
	isSet bool
}

func (v NullableStoragePureProtectionGroupSnapshot) Get() *StoragePureProtectionGroupSnapshot {
	return v.value
}

func (v *NullableStoragePureProtectionGroupSnapshot) Set(val *StoragePureProtectionGroupSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureProtectionGroupSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureProtectionGroupSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureProtectionGroupSnapshot(val *StoragePureProtectionGroupSnapshot) *NullableStoragePureProtectionGroupSnapshot {
	return &NullableStoragePureProtectionGroupSnapshot{value: val, isSet: true}
}

func (v NullableStoragePureProtectionGroupSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureProtectionGroupSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
