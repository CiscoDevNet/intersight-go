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

// checks if the StorageNetAppAggregateEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppAggregateEvent{}

// StorageNetAppAggregateEvent An event where the impacted resource type is an aggregate.
type StorageNetAppAggregateEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                     `json:"ObjectType"`
	Aggregate            NullableStorageNetAppAggregateRelationship `json:"Aggregate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppAggregateEvent StorageNetAppAggregateEvent

// NewStorageNetAppAggregateEvent instantiates a new StorageNetAppAggregateEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppAggregateEvent(classId string, objectType string) *StorageNetAppAggregateEvent {
	this := StorageNetAppAggregateEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppAggregateEventWithDefaults instantiates a new StorageNetAppAggregateEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppAggregateEventWithDefaults() *StorageNetAppAggregateEvent {
	this := StorageNetAppAggregateEvent{}
	var classId string = "storage.NetAppAggregateEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppAggregateEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppAggregateEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppAggregateEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppAggregateEvent" of the ClassId field.
func (o *StorageNetAppAggregateEvent) GetDefaultClassId() interface{} {
	return "storage.NetAppAggregateEvent"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppAggregateEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppAggregateEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppAggregateEvent" of the ObjectType field.
func (o *StorageNetAppAggregateEvent) GetDefaultObjectType() interface{} {
	return "storage.NetAppAggregateEvent"
}

// GetAggregate returns the Aggregate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppAggregateEvent) GetAggregate() StorageNetAppAggregateRelationship {
	if o == nil || IsNil(o.Aggregate.Get()) {
		var ret StorageNetAppAggregateRelationship
		return ret
	}
	return *o.Aggregate.Get()
}

// GetAggregateOk returns a tuple with the Aggregate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppAggregateEvent) GetAggregateOk() (*StorageNetAppAggregateRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aggregate.Get(), o.Aggregate.IsSet()
}

// HasAggregate returns a boolean if a field has been set.
func (o *StorageNetAppAggregateEvent) HasAggregate() bool {
	if o != nil && o.Aggregate.IsSet() {
		return true
	}

	return false
}

// SetAggregate gets a reference to the given NullableStorageNetAppAggregateRelationship and assigns it to the Aggregate field.
func (o *StorageNetAppAggregateEvent) SetAggregate(v StorageNetAppAggregateRelationship) {
	o.Aggregate.Set(&v)
}

// SetAggregateNil sets the value for Aggregate to be an explicit nil
func (o *StorageNetAppAggregateEvent) SetAggregateNil() {
	o.Aggregate.Set(nil)
}

// UnsetAggregate ensures that no value is present for Aggregate, not even an explicit nil
func (o *StorageNetAppAggregateEvent) UnsetAggregate() {
	o.Aggregate.Unset()
}

func (o StorageNetAppAggregateEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppAggregateEvent) ToMap() (map[string]interface{}, error) {
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
	if o.Aggregate.IsSet() {
		toSerialize["Aggregate"] = o.Aggregate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppAggregateEvent) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppAggregateEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                     `json:"ObjectType"`
		Aggregate  NullableStorageNetAppAggregateRelationship `json:"Aggregate,omitempty"`
	}

	varStorageNetAppAggregateEventWithoutEmbeddedStruct := StorageNetAppAggregateEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppAggregateEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppAggregateEvent := _StorageNetAppAggregateEvent{}
		varStorageNetAppAggregateEvent.ClassId = varStorageNetAppAggregateEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppAggregateEvent.ObjectType = varStorageNetAppAggregateEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppAggregateEvent.Aggregate = varStorageNetAppAggregateEventWithoutEmbeddedStruct.Aggregate
		*o = StorageNetAppAggregateEvent(varStorageNetAppAggregateEvent)
	} else {
		return err
	}

	varStorageNetAppAggregateEvent := _StorageNetAppAggregateEvent{}

	err = json.Unmarshal(data, &varStorageNetAppAggregateEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppAggregateEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Aggregate")

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

type NullableStorageNetAppAggregateEvent struct {
	value *StorageNetAppAggregateEvent
	isSet bool
}

func (v NullableStorageNetAppAggregateEvent) Get() *StorageNetAppAggregateEvent {
	return v.value
}

func (v *NullableStorageNetAppAggregateEvent) Set(val *StorageNetAppAggregateEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppAggregateEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppAggregateEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppAggregateEvent(val *StorageNetAppAggregateEvent) *NullableStorageNetAppAggregateEvent {
	return &NullableStorageNetAppAggregateEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppAggregateEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppAggregateEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
