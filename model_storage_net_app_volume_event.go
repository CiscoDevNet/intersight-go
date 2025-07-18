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

// checks if the StorageNetAppVolumeEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppVolumeEvent{}

// StorageNetAppVolumeEvent An event where the impacted resource type is a volume.
type StorageNetAppVolumeEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                  `json:"ObjectType"`
	Volume               NullableStorageNetAppVolumeRelationship `json:"Volume,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppVolumeEvent StorageNetAppVolumeEvent

// NewStorageNetAppVolumeEvent instantiates a new StorageNetAppVolumeEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppVolumeEvent(classId string, objectType string) *StorageNetAppVolumeEvent {
	this := StorageNetAppVolumeEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppVolumeEventWithDefaults instantiates a new StorageNetAppVolumeEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppVolumeEventWithDefaults() *StorageNetAppVolumeEvent {
	this := StorageNetAppVolumeEvent{}
	var classId string = "storage.NetAppVolumeEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppVolumeEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppVolumeEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppVolumeEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppVolumeEvent" of the ClassId field.
func (o *StorageNetAppVolumeEvent) GetDefaultClassId() interface{} {
	return "storage.NetAppVolumeEvent"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppVolumeEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppVolumeEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppVolumeEvent" of the ObjectType field.
func (o *StorageNetAppVolumeEvent) GetDefaultObjectType() interface{} {
	return "storage.NetAppVolumeEvent"
}

// GetVolume returns the Volume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppVolumeEvent) GetVolume() StorageNetAppVolumeRelationship {
	if o == nil || IsNil(o.Volume.Get()) {
		var ret StorageNetAppVolumeRelationship
		return ret
	}
	return *o.Volume.Get()
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppVolumeEvent) GetVolumeOk() (*StorageNetAppVolumeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Volume.Get(), o.Volume.IsSet()
}

// HasVolume returns a boolean if a field has been set.
func (o *StorageNetAppVolumeEvent) HasVolume() bool {
	if o != nil && o.Volume.IsSet() {
		return true
	}

	return false
}

// SetVolume gets a reference to the given NullableStorageNetAppVolumeRelationship and assigns it to the Volume field.
func (o *StorageNetAppVolumeEvent) SetVolume(v StorageNetAppVolumeRelationship) {
	o.Volume.Set(&v)
}

// SetVolumeNil sets the value for Volume to be an explicit nil
func (o *StorageNetAppVolumeEvent) SetVolumeNil() {
	o.Volume.Set(nil)
}

// UnsetVolume ensures that no value is present for Volume, not even an explicit nil
func (o *StorageNetAppVolumeEvent) UnsetVolume() {
	o.Volume.Unset()
}

func (o StorageNetAppVolumeEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppVolumeEvent) ToMap() (map[string]interface{}, error) {
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
	if o.Volume.IsSet() {
		toSerialize["Volume"] = o.Volume.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppVolumeEvent) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppVolumeEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                  `json:"ObjectType"`
		Volume     NullableStorageNetAppVolumeRelationship `json:"Volume,omitempty"`
	}

	varStorageNetAppVolumeEventWithoutEmbeddedStruct := StorageNetAppVolumeEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppVolumeEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppVolumeEvent := _StorageNetAppVolumeEvent{}
		varStorageNetAppVolumeEvent.ClassId = varStorageNetAppVolumeEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppVolumeEvent.ObjectType = varStorageNetAppVolumeEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppVolumeEvent.Volume = varStorageNetAppVolumeEventWithoutEmbeddedStruct.Volume
		*o = StorageNetAppVolumeEvent(varStorageNetAppVolumeEvent)
	} else {
		return err
	}

	varStorageNetAppVolumeEvent := _StorageNetAppVolumeEvent{}

	err = json.Unmarshal(data, &varStorageNetAppVolumeEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppVolumeEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Volume")

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

type NullableStorageNetAppVolumeEvent struct {
	value *StorageNetAppVolumeEvent
	isSet bool
}

func (v NullableStorageNetAppVolumeEvent) Get() *StorageNetAppVolumeEvent {
	return v.value
}

func (v *NullableStorageNetAppVolumeEvent) Set(val *StorageNetAppVolumeEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppVolumeEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppVolumeEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppVolumeEvent(val *StorageNetAppVolumeEvent) *NullableStorageNetAppVolumeEvent {
	return &NullableStorageNetAppVolumeEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppVolumeEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppVolumeEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
