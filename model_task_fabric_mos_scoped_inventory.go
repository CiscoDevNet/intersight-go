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

// checks if the TaskFabricMosScopedInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskFabricMosScopedInventory{}

// TaskFabricMosScopedInventory API to trigger on-demand inventory to update MDS objects in Intersight.
type TaskFabricMosScopedInventory struct {
	ConnectorScopedInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                      `json:"ObjectType"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaskFabricMosScopedInventory TaskFabricMosScopedInventory

// NewTaskFabricMosScopedInventory instantiates a new TaskFabricMosScopedInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskFabricMosScopedInventory(classId string, objectType string) *TaskFabricMosScopedInventory {
	this := TaskFabricMosScopedInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTaskFabricMosScopedInventoryWithDefaults instantiates a new TaskFabricMosScopedInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskFabricMosScopedInventoryWithDefaults() *TaskFabricMosScopedInventory {
	this := TaskFabricMosScopedInventory{}
	var classId string = "task.FabricMosScopedInventory"
	this.ClassId = classId
	var objectType string = "task.FabricMosScopedInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TaskFabricMosScopedInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TaskFabricMosScopedInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TaskFabricMosScopedInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "task.FabricMosScopedInventory" of the ClassId field.
func (o *TaskFabricMosScopedInventory) GetDefaultClassId() interface{} {
	return "task.FabricMosScopedInventory"
}

// GetObjectType returns the ObjectType field value
func (o *TaskFabricMosScopedInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TaskFabricMosScopedInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TaskFabricMosScopedInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "task.FabricMosScopedInventory" of the ObjectType field.
func (o *TaskFabricMosScopedInventory) GetDefaultObjectType() interface{} {
	return "task.FabricMosScopedInventory"
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskFabricMosScopedInventory) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskFabricMosScopedInventory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *TaskFabricMosScopedInventory) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *TaskFabricMosScopedInventory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *TaskFabricMosScopedInventory) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *TaskFabricMosScopedInventory) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o TaskFabricMosScopedInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskFabricMosScopedInventory) ToMap() (map[string]interface{}, error) {
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

func (o *TaskFabricMosScopedInventory) UnmarshalJSON(data []byte) (err error) {
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
	type TaskFabricMosScopedInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string                                      `json:"ObjectType"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varTaskFabricMosScopedInventoryWithoutEmbeddedStruct := TaskFabricMosScopedInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varTaskFabricMosScopedInventoryWithoutEmbeddedStruct)
	if err == nil {
		varTaskFabricMosScopedInventory := _TaskFabricMosScopedInventory{}
		varTaskFabricMosScopedInventory.ClassId = varTaskFabricMosScopedInventoryWithoutEmbeddedStruct.ClassId
		varTaskFabricMosScopedInventory.ObjectType = varTaskFabricMosScopedInventoryWithoutEmbeddedStruct.ObjectType
		varTaskFabricMosScopedInventory.RegisteredDevice = varTaskFabricMosScopedInventoryWithoutEmbeddedStruct.RegisteredDevice
		*o = TaskFabricMosScopedInventory(varTaskFabricMosScopedInventory)
	} else {
		return err
	}

	varTaskFabricMosScopedInventory := _TaskFabricMosScopedInventory{}

	err = json.Unmarshal(data, &varTaskFabricMosScopedInventory)
	if err == nil {
		o.ConnectorScopedInventory = varTaskFabricMosScopedInventory.ConnectorScopedInventory
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

type NullableTaskFabricMosScopedInventory struct {
	value *TaskFabricMosScopedInventory
	isSet bool
}

func (v NullableTaskFabricMosScopedInventory) Get() *TaskFabricMosScopedInventory {
	return v.value
}

func (v *NullableTaskFabricMosScopedInventory) Set(val *TaskFabricMosScopedInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskFabricMosScopedInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskFabricMosScopedInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskFabricMosScopedInventory(val *TaskFabricMosScopedInventory) *NullableTaskFabricMosScopedInventory {
	return &NullableTaskFabricMosScopedInventory{value: val, isSet: true}
}

func (v NullableTaskFabricMosScopedInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskFabricMosScopedInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
