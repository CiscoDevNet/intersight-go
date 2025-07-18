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

// checks if the EquipmentAbstractDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentAbstractDevice{}

// EquipmentAbstractDevice Common attributes for inventory device in Intersight.
type EquipmentAbstractDevice struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType     string                          `json:"ObjectType"`
	DeviceLocation NullableEquipmentDeviceLocation `json:"DeviceLocation,omitempty"`
	// The categorization of the device type. Optional parameter to categorize devices by product type. For example, Meraki device types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, and secureConnect.
	DeviceType *string `json:"DeviceType,omitempty"`
	// The hardware version of the device.
	HardwareVersion *string `json:"HardwareVersion,omitempty"`
	// Administrator defined name for the device.
	Name *string `json:"Name,omitempty"`
	// Unique identity of the device.
	Uuid *string `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	// Current running software version of the device.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentAbstractDevice EquipmentAbstractDevice

// NewEquipmentAbstractDevice instantiates a new EquipmentAbstractDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentAbstractDevice(classId string, objectType string) *EquipmentAbstractDevice {
	this := EquipmentAbstractDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentAbstractDeviceWithDefaults instantiates a new EquipmentAbstractDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentAbstractDeviceWithDefaults() *EquipmentAbstractDevice {
	this := EquipmentAbstractDevice{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentAbstractDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentAbstractDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentAbstractDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentAbstractDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentAbstractDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentAbstractDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceLocation returns the DeviceLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentAbstractDevice) GetDeviceLocation() EquipmentDeviceLocation {
	if o == nil || IsNil(o.DeviceLocation.Get()) {
		var ret EquipmentDeviceLocation
		return ret
	}
	return *o.DeviceLocation.Get()
}

// GetDeviceLocationOk returns a tuple with the DeviceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentAbstractDevice) GetDeviceLocationOk() (*EquipmentDeviceLocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceLocation.Get(), o.DeviceLocation.IsSet()
}

// HasDeviceLocation returns a boolean if a field has been set.
func (o *EquipmentAbstractDevice) HasDeviceLocation() bool {
	if o != nil && o.DeviceLocation.IsSet() {
		return true
	}

	return false
}

// SetDeviceLocation gets a reference to the given NullableEquipmentDeviceLocation and assigns it to the DeviceLocation field.
func (o *EquipmentAbstractDevice) SetDeviceLocation(v EquipmentDeviceLocation) {
	o.DeviceLocation.Set(&v)
}

// SetDeviceLocationNil sets the value for DeviceLocation to be an explicit nil
func (o *EquipmentAbstractDevice) SetDeviceLocationNil() {
	o.DeviceLocation.Set(nil)
}

// UnsetDeviceLocation ensures that no value is present for DeviceLocation, not even an explicit nil
func (o *EquipmentAbstractDevice) UnsetDeviceLocation() {
	o.DeviceLocation.Unset()
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *EquipmentAbstractDevice) GetDeviceType() string {
	if o == nil || IsNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentAbstractDevice) GetDeviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *EquipmentAbstractDevice) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *EquipmentAbstractDevice) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetHardwareVersion returns the HardwareVersion field value if set, zero value otherwise.
func (o *EquipmentAbstractDevice) GetHardwareVersion() string {
	if o == nil || IsNil(o.HardwareVersion) {
		var ret string
		return ret
	}
	return *o.HardwareVersion
}

// GetHardwareVersionOk returns a tuple with the HardwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentAbstractDevice) GetHardwareVersionOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareVersion) {
		return nil, false
	}
	return o.HardwareVersion, true
}

// HasHardwareVersion returns a boolean if a field has been set.
func (o *EquipmentAbstractDevice) HasHardwareVersion() bool {
	if o != nil && !IsNil(o.HardwareVersion) {
		return true
	}

	return false
}

// SetHardwareVersion gets a reference to the given string and assigns it to the HardwareVersion field.
func (o *EquipmentAbstractDevice) SetHardwareVersion(v string) {
	o.HardwareVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EquipmentAbstractDevice) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentAbstractDevice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EquipmentAbstractDevice) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EquipmentAbstractDevice) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *EquipmentAbstractDevice) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentAbstractDevice) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *EquipmentAbstractDevice) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *EquipmentAbstractDevice) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EquipmentAbstractDevice) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentAbstractDevice) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EquipmentAbstractDevice) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EquipmentAbstractDevice) SetVersion(v string) {
	o.Version = &v
}

func (o EquipmentAbstractDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentAbstractDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if o.DeviceLocation.IsSet() {
		toSerialize["DeviceLocation"] = o.DeviceLocation.Get()
	}
	if !IsNil(o.DeviceType) {
		toSerialize["DeviceType"] = o.DeviceType
	}
	if !IsNil(o.HardwareVersion) {
		toSerialize["HardwareVersion"] = o.HardwareVersion
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentAbstractDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type EquipmentAbstractDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType     string                          `json:"ObjectType"`
		DeviceLocation NullableEquipmentDeviceLocation `json:"DeviceLocation,omitempty"`
		// The categorization of the device type. Optional parameter to categorize devices by product type. For example, Meraki device types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, and secureConnect.
		DeviceType *string `json:"DeviceType,omitempty"`
		// The hardware version of the device.
		HardwareVersion *string `json:"HardwareVersion,omitempty"`
		// Administrator defined name for the device.
		Name *string `json:"Name,omitempty"`
		// Unique identity of the device.
		Uuid *string `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
		// Current running software version of the device.
		Version *string `json:"Version,omitempty"`
	}

	varEquipmentAbstractDeviceWithoutEmbeddedStruct := EquipmentAbstractDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentAbstractDeviceWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentAbstractDevice := _EquipmentAbstractDevice{}
		varEquipmentAbstractDevice.ClassId = varEquipmentAbstractDeviceWithoutEmbeddedStruct.ClassId
		varEquipmentAbstractDevice.ObjectType = varEquipmentAbstractDeviceWithoutEmbeddedStruct.ObjectType
		varEquipmentAbstractDevice.DeviceLocation = varEquipmentAbstractDeviceWithoutEmbeddedStruct.DeviceLocation
		varEquipmentAbstractDevice.DeviceType = varEquipmentAbstractDeviceWithoutEmbeddedStruct.DeviceType
		varEquipmentAbstractDevice.HardwareVersion = varEquipmentAbstractDeviceWithoutEmbeddedStruct.HardwareVersion
		varEquipmentAbstractDevice.Name = varEquipmentAbstractDeviceWithoutEmbeddedStruct.Name
		varEquipmentAbstractDevice.Uuid = varEquipmentAbstractDeviceWithoutEmbeddedStruct.Uuid
		varEquipmentAbstractDevice.Version = varEquipmentAbstractDeviceWithoutEmbeddedStruct.Version
		*o = EquipmentAbstractDevice(varEquipmentAbstractDevice)
	} else {
		return err
	}

	varEquipmentAbstractDevice := _EquipmentAbstractDevice{}

	err = json.Unmarshal(data, &varEquipmentAbstractDevice)
	if err == nil {
		o.EquipmentBase = varEquipmentAbstractDevice.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceLocation")
		delete(additionalProperties, "DeviceType")
		delete(additionalProperties, "HardwareVersion")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Version")

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

type NullableEquipmentAbstractDevice struct {
	value *EquipmentAbstractDevice
	isSet bool
}

func (v NullableEquipmentAbstractDevice) Get() *EquipmentAbstractDevice {
	return v.value
}

func (v *NullableEquipmentAbstractDevice) Set(val *EquipmentAbstractDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentAbstractDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentAbstractDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentAbstractDevice(val *EquipmentAbstractDevice) *NullableEquipmentAbstractDevice {
	return &NullableEquipmentAbstractDevice{value: val, isSet: true}
}

func (v NullableEquipmentAbstractDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentAbstractDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
