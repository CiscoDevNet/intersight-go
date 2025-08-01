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

// checks if the HyperflexFeatureLimitInternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexFeatureLimitInternal{}

// HyperflexFeatureLimitInternal The HyperFlex installer feature limits for internal system use.
type HyperflexFeatureLimitInternal struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                  `json:"ObjectType"`
	FeatureLimitEntries  []HyperflexFeatureLimitEntry            `json:"FeatureLimitEntries,omitempty"`
	AppCatalog           NullableHyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexFeatureLimitInternal HyperflexFeatureLimitInternal

// NewHyperflexFeatureLimitInternal instantiates a new HyperflexFeatureLimitInternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexFeatureLimitInternal(classId string, objectType string) *HyperflexFeatureLimitInternal {
	this := HyperflexFeatureLimitInternal{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexFeatureLimitInternalWithDefaults instantiates a new HyperflexFeatureLimitInternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexFeatureLimitInternalWithDefaults() *HyperflexFeatureLimitInternal {
	this := HyperflexFeatureLimitInternal{}
	var classId string = "hyperflex.FeatureLimitInternal"
	this.ClassId = classId
	var objectType string = "hyperflex.FeatureLimitInternal"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexFeatureLimitInternal) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexFeatureLimitInternal) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexFeatureLimitInternal) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.FeatureLimitInternal" of the ClassId field.
func (o *HyperflexFeatureLimitInternal) GetDefaultClassId() interface{} {
	return "hyperflex.FeatureLimitInternal"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexFeatureLimitInternal) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexFeatureLimitInternal) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexFeatureLimitInternal) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.FeatureLimitInternal" of the ObjectType field.
func (o *HyperflexFeatureLimitInternal) GetDefaultObjectType() interface{} {
	return "hyperflex.FeatureLimitInternal"
}

// GetFeatureLimitEntries returns the FeatureLimitEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexFeatureLimitInternal) GetFeatureLimitEntries() []HyperflexFeatureLimitEntry {
	if o == nil {
		var ret []HyperflexFeatureLimitEntry
		return ret
	}
	return o.FeatureLimitEntries
}

// GetFeatureLimitEntriesOk returns a tuple with the FeatureLimitEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexFeatureLimitInternal) GetFeatureLimitEntriesOk() ([]HyperflexFeatureLimitEntry, bool) {
	if o == nil || IsNil(o.FeatureLimitEntries) {
		return nil, false
	}
	return o.FeatureLimitEntries, true
}

// HasFeatureLimitEntries returns a boolean if a field has been set.
func (o *HyperflexFeatureLimitInternal) HasFeatureLimitEntries() bool {
	if o != nil && !IsNil(o.FeatureLimitEntries) {
		return true
	}

	return false
}

// SetFeatureLimitEntries gets a reference to the given []HyperflexFeatureLimitEntry and assigns it to the FeatureLimitEntries field.
func (o *HyperflexFeatureLimitInternal) SetFeatureLimitEntries(v []HyperflexFeatureLimitEntry) {
	o.FeatureLimitEntries = v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexFeatureLimitInternal) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || IsNil(o.AppCatalog.Get()) {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog.Get()
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexFeatureLimitInternal) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppCatalog.Get(), o.AppCatalog.IsSet()
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HyperflexFeatureLimitInternal) HasAppCatalog() bool {
	if o != nil && o.AppCatalog.IsSet() {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given NullableHyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HyperflexFeatureLimitInternal) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog.Set(&v)
}

// SetAppCatalogNil sets the value for AppCatalog to be an explicit nil
func (o *HyperflexFeatureLimitInternal) SetAppCatalogNil() {
	o.AppCatalog.Set(nil)
}

// UnsetAppCatalog ensures that no value is present for AppCatalog, not even an explicit nil
func (o *HyperflexFeatureLimitInternal) UnsetAppCatalog() {
	o.AppCatalog.Unset()
}

func (o HyperflexFeatureLimitInternal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexFeatureLimitInternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.FeatureLimitEntries != nil {
		toSerialize["FeatureLimitEntries"] = o.FeatureLimitEntries
	}
	if o.AppCatalog.IsSet() {
		toSerialize["AppCatalog"] = o.AppCatalog.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexFeatureLimitInternal) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexFeatureLimitInternalWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType          string                                  `json:"ObjectType"`
		FeatureLimitEntries []HyperflexFeatureLimitEntry            `json:"FeatureLimitEntries,omitempty"`
		AppCatalog          NullableHyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	}

	varHyperflexFeatureLimitInternalWithoutEmbeddedStruct := HyperflexFeatureLimitInternalWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexFeatureLimitInternalWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexFeatureLimitInternal := _HyperflexFeatureLimitInternal{}
		varHyperflexFeatureLimitInternal.ClassId = varHyperflexFeatureLimitInternalWithoutEmbeddedStruct.ClassId
		varHyperflexFeatureLimitInternal.ObjectType = varHyperflexFeatureLimitInternalWithoutEmbeddedStruct.ObjectType
		varHyperflexFeatureLimitInternal.FeatureLimitEntries = varHyperflexFeatureLimitInternalWithoutEmbeddedStruct.FeatureLimitEntries
		varHyperflexFeatureLimitInternal.AppCatalog = varHyperflexFeatureLimitInternalWithoutEmbeddedStruct.AppCatalog
		*o = HyperflexFeatureLimitInternal(varHyperflexFeatureLimitInternal)
	} else {
		return err
	}

	varHyperflexFeatureLimitInternal := _HyperflexFeatureLimitInternal{}

	err = json.Unmarshal(data, &varHyperflexFeatureLimitInternal)
	if err == nil {
		o.MoBaseMo = varHyperflexFeatureLimitInternal.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FeatureLimitEntries")
		delete(additionalProperties, "AppCatalog")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableHyperflexFeatureLimitInternal struct {
	value *HyperflexFeatureLimitInternal
	isSet bool
}

func (v NullableHyperflexFeatureLimitInternal) Get() *HyperflexFeatureLimitInternal {
	return v.value
}

func (v *NullableHyperflexFeatureLimitInternal) Set(val *HyperflexFeatureLimitInternal) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexFeatureLimitInternal) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexFeatureLimitInternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexFeatureLimitInternal(val *HyperflexFeatureLimitInternal) *NullableHyperflexFeatureLimitInternal {
	return &NullableHyperflexFeatureLimitInternal{value: val, isSet: true}
}

func (v NullableHyperflexFeatureLimitInternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexFeatureLimitInternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
