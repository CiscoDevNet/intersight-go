/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-17769
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

// checks if the FmcDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FmcDevice{}

// FmcDevice Details of the FMC Devices.
type FmcDevice struct {
	FmcInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique Identifier of the Device.
	DeviceId *string `json:"DeviceId,omitempty"`
	// Unique Identifier of the Domain.
	DomainId *string `json:"DomainId,omitempty"`
	// A user provided name of the Device.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FmcDevice FmcDevice

// NewFmcDevice instantiates a new FmcDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFmcDevice(classId string, objectType string) *FmcDevice {
	this := FmcDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFmcDeviceWithDefaults instantiates a new FmcDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFmcDeviceWithDefaults() *FmcDevice {
	this := FmcDevice{}
	var classId string = "fmc.Device"
	this.ClassId = classId
	var objectType string = "fmc.Device"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FmcDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FmcDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FmcDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FmcDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FmcDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FmcDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *FmcDevice) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FmcDevice) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *FmcDevice) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *FmcDevice) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *FmcDevice) GetDomainId() string {
	if o == nil || IsNil(o.DomainId) {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FmcDevice) GetDomainIdOk() (*string, bool) {
	if o == nil || IsNil(o.DomainId) {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *FmcDevice) HasDomainId() bool {
	if o != nil && !IsNil(o.DomainId) {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *FmcDevice) SetDomainId(v string) {
	o.DomainId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FmcDevice) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FmcDevice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FmcDevice) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FmcDevice) SetName(v string) {
	o.Name = &v
}

func (o FmcDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FmcDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFmcInventoryEntity, errFmcInventoryEntity := json.Marshal(o.FmcInventoryEntity)
	if errFmcInventoryEntity != nil {
		return map[string]interface{}{}, errFmcInventoryEntity
	}
	errFmcInventoryEntity = json.Unmarshal([]byte(serializedFmcInventoryEntity), &toSerialize)
	if errFmcInventoryEntity != nil {
		return map[string]interface{}{}, errFmcInventoryEntity
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DeviceId) {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if !IsNil(o.DomainId) {
		toSerialize["DomainId"] = o.DomainId
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FmcDevice) UnmarshalJSON(data []byte) (err error) {
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

	type FmcDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Unique Identifier of the Device.
		DeviceId *string `json:"DeviceId,omitempty"`
		// Unique Identifier of the Domain.
		DomainId *string `json:"DomainId,omitempty"`
		// A user provided name of the Device.
		Name *string `json:"Name,omitempty"`
	}

	varFmcDeviceWithoutEmbeddedStruct := FmcDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFmcDeviceWithoutEmbeddedStruct)
	if err == nil {
		varFmcDevice := _FmcDevice{}
		varFmcDevice.ClassId = varFmcDeviceWithoutEmbeddedStruct.ClassId
		varFmcDevice.ObjectType = varFmcDeviceWithoutEmbeddedStruct.ObjectType
		varFmcDevice.DeviceId = varFmcDeviceWithoutEmbeddedStruct.DeviceId
		varFmcDevice.DomainId = varFmcDeviceWithoutEmbeddedStruct.DomainId
		varFmcDevice.Name = varFmcDeviceWithoutEmbeddedStruct.Name
		*o = FmcDevice(varFmcDevice)
	} else {
		return err
	}

	varFmcDevice := _FmcDevice{}

	err = json.Unmarshal(data, &varFmcDevice)
	if err == nil {
		o.FmcInventoryEntity = varFmcDevice.FmcInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "DomainId")
		delete(additionalProperties, "Name")

		// remove fields from embedded structs
		reflectFmcInventoryEntity := reflect.ValueOf(o.FmcInventoryEntity)
		for i := 0; i < reflectFmcInventoryEntity.Type().NumField(); i++ {
			t := reflectFmcInventoryEntity.Type().Field(i)

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

type NullableFmcDevice struct {
	value *FmcDevice
	isSet bool
}

func (v NullableFmcDevice) Get() *FmcDevice {
	return v.value
}

func (v *NullableFmcDevice) Set(val *FmcDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableFmcDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableFmcDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFmcDevice(val *FmcDevice) *NullableFmcDevice {
	return &NullableFmcDevice{value: val, isSet: true}
}

func (v NullableFmcDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFmcDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}