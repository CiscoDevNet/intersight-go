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

// checks if the StorageItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageItem{}

// StorageItem The local Storage information.
type StorageItem struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The alarmType of the Local storage.
	AlarmType *string `json:"AlarmType,omitempty"`
	// The name of the Local storage.
	Name *string `json:"Name,omitempty"`
	// The operState of the Local storage.
	OperState *string `json:"OperState,omitempty"`
	// The size (MiB) of the Local storage.
	Size *string `json:"Size,omitempty"`
	// The used percent of the Local storage.
	Used                   *string                                     `json:"Used,omitempty"`
	InventoryDeviceInfo    NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	NetworkElement         NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice       NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StorageControllerDrive NullableStorageControllerDriveRelationship  `json:"StorageControllerDrive,omitempty"`
	// An array of relationships to storageFileItem resources.
	StorageFiles         []StorageFileItemRelationship `json:"StorageFiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageItem StorageItem

// NewStorageItem instantiates a new StorageItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageItem(classId string, objectType string) *StorageItem {
	this := StorageItem{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageItemWithDefaults instantiates a new StorageItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageItemWithDefaults() *StorageItem {
	this := StorageItem{}
	var classId string = "storage.Item"
	this.ClassId = classId
	var objectType string = "storage.Item"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageItem) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageItem) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageItem) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.Item" of the ClassId field.
func (o *StorageItem) GetDefaultClassId() interface{} {
	return "storage.Item"
}

// GetObjectType returns the ObjectType field value
func (o *StorageItem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageItem) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageItem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.Item" of the ObjectType field.
func (o *StorageItem) GetDefaultObjectType() interface{} {
	return "storage.Item"
}

// GetAlarmType returns the AlarmType field value if set, zero value otherwise.
func (o *StorageItem) GetAlarmType() string {
	if o == nil || IsNil(o.AlarmType) {
		var ret string
		return ret
	}
	return *o.AlarmType
}

// GetAlarmTypeOk returns a tuple with the AlarmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetAlarmTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AlarmType) {
		return nil, false
	}
	return o.AlarmType, true
}

// HasAlarmType returns a boolean if a field has been set.
func (o *StorageItem) HasAlarmType() bool {
	if o != nil && !IsNil(o.AlarmType) {
		return true
	}

	return false
}

// SetAlarmType gets a reference to the given string and assigns it to the AlarmType field.
func (o *StorageItem) SetAlarmType(v string) {
	o.AlarmType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageItem) SetName(v string) {
	o.Name = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *StorageItem) GetOperState() string {
	if o == nil || IsNil(o.OperState) {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetOperStateOk() (*string, bool) {
	if o == nil || IsNil(o.OperState) {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *StorageItem) HasOperState() bool {
	if o != nil && !IsNil(o.OperState) {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *StorageItem) SetOperState(v string) {
	o.OperState = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageItem) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageItem) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *StorageItem) SetSize(v string) {
	o.Size = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *StorageItem) GetUsed() string {
	if o == nil || IsNil(o.Used) {
		var ret string
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageItem) GetUsedOk() (*string, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *StorageItem) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given string and assigns it to the Used field.
func (o *StorageItem) SetUsed(v string) {
	o.Used = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageItem) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageItem) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageItem) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageItem) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *StorageItem) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *StorageItem) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageItem) GetNetworkElement() NetworkElementRelationship {
	if o == nil || IsNil(o.NetworkElement.Get()) {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement.Get()
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageItem) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkElement.Get(), o.NetworkElement.IsSet()
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *StorageItem) HasNetworkElement() bool {
	if o != nil && o.NetworkElement.IsSet() {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NullableNetworkElementRelationship and assigns it to the NetworkElement field.
func (o *StorageItem) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement.Set(&v)
}

// SetNetworkElementNil sets the value for NetworkElement to be an explicit nil
func (o *StorageItem) SetNetworkElementNil() {
	o.NetworkElement.Set(nil)
}

// UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
func (o *StorageItem) UnsetNetworkElement() {
	o.NetworkElement.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageItem) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageItem) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageItem) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageItem) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageItem) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageItem) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetStorageControllerDrive returns the StorageControllerDrive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageItem) GetStorageControllerDrive() StorageControllerDriveRelationship {
	if o == nil || IsNil(o.StorageControllerDrive.Get()) {
		var ret StorageControllerDriveRelationship
		return ret
	}
	return *o.StorageControllerDrive.Get()
}

// GetStorageControllerDriveOk returns a tuple with the StorageControllerDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageItem) GetStorageControllerDriveOk() (*StorageControllerDriveRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageControllerDrive.Get(), o.StorageControllerDrive.IsSet()
}

// HasStorageControllerDrive returns a boolean if a field has been set.
func (o *StorageItem) HasStorageControllerDrive() bool {
	if o != nil && o.StorageControllerDrive.IsSet() {
		return true
	}

	return false
}

// SetStorageControllerDrive gets a reference to the given NullableStorageControllerDriveRelationship and assigns it to the StorageControllerDrive field.
func (o *StorageItem) SetStorageControllerDrive(v StorageControllerDriveRelationship) {
	o.StorageControllerDrive.Set(&v)
}

// SetStorageControllerDriveNil sets the value for StorageControllerDrive to be an explicit nil
func (o *StorageItem) SetStorageControllerDriveNil() {
	o.StorageControllerDrive.Set(nil)
}

// UnsetStorageControllerDrive ensures that no value is present for StorageControllerDrive, not even an explicit nil
func (o *StorageItem) UnsetStorageControllerDrive() {
	o.StorageControllerDrive.Unset()
}

// GetStorageFiles returns the StorageFiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageItem) GetStorageFiles() []StorageFileItemRelationship {
	if o == nil {
		var ret []StorageFileItemRelationship
		return ret
	}
	return o.StorageFiles
}

// GetStorageFilesOk returns a tuple with the StorageFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageItem) GetStorageFilesOk() ([]StorageFileItemRelationship, bool) {
	if o == nil || IsNil(o.StorageFiles) {
		return nil, false
	}
	return o.StorageFiles, true
}

// HasStorageFiles returns a boolean if a field has been set.
func (o *StorageItem) HasStorageFiles() bool {
	if o != nil && !IsNil(o.StorageFiles) {
		return true
	}

	return false
}

// SetStorageFiles gets a reference to the given []StorageFileItemRelationship and assigns it to the StorageFiles field.
func (o *StorageItem) SetStorageFiles(v []StorageFileItemRelationship) {
	o.StorageFiles = v
}

func (o StorageItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AlarmType) {
		toSerialize["AlarmType"] = o.AlarmType
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.OperState) {
		toSerialize["OperState"] = o.OperState
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	if !IsNil(o.Used) {
		toSerialize["Used"] = o.Used
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.NetworkElement.IsSet() {
		toSerialize["NetworkElement"] = o.NetworkElement.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.StorageControllerDrive.IsSet() {
		toSerialize["StorageControllerDrive"] = o.StorageControllerDrive.Get()
	}
	if o.StorageFiles != nil {
		toSerialize["StorageFiles"] = o.StorageFiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageItem) UnmarshalJSON(data []byte) (err error) {
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
	type StorageItemWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The alarmType of the Local storage.
		AlarmType *string `json:"AlarmType,omitempty"`
		// The name of the Local storage.
		Name *string `json:"Name,omitempty"`
		// The operState of the Local storage.
		OperState *string `json:"OperState,omitempty"`
		// The size (MiB) of the Local storage.
		Size *string `json:"Size,omitempty"`
		// The used percent of the Local storage.
		Used                   *string                                     `json:"Used,omitempty"`
		InventoryDeviceInfo    NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		NetworkElement         NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice       NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		StorageControllerDrive NullableStorageControllerDriveRelationship  `json:"StorageControllerDrive,omitempty"`
		// An array of relationships to storageFileItem resources.
		StorageFiles []StorageFileItemRelationship `json:"StorageFiles,omitempty"`
	}

	varStorageItemWithoutEmbeddedStruct := StorageItemWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageItemWithoutEmbeddedStruct)
	if err == nil {
		varStorageItem := _StorageItem{}
		varStorageItem.ClassId = varStorageItemWithoutEmbeddedStruct.ClassId
		varStorageItem.ObjectType = varStorageItemWithoutEmbeddedStruct.ObjectType
		varStorageItem.AlarmType = varStorageItemWithoutEmbeddedStruct.AlarmType
		varStorageItem.Name = varStorageItemWithoutEmbeddedStruct.Name
		varStorageItem.OperState = varStorageItemWithoutEmbeddedStruct.OperState
		varStorageItem.Size = varStorageItemWithoutEmbeddedStruct.Size
		varStorageItem.Used = varStorageItemWithoutEmbeddedStruct.Used
		varStorageItem.InventoryDeviceInfo = varStorageItemWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageItem.NetworkElement = varStorageItemWithoutEmbeddedStruct.NetworkElement
		varStorageItem.RegisteredDevice = varStorageItemWithoutEmbeddedStruct.RegisteredDevice
		varStorageItem.StorageControllerDrive = varStorageItemWithoutEmbeddedStruct.StorageControllerDrive
		varStorageItem.StorageFiles = varStorageItemWithoutEmbeddedStruct.StorageFiles
		*o = StorageItem(varStorageItem)
	} else {
		return err
	}

	varStorageItem := _StorageItem{}

	err = json.Unmarshal(data, &varStorageItem)
	if err == nil {
		o.InventoryBase = varStorageItem.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlarmType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "Used")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageControllerDrive")
		delete(additionalProperties, "StorageFiles")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableStorageItem struct {
	value *StorageItem
	isSet bool
}

func (v NullableStorageItem) Get() *StorageItem {
	return v.value
}

func (v *NullableStorageItem) Set(val *StorageItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageItem(val *StorageItem) *NullableStorageItem {
	return &NullableStorageItem{value: val, isSet: true}
}

func (v NullableStorageItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
