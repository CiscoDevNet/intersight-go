/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2024112619
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

// checks if the HciIpAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HciIpAddress{}

// HciIpAddress An IP address that can be either IPv4 or IPv6.
type HciIpAddress struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An IPv4 address in this IpAddress.
	Ipv4Address *string `json:"Ipv4Address,omitempty"`
	// The prefix length of the IPv4 address.
	Ipv4PrefixLength *int32 `json:"Ipv4PrefixLength,omitempty"`
	// An IPv6 address in this IpAddress.
	Ipv6Address *string `json:"Ipv6Address,omitempty"`
	// The prefix length of the IPv6 address.
	Ipv6PrefixLength     *int32 `json:"Ipv6PrefixLength,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HciIpAddress HciIpAddress

// NewHciIpAddress instantiates a new HciIpAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHciIpAddress(classId string, objectType string) *HciIpAddress {
	this := HciIpAddress{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHciIpAddressWithDefaults instantiates a new HciIpAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHciIpAddressWithDefaults() *HciIpAddress {
	this := HciIpAddress{}
	var classId string = "hci.IpAddress"
	this.ClassId = classId
	var objectType string = "hci.IpAddress"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HciIpAddress) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HciIpAddress) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HciIpAddress) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hci.IpAddress" of the ClassId field.
func (o *HciIpAddress) GetDefaultClassId() interface{} {
	return "hci.IpAddress"
}

// GetObjectType returns the ObjectType field value
func (o *HciIpAddress) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HciIpAddress) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HciIpAddress) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hci.IpAddress" of the ObjectType field.
func (o *HciIpAddress) GetDefaultObjectType() interface{} {
	return "hci.IpAddress"
}

// GetIpv4Address returns the Ipv4Address field value if set, zero value otherwise.
func (o *HciIpAddress) GetIpv4Address() string {
	if o == nil || IsNil(o.Ipv4Address) {
		var ret string
		return ret
	}
	return *o.Ipv4Address
}

// GetIpv4AddressOk returns a tuple with the Ipv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciIpAddress) GetIpv4AddressOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Address) {
		return nil, false
	}
	return o.Ipv4Address, true
}

// HasIpv4Address returns a boolean if a field has been set.
func (o *HciIpAddress) HasIpv4Address() bool {
	if o != nil && !IsNil(o.Ipv4Address) {
		return true
	}

	return false
}

// SetIpv4Address gets a reference to the given string and assigns it to the Ipv4Address field.
func (o *HciIpAddress) SetIpv4Address(v string) {
	o.Ipv4Address = &v
}

// GetIpv4PrefixLength returns the Ipv4PrefixLength field value if set, zero value otherwise.
func (o *HciIpAddress) GetIpv4PrefixLength() int32 {
	if o == nil || IsNil(o.Ipv4PrefixLength) {
		var ret int32
		return ret
	}
	return *o.Ipv4PrefixLength
}

// GetIpv4PrefixLengthOk returns a tuple with the Ipv4PrefixLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciIpAddress) GetIpv4PrefixLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Ipv4PrefixLength) {
		return nil, false
	}
	return o.Ipv4PrefixLength, true
}

// HasIpv4PrefixLength returns a boolean if a field has been set.
func (o *HciIpAddress) HasIpv4PrefixLength() bool {
	if o != nil && !IsNil(o.Ipv4PrefixLength) {
		return true
	}

	return false
}

// SetIpv4PrefixLength gets a reference to the given int32 and assigns it to the Ipv4PrefixLength field.
func (o *HciIpAddress) SetIpv4PrefixLength(v int32) {
	o.Ipv4PrefixLength = &v
}

// GetIpv6Address returns the Ipv6Address field value if set, zero value otherwise.
func (o *HciIpAddress) GetIpv6Address() string {
	if o == nil || IsNil(o.Ipv6Address) {
		var ret string
		return ret
	}
	return *o.Ipv6Address
}

// GetIpv6AddressOk returns a tuple with the Ipv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciIpAddress) GetIpv6AddressOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv6Address) {
		return nil, false
	}
	return o.Ipv6Address, true
}

// HasIpv6Address returns a boolean if a field has been set.
func (o *HciIpAddress) HasIpv6Address() bool {
	if o != nil && !IsNil(o.Ipv6Address) {
		return true
	}

	return false
}

// SetIpv6Address gets a reference to the given string and assigns it to the Ipv6Address field.
func (o *HciIpAddress) SetIpv6Address(v string) {
	o.Ipv6Address = &v
}

// GetIpv6PrefixLength returns the Ipv6PrefixLength field value if set, zero value otherwise.
func (o *HciIpAddress) GetIpv6PrefixLength() int32 {
	if o == nil || IsNil(o.Ipv6PrefixLength) {
		var ret int32
		return ret
	}
	return *o.Ipv6PrefixLength
}

// GetIpv6PrefixLengthOk returns a tuple with the Ipv6PrefixLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciIpAddress) GetIpv6PrefixLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Ipv6PrefixLength) {
		return nil, false
	}
	return o.Ipv6PrefixLength, true
}

// HasIpv6PrefixLength returns a boolean if a field has been set.
func (o *HciIpAddress) HasIpv6PrefixLength() bool {
	if o != nil && !IsNil(o.Ipv6PrefixLength) {
		return true
	}

	return false
}

// SetIpv6PrefixLength gets a reference to the given int32 and assigns it to the Ipv6PrefixLength field.
func (o *HciIpAddress) SetIpv6PrefixLength(v int32) {
	o.Ipv6PrefixLength = &v
}

func (o HciIpAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HciIpAddress) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Ipv4Address) {
		toSerialize["Ipv4Address"] = o.Ipv4Address
	}
	if !IsNil(o.Ipv4PrefixLength) {
		toSerialize["Ipv4PrefixLength"] = o.Ipv4PrefixLength
	}
	if !IsNil(o.Ipv6Address) {
		toSerialize["Ipv6Address"] = o.Ipv6Address
	}
	if !IsNil(o.Ipv6PrefixLength) {
		toSerialize["Ipv6PrefixLength"] = o.Ipv6PrefixLength
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HciIpAddress) UnmarshalJSON(data []byte) (err error) {
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
	type HciIpAddressWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// An IPv4 address in this IpAddress.
		Ipv4Address *string `json:"Ipv4Address,omitempty"`
		// The prefix length of the IPv4 address.
		Ipv4PrefixLength *int32 `json:"Ipv4PrefixLength,omitempty"`
		// An IPv6 address in this IpAddress.
		Ipv6Address *string `json:"Ipv6Address,omitempty"`
		// The prefix length of the IPv6 address.
		Ipv6PrefixLength *int32 `json:"Ipv6PrefixLength,omitempty"`
	}

	varHciIpAddressWithoutEmbeddedStruct := HciIpAddressWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHciIpAddressWithoutEmbeddedStruct)
	if err == nil {
		varHciIpAddress := _HciIpAddress{}
		varHciIpAddress.ClassId = varHciIpAddressWithoutEmbeddedStruct.ClassId
		varHciIpAddress.ObjectType = varHciIpAddressWithoutEmbeddedStruct.ObjectType
		varHciIpAddress.Ipv4Address = varHciIpAddressWithoutEmbeddedStruct.Ipv4Address
		varHciIpAddress.Ipv4PrefixLength = varHciIpAddressWithoutEmbeddedStruct.Ipv4PrefixLength
		varHciIpAddress.Ipv6Address = varHciIpAddressWithoutEmbeddedStruct.Ipv6Address
		varHciIpAddress.Ipv6PrefixLength = varHciIpAddressWithoutEmbeddedStruct.Ipv6PrefixLength
		*o = HciIpAddress(varHciIpAddress)
	} else {
		return err
	}

	varHciIpAddress := _HciIpAddress{}

	err = json.Unmarshal(data, &varHciIpAddress)
	if err == nil {
		o.MoBaseComplexType = varHciIpAddress.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Ipv4Address")
		delete(additionalProperties, "Ipv4PrefixLength")
		delete(additionalProperties, "Ipv6Address")
		delete(additionalProperties, "Ipv6PrefixLength")

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

type NullableHciIpAddress struct {
	value *HciIpAddress
	isSet bool
}

func (v NullableHciIpAddress) Get() *HciIpAddress {
	return v.value
}

func (v *NullableHciIpAddress) Set(val *HciIpAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableHciIpAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableHciIpAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHciIpAddress(val *HciIpAddress) *NullableHciIpAddress {
	return &NullableHciIpAddress{value: val, isSet: true}
}

func (v NullableHciIpAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHciIpAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}