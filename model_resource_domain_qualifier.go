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

// checks if the ResourceDomainQualifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceDomainQualifier{}

// ResourceDomainQualifier Qualify or filter the server based on domain properties such as DomainNames and Fabric Interconnect product IDs (PIDs).
type ResourceDomainQualifier struct {
	ResourceResourceQualifier
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType             string   `json:"ObjectType"`
	DomainNames            []string `json:"DomainNames,omitempty"`
	FabricInterConnectPids []string `json:"FabricInterConnectPids,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _ResourceDomainQualifier ResourceDomainQualifier

// NewResourceDomainQualifier instantiates a new ResourceDomainQualifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceDomainQualifier(classId string, objectType string) *ResourceDomainQualifier {
	this := ResourceDomainQualifier{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourceDomainQualifierWithDefaults instantiates a new ResourceDomainQualifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceDomainQualifierWithDefaults() *ResourceDomainQualifier {
	this := ResourceDomainQualifier{}
	var classId string = "resource.DomainQualifier"
	this.ClassId = classId
	var objectType string = "resource.DomainQualifier"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceDomainQualifier) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceDomainQualifier) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceDomainQualifier) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "resource.DomainQualifier" of the ClassId field.
func (o *ResourceDomainQualifier) GetDefaultClassId() interface{} {
	return "resource.DomainQualifier"
}

// GetObjectType returns the ObjectType field value
func (o *ResourceDomainQualifier) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceDomainQualifier) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceDomainQualifier) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "resource.DomainQualifier" of the ObjectType field.
func (o *ResourceDomainQualifier) GetDefaultObjectType() interface{} {
	return "resource.DomainQualifier"
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceDomainQualifier) GetDomainNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceDomainQualifier) GetDomainNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainNames) {
		return nil, false
	}
	return o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *ResourceDomainQualifier) HasDomainNames() bool {
	if o != nil && !IsNil(o.DomainNames) {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *ResourceDomainQualifier) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetFabricInterConnectPids returns the FabricInterConnectPids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceDomainQualifier) GetFabricInterConnectPids() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FabricInterConnectPids
}

// GetFabricInterConnectPidsOk returns a tuple with the FabricInterConnectPids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceDomainQualifier) GetFabricInterConnectPidsOk() ([]string, bool) {
	if o == nil || IsNil(o.FabricInterConnectPids) {
		return nil, false
	}
	return o.FabricInterConnectPids, true
}

// HasFabricInterConnectPids returns a boolean if a field has been set.
func (o *ResourceDomainQualifier) HasFabricInterConnectPids() bool {
	if o != nil && !IsNil(o.FabricInterConnectPids) {
		return true
	}

	return false
}

// SetFabricInterConnectPids gets a reference to the given []string and assigns it to the FabricInterConnectPids field.
func (o *ResourceDomainQualifier) SetFabricInterConnectPids(v []string) {
	o.FabricInterConnectPids = v
}

func (o ResourceDomainQualifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceDomainQualifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedResourceResourceQualifier, errResourceResourceQualifier := json.Marshal(o.ResourceResourceQualifier)
	if errResourceResourceQualifier != nil {
		return map[string]interface{}{}, errResourceResourceQualifier
	}
	errResourceResourceQualifier = json.Unmarshal([]byte(serializedResourceResourceQualifier), &toSerialize)
	if errResourceResourceQualifier != nil {
		return map[string]interface{}{}, errResourceResourceQualifier
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.DomainNames != nil {
		toSerialize["DomainNames"] = o.DomainNames
	}
	if o.FabricInterConnectPids != nil {
		toSerialize["FabricInterConnectPids"] = o.FabricInterConnectPids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceDomainQualifier) UnmarshalJSON(data []byte) (err error) {
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
	type ResourceDomainQualifierWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType             string   `json:"ObjectType"`
		DomainNames            []string `json:"DomainNames,omitempty"`
		FabricInterConnectPids []string `json:"FabricInterConnectPids,omitempty"`
	}

	varResourceDomainQualifierWithoutEmbeddedStruct := ResourceDomainQualifierWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varResourceDomainQualifierWithoutEmbeddedStruct)
	if err == nil {
		varResourceDomainQualifier := _ResourceDomainQualifier{}
		varResourceDomainQualifier.ClassId = varResourceDomainQualifierWithoutEmbeddedStruct.ClassId
		varResourceDomainQualifier.ObjectType = varResourceDomainQualifierWithoutEmbeddedStruct.ObjectType
		varResourceDomainQualifier.DomainNames = varResourceDomainQualifierWithoutEmbeddedStruct.DomainNames
		varResourceDomainQualifier.FabricInterConnectPids = varResourceDomainQualifierWithoutEmbeddedStruct.FabricInterConnectPids
		*o = ResourceDomainQualifier(varResourceDomainQualifier)
	} else {
		return err
	}

	varResourceDomainQualifier := _ResourceDomainQualifier{}

	err = json.Unmarshal(data, &varResourceDomainQualifier)
	if err == nil {
		o.ResourceResourceQualifier = varResourceDomainQualifier.ResourceResourceQualifier
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DomainNames")
		delete(additionalProperties, "FabricInterConnectPids")

		// remove fields from embedded structs
		reflectResourceResourceQualifier := reflect.ValueOf(o.ResourceResourceQualifier)
		for i := 0; i < reflectResourceResourceQualifier.Type().NumField(); i++ {
			t := reflectResourceResourceQualifier.Type().Field(i)

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

type NullableResourceDomainQualifier struct {
	value *ResourceDomainQualifier
	isSet bool
}

func (v NullableResourceDomainQualifier) Get() *ResourceDomainQualifier {
	return v.value
}

func (v *NullableResourceDomainQualifier) Set(val *ResourceDomainQualifier) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceDomainQualifier) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceDomainQualifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceDomainQualifier(val *ResourceDomainQualifier) *NullableResourceDomainQualifier {
	return &NullableResourceDomainQualifier{value: val, isSet: true}
}

func (v NullableResourceDomainQualifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceDomainQualifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}