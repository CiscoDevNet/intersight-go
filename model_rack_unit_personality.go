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

// checks if the RackUnitPersonality type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackUnitPersonality{}

// RackUnitPersonality This can be used internally to model a server based on a defined personality without having to reprogram the server PID.
type RackUnitPersonality struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Additional info about the added software personality.
	AdditionalInfo *string `json:"AdditionalInfo,omitempty"`
	// Name of the software personality.
	Name *string `json:"Name,omitempty"`
	// Unique identity of added software personality.
	PersonalityId        *int64                                      `json:"PersonalityId,omitempty"`
	ComputeRackUnit      NullableComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RackUnitPersonality RackUnitPersonality

// NewRackUnitPersonality instantiates a new RackUnitPersonality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackUnitPersonality(classId string, objectType string) *RackUnitPersonality {
	this := RackUnitPersonality{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRackUnitPersonalityWithDefaults instantiates a new RackUnitPersonality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackUnitPersonalityWithDefaults() *RackUnitPersonality {
	this := RackUnitPersonality{}
	var classId string = "rack.UnitPersonality"
	this.ClassId = classId
	var objectType string = "rack.UnitPersonality"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RackUnitPersonality) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RackUnitPersonality) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RackUnitPersonality) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "rack.UnitPersonality" of the ClassId field.
func (o *RackUnitPersonality) GetDefaultClassId() interface{} {
	return "rack.UnitPersonality"
}

// GetObjectType returns the ObjectType field value
func (o *RackUnitPersonality) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RackUnitPersonality) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RackUnitPersonality) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "rack.UnitPersonality" of the ObjectType field.
func (o *RackUnitPersonality) GetDefaultObjectType() interface{} {
	return "rack.UnitPersonality"
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *RackUnitPersonality) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnitPersonality) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *RackUnitPersonality) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *RackUnitPersonality) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RackUnitPersonality) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnitPersonality) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RackUnitPersonality) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RackUnitPersonality) SetName(v string) {
	o.Name = &v
}

// GetPersonalityId returns the PersonalityId field value if set, zero value otherwise.
func (o *RackUnitPersonality) GetPersonalityId() int64 {
	if o == nil || IsNil(o.PersonalityId) {
		var ret int64
		return ret
	}
	return *o.PersonalityId
}

// GetPersonalityIdOk returns a tuple with the PersonalityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnitPersonality) GetPersonalityIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PersonalityId) {
		return nil, false
	}
	return o.PersonalityId, true
}

// HasPersonalityId returns a boolean if a field has been set.
func (o *RackUnitPersonality) HasPersonalityId() bool {
	if o != nil && !IsNil(o.PersonalityId) {
		return true
	}

	return false
}

// SetPersonalityId gets a reference to the given int64 and assigns it to the PersonalityId field.
func (o *RackUnitPersonality) SetPersonalityId(v int64) {
	o.PersonalityId = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackUnitPersonality) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || IsNil(o.ComputeRackUnit.Get()) {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit.Get()
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackUnitPersonality) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeRackUnit.Get(), o.ComputeRackUnit.IsSet()
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *RackUnitPersonality) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit.IsSet() {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given NullableComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *RackUnitPersonality) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit.Set(&v)
}

// SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil
func (o *RackUnitPersonality) SetComputeRackUnitNil() {
	o.ComputeRackUnit.Set(nil)
}

// UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
func (o *RackUnitPersonality) UnsetComputeRackUnit() {
	o.ComputeRackUnit.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackUnitPersonality) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackUnitPersonality) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *RackUnitPersonality) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *RackUnitPersonality) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *RackUnitPersonality) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *RackUnitPersonality) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackUnitPersonality) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackUnitPersonality) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *RackUnitPersonality) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *RackUnitPersonality) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *RackUnitPersonality) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *RackUnitPersonality) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o RackUnitPersonality) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackUnitPersonality) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AdditionalInfo) {
		toSerialize["AdditionalInfo"] = o.AdditionalInfo
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PersonalityId) {
		toSerialize["PersonalityId"] = o.PersonalityId
	}
	if o.ComputeRackUnit.IsSet() {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit.Get()
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RackUnitPersonality) UnmarshalJSON(data []byte) (err error) {
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
	type RackUnitPersonalityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Additional info about the added software personality.
		AdditionalInfo *string `json:"AdditionalInfo,omitempty"`
		// Name of the software personality.
		Name *string `json:"Name,omitempty"`
		// Unique identity of added software personality.
		PersonalityId       *int64                                      `json:"PersonalityId,omitempty"`
		ComputeRackUnit     NullableComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varRackUnitPersonalityWithoutEmbeddedStruct := RackUnitPersonalityWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varRackUnitPersonalityWithoutEmbeddedStruct)
	if err == nil {
		varRackUnitPersonality := _RackUnitPersonality{}
		varRackUnitPersonality.ClassId = varRackUnitPersonalityWithoutEmbeddedStruct.ClassId
		varRackUnitPersonality.ObjectType = varRackUnitPersonalityWithoutEmbeddedStruct.ObjectType
		varRackUnitPersonality.AdditionalInfo = varRackUnitPersonalityWithoutEmbeddedStruct.AdditionalInfo
		varRackUnitPersonality.Name = varRackUnitPersonalityWithoutEmbeddedStruct.Name
		varRackUnitPersonality.PersonalityId = varRackUnitPersonalityWithoutEmbeddedStruct.PersonalityId
		varRackUnitPersonality.ComputeRackUnit = varRackUnitPersonalityWithoutEmbeddedStruct.ComputeRackUnit
		varRackUnitPersonality.InventoryDeviceInfo = varRackUnitPersonalityWithoutEmbeddedStruct.InventoryDeviceInfo
		varRackUnitPersonality.RegisteredDevice = varRackUnitPersonalityWithoutEmbeddedStruct.RegisteredDevice
		*o = RackUnitPersonality(varRackUnitPersonality)
	} else {
		return err
	}

	varRackUnitPersonality := _RackUnitPersonality{}

	err = json.Unmarshal(data, &varRackUnitPersonality)
	if err == nil {
		o.InventoryBase = varRackUnitPersonality.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdditionalInfo")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PersonalityId")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
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

type NullableRackUnitPersonality struct {
	value *RackUnitPersonality
	isSet bool
}

func (v NullableRackUnitPersonality) Get() *RackUnitPersonality {
	return v.value
}

func (v *NullableRackUnitPersonality) Set(val *RackUnitPersonality) {
	v.value = val
	v.isSet = true
}

func (v NullableRackUnitPersonality) IsSet() bool {
	return v.isSet
}

func (v *NullableRackUnitPersonality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackUnitPersonality(val *RackUnitPersonality) *NullableRackUnitPersonality {
	return &NullableRackUnitPersonality{value: val, isSet: true}
}

func (v NullableRackUnitPersonality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackUnitPersonality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
