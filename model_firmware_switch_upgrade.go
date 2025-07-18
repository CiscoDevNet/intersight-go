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

// checks if the FirmwareSwitchUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareSwitchUpgrade{}

// FirmwareSwitchUpgrade Firmware upgrade operation for Fabric Interconnects that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
type FirmwareSwitchUpgrade struct {
	FirmwareUpgradeBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The flag to enable or disable fabric evacuation during the switch firmware upgrade. In case of IMM, it is mandatory to have the Fabric Interconnects associated with domain profile for fabric evacuation to happen.
	EnableFabricEvacuation *bool `json:"EnableFabricEvacuation,omitempty"`
	// The flag to enable or disable the option to wait for IO paths connectivity during the switch firmware upgrade.
	SkipWaitForIoPathConnectivity *bool                                       `json:"SkipWaitForIoPathConnectivity,omitempty"`
	Device                        NullableAssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	// An array of relationships to networkElement resources.
	NetworkElements      []NetworkElementRelationship `json:"NetworkElements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareSwitchUpgrade FirmwareSwitchUpgrade

// NewFirmwareSwitchUpgrade instantiates a new FirmwareSwitchUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareSwitchUpgrade(classId string, objectType string) *FirmwareSwitchUpgrade {
	this := FirmwareSwitchUpgrade{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "NONE"
	this.Status = &status
	var upgradeType string = "direct_upgrade"
	this.UpgradeType = &upgradeType
	var enableFabricEvacuation bool = true
	this.EnableFabricEvacuation = &enableFabricEvacuation
	var skipWaitForIoPathConnectivity bool = false
	this.SkipWaitForIoPathConnectivity = &skipWaitForIoPathConnectivity
	return &this
}

// NewFirmwareSwitchUpgradeWithDefaults instantiates a new FirmwareSwitchUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareSwitchUpgradeWithDefaults() *FirmwareSwitchUpgrade {
	this := FirmwareSwitchUpgrade{}
	var classId string = "firmware.SwitchUpgrade"
	this.ClassId = classId
	var objectType string = "firmware.SwitchUpgrade"
	this.ObjectType = objectType
	var enableFabricEvacuation bool = true
	this.EnableFabricEvacuation = &enableFabricEvacuation
	var skipWaitForIoPathConnectivity bool = false
	this.SkipWaitForIoPathConnectivity = &skipWaitForIoPathConnectivity
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareSwitchUpgrade) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareSwitchUpgrade) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareSwitchUpgrade) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "firmware.SwitchUpgrade" of the ClassId field.
func (o *FirmwareSwitchUpgrade) GetDefaultClassId() interface{} {
	return "firmware.SwitchUpgrade"
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareSwitchUpgrade) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareSwitchUpgrade) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareSwitchUpgrade) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "firmware.SwitchUpgrade" of the ObjectType field.
func (o *FirmwareSwitchUpgrade) GetDefaultObjectType() interface{} {
	return "firmware.SwitchUpgrade"
}

// GetEnableFabricEvacuation returns the EnableFabricEvacuation field value if set, zero value otherwise.
func (o *FirmwareSwitchUpgrade) GetEnableFabricEvacuation() bool {
	if o == nil || IsNil(o.EnableFabricEvacuation) {
		var ret bool
		return ret
	}
	return *o.EnableFabricEvacuation
}

// GetEnableFabricEvacuationOk returns a tuple with the EnableFabricEvacuation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareSwitchUpgrade) GetEnableFabricEvacuationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableFabricEvacuation) {
		return nil, false
	}
	return o.EnableFabricEvacuation, true
}

// HasEnableFabricEvacuation returns a boolean if a field has been set.
func (o *FirmwareSwitchUpgrade) HasEnableFabricEvacuation() bool {
	if o != nil && !IsNil(o.EnableFabricEvacuation) {
		return true
	}

	return false
}

// SetEnableFabricEvacuation gets a reference to the given bool and assigns it to the EnableFabricEvacuation field.
func (o *FirmwareSwitchUpgrade) SetEnableFabricEvacuation(v bool) {
	o.EnableFabricEvacuation = &v
}

// GetSkipWaitForIoPathConnectivity returns the SkipWaitForIoPathConnectivity field value if set, zero value otherwise.
func (o *FirmwareSwitchUpgrade) GetSkipWaitForIoPathConnectivity() bool {
	if o == nil || IsNil(o.SkipWaitForIoPathConnectivity) {
		var ret bool
		return ret
	}
	return *o.SkipWaitForIoPathConnectivity
}

// GetSkipWaitForIoPathConnectivityOk returns a tuple with the SkipWaitForIoPathConnectivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareSwitchUpgrade) GetSkipWaitForIoPathConnectivityOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipWaitForIoPathConnectivity) {
		return nil, false
	}
	return o.SkipWaitForIoPathConnectivity, true
}

// HasSkipWaitForIoPathConnectivity returns a boolean if a field has been set.
func (o *FirmwareSwitchUpgrade) HasSkipWaitForIoPathConnectivity() bool {
	if o != nil && !IsNil(o.SkipWaitForIoPathConnectivity) {
		return true
	}

	return false
}

// SetSkipWaitForIoPathConnectivity gets a reference to the given bool and assigns it to the SkipWaitForIoPathConnectivity field.
func (o *FirmwareSwitchUpgrade) SetSkipWaitForIoPathConnectivity(v bool) {
	o.SkipWaitForIoPathConnectivity = &v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareSwitchUpgrade) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.Device.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareSwitchUpgrade) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *FirmwareSwitchUpgrade) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *FirmwareSwitchUpgrade) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *FirmwareSwitchUpgrade) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *FirmwareSwitchUpgrade) UnsetDevice() {
	o.Device.Unset()
}

// GetNetworkElements returns the NetworkElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareSwitchUpgrade) GetNetworkElements() []NetworkElementRelationship {
	if o == nil {
		var ret []NetworkElementRelationship
		return ret
	}
	return o.NetworkElements
}

// GetNetworkElementsOk returns a tuple with the NetworkElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareSwitchUpgrade) GetNetworkElementsOk() ([]NetworkElementRelationship, bool) {
	if o == nil || IsNil(o.NetworkElements) {
		return nil, false
	}
	return o.NetworkElements, true
}

// HasNetworkElements returns a boolean if a field has been set.
func (o *FirmwareSwitchUpgrade) HasNetworkElements() bool {
	if o != nil && !IsNil(o.NetworkElements) {
		return true
	}

	return false
}

// SetNetworkElements gets a reference to the given []NetworkElementRelationship and assigns it to the NetworkElements field.
func (o *FirmwareSwitchUpgrade) SetNetworkElements(v []NetworkElementRelationship) {
	o.NetworkElements = v
}

func (o FirmwareSwitchUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareSwitchUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareUpgradeBase, errFirmwareUpgradeBase := json.Marshal(o.FirmwareUpgradeBase)
	if errFirmwareUpgradeBase != nil {
		return map[string]interface{}{}, errFirmwareUpgradeBase
	}
	errFirmwareUpgradeBase = json.Unmarshal([]byte(serializedFirmwareUpgradeBase), &toSerialize)
	if errFirmwareUpgradeBase != nil {
		return map[string]interface{}{}, errFirmwareUpgradeBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.EnableFabricEvacuation) {
		toSerialize["EnableFabricEvacuation"] = o.EnableFabricEvacuation
	}
	if !IsNil(o.SkipWaitForIoPathConnectivity) {
		toSerialize["SkipWaitForIoPathConnectivity"] = o.SkipWaitForIoPathConnectivity
	}
	if o.Device.IsSet() {
		toSerialize["Device"] = o.Device.Get()
	}
	if o.NetworkElements != nil {
		toSerialize["NetworkElements"] = o.NetworkElements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirmwareSwitchUpgrade) UnmarshalJSON(data []byte) (err error) {
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
	type FirmwareSwitchUpgradeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The flag to enable or disable fabric evacuation during the switch firmware upgrade. In case of IMM, it is mandatory to have the Fabric Interconnects associated with domain profile for fabric evacuation to happen.
		EnableFabricEvacuation *bool `json:"EnableFabricEvacuation,omitempty"`
		// The flag to enable or disable the option to wait for IO paths connectivity during the switch firmware upgrade.
		SkipWaitForIoPathConnectivity *bool                                       `json:"SkipWaitForIoPathConnectivity,omitempty"`
		Device                        NullableAssetDeviceRegistrationRelationship `json:"Device,omitempty"`
		// An array of relationships to networkElement resources.
		NetworkElements []NetworkElementRelationship `json:"NetworkElements,omitempty"`
	}

	varFirmwareSwitchUpgradeWithoutEmbeddedStruct := FirmwareSwitchUpgradeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFirmwareSwitchUpgradeWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareSwitchUpgrade := _FirmwareSwitchUpgrade{}
		varFirmwareSwitchUpgrade.ClassId = varFirmwareSwitchUpgradeWithoutEmbeddedStruct.ClassId
		varFirmwareSwitchUpgrade.ObjectType = varFirmwareSwitchUpgradeWithoutEmbeddedStruct.ObjectType
		varFirmwareSwitchUpgrade.EnableFabricEvacuation = varFirmwareSwitchUpgradeWithoutEmbeddedStruct.EnableFabricEvacuation
		varFirmwareSwitchUpgrade.SkipWaitForIoPathConnectivity = varFirmwareSwitchUpgradeWithoutEmbeddedStruct.SkipWaitForIoPathConnectivity
		varFirmwareSwitchUpgrade.Device = varFirmwareSwitchUpgradeWithoutEmbeddedStruct.Device
		varFirmwareSwitchUpgrade.NetworkElements = varFirmwareSwitchUpgradeWithoutEmbeddedStruct.NetworkElements
		*o = FirmwareSwitchUpgrade(varFirmwareSwitchUpgrade)
	} else {
		return err
	}

	varFirmwareSwitchUpgrade := _FirmwareSwitchUpgrade{}

	err = json.Unmarshal(data, &varFirmwareSwitchUpgrade)
	if err == nil {
		o.FirmwareUpgradeBase = varFirmwareSwitchUpgrade.FirmwareUpgradeBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnableFabricEvacuation")
		delete(additionalProperties, "SkipWaitForIoPathConnectivity")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "NetworkElements")

		// remove fields from embedded structs
		reflectFirmwareUpgradeBase := reflect.ValueOf(o.FirmwareUpgradeBase)
		for i := 0; i < reflectFirmwareUpgradeBase.Type().NumField(); i++ {
			t := reflectFirmwareUpgradeBase.Type().Field(i)

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

type NullableFirmwareSwitchUpgrade struct {
	value *FirmwareSwitchUpgrade
	isSet bool
}

func (v NullableFirmwareSwitchUpgrade) Get() *FirmwareSwitchUpgrade {
	return v.value
}

func (v *NullableFirmwareSwitchUpgrade) Set(val *FirmwareSwitchUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareSwitchUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareSwitchUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareSwitchUpgrade(val *FirmwareSwitchUpgrade) *NullableFirmwareSwitchUpgrade {
	return &NullableFirmwareSwitchUpgrade{value: val, isSet: true}
}

func (v NullableFirmwareSwitchUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareSwitchUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
