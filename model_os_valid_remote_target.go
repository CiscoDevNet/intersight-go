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

// checks if the OsValidRemoteTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsValidRemoteTarget{}

// OsValidRemoteTarget ValidRemoteTarget is used to fetch all the valid remote targets for the servers from its server profile. The List of Install targets includes iSCSI and Fibre Channel paths.
type OsValidRemoteTarget struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType       string                   `json:"ObjectType"`
	FibreChannelLuns []OsFibreChannelResponse `json:"FibreChannelLuns,omitempty"`
	IscsiLuns        []OsIscsiLunResponse     `json:"IscsiLuns,omitempty"`
	// Flag to denote the source of the request. If the call is from Orchestration UI, only the flat list of Install targets can be sent as response.
	Src                  *string                             `json:"Src,omitempty"`
	Server               NullableComputePhysicalRelationship `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsValidRemoteTarget OsValidRemoteTarget

// NewOsValidRemoteTarget instantiates a new OsValidRemoteTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsValidRemoteTarget(classId string, objectType string) *OsValidRemoteTarget {
	this := OsValidRemoteTarget{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsValidRemoteTargetWithDefaults instantiates a new OsValidRemoteTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsValidRemoteTargetWithDefaults() *OsValidRemoteTarget {
	this := OsValidRemoteTarget{}
	var classId string = "os.ValidRemoteTarget"
	this.ClassId = classId
	var objectType string = "os.ValidRemoteTarget"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsValidRemoteTarget) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsValidRemoteTarget) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsValidRemoteTarget) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "os.ValidRemoteTarget" of the ClassId field.
func (o *OsValidRemoteTarget) GetDefaultClassId() interface{} {
	return "os.ValidRemoteTarget"
}

// GetObjectType returns the ObjectType field value
func (o *OsValidRemoteTarget) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsValidRemoteTarget) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsValidRemoteTarget) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "os.ValidRemoteTarget" of the ObjectType field.
func (o *OsValidRemoteTarget) GetDefaultObjectType() interface{} {
	return "os.ValidRemoteTarget"
}

// GetFibreChannelLuns returns the FibreChannelLuns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidRemoteTarget) GetFibreChannelLuns() []OsFibreChannelResponse {
	if o == nil {
		var ret []OsFibreChannelResponse
		return ret
	}
	return o.FibreChannelLuns
}

// GetFibreChannelLunsOk returns a tuple with the FibreChannelLuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidRemoteTarget) GetFibreChannelLunsOk() ([]OsFibreChannelResponse, bool) {
	if o == nil || IsNil(o.FibreChannelLuns) {
		return nil, false
	}
	return o.FibreChannelLuns, true
}

// HasFibreChannelLuns returns a boolean if a field has been set.
func (o *OsValidRemoteTarget) HasFibreChannelLuns() bool {
	if o != nil && !IsNil(o.FibreChannelLuns) {
		return true
	}

	return false
}

// SetFibreChannelLuns gets a reference to the given []OsFibreChannelResponse and assigns it to the FibreChannelLuns field.
func (o *OsValidRemoteTarget) SetFibreChannelLuns(v []OsFibreChannelResponse) {
	o.FibreChannelLuns = v
}

// GetIscsiLuns returns the IscsiLuns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidRemoteTarget) GetIscsiLuns() []OsIscsiLunResponse {
	if o == nil {
		var ret []OsIscsiLunResponse
		return ret
	}
	return o.IscsiLuns
}

// GetIscsiLunsOk returns a tuple with the IscsiLuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidRemoteTarget) GetIscsiLunsOk() ([]OsIscsiLunResponse, bool) {
	if o == nil || IsNil(o.IscsiLuns) {
		return nil, false
	}
	return o.IscsiLuns, true
}

// HasIscsiLuns returns a boolean if a field has been set.
func (o *OsValidRemoteTarget) HasIscsiLuns() bool {
	if o != nil && !IsNil(o.IscsiLuns) {
		return true
	}

	return false
}

// SetIscsiLuns gets a reference to the given []OsIscsiLunResponse and assigns it to the IscsiLuns field.
func (o *OsValidRemoteTarget) SetIscsiLuns(v []OsIscsiLunResponse) {
	o.IscsiLuns = v
}

// GetSrc returns the Src field value if set, zero value otherwise.
func (o *OsValidRemoteTarget) GetSrc() string {
	if o == nil || IsNil(o.Src) {
		var ret string
		return ret
	}
	return *o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsValidRemoteTarget) GetSrcOk() (*string, bool) {
	if o == nil || IsNil(o.Src) {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *OsValidRemoteTarget) HasSrc() bool {
	if o != nil && !IsNil(o.Src) {
		return true
	}

	return false
}

// SetSrc gets a reference to the given string and assigns it to the Src field.
func (o *OsValidRemoteTarget) SetSrc(v string) {
	o.Src = &v
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidRemoteTarget) GetServer() ComputePhysicalRelationship {
	if o == nil || IsNil(o.Server.Get()) {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server.Get()
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidRemoteTarget) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Server.Get(), o.Server.IsSet()
}

// HasServer returns a boolean if a field has been set.
func (o *OsValidRemoteTarget) HasServer() bool {
	if o != nil && o.Server.IsSet() {
		return true
	}

	return false
}

// SetServer gets a reference to the given NullableComputePhysicalRelationship and assigns it to the Server field.
func (o *OsValidRemoteTarget) SetServer(v ComputePhysicalRelationship) {
	o.Server.Set(&v)
}

// SetServerNil sets the value for Server to be an explicit nil
func (o *OsValidRemoteTarget) SetServerNil() {
	o.Server.Set(nil)
}

// UnsetServer ensures that no value is present for Server, not even an explicit nil
func (o *OsValidRemoteTarget) UnsetServer() {
	o.Server.Unset()
}

func (o OsValidRemoteTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsValidRemoteTarget) ToMap() (map[string]interface{}, error) {
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
	if o.FibreChannelLuns != nil {
		toSerialize["FibreChannelLuns"] = o.FibreChannelLuns
	}
	if o.IscsiLuns != nil {
		toSerialize["IscsiLuns"] = o.IscsiLuns
	}
	if !IsNil(o.Src) {
		toSerialize["Src"] = o.Src
	}
	if o.Server.IsSet() {
		toSerialize["Server"] = o.Server.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OsValidRemoteTarget) UnmarshalJSON(data []byte) (err error) {
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
	type OsValidRemoteTargetWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string                   `json:"ObjectType"`
		FibreChannelLuns []OsFibreChannelResponse `json:"FibreChannelLuns,omitempty"`
		IscsiLuns        []OsIscsiLunResponse     `json:"IscsiLuns,omitempty"`
		// Flag to denote the source of the request. If the call is from Orchestration UI, only the flat list of Install targets can be sent as response.
		Src    *string                             `json:"Src,omitempty"`
		Server NullableComputePhysicalRelationship `json:"Server,omitempty"`
	}

	varOsValidRemoteTargetWithoutEmbeddedStruct := OsValidRemoteTargetWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOsValidRemoteTargetWithoutEmbeddedStruct)
	if err == nil {
		varOsValidRemoteTarget := _OsValidRemoteTarget{}
		varOsValidRemoteTarget.ClassId = varOsValidRemoteTargetWithoutEmbeddedStruct.ClassId
		varOsValidRemoteTarget.ObjectType = varOsValidRemoteTargetWithoutEmbeddedStruct.ObjectType
		varOsValidRemoteTarget.FibreChannelLuns = varOsValidRemoteTargetWithoutEmbeddedStruct.FibreChannelLuns
		varOsValidRemoteTarget.IscsiLuns = varOsValidRemoteTargetWithoutEmbeddedStruct.IscsiLuns
		varOsValidRemoteTarget.Src = varOsValidRemoteTargetWithoutEmbeddedStruct.Src
		varOsValidRemoteTarget.Server = varOsValidRemoteTargetWithoutEmbeddedStruct.Server
		*o = OsValidRemoteTarget(varOsValidRemoteTarget)
	} else {
		return err
	}

	varOsValidRemoteTarget := _OsValidRemoteTarget{}

	err = json.Unmarshal(data, &varOsValidRemoteTarget)
	if err == nil {
		o.MoBaseMo = varOsValidRemoteTarget.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FibreChannelLuns")
		delete(additionalProperties, "IscsiLuns")
		delete(additionalProperties, "Src")
		delete(additionalProperties, "Server")

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

type NullableOsValidRemoteTarget struct {
	value *OsValidRemoteTarget
	isSet bool
}

func (v NullableOsValidRemoteTarget) Get() *OsValidRemoteTarget {
	return v.value
}

func (v *NullableOsValidRemoteTarget) Set(val *OsValidRemoteTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableOsValidRemoteTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableOsValidRemoteTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsValidRemoteTarget(val *OsValidRemoteTarget) *NullableOsValidRemoteTarget {
	return &NullableOsValidRemoteTarget{value: val, isSet: true}
}

func (v NullableOsValidRemoteTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsValidRemoteTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
