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

// checks if the HyperflexFilePath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexFilePath{}

// HyperflexFilePath File path information for this snapshot.
type HyperflexFilePath struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                         `json:"ObjectType"`
	DsInfo     NullableHyperflexDatastoreInfo `json:"DsInfo,omitempty"`
	// Relative file path within the datastore.
	RelativeFilePath     *string `json:"RelativeFilePath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexFilePath HyperflexFilePath

// NewHyperflexFilePath instantiates a new HyperflexFilePath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexFilePath(classId string, objectType string) *HyperflexFilePath {
	this := HyperflexFilePath{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexFilePathWithDefaults instantiates a new HyperflexFilePath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexFilePathWithDefaults() *HyperflexFilePath {
	this := HyperflexFilePath{}
	var classId string = "hyperflex.FilePath"
	this.ClassId = classId
	var objectType string = "hyperflex.FilePath"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexFilePath) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexFilePath) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexFilePath) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.FilePath" of the ClassId field.
func (o *HyperflexFilePath) GetDefaultClassId() interface{} {
	return "hyperflex.FilePath"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexFilePath) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexFilePath) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexFilePath) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.FilePath" of the ObjectType field.
func (o *HyperflexFilePath) GetDefaultObjectType() interface{} {
	return "hyperflex.FilePath"
}

// GetDsInfo returns the DsInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexFilePath) GetDsInfo() HyperflexDatastoreInfo {
	if o == nil || IsNil(o.DsInfo.Get()) {
		var ret HyperflexDatastoreInfo
		return ret
	}
	return *o.DsInfo.Get()
}

// GetDsInfoOk returns a tuple with the DsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexFilePath) GetDsInfoOk() (*HyperflexDatastoreInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.DsInfo.Get(), o.DsInfo.IsSet()
}

// HasDsInfo returns a boolean if a field has been set.
func (o *HyperflexFilePath) HasDsInfo() bool {
	if o != nil && o.DsInfo.IsSet() {
		return true
	}

	return false
}

// SetDsInfo gets a reference to the given NullableHyperflexDatastoreInfo and assigns it to the DsInfo field.
func (o *HyperflexFilePath) SetDsInfo(v HyperflexDatastoreInfo) {
	o.DsInfo.Set(&v)
}

// SetDsInfoNil sets the value for DsInfo to be an explicit nil
func (o *HyperflexFilePath) SetDsInfoNil() {
	o.DsInfo.Set(nil)
}

// UnsetDsInfo ensures that no value is present for DsInfo, not even an explicit nil
func (o *HyperflexFilePath) UnsetDsInfo() {
	o.DsInfo.Unset()
}

// GetRelativeFilePath returns the RelativeFilePath field value if set, zero value otherwise.
func (o *HyperflexFilePath) GetRelativeFilePath() string {
	if o == nil || IsNil(o.RelativeFilePath) {
		var ret string
		return ret
	}
	return *o.RelativeFilePath
}

// GetRelativeFilePathOk returns a tuple with the RelativeFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexFilePath) GetRelativeFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.RelativeFilePath) {
		return nil, false
	}
	return o.RelativeFilePath, true
}

// HasRelativeFilePath returns a boolean if a field has been set.
func (o *HyperflexFilePath) HasRelativeFilePath() bool {
	if o != nil && !IsNil(o.RelativeFilePath) {
		return true
	}

	return false
}

// SetRelativeFilePath gets a reference to the given string and assigns it to the RelativeFilePath field.
func (o *HyperflexFilePath) SetRelativeFilePath(v string) {
	o.RelativeFilePath = &v
}

func (o HyperflexFilePath) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexFilePath) ToMap() (map[string]interface{}, error) {
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
	if o.DsInfo.IsSet() {
		toSerialize["DsInfo"] = o.DsInfo.Get()
	}
	if !IsNil(o.RelativeFilePath) {
		toSerialize["RelativeFilePath"] = o.RelativeFilePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexFilePath) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexFilePathWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                         `json:"ObjectType"`
		DsInfo     NullableHyperflexDatastoreInfo `json:"DsInfo,omitempty"`
		// Relative file path within the datastore.
		RelativeFilePath *string `json:"RelativeFilePath,omitempty"`
	}

	varHyperflexFilePathWithoutEmbeddedStruct := HyperflexFilePathWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexFilePathWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexFilePath := _HyperflexFilePath{}
		varHyperflexFilePath.ClassId = varHyperflexFilePathWithoutEmbeddedStruct.ClassId
		varHyperflexFilePath.ObjectType = varHyperflexFilePathWithoutEmbeddedStruct.ObjectType
		varHyperflexFilePath.DsInfo = varHyperflexFilePathWithoutEmbeddedStruct.DsInfo
		varHyperflexFilePath.RelativeFilePath = varHyperflexFilePathWithoutEmbeddedStruct.RelativeFilePath
		*o = HyperflexFilePath(varHyperflexFilePath)
	} else {
		return err
	}

	varHyperflexFilePath := _HyperflexFilePath{}

	err = json.Unmarshal(data, &varHyperflexFilePath)
	if err == nil {
		o.MoBaseComplexType = varHyperflexFilePath.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DsInfo")
		delete(additionalProperties, "RelativeFilePath")

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

type NullableHyperflexFilePath struct {
	value *HyperflexFilePath
	isSet bool
}

func (v NullableHyperflexFilePath) Get() *HyperflexFilePath {
	return v.value
}

func (v *NullableHyperflexFilePath) Set(val *HyperflexFilePath) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexFilePath) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexFilePath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexFilePath(val *HyperflexFilePath) *NullableHyperflexFilePath {
	return &NullableHyperflexFilePath{value: val, isSet: true}
}

func (v NullableHyperflexFilePath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexFilePath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
