/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-17956
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

// checks if the CloudBaseNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudBaseNetwork{}

// CloudBaseNetwork A base network object that is extended by Cloud Network objects.
type CloudBaseNetwork struct {
	VirtualizationBaseNetwork
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType  string                   `json:"ObjectType"`
	BillingUnit NullableCloudBillingUnit `json:"BillingUnit,omitempty"`
	// CIDR scheme for defining an IP block.
	Cidr                 *string                       `json:"Cidr,omitempty"`
	RegionInfo           NullableCloudCloudRegion      `json:"RegionInfo,omitempty"`
	ZoneInfo             NullableCloudAvailabilityZone `json:"ZoneInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudBaseNetwork CloudBaseNetwork

// NewCloudBaseNetwork instantiates a new CloudBaseNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBaseNetwork(classId string, objectType string) *CloudBaseNetwork {
	this := CloudBaseNetwork{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudBaseNetworkWithDefaults instantiates a new CloudBaseNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBaseNetworkWithDefaults() *CloudBaseNetwork {
	this := CloudBaseNetwork{}
	var classId string = "cloud.AwsSubnet"
	this.ClassId = classId
	var objectType string = "cloud.AwsSubnet"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudBaseNetwork) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudBaseNetwork) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudBaseNetwork) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudBaseNetwork) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudBaseNetwork) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudBaseNetwork) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBillingUnit returns the BillingUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseNetwork) GetBillingUnit() CloudBillingUnit {
	if o == nil || IsNil(o.BillingUnit.Get()) {
		var ret CloudBillingUnit
		return ret
	}
	return *o.BillingUnit.Get()
}

// GetBillingUnitOk returns a tuple with the BillingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseNetwork) GetBillingUnitOk() (*CloudBillingUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingUnit.Get(), o.BillingUnit.IsSet()
}

// HasBillingUnit returns a boolean if a field has been set.
func (o *CloudBaseNetwork) HasBillingUnit() bool {
	if o != nil && o.BillingUnit.IsSet() {
		return true
	}

	return false
}

// SetBillingUnit gets a reference to the given NullableCloudBillingUnit and assigns it to the BillingUnit field.
func (o *CloudBaseNetwork) SetBillingUnit(v CloudBillingUnit) {
	o.BillingUnit.Set(&v)
}

// SetBillingUnitNil sets the value for BillingUnit to be an explicit nil
func (o *CloudBaseNetwork) SetBillingUnitNil() {
	o.BillingUnit.Set(nil)
}

// UnsetBillingUnit ensures that no value is present for BillingUnit, not even an explicit nil
func (o *CloudBaseNetwork) UnsetBillingUnit() {
	o.BillingUnit.Unset()
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *CloudBaseNetwork) GetCidr() string {
	if o == nil || IsNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseNetwork) GetCidrOk() (*string, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *CloudBaseNetwork) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *CloudBaseNetwork) SetCidr(v string) {
	o.Cidr = &v
}

// GetRegionInfo returns the RegionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseNetwork) GetRegionInfo() CloudCloudRegion {
	if o == nil || IsNil(o.RegionInfo.Get()) {
		var ret CloudCloudRegion
		return ret
	}
	return *o.RegionInfo.Get()
}

// GetRegionInfoOk returns a tuple with the RegionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseNetwork) GetRegionInfoOk() (*CloudCloudRegion, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegionInfo.Get(), o.RegionInfo.IsSet()
}

// HasRegionInfo returns a boolean if a field has been set.
func (o *CloudBaseNetwork) HasRegionInfo() bool {
	if o != nil && o.RegionInfo.IsSet() {
		return true
	}

	return false
}

// SetRegionInfo gets a reference to the given NullableCloudCloudRegion and assigns it to the RegionInfo field.
func (o *CloudBaseNetwork) SetRegionInfo(v CloudCloudRegion) {
	o.RegionInfo.Set(&v)
}

// SetRegionInfoNil sets the value for RegionInfo to be an explicit nil
func (o *CloudBaseNetwork) SetRegionInfoNil() {
	o.RegionInfo.Set(nil)
}

// UnsetRegionInfo ensures that no value is present for RegionInfo, not even an explicit nil
func (o *CloudBaseNetwork) UnsetRegionInfo() {
	o.RegionInfo.Unset()
}

// GetZoneInfo returns the ZoneInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseNetwork) GetZoneInfo() CloudAvailabilityZone {
	if o == nil || IsNil(o.ZoneInfo.Get()) {
		var ret CloudAvailabilityZone
		return ret
	}
	return *o.ZoneInfo.Get()
}

// GetZoneInfoOk returns a tuple with the ZoneInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseNetwork) GetZoneInfoOk() (*CloudAvailabilityZone, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZoneInfo.Get(), o.ZoneInfo.IsSet()
}

// HasZoneInfo returns a boolean if a field has been set.
func (o *CloudBaseNetwork) HasZoneInfo() bool {
	if o != nil && o.ZoneInfo.IsSet() {
		return true
	}

	return false
}

// SetZoneInfo gets a reference to the given NullableCloudAvailabilityZone and assigns it to the ZoneInfo field.
func (o *CloudBaseNetwork) SetZoneInfo(v CloudAvailabilityZone) {
	o.ZoneInfo.Set(&v)
}

// SetZoneInfoNil sets the value for ZoneInfo to be an explicit nil
func (o *CloudBaseNetwork) SetZoneInfoNil() {
	o.ZoneInfo.Set(nil)
}

// UnsetZoneInfo ensures that no value is present for ZoneInfo, not even an explicit nil
func (o *CloudBaseNetwork) UnsetZoneInfo() {
	o.ZoneInfo.Unset()
}

func (o CloudBaseNetwork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudBaseNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseNetwork, errVirtualizationBaseNetwork := json.Marshal(o.VirtualizationBaseNetwork)
	if errVirtualizationBaseNetwork != nil {
		return map[string]interface{}{}, errVirtualizationBaseNetwork
	}
	errVirtualizationBaseNetwork = json.Unmarshal([]byte(serializedVirtualizationBaseNetwork), &toSerialize)
	if errVirtualizationBaseNetwork != nil {
		return map[string]interface{}{}, errVirtualizationBaseNetwork
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if o.BillingUnit.IsSet() {
		toSerialize["BillingUnit"] = o.BillingUnit.Get()
	}
	if !IsNil(o.Cidr) {
		toSerialize["Cidr"] = o.Cidr
	}
	if o.RegionInfo.IsSet() {
		toSerialize["RegionInfo"] = o.RegionInfo.Get()
	}
	if o.ZoneInfo.IsSet() {
		toSerialize["ZoneInfo"] = o.ZoneInfo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudBaseNetwork) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type CloudBaseNetworkWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType  string                   `json:"ObjectType"`
		BillingUnit NullableCloudBillingUnit `json:"BillingUnit,omitempty"`
		// CIDR scheme for defining an IP block.
		Cidr       *string                       `json:"Cidr,omitempty"`
		RegionInfo NullableCloudCloudRegion      `json:"RegionInfo,omitempty"`
		ZoneInfo   NullableCloudAvailabilityZone `json:"ZoneInfo,omitempty"`
	}

	varCloudBaseNetworkWithoutEmbeddedStruct := CloudBaseNetworkWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCloudBaseNetworkWithoutEmbeddedStruct)
	if err == nil {
		varCloudBaseNetwork := _CloudBaseNetwork{}
		varCloudBaseNetwork.ClassId = varCloudBaseNetworkWithoutEmbeddedStruct.ClassId
		varCloudBaseNetwork.ObjectType = varCloudBaseNetworkWithoutEmbeddedStruct.ObjectType
		varCloudBaseNetwork.BillingUnit = varCloudBaseNetworkWithoutEmbeddedStruct.BillingUnit
		varCloudBaseNetwork.Cidr = varCloudBaseNetworkWithoutEmbeddedStruct.Cidr
		varCloudBaseNetwork.RegionInfo = varCloudBaseNetworkWithoutEmbeddedStruct.RegionInfo
		varCloudBaseNetwork.ZoneInfo = varCloudBaseNetworkWithoutEmbeddedStruct.ZoneInfo
		*o = CloudBaseNetwork(varCloudBaseNetwork)
	} else {
		return err
	}

	varCloudBaseNetwork := _CloudBaseNetwork{}

	err = json.Unmarshal(data, &varCloudBaseNetwork)
	if err == nil {
		o.VirtualizationBaseNetwork = varCloudBaseNetwork.VirtualizationBaseNetwork
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BillingUnit")
		delete(additionalProperties, "Cidr")
		delete(additionalProperties, "RegionInfo")
		delete(additionalProperties, "ZoneInfo")

		// remove fields from embedded structs
		reflectVirtualizationBaseNetwork := reflect.ValueOf(o.VirtualizationBaseNetwork)
		for i := 0; i < reflectVirtualizationBaseNetwork.Type().NumField(); i++ {
			t := reflectVirtualizationBaseNetwork.Type().Field(i)

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

type NullableCloudBaseNetwork struct {
	value *CloudBaseNetwork
	isSet bool
}

func (v NullableCloudBaseNetwork) Get() *CloudBaseNetwork {
	return v.value
}

func (v *NullableCloudBaseNetwork) Set(val *CloudBaseNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudBaseNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudBaseNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudBaseNetwork(val *CloudBaseNetwork) *NullableCloudBaseNetwork {
	return &NullableCloudBaseNetwork{value: val, isSet: true}
}

func (v NullableCloudBaseNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudBaseNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
