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

// checks if the KvmTunneledKvmPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KvmTunneledKvmPolicy{}

// KvmTunneledKvmPolicy Policy to control Tunelled vKVM for a specific account.
type KvmTunneledKvmPolicy struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable or Disable configuration of tunneled KVM for a specific account.
	TunneledKvmConfiguration *bool `json:"TunneledKvmConfiguration,omitempty"`
	// Enable or Disable launching tunneled KVM for a specific account.
	TunneledKvmLaunch    *bool                          `json:"TunneledKvmLaunch,omitempty"`
	Account              NullableIamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmTunneledKvmPolicy KvmTunneledKvmPolicy

// NewKvmTunneledKvmPolicy instantiates a new KvmTunneledKvmPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmTunneledKvmPolicy(classId string, objectType string) *KvmTunneledKvmPolicy {
	this := KvmTunneledKvmPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var tunneledKvmConfiguration bool = false
	this.TunneledKvmConfiguration = &tunneledKvmConfiguration
	var tunneledKvmLaunch bool = false
	this.TunneledKvmLaunch = &tunneledKvmLaunch
	return &this
}

// NewKvmTunneledKvmPolicyWithDefaults instantiates a new KvmTunneledKvmPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmTunneledKvmPolicyWithDefaults() *KvmTunneledKvmPolicy {
	this := KvmTunneledKvmPolicy{}
	var classId string = "kvm.TunneledKvmPolicy"
	this.ClassId = classId
	var objectType string = "kvm.TunneledKvmPolicy"
	this.ObjectType = objectType
	var tunneledKvmConfiguration bool = false
	this.TunneledKvmConfiguration = &tunneledKvmConfiguration
	var tunneledKvmLaunch bool = false
	this.TunneledKvmLaunch = &tunneledKvmLaunch
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmTunneledKvmPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmTunneledKvmPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmTunneledKvmPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kvm.TunneledKvmPolicy" of the ClassId field.
func (o *KvmTunneledKvmPolicy) GetDefaultClassId() interface{} {
	return "kvm.TunneledKvmPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *KvmTunneledKvmPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmTunneledKvmPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmTunneledKvmPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kvm.TunneledKvmPolicy" of the ObjectType field.
func (o *KvmTunneledKvmPolicy) GetDefaultObjectType() interface{} {
	return "kvm.TunneledKvmPolicy"
}

// GetTunneledKvmConfiguration returns the TunneledKvmConfiguration field value if set, zero value otherwise.
func (o *KvmTunneledKvmPolicy) GetTunneledKvmConfiguration() bool {
	if o == nil || IsNil(o.TunneledKvmConfiguration) {
		var ret bool
		return ret
	}
	return *o.TunneledKvmConfiguration
}

// GetTunneledKvmConfigurationOk returns a tuple with the TunneledKvmConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmTunneledKvmPolicy) GetTunneledKvmConfigurationOk() (*bool, bool) {
	if o == nil || IsNil(o.TunneledKvmConfiguration) {
		return nil, false
	}
	return o.TunneledKvmConfiguration, true
}

// HasTunneledKvmConfiguration returns a boolean if a field has been set.
func (o *KvmTunneledKvmPolicy) HasTunneledKvmConfiguration() bool {
	if o != nil && !IsNil(o.TunneledKvmConfiguration) {
		return true
	}

	return false
}

// SetTunneledKvmConfiguration gets a reference to the given bool and assigns it to the TunneledKvmConfiguration field.
func (o *KvmTunneledKvmPolicy) SetTunneledKvmConfiguration(v bool) {
	o.TunneledKvmConfiguration = &v
}

// GetTunneledKvmLaunch returns the TunneledKvmLaunch field value if set, zero value otherwise.
func (o *KvmTunneledKvmPolicy) GetTunneledKvmLaunch() bool {
	if o == nil || IsNil(o.TunneledKvmLaunch) {
		var ret bool
		return ret
	}
	return *o.TunneledKvmLaunch
}

// GetTunneledKvmLaunchOk returns a tuple with the TunneledKvmLaunch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmTunneledKvmPolicy) GetTunneledKvmLaunchOk() (*bool, bool) {
	if o == nil || IsNil(o.TunneledKvmLaunch) {
		return nil, false
	}
	return o.TunneledKvmLaunch, true
}

// HasTunneledKvmLaunch returns a boolean if a field has been set.
func (o *KvmTunneledKvmPolicy) HasTunneledKvmLaunch() bool {
	if o != nil && !IsNil(o.TunneledKvmLaunch) {
		return true
	}

	return false
}

// SetTunneledKvmLaunch gets a reference to the given bool and assigns it to the TunneledKvmLaunch field.
func (o *KvmTunneledKvmPolicy) SetTunneledKvmLaunch(v bool) {
	o.TunneledKvmLaunch = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmTunneledKvmPolicy) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmTunneledKvmPolicy) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *KvmTunneledKvmPolicy) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *KvmTunneledKvmPolicy) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *KvmTunneledKvmPolicy) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *KvmTunneledKvmPolicy) UnsetAccount() {
	o.Account.Unset()
}

func (o KvmTunneledKvmPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KvmTunneledKvmPolicy) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.TunneledKvmConfiguration) {
		toSerialize["TunneledKvmConfiguration"] = o.TunneledKvmConfiguration
	}
	if !IsNil(o.TunneledKvmLaunch) {
		toSerialize["TunneledKvmLaunch"] = o.TunneledKvmLaunch
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KvmTunneledKvmPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type KvmTunneledKvmPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enable or Disable configuration of tunneled KVM for a specific account.
		TunneledKvmConfiguration *bool `json:"TunneledKvmConfiguration,omitempty"`
		// Enable or Disable launching tunneled KVM for a specific account.
		TunneledKvmLaunch *bool                          `json:"TunneledKvmLaunch,omitempty"`
		Account           NullableIamAccountRelationship `json:"Account,omitempty"`
	}

	varKvmTunneledKvmPolicyWithoutEmbeddedStruct := KvmTunneledKvmPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKvmTunneledKvmPolicyWithoutEmbeddedStruct)
	if err == nil {
		varKvmTunneledKvmPolicy := _KvmTunneledKvmPolicy{}
		varKvmTunneledKvmPolicy.ClassId = varKvmTunneledKvmPolicyWithoutEmbeddedStruct.ClassId
		varKvmTunneledKvmPolicy.ObjectType = varKvmTunneledKvmPolicyWithoutEmbeddedStruct.ObjectType
		varKvmTunneledKvmPolicy.TunneledKvmConfiguration = varKvmTunneledKvmPolicyWithoutEmbeddedStruct.TunneledKvmConfiguration
		varKvmTunneledKvmPolicy.TunneledKvmLaunch = varKvmTunneledKvmPolicyWithoutEmbeddedStruct.TunneledKvmLaunch
		varKvmTunneledKvmPolicy.Account = varKvmTunneledKvmPolicyWithoutEmbeddedStruct.Account
		*o = KvmTunneledKvmPolicy(varKvmTunneledKvmPolicy)
	} else {
		return err
	}

	varKvmTunneledKvmPolicy := _KvmTunneledKvmPolicy{}

	err = json.Unmarshal(data, &varKvmTunneledKvmPolicy)
	if err == nil {
		o.MoBaseMo = varKvmTunneledKvmPolicy.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TunneledKvmConfiguration")
		delete(additionalProperties, "TunneledKvmLaunch")
		delete(additionalProperties, "Account")

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

type NullableKvmTunneledKvmPolicy struct {
	value *KvmTunneledKvmPolicy
	isSet bool
}

func (v NullableKvmTunneledKvmPolicy) Get() *KvmTunneledKvmPolicy {
	return v.value
}

func (v *NullableKvmTunneledKvmPolicy) Set(val *KvmTunneledKvmPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmTunneledKvmPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmTunneledKvmPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmTunneledKvmPolicy(val *KvmTunneledKvmPolicy) *NullableKvmTunneledKvmPolicy {
	return &NullableKvmTunneledKvmPolicy{value: val, isSet: true}
}

func (v NullableKvmTunneledKvmPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmTunneledKvmPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
