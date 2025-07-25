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

// checks if the TaskNetAppScopedInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskNetAppScopedInventory{}

// TaskNetAppScopedInventory API to trigger on-demand NetAppStorage inventory to update modified objects in Intersight report.
type TaskNetAppScopedInventory struct {
	ConnectorScopedInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                      `json:"ObjectType"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaskNetAppScopedInventory TaskNetAppScopedInventory

// NewTaskNetAppScopedInventory instantiates a new TaskNetAppScopedInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskNetAppScopedInventory(classId string, objectType string) *TaskNetAppScopedInventory {
	this := TaskNetAppScopedInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTaskNetAppScopedInventoryWithDefaults instantiates a new TaskNetAppScopedInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskNetAppScopedInventoryWithDefaults() *TaskNetAppScopedInventory {
	this := TaskNetAppScopedInventory{}
	var classId string = "task.NetAppScopedInventory"
	this.ClassId = classId
	var objectType string = "task.NetAppScopedInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TaskNetAppScopedInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TaskNetAppScopedInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TaskNetAppScopedInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "task.NetAppScopedInventory" of the ClassId field.
func (o *TaskNetAppScopedInventory) GetDefaultClassId() interface{} {
	return "task.NetAppScopedInventory"
}

// GetObjectType returns the ObjectType field value
func (o *TaskNetAppScopedInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TaskNetAppScopedInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TaskNetAppScopedInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "task.NetAppScopedInventory" of the ObjectType field.
func (o *TaskNetAppScopedInventory) GetDefaultObjectType() interface{} {
	return "task.NetAppScopedInventory"
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskNetAppScopedInventory) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskNetAppScopedInventory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *TaskNetAppScopedInventory) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *TaskNetAppScopedInventory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *TaskNetAppScopedInventory) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *TaskNetAppScopedInventory) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o TaskNetAppScopedInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskNetAppScopedInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorScopedInventory, errConnectorScopedInventory := json.Marshal(o.ConnectorScopedInventory)
	if errConnectorScopedInventory != nil {
		return map[string]interface{}{}, errConnectorScopedInventory
	}
	errConnectorScopedInventory = json.Unmarshal([]byte(serializedConnectorScopedInventory), &toSerialize)
	if errConnectorScopedInventory != nil {
		return map[string]interface{}{}, errConnectorScopedInventory
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaskNetAppScopedInventory) UnmarshalJSON(data []byte) (err error) {
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
	type TaskNetAppScopedInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string                                      `json:"ObjectType"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varTaskNetAppScopedInventoryWithoutEmbeddedStruct := TaskNetAppScopedInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varTaskNetAppScopedInventoryWithoutEmbeddedStruct)
	if err == nil {
		varTaskNetAppScopedInventory := _TaskNetAppScopedInventory{}
		varTaskNetAppScopedInventory.ClassId = varTaskNetAppScopedInventoryWithoutEmbeddedStruct.ClassId
		varTaskNetAppScopedInventory.ObjectType = varTaskNetAppScopedInventoryWithoutEmbeddedStruct.ObjectType
		varTaskNetAppScopedInventory.RegisteredDevice = varTaskNetAppScopedInventoryWithoutEmbeddedStruct.RegisteredDevice
		*o = TaskNetAppScopedInventory(varTaskNetAppScopedInventory)
	} else {
		return err
	}

	varTaskNetAppScopedInventory := _TaskNetAppScopedInventory{}

	err = json.Unmarshal(data, &varTaskNetAppScopedInventory)
	if err == nil {
		o.ConnectorScopedInventory = varTaskNetAppScopedInventory.ConnectorScopedInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectConnectorScopedInventory := reflect.ValueOf(o.ConnectorScopedInventory)
		for i := 0; i < reflectConnectorScopedInventory.Type().NumField(); i++ {
			t := reflectConnectorScopedInventory.Type().Field(i)

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

type NullableTaskNetAppScopedInventory struct {
	value *TaskNetAppScopedInventory
	isSet bool
}

func (v NullableTaskNetAppScopedInventory) Get() *TaskNetAppScopedInventory {
	return v.value
}

func (v *NullableTaskNetAppScopedInventory) Set(val *TaskNetAppScopedInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskNetAppScopedInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskNetAppScopedInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskNetAppScopedInventory(val *TaskNetAppScopedInventory) *NullableTaskNetAppScopedInventory {
	return &NullableTaskNetAppScopedInventory{value: val, isSet: true}
}

func (v NullableTaskNetAppScopedInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskNetAppScopedInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
