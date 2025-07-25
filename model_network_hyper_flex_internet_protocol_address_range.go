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

// checks if the NetworkHyperFlexInternetProtocolAddressRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkHyperFlexInternetProtocolAddressRange{}

// NetworkHyperFlexInternetProtocolAddressRange The HyperFlex internet protocol address range.
type NetworkHyperFlexInternetProtocolAddressRange struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Begining IP address. IPv4 only.
	BeginAddress *string `json:"BeginAddress,omitempty"`
	// Ending IP address. IPv4 only.
	EndAddress           *string `json:"EndAddress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkHyperFlexInternetProtocolAddressRange NetworkHyperFlexInternetProtocolAddressRange

// NewNetworkHyperFlexInternetProtocolAddressRange instantiates a new NetworkHyperFlexInternetProtocolAddressRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkHyperFlexInternetProtocolAddressRange(classId string, objectType string) *NetworkHyperFlexInternetProtocolAddressRange {
	this := NetworkHyperFlexInternetProtocolAddressRange{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkHyperFlexInternetProtocolAddressRangeWithDefaults instantiates a new NetworkHyperFlexInternetProtocolAddressRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkHyperFlexInternetProtocolAddressRangeWithDefaults() *NetworkHyperFlexInternetProtocolAddressRange {
	this := NetworkHyperFlexInternetProtocolAddressRange{}
	var classId string = "network.HyperFlexInternetProtocolAddressRange"
	this.ClassId = classId
	var objectType string = "network.HyperFlexInternetProtocolAddressRange"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkHyperFlexInternetProtocolAddressRange) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkHyperFlexInternetProtocolAddressRange) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkHyperFlexInternetProtocolAddressRange) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "network.HyperFlexInternetProtocolAddressRange" of the ClassId field.
func (o *NetworkHyperFlexInternetProtocolAddressRange) GetDefaultClassId() interface{} {
	return "network.HyperFlexInternetProtocolAddressRange"
}

// GetObjectType returns the ObjectType field value
func (o *NetworkHyperFlexInternetProtocolAddressRange) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkHyperFlexInternetProtocolAddressRange) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkHyperFlexInternetProtocolAddressRange) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "network.HyperFlexInternetProtocolAddressRange" of the ObjectType field.
func (o *NetworkHyperFlexInternetProtocolAddressRange) GetDefaultObjectType() interface{} {
	return "network.HyperFlexInternetProtocolAddressRange"
}

// GetBeginAddress returns the BeginAddress field value if set, zero value otherwise.
func (o *NetworkHyperFlexInternetProtocolAddressRange) GetBeginAddress() string {
	if o == nil || IsNil(o.BeginAddress) {
		var ret string
		return ret
	}
	return *o.BeginAddress
}

// GetBeginAddressOk returns a tuple with the BeginAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHyperFlexInternetProtocolAddressRange) GetBeginAddressOk() (*string, bool) {
	if o == nil || IsNil(o.BeginAddress) {
		return nil, false
	}
	return o.BeginAddress, true
}

// HasBeginAddress returns a boolean if a field has been set.
func (o *NetworkHyperFlexInternetProtocolAddressRange) HasBeginAddress() bool {
	if o != nil && !IsNil(o.BeginAddress) {
		return true
	}

	return false
}

// SetBeginAddress gets a reference to the given string and assigns it to the BeginAddress field.
func (o *NetworkHyperFlexInternetProtocolAddressRange) SetBeginAddress(v string) {
	o.BeginAddress = &v
}

// GetEndAddress returns the EndAddress field value if set, zero value otherwise.
func (o *NetworkHyperFlexInternetProtocolAddressRange) GetEndAddress() string {
	if o == nil || IsNil(o.EndAddress) {
		var ret string
		return ret
	}
	return *o.EndAddress
}

// GetEndAddressOk returns a tuple with the EndAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHyperFlexInternetProtocolAddressRange) GetEndAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EndAddress) {
		return nil, false
	}
	return o.EndAddress, true
}

// HasEndAddress returns a boolean if a field has been set.
func (o *NetworkHyperFlexInternetProtocolAddressRange) HasEndAddress() bool {
	if o != nil && !IsNil(o.EndAddress) {
		return true
	}

	return false
}

// SetEndAddress gets a reference to the given string and assigns it to the EndAddress field.
func (o *NetworkHyperFlexInternetProtocolAddressRange) SetEndAddress(v string) {
	o.EndAddress = &v
}

func (o NetworkHyperFlexInternetProtocolAddressRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkHyperFlexInternetProtocolAddressRange) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BeginAddress) {
		toSerialize["BeginAddress"] = o.BeginAddress
	}
	if !IsNil(o.EndAddress) {
		toSerialize["EndAddress"] = o.EndAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkHyperFlexInternetProtocolAddressRange) UnmarshalJSON(data []byte) (err error) {
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
	type NetworkHyperFlexInternetProtocolAddressRangeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Begining IP address. IPv4 only.
		BeginAddress *string `json:"BeginAddress,omitempty"`
		// Ending IP address. IPv4 only.
		EndAddress *string `json:"EndAddress,omitempty"`
	}

	varNetworkHyperFlexInternetProtocolAddressRangeWithoutEmbeddedStruct := NetworkHyperFlexInternetProtocolAddressRangeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNetworkHyperFlexInternetProtocolAddressRangeWithoutEmbeddedStruct)
	if err == nil {
		varNetworkHyperFlexInternetProtocolAddressRange := _NetworkHyperFlexInternetProtocolAddressRange{}
		varNetworkHyperFlexInternetProtocolAddressRange.ClassId = varNetworkHyperFlexInternetProtocolAddressRangeWithoutEmbeddedStruct.ClassId
		varNetworkHyperFlexInternetProtocolAddressRange.ObjectType = varNetworkHyperFlexInternetProtocolAddressRangeWithoutEmbeddedStruct.ObjectType
		varNetworkHyperFlexInternetProtocolAddressRange.BeginAddress = varNetworkHyperFlexInternetProtocolAddressRangeWithoutEmbeddedStruct.BeginAddress
		varNetworkHyperFlexInternetProtocolAddressRange.EndAddress = varNetworkHyperFlexInternetProtocolAddressRangeWithoutEmbeddedStruct.EndAddress
		*o = NetworkHyperFlexInternetProtocolAddressRange(varNetworkHyperFlexInternetProtocolAddressRange)
	} else {
		return err
	}

	varNetworkHyperFlexInternetProtocolAddressRange := _NetworkHyperFlexInternetProtocolAddressRange{}

	err = json.Unmarshal(data, &varNetworkHyperFlexInternetProtocolAddressRange)
	if err == nil {
		o.MoBaseComplexType = varNetworkHyperFlexInternetProtocolAddressRange.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BeginAddress")
		delete(additionalProperties, "EndAddress")

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

type NullableNetworkHyperFlexInternetProtocolAddressRange struct {
	value *NetworkHyperFlexInternetProtocolAddressRange
	isSet bool
}

func (v NullableNetworkHyperFlexInternetProtocolAddressRange) Get() *NetworkHyperFlexInternetProtocolAddressRange {
	return v.value
}

func (v *NullableNetworkHyperFlexInternetProtocolAddressRange) Set(val *NetworkHyperFlexInternetProtocolAddressRange) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkHyperFlexInternetProtocolAddressRange) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkHyperFlexInternetProtocolAddressRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkHyperFlexInternetProtocolAddressRange(val *NetworkHyperFlexInternetProtocolAddressRange) *NullableNetworkHyperFlexInternetProtocolAddressRange {
	return &NullableNetworkHyperFlexInternetProtocolAddressRange{value: val, isSet: true}
}

func (v NullableNetworkHyperFlexInternetProtocolAddressRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkHyperFlexInternetProtocolAddressRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
