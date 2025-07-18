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

// checks if the EquipmentBaseSensor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentBaseSensor{}

// EquipmentBaseSensor Information for a particular sensor on a storage array controller.
type EquipmentBaseSensor struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The name of the storage array controller that the sensor applies to.
	ControllerName *string `json:"ControllerName,omitempty"`
	// The name of a specific sensor.
	Name *string `json:"Name,omitempty"`
	// The state of the specified sensor.
	State *string `json:"State,omitempty"`
	// The type of the specified sensor.
	Type *string `json:"Type,omitempty"`
	// The units that correspond to the value of the sensor, if applicable.
	Units *string `json:"Units,omitempty"`
	// The value of the specified sensor.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentBaseSensor EquipmentBaseSensor

// NewEquipmentBaseSensor instantiates a new EquipmentBaseSensor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentBaseSensor(classId string, objectType string) *EquipmentBaseSensor {
	this := EquipmentBaseSensor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentBaseSensorWithDefaults instantiates a new EquipmentBaseSensor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentBaseSensorWithDefaults() *EquipmentBaseSensor {
	this := EquipmentBaseSensor{}
	var classId string = "storage.NetAppSensor"
	this.ClassId = classId
	var objectType string = "storage.NetAppSensor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentBaseSensor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentBaseSensor) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppSensor" of the ClassId field.
func (o *EquipmentBaseSensor) GetDefaultClassId() interface{} {
	return "storage.NetAppSensor"
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentBaseSensor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentBaseSensor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppSensor" of the ObjectType field.
func (o *EquipmentBaseSensor) GetDefaultObjectType() interface{} {
	return "storage.NetAppSensor"
}

// GetControllerName returns the ControllerName field value if set, zero value otherwise.
func (o *EquipmentBaseSensor) GetControllerName() string {
	if o == nil || IsNil(o.ControllerName) {
		var ret string
		return ret
	}
	return *o.ControllerName
}

// GetControllerNameOk returns a tuple with the ControllerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensor) GetControllerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerName) {
		return nil, false
	}
	return o.ControllerName, true
}

// HasControllerName returns a boolean if a field has been set.
func (o *EquipmentBaseSensor) HasControllerName() bool {
	if o != nil && !IsNil(o.ControllerName) {
		return true
	}

	return false
}

// SetControllerName gets a reference to the given string and assigns it to the ControllerName field.
func (o *EquipmentBaseSensor) SetControllerName(v string) {
	o.ControllerName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EquipmentBaseSensor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EquipmentBaseSensor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EquipmentBaseSensor) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *EquipmentBaseSensor) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensor) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *EquipmentBaseSensor) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *EquipmentBaseSensor) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EquipmentBaseSensor) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensor) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EquipmentBaseSensor) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EquipmentBaseSensor) SetType(v string) {
	o.Type = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *EquipmentBaseSensor) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensor) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *EquipmentBaseSensor) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *EquipmentBaseSensor) SetUnits(v string) {
	o.Units = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EquipmentBaseSensor) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBaseSensor) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EquipmentBaseSensor) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EquipmentBaseSensor) SetValue(v string) {
	o.Value = &v
}

func (o EquipmentBaseSensor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentBaseSensor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ControllerName) {
		toSerialize["ControllerName"] = o.ControllerName
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Units) {
		toSerialize["Units"] = o.Units
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentBaseSensor) UnmarshalJSON(data []byte) (err error) {
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
	type EquipmentBaseSensorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The name of the storage array controller that the sensor applies to.
		ControllerName *string `json:"ControllerName,omitempty"`
		// The name of a specific sensor.
		Name *string `json:"Name,omitempty"`
		// The state of the specified sensor.
		State *string `json:"State,omitempty"`
		// The type of the specified sensor.
		Type *string `json:"Type,omitempty"`
		// The units that correspond to the value of the sensor, if applicable.
		Units *string `json:"Units,omitempty"`
		// The value of the specified sensor.
		Value *string `json:"Value,omitempty"`
	}

	varEquipmentBaseSensorWithoutEmbeddedStruct := EquipmentBaseSensorWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentBaseSensorWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentBaseSensor := _EquipmentBaseSensor{}
		varEquipmentBaseSensor.ClassId = varEquipmentBaseSensorWithoutEmbeddedStruct.ClassId
		varEquipmentBaseSensor.ObjectType = varEquipmentBaseSensorWithoutEmbeddedStruct.ObjectType
		varEquipmentBaseSensor.ControllerName = varEquipmentBaseSensorWithoutEmbeddedStruct.ControllerName
		varEquipmentBaseSensor.Name = varEquipmentBaseSensorWithoutEmbeddedStruct.Name
		varEquipmentBaseSensor.State = varEquipmentBaseSensorWithoutEmbeddedStruct.State
		varEquipmentBaseSensor.Type = varEquipmentBaseSensorWithoutEmbeddedStruct.Type
		varEquipmentBaseSensor.Units = varEquipmentBaseSensorWithoutEmbeddedStruct.Units
		varEquipmentBaseSensor.Value = varEquipmentBaseSensorWithoutEmbeddedStruct.Value
		*o = EquipmentBaseSensor(varEquipmentBaseSensor)
	} else {
		return err
	}

	varEquipmentBaseSensor := _EquipmentBaseSensor{}

	err = json.Unmarshal(data, &varEquipmentBaseSensor)
	if err == nil {
		o.MoBaseMo = varEquipmentBaseSensor.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ControllerName")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Units")
		delete(additionalProperties, "Value")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableEquipmentBaseSensor struct {
	value *EquipmentBaseSensor
	isSet bool
}

func (v NullableEquipmentBaseSensor) Get() *EquipmentBaseSensor {
	return v.value
}

func (v *NullableEquipmentBaseSensor) Set(val *EquipmentBaseSensor) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentBaseSensor) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentBaseSensor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentBaseSensor(val *EquipmentBaseSensor) *NullableEquipmentBaseSensor {
	return &NullableEquipmentBaseSensor{value: val, isSet: true}
}

func (v NullableEquipmentBaseSensor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentBaseSensor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
