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

// checks if the AssetTargetSignature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetTargetSignature{}

// AssetTargetSignature A signature generated by Intersight on creation of a target.
type AssetTargetSignature struct {
	AssetClaimSignature
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The moid of the associated target.
	TargetId             *string `json:"TargetId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetTargetSignature AssetTargetSignature

// NewAssetTargetSignature instantiates a new AssetTargetSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTargetSignature(classId string, objectType string) *AssetTargetSignature {
	this := AssetTargetSignature{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetTargetSignatureWithDefaults instantiates a new AssetTargetSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTargetSignatureWithDefaults() *AssetTargetSignature {
	this := AssetTargetSignature{}
	var classId string = "asset.TargetSignature"
	this.ClassId = classId
	var objectType string = "asset.TargetSignature"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetTargetSignature) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetTargetSignature) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetTargetSignature) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "asset.TargetSignature" of the ClassId field.
func (o *AssetTargetSignature) GetDefaultClassId() interface{} {
	return "asset.TargetSignature"
}

// GetObjectType returns the ObjectType field value
func (o *AssetTargetSignature) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetTargetSignature) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetTargetSignature) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "asset.TargetSignature" of the ObjectType field.
func (o *AssetTargetSignature) GetDefaultObjectType() interface{} {
	return "asset.TargetSignature"
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *AssetTargetSignature) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetSignature) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *AssetTargetSignature) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *AssetTargetSignature) SetTargetId(v string) {
	o.TargetId = &v
}

func (o AssetTargetSignature) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetTargetSignature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetClaimSignature, errAssetClaimSignature := json.Marshal(o.AssetClaimSignature)
	if errAssetClaimSignature != nil {
		return map[string]interface{}{}, errAssetClaimSignature
	}
	errAssetClaimSignature = json.Unmarshal([]byte(serializedAssetClaimSignature), &toSerialize)
	if errAssetClaimSignature != nil {
		return map[string]interface{}{}, errAssetClaimSignature
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.TargetId) {
		toSerialize["TargetId"] = o.TargetId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetTargetSignature) UnmarshalJSON(data []byte) (err error) {
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
	type AssetTargetSignatureWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The moid of the associated target.
		TargetId *string `json:"TargetId,omitempty"`
	}

	varAssetTargetSignatureWithoutEmbeddedStruct := AssetTargetSignatureWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetTargetSignatureWithoutEmbeddedStruct)
	if err == nil {
		varAssetTargetSignature := _AssetTargetSignature{}
		varAssetTargetSignature.ClassId = varAssetTargetSignatureWithoutEmbeddedStruct.ClassId
		varAssetTargetSignature.ObjectType = varAssetTargetSignatureWithoutEmbeddedStruct.ObjectType
		varAssetTargetSignature.TargetId = varAssetTargetSignatureWithoutEmbeddedStruct.TargetId
		*o = AssetTargetSignature(varAssetTargetSignature)
	} else {
		return err
	}

	varAssetTargetSignature := _AssetTargetSignature{}

	err = json.Unmarshal(data, &varAssetTargetSignature)
	if err == nil {
		o.AssetClaimSignature = varAssetTargetSignature.AssetClaimSignature
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetId")

		// remove fields from embedded structs
		reflectAssetClaimSignature := reflect.ValueOf(o.AssetClaimSignature)
		for i := 0; i < reflectAssetClaimSignature.Type().NumField(); i++ {
			t := reflectAssetClaimSignature.Type().Field(i)

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

type NullableAssetTargetSignature struct {
	value *AssetTargetSignature
	isSet bool
}

func (v NullableAssetTargetSignature) Get() *AssetTargetSignature {
	return v.value
}

func (v *NullableAssetTargetSignature) Set(val *AssetTargetSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTargetSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTargetSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTargetSignature(val *AssetTargetSignature) *NullableAssetTargetSignature {
	return &NullableAssetTargetSignature{value: val, isSet: true}
}

func (v NullableAssetTargetSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTargetSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
