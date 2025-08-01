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

// checks if the CapabilityCimcFirmwareDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityCimcFirmwareDescriptor{}

// CapabilityCimcFirmwareDescriptor Descriptor that identifies the server's redfish integration capability using cimc firmware info.
type CapabilityCimcFirmwareDescriptor struct {
	CapabilityHardwareDescriptor
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the server uses ep proxy to communicate with the adapter.
	AdapterEpProxyEnabled *bool `json:"AdapterEpProxyEnabled,omitempty"`
	// Minimum server firmware version for local users password properties feature support.
	LocalUserPswdSupportedVer *string `json:"LocalUserPswdSupportedVer,omitempty"`
	// Revision information for the server.
	Revision *string `json:"Revision,omitempty"`
	// Minimum server firmware version for UUID feature support.
	UuidSupportedVer     *string `json:"UuidSupportedVer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityCimcFirmwareDescriptor CapabilityCimcFirmwareDescriptor

// NewCapabilityCimcFirmwareDescriptor instantiates a new CapabilityCimcFirmwareDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityCimcFirmwareDescriptor(classId string, objectType string) *CapabilityCimcFirmwareDescriptor {
	this := CapabilityCimcFirmwareDescriptor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityCimcFirmwareDescriptorWithDefaults instantiates a new CapabilityCimcFirmwareDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityCimcFirmwareDescriptorWithDefaults() *CapabilityCimcFirmwareDescriptor {
	this := CapabilityCimcFirmwareDescriptor{}
	var classId string = "capability.CimcFirmwareDescriptor"
	this.ClassId = classId
	var objectType string = "capability.CimcFirmwareDescriptor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityCimcFirmwareDescriptor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityCimcFirmwareDescriptor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityCimcFirmwareDescriptor) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.CimcFirmwareDescriptor" of the ClassId field.
func (o *CapabilityCimcFirmwareDescriptor) GetDefaultClassId() interface{} {
	return "capability.CimcFirmwareDescriptor"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityCimcFirmwareDescriptor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityCimcFirmwareDescriptor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityCimcFirmwareDescriptor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.CimcFirmwareDescriptor" of the ObjectType field.
func (o *CapabilityCimcFirmwareDescriptor) GetDefaultObjectType() interface{} {
	return "capability.CimcFirmwareDescriptor"
}

// GetAdapterEpProxyEnabled returns the AdapterEpProxyEnabled field value if set, zero value otherwise.
func (o *CapabilityCimcFirmwareDescriptor) GetAdapterEpProxyEnabled() bool {
	if o == nil || IsNil(o.AdapterEpProxyEnabled) {
		var ret bool
		return ret
	}
	return *o.AdapterEpProxyEnabled
}

// GetAdapterEpProxyEnabledOk returns a tuple with the AdapterEpProxyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCimcFirmwareDescriptor) GetAdapterEpProxyEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AdapterEpProxyEnabled) {
		return nil, false
	}
	return o.AdapterEpProxyEnabled, true
}

// HasAdapterEpProxyEnabled returns a boolean if a field has been set.
func (o *CapabilityCimcFirmwareDescriptor) HasAdapterEpProxyEnabled() bool {
	if o != nil && !IsNil(o.AdapterEpProxyEnabled) {
		return true
	}

	return false
}

// SetAdapterEpProxyEnabled gets a reference to the given bool and assigns it to the AdapterEpProxyEnabled field.
func (o *CapabilityCimcFirmwareDescriptor) SetAdapterEpProxyEnabled(v bool) {
	o.AdapterEpProxyEnabled = &v
}

// GetLocalUserPswdSupportedVer returns the LocalUserPswdSupportedVer field value if set, zero value otherwise.
func (o *CapabilityCimcFirmwareDescriptor) GetLocalUserPswdSupportedVer() string {
	if o == nil || IsNil(o.LocalUserPswdSupportedVer) {
		var ret string
		return ret
	}
	return *o.LocalUserPswdSupportedVer
}

// GetLocalUserPswdSupportedVerOk returns a tuple with the LocalUserPswdSupportedVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCimcFirmwareDescriptor) GetLocalUserPswdSupportedVerOk() (*string, bool) {
	if o == nil || IsNil(o.LocalUserPswdSupportedVer) {
		return nil, false
	}
	return o.LocalUserPswdSupportedVer, true
}

// HasLocalUserPswdSupportedVer returns a boolean if a field has been set.
func (o *CapabilityCimcFirmwareDescriptor) HasLocalUserPswdSupportedVer() bool {
	if o != nil && !IsNil(o.LocalUserPswdSupportedVer) {
		return true
	}

	return false
}

// SetLocalUserPswdSupportedVer gets a reference to the given string and assigns it to the LocalUserPswdSupportedVer field.
func (o *CapabilityCimcFirmwareDescriptor) SetLocalUserPswdSupportedVer(v string) {
	o.LocalUserPswdSupportedVer = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *CapabilityCimcFirmwareDescriptor) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCimcFirmwareDescriptor) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *CapabilityCimcFirmwareDescriptor) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *CapabilityCimcFirmwareDescriptor) SetRevision(v string) {
	o.Revision = &v
}

// GetUuidSupportedVer returns the UuidSupportedVer field value if set, zero value otherwise.
func (o *CapabilityCimcFirmwareDescriptor) GetUuidSupportedVer() string {
	if o == nil || IsNil(o.UuidSupportedVer) {
		var ret string
		return ret
	}
	return *o.UuidSupportedVer
}

// GetUuidSupportedVerOk returns a tuple with the UuidSupportedVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCimcFirmwareDescriptor) GetUuidSupportedVerOk() (*string, bool) {
	if o == nil || IsNil(o.UuidSupportedVer) {
		return nil, false
	}
	return o.UuidSupportedVer, true
}

// HasUuidSupportedVer returns a boolean if a field has been set.
func (o *CapabilityCimcFirmwareDescriptor) HasUuidSupportedVer() bool {
	if o != nil && !IsNil(o.UuidSupportedVer) {
		return true
	}

	return false
}

// SetUuidSupportedVer gets a reference to the given string and assigns it to the UuidSupportedVer field.
func (o *CapabilityCimcFirmwareDescriptor) SetUuidSupportedVer(v string) {
	o.UuidSupportedVer = &v
}

func (o CapabilityCimcFirmwareDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityCimcFirmwareDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityHardwareDescriptor, errCapabilityHardwareDescriptor := json.Marshal(o.CapabilityHardwareDescriptor)
	if errCapabilityHardwareDescriptor != nil {
		return map[string]interface{}{}, errCapabilityHardwareDescriptor
	}
	errCapabilityHardwareDescriptor = json.Unmarshal([]byte(serializedCapabilityHardwareDescriptor), &toSerialize)
	if errCapabilityHardwareDescriptor != nil {
		return map[string]interface{}{}, errCapabilityHardwareDescriptor
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AdapterEpProxyEnabled) {
		toSerialize["AdapterEpProxyEnabled"] = o.AdapterEpProxyEnabled
	}
	if !IsNil(o.LocalUserPswdSupportedVer) {
		toSerialize["LocalUserPswdSupportedVer"] = o.LocalUserPswdSupportedVer
	}
	if !IsNil(o.Revision) {
		toSerialize["Revision"] = o.Revision
	}
	if !IsNil(o.UuidSupportedVer) {
		toSerialize["UuidSupportedVer"] = o.UuidSupportedVer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityCimcFirmwareDescriptor) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the server uses ep proxy to communicate with the adapter.
		AdapterEpProxyEnabled *bool `json:"AdapterEpProxyEnabled,omitempty"`
		// Minimum server firmware version for local users password properties feature support.
		LocalUserPswdSupportedVer *string `json:"LocalUserPswdSupportedVer,omitempty"`
		// Revision information for the server.
		Revision *string `json:"Revision,omitempty"`
		// Minimum server firmware version for UUID feature support.
		UuidSupportedVer *string `json:"UuidSupportedVer,omitempty"`
	}

	varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct := CapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityCimcFirmwareDescriptor := _CapabilityCimcFirmwareDescriptor{}
		varCapabilityCimcFirmwareDescriptor.ClassId = varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct.ClassId
		varCapabilityCimcFirmwareDescriptor.ObjectType = varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct.ObjectType
		varCapabilityCimcFirmwareDescriptor.AdapterEpProxyEnabled = varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct.AdapterEpProxyEnabled
		varCapabilityCimcFirmwareDescriptor.LocalUserPswdSupportedVer = varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct.LocalUserPswdSupportedVer
		varCapabilityCimcFirmwareDescriptor.Revision = varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct.Revision
		varCapabilityCimcFirmwareDescriptor.UuidSupportedVer = varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct.UuidSupportedVer
		*o = CapabilityCimcFirmwareDescriptor(varCapabilityCimcFirmwareDescriptor)
	} else {
		return err
	}

	varCapabilityCimcFirmwareDescriptor := _CapabilityCimcFirmwareDescriptor{}

	err = json.Unmarshal(data, &varCapabilityCimcFirmwareDescriptor)
	if err == nil {
		o.CapabilityHardwareDescriptor = varCapabilityCimcFirmwareDescriptor.CapabilityHardwareDescriptor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterEpProxyEnabled")
		delete(additionalProperties, "LocalUserPswdSupportedVer")
		delete(additionalProperties, "Revision")
		delete(additionalProperties, "UuidSupportedVer")

		// remove fields from embedded structs
		reflectCapabilityHardwareDescriptor := reflect.ValueOf(o.CapabilityHardwareDescriptor)
		for i := 0; i < reflectCapabilityHardwareDescriptor.Type().NumField(); i++ {
			t := reflectCapabilityHardwareDescriptor.Type().Field(i)

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

type NullableCapabilityCimcFirmwareDescriptor struct {
	value *CapabilityCimcFirmwareDescriptor
	isSet bool
}

func (v NullableCapabilityCimcFirmwareDescriptor) Get() *CapabilityCimcFirmwareDescriptor {
	return v.value
}

func (v *NullableCapabilityCimcFirmwareDescriptor) Set(val *CapabilityCimcFirmwareDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityCimcFirmwareDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityCimcFirmwareDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityCimcFirmwareDescriptor(val *CapabilityCimcFirmwareDescriptor) *NullableCapabilityCimcFirmwareDescriptor {
	return &NullableCapabilityCimcFirmwareDescriptor{value: val, isSet: true}
}

func (v NullableCapabilityCimcFirmwareDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityCimcFirmwareDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
