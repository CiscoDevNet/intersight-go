/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-15473
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

// checks if the NetworkVfc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkVfc{}

// NetworkVfc Vfc configured on a Fabric Interconnect.
type NetworkVfc struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description for the vHBA.
	Description *string `json:"Description,omitempty"`
	// Vfc Identifier on a Fabric Interconnect.
	VfcId                  *int64                                      `json:"VfcId,omitempty"`
	AdapterHostFcInterface NullableAdapterHostFcInterfaceRelationship  `json:"AdapterHostFcInterface,omitempty"`
	NetworkElement         NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice       NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _NetworkVfc NetworkVfc

// NewNetworkVfc instantiates a new NetworkVfc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkVfc(classId string, objectType string) *NetworkVfc {
	this := NetworkVfc{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkVfcWithDefaults instantiates a new NetworkVfc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkVfcWithDefaults() *NetworkVfc {
	this := NetworkVfc{}
	var classId string = "network.Vfc"
	this.ClassId = classId
	var objectType string = "network.Vfc"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkVfc) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkVfc) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkVfc) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkVfc) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkVfc) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkVfc) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworkVfc) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVfc) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkVfc) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworkVfc) SetDescription(v string) {
	o.Description = &v
}

// GetVfcId returns the VfcId field value if set, zero value otherwise.
func (o *NetworkVfc) GetVfcId() int64 {
	if o == nil || IsNil(o.VfcId) {
		var ret int64
		return ret
	}
	return *o.VfcId
}

// GetVfcIdOk returns a tuple with the VfcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVfc) GetVfcIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VfcId) {
		return nil, false
	}
	return o.VfcId, true
}

// HasVfcId returns a boolean if a field has been set.
func (o *NetworkVfc) HasVfcId() bool {
	if o != nil && !IsNil(o.VfcId) {
		return true
	}

	return false
}

// SetVfcId gets a reference to the given int64 and assigns it to the VfcId field.
func (o *NetworkVfc) SetVfcId(v int64) {
	o.VfcId = &v
}

// GetAdapterHostFcInterface returns the AdapterHostFcInterface field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkVfc) GetAdapterHostFcInterface() AdapterHostFcInterfaceRelationship {
	if o == nil || IsNil(o.AdapterHostFcInterface.Get()) {
		var ret AdapterHostFcInterfaceRelationship
		return ret
	}
	return *o.AdapterHostFcInterface.Get()
}

// GetAdapterHostFcInterfaceOk returns a tuple with the AdapterHostFcInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkVfc) GetAdapterHostFcInterfaceOk() (*AdapterHostFcInterfaceRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdapterHostFcInterface.Get(), o.AdapterHostFcInterface.IsSet()
}

// HasAdapterHostFcInterface returns a boolean if a field has been set.
func (o *NetworkVfc) HasAdapterHostFcInterface() bool {
	if o != nil && o.AdapterHostFcInterface.IsSet() {
		return true
	}

	return false
}

// SetAdapterHostFcInterface gets a reference to the given NullableAdapterHostFcInterfaceRelationship and assigns it to the AdapterHostFcInterface field.
func (o *NetworkVfc) SetAdapterHostFcInterface(v AdapterHostFcInterfaceRelationship) {
	o.AdapterHostFcInterface.Set(&v)
}

// SetAdapterHostFcInterfaceNil sets the value for AdapterHostFcInterface to be an explicit nil
func (o *NetworkVfc) SetAdapterHostFcInterfaceNil() {
	o.AdapterHostFcInterface.Set(nil)
}

// UnsetAdapterHostFcInterface ensures that no value is present for AdapterHostFcInterface, not even an explicit nil
func (o *NetworkVfc) UnsetAdapterHostFcInterface() {
	o.AdapterHostFcInterface.Unset()
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkVfc) GetNetworkElement() NetworkElementRelationship {
	if o == nil || IsNil(o.NetworkElement.Get()) {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement.Get()
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkVfc) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkElement.Get(), o.NetworkElement.IsSet()
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkVfc) HasNetworkElement() bool {
	if o != nil && o.NetworkElement.IsSet() {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NullableNetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkVfc) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement.Set(&v)
}

// SetNetworkElementNil sets the value for NetworkElement to be an explicit nil
func (o *NetworkVfc) SetNetworkElementNil() {
	o.NetworkElement.Set(nil)
}

// UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
func (o *NetworkVfc) UnsetNetworkElement() {
	o.NetworkElement.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkVfc) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkVfc) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkVfc) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkVfc) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NetworkVfc) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NetworkVfc) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NetworkVfc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkVfc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.VfcId) {
		toSerialize["VfcId"] = o.VfcId
	}
	if o.AdapterHostFcInterface.IsSet() {
		toSerialize["AdapterHostFcInterface"] = o.AdapterHostFcInterface.Get()
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

func (o *NetworkVfc) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type NetworkVfcWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description for the vHBA.
		Description *string `json:"Description,omitempty"`
		// Vfc Identifier on a Fabric Interconnect.
		VfcId                  *int64                                      `json:"VfcId,omitempty"`
		AdapterHostFcInterface NullableAdapterHostFcInterfaceRelationship  `json:"AdapterHostFcInterface,omitempty"`
		NetworkElement         NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice       NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNetworkVfcWithoutEmbeddedStruct := NetworkVfcWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNetworkVfcWithoutEmbeddedStruct)
	if err == nil {
		varNetworkVfc := _NetworkVfc{}
		varNetworkVfc.ClassId = varNetworkVfcWithoutEmbeddedStruct.ClassId
		varNetworkVfc.ObjectType = varNetworkVfcWithoutEmbeddedStruct.ObjectType
		varNetworkVfc.Description = varNetworkVfcWithoutEmbeddedStruct.Description
		varNetworkVfc.VfcId = varNetworkVfcWithoutEmbeddedStruct.VfcId
		varNetworkVfc.AdapterHostFcInterface = varNetworkVfcWithoutEmbeddedStruct.AdapterHostFcInterface
		varNetworkVfc.NetworkElement = varNetworkVfcWithoutEmbeddedStruct.NetworkElement
		varNetworkVfc.RegisteredDevice = varNetworkVfcWithoutEmbeddedStruct.RegisteredDevice
		*o = NetworkVfc(varNetworkVfc)
	} else {
		return err
	}

	varNetworkVfc := _NetworkVfc{}

	err = json.Unmarshal(data, &varNetworkVfc)
	if err == nil {
		o.InventoryBase = varNetworkVfc.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "VfcId")
		delete(additionalProperties, "AdapterHostFcInterface")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNetworkVfc struct {
	value *NetworkVfc
	isSet bool
}

func (v NullableNetworkVfc) Get() *NetworkVfc {
	return v.value
}

func (v *NullableNetworkVfc) Set(val *NetworkVfc) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkVfc) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkVfc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkVfc(val *NetworkVfc) *NullableNetworkVfc {
	return &NullableNetworkVfc{value: val, isSet: true}
}

func (v NullableNetworkVfc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkVfc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
