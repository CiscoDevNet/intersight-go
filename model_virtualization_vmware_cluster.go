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

// checks if the VirtualizationVmwareCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationVmwareCluster{}

// VirtualizationVmwareCluster A real cluster of resources within a data center in VMware. A cluster is a convenient grouping of resources such as Host, Datastore, etc.
type VirtualizationVmwareCluster struct {
	VirtualizationBaseCluster
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                    `json:"ObjectType"`
	AttachedResourceTags []VirtualizationVmwareAttachedResourceTag `json:"AttachedResourceTags,omitempty"`
	// CPU over commitment associated with this cluster.
	CpuOverCommitment *int64 `json:"CpuOverCommitment,omitempty"`
	// Count of all datastores associated with this cluster.
	DatastoreCount *int64 `json:"DatastoreCount,omitempty"`
	// Inventory path of the cluster.
	InventoryPath *string `json:"InventoryPath,omitempty"`
	// Every cluster has an option to enable proactive HA in vCenter. Set to true when the vCenter admin has enabled proactive HA for the cluster.
	ProactiveHaEnabled   *bool                                              `json:"ProactiveHaEnabled,omitempty"`
	Datacenter           NullableVirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship        `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareCluster VirtualizationVmwareCluster

// NewVirtualizationVmwareCluster instantiates a new VirtualizationVmwareCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareCluster(classId string, objectType string) *VirtualizationVmwareCluster {
	this := VirtualizationVmwareCluster{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// NewVirtualizationVmwareClusterWithDefaults instantiates a new VirtualizationVmwareCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareClusterWithDefaults() *VirtualizationVmwareCluster {
	this := VirtualizationVmwareCluster{}
	var classId string = "virtualization.VmwareCluster"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareCluster"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareCluster) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareCluster) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "virtualization.VmwareCluster" of the ClassId field.
func (o *VirtualizationVmwareCluster) GetDefaultClassId() interface{} {
	return "virtualization.VmwareCluster"
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareCluster) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareCluster) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "virtualization.VmwareCluster" of the ObjectType field.
func (o *VirtualizationVmwareCluster) GetDefaultObjectType() interface{} {
	return "virtualization.VmwareCluster"
}

// GetAttachedResourceTags returns the AttachedResourceTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareCluster) GetAttachedResourceTags() []VirtualizationVmwareAttachedResourceTag {
	if o == nil {
		var ret []VirtualizationVmwareAttachedResourceTag
		return ret
	}
	return o.AttachedResourceTags
}

// GetAttachedResourceTagsOk returns a tuple with the AttachedResourceTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareCluster) GetAttachedResourceTagsOk() ([]VirtualizationVmwareAttachedResourceTag, bool) {
	if o == nil || IsNil(o.AttachedResourceTags) {
		return nil, false
	}
	return o.AttachedResourceTags, true
}

// HasAttachedResourceTags returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasAttachedResourceTags() bool {
	if o != nil && !IsNil(o.AttachedResourceTags) {
		return true
	}

	return false
}

// SetAttachedResourceTags gets a reference to the given []VirtualizationVmwareAttachedResourceTag and assigns it to the AttachedResourceTags field.
func (o *VirtualizationVmwareCluster) SetAttachedResourceTags(v []VirtualizationVmwareAttachedResourceTag) {
	o.AttachedResourceTags = v
}

// GetCpuOverCommitment returns the CpuOverCommitment field value if set, zero value otherwise.
func (o *VirtualizationVmwareCluster) GetCpuOverCommitment() int64 {
	if o == nil || IsNil(o.CpuOverCommitment) {
		var ret int64
		return ret
	}
	return *o.CpuOverCommitment
}

// GetCpuOverCommitmentOk returns a tuple with the CpuOverCommitment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetCpuOverCommitmentOk() (*int64, bool) {
	if o == nil || IsNil(o.CpuOverCommitment) {
		return nil, false
	}
	return o.CpuOverCommitment, true
}

// HasCpuOverCommitment returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasCpuOverCommitment() bool {
	if o != nil && !IsNil(o.CpuOverCommitment) {
		return true
	}

	return false
}

// SetCpuOverCommitment gets a reference to the given int64 and assigns it to the CpuOverCommitment field.
func (o *VirtualizationVmwareCluster) SetCpuOverCommitment(v int64) {
	o.CpuOverCommitment = &v
}

// GetDatastoreCount returns the DatastoreCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareCluster) GetDatastoreCount() int64 {
	if o == nil || IsNil(o.DatastoreCount) {
		var ret int64
		return ret
	}
	return *o.DatastoreCount
}

// GetDatastoreCountOk returns a tuple with the DatastoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetDatastoreCountOk() (*int64, bool) {
	if o == nil || IsNil(o.DatastoreCount) {
		return nil, false
	}
	return o.DatastoreCount, true
}

// HasDatastoreCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasDatastoreCount() bool {
	if o != nil && !IsNil(o.DatastoreCount) {
		return true
	}

	return false
}

// SetDatastoreCount gets a reference to the given int64 and assigns it to the DatastoreCount field.
func (o *VirtualizationVmwareCluster) SetDatastoreCount(v int64) {
	o.DatastoreCount = &v
}

// GetInventoryPath returns the InventoryPath field value if set, zero value otherwise.
func (o *VirtualizationVmwareCluster) GetInventoryPath() string {
	if o == nil || IsNil(o.InventoryPath) {
		var ret string
		return ret
	}
	return *o.InventoryPath
}

// GetInventoryPathOk returns a tuple with the InventoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetInventoryPathOk() (*string, bool) {
	if o == nil || IsNil(o.InventoryPath) {
		return nil, false
	}
	return o.InventoryPath, true
}

// HasInventoryPath returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasInventoryPath() bool {
	if o != nil && !IsNil(o.InventoryPath) {
		return true
	}

	return false
}

// SetInventoryPath gets a reference to the given string and assigns it to the InventoryPath field.
func (o *VirtualizationVmwareCluster) SetInventoryPath(v string) {
	o.InventoryPath = &v
}

// GetProactiveHaEnabled returns the ProactiveHaEnabled field value if set, zero value otherwise.
func (o *VirtualizationVmwareCluster) GetProactiveHaEnabled() bool {
	if o == nil || IsNil(o.ProactiveHaEnabled) {
		var ret bool
		return ret
	}
	return *o.ProactiveHaEnabled
}

// GetProactiveHaEnabledOk returns a tuple with the ProactiveHaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetProactiveHaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ProactiveHaEnabled) {
		return nil, false
	}
	return o.ProactiveHaEnabled, true
}

// HasProactiveHaEnabled returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasProactiveHaEnabled() bool {
	if o != nil && !IsNil(o.ProactiveHaEnabled) {
		return true
	}

	return false
}

// SetProactiveHaEnabled gets a reference to the given bool and assigns it to the ProactiveHaEnabled field.
func (o *VirtualizationVmwareCluster) SetProactiveHaEnabled(v bool) {
	o.ProactiveHaEnabled = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareCluster) GetDatacenter() VirtualizationVmwareDatacenterRelationship {
	if o == nil || IsNil(o.Datacenter.Get()) {
		var ret VirtualizationVmwareDatacenterRelationship
		return ret
	}
	return *o.Datacenter.Get()
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareCluster) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Datacenter.Get(), o.Datacenter.IsSet()
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasDatacenter() bool {
	if o != nil && o.Datacenter.IsSet() {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given NullableVirtualizationVmwareDatacenterRelationship and assigns it to the Datacenter field.
func (o *VirtualizationVmwareCluster) SetDatacenter(v VirtualizationVmwareDatacenterRelationship) {
	o.Datacenter.Set(&v)
}

// SetDatacenterNil sets the value for Datacenter to be an explicit nil
func (o *VirtualizationVmwareCluster) SetDatacenterNil() {
	o.Datacenter.Set(nil)
}

// UnsetDatacenter ensures that no value is present for Datacenter, not even an explicit nil
func (o *VirtualizationVmwareCluster) UnsetDatacenter() {
	o.Datacenter.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *VirtualizationVmwareCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *VirtualizationVmwareCluster) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *VirtualizationVmwareCluster) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o VirtualizationVmwareCluster) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationVmwareCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseCluster, errVirtualizationBaseCluster := json.Marshal(o.VirtualizationBaseCluster)
	if errVirtualizationBaseCluster != nil {
		return map[string]interface{}{}, errVirtualizationBaseCluster
	}
	errVirtualizationBaseCluster = json.Unmarshal([]byte(serializedVirtualizationBaseCluster), &toSerialize)
	if errVirtualizationBaseCluster != nil {
		return map[string]interface{}{}, errVirtualizationBaseCluster
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.AttachedResourceTags != nil {
		toSerialize["AttachedResourceTags"] = o.AttachedResourceTags
	}
	if !IsNil(o.CpuOverCommitment) {
		toSerialize["CpuOverCommitment"] = o.CpuOverCommitment
	}
	if !IsNil(o.DatastoreCount) {
		toSerialize["DatastoreCount"] = o.DatastoreCount
	}
	if !IsNil(o.InventoryPath) {
		toSerialize["InventoryPath"] = o.InventoryPath
	}
	if !IsNil(o.ProactiveHaEnabled) {
		toSerialize["ProactiveHaEnabled"] = o.ProactiveHaEnabled
	}
	if o.Datacenter.IsSet() {
		toSerialize["Datacenter"] = o.Datacenter.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualizationVmwareCluster) UnmarshalJSON(data []byte) (err error) {
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
	type VirtualizationVmwareClusterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string                                    `json:"ObjectType"`
		AttachedResourceTags []VirtualizationVmwareAttachedResourceTag `json:"AttachedResourceTags,omitempty"`
		// CPU over commitment associated with this cluster.
		CpuOverCommitment *int64 `json:"CpuOverCommitment,omitempty"`
		// Count of all datastores associated with this cluster.
		DatastoreCount *int64 `json:"DatastoreCount,omitempty"`
		// Inventory path of the cluster.
		InventoryPath *string `json:"InventoryPath,omitempty"`
		// Every cluster has an option to enable proactive HA in vCenter. Set to true when the vCenter admin has enabled proactive HA for the cluster.
		ProactiveHaEnabled *bool                                              `json:"ProactiveHaEnabled,omitempty"`
		Datacenter         NullableVirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
		RegisteredDevice   NullableAssetDeviceRegistrationRelationship        `json:"RegisteredDevice,omitempty"`
	}

	varVirtualizationVmwareClusterWithoutEmbeddedStruct := VirtualizationVmwareClusterWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVirtualizationVmwareClusterWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareCluster := _VirtualizationVmwareCluster{}
		varVirtualizationVmwareCluster.ClassId = varVirtualizationVmwareClusterWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareCluster.ObjectType = varVirtualizationVmwareClusterWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareCluster.AttachedResourceTags = varVirtualizationVmwareClusterWithoutEmbeddedStruct.AttachedResourceTags
		varVirtualizationVmwareCluster.CpuOverCommitment = varVirtualizationVmwareClusterWithoutEmbeddedStruct.CpuOverCommitment
		varVirtualizationVmwareCluster.DatastoreCount = varVirtualizationVmwareClusterWithoutEmbeddedStruct.DatastoreCount
		varVirtualizationVmwareCluster.InventoryPath = varVirtualizationVmwareClusterWithoutEmbeddedStruct.InventoryPath
		varVirtualizationVmwareCluster.ProactiveHaEnabled = varVirtualizationVmwareClusterWithoutEmbeddedStruct.ProactiveHaEnabled
		varVirtualizationVmwareCluster.Datacenter = varVirtualizationVmwareClusterWithoutEmbeddedStruct.Datacenter
		varVirtualizationVmwareCluster.RegisteredDevice = varVirtualizationVmwareClusterWithoutEmbeddedStruct.RegisteredDevice
		*o = VirtualizationVmwareCluster(varVirtualizationVmwareCluster)
	} else {
		return err
	}

	varVirtualizationVmwareCluster := _VirtualizationVmwareCluster{}

	err = json.Unmarshal(data, &varVirtualizationVmwareCluster)
	if err == nil {
		o.VirtualizationBaseCluster = varVirtualizationVmwareCluster.VirtualizationBaseCluster
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AttachedResourceTags")
		delete(additionalProperties, "CpuOverCommitment")
		delete(additionalProperties, "DatastoreCount")
		delete(additionalProperties, "InventoryPath")
		delete(additionalProperties, "ProactiveHaEnabled")
		delete(additionalProperties, "Datacenter")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectVirtualizationBaseCluster := reflect.ValueOf(o.VirtualizationBaseCluster)
		for i := 0; i < reflectVirtualizationBaseCluster.Type().NumField(); i++ {
			t := reflectVirtualizationBaseCluster.Type().Field(i)

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

type NullableVirtualizationVmwareCluster struct {
	value *VirtualizationVmwareCluster
	isSet bool
}

func (v NullableVirtualizationVmwareCluster) Get() *VirtualizationVmwareCluster {
	return v.value
}

func (v *NullableVirtualizationVmwareCluster) Set(val *VirtualizationVmwareCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareCluster(val *VirtualizationVmwareCluster) *NullableVirtualizationVmwareCluster {
	return &NullableVirtualizationVmwareCluster{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
