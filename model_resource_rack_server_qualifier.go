/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2024112619
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

// checks if the ResourceRackServerQualifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRackServerQualifier{}

// ResourceRackServerQualifier The rack servers are qualified for the pool based on properties of RackServer, such as RackId, RackPID, etc., which are captured as part of the RackServerQualifier.
type ResourceRackServerQualifier struct {
	ResourceServerQualifier
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                      `json:"ObjectType"`
	RackIdRange          []ResourceRackIdRangeFilter `json:"RackIdRange,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceRackServerQualifier ResourceRackServerQualifier

// NewResourceRackServerQualifier instantiates a new ResourceRackServerQualifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRackServerQualifier(classId string, objectType string) *ResourceRackServerQualifier {
	this := ResourceRackServerQualifier{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourceRackServerQualifierWithDefaults instantiates a new ResourceRackServerQualifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRackServerQualifierWithDefaults() *ResourceRackServerQualifier {
	this := ResourceRackServerQualifier{}
	var classId string = "resource.RackServerQualifier"
	this.ClassId = classId
	var objectType string = "resource.RackServerQualifier"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceRackServerQualifier) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceRackServerQualifier) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceRackServerQualifier) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "resource.RackServerQualifier" of the ClassId field.
func (o *ResourceRackServerQualifier) GetDefaultClassId() interface{} {
	return "resource.RackServerQualifier"
}

// GetObjectType returns the ObjectType field value
func (o *ResourceRackServerQualifier) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceRackServerQualifier) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceRackServerQualifier) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "resource.RackServerQualifier" of the ObjectType field.
func (o *ResourceRackServerQualifier) GetDefaultObjectType() interface{} {
	return "resource.RackServerQualifier"
}

// GetRackIdRange returns the RackIdRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceRackServerQualifier) GetRackIdRange() []ResourceRackIdRangeFilter {
	if o == nil {
		var ret []ResourceRackIdRangeFilter
		return ret
	}
	return o.RackIdRange
}

// GetRackIdRangeOk returns a tuple with the RackIdRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceRackServerQualifier) GetRackIdRangeOk() ([]ResourceRackIdRangeFilter, bool) {
	if o == nil || IsNil(o.RackIdRange) {
		return nil, false
	}
	return o.RackIdRange, true
}

// HasRackIdRange returns a boolean if a field has been set.
func (o *ResourceRackServerQualifier) HasRackIdRange() bool {
	if o != nil && !IsNil(o.RackIdRange) {
		return true
	}

	return false
}

// SetRackIdRange gets a reference to the given []ResourceRackIdRangeFilter and assigns it to the RackIdRange field.
func (o *ResourceRackServerQualifier) SetRackIdRange(v []ResourceRackIdRangeFilter) {
	o.RackIdRange = v
}

func (o ResourceRackServerQualifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRackServerQualifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedResourceServerQualifier, errResourceServerQualifier := json.Marshal(o.ResourceServerQualifier)
	if errResourceServerQualifier != nil {
		return map[string]interface{}{}, errResourceServerQualifier
	}
	errResourceServerQualifier = json.Unmarshal([]byte(serializedResourceServerQualifier), &toSerialize)
	if errResourceServerQualifier != nil {
		return map[string]interface{}{}, errResourceServerQualifier
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.RackIdRange != nil {
		toSerialize["RackIdRange"] = o.RackIdRange
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceRackServerQualifier) UnmarshalJSON(data []byte) (err error) {
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
	type ResourceRackServerQualifierWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string                      `json:"ObjectType"`
		RackIdRange []ResourceRackIdRangeFilter `json:"RackIdRange,omitempty"`
	}

	varResourceRackServerQualifierWithoutEmbeddedStruct := ResourceRackServerQualifierWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varResourceRackServerQualifierWithoutEmbeddedStruct)
	if err == nil {
		varResourceRackServerQualifier := _ResourceRackServerQualifier{}
		varResourceRackServerQualifier.ClassId = varResourceRackServerQualifierWithoutEmbeddedStruct.ClassId
		varResourceRackServerQualifier.ObjectType = varResourceRackServerQualifierWithoutEmbeddedStruct.ObjectType
		varResourceRackServerQualifier.RackIdRange = varResourceRackServerQualifierWithoutEmbeddedStruct.RackIdRange
		*o = ResourceRackServerQualifier(varResourceRackServerQualifier)
	} else {
		return err
	}

	varResourceRackServerQualifier := _ResourceRackServerQualifier{}

	err = json.Unmarshal(data, &varResourceRackServerQualifier)
	if err == nil {
		o.ResourceServerQualifier = varResourceRackServerQualifier.ResourceServerQualifier
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RackIdRange")

		// remove fields from embedded structs
		reflectResourceServerQualifier := reflect.ValueOf(o.ResourceServerQualifier)
		for i := 0; i < reflectResourceServerQualifier.Type().NumField(); i++ {
			t := reflectResourceServerQualifier.Type().Field(i)

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

type NullableResourceRackServerQualifier struct {
	value *ResourceRackServerQualifier
	isSet bool
}

func (v NullableResourceRackServerQualifier) Get() *ResourceRackServerQualifier {
	return v.value
}

func (v *NullableResourceRackServerQualifier) Set(val *ResourceRackServerQualifier) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRackServerQualifier) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRackServerQualifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRackServerQualifier(val *ResourceRackServerQualifier) *NullableResourceRackServerQualifier {
	return &NullableResourceRackServerQualifier{value: val, isSet: true}
}

func (v NullableResourceRackServerQualifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRackServerQualifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}