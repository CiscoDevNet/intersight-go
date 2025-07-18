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

// checks if the FabricMulticastPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricMulticastPolicy{}

// FabricMulticastPolicy A policy to configure Multicast settings for all the Virtual LAN networks.
type FabricMulticastPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Used to define the IGMP Querier IP address.
	QuerierIpAddress *string `json:"QuerierIpAddress,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	// Used to define the IGMP Querier IP address of the peer switch.
	QuerierIpAddressPeer *string `json:"QuerierIpAddressPeer,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	// Administrative state of the IGMP Querier for this VLAN. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
	QuerierState *string `json:"QuerierState,omitempty"`
	// Administrative state of the IGMP Snooping for this VLAN. * `Enabled` - Admin configured Enabled State. * `Disabled` - Admin configured Disabled State.
	SnoopingState *string `json:"SnoopingState,omitempty"`
	// Administrative state of the IGMP source IP proxy for this VLAN. * `Enabled` - Admin configured Enabled State. * `Disabled` - Admin configured Disabled State.
	SrcIpProxy           *string                                      `json:"SrcIpProxy,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricMulticastPolicy FabricMulticastPolicy

// NewFabricMulticastPolicy instantiates a new FabricMulticastPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricMulticastPolicy(classId string, objectType string) *FabricMulticastPolicy {
	this := FabricMulticastPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var querierState string = "Disabled"
	this.QuerierState = &querierState
	var snoopingState string = "Enabled"
	this.SnoopingState = &snoopingState
	var srcIpProxy string = "Enabled"
	this.SrcIpProxy = &srcIpProxy
	return &this
}

// NewFabricMulticastPolicyWithDefaults instantiates a new FabricMulticastPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricMulticastPolicyWithDefaults() *FabricMulticastPolicy {
	this := FabricMulticastPolicy{}
	var classId string = "fabric.MulticastPolicy"
	this.ClassId = classId
	var objectType string = "fabric.MulticastPolicy"
	this.ObjectType = objectType
	var querierState string = "Disabled"
	this.QuerierState = &querierState
	var snoopingState string = "Enabled"
	this.SnoopingState = &snoopingState
	var srcIpProxy string = "Enabled"
	this.SrcIpProxy = &srcIpProxy
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricMulticastPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricMulticastPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.MulticastPolicy" of the ClassId field.
func (o *FabricMulticastPolicy) GetDefaultClassId() interface{} {
	return "fabric.MulticastPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *FabricMulticastPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricMulticastPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.MulticastPolicy" of the ObjectType field.
func (o *FabricMulticastPolicy) GetDefaultObjectType() interface{} {
	return "fabric.MulticastPolicy"
}

// GetQuerierIpAddress returns the QuerierIpAddress field value if set, zero value otherwise.
func (o *FabricMulticastPolicy) GetQuerierIpAddress() string {
	if o == nil || IsNil(o.QuerierIpAddress) {
		var ret string
		return ret
	}
	return *o.QuerierIpAddress
}

// GetQuerierIpAddressOk returns a tuple with the QuerierIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicy) GetQuerierIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.QuerierIpAddress) {
		return nil, false
	}
	return o.QuerierIpAddress, true
}

// HasQuerierIpAddress returns a boolean if a field has been set.
func (o *FabricMulticastPolicy) HasQuerierIpAddress() bool {
	if o != nil && !IsNil(o.QuerierIpAddress) {
		return true
	}

	return false
}

// SetQuerierIpAddress gets a reference to the given string and assigns it to the QuerierIpAddress field.
func (o *FabricMulticastPolicy) SetQuerierIpAddress(v string) {
	o.QuerierIpAddress = &v
}

// GetQuerierIpAddressPeer returns the QuerierIpAddressPeer field value if set, zero value otherwise.
func (o *FabricMulticastPolicy) GetQuerierIpAddressPeer() string {
	if o == nil || IsNil(o.QuerierIpAddressPeer) {
		var ret string
		return ret
	}
	return *o.QuerierIpAddressPeer
}

// GetQuerierIpAddressPeerOk returns a tuple with the QuerierIpAddressPeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicy) GetQuerierIpAddressPeerOk() (*string, bool) {
	if o == nil || IsNil(o.QuerierIpAddressPeer) {
		return nil, false
	}
	return o.QuerierIpAddressPeer, true
}

// HasQuerierIpAddressPeer returns a boolean if a field has been set.
func (o *FabricMulticastPolicy) HasQuerierIpAddressPeer() bool {
	if o != nil && !IsNil(o.QuerierIpAddressPeer) {
		return true
	}

	return false
}

// SetQuerierIpAddressPeer gets a reference to the given string and assigns it to the QuerierIpAddressPeer field.
func (o *FabricMulticastPolicy) SetQuerierIpAddressPeer(v string) {
	o.QuerierIpAddressPeer = &v
}

// GetQuerierState returns the QuerierState field value if set, zero value otherwise.
func (o *FabricMulticastPolicy) GetQuerierState() string {
	if o == nil || IsNil(o.QuerierState) {
		var ret string
		return ret
	}
	return *o.QuerierState
}

// GetQuerierStateOk returns a tuple with the QuerierState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicy) GetQuerierStateOk() (*string, bool) {
	if o == nil || IsNil(o.QuerierState) {
		return nil, false
	}
	return o.QuerierState, true
}

// HasQuerierState returns a boolean if a field has been set.
func (o *FabricMulticastPolicy) HasQuerierState() bool {
	if o != nil && !IsNil(o.QuerierState) {
		return true
	}

	return false
}

// SetQuerierState gets a reference to the given string and assigns it to the QuerierState field.
func (o *FabricMulticastPolicy) SetQuerierState(v string) {
	o.QuerierState = &v
}

// GetSnoopingState returns the SnoopingState field value if set, zero value otherwise.
func (o *FabricMulticastPolicy) GetSnoopingState() string {
	if o == nil || IsNil(o.SnoopingState) {
		var ret string
		return ret
	}
	return *o.SnoopingState
}

// GetSnoopingStateOk returns a tuple with the SnoopingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicy) GetSnoopingStateOk() (*string, bool) {
	if o == nil || IsNil(o.SnoopingState) {
		return nil, false
	}
	return o.SnoopingState, true
}

// HasSnoopingState returns a boolean if a field has been set.
func (o *FabricMulticastPolicy) HasSnoopingState() bool {
	if o != nil && !IsNil(o.SnoopingState) {
		return true
	}

	return false
}

// SetSnoopingState gets a reference to the given string and assigns it to the SnoopingState field.
func (o *FabricMulticastPolicy) SetSnoopingState(v string) {
	o.SnoopingState = &v
}

// GetSrcIpProxy returns the SrcIpProxy field value if set, zero value otherwise.
func (o *FabricMulticastPolicy) GetSrcIpProxy() string {
	if o == nil || IsNil(o.SrcIpProxy) {
		var ret string
		return ret
	}
	return *o.SrcIpProxy
}

// GetSrcIpProxyOk returns a tuple with the SrcIpProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicy) GetSrcIpProxyOk() (*string, bool) {
	if o == nil || IsNil(o.SrcIpProxy) {
		return nil, false
	}
	return o.SrcIpProxy, true
}

// HasSrcIpProxy returns a boolean if a field has been set.
func (o *FabricMulticastPolicy) HasSrcIpProxy() bool {
	if o != nil && !IsNil(o.SrcIpProxy) {
		return true
	}

	return false
}

// SetSrcIpProxy gets a reference to the given string and assigns it to the SrcIpProxy field.
func (o *FabricMulticastPolicy) SetSrcIpProxy(v string) {
	o.SrcIpProxy = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricMulticastPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricMulticastPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricMulticastPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricMulticastPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *FabricMulticastPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *FabricMulticastPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

func (o FabricMulticastPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricMulticastPolicy) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.QuerierIpAddress) {
		toSerialize["QuerierIpAddress"] = o.QuerierIpAddress
	}
	if !IsNil(o.QuerierIpAddressPeer) {
		toSerialize["QuerierIpAddressPeer"] = o.QuerierIpAddressPeer
	}
	if !IsNil(o.QuerierState) {
		toSerialize["QuerierState"] = o.QuerierState
	}
	if !IsNil(o.SnoopingState) {
		toSerialize["SnoopingState"] = o.SnoopingState
	}
	if !IsNil(o.SrcIpProxy) {
		toSerialize["SrcIpProxy"] = o.SrcIpProxy
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricMulticastPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type FabricMulticastPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Used to define the IGMP Querier IP address.
		QuerierIpAddress *string `json:"QuerierIpAddress,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
		// Used to define the IGMP Querier IP address of the peer switch.
		QuerierIpAddressPeer *string `json:"QuerierIpAddressPeer,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
		// Administrative state of the IGMP Querier for this VLAN. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
		QuerierState *string `json:"QuerierState,omitempty"`
		// Administrative state of the IGMP Snooping for this VLAN. * `Enabled` - Admin configured Enabled State. * `Disabled` - Admin configured Disabled State.
		SnoopingState *string `json:"SnoopingState,omitempty"`
		// Administrative state of the IGMP source IP proxy for this VLAN. * `Enabled` - Admin configured Enabled State. * `Disabled` - Admin configured Disabled State.
		SrcIpProxy   *string                                      `json:"SrcIpProxy,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varFabricMulticastPolicyWithoutEmbeddedStruct := FabricMulticastPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricMulticastPolicyWithoutEmbeddedStruct)
	if err == nil {
		varFabricMulticastPolicy := _FabricMulticastPolicy{}
		varFabricMulticastPolicy.ClassId = varFabricMulticastPolicyWithoutEmbeddedStruct.ClassId
		varFabricMulticastPolicy.ObjectType = varFabricMulticastPolicyWithoutEmbeddedStruct.ObjectType
		varFabricMulticastPolicy.QuerierIpAddress = varFabricMulticastPolicyWithoutEmbeddedStruct.QuerierIpAddress
		varFabricMulticastPolicy.QuerierIpAddressPeer = varFabricMulticastPolicyWithoutEmbeddedStruct.QuerierIpAddressPeer
		varFabricMulticastPolicy.QuerierState = varFabricMulticastPolicyWithoutEmbeddedStruct.QuerierState
		varFabricMulticastPolicy.SnoopingState = varFabricMulticastPolicyWithoutEmbeddedStruct.SnoopingState
		varFabricMulticastPolicy.SrcIpProxy = varFabricMulticastPolicyWithoutEmbeddedStruct.SrcIpProxy
		varFabricMulticastPolicy.Organization = varFabricMulticastPolicyWithoutEmbeddedStruct.Organization
		*o = FabricMulticastPolicy(varFabricMulticastPolicy)
	} else {
		return err
	}

	varFabricMulticastPolicy := _FabricMulticastPolicy{}

	err = json.Unmarshal(data, &varFabricMulticastPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varFabricMulticastPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "QuerierIpAddress")
		delete(additionalProperties, "QuerierIpAddressPeer")
		delete(additionalProperties, "QuerierState")
		delete(additionalProperties, "SnoopingState")
		delete(additionalProperties, "SrcIpProxy")
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

type NullableFabricMulticastPolicy struct {
	value *FabricMulticastPolicy
	isSet bool
}

func (v NullableFabricMulticastPolicy) Get() *FabricMulticastPolicy {
	return v.value
}

func (v *NullableFabricMulticastPolicy) Set(val *FabricMulticastPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricMulticastPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricMulticastPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricMulticastPolicy(val *FabricMulticastPolicy) *NullableFabricMulticastPolicy {
	return &NullableFabricMulticastPolicy{value: val, isSet: true}
}

func (v NullableFabricMulticastPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricMulticastPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
