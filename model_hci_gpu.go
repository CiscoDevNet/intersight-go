/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2024112619
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

// checks if the HciGpu type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HciGpu{}

// HciGpu A Gpu associated with a node.
type HciGpu struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The unique identifier of the cluster.
	ClusterExtId *string `json:"ClusterExtId,omitempty"`
	// The name of the cluster this GPU belongs to.
	ClusterName *string `json:"ClusterName,omitempty"`
	// The unique identifier of the controller VM.
	ControllerVmId *string `json:"ControllerVmId,omitempty"`
	// The unique identifier of the gpu.
	GpuExtId *string `json:"GpuExtId,omitempty"`
	// The unique identifier of the node.
	NodeExtId *string `json:"NodeExtId,omitempty"`
	// The number of vGPUs allocated.
	NumberOfVgpusAllocated *int64                                      `json:"NumberOfVgpusAllocated,omitempty"`
	VirtualGpuConfig       NullableHciVirtualGpuConfig                 `json:"VirtualGpuConfig,omitempty"`
	Node                   NullableHciNodeRelationship                 `json:"Node,omitempty"`
	RegisteredDevice       NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _HciGpu HciGpu

// NewHciGpu instantiates a new HciGpu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHciGpu(classId string, objectType string) *HciGpu {
	this := HciGpu{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHciGpuWithDefaults instantiates a new HciGpu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHciGpuWithDefaults() *HciGpu {
	this := HciGpu{}
	var classId string = "hci.Gpu"
	this.ClassId = classId
	var objectType string = "hci.Gpu"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HciGpu) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HciGpu) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HciGpu) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hci.Gpu" of the ClassId field.
func (o *HciGpu) GetDefaultClassId() interface{} {
	return "hci.Gpu"
}

// GetObjectType returns the ObjectType field value
func (o *HciGpu) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HciGpu) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HciGpu) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hci.Gpu" of the ObjectType field.
func (o *HciGpu) GetDefaultObjectType() interface{} {
	return "hci.Gpu"
}

// GetClusterExtId returns the ClusterExtId field value if set, zero value otherwise.
func (o *HciGpu) GetClusterExtId() string {
	if o == nil || IsNil(o.ClusterExtId) {
		var ret string
		return ret
	}
	return *o.ClusterExtId
}

// GetClusterExtIdOk returns a tuple with the ClusterExtId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciGpu) GetClusterExtIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterExtId) {
		return nil, false
	}
	return o.ClusterExtId, true
}

// HasClusterExtId returns a boolean if a field has been set.
func (o *HciGpu) HasClusterExtId() bool {
	if o != nil && !IsNil(o.ClusterExtId) {
		return true
	}

	return false
}

// SetClusterExtId gets a reference to the given string and assigns it to the ClusterExtId field.
func (o *HciGpu) SetClusterExtId(v string) {
	o.ClusterExtId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *HciGpu) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciGpu) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *HciGpu) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *HciGpu) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetControllerVmId returns the ControllerVmId field value if set, zero value otherwise.
func (o *HciGpu) GetControllerVmId() string {
	if o == nil || IsNil(o.ControllerVmId) {
		var ret string
		return ret
	}
	return *o.ControllerVmId
}

// GetControllerVmIdOk returns a tuple with the ControllerVmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciGpu) GetControllerVmIdOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerVmId) {
		return nil, false
	}
	return o.ControllerVmId, true
}

// HasControllerVmId returns a boolean if a field has been set.
func (o *HciGpu) HasControllerVmId() bool {
	if o != nil && !IsNil(o.ControllerVmId) {
		return true
	}

	return false
}

// SetControllerVmId gets a reference to the given string and assigns it to the ControllerVmId field.
func (o *HciGpu) SetControllerVmId(v string) {
	o.ControllerVmId = &v
}

// GetGpuExtId returns the GpuExtId field value if set, zero value otherwise.
func (o *HciGpu) GetGpuExtId() string {
	if o == nil || IsNil(o.GpuExtId) {
		var ret string
		return ret
	}
	return *o.GpuExtId
}

// GetGpuExtIdOk returns a tuple with the GpuExtId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciGpu) GetGpuExtIdOk() (*string, bool) {
	if o == nil || IsNil(o.GpuExtId) {
		return nil, false
	}
	return o.GpuExtId, true
}

// HasGpuExtId returns a boolean if a field has been set.
func (o *HciGpu) HasGpuExtId() bool {
	if o != nil && !IsNil(o.GpuExtId) {
		return true
	}

	return false
}

// SetGpuExtId gets a reference to the given string and assigns it to the GpuExtId field.
func (o *HciGpu) SetGpuExtId(v string) {
	o.GpuExtId = &v
}

// GetNodeExtId returns the NodeExtId field value if set, zero value otherwise.
func (o *HciGpu) GetNodeExtId() string {
	if o == nil || IsNil(o.NodeExtId) {
		var ret string
		return ret
	}
	return *o.NodeExtId
}

// GetNodeExtIdOk returns a tuple with the NodeExtId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciGpu) GetNodeExtIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeExtId) {
		return nil, false
	}
	return o.NodeExtId, true
}

// HasNodeExtId returns a boolean if a field has been set.
func (o *HciGpu) HasNodeExtId() bool {
	if o != nil && !IsNil(o.NodeExtId) {
		return true
	}

	return false
}

// SetNodeExtId gets a reference to the given string and assigns it to the NodeExtId field.
func (o *HciGpu) SetNodeExtId(v string) {
	o.NodeExtId = &v
}

// GetNumberOfVgpusAllocated returns the NumberOfVgpusAllocated field value if set, zero value otherwise.
func (o *HciGpu) GetNumberOfVgpusAllocated() int64 {
	if o == nil || IsNil(o.NumberOfVgpusAllocated) {
		var ret int64
		return ret
	}
	return *o.NumberOfVgpusAllocated
}

// GetNumberOfVgpusAllocatedOk returns a tuple with the NumberOfVgpusAllocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciGpu) GetNumberOfVgpusAllocatedOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfVgpusAllocated) {
		return nil, false
	}
	return o.NumberOfVgpusAllocated, true
}

// HasNumberOfVgpusAllocated returns a boolean if a field has been set.
func (o *HciGpu) HasNumberOfVgpusAllocated() bool {
	if o != nil && !IsNil(o.NumberOfVgpusAllocated) {
		return true
	}

	return false
}

// SetNumberOfVgpusAllocated gets a reference to the given int64 and assigns it to the NumberOfVgpusAllocated field.
func (o *HciGpu) SetNumberOfVgpusAllocated(v int64) {
	o.NumberOfVgpusAllocated = &v
}

// GetVirtualGpuConfig returns the VirtualGpuConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HciGpu) GetVirtualGpuConfig() HciVirtualGpuConfig {
	if o == nil || IsNil(o.VirtualGpuConfig.Get()) {
		var ret HciVirtualGpuConfig
		return ret
	}
	return *o.VirtualGpuConfig.Get()
}

// GetVirtualGpuConfigOk returns a tuple with the VirtualGpuConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HciGpu) GetVirtualGpuConfigOk() (*HciVirtualGpuConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualGpuConfig.Get(), o.VirtualGpuConfig.IsSet()
}

// HasVirtualGpuConfig returns a boolean if a field has been set.
func (o *HciGpu) HasVirtualGpuConfig() bool {
	if o != nil && o.VirtualGpuConfig.IsSet() {
		return true
	}

	return false
}

// SetVirtualGpuConfig gets a reference to the given NullableHciVirtualGpuConfig and assigns it to the VirtualGpuConfig field.
func (o *HciGpu) SetVirtualGpuConfig(v HciVirtualGpuConfig) {
	o.VirtualGpuConfig.Set(&v)
}

// SetVirtualGpuConfigNil sets the value for VirtualGpuConfig to be an explicit nil
func (o *HciGpu) SetVirtualGpuConfigNil() {
	o.VirtualGpuConfig.Set(nil)
}

// UnsetVirtualGpuConfig ensures that no value is present for VirtualGpuConfig, not even an explicit nil
func (o *HciGpu) UnsetVirtualGpuConfig() {
	o.VirtualGpuConfig.Unset()
}

// GetNode returns the Node field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HciGpu) GetNode() HciNodeRelationship {
	if o == nil || IsNil(o.Node.Get()) {
		var ret HciNodeRelationship
		return ret
	}
	return *o.Node.Get()
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HciGpu) GetNodeOk() (*HciNodeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Node.Get(), o.Node.IsSet()
}

// HasNode returns a boolean if a field has been set.
func (o *HciGpu) HasNode() bool {
	if o != nil && o.Node.IsSet() {
		return true
	}

	return false
}

// SetNode gets a reference to the given NullableHciNodeRelationship and assigns it to the Node field.
func (o *HciGpu) SetNode(v HciNodeRelationship) {
	o.Node.Set(&v)
}

// SetNodeNil sets the value for Node to be an explicit nil
func (o *HciGpu) SetNodeNil() {
	o.Node.Set(nil)
}

// UnsetNode ensures that no value is present for Node, not even an explicit nil
func (o *HciGpu) UnsetNode() {
	o.Node.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HciGpu) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HciGpu) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HciGpu) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HciGpu) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *HciGpu) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *HciGpu) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o HciGpu) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HciGpu) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ClusterExtId) {
		toSerialize["ClusterExtId"] = o.ClusterExtId
	}
	if !IsNil(o.ClusterName) {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if !IsNil(o.ControllerVmId) {
		toSerialize["ControllerVmId"] = o.ControllerVmId
	}
	if !IsNil(o.GpuExtId) {
		toSerialize["GpuExtId"] = o.GpuExtId
	}
	if !IsNil(o.NodeExtId) {
		toSerialize["NodeExtId"] = o.NodeExtId
	}
	if !IsNil(o.NumberOfVgpusAllocated) {
		toSerialize["NumberOfVgpusAllocated"] = o.NumberOfVgpusAllocated
	}
	if o.VirtualGpuConfig.IsSet() {
		toSerialize["VirtualGpuConfig"] = o.VirtualGpuConfig.Get()
	}
	if o.Node.IsSet() {
		toSerialize["Node"] = o.Node.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HciGpu) UnmarshalJSON(data []byte) (err error) {
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
	type HciGpuWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The unique identifier of the cluster.
		ClusterExtId *string `json:"ClusterExtId,omitempty"`
		// The name of the cluster this GPU belongs to.
		ClusterName *string `json:"ClusterName,omitempty"`
		// The unique identifier of the controller VM.
		ControllerVmId *string `json:"ControllerVmId,omitempty"`
		// The unique identifier of the gpu.
		GpuExtId *string `json:"GpuExtId,omitempty"`
		// The unique identifier of the node.
		NodeExtId *string `json:"NodeExtId,omitempty"`
		// The number of vGPUs allocated.
		NumberOfVgpusAllocated *int64                                      `json:"NumberOfVgpusAllocated,omitempty"`
		VirtualGpuConfig       NullableHciVirtualGpuConfig                 `json:"VirtualGpuConfig,omitempty"`
		Node                   NullableHciNodeRelationship                 `json:"Node,omitempty"`
		RegisteredDevice       NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varHciGpuWithoutEmbeddedStruct := HciGpuWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHciGpuWithoutEmbeddedStruct)
	if err == nil {
		varHciGpu := _HciGpu{}
		varHciGpu.ClassId = varHciGpuWithoutEmbeddedStruct.ClassId
		varHciGpu.ObjectType = varHciGpuWithoutEmbeddedStruct.ObjectType
		varHciGpu.ClusterExtId = varHciGpuWithoutEmbeddedStruct.ClusterExtId
		varHciGpu.ClusterName = varHciGpuWithoutEmbeddedStruct.ClusterName
		varHciGpu.ControllerVmId = varHciGpuWithoutEmbeddedStruct.ControllerVmId
		varHciGpu.GpuExtId = varHciGpuWithoutEmbeddedStruct.GpuExtId
		varHciGpu.NodeExtId = varHciGpuWithoutEmbeddedStruct.NodeExtId
		varHciGpu.NumberOfVgpusAllocated = varHciGpuWithoutEmbeddedStruct.NumberOfVgpusAllocated
		varHciGpu.VirtualGpuConfig = varHciGpuWithoutEmbeddedStruct.VirtualGpuConfig
		varHciGpu.Node = varHciGpuWithoutEmbeddedStruct.Node
		varHciGpu.RegisteredDevice = varHciGpuWithoutEmbeddedStruct.RegisteredDevice
		*o = HciGpu(varHciGpu)
	} else {
		return err
	}

	varHciGpu := _HciGpu{}

	err = json.Unmarshal(data, &varHciGpu)
	if err == nil {
		o.MoBaseMo = varHciGpu.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterExtId")
		delete(additionalProperties, "ClusterName")
		delete(additionalProperties, "ControllerVmId")
		delete(additionalProperties, "GpuExtId")
		delete(additionalProperties, "NodeExtId")
		delete(additionalProperties, "NumberOfVgpusAllocated")
		delete(additionalProperties, "VirtualGpuConfig")
		delete(additionalProperties, "Node")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableHciGpu struct {
	value *HciGpu
	isSet bool
}

func (v NullableHciGpu) Get() *HciGpu {
	return v.value
}

func (v *NullableHciGpu) Set(val *HciGpu) {
	v.value = val
	v.isSet = true
}

func (v NullableHciGpu) IsSet() bool {
	return v.isSet
}

func (v *NullableHciGpu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHciGpu(val *HciGpu) *NullableHciGpu {
	return &NullableHciGpu{value: val, isSet: true}
}

func (v NullableHciGpu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHciGpu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}