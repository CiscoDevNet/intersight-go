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

// checks if the CommTagDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommTagDefinition{}

// CommTagDefinition An arbitrary string. Tag can be assigned to any managed object. The key will be copied to the managed object. The tag key is a string representation of the tag. The tag key can be hierarchical or key-value pair.
type CommTagDefinition struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If this flag is enabled, the tag will be propagated to related managed objects. This is currently set to true by default for hierarchical tags. Propagation is managed by the system and cannot be configured by users.
	EnablePropagation *bool `json:"EnablePropagation,omitempty"`
	// The string representation of the tag key. If the tag is of hierarchical type, then \"/\" will be interpreted as hierarchy delimiters. It can contain alphabets, numbers, \"_\", \"-\". Key cannot start with \"_\", \"-\" or \"/\". The tag key must be unique within the account. The tag key is case sensitive and must not be empty.
	Key *string `json:"Key,omitempty" validate:"regexp=^[A-Za-z0-9]([A-Za-z0-9_.-]{0,48}[A-Za-z0-9])?(\\/[A-Za-z0-9]([A-Za-z0-9_.-]{0,48}[A-Za-z0-9])?)*$"`
	// An enum type that defines the type of tag. Only hierarchical tags are supported for now, and the type is set to hierarchical by default. * `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional. * `PathTag` - Key contain path information. Value is not present for these tags. The hierarchy is created by using the '/' character as a delimiter.For example, if the tag is \"A/B/C\", then \"A\" is the parent tag, \"B\" is the child tag of \"A\" and \"C\" is the child tag of \"B\".
	Type                 *string                               `json:"Type,omitempty"`
	Account              NullableIamAccountRelationship        `json:"Account,omitempty"`
	ParentTag            NullableCommTagDefinitionRelationship `json:"ParentTag,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommTagDefinition CommTagDefinition

// NewCommTagDefinition instantiates a new CommTagDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommTagDefinition(classId string, objectType string) *CommTagDefinition {
	this := CommTagDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "KeyValue"
	this.Type = &type_
	return &this
}

// NewCommTagDefinitionWithDefaults instantiates a new CommTagDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommTagDefinitionWithDefaults() *CommTagDefinition {
	this := CommTagDefinition{}
	var classId string = "comm.TagDefinition"
	this.ClassId = classId
	var objectType string = "comm.TagDefinition"
	this.ObjectType = objectType
	var type_ string = "KeyValue"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *CommTagDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CommTagDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CommTagDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "comm.TagDefinition" of the ClassId field.
func (o *CommTagDefinition) GetDefaultClassId() interface{} {
	return "comm.TagDefinition"
}

// GetObjectType returns the ObjectType field value
func (o *CommTagDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CommTagDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CommTagDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "comm.TagDefinition" of the ObjectType field.
func (o *CommTagDefinition) GetDefaultObjectType() interface{} {
	return "comm.TagDefinition"
}

// GetEnablePropagation returns the EnablePropagation field value if set, zero value otherwise.
func (o *CommTagDefinition) GetEnablePropagation() bool {
	if o == nil || IsNil(o.EnablePropagation) {
		var ret bool
		return ret
	}
	return *o.EnablePropagation
}

// GetEnablePropagationOk returns a tuple with the EnablePropagation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommTagDefinition) GetEnablePropagationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePropagation) {
		return nil, false
	}
	return o.EnablePropagation, true
}

// HasEnablePropagation returns a boolean if a field has been set.
func (o *CommTagDefinition) HasEnablePropagation() bool {
	if o != nil && !IsNil(o.EnablePropagation) {
		return true
	}

	return false
}

// SetEnablePropagation gets a reference to the given bool and assigns it to the EnablePropagation field.
func (o *CommTagDefinition) SetEnablePropagation(v bool) {
	o.EnablePropagation = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CommTagDefinition) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommTagDefinition) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CommTagDefinition) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CommTagDefinition) SetKey(v string) {
	o.Key = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CommTagDefinition) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommTagDefinition) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CommTagDefinition) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CommTagDefinition) SetType(v string) {
	o.Type = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommTagDefinition) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommTagDefinition) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *CommTagDefinition) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *CommTagDefinition) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *CommTagDefinition) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *CommTagDefinition) UnsetAccount() {
	o.Account.Unset()
}

// GetParentTag returns the ParentTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommTagDefinition) GetParentTag() CommTagDefinitionRelationship {
	if o == nil || IsNil(o.ParentTag.Get()) {
		var ret CommTagDefinitionRelationship
		return ret
	}
	return *o.ParentTag.Get()
}

// GetParentTagOk returns a tuple with the ParentTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommTagDefinition) GetParentTagOk() (*CommTagDefinitionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentTag.Get(), o.ParentTag.IsSet()
}

// HasParentTag returns a boolean if a field has been set.
func (o *CommTagDefinition) HasParentTag() bool {
	if o != nil && o.ParentTag.IsSet() {
		return true
	}

	return false
}

// SetParentTag gets a reference to the given NullableCommTagDefinitionRelationship and assigns it to the ParentTag field.
func (o *CommTagDefinition) SetParentTag(v CommTagDefinitionRelationship) {
	o.ParentTag.Set(&v)
}

// SetParentTagNil sets the value for ParentTag to be an explicit nil
func (o *CommTagDefinition) SetParentTagNil() {
	o.ParentTag.Set(nil)
}

// UnsetParentTag ensures that no value is present for ParentTag, not even an explicit nil
func (o *CommTagDefinition) UnsetParentTag() {
	o.ParentTag.Unset()
}

func (o CommTagDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommTagDefinition) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EnablePropagation) {
		toSerialize["EnablePropagation"] = o.EnablePropagation
	}
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}
	if o.ParentTag.IsSet() {
		toSerialize["ParentTag"] = o.ParentTag.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommTagDefinition) UnmarshalJSON(data []byte) (err error) {
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
	type CommTagDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// If this flag is enabled, the tag will be propagated to related managed objects. This is currently set to true by default for hierarchical tags. Propagation is managed by the system and cannot be configured by users.
		EnablePropagation *bool `json:"EnablePropagation,omitempty"`
		// The string representation of the tag key. If the tag is of hierarchical type, then \"/\" will be interpreted as hierarchy delimiters. It can contain alphabets, numbers, \"_\", \"-\". Key cannot start with \"_\", \"-\" or \"/\". The tag key must be unique within the account. The tag key is case sensitive and must not be empty.
		Key *string `json:"Key,omitempty" validate:"regexp=^[A-Za-z0-9]([A-Za-z0-9_.-]{0,48}[A-Za-z0-9])?(\\/[A-Za-z0-9]([A-Za-z0-9_.-]{0,48}[A-Za-z0-9])?)*$"`
		// An enum type that defines the type of tag. Only hierarchical tags are supported for now, and the type is set to hierarchical by default. * `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional. * `PathTag` - Key contain path information. Value is not present for these tags. The hierarchy is created by using the '/' character as a delimiter.For example, if the tag is \"A/B/C\", then \"A\" is the parent tag, \"B\" is the child tag of \"A\" and \"C\" is the child tag of \"B\".
		Type      *string                               `json:"Type,omitempty"`
		Account   NullableIamAccountRelationship        `json:"Account,omitempty"`
		ParentTag NullableCommTagDefinitionRelationship `json:"ParentTag,omitempty"`
	}

	varCommTagDefinitionWithoutEmbeddedStruct := CommTagDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCommTagDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varCommTagDefinition := _CommTagDefinition{}
		varCommTagDefinition.ClassId = varCommTagDefinitionWithoutEmbeddedStruct.ClassId
		varCommTagDefinition.ObjectType = varCommTagDefinitionWithoutEmbeddedStruct.ObjectType
		varCommTagDefinition.EnablePropagation = varCommTagDefinitionWithoutEmbeddedStruct.EnablePropagation
		varCommTagDefinition.Key = varCommTagDefinitionWithoutEmbeddedStruct.Key
		varCommTagDefinition.Type = varCommTagDefinitionWithoutEmbeddedStruct.Type
		varCommTagDefinition.Account = varCommTagDefinitionWithoutEmbeddedStruct.Account
		varCommTagDefinition.ParentTag = varCommTagDefinitionWithoutEmbeddedStruct.ParentTag
		*o = CommTagDefinition(varCommTagDefinition)
	} else {
		return err
	}

	varCommTagDefinition := _CommTagDefinition{}

	err = json.Unmarshal(data, &varCommTagDefinition)
	if err == nil {
		o.MoBaseMo = varCommTagDefinition.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnablePropagation")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "ParentTag")

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

type NullableCommTagDefinition struct {
	value *CommTagDefinition
	isSet bool
}

func (v NullableCommTagDefinition) Get() *CommTagDefinition {
	return v.value
}

func (v *NullableCommTagDefinition) Set(val *CommTagDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableCommTagDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableCommTagDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommTagDefinition(val *CommTagDefinition) *NullableCommTagDefinition {
	return &NullableCommTagDefinition{value: val, isSet: true}
}

func (v NullableCommTagDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommTagDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
