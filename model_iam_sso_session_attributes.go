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

// checks if the IamSsoSessionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamSsoSessionAttributes{}

// IamSsoSessionAttributes Session attributes required to maintain states of SP and IdP.
type IamSsoSessionAttributes struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Captures if the user is allowed to switch roles or accounts after login and is used by UI to show/hide Switch Role Or Account option. It is set to true when Intersight identifies that the user has multiple roles or accounts during login based on user attributes and Intersight account IAM configuration. For IdP initiated SSO logins, this field is always set to false. Any changes to user roles or accounts after a session is created will not change this field. If switch role/account option is selected, user is re-authenticated with IdP and user roles and accounts are identified and AllowSwitchRoleOrAccount property is updated accordingly.
	AllowSwitchRoleOrAccount *bool `json:"AllowSwitchRoleOrAccount,omitempty"`
	// SAML SessionNotOnOrAfter attribute sent by IdP in the assertion. IdP uses this to control for how long SP session maybe. SP does not issue SLO if the session is not valid.
	IdpSessionExpiration *string `json:"IdpSessionExpiration,omitempty"`
	// SAML SessionIndex attribute sent by IdP in the assertion. This has to be sent back to IdP in LogoutRequest.
	IdpSessionIndex *string `json:"IdpSessionIndex,omitempty"`
	// Sign-in is SP-Intitiated or IdP-Intitiated.
	IsIdpInitiatedSso *bool `json:"IsIdpInitiatedSso,omitempty"`
	// Select Account Page shows partial list of accounts when processing data from some regions fails.
	IsPartialAccountsListed *bool `json:"IsPartialAccountsListed,omitempty"`
	// SAML Subject NameID attribute sent by IdP in the assertion. This has to be sent back to IdP in LogoutRequest.
	SubjectName          *string `json:"SubjectName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamSsoSessionAttributes IamSsoSessionAttributes

// NewIamSsoSessionAttributes instantiates a new IamSsoSessionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamSsoSessionAttributes(classId string, objectType string) *IamSsoSessionAttributes {
	this := IamSsoSessionAttributes{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamSsoSessionAttributesWithDefaults instantiates a new IamSsoSessionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamSsoSessionAttributesWithDefaults() *IamSsoSessionAttributes {
	this := IamSsoSessionAttributes{}
	var classId string = "iam.SsoSessionAttributes"
	this.ClassId = classId
	var objectType string = "iam.SsoSessionAttributes"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamSsoSessionAttributes) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamSsoSessionAttributes) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamSsoSessionAttributes) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.SsoSessionAttributes" of the ClassId field.
func (o *IamSsoSessionAttributes) GetDefaultClassId() interface{} {
	return "iam.SsoSessionAttributes"
}

// GetObjectType returns the ObjectType field value
func (o *IamSsoSessionAttributes) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamSsoSessionAttributes) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamSsoSessionAttributes) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.SsoSessionAttributes" of the ObjectType field.
func (o *IamSsoSessionAttributes) GetDefaultObjectType() interface{} {
	return "iam.SsoSessionAttributes"
}

// GetAllowSwitchRoleOrAccount returns the AllowSwitchRoleOrAccount field value if set, zero value otherwise.
func (o *IamSsoSessionAttributes) GetAllowSwitchRoleOrAccount() bool {
	if o == nil || IsNil(o.AllowSwitchRoleOrAccount) {
		var ret bool
		return ret
	}
	return *o.AllowSwitchRoleOrAccount
}

// GetAllowSwitchRoleOrAccountOk returns a tuple with the AllowSwitchRoleOrAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSsoSessionAttributes) GetAllowSwitchRoleOrAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSwitchRoleOrAccount) {
		return nil, false
	}
	return o.AllowSwitchRoleOrAccount, true
}

// HasAllowSwitchRoleOrAccount returns a boolean if a field has been set.
func (o *IamSsoSessionAttributes) HasAllowSwitchRoleOrAccount() bool {
	if o != nil && !IsNil(o.AllowSwitchRoleOrAccount) {
		return true
	}

	return false
}

// SetAllowSwitchRoleOrAccount gets a reference to the given bool and assigns it to the AllowSwitchRoleOrAccount field.
func (o *IamSsoSessionAttributes) SetAllowSwitchRoleOrAccount(v bool) {
	o.AllowSwitchRoleOrAccount = &v
}

// GetIdpSessionExpiration returns the IdpSessionExpiration field value if set, zero value otherwise.
func (o *IamSsoSessionAttributes) GetIdpSessionExpiration() string {
	if o == nil || IsNil(o.IdpSessionExpiration) {
		var ret string
		return ret
	}
	return *o.IdpSessionExpiration
}

// GetIdpSessionExpirationOk returns a tuple with the IdpSessionExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSsoSessionAttributes) GetIdpSessionExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.IdpSessionExpiration) {
		return nil, false
	}
	return o.IdpSessionExpiration, true
}

// HasIdpSessionExpiration returns a boolean if a field has been set.
func (o *IamSsoSessionAttributes) HasIdpSessionExpiration() bool {
	if o != nil && !IsNil(o.IdpSessionExpiration) {
		return true
	}

	return false
}

// SetIdpSessionExpiration gets a reference to the given string and assigns it to the IdpSessionExpiration field.
func (o *IamSsoSessionAttributes) SetIdpSessionExpiration(v string) {
	o.IdpSessionExpiration = &v
}

// GetIdpSessionIndex returns the IdpSessionIndex field value if set, zero value otherwise.
func (o *IamSsoSessionAttributes) GetIdpSessionIndex() string {
	if o == nil || IsNil(o.IdpSessionIndex) {
		var ret string
		return ret
	}
	return *o.IdpSessionIndex
}

// GetIdpSessionIndexOk returns a tuple with the IdpSessionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSsoSessionAttributes) GetIdpSessionIndexOk() (*string, bool) {
	if o == nil || IsNil(o.IdpSessionIndex) {
		return nil, false
	}
	return o.IdpSessionIndex, true
}

// HasIdpSessionIndex returns a boolean if a field has been set.
func (o *IamSsoSessionAttributes) HasIdpSessionIndex() bool {
	if o != nil && !IsNil(o.IdpSessionIndex) {
		return true
	}

	return false
}

// SetIdpSessionIndex gets a reference to the given string and assigns it to the IdpSessionIndex field.
func (o *IamSsoSessionAttributes) SetIdpSessionIndex(v string) {
	o.IdpSessionIndex = &v
}

// GetIsIdpInitiatedSso returns the IsIdpInitiatedSso field value if set, zero value otherwise.
func (o *IamSsoSessionAttributes) GetIsIdpInitiatedSso() bool {
	if o == nil || IsNil(o.IsIdpInitiatedSso) {
		var ret bool
		return ret
	}
	return *o.IsIdpInitiatedSso
}

// GetIsIdpInitiatedSsoOk returns a tuple with the IsIdpInitiatedSso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSsoSessionAttributes) GetIsIdpInitiatedSsoOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIdpInitiatedSso) {
		return nil, false
	}
	return o.IsIdpInitiatedSso, true
}

// HasIsIdpInitiatedSso returns a boolean if a field has been set.
func (o *IamSsoSessionAttributes) HasIsIdpInitiatedSso() bool {
	if o != nil && !IsNil(o.IsIdpInitiatedSso) {
		return true
	}

	return false
}

// SetIsIdpInitiatedSso gets a reference to the given bool and assigns it to the IsIdpInitiatedSso field.
func (o *IamSsoSessionAttributes) SetIsIdpInitiatedSso(v bool) {
	o.IsIdpInitiatedSso = &v
}

// GetIsPartialAccountsListed returns the IsPartialAccountsListed field value if set, zero value otherwise.
func (o *IamSsoSessionAttributes) GetIsPartialAccountsListed() bool {
	if o == nil || IsNil(o.IsPartialAccountsListed) {
		var ret bool
		return ret
	}
	return *o.IsPartialAccountsListed
}

// GetIsPartialAccountsListedOk returns a tuple with the IsPartialAccountsListed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSsoSessionAttributes) GetIsPartialAccountsListedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPartialAccountsListed) {
		return nil, false
	}
	return o.IsPartialAccountsListed, true
}

// HasIsPartialAccountsListed returns a boolean if a field has been set.
func (o *IamSsoSessionAttributes) HasIsPartialAccountsListed() bool {
	if o != nil && !IsNil(o.IsPartialAccountsListed) {
		return true
	}

	return false
}

// SetIsPartialAccountsListed gets a reference to the given bool and assigns it to the IsPartialAccountsListed field.
func (o *IamSsoSessionAttributes) SetIsPartialAccountsListed(v bool) {
	o.IsPartialAccountsListed = &v
}

// GetSubjectName returns the SubjectName field value if set, zero value otherwise.
func (o *IamSsoSessionAttributes) GetSubjectName() string {
	if o == nil || IsNil(o.SubjectName) {
		var ret string
		return ret
	}
	return *o.SubjectName
}

// GetSubjectNameOk returns a tuple with the SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSsoSessionAttributes) GetSubjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectName) {
		return nil, false
	}
	return o.SubjectName, true
}

// HasSubjectName returns a boolean if a field has been set.
func (o *IamSsoSessionAttributes) HasSubjectName() bool {
	if o != nil && !IsNil(o.SubjectName) {
		return true
	}

	return false
}

// SetSubjectName gets a reference to the given string and assigns it to the SubjectName field.
func (o *IamSsoSessionAttributes) SetSubjectName(v string) {
	o.SubjectName = &v
}

func (o IamSsoSessionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamSsoSessionAttributes) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AllowSwitchRoleOrAccount) {
		toSerialize["AllowSwitchRoleOrAccount"] = o.AllowSwitchRoleOrAccount
	}
	if !IsNil(o.IdpSessionExpiration) {
		toSerialize["IdpSessionExpiration"] = o.IdpSessionExpiration
	}
	if !IsNil(o.IdpSessionIndex) {
		toSerialize["IdpSessionIndex"] = o.IdpSessionIndex
	}
	if !IsNil(o.IsIdpInitiatedSso) {
		toSerialize["IsIdpInitiatedSso"] = o.IsIdpInitiatedSso
	}
	if !IsNil(o.IsPartialAccountsListed) {
		toSerialize["IsPartialAccountsListed"] = o.IsPartialAccountsListed
	}
	if !IsNil(o.SubjectName) {
		toSerialize["SubjectName"] = o.SubjectName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamSsoSessionAttributes) UnmarshalJSON(data []byte) (err error) {
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
	type IamSsoSessionAttributesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Captures if the user is allowed to switch roles or accounts after login and is used by UI to show/hide Switch Role Or Account option. It is set to true when Intersight identifies that the user has multiple roles or accounts during login based on user attributes and Intersight account IAM configuration. For IdP initiated SSO logins, this field is always set to false. Any changes to user roles or accounts after a session is created will not change this field. If switch role/account option is selected, user is re-authenticated with IdP and user roles and accounts are identified and AllowSwitchRoleOrAccount property is updated accordingly.
		AllowSwitchRoleOrAccount *bool `json:"AllowSwitchRoleOrAccount,omitempty"`
		// SAML SessionNotOnOrAfter attribute sent by IdP in the assertion. IdP uses this to control for how long SP session maybe. SP does not issue SLO if the session is not valid.
		IdpSessionExpiration *string `json:"IdpSessionExpiration,omitempty"`
		// SAML SessionIndex attribute sent by IdP in the assertion. This has to be sent back to IdP in LogoutRequest.
		IdpSessionIndex *string `json:"IdpSessionIndex,omitempty"`
		// Sign-in is SP-Intitiated or IdP-Intitiated.
		IsIdpInitiatedSso *bool `json:"IsIdpInitiatedSso,omitempty"`
		// Select Account Page shows partial list of accounts when processing data from some regions fails.
		IsPartialAccountsListed *bool `json:"IsPartialAccountsListed,omitempty"`
		// SAML Subject NameID attribute sent by IdP in the assertion. This has to be sent back to IdP in LogoutRequest.
		SubjectName *string `json:"SubjectName,omitempty"`
	}

	varIamSsoSessionAttributesWithoutEmbeddedStruct := IamSsoSessionAttributesWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamSsoSessionAttributesWithoutEmbeddedStruct)
	if err == nil {
		varIamSsoSessionAttributes := _IamSsoSessionAttributes{}
		varIamSsoSessionAttributes.ClassId = varIamSsoSessionAttributesWithoutEmbeddedStruct.ClassId
		varIamSsoSessionAttributes.ObjectType = varIamSsoSessionAttributesWithoutEmbeddedStruct.ObjectType
		varIamSsoSessionAttributes.AllowSwitchRoleOrAccount = varIamSsoSessionAttributesWithoutEmbeddedStruct.AllowSwitchRoleOrAccount
		varIamSsoSessionAttributes.IdpSessionExpiration = varIamSsoSessionAttributesWithoutEmbeddedStruct.IdpSessionExpiration
		varIamSsoSessionAttributes.IdpSessionIndex = varIamSsoSessionAttributesWithoutEmbeddedStruct.IdpSessionIndex
		varIamSsoSessionAttributes.IsIdpInitiatedSso = varIamSsoSessionAttributesWithoutEmbeddedStruct.IsIdpInitiatedSso
		varIamSsoSessionAttributes.IsPartialAccountsListed = varIamSsoSessionAttributesWithoutEmbeddedStruct.IsPartialAccountsListed
		varIamSsoSessionAttributes.SubjectName = varIamSsoSessionAttributesWithoutEmbeddedStruct.SubjectName
		*o = IamSsoSessionAttributes(varIamSsoSessionAttributes)
	} else {
		return err
	}

	varIamSsoSessionAttributes := _IamSsoSessionAttributes{}

	err = json.Unmarshal(data, &varIamSsoSessionAttributes)
	if err == nil {
		o.MoBaseComplexType = varIamSsoSessionAttributes.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllowSwitchRoleOrAccount")
		delete(additionalProperties, "IdpSessionExpiration")
		delete(additionalProperties, "IdpSessionIndex")
		delete(additionalProperties, "IsIdpInitiatedSso")
		delete(additionalProperties, "IsPartialAccountsListed")
		delete(additionalProperties, "SubjectName")

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

type NullableIamSsoSessionAttributes struct {
	value *IamSsoSessionAttributes
	isSet bool
}

func (v NullableIamSsoSessionAttributes) Get() *IamSsoSessionAttributes {
	return v.value
}

func (v *NullableIamSsoSessionAttributes) Set(val *IamSsoSessionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSsoSessionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSsoSessionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSsoSessionAttributes(val *IamSsoSessionAttributes) *NullableIamSsoSessionAttributes {
	return &NullableIamSsoSessionAttributes{value: val, isSet: true}
}

func (v NullableIamSsoSessionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSsoSessionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
