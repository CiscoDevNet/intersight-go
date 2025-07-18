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

// checks if the StorageControllerDrive type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageControllerDrive{}

// StorageControllerDrive The Local Storage present in a server.
type StorageControllerDrive struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The type of Local Storage like FlexMMC, USB Drive. * `Unknown` - Storage partition type is not known. * `FlexMMC` - The FlexMMC type of storage local drive. * `USB` - The USB type of storage drive.
	ControllerType *string `json:"ControllerType,omitempty"`
	// The description of local storage drive.
	Description *string `json:"Description,omitempty"`
	// Name of the local Storage.
	Name *string `json:"Name,omitempty"`
	// Total Partition count in local storage.
	PartitionCount *int64 `json:"PartitionCount,omitempty"`
	// The Id of the local Storage.
	StorageId *string `json:"StorageId,omitempty"`
	// The type of storage like internal or external. * `Unknown` - Not any of the known Storage Types. * `Internal` - The internal storage type. * `External` - The external storage type.
	Type             *string                                     `json:"Type,omitempty"`
	ComputeBoard     NullableComputeBoardRelationship            `json:"ComputeBoard,omitempty"`
	RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to storageItem resources.
	StorageItem          []StorageItemRelationship `json:"StorageItem,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageControllerDrive StorageControllerDrive

// NewStorageControllerDrive instantiates a new StorageControllerDrive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageControllerDrive(classId string, objectType string) *StorageControllerDrive {
	this := StorageControllerDrive{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageControllerDriveWithDefaults instantiates a new StorageControllerDrive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageControllerDriveWithDefaults() *StorageControllerDrive {
	this := StorageControllerDrive{}
	var classId string = "storage.ControllerDrive"
	this.ClassId = classId
	var objectType string = "storage.ControllerDrive"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageControllerDrive) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageControllerDrive) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageControllerDrive) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.ControllerDrive" of the ClassId field.
func (o *StorageControllerDrive) GetDefaultClassId() interface{} {
	return "storage.ControllerDrive"
}

// GetObjectType returns the ObjectType field value
func (o *StorageControllerDrive) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageControllerDrive) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageControllerDrive) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.ControllerDrive" of the ObjectType field.
func (o *StorageControllerDrive) GetDefaultObjectType() interface{} {
	return "storage.ControllerDrive"
}

// GetControllerType returns the ControllerType field value if set, zero value otherwise.
func (o *StorageControllerDrive) GetControllerType() string {
	if o == nil || IsNil(o.ControllerType) {
		var ret string
		return ret
	}
	return *o.ControllerType
}

// GetControllerTypeOk returns a tuple with the ControllerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageControllerDrive) GetControllerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerType) {
		return nil, false
	}
	return o.ControllerType, true
}

// HasControllerType returns a boolean if a field has been set.
func (o *StorageControllerDrive) HasControllerType() bool {
	if o != nil && !IsNil(o.ControllerType) {
		return true
	}

	return false
}

// SetControllerType gets a reference to the given string and assigns it to the ControllerType field.
func (o *StorageControllerDrive) SetControllerType(v string) {
	o.ControllerType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageControllerDrive) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageControllerDrive) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageControllerDrive) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageControllerDrive) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageControllerDrive) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageControllerDrive) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageControllerDrive) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageControllerDrive) SetName(v string) {
	o.Name = &v
}

// GetPartitionCount returns the PartitionCount field value if set, zero value otherwise.
func (o *StorageControllerDrive) GetPartitionCount() int64 {
	if o == nil || IsNil(o.PartitionCount) {
		var ret int64
		return ret
	}
	return *o.PartitionCount
}

// GetPartitionCountOk returns a tuple with the PartitionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageControllerDrive) GetPartitionCountOk() (*int64, bool) {
	if o == nil || IsNil(o.PartitionCount) {
		return nil, false
	}
	return o.PartitionCount, true
}

// HasPartitionCount returns a boolean if a field has been set.
func (o *StorageControllerDrive) HasPartitionCount() bool {
	if o != nil && !IsNil(o.PartitionCount) {
		return true
	}

	return false
}

// SetPartitionCount gets a reference to the given int64 and assigns it to the PartitionCount field.
func (o *StorageControllerDrive) SetPartitionCount(v int64) {
	o.PartitionCount = &v
}

// GetStorageId returns the StorageId field value if set, zero value otherwise.
func (o *StorageControllerDrive) GetStorageId() string {
	if o == nil || IsNil(o.StorageId) {
		var ret string
		return ret
	}
	return *o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageControllerDrive) GetStorageIdOk() (*string, bool) {
	if o == nil || IsNil(o.StorageId) {
		return nil, false
	}
	return o.StorageId, true
}

// HasStorageId returns a boolean if a field has been set.
func (o *StorageControllerDrive) HasStorageId() bool {
	if o != nil && !IsNil(o.StorageId) {
		return true
	}

	return false
}

// SetStorageId gets a reference to the given string and assigns it to the StorageId field.
func (o *StorageControllerDrive) SetStorageId(v string) {
	o.StorageId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageControllerDrive) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageControllerDrive) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageControllerDrive) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageControllerDrive) SetType(v string) {
	o.Type = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageControllerDrive) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || IsNil(o.ComputeBoard.Get()) {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard.Get()
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageControllerDrive) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeBoard.Get(), o.ComputeBoard.IsSet()
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *StorageControllerDrive) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard.IsSet() {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given NullableComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *StorageControllerDrive) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard.Set(&v)
}

// SetComputeBoardNil sets the value for ComputeBoard to be an explicit nil
func (o *StorageControllerDrive) SetComputeBoardNil() {
	o.ComputeBoard.Set(nil)
}

// UnsetComputeBoard ensures that no value is present for ComputeBoard, not even an explicit nil
func (o *StorageControllerDrive) UnsetComputeBoard() {
	o.ComputeBoard.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageControllerDrive) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageControllerDrive) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageControllerDrive) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageControllerDrive) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageControllerDrive) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageControllerDrive) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetStorageItem returns the StorageItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageControllerDrive) GetStorageItem() []StorageItemRelationship {
	if o == nil {
		var ret []StorageItemRelationship
		return ret
	}
	return o.StorageItem
}

// GetStorageItemOk returns a tuple with the StorageItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageControllerDrive) GetStorageItemOk() ([]StorageItemRelationship, bool) {
	if o == nil || IsNil(o.StorageItem) {
		return nil, false
	}
	return o.StorageItem, true
}

// HasStorageItem returns a boolean if a field has been set.
func (o *StorageControllerDrive) HasStorageItem() bool {
	if o != nil && !IsNil(o.StorageItem) {
		return true
	}

	return false
}

// SetStorageItem gets a reference to the given []StorageItemRelationship and assigns it to the StorageItem field.
func (o *StorageControllerDrive) SetStorageItem(v []StorageItemRelationship) {
	o.StorageItem = v
}

func (o StorageControllerDrive) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageControllerDrive) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ControllerType) {
		toSerialize["ControllerType"] = o.ControllerType
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PartitionCount) {
		toSerialize["PartitionCount"] = o.PartitionCount
	}
	if !IsNil(o.StorageId) {
		toSerialize["StorageId"] = o.StorageId
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.ComputeBoard.IsSet() {
		toSerialize["ComputeBoard"] = o.ComputeBoard.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.StorageItem != nil {
		toSerialize["StorageItem"] = o.StorageItem
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageControllerDrive) UnmarshalJSON(data []byte) (err error) {
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
	type StorageControllerDriveWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The type of Local Storage like FlexMMC, USB Drive. * `Unknown` - Storage partition type is not known. * `FlexMMC` - The FlexMMC type of storage local drive. * `USB` - The USB type of storage drive.
		ControllerType *string `json:"ControllerType,omitempty"`
		// The description of local storage drive.
		Description *string `json:"Description,omitempty"`
		// Name of the local Storage.
		Name *string `json:"Name,omitempty"`
		// Total Partition count in local storage.
		PartitionCount *int64 `json:"PartitionCount,omitempty"`
		// The Id of the local Storage.
		StorageId *string `json:"StorageId,omitempty"`
		// The type of storage like internal or external. * `Unknown` - Not any of the known Storage Types. * `Internal` - The internal storage type. * `External` - The external storage type.
		Type             *string                                     `json:"Type,omitempty"`
		ComputeBoard     NullableComputeBoardRelationship            `json:"ComputeBoard,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		// An array of relationships to storageItem resources.
		StorageItem []StorageItemRelationship `json:"StorageItem,omitempty"`
	}

	varStorageControllerDriveWithoutEmbeddedStruct := StorageControllerDriveWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageControllerDriveWithoutEmbeddedStruct)
	if err == nil {
		varStorageControllerDrive := _StorageControllerDrive{}
		varStorageControllerDrive.ClassId = varStorageControllerDriveWithoutEmbeddedStruct.ClassId
		varStorageControllerDrive.ObjectType = varStorageControllerDriveWithoutEmbeddedStruct.ObjectType
		varStorageControllerDrive.ControllerType = varStorageControllerDriveWithoutEmbeddedStruct.ControllerType
		varStorageControllerDrive.Description = varStorageControllerDriveWithoutEmbeddedStruct.Description
		varStorageControllerDrive.Name = varStorageControllerDriveWithoutEmbeddedStruct.Name
		varStorageControllerDrive.PartitionCount = varStorageControllerDriveWithoutEmbeddedStruct.PartitionCount
		varStorageControllerDrive.StorageId = varStorageControllerDriveWithoutEmbeddedStruct.StorageId
		varStorageControllerDrive.Type = varStorageControllerDriveWithoutEmbeddedStruct.Type
		varStorageControllerDrive.ComputeBoard = varStorageControllerDriveWithoutEmbeddedStruct.ComputeBoard
		varStorageControllerDrive.RegisteredDevice = varStorageControllerDriveWithoutEmbeddedStruct.RegisteredDevice
		varStorageControllerDrive.StorageItem = varStorageControllerDriveWithoutEmbeddedStruct.StorageItem
		*o = StorageControllerDrive(varStorageControllerDrive)
	} else {
		return err
	}

	varStorageControllerDrive := _StorageControllerDrive{}

	err = json.Unmarshal(data, &varStorageControllerDrive)
	if err == nil {
		o.InventoryBase = varStorageControllerDrive.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ControllerType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PartitionCount")
		delete(additionalProperties, "StorageId")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageItem")

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

type NullableStorageControllerDrive struct {
	value *StorageControllerDrive
	isSet bool
}

func (v NullableStorageControllerDrive) Get() *StorageControllerDrive {
	return v.value
}

func (v *NullableStorageControllerDrive) Set(val *StorageControllerDrive) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageControllerDrive) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageControllerDrive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageControllerDrive(val *StorageControllerDrive) *NullableStorageControllerDrive {
	return &NullableStorageControllerDrive{value: val, isSet: true}
}

func (v NullableStorageControllerDrive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageControllerDrive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
