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

// checks if the IssueDeviceTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueDeviceTag{}

// IssueDeviceTag An arbitrary key and value pair that can be used to tag REST resources and organize managed objects by assigning meta-data tags to any object.
type IssueDeviceTag struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The string representation of a tag key.
	Key *string `json:"Key,omitempty"`
	// The string representation of a tag value.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IssueDeviceTag IssueDeviceTag

// NewIssueDeviceTag instantiates a new IssueDeviceTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueDeviceTag(classId string, objectType string) *IssueDeviceTag {
	this := IssueDeviceTag{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIssueDeviceTagWithDefaults instantiates a new IssueDeviceTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueDeviceTagWithDefaults() *IssueDeviceTag {
	this := IssueDeviceTag{}
	var classId string = "issue.DeviceTag"
	this.ClassId = classId
	var objectType string = "issue.DeviceTag"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IssueDeviceTag) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IssueDeviceTag) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IssueDeviceTag) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "issue.DeviceTag" of the ClassId field.
func (o *IssueDeviceTag) GetDefaultClassId() interface{} {
	return "issue.DeviceTag"
}

// GetObjectType returns the ObjectType field value
func (o *IssueDeviceTag) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IssueDeviceTag) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IssueDeviceTag) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "issue.DeviceTag" of the ObjectType field.
func (o *IssueDeviceTag) GetDefaultObjectType() interface{} {
	return "issue.DeviceTag"
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *IssueDeviceTag) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDeviceTag) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *IssueDeviceTag) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *IssueDeviceTag) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IssueDeviceTag) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDeviceTag) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IssueDeviceTag) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IssueDeviceTag) SetValue(v string) {
	o.Value = &v
}

func (o IssueDeviceTag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueDeviceTag) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IssueDeviceTag) UnmarshalJSON(data []byte) (err error) {
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
	type IssueDeviceTagWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The string representation of a tag key.
		Key *string `json:"Key,omitempty"`
		// The string representation of a tag value.
		Value *string `json:"Value,omitempty"`
	}

	varIssueDeviceTagWithoutEmbeddedStruct := IssueDeviceTagWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIssueDeviceTagWithoutEmbeddedStruct)
	if err == nil {
		varIssueDeviceTag := _IssueDeviceTag{}
		varIssueDeviceTag.ClassId = varIssueDeviceTagWithoutEmbeddedStruct.ClassId
		varIssueDeviceTag.ObjectType = varIssueDeviceTagWithoutEmbeddedStruct.ObjectType
		varIssueDeviceTag.Key = varIssueDeviceTagWithoutEmbeddedStruct.Key
		varIssueDeviceTag.Value = varIssueDeviceTagWithoutEmbeddedStruct.Value
		*o = IssueDeviceTag(varIssueDeviceTag)
	} else {
		return err
	}

	varIssueDeviceTag := _IssueDeviceTag{}

	err = json.Unmarshal(data, &varIssueDeviceTag)
	if err == nil {
		o.MoBaseComplexType = varIssueDeviceTag.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Value")

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

type NullableIssueDeviceTag struct {
	value *IssueDeviceTag
	isSet bool
}

func (v NullableIssueDeviceTag) Get() *IssueDeviceTag {
	return v.value
}

func (v *NullableIssueDeviceTag) Set(val *IssueDeviceTag) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueDeviceTag) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueDeviceTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueDeviceTag(val *IssueDeviceTag) *NullableIssueDeviceTag {
	return &NullableIssueDeviceTag{value: val, isSet: true}
}

func (v NullableIssueDeviceTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueDeviceTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
