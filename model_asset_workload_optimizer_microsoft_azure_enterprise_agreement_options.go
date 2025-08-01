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

// checks if the AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions{}

// AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions Captures configuration specific to the Microsoft Azure Enterprise Agreement target for the Workload Optimizer service.
type AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enrollment number for Azure EA.
	EnrollmentNumber     *string `json:"EnrollmentNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions

// NewAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions instantiates a new AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions(classId string, objectType string) *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions {
	this := AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithDefaults() *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions {
	this := AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions{}
	var classId string = "asset.WorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "asset.WorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions" of the ClassId field.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetDefaultClassId() interface{} {
	return "asset.WorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions"
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "asset.WorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions" of the ObjectType field.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetDefaultObjectType() interface{} {
	return "asset.WorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions"
}

// GetEnrollmentNumber returns the EnrollmentNumber field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetEnrollmentNumber() string {
	if o == nil || IsNil(o.EnrollmentNumber) {
		var ret string
		return ret
	}
	return *o.EnrollmentNumber
}

// GetEnrollmentNumberOk returns a tuple with the EnrollmentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetEnrollmentNumberOk() (*string, bool) {
	if o == nil || IsNil(o.EnrollmentNumber) {
		return nil, false
	}
	return o.EnrollmentNumber, true
}

// HasEnrollmentNumber returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) HasEnrollmentNumber() bool {
	if o != nil && !IsNil(o.EnrollmentNumber) {
		return true
	}

	return false
}

// SetEnrollmentNumber gets a reference to the given string and assigns it to the EnrollmentNumber field.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) SetEnrollmentNumber(v string) {
	o.EnrollmentNumber = &v
}

func (o AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return map[string]interface{}{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return map[string]interface{}{}, errAssetServiceOptions
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.EnrollmentNumber) {
		toSerialize["EnrollmentNumber"] = o.EnrollmentNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) UnmarshalJSON(data []byte) (err error) {
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
	type AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enrollment number for Azure EA.
		EnrollmentNumber *string `json:"EnrollmentNumber,omitempty"`
	}

	varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct := AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions := _AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions{}
		varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions.ClassId = varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct.ClassId
		varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions.ObjectType = varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct.ObjectType
		varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions.EnrollmentNumber = varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct.EnrollmentNumber
		*o = AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions(varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions)
	} else {
		return err
	}

	varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions := _AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions{}

	err = json.Unmarshal(data, &varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnrollmentNumber")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

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

type NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions struct {
	value *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions
	isSet bool
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) Get() *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) Set(val *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions(val *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) *NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions {
	return &NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
