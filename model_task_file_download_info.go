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

// checks if the TaskFileDownloadInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskFileDownloadInfo{}

// TaskFileDownloadInfo Specifies a download path or location supports S3 currently.
type TaskFileDownloadInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When the type of the download is inline, it will read the file from the contents here.
	Contents *string `json:"Contents,omitempty"`
	// The path of the download from the specified source location for type S3, then this is the object key.
	Path *string `json:"Path,omitempty"`
	// The source of the download location and if type is S3, then this is the bucket name. In case of MoRef download type  the source will have the Moid of the workflow definition.
	Source *string `json:"Source,omitempty"`
	// The type of download location is captured in type. * `S3` - Download workflow from S3. * `Local` - Download workflow from local path. * `Inline` - Workflow is provided inline. * `MoRef` - Start an existing workflow registered with given Moid.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaskFileDownloadInfo TaskFileDownloadInfo

// NewTaskFileDownloadInfo instantiates a new TaskFileDownloadInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskFileDownloadInfo(classId string, objectType string) *TaskFileDownloadInfo {
	this := TaskFileDownloadInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "S3"
	this.Type = &type_
	return &this
}

// NewTaskFileDownloadInfoWithDefaults instantiates a new TaskFileDownloadInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskFileDownloadInfoWithDefaults() *TaskFileDownloadInfo {
	this := TaskFileDownloadInfo{}
	var classId string = "task.FileDownloadInfo"
	this.ClassId = classId
	var objectType string = "task.FileDownloadInfo"
	this.ObjectType = objectType
	var type_ string = "S3"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *TaskFileDownloadInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TaskFileDownloadInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TaskFileDownloadInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "task.FileDownloadInfo" of the ClassId field.
func (o *TaskFileDownloadInfo) GetDefaultClassId() interface{} {
	return "task.FileDownloadInfo"
}

// GetObjectType returns the ObjectType field value
func (o *TaskFileDownloadInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TaskFileDownloadInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TaskFileDownloadInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "task.FileDownloadInfo" of the ObjectType field.
func (o *TaskFileDownloadInfo) GetDefaultObjectType() interface{} {
	return "task.FileDownloadInfo"
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *TaskFileDownloadInfo) GetContents() string {
	if o == nil || IsNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskFileDownloadInfo) GetContentsOk() (*string, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *TaskFileDownloadInfo) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *TaskFileDownloadInfo) SetContents(v string) {
	o.Contents = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *TaskFileDownloadInfo) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskFileDownloadInfo) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *TaskFileDownloadInfo) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *TaskFileDownloadInfo) SetPath(v string) {
	o.Path = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TaskFileDownloadInfo) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskFileDownloadInfo) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TaskFileDownloadInfo) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *TaskFileDownloadInfo) SetSource(v string) {
	o.Source = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaskFileDownloadInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskFileDownloadInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaskFileDownloadInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TaskFileDownloadInfo) SetType(v string) {
	o.Type = &v
}

func (o TaskFileDownloadInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskFileDownloadInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Contents) {
		toSerialize["Contents"] = o.Contents
	}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if !IsNil(o.Source) {
		toSerialize["Source"] = o.Source
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaskFileDownloadInfo) UnmarshalJSON(data []byte) (err error) {
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
	type TaskFileDownloadInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// When the type of the download is inline, it will read the file from the contents here.
		Contents *string `json:"Contents,omitempty"`
		// The path of the download from the specified source location for type S3, then this is the object key.
		Path *string `json:"Path,omitempty"`
		// The source of the download location and if type is S3, then this is the bucket name. In case of MoRef download type  the source will have the Moid of the workflow definition.
		Source *string `json:"Source,omitempty"`
		// The type of download location is captured in type. * `S3` - Download workflow from S3. * `Local` - Download workflow from local path. * `Inline` - Workflow is provided inline. * `MoRef` - Start an existing workflow registered with given Moid.
		Type *string `json:"Type,omitempty"`
	}

	varTaskFileDownloadInfoWithoutEmbeddedStruct := TaskFileDownloadInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varTaskFileDownloadInfoWithoutEmbeddedStruct)
	if err == nil {
		varTaskFileDownloadInfo := _TaskFileDownloadInfo{}
		varTaskFileDownloadInfo.ClassId = varTaskFileDownloadInfoWithoutEmbeddedStruct.ClassId
		varTaskFileDownloadInfo.ObjectType = varTaskFileDownloadInfoWithoutEmbeddedStruct.ObjectType
		varTaskFileDownloadInfo.Contents = varTaskFileDownloadInfoWithoutEmbeddedStruct.Contents
		varTaskFileDownloadInfo.Path = varTaskFileDownloadInfoWithoutEmbeddedStruct.Path
		varTaskFileDownloadInfo.Source = varTaskFileDownloadInfoWithoutEmbeddedStruct.Source
		varTaskFileDownloadInfo.Type = varTaskFileDownloadInfoWithoutEmbeddedStruct.Type
		*o = TaskFileDownloadInfo(varTaskFileDownloadInfo)
	} else {
		return err
	}

	varTaskFileDownloadInfo := _TaskFileDownloadInfo{}

	err = json.Unmarshal(data, &varTaskFileDownloadInfo)
	if err == nil {
		o.MoBaseComplexType = varTaskFileDownloadInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Contents")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "Source")
		delete(additionalProperties, "Type")

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

type NullableTaskFileDownloadInfo struct {
	value *TaskFileDownloadInfo
	isSet bool
}

func (v NullableTaskFileDownloadInfo) Get() *TaskFileDownloadInfo {
	return v.value
}

func (v *NullableTaskFileDownloadInfo) Set(val *TaskFileDownloadInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskFileDownloadInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskFileDownloadInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskFileDownloadInfo(val *TaskFileDownloadInfo) *NullableTaskFileDownloadInfo {
	return &NullableTaskFileDownloadInfo{value: val, isSet: true}
}

func (v NullableTaskFileDownloadInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskFileDownloadInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
