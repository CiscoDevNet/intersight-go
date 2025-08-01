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

// checks if the StorageFlexFlashPhysicalDrive type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageFlexFlashPhysicalDrive{}

// StorageFlexFlashPhysicalDrive Physical Drive repersenting a SD Card.
type StorageFlexFlashPhysicalDrive struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The status of the flex flash physical drive.
	CardStatus *string `json:"CardStatus,omitempty"`
	// The card type of the flex flash physical drive.
	CardType *string `json:"CardType,omitempty"`
	// The OEM Identifier of the flex flash physical drive.
	OemId *string `json:"OemId,omitempty"`
	// The drive status of the flex flash physical drive.
	PdStatus                   *string                                        `json:"PdStatus,omitempty"`
	InventoryDeviceInfo        NullableInventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice           NullableAssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	StorageFlexFlashController NullableStorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _StorageFlexFlashPhysicalDrive StorageFlexFlashPhysicalDrive

// NewStorageFlexFlashPhysicalDrive instantiates a new StorageFlexFlashPhysicalDrive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexFlashPhysicalDrive(classId string, objectType string) *StorageFlexFlashPhysicalDrive {
	this := StorageFlexFlashPhysicalDrive{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageFlexFlashPhysicalDriveWithDefaults instantiates a new StorageFlexFlashPhysicalDrive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexFlashPhysicalDriveWithDefaults() *StorageFlexFlashPhysicalDrive {
	this := StorageFlexFlashPhysicalDrive{}
	var classId string = "storage.FlexFlashPhysicalDrive"
	this.ClassId = classId
	var objectType string = "storage.FlexFlashPhysicalDrive"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageFlexFlashPhysicalDrive) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageFlexFlashPhysicalDrive) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.FlexFlashPhysicalDrive" of the ClassId field.
func (o *StorageFlexFlashPhysicalDrive) GetDefaultClassId() interface{} {
	return "storage.FlexFlashPhysicalDrive"
}

// GetObjectType returns the ObjectType field value
func (o *StorageFlexFlashPhysicalDrive) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageFlexFlashPhysicalDrive) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.FlexFlashPhysicalDrive" of the ObjectType field.
func (o *StorageFlexFlashPhysicalDrive) GetDefaultObjectType() interface{} {
	return "storage.FlexFlashPhysicalDrive"
}

// GetCardStatus returns the CardStatus field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDrive) GetCardStatus() string {
	if o == nil || IsNil(o.CardStatus) {
		var ret string
		return ret
	}
	return *o.CardStatus
}

// GetCardStatusOk returns a tuple with the CardStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetCardStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CardStatus) {
		return nil, false
	}
	return o.CardStatus, true
}

// HasCardStatus returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasCardStatus() bool {
	if o != nil && !IsNil(o.CardStatus) {
		return true
	}

	return false
}

// SetCardStatus gets a reference to the given string and assigns it to the CardStatus field.
func (o *StorageFlexFlashPhysicalDrive) SetCardStatus(v string) {
	o.CardStatus = &v
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDrive) GetCardType() string {
	if o == nil || IsNil(o.CardType) {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CardType) {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasCardType() bool {
	if o != nil && !IsNil(o.CardType) {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *StorageFlexFlashPhysicalDrive) SetCardType(v string) {
	o.CardType = &v
}

// GetOemId returns the OemId field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDrive) GetOemId() string {
	if o == nil || IsNil(o.OemId) {
		var ret string
		return ret
	}
	return *o.OemId
}

// GetOemIdOk returns a tuple with the OemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetOemIdOk() (*string, bool) {
	if o == nil || IsNil(o.OemId) {
		return nil, false
	}
	return o.OemId, true
}

// HasOemId returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasOemId() bool {
	if o != nil && !IsNil(o.OemId) {
		return true
	}

	return false
}

// SetOemId gets a reference to the given string and assigns it to the OemId field.
func (o *StorageFlexFlashPhysicalDrive) SetOemId(v string) {
	o.OemId = &v
}

// GetPdStatus returns the PdStatus field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDrive) GetPdStatus() string {
	if o == nil || IsNil(o.PdStatus) {
		var ret string
		return ret
	}
	return *o.PdStatus
}

// GetPdStatusOk returns a tuple with the PdStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetPdStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PdStatus) {
		return nil, false
	}
	return o.PdStatus, true
}

// HasPdStatus returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasPdStatus() bool {
	if o != nil && !IsNil(o.PdStatus) {
		return true
	}

	return false
}

// SetPdStatus gets a reference to the given string and assigns it to the PdStatus field.
func (o *StorageFlexFlashPhysicalDrive) SetPdStatus(v string) {
	o.PdStatus = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashPhysicalDrive) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashPhysicalDrive) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexFlashPhysicalDrive) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *StorageFlexFlashPhysicalDrive) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *StorageFlexFlashPhysicalDrive) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashPhysicalDrive) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashPhysicalDrive) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexFlashPhysicalDrive) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageFlexFlashPhysicalDrive) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageFlexFlashPhysicalDrive) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetStorageFlexFlashController returns the StorageFlexFlashController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashPhysicalDrive) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship {
	if o == nil || IsNil(o.StorageFlexFlashController.Get()) {
		var ret StorageFlexFlashControllerRelationship
		return ret
	}
	return *o.StorageFlexFlashController.Get()
}

// GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashPhysicalDrive) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageFlexFlashController.Get(), o.StorageFlexFlashController.IsSet()
}

// HasStorageFlexFlashController returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasStorageFlexFlashController() bool {
	if o != nil && o.StorageFlexFlashController.IsSet() {
		return true
	}

	return false
}

// SetStorageFlexFlashController gets a reference to the given NullableStorageFlexFlashControllerRelationship and assigns it to the StorageFlexFlashController field.
func (o *StorageFlexFlashPhysicalDrive) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship) {
	o.StorageFlexFlashController.Set(&v)
}

// SetStorageFlexFlashControllerNil sets the value for StorageFlexFlashController to be an explicit nil
func (o *StorageFlexFlashPhysicalDrive) SetStorageFlexFlashControllerNil() {
	o.StorageFlexFlashController.Set(nil)
}

// UnsetStorageFlexFlashController ensures that no value is present for StorageFlexFlashController, not even an explicit nil
func (o *StorageFlexFlashPhysicalDrive) UnsetStorageFlexFlashController() {
	o.StorageFlexFlashController.Unset()
}

func (o StorageFlexFlashPhysicalDrive) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageFlexFlashPhysicalDrive) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.CardStatus) {
		toSerialize["CardStatus"] = o.CardStatus
	}
	if !IsNil(o.CardType) {
		toSerialize["CardType"] = o.CardType
	}
	if !IsNil(o.OemId) {
		toSerialize["OemId"] = o.OemId
	}
	if !IsNil(o.PdStatus) {
		toSerialize["PdStatus"] = o.PdStatus
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.StorageFlexFlashController.IsSet() {
		toSerialize["StorageFlexFlashController"] = o.StorageFlexFlashController.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageFlexFlashPhysicalDrive) UnmarshalJSON(data []byte) (err error) {
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
	type StorageFlexFlashPhysicalDriveWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The status of the flex flash physical drive.
		CardStatus *string `json:"CardStatus,omitempty"`
		// The card type of the flex flash physical drive.
		CardType *string `json:"CardType,omitempty"`
		// The OEM Identifier of the flex flash physical drive.
		OemId *string `json:"OemId,omitempty"`
		// The drive status of the flex flash physical drive.
		PdStatus                   *string                                        `json:"PdStatus,omitempty"`
		InventoryDeviceInfo        NullableInventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice           NullableAssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
		StorageFlexFlashController NullableStorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
	}

	varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct := StorageFlexFlashPhysicalDriveWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct)
	if err == nil {
		varStorageFlexFlashPhysicalDrive := _StorageFlexFlashPhysicalDrive{}
		varStorageFlexFlashPhysicalDrive.ClassId = varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct.ClassId
		varStorageFlexFlashPhysicalDrive.ObjectType = varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct.ObjectType
		varStorageFlexFlashPhysicalDrive.CardStatus = varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct.CardStatus
		varStorageFlexFlashPhysicalDrive.CardType = varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct.CardType
		varStorageFlexFlashPhysicalDrive.OemId = varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct.OemId
		varStorageFlexFlashPhysicalDrive.PdStatus = varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct.PdStatus
		varStorageFlexFlashPhysicalDrive.InventoryDeviceInfo = varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageFlexFlashPhysicalDrive.RegisteredDevice = varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct.RegisteredDevice
		varStorageFlexFlashPhysicalDrive.StorageFlexFlashController = varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct.StorageFlexFlashController
		*o = StorageFlexFlashPhysicalDrive(varStorageFlexFlashPhysicalDrive)
	} else {
		return err
	}

	varStorageFlexFlashPhysicalDrive := _StorageFlexFlashPhysicalDrive{}

	err = json.Unmarshal(data, &varStorageFlexFlashPhysicalDrive)
	if err == nil {
		o.EquipmentBase = varStorageFlexFlashPhysicalDrive.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CardStatus")
		delete(additionalProperties, "CardType")
		delete(additionalProperties, "OemId")
		delete(additionalProperties, "PdStatus")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageFlexFlashController")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableStorageFlexFlashPhysicalDrive struct {
	value *StorageFlexFlashPhysicalDrive
	isSet bool
}

func (v NullableStorageFlexFlashPhysicalDrive) Get() *StorageFlexFlashPhysicalDrive {
	return v.value
}

func (v *NullableStorageFlexFlashPhysicalDrive) Set(val *StorageFlexFlashPhysicalDrive) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexFlashPhysicalDrive) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexFlashPhysicalDrive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexFlashPhysicalDrive(val *StorageFlexFlashPhysicalDrive) *NullableStorageFlexFlashPhysicalDrive {
	return &NullableStorageFlexFlashPhysicalDrive{value: val, isSet: true}
}

func (v NullableStorageFlexFlashPhysicalDrive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexFlashPhysicalDrive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
