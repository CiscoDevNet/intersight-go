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

// checks if the HyperflexWwxnPrefixRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexWwxnPrefixRange{}

// HyperflexWwxnPrefixRange A range of WWxN prefixes. The range is inclusive and comprised of a start and end WWxN addresses. A single address can be specified by setting it as the start and end of the range.
type HyperflexWwxnPrefixRange struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The end WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX.
	EndAddr *string `json:"EndAddr,omitempty" validate:"regexp=^$|^20:00:00:25:B5:[0-9a-fA-F]{2}$"`
	// The start WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX.
	StartAddr            *string `json:"StartAddr,omitempty" validate:"regexp=^$|^20:00:00:25:B5:[0-9a-fA-F]{2}$"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexWwxnPrefixRange HyperflexWwxnPrefixRange

// NewHyperflexWwxnPrefixRange instantiates a new HyperflexWwxnPrefixRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexWwxnPrefixRange(classId string, objectType string) *HyperflexWwxnPrefixRange {
	this := HyperflexWwxnPrefixRange{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexWwxnPrefixRangeWithDefaults instantiates a new HyperflexWwxnPrefixRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexWwxnPrefixRangeWithDefaults() *HyperflexWwxnPrefixRange {
	this := HyperflexWwxnPrefixRange{}
	var classId string = "hyperflex.WwxnPrefixRange"
	this.ClassId = classId
	var objectType string = "hyperflex.WwxnPrefixRange"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexWwxnPrefixRange) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexWwxnPrefixRange) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexWwxnPrefixRange) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.WwxnPrefixRange" of the ClassId field.
func (o *HyperflexWwxnPrefixRange) GetDefaultClassId() interface{} {
	return "hyperflex.WwxnPrefixRange"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexWwxnPrefixRange) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexWwxnPrefixRange) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexWwxnPrefixRange) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.WwxnPrefixRange" of the ObjectType field.
func (o *HyperflexWwxnPrefixRange) GetDefaultObjectType() interface{} {
	return "hyperflex.WwxnPrefixRange"
}

// GetEndAddr returns the EndAddr field value if set, zero value otherwise.
func (o *HyperflexWwxnPrefixRange) GetEndAddr() string {
	if o == nil || IsNil(o.EndAddr) {
		var ret string
		return ret
	}
	return *o.EndAddr
}

// GetEndAddrOk returns a tuple with the EndAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexWwxnPrefixRange) GetEndAddrOk() (*string, bool) {
	if o == nil || IsNil(o.EndAddr) {
		return nil, false
	}
	return o.EndAddr, true
}

// HasEndAddr returns a boolean if a field has been set.
func (o *HyperflexWwxnPrefixRange) HasEndAddr() bool {
	if o != nil && !IsNil(o.EndAddr) {
		return true
	}

	return false
}

// SetEndAddr gets a reference to the given string and assigns it to the EndAddr field.
func (o *HyperflexWwxnPrefixRange) SetEndAddr(v string) {
	o.EndAddr = &v
}

// GetStartAddr returns the StartAddr field value if set, zero value otherwise.
func (o *HyperflexWwxnPrefixRange) GetStartAddr() string {
	if o == nil || IsNil(o.StartAddr) {
		var ret string
		return ret
	}
	return *o.StartAddr
}

// GetStartAddrOk returns a tuple with the StartAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexWwxnPrefixRange) GetStartAddrOk() (*string, bool) {
	if o == nil || IsNil(o.StartAddr) {
		return nil, false
	}
	return o.StartAddr, true
}

// HasStartAddr returns a boolean if a field has been set.
func (o *HyperflexWwxnPrefixRange) HasStartAddr() bool {
	if o != nil && !IsNil(o.StartAddr) {
		return true
	}

	return false
}

// SetStartAddr gets a reference to the given string and assigns it to the StartAddr field.
func (o *HyperflexWwxnPrefixRange) SetStartAddr(v string) {
	o.StartAddr = &v
}

func (o HyperflexWwxnPrefixRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexWwxnPrefixRange) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EndAddr) {
		toSerialize["EndAddr"] = o.EndAddr
	}
	if !IsNil(o.StartAddr) {
		toSerialize["StartAddr"] = o.StartAddr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexWwxnPrefixRange) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexWwxnPrefixRangeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The end WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX.
		EndAddr *string `json:"EndAddr,omitempty" validate:"regexp=^$|^20:00:00:25:B5:[0-9a-fA-F]{2}$"`
		// The start WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX.
		StartAddr *string `json:"StartAddr,omitempty" validate:"regexp=^$|^20:00:00:25:B5:[0-9a-fA-F]{2}$"`
	}

	varHyperflexWwxnPrefixRangeWithoutEmbeddedStruct := HyperflexWwxnPrefixRangeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexWwxnPrefixRangeWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexWwxnPrefixRange := _HyperflexWwxnPrefixRange{}
		varHyperflexWwxnPrefixRange.ClassId = varHyperflexWwxnPrefixRangeWithoutEmbeddedStruct.ClassId
		varHyperflexWwxnPrefixRange.ObjectType = varHyperflexWwxnPrefixRangeWithoutEmbeddedStruct.ObjectType
		varHyperflexWwxnPrefixRange.EndAddr = varHyperflexWwxnPrefixRangeWithoutEmbeddedStruct.EndAddr
		varHyperflexWwxnPrefixRange.StartAddr = varHyperflexWwxnPrefixRangeWithoutEmbeddedStruct.StartAddr
		*o = HyperflexWwxnPrefixRange(varHyperflexWwxnPrefixRange)
	} else {
		return err
	}

	varHyperflexWwxnPrefixRange := _HyperflexWwxnPrefixRange{}

	err = json.Unmarshal(data, &varHyperflexWwxnPrefixRange)
	if err == nil {
		o.MoBaseComplexType = varHyperflexWwxnPrefixRange.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndAddr")
		delete(additionalProperties, "StartAddr")

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

type NullableHyperflexWwxnPrefixRange struct {
	value *HyperflexWwxnPrefixRange
	isSet bool
}

func (v NullableHyperflexWwxnPrefixRange) Get() *HyperflexWwxnPrefixRange {
	return v.value
}

func (v *NullableHyperflexWwxnPrefixRange) Set(val *HyperflexWwxnPrefixRange) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexWwxnPrefixRange) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexWwxnPrefixRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexWwxnPrefixRange(val *HyperflexWwxnPrefixRange) *NullableHyperflexWwxnPrefixRange {
	return &NullableHyperflexWwxnPrefixRange{value: val, isSet: true}
}

func (v NullableHyperflexWwxnPrefixRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexWwxnPrefixRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
