/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

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
	PdStatus                   *string                                 `json:"PdStatus,omitempty"`
	InventoryDeviceInfo        *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice           *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	StorageFlexFlashController *StorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
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

// GetCardStatus returns the CardStatus field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDrive) GetCardStatus() string {
	if o == nil || o.CardStatus == nil {
		var ret string
		return ret
	}
	return *o.CardStatus
}

// GetCardStatusOk returns a tuple with the CardStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetCardStatusOk() (*string, bool) {
	if o == nil || o.CardStatus == nil {
		return nil, false
	}
	return o.CardStatus, true
}

// HasCardStatus returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasCardStatus() bool {
	if o != nil && o.CardStatus != nil {
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
	if o == nil || o.CardType == nil {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetCardTypeOk() (*string, bool) {
	if o == nil || o.CardType == nil {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasCardType() bool {
	if o != nil && o.CardType != nil {
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
	if o == nil || o.OemId == nil {
		var ret string
		return ret
	}
	return *o.OemId
}

// GetOemIdOk returns a tuple with the OemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetOemIdOk() (*string, bool) {
	if o == nil || o.OemId == nil {
		return nil, false
	}
	return o.OemId, true
}

// HasOemId returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasOemId() bool {
	if o != nil && o.OemId != nil {
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
	if o == nil || o.PdStatus == nil {
		var ret string
		return ret
	}
	return *o.PdStatus
}

// GetPdStatusOk returns a tuple with the PdStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetPdStatusOk() (*string, bool) {
	if o == nil || o.PdStatus == nil {
		return nil, false
	}
	return o.PdStatus, true
}

// HasPdStatus returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasPdStatus() bool {
	if o != nil && o.PdStatus != nil {
		return true
	}

	return false
}

// SetPdStatus gets a reference to the given string and assigns it to the PdStatus field.
func (o *StorageFlexFlashPhysicalDrive) SetPdStatus(v string) {
	o.PdStatus = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDrive) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexFlashPhysicalDrive) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDrive) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexFlashPhysicalDrive) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageFlexFlashController returns the StorageFlexFlashController field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDrive) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship {
	if o == nil || o.StorageFlexFlashController == nil {
		var ret StorageFlexFlashControllerRelationship
		return ret
	}
	return *o.StorageFlexFlashController
}

// GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDrive) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool) {
	if o == nil || o.StorageFlexFlashController == nil {
		return nil, false
	}
	return o.StorageFlexFlashController, true
}

// HasStorageFlexFlashController returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDrive) HasStorageFlexFlashController() bool {
	if o != nil && o.StorageFlexFlashController != nil {
		return true
	}

	return false
}

// SetStorageFlexFlashController gets a reference to the given StorageFlexFlashControllerRelationship and assigns it to the StorageFlexFlashController field.
func (o *StorageFlexFlashPhysicalDrive) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship) {
	o.StorageFlexFlashController = &v
}

func (o StorageFlexFlashPhysicalDrive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CardStatus != nil {
		toSerialize["CardStatus"] = o.CardStatus
	}
	if o.CardType != nil {
		toSerialize["CardType"] = o.CardType
	}
	if o.OemId != nil {
		toSerialize["OemId"] = o.OemId
	}
	if o.PdStatus != nil {
		toSerialize["PdStatus"] = o.PdStatus
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageFlexFlashController != nil {
		toSerialize["StorageFlexFlashController"] = o.StorageFlexFlashController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageFlexFlashPhysicalDrive) UnmarshalJSON(bytes []byte) (err error) {
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
		PdStatus                   *string                                 `json:"PdStatus,omitempty"`
		InventoryDeviceInfo        *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice           *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
		StorageFlexFlashController *StorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
	}

	varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct := StorageFlexFlashPhysicalDriveWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageFlexFlashPhysicalDriveWithoutEmbeddedStruct)
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

	err = json.Unmarshal(bytes, &varStorageFlexFlashPhysicalDrive)
	if err == nil {
		o.EquipmentBase = varStorageFlexFlashPhysicalDrive.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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