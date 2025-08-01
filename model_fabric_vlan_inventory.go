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

// checks if the FabricVlanInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricVlanInventory{}

// FabricVlanInventory Configuration object inventoried from the device for VLAN.
type FabricVlanInventory struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin State of the VLAN attribute in the switch.
	AdminState *string `json:"AdminState,omitempty"`
	// Name of the VLAN Inventory for the switch.
	Name *string `json:"Name,omitempty"`
	// Ports associated to the VLAN.
	Port *string `json:"Port,omitempty"`
	// Status of the VLAN attribute in the switch.
	Status *string `json:"Status,omitempty"`
	// The Identifier for VLAN in the switch.
	VlanId               *string                                     `json:"VlanId,omitempty"`
	NetworkElement       NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricVlanInventory FabricVlanInventory

// NewFabricVlanInventory instantiates a new FabricVlanInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricVlanInventory(classId string, objectType string) *FabricVlanInventory {
	this := FabricVlanInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricVlanInventoryWithDefaults instantiates a new FabricVlanInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricVlanInventoryWithDefaults() *FabricVlanInventory {
	this := FabricVlanInventory{}
	var classId string = "fabric.VlanInventory"
	this.ClassId = classId
	var objectType string = "fabric.VlanInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricVlanInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricVlanInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricVlanInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.VlanInventory" of the ClassId field.
func (o *FabricVlanInventory) GetDefaultClassId() interface{} {
	return "fabric.VlanInventory"
}

// GetObjectType returns the ObjectType field value
func (o *FabricVlanInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricVlanInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricVlanInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.VlanInventory" of the ObjectType field.
func (o *FabricVlanInventory) GetDefaultObjectType() interface{} {
	return "fabric.VlanInventory"
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *FabricVlanInventory) GetAdminState() string {
	if o == nil || IsNil(o.AdminState) {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanInventory) GetAdminStateOk() (*string, bool) {
	if o == nil || IsNil(o.AdminState) {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *FabricVlanInventory) HasAdminState() bool {
	if o != nil && !IsNil(o.AdminState) {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *FabricVlanInventory) SetAdminState(v string) {
	o.AdminState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricVlanInventory) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanInventory) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricVlanInventory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricVlanInventory) SetName(v string) {
	o.Name = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *FabricVlanInventory) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanInventory) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *FabricVlanInventory) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *FabricVlanInventory) SetPort(v string) {
	o.Port = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FabricVlanInventory) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanInventory) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FabricVlanInventory) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FabricVlanInventory) SetStatus(v string) {
	o.Status = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *FabricVlanInventory) GetVlanId() string {
	if o == nil || IsNil(o.VlanId) {
		var ret string
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanInventory) GetVlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *FabricVlanInventory) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given string and assigns it to the VlanId field.
func (o *FabricVlanInventory) SetVlanId(v string) {
	o.VlanId = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricVlanInventory) GetNetworkElement() NetworkElementRelationship {
	if o == nil || IsNil(o.NetworkElement.Get()) {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement.Get()
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricVlanInventory) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkElement.Get(), o.NetworkElement.IsSet()
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *FabricVlanInventory) HasNetworkElement() bool {
	if o != nil && o.NetworkElement.IsSet() {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NullableNetworkElementRelationship and assigns it to the NetworkElement field.
func (o *FabricVlanInventory) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement.Set(&v)
}

// SetNetworkElementNil sets the value for NetworkElement to be an explicit nil
func (o *FabricVlanInventory) SetNetworkElementNil() {
	o.NetworkElement.Set(nil)
}

// UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
func (o *FabricVlanInventory) UnsetNetworkElement() {
	o.NetworkElement.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricVlanInventory) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricVlanInventory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *FabricVlanInventory) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *FabricVlanInventory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *FabricVlanInventory) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *FabricVlanInventory) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o FabricVlanInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricVlanInventory) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AdminState) {
		toSerialize["AdminState"] = o.AdminState
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Port) {
		toSerialize["Port"] = o.Port
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.VlanId) {
		toSerialize["VlanId"] = o.VlanId
	}
	if o.NetworkElement.IsSet() {
		toSerialize["NetworkElement"] = o.NetworkElement.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricVlanInventory) UnmarshalJSON(data []byte) (err error) {
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
	type FabricVlanInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin State of the VLAN attribute in the switch.
		AdminState *string `json:"AdminState,omitempty"`
		// Name of the VLAN Inventory for the switch.
		Name *string `json:"Name,omitempty"`
		// Ports associated to the VLAN.
		Port *string `json:"Port,omitempty"`
		// Status of the VLAN attribute in the switch.
		Status *string `json:"Status,omitempty"`
		// The Identifier for VLAN in the switch.
		VlanId           *string                                     `json:"VlanId,omitempty"`
		NetworkElement   NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varFabricVlanInventoryWithoutEmbeddedStruct := FabricVlanInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricVlanInventoryWithoutEmbeddedStruct)
	if err == nil {
		varFabricVlanInventory := _FabricVlanInventory{}
		varFabricVlanInventory.ClassId = varFabricVlanInventoryWithoutEmbeddedStruct.ClassId
		varFabricVlanInventory.ObjectType = varFabricVlanInventoryWithoutEmbeddedStruct.ObjectType
		varFabricVlanInventory.AdminState = varFabricVlanInventoryWithoutEmbeddedStruct.AdminState
		varFabricVlanInventory.Name = varFabricVlanInventoryWithoutEmbeddedStruct.Name
		varFabricVlanInventory.Port = varFabricVlanInventoryWithoutEmbeddedStruct.Port
		varFabricVlanInventory.Status = varFabricVlanInventoryWithoutEmbeddedStruct.Status
		varFabricVlanInventory.VlanId = varFabricVlanInventoryWithoutEmbeddedStruct.VlanId
		varFabricVlanInventory.NetworkElement = varFabricVlanInventoryWithoutEmbeddedStruct.NetworkElement
		varFabricVlanInventory.RegisteredDevice = varFabricVlanInventoryWithoutEmbeddedStruct.RegisteredDevice
		*o = FabricVlanInventory(varFabricVlanInventory)
	} else {
		return err
	}

	varFabricVlanInventory := _FabricVlanInventory{}

	err = json.Unmarshal(data, &varFabricVlanInventory)
	if err == nil {
		o.MoBaseMo = varFabricVlanInventory.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "VlanId")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableFabricVlanInventory struct {
	value *FabricVlanInventory
	isSet bool
}

func (v NullableFabricVlanInventory) Get() *FabricVlanInventory {
	return v.value
}

func (v *NullableFabricVlanInventory) Set(val *FabricVlanInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricVlanInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricVlanInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricVlanInventory(val *FabricVlanInventory) *NullableFabricVlanInventory {
	return &NullableFabricVlanInventory{value: val, isSet: true}
}

func (v NullableFabricVlanInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricVlanInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
