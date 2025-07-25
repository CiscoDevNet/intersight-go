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

// checks if the VnicVmqSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicVmqSettings{}

// VnicVmqSettings Virtual Machine Queue Settings for the virtual interface that allow efficient transfer of network traffic to the guest OS.
type VnicVmqSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables VMQ feature on the virtual interface.
	Enabled *bool `json:"Enabled,omitempty"`
	// Enables Virtual Machine Multi-Queue feature on the virtual interface. VMMQ allows configuration of multiple I/O queues for a single VM and thus distributes traffic across multiple CPU cores in a VM.
	MultiQueueSupport *bool `json:"MultiQueueSupport,omitempty"`
	// The number of interrupt resources to be allocated. Recommended value is the number of CPU threads or logical processors available in the server.
	NumInterrupts *int64 `json:"NumInterrupts,omitempty"`
	// The number of sub vNICs to be created.
	NumSubVnics *int64 `json:"NumSubVnics,omitempty"`
	// The number of hardware Virtual Machine Queues to be allocated. The number of VMQs per adapter must be one more than the maximum number of VM NICs.
	NumVmqs *int64 `json:"NumVmqs,omitempty"`
	// Ethernet Adapter policy to be associated with the Sub vNICs. The Transmit Queue and Receive Queue resource value of VMMQ adapter policy should be greater than or equal to the configured number of sub vNICs.
	VmmqAdapterPolicy    *string `json:"VmmqAdapterPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicVmqSettings VnicVmqSettings

// NewVnicVmqSettings instantiates a new VnicVmqSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicVmqSettings(classId string, objectType string) *VnicVmqSettings {
	this := VnicVmqSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	var multiQueueSupport bool = false
	this.MultiQueueSupport = &multiQueueSupport
	var numInterrupts int64 = 16
	this.NumInterrupts = &numInterrupts
	var numSubVnics int64 = 64
	this.NumSubVnics = &numSubVnics
	var numVmqs int64 = 4
	this.NumVmqs = &numVmqs
	return &this
}

// NewVnicVmqSettingsWithDefaults instantiates a new VnicVmqSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicVmqSettingsWithDefaults() *VnicVmqSettings {
	this := VnicVmqSettings{}
	var classId string = "vnic.VmqSettings"
	this.ClassId = classId
	var objectType string = "vnic.VmqSettings"
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	var multiQueueSupport bool = false
	this.MultiQueueSupport = &multiQueueSupport
	var numInterrupts int64 = 16
	this.NumInterrupts = &numInterrupts
	var numSubVnics int64 = 64
	this.NumSubVnics = &numSubVnics
	var numVmqs int64 = 4
	this.NumVmqs = &numVmqs
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicVmqSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicVmqSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicVmqSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.VmqSettings" of the ClassId field.
func (o *VnicVmqSettings) GetDefaultClassId() interface{} {
	return "vnic.VmqSettings"
}

// GetObjectType returns the ObjectType field value
func (o *VnicVmqSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicVmqSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicVmqSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.VmqSettings" of the ObjectType field.
func (o *VnicVmqSettings) GetDefaultObjectType() interface{} {
	return "vnic.VmqSettings"
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VnicVmqSettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVmqSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VnicVmqSettings) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VnicVmqSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMultiQueueSupport returns the MultiQueueSupport field value if set, zero value otherwise.
func (o *VnicVmqSettings) GetMultiQueueSupport() bool {
	if o == nil || IsNil(o.MultiQueueSupport) {
		var ret bool
		return ret
	}
	return *o.MultiQueueSupport
}

// GetMultiQueueSupportOk returns a tuple with the MultiQueueSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVmqSettings) GetMultiQueueSupportOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiQueueSupport) {
		return nil, false
	}
	return o.MultiQueueSupport, true
}

// HasMultiQueueSupport returns a boolean if a field has been set.
func (o *VnicVmqSettings) HasMultiQueueSupport() bool {
	if o != nil && !IsNil(o.MultiQueueSupport) {
		return true
	}

	return false
}

// SetMultiQueueSupport gets a reference to the given bool and assigns it to the MultiQueueSupport field.
func (o *VnicVmqSettings) SetMultiQueueSupport(v bool) {
	o.MultiQueueSupport = &v
}

// GetNumInterrupts returns the NumInterrupts field value if set, zero value otherwise.
func (o *VnicVmqSettings) GetNumInterrupts() int64 {
	if o == nil || IsNil(o.NumInterrupts) {
		var ret int64
		return ret
	}
	return *o.NumInterrupts
}

// GetNumInterruptsOk returns a tuple with the NumInterrupts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVmqSettings) GetNumInterruptsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumInterrupts) {
		return nil, false
	}
	return o.NumInterrupts, true
}

// HasNumInterrupts returns a boolean if a field has been set.
func (o *VnicVmqSettings) HasNumInterrupts() bool {
	if o != nil && !IsNil(o.NumInterrupts) {
		return true
	}

	return false
}

// SetNumInterrupts gets a reference to the given int64 and assigns it to the NumInterrupts field.
func (o *VnicVmqSettings) SetNumInterrupts(v int64) {
	o.NumInterrupts = &v
}

// GetNumSubVnics returns the NumSubVnics field value if set, zero value otherwise.
func (o *VnicVmqSettings) GetNumSubVnics() int64 {
	if o == nil || IsNil(o.NumSubVnics) {
		var ret int64
		return ret
	}
	return *o.NumSubVnics
}

// GetNumSubVnicsOk returns a tuple with the NumSubVnics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVmqSettings) GetNumSubVnicsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumSubVnics) {
		return nil, false
	}
	return o.NumSubVnics, true
}

// HasNumSubVnics returns a boolean if a field has been set.
func (o *VnicVmqSettings) HasNumSubVnics() bool {
	if o != nil && !IsNil(o.NumSubVnics) {
		return true
	}

	return false
}

// SetNumSubVnics gets a reference to the given int64 and assigns it to the NumSubVnics field.
func (o *VnicVmqSettings) SetNumSubVnics(v int64) {
	o.NumSubVnics = &v
}

// GetNumVmqs returns the NumVmqs field value if set, zero value otherwise.
func (o *VnicVmqSettings) GetNumVmqs() int64 {
	if o == nil || IsNil(o.NumVmqs) {
		var ret int64
		return ret
	}
	return *o.NumVmqs
}

// GetNumVmqsOk returns a tuple with the NumVmqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVmqSettings) GetNumVmqsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumVmqs) {
		return nil, false
	}
	return o.NumVmqs, true
}

// HasNumVmqs returns a boolean if a field has been set.
func (o *VnicVmqSettings) HasNumVmqs() bool {
	if o != nil && !IsNil(o.NumVmqs) {
		return true
	}

	return false
}

// SetNumVmqs gets a reference to the given int64 and assigns it to the NumVmqs field.
func (o *VnicVmqSettings) SetNumVmqs(v int64) {
	o.NumVmqs = &v
}

// GetVmmqAdapterPolicy returns the VmmqAdapterPolicy field value if set, zero value otherwise.
func (o *VnicVmqSettings) GetVmmqAdapterPolicy() string {
	if o == nil || IsNil(o.VmmqAdapterPolicy) {
		var ret string
		return ret
	}
	return *o.VmmqAdapterPolicy
}

// GetVmmqAdapterPolicyOk returns a tuple with the VmmqAdapterPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVmqSettings) GetVmmqAdapterPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.VmmqAdapterPolicy) {
		return nil, false
	}
	return o.VmmqAdapterPolicy, true
}

// HasVmmqAdapterPolicy returns a boolean if a field has been set.
func (o *VnicVmqSettings) HasVmmqAdapterPolicy() bool {
	if o != nil && !IsNil(o.VmmqAdapterPolicy) {
		return true
	}

	return false
}

// SetVmmqAdapterPolicy gets a reference to the given string and assigns it to the VmmqAdapterPolicy field.
func (o *VnicVmqSettings) SetVmmqAdapterPolicy(v string) {
	o.VmmqAdapterPolicy = &v
}

func (o VnicVmqSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicVmqSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.MultiQueueSupport) {
		toSerialize["MultiQueueSupport"] = o.MultiQueueSupport
	}
	if !IsNil(o.NumInterrupts) {
		toSerialize["NumInterrupts"] = o.NumInterrupts
	}
	if !IsNil(o.NumSubVnics) {
		toSerialize["NumSubVnics"] = o.NumSubVnics
	}
	if !IsNil(o.NumVmqs) {
		toSerialize["NumVmqs"] = o.NumVmqs
	}
	if !IsNil(o.VmmqAdapterPolicy) {
		toSerialize["VmmqAdapterPolicy"] = o.VmmqAdapterPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicVmqSettings) UnmarshalJSON(data []byte) (err error) {
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
	type VnicVmqSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enables VMQ feature on the virtual interface.
		Enabled *bool `json:"Enabled,omitempty"`
		// Enables Virtual Machine Multi-Queue feature on the virtual interface. VMMQ allows configuration of multiple I/O queues for a single VM and thus distributes traffic across multiple CPU cores in a VM.
		MultiQueueSupport *bool `json:"MultiQueueSupport,omitempty"`
		// The number of interrupt resources to be allocated. Recommended value is the number of CPU threads or logical processors available in the server.
		NumInterrupts *int64 `json:"NumInterrupts,omitempty"`
		// The number of sub vNICs to be created.
		NumSubVnics *int64 `json:"NumSubVnics,omitempty"`
		// The number of hardware Virtual Machine Queues to be allocated. The number of VMQs per adapter must be one more than the maximum number of VM NICs.
		NumVmqs *int64 `json:"NumVmqs,omitempty"`
		// Ethernet Adapter policy to be associated with the Sub vNICs. The Transmit Queue and Receive Queue resource value of VMMQ adapter policy should be greater than or equal to the configured number of sub vNICs.
		VmmqAdapterPolicy *string `json:"VmmqAdapterPolicy,omitempty"`
	}

	varVnicVmqSettingsWithoutEmbeddedStruct := VnicVmqSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicVmqSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicVmqSettings := _VnicVmqSettings{}
		varVnicVmqSettings.ClassId = varVnicVmqSettingsWithoutEmbeddedStruct.ClassId
		varVnicVmqSettings.ObjectType = varVnicVmqSettingsWithoutEmbeddedStruct.ObjectType
		varVnicVmqSettings.Enabled = varVnicVmqSettingsWithoutEmbeddedStruct.Enabled
		varVnicVmqSettings.MultiQueueSupport = varVnicVmqSettingsWithoutEmbeddedStruct.MultiQueueSupport
		varVnicVmqSettings.NumInterrupts = varVnicVmqSettingsWithoutEmbeddedStruct.NumInterrupts
		varVnicVmqSettings.NumSubVnics = varVnicVmqSettingsWithoutEmbeddedStruct.NumSubVnics
		varVnicVmqSettings.NumVmqs = varVnicVmqSettingsWithoutEmbeddedStruct.NumVmqs
		varVnicVmqSettings.VmmqAdapterPolicy = varVnicVmqSettingsWithoutEmbeddedStruct.VmmqAdapterPolicy
		*o = VnicVmqSettings(varVnicVmqSettings)
	} else {
		return err
	}

	varVnicVmqSettings := _VnicVmqSettings{}

	err = json.Unmarshal(data, &varVnicVmqSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicVmqSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "MultiQueueSupport")
		delete(additionalProperties, "NumInterrupts")
		delete(additionalProperties, "NumSubVnics")
		delete(additionalProperties, "NumVmqs")
		delete(additionalProperties, "VmmqAdapterPolicy")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableVnicVmqSettings struct {
	value *VnicVmqSettings
	isSet bool
}

func (v NullableVnicVmqSettings) Get() *VnicVmqSettings {
	return v.value
}

func (v *NullableVnicVmqSettings) Set(val *VnicVmqSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicVmqSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicVmqSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicVmqSettings(val *VnicVmqSettings) *NullableVnicVmqSettings {
	return &NullableVnicVmqSettings{value: val, isSet: true}
}

func (v NullableVnicVmqSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicVmqSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
