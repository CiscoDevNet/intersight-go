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

// checks if the TopSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TopSystem{}

// TopSystem Root container for all UCSM / CIMC MOs.
type TopSystem struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The IPv4 address of system.
	Ipv4Address *string `json:"Ipv4Address,omitempty"`
	// The IPv6 address of system.
	Ipv6Address *string `json:"Ipv6Address,omitempty"`
	// The current mode of the system.
	Mode *string `json:"Mode,omitempty"`
	// The admin configured name of the system.
	Name *string `json:"Name,omitempty"`
	// The operational timezone of the system, empty indicates no timezone has been set specifically.
	TimeZone *string `json:"TimeZone,omitempty"`
	// An array of relationships to computeBlade resources.
	ComputeBlades []ComputeBladeRelationship `json:"ComputeBlades,omitempty"`
	// An array of relationships to computeRackUnit resources.
	ComputeRackUnits     []ComputeRackUnitRelationship            `json:"ComputeRackUnits,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship  `json:"InventoryDeviceInfo,omitempty"`
	ManagementController NullableManagementControllerRelationship `json:"ManagementController,omitempty"`
	// An array of relationships to networkElement resources.
	NetworkElements      []NetworkElementRelationship                `json:"NetworkElements,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TopSystem TopSystem

// NewTopSystem instantiates a new TopSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopSystem(classId string, objectType string) *TopSystem {
	this := TopSystem{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTopSystemWithDefaults instantiates a new TopSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopSystemWithDefaults() *TopSystem {
	this := TopSystem{}
	var classId string = "top.System"
	this.ClassId = classId
	var objectType string = "top.System"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TopSystem) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TopSystem) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TopSystem) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "top.System" of the ClassId field.
func (o *TopSystem) GetDefaultClassId() interface{} {
	return "top.System"
}

// GetObjectType returns the ObjectType field value
func (o *TopSystem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TopSystem) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TopSystem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "top.System" of the ObjectType field.
func (o *TopSystem) GetDefaultObjectType() interface{} {
	return "top.System"
}

// GetIpv4Address returns the Ipv4Address field value if set, zero value otherwise.
func (o *TopSystem) GetIpv4Address() string {
	if o == nil || IsNil(o.Ipv4Address) {
		var ret string
		return ret
	}
	return *o.Ipv4Address
}

// GetIpv4AddressOk returns a tuple with the Ipv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystem) GetIpv4AddressOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Address) {
		return nil, false
	}
	return o.Ipv4Address, true
}

// HasIpv4Address returns a boolean if a field has been set.
func (o *TopSystem) HasIpv4Address() bool {
	if o != nil && !IsNil(o.Ipv4Address) {
		return true
	}

	return false
}

// SetIpv4Address gets a reference to the given string and assigns it to the Ipv4Address field.
func (o *TopSystem) SetIpv4Address(v string) {
	o.Ipv4Address = &v
}

// GetIpv6Address returns the Ipv6Address field value if set, zero value otherwise.
func (o *TopSystem) GetIpv6Address() string {
	if o == nil || IsNil(o.Ipv6Address) {
		var ret string
		return ret
	}
	return *o.Ipv6Address
}

// GetIpv6AddressOk returns a tuple with the Ipv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystem) GetIpv6AddressOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv6Address) {
		return nil, false
	}
	return o.Ipv6Address, true
}

// HasIpv6Address returns a boolean if a field has been set.
func (o *TopSystem) HasIpv6Address() bool {
	if o != nil && !IsNil(o.Ipv6Address) {
		return true
	}

	return false
}

// SetIpv6Address gets a reference to the given string and assigns it to the Ipv6Address field.
func (o *TopSystem) SetIpv6Address(v string) {
	o.Ipv6Address = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *TopSystem) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystem) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *TopSystem) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *TopSystem) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TopSystem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TopSystem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TopSystem) SetName(v string) {
	o.Name = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *TopSystem) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystem) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *TopSystem) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *TopSystem) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetComputeBlades returns the ComputeBlades field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopSystem) GetComputeBlades() []ComputeBladeRelationship {
	if o == nil {
		var ret []ComputeBladeRelationship
		return ret
	}
	return o.ComputeBlades
}

// GetComputeBladesOk returns a tuple with the ComputeBlades field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopSystem) GetComputeBladesOk() ([]ComputeBladeRelationship, bool) {
	if o == nil || IsNil(o.ComputeBlades) {
		return nil, false
	}
	return o.ComputeBlades, true
}

// HasComputeBlades returns a boolean if a field has been set.
func (o *TopSystem) HasComputeBlades() bool {
	if o != nil && !IsNil(o.ComputeBlades) {
		return true
	}

	return false
}

// SetComputeBlades gets a reference to the given []ComputeBladeRelationship and assigns it to the ComputeBlades field.
func (o *TopSystem) SetComputeBlades(v []ComputeBladeRelationship) {
	o.ComputeBlades = v
}

// GetComputeRackUnits returns the ComputeRackUnits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopSystem) GetComputeRackUnits() []ComputeRackUnitRelationship {
	if o == nil {
		var ret []ComputeRackUnitRelationship
		return ret
	}
	return o.ComputeRackUnits
}

// GetComputeRackUnitsOk returns a tuple with the ComputeRackUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopSystem) GetComputeRackUnitsOk() ([]ComputeRackUnitRelationship, bool) {
	if o == nil || IsNil(o.ComputeRackUnits) {
		return nil, false
	}
	return o.ComputeRackUnits, true
}

// HasComputeRackUnits returns a boolean if a field has been set.
func (o *TopSystem) HasComputeRackUnits() bool {
	if o != nil && !IsNil(o.ComputeRackUnits) {
		return true
	}

	return false
}

// SetComputeRackUnits gets a reference to the given []ComputeRackUnitRelationship and assigns it to the ComputeRackUnits field.
func (o *TopSystem) SetComputeRackUnits(v []ComputeRackUnitRelationship) {
	o.ComputeRackUnits = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopSystem) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopSystem) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *TopSystem) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *TopSystem) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *TopSystem) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *TopSystem) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetManagementController returns the ManagementController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopSystem) GetManagementController() ManagementControllerRelationship {
	if o == nil || IsNil(o.ManagementController.Get()) {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.ManagementController.Get()
}

// GetManagementControllerOk returns a tuple with the ManagementController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopSystem) GetManagementControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManagementController.Get(), o.ManagementController.IsSet()
}

// HasManagementController returns a boolean if a field has been set.
func (o *TopSystem) HasManagementController() bool {
	if o != nil && o.ManagementController.IsSet() {
		return true
	}

	return false
}

// SetManagementController gets a reference to the given NullableManagementControllerRelationship and assigns it to the ManagementController field.
func (o *TopSystem) SetManagementController(v ManagementControllerRelationship) {
	o.ManagementController.Set(&v)
}

// SetManagementControllerNil sets the value for ManagementController to be an explicit nil
func (o *TopSystem) SetManagementControllerNil() {
	o.ManagementController.Set(nil)
}

// UnsetManagementController ensures that no value is present for ManagementController, not even an explicit nil
func (o *TopSystem) UnsetManagementController() {
	o.ManagementController.Unset()
}

// GetNetworkElements returns the NetworkElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopSystem) GetNetworkElements() []NetworkElementRelationship {
	if o == nil {
		var ret []NetworkElementRelationship
		return ret
	}
	return o.NetworkElements
}

// GetNetworkElementsOk returns a tuple with the NetworkElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopSystem) GetNetworkElementsOk() ([]NetworkElementRelationship, bool) {
	if o == nil || IsNil(o.NetworkElements) {
		return nil, false
	}
	return o.NetworkElements, true
}

// HasNetworkElements returns a boolean if a field has been set.
func (o *TopSystem) HasNetworkElements() bool {
	if o != nil && !IsNil(o.NetworkElements) {
		return true
	}

	return false
}

// SetNetworkElements gets a reference to the given []NetworkElementRelationship and assigns it to the NetworkElements field.
func (o *TopSystem) SetNetworkElements(v []NetworkElementRelationship) {
	o.NetworkElements = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopSystem) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopSystem) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *TopSystem) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *TopSystem) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *TopSystem) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *TopSystem) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o TopSystem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopSystem) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Ipv4Address) {
		toSerialize["Ipv4Address"] = o.Ipv4Address
	}
	if !IsNil(o.Ipv6Address) {
		toSerialize["Ipv6Address"] = o.Ipv6Address
	}
	if !IsNil(o.Mode) {
		toSerialize["Mode"] = o.Mode
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.TimeZone) {
		toSerialize["TimeZone"] = o.TimeZone
	}
	if o.ComputeBlades != nil {
		toSerialize["ComputeBlades"] = o.ComputeBlades
	}
	if o.ComputeRackUnits != nil {
		toSerialize["ComputeRackUnits"] = o.ComputeRackUnits
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.ManagementController.IsSet() {
		toSerialize["ManagementController"] = o.ManagementController.Get()
	}
	if o.NetworkElements != nil {
		toSerialize["NetworkElements"] = o.NetworkElements
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TopSystem) UnmarshalJSON(data []byte) (err error) {
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
	type TopSystemWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The IPv4 address of system.
		Ipv4Address *string `json:"Ipv4Address,omitempty"`
		// The IPv6 address of system.
		Ipv6Address *string `json:"Ipv6Address,omitempty"`
		// The current mode of the system.
		Mode *string `json:"Mode,omitempty"`
		// The admin configured name of the system.
		Name *string `json:"Name,omitempty"`
		// The operational timezone of the system, empty indicates no timezone has been set specifically.
		TimeZone *string `json:"TimeZone,omitempty"`
		// An array of relationships to computeBlade resources.
		ComputeBlades []ComputeBladeRelationship `json:"ComputeBlades,omitempty"`
		// An array of relationships to computeRackUnit resources.
		ComputeRackUnits     []ComputeRackUnitRelationship            `json:"ComputeRackUnits,omitempty"`
		InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship  `json:"InventoryDeviceInfo,omitempty"`
		ManagementController NullableManagementControllerRelationship `json:"ManagementController,omitempty"`
		// An array of relationships to networkElement resources.
		NetworkElements  []NetworkElementRelationship                `json:"NetworkElements,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varTopSystemWithoutEmbeddedStruct := TopSystemWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varTopSystemWithoutEmbeddedStruct)
	if err == nil {
		varTopSystem := _TopSystem{}
		varTopSystem.ClassId = varTopSystemWithoutEmbeddedStruct.ClassId
		varTopSystem.ObjectType = varTopSystemWithoutEmbeddedStruct.ObjectType
		varTopSystem.Ipv4Address = varTopSystemWithoutEmbeddedStruct.Ipv4Address
		varTopSystem.Ipv6Address = varTopSystemWithoutEmbeddedStruct.Ipv6Address
		varTopSystem.Mode = varTopSystemWithoutEmbeddedStruct.Mode
		varTopSystem.Name = varTopSystemWithoutEmbeddedStruct.Name
		varTopSystem.TimeZone = varTopSystemWithoutEmbeddedStruct.TimeZone
		varTopSystem.ComputeBlades = varTopSystemWithoutEmbeddedStruct.ComputeBlades
		varTopSystem.ComputeRackUnits = varTopSystemWithoutEmbeddedStruct.ComputeRackUnits
		varTopSystem.InventoryDeviceInfo = varTopSystemWithoutEmbeddedStruct.InventoryDeviceInfo
		varTopSystem.ManagementController = varTopSystemWithoutEmbeddedStruct.ManagementController
		varTopSystem.NetworkElements = varTopSystemWithoutEmbeddedStruct.NetworkElements
		varTopSystem.RegisteredDevice = varTopSystemWithoutEmbeddedStruct.RegisteredDevice
		*o = TopSystem(varTopSystem)
	} else {
		return err
	}

	varTopSystem := _TopSystem{}

	err = json.Unmarshal(data, &varTopSystem)
	if err == nil {
		o.InventoryBase = varTopSystem.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Ipv4Address")
		delete(additionalProperties, "Ipv6Address")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "TimeZone")
		delete(additionalProperties, "ComputeBlades")
		delete(additionalProperties, "ComputeRackUnits")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "ManagementController")
		delete(additionalProperties, "NetworkElements")
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

type NullableTopSystem struct {
	value *TopSystem
	isSet bool
}

func (v NullableTopSystem) Get() *TopSystem {
	return v.value
}

func (v *NullableTopSystem) Set(val *TopSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableTopSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableTopSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopSystem(val *TopSystem) *NullableTopSystem {
	return &NullableTopSystem{value: val, isSet: true}
}

func (v NullableTopSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
