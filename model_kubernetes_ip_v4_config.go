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

// checks if the KubernetesIpV4Config type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesIpV4Config{}

// KubernetesIpV4Config Network interface configuration data for IPv4 interfaces.
type KubernetesIpV4Config struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IPv4 Address in CIDR format.
	Ip                   *string  `json:"Ip,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\/([0-9]|[1-2][0-9]|3[0-2])$"`
	Lease                *MoMoRef `json:"Lease,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesIpV4Config KubernetesIpV4Config

// NewKubernetesIpV4Config instantiates a new KubernetesIpV4Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesIpV4Config(classId string, objectType string) *KubernetesIpV4Config {
	this := KubernetesIpV4Config{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesIpV4ConfigWithDefaults instantiates a new KubernetesIpV4Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesIpV4ConfigWithDefaults() *KubernetesIpV4Config {
	this := KubernetesIpV4Config{}
	var classId string = "kubernetes.IpV4Config"
	this.ClassId = classId
	var objectType string = "kubernetes.IpV4Config"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesIpV4Config) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesIpV4Config) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesIpV4Config) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.IpV4Config" of the ClassId field.
func (o *KubernetesIpV4Config) GetDefaultClassId() interface{} {
	return "kubernetes.IpV4Config"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesIpV4Config) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesIpV4Config) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesIpV4Config) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.IpV4Config" of the ObjectType field.
func (o *KubernetesIpV4Config) GetDefaultObjectType() interface{} {
	return "kubernetes.IpV4Config"
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *KubernetesIpV4Config) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesIpV4Config) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *KubernetesIpV4Config) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *KubernetesIpV4Config) SetIp(v string) {
	o.Ip = &v
}

// GetLease returns the Lease field value if set, zero value otherwise.
func (o *KubernetesIpV4Config) GetLease() MoMoRef {
	if o == nil || IsNil(o.Lease) {
		var ret MoMoRef
		return ret
	}
	return *o.Lease
}

// GetLeaseOk returns a tuple with the Lease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesIpV4Config) GetLeaseOk() (*MoMoRef, bool) {
	if o == nil || IsNil(o.Lease) {
		return nil, false
	}
	return o.Lease, true
}

// HasLease returns a boolean if a field has been set.
func (o *KubernetesIpV4Config) HasLease() bool {
	if o != nil && !IsNil(o.Lease) {
		return true
	}

	return false
}

// SetLease gets a reference to the given MoMoRef and assigns it to the Lease field.
func (o *KubernetesIpV4Config) SetLease(v MoMoRef) {
	o.Lease = &v
}

func (o KubernetesIpV4Config) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesIpV4Config) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Ip) {
		toSerialize["Ip"] = o.Ip
	}
	if !IsNil(o.Lease) {
		toSerialize["Lease"] = o.Lease
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesIpV4Config) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesIpV4ConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IPv4 Address in CIDR format.
		Ip    *string  `json:"Ip,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\/([0-9]|[1-2][0-9]|3[0-2])$"`
		Lease *MoMoRef `json:"Lease,omitempty"`
	}

	varKubernetesIpV4ConfigWithoutEmbeddedStruct := KubernetesIpV4ConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesIpV4ConfigWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesIpV4Config := _KubernetesIpV4Config{}
		varKubernetesIpV4Config.ClassId = varKubernetesIpV4ConfigWithoutEmbeddedStruct.ClassId
		varKubernetesIpV4Config.ObjectType = varKubernetesIpV4ConfigWithoutEmbeddedStruct.ObjectType
		varKubernetesIpV4Config.Ip = varKubernetesIpV4ConfigWithoutEmbeddedStruct.Ip
		varKubernetesIpV4Config.Lease = varKubernetesIpV4ConfigWithoutEmbeddedStruct.Lease
		*o = KubernetesIpV4Config(varKubernetesIpV4Config)
	} else {
		return err
	}

	varKubernetesIpV4Config := _KubernetesIpV4Config{}

	err = json.Unmarshal(data, &varKubernetesIpV4Config)
	if err == nil {
		o.MoBaseComplexType = varKubernetesIpV4Config.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "Lease")

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

type NullableKubernetesIpV4Config struct {
	value *KubernetesIpV4Config
	isSet bool
}

func (v NullableKubernetesIpV4Config) Get() *KubernetesIpV4Config {
	return v.value
}

func (v *NullableKubernetesIpV4Config) Set(val *KubernetesIpV4Config) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesIpV4Config) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesIpV4Config) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesIpV4Config(val *KubernetesIpV4Config) *NullableKubernetesIpV4Config {
	return &NullableKubernetesIpV4Config{value: val, isSet: true}
}

func (v NullableKubernetesIpV4Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesIpV4Config) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
