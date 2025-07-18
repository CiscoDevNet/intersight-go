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

// checks if the AssetDeviceStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetDeviceStatistics{}

// AssetDeviceStatistics Type for saving device statistics related information for HyperFlex and UCS devices. It is used in asset.DeploymentDevice object to save the current device statistics.
type AssetDeviceStatistics struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Deployment type of HyperFlex cluster.
	ClusterDeploymentType *string `json:"ClusterDeploymentType,omitempty"`
	// Reference to HyperFlex cluster target device ID.
	ClusterDeviceMoid *string `json:"ClusterDeviceMoid,omitempty"`
	// Name of the cluster. It is specified only for HyperFlex based devices.
	ClusterName *string `json:"ClusterName,omitempty"`
	// Data replication factor of HyperFlex cluster.
	ClusterReplicationFactor *int64 `json:"ClusterReplicationFactor,omitempty"`
	// The status of the persistent connection between the device connector and Intersight, for HyperFlex or UCS device. 1 represents being connected and 0 represents being disconnected.
	Connected *int64 `json:"Connected,omitempty"`
	// Defines the average proportion of resources used by the device within the cluster. example in a cluster having 3 nodes, the membershipRatio of each node is 1/3 or 0.33. It is specified only for HyperFlex based devices.
	MembershipRatio *float32 `json:"MembershipRatio,omitempty"`
	// Memory Reliability, availability and serviceability (RAS) factor.
	MemoryMirroringFactor *float32            `json:"MemoryMirroringFactor,omitempty"`
	VmHost                NullableAssetVmHost `json:"VmHost,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _AssetDeviceStatistics AssetDeviceStatistics

// NewAssetDeviceStatistics instantiates a new AssetDeviceStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceStatistics(classId string, objectType string) *AssetDeviceStatistics {
	this := AssetDeviceStatistics{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetDeviceStatisticsWithDefaults instantiates a new AssetDeviceStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceStatisticsWithDefaults() *AssetDeviceStatistics {
	this := AssetDeviceStatistics{}
	var classId string = "asset.DeviceStatistics"
	this.ClassId = classId
	var objectType string = "asset.DeviceStatistics"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeviceStatistics) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeviceStatistics) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "asset.DeviceStatistics" of the ClassId field.
func (o *AssetDeviceStatistics) GetDefaultClassId() interface{} {
	return "asset.DeviceStatistics"
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeviceStatistics) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeviceStatistics) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "asset.DeviceStatistics" of the ObjectType field.
func (o *AssetDeviceStatistics) GetDefaultObjectType() interface{} {
	return "asset.DeviceStatistics"
}

// GetClusterDeploymentType returns the ClusterDeploymentType field value if set, zero value otherwise.
func (o *AssetDeviceStatistics) GetClusterDeploymentType() string {
	if o == nil || IsNil(o.ClusterDeploymentType) {
		var ret string
		return ret
	}
	return *o.ClusterDeploymentType
}

// GetClusterDeploymentTypeOk returns a tuple with the ClusterDeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetClusterDeploymentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterDeploymentType) {
		return nil, false
	}
	return o.ClusterDeploymentType, true
}

// HasClusterDeploymentType returns a boolean if a field has been set.
func (o *AssetDeviceStatistics) HasClusterDeploymentType() bool {
	if o != nil && !IsNil(o.ClusterDeploymentType) {
		return true
	}

	return false
}

// SetClusterDeploymentType gets a reference to the given string and assigns it to the ClusterDeploymentType field.
func (o *AssetDeviceStatistics) SetClusterDeploymentType(v string) {
	o.ClusterDeploymentType = &v
}

// GetClusterDeviceMoid returns the ClusterDeviceMoid field value if set, zero value otherwise.
func (o *AssetDeviceStatistics) GetClusterDeviceMoid() string {
	if o == nil || IsNil(o.ClusterDeviceMoid) {
		var ret string
		return ret
	}
	return *o.ClusterDeviceMoid
}

// GetClusterDeviceMoidOk returns a tuple with the ClusterDeviceMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetClusterDeviceMoidOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterDeviceMoid) {
		return nil, false
	}
	return o.ClusterDeviceMoid, true
}

// HasClusterDeviceMoid returns a boolean if a field has been set.
func (o *AssetDeviceStatistics) HasClusterDeviceMoid() bool {
	if o != nil && !IsNil(o.ClusterDeviceMoid) {
		return true
	}

	return false
}

// SetClusterDeviceMoid gets a reference to the given string and assigns it to the ClusterDeviceMoid field.
func (o *AssetDeviceStatistics) SetClusterDeviceMoid(v string) {
	o.ClusterDeviceMoid = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *AssetDeviceStatistics) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *AssetDeviceStatistics) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *AssetDeviceStatistics) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterReplicationFactor returns the ClusterReplicationFactor field value if set, zero value otherwise.
func (o *AssetDeviceStatistics) GetClusterReplicationFactor() int64 {
	if o == nil || IsNil(o.ClusterReplicationFactor) {
		var ret int64
		return ret
	}
	return *o.ClusterReplicationFactor
}

// GetClusterReplicationFactorOk returns a tuple with the ClusterReplicationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetClusterReplicationFactorOk() (*int64, bool) {
	if o == nil || IsNil(o.ClusterReplicationFactor) {
		return nil, false
	}
	return o.ClusterReplicationFactor, true
}

// HasClusterReplicationFactor returns a boolean if a field has been set.
func (o *AssetDeviceStatistics) HasClusterReplicationFactor() bool {
	if o != nil && !IsNil(o.ClusterReplicationFactor) {
		return true
	}

	return false
}

// SetClusterReplicationFactor gets a reference to the given int64 and assigns it to the ClusterReplicationFactor field.
func (o *AssetDeviceStatistics) SetClusterReplicationFactor(v int64) {
	o.ClusterReplicationFactor = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *AssetDeviceStatistics) GetConnected() int64 {
	if o == nil || IsNil(o.Connected) {
		var ret int64
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetConnectedOk() (*int64, bool) {
	if o == nil || IsNil(o.Connected) {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *AssetDeviceStatistics) HasConnected() bool {
	if o != nil && !IsNil(o.Connected) {
		return true
	}

	return false
}

// SetConnected gets a reference to the given int64 and assigns it to the Connected field.
func (o *AssetDeviceStatistics) SetConnected(v int64) {
	o.Connected = &v
}

// GetMembershipRatio returns the MembershipRatio field value if set, zero value otherwise.
func (o *AssetDeviceStatistics) GetMembershipRatio() float32 {
	if o == nil || IsNil(o.MembershipRatio) {
		var ret float32
		return ret
	}
	return *o.MembershipRatio
}

// GetMembershipRatioOk returns a tuple with the MembershipRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetMembershipRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.MembershipRatio) {
		return nil, false
	}
	return o.MembershipRatio, true
}

// HasMembershipRatio returns a boolean if a field has been set.
func (o *AssetDeviceStatistics) HasMembershipRatio() bool {
	if o != nil && !IsNil(o.MembershipRatio) {
		return true
	}

	return false
}

// SetMembershipRatio gets a reference to the given float32 and assigns it to the MembershipRatio field.
func (o *AssetDeviceStatistics) SetMembershipRatio(v float32) {
	o.MembershipRatio = &v
}

// GetMemoryMirroringFactor returns the MemoryMirroringFactor field value if set, zero value otherwise.
func (o *AssetDeviceStatistics) GetMemoryMirroringFactor() float32 {
	if o == nil || IsNil(o.MemoryMirroringFactor) {
		var ret float32
		return ret
	}
	return *o.MemoryMirroringFactor
}

// GetMemoryMirroringFactorOk returns a tuple with the MemoryMirroringFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetMemoryMirroringFactorOk() (*float32, bool) {
	if o == nil || IsNil(o.MemoryMirroringFactor) {
		return nil, false
	}
	return o.MemoryMirroringFactor, true
}

// HasMemoryMirroringFactor returns a boolean if a field has been set.
func (o *AssetDeviceStatistics) HasMemoryMirroringFactor() bool {
	if o != nil && !IsNil(o.MemoryMirroringFactor) {
		return true
	}

	return false
}

// SetMemoryMirroringFactor gets a reference to the given float32 and assigns it to the MemoryMirroringFactor field.
func (o *AssetDeviceStatistics) SetMemoryMirroringFactor(v float32) {
	o.MemoryMirroringFactor = &v
}

// GetVmHost returns the VmHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceStatistics) GetVmHost() AssetVmHost {
	if o == nil || IsNil(o.VmHost.Get()) {
		var ret AssetVmHost
		return ret
	}
	return *o.VmHost.Get()
}

// GetVmHostOk returns a tuple with the VmHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceStatistics) GetVmHostOk() (*AssetVmHost, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmHost.Get(), o.VmHost.IsSet()
}

// HasVmHost returns a boolean if a field has been set.
func (o *AssetDeviceStatistics) HasVmHost() bool {
	if o != nil && o.VmHost.IsSet() {
		return true
	}

	return false
}

// SetVmHost gets a reference to the given NullableAssetVmHost and assigns it to the VmHost field.
func (o *AssetDeviceStatistics) SetVmHost(v AssetVmHost) {
	o.VmHost.Set(&v)
}

// SetVmHostNil sets the value for VmHost to be an explicit nil
func (o *AssetDeviceStatistics) SetVmHostNil() {
	o.VmHost.Set(nil)
}

// UnsetVmHost ensures that no value is present for VmHost, not even an explicit nil
func (o *AssetDeviceStatistics) UnsetVmHost() {
	o.VmHost.Unset()
}

func (o AssetDeviceStatistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetDeviceStatistics) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ClusterDeploymentType) {
		toSerialize["ClusterDeploymentType"] = o.ClusterDeploymentType
	}
	if !IsNil(o.ClusterDeviceMoid) {
		toSerialize["ClusterDeviceMoid"] = o.ClusterDeviceMoid
	}
	if !IsNil(o.ClusterName) {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if !IsNil(o.ClusterReplicationFactor) {
		toSerialize["ClusterReplicationFactor"] = o.ClusterReplicationFactor
	}
	if !IsNil(o.Connected) {
		toSerialize["Connected"] = o.Connected
	}
	if !IsNil(o.MembershipRatio) {
		toSerialize["MembershipRatio"] = o.MembershipRatio
	}
	if !IsNil(o.MemoryMirroringFactor) {
		toSerialize["MemoryMirroringFactor"] = o.MemoryMirroringFactor
	}
	if o.VmHost.IsSet() {
		toSerialize["VmHost"] = o.VmHost.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetDeviceStatistics) UnmarshalJSON(data []byte) (err error) {
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
	type AssetDeviceStatisticsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Deployment type of HyperFlex cluster.
		ClusterDeploymentType *string `json:"ClusterDeploymentType,omitempty"`
		// Reference to HyperFlex cluster target device ID.
		ClusterDeviceMoid *string `json:"ClusterDeviceMoid,omitempty"`
		// Name of the cluster. It is specified only for HyperFlex based devices.
		ClusterName *string `json:"ClusterName,omitempty"`
		// Data replication factor of HyperFlex cluster.
		ClusterReplicationFactor *int64 `json:"ClusterReplicationFactor,omitempty"`
		// The status of the persistent connection between the device connector and Intersight, for HyperFlex or UCS device. 1 represents being connected and 0 represents being disconnected.
		Connected *int64 `json:"Connected,omitempty"`
		// Defines the average proportion of resources used by the device within the cluster. example in a cluster having 3 nodes, the membershipRatio of each node is 1/3 or 0.33. It is specified only for HyperFlex based devices.
		MembershipRatio *float32 `json:"MembershipRatio,omitempty"`
		// Memory Reliability, availability and serviceability (RAS) factor.
		MemoryMirroringFactor *float32            `json:"MemoryMirroringFactor,omitempty"`
		VmHost                NullableAssetVmHost `json:"VmHost,omitempty"`
	}

	varAssetDeviceStatisticsWithoutEmbeddedStruct := AssetDeviceStatisticsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetDeviceStatisticsWithoutEmbeddedStruct)
	if err == nil {
		varAssetDeviceStatistics := _AssetDeviceStatistics{}
		varAssetDeviceStatistics.ClassId = varAssetDeviceStatisticsWithoutEmbeddedStruct.ClassId
		varAssetDeviceStatistics.ObjectType = varAssetDeviceStatisticsWithoutEmbeddedStruct.ObjectType
		varAssetDeviceStatistics.ClusterDeploymentType = varAssetDeviceStatisticsWithoutEmbeddedStruct.ClusterDeploymentType
		varAssetDeviceStatistics.ClusterDeviceMoid = varAssetDeviceStatisticsWithoutEmbeddedStruct.ClusterDeviceMoid
		varAssetDeviceStatistics.ClusterName = varAssetDeviceStatisticsWithoutEmbeddedStruct.ClusterName
		varAssetDeviceStatistics.ClusterReplicationFactor = varAssetDeviceStatisticsWithoutEmbeddedStruct.ClusterReplicationFactor
		varAssetDeviceStatistics.Connected = varAssetDeviceStatisticsWithoutEmbeddedStruct.Connected
		varAssetDeviceStatistics.MembershipRatio = varAssetDeviceStatisticsWithoutEmbeddedStruct.MembershipRatio
		varAssetDeviceStatistics.MemoryMirroringFactor = varAssetDeviceStatisticsWithoutEmbeddedStruct.MemoryMirroringFactor
		varAssetDeviceStatistics.VmHost = varAssetDeviceStatisticsWithoutEmbeddedStruct.VmHost
		*o = AssetDeviceStatistics(varAssetDeviceStatistics)
	} else {
		return err
	}

	varAssetDeviceStatistics := _AssetDeviceStatistics{}

	err = json.Unmarshal(data, &varAssetDeviceStatistics)
	if err == nil {
		o.MoBaseComplexType = varAssetDeviceStatistics.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterDeploymentType")
		delete(additionalProperties, "ClusterDeviceMoid")
		delete(additionalProperties, "ClusterName")
		delete(additionalProperties, "ClusterReplicationFactor")
		delete(additionalProperties, "Connected")
		delete(additionalProperties, "MembershipRatio")
		delete(additionalProperties, "MemoryMirroringFactor")
		delete(additionalProperties, "VmHost")

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

type NullableAssetDeviceStatistics struct {
	value *AssetDeviceStatistics
	isSet bool
}

func (v NullableAssetDeviceStatistics) Get() *AssetDeviceStatistics {
	return v.value
}

func (v *NullableAssetDeviceStatistics) Set(val *AssetDeviceStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceStatistics(val *AssetDeviceStatistics) *NullableAssetDeviceStatistics {
	return &NullableAssetDeviceStatistics{value: val, isSet: true}
}

func (v NullableAssetDeviceStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
