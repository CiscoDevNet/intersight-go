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

// checks if the MarketplaceUseCaseVersionResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplaceUseCaseVersionResource{}

// MarketplaceUseCaseVersionResource A MO describing the resources that belong to a UseCaseVersion.
type MarketplaceUseCaseVersionResource struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A string ID for each use case.
	ResourceId *string `json:"ResourceId,omitempty"`
	// A string resource type for each use case.
	ResourceType         *string `json:"ResourceType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarketplaceUseCaseVersionResource MarketplaceUseCaseVersionResource

// NewMarketplaceUseCaseVersionResource instantiates a new MarketplaceUseCaseVersionResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceUseCaseVersionResource(classId string, objectType string) *MarketplaceUseCaseVersionResource {
	this := MarketplaceUseCaseVersionResource{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMarketplaceUseCaseVersionResourceWithDefaults instantiates a new MarketplaceUseCaseVersionResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceUseCaseVersionResourceWithDefaults() *MarketplaceUseCaseVersionResource {
	this := MarketplaceUseCaseVersionResource{}
	var classId string = "marketplace.UseCaseVersionResource"
	this.ClassId = classId
	var objectType string = "marketplace.UseCaseVersionResource"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MarketplaceUseCaseVersionResource) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseVersionResource) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MarketplaceUseCaseVersionResource) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "marketplace.UseCaseVersionResource" of the ClassId field.
func (o *MarketplaceUseCaseVersionResource) GetDefaultClassId() interface{} {
	return "marketplace.UseCaseVersionResource"
}

// GetObjectType returns the ObjectType field value
func (o *MarketplaceUseCaseVersionResource) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseVersionResource) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MarketplaceUseCaseVersionResource) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "marketplace.UseCaseVersionResource" of the ObjectType field.
func (o *MarketplaceUseCaseVersionResource) GetDefaultObjectType() interface{} {
	return "marketplace.UseCaseVersionResource"
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *MarketplaceUseCaseVersionResource) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseVersionResource) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *MarketplaceUseCaseVersionResource) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *MarketplaceUseCaseVersionResource) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *MarketplaceUseCaseVersionResource) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseVersionResource) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *MarketplaceUseCaseVersionResource) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *MarketplaceUseCaseVersionResource) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o MarketplaceUseCaseVersionResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketplaceUseCaseVersionResource) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ResourceId) {
		toSerialize["ResourceId"] = o.ResourceId
	}
	if !IsNil(o.ResourceType) {
		toSerialize["ResourceType"] = o.ResourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarketplaceUseCaseVersionResource) UnmarshalJSON(data []byte) (err error) {
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
	type MarketplaceUseCaseVersionResourceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// A string ID for each use case.
		ResourceId *string `json:"ResourceId,omitempty"`
		// A string resource type for each use case.
		ResourceType *string `json:"ResourceType,omitempty"`
	}

	varMarketplaceUseCaseVersionResourceWithoutEmbeddedStruct := MarketplaceUseCaseVersionResourceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMarketplaceUseCaseVersionResourceWithoutEmbeddedStruct)
	if err == nil {
		varMarketplaceUseCaseVersionResource := _MarketplaceUseCaseVersionResource{}
		varMarketplaceUseCaseVersionResource.ClassId = varMarketplaceUseCaseVersionResourceWithoutEmbeddedStruct.ClassId
		varMarketplaceUseCaseVersionResource.ObjectType = varMarketplaceUseCaseVersionResourceWithoutEmbeddedStruct.ObjectType
		varMarketplaceUseCaseVersionResource.ResourceId = varMarketplaceUseCaseVersionResourceWithoutEmbeddedStruct.ResourceId
		varMarketplaceUseCaseVersionResource.ResourceType = varMarketplaceUseCaseVersionResourceWithoutEmbeddedStruct.ResourceType
		*o = MarketplaceUseCaseVersionResource(varMarketplaceUseCaseVersionResource)
	} else {
		return err
	}

	varMarketplaceUseCaseVersionResource := _MarketplaceUseCaseVersionResource{}

	err = json.Unmarshal(data, &varMarketplaceUseCaseVersionResource)
	if err == nil {
		o.MoBaseComplexType = varMarketplaceUseCaseVersionResource.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ResourceId")
		delete(additionalProperties, "ResourceType")

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

type NullableMarketplaceUseCaseVersionResource struct {
	value *MarketplaceUseCaseVersionResource
	isSet bool
}

func (v NullableMarketplaceUseCaseVersionResource) Get() *MarketplaceUseCaseVersionResource {
	return v.value
}

func (v *NullableMarketplaceUseCaseVersionResource) Set(val *MarketplaceUseCaseVersionResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceUseCaseVersionResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceUseCaseVersionResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceUseCaseVersionResource(val *MarketplaceUseCaseVersionResource) *NullableMarketplaceUseCaseVersionResource {
	return &NullableMarketplaceUseCaseVersionResource{value: val, isSet: true}
}

func (v NullableMarketplaceUseCaseVersionResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceUseCaseVersionResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
