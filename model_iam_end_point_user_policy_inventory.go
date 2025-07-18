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

// checks if the IamEndPointUserPolicyInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamEndPointUserPolicyInventory{}

// IamEndPointUserPolicyInventory Enables creation of local users on endpoints.
type IamEndPointUserPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType         string                                `json:"ObjectType"`
	PasswordProperties NullableIamEndPointPasswordProperties `json:"PasswordProperties,omitempty"`
	// An array of relationships to iamEndPointUserRoleInventory resources.
	EndPointUserRoles    []IamEndPointUserRoleInventoryRelationship `json:"EndPointUserRoles,omitempty"`
	TargetMo             NullableMoBaseMoRelationship               `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamEndPointUserPolicyInventory IamEndPointUserPolicyInventory

// NewIamEndPointUserPolicyInventory instantiates a new IamEndPointUserPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointUserPolicyInventory(classId string, objectType string) *IamEndPointUserPolicyInventory {
	this := IamEndPointUserPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamEndPointUserPolicyInventoryWithDefaults instantiates a new IamEndPointUserPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointUserPolicyInventoryWithDefaults() *IamEndPointUserPolicyInventory {
	this := IamEndPointUserPolicyInventory{}
	var classId string = "iam.EndPointUserPolicyInventory"
	this.ClassId = classId
	var objectType string = "iam.EndPointUserPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamEndPointUserPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamEndPointUserPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.EndPointUserPolicyInventory" of the ClassId field.
func (o *IamEndPointUserPolicyInventory) GetDefaultClassId() interface{} {
	return "iam.EndPointUserPolicyInventory"
}

// GetObjectType returns the ObjectType field value
func (o *IamEndPointUserPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamEndPointUserPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.EndPointUserPolicyInventory" of the ObjectType field.
func (o *IamEndPointUserPolicyInventory) GetDefaultObjectType() interface{} {
	return "iam.EndPointUserPolicyInventory"
}

// GetPasswordProperties returns the PasswordProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserPolicyInventory) GetPasswordProperties() IamEndPointPasswordProperties {
	if o == nil || IsNil(o.PasswordProperties.Get()) {
		var ret IamEndPointPasswordProperties
		return ret
	}
	return *o.PasswordProperties.Get()
}

// GetPasswordPropertiesOk returns a tuple with the PasswordProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserPolicyInventory) GetPasswordPropertiesOk() (*IamEndPointPasswordProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordProperties.Get(), o.PasswordProperties.IsSet()
}

// HasPasswordProperties returns a boolean if a field has been set.
func (o *IamEndPointUserPolicyInventory) HasPasswordProperties() bool {
	if o != nil && o.PasswordProperties.IsSet() {
		return true
	}

	return false
}

// SetPasswordProperties gets a reference to the given NullableIamEndPointPasswordProperties and assigns it to the PasswordProperties field.
func (o *IamEndPointUserPolicyInventory) SetPasswordProperties(v IamEndPointPasswordProperties) {
	o.PasswordProperties.Set(&v)
}

// SetPasswordPropertiesNil sets the value for PasswordProperties to be an explicit nil
func (o *IamEndPointUserPolicyInventory) SetPasswordPropertiesNil() {
	o.PasswordProperties.Set(nil)
}

// UnsetPasswordProperties ensures that no value is present for PasswordProperties, not even an explicit nil
func (o *IamEndPointUserPolicyInventory) UnsetPasswordProperties() {
	o.PasswordProperties.Unset()
}

// GetEndPointUserRoles returns the EndPointUserRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserPolicyInventory) GetEndPointUserRoles() []IamEndPointUserRoleInventoryRelationship {
	if o == nil {
		var ret []IamEndPointUserRoleInventoryRelationship
		return ret
	}
	return o.EndPointUserRoles
}

// GetEndPointUserRolesOk returns a tuple with the EndPointUserRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserPolicyInventory) GetEndPointUserRolesOk() ([]IamEndPointUserRoleInventoryRelationship, bool) {
	if o == nil || IsNil(o.EndPointUserRoles) {
		return nil, false
	}
	return o.EndPointUserRoles, true
}

// HasEndPointUserRoles returns a boolean if a field has been set.
func (o *IamEndPointUserPolicyInventory) HasEndPointUserRoles() bool {
	if o != nil && !IsNil(o.EndPointUserRoles) {
		return true
	}

	return false
}

// SetEndPointUserRoles gets a reference to the given []IamEndPointUserRoleInventoryRelationship and assigns it to the EndPointUserRoles field.
func (o *IamEndPointUserPolicyInventory) SetEndPointUserRoles(v []IamEndPointUserRoleInventoryRelationship) {
	o.EndPointUserRoles = v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || IsNil(o.TargetMo.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo.Get()
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetMo.Get(), o.TargetMo.IsSet()
}

// HasTargetMo returns a boolean if a field has been set.
func (o *IamEndPointUserPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo.IsSet() {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given NullableMoBaseMoRelationship and assigns it to the TargetMo field.
func (o *IamEndPointUserPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo.Set(&v)
}

// SetTargetMoNil sets the value for TargetMo to be an explicit nil
func (o *IamEndPointUserPolicyInventory) SetTargetMoNil() {
	o.TargetMo.Set(nil)
}

// UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil
func (o *IamEndPointUserPolicyInventory) UnsetTargetMo() {
	o.TargetMo.Unset()
}

func (o IamEndPointUserPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamEndPointUserPolicyInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.PasswordProperties.IsSet() {
		toSerialize["PasswordProperties"] = o.PasswordProperties.Get()
	}
	if o.EndPointUserRoles != nil {
		toSerialize["EndPointUserRoles"] = o.EndPointUserRoles
	}
	if o.TargetMo.IsSet() {
		toSerialize["TargetMo"] = o.TargetMo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamEndPointUserPolicyInventory) UnmarshalJSON(data []byte) (err error) {
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
	type IamEndPointUserPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType         string                                `json:"ObjectType"`
		PasswordProperties NullableIamEndPointPasswordProperties `json:"PasswordProperties,omitempty"`
		// An array of relationships to iamEndPointUserRoleInventory resources.
		EndPointUserRoles []IamEndPointUserRoleInventoryRelationship `json:"EndPointUserRoles,omitempty"`
		TargetMo          NullableMoBaseMoRelationship               `json:"TargetMo,omitempty"`
	}

	varIamEndPointUserPolicyInventoryWithoutEmbeddedStruct := IamEndPointUserPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamEndPointUserPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varIamEndPointUserPolicyInventory := _IamEndPointUserPolicyInventory{}
		varIamEndPointUserPolicyInventory.ClassId = varIamEndPointUserPolicyInventoryWithoutEmbeddedStruct.ClassId
		varIamEndPointUserPolicyInventory.ObjectType = varIamEndPointUserPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varIamEndPointUserPolicyInventory.PasswordProperties = varIamEndPointUserPolicyInventoryWithoutEmbeddedStruct.PasswordProperties
		varIamEndPointUserPolicyInventory.EndPointUserRoles = varIamEndPointUserPolicyInventoryWithoutEmbeddedStruct.EndPointUserRoles
		varIamEndPointUserPolicyInventory.TargetMo = varIamEndPointUserPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = IamEndPointUserPolicyInventory(varIamEndPointUserPolicyInventory)
	} else {
		return err
	}

	varIamEndPointUserPolicyInventory := _IamEndPointUserPolicyInventory{}

	err = json.Unmarshal(data, &varIamEndPointUserPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varIamEndPointUserPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PasswordProperties")
		delete(additionalProperties, "EndPointUserRoles")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableIamEndPointUserPolicyInventory struct {
	value *IamEndPointUserPolicyInventory
	isSet bool
}

func (v NullableIamEndPointUserPolicyInventory) Get() *IamEndPointUserPolicyInventory {
	return v.value
}

func (v *NullableIamEndPointUserPolicyInventory) Set(val *IamEndPointUserPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointUserPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointUserPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointUserPolicyInventory(val *IamEndPointUserPolicyInventory) *NullableIamEndPointUserPolicyInventory {
	return &NullableIamEndPointUserPolicyInventory{value: val, isSet: true}
}

func (v NullableIamEndPointUserPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointUserPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
