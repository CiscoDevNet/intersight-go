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

// checks if the HciEsxiVmDisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HciEsxiVmDisk{}

// HciEsxiVmDisk A disk associated with an ESXi VM.
type HciEsxiVmDisk struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The bus type of the disk. Possible values are 'IDE', 'SCSI', 'SATA', 'NVME'.
	BusType *string `json:"BusType,omitempty"`
	// The unique identifier of the disk.
	DiskExtId *string `json:"DiskExtId,omitempty"`
	// The size of the disk in bytes.
	DiskSizeBytes *int64 `json:"DiskSizeBytes,omitempty"`
	// The index of the disk, similar to a slot number on physical machine.
	Index *int32 `json:"Index,omitempty"`
	// Indicates whether the virtual disk is pinned to the hot tier or not.
	IsFlashModeEnabled *bool `json:"IsFlashModeEnabled,omitempty"`
	// The extId of the storage container which backs this disk.
	StorageContainerExtId *string `json:"StorageContainerExtId,omitempty"`
	// The unique identifier of the VM.
	VmExtId              *string                                     `json:"VmExtId,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	Vm                   NullableHciEsxiVmRelationship               `json:"Vm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HciEsxiVmDisk HciEsxiVmDisk

// NewHciEsxiVmDisk instantiates a new HciEsxiVmDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHciEsxiVmDisk(classId string, objectType string) *HciEsxiVmDisk {
	this := HciEsxiVmDisk{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHciEsxiVmDiskWithDefaults instantiates a new HciEsxiVmDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHciEsxiVmDiskWithDefaults() *HciEsxiVmDisk {
	this := HciEsxiVmDisk{}
	var classId string = "hci.EsxiVmDisk"
	this.ClassId = classId
	var objectType string = "hci.EsxiVmDisk"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HciEsxiVmDisk) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HciEsxiVmDisk) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HciEsxiVmDisk) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hci.EsxiVmDisk" of the ClassId field.
func (o *HciEsxiVmDisk) GetDefaultClassId() interface{} {
	return "hci.EsxiVmDisk"
}

// GetObjectType returns the ObjectType field value
func (o *HciEsxiVmDisk) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HciEsxiVmDisk) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HciEsxiVmDisk) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hci.EsxiVmDisk" of the ObjectType field.
func (o *HciEsxiVmDisk) GetDefaultObjectType() interface{} {
	return "hci.EsxiVmDisk"
}

// GetBusType returns the BusType field value if set, zero value otherwise.
func (o *HciEsxiVmDisk) GetBusType() string {
	if o == nil || IsNil(o.BusType) {
		var ret string
		return ret
	}
	return *o.BusType
}

// GetBusTypeOk returns a tuple with the BusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEsxiVmDisk) GetBusTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BusType) {
		return nil, false
	}
	return o.BusType, true
}

// HasBusType returns a boolean if a field has been set.
func (o *HciEsxiVmDisk) HasBusType() bool {
	if o != nil && !IsNil(o.BusType) {
		return true
	}

	return false
}

// SetBusType gets a reference to the given string and assigns it to the BusType field.
func (o *HciEsxiVmDisk) SetBusType(v string) {
	o.BusType = &v
}

// GetDiskExtId returns the DiskExtId field value if set, zero value otherwise.
func (o *HciEsxiVmDisk) GetDiskExtId() string {
	if o == nil || IsNil(o.DiskExtId) {
		var ret string
		return ret
	}
	return *o.DiskExtId
}

// GetDiskExtIdOk returns a tuple with the DiskExtId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEsxiVmDisk) GetDiskExtIdOk() (*string, bool) {
	if o == nil || IsNil(o.DiskExtId) {
		return nil, false
	}
	return o.DiskExtId, true
}

// HasDiskExtId returns a boolean if a field has been set.
func (o *HciEsxiVmDisk) HasDiskExtId() bool {
	if o != nil && !IsNil(o.DiskExtId) {
		return true
	}

	return false
}

// SetDiskExtId gets a reference to the given string and assigns it to the DiskExtId field.
func (o *HciEsxiVmDisk) SetDiskExtId(v string) {
	o.DiskExtId = &v
}

// GetDiskSizeBytes returns the DiskSizeBytes field value if set, zero value otherwise.
func (o *HciEsxiVmDisk) GetDiskSizeBytes() int64 {
	if o == nil || IsNil(o.DiskSizeBytes) {
		var ret int64
		return ret
	}
	return *o.DiskSizeBytes
}

// GetDiskSizeBytesOk returns a tuple with the DiskSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEsxiVmDisk) GetDiskSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DiskSizeBytes) {
		return nil, false
	}
	return o.DiskSizeBytes, true
}

// HasDiskSizeBytes returns a boolean if a field has been set.
func (o *HciEsxiVmDisk) HasDiskSizeBytes() bool {
	if o != nil && !IsNil(o.DiskSizeBytes) {
		return true
	}

	return false
}

// SetDiskSizeBytes gets a reference to the given int64 and assigns it to the DiskSizeBytes field.
func (o *HciEsxiVmDisk) SetDiskSizeBytes(v int64) {
	o.DiskSizeBytes = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *HciEsxiVmDisk) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEsxiVmDisk) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *HciEsxiVmDisk) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *HciEsxiVmDisk) SetIndex(v int32) {
	o.Index = &v
}

// GetIsFlashModeEnabled returns the IsFlashModeEnabled field value if set, zero value otherwise.
func (o *HciEsxiVmDisk) GetIsFlashModeEnabled() bool {
	if o == nil || IsNil(o.IsFlashModeEnabled) {
		var ret bool
		return ret
	}
	return *o.IsFlashModeEnabled
}

// GetIsFlashModeEnabledOk returns a tuple with the IsFlashModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEsxiVmDisk) GetIsFlashModeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFlashModeEnabled) {
		return nil, false
	}
	return o.IsFlashModeEnabled, true
}

// HasIsFlashModeEnabled returns a boolean if a field has been set.
func (o *HciEsxiVmDisk) HasIsFlashModeEnabled() bool {
	if o != nil && !IsNil(o.IsFlashModeEnabled) {
		return true
	}

	return false
}

// SetIsFlashModeEnabled gets a reference to the given bool and assigns it to the IsFlashModeEnabled field.
func (o *HciEsxiVmDisk) SetIsFlashModeEnabled(v bool) {
	o.IsFlashModeEnabled = &v
}

// GetStorageContainerExtId returns the StorageContainerExtId field value if set, zero value otherwise.
func (o *HciEsxiVmDisk) GetStorageContainerExtId() string {
	if o == nil || IsNil(o.StorageContainerExtId) {
		var ret string
		return ret
	}
	return *o.StorageContainerExtId
}

// GetStorageContainerExtIdOk returns a tuple with the StorageContainerExtId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEsxiVmDisk) GetStorageContainerExtIdOk() (*string, bool) {
	if o == nil || IsNil(o.StorageContainerExtId) {
		return nil, false
	}
	return o.StorageContainerExtId, true
}

// HasStorageContainerExtId returns a boolean if a field has been set.
func (o *HciEsxiVmDisk) HasStorageContainerExtId() bool {
	if o != nil && !IsNil(o.StorageContainerExtId) {
		return true
	}

	return false
}

// SetStorageContainerExtId gets a reference to the given string and assigns it to the StorageContainerExtId field.
func (o *HciEsxiVmDisk) SetStorageContainerExtId(v string) {
	o.StorageContainerExtId = &v
}

// GetVmExtId returns the VmExtId field value if set, zero value otherwise.
func (o *HciEsxiVmDisk) GetVmExtId() string {
	if o == nil || IsNil(o.VmExtId) {
		var ret string
		return ret
	}
	return *o.VmExtId
}

// GetVmExtIdOk returns a tuple with the VmExtId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEsxiVmDisk) GetVmExtIdOk() (*string, bool) {
	if o == nil || IsNil(o.VmExtId) {
		return nil, false
	}
	return o.VmExtId, true
}

// HasVmExtId returns a boolean if a field has been set.
func (o *HciEsxiVmDisk) HasVmExtId() bool {
	if o != nil && !IsNil(o.VmExtId) {
		return true
	}

	return false
}

// SetVmExtId gets a reference to the given string and assigns it to the VmExtId field.
func (o *HciEsxiVmDisk) SetVmExtId(v string) {
	o.VmExtId = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HciEsxiVmDisk) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HciEsxiVmDisk) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HciEsxiVmDisk) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HciEsxiVmDisk) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *HciEsxiVmDisk) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *HciEsxiVmDisk) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetVm returns the Vm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HciEsxiVmDisk) GetVm() HciEsxiVmRelationship {
	if o == nil || IsNil(o.Vm.Get()) {
		var ret HciEsxiVmRelationship
		return ret
	}
	return *o.Vm.Get()
}

// GetVmOk returns a tuple with the Vm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HciEsxiVmDisk) GetVmOk() (*HciEsxiVmRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vm.Get(), o.Vm.IsSet()
}

// HasVm returns a boolean if a field has been set.
func (o *HciEsxiVmDisk) HasVm() bool {
	if o != nil && o.Vm.IsSet() {
		return true
	}

	return false
}

// SetVm gets a reference to the given NullableHciEsxiVmRelationship and assigns it to the Vm field.
func (o *HciEsxiVmDisk) SetVm(v HciEsxiVmRelationship) {
	o.Vm.Set(&v)
}

// SetVmNil sets the value for Vm to be an explicit nil
func (o *HciEsxiVmDisk) SetVmNil() {
	o.Vm.Set(nil)
}

// UnsetVm ensures that no value is present for Vm, not even an explicit nil
func (o *HciEsxiVmDisk) UnsetVm() {
	o.Vm.Unset()
}

func (o HciEsxiVmDisk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HciEsxiVmDisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.BusType) {
		toSerialize["BusType"] = o.BusType
	}
	if !IsNil(o.DiskExtId) {
		toSerialize["DiskExtId"] = o.DiskExtId
	}
	if !IsNil(o.DiskSizeBytes) {
		toSerialize["DiskSizeBytes"] = o.DiskSizeBytes
	}
	if !IsNil(o.Index) {
		toSerialize["Index"] = o.Index
	}
	if !IsNil(o.IsFlashModeEnabled) {
		toSerialize["IsFlashModeEnabled"] = o.IsFlashModeEnabled
	}
	if !IsNil(o.StorageContainerExtId) {
		toSerialize["StorageContainerExtId"] = o.StorageContainerExtId
	}
	if !IsNil(o.VmExtId) {
		toSerialize["VmExtId"] = o.VmExtId
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.Vm.IsSet() {
		toSerialize["Vm"] = o.Vm.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HciEsxiVmDisk) UnmarshalJSON(data []byte) (err error) {
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
	type HciEsxiVmDiskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The bus type of the disk. Possible values are 'IDE', 'SCSI', 'SATA', 'NVME'.
		BusType *string `json:"BusType,omitempty"`
		// The unique identifier of the disk.
		DiskExtId *string `json:"DiskExtId,omitempty"`
		// The size of the disk in bytes.
		DiskSizeBytes *int64 `json:"DiskSizeBytes,omitempty"`
		// The index of the disk, similar to a slot number on physical machine.
		Index *int32 `json:"Index,omitempty"`
		// Indicates whether the virtual disk is pinned to the hot tier or not.
		IsFlashModeEnabled *bool `json:"IsFlashModeEnabled,omitempty"`
		// The extId of the storage container which backs this disk.
		StorageContainerExtId *string `json:"StorageContainerExtId,omitempty"`
		// The unique identifier of the VM.
		VmExtId          *string                                     `json:"VmExtId,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		Vm               NullableHciEsxiVmRelationship               `json:"Vm,omitempty"`
	}

	varHciEsxiVmDiskWithoutEmbeddedStruct := HciEsxiVmDiskWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHciEsxiVmDiskWithoutEmbeddedStruct)
	if err == nil {
		varHciEsxiVmDisk := _HciEsxiVmDisk{}
		varHciEsxiVmDisk.ClassId = varHciEsxiVmDiskWithoutEmbeddedStruct.ClassId
		varHciEsxiVmDisk.ObjectType = varHciEsxiVmDiskWithoutEmbeddedStruct.ObjectType
		varHciEsxiVmDisk.BusType = varHciEsxiVmDiskWithoutEmbeddedStruct.BusType
		varHciEsxiVmDisk.DiskExtId = varHciEsxiVmDiskWithoutEmbeddedStruct.DiskExtId
		varHciEsxiVmDisk.DiskSizeBytes = varHciEsxiVmDiskWithoutEmbeddedStruct.DiskSizeBytes
		varHciEsxiVmDisk.Index = varHciEsxiVmDiskWithoutEmbeddedStruct.Index
		varHciEsxiVmDisk.IsFlashModeEnabled = varHciEsxiVmDiskWithoutEmbeddedStruct.IsFlashModeEnabled
		varHciEsxiVmDisk.StorageContainerExtId = varHciEsxiVmDiskWithoutEmbeddedStruct.StorageContainerExtId
		varHciEsxiVmDisk.VmExtId = varHciEsxiVmDiskWithoutEmbeddedStruct.VmExtId
		varHciEsxiVmDisk.RegisteredDevice = varHciEsxiVmDiskWithoutEmbeddedStruct.RegisteredDevice
		varHciEsxiVmDisk.Vm = varHciEsxiVmDiskWithoutEmbeddedStruct.Vm
		*o = HciEsxiVmDisk(varHciEsxiVmDisk)
	} else {
		return err
	}

	varHciEsxiVmDisk := _HciEsxiVmDisk{}

	err = json.Unmarshal(data, &varHciEsxiVmDisk)
	if err == nil {
		o.MoBaseMo = varHciEsxiVmDisk.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BusType")
		delete(additionalProperties, "DiskExtId")
		delete(additionalProperties, "DiskSizeBytes")
		delete(additionalProperties, "Index")
		delete(additionalProperties, "IsFlashModeEnabled")
		delete(additionalProperties, "StorageContainerExtId")
		delete(additionalProperties, "VmExtId")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Vm")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableHciEsxiVmDisk struct {
	value *HciEsxiVmDisk
	isSet bool
}

func (v NullableHciEsxiVmDisk) Get() *HciEsxiVmDisk {
	return v.value
}

func (v *NullableHciEsxiVmDisk) Set(val *HciEsxiVmDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableHciEsxiVmDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableHciEsxiVmDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHciEsxiVmDisk(val *HciEsxiVmDisk) *NullableHciEsxiVmDisk {
	return &NullableHciEsxiVmDisk{value: val, isSet: true}
}

func (v NullableHciEsxiVmDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHciEsxiVmDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
