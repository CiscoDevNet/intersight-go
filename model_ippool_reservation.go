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

// checks if the IppoolReservation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IppoolReservation{}

// IppoolReservation The IP reservation object, used to hold reserved addresses.
type IppoolReservation struct {
	PoolReservation
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP identity to be reserved.
	Identity *string `json:"Identity,omitempty"`
	// Type of the IP address that needs to be reserved. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
	IpType *string `json:"IpType,omitempty"`
	// IPv4 Address to be reserved.
	IpV4Address *string `json:"IpV4Address,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	// IPv6 Address to be reserved.
	IpV6Address *string `json:"IpV6Address,omitempty" validate:"regexp=^$|^(([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:[0-9A-Fa-f]{0,4}|:[0-9A-Fa-f]{1,4})?|(:[0-9A-Fa-f]{1,4}){0,2})|(:[0-9A-Fa-f]{1,4}){0,3})|(:[0-9A-Fa-f]{1,4}){0,4})|:(:[0-9A-Fa-f]{1,4}){0,5})((:[0-9A-Fa-f]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\\\\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9A-Fa-f]{1,4}:){1,6}|:):[0-9A-Fa-f]{0,4}|([0-9A-Fa-f]{1,4}:){7}:)$"`
	// The moid of the Virtual Routing and Forwarding MO.
	VrfMoid              *string                                      `json:"VrfMoid,omitempty"`
	BlockHead            NullableIppoolShadowBlockRelationship        `json:"BlockHead,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	Pool                 NullableIppoolPoolRelationship               `json:"Pool,omitempty"`
	PoolMember           NullableIppoolPoolMemberRelationship         `json:"PoolMember,omitempty"`
	ShadowPool           NullableIppoolShadowPoolRelationship         `json:"ShadowPool,omitempty"`
	Universe             NullableIppoolUniverseRelationship           `json:"Universe,omitempty"`
	Vrf                  NullableVrfVrfRelationship                   `json:"Vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolReservation IppoolReservation

// NewIppoolReservation instantiates a new IppoolReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolReservation(classId string, objectType string) *IppoolReservation {
	this := IppoolReservation{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allocationType string = "dynamic"
	this.AllocationType = &allocationType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// NewIppoolReservationWithDefaults instantiates a new IppoolReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolReservationWithDefaults() *IppoolReservation {
	this := IppoolReservation{}
	var classId string = "ippool.Reservation"
	this.ClassId = classId
	var objectType string = "ippool.Reservation"
	this.ObjectType = objectType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolReservation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolReservation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolReservation) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "ippool.Reservation" of the ClassId field.
func (o *IppoolReservation) GetDefaultClassId() interface{} {
	return "ippool.Reservation"
}

// GetObjectType returns the ObjectType field value
func (o *IppoolReservation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolReservation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolReservation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "ippool.Reservation" of the ObjectType field.
func (o *IppoolReservation) GetDefaultObjectType() interface{} {
	return "ippool.Reservation"
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *IppoolReservation) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservation) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *IppoolReservation) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *IppoolReservation) SetIdentity(v string) {
	o.Identity = &v
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *IppoolReservation) GetIpType() string {
	if o == nil || IsNil(o.IpType) {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservation) GetIpTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IpType) {
		return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *IppoolReservation) HasIpType() bool {
	if o != nil && !IsNil(o.IpType) {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *IppoolReservation) SetIpType(v string) {
	o.IpType = &v
}

// GetIpV4Address returns the IpV4Address field value if set, zero value otherwise.
func (o *IppoolReservation) GetIpV4Address() string {
	if o == nil || IsNil(o.IpV4Address) {
		var ret string
		return ret
	}
	return *o.IpV4Address
}

// GetIpV4AddressOk returns a tuple with the IpV4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservation) GetIpV4AddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpV4Address) {
		return nil, false
	}
	return o.IpV4Address, true
}

// HasIpV4Address returns a boolean if a field has been set.
func (o *IppoolReservation) HasIpV4Address() bool {
	if o != nil && !IsNil(o.IpV4Address) {
		return true
	}

	return false
}

// SetIpV4Address gets a reference to the given string and assigns it to the IpV4Address field.
func (o *IppoolReservation) SetIpV4Address(v string) {
	o.IpV4Address = &v
}

// GetIpV6Address returns the IpV6Address field value if set, zero value otherwise.
func (o *IppoolReservation) GetIpV6Address() string {
	if o == nil || IsNil(o.IpV6Address) {
		var ret string
		return ret
	}
	return *o.IpV6Address
}

// GetIpV6AddressOk returns a tuple with the IpV6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservation) GetIpV6AddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpV6Address) {
		return nil, false
	}
	return o.IpV6Address, true
}

// HasIpV6Address returns a boolean if a field has been set.
func (o *IppoolReservation) HasIpV6Address() bool {
	if o != nil && !IsNil(o.IpV6Address) {
		return true
	}

	return false
}

// SetIpV6Address gets a reference to the given string and assigns it to the IpV6Address field.
func (o *IppoolReservation) SetIpV6Address(v string) {
	o.IpV6Address = &v
}

// GetVrfMoid returns the VrfMoid field value if set, zero value otherwise.
func (o *IppoolReservation) GetVrfMoid() string {
	if o == nil || IsNil(o.VrfMoid) {
		var ret string
		return ret
	}
	return *o.VrfMoid
}

// GetVrfMoidOk returns a tuple with the VrfMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservation) GetVrfMoidOk() (*string, bool) {
	if o == nil || IsNil(o.VrfMoid) {
		return nil, false
	}
	return o.VrfMoid, true
}

// HasVrfMoid returns a boolean if a field has been set.
func (o *IppoolReservation) HasVrfMoid() bool {
	if o != nil && !IsNil(o.VrfMoid) {
		return true
	}

	return false
}

// SetVrfMoid gets a reference to the given string and assigns it to the VrfMoid field.
func (o *IppoolReservation) SetVrfMoid(v string) {
	o.VrfMoid = &v
}

// GetBlockHead returns the BlockHead field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolReservation) GetBlockHead() IppoolShadowBlockRelationship {
	if o == nil || IsNil(o.BlockHead.Get()) {
		var ret IppoolShadowBlockRelationship
		return ret
	}
	return *o.BlockHead.Get()
}

// GetBlockHeadOk returns a tuple with the BlockHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolReservation) GetBlockHeadOk() (*IppoolShadowBlockRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockHead.Get(), o.BlockHead.IsSet()
}

// HasBlockHead returns a boolean if a field has been set.
func (o *IppoolReservation) HasBlockHead() bool {
	if o != nil && o.BlockHead.IsSet() {
		return true
	}

	return false
}

// SetBlockHead gets a reference to the given NullableIppoolShadowBlockRelationship and assigns it to the BlockHead field.
func (o *IppoolReservation) SetBlockHead(v IppoolShadowBlockRelationship) {
	o.BlockHead.Set(&v)
}

// SetBlockHeadNil sets the value for BlockHead to be an explicit nil
func (o *IppoolReservation) SetBlockHeadNil() {
	o.BlockHead.Set(nil)
}

// UnsetBlockHead ensures that no value is present for BlockHead, not even an explicit nil
func (o *IppoolReservation) UnsetBlockHead() {
	o.BlockHead.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolReservation) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolReservation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *IppoolReservation) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *IppoolReservation) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *IppoolReservation) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *IppoolReservation) UnsetOrganization() {
	o.Organization.Unset()
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolReservation) GetPool() IppoolPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolReservation) GetPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolReservation) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableIppoolPoolRelationship and assigns it to the Pool field.
func (o *IppoolReservation) SetPool(v IppoolPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *IppoolReservation) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *IppoolReservation) UnsetPool() {
	o.Pool.Unset()
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolReservation) GetPoolMember() IppoolPoolMemberRelationship {
	if o == nil || IsNil(o.PoolMember.Get()) {
		var ret IppoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember.Get()
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolReservation) GetPoolMemberOk() (*IppoolPoolMemberRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoolMember.Get(), o.PoolMember.IsSet()
}

// HasPoolMember returns a boolean if a field has been set.
func (o *IppoolReservation) HasPoolMember() bool {
	if o != nil && o.PoolMember.IsSet() {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given NullableIppoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *IppoolReservation) SetPoolMember(v IppoolPoolMemberRelationship) {
	o.PoolMember.Set(&v)
}

// SetPoolMemberNil sets the value for PoolMember to be an explicit nil
func (o *IppoolReservation) SetPoolMemberNil() {
	o.PoolMember.Set(nil)
}

// UnsetPoolMember ensures that no value is present for PoolMember, not even an explicit nil
func (o *IppoolReservation) UnsetPoolMember() {
	o.PoolMember.Unset()
}

// GetShadowPool returns the ShadowPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolReservation) GetShadowPool() IppoolShadowPoolRelationship {
	if o == nil || IsNil(o.ShadowPool.Get()) {
		var ret IppoolShadowPoolRelationship
		return ret
	}
	return *o.ShadowPool.Get()
}

// GetShadowPoolOk returns a tuple with the ShadowPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolReservation) GetShadowPoolOk() (*IppoolShadowPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShadowPool.Get(), o.ShadowPool.IsSet()
}

// HasShadowPool returns a boolean if a field has been set.
func (o *IppoolReservation) HasShadowPool() bool {
	if o != nil && o.ShadowPool.IsSet() {
		return true
	}

	return false
}

// SetShadowPool gets a reference to the given NullableIppoolShadowPoolRelationship and assigns it to the ShadowPool field.
func (o *IppoolReservation) SetShadowPool(v IppoolShadowPoolRelationship) {
	o.ShadowPool.Set(&v)
}

// SetShadowPoolNil sets the value for ShadowPool to be an explicit nil
func (o *IppoolReservation) SetShadowPoolNil() {
	o.ShadowPool.Set(nil)
}

// UnsetShadowPool ensures that no value is present for ShadowPool, not even an explicit nil
func (o *IppoolReservation) UnsetShadowPool() {
	o.ShadowPool.Unset()
}

// GetUniverse returns the Universe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolReservation) GetUniverse() IppoolUniverseRelationship {
	if o == nil || IsNil(o.Universe.Get()) {
		var ret IppoolUniverseRelationship
		return ret
	}
	return *o.Universe.Get()
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolReservation) GetUniverseOk() (*IppoolUniverseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Universe.Get(), o.Universe.IsSet()
}

// HasUniverse returns a boolean if a field has been set.
func (o *IppoolReservation) HasUniverse() bool {
	if o != nil && o.Universe.IsSet() {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given NullableIppoolUniverseRelationship and assigns it to the Universe field.
func (o *IppoolReservation) SetUniverse(v IppoolUniverseRelationship) {
	o.Universe.Set(&v)
}

// SetUniverseNil sets the value for Universe to be an explicit nil
func (o *IppoolReservation) SetUniverseNil() {
	o.Universe.Set(nil)
}

// UnsetUniverse ensures that no value is present for Universe, not even an explicit nil
func (o *IppoolReservation) UnsetUniverse() {
	o.Universe.Unset()
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolReservation) GetVrf() VrfVrfRelationship {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolReservation) GetVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *IppoolReservation) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableVrfVrfRelationship and assigns it to the Vrf field.
func (o *IppoolReservation) SetVrf(v VrfVrfRelationship) {
	o.Vrf.Set(&v)
}

// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *IppoolReservation) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *IppoolReservation) UnsetVrf() {
	o.Vrf.Unset()
}

func (o IppoolReservation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IppoolReservation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolReservation, errPoolReservation := json.Marshal(o.PoolReservation)
	if errPoolReservation != nil {
		return map[string]interface{}{}, errPoolReservation
	}
	errPoolReservation = json.Unmarshal([]byte(serializedPoolReservation), &toSerialize)
	if errPoolReservation != nil {
		return map[string]interface{}{}, errPoolReservation
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Identity) {
		toSerialize["Identity"] = o.Identity
	}
	if !IsNil(o.IpType) {
		toSerialize["IpType"] = o.IpType
	}
	if !IsNil(o.IpV4Address) {
		toSerialize["IpV4Address"] = o.IpV4Address
	}
	if !IsNil(o.IpV6Address) {
		toSerialize["IpV6Address"] = o.IpV6Address
	}
	if !IsNil(o.VrfMoid) {
		toSerialize["VrfMoid"] = o.VrfMoid
	}
	if o.BlockHead.IsSet() {
		toSerialize["BlockHead"] = o.BlockHead.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.Pool.IsSet() {
		toSerialize["Pool"] = o.Pool.Get()
	}
	if o.PoolMember.IsSet() {
		toSerialize["PoolMember"] = o.PoolMember.Get()
	}
	if o.ShadowPool.IsSet() {
		toSerialize["ShadowPool"] = o.ShadowPool.Get()
	}
	if o.Universe.IsSet() {
		toSerialize["Universe"] = o.Universe.Get()
	}
	if o.Vrf.IsSet() {
		toSerialize["Vrf"] = o.Vrf.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IppoolReservation) UnmarshalJSON(data []byte) (err error) {
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
	type IppoolReservationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP identity to be reserved.
		Identity *string `json:"Identity,omitempty"`
		// Type of the IP address that needs to be reserved. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
		IpType *string `json:"IpType,omitempty"`
		// IPv4 Address to be reserved.
		IpV4Address *string `json:"IpV4Address,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
		// IPv6 Address to be reserved.
		IpV6Address *string `json:"IpV6Address,omitempty" validate:"regexp=^$|^(([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:[0-9A-Fa-f]{0,4}|:[0-9A-Fa-f]{1,4})?|(:[0-9A-Fa-f]{1,4}){0,2})|(:[0-9A-Fa-f]{1,4}){0,3})|(:[0-9A-Fa-f]{1,4}){0,4})|:(:[0-9A-Fa-f]{1,4}){0,5})((:[0-9A-Fa-f]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\\\\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9A-Fa-f]{1,4}:){1,6}|:):[0-9A-Fa-f]{0,4}|([0-9A-Fa-f]{1,4}:){7}:)$"`
		// The moid of the Virtual Routing and Forwarding MO.
		VrfMoid      *string                                      `json:"VrfMoid,omitempty"`
		BlockHead    NullableIppoolShadowBlockRelationship        `json:"BlockHead,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		Pool         NullableIppoolPoolRelationship               `json:"Pool,omitempty"`
		PoolMember   NullableIppoolPoolMemberRelationship         `json:"PoolMember,omitempty"`
		ShadowPool   NullableIppoolShadowPoolRelationship         `json:"ShadowPool,omitempty"`
		Universe     NullableIppoolUniverseRelationship           `json:"Universe,omitempty"`
		Vrf          NullableVrfVrfRelationship                   `json:"Vrf,omitempty"`
	}

	varIppoolReservationWithoutEmbeddedStruct := IppoolReservationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIppoolReservationWithoutEmbeddedStruct)
	if err == nil {
		varIppoolReservation := _IppoolReservation{}
		varIppoolReservation.ClassId = varIppoolReservationWithoutEmbeddedStruct.ClassId
		varIppoolReservation.ObjectType = varIppoolReservationWithoutEmbeddedStruct.ObjectType
		varIppoolReservation.Identity = varIppoolReservationWithoutEmbeddedStruct.Identity
		varIppoolReservation.IpType = varIppoolReservationWithoutEmbeddedStruct.IpType
		varIppoolReservation.IpV4Address = varIppoolReservationWithoutEmbeddedStruct.IpV4Address
		varIppoolReservation.IpV6Address = varIppoolReservationWithoutEmbeddedStruct.IpV6Address
		varIppoolReservation.VrfMoid = varIppoolReservationWithoutEmbeddedStruct.VrfMoid
		varIppoolReservation.BlockHead = varIppoolReservationWithoutEmbeddedStruct.BlockHead
		varIppoolReservation.Organization = varIppoolReservationWithoutEmbeddedStruct.Organization
		varIppoolReservation.Pool = varIppoolReservationWithoutEmbeddedStruct.Pool
		varIppoolReservation.PoolMember = varIppoolReservationWithoutEmbeddedStruct.PoolMember
		varIppoolReservation.ShadowPool = varIppoolReservationWithoutEmbeddedStruct.ShadowPool
		varIppoolReservation.Universe = varIppoolReservationWithoutEmbeddedStruct.Universe
		varIppoolReservation.Vrf = varIppoolReservationWithoutEmbeddedStruct.Vrf
		*o = IppoolReservation(varIppoolReservation)
	} else {
		return err
	}

	varIppoolReservation := _IppoolReservation{}

	err = json.Unmarshal(data, &varIppoolReservation)
	if err == nil {
		o.PoolReservation = varIppoolReservation.PoolReservation
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "IpType")
		delete(additionalProperties, "IpV4Address")
		delete(additionalProperties, "IpV6Address")
		delete(additionalProperties, "VrfMoid")
		delete(additionalProperties, "BlockHead")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "ShadowPool")
		delete(additionalProperties, "Universe")
		delete(additionalProperties, "Vrf")

		// remove fields from embedded structs
		reflectPoolReservation := reflect.ValueOf(o.PoolReservation)
		for i := 0; i < reflectPoolReservation.Type().NumField(); i++ {
			t := reflectPoolReservation.Type().Field(i)

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

type NullableIppoolReservation struct {
	value *IppoolReservation
	isSet bool
}

func (v NullableIppoolReservation) Get() *IppoolReservation {
	return v.value
}

func (v *NullableIppoolReservation) Set(val *IppoolReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolReservation(val *IppoolReservation) *NullableIppoolReservation {
	return &NullableIppoolReservation{value: val, isSet: true}
}

func (v NullableIppoolReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
