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

// checks if the FabricSpanSourceVnicEthIf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricSpanSourceVnicEthIf{}

// FabricSpanSourceVnicEthIf Configures a VNIC as the SPAN Source for a given SPAN session. For failover VNICs, VNIC must be added as a source to a SPAN session on the peer Fabric Interconnect to ensure that traffic is mirrored on both Fabric Interconnects.
type FabricSpanSourceVnicEthIf struct {
	FabricAbstractSpanSourceVirtualIf
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                        `json:"ObjectType"`
	Vnic                 NullableVnicEthIfRelationship `json:"Vnic,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricSpanSourceVnicEthIf FabricSpanSourceVnicEthIf

// NewFabricSpanSourceVnicEthIf instantiates a new FabricSpanSourceVnicEthIf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSpanSourceVnicEthIf(classId string, objectType string) *FabricSpanSourceVnicEthIf {
	this := FabricSpanSourceVnicEthIf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var direction string = "Receive"
	this.Direction = &direction
	return &this
}

// NewFabricSpanSourceVnicEthIfWithDefaults instantiates a new FabricSpanSourceVnicEthIf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSpanSourceVnicEthIfWithDefaults() *FabricSpanSourceVnicEthIf {
	this := FabricSpanSourceVnicEthIf{}
	var classId string = "fabric.SpanSourceVnicEthIf"
	this.ClassId = classId
	var objectType string = "fabric.SpanSourceVnicEthIf"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSpanSourceVnicEthIf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSpanSourceVnicEthIf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSpanSourceVnicEthIf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricSpanSourceVnicEthIf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSpanSourceVnicEthIf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSpanSourceVnicEthIf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVnic returns the Vnic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSpanSourceVnicEthIf) GetVnic() VnicEthIfRelationship {
	if o == nil || IsNil(o.Vnic.Get()) {
		var ret VnicEthIfRelationship
		return ret
	}
	return *o.Vnic.Get()
}

// GetVnicOk returns a tuple with the Vnic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSpanSourceVnicEthIf) GetVnicOk() (*VnicEthIfRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vnic.Get(), o.Vnic.IsSet()
}

// HasVnic returns a boolean if a field has been set.
func (o *FabricSpanSourceVnicEthIf) HasVnic() bool {
	if o != nil && o.Vnic.IsSet() {
		return true
	}

	return false
}

// SetVnic gets a reference to the given NullableVnicEthIfRelationship and assigns it to the Vnic field.
func (o *FabricSpanSourceVnicEthIf) SetVnic(v VnicEthIfRelationship) {
	o.Vnic.Set(&v)
}

// SetVnicNil sets the value for Vnic to be an explicit nil
func (o *FabricSpanSourceVnicEthIf) SetVnicNil() {
	o.Vnic.Set(nil)
}

// UnsetVnic ensures that no value is present for Vnic, not even an explicit nil
func (o *FabricSpanSourceVnicEthIf) UnsetVnic() {
	o.Vnic.Unset()
}

func (o FabricSpanSourceVnicEthIf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricSpanSourceVnicEthIf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricAbstractSpanSourceVirtualIf, errFabricAbstractSpanSourceVirtualIf := json.Marshal(o.FabricAbstractSpanSourceVirtualIf)
	if errFabricAbstractSpanSourceVirtualIf != nil {
		return map[string]interface{}{}, errFabricAbstractSpanSourceVirtualIf
	}
	errFabricAbstractSpanSourceVirtualIf = json.Unmarshal([]byte(serializedFabricAbstractSpanSourceVirtualIf), &toSerialize)
	if errFabricAbstractSpanSourceVirtualIf != nil {
		return map[string]interface{}{}, errFabricAbstractSpanSourceVirtualIf
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if o.Vnic.IsSet() {
		toSerialize["Vnic"] = o.Vnic.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricSpanSourceVnicEthIf) UnmarshalJSON(data []byte) (err error) {
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

	type FabricSpanSourceVnicEthIfWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                        `json:"ObjectType"`
		Vnic       NullableVnicEthIfRelationship `json:"Vnic,omitempty"`
	}

	varFabricSpanSourceVnicEthIfWithoutEmbeddedStruct := FabricSpanSourceVnicEthIfWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricSpanSourceVnicEthIfWithoutEmbeddedStruct)
	if err == nil {
		varFabricSpanSourceVnicEthIf := _FabricSpanSourceVnicEthIf{}
		varFabricSpanSourceVnicEthIf.ClassId = varFabricSpanSourceVnicEthIfWithoutEmbeddedStruct.ClassId
		varFabricSpanSourceVnicEthIf.ObjectType = varFabricSpanSourceVnicEthIfWithoutEmbeddedStruct.ObjectType
		varFabricSpanSourceVnicEthIf.Vnic = varFabricSpanSourceVnicEthIfWithoutEmbeddedStruct.Vnic
		*o = FabricSpanSourceVnicEthIf(varFabricSpanSourceVnicEthIf)
	} else {
		return err
	}

	varFabricSpanSourceVnicEthIf := _FabricSpanSourceVnicEthIf{}

	err = json.Unmarshal(data, &varFabricSpanSourceVnicEthIf)
	if err == nil {
		o.FabricAbstractSpanSourceVirtualIf = varFabricSpanSourceVnicEthIf.FabricAbstractSpanSourceVirtualIf
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Vnic")

		// remove fields from embedded structs
		reflectFabricAbstractSpanSourceVirtualIf := reflect.ValueOf(o.FabricAbstractSpanSourceVirtualIf)
		for i := 0; i < reflectFabricAbstractSpanSourceVirtualIf.Type().NumField(); i++ {
			t := reflectFabricAbstractSpanSourceVirtualIf.Type().Field(i)

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

type NullableFabricSpanSourceVnicEthIf struct {
	value *FabricSpanSourceVnicEthIf
	isSet bool
}

func (v NullableFabricSpanSourceVnicEthIf) Get() *FabricSpanSourceVnicEthIf {
	return v.value
}

func (v *NullableFabricSpanSourceVnicEthIf) Set(val *FabricSpanSourceVnicEthIf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSpanSourceVnicEthIf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSpanSourceVnicEthIf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSpanSourceVnicEthIf(val *FabricSpanSourceVnicEthIf) *NullableFabricSpanSourceVnicEthIf {
	return &NullableFabricSpanSourceVnicEthIf{value: val, isSet: true}
}

func (v NullableFabricSpanSourceVnicEthIf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSpanSourceVnicEthIf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}