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

// checks if the KubernetesCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesCluster{}

// KubernetesCluster Inventories a Kubernetes cluster state. A Cluster object is automatically created when a Kubernetes API server is configured for a cluster.
type KubernetesCluster struct {
	ComputeBaseCluster
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of the endpoint connection of this Kubernetes cluster. * `` - The target details have been persisted but Intersight has not yet attempted to connect to the target. * `Connected` - Intersight is able to establish a connection to the target and initiate management activities. * `NotConnected` - Intersight is unable to establish a connection to the target. * `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established. * `UnclaimInProgress` - Unclaim of the target is in progress. Intersight is able to connect to the target and all management operations are supported. * `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight. * `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// Kubeconfig for the cluster to collect inventory for.
	KubeConfig          *string                                           `json:"KubeConfig,omitempty"`
	ClusterAddonProfile NullableKubernetesClusterAddonProfileRelationship `json:"ClusterAddonProfile,omitempty"`
	Organization        NullableOrganizationOrganizationRelationship      `json:"Organization,omitempty"`
	// An array of relationships to assetDeviceRegistration resources.
	RegisteredDevices    []AssetDeviceRegistrationRelationship `json:"RegisteredDevices,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesCluster KubernetesCluster

// NewKubernetesCluster instantiates a new KubernetesCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCluster(classId string, objectType string) *KubernetesCluster {
	this := KubernetesCluster{}
	this.ClassId = classId
	this.ObjectType = objectType
	var connectionStatus string = ""
	this.ConnectionStatus = &connectionStatus
	return &this
}

// NewKubernetesClusterWithDefaults instantiates a new KubernetesCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterWithDefaults() *KubernetesCluster {
	this := KubernetesCluster{}
	var classId string = "kubernetes.Cluster"
	this.ClassId = classId
	var objectType string = "kubernetes.Cluster"
	this.ObjectType = objectType
	var connectionStatus string = ""
	this.ConnectionStatus = &connectionStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesCluster) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesCluster) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.Cluster" of the ClassId field.
func (o *KubernetesCluster) GetDefaultClassId() interface{} {
	return "kubernetes.Cluster"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesCluster) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesCluster) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.Cluster" of the ObjectType field.
func (o *KubernetesCluster) GetDefaultObjectType() interface{} {
	return "kubernetes.Cluster"
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *KubernetesCluster) GetConnectionStatus() string {
	if o == nil || IsNil(o.ConnectionStatus) {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetConnectionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionStatus) {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *KubernetesCluster) HasConnectionStatus() bool {
	if o != nil && !IsNil(o.ConnectionStatus) {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *KubernetesCluster) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetKubeConfig returns the KubeConfig field value if set, zero value otherwise.
func (o *KubernetesCluster) GetKubeConfig() string {
	if o == nil || IsNil(o.KubeConfig) {
		var ret string
		return ret
	}
	return *o.KubeConfig
}

// GetKubeConfigOk returns a tuple with the KubeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetKubeConfigOk() (*string, bool) {
	if o == nil || IsNil(o.KubeConfig) {
		return nil, false
	}
	return o.KubeConfig, true
}

// HasKubeConfig returns a boolean if a field has been set.
func (o *KubernetesCluster) HasKubeConfig() bool {
	if o != nil && !IsNil(o.KubeConfig) {
		return true
	}

	return false
}

// SetKubeConfig gets a reference to the given string and assigns it to the KubeConfig field.
func (o *KubernetesCluster) SetKubeConfig(v string) {
	o.KubeConfig = &v
}

// GetClusterAddonProfile returns the ClusterAddonProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesCluster) GetClusterAddonProfile() KubernetesClusterAddonProfileRelationship {
	if o == nil || IsNil(o.ClusterAddonProfile.Get()) {
		var ret KubernetesClusterAddonProfileRelationship
		return ret
	}
	return *o.ClusterAddonProfile.Get()
}

// GetClusterAddonProfileOk returns a tuple with the ClusterAddonProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCluster) GetClusterAddonProfileOk() (*KubernetesClusterAddonProfileRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterAddonProfile.Get(), o.ClusterAddonProfile.IsSet()
}

// HasClusterAddonProfile returns a boolean if a field has been set.
func (o *KubernetesCluster) HasClusterAddonProfile() bool {
	if o != nil && o.ClusterAddonProfile.IsSet() {
		return true
	}

	return false
}

// SetClusterAddonProfile gets a reference to the given NullableKubernetesClusterAddonProfileRelationship and assigns it to the ClusterAddonProfile field.
func (o *KubernetesCluster) SetClusterAddonProfile(v KubernetesClusterAddonProfileRelationship) {
	o.ClusterAddonProfile.Set(&v)
}

// SetClusterAddonProfileNil sets the value for ClusterAddonProfile to be an explicit nil
func (o *KubernetesCluster) SetClusterAddonProfileNil() {
	o.ClusterAddonProfile.Set(nil)
}

// UnsetClusterAddonProfile ensures that no value is present for ClusterAddonProfile, not even an explicit nil
func (o *KubernetesCluster) UnsetClusterAddonProfile() {
	o.ClusterAddonProfile.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesCluster) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCluster) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesCluster) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesCluster) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *KubernetesCluster) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *KubernetesCluster) UnsetOrganization() {
	o.Organization.Unset()
}

// GetRegisteredDevices returns the RegisteredDevices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesCluster) GetRegisteredDevices() []AssetDeviceRegistrationRelationship {
	if o == nil {
		var ret []AssetDeviceRegistrationRelationship
		return ret
	}
	return o.RegisteredDevices
}

// GetRegisteredDevicesOk returns a tuple with the RegisteredDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCluster) GetRegisteredDevicesOk() ([]AssetDeviceRegistrationRelationship, bool) {
	if o == nil || IsNil(o.RegisteredDevices) {
		return nil, false
	}
	return o.RegisteredDevices, true
}

// HasRegisteredDevices returns a boolean if a field has been set.
func (o *KubernetesCluster) HasRegisteredDevices() bool {
	if o != nil && !IsNil(o.RegisteredDevices) {
		return true
	}

	return false
}

// SetRegisteredDevices gets a reference to the given []AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevices field.
func (o *KubernetesCluster) SetRegisteredDevices(v []AssetDeviceRegistrationRelationship) {
	o.RegisteredDevices = v
}

func (o KubernetesCluster) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedComputeBaseCluster, errComputeBaseCluster := json.Marshal(o.ComputeBaseCluster)
	if errComputeBaseCluster != nil {
		return map[string]interface{}{}, errComputeBaseCluster
	}
	errComputeBaseCluster = json.Unmarshal([]byte(serializedComputeBaseCluster), &toSerialize)
	if errComputeBaseCluster != nil {
		return map[string]interface{}{}, errComputeBaseCluster
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ConnectionStatus) {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if !IsNil(o.KubeConfig) {
		toSerialize["KubeConfig"] = o.KubeConfig
	}
	if o.ClusterAddonProfile.IsSet() {
		toSerialize["ClusterAddonProfile"] = o.ClusterAddonProfile.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.RegisteredDevices != nil {
		toSerialize["RegisteredDevices"] = o.RegisteredDevices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesCluster) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesClusterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Status of the endpoint connection of this Kubernetes cluster. * `` - The target details have been persisted but Intersight has not yet attempted to connect to the target. * `Connected` - Intersight is able to establish a connection to the target and initiate management activities. * `NotConnected` - Intersight is unable to establish a connection to the target. * `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established. * `UnclaimInProgress` - Unclaim of the target is in progress. Intersight is able to connect to the target and all management operations are supported. * `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight. * `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.
		ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
		// Kubeconfig for the cluster to collect inventory for.
		KubeConfig          *string                                           `json:"KubeConfig,omitempty"`
		ClusterAddonProfile NullableKubernetesClusterAddonProfileRelationship `json:"ClusterAddonProfile,omitempty"`
		Organization        NullableOrganizationOrganizationRelationship      `json:"Organization,omitempty"`
		// An array of relationships to assetDeviceRegistration resources.
		RegisteredDevices []AssetDeviceRegistrationRelationship `json:"RegisteredDevices,omitempty"`
	}

	varKubernetesClusterWithoutEmbeddedStruct := KubernetesClusterWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesClusterWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesCluster := _KubernetesCluster{}
		varKubernetesCluster.ClassId = varKubernetesClusterWithoutEmbeddedStruct.ClassId
		varKubernetesCluster.ObjectType = varKubernetesClusterWithoutEmbeddedStruct.ObjectType
		varKubernetesCluster.ConnectionStatus = varKubernetesClusterWithoutEmbeddedStruct.ConnectionStatus
		varKubernetesCluster.KubeConfig = varKubernetesClusterWithoutEmbeddedStruct.KubeConfig
		varKubernetesCluster.ClusterAddonProfile = varKubernetesClusterWithoutEmbeddedStruct.ClusterAddonProfile
		varKubernetesCluster.Organization = varKubernetesClusterWithoutEmbeddedStruct.Organization
		varKubernetesCluster.RegisteredDevices = varKubernetesClusterWithoutEmbeddedStruct.RegisteredDevices
		*o = KubernetesCluster(varKubernetesCluster)
	} else {
		return err
	}

	varKubernetesCluster := _KubernetesCluster{}

	err = json.Unmarshal(data, &varKubernetesCluster)
	if err == nil {
		o.ComputeBaseCluster = varKubernetesCluster.ComputeBaseCluster
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "KubeConfig")
		delete(additionalProperties, "ClusterAddonProfile")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "RegisteredDevices")

		// remove fields from embedded structs
		reflectComputeBaseCluster := reflect.ValueOf(o.ComputeBaseCluster)
		for i := 0; i < reflectComputeBaseCluster.Type().NumField(); i++ {
			t := reflectComputeBaseCluster.Type().Field(i)

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

type NullableKubernetesCluster struct {
	value *KubernetesCluster
	isSet bool
}

func (v NullableKubernetesCluster) Get() *KubernetesCluster {
	return v.value
}

func (v *NullableKubernetesCluster) Set(val *KubernetesCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCluster(val *KubernetesCluster) *NullableKubernetesCluster {
	return &NullableKubernetesCluster{value: val, isSet: true}
}

func (v NullableKubernetesCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
