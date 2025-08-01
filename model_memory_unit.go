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

// checks if the MemoryUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemoryUnit{}

// MemoryUnit This represents a memory DIMM on a server.
type MemoryUnit struct {
	MemoryAbstractUnit
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field displays the description of the DIMM.
	Description *string `json:"Description,omitempty"`
	// This field indicates whether the DIMM is supported on the server or not.
	IsPlatformSupported *bool `json:"IsPlatformSupported,omitempty"`
	// This represents the ID of a regular DIMM on a server.
	// Deprecated
	MemoryId *int64 `json:"MemoryId,omitempty"`
	// This field displays the part number of the DIMM.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field displays the product ID of the DIMM.
	Pid                  *string                                     `json:"Pid,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	MemoryArray          NullableMemoryArrayRelationship             `json:"MemoryArray,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemoryUnit MemoryUnit

// NewMemoryUnit instantiates a new MemoryUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryUnit(classId string, objectType string) *MemoryUnit {
	this := MemoryUnit{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryUnitWithDefaults instantiates a new MemoryUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryUnitWithDefaults() *MemoryUnit {
	this := MemoryUnit{}
	var classId string = "memory.Unit"
	this.ClassId = classId
	var objectType string = "memory.Unit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryUnit) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryUnit) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryUnit) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "memory.Unit" of the ClassId field.
func (o *MemoryUnit) GetDefaultClassId() interface{} {
	return "memory.Unit"
}

// GetObjectType returns the ObjectType field value
func (o *MemoryUnit) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryUnit) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryUnit) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "memory.Unit" of the ObjectType field.
func (o *MemoryUnit) GetDefaultObjectType() interface{} {
	return "memory.Unit"
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MemoryUnit) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryUnit) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MemoryUnit) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MemoryUnit) SetDescription(v string) {
	o.Description = &v
}

// GetIsPlatformSupported returns the IsPlatformSupported field value if set, zero value otherwise.
func (o *MemoryUnit) GetIsPlatformSupported() bool {
	if o == nil || IsNil(o.IsPlatformSupported) {
		var ret bool
		return ret
	}
	return *o.IsPlatformSupported
}

// GetIsPlatformSupportedOk returns a tuple with the IsPlatformSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryUnit) GetIsPlatformSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPlatformSupported) {
		return nil, false
	}
	return o.IsPlatformSupported, true
}

// HasIsPlatformSupported returns a boolean if a field has been set.
func (o *MemoryUnit) HasIsPlatformSupported() bool {
	if o != nil && !IsNil(o.IsPlatformSupported) {
		return true
	}

	return false
}

// SetIsPlatformSupported gets a reference to the given bool and assigns it to the IsPlatformSupported field.
func (o *MemoryUnit) SetIsPlatformSupported(v bool) {
	o.IsPlatformSupported = &v
}

// GetMemoryId returns the MemoryId field value if set, zero value otherwise.
// Deprecated
func (o *MemoryUnit) GetMemoryId() int64 {
	if o == nil || IsNil(o.MemoryId) {
		var ret int64
		return ret
	}
	return *o.MemoryId
}

// GetMemoryIdOk returns a tuple with the MemoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *MemoryUnit) GetMemoryIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MemoryId) {
		return nil, false
	}
	return o.MemoryId, true
}

// HasMemoryId returns a boolean if a field has been set.
func (o *MemoryUnit) HasMemoryId() bool {
	if o != nil && !IsNil(o.MemoryId) {
		return true
	}

	return false
}

// SetMemoryId gets a reference to the given int64 and assigns it to the MemoryId field.
// Deprecated
func (o *MemoryUnit) SetMemoryId(v int64) {
	o.MemoryId = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *MemoryUnit) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryUnit) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *MemoryUnit) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *MemoryUnit) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *MemoryUnit) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryUnit) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *MemoryUnit) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *MemoryUnit) SetPid(v string) {
	o.Pid = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryUnit) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryUnit) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryUnit) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryUnit) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *MemoryUnit) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *MemoryUnit) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetMemoryArray returns the MemoryArray field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryUnit) GetMemoryArray() MemoryArrayRelationship {
	if o == nil || IsNil(o.MemoryArray.Get()) {
		var ret MemoryArrayRelationship
		return ret
	}
	return *o.MemoryArray.Get()
}

// GetMemoryArrayOk returns a tuple with the MemoryArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryUnit) GetMemoryArrayOk() (*MemoryArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryArray.Get(), o.MemoryArray.IsSet()
}

// HasMemoryArray returns a boolean if a field has been set.
func (o *MemoryUnit) HasMemoryArray() bool {
	if o != nil && o.MemoryArray.IsSet() {
		return true
	}

	return false
}

// SetMemoryArray gets a reference to the given NullableMemoryArrayRelationship and assigns it to the MemoryArray field.
func (o *MemoryUnit) SetMemoryArray(v MemoryArrayRelationship) {
	o.MemoryArray.Set(&v)
}

// SetMemoryArrayNil sets the value for MemoryArray to be an explicit nil
func (o *MemoryUnit) SetMemoryArrayNil() {
	o.MemoryArray.Set(nil)
}

// UnsetMemoryArray ensures that no value is present for MemoryArray, not even an explicit nil
func (o *MemoryUnit) UnsetMemoryArray() {
	o.MemoryArray.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryUnit) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *MemoryUnit) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *MemoryUnit) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o MemoryUnit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemoryUnit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMemoryAbstractUnit, errMemoryAbstractUnit := json.Marshal(o.MemoryAbstractUnit)
	if errMemoryAbstractUnit != nil {
		return map[string]interface{}{}, errMemoryAbstractUnit
	}
	errMemoryAbstractUnit = json.Unmarshal([]byte(serializedMemoryAbstractUnit), &toSerialize)
	if errMemoryAbstractUnit != nil {
		return map[string]interface{}{}, errMemoryAbstractUnit
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.IsPlatformSupported) {
		toSerialize["IsPlatformSupported"] = o.IsPlatformSupported
	}
	if !IsNil(o.MemoryId) {
		toSerialize["MemoryId"] = o.MemoryId
	}
	if !IsNil(o.PartNumber) {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if !IsNil(o.Pid) {
		toSerialize["Pid"] = o.Pid
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.MemoryArray.IsSet() {
		toSerialize["MemoryArray"] = o.MemoryArray.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemoryUnit) UnmarshalJSON(data []byte) (err error) {
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
	type MemoryUnitWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This field displays the description of the DIMM.
		Description *string `json:"Description,omitempty"`
		// This field indicates whether the DIMM is supported on the server or not.
		IsPlatformSupported *bool `json:"IsPlatformSupported,omitempty"`
		// This represents the ID of a regular DIMM on a server.
		// Deprecated
		MemoryId *int64 `json:"MemoryId,omitempty"`
		// This field displays the part number of the DIMM.
		PartNumber *string `json:"PartNumber,omitempty"`
		// This field displays the product ID of the DIMM.
		Pid                 *string                                     `json:"Pid,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		MemoryArray         NullableMemoryArrayRelationship             `json:"MemoryArray,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varMemoryUnitWithoutEmbeddedStruct := MemoryUnitWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMemoryUnitWithoutEmbeddedStruct)
	if err == nil {
		varMemoryUnit := _MemoryUnit{}
		varMemoryUnit.ClassId = varMemoryUnitWithoutEmbeddedStruct.ClassId
		varMemoryUnit.ObjectType = varMemoryUnitWithoutEmbeddedStruct.ObjectType
		varMemoryUnit.Description = varMemoryUnitWithoutEmbeddedStruct.Description
		varMemoryUnit.IsPlatformSupported = varMemoryUnitWithoutEmbeddedStruct.IsPlatformSupported
		varMemoryUnit.MemoryId = varMemoryUnitWithoutEmbeddedStruct.MemoryId
		varMemoryUnit.PartNumber = varMemoryUnitWithoutEmbeddedStruct.PartNumber
		varMemoryUnit.Pid = varMemoryUnitWithoutEmbeddedStruct.Pid
		varMemoryUnit.InventoryDeviceInfo = varMemoryUnitWithoutEmbeddedStruct.InventoryDeviceInfo
		varMemoryUnit.MemoryArray = varMemoryUnitWithoutEmbeddedStruct.MemoryArray
		varMemoryUnit.RegisteredDevice = varMemoryUnitWithoutEmbeddedStruct.RegisteredDevice
		*o = MemoryUnit(varMemoryUnit)
	} else {
		return err
	}

	varMemoryUnit := _MemoryUnit{}

	err = json.Unmarshal(data, &varMemoryUnit)
	if err == nil {
		o.MemoryAbstractUnit = varMemoryUnit.MemoryAbstractUnit
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "IsPlatformSupported")
		delete(additionalProperties, "MemoryId")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "MemoryArray")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMemoryAbstractUnit := reflect.ValueOf(o.MemoryAbstractUnit)
		for i := 0; i < reflectMemoryAbstractUnit.Type().NumField(); i++ {
			t := reflectMemoryAbstractUnit.Type().Field(i)

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

type NullableMemoryUnit struct {
	value *MemoryUnit
	isSet bool
}

func (v NullableMemoryUnit) Get() *MemoryUnit {
	return v.value
}

func (v *NullableMemoryUnit) Set(val *MemoryUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryUnit(val *MemoryUnit) *NullableMemoryUnit {
	return &NullableMemoryUnit{value: val, isSet: true}
}

func (v NullableMemoryUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
