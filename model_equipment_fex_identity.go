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

// checks if the EquipmentFexIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentFexIdentity{}

// EquipmentFexIdentity FexIdentity Object conatains basic information of fabric extender. moduleId is uniquely allocated for the combination of vendor, model and serial number of the chassis.
type EquipmentFexIdentity struct {
	EquipmentPhysicalIdentity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B.
	SwitchId             *int64                             `json:"SwitchId,omitempty"`
	Fex                  NullableEquipmentFexRelationship   `json:"Fex,omitempty"`
	NetworkElement       NullableNetworkElementRelationship `json:"NetworkElement,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentFexIdentity EquipmentFexIdentity

// NewEquipmentFexIdentity instantiates a new EquipmentFexIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentFexIdentity(classId string, objectType string) *EquipmentFexIdentity {
	this := EquipmentFexIdentity{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminAction string = "None"
	this.AdminAction = &adminAction
	return &this
}

// NewEquipmentFexIdentityWithDefaults instantiates a new EquipmentFexIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentFexIdentityWithDefaults() *EquipmentFexIdentity {
	this := EquipmentFexIdentity{}
	var classId string = "equipment.FexIdentity"
	this.ClassId = classId
	var objectType string = "equipment.FexIdentity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentFexIdentity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentFexIdentity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentFexIdentity) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "equipment.FexIdentity" of the ClassId field.
func (o *EquipmentFexIdentity) GetDefaultClassId() interface{} {
	return "equipment.FexIdentity"
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentFexIdentity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentFexIdentity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentFexIdentity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "equipment.FexIdentity" of the ObjectType field.
func (o *EquipmentFexIdentity) GetDefaultObjectType() interface{} {
	return "equipment.FexIdentity"
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *EquipmentFexIdentity) GetSwitchId() int64 {
	if o == nil || IsNil(o.SwitchId) {
		var ret int64
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexIdentity) GetSwitchIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SwitchId) {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *EquipmentFexIdentity) HasSwitchId() bool {
	if o != nil && !IsNil(o.SwitchId) {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given int64 and assigns it to the SwitchId field.
func (o *EquipmentFexIdentity) SetSwitchId(v int64) {
	o.SwitchId = &v
}

// GetFex returns the Fex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentFexIdentity) GetFex() EquipmentFexRelationship {
	if o == nil || IsNil(o.Fex.Get()) {
		var ret EquipmentFexRelationship
		return ret
	}
	return *o.Fex.Get()
}

// GetFexOk returns a tuple with the Fex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentFexIdentity) GetFexOk() (*EquipmentFexRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fex.Get(), o.Fex.IsSet()
}

// HasFex returns a boolean if a field has been set.
func (o *EquipmentFexIdentity) HasFex() bool {
	if o != nil && o.Fex.IsSet() {
		return true
	}

	return false
}

// SetFex gets a reference to the given NullableEquipmentFexRelationship and assigns it to the Fex field.
func (o *EquipmentFexIdentity) SetFex(v EquipmentFexRelationship) {
	o.Fex.Set(&v)
}

// SetFexNil sets the value for Fex to be an explicit nil
func (o *EquipmentFexIdentity) SetFexNil() {
	o.Fex.Set(nil)
}

// UnsetFex ensures that no value is present for Fex, not even an explicit nil
func (o *EquipmentFexIdentity) UnsetFex() {
	o.Fex.Unset()
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentFexIdentity) GetNetworkElement() NetworkElementRelationship {
	if o == nil || IsNil(o.NetworkElement.Get()) {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement.Get()
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentFexIdentity) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkElement.Get(), o.NetworkElement.IsSet()
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *EquipmentFexIdentity) HasNetworkElement() bool {
	if o != nil && o.NetworkElement.IsSet() {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NullableNetworkElementRelationship and assigns it to the NetworkElement field.
func (o *EquipmentFexIdentity) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement.Set(&v)
}

// SetNetworkElementNil sets the value for NetworkElement to be an explicit nil
func (o *EquipmentFexIdentity) SetNetworkElementNil() {
	o.NetworkElement.Set(nil)
}

// UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
func (o *EquipmentFexIdentity) UnsetNetworkElement() {
	o.NetworkElement.Unset()
}

func (o EquipmentFexIdentity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentFexIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentPhysicalIdentity, errEquipmentPhysicalIdentity := json.Marshal(o.EquipmentPhysicalIdentity)
	if errEquipmentPhysicalIdentity != nil {
		return map[string]interface{}{}, errEquipmentPhysicalIdentity
	}
	errEquipmentPhysicalIdentity = json.Unmarshal([]byte(serializedEquipmentPhysicalIdentity), &toSerialize)
	if errEquipmentPhysicalIdentity != nil {
		return map[string]interface{}{}, errEquipmentPhysicalIdentity
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.SwitchId) {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.Fex.IsSet() {
		toSerialize["Fex"] = o.Fex.Get()
	}
	if o.NetworkElement.IsSet() {
		toSerialize["NetworkElement"] = o.NetworkElement.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentFexIdentity) UnmarshalJSON(data []byte) (err error) {
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
	type EquipmentFexIdentityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B.
		SwitchId       *int64                             `json:"SwitchId,omitempty"`
		Fex            NullableEquipmentFexRelationship   `json:"Fex,omitempty"`
		NetworkElement NullableNetworkElementRelationship `json:"NetworkElement,omitempty"`
	}

	varEquipmentFexIdentityWithoutEmbeddedStruct := EquipmentFexIdentityWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentFexIdentityWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentFexIdentity := _EquipmentFexIdentity{}
		varEquipmentFexIdentity.ClassId = varEquipmentFexIdentityWithoutEmbeddedStruct.ClassId
		varEquipmentFexIdentity.ObjectType = varEquipmentFexIdentityWithoutEmbeddedStruct.ObjectType
		varEquipmentFexIdentity.SwitchId = varEquipmentFexIdentityWithoutEmbeddedStruct.SwitchId
		varEquipmentFexIdentity.Fex = varEquipmentFexIdentityWithoutEmbeddedStruct.Fex
		varEquipmentFexIdentity.NetworkElement = varEquipmentFexIdentityWithoutEmbeddedStruct.NetworkElement
		*o = EquipmentFexIdentity(varEquipmentFexIdentity)
	} else {
		return err
	}

	varEquipmentFexIdentity := _EquipmentFexIdentity{}

	err = json.Unmarshal(data, &varEquipmentFexIdentity)
	if err == nil {
		o.EquipmentPhysicalIdentity = varEquipmentFexIdentity.EquipmentPhysicalIdentity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "Fex")
		delete(additionalProperties, "NetworkElement")

		// remove fields from embedded structs
		reflectEquipmentPhysicalIdentity := reflect.ValueOf(o.EquipmentPhysicalIdentity)
		for i := 0; i < reflectEquipmentPhysicalIdentity.Type().NumField(); i++ {
			t := reflectEquipmentPhysicalIdentity.Type().Field(i)

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

type NullableEquipmentFexIdentity struct {
	value *EquipmentFexIdentity
	isSet bool
}

func (v NullableEquipmentFexIdentity) Get() *EquipmentFexIdentity {
	return v.value
}

func (v *NullableEquipmentFexIdentity) Set(val *EquipmentFexIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentFexIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentFexIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentFexIdentity(val *EquipmentFexIdentity) *NullableEquipmentFexIdentity {
	return &NullableEquipmentFexIdentity{value: val, isSet: true}
}

func (v NullableEquipmentFexIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentFexIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
