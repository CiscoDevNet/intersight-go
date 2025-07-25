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

// checks if the FabricMacSecEaPol type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricMacSecEaPol{}

// FabricMacSecEaPol Extensible authentication protocol over LAN (EAPoL). MACsec transmits MACsec key agreement (MKA) protocol data units (PDUs) using EAPoL packets to establish a secure session.
type FabricMacSecEaPol struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Ethertype to use in extensible authentication protocol over LAN (EAPoL) frames for MACsec key agreement (MKA) protocol data units (PDUs). The range is between 0x600 - 0xffff.
	EaPolEthertype *string `json:"EaPolEthertype,omitempty" validate:"regexp=^0x([6-9A-Fa-f][0-9A-Fa-f]{2}|[A-Fa-f0-9][0-9A-Fa-f]{3})$"`
	// MAC address to use in extensible authentication protocol over LAN (EAPoL) for MACsec key agreement (MKA) protocol data units (PDUs). EAPol mac address should not be equal to all-zero (0000.0000.0000).
	EaPolMacAddress      *string `json:"EaPolMacAddress,omitempty" validate:"regexp=^[0-9A-Fa-f]{4}\\\\.[0-9A-Fa-f]{4}\\\\.[0-9A-Fa-f]{4}$"`
	AdditionalProperties map[string]interface{}
}

type _FabricMacSecEaPol FabricMacSecEaPol

// NewFabricMacSecEaPol instantiates a new FabricMacSecEaPol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricMacSecEaPol(classId string, objectType string) *FabricMacSecEaPol {
	this := FabricMacSecEaPol{}
	this.ClassId = classId
	this.ObjectType = objectType
	var eaPolEthertype string = "34958"
	this.EaPolEthertype = &eaPolEthertype
	var eaPolMacAddress string = "0180.C200.0003"
	this.EaPolMacAddress = &eaPolMacAddress
	return &this
}

// NewFabricMacSecEaPolWithDefaults instantiates a new FabricMacSecEaPol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricMacSecEaPolWithDefaults() *FabricMacSecEaPol {
	this := FabricMacSecEaPol{}
	var classId string = "fabric.MacSecEaPol"
	this.ClassId = classId
	var objectType string = "fabric.MacSecEaPol"
	this.ObjectType = objectType
	var eaPolEthertype string = "34958"
	this.EaPolEthertype = &eaPolEthertype
	var eaPolMacAddress string = "0180.C200.0003"
	this.EaPolMacAddress = &eaPolMacAddress
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricMacSecEaPol) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricMacSecEaPol) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricMacSecEaPol) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.MacSecEaPol" of the ClassId field.
func (o *FabricMacSecEaPol) GetDefaultClassId() interface{} {
	return "fabric.MacSecEaPol"
}

// GetObjectType returns the ObjectType field value
func (o *FabricMacSecEaPol) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricMacSecEaPol) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricMacSecEaPol) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.MacSecEaPol" of the ObjectType field.
func (o *FabricMacSecEaPol) GetDefaultObjectType() interface{} {
	return "fabric.MacSecEaPol"
}

// GetEaPolEthertype returns the EaPolEthertype field value if set, zero value otherwise.
func (o *FabricMacSecEaPol) GetEaPolEthertype() string {
	if o == nil || IsNil(o.EaPolEthertype) {
		var ret string
		return ret
	}
	return *o.EaPolEthertype
}

// GetEaPolEthertypeOk returns a tuple with the EaPolEthertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMacSecEaPol) GetEaPolEthertypeOk() (*string, bool) {
	if o == nil || IsNil(o.EaPolEthertype) {
		return nil, false
	}
	return o.EaPolEthertype, true
}

// HasEaPolEthertype returns a boolean if a field has been set.
func (o *FabricMacSecEaPol) HasEaPolEthertype() bool {
	if o != nil && !IsNil(o.EaPolEthertype) {
		return true
	}

	return false
}

// SetEaPolEthertype gets a reference to the given string and assigns it to the EaPolEthertype field.
func (o *FabricMacSecEaPol) SetEaPolEthertype(v string) {
	o.EaPolEthertype = &v
}

// GetEaPolMacAddress returns the EaPolMacAddress field value if set, zero value otherwise.
func (o *FabricMacSecEaPol) GetEaPolMacAddress() string {
	if o == nil || IsNil(o.EaPolMacAddress) {
		var ret string
		return ret
	}
	return *o.EaPolMacAddress
}

// GetEaPolMacAddressOk returns a tuple with the EaPolMacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMacSecEaPol) GetEaPolMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EaPolMacAddress) {
		return nil, false
	}
	return o.EaPolMacAddress, true
}

// HasEaPolMacAddress returns a boolean if a field has been set.
func (o *FabricMacSecEaPol) HasEaPolMacAddress() bool {
	if o != nil && !IsNil(o.EaPolMacAddress) {
		return true
	}

	return false
}

// SetEaPolMacAddress gets a reference to the given string and assigns it to the EaPolMacAddress field.
func (o *FabricMacSecEaPol) SetEaPolMacAddress(v string) {
	o.EaPolMacAddress = &v
}

func (o FabricMacSecEaPol) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricMacSecEaPol) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.EaPolEthertype) {
		toSerialize["EaPolEthertype"] = o.EaPolEthertype
	}
	if !IsNil(o.EaPolMacAddress) {
		toSerialize["EaPolMacAddress"] = o.EaPolMacAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricMacSecEaPol) UnmarshalJSON(data []byte) (err error) {
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
	type FabricMacSecEaPolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Ethertype to use in extensible authentication protocol over LAN (EAPoL) frames for MACsec key agreement (MKA) protocol data units (PDUs). The range is between 0x600 - 0xffff.
		EaPolEthertype *string `json:"EaPolEthertype,omitempty" validate:"regexp=^0x([6-9A-Fa-f][0-9A-Fa-f]{2}|[A-Fa-f0-9][0-9A-Fa-f]{3})$"`
		// MAC address to use in extensible authentication protocol over LAN (EAPoL) for MACsec key agreement (MKA) protocol data units (PDUs). EAPol mac address should not be equal to all-zero (0000.0000.0000).
		EaPolMacAddress *string `json:"EaPolMacAddress,omitempty" validate:"regexp=^[0-9A-Fa-f]{4}\\\\.[0-9A-Fa-f]{4}\\\\.[0-9A-Fa-f]{4}$"`
	}

	varFabricMacSecEaPolWithoutEmbeddedStruct := FabricMacSecEaPolWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricMacSecEaPolWithoutEmbeddedStruct)
	if err == nil {
		varFabricMacSecEaPol := _FabricMacSecEaPol{}
		varFabricMacSecEaPol.ClassId = varFabricMacSecEaPolWithoutEmbeddedStruct.ClassId
		varFabricMacSecEaPol.ObjectType = varFabricMacSecEaPolWithoutEmbeddedStruct.ObjectType
		varFabricMacSecEaPol.EaPolEthertype = varFabricMacSecEaPolWithoutEmbeddedStruct.EaPolEthertype
		varFabricMacSecEaPol.EaPolMacAddress = varFabricMacSecEaPolWithoutEmbeddedStruct.EaPolMacAddress
		*o = FabricMacSecEaPol(varFabricMacSecEaPol)
	} else {
		return err
	}

	varFabricMacSecEaPol := _FabricMacSecEaPol{}

	err = json.Unmarshal(data, &varFabricMacSecEaPol)
	if err == nil {
		o.MoBaseComplexType = varFabricMacSecEaPol.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EaPolEthertype")
		delete(additionalProperties, "EaPolMacAddress")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableFabricMacSecEaPol struct {
	value *FabricMacSecEaPol
	isSet bool
}

func (v NullableFabricMacSecEaPol) Get() *FabricMacSecEaPol {
	return v.value
}

func (v *NullableFabricMacSecEaPol) Set(val *FabricMacSecEaPol) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricMacSecEaPol) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricMacSecEaPol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricMacSecEaPol(val *FabricMacSecEaPol) *NullableFabricMacSecEaPol {
	return &NullableFabricMacSecEaPol{value: val, isSet: true}
}

func (v NullableFabricMacSecEaPol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricMacSecEaPol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
