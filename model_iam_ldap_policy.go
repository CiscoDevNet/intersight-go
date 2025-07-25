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

// checks if the IamLdapPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamLdapPolicy{}

// IamLdapPolicy LDAP Policy configurations.
type IamLdapPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType     string                        `json:"ObjectType"`
	BaseProperties NullableIamLdapBaseProperties `json:"BaseProperties,omitempty"`
	DnsParameters  NullableIamLdapDnsParameters  `json:"DnsParameters,omitempty"`
	// Enables DNS to access LDAP servers.
	EnableDns *bool `json:"EnableDns,omitempty"`
	// LDAP server performs authentication.
	Enabled *bool `json:"Enabled,omitempty"`
	// Search precedence between local user database and LDAP user database. * `LocalUserDb` - Precedence is given to local user database while searching. * `LDAPUserDb` - Precedence is given to LADP user database while searching.
	UserSearchPrecedence *string                        `json:"UserSearchPrecedence,omitempty"`
	ApplianceAccount     NullableIamAccountRelationship `json:"ApplianceAccount,omitempty"`
	// An array of relationships to iamLdapGroup resources.
	Groups       []IamLdapGroupRelationship                   `json:"Groups,omitempty"`
	Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	// Deprecated
	Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	// An array of relationships to iamLdapProvider resources.
	Providers            []IamLdapProviderRelationship `json:"Providers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamLdapPolicy IamLdapPolicy

// NewIamLdapPolicy instantiates a new IamLdapPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamLdapPolicy(classId string, objectType string) *IamLdapPolicy {
	this := IamLdapPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var userSearchPrecedence string = "LocalUserDb"
	this.UserSearchPrecedence = &userSearchPrecedence
	return &this
}

// NewIamLdapPolicyWithDefaults instantiates a new IamLdapPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamLdapPolicyWithDefaults() *IamLdapPolicy {
	this := IamLdapPolicy{}
	var classId string = "iam.LdapPolicy"
	this.ClassId = classId
	var objectType string = "iam.LdapPolicy"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var userSearchPrecedence string = "LocalUserDb"
	this.UserSearchPrecedence = &userSearchPrecedence
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamLdapPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamLdapPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamLdapPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.LdapPolicy" of the ClassId field.
func (o *IamLdapPolicy) GetDefaultClassId() interface{} {
	return "iam.LdapPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *IamLdapPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamLdapPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamLdapPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.LdapPolicy" of the ObjectType field.
func (o *IamLdapPolicy) GetDefaultObjectType() interface{} {
	return "iam.LdapPolicy"
}

// GetBaseProperties returns the BaseProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamLdapPolicy) GetBaseProperties() IamLdapBaseProperties {
	if o == nil || IsNil(o.BaseProperties.Get()) {
		var ret IamLdapBaseProperties
		return ret
	}
	return *o.BaseProperties.Get()
}

// GetBasePropertiesOk returns a tuple with the BaseProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamLdapPolicy) GetBasePropertiesOk() (*IamLdapBaseProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseProperties.Get(), o.BaseProperties.IsSet()
}

// HasBaseProperties returns a boolean if a field has been set.
func (o *IamLdapPolicy) HasBaseProperties() bool {
	if o != nil && o.BaseProperties.IsSet() {
		return true
	}

	return false
}

// SetBaseProperties gets a reference to the given NullableIamLdapBaseProperties and assigns it to the BaseProperties field.
func (o *IamLdapPolicy) SetBaseProperties(v IamLdapBaseProperties) {
	o.BaseProperties.Set(&v)
}

// SetBasePropertiesNil sets the value for BaseProperties to be an explicit nil
func (o *IamLdapPolicy) SetBasePropertiesNil() {
	o.BaseProperties.Set(nil)
}

// UnsetBaseProperties ensures that no value is present for BaseProperties, not even an explicit nil
func (o *IamLdapPolicy) UnsetBaseProperties() {
	o.BaseProperties.Unset()
}

// GetDnsParameters returns the DnsParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamLdapPolicy) GetDnsParameters() IamLdapDnsParameters {
	if o == nil || IsNil(o.DnsParameters.Get()) {
		var ret IamLdapDnsParameters
		return ret
	}
	return *o.DnsParameters.Get()
}

// GetDnsParametersOk returns a tuple with the DnsParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamLdapPolicy) GetDnsParametersOk() (*IamLdapDnsParameters, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsParameters.Get(), o.DnsParameters.IsSet()
}

// HasDnsParameters returns a boolean if a field has been set.
func (o *IamLdapPolicy) HasDnsParameters() bool {
	if o != nil && o.DnsParameters.IsSet() {
		return true
	}

	return false
}

// SetDnsParameters gets a reference to the given NullableIamLdapDnsParameters and assigns it to the DnsParameters field.
func (o *IamLdapPolicy) SetDnsParameters(v IamLdapDnsParameters) {
	o.DnsParameters.Set(&v)
}

// SetDnsParametersNil sets the value for DnsParameters to be an explicit nil
func (o *IamLdapPolicy) SetDnsParametersNil() {
	o.DnsParameters.Set(nil)
}

// UnsetDnsParameters ensures that no value is present for DnsParameters, not even an explicit nil
func (o *IamLdapPolicy) UnsetDnsParameters() {
	o.DnsParameters.Unset()
}

// GetEnableDns returns the EnableDns field value if set, zero value otherwise.
func (o *IamLdapPolicy) GetEnableDns() bool {
	if o == nil || IsNil(o.EnableDns) {
		var ret bool
		return ret
	}
	return *o.EnableDns
}

// GetEnableDnsOk returns a tuple with the EnableDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicy) GetEnableDnsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDns) {
		return nil, false
	}
	return o.EnableDns, true
}

// HasEnableDns returns a boolean if a field has been set.
func (o *IamLdapPolicy) HasEnableDns() bool {
	if o != nil && !IsNil(o.EnableDns) {
		return true
	}

	return false
}

// SetEnableDns gets a reference to the given bool and assigns it to the EnableDns field.
func (o *IamLdapPolicy) SetEnableDns(v bool) {
	o.EnableDns = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IamLdapPolicy) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IamLdapPolicy) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IamLdapPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUserSearchPrecedence returns the UserSearchPrecedence field value if set, zero value otherwise.
func (o *IamLdapPolicy) GetUserSearchPrecedence() string {
	if o == nil || IsNil(o.UserSearchPrecedence) {
		var ret string
		return ret
	}
	return *o.UserSearchPrecedence
}

// GetUserSearchPrecedenceOk returns a tuple with the UserSearchPrecedence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapPolicy) GetUserSearchPrecedenceOk() (*string, bool) {
	if o == nil || IsNil(o.UserSearchPrecedence) {
		return nil, false
	}
	return o.UserSearchPrecedence, true
}

// HasUserSearchPrecedence returns a boolean if a field has been set.
func (o *IamLdapPolicy) HasUserSearchPrecedence() bool {
	if o != nil && !IsNil(o.UserSearchPrecedence) {
		return true
	}

	return false
}

// SetUserSearchPrecedence gets a reference to the given string and assigns it to the UserSearchPrecedence field.
func (o *IamLdapPolicy) SetUserSearchPrecedence(v string) {
	o.UserSearchPrecedence = &v
}

// GetApplianceAccount returns the ApplianceAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamLdapPolicy) GetApplianceAccount() IamAccountRelationship {
	if o == nil || IsNil(o.ApplianceAccount.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.ApplianceAccount.Get()
}

// GetApplianceAccountOk returns a tuple with the ApplianceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamLdapPolicy) GetApplianceAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplianceAccount.Get(), o.ApplianceAccount.IsSet()
}

// HasApplianceAccount returns a boolean if a field has been set.
func (o *IamLdapPolicy) HasApplianceAccount() bool {
	if o != nil && o.ApplianceAccount.IsSet() {
		return true
	}

	return false
}

// SetApplianceAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the ApplianceAccount field.
func (o *IamLdapPolicy) SetApplianceAccount(v IamAccountRelationship) {
	o.ApplianceAccount.Set(&v)
}

// SetApplianceAccountNil sets the value for ApplianceAccount to be an explicit nil
func (o *IamLdapPolicy) SetApplianceAccountNil() {
	o.ApplianceAccount.Set(nil)
}

// UnsetApplianceAccount ensures that no value is present for ApplianceAccount, not even an explicit nil
func (o *IamLdapPolicy) UnsetApplianceAccount() {
	o.ApplianceAccount.Unset()
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamLdapPolicy) GetGroups() []IamLdapGroupRelationship {
	if o == nil {
		var ret []IamLdapGroupRelationship
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamLdapPolicy) GetGroupsOk() ([]IamLdapGroupRelationship, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *IamLdapPolicy) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []IamLdapGroupRelationship and assigns it to the Groups field.
func (o *IamLdapPolicy) SetGroups(v []IamLdapGroupRelationship) {
	o.Groups = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamLdapPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamLdapPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *IamLdapPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *IamLdapPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *IamLdapPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *IamLdapPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *IamLdapPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
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
func (o *IamLdapPolicy) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *IamLdapPolicy) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
// Deprecated
func (o *IamLdapPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

// GetProviders returns the Providers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamLdapPolicy) GetProviders() []IamLdapProviderRelationship {
	if o == nil {
		var ret []IamLdapProviderRelationship
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamLdapPolicy) GetProvidersOk() ([]IamLdapProviderRelationship, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *IamLdapPolicy) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []IamLdapProviderRelationship and assigns it to the Providers field.
func (o *IamLdapPolicy) SetProviders(v []IamLdapProviderRelationship) {
	o.Providers = v
}

func (o IamLdapPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamLdapPolicy) ToMap() (map[string]interface{}, error) {
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
	if o.BaseProperties.IsSet() {
		toSerialize["BaseProperties"] = o.BaseProperties.Get()
	}
	if o.DnsParameters.IsSet() {
		toSerialize["DnsParameters"] = o.DnsParameters.Get()
	}
	if !IsNil(o.EnableDns) {
		toSerialize["EnableDns"] = o.EnableDns
	}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.UserSearchPrecedence) {
		toSerialize["UserSearchPrecedence"] = o.UserSearchPrecedence
	}
	if o.ApplianceAccount.IsSet() {
		toSerialize["ApplianceAccount"] = o.ApplianceAccount.Get()
	}
	if o.Groups != nil {
		toSerialize["Groups"] = o.Groups
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	if o.Providers != nil {
		toSerialize["Providers"] = o.Providers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamLdapPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type IamLdapPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType     string                        `json:"ObjectType"`
		BaseProperties NullableIamLdapBaseProperties `json:"BaseProperties,omitempty"`
		DnsParameters  NullableIamLdapDnsParameters  `json:"DnsParameters,omitempty"`
		// Enables DNS to access LDAP servers.
		EnableDns *bool `json:"EnableDns,omitempty"`
		// LDAP server performs authentication.
		Enabled *bool `json:"Enabled,omitempty"`
		// Search precedence between local user database and LDAP user database. * `LocalUserDb` - Precedence is given to local user database while searching. * `LDAPUserDb` - Precedence is given to LADP user database while searching.
		UserSearchPrecedence *string                        `json:"UserSearchPrecedence,omitempty"`
		ApplianceAccount     NullableIamAccountRelationship `json:"ApplianceAccount,omitempty"`
		// An array of relationships to iamLdapGroup resources.
		Groups       []IamLdapGroupRelationship                   `json:"Groups,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		// Deprecated
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
		// An array of relationships to iamLdapProvider resources.
		Providers []IamLdapProviderRelationship `json:"Providers,omitempty"`
	}

	varIamLdapPolicyWithoutEmbeddedStruct := IamLdapPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamLdapPolicyWithoutEmbeddedStruct)
	if err == nil {
		varIamLdapPolicy := _IamLdapPolicy{}
		varIamLdapPolicy.ClassId = varIamLdapPolicyWithoutEmbeddedStruct.ClassId
		varIamLdapPolicy.ObjectType = varIamLdapPolicyWithoutEmbeddedStruct.ObjectType
		varIamLdapPolicy.BaseProperties = varIamLdapPolicyWithoutEmbeddedStruct.BaseProperties
		varIamLdapPolicy.DnsParameters = varIamLdapPolicyWithoutEmbeddedStruct.DnsParameters
		varIamLdapPolicy.EnableDns = varIamLdapPolicyWithoutEmbeddedStruct.EnableDns
		varIamLdapPolicy.Enabled = varIamLdapPolicyWithoutEmbeddedStruct.Enabled
		varIamLdapPolicy.UserSearchPrecedence = varIamLdapPolicyWithoutEmbeddedStruct.UserSearchPrecedence
		varIamLdapPolicy.ApplianceAccount = varIamLdapPolicyWithoutEmbeddedStruct.ApplianceAccount
		varIamLdapPolicy.Groups = varIamLdapPolicyWithoutEmbeddedStruct.Groups
		varIamLdapPolicy.Organization = varIamLdapPolicyWithoutEmbeddedStruct.Organization
		varIamLdapPolicy.Profiles = varIamLdapPolicyWithoutEmbeddedStruct.Profiles
		varIamLdapPolicy.Providers = varIamLdapPolicyWithoutEmbeddedStruct.Providers
		*o = IamLdapPolicy(varIamLdapPolicy)
	} else {
		return err
	}

	varIamLdapPolicy := _IamLdapPolicy{}

	err = json.Unmarshal(data, &varIamLdapPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varIamLdapPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BaseProperties")
		delete(additionalProperties, "DnsParameters")
		delete(additionalProperties, "EnableDns")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "UserSearchPrecedence")
		delete(additionalProperties, "ApplianceAccount")
		delete(additionalProperties, "Groups")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		delete(additionalProperties, "Providers")

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

type NullableIamLdapPolicy struct {
	value *IamLdapPolicy
	isSet bool
}

func (v NullableIamLdapPolicy) Get() *IamLdapPolicy {
	return v.value
}

func (v *NullableIamLdapPolicy) Set(val *IamLdapPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableIamLdapPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableIamLdapPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamLdapPolicy(val *IamLdapPolicy) *NullableIamLdapPolicy {
	return &NullableIamLdapPolicy{value: val, isSet: true}
}

func (v NullableIamLdapPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamLdapPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
