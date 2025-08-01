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

// checks if the KubernetesAddonPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesAddonPolicy{}

// KubernetesAddonPolicy A policy that defines which AddonDefinitions to use.
type KubernetesAddonPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                        `json:"ObjectType"`
	AddonConfiguration   NullableKubernetesAddonConfiguration          `json:"AddonConfiguration,omitempty"`
	AddonDefinition      NullableKubernetesAddonDefinitionRelationship `json:"AddonDefinition,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAddonPolicy KubernetesAddonPolicy

// NewKubernetesAddonPolicy instantiates a new KubernetesAddonPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAddonPolicy(classId string, objectType string) *KubernetesAddonPolicy {
	this := KubernetesAddonPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAddonPolicyWithDefaults instantiates a new KubernetesAddonPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAddonPolicyWithDefaults() *KubernetesAddonPolicy {
	this := KubernetesAddonPolicy{}
	var classId string = "kubernetes.AddonPolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.AddonPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAddonPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAddonPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.AddonPolicy" of the ClassId field.
func (o *KubernetesAddonPolicy) GetDefaultClassId() interface{} {
	return "kubernetes.AddonPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAddonPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAddonPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.AddonPolicy" of the ObjectType field.
func (o *KubernetesAddonPolicy) GetDefaultObjectType() interface{} {
	return "kubernetes.AddonPolicy"
}

// GetAddonConfiguration returns the AddonConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAddonPolicy) GetAddonConfiguration() KubernetesAddonConfiguration {
	if o == nil || IsNil(o.AddonConfiguration.Get()) {
		var ret KubernetesAddonConfiguration
		return ret
	}
	return *o.AddonConfiguration.Get()
}

// GetAddonConfigurationOk returns a tuple with the AddonConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAddonPolicy) GetAddonConfigurationOk() (*KubernetesAddonConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddonConfiguration.Get(), o.AddonConfiguration.IsSet()
}

// HasAddonConfiguration returns a boolean if a field has been set.
func (o *KubernetesAddonPolicy) HasAddonConfiguration() bool {
	if o != nil && o.AddonConfiguration.IsSet() {
		return true
	}

	return false
}

// SetAddonConfiguration gets a reference to the given NullableKubernetesAddonConfiguration and assigns it to the AddonConfiguration field.
func (o *KubernetesAddonPolicy) SetAddonConfiguration(v KubernetesAddonConfiguration) {
	o.AddonConfiguration.Set(&v)
}

// SetAddonConfigurationNil sets the value for AddonConfiguration to be an explicit nil
func (o *KubernetesAddonPolicy) SetAddonConfigurationNil() {
	o.AddonConfiguration.Set(nil)
}

// UnsetAddonConfiguration ensures that no value is present for AddonConfiguration, not even an explicit nil
func (o *KubernetesAddonPolicy) UnsetAddonConfiguration() {
	o.AddonConfiguration.Unset()
}

// GetAddonDefinition returns the AddonDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAddonPolicy) GetAddonDefinition() KubernetesAddonDefinitionRelationship {
	if o == nil || IsNil(o.AddonDefinition.Get()) {
		var ret KubernetesAddonDefinitionRelationship
		return ret
	}
	return *o.AddonDefinition.Get()
}

// GetAddonDefinitionOk returns a tuple with the AddonDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAddonPolicy) GetAddonDefinitionOk() (*KubernetesAddonDefinitionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddonDefinition.Get(), o.AddonDefinition.IsSet()
}

// HasAddonDefinition returns a boolean if a field has been set.
func (o *KubernetesAddonPolicy) HasAddonDefinition() bool {
	if o != nil && o.AddonDefinition.IsSet() {
		return true
	}

	return false
}

// SetAddonDefinition gets a reference to the given NullableKubernetesAddonDefinitionRelationship and assigns it to the AddonDefinition field.
func (o *KubernetesAddonPolicy) SetAddonDefinition(v KubernetesAddonDefinitionRelationship) {
	o.AddonDefinition.Set(&v)
}

// SetAddonDefinitionNil sets the value for AddonDefinition to be an explicit nil
func (o *KubernetesAddonPolicy) SetAddonDefinitionNil() {
	o.AddonDefinition.Set(nil)
}

// UnsetAddonDefinition ensures that no value is present for AddonDefinition, not even an explicit nil
func (o *KubernetesAddonPolicy) UnsetAddonDefinition() {
	o.AddonDefinition.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAddonPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAddonPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesAddonPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesAddonPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *KubernetesAddonPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *KubernetesAddonPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

func (o KubernetesAddonPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesAddonPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.AddonConfiguration.IsSet() {
		toSerialize["AddonConfiguration"] = o.AddonConfiguration.Get()
	}
	if o.AddonDefinition.IsSet() {
		toSerialize["AddonDefinition"] = o.AddonDefinition.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesAddonPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesAddonPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType         string                                        `json:"ObjectType"`
		AddonConfiguration NullableKubernetesAddonConfiguration          `json:"AddonConfiguration,omitempty"`
		AddonDefinition    NullableKubernetesAddonDefinitionRelationship `json:"AddonDefinition,omitempty"`
		Organization       NullableOrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	}

	varKubernetesAddonPolicyWithoutEmbeddedStruct := KubernetesAddonPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesAddonPolicyWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesAddonPolicy := _KubernetesAddonPolicy{}
		varKubernetesAddonPolicy.ClassId = varKubernetesAddonPolicyWithoutEmbeddedStruct.ClassId
		varKubernetesAddonPolicy.ObjectType = varKubernetesAddonPolicyWithoutEmbeddedStruct.ObjectType
		varKubernetesAddonPolicy.AddonConfiguration = varKubernetesAddonPolicyWithoutEmbeddedStruct.AddonConfiguration
		varKubernetesAddonPolicy.AddonDefinition = varKubernetesAddonPolicyWithoutEmbeddedStruct.AddonDefinition
		varKubernetesAddonPolicy.Organization = varKubernetesAddonPolicyWithoutEmbeddedStruct.Organization
		*o = KubernetesAddonPolicy(varKubernetesAddonPolicy)
	} else {
		return err
	}

	varKubernetesAddonPolicy := _KubernetesAddonPolicy{}

	err = json.Unmarshal(data, &varKubernetesAddonPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varKubernetesAddonPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AddonConfiguration")
		delete(additionalProperties, "AddonDefinition")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableKubernetesAddonPolicy struct {
	value *KubernetesAddonPolicy
	isSet bool
}

func (v NullableKubernetesAddonPolicy) Get() *KubernetesAddonPolicy {
	return v.value
}

func (v *NullableKubernetesAddonPolicy) Set(val *KubernetesAddonPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAddonPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAddonPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAddonPolicy(val *KubernetesAddonPolicy) *NullableKubernetesAddonPolicy {
	return &NullableKubernetesAddonPolicy{value: val, isSet: true}
}

func (v NullableKubernetesAddonPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAddonPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
