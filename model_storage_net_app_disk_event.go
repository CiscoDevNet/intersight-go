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

// checks if the StorageNetAppDiskEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppDiskEvent{}

// StorageNetAppDiskEvent An event where the impacted resource type is a disk.
type StorageNetAppDiskEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                    `json:"ObjectType"`
	Disk                 NullableStorageNetAppBaseDiskRelationship `json:"Disk,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppDiskEvent StorageNetAppDiskEvent

// NewStorageNetAppDiskEvent instantiates a new StorageNetAppDiskEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppDiskEvent(classId string, objectType string) *StorageNetAppDiskEvent {
	this := StorageNetAppDiskEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppDiskEventWithDefaults instantiates a new StorageNetAppDiskEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppDiskEventWithDefaults() *StorageNetAppDiskEvent {
	this := StorageNetAppDiskEvent{}
	var classId string = "storage.NetAppDiskEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppDiskEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppDiskEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppDiskEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppDiskEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppDiskEvent" of the ClassId field.
func (o *StorageNetAppDiskEvent) GetDefaultClassId() interface{} {
	return "storage.NetAppDiskEvent"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppDiskEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppDiskEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppDiskEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppDiskEvent" of the ObjectType field.
func (o *StorageNetAppDiskEvent) GetDefaultObjectType() interface{} {
	return "storage.NetAppDiskEvent"
}

// GetDisk returns the Disk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppDiskEvent) GetDisk() StorageNetAppBaseDiskRelationship {
	if o == nil || IsNil(o.Disk.Get()) {
		var ret StorageNetAppBaseDiskRelationship
		return ret
	}
	return *o.Disk.Get()
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppDiskEvent) GetDiskOk() (*StorageNetAppBaseDiskRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disk.Get(), o.Disk.IsSet()
}

// HasDisk returns a boolean if a field has been set.
func (o *StorageNetAppDiskEvent) HasDisk() bool {
	if o != nil && o.Disk.IsSet() {
		return true
	}

	return false
}

// SetDisk gets a reference to the given NullableStorageNetAppBaseDiskRelationship and assigns it to the Disk field.
func (o *StorageNetAppDiskEvent) SetDisk(v StorageNetAppBaseDiskRelationship) {
	o.Disk.Set(&v)
}

// SetDiskNil sets the value for Disk to be an explicit nil
func (o *StorageNetAppDiskEvent) SetDiskNil() {
	o.Disk.Set(nil)
}

// UnsetDisk ensures that no value is present for Disk, not even an explicit nil
func (o *StorageNetAppDiskEvent) UnsetDisk() {
	o.Disk.Unset()
}

func (o StorageNetAppDiskEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppDiskEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageNetAppBaseEvent, errStorageNetAppBaseEvent := json.Marshal(o.StorageNetAppBaseEvent)
	if errStorageNetAppBaseEvent != nil {
		return map[string]interface{}{}, errStorageNetAppBaseEvent
	}
	errStorageNetAppBaseEvent = json.Unmarshal([]byte(serializedStorageNetAppBaseEvent), &toSerialize)
	if errStorageNetAppBaseEvent != nil {
		return map[string]interface{}{}, errStorageNetAppBaseEvent
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Disk.IsSet() {
		toSerialize["Disk"] = o.Disk.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppDiskEvent) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppDiskEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                    `json:"ObjectType"`
		Disk       NullableStorageNetAppBaseDiskRelationship `json:"Disk,omitempty"`
	}

	varStorageNetAppDiskEventWithoutEmbeddedStruct := StorageNetAppDiskEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppDiskEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppDiskEvent := _StorageNetAppDiskEvent{}
		varStorageNetAppDiskEvent.ClassId = varStorageNetAppDiskEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppDiskEvent.ObjectType = varStorageNetAppDiskEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppDiskEvent.Disk = varStorageNetAppDiskEventWithoutEmbeddedStruct.Disk
		*o = StorageNetAppDiskEvent(varStorageNetAppDiskEvent)
	} else {
		return err
	}

	varStorageNetAppDiskEvent := _StorageNetAppDiskEvent{}

	err = json.Unmarshal(data, &varStorageNetAppDiskEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppDiskEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Disk")

		// remove fields from embedded structs
		reflectStorageNetAppBaseEvent := reflect.ValueOf(o.StorageNetAppBaseEvent)
		for i := 0; i < reflectStorageNetAppBaseEvent.Type().NumField(); i++ {
			t := reflectStorageNetAppBaseEvent.Type().Field(i)

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

type NullableStorageNetAppDiskEvent struct {
	value *StorageNetAppDiskEvent
	isSet bool
}

func (v NullableStorageNetAppDiskEvent) Get() *StorageNetAppDiskEvent {
	return v.value
}

func (v *NullableStorageNetAppDiskEvent) Set(val *StorageNetAppDiskEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppDiskEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppDiskEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppDiskEvent(val *StorageNetAppDiskEvent) *NullableStorageNetAppDiskEvent {
	return &NullableStorageNetAppDiskEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppDiskEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppDiskEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
