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

// checks if the CapabilityPortRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityPortRange{}

// CapabilityPortRange Type to represent ranges of ports supporting certain speeds or features on a switch.
type CapabilityPortRange struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Ending Port ID in this range of ports.
	EndPortId *int64 `json:"EndPortId,omitempty"`
	// Ending Slot ID in this range of ports.
	EndSlotId *int64 `json:"EndSlotId,omitempty"`
	// Starting Port ID in this range of ports.
	StartPortId *int64 `json:"StartPortId,omitempty"`
	// Starting Slot ID in this range of ports.
	StartSlotId          *int64 `json:"StartSlotId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityPortRange CapabilityPortRange

// NewCapabilityPortRange instantiates a new CapabilityPortRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityPortRange(classId string, objectType string) *CapabilityPortRange {
	this := CapabilityPortRange{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityPortRangeWithDefaults instantiates a new CapabilityPortRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityPortRangeWithDefaults() *CapabilityPortRange {
	this := CapabilityPortRange{}
	var classId string = "capability.PortRange"
	this.ClassId = classId
	var objectType string = "capability.PortRange"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityPortRange) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityPortRange) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityPortRange) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.PortRange" of the ClassId field.
func (o *CapabilityPortRange) GetDefaultClassId() interface{} {
	return "capability.PortRange"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityPortRange) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityPortRange) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityPortRange) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.PortRange" of the ObjectType field.
func (o *CapabilityPortRange) GetDefaultObjectType() interface{} {
	return "capability.PortRange"
}

// GetEndPortId returns the EndPortId field value if set, zero value otherwise.
func (o *CapabilityPortRange) GetEndPortId() int64 {
	if o == nil || IsNil(o.EndPortId) {
		var ret int64
		return ret
	}
	return *o.EndPortId
}

// GetEndPortIdOk returns a tuple with the EndPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityPortRange) GetEndPortIdOk() (*int64, bool) {
	if o == nil || IsNil(o.EndPortId) {
		return nil, false
	}
	return o.EndPortId, true
}

// HasEndPortId returns a boolean if a field has been set.
func (o *CapabilityPortRange) HasEndPortId() bool {
	if o != nil && !IsNil(o.EndPortId) {
		return true
	}

	return false
}

// SetEndPortId gets a reference to the given int64 and assigns it to the EndPortId field.
func (o *CapabilityPortRange) SetEndPortId(v int64) {
	o.EndPortId = &v
}

// GetEndSlotId returns the EndSlotId field value if set, zero value otherwise.
func (o *CapabilityPortRange) GetEndSlotId() int64 {
	if o == nil || IsNil(o.EndSlotId) {
		var ret int64
		return ret
	}
	return *o.EndSlotId
}

// GetEndSlotIdOk returns a tuple with the EndSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityPortRange) GetEndSlotIdOk() (*int64, bool) {
	if o == nil || IsNil(o.EndSlotId) {
		return nil, false
	}
	return o.EndSlotId, true
}

// HasEndSlotId returns a boolean if a field has been set.
func (o *CapabilityPortRange) HasEndSlotId() bool {
	if o != nil && !IsNil(o.EndSlotId) {
		return true
	}

	return false
}

// SetEndSlotId gets a reference to the given int64 and assigns it to the EndSlotId field.
func (o *CapabilityPortRange) SetEndSlotId(v int64) {
	o.EndSlotId = &v
}

// GetStartPortId returns the StartPortId field value if set, zero value otherwise.
func (o *CapabilityPortRange) GetStartPortId() int64 {
	if o == nil || IsNil(o.StartPortId) {
		var ret int64
		return ret
	}
	return *o.StartPortId
}

// GetStartPortIdOk returns a tuple with the StartPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityPortRange) GetStartPortIdOk() (*int64, bool) {
	if o == nil || IsNil(o.StartPortId) {
		return nil, false
	}
	return o.StartPortId, true
}

// HasStartPortId returns a boolean if a field has been set.
func (o *CapabilityPortRange) HasStartPortId() bool {
	if o != nil && !IsNil(o.StartPortId) {
		return true
	}

	return false
}

// SetStartPortId gets a reference to the given int64 and assigns it to the StartPortId field.
func (o *CapabilityPortRange) SetStartPortId(v int64) {
	o.StartPortId = &v
}

// GetStartSlotId returns the StartSlotId field value if set, zero value otherwise.
func (o *CapabilityPortRange) GetStartSlotId() int64 {
	if o == nil || IsNil(o.StartSlotId) {
		var ret int64
		return ret
	}
	return *o.StartSlotId
}

// GetStartSlotIdOk returns a tuple with the StartSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityPortRange) GetStartSlotIdOk() (*int64, bool) {
	if o == nil || IsNil(o.StartSlotId) {
		return nil, false
	}
	return o.StartSlotId, true
}

// HasStartSlotId returns a boolean if a field has been set.
func (o *CapabilityPortRange) HasStartSlotId() bool {
	if o != nil && !IsNil(o.StartSlotId) {
		return true
	}

	return false
}

// SetStartSlotId gets a reference to the given int64 and assigns it to the StartSlotId field.
func (o *CapabilityPortRange) SetStartSlotId(v int64) {
	o.StartSlotId = &v
}

func (o CapabilityPortRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityPortRange) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EndPortId) {
		toSerialize["EndPortId"] = o.EndPortId
	}
	if !IsNil(o.EndSlotId) {
		toSerialize["EndSlotId"] = o.EndSlotId
	}
	if !IsNil(o.StartPortId) {
		toSerialize["StartPortId"] = o.StartPortId
	}
	if !IsNil(o.StartSlotId) {
		toSerialize["StartSlotId"] = o.StartSlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityPortRange) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityPortRangeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Ending Port ID in this range of ports.
		EndPortId *int64 `json:"EndPortId,omitempty"`
		// Ending Slot ID in this range of ports.
		EndSlotId *int64 `json:"EndSlotId,omitempty"`
		// Starting Port ID in this range of ports.
		StartPortId *int64 `json:"StartPortId,omitempty"`
		// Starting Slot ID in this range of ports.
		StartSlotId *int64 `json:"StartSlotId,omitempty"`
	}

	varCapabilityPortRangeWithoutEmbeddedStruct := CapabilityPortRangeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityPortRangeWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityPortRange := _CapabilityPortRange{}
		varCapabilityPortRange.ClassId = varCapabilityPortRangeWithoutEmbeddedStruct.ClassId
		varCapabilityPortRange.ObjectType = varCapabilityPortRangeWithoutEmbeddedStruct.ObjectType
		varCapabilityPortRange.EndPortId = varCapabilityPortRangeWithoutEmbeddedStruct.EndPortId
		varCapabilityPortRange.EndSlotId = varCapabilityPortRangeWithoutEmbeddedStruct.EndSlotId
		varCapabilityPortRange.StartPortId = varCapabilityPortRangeWithoutEmbeddedStruct.StartPortId
		varCapabilityPortRange.StartSlotId = varCapabilityPortRangeWithoutEmbeddedStruct.StartSlotId
		*o = CapabilityPortRange(varCapabilityPortRange)
	} else {
		return err
	}

	varCapabilityPortRange := _CapabilityPortRange{}

	err = json.Unmarshal(data, &varCapabilityPortRange)
	if err == nil {
		o.MoBaseComplexType = varCapabilityPortRange.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndPortId")
		delete(additionalProperties, "EndSlotId")
		delete(additionalProperties, "StartPortId")
		delete(additionalProperties, "StartSlotId")

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

type NullableCapabilityPortRange struct {
	value *CapabilityPortRange
	isSet bool
}

func (v NullableCapabilityPortRange) Get() *CapabilityPortRange {
	return v.value
}

func (v *NullableCapabilityPortRange) Set(val *CapabilityPortRange) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityPortRange) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityPortRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityPortRange(val *CapabilityPortRange) *NullableCapabilityPortRange {
	return &NullableCapabilityPortRange{value: val, isSet: true}
}

func (v NullableCapabilityPortRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityPortRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
