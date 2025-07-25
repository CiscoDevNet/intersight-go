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

// checks if the CapabilityAdapterUpgradeSupportMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityAdapterUpgradeSupportMeta{}

// CapabilityAdapterUpgradeSupportMeta Internal meta-data to enable adapter upgrade related decision making.
type CapabilityAdapterUpgradeSupportMeta struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Information related to the list of adapters.
	Description *string `json:"Description,omitempty"`
	// Series name of the adapter which will be supported in firmware operation.
	SeriesId             *string  `json:"SeriesId,omitempty"`
	SupportedModels      []string `json:"SupportedModels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityAdapterUpgradeSupportMeta CapabilityAdapterUpgradeSupportMeta

// NewCapabilityAdapterUpgradeSupportMeta instantiates a new CapabilityAdapterUpgradeSupportMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityAdapterUpgradeSupportMeta(classId string, objectType string) *CapabilityAdapterUpgradeSupportMeta {
	this := CapabilityAdapterUpgradeSupportMeta{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityAdapterUpgradeSupportMetaWithDefaults instantiates a new CapabilityAdapterUpgradeSupportMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityAdapterUpgradeSupportMetaWithDefaults() *CapabilityAdapterUpgradeSupportMeta {
	this := CapabilityAdapterUpgradeSupportMeta{}
	var classId string = "capability.AdapterUpgradeSupportMeta"
	this.ClassId = classId
	var objectType string = "capability.AdapterUpgradeSupportMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityAdapterUpgradeSupportMeta) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUpgradeSupportMeta) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityAdapterUpgradeSupportMeta) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.AdapterUpgradeSupportMeta" of the ClassId field.
func (o *CapabilityAdapterUpgradeSupportMeta) GetDefaultClassId() interface{} {
	return "capability.AdapterUpgradeSupportMeta"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityAdapterUpgradeSupportMeta) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUpgradeSupportMeta) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityAdapterUpgradeSupportMeta) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.AdapterUpgradeSupportMeta" of the ObjectType field.
func (o *CapabilityAdapterUpgradeSupportMeta) GetDefaultObjectType() interface{} {
	return "capability.AdapterUpgradeSupportMeta"
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CapabilityAdapterUpgradeSupportMeta) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUpgradeSupportMeta) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CapabilityAdapterUpgradeSupportMeta) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CapabilityAdapterUpgradeSupportMeta) SetDescription(v string) {
	o.Description = &v
}

// GetSeriesId returns the SeriesId field value if set, zero value otherwise.
func (o *CapabilityAdapterUpgradeSupportMeta) GetSeriesId() string {
	if o == nil || IsNil(o.SeriesId) {
		var ret string
		return ret
	}
	return *o.SeriesId
}

// GetSeriesIdOk returns a tuple with the SeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUpgradeSupportMeta) GetSeriesIdOk() (*string, bool) {
	if o == nil || IsNil(o.SeriesId) {
		return nil, false
	}
	return o.SeriesId, true
}

// HasSeriesId returns a boolean if a field has been set.
func (o *CapabilityAdapterUpgradeSupportMeta) HasSeriesId() bool {
	if o != nil && !IsNil(o.SeriesId) {
		return true
	}

	return false
}

// SetSeriesId gets a reference to the given string and assigns it to the SeriesId field.
func (o *CapabilityAdapterUpgradeSupportMeta) SetSeriesId(v string) {
	o.SeriesId = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityAdapterUpgradeSupportMeta) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityAdapterUpgradeSupportMeta) GetSupportedModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedModels) {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *CapabilityAdapterUpgradeSupportMeta) HasSupportedModels() bool {
	if o != nil && !IsNil(o.SupportedModels) {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *CapabilityAdapterUpgradeSupportMeta) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

func (o CapabilityAdapterUpgradeSupportMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityAdapterUpgradeSupportMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.SeriesId) {
		toSerialize["SeriesId"] = o.SeriesId
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityAdapterUpgradeSupportMeta) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityAdapterUpgradeSupportMetaWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Information related to the list of adapters.
		Description *string `json:"Description,omitempty"`
		// Series name of the adapter which will be supported in firmware operation.
		SeriesId        *string  `json:"SeriesId,omitempty"`
		SupportedModels []string `json:"SupportedModels,omitempty"`
	}

	varCapabilityAdapterUpgradeSupportMetaWithoutEmbeddedStruct := CapabilityAdapterUpgradeSupportMetaWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityAdapterUpgradeSupportMetaWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityAdapterUpgradeSupportMeta := _CapabilityAdapterUpgradeSupportMeta{}
		varCapabilityAdapterUpgradeSupportMeta.ClassId = varCapabilityAdapterUpgradeSupportMetaWithoutEmbeddedStruct.ClassId
		varCapabilityAdapterUpgradeSupportMeta.ObjectType = varCapabilityAdapterUpgradeSupportMetaWithoutEmbeddedStruct.ObjectType
		varCapabilityAdapterUpgradeSupportMeta.Description = varCapabilityAdapterUpgradeSupportMetaWithoutEmbeddedStruct.Description
		varCapabilityAdapterUpgradeSupportMeta.SeriesId = varCapabilityAdapterUpgradeSupportMetaWithoutEmbeddedStruct.SeriesId
		varCapabilityAdapterUpgradeSupportMeta.SupportedModels = varCapabilityAdapterUpgradeSupportMetaWithoutEmbeddedStruct.SupportedModels
		*o = CapabilityAdapterUpgradeSupportMeta(varCapabilityAdapterUpgradeSupportMeta)
	} else {
		return err
	}

	varCapabilityAdapterUpgradeSupportMeta := _CapabilityAdapterUpgradeSupportMeta{}

	err = json.Unmarshal(data, &varCapabilityAdapterUpgradeSupportMeta)
	if err == nil {
		o.CapabilityCapability = varCapabilityAdapterUpgradeSupportMeta.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "SeriesId")
		delete(additionalProperties, "SupportedModels")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableCapabilityAdapterUpgradeSupportMeta struct {
	value *CapabilityAdapterUpgradeSupportMeta
	isSet bool
}

func (v NullableCapabilityAdapterUpgradeSupportMeta) Get() *CapabilityAdapterUpgradeSupportMeta {
	return v.value
}

func (v *NullableCapabilityAdapterUpgradeSupportMeta) Set(val *CapabilityAdapterUpgradeSupportMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityAdapterUpgradeSupportMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityAdapterUpgradeSupportMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityAdapterUpgradeSupportMeta(val *CapabilityAdapterUpgradeSupportMeta) *NullableCapabilityAdapterUpgradeSupportMeta {
	return &NullableCapabilityAdapterUpgradeSupportMeta{value: val, isSet: true}
}

func (v NullableCapabilityAdapterUpgradeSupportMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityAdapterUpgradeSupportMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
