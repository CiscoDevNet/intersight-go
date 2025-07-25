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

// checks if the KubernetesNvidiaGpuProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesNvidiaGpuProduct{}

// KubernetesNvidiaGpuProduct Information of a Nvidia GPU product.
type KubernetesNvidiaGpuProduct struct {
	KubernetesBaseGpuProduct
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// True if this Nvidia GPU supports MIG.
	MigCapable           *bool    `json:"MigCapable,omitempty"`
	MigProfiles          []string `json:"MigProfiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNvidiaGpuProduct KubernetesNvidiaGpuProduct

// NewKubernetesNvidiaGpuProduct instantiates a new KubernetesNvidiaGpuProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNvidiaGpuProduct(classId string, objectType string) *KubernetesNvidiaGpuProduct {
	this := KubernetesNvidiaGpuProduct{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesNvidiaGpuProductWithDefaults instantiates a new KubernetesNvidiaGpuProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNvidiaGpuProductWithDefaults() *KubernetesNvidiaGpuProduct {
	this := KubernetesNvidiaGpuProduct{}
	var classId string = "kubernetes.NvidiaGpuProduct"
	this.ClassId = classId
	var objectType string = "kubernetes.NvidiaGpuProduct"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNvidiaGpuProduct) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNvidiaGpuProduct) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNvidiaGpuProduct) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.NvidiaGpuProduct" of the ClassId field.
func (o *KubernetesNvidiaGpuProduct) GetDefaultClassId() interface{} {
	return "kubernetes.NvidiaGpuProduct"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNvidiaGpuProduct) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNvidiaGpuProduct) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNvidiaGpuProduct) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.NvidiaGpuProduct" of the ObjectType field.
func (o *KubernetesNvidiaGpuProduct) GetDefaultObjectType() interface{} {
	return "kubernetes.NvidiaGpuProduct"
}

// GetMigCapable returns the MigCapable field value if set, zero value otherwise.
func (o *KubernetesNvidiaGpuProduct) GetMigCapable() bool {
	if o == nil || IsNil(o.MigCapable) {
		var ret bool
		return ret
	}
	return *o.MigCapable
}

// GetMigCapableOk returns a tuple with the MigCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNvidiaGpuProduct) GetMigCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.MigCapable) {
		return nil, false
	}
	return o.MigCapable, true
}

// HasMigCapable returns a boolean if a field has been set.
func (o *KubernetesNvidiaGpuProduct) HasMigCapable() bool {
	if o != nil && !IsNil(o.MigCapable) {
		return true
	}

	return false
}

// SetMigCapable gets a reference to the given bool and assigns it to the MigCapable field.
func (o *KubernetesNvidiaGpuProduct) SetMigCapable(v bool) {
	o.MigCapable = &v
}

// GetMigProfiles returns the MigProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNvidiaGpuProduct) GetMigProfiles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MigProfiles
}

// GetMigProfilesOk returns a tuple with the MigProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNvidiaGpuProduct) GetMigProfilesOk() ([]string, bool) {
	if o == nil || IsNil(o.MigProfiles) {
		return nil, false
	}
	return o.MigProfiles, true
}

// HasMigProfiles returns a boolean if a field has been set.
func (o *KubernetesNvidiaGpuProduct) HasMigProfiles() bool {
	if o != nil && !IsNil(o.MigProfiles) {
		return true
	}

	return false
}

// SetMigProfiles gets a reference to the given []string and assigns it to the MigProfiles field.
func (o *KubernetesNvidiaGpuProduct) SetMigProfiles(v []string) {
	o.MigProfiles = v
}

func (o KubernetesNvidiaGpuProduct) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesNvidiaGpuProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesBaseGpuProduct, errKubernetesBaseGpuProduct := json.Marshal(o.KubernetesBaseGpuProduct)
	if errKubernetesBaseGpuProduct != nil {
		return map[string]interface{}{}, errKubernetesBaseGpuProduct
	}
	errKubernetesBaseGpuProduct = json.Unmarshal([]byte(serializedKubernetesBaseGpuProduct), &toSerialize)
	if errKubernetesBaseGpuProduct != nil {
		return map[string]interface{}{}, errKubernetesBaseGpuProduct
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.MigCapable) {
		toSerialize["MigCapable"] = o.MigCapable
	}
	if o.MigProfiles != nil {
		toSerialize["MigProfiles"] = o.MigProfiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesNvidiaGpuProduct) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesNvidiaGpuProductWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// True if this Nvidia GPU supports MIG.
		MigCapable  *bool    `json:"MigCapable,omitempty"`
		MigProfiles []string `json:"MigProfiles,omitempty"`
	}

	varKubernetesNvidiaGpuProductWithoutEmbeddedStruct := KubernetesNvidiaGpuProductWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesNvidiaGpuProductWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesNvidiaGpuProduct := _KubernetesNvidiaGpuProduct{}
		varKubernetesNvidiaGpuProduct.ClassId = varKubernetesNvidiaGpuProductWithoutEmbeddedStruct.ClassId
		varKubernetesNvidiaGpuProduct.ObjectType = varKubernetesNvidiaGpuProductWithoutEmbeddedStruct.ObjectType
		varKubernetesNvidiaGpuProduct.MigCapable = varKubernetesNvidiaGpuProductWithoutEmbeddedStruct.MigCapable
		varKubernetesNvidiaGpuProduct.MigProfiles = varKubernetesNvidiaGpuProductWithoutEmbeddedStruct.MigProfiles
		*o = KubernetesNvidiaGpuProduct(varKubernetesNvidiaGpuProduct)
	} else {
		return err
	}

	varKubernetesNvidiaGpuProduct := _KubernetesNvidiaGpuProduct{}

	err = json.Unmarshal(data, &varKubernetesNvidiaGpuProduct)
	if err == nil {
		o.KubernetesBaseGpuProduct = varKubernetesNvidiaGpuProduct.KubernetesBaseGpuProduct
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MigCapable")
		delete(additionalProperties, "MigProfiles")

		// remove fields from embedded structs
		reflectKubernetesBaseGpuProduct := reflect.ValueOf(o.KubernetesBaseGpuProduct)
		for i := 0; i < reflectKubernetesBaseGpuProduct.Type().NumField(); i++ {
			t := reflectKubernetesBaseGpuProduct.Type().Field(i)

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

type NullableKubernetesNvidiaGpuProduct struct {
	value *KubernetesNvidiaGpuProduct
	isSet bool
}

func (v NullableKubernetesNvidiaGpuProduct) Get() *KubernetesNvidiaGpuProduct {
	return v.value
}

func (v *NullableKubernetesNvidiaGpuProduct) Set(val *KubernetesNvidiaGpuProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNvidiaGpuProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNvidiaGpuProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNvidiaGpuProduct(val *KubernetesNvidiaGpuProduct) *NullableKubernetesNvidiaGpuProduct {
	return &NullableKubernetesNvidiaGpuProduct{value: val, isSet: true}
}

func (v NullableKubernetesNvidiaGpuProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNvidiaGpuProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
