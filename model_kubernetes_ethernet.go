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

// checks if the KubernetesEthernet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesEthernet{}

// KubernetesEthernet Configuration to apply to a Physical Network Interface Card.
type KubernetesEthernet struct {
	KubernetesNetworkInterface
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                            `json:"ObjectType"`
	Matcher    NullableKubernetesEthernetMatcher `json:"Matcher,omitempty"`
	// If the infrastructure network is selectable, this indicates which network to attach to the port.
	ProviderName         *string `json:"ProviderName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesEthernet KubernetesEthernet

// NewKubernetesEthernet instantiates a new KubernetesEthernet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesEthernet(classId string, objectType string) *KubernetesEthernet {
	this := KubernetesEthernet{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesEthernetWithDefaults instantiates a new KubernetesEthernet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesEthernetWithDefaults() *KubernetesEthernet {
	this := KubernetesEthernet{}
	var classId string = "kubernetes.Ethernet"
	this.ClassId = classId
	var objectType string = "kubernetes.Ethernet"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesEthernet) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesEthernet) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesEthernet) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.Ethernet" of the ClassId field.
func (o *KubernetesEthernet) GetDefaultClassId() interface{} {
	return "kubernetes.Ethernet"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesEthernet) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesEthernet) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesEthernet) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.Ethernet" of the ObjectType field.
func (o *KubernetesEthernet) GetDefaultObjectType() interface{} {
	return "kubernetes.Ethernet"
}

// GetMatcher returns the Matcher field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesEthernet) GetMatcher() KubernetesEthernetMatcher {
	if o == nil || IsNil(o.Matcher.Get()) {
		var ret KubernetesEthernetMatcher
		return ret
	}
	return *o.Matcher.Get()
}

// GetMatcherOk returns a tuple with the Matcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesEthernet) GetMatcherOk() (*KubernetesEthernetMatcher, bool) {
	if o == nil {
		return nil, false
	}
	return o.Matcher.Get(), o.Matcher.IsSet()
}

// HasMatcher returns a boolean if a field has been set.
func (o *KubernetesEthernet) HasMatcher() bool {
	if o != nil && o.Matcher.IsSet() {
		return true
	}

	return false
}

// SetMatcher gets a reference to the given NullableKubernetesEthernetMatcher and assigns it to the Matcher field.
func (o *KubernetesEthernet) SetMatcher(v KubernetesEthernetMatcher) {
	o.Matcher.Set(&v)
}

// SetMatcherNil sets the value for Matcher to be an explicit nil
func (o *KubernetesEthernet) SetMatcherNil() {
	o.Matcher.Set(nil)
}

// UnsetMatcher ensures that no value is present for Matcher, not even an explicit nil
func (o *KubernetesEthernet) UnsetMatcher() {
	o.Matcher.Unset()
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *KubernetesEthernet) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEthernet) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *KubernetesEthernet) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *KubernetesEthernet) SetProviderName(v string) {
	o.ProviderName = &v
}

func (o KubernetesEthernet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesEthernet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesNetworkInterface, errKubernetesNetworkInterface := json.Marshal(o.KubernetesNetworkInterface)
	if errKubernetesNetworkInterface != nil {
		return map[string]interface{}{}, errKubernetesNetworkInterface
	}
	errKubernetesNetworkInterface = json.Unmarshal([]byte(serializedKubernetesNetworkInterface), &toSerialize)
	if errKubernetesNetworkInterface != nil {
		return map[string]interface{}{}, errKubernetesNetworkInterface
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Matcher.IsSet() {
		toSerialize["Matcher"] = o.Matcher.Get()
	}
	if !IsNil(o.ProviderName) {
		toSerialize["ProviderName"] = o.ProviderName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesEthernet) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesEthernetWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                            `json:"ObjectType"`
		Matcher    NullableKubernetesEthernetMatcher `json:"Matcher,omitempty"`
		// If the infrastructure network is selectable, this indicates which network to attach to the port.
		ProviderName *string `json:"ProviderName,omitempty"`
	}

	varKubernetesEthernetWithoutEmbeddedStruct := KubernetesEthernetWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesEthernetWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesEthernet := _KubernetesEthernet{}
		varKubernetesEthernet.ClassId = varKubernetesEthernetWithoutEmbeddedStruct.ClassId
		varKubernetesEthernet.ObjectType = varKubernetesEthernetWithoutEmbeddedStruct.ObjectType
		varKubernetesEthernet.Matcher = varKubernetesEthernetWithoutEmbeddedStruct.Matcher
		varKubernetesEthernet.ProviderName = varKubernetesEthernetWithoutEmbeddedStruct.ProviderName
		*o = KubernetesEthernet(varKubernetesEthernet)
	} else {
		return err
	}

	varKubernetesEthernet := _KubernetesEthernet{}

	err = json.Unmarshal(data, &varKubernetesEthernet)
	if err == nil {
		o.KubernetesNetworkInterface = varKubernetesEthernet.KubernetesNetworkInterface
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Matcher")
		delete(additionalProperties, "ProviderName")

		// remove fields from embedded structs
		reflectKubernetesNetworkInterface := reflect.ValueOf(o.KubernetesNetworkInterface)
		for i := 0; i < reflectKubernetesNetworkInterface.Type().NumField(); i++ {
			t := reflectKubernetesNetworkInterface.Type().Field(i)

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

type NullableKubernetesEthernet struct {
	value *KubernetesEthernet
	isSet bool
}

func (v NullableKubernetesEthernet) Get() *KubernetesEthernet {
	return v.value
}

func (v *NullableKubernetesEthernet) Set(val *KubernetesEthernet) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesEthernet) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesEthernet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesEthernet(val *KubernetesEthernet) *NullableKubernetesEthernet {
	return &NullableKubernetesEthernet{value: val, isSet: true}
}

func (v NullableKubernetesEthernet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesEthernet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
