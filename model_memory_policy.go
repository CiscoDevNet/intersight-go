/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-18775
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

// checks if the MemoryPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemoryPolicy{}

// MemoryPolicy Memory Management policy models a configuration that can be applied to servers to manage memory features.
type MemoryPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable DIMM Blocklisting on the server.
	EnableDimmBlocklisting *bool                                        `json:"EnableDimmBlocklisting,omitempty"`
	Organization           NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemoryPolicy MemoryPolicy

// NewMemoryPolicy instantiates a new MemoryPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPolicy(classId string, objectType string) *MemoryPolicy {
	this := MemoryPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enableDimmBlocklisting bool = false
	this.EnableDimmBlocklisting = &enableDimmBlocklisting
	return &this
}

// NewMemoryPolicyWithDefaults instantiates a new MemoryPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPolicyWithDefaults() *MemoryPolicy {
	this := MemoryPolicy{}
	var classId string = "memory.Policy"
	this.ClassId = classId
	var objectType string = "memory.Policy"
	this.ObjectType = objectType
	var enableDimmBlocklisting bool = false
	this.EnableDimmBlocklisting = &enableDimmBlocklisting
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "memory.Policy" of the ClassId field.
func (o *MemoryPolicy) GetDefaultClassId() interface{} {
	return "memory.Policy"
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "memory.Policy" of the ObjectType field.
func (o *MemoryPolicy) GetDefaultObjectType() interface{} {
	return "memory.Policy"
}

// GetEnableDimmBlocklisting returns the EnableDimmBlocklisting field value if set, zero value otherwise.
func (o *MemoryPolicy) GetEnableDimmBlocklisting() bool {
	if o == nil || IsNil(o.EnableDimmBlocklisting) {
		var ret bool
		return ret
	}
	return *o.EnableDimmBlocklisting
}

// GetEnableDimmBlocklistingOk returns a tuple with the EnableDimmBlocklisting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPolicy) GetEnableDimmBlocklistingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDimmBlocklisting) {
		return nil, false
	}
	return o.EnableDimmBlocklisting, true
}

// HasEnableDimmBlocklisting returns a boolean if a field has been set.
func (o *MemoryPolicy) HasEnableDimmBlocklisting() bool {
	if o != nil && !IsNil(o.EnableDimmBlocklisting) {
		return true
	}

	return false
}

// SetEnableDimmBlocklisting gets a reference to the given bool and assigns it to the EnableDimmBlocklisting field.
func (o *MemoryPolicy) SetEnableDimmBlocklisting(v bool) {
	o.EnableDimmBlocklisting = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *MemoryPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *MemoryPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *MemoryPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *MemoryPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPolicy) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *MemoryPolicy) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *MemoryPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o MemoryPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemoryPolicy) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EnableDimmBlocklisting) {
		toSerialize["EnableDimmBlocklisting"] = o.EnableDimmBlocklisting
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemoryPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type MemoryPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enable DIMM Blocklisting on the server.
		EnableDimmBlocklisting *bool                                        `json:"EnableDimmBlocklisting,omitempty"`
		Organization           NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varMemoryPolicyWithoutEmbeddedStruct := MemoryPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMemoryPolicyWithoutEmbeddedStruct)
	if err == nil {
		varMemoryPolicy := _MemoryPolicy{}
		varMemoryPolicy.ClassId = varMemoryPolicyWithoutEmbeddedStruct.ClassId
		varMemoryPolicy.ObjectType = varMemoryPolicyWithoutEmbeddedStruct.ObjectType
		varMemoryPolicy.EnableDimmBlocklisting = varMemoryPolicyWithoutEmbeddedStruct.EnableDimmBlocklisting
		varMemoryPolicy.Organization = varMemoryPolicyWithoutEmbeddedStruct.Organization
		varMemoryPolicy.Profiles = varMemoryPolicyWithoutEmbeddedStruct.Profiles
		*o = MemoryPolicy(varMemoryPolicy)
	} else {
		return err
	}

	varMemoryPolicy := _MemoryPolicy{}

	err = json.Unmarshal(data, &varMemoryPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varMemoryPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnableDimmBlocklisting")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

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

type NullableMemoryPolicy struct {
	value *MemoryPolicy
	isSet bool
}

func (v NullableMemoryPolicy) Get() *MemoryPolicy {
	return v.value
}

func (v *NullableMemoryPolicy) Set(val *MemoryPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPolicy(val *MemoryPolicy) *NullableMemoryPolicy {
	return &NullableMemoryPolicy{value: val, isSet: true}
}

func (v NullableMemoryPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}