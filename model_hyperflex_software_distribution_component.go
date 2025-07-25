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

// checks if the HyperflexSoftwareDistributionComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexSoftwareDistributionComponent{}

// HyperflexSoftwareDistributionComponent A HyperFlex Software Distribution Component.
type HyperflexSoftwareDistributionComponent struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The bucket name where the files are present, if source is external cloud store.
	BucketName *string `json:"BucketName,omitempty"`
	// The HyperFlex Software Distribution Component Identifier.
	ComponentId *string `json:"ComponentId,omitempty"`
	// The HyperFlex Software Distribution Component Name.
	ComponentName *string `json:"ComponentName,omitempty"`
	// File location on the cloud storage.
	FilePath        *string  `json:"FilePath,omitempty"`
	FilesToDownload []string `json:"FilesToDownload,omitempty"`
	// The HyperFlex Software Distribution Component Version.
	Version                     *string                                                  `json:"Version,omitempty"`
	SoftwareDistributionVersion NullableHyperflexSoftwareDistributionVersionRelationship `json:"SoftwareDistributionVersion,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _HyperflexSoftwareDistributionComponent HyperflexSoftwareDistributionComponent

// NewHyperflexSoftwareDistributionComponent instantiates a new HyperflexSoftwareDistributionComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSoftwareDistributionComponent(classId string, objectType string) *HyperflexSoftwareDistributionComponent {
	this := HyperflexSoftwareDistributionComponent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSoftwareDistributionComponentWithDefaults instantiates a new HyperflexSoftwareDistributionComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSoftwareDistributionComponentWithDefaults() *HyperflexSoftwareDistributionComponent {
	this := HyperflexSoftwareDistributionComponent{}
	var classId string = "hyperflex.SoftwareDistributionComponent"
	this.ClassId = classId
	var objectType string = "hyperflex.SoftwareDistributionComponent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSoftwareDistributionComponent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSoftwareDistributionComponent) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.SoftwareDistributionComponent" of the ClassId field.
func (o *HyperflexSoftwareDistributionComponent) GetDefaultClassId() interface{} {
	return "hyperflex.SoftwareDistributionComponent"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSoftwareDistributionComponent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSoftwareDistributionComponent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.SoftwareDistributionComponent" of the ObjectType field.
func (o *HyperflexSoftwareDistributionComponent) GetDefaultObjectType() interface{} {
	return "hyperflex.SoftwareDistributionComponent"
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponent) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *HyperflexSoftwareDistributionComponent) SetBucketName(v string) {
	o.BucketName = &v
}

// GetComponentId returns the ComponentId field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponent) GetComponentId() string {
	if o == nil || IsNil(o.ComponentId) {
		var ret string
		return ret
	}
	return *o.ComponentId
}

// GetComponentIdOk returns a tuple with the ComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetComponentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentId) {
		return nil, false
	}
	return o.ComponentId, true
}

// HasComponentId returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasComponentId() bool {
	if o != nil && !IsNil(o.ComponentId) {
		return true
	}

	return false
}

// SetComponentId gets a reference to the given string and assigns it to the ComponentId field.
func (o *HyperflexSoftwareDistributionComponent) SetComponentId(v string) {
	o.ComponentId = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponent) GetComponentName() string {
	if o == nil || IsNil(o.ComponentName) {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetComponentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentName) {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasComponentName() bool {
	if o != nil && !IsNil(o.ComponentName) {
		return true
	}

	return false
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *HyperflexSoftwareDistributionComponent) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponent) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *HyperflexSoftwareDistributionComponent) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFilesToDownload returns the FilesToDownload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSoftwareDistributionComponent) GetFilesToDownload() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FilesToDownload
}

// GetFilesToDownloadOk returns a tuple with the FilesToDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSoftwareDistributionComponent) GetFilesToDownloadOk() ([]string, bool) {
	if o == nil || IsNil(o.FilesToDownload) {
		return nil, false
	}
	return o.FilesToDownload, true
}

// HasFilesToDownload returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasFilesToDownload() bool {
	if o != nil && !IsNil(o.FilesToDownload) {
		return true
	}

	return false
}

// SetFilesToDownload gets a reference to the given []string and assigns it to the FilesToDownload field.
func (o *HyperflexSoftwareDistributionComponent) SetFilesToDownload(v []string) {
	o.FilesToDownload = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponent) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexSoftwareDistributionComponent) SetVersion(v string) {
	o.Version = &v
}

// GetSoftwareDistributionVersion returns the SoftwareDistributionVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSoftwareDistributionComponent) GetSoftwareDistributionVersion() HyperflexSoftwareDistributionVersionRelationship {
	if o == nil || IsNil(o.SoftwareDistributionVersion.Get()) {
		var ret HyperflexSoftwareDistributionVersionRelationship
		return ret
	}
	return *o.SoftwareDistributionVersion.Get()
}

// GetSoftwareDistributionVersionOk returns a tuple with the SoftwareDistributionVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSoftwareDistributionComponent) GetSoftwareDistributionVersionOk() (*HyperflexSoftwareDistributionVersionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SoftwareDistributionVersion.Get(), o.SoftwareDistributionVersion.IsSet()
}

// HasSoftwareDistributionVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasSoftwareDistributionVersion() bool {
	if o != nil && o.SoftwareDistributionVersion.IsSet() {
		return true
	}

	return false
}

// SetSoftwareDistributionVersion gets a reference to the given NullableHyperflexSoftwareDistributionVersionRelationship and assigns it to the SoftwareDistributionVersion field.
func (o *HyperflexSoftwareDistributionComponent) SetSoftwareDistributionVersion(v HyperflexSoftwareDistributionVersionRelationship) {
	o.SoftwareDistributionVersion.Set(&v)
}

// SetSoftwareDistributionVersionNil sets the value for SoftwareDistributionVersion to be an explicit nil
func (o *HyperflexSoftwareDistributionComponent) SetSoftwareDistributionVersionNil() {
	o.SoftwareDistributionVersion.Set(nil)
}

// UnsetSoftwareDistributionVersion ensures that no value is present for SoftwareDistributionVersion, not even an explicit nil
func (o *HyperflexSoftwareDistributionComponent) UnsetSoftwareDistributionVersion() {
	o.SoftwareDistributionVersion.Unset()
}

func (o HyperflexSoftwareDistributionComponent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexSoftwareDistributionComponent) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BucketName) {
		toSerialize["BucketName"] = o.BucketName
	}
	if !IsNil(o.ComponentId) {
		toSerialize["ComponentId"] = o.ComponentId
	}
	if !IsNil(o.ComponentName) {
		toSerialize["ComponentName"] = o.ComponentName
	}
	if !IsNil(o.FilePath) {
		toSerialize["FilePath"] = o.FilePath
	}
	if o.FilesToDownload != nil {
		toSerialize["FilesToDownload"] = o.FilesToDownload
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if o.SoftwareDistributionVersion.IsSet() {
		toSerialize["SoftwareDistributionVersion"] = o.SoftwareDistributionVersion.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexSoftwareDistributionComponent) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexSoftwareDistributionComponentWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The bucket name where the files are present, if source is external cloud store.
		BucketName *string `json:"BucketName,omitempty"`
		// The HyperFlex Software Distribution Component Identifier.
		ComponentId *string `json:"ComponentId,omitempty"`
		// The HyperFlex Software Distribution Component Name.
		ComponentName *string `json:"ComponentName,omitempty"`
		// File location on the cloud storage.
		FilePath        *string  `json:"FilePath,omitempty"`
		FilesToDownload []string `json:"FilesToDownload,omitempty"`
		// The HyperFlex Software Distribution Component Version.
		Version                     *string                                                  `json:"Version,omitempty"`
		SoftwareDistributionVersion NullableHyperflexSoftwareDistributionVersionRelationship `json:"SoftwareDistributionVersion,omitempty"`
	}

	varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct := HyperflexSoftwareDistributionComponentWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexSoftwareDistributionComponent := _HyperflexSoftwareDistributionComponent{}
		varHyperflexSoftwareDistributionComponent.ClassId = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.ClassId
		varHyperflexSoftwareDistributionComponent.ObjectType = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.ObjectType
		varHyperflexSoftwareDistributionComponent.BucketName = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.BucketName
		varHyperflexSoftwareDistributionComponent.ComponentId = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.ComponentId
		varHyperflexSoftwareDistributionComponent.ComponentName = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.ComponentName
		varHyperflexSoftwareDistributionComponent.FilePath = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.FilePath
		varHyperflexSoftwareDistributionComponent.FilesToDownload = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.FilesToDownload
		varHyperflexSoftwareDistributionComponent.Version = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.Version
		varHyperflexSoftwareDistributionComponent.SoftwareDistributionVersion = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.SoftwareDistributionVersion
		*o = HyperflexSoftwareDistributionComponent(varHyperflexSoftwareDistributionComponent)
	} else {
		return err
	}

	varHyperflexSoftwareDistributionComponent := _HyperflexSoftwareDistributionComponent{}

	err = json.Unmarshal(data, &varHyperflexSoftwareDistributionComponent)
	if err == nil {
		o.MoBaseMo = varHyperflexSoftwareDistributionComponent.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BucketName")
		delete(additionalProperties, "ComponentId")
		delete(additionalProperties, "ComponentName")
		delete(additionalProperties, "FilePath")
		delete(additionalProperties, "FilesToDownload")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "SoftwareDistributionVersion")

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

type NullableHyperflexSoftwareDistributionComponent struct {
	value *HyperflexSoftwareDistributionComponent
	isSet bool
}

func (v NullableHyperflexSoftwareDistributionComponent) Get() *HyperflexSoftwareDistributionComponent {
	return v.value
}

func (v *NullableHyperflexSoftwareDistributionComponent) Set(val *HyperflexSoftwareDistributionComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSoftwareDistributionComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSoftwareDistributionComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSoftwareDistributionComponent(val *HyperflexSoftwareDistributionComponent) *NullableHyperflexSoftwareDistributionComponent {
	return &NullableHyperflexSoftwareDistributionComponent{value: val, isSet: true}
}

func (v NullableHyperflexSoftwareDistributionComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSoftwareDistributionComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
