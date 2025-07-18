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

// checks if the AssetVirtualizationAmazonWebServiceOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetVirtualizationAmazonWebServiceOptions{}

// AssetVirtualizationAmazonWebServiceOptions Captures Virtualization Service configuration specific to AWS.
type AssetVirtualizationAmazonWebServiceOptions struct {
	AssetVirtualizationCloudOptions
	AdditionalProperties map[string]interface{}
}

type _AssetVirtualizationAmazonWebServiceOptions AssetVirtualizationAmazonWebServiceOptions

// NewAssetVirtualizationAmazonWebServiceOptions instantiates a new AssetVirtualizationAmazonWebServiceOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetVirtualizationAmazonWebServiceOptions(classId string, objectType string) *AssetVirtualizationAmazonWebServiceOptions {
	this := AssetVirtualizationAmazonWebServiceOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetVirtualizationAmazonWebServiceOptionsWithDefaults instantiates a new AssetVirtualizationAmazonWebServiceOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetVirtualizationAmazonWebServiceOptionsWithDefaults() *AssetVirtualizationAmazonWebServiceOptions {
	this := AssetVirtualizationAmazonWebServiceOptions{}
	return &this
}

func (o AssetVirtualizationAmazonWebServiceOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetVirtualizationAmazonWebServiceOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetVirtualizationCloudOptions, errAssetVirtualizationCloudOptions := json.Marshal(o.AssetVirtualizationCloudOptions)
	if errAssetVirtualizationCloudOptions != nil {
		return map[string]interface{}{}, errAssetVirtualizationCloudOptions
	}
	errAssetVirtualizationCloudOptions = json.Unmarshal([]byte(serializedAssetVirtualizationCloudOptions), &toSerialize)
	if errAssetVirtualizationCloudOptions != nil {
		return map[string]interface{}{}, errAssetVirtualizationCloudOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetVirtualizationAmazonWebServiceOptions) UnmarshalJSON(data []byte) (err error) {
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
	type AssetVirtualizationAmazonWebServiceOptionsWithoutEmbeddedStruct struct {
	}

	varAssetVirtualizationAmazonWebServiceOptionsWithoutEmbeddedStruct := AssetVirtualizationAmazonWebServiceOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetVirtualizationAmazonWebServiceOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetVirtualizationAmazonWebServiceOptions := _AssetVirtualizationAmazonWebServiceOptions{}
		*o = AssetVirtualizationAmazonWebServiceOptions(varAssetVirtualizationAmazonWebServiceOptions)
	} else {
		return err
	}

	varAssetVirtualizationAmazonWebServiceOptions := _AssetVirtualizationAmazonWebServiceOptions{}

	err = json.Unmarshal(data, &varAssetVirtualizationAmazonWebServiceOptions)
	if err == nil {
		o.AssetVirtualizationCloudOptions = varAssetVirtualizationAmazonWebServiceOptions.AssetVirtualizationCloudOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectAssetVirtualizationCloudOptions := reflect.ValueOf(o.AssetVirtualizationCloudOptions)
		for i := 0; i < reflectAssetVirtualizationCloudOptions.Type().NumField(); i++ {
			t := reflectAssetVirtualizationCloudOptions.Type().Field(i)

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

type NullableAssetVirtualizationAmazonWebServiceOptions struct {
	value *AssetVirtualizationAmazonWebServiceOptions
	isSet bool
}

func (v NullableAssetVirtualizationAmazonWebServiceOptions) Get() *AssetVirtualizationAmazonWebServiceOptions {
	return v.value
}

func (v *NullableAssetVirtualizationAmazonWebServiceOptions) Set(val *AssetVirtualizationAmazonWebServiceOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetVirtualizationAmazonWebServiceOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetVirtualizationAmazonWebServiceOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetVirtualizationAmazonWebServiceOptions(val *AssetVirtualizationAmazonWebServiceOptions) *NullableAssetVirtualizationAmazonWebServiceOptions {
	return &NullableAssetVirtualizationAmazonWebServiceOptions{value: val, isSet: true}
}

func (v NullableAssetVirtualizationAmazonWebServiceOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetVirtualizationAmazonWebServiceOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
