/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9235
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// StorageNetAppCifsAcl Access Control List entry for the CIFS share.
type StorageNetAppCifsAcl struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Access rights that a user or group has on the defined CIFS share.
	Permission *string `json:"Permission,omitempty"`
	// The type of operating system for the CIFS share.
	Type *string `json:"Type,omitempty"`
	// User or group name to add to the access control list of a CIFS share.
	UserOrGroup          *string `json:"UserOrGroup,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppCifsAcl StorageNetAppCifsAcl

// NewStorageNetAppCifsAcl instantiates a new StorageNetAppCifsAcl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppCifsAcl(classId string, objectType string) *StorageNetAppCifsAcl {
	this := StorageNetAppCifsAcl{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppCifsAclWithDefaults instantiates a new StorageNetAppCifsAcl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppCifsAclWithDefaults() *StorageNetAppCifsAcl {
	this := StorageNetAppCifsAcl{}
	var classId string = "storage.NetAppCifsAcl"
	this.ClassId = classId
	var objectType string = "storage.NetAppCifsAcl"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppCifsAcl) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsAcl) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppCifsAcl) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppCifsAcl) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsAcl) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppCifsAcl) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *StorageNetAppCifsAcl) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsAcl) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *StorageNetAppCifsAcl) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *StorageNetAppCifsAcl) SetPermission(v string) {
	o.Permission = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageNetAppCifsAcl) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsAcl) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageNetAppCifsAcl) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageNetAppCifsAcl) SetType(v string) {
	o.Type = &v
}

// GetUserOrGroup returns the UserOrGroup field value if set, zero value otherwise.
func (o *StorageNetAppCifsAcl) GetUserOrGroup() string {
	if o == nil || o.UserOrGroup == nil {
		var ret string
		return ret
	}
	return *o.UserOrGroup
}

// GetUserOrGroupOk returns a tuple with the UserOrGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsAcl) GetUserOrGroupOk() (*string, bool) {
	if o == nil || o.UserOrGroup == nil {
		return nil, false
	}
	return o.UserOrGroup, true
}

// HasUserOrGroup returns a boolean if a field has been set.
func (o *StorageNetAppCifsAcl) HasUserOrGroup() bool {
	if o != nil && o.UserOrGroup != nil {
		return true
	}

	return false
}

// SetUserOrGroup gets a reference to the given string and assigns it to the UserOrGroup field.
func (o *StorageNetAppCifsAcl) SetUserOrGroup(v string) {
	o.UserOrGroup = &v
}

func (o StorageNetAppCifsAcl) MarshalJSON() ([]byte, error) {
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
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.UserOrGroup != nil {
		toSerialize["UserOrGroup"] = o.UserOrGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppCifsAcl) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppCifsAclWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Access rights that a user or group has on the defined CIFS share.
		Permission *string `json:"Permission,omitempty"`
		// The type of operating system for the CIFS share.
		Type *string `json:"Type,omitempty"`
		// User or group name to add to the access control list of a CIFS share.
		UserOrGroup *string `json:"UserOrGroup,omitempty"`
	}

	varStorageNetAppCifsAclWithoutEmbeddedStruct := StorageNetAppCifsAclWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppCifsAclWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppCifsAcl := _StorageNetAppCifsAcl{}
		varStorageNetAppCifsAcl.ClassId = varStorageNetAppCifsAclWithoutEmbeddedStruct.ClassId
		varStorageNetAppCifsAcl.ObjectType = varStorageNetAppCifsAclWithoutEmbeddedStruct.ObjectType
		varStorageNetAppCifsAcl.Permission = varStorageNetAppCifsAclWithoutEmbeddedStruct.Permission
		varStorageNetAppCifsAcl.Type = varStorageNetAppCifsAclWithoutEmbeddedStruct.Type
		varStorageNetAppCifsAcl.UserOrGroup = varStorageNetAppCifsAclWithoutEmbeddedStruct.UserOrGroup
		*o = StorageNetAppCifsAcl(varStorageNetAppCifsAcl)
	} else {
		return err
	}

	varStorageNetAppCifsAcl := _StorageNetAppCifsAcl{}

	err = json.Unmarshal(bytes, &varStorageNetAppCifsAcl)
	if err == nil {
		o.MoBaseComplexType = varStorageNetAppCifsAcl.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "UserOrGroup")

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

type NullableStorageNetAppCifsAcl struct {
	value *StorageNetAppCifsAcl
	isSet bool
}

func (v NullableStorageNetAppCifsAcl) Get() *StorageNetAppCifsAcl {
	return v.value
}

func (v *NullableStorageNetAppCifsAcl) Set(val *StorageNetAppCifsAcl) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppCifsAcl) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppCifsAcl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppCifsAcl(val *StorageNetAppCifsAcl) *NullableStorageNetAppCifsAcl {
	return &NullableStorageNetAppCifsAcl{value: val, isSet: true}
}

func (v NullableStorageNetAppCifsAcl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppCifsAcl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}