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

// checks if the VnicSanConnectivityPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicSanConnectivityPolicy{}

// VnicSanConnectivityPolicy SAN connectivity policy determines the network storage resources and the connections between the server and the SAN on the network. This policy enables configuration of vHBAs that the servers use to communicate with the storage network.
type VnicSanConnectivityPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The mode used for placement of vNICs on network adapters. It can either be Auto or Custom. * `custom` - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user. * `auto` - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system.
	PlacementMode *string `json:"PlacementMode,omitempty"`
	// The WWNN address for the server node must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN's in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx.
	StaticWwnnAddress *string `json:"StaticWwnnAddress,omitempty" validate:"regexp=^$|((^20|5[0-9a-fA-F]{1}):([0-9a-fA-F]{2}:){6}([0-9a-fA-F]{2})$)"`
	// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
	TargetPlatform *string `json:"TargetPlatform,omitempty"`
	// Type of allocation selected to assign a WWNN address for the server node. * `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface.
	WwnnAddressType *string `json:"WwnnAddressType,omitempty"`
	// An array of relationships to vnicFcIf resources.
	FcIfs        []VnicFcIfRelationship                       `json:"FcIfs,omitempty"`
	Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	// Deprecated
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	WwnnPool             NullableFcpoolPoolRelationship            `json:"WwnnPool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicSanConnectivityPolicy VnicSanConnectivityPolicy

// NewVnicSanConnectivityPolicy instantiates a new VnicSanConnectivityPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicSanConnectivityPolicy(classId string, objectType string) *VnicSanConnectivityPolicy {
	this := VnicSanConnectivityPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var placementMode string = "custom"
	this.PlacementMode = &placementMode
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	var wwnnAddressType string = "POOL"
	this.WwnnAddressType = &wwnnAddressType
	return &this
}

// NewVnicSanConnectivityPolicyWithDefaults instantiates a new VnicSanConnectivityPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicSanConnectivityPolicyWithDefaults() *VnicSanConnectivityPolicy {
	this := VnicSanConnectivityPolicy{}
	var classId string = "vnic.SanConnectivityPolicy"
	this.ClassId = classId
	var objectType string = "vnic.SanConnectivityPolicy"
	this.ObjectType = objectType
	var placementMode string = "custom"
	this.PlacementMode = &placementMode
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	var wwnnAddressType string = "POOL"
	this.WwnnAddressType = &wwnnAddressType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicSanConnectivityPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicSanConnectivityPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.SanConnectivityPolicy" of the ClassId field.
func (o *VnicSanConnectivityPolicy) GetDefaultClassId() interface{} {
	return "vnic.SanConnectivityPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *VnicSanConnectivityPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicSanConnectivityPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.SanConnectivityPolicy" of the ObjectType field.
func (o *VnicSanConnectivityPolicy) GetDefaultObjectType() interface{} {
	return "vnic.SanConnectivityPolicy"
}

// GetPlacementMode returns the PlacementMode field value if set, zero value otherwise.
func (o *VnicSanConnectivityPolicy) GetPlacementMode() string {
	if o == nil || IsNil(o.PlacementMode) {
		var ret string
		return ret
	}
	return *o.PlacementMode
}

// GetPlacementModeOk returns a tuple with the PlacementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicy) GetPlacementModeOk() (*string, bool) {
	if o == nil || IsNil(o.PlacementMode) {
		return nil, false
	}
	return o.PlacementMode, true
}

// HasPlacementMode returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicy) HasPlacementMode() bool {
	if o != nil && !IsNil(o.PlacementMode) {
		return true
	}

	return false
}

// SetPlacementMode gets a reference to the given string and assigns it to the PlacementMode field.
func (o *VnicSanConnectivityPolicy) SetPlacementMode(v string) {
	o.PlacementMode = &v
}

// GetStaticWwnnAddress returns the StaticWwnnAddress field value if set, zero value otherwise.
func (o *VnicSanConnectivityPolicy) GetStaticWwnnAddress() string {
	if o == nil || IsNil(o.StaticWwnnAddress) {
		var ret string
		return ret
	}
	return *o.StaticWwnnAddress
}

// GetStaticWwnnAddressOk returns a tuple with the StaticWwnnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicy) GetStaticWwnnAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StaticWwnnAddress) {
		return nil, false
	}
	return o.StaticWwnnAddress, true
}

// HasStaticWwnnAddress returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicy) HasStaticWwnnAddress() bool {
	if o != nil && !IsNil(o.StaticWwnnAddress) {
		return true
	}

	return false
}

// SetStaticWwnnAddress gets a reference to the given string and assigns it to the StaticWwnnAddress field.
func (o *VnicSanConnectivityPolicy) SetStaticWwnnAddress(v string) {
	o.StaticWwnnAddress = &v
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *VnicSanConnectivityPolicy) GetTargetPlatform() string {
	if o == nil || IsNil(o.TargetPlatform) {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicy) GetTargetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPlatform) {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicy) HasTargetPlatform() bool {
	if o != nil && !IsNil(o.TargetPlatform) {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *VnicSanConnectivityPolicy) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetWwnnAddressType returns the WwnnAddressType field value if set, zero value otherwise.
func (o *VnicSanConnectivityPolicy) GetWwnnAddressType() string {
	if o == nil || IsNil(o.WwnnAddressType) {
		var ret string
		return ret
	}
	return *o.WwnnAddressType
}

// GetWwnnAddressTypeOk returns a tuple with the WwnnAddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicy) GetWwnnAddressTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WwnnAddressType) {
		return nil, false
	}
	return o.WwnnAddressType, true
}

// HasWwnnAddressType returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicy) HasWwnnAddressType() bool {
	if o != nil && !IsNil(o.WwnnAddressType) {
		return true
	}

	return false
}

// SetWwnnAddressType gets a reference to the given string and assigns it to the WwnnAddressType field.
func (o *VnicSanConnectivityPolicy) SetWwnnAddressType(v string) {
	o.WwnnAddressType = &v
}

// GetFcIfs returns the FcIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicSanConnectivityPolicy) GetFcIfs() []VnicFcIfRelationship {
	if o == nil {
		var ret []VnicFcIfRelationship
		return ret
	}
	return o.FcIfs
}

// GetFcIfsOk returns a tuple with the FcIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicSanConnectivityPolicy) GetFcIfsOk() ([]VnicFcIfRelationship, bool) {
	if o == nil || IsNil(o.FcIfs) {
		return nil, false
	}
	return o.FcIfs, true
}

// HasFcIfs returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicy) HasFcIfs() bool {
	if o != nil && !IsNil(o.FcIfs) {
		return true
	}

	return false
}

// SetFcIfs gets a reference to the given []VnicFcIfRelationship and assigns it to the FcIfs field.
func (o *VnicSanConnectivityPolicy) SetFcIfs(v []VnicFcIfRelationship) {
	o.FcIfs = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicSanConnectivityPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicSanConnectivityPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicSanConnectivityPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *VnicSanConnectivityPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *VnicSanConnectivityPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *VnicSanConnectivityPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *VnicSanConnectivityPolicy) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicy) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
// Deprecated
func (o *VnicSanConnectivityPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

// GetWwnnPool returns the WwnnPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicSanConnectivityPolicy) GetWwnnPool() FcpoolPoolRelationship {
	if o == nil || IsNil(o.WwnnPool.Get()) {
		var ret FcpoolPoolRelationship
		return ret
	}
	return *o.WwnnPool.Get()
}

// GetWwnnPoolOk returns a tuple with the WwnnPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicSanConnectivityPolicy) GetWwnnPoolOk() (*FcpoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.WwnnPool.Get(), o.WwnnPool.IsSet()
}

// HasWwnnPool returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicy) HasWwnnPool() bool {
	if o != nil && o.WwnnPool.IsSet() {
		return true
	}

	return false
}

// SetWwnnPool gets a reference to the given NullableFcpoolPoolRelationship and assigns it to the WwnnPool field.
func (o *VnicSanConnectivityPolicy) SetWwnnPool(v FcpoolPoolRelationship) {
	o.WwnnPool.Set(&v)
}

// SetWwnnPoolNil sets the value for WwnnPool to be an explicit nil
func (o *VnicSanConnectivityPolicy) SetWwnnPoolNil() {
	o.WwnnPool.Set(nil)
}

// UnsetWwnnPool ensures that no value is present for WwnnPool, not even an explicit nil
func (o *VnicSanConnectivityPolicy) UnsetWwnnPool() {
	o.WwnnPool.Unset()
}

func (o VnicSanConnectivityPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicSanConnectivityPolicy) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PlacementMode) {
		toSerialize["PlacementMode"] = o.PlacementMode
	}
	if !IsNil(o.StaticWwnnAddress) {
		toSerialize["StaticWwnnAddress"] = o.StaticWwnnAddress
	}
	if !IsNil(o.TargetPlatform) {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if !IsNil(o.WwnnAddressType) {
		toSerialize["WwnnAddressType"] = o.WwnnAddressType
	}
	if o.FcIfs != nil {
		toSerialize["FcIfs"] = o.FcIfs
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	if o.WwnnPool.IsSet() {
		toSerialize["WwnnPool"] = o.WwnnPool.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicSanConnectivityPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type VnicSanConnectivityPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The mode used for placement of vNICs on network adapters. It can either be Auto or Custom. * `custom` - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user. * `auto` - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system.
		PlacementMode *string `json:"PlacementMode,omitempty"`
		// The WWNN address for the server node must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN's in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx.
		StaticWwnnAddress *string `json:"StaticWwnnAddress,omitempty" validate:"regexp=^$|((^20|5[0-9a-fA-F]{1}):([0-9a-fA-F]{2}:){6}([0-9a-fA-F]{2})$)"`
		// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
		TargetPlatform *string `json:"TargetPlatform,omitempty"`
		// Type of allocation selected to assign a WWNN address for the server node. * `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface.
		WwnnAddressType *string `json:"WwnnAddressType,omitempty"`
		// An array of relationships to vnicFcIf resources.
		FcIfs        []VnicFcIfRelationship                       `json:"FcIfs,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		// Deprecated
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
		WwnnPool NullableFcpoolPoolRelationship            `json:"WwnnPool,omitempty"`
	}

	varVnicSanConnectivityPolicyWithoutEmbeddedStruct := VnicSanConnectivityPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicSanConnectivityPolicyWithoutEmbeddedStruct)
	if err == nil {
		varVnicSanConnectivityPolicy := _VnicSanConnectivityPolicy{}
		varVnicSanConnectivityPolicy.ClassId = varVnicSanConnectivityPolicyWithoutEmbeddedStruct.ClassId
		varVnicSanConnectivityPolicy.ObjectType = varVnicSanConnectivityPolicyWithoutEmbeddedStruct.ObjectType
		varVnicSanConnectivityPolicy.PlacementMode = varVnicSanConnectivityPolicyWithoutEmbeddedStruct.PlacementMode
		varVnicSanConnectivityPolicy.StaticWwnnAddress = varVnicSanConnectivityPolicyWithoutEmbeddedStruct.StaticWwnnAddress
		varVnicSanConnectivityPolicy.TargetPlatform = varVnicSanConnectivityPolicyWithoutEmbeddedStruct.TargetPlatform
		varVnicSanConnectivityPolicy.WwnnAddressType = varVnicSanConnectivityPolicyWithoutEmbeddedStruct.WwnnAddressType
		varVnicSanConnectivityPolicy.FcIfs = varVnicSanConnectivityPolicyWithoutEmbeddedStruct.FcIfs
		varVnicSanConnectivityPolicy.Organization = varVnicSanConnectivityPolicyWithoutEmbeddedStruct.Organization
		varVnicSanConnectivityPolicy.Profiles = varVnicSanConnectivityPolicyWithoutEmbeddedStruct.Profiles
		varVnicSanConnectivityPolicy.WwnnPool = varVnicSanConnectivityPolicyWithoutEmbeddedStruct.WwnnPool
		*o = VnicSanConnectivityPolicy(varVnicSanConnectivityPolicy)
	} else {
		return err
	}

	varVnicSanConnectivityPolicy := _VnicSanConnectivityPolicy{}

	err = json.Unmarshal(data, &varVnicSanConnectivityPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varVnicSanConnectivityPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PlacementMode")
		delete(additionalProperties, "StaticWwnnAddress")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "WwnnAddressType")
		delete(additionalProperties, "FcIfs")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		delete(additionalProperties, "WwnnPool")

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

type NullableVnicSanConnectivityPolicy struct {
	value *VnicSanConnectivityPolicy
	isSet bool
}

func (v NullableVnicSanConnectivityPolicy) Get() *VnicSanConnectivityPolicy {
	return v.value
}

func (v *NullableVnicSanConnectivityPolicy) Set(val *VnicSanConnectivityPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicSanConnectivityPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicSanConnectivityPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicSanConnectivityPolicy(val *VnicSanConnectivityPolicy) *NullableVnicSanConnectivityPolicy {
	return &NullableVnicSanConnectivityPolicy{value: val, isSet: true}
}

func (v NullableVnicSanConnectivityPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicSanConnectivityPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
