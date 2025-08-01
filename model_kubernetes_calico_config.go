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

// checks if the KubernetesCalicoConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesCalicoConfig{}

// KubernetesCalicoConfig Configuration of Calico CNI plugin.
type KubernetesCalicoConfig struct {
	KubernetesCniConfig
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP version that can take on values v4 or v6. * `v4` - This refers to the IPv4 address. * `v6` - This refers to the IPv6 address.
	IpVersion *string `json:"IpVersion,omitempty"`
	// Workload interface maximum transmission unit (MTU).
	Mtu                  *int64 `json:"Mtu,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesCalicoConfig KubernetesCalicoConfig

// NewKubernetesCalicoConfig instantiates a new KubernetesCalicoConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCalicoConfig(classId string, objectType string) *KubernetesCalicoConfig {
	this := KubernetesCalicoConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ipVersion string = "v4"
	this.IpVersion = &ipVersion
	return &this
}

// NewKubernetesCalicoConfigWithDefaults instantiates a new KubernetesCalicoConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesCalicoConfigWithDefaults() *KubernetesCalicoConfig {
	this := KubernetesCalicoConfig{}
	var classId string = "kubernetes.CalicoConfig"
	this.ClassId = classId
	var objectType string = "kubernetes.CalicoConfig"
	this.ObjectType = objectType
	var ipVersion string = "v4"
	this.IpVersion = &ipVersion
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesCalicoConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesCalicoConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesCalicoConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.CalicoConfig" of the ClassId field.
func (o *KubernetesCalicoConfig) GetDefaultClassId() interface{} {
	return "kubernetes.CalicoConfig"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesCalicoConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesCalicoConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesCalicoConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.CalicoConfig" of the ObjectType field.
func (o *KubernetesCalicoConfig) GetDefaultObjectType() interface{} {
	return "kubernetes.CalicoConfig"
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *KubernetesCalicoConfig) GetIpVersion() string {
	if o == nil || IsNil(o.IpVersion) {
		var ret string
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCalicoConfig) GetIpVersionOk() (*string, bool) {
	if o == nil || IsNil(o.IpVersion) {
		return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *KubernetesCalicoConfig) HasIpVersion() bool {
	if o != nil && !IsNil(o.IpVersion) {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given string and assigns it to the IpVersion field.
func (o *KubernetesCalicoConfig) SetIpVersion(v string) {
	o.IpVersion = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *KubernetesCalicoConfig) GetMtu() int64 {
	if o == nil || IsNil(o.Mtu) {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCalicoConfig) GetMtuOk() (*int64, bool) {
	if o == nil || IsNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *KubernetesCalicoConfig) HasMtu() bool {
	if o != nil && !IsNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *KubernetesCalicoConfig) SetMtu(v int64) {
	o.Mtu = &v
}

func (o KubernetesCalicoConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesCalicoConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesCniConfig, errKubernetesCniConfig := json.Marshal(o.KubernetesCniConfig)
	if errKubernetesCniConfig != nil {
		return map[string]interface{}{}, errKubernetesCniConfig
	}
	errKubernetesCniConfig = json.Unmarshal([]byte(serializedKubernetesCniConfig), &toSerialize)
	if errKubernetesCniConfig != nil {
		return map[string]interface{}{}, errKubernetesCniConfig
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.IpVersion) {
		toSerialize["IpVersion"] = o.IpVersion
	}
	if !IsNil(o.Mtu) {
		toSerialize["Mtu"] = o.Mtu
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesCalicoConfig) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesCalicoConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP version that can take on values v4 or v6. * `v4` - This refers to the IPv4 address. * `v6` - This refers to the IPv6 address.
		IpVersion *string `json:"IpVersion,omitempty"`
		// Workload interface maximum transmission unit (MTU).
		Mtu *int64 `json:"Mtu,omitempty"`
	}

	varKubernetesCalicoConfigWithoutEmbeddedStruct := KubernetesCalicoConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesCalicoConfigWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesCalicoConfig := _KubernetesCalicoConfig{}
		varKubernetesCalicoConfig.ClassId = varKubernetesCalicoConfigWithoutEmbeddedStruct.ClassId
		varKubernetesCalicoConfig.ObjectType = varKubernetesCalicoConfigWithoutEmbeddedStruct.ObjectType
		varKubernetesCalicoConfig.IpVersion = varKubernetesCalicoConfigWithoutEmbeddedStruct.IpVersion
		varKubernetesCalicoConfig.Mtu = varKubernetesCalicoConfigWithoutEmbeddedStruct.Mtu
		*o = KubernetesCalicoConfig(varKubernetesCalicoConfig)
	} else {
		return err
	}

	varKubernetesCalicoConfig := _KubernetesCalicoConfig{}

	err = json.Unmarshal(data, &varKubernetesCalicoConfig)
	if err == nil {
		o.KubernetesCniConfig = varKubernetesCalicoConfig.KubernetesCniConfig
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpVersion")
		delete(additionalProperties, "Mtu")

		// remove fields from embedded structs
		reflectKubernetesCniConfig := reflect.ValueOf(o.KubernetesCniConfig)
		for i := 0; i < reflectKubernetesCniConfig.Type().NumField(); i++ {
			t := reflectKubernetesCniConfig.Type().Field(i)

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

type NullableKubernetesCalicoConfig struct {
	value *KubernetesCalicoConfig
	isSet bool
}

func (v NullableKubernetesCalicoConfig) Get() *KubernetesCalicoConfig {
	return v.value
}

func (v *NullableKubernetesCalicoConfig) Set(val *KubernetesCalicoConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCalicoConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCalicoConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCalicoConfig(val *KubernetesCalicoConfig) *NullableKubernetesCalicoConfig {
	return &NullableKubernetesCalicoConfig{value: val, isSet: true}
}

func (v NullableKubernetesCalicoConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCalicoConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
