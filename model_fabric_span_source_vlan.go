/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-17769
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

// checks if the FabricSpanSourceVlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricSpanSourceVlan{}

// FabricSpanSourceVlan Configures a VLAN as the SPAN Source for a given SPAN session. For VLAN that is common to both Fabric Interconnects, VLAN must be added as a source to a SPAN session on the peer Fabric Interconnect to ensure that traffic is mirrored on both Fabric Interconnects.
type FabricSpanSourceVlan struct {
	FabricAbstractSpanSource
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// SPAN source VLAN Identifier.
	VlanId               *int64                                `json:"VlanId,omitempty"`
	SpanSession          NullableFabricSpanSessionRelationship `json:"SpanSession,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricSpanSourceVlan FabricSpanSourceVlan

// NewFabricSpanSourceVlan instantiates a new FabricSpanSourceVlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSpanSourceVlan(classId string, objectType string) *FabricSpanSourceVlan {
	this := FabricSpanSourceVlan{}
	this.ClassId = classId
	this.ObjectType = objectType
	var direction string = "Receive"
	this.Direction = &direction
	return &this
}

// NewFabricSpanSourceVlanWithDefaults instantiates a new FabricSpanSourceVlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSpanSourceVlanWithDefaults() *FabricSpanSourceVlan {
	this := FabricSpanSourceVlan{}
	var classId string = "fabric.SpanSourceVlan"
	this.ClassId = classId
	var objectType string = "fabric.SpanSourceVlan"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSpanSourceVlan) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSpanSourceVlan) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSpanSourceVlan) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricSpanSourceVlan) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSpanSourceVlan) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSpanSourceVlan) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *FabricSpanSourceVlan) GetVlanId() int64 {
	if o == nil || IsNil(o.VlanId) {
		var ret int64
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSpanSourceVlan) GetVlanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *FabricSpanSourceVlan) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int64 and assigns it to the VlanId field.
func (o *FabricSpanSourceVlan) SetVlanId(v int64) {
	o.VlanId = &v
}

// GetSpanSession returns the SpanSession field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSpanSourceVlan) GetSpanSession() FabricSpanSessionRelationship {
	if o == nil || IsNil(o.SpanSession.Get()) {
		var ret FabricSpanSessionRelationship
		return ret
	}
	return *o.SpanSession.Get()
}

// GetSpanSessionOk returns a tuple with the SpanSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSpanSourceVlan) GetSpanSessionOk() (*FabricSpanSessionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpanSession.Get(), o.SpanSession.IsSet()
}

// HasSpanSession returns a boolean if a field has been set.
func (o *FabricSpanSourceVlan) HasSpanSession() bool {
	if o != nil && o.SpanSession.IsSet() {
		return true
	}

	return false
}

// SetSpanSession gets a reference to the given NullableFabricSpanSessionRelationship and assigns it to the SpanSession field.
func (o *FabricSpanSourceVlan) SetSpanSession(v FabricSpanSessionRelationship) {
	o.SpanSession.Set(&v)
}

// SetSpanSessionNil sets the value for SpanSession to be an explicit nil
func (o *FabricSpanSourceVlan) SetSpanSessionNil() {
	o.SpanSession.Set(nil)
}

// UnsetSpanSession ensures that no value is present for SpanSession, not even an explicit nil
func (o *FabricSpanSourceVlan) UnsetSpanSession() {
	o.SpanSession.Unset()
}

func (o FabricSpanSourceVlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricSpanSourceVlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricAbstractSpanSource, errFabricAbstractSpanSource := json.Marshal(o.FabricAbstractSpanSource)
	if errFabricAbstractSpanSource != nil {
		return map[string]interface{}{}, errFabricAbstractSpanSource
	}
	errFabricAbstractSpanSource = json.Unmarshal([]byte(serializedFabricAbstractSpanSource), &toSerialize)
	if errFabricAbstractSpanSource != nil {
		return map[string]interface{}{}, errFabricAbstractSpanSource
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.VlanId) {
		toSerialize["VlanId"] = o.VlanId
	}
	if o.SpanSession.IsSet() {
		toSerialize["SpanSession"] = o.SpanSession.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricSpanSourceVlan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type FabricSpanSourceVlanWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// SPAN source VLAN Identifier.
		VlanId      *int64                                `json:"VlanId,omitempty"`
		SpanSession NullableFabricSpanSessionRelationship `json:"SpanSession,omitempty"`
	}

	varFabricSpanSourceVlanWithoutEmbeddedStruct := FabricSpanSourceVlanWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricSpanSourceVlanWithoutEmbeddedStruct)
	if err == nil {
		varFabricSpanSourceVlan := _FabricSpanSourceVlan{}
		varFabricSpanSourceVlan.ClassId = varFabricSpanSourceVlanWithoutEmbeddedStruct.ClassId
		varFabricSpanSourceVlan.ObjectType = varFabricSpanSourceVlanWithoutEmbeddedStruct.ObjectType
		varFabricSpanSourceVlan.VlanId = varFabricSpanSourceVlanWithoutEmbeddedStruct.VlanId
		varFabricSpanSourceVlan.SpanSession = varFabricSpanSourceVlanWithoutEmbeddedStruct.SpanSession
		*o = FabricSpanSourceVlan(varFabricSpanSourceVlan)
	} else {
		return err
	}

	varFabricSpanSourceVlan := _FabricSpanSourceVlan{}

	err = json.Unmarshal(data, &varFabricSpanSourceVlan)
	if err == nil {
		o.FabricAbstractSpanSource = varFabricSpanSourceVlan.FabricAbstractSpanSource
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "VlanId")
		delete(additionalProperties, "SpanSession")

		// remove fields from embedded structs
		reflectFabricAbstractSpanSource := reflect.ValueOf(o.FabricAbstractSpanSource)
		for i := 0; i < reflectFabricAbstractSpanSource.Type().NumField(); i++ {
			t := reflectFabricAbstractSpanSource.Type().Field(i)

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

type NullableFabricSpanSourceVlan struct {
	value *FabricSpanSourceVlan
	isSet bool
}

func (v NullableFabricSpanSourceVlan) Get() *FabricSpanSourceVlan {
	return v.value
}

func (v *NullableFabricSpanSourceVlan) Set(val *FabricSpanSourceVlan) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSpanSourceVlan) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSpanSourceVlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSpanSourceVlan(val *FabricSpanSourceVlan) *NullableFabricSpanSourceVlan {
	return &NullableFabricSpanSourceVlan{value: val, isSet: true}
}

func (v NullableFabricSpanSourceVlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSpanSourceVlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}