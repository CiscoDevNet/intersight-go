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

// checks if the OsPhysicalDiskResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsPhysicalDiskResponse{}

// OsPhysicalDiskResponse Physical Disk target entry for the UI.
type OsPhysicalDiskResponse struct {
	OsInstallTargetResponse
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Bootable field of the Physical Drive target.
	Bootable *string `json:"Bootable,omitempty"`
	// The Physical Disk Name to be used as Install Target.
	Name *string `json:"Name,omitempty"`
	// Serial Number of the Physical Disk Target.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// The Storage Controller associated to the physical disk.
	StorageControllerSlotId *string `json:"StorageControllerSlotId,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _OsPhysicalDiskResponse OsPhysicalDiskResponse

// NewOsPhysicalDiskResponse instantiates a new OsPhysicalDiskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsPhysicalDiskResponse(classId string, objectType string) *OsPhysicalDiskResponse {
	this := OsPhysicalDiskResponse{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsPhysicalDiskResponseWithDefaults instantiates a new OsPhysicalDiskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsPhysicalDiskResponseWithDefaults() *OsPhysicalDiskResponse {
	this := OsPhysicalDiskResponse{}
	var classId string = "os.PhysicalDiskResponse"
	this.ClassId = classId
	var objectType string = "os.PhysicalDiskResponse"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsPhysicalDiskResponse) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsPhysicalDiskResponse) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsPhysicalDiskResponse) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "os.PhysicalDiskResponse" of the ClassId field.
func (o *OsPhysicalDiskResponse) GetDefaultClassId() interface{} {
	return "os.PhysicalDiskResponse"
}

// GetObjectType returns the ObjectType field value
func (o *OsPhysicalDiskResponse) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsPhysicalDiskResponse) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsPhysicalDiskResponse) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "os.PhysicalDiskResponse" of the ObjectType field.
func (o *OsPhysicalDiskResponse) GetDefaultObjectType() interface{} {
	return "os.PhysicalDiskResponse"
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *OsPhysicalDiskResponse) GetBootable() string {
	if o == nil || IsNil(o.Bootable) {
		var ret string
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPhysicalDiskResponse) GetBootableOk() (*string, bool) {
	if o == nil || IsNil(o.Bootable) {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *OsPhysicalDiskResponse) HasBootable() bool {
	if o != nil && !IsNil(o.Bootable) {
		return true
	}

	return false
}

// SetBootable gets a reference to the given string and assigns it to the Bootable field.
func (o *OsPhysicalDiskResponse) SetBootable(v string) {
	o.Bootable = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsPhysicalDiskResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPhysicalDiskResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsPhysicalDiskResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsPhysicalDiskResponse) SetName(v string) {
	o.Name = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *OsPhysicalDiskResponse) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPhysicalDiskResponse) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *OsPhysicalDiskResponse) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *OsPhysicalDiskResponse) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetStorageControllerSlotId returns the StorageControllerSlotId field value if set, zero value otherwise.
func (o *OsPhysicalDiskResponse) GetStorageControllerSlotId() string {
	if o == nil || IsNil(o.StorageControllerSlotId) {
		var ret string
		return ret
	}
	return *o.StorageControllerSlotId
}

// GetStorageControllerSlotIdOk returns a tuple with the StorageControllerSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPhysicalDiskResponse) GetStorageControllerSlotIdOk() (*string, bool) {
	if o == nil || IsNil(o.StorageControllerSlotId) {
		return nil, false
	}
	return o.StorageControllerSlotId, true
}

// HasStorageControllerSlotId returns a boolean if a field has been set.
func (o *OsPhysicalDiskResponse) HasStorageControllerSlotId() bool {
	if o != nil && !IsNil(o.StorageControllerSlotId) {
		return true
	}

	return false
}

// SetStorageControllerSlotId gets a reference to the given string and assigns it to the StorageControllerSlotId field.
func (o *OsPhysicalDiskResponse) SetStorageControllerSlotId(v string) {
	o.StorageControllerSlotId = &v
}

func (o OsPhysicalDiskResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsPhysicalDiskResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedOsInstallTargetResponse, errOsInstallTargetResponse := json.Marshal(o.OsInstallTargetResponse)
	if errOsInstallTargetResponse != nil {
		return map[string]interface{}{}, errOsInstallTargetResponse
	}
	errOsInstallTargetResponse = json.Unmarshal([]byte(serializedOsInstallTargetResponse), &toSerialize)
	if errOsInstallTargetResponse != nil {
		return map[string]interface{}{}, errOsInstallTargetResponse
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Bootable) {
		toSerialize["Bootable"] = o.Bootable
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if !IsNil(o.StorageControllerSlotId) {
		toSerialize["StorageControllerSlotId"] = o.StorageControllerSlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OsPhysicalDiskResponse) UnmarshalJSON(data []byte) (err error) {
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
	type OsPhysicalDiskResponseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Bootable field of the Physical Drive target.
		Bootable *string `json:"Bootable,omitempty"`
		// The Physical Disk Name to be used as Install Target.
		Name *string `json:"Name,omitempty"`
		// Serial Number of the Physical Disk Target.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// The Storage Controller associated to the physical disk.
		StorageControllerSlotId *string `json:"StorageControllerSlotId,omitempty"`
	}

	varOsPhysicalDiskResponseWithoutEmbeddedStruct := OsPhysicalDiskResponseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOsPhysicalDiskResponseWithoutEmbeddedStruct)
	if err == nil {
		varOsPhysicalDiskResponse := _OsPhysicalDiskResponse{}
		varOsPhysicalDiskResponse.ClassId = varOsPhysicalDiskResponseWithoutEmbeddedStruct.ClassId
		varOsPhysicalDiskResponse.ObjectType = varOsPhysicalDiskResponseWithoutEmbeddedStruct.ObjectType
		varOsPhysicalDiskResponse.Bootable = varOsPhysicalDiskResponseWithoutEmbeddedStruct.Bootable
		varOsPhysicalDiskResponse.Name = varOsPhysicalDiskResponseWithoutEmbeddedStruct.Name
		varOsPhysicalDiskResponse.SerialNumber = varOsPhysicalDiskResponseWithoutEmbeddedStruct.SerialNumber
		varOsPhysicalDiskResponse.StorageControllerSlotId = varOsPhysicalDiskResponseWithoutEmbeddedStruct.StorageControllerSlotId
		*o = OsPhysicalDiskResponse(varOsPhysicalDiskResponse)
	} else {
		return err
	}

	varOsPhysicalDiskResponse := _OsPhysicalDiskResponse{}

	err = json.Unmarshal(data, &varOsPhysicalDiskResponse)
	if err == nil {
		o.OsInstallTargetResponse = varOsPhysicalDiskResponse.OsInstallTargetResponse
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootable")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "StorageControllerSlotId")

		// remove fields from embedded structs
		reflectOsInstallTargetResponse := reflect.ValueOf(o.OsInstallTargetResponse)
		for i := 0; i < reflectOsInstallTargetResponse.Type().NumField(); i++ {
			t := reflectOsInstallTargetResponse.Type().Field(i)

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

type NullableOsPhysicalDiskResponse struct {
	value *OsPhysicalDiskResponse
	isSet bool
}

func (v NullableOsPhysicalDiskResponse) Get() *OsPhysicalDiskResponse {
	return v.value
}

func (v *NullableOsPhysicalDiskResponse) Set(val *OsPhysicalDiskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOsPhysicalDiskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOsPhysicalDiskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsPhysicalDiskResponse(val *OsPhysicalDiskResponse) *NullableOsPhysicalDiskResponse {
	return &NullableOsPhysicalDiskResponse{value: val, isSet: true}
}

func (v NullableOsPhysicalDiskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsPhysicalDiskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
