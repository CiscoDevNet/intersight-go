/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2024101709
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

// checks if the CapabilityDomainPolicyRequirement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityDomainPolicyRequirement{}

// CapabilityDomainPolicyRequirement Version Constraint requirement for a domian policy.
type CapabilityDomainPolicyRequirement struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType                   string                                           `json:"ObjectType"`
	CertificatePolicyConstraints NullableCapabilityCertificatePropertyConstraints `json:"CertificatePolicyConstraints,omitempty"`
	LdapConstraints              NullableCapabilityLdapBasePropertyConstraints    `json:"LdapConstraints,omitempty"`
	// Minimum Bundle version from which it is supported.
	MinBundleVersion *string `json:"MinBundleVersion,omitempty"`
	// Minimum Version from which policy is supported.
	MinVersion *string `json:"MinVersion,omitempty"`
	// Type of the platform for which version compatibility is specified. Example - 3GFI, 4GFI, etc.
	Model *string `json:"Model,omitempty"`
	// Policy Name for which version compatibility is specified. Example - snmp.Policy, ldap.Policy.
	PolicyName           *string `json:"PolicyName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityDomainPolicyRequirement CapabilityDomainPolicyRequirement

// NewCapabilityDomainPolicyRequirement instantiates a new CapabilityDomainPolicyRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityDomainPolicyRequirement(classId string, objectType string) *CapabilityDomainPolicyRequirement {
	this := CapabilityDomainPolicyRequirement{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityDomainPolicyRequirementWithDefaults instantiates a new CapabilityDomainPolicyRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityDomainPolicyRequirementWithDefaults() *CapabilityDomainPolicyRequirement {
	this := CapabilityDomainPolicyRequirement{}
	var classId string = "capability.DomainPolicyRequirement"
	this.ClassId = classId
	var objectType string = "capability.DomainPolicyRequirement"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityDomainPolicyRequirement) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityDomainPolicyRequirement) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityDomainPolicyRequirement) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.DomainPolicyRequirement" of the ClassId field.
func (o *CapabilityDomainPolicyRequirement) GetDefaultClassId() interface{} {
	return "capability.DomainPolicyRequirement"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityDomainPolicyRequirement) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityDomainPolicyRequirement) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityDomainPolicyRequirement) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.DomainPolicyRequirement" of the ObjectType field.
func (o *CapabilityDomainPolicyRequirement) GetDefaultObjectType() interface{} {
	return "capability.DomainPolicyRequirement"
}

// GetCertificatePolicyConstraints returns the CertificatePolicyConstraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityDomainPolicyRequirement) GetCertificatePolicyConstraints() CapabilityCertificatePropertyConstraints {
	if o == nil || IsNil(o.CertificatePolicyConstraints.Get()) {
		var ret CapabilityCertificatePropertyConstraints
		return ret
	}
	return *o.CertificatePolicyConstraints.Get()
}

// GetCertificatePolicyConstraintsOk returns a tuple with the CertificatePolicyConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityDomainPolicyRequirement) GetCertificatePolicyConstraintsOk() (*CapabilityCertificatePropertyConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificatePolicyConstraints.Get(), o.CertificatePolicyConstraints.IsSet()
}

// HasCertificatePolicyConstraints returns a boolean if a field has been set.
func (o *CapabilityDomainPolicyRequirement) HasCertificatePolicyConstraints() bool {
	if o != nil && o.CertificatePolicyConstraints.IsSet() {
		return true
	}

	return false
}

// SetCertificatePolicyConstraints gets a reference to the given NullableCapabilityCertificatePropertyConstraints and assigns it to the CertificatePolicyConstraints field.
func (o *CapabilityDomainPolicyRequirement) SetCertificatePolicyConstraints(v CapabilityCertificatePropertyConstraints) {
	o.CertificatePolicyConstraints.Set(&v)
}

// SetCertificatePolicyConstraintsNil sets the value for CertificatePolicyConstraints to be an explicit nil
func (o *CapabilityDomainPolicyRequirement) SetCertificatePolicyConstraintsNil() {
	o.CertificatePolicyConstraints.Set(nil)
}

// UnsetCertificatePolicyConstraints ensures that no value is present for CertificatePolicyConstraints, not even an explicit nil
func (o *CapabilityDomainPolicyRequirement) UnsetCertificatePolicyConstraints() {
	o.CertificatePolicyConstraints.Unset()
}

// GetLdapConstraints returns the LdapConstraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityDomainPolicyRequirement) GetLdapConstraints() CapabilityLdapBasePropertyConstraints {
	if o == nil || IsNil(o.LdapConstraints.Get()) {
		var ret CapabilityLdapBasePropertyConstraints
		return ret
	}
	return *o.LdapConstraints.Get()
}

// GetLdapConstraintsOk returns a tuple with the LdapConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityDomainPolicyRequirement) GetLdapConstraintsOk() (*CapabilityLdapBasePropertyConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.LdapConstraints.Get(), o.LdapConstraints.IsSet()
}

// HasLdapConstraints returns a boolean if a field has been set.
func (o *CapabilityDomainPolicyRequirement) HasLdapConstraints() bool {
	if o != nil && o.LdapConstraints.IsSet() {
		return true
	}

	return false
}

// SetLdapConstraints gets a reference to the given NullableCapabilityLdapBasePropertyConstraints and assigns it to the LdapConstraints field.
func (o *CapabilityDomainPolicyRequirement) SetLdapConstraints(v CapabilityLdapBasePropertyConstraints) {
	o.LdapConstraints.Set(&v)
}

// SetLdapConstraintsNil sets the value for LdapConstraints to be an explicit nil
func (o *CapabilityDomainPolicyRequirement) SetLdapConstraintsNil() {
	o.LdapConstraints.Set(nil)
}

// UnsetLdapConstraints ensures that no value is present for LdapConstraints, not even an explicit nil
func (o *CapabilityDomainPolicyRequirement) UnsetLdapConstraints() {
	o.LdapConstraints.Unset()
}

// GetMinBundleVersion returns the MinBundleVersion field value if set, zero value otherwise.
func (o *CapabilityDomainPolicyRequirement) GetMinBundleVersion() string {
	if o == nil || IsNil(o.MinBundleVersion) {
		var ret string
		return ret
	}
	return *o.MinBundleVersion
}

// GetMinBundleVersionOk returns a tuple with the MinBundleVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDomainPolicyRequirement) GetMinBundleVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinBundleVersion) {
		return nil, false
	}
	return o.MinBundleVersion, true
}

// HasMinBundleVersion returns a boolean if a field has been set.
func (o *CapabilityDomainPolicyRequirement) HasMinBundleVersion() bool {
	if o != nil && !IsNil(o.MinBundleVersion) {
		return true
	}

	return false
}

// SetMinBundleVersion gets a reference to the given string and assigns it to the MinBundleVersion field.
func (o *CapabilityDomainPolicyRequirement) SetMinBundleVersion(v string) {
	o.MinBundleVersion = &v
}

// GetMinVersion returns the MinVersion field value if set, zero value otherwise.
func (o *CapabilityDomainPolicyRequirement) GetMinVersion() string {
	if o == nil || IsNil(o.MinVersion) {
		var ret string
		return ret
	}
	return *o.MinVersion
}

// GetMinVersionOk returns a tuple with the MinVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDomainPolicyRequirement) GetMinVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinVersion) {
		return nil, false
	}
	return o.MinVersion, true
}

// HasMinVersion returns a boolean if a field has been set.
func (o *CapabilityDomainPolicyRequirement) HasMinVersion() bool {
	if o != nil && !IsNil(o.MinVersion) {
		return true
	}

	return false
}

// SetMinVersion gets a reference to the given string and assigns it to the MinVersion field.
func (o *CapabilityDomainPolicyRequirement) SetMinVersion(v string) {
	o.MinVersion = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *CapabilityDomainPolicyRequirement) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDomainPolicyRequirement) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *CapabilityDomainPolicyRequirement) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *CapabilityDomainPolicyRequirement) SetModel(v string) {
	o.Model = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *CapabilityDomainPolicyRequirement) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName) {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDomainPolicyRequirement) GetPolicyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyName) {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *CapabilityDomainPolicyRequirement) HasPolicyName() bool {
	if o != nil && !IsNil(o.PolicyName) {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *CapabilityDomainPolicyRequirement) SetPolicyName(v string) {
	o.PolicyName = &v
}

func (o CapabilityDomainPolicyRequirement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityDomainPolicyRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.CertificatePolicyConstraints.IsSet() {
		toSerialize["CertificatePolicyConstraints"] = o.CertificatePolicyConstraints.Get()
	}
	if o.LdapConstraints.IsSet() {
		toSerialize["LdapConstraints"] = o.LdapConstraints.Get()
	}
	if !IsNil(o.MinBundleVersion) {
		toSerialize["MinBundleVersion"] = o.MinBundleVersion
	}
	if !IsNil(o.MinVersion) {
		toSerialize["MinVersion"] = o.MinVersion
	}
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	if !IsNil(o.PolicyName) {
		toSerialize["PolicyName"] = o.PolicyName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityDomainPolicyRequirement) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityDomainPolicyRequirementWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType                   string                                           `json:"ObjectType"`
		CertificatePolicyConstraints NullableCapabilityCertificatePropertyConstraints `json:"CertificatePolicyConstraints,omitempty"`
		LdapConstraints              NullableCapabilityLdapBasePropertyConstraints    `json:"LdapConstraints,omitempty"`
		// Minimum Bundle version from which it is supported.
		MinBundleVersion *string `json:"MinBundleVersion,omitempty"`
		// Minimum Version from which policy is supported.
		MinVersion *string `json:"MinVersion,omitempty"`
		// Type of the platform for which version compatibility is specified. Example - 3GFI, 4GFI, etc.
		Model *string `json:"Model,omitempty"`
		// Policy Name for which version compatibility is specified. Example - snmp.Policy, ldap.Policy.
		PolicyName *string `json:"PolicyName,omitempty"`
	}

	varCapabilityDomainPolicyRequirementWithoutEmbeddedStruct := CapabilityDomainPolicyRequirementWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityDomainPolicyRequirementWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityDomainPolicyRequirement := _CapabilityDomainPolicyRequirement{}
		varCapabilityDomainPolicyRequirement.ClassId = varCapabilityDomainPolicyRequirementWithoutEmbeddedStruct.ClassId
		varCapabilityDomainPolicyRequirement.ObjectType = varCapabilityDomainPolicyRequirementWithoutEmbeddedStruct.ObjectType
		varCapabilityDomainPolicyRequirement.CertificatePolicyConstraints = varCapabilityDomainPolicyRequirementWithoutEmbeddedStruct.CertificatePolicyConstraints
		varCapabilityDomainPolicyRequirement.LdapConstraints = varCapabilityDomainPolicyRequirementWithoutEmbeddedStruct.LdapConstraints
		varCapabilityDomainPolicyRequirement.MinBundleVersion = varCapabilityDomainPolicyRequirementWithoutEmbeddedStruct.MinBundleVersion
		varCapabilityDomainPolicyRequirement.MinVersion = varCapabilityDomainPolicyRequirementWithoutEmbeddedStruct.MinVersion
		varCapabilityDomainPolicyRequirement.Model = varCapabilityDomainPolicyRequirementWithoutEmbeddedStruct.Model
		varCapabilityDomainPolicyRequirement.PolicyName = varCapabilityDomainPolicyRequirementWithoutEmbeddedStruct.PolicyName
		*o = CapabilityDomainPolicyRequirement(varCapabilityDomainPolicyRequirement)
	} else {
		return err
	}

	varCapabilityDomainPolicyRequirement := _CapabilityDomainPolicyRequirement{}

	err = json.Unmarshal(data, &varCapabilityDomainPolicyRequirement)
	if err == nil {
		o.CapabilityCapability = varCapabilityDomainPolicyRequirement.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CertificatePolicyConstraints")
		delete(additionalProperties, "LdapConstraints")
		delete(additionalProperties, "MinBundleVersion")
		delete(additionalProperties, "MinVersion")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "PolicyName")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableCapabilityDomainPolicyRequirement struct {
	value *CapabilityDomainPolicyRequirement
	isSet bool
}

func (v NullableCapabilityDomainPolicyRequirement) Get() *CapabilityDomainPolicyRequirement {
	return v.value
}

func (v *NullableCapabilityDomainPolicyRequirement) Set(val *CapabilityDomainPolicyRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityDomainPolicyRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityDomainPolicyRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityDomainPolicyRequirement(val *CapabilityDomainPolicyRequirement) *NullableCapabilityDomainPolicyRequirement {
	return &NullableCapabilityDomainPolicyRequirement{value: val, isSet: true}
}

func (v NullableCapabilityDomainPolicyRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityDomainPolicyRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}