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

// checks if the StorageHitachiDisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageHitachiDisk{}

// StorageHitachiDisk Disk entity associated with Hitachi storage array.
type StorageHitachiDisk struct {
	StorageBaseArrayDisk
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Drive type code of the Hitachi Disk.
	DriveTypeCode *string `json:"DriveTypeCode,omitempty"`
	// Parity group number. When the drive does not belong to any parity group, an empty character string is output.
	ParityGroupId *string `json:"ParityGroupId,omitempty"`
	// Drive type of the Hitachi Disk. * `N/A` - Drive Type is not available. * `SAS` - SAS stands for Serial Attached SCSI. * `SSD(MLC)` - SSD (MLC) stands for Multiple Level Cell. * `SSD(FMC)` - SSD (FMC) stands for Flash Memory Compressed. * `SSD(FMD)` - SSD (FMD) stands for Flash Module Drive. * `SSD(SLC)` - SSD (SLC) stands for Single Level Cell. * `SSD` - SSD stands for Solid-State Drive. * `SSD(RI)` - SSD (RI) stands for Read Intensive. * `SCM` - SCM stands for Storage Class Memory.
	TypeDetail *string `json:"TypeDetail,omitempty"`
	// Purpose for which the drive is used.
	Usage                *string                                       `json:"Usage,omitempty"`
	Array                NullableStorageHitachiArrayRelationship       `json:"Array,omitempty"`
	ParityGroup          NullableStorageHitachiParityGroupRelationship `json:"ParityGroup,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship   `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiDisk StorageHitachiDisk

// NewStorageHitachiDisk instantiates a new StorageHitachiDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiDisk(classId string, objectType string) *StorageHitachiDisk {
	this := StorageHitachiDisk{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiDiskWithDefaults instantiates a new StorageHitachiDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiDiskWithDefaults() *StorageHitachiDisk {
	this := StorageHitachiDisk{}
	var classId string = "storage.HitachiDisk"
	this.ClassId = classId
	var objectType string = "storage.HitachiDisk"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiDisk) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiDisk) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiDisk) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.HitachiDisk" of the ClassId field.
func (o *StorageHitachiDisk) GetDefaultClassId() interface{} {
	return "storage.HitachiDisk"
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiDisk) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiDisk) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiDisk) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.HitachiDisk" of the ObjectType field.
func (o *StorageHitachiDisk) GetDefaultObjectType() interface{} {
	return "storage.HitachiDisk"
}

// GetDriveTypeCode returns the DriveTypeCode field value if set, zero value otherwise.
func (o *StorageHitachiDisk) GetDriveTypeCode() string {
	if o == nil || IsNil(o.DriveTypeCode) {
		var ret string
		return ret
	}
	return *o.DriveTypeCode
}

// GetDriveTypeCodeOk returns a tuple with the DriveTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiDisk) GetDriveTypeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DriveTypeCode) {
		return nil, false
	}
	return o.DriveTypeCode, true
}

// HasDriveTypeCode returns a boolean if a field has been set.
func (o *StorageHitachiDisk) HasDriveTypeCode() bool {
	if o != nil && !IsNil(o.DriveTypeCode) {
		return true
	}

	return false
}

// SetDriveTypeCode gets a reference to the given string and assigns it to the DriveTypeCode field.
func (o *StorageHitachiDisk) SetDriveTypeCode(v string) {
	o.DriveTypeCode = &v
}

// GetParityGroupId returns the ParityGroupId field value if set, zero value otherwise.
func (o *StorageHitachiDisk) GetParityGroupId() string {
	if o == nil || IsNil(o.ParityGroupId) {
		var ret string
		return ret
	}
	return *o.ParityGroupId
}

// GetParityGroupIdOk returns a tuple with the ParityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiDisk) GetParityGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParityGroupId) {
		return nil, false
	}
	return o.ParityGroupId, true
}

// HasParityGroupId returns a boolean if a field has been set.
func (o *StorageHitachiDisk) HasParityGroupId() bool {
	if o != nil && !IsNil(o.ParityGroupId) {
		return true
	}

	return false
}

// SetParityGroupId gets a reference to the given string and assigns it to the ParityGroupId field.
func (o *StorageHitachiDisk) SetParityGroupId(v string) {
	o.ParityGroupId = &v
}

// GetTypeDetail returns the TypeDetail field value if set, zero value otherwise.
func (o *StorageHitachiDisk) GetTypeDetail() string {
	if o == nil || IsNil(o.TypeDetail) {
		var ret string
		return ret
	}
	return *o.TypeDetail
}

// GetTypeDetailOk returns a tuple with the TypeDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiDisk) GetTypeDetailOk() (*string, bool) {
	if o == nil || IsNil(o.TypeDetail) {
		return nil, false
	}
	return o.TypeDetail, true
}

// HasTypeDetail returns a boolean if a field has been set.
func (o *StorageHitachiDisk) HasTypeDetail() bool {
	if o != nil && !IsNil(o.TypeDetail) {
		return true
	}

	return false
}

// SetTypeDetail gets a reference to the given string and assigns it to the TypeDetail field.
func (o *StorageHitachiDisk) SetTypeDetail(v string) {
	o.TypeDetail = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *StorageHitachiDisk) GetUsage() string {
	if o == nil || IsNil(o.Usage) {
		var ret string
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiDisk) GetUsageOk() (*string, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *StorageHitachiDisk) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given string and assigns it to the Usage field.
func (o *StorageHitachiDisk) SetUsage(v string) {
	o.Usage = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiDisk) GetArray() StorageHitachiArrayRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiDisk) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiDisk) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiDisk) SetArray(v StorageHitachiArrayRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StorageHitachiDisk) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StorageHitachiDisk) UnsetArray() {
	o.Array.Unset()
}

// GetParityGroup returns the ParityGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiDisk) GetParityGroup() StorageHitachiParityGroupRelationship {
	if o == nil || IsNil(o.ParityGroup.Get()) {
		var ret StorageHitachiParityGroupRelationship
		return ret
	}
	return *o.ParityGroup.Get()
}

// GetParityGroupOk returns a tuple with the ParityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiDisk) GetParityGroupOk() (*StorageHitachiParityGroupRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParityGroup.Get(), o.ParityGroup.IsSet()
}

// HasParityGroup returns a boolean if a field has been set.
func (o *StorageHitachiDisk) HasParityGroup() bool {
	if o != nil && o.ParityGroup.IsSet() {
		return true
	}

	return false
}

// SetParityGroup gets a reference to the given NullableStorageHitachiParityGroupRelationship and assigns it to the ParityGroup field.
func (o *StorageHitachiDisk) SetParityGroup(v StorageHitachiParityGroupRelationship) {
	o.ParityGroup.Set(&v)
}

// SetParityGroupNil sets the value for ParityGroup to be an explicit nil
func (o *StorageHitachiDisk) SetParityGroupNil() {
	o.ParityGroup.Set(nil)
}

// UnsetParityGroup ensures that no value is present for ParityGroup, not even an explicit nil
func (o *StorageHitachiDisk) UnsetParityGroup() {
	o.ParityGroup.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiDisk) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiDisk) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiDisk) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiDisk) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageHitachiDisk) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageHitachiDisk) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o StorageHitachiDisk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageHitachiDisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseArrayDisk, errStorageBaseArrayDisk := json.Marshal(o.StorageBaseArrayDisk)
	if errStorageBaseArrayDisk != nil {
		return map[string]interface{}{}, errStorageBaseArrayDisk
	}
	errStorageBaseArrayDisk = json.Unmarshal([]byte(serializedStorageBaseArrayDisk), &toSerialize)
	if errStorageBaseArrayDisk != nil {
		return map[string]interface{}{}, errStorageBaseArrayDisk
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DriveTypeCode) {
		toSerialize["DriveTypeCode"] = o.DriveTypeCode
	}
	if !IsNil(o.ParityGroupId) {
		toSerialize["ParityGroupId"] = o.ParityGroupId
	}
	if !IsNil(o.TypeDetail) {
		toSerialize["TypeDetail"] = o.TypeDetail
	}
	if !IsNil(o.Usage) {
		toSerialize["Usage"] = o.Usage
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}
	if o.ParityGroup.IsSet() {
		toSerialize["ParityGroup"] = o.ParityGroup.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageHitachiDisk) UnmarshalJSON(data []byte) (err error) {
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
	type StorageHitachiDiskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Drive type code of the Hitachi Disk.
		DriveTypeCode *string `json:"DriveTypeCode,omitempty"`
		// Parity group number. When the drive does not belong to any parity group, an empty character string is output.
		ParityGroupId *string `json:"ParityGroupId,omitempty"`
		// Drive type of the Hitachi Disk. * `N/A` - Drive Type is not available. * `SAS` - SAS stands for Serial Attached SCSI. * `SSD(MLC)` - SSD (MLC) stands for Multiple Level Cell. * `SSD(FMC)` - SSD (FMC) stands for Flash Memory Compressed. * `SSD(FMD)` - SSD (FMD) stands for Flash Module Drive. * `SSD(SLC)` - SSD (SLC) stands for Single Level Cell. * `SSD` - SSD stands for Solid-State Drive. * `SSD(RI)` - SSD (RI) stands for Read Intensive. * `SCM` - SCM stands for Storage Class Memory.
		TypeDetail *string `json:"TypeDetail,omitempty"`
		// Purpose for which the drive is used.
		Usage            *string                                       `json:"Usage,omitempty"`
		Array            NullableStorageHitachiArrayRelationship       `json:"Array,omitempty"`
		ParityGroup      NullableStorageHitachiParityGroupRelationship `json:"ParityGroup,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship   `json:"RegisteredDevice,omitempty"`
	}

	varStorageHitachiDiskWithoutEmbeddedStruct := StorageHitachiDiskWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageHitachiDiskWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiDisk := _StorageHitachiDisk{}
		varStorageHitachiDisk.ClassId = varStorageHitachiDiskWithoutEmbeddedStruct.ClassId
		varStorageHitachiDisk.ObjectType = varStorageHitachiDiskWithoutEmbeddedStruct.ObjectType
		varStorageHitachiDisk.DriveTypeCode = varStorageHitachiDiskWithoutEmbeddedStruct.DriveTypeCode
		varStorageHitachiDisk.ParityGroupId = varStorageHitachiDiskWithoutEmbeddedStruct.ParityGroupId
		varStorageHitachiDisk.TypeDetail = varStorageHitachiDiskWithoutEmbeddedStruct.TypeDetail
		varStorageHitachiDisk.Usage = varStorageHitachiDiskWithoutEmbeddedStruct.Usage
		varStorageHitachiDisk.Array = varStorageHitachiDiskWithoutEmbeddedStruct.Array
		varStorageHitachiDisk.ParityGroup = varStorageHitachiDiskWithoutEmbeddedStruct.ParityGroup
		varStorageHitachiDisk.RegisteredDevice = varStorageHitachiDiskWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageHitachiDisk(varStorageHitachiDisk)
	} else {
		return err
	}

	varStorageHitachiDisk := _StorageHitachiDisk{}

	err = json.Unmarshal(data, &varStorageHitachiDisk)
	if err == nil {
		o.StorageBaseArrayDisk = varStorageHitachiDisk.StorageBaseArrayDisk
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriveTypeCode")
		delete(additionalProperties, "ParityGroupId")
		delete(additionalProperties, "TypeDetail")
		delete(additionalProperties, "Usage")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ParityGroup")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseArrayDisk := reflect.ValueOf(o.StorageBaseArrayDisk)
		for i := 0; i < reflectStorageBaseArrayDisk.Type().NumField(); i++ {
			t := reflectStorageBaseArrayDisk.Type().Field(i)

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

type NullableStorageHitachiDisk struct {
	value *StorageHitachiDisk
	isSet bool
}

func (v NullableStorageHitachiDisk) Get() *StorageHitachiDisk {
	return v.value
}

func (v *NullableStorageHitachiDisk) Set(val *StorageHitachiDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiDisk(val *StorageHitachiDisk) *NullableStorageHitachiDisk {
	return &NullableStorageHitachiDisk{value: val, isSet: true}
}

func (v NullableStorageHitachiDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
