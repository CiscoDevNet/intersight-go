/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// StorageKeySetting Models the local key configurarion required for disks encryptions.
type StorageKeySetting struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Method to be used for fetching the encryption key. * `None` - Drive encryption not configured. * `Manual` - Drive encryption using manual key. * `Kmip` - Remote encryption using KMIP.
	KeyType              *string                         `json:"KeyType,omitempty"`
	ManualKey            NullableStorageLocalKeySetting  `json:"ManualKey,omitempty"`
	RemoteKey            NullableStorageRemoteKeySetting `json:"RemoteKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageKeySetting StorageKeySetting

// NewStorageKeySetting instantiates a new StorageKeySetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageKeySetting(classId string, objectType string) *StorageKeySetting {
	this := StorageKeySetting{}
	this.ClassId = classId
	this.ObjectType = objectType
	var keyType string = "None"
	this.KeyType = &keyType
	return &this
}

// NewStorageKeySettingWithDefaults instantiates a new StorageKeySetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageKeySettingWithDefaults() *StorageKeySetting {
	this := StorageKeySetting{}
	var classId string = "storage.KeySetting"
	this.ClassId = classId
	var objectType string = "storage.KeySetting"
	this.ObjectType = objectType
	var keyType string = "None"
	this.KeyType = &keyType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageKeySetting) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageKeySetting) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageKeySetting) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageKeySetting) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageKeySetting) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageKeySetting) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *StorageKeySetting) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageKeySetting) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *StorageKeySetting) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *StorageKeySetting) SetKeyType(v string) {
	o.KeyType = &v
}

// GetManualKey returns the ManualKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageKeySetting) GetManualKey() StorageLocalKeySetting {
	if o == nil || o.ManualKey.Get() == nil {
		var ret StorageLocalKeySetting
		return ret
	}
	return *o.ManualKey.Get()
}

// GetManualKeyOk returns a tuple with the ManualKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageKeySetting) GetManualKeyOk() (*StorageLocalKeySetting, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManualKey.Get(), o.ManualKey.IsSet()
}

// HasManualKey returns a boolean if a field has been set.
func (o *StorageKeySetting) HasManualKey() bool {
	if o != nil && o.ManualKey.IsSet() {
		return true
	}

	return false
}

// SetManualKey gets a reference to the given NullableStorageLocalKeySetting and assigns it to the ManualKey field.
func (o *StorageKeySetting) SetManualKey(v StorageLocalKeySetting) {
	o.ManualKey.Set(&v)
}

// SetManualKeyNil sets the value for ManualKey to be an explicit nil
func (o *StorageKeySetting) SetManualKeyNil() {
	o.ManualKey.Set(nil)
}

// UnsetManualKey ensures that no value is present for ManualKey, not even an explicit nil
func (o *StorageKeySetting) UnsetManualKey() {
	o.ManualKey.Unset()
}

// GetRemoteKey returns the RemoteKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageKeySetting) GetRemoteKey() StorageRemoteKeySetting {
	if o == nil || o.RemoteKey.Get() == nil {
		var ret StorageRemoteKeySetting
		return ret
	}
	return *o.RemoteKey.Get()
}

// GetRemoteKeyOk returns a tuple with the RemoteKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageKeySetting) GetRemoteKeyOk() (*StorageRemoteKeySetting, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteKey.Get(), o.RemoteKey.IsSet()
}

// HasRemoteKey returns a boolean if a field has been set.
func (o *StorageKeySetting) HasRemoteKey() bool {
	if o != nil && o.RemoteKey.IsSet() {
		return true
	}

	return false
}

// SetRemoteKey gets a reference to the given NullableStorageRemoteKeySetting and assigns it to the RemoteKey field.
func (o *StorageKeySetting) SetRemoteKey(v StorageRemoteKeySetting) {
	o.RemoteKey.Set(&v)
}

// SetRemoteKeyNil sets the value for RemoteKey to be an explicit nil
func (o *StorageKeySetting) SetRemoteKeyNil() {
	o.RemoteKey.Set(nil)
}

// UnsetRemoteKey ensures that no value is present for RemoteKey, not even an explicit nil
func (o *StorageKeySetting) UnsetRemoteKey() {
	o.RemoteKey.Unset()
}

func (o StorageKeySetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.KeyType != nil {
		toSerialize["KeyType"] = o.KeyType
	}
	if o.ManualKey.IsSet() {
		toSerialize["ManualKey"] = o.ManualKey.Get()
	}
	if o.RemoteKey.IsSet() {
		toSerialize["RemoteKey"] = o.RemoteKey.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageKeySetting) UnmarshalJSON(bytes []byte) (err error) {
	type StorageKeySettingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Method to be used for fetching the encryption key. * `None` - Drive encryption not configured. * `Manual` - Drive encryption using manual key. * `Kmip` - Remote encryption using KMIP.
		KeyType   *string                         `json:"KeyType,omitempty"`
		ManualKey NullableStorageLocalKeySetting  `json:"ManualKey,omitempty"`
		RemoteKey NullableStorageRemoteKeySetting `json:"RemoteKey,omitempty"`
	}

	varStorageKeySettingWithoutEmbeddedStruct := StorageKeySettingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageKeySettingWithoutEmbeddedStruct)
	if err == nil {
		varStorageKeySetting := _StorageKeySetting{}
		varStorageKeySetting.ClassId = varStorageKeySettingWithoutEmbeddedStruct.ClassId
		varStorageKeySetting.ObjectType = varStorageKeySettingWithoutEmbeddedStruct.ObjectType
		varStorageKeySetting.KeyType = varStorageKeySettingWithoutEmbeddedStruct.KeyType
		varStorageKeySetting.ManualKey = varStorageKeySettingWithoutEmbeddedStruct.ManualKey
		varStorageKeySetting.RemoteKey = varStorageKeySettingWithoutEmbeddedStruct.RemoteKey
		*o = StorageKeySetting(varStorageKeySetting)
	} else {
		return err
	}

	varStorageKeySetting := _StorageKeySetting{}

	err = json.Unmarshal(bytes, &varStorageKeySetting)
	if err == nil {
		o.MoBaseComplexType = varStorageKeySetting.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KeyType")
		delete(additionalProperties, "ManualKey")
		delete(additionalProperties, "RemoteKey")

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

type NullableStorageKeySetting struct {
	value *StorageKeySetting
	isSet bool
}

func (v NullableStorageKeySetting) Get() *StorageKeySetting {
	return v.value
}

func (v *NullableStorageKeySetting) Set(val *StorageKeySetting) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageKeySetting) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageKeySetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageKeySetting(val *StorageKeySetting) *NullableStorageKeySetting {
	return &NullableStorageKeySetting{value: val, isSet: true}
}

func (v NullableStorageKeySetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageKeySetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}