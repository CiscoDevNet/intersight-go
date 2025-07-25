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

// checks if the CapabilityUnsupportedFeatureConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityUnsupportedFeatureConfig{}

// CapabilityUnsupportedFeatureConfig UnsUpported Feature Configuration specific to the adapter feature and generation.
type CapabilityUnsupportedFeatureConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The adapter generations that support this feature. * `4` - Fourth generation adapters (14xx). The PIDs of these adapters end with the string 04. * `2` - Second generation VIC adapters (12xx). The PIDs of these adapters end with the string 02. * `3` - Third generation adapters (13xx). The PIDs of these adapters end with the string 03. * `5` - Fifth generation adapters (15xx). The PIDs of these adapters contain the V5 string.
	Generation           *int32   `json:"Generation,omitempty"`
	UnsupportdFeatures   []string `json:"UnsupportdFeatures,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityUnsupportedFeatureConfig CapabilityUnsupportedFeatureConfig

// NewCapabilityUnsupportedFeatureConfig instantiates a new CapabilityUnsupportedFeatureConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityUnsupportedFeatureConfig(classId string, objectType string) *CapabilityUnsupportedFeatureConfig {
	this := CapabilityUnsupportedFeatureConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var generation int32 = 4
	this.Generation = &generation
	return &this
}

// NewCapabilityUnsupportedFeatureConfigWithDefaults instantiates a new CapabilityUnsupportedFeatureConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityUnsupportedFeatureConfigWithDefaults() *CapabilityUnsupportedFeatureConfig {
	this := CapabilityUnsupportedFeatureConfig{}
	var classId string = "capability.UnsupportedFeatureConfig"
	this.ClassId = classId
	var objectType string = "capability.UnsupportedFeatureConfig"
	this.ObjectType = objectType
	var generation int32 = 4
	this.Generation = &generation
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityUnsupportedFeatureConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityUnsupportedFeatureConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityUnsupportedFeatureConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.UnsupportedFeatureConfig" of the ClassId field.
func (o *CapabilityUnsupportedFeatureConfig) GetDefaultClassId() interface{} {
	return "capability.UnsupportedFeatureConfig"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityUnsupportedFeatureConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityUnsupportedFeatureConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityUnsupportedFeatureConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.UnsupportedFeatureConfig" of the ObjectType field.
func (o *CapabilityUnsupportedFeatureConfig) GetDefaultObjectType() interface{} {
	return "capability.UnsupportedFeatureConfig"
}

// GetGeneration returns the Generation field value if set, zero value otherwise.
func (o *CapabilityUnsupportedFeatureConfig) GetGeneration() int32 {
	if o == nil || IsNil(o.Generation) {
		var ret int32
		return ret
	}
	return *o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityUnsupportedFeatureConfig) GetGenerationOk() (*int32, bool) {
	if o == nil || IsNil(o.Generation) {
		return nil, false
	}
	return o.Generation, true
}

// HasGeneration returns a boolean if a field has been set.
func (o *CapabilityUnsupportedFeatureConfig) HasGeneration() bool {
	if o != nil && !IsNil(o.Generation) {
		return true
	}

	return false
}

// SetGeneration gets a reference to the given int32 and assigns it to the Generation field.
func (o *CapabilityUnsupportedFeatureConfig) SetGeneration(v int32) {
	o.Generation = &v
}

// GetUnsupportdFeatures returns the UnsupportdFeatures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityUnsupportedFeatureConfig) GetUnsupportdFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UnsupportdFeatures
}

// GetUnsupportdFeaturesOk returns a tuple with the UnsupportdFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityUnsupportedFeatureConfig) GetUnsupportdFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.UnsupportdFeatures) {
		return nil, false
	}
	return o.UnsupportdFeatures, true
}

// HasUnsupportdFeatures returns a boolean if a field has been set.
func (o *CapabilityUnsupportedFeatureConfig) HasUnsupportdFeatures() bool {
	if o != nil && !IsNil(o.UnsupportdFeatures) {
		return true
	}

	return false
}

// SetUnsupportdFeatures gets a reference to the given []string and assigns it to the UnsupportdFeatures field.
func (o *CapabilityUnsupportedFeatureConfig) SetUnsupportdFeatures(v []string) {
	o.UnsupportdFeatures = v
}

func (o CapabilityUnsupportedFeatureConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityUnsupportedFeatureConfig) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Generation) {
		toSerialize["Generation"] = o.Generation
	}
	if o.UnsupportdFeatures != nil {
		toSerialize["UnsupportdFeatures"] = o.UnsupportdFeatures
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityUnsupportedFeatureConfig) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityUnsupportedFeatureConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The adapter generations that support this feature. * `4` - Fourth generation adapters (14xx). The PIDs of these adapters end with the string 04. * `2` - Second generation VIC adapters (12xx). The PIDs of these adapters end with the string 02. * `3` - Third generation adapters (13xx). The PIDs of these adapters end with the string 03. * `5` - Fifth generation adapters (15xx). The PIDs of these adapters contain the V5 string.
		Generation         *int32   `json:"Generation,omitempty"`
		UnsupportdFeatures []string `json:"UnsupportdFeatures,omitempty"`
	}

	varCapabilityUnsupportedFeatureConfigWithoutEmbeddedStruct := CapabilityUnsupportedFeatureConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityUnsupportedFeatureConfigWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityUnsupportedFeatureConfig := _CapabilityUnsupportedFeatureConfig{}
		varCapabilityUnsupportedFeatureConfig.ClassId = varCapabilityUnsupportedFeatureConfigWithoutEmbeddedStruct.ClassId
		varCapabilityUnsupportedFeatureConfig.ObjectType = varCapabilityUnsupportedFeatureConfigWithoutEmbeddedStruct.ObjectType
		varCapabilityUnsupportedFeatureConfig.Generation = varCapabilityUnsupportedFeatureConfigWithoutEmbeddedStruct.Generation
		varCapabilityUnsupportedFeatureConfig.UnsupportdFeatures = varCapabilityUnsupportedFeatureConfigWithoutEmbeddedStruct.UnsupportdFeatures
		*o = CapabilityUnsupportedFeatureConfig(varCapabilityUnsupportedFeatureConfig)
	} else {
		return err
	}

	varCapabilityUnsupportedFeatureConfig := _CapabilityUnsupportedFeatureConfig{}

	err = json.Unmarshal(data, &varCapabilityUnsupportedFeatureConfig)
	if err == nil {
		o.MoBaseComplexType = varCapabilityUnsupportedFeatureConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Generation")
		delete(additionalProperties, "UnsupportdFeatures")

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

type NullableCapabilityUnsupportedFeatureConfig struct {
	value *CapabilityUnsupportedFeatureConfig
	isSet bool
}

func (v NullableCapabilityUnsupportedFeatureConfig) Get() *CapabilityUnsupportedFeatureConfig {
	return v.value
}

func (v *NullableCapabilityUnsupportedFeatureConfig) Set(val *CapabilityUnsupportedFeatureConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityUnsupportedFeatureConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityUnsupportedFeatureConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityUnsupportedFeatureConfig(val *CapabilityUnsupportedFeatureConfig) *NullableCapabilityUnsupportedFeatureConfig {
	return &NullableCapabilityUnsupportedFeatureConfig{value: val, isSet: true}
}

func (v NullableCapabilityUnsupportedFeatureConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityUnsupportedFeatureConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
