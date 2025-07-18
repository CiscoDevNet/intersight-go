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

// checks if the ContentTextParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentTextParameter{}

// ContentTextParameter Concrete implementation of BaseParameter for Text content.
type ContentTextParameter struct {
	ContentBaseParameter
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Data to be extracted from text content can be simple type or complex type or collection of simple/complex types. Complex types are group of simple or complex type. Delimiter is required to stop parsing list and complex data types. isDelimiter specifies whether given TextParameter is a delimiter or regular rule to capture the text data.
	IsDelimiter *bool `json:"IsDelimiter,omitempty"`
	// Set to true of the next value to capture resides on the same text line of current match. By default textFSM engine gets the next text line on finding the first match.
	IsNextCaptureOnSameLine *bool `json:"IsNextCaptureOnSameLine,omitempty"`
	// Regular expression of the line containing the data to be extracted from text content.
	RegexLine            *string `json:"RegexLine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContentTextParameter ContentTextParameter

// NewContentTextParameter instantiates a new ContentTextParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentTextParameter(classId string, objectType string) *ContentTextParameter {
	this := ContentTextParameter{}
	this.ClassId = classId
	this.ObjectType = objectType
	var itemType string = "simple"
	this.ItemType = &itemType
	var type_ string = "simple"
	this.Type = &type_
	var isDelimiter bool = false
	this.IsDelimiter = &isDelimiter
	var isNextCaptureOnSameLine bool = false
	this.IsNextCaptureOnSameLine = &isNextCaptureOnSameLine
	return &this
}

// NewContentTextParameterWithDefaults instantiates a new ContentTextParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentTextParameterWithDefaults() *ContentTextParameter {
	this := ContentTextParameter{}
	var classId string = "content.TextParameter"
	this.ClassId = classId
	var objectType string = "content.TextParameter"
	this.ObjectType = objectType
	var isDelimiter bool = false
	this.IsDelimiter = &isDelimiter
	var isNextCaptureOnSameLine bool = false
	this.IsNextCaptureOnSameLine = &isNextCaptureOnSameLine
	return &this
}

// GetClassId returns the ClassId field value
func (o *ContentTextParameter) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ContentTextParameter) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ContentTextParameter) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "content.TextParameter" of the ClassId field.
func (o *ContentTextParameter) GetDefaultClassId() interface{} {
	return "content.TextParameter"
}

// GetObjectType returns the ObjectType field value
func (o *ContentTextParameter) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ContentTextParameter) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ContentTextParameter) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "content.TextParameter" of the ObjectType field.
func (o *ContentTextParameter) GetDefaultObjectType() interface{} {
	return "content.TextParameter"
}

// GetIsDelimiter returns the IsDelimiter field value if set, zero value otherwise.
func (o *ContentTextParameter) GetIsDelimiter() bool {
	if o == nil || IsNil(o.IsDelimiter) {
		var ret bool
		return ret
	}
	return *o.IsDelimiter
}

// GetIsDelimiterOk returns a tuple with the IsDelimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextParameter) GetIsDelimiterOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDelimiter) {
		return nil, false
	}
	return o.IsDelimiter, true
}

// HasIsDelimiter returns a boolean if a field has been set.
func (o *ContentTextParameter) HasIsDelimiter() bool {
	if o != nil && !IsNil(o.IsDelimiter) {
		return true
	}

	return false
}

// SetIsDelimiter gets a reference to the given bool and assigns it to the IsDelimiter field.
func (o *ContentTextParameter) SetIsDelimiter(v bool) {
	o.IsDelimiter = &v
}

// GetIsNextCaptureOnSameLine returns the IsNextCaptureOnSameLine field value if set, zero value otherwise.
func (o *ContentTextParameter) GetIsNextCaptureOnSameLine() bool {
	if o == nil || IsNil(o.IsNextCaptureOnSameLine) {
		var ret bool
		return ret
	}
	return *o.IsNextCaptureOnSameLine
}

// GetIsNextCaptureOnSameLineOk returns a tuple with the IsNextCaptureOnSameLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextParameter) GetIsNextCaptureOnSameLineOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNextCaptureOnSameLine) {
		return nil, false
	}
	return o.IsNextCaptureOnSameLine, true
}

// HasIsNextCaptureOnSameLine returns a boolean if a field has been set.
func (o *ContentTextParameter) HasIsNextCaptureOnSameLine() bool {
	if o != nil && !IsNil(o.IsNextCaptureOnSameLine) {
		return true
	}

	return false
}

// SetIsNextCaptureOnSameLine gets a reference to the given bool and assigns it to the IsNextCaptureOnSameLine field.
func (o *ContentTextParameter) SetIsNextCaptureOnSameLine(v bool) {
	o.IsNextCaptureOnSameLine = &v
}

// GetRegexLine returns the RegexLine field value if set, zero value otherwise.
func (o *ContentTextParameter) GetRegexLine() string {
	if o == nil || IsNil(o.RegexLine) {
		var ret string
		return ret
	}
	return *o.RegexLine
}

// GetRegexLineOk returns a tuple with the RegexLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextParameter) GetRegexLineOk() (*string, bool) {
	if o == nil || IsNil(o.RegexLine) {
		return nil, false
	}
	return o.RegexLine, true
}

// HasRegexLine returns a boolean if a field has been set.
func (o *ContentTextParameter) HasRegexLine() bool {
	if o != nil && !IsNil(o.RegexLine) {
		return true
	}

	return false
}

// SetRegexLine gets a reference to the given string and assigns it to the RegexLine field.
func (o *ContentTextParameter) SetRegexLine(v string) {
	o.RegexLine = &v
}

func (o ContentTextParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentTextParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedContentBaseParameter, errContentBaseParameter := json.Marshal(o.ContentBaseParameter)
	if errContentBaseParameter != nil {
		return map[string]interface{}{}, errContentBaseParameter
	}
	errContentBaseParameter = json.Unmarshal([]byte(serializedContentBaseParameter), &toSerialize)
	if errContentBaseParameter != nil {
		return map[string]interface{}{}, errContentBaseParameter
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.IsDelimiter) {
		toSerialize["IsDelimiter"] = o.IsDelimiter
	}
	if !IsNil(o.IsNextCaptureOnSameLine) {
		toSerialize["IsNextCaptureOnSameLine"] = o.IsNextCaptureOnSameLine
	}
	if !IsNil(o.RegexLine) {
		toSerialize["RegexLine"] = o.RegexLine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContentTextParameter) UnmarshalJSON(data []byte) (err error) {
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
	type ContentTextParameterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Data to be extracted from text content can be simple type or complex type or collection of simple/complex types. Complex types are group of simple or complex type. Delimiter is required to stop parsing list and complex data types. isDelimiter specifies whether given TextParameter is a delimiter or regular rule to capture the text data.
		IsDelimiter *bool `json:"IsDelimiter,omitempty"`
		// Set to true of the next value to capture resides on the same text line of current match. By default textFSM engine gets the next text line on finding the first match.
		IsNextCaptureOnSameLine *bool `json:"IsNextCaptureOnSameLine,omitempty"`
		// Regular expression of the line containing the data to be extracted from text content.
		RegexLine *string `json:"RegexLine,omitempty"`
	}

	varContentTextParameterWithoutEmbeddedStruct := ContentTextParameterWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varContentTextParameterWithoutEmbeddedStruct)
	if err == nil {
		varContentTextParameter := _ContentTextParameter{}
		varContentTextParameter.ClassId = varContentTextParameterWithoutEmbeddedStruct.ClassId
		varContentTextParameter.ObjectType = varContentTextParameterWithoutEmbeddedStruct.ObjectType
		varContentTextParameter.IsDelimiter = varContentTextParameterWithoutEmbeddedStruct.IsDelimiter
		varContentTextParameter.IsNextCaptureOnSameLine = varContentTextParameterWithoutEmbeddedStruct.IsNextCaptureOnSameLine
		varContentTextParameter.RegexLine = varContentTextParameterWithoutEmbeddedStruct.RegexLine
		*o = ContentTextParameter(varContentTextParameter)
	} else {
		return err
	}

	varContentTextParameter := _ContentTextParameter{}

	err = json.Unmarshal(data, &varContentTextParameter)
	if err == nil {
		o.ContentBaseParameter = varContentTextParameter.ContentBaseParameter
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsDelimiter")
		delete(additionalProperties, "IsNextCaptureOnSameLine")
		delete(additionalProperties, "RegexLine")

		// remove fields from embedded structs
		reflectContentBaseParameter := reflect.ValueOf(o.ContentBaseParameter)
		for i := 0; i < reflectContentBaseParameter.Type().NumField(); i++ {
			t := reflectContentBaseParameter.Type().Field(i)

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

type NullableContentTextParameter struct {
	value *ContentTextParameter
	isSet bool
}

func (v NullableContentTextParameter) Get() *ContentTextParameter {
	return v.value
}

func (v *NullableContentTextParameter) Set(val *ContentTextParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableContentTextParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableContentTextParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentTextParameter(val *ContentTextParameter) *NullableContentTextParameter {
	return &NullableContentTextParameter{value: val, isSet: true}
}

func (v NullableContentTextParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentTextParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
