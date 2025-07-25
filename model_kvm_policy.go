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

// checks if the KvmPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KvmPolicy{}

// KvmPolicy Policy to configure KVM Launch settings.
type KvmPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If enabled, displays KVM session on any monitor attached to the server.
	EnableLocalServerVideo *bool `json:"EnableLocalServerVideo,omitempty"`
	// If enabled, encrypts all video information sent through KVM. Please note that this can no longer be disabled for servers running versions 4.2 and above.
	EnableVideoEncryption *bool `json:"EnableVideoEncryption,omitempty"`
	// State of the vKVM service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// The maximum number of concurrent KVM sessions allowed.
	MaximumSessions *int64 `json:"MaximumSessions,omitempty"`
	// The port used for KVM communication.
	RemotePort *int64 `json:"RemotePort,omitempty"`
	// Enables Tunneled vKVM on the endpoint. Applicable only for Device Connectors that support Tunneled vKVM.
	TunneledKvmEnabled *bool                                        `json:"TunneledKvmEnabled,omitempty"`
	Organization       NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	// Deprecated
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmPolicy KvmPolicy

// NewKvmPolicy instantiates a new KvmPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmPolicy(classId string, objectType string) *KvmPolicy {
	this := KvmPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enableLocalServerVideo bool = true
	this.EnableLocalServerVideo = &enableLocalServerVideo
	var enableVideoEncryption bool = true
	this.EnableVideoEncryption = &enableVideoEncryption
	var enabled bool = true
	this.Enabled = &enabled
	var maximumSessions int64 = 4
	this.MaximumSessions = &maximumSessions
	var remotePort int64 = 2068
	this.RemotePort = &remotePort
	var tunneledKvmEnabled bool = false
	this.TunneledKvmEnabled = &tunneledKvmEnabled
	return &this
}

// NewKvmPolicyWithDefaults instantiates a new KvmPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmPolicyWithDefaults() *KvmPolicy {
	this := KvmPolicy{}
	var classId string = "kvm.Policy"
	this.ClassId = classId
	var objectType string = "kvm.Policy"
	this.ObjectType = objectType
	var enableLocalServerVideo bool = true
	this.EnableLocalServerVideo = &enableLocalServerVideo
	var enableVideoEncryption bool = true
	this.EnableVideoEncryption = &enableVideoEncryption
	var enabled bool = true
	this.Enabled = &enabled
	var maximumSessions int64 = 4
	this.MaximumSessions = &maximumSessions
	var remotePort int64 = 2068
	this.RemotePort = &remotePort
	var tunneledKvmEnabled bool = false
	this.TunneledKvmEnabled = &tunneledKvmEnabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kvm.Policy" of the ClassId field.
func (o *KvmPolicy) GetDefaultClassId() interface{} {
	return "kvm.Policy"
}

// GetObjectType returns the ObjectType field value
func (o *KvmPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kvm.Policy" of the ObjectType field.
func (o *KvmPolicy) GetDefaultObjectType() interface{} {
	return "kvm.Policy"
}

// GetEnableLocalServerVideo returns the EnableLocalServerVideo field value if set, zero value otherwise.
func (o *KvmPolicy) GetEnableLocalServerVideo() bool {
	if o == nil || IsNil(o.EnableLocalServerVideo) {
		var ret bool
		return ret
	}
	return *o.EnableLocalServerVideo
}

// GetEnableLocalServerVideoOk returns a tuple with the EnableLocalServerVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetEnableLocalServerVideoOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableLocalServerVideo) {
		return nil, false
	}
	return o.EnableLocalServerVideo, true
}

// HasEnableLocalServerVideo returns a boolean if a field has been set.
func (o *KvmPolicy) HasEnableLocalServerVideo() bool {
	if o != nil && !IsNil(o.EnableLocalServerVideo) {
		return true
	}

	return false
}

// SetEnableLocalServerVideo gets a reference to the given bool and assigns it to the EnableLocalServerVideo field.
func (o *KvmPolicy) SetEnableLocalServerVideo(v bool) {
	o.EnableLocalServerVideo = &v
}

// GetEnableVideoEncryption returns the EnableVideoEncryption field value if set, zero value otherwise.
func (o *KvmPolicy) GetEnableVideoEncryption() bool {
	if o == nil || IsNil(o.EnableVideoEncryption) {
		var ret bool
		return ret
	}
	return *o.EnableVideoEncryption
}

// GetEnableVideoEncryptionOk returns a tuple with the EnableVideoEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetEnableVideoEncryptionOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableVideoEncryption) {
		return nil, false
	}
	return o.EnableVideoEncryption, true
}

// HasEnableVideoEncryption returns a boolean if a field has been set.
func (o *KvmPolicy) HasEnableVideoEncryption() bool {
	if o != nil && !IsNil(o.EnableVideoEncryption) {
		return true
	}

	return false
}

// SetEnableVideoEncryption gets a reference to the given bool and assigns it to the EnableVideoEncryption field.
func (o *KvmPolicy) SetEnableVideoEncryption(v bool) {
	o.EnableVideoEncryption = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KvmPolicy) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KvmPolicy) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KvmPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaximumSessions returns the MaximumSessions field value if set, zero value otherwise.
func (o *KvmPolicy) GetMaximumSessions() int64 {
	if o == nil || IsNil(o.MaximumSessions) {
		var ret int64
		return ret
	}
	return *o.MaximumSessions
}

// GetMaximumSessionsOk returns a tuple with the MaximumSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetMaximumSessionsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumSessions) {
		return nil, false
	}
	return o.MaximumSessions, true
}

// HasMaximumSessions returns a boolean if a field has been set.
func (o *KvmPolicy) HasMaximumSessions() bool {
	if o != nil && !IsNil(o.MaximumSessions) {
		return true
	}

	return false
}

// SetMaximumSessions gets a reference to the given int64 and assigns it to the MaximumSessions field.
func (o *KvmPolicy) SetMaximumSessions(v int64) {
	o.MaximumSessions = &v
}

// GetRemotePort returns the RemotePort field value if set, zero value otherwise.
func (o *KvmPolicy) GetRemotePort() int64 {
	if o == nil || IsNil(o.RemotePort) {
		var ret int64
		return ret
	}
	return *o.RemotePort
}

// GetRemotePortOk returns a tuple with the RemotePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetRemotePortOk() (*int64, bool) {
	if o == nil || IsNil(o.RemotePort) {
		return nil, false
	}
	return o.RemotePort, true
}

// HasRemotePort returns a boolean if a field has been set.
func (o *KvmPolicy) HasRemotePort() bool {
	if o != nil && !IsNil(o.RemotePort) {
		return true
	}

	return false
}

// SetRemotePort gets a reference to the given int64 and assigns it to the RemotePort field.
func (o *KvmPolicy) SetRemotePort(v int64) {
	o.RemotePort = &v
}

// GetTunneledKvmEnabled returns the TunneledKvmEnabled field value if set, zero value otherwise.
func (o *KvmPolicy) GetTunneledKvmEnabled() bool {
	if o == nil || IsNil(o.TunneledKvmEnabled) {
		var ret bool
		return ret
	}
	return *o.TunneledKvmEnabled
}

// GetTunneledKvmEnabledOk returns a tuple with the TunneledKvmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicy) GetTunneledKvmEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TunneledKvmEnabled) {
		return nil, false
	}
	return o.TunneledKvmEnabled, true
}

// HasTunneledKvmEnabled returns a boolean if a field has been set.
func (o *KvmPolicy) HasTunneledKvmEnabled() bool {
	if o != nil && !IsNil(o.TunneledKvmEnabled) {
		return true
	}

	return false
}

// SetTunneledKvmEnabled gets a reference to the given bool and assigns it to the TunneledKvmEnabled field.
func (o *KvmPolicy) SetTunneledKvmEnabled(v bool) {
	o.TunneledKvmEnabled = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *KvmPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KvmPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *KvmPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *KvmPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *KvmPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *KvmPolicy) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *KvmPolicy) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
// Deprecated
func (o *KvmPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o KvmPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KvmPolicy) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EnableLocalServerVideo) {
		toSerialize["EnableLocalServerVideo"] = o.EnableLocalServerVideo
	}
	if !IsNil(o.EnableVideoEncryption) {
		toSerialize["EnableVideoEncryption"] = o.EnableVideoEncryption
	}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.MaximumSessions) {
		toSerialize["MaximumSessions"] = o.MaximumSessions
	}
	if !IsNil(o.RemotePort) {
		toSerialize["RemotePort"] = o.RemotePort
	}
	if !IsNil(o.TunneledKvmEnabled) {
		toSerialize["TunneledKvmEnabled"] = o.TunneledKvmEnabled
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

func (o *KvmPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type KvmPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// If enabled, displays KVM session on any monitor attached to the server.
		EnableLocalServerVideo *bool `json:"EnableLocalServerVideo,omitempty"`
		// If enabled, encrypts all video information sent through KVM. Please note that this can no longer be disabled for servers running versions 4.2 and above.
		EnableVideoEncryption *bool `json:"EnableVideoEncryption,omitempty"`
		// State of the vKVM service on the endpoint.
		Enabled *bool `json:"Enabled,omitempty"`
		// The maximum number of concurrent KVM sessions allowed.
		MaximumSessions *int64 `json:"MaximumSessions,omitempty"`
		// The port used for KVM communication.
		RemotePort *int64 `json:"RemotePort,omitempty"`
		// Enables Tunneled vKVM on the endpoint. Applicable only for Device Connectors that support Tunneled vKVM.
		TunneledKvmEnabled *bool                                        `json:"TunneledKvmEnabled,omitempty"`
		Organization       NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		// Deprecated
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varKvmPolicyWithoutEmbeddedStruct := KvmPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKvmPolicyWithoutEmbeddedStruct)
	if err == nil {
		varKvmPolicy := _KvmPolicy{}
		varKvmPolicy.ClassId = varKvmPolicyWithoutEmbeddedStruct.ClassId
		varKvmPolicy.ObjectType = varKvmPolicyWithoutEmbeddedStruct.ObjectType
		varKvmPolicy.EnableLocalServerVideo = varKvmPolicyWithoutEmbeddedStruct.EnableLocalServerVideo
		varKvmPolicy.EnableVideoEncryption = varKvmPolicyWithoutEmbeddedStruct.EnableVideoEncryption
		varKvmPolicy.Enabled = varKvmPolicyWithoutEmbeddedStruct.Enabled
		varKvmPolicy.MaximumSessions = varKvmPolicyWithoutEmbeddedStruct.MaximumSessions
		varKvmPolicy.RemotePort = varKvmPolicyWithoutEmbeddedStruct.RemotePort
		varKvmPolicy.TunneledKvmEnabled = varKvmPolicyWithoutEmbeddedStruct.TunneledKvmEnabled
		varKvmPolicy.Organization = varKvmPolicyWithoutEmbeddedStruct.Organization
		varKvmPolicy.Profiles = varKvmPolicyWithoutEmbeddedStruct.Profiles
		*o = KvmPolicy(varKvmPolicy)
	} else {
		return err
	}

	varKvmPolicy := _KvmPolicy{}

	err = json.Unmarshal(data, &varKvmPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varKvmPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnableLocalServerVideo")
		delete(additionalProperties, "EnableVideoEncryption")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "MaximumSessions")
		delete(additionalProperties, "RemotePort")
		delete(additionalProperties, "TunneledKvmEnabled")
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

type NullableKvmPolicy struct {
	value *KvmPolicy
	isSet bool
}

func (v NullableKvmPolicy) Get() *KvmPolicy {
	return v.value
}

func (v *NullableKvmPolicy) Set(val *KvmPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmPolicy(val *KvmPolicy) *NullableKvmPolicy {
	return &NullableKvmPolicy{value: val, isSet: true}
}

func (v NullableKvmPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
