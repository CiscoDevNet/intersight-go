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

// checks if the ChassisIomProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChassisIomProfile{}

// ChassisIomProfile A profile specifying configuration settings for IOM.
type ChassisIomProfile struct {
	PolicyAbstractConfigProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                     `json:"ObjectType"`
	ConfigChanges NullablePolicyConfigChange `json:"ConfigChanges,omitempty"`
	// IOM in chassis for which IOM profile is applicable. or which is attached to a Fabric Interconnect managed by Intersight. * `IOMA` - IOM on left side of chassis. * `IOMB` - IOM on right side of chassis.
	IomEntity *string `json:"IomEntity,omitempty"`
	// An array of relationships to chassisConfigChangeDetail resources.
	ConfigChangeDetails  []ChassisConfigChangeDetailRelationship      `json:"ConfigChangeDetails,omitempty"`
	ConfigResult         NullableChassisConfigResultRelationship      `json:"ConfigResult,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	Profile              NullableChassisBaseProfileRelationship       `json:"Profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChassisIomProfile ChassisIomProfile

// NewChassisIomProfile instantiates a new ChassisIomProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChassisIomProfile(classId string, objectType string) *ChassisIomProfile {
	this := ChassisIomProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	var action string = "No-op"
	this.Action = &action
	var iomEntity string = "IOMA"
	this.IomEntity = &iomEntity
	return &this
}

// NewChassisIomProfileWithDefaults instantiates a new ChassisIomProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChassisIomProfileWithDefaults() *ChassisIomProfile {
	this := ChassisIomProfile{}
	var classId string = "chassis.IomProfile"
	this.ClassId = classId
	var objectType string = "chassis.IomProfile"
	this.ObjectType = objectType
	var iomEntity string = "IOMA"
	this.IomEntity = &iomEntity
	return &this
}

// GetClassId returns the ClassId field value
func (o *ChassisIomProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ChassisIomProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ChassisIomProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "chassis.IomProfile" of the ClassId field.
func (o *ChassisIomProfile) GetDefaultClassId() interface{} {
	return "chassis.IomProfile"
}

// GetObjectType returns the ObjectType field value
func (o *ChassisIomProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ChassisIomProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ChassisIomProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "chassis.IomProfile" of the ObjectType field.
func (o *ChassisIomProfile) GetDefaultObjectType() interface{} {
	return "chassis.IomProfile"
}

// GetConfigChanges returns the ConfigChanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisIomProfile) GetConfigChanges() PolicyConfigChange {
	if o == nil || IsNil(o.ConfigChanges.Get()) {
		var ret PolicyConfigChange
		return ret
	}
	return *o.ConfigChanges.Get()
}

// GetConfigChangesOk returns a tuple with the ConfigChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisIomProfile) GetConfigChangesOk() (*PolicyConfigChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChanges.Get(), o.ConfigChanges.IsSet()
}

// HasConfigChanges returns a boolean if a field has been set.
func (o *ChassisIomProfile) HasConfigChanges() bool {
	if o != nil && o.ConfigChanges.IsSet() {
		return true
	}

	return false
}

// SetConfigChanges gets a reference to the given NullablePolicyConfigChange and assigns it to the ConfigChanges field.
func (o *ChassisIomProfile) SetConfigChanges(v PolicyConfigChange) {
	o.ConfigChanges.Set(&v)
}

// SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil
func (o *ChassisIomProfile) SetConfigChangesNil() {
	o.ConfigChanges.Set(nil)
}

// UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
func (o *ChassisIomProfile) UnsetConfigChanges() {
	o.ConfigChanges.Unset()
}

// GetIomEntity returns the IomEntity field value if set, zero value otherwise.
func (o *ChassisIomProfile) GetIomEntity() string {
	if o == nil || IsNil(o.IomEntity) {
		var ret string
		return ret
	}
	return *o.IomEntity
}

// GetIomEntityOk returns a tuple with the IomEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisIomProfile) GetIomEntityOk() (*string, bool) {
	if o == nil || IsNil(o.IomEntity) {
		return nil, false
	}
	return o.IomEntity, true
}

// HasIomEntity returns a boolean if a field has been set.
func (o *ChassisIomProfile) HasIomEntity() bool {
	if o != nil && !IsNil(o.IomEntity) {
		return true
	}

	return false
}

// SetIomEntity gets a reference to the given string and assigns it to the IomEntity field.
func (o *ChassisIomProfile) SetIomEntity(v string) {
	o.IomEntity = &v
}

// GetConfigChangeDetails returns the ConfigChangeDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisIomProfile) GetConfigChangeDetails() []ChassisConfigChangeDetailRelationship {
	if o == nil {
		var ret []ChassisConfigChangeDetailRelationship
		return ret
	}
	return o.ConfigChangeDetails
}

// GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisIomProfile) GetConfigChangeDetailsOk() ([]ChassisConfigChangeDetailRelationship, bool) {
	if o == nil || IsNil(o.ConfigChangeDetails) {
		return nil, false
	}
	return o.ConfigChangeDetails, true
}

// HasConfigChangeDetails returns a boolean if a field has been set.
func (o *ChassisIomProfile) HasConfigChangeDetails() bool {
	if o != nil && !IsNil(o.ConfigChangeDetails) {
		return true
	}

	return false
}

// SetConfigChangeDetails gets a reference to the given []ChassisConfigChangeDetailRelationship and assigns it to the ConfigChangeDetails field.
func (o *ChassisIomProfile) SetConfigChangeDetails(v []ChassisConfigChangeDetailRelationship) {
	o.ConfigChangeDetails = v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisIomProfile) GetConfigResult() ChassisConfigResultRelationship {
	if o == nil || IsNil(o.ConfigResult.Get()) {
		var ret ChassisConfigResultRelationship
		return ret
	}
	return *o.ConfigResult.Get()
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisIomProfile) GetConfigResultOk() (*ChassisConfigResultRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigResult.Get(), o.ConfigResult.IsSet()
}

// HasConfigResult returns a boolean if a field has been set.
func (o *ChassisIomProfile) HasConfigResult() bool {
	if o != nil && o.ConfigResult.IsSet() {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given NullableChassisConfigResultRelationship and assigns it to the ConfigResult field.
func (o *ChassisIomProfile) SetConfigResult(v ChassisConfigResultRelationship) {
	o.ConfigResult.Set(&v)
}

// SetConfigResultNil sets the value for ConfigResult to be an explicit nil
func (o *ChassisIomProfile) SetConfigResultNil() {
	o.ConfigResult.Set(nil)
}

// UnsetConfigResult ensures that no value is present for ConfigResult, not even an explicit nil
func (o *ChassisIomProfile) UnsetConfigResult() {
	o.ConfigResult.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisIomProfile) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisIomProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *ChassisIomProfile) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ChassisIomProfile) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *ChassisIomProfile) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *ChassisIomProfile) UnsetOrganization() {
	o.Organization.Unset()
}

// GetProfile returns the Profile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisIomProfile) GetProfile() ChassisBaseProfileRelationship {
	if o == nil || IsNil(o.Profile.Get()) {
		var ret ChassisBaseProfileRelationship
		return ret
	}
	return *o.Profile.Get()
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisIomProfile) GetProfileOk() (*ChassisBaseProfileRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profile.Get(), o.Profile.IsSet()
}

// HasProfile returns a boolean if a field has been set.
func (o *ChassisIomProfile) HasProfile() bool {
	if o != nil && o.Profile.IsSet() {
		return true
	}

	return false
}

// SetProfile gets a reference to the given NullableChassisBaseProfileRelationship and assigns it to the Profile field.
func (o *ChassisIomProfile) SetProfile(v ChassisBaseProfileRelationship) {
	o.Profile.Set(&v)
}

// SetProfileNil sets the value for Profile to be an explicit nil
func (o *ChassisIomProfile) SetProfileNil() {
	o.Profile.Set(nil)
}

// UnsetProfile ensures that no value is present for Profile, not even an explicit nil
func (o *ChassisIomProfile) UnsetProfile() {
	o.Profile.Unset()
}

func (o ChassisIomProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChassisIomProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigProfile, errPolicyAbstractConfigProfile := json.Marshal(o.PolicyAbstractConfigProfile)
	if errPolicyAbstractConfigProfile != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigProfile
	}
	errPolicyAbstractConfigProfile = json.Unmarshal([]byte(serializedPolicyAbstractConfigProfile), &toSerialize)
	if errPolicyAbstractConfigProfile != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigProfile
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.ConfigChanges.IsSet() {
		toSerialize["ConfigChanges"] = o.ConfigChanges.Get()
	}
	if !IsNil(o.IomEntity) {
		toSerialize["IomEntity"] = o.IomEntity
	}
	if o.ConfigChangeDetails != nil {
		toSerialize["ConfigChangeDetails"] = o.ConfigChangeDetails
	}
	if o.ConfigResult.IsSet() {
		toSerialize["ConfigResult"] = o.ConfigResult.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.Profile.IsSet() {
		toSerialize["Profile"] = o.Profile.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChassisIomProfile) UnmarshalJSON(data []byte) (err error) {
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
	type ChassisIomProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string                     `json:"ObjectType"`
		ConfigChanges NullablePolicyConfigChange `json:"ConfigChanges,omitempty"`
		// IOM in chassis for which IOM profile is applicable. or which is attached to a Fabric Interconnect managed by Intersight. * `IOMA` - IOM on left side of chassis. * `IOMB` - IOM on right side of chassis.
		IomEntity *string `json:"IomEntity,omitempty"`
		// An array of relationships to chassisConfigChangeDetail resources.
		ConfigChangeDetails []ChassisConfigChangeDetailRelationship      `json:"ConfigChangeDetails,omitempty"`
		ConfigResult        NullableChassisConfigResultRelationship      `json:"ConfigResult,omitempty"`
		Organization        NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		Profile             NullableChassisBaseProfileRelationship       `json:"Profile,omitempty"`
	}

	varChassisIomProfileWithoutEmbeddedStruct := ChassisIomProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varChassisIomProfileWithoutEmbeddedStruct)
	if err == nil {
		varChassisIomProfile := _ChassisIomProfile{}
		varChassisIomProfile.ClassId = varChassisIomProfileWithoutEmbeddedStruct.ClassId
		varChassisIomProfile.ObjectType = varChassisIomProfileWithoutEmbeddedStruct.ObjectType
		varChassisIomProfile.ConfigChanges = varChassisIomProfileWithoutEmbeddedStruct.ConfigChanges
		varChassisIomProfile.IomEntity = varChassisIomProfileWithoutEmbeddedStruct.IomEntity
		varChassisIomProfile.ConfigChangeDetails = varChassisIomProfileWithoutEmbeddedStruct.ConfigChangeDetails
		varChassisIomProfile.ConfigResult = varChassisIomProfileWithoutEmbeddedStruct.ConfigResult
		varChassisIomProfile.Organization = varChassisIomProfileWithoutEmbeddedStruct.Organization
		varChassisIomProfile.Profile = varChassisIomProfileWithoutEmbeddedStruct.Profile
		*o = ChassisIomProfile(varChassisIomProfile)
	} else {
		return err
	}

	varChassisIomProfile := _ChassisIomProfile{}

	err = json.Unmarshal(data, &varChassisIomProfile)
	if err == nil {
		o.PolicyAbstractConfigProfile = varChassisIomProfile.PolicyAbstractConfigProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigChanges")
		delete(additionalProperties, "IomEntity")
		delete(additionalProperties, "ConfigChangeDetails")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profile")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigProfile := reflect.ValueOf(o.PolicyAbstractConfigProfile)
		for i := 0; i < reflectPolicyAbstractConfigProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigProfile.Type().Field(i)

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

type NullableChassisIomProfile struct {
	value *ChassisIomProfile
	isSet bool
}

func (v NullableChassisIomProfile) Get() *ChassisIomProfile {
	return v.value
}

func (v *NullableChassisIomProfile) Set(val *ChassisIomProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisIomProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisIomProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisIomProfile(val *ChassisIomProfile) *NullableChassisIomProfile {
	return &NullableChassisIomProfile{value: val, isSet: true}
}

func (v NullableChassisIomProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisIomProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
