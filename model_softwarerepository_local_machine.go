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

// checks if the SoftwarerepositoryLocalMachine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwarerepositoryLocalMachine{}

// SoftwarerepositoryLocalMachine The user's local machine that is being used as the source for an image.
type SoftwarerepositoryLocalMachine struct {
	SoftwarerepositoryFileServer
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When the import action in the file MO is updated with 'GeneratePreSignedDownloadUrl', Intersight returns a pre-signed URL in this property as part of the patch response. The user is expected to subsequently download the file using this URL.
	DownloadUrl *string `json:"DownloadUrl,omitempty"`
	// The chunk size (in bytes) for each part of the file to be uploaded.
	PartSize *int64 `json:"PartSize,omitempty"`
	// When the import action in file MO is updated with 'GeneratePreSignedUploadUrl', Intersight shall return a upload Id in this property as part of the PATCH response.
	UploadId *string `json:"UploadId,omitempty"`
	// When a file MO is created with 'LocalMachine' as the source, Intersight returns a pre-signed URL in this property as part of the POST response. The user is expected to subsequently upload the file content using this URL. Once the upload is completed, the user is expected to patch the uploader object's transfer state to success.
	UploadUrl            *string  `json:"UploadUrl,omitempty"`
	UploadUrls           []string `json:"UploadUrls,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryLocalMachine SoftwarerepositoryLocalMachine

// NewSoftwarerepositoryLocalMachine instantiates a new SoftwarerepositoryLocalMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryLocalMachine(classId string, objectType string) *SoftwarerepositoryLocalMachine {
	this := SoftwarerepositoryLocalMachine{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryLocalMachineWithDefaults instantiates a new SoftwarerepositoryLocalMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryLocalMachineWithDefaults() *SoftwarerepositoryLocalMachine {
	this := SoftwarerepositoryLocalMachine{}
	var classId string = "softwarerepository.LocalMachine"
	this.ClassId = classId
	var objectType string = "softwarerepository.LocalMachine"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryLocalMachine) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachine) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryLocalMachine) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "softwarerepository.LocalMachine" of the ClassId field.
func (o *SoftwarerepositoryLocalMachine) GetDefaultClassId() interface{} {
	return "softwarerepository.LocalMachine"
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryLocalMachine) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachine) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryLocalMachine) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "softwarerepository.LocalMachine" of the ObjectType field.
func (o *SoftwarerepositoryLocalMachine) GetDefaultObjectType() interface{} {
	return "softwarerepository.LocalMachine"
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *SoftwarerepositoryLocalMachine) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachine) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachine) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *SoftwarerepositoryLocalMachine) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetPartSize returns the PartSize field value if set, zero value otherwise.
func (o *SoftwarerepositoryLocalMachine) GetPartSize() int64 {
	if o == nil || IsNil(o.PartSize) {
		var ret int64
		return ret
	}
	return *o.PartSize
}

// GetPartSizeOk returns a tuple with the PartSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachine) GetPartSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PartSize) {
		return nil, false
	}
	return o.PartSize, true
}

// HasPartSize returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachine) HasPartSize() bool {
	if o != nil && !IsNil(o.PartSize) {
		return true
	}

	return false
}

// SetPartSize gets a reference to the given int64 and assigns it to the PartSize field.
func (o *SoftwarerepositoryLocalMachine) SetPartSize(v int64) {
	o.PartSize = &v
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *SoftwarerepositoryLocalMachine) GetUploadId() string {
	if o == nil || IsNil(o.UploadId) {
		var ret string
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachine) GetUploadIdOk() (*string, bool) {
	if o == nil || IsNil(o.UploadId) {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachine) HasUploadId() bool {
	if o != nil && !IsNil(o.UploadId) {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given string and assigns it to the UploadId field.
func (o *SoftwarerepositoryLocalMachine) SetUploadId(v string) {
	o.UploadId = &v
}

// GetUploadUrl returns the UploadUrl field value if set, zero value otherwise.
func (o *SoftwarerepositoryLocalMachine) GetUploadUrl() string {
	if o == nil || IsNil(o.UploadUrl) {
		var ret string
		return ret
	}
	return *o.UploadUrl
}

// GetUploadUrlOk returns a tuple with the UploadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachine) GetUploadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.UploadUrl) {
		return nil, false
	}
	return o.UploadUrl, true
}

// HasUploadUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachine) HasUploadUrl() bool {
	if o != nil && !IsNil(o.UploadUrl) {
		return true
	}

	return false
}

// SetUploadUrl gets a reference to the given string and assigns it to the UploadUrl field.
func (o *SoftwarerepositoryLocalMachine) SetUploadUrl(v string) {
	o.UploadUrl = &v
}

// GetUploadUrls returns the UploadUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryLocalMachine) GetUploadUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UploadUrls
}

// GetUploadUrlsOk returns a tuple with the UploadUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryLocalMachine) GetUploadUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.UploadUrls) {
		return nil, false
	}
	return o.UploadUrls, true
}

// HasUploadUrls returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachine) HasUploadUrls() bool {
	if o != nil && !IsNil(o.UploadUrls) {
		return true
	}

	return false
}

// SetUploadUrls gets a reference to the given []string and assigns it to the UploadUrls field.
func (o *SoftwarerepositoryLocalMachine) SetUploadUrls(v []string) {
	o.UploadUrls = v
}

func (o SoftwarerepositoryLocalMachine) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwarerepositoryLocalMachine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSoftwarerepositoryFileServer, errSoftwarerepositoryFileServer := json.Marshal(o.SoftwarerepositoryFileServer)
	if errSoftwarerepositoryFileServer != nil {
		return map[string]interface{}{}, errSoftwarerepositoryFileServer
	}
	errSoftwarerepositoryFileServer = json.Unmarshal([]byte(serializedSoftwarerepositoryFileServer), &toSerialize)
	if errSoftwarerepositoryFileServer != nil {
		return map[string]interface{}{}, errSoftwarerepositoryFileServer
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DownloadUrl) {
		toSerialize["DownloadUrl"] = o.DownloadUrl
	}
	if !IsNil(o.PartSize) {
		toSerialize["PartSize"] = o.PartSize
	}
	if !IsNil(o.UploadId) {
		toSerialize["UploadId"] = o.UploadId
	}
	if !IsNil(o.UploadUrl) {
		toSerialize["UploadUrl"] = o.UploadUrl
	}
	if o.UploadUrls != nil {
		toSerialize["UploadUrls"] = o.UploadUrls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SoftwarerepositoryLocalMachine) UnmarshalJSON(data []byte) (err error) {
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
	type SoftwarerepositoryLocalMachineWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// When the import action in the file MO is updated with 'GeneratePreSignedDownloadUrl', Intersight returns a pre-signed URL in this property as part of the patch response. The user is expected to subsequently download the file using this URL.
		DownloadUrl *string `json:"DownloadUrl,omitempty"`
		// The chunk size (in bytes) for each part of the file to be uploaded.
		PartSize *int64 `json:"PartSize,omitempty"`
		// When the import action in file MO is updated with 'GeneratePreSignedUploadUrl', Intersight shall return a upload Id in this property as part of the PATCH response.
		UploadId *string `json:"UploadId,omitempty"`
		// When a file MO is created with 'LocalMachine' as the source, Intersight returns a pre-signed URL in this property as part of the POST response. The user is expected to subsequently upload the file content using this URL. Once the upload is completed, the user is expected to patch the uploader object's transfer state to success.
		UploadUrl  *string  `json:"UploadUrl,omitempty"`
		UploadUrls []string `json:"UploadUrls,omitempty"`
	}

	varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct := SoftwarerepositoryLocalMachineWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryLocalMachine := _SoftwarerepositoryLocalMachine{}
		varSoftwarerepositoryLocalMachine.ClassId = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.ClassId
		varSoftwarerepositoryLocalMachine.ObjectType = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.ObjectType
		varSoftwarerepositoryLocalMachine.DownloadUrl = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.DownloadUrl
		varSoftwarerepositoryLocalMachine.PartSize = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.PartSize
		varSoftwarerepositoryLocalMachine.UploadId = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.UploadId
		varSoftwarerepositoryLocalMachine.UploadUrl = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.UploadUrl
		varSoftwarerepositoryLocalMachine.UploadUrls = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.UploadUrls
		*o = SoftwarerepositoryLocalMachine(varSoftwarerepositoryLocalMachine)
	} else {
		return err
	}

	varSoftwarerepositoryLocalMachine := _SoftwarerepositoryLocalMachine{}

	err = json.Unmarshal(data, &varSoftwarerepositoryLocalMachine)
	if err == nil {
		o.SoftwarerepositoryFileServer = varSoftwarerepositoryLocalMachine.SoftwarerepositoryFileServer
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DownloadUrl")
		delete(additionalProperties, "PartSize")
		delete(additionalProperties, "UploadId")
		delete(additionalProperties, "UploadUrl")
		delete(additionalProperties, "UploadUrls")

		// remove fields from embedded structs
		reflectSoftwarerepositoryFileServer := reflect.ValueOf(o.SoftwarerepositoryFileServer)
		for i := 0; i < reflectSoftwarerepositoryFileServer.Type().NumField(); i++ {
			t := reflectSoftwarerepositoryFileServer.Type().Field(i)

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

type NullableSoftwarerepositoryLocalMachine struct {
	value *SoftwarerepositoryLocalMachine
	isSet bool
}

func (v NullableSoftwarerepositoryLocalMachine) Get() *SoftwarerepositoryLocalMachine {
	return v.value
}

func (v *NullableSoftwarerepositoryLocalMachine) Set(val *SoftwarerepositoryLocalMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryLocalMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryLocalMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryLocalMachine(val *SoftwarerepositoryLocalMachine) *NullableSoftwarerepositoryLocalMachine {
	return &NullableSoftwarerepositoryLocalMachine{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryLocalMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryLocalMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
