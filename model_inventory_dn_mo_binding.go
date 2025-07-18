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

// checks if the InventoryDnMoBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryDnMoBinding{}

// InventoryDnMoBinding DnMoBinding provides a binding between a Intersight MO and a UCSM MO which has a DN.
type InventoryDnMoBinding struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Distinguished Name for this object, used to uniquely identify this object.
	Dn *string `json:"Dn,omitempty"`
	// The MO ID of the target MO for this particular Distinguished Name (dn).
	TargetMoId *string `json:"TargetMoId,omitempty"`
	// The type of the target MO for this particular Distinguished Name (dn).
	TargetMoType         *string                                     `json:"TargetMoType,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryDnMoBinding InventoryDnMoBinding

// NewInventoryDnMoBinding instantiates a new InventoryDnMoBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryDnMoBinding(classId string, objectType string) *InventoryDnMoBinding {
	this := InventoryDnMoBinding{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryDnMoBindingWithDefaults instantiates a new InventoryDnMoBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryDnMoBindingWithDefaults() *InventoryDnMoBinding {
	this := InventoryDnMoBinding{}
	var classId string = "inventory.DnMoBinding"
	this.ClassId = classId
	var objectType string = "inventory.DnMoBinding"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryDnMoBinding) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBinding) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryDnMoBinding) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "inventory.DnMoBinding" of the ClassId field.
func (o *InventoryDnMoBinding) GetDefaultClassId() interface{} {
	return "inventory.DnMoBinding"
}

// GetObjectType returns the ObjectType field value
func (o *InventoryDnMoBinding) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBinding) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryDnMoBinding) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "inventory.DnMoBinding" of the ObjectType field.
func (o *InventoryDnMoBinding) GetDefaultObjectType() interface{} {
	return "inventory.DnMoBinding"
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *InventoryDnMoBinding) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBinding) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *InventoryDnMoBinding) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *InventoryDnMoBinding) SetDn(v string) {
	o.Dn = &v
}

// GetTargetMoId returns the TargetMoId field value if set, zero value otherwise.
func (o *InventoryDnMoBinding) GetTargetMoId() string {
	if o == nil || IsNil(o.TargetMoId) {
		var ret string
		return ret
	}
	return *o.TargetMoId
}

// GetTargetMoIdOk returns a tuple with the TargetMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBinding) GetTargetMoIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetMoId) {
		return nil, false
	}
	return o.TargetMoId, true
}

// HasTargetMoId returns a boolean if a field has been set.
func (o *InventoryDnMoBinding) HasTargetMoId() bool {
	if o != nil && !IsNil(o.TargetMoId) {
		return true
	}

	return false
}

// SetTargetMoId gets a reference to the given string and assigns it to the TargetMoId field.
func (o *InventoryDnMoBinding) SetTargetMoId(v string) {
	o.TargetMoId = &v
}

// GetTargetMoType returns the TargetMoType field value if set, zero value otherwise.
func (o *InventoryDnMoBinding) GetTargetMoType() string {
	if o == nil || IsNil(o.TargetMoType) {
		var ret string
		return ret
	}
	return *o.TargetMoType
}

// GetTargetMoTypeOk returns a tuple with the TargetMoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBinding) GetTargetMoTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetMoType) {
		return nil, false
	}
	return o.TargetMoType, true
}

// HasTargetMoType returns a boolean if a field has been set.
func (o *InventoryDnMoBinding) HasTargetMoType() bool {
	if o != nil && !IsNil(o.TargetMoType) {
		return true
	}

	return false
}

// SetTargetMoType gets a reference to the given string and assigns it to the TargetMoType field.
func (o *InventoryDnMoBinding) SetTargetMoType(v string) {
	o.TargetMoType = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryDnMoBinding) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryDnMoBinding) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *InventoryDnMoBinding) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *InventoryDnMoBinding) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *InventoryDnMoBinding) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *InventoryDnMoBinding) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o InventoryDnMoBinding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryDnMoBinding) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if !IsNil(o.TargetMoId) {
		toSerialize["TargetMoId"] = o.TargetMoId
	}
	if !IsNil(o.TargetMoType) {
		toSerialize["TargetMoType"] = o.TargetMoType
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryDnMoBinding) UnmarshalJSON(data []byte) (err error) {
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
	type InventoryDnMoBindingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The Distinguished Name for this object, used to uniquely identify this object.
		Dn *string `json:"Dn,omitempty"`
		// The MO ID of the target MO for this particular Distinguished Name (dn).
		TargetMoId *string `json:"TargetMoId,omitempty"`
		// The type of the target MO for this particular Distinguished Name (dn).
		TargetMoType     *string                                     `json:"TargetMoType,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varInventoryDnMoBindingWithoutEmbeddedStruct := InventoryDnMoBindingWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varInventoryDnMoBindingWithoutEmbeddedStruct)
	if err == nil {
		varInventoryDnMoBinding := _InventoryDnMoBinding{}
		varInventoryDnMoBinding.ClassId = varInventoryDnMoBindingWithoutEmbeddedStruct.ClassId
		varInventoryDnMoBinding.ObjectType = varInventoryDnMoBindingWithoutEmbeddedStruct.ObjectType
		varInventoryDnMoBinding.Dn = varInventoryDnMoBindingWithoutEmbeddedStruct.Dn
		varInventoryDnMoBinding.TargetMoId = varInventoryDnMoBindingWithoutEmbeddedStruct.TargetMoId
		varInventoryDnMoBinding.TargetMoType = varInventoryDnMoBindingWithoutEmbeddedStruct.TargetMoType
		varInventoryDnMoBinding.RegisteredDevice = varInventoryDnMoBindingWithoutEmbeddedStruct.RegisteredDevice
		*o = InventoryDnMoBinding(varInventoryDnMoBinding)
	} else {
		return err
	}

	varInventoryDnMoBinding := _InventoryDnMoBinding{}

	err = json.Unmarshal(data, &varInventoryDnMoBinding)
	if err == nil {
		o.MoBaseMo = varInventoryDnMoBinding.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "TargetMoId")
		delete(additionalProperties, "TargetMoType")
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

type NullableInventoryDnMoBinding struct {
	value *InventoryDnMoBinding
	isSet bool
}

func (v NullableInventoryDnMoBinding) Get() *InventoryDnMoBinding {
	return v.value
}

func (v *NullableInventoryDnMoBinding) Set(val *InventoryDnMoBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryDnMoBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryDnMoBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryDnMoBinding(val *InventoryDnMoBinding) *NullableInventoryDnMoBinding {
	return &NullableInventoryDnMoBinding{value: val, isSet: true}
}

func (v NullableInventoryDnMoBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryDnMoBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
