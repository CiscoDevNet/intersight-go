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

// checks if the ResourcepoolQualificationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourcepoolQualificationType{}

// ResourcepoolQualificationType The QualificationType denotes the method by which a resource is added to a pool. It serves as an identifier to distinguish between static and dynamic resource additions. 'Static' indicates that the resource is added with predetermined specifications. Conversely, 'Dynamic' suggests that the resource is added based on specific conditions or qualifiers. In the dynamic case, this includes the qualification policies under which the resource is discovered.
type ResourcepoolQualificationType struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The resource is added statically or dynamically. If the value is true, it signifies a static addition, implying that the resource was added with fixed specifications. Conversely, if the value is false, it suggests a dynamic addition, indicating that the resource was incorporated based on certain conditions or qualifiers.
	IsStaticAddition      *bool     `json:"IsStaticAddition,omitempty"`
	QualificationPolicies []MoMoRef `json:"QualificationPolicies,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _ResourcepoolQualificationType ResourcepoolQualificationType

// NewResourcepoolQualificationType instantiates a new ResourcepoolQualificationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcepoolQualificationType(classId string, objectType string) *ResourcepoolQualificationType {
	this := ResourcepoolQualificationType{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourcepoolQualificationTypeWithDefaults instantiates a new ResourcepoolQualificationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcepoolQualificationTypeWithDefaults() *ResourcepoolQualificationType {
	this := ResourcepoolQualificationType{}
	var classId string = "resourcepool.QualificationType"
	this.ClassId = classId
	var objectType string = "resourcepool.QualificationType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourcepoolQualificationType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolQualificationType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourcepoolQualificationType) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "resourcepool.QualificationType" of the ClassId field.
func (o *ResourcepoolQualificationType) GetDefaultClassId() interface{} {
	return "resourcepool.QualificationType"
}

// GetObjectType returns the ObjectType field value
func (o *ResourcepoolQualificationType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolQualificationType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourcepoolQualificationType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "resourcepool.QualificationType" of the ObjectType field.
func (o *ResourcepoolQualificationType) GetDefaultObjectType() interface{} {
	return "resourcepool.QualificationType"
}

// GetIsStaticAddition returns the IsStaticAddition field value if set, zero value otherwise.
func (o *ResourcepoolQualificationType) GetIsStaticAddition() bool {
	if o == nil || IsNil(o.IsStaticAddition) {
		var ret bool
		return ret
	}
	return *o.IsStaticAddition
}

// GetIsStaticAdditionOk returns a tuple with the IsStaticAddition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolQualificationType) GetIsStaticAdditionOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStaticAddition) {
		return nil, false
	}
	return o.IsStaticAddition, true
}

// HasIsStaticAddition returns a boolean if a field has been set.
func (o *ResourcepoolQualificationType) HasIsStaticAddition() bool {
	if o != nil && !IsNil(o.IsStaticAddition) {
		return true
	}

	return false
}

// SetIsStaticAddition gets a reference to the given bool and assigns it to the IsStaticAddition field.
func (o *ResourcepoolQualificationType) SetIsStaticAddition(v bool) {
	o.IsStaticAddition = &v
}

// GetQualificationPolicies returns the QualificationPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolQualificationType) GetQualificationPolicies() []MoMoRef {
	if o == nil {
		var ret []MoMoRef
		return ret
	}
	return o.QualificationPolicies
}

// GetQualificationPoliciesOk returns a tuple with the QualificationPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolQualificationType) GetQualificationPoliciesOk() ([]MoMoRef, bool) {
	if o == nil || IsNil(o.QualificationPolicies) {
		return nil, false
	}
	return o.QualificationPolicies, true
}

// HasQualificationPolicies returns a boolean if a field has been set.
func (o *ResourcepoolQualificationType) HasQualificationPolicies() bool {
	if o != nil && !IsNil(o.QualificationPolicies) {
		return true
	}

	return false
}

// SetQualificationPolicies gets a reference to the given []MoMoRef and assigns it to the QualificationPolicies field.
func (o *ResourcepoolQualificationType) SetQualificationPolicies(v []MoMoRef) {
	o.QualificationPolicies = v
}

func (o ResourcepoolQualificationType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcepoolQualificationType) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsStaticAddition) {
		toSerialize["IsStaticAddition"] = o.IsStaticAddition
	}
	if o.QualificationPolicies != nil {
		toSerialize["QualificationPolicies"] = o.QualificationPolicies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourcepoolQualificationType) UnmarshalJSON(data []byte) (err error) {
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
	type ResourcepoolQualificationTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The resource is added statically or dynamically. If the value is true, it signifies a static addition, implying that the resource was added with fixed specifications. Conversely, if the value is false, it suggests a dynamic addition, indicating that the resource was incorporated based on certain conditions or qualifiers.
		IsStaticAddition      *bool     `json:"IsStaticAddition,omitempty"`
		QualificationPolicies []MoMoRef `json:"QualificationPolicies,omitempty"`
	}

	varResourcepoolQualificationTypeWithoutEmbeddedStruct := ResourcepoolQualificationTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varResourcepoolQualificationTypeWithoutEmbeddedStruct)
	if err == nil {
		varResourcepoolQualificationType := _ResourcepoolQualificationType{}
		varResourcepoolQualificationType.ClassId = varResourcepoolQualificationTypeWithoutEmbeddedStruct.ClassId
		varResourcepoolQualificationType.ObjectType = varResourcepoolQualificationTypeWithoutEmbeddedStruct.ObjectType
		varResourcepoolQualificationType.IsStaticAddition = varResourcepoolQualificationTypeWithoutEmbeddedStruct.IsStaticAddition
		varResourcepoolQualificationType.QualificationPolicies = varResourcepoolQualificationTypeWithoutEmbeddedStruct.QualificationPolicies
		*o = ResourcepoolQualificationType(varResourcepoolQualificationType)
	} else {
		return err
	}

	varResourcepoolQualificationType := _ResourcepoolQualificationType{}

	err = json.Unmarshal(data, &varResourcepoolQualificationType)
	if err == nil {
		o.MoBaseComplexType = varResourcepoolQualificationType.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsStaticAddition")
		delete(additionalProperties, "QualificationPolicies")

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

type NullableResourcepoolQualificationType struct {
	value *ResourcepoolQualificationType
	isSet bool
}

func (v NullableResourcepoolQualificationType) Get() *ResourcepoolQualificationType {
	return v.value
}

func (v *NullableResourcepoolQualificationType) Set(val *ResourcepoolQualificationType) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcepoolQualificationType) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcepoolQualificationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcepoolQualificationType(val *ResourcepoolQualificationType) *NullableResourcepoolQualificationType {
	return &NullableResourcepoolQualificationType{value: val, isSet: true}
}

func (v NullableResourcepoolQualificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcepoolQualificationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}