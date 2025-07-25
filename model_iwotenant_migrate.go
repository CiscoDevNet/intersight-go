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

// checks if the IwotenantMigrate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IwotenantMigrate{}

// IwotenantMigrate Api used to transfer customer's IWO data.
type IwotenantMigrate struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Password that satisfies Cisco's Password Complexity security requirement. Minimum password length, which MUST be nonzero and up to the current maximum length Contain both upper-case and lower-case letters Contain at least one number (for example, 0-9) Contain at least one special character.
	Password *string `json:"Password,omitempty"`
	// S3 bucket URL for uploading the migration data.
	Url                  *string                        `json:"Url,omitempty"`
	Account              NullableIamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IwotenantMigrate IwotenantMigrate

// NewIwotenantMigrate instantiates a new IwotenantMigrate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIwotenantMigrate(classId string, objectType string) *IwotenantMigrate {
	this := IwotenantMigrate{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIwotenantMigrateWithDefaults instantiates a new IwotenantMigrate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIwotenantMigrateWithDefaults() *IwotenantMigrate {
	this := IwotenantMigrate{}
	var classId string = "iwotenant.Migrate"
	this.ClassId = classId
	var objectType string = "iwotenant.Migrate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IwotenantMigrate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IwotenantMigrate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IwotenantMigrate) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iwotenant.Migrate" of the ClassId field.
func (o *IwotenantMigrate) GetDefaultClassId() interface{} {
	return "iwotenant.Migrate"
}

// GetObjectType returns the ObjectType field value
func (o *IwotenantMigrate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IwotenantMigrate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IwotenantMigrate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iwotenant.Migrate" of the ObjectType field.
func (o *IwotenantMigrate) GetDefaultObjectType() interface{} {
	return "iwotenant.Migrate"
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *IwotenantMigrate) GetIsPasswordSet() bool {
	if o == nil || IsNil(o.IsPasswordSet) {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMigrate) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPasswordSet) {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *IwotenantMigrate) HasIsPasswordSet() bool {
	if o != nil && !IsNil(o.IsPasswordSet) {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *IwotenantMigrate) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IwotenantMigrate) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMigrate) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IwotenantMigrate) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IwotenantMigrate) SetPassword(v string) {
	o.Password = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *IwotenantMigrate) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMigrate) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *IwotenantMigrate) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *IwotenantMigrate) SetUrl(v string) {
	o.Url = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IwotenantMigrate) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IwotenantMigrate) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *IwotenantMigrate) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *IwotenantMigrate) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *IwotenantMigrate) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *IwotenantMigrate) UnsetAccount() {
	o.Account.Unset()
}

func (o IwotenantMigrate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IwotenantMigrate) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsPasswordSet) {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if !IsNil(o.Password) {
		toSerialize["Password"] = o.Password
	}
	if !IsNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IwotenantMigrate) UnmarshalJSON(data []byte) (err error) {
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
	type IwotenantMigrateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// Password that satisfies Cisco's Password Complexity security requirement. Minimum password length, which MUST be nonzero and up to the current maximum length Contain both upper-case and lower-case letters Contain at least one number (for example, 0-9) Contain at least one special character.
		Password *string `json:"Password,omitempty"`
		// S3 bucket URL for uploading the migration data.
		Url     *string                        `json:"Url,omitempty"`
		Account NullableIamAccountRelationship `json:"Account,omitempty"`
	}

	varIwotenantMigrateWithoutEmbeddedStruct := IwotenantMigrateWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIwotenantMigrateWithoutEmbeddedStruct)
	if err == nil {
		varIwotenantMigrate := _IwotenantMigrate{}
		varIwotenantMigrate.ClassId = varIwotenantMigrateWithoutEmbeddedStruct.ClassId
		varIwotenantMigrate.ObjectType = varIwotenantMigrateWithoutEmbeddedStruct.ObjectType
		varIwotenantMigrate.IsPasswordSet = varIwotenantMigrateWithoutEmbeddedStruct.IsPasswordSet
		varIwotenantMigrate.Password = varIwotenantMigrateWithoutEmbeddedStruct.Password
		varIwotenantMigrate.Url = varIwotenantMigrateWithoutEmbeddedStruct.Url
		varIwotenantMigrate.Account = varIwotenantMigrateWithoutEmbeddedStruct.Account
		*o = IwotenantMigrate(varIwotenantMigrate)
	} else {
		return err
	}

	varIwotenantMigrate := _IwotenantMigrate{}

	err = json.Unmarshal(data, &varIwotenantMigrate)
	if err == nil {
		o.MoBaseMo = varIwotenantMigrate.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Url")
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

type NullableIwotenantMigrate struct {
	value *IwotenantMigrate
	isSet bool
}

func (v NullableIwotenantMigrate) Get() *IwotenantMigrate {
	return v.value
}

func (v *NullableIwotenantMigrate) Set(val *IwotenantMigrate) {
	v.value = val
	v.isSet = true
}

func (v NullableIwotenantMigrate) IsSet() bool {
	return v.isSet
}

func (v *NullableIwotenantMigrate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIwotenantMigrate(val *IwotenantMigrate) *NullableIwotenantMigrate {
	return &NullableIwotenantMigrate{value: val, isSet: true}
}

func (v NullableIwotenantMigrate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIwotenantMigrate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
