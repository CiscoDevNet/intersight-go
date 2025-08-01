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

// checks if the KubernetesContainerRuntimePolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesContainerRuntimePolicy{}

// KubernetesContainerRuntimePolicy A policy specifying container runtime configuration, such as docker proxy, no proxy and bridge network IP.
type KubernetesContainerRuntimePolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Bridge IP (--bip) including Prefix (e.g., 172.17.0.5/24) that Docker will use for the default bridge network (docker0). Containers will connect to this if no other network is configured, not used by kubernetes pods because their network is managed by CNI. However this address space must not collide with other CIDRs on your networks, including the cluster's Service CIDR, Pod Network CIDR and IP Pools.
	DockerBridgeNetworkCidr *string                       `json:"DockerBridgeNetworkCidr,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\/([0-9]|[1-2][0-9]|3[0-2])$"`
	DockerHttpProxy         NullableKubernetesProxyConfig `json:"DockerHttpProxy,omitempty"`
	DockerHttpsProxy        NullableKubernetesProxyConfig `json:"DockerHttpsProxy,omitempty"`
	DockerNoProxy           []string                      `json:"DockerNoProxy,omitempty"`
	// An array of relationships to kubernetesClusterProfile resources.
	ClusterProfiles      []KubernetesClusterProfileRelationship       `json:"ClusterProfiles,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesContainerRuntimePolicy KubernetesContainerRuntimePolicy

// NewKubernetesContainerRuntimePolicy instantiates a new KubernetesContainerRuntimePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesContainerRuntimePolicy(classId string, objectType string) *KubernetesContainerRuntimePolicy {
	this := KubernetesContainerRuntimePolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesContainerRuntimePolicyWithDefaults instantiates a new KubernetesContainerRuntimePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesContainerRuntimePolicyWithDefaults() *KubernetesContainerRuntimePolicy {
	this := KubernetesContainerRuntimePolicy{}
	var classId string = "kubernetes.ContainerRuntimePolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.ContainerRuntimePolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesContainerRuntimePolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainerRuntimePolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesContainerRuntimePolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.ContainerRuntimePolicy" of the ClassId field.
func (o *KubernetesContainerRuntimePolicy) GetDefaultClassId() interface{} {
	return "kubernetes.ContainerRuntimePolicy"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesContainerRuntimePolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainerRuntimePolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesContainerRuntimePolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.ContainerRuntimePolicy" of the ObjectType field.
func (o *KubernetesContainerRuntimePolicy) GetDefaultObjectType() interface{} {
	return "kubernetes.ContainerRuntimePolicy"
}

// GetDockerBridgeNetworkCidr returns the DockerBridgeNetworkCidr field value if set, zero value otherwise.
func (o *KubernetesContainerRuntimePolicy) GetDockerBridgeNetworkCidr() string {
	if o == nil || IsNil(o.DockerBridgeNetworkCidr) {
		var ret string
		return ret
	}
	return *o.DockerBridgeNetworkCidr
}

// GetDockerBridgeNetworkCidrOk returns a tuple with the DockerBridgeNetworkCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesContainerRuntimePolicy) GetDockerBridgeNetworkCidrOk() (*string, bool) {
	if o == nil || IsNil(o.DockerBridgeNetworkCidr) {
		return nil, false
	}
	return o.DockerBridgeNetworkCidr, true
}

// HasDockerBridgeNetworkCidr returns a boolean if a field has been set.
func (o *KubernetesContainerRuntimePolicy) HasDockerBridgeNetworkCidr() bool {
	if o != nil && !IsNil(o.DockerBridgeNetworkCidr) {
		return true
	}

	return false
}

// SetDockerBridgeNetworkCidr gets a reference to the given string and assigns it to the DockerBridgeNetworkCidr field.
func (o *KubernetesContainerRuntimePolicy) SetDockerBridgeNetworkCidr(v string) {
	o.DockerBridgeNetworkCidr = &v
}

// GetDockerHttpProxy returns the DockerHttpProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesContainerRuntimePolicy) GetDockerHttpProxy() KubernetesProxyConfig {
	if o == nil || IsNil(o.DockerHttpProxy.Get()) {
		var ret KubernetesProxyConfig
		return ret
	}
	return *o.DockerHttpProxy.Get()
}

// GetDockerHttpProxyOk returns a tuple with the DockerHttpProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesContainerRuntimePolicy) GetDockerHttpProxyOk() (*KubernetesProxyConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.DockerHttpProxy.Get(), o.DockerHttpProxy.IsSet()
}

// HasDockerHttpProxy returns a boolean if a field has been set.
func (o *KubernetesContainerRuntimePolicy) HasDockerHttpProxy() bool {
	if o != nil && o.DockerHttpProxy.IsSet() {
		return true
	}

	return false
}

// SetDockerHttpProxy gets a reference to the given NullableKubernetesProxyConfig and assigns it to the DockerHttpProxy field.
func (o *KubernetesContainerRuntimePolicy) SetDockerHttpProxy(v KubernetesProxyConfig) {
	o.DockerHttpProxy.Set(&v)
}

// SetDockerHttpProxyNil sets the value for DockerHttpProxy to be an explicit nil
func (o *KubernetesContainerRuntimePolicy) SetDockerHttpProxyNil() {
	o.DockerHttpProxy.Set(nil)
}

// UnsetDockerHttpProxy ensures that no value is present for DockerHttpProxy, not even an explicit nil
func (o *KubernetesContainerRuntimePolicy) UnsetDockerHttpProxy() {
	o.DockerHttpProxy.Unset()
}

// GetDockerHttpsProxy returns the DockerHttpsProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesContainerRuntimePolicy) GetDockerHttpsProxy() KubernetesProxyConfig {
	if o == nil || IsNil(o.DockerHttpsProxy.Get()) {
		var ret KubernetesProxyConfig
		return ret
	}
	return *o.DockerHttpsProxy.Get()
}

// GetDockerHttpsProxyOk returns a tuple with the DockerHttpsProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesContainerRuntimePolicy) GetDockerHttpsProxyOk() (*KubernetesProxyConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.DockerHttpsProxy.Get(), o.DockerHttpsProxy.IsSet()
}

// HasDockerHttpsProxy returns a boolean if a field has been set.
func (o *KubernetesContainerRuntimePolicy) HasDockerHttpsProxy() bool {
	if o != nil && o.DockerHttpsProxy.IsSet() {
		return true
	}

	return false
}

// SetDockerHttpsProxy gets a reference to the given NullableKubernetesProxyConfig and assigns it to the DockerHttpsProxy field.
func (o *KubernetesContainerRuntimePolicy) SetDockerHttpsProxy(v KubernetesProxyConfig) {
	o.DockerHttpsProxy.Set(&v)
}

// SetDockerHttpsProxyNil sets the value for DockerHttpsProxy to be an explicit nil
func (o *KubernetesContainerRuntimePolicy) SetDockerHttpsProxyNil() {
	o.DockerHttpsProxy.Set(nil)
}

// UnsetDockerHttpsProxy ensures that no value is present for DockerHttpsProxy, not even an explicit nil
func (o *KubernetesContainerRuntimePolicy) UnsetDockerHttpsProxy() {
	o.DockerHttpsProxy.Unset()
}

// GetDockerNoProxy returns the DockerNoProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesContainerRuntimePolicy) GetDockerNoProxy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DockerNoProxy
}

// GetDockerNoProxyOk returns a tuple with the DockerNoProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesContainerRuntimePolicy) GetDockerNoProxyOk() ([]string, bool) {
	if o == nil || IsNil(o.DockerNoProxy) {
		return nil, false
	}
	return o.DockerNoProxy, true
}

// HasDockerNoProxy returns a boolean if a field has been set.
func (o *KubernetesContainerRuntimePolicy) HasDockerNoProxy() bool {
	if o != nil && !IsNil(o.DockerNoProxy) {
		return true
	}

	return false
}

// SetDockerNoProxy gets a reference to the given []string and assigns it to the DockerNoProxy field.
func (o *KubernetesContainerRuntimePolicy) SetDockerNoProxy(v []string) {
	o.DockerNoProxy = v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesContainerRuntimePolicy) GetClusterProfiles() []KubernetesClusterProfileRelationship {
	if o == nil {
		var ret []KubernetesClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesContainerRuntimePolicy) GetClusterProfilesOk() ([]KubernetesClusterProfileRelationship, bool) {
	if o == nil || IsNil(o.ClusterProfiles) {
		return nil, false
	}
	return o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *KubernetesContainerRuntimePolicy) HasClusterProfiles() bool {
	if o != nil && !IsNil(o.ClusterProfiles) {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []KubernetesClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *KubernetesContainerRuntimePolicy) SetClusterProfiles(v []KubernetesClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesContainerRuntimePolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesContainerRuntimePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesContainerRuntimePolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesContainerRuntimePolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *KubernetesContainerRuntimePolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *KubernetesContainerRuntimePolicy) UnsetOrganization() {
	o.Organization.Unset()
}

func (o KubernetesContainerRuntimePolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesContainerRuntimePolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DockerBridgeNetworkCidr) {
		toSerialize["DockerBridgeNetworkCidr"] = o.DockerBridgeNetworkCidr
	}
	if o.DockerHttpProxy.IsSet() {
		toSerialize["DockerHttpProxy"] = o.DockerHttpProxy.Get()
	}
	if o.DockerHttpsProxy.IsSet() {
		toSerialize["DockerHttpsProxy"] = o.DockerHttpsProxy.Get()
	}
	if o.DockerNoProxy != nil {
		toSerialize["DockerNoProxy"] = o.DockerNoProxy
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesContainerRuntimePolicy) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesContainerRuntimePolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Bridge IP (--bip) including Prefix (e.g., 172.17.0.5/24) that Docker will use for the default bridge network (docker0). Containers will connect to this if no other network is configured, not used by kubernetes pods because their network is managed by CNI. However this address space must not collide with other CIDRs on your networks, including the cluster's Service CIDR, Pod Network CIDR and IP Pools.
		DockerBridgeNetworkCidr *string                       `json:"DockerBridgeNetworkCidr,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\/([0-9]|[1-2][0-9]|3[0-2])$"`
		DockerHttpProxy         NullableKubernetesProxyConfig `json:"DockerHttpProxy,omitempty"`
		DockerHttpsProxy        NullableKubernetesProxyConfig `json:"DockerHttpsProxy,omitempty"`
		DockerNoProxy           []string                      `json:"DockerNoProxy,omitempty"`
		// An array of relationships to kubernetesClusterProfile resources.
		ClusterProfiles []KubernetesClusterProfileRelationship       `json:"ClusterProfiles,omitempty"`
		Organization    NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varKubernetesContainerRuntimePolicyWithoutEmbeddedStruct := KubernetesContainerRuntimePolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesContainerRuntimePolicyWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesContainerRuntimePolicy := _KubernetesContainerRuntimePolicy{}
		varKubernetesContainerRuntimePolicy.ClassId = varKubernetesContainerRuntimePolicyWithoutEmbeddedStruct.ClassId
		varKubernetesContainerRuntimePolicy.ObjectType = varKubernetesContainerRuntimePolicyWithoutEmbeddedStruct.ObjectType
		varKubernetesContainerRuntimePolicy.DockerBridgeNetworkCidr = varKubernetesContainerRuntimePolicyWithoutEmbeddedStruct.DockerBridgeNetworkCidr
		varKubernetesContainerRuntimePolicy.DockerHttpProxy = varKubernetesContainerRuntimePolicyWithoutEmbeddedStruct.DockerHttpProxy
		varKubernetesContainerRuntimePolicy.DockerHttpsProxy = varKubernetesContainerRuntimePolicyWithoutEmbeddedStruct.DockerHttpsProxy
		varKubernetesContainerRuntimePolicy.DockerNoProxy = varKubernetesContainerRuntimePolicyWithoutEmbeddedStruct.DockerNoProxy
		varKubernetesContainerRuntimePolicy.ClusterProfiles = varKubernetesContainerRuntimePolicyWithoutEmbeddedStruct.ClusterProfiles
		varKubernetesContainerRuntimePolicy.Organization = varKubernetesContainerRuntimePolicyWithoutEmbeddedStruct.Organization
		*o = KubernetesContainerRuntimePolicy(varKubernetesContainerRuntimePolicy)
	} else {
		return err
	}

	varKubernetesContainerRuntimePolicy := _KubernetesContainerRuntimePolicy{}

	err = json.Unmarshal(data, &varKubernetesContainerRuntimePolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varKubernetesContainerRuntimePolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DockerBridgeNetworkCidr")
		delete(additionalProperties, "DockerHttpProxy")
		delete(additionalProperties, "DockerHttpsProxy")
		delete(additionalProperties, "DockerNoProxy")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableKubernetesContainerRuntimePolicy struct {
	value *KubernetesContainerRuntimePolicy
	isSet bool
}

func (v NullableKubernetesContainerRuntimePolicy) Get() *KubernetesContainerRuntimePolicy {
	return v.value
}

func (v *NullableKubernetesContainerRuntimePolicy) Set(val *KubernetesContainerRuntimePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesContainerRuntimePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesContainerRuntimePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesContainerRuntimePolicy(val *KubernetesContainerRuntimePolicy) *NullableKubernetesContainerRuntimePolicy {
	return &NullableKubernetesContainerRuntimePolicy{value: val, isSet: true}
}

func (v NullableKubernetesContainerRuntimePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesContainerRuntimePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
