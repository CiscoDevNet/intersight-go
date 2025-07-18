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

// checks if the FabricPortChannelRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricPortChannelRole{}

// FabricPortChannelRole Configuration object sent by user to apply on a port-channel.
type FabricPortChannelRole struct {
	FabricAbstractInterfaceRole
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Unique Identifier of the port-channel, local to this switch.
	PcId                 *int64                               `json:"PcId,omitempty"`
	Ports                []FabricPortIdentifier               `json:"Ports,omitempty"`
	PortPolicy           NullableFabricPortPolicyRelationship `json:"PortPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricPortChannelRole FabricPortChannelRole

// NewFabricPortChannelRole instantiates a new FabricPortChannelRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricPortChannelRole(classId string, objectType string) *FabricPortChannelRole {
	this := FabricPortChannelRole{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricPortChannelRoleWithDefaults instantiates a new FabricPortChannelRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricPortChannelRoleWithDefaults() *FabricPortChannelRole {
	this := FabricPortChannelRole{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricPortChannelRole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricPortChannelRole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricPortChannelRole) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricPortChannelRole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricPortChannelRole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricPortChannelRole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPcId returns the PcId field value if set, zero value otherwise.
func (o *FabricPortChannelRole) GetPcId() int64 {
	if o == nil || IsNil(o.PcId) {
		var ret int64
		return ret
	}
	return *o.PcId
}

// GetPcIdOk returns a tuple with the PcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortChannelRole) GetPcIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PcId) {
		return nil, false
	}
	return o.PcId, true
}

// HasPcId returns a boolean if a field has been set.
func (o *FabricPortChannelRole) HasPcId() bool {
	if o != nil && !IsNil(o.PcId) {
		return true
	}

	return false
}

// SetPcId gets a reference to the given int64 and assigns it to the PcId field.
func (o *FabricPortChannelRole) SetPcId(v int64) {
	o.PcId = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricPortChannelRole) GetPorts() []FabricPortIdentifier {
	if o == nil {
		var ret []FabricPortIdentifier
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricPortChannelRole) GetPortsOk() ([]FabricPortIdentifier, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *FabricPortChannelRole) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []FabricPortIdentifier and assigns it to the Ports field.
func (o *FabricPortChannelRole) SetPorts(v []FabricPortIdentifier) {
	o.Ports = v
}

// GetPortPolicy returns the PortPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricPortChannelRole) GetPortPolicy() FabricPortPolicyRelationship {
	if o == nil || IsNil(o.PortPolicy.Get()) {
		var ret FabricPortPolicyRelationship
		return ret
	}
	return *o.PortPolicy.Get()
}

// GetPortPolicyOk returns a tuple with the PortPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricPortChannelRole) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortPolicy.Get(), o.PortPolicy.IsSet()
}

// HasPortPolicy returns a boolean if a field has been set.
func (o *FabricPortChannelRole) HasPortPolicy() bool {
	if o != nil && o.PortPolicy.IsSet() {
		return true
	}

	return false
}

// SetPortPolicy gets a reference to the given NullableFabricPortPolicyRelationship and assigns it to the PortPolicy field.
func (o *FabricPortChannelRole) SetPortPolicy(v FabricPortPolicyRelationship) {
	o.PortPolicy.Set(&v)
}

// SetPortPolicyNil sets the value for PortPolicy to be an explicit nil
func (o *FabricPortChannelRole) SetPortPolicyNil() {
	o.PortPolicy.Set(nil)
}

// UnsetPortPolicy ensures that no value is present for PortPolicy, not even an explicit nil
func (o *FabricPortChannelRole) UnsetPortPolicy() {
	o.PortPolicy.Unset()
}

func (o FabricPortChannelRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricPortChannelRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricAbstractInterfaceRole, errFabricAbstractInterfaceRole := json.Marshal(o.FabricAbstractInterfaceRole)
	if errFabricAbstractInterfaceRole != nil {
		return map[string]interface{}{}, errFabricAbstractInterfaceRole
	}
	errFabricAbstractInterfaceRole = json.Unmarshal([]byte(serializedFabricAbstractInterfaceRole), &toSerialize)
	if errFabricAbstractInterfaceRole != nil {
		return map[string]interface{}{}, errFabricAbstractInterfaceRole
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.PcId) {
		toSerialize["PcId"] = o.PcId
	}
	if o.Ports != nil {
		toSerialize["Ports"] = o.Ports
	}
	if o.PortPolicy.IsSet() {
		toSerialize["PortPolicy"] = o.PortPolicy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricPortChannelRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type FabricPortChannelRoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Unique Identifier of the port-channel, local to this switch.
		PcId       *int64                               `json:"PcId,omitempty"`
		Ports      []FabricPortIdentifier               `json:"Ports,omitempty"`
		PortPolicy NullableFabricPortPolicyRelationship `json:"PortPolicy,omitempty"`
	}

	varFabricPortChannelRoleWithoutEmbeddedStruct := FabricPortChannelRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricPortChannelRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricPortChannelRole := _FabricPortChannelRole{}
		varFabricPortChannelRole.ClassId = varFabricPortChannelRoleWithoutEmbeddedStruct.ClassId
		varFabricPortChannelRole.ObjectType = varFabricPortChannelRoleWithoutEmbeddedStruct.ObjectType
		varFabricPortChannelRole.PcId = varFabricPortChannelRoleWithoutEmbeddedStruct.PcId
		varFabricPortChannelRole.Ports = varFabricPortChannelRoleWithoutEmbeddedStruct.Ports
		varFabricPortChannelRole.PortPolicy = varFabricPortChannelRoleWithoutEmbeddedStruct.PortPolicy
		*o = FabricPortChannelRole(varFabricPortChannelRole)
	} else {
		return err
	}

	varFabricPortChannelRole := _FabricPortChannelRole{}

	err = json.Unmarshal(data, &varFabricPortChannelRole)
	if err == nil {
		o.FabricAbstractInterfaceRole = varFabricPortChannelRole.FabricAbstractInterfaceRole
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PcId")
		delete(additionalProperties, "Ports")
		delete(additionalProperties, "PortPolicy")

		// remove fields from embedded structs
		reflectFabricAbstractInterfaceRole := reflect.ValueOf(o.FabricAbstractInterfaceRole)
		for i := 0; i < reflectFabricAbstractInterfaceRole.Type().NumField(); i++ {
			t := reflectFabricAbstractInterfaceRole.Type().Field(i)

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

type NullableFabricPortChannelRole struct {
	value *FabricPortChannelRole
	isSet bool
}

func (v NullableFabricPortChannelRole) Get() *FabricPortChannelRole {
	return v.value
}

func (v *NullableFabricPortChannelRole) Set(val *FabricPortChannelRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricPortChannelRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricPortChannelRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricPortChannelRole(val *FabricPortChannelRole) *NullableFabricPortChannelRole {
	return &NullableFabricPortChannelRole{value: val, isSet: true}
}

func (v NullableFabricPortChannelRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricPortChannelRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
