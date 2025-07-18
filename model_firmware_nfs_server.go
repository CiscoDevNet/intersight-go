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

// checks if the FirmwareNfsServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareNfsServer{}

// FirmwareNfsServer Network share file server.
type FirmwareNfsServer struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile.
	FileLocation *string `json:"FileLocation,omitempty"`
	// Mount option as configured on the NFS Server. For example:nolock.
	MountOptions *string `json:"MountOptions,omitempty"`
	// Filename of the image in the remote share location. For example:ucs-c220m5-huu-3.1.2c.iso.
	RemoteFile *string `json:"RemoteFile,omitempty"`
	// NFS Server Hostname or IP Address. For example:NFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade.
	RemoteIp *string `json:"RemoteIp,omitempty"`
	// Directory where the image is stored. For example:/share/subfolder.
	RemoteShare          *string `json:"RemoteShare,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareNfsServer FirmwareNfsServer

// NewFirmwareNfsServer instantiates a new FirmwareNfsServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareNfsServer(classId string, objectType string) *FirmwareNfsServer {
	this := FirmwareNfsServer{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareNfsServerWithDefaults instantiates a new FirmwareNfsServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareNfsServerWithDefaults() *FirmwareNfsServer {
	this := FirmwareNfsServer{}
	var classId string = "firmware.NfsServer"
	this.ClassId = classId
	var objectType string = "firmware.NfsServer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareNfsServer) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareNfsServer) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "firmware.NfsServer" of the ClassId field.
func (o *FirmwareNfsServer) GetDefaultClassId() interface{} {
	return "firmware.NfsServer"
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareNfsServer) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareNfsServer) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "firmware.NfsServer" of the ObjectType field.
func (o *FirmwareNfsServer) GetDefaultObjectType() interface{} {
	return "firmware.NfsServer"
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *FirmwareNfsServer) GetFileLocation() string {
	if o == nil || IsNil(o.FileLocation) {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetFileLocationOk() (*string, bool) {
	if o == nil || IsNil(o.FileLocation) {
		return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *FirmwareNfsServer) HasFileLocation() bool {
	if o != nil && !IsNil(o.FileLocation) {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *FirmwareNfsServer) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetMountOptions returns the MountOptions field value if set, zero value otherwise.
func (o *FirmwareNfsServer) GetMountOptions() string {
	if o == nil || IsNil(o.MountOptions) {
		var ret string
		return ret
	}
	return *o.MountOptions
}

// GetMountOptionsOk returns a tuple with the MountOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetMountOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.MountOptions) {
		return nil, false
	}
	return o.MountOptions, true
}

// HasMountOptions returns a boolean if a field has been set.
func (o *FirmwareNfsServer) HasMountOptions() bool {
	if o != nil && !IsNil(o.MountOptions) {
		return true
	}

	return false
}

// SetMountOptions gets a reference to the given string and assigns it to the MountOptions field.
func (o *FirmwareNfsServer) SetMountOptions(v string) {
	o.MountOptions = &v
}

// GetRemoteFile returns the RemoteFile field value if set, zero value otherwise.
func (o *FirmwareNfsServer) GetRemoteFile() string {
	if o == nil || IsNil(o.RemoteFile) {
		var ret string
		return ret
	}
	return *o.RemoteFile
}

// GetRemoteFileOk returns a tuple with the RemoteFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetRemoteFileOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteFile) {
		return nil, false
	}
	return o.RemoteFile, true
}

// HasRemoteFile returns a boolean if a field has been set.
func (o *FirmwareNfsServer) HasRemoteFile() bool {
	if o != nil && !IsNil(o.RemoteFile) {
		return true
	}

	return false
}

// SetRemoteFile gets a reference to the given string and assigns it to the RemoteFile field.
func (o *FirmwareNfsServer) SetRemoteFile(v string) {
	o.RemoteFile = &v
}

// GetRemoteIp returns the RemoteIp field value if set, zero value otherwise.
func (o *FirmwareNfsServer) GetRemoteIp() string {
	if o == nil || IsNil(o.RemoteIp) {
		var ret string
		return ret
	}
	return *o.RemoteIp
}

// GetRemoteIpOk returns a tuple with the RemoteIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetRemoteIpOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteIp) {
		return nil, false
	}
	return o.RemoteIp, true
}

// HasRemoteIp returns a boolean if a field has been set.
func (o *FirmwareNfsServer) HasRemoteIp() bool {
	if o != nil && !IsNil(o.RemoteIp) {
		return true
	}

	return false
}

// SetRemoteIp gets a reference to the given string and assigns it to the RemoteIp field.
func (o *FirmwareNfsServer) SetRemoteIp(v string) {
	o.RemoteIp = &v
}

// GetRemoteShare returns the RemoteShare field value if set, zero value otherwise.
func (o *FirmwareNfsServer) GetRemoteShare() string {
	if o == nil || IsNil(o.RemoteShare) {
		var ret string
		return ret
	}
	return *o.RemoteShare
}

// GetRemoteShareOk returns a tuple with the RemoteShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetRemoteShareOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteShare) {
		return nil, false
	}
	return o.RemoteShare, true
}

// HasRemoteShare returns a boolean if a field has been set.
func (o *FirmwareNfsServer) HasRemoteShare() bool {
	if o != nil && !IsNil(o.RemoteShare) {
		return true
	}

	return false
}

// SetRemoteShare gets a reference to the given string and assigns it to the RemoteShare field.
func (o *FirmwareNfsServer) SetRemoteShare(v string) {
	o.RemoteShare = &v
}

func (o FirmwareNfsServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareNfsServer) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.FileLocation) {
		toSerialize["FileLocation"] = o.FileLocation
	}
	if !IsNil(o.MountOptions) {
		toSerialize["MountOptions"] = o.MountOptions
	}
	if !IsNil(o.RemoteFile) {
		toSerialize["RemoteFile"] = o.RemoteFile
	}
	if !IsNil(o.RemoteIp) {
		toSerialize["RemoteIp"] = o.RemoteIp
	}
	if !IsNil(o.RemoteShare) {
		toSerialize["RemoteShare"] = o.RemoteShare
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirmwareNfsServer) UnmarshalJSON(data []byte) (err error) {
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
	type FirmwareNfsServerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile.
		FileLocation *string `json:"FileLocation,omitempty"`
		// Mount option as configured on the NFS Server. For example:nolock.
		MountOptions *string `json:"MountOptions,omitempty"`
		// Filename of the image in the remote share location. For example:ucs-c220m5-huu-3.1.2c.iso.
		RemoteFile *string `json:"RemoteFile,omitempty"`
		// NFS Server Hostname or IP Address. For example:NFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade.
		RemoteIp *string `json:"RemoteIp,omitempty"`
		// Directory where the image is stored. For example:/share/subfolder.
		RemoteShare *string `json:"RemoteShare,omitempty"`
	}

	varFirmwareNfsServerWithoutEmbeddedStruct := FirmwareNfsServerWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFirmwareNfsServerWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareNfsServer := _FirmwareNfsServer{}
		varFirmwareNfsServer.ClassId = varFirmwareNfsServerWithoutEmbeddedStruct.ClassId
		varFirmwareNfsServer.ObjectType = varFirmwareNfsServerWithoutEmbeddedStruct.ObjectType
		varFirmwareNfsServer.FileLocation = varFirmwareNfsServerWithoutEmbeddedStruct.FileLocation
		varFirmwareNfsServer.MountOptions = varFirmwareNfsServerWithoutEmbeddedStruct.MountOptions
		varFirmwareNfsServer.RemoteFile = varFirmwareNfsServerWithoutEmbeddedStruct.RemoteFile
		varFirmwareNfsServer.RemoteIp = varFirmwareNfsServerWithoutEmbeddedStruct.RemoteIp
		varFirmwareNfsServer.RemoteShare = varFirmwareNfsServerWithoutEmbeddedStruct.RemoteShare
		*o = FirmwareNfsServer(varFirmwareNfsServer)
	} else {
		return err
	}

	varFirmwareNfsServer := _FirmwareNfsServer{}

	err = json.Unmarshal(data, &varFirmwareNfsServer)
	if err == nil {
		o.MoBaseComplexType = varFirmwareNfsServer.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileLocation")
		delete(additionalProperties, "MountOptions")
		delete(additionalProperties, "RemoteFile")
		delete(additionalProperties, "RemoteIp")
		delete(additionalProperties, "RemoteShare")

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

type NullableFirmwareNfsServer struct {
	value *FirmwareNfsServer
	isSet bool
}

func (v NullableFirmwareNfsServer) Get() *FirmwareNfsServer {
	return v.value
}

func (v *NullableFirmwareNfsServer) Set(val *FirmwareNfsServer) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareNfsServer) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareNfsServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareNfsServer(val *FirmwareNfsServer) *NullableFirmwareNfsServer {
	return &NullableFirmwareNfsServer{value: val, isSet: true}
}

func (v NullableFirmwareNfsServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareNfsServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
