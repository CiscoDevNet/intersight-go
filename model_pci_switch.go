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

// checks if the PciSwitch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PciSwitch{}

// PciSwitch PCI Switch present in a server connected to two GPUs and one PCIe adapter.
type PciSwitch struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The device id of the switch.
	DeviceId *string `json:"DeviceId,omitempty"`
	// The composite health of the switch.
	Health *string `json:"Health,omitempty"`
	// The number of GPUs and PCI adapters connected the switch.
	NumOfAdaptors *string `json:"NumOfAdaptors,omitempty"`
	// The PCI address of the switch.
	PciAddress *string `json:"PciAddress,omitempty"`
	// The PCI slot name of the switch.
	PciSlot *string `json:"PciSlot,omitempty"`
	// The model information for the switch.
	ProductName *string `json:"ProductName,omitempty"`
	// The product revision of the switch.
	ProductRevision *string `json:"ProductRevision,omitempty"`
	// The sub device id of the switch.
	SubDeviceId *string `json:"SubDeviceId,omitempty"`
	// The sub vendor id of the switch.
	SubVendorId *string `json:"SubVendorId,omitempty"`
	// The current temperature of the switch.
	Temperature *string `json:"Temperature,omitempty"`
	// The type information of the switch.
	Type *string `json:"Type,omitempty"`
	// The vendor id of the switch.
	VendorId            *string                                 `json:"VendorId,omitempty"`
	ComputeBoard        NullableComputeBoardRelationship        `json:"ComputeBoard,omitempty"`
	InventoryDeviceInfo NullableInventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to pciLink resources.
	Links            []PciLinkRelationship                       `json:"Links,omitempty"`
	RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to firmwareRunningFirmware resources.
	RunningFirmware      []FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PciSwitch PciSwitch

// NewPciSwitch instantiates a new PciSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPciSwitch(classId string, objectType string) *PciSwitch {
	this := PciSwitch{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPciSwitchWithDefaults instantiates a new PciSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPciSwitchWithDefaults() *PciSwitch {
	this := PciSwitch{}
	var classId string = "pci.Switch"
	this.ClassId = classId
	var objectType string = "pci.Switch"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PciSwitch) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PciSwitch) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "pci.Switch" of the ClassId field.
func (o *PciSwitch) GetDefaultClassId() interface{} {
	return "pci.Switch"
}

// GetObjectType returns the ObjectType field value
func (o *PciSwitch) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PciSwitch) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "pci.Switch" of the ObjectType field.
func (o *PciSwitch) GetDefaultObjectType() interface{} {
	return "pci.Switch"
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *PciSwitch) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *PciSwitch) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *PciSwitch) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *PciSwitch) GetHealth() string {
	if o == nil || IsNil(o.Health) {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetHealthOk() (*string, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *PciSwitch) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *PciSwitch) SetHealth(v string) {
	o.Health = &v
}

// GetNumOfAdaptors returns the NumOfAdaptors field value if set, zero value otherwise.
func (o *PciSwitch) GetNumOfAdaptors() string {
	if o == nil || IsNil(o.NumOfAdaptors) {
		var ret string
		return ret
	}
	return *o.NumOfAdaptors
}

// GetNumOfAdaptorsOk returns a tuple with the NumOfAdaptors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetNumOfAdaptorsOk() (*string, bool) {
	if o == nil || IsNil(o.NumOfAdaptors) {
		return nil, false
	}
	return o.NumOfAdaptors, true
}

// HasNumOfAdaptors returns a boolean if a field has been set.
func (o *PciSwitch) HasNumOfAdaptors() bool {
	if o != nil && !IsNil(o.NumOfAdaptors) {
		return true
	}

	return false
}

// SetNumOfAdaptors gets a reference to the given string and assigns it to the NumOfAdaptors field.
func (o *PciSwitch) SetNumOfAdaptors(v string) {
	o.NumOfAdaptors = &v
}

// GetPciAddress returns the PciAddress field value if set, zero value otherwise.
func (o *PciSwitch) GetPciAddress() string {
	if o == nil || IsNil(o.PciAddress) {
		var ret string
		return ret
	}
	return *o.PciAddress
}

// GetPciAddressOk returns a tuple with the PciAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetPciAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PciAddress) {
		return nil, false
	}
	return o.PciAddress, true
}

// HasPciAddress returns a boolean if a field has been set.
func (o *PciSwitch) HasPciAddress() bool {
	if o != nil && !IsNil(o.PciAddress) {
		return true
	}

	return false
}

// SetPciAddress gets a reference to the given string and assigns it to the PciAddress field.
func (o *PciSwitch) SetPciAddress(v string) {
	o.PciAddress = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *PciSwitch) GetPciSlot() string {
	if o == nil || IsNil(o.PciSlot) {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetPciSlotOk() (*string, bool) {
	if o == nil || IsNil(o.PciSlot) {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *PciSwitch) HasPciSlot() bool {
	if o != nil && !IsNil(o.PciSlot) {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *PciSwitch) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *PciSwitch) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *PciSwitch) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *PciSwitch) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductRevision returns the ProductRevision field value if set, zero value otherwise.
func (o *PciSwitch) GetProductRevision() string {
	if o == nil || IsNil(o.ProductRevision) {
		var ret string
		return ret
	}
	return *o.ProductRevision
}

// GetProductRevisionOk returns a tuple with the ProductRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetProductRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductRevision) {
		return nil, false
	}
	return o.ProductRevision, true
}

// HasProductRevision returns a boolean if a field has been set.
func (o *PciSwitch) HasProductRevision() bool {
	if o != nil && !IsNil(o.ProductRevision) {
		return true
	}

	return false
}

// SetProductRevision gets a reference to the given string and assigns it to the ProductRevision field.
func (o *PciSwitch) SetProductRevision(v string) {
	o.ProductRevision = &v
}

// GetSubDeviceId returns the SubDeviceId field value if set, zero value otherwise.
func (o *PciSwitch) GetSubDeviceId() string {
	if o == nil || IsNil(o.SubDeviceId) {
		var ret string
		return ret
	}
	return *o.SubDeviceId
}

// GetSubDeviceIdOk returns a tuple with the SubDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetSubDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubDeviceId) {
		return nil, false
	}
	return o.SubDeviceId, true
}

// HasSubDeviceId returns a boolean if a field has been set.
func (o *PciSwitch) HasSubDeviceId() bool {
	if o != nil && !IsNil(o.SubDeviceId) {
		return true
	}

	return false
}

// SetSubDeviceId gets a reference to the given string and assigns it to the SubDeviceId field.
func (o *PciSwitch) SetSubDeviceId(v string) {
	o.SubDeviceId = &v
}

// GetSubVendorId returns the SubVendorId field value if set, zero value otherwise.
func (o *PciSwitch) GetSubVendorId() string {
	if o == nil || IsNil(o.SubVendorId) {
		var ret string
		return ret
	}
	return *o.SubVendorId
}

// GetSubVendorIdOk returns a tuple with the SubVendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetSubVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubVendorId) {
		return nil, false
	}
	return o.SubVendorId, true
}

// HasSubVendorId returns a boolean if a field has been set.
func (o *PciSwitch) HasSubVendorId() bool {
	if o != nil && !IsNil(o.SubVendorId) {
		return true
	}

	return false
}

// SetSubVendorId gets a reference to the given string and assigns it to the SubVendorId field.
func (o *PciSwitch) SetSubVendorId(v string) {
	o.SubVendorId = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *PciSwitch) GetTemperature() string {
	if o == nil || IsNil(o.Temperature) {
		var ret string
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetTemperatureOk() (*string, bool) {
	if o == nil || IsNil(o.Temperature) {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *PciSwitch) HasTemperature() bool {
	if o != nil && !IsNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given string and assigns it to the Temperature field.
func (o *PciSwitch) SetTemperature(v string) {
	o.Temperature = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PciSwitch) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PciSwitch) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PciSwitch) SetType(v string) {
	o.Type = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *PciSwitch) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *PciSwitch) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *PciSwitch) SetVendorId(v string) {
	o.VendorId = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciSwitch) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || IsNil(o.ComputeBoard.Get()) {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard.Get()
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciSwitch) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeBoard.Get(), o.ComputeBoard.IsSet()
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *PciSwitch) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard.IsSet() {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given NullableComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *PciSwitch) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard.Set(&v)
}

// SetComputeBoardNil sets the value for ComputeBoard to be an explicit nil
func (o *PciSwitch) SetComputeBoardNil() {
	o.ComputeBoard.Set(nil)
}

// UnsetComputeBoard ensures that no value is present for ComputeBoard, not even an explicit nil
func (o *PciSwitch) UnsetComputeBoard() {
	o.ComputeBoard.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciSwitch) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciSwitch) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *PciSwitch) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *PciSwitch) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *PciSwitch) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *PciSwitch) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciSwitch) GetLinks() []PciLinkRelationship {
	if o == nil {
		var ret []PciLinkRelationship
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciSwitch) GetLinksOk() ([]PciLinkRelationship, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PciSwitch) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []PciLinkRelationship and assigns it to the Links field.
func (o *PciSwitch) SetLinks(v []PciLinkRelationship) {
	o.Links = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciSwitch) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciSwitch) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PciSwitch) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PciSwitch) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *PciSwitch) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *PciSwitch) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetRunningFirmware returns the RunningFirmware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciSwitch) GetRunningFirmware() []FirmwareRunningFirmwareRelationship {
	if o == nil {
		var ret []FirmwareRunningFirmwareRelationship
		return ret
	}
	return o.RunningFirmware
}

// GetRunningFirmwareOk returns a tuple with the RunningFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciSwitch) GetRunningFirmwareOk() ([]FirmwareRunningFirmwareRelationship, bool) {
	if o == nil || IsNil(o.RunningFirmware) {
		return nil, false
	}
	return o.RunningFirmware, true
}

// HasRunningFirmware returns a boolean if a field has been set.
func (o *PciSwitch) HasRunningFirmware() bool {
	if o != nil && !IsNil(o.RunningFirmware) {
		return true
	}

	return false
}

// SetRunningFirmware gets a reference to the given []FirmwareRunningFirmwareRelationship and assigns it to the RunningFirmware field.
func (o *PciSwitch) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship) {
	o.RunningFirmware = v
}

func (o PciSwitch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PciSwitch) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DeviceId) {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if !IsNil(o.Health) {
		toSerialize["Health"] = o.Health
	}
	if !IsNil(o.NumOfAdaptors) {
		toSerialize["NumOfAdaptors"] = o.NumOfAdaptors
	}
	if !IsNil(o.PciAddress) {
		toSerialize["PciAddress"] = o.PciAddress
	}
	if !IsNil(o.PciSlot) {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if !IsNil(o.ProductName) {
		toSerialize["ProductName"] = o.ProductName
	}
	if !IsNil(o.ProductRevision) {
		toSerialize["ProductRevision"] = o.ProductRevision
	}
	if !IsNil(o.SubDeviceId) {
		toSerialize["SubDeviceId"] = o.SubDeviceId
	}
	if !IsNil(o.SubVendorId) {
		toSerialize["SubVendorId"] = o.SubVendorId
	}
	if !IsNil(o.Temperature) {
		toSerialize["Temperature"] = o.Temperature
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.VendorId) {
		toSerialize["VendorId"] = o.VendorId
	}
	if o.ComputeBoard.IsSet() {
		toSerialize["ComputeBoard"] = o.ComputeBoard.Get()
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.Links != nil {
		toSerialize["Links"] = o.Links
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.RunningFirmware != nil {
		toSerialize["RunningFirmware"] = o.RunningFirmware
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PciSwitch) UnmarshalJSON(data []byte) (err error) {
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
	type PciSwitchWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The device id of the switch.
		DeviceId *string `json:"DeviceId,omitempty"`
		// The composite health of the switch.
		Health *string `json:"Health,omitempty"`
		// The number of GPUs and PCI adapters connected the switch.
		NumOfAdaptors *string `json:"NumOfAdaptors,omitempty"`
		// The PCI address of the switch.
		PciAddress *string `json:"PciAddress,omitempty"`
		// The PCI slot name of the switch.
		PciSlot *string `json:"PciSlot,omitempty"`
		// The model information for the switch.
		ProductName *string `json:"ProductName,omitempty"`
		// The product revision of the switch.
		ProductRevision *string `json:"ProductRevision,omitempty"`
		// The sub device id of the switch.
		SubDeviceId *string `json:"SubDeviceId,omitempty"`
		// The sub vendor id of the switch.
		SubVendorId *string `json:"SubVendorId,omitempty"`
		// The current temperature of the switch.
		Temperature *string `json:"Temperature,omitempty"`
		// The type information of the switch.
		Type *string `json:"Type,omitempty"`
		// The vendor id of the switch.
		VendorId            *string                                 `json:"VendorId,omitempty"`
		ComputeBoard        NullableComputeBoardRelationship        `json:"ComputeBoard,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		// An array of relationships to pciLink resources.
		Links            []PciLinkRelationship                       `json:"Links,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		// An array of relationships to firmwareRunningFirmware resources.
		RunningFirmware []FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
	}

	varPciSwitchWithoutEmbeddedStruct := PciSwitchWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPciSwitchWithoutEmbeddedStruct)
	if err == nil {
		varPciSwitch := _PciSwitch{}
		varPciSwitch.ClassId = varPciSwitchWithoutEmbeddedStruct.ClassId
		varPciSwitch.ObjectType = varPciSwitchWithoutEmbeddedStruct.ObjectType
		varPciSwitch.DeviceId = varPciSwitchWithoutEmbeddedStruct.DeviceId
		varPciSwitch.Health = varPciSwitchWithoutEmbeddedStruct.Health
		varPciSwitch.NumOfAdaptors = varPciSwitchWithoutEmbeddedStruct.NumOfAdaptors
		varPciSwitch.PciAddress = varPciSwitchWithoutEmbeddedStruct.PciAddress
		varPciSwitch.PciSlot = varPciSwitchWithoutEmbeddedStruct.PciSlot
		varPciSwitch.ProductName = varPciSwitchWithoutEmbeddedStruct.ProductName
		varPciSwitch.ProductRevision = varPciSwitchWithoutEmbeddedStruct.ProductRevision
		varPciSwitch.SubDeviceId = varPciSwitchWithoutEmbeddedStruct.SubDeviceId
		varPciSwitch.SubVendorId = varPciSwitchWithoutEmbeddedStruct.SubVendorId
		varPciSwitch.Temperature = varPciSwitchWithoutEmbeddedStruct.Temperature
		varPciSwitch.Type = varPciSwitchWithoutEmbeddedStruct.Type
		varPciSwitch.VendorId = varPciSwitchWithoutEmbeddedStruct.VendorId
		varPciSwitch.ComputeBoard = varPciSwitchWithoutEmbeddedStruct.ComputeBoard
		varPciSwitch.InventoryDeviceInfo = varPciSwitchWithoutEmbeddedStruct.InventoryDeviceInfo
		varPciSwitch.Links = varPciSwitchWithoutEmbeddedStruct.Links
		varPciSwitch.RegisteredDevice = varPciSwitchWithoutEmbeddedStruct.RegisteredDevice
		varPciSwitch.RunningFirmware = varPciSwitchWithoutEmbeddedStruct.RunningFirmware
		*o = PciSwitch(varPciSwitch)
	} else {
		return err
	}

	varPciSwitch := _PciSwitch{}

	err = json.Unmarshal(data, &varPciSwitch)
	if err == nil {
		o.EquipmentBase = varPciSwitch.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "NumOfAdaptors")
		delete(additionalProperties, "PciAddress")
		delete(additionalProperties, "PciSlot")
		delete(additionalProperties, "ProductName")
		delete(additionalProperties, "ProductRevision")
		delete(additionalProperties, "SubDeviceId")
		delete(additionalProperties, "SubVendorId")
		delete(additionalProperties, "Temperature")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VendorId")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "Links")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "RunningFirmware")

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

type NullablePciSwitch struct {
	value *PciSwitch
	isSet bool
}

func (v NullablePciSwitch) Get() *PciSwitch {
	return v.value
}

func (v *NullablePciSwitch) Set(val *PciSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullablePciSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullablePciSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePciSwitch(val *PciSwitch) *NullablePciSwitch {
	return &NullablePciSwitch{value: val, isSet: true}
}

func (v NullablePciSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePciSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
