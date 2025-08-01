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

// checks if the MerakiNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerakiNetwork{}

// MerakiNetwork A network belonging to an organization.
type MerakiNetwork struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Meraki network name seamlessly uniting devices for effortless connectivity.
	Name *string `json:"Name,omitempty"`
	// The unique Meraki network id.
	NetworkId   *string  `json:"NetworkId,omitempty"`
	NetworkTags []string `json:"NetworkTags,omitempty"`
	// The unique Meraki organization id.
	OrganizationId       *string                                     `json:"OrganizationId,omitempty"`
	ProductTypes         []string                                    `json:"ProductTypes,omitempty"`
	Organization         NullableMerakiOrganizationRelationship      `json:"Organization,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MerakiNetwork MerakiNetwork

// NewMerakiNetwork instantiates a new MerakiNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerakiNetwork(classId string, objectType string) *MerakiNetwork {
	this := MerakiNetwork{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMerakiNetworkWithDefaults instantiates a new MerakiNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerakiNetworkWithDefaults() *MerakiNetwork {
	this := MerakiNetwork{}
	var classId string = "meraki.Network"
	this.ClassId = classId
	var objectType string = "meraki.Network"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MerakiNetwork) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MerakiNetwork) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MerakiNetwork) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "meraki.Network" of the ClassId field.
func (o *MerakiNetwork) GetDefaultClassId() interface{} {
	return "meraki.Network"
}

// GetObjectType returns the ObjectType field value
func (o *MerakiNetwork) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MerakiNetwork) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MerakiNetwork) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "meraki.Network" of the ObjectType field.
func (o *MerakiNetwork) GetDefaultObjectType() interface{} {
	return "meraki.Network"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MerakiNetwork) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerakiNetwork) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MerakiNetwork) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MerakiNetwork) SetName(v string) {
	o.Name = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *MerakiNetwork) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerakiNetwork) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *MerakiNetwork) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *MerakiNetwork) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetNetworkTags returns the NetworkTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerakiNetwork) GetNetworkTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerakiNetwork) GetNetworkTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.NetworkTags) {
		return nil, false
	}
	return o.NetworkTags, true
}

// HasNetworkTags returns a boolean if a field has been set.
func (o *MerakiNetwork) HasNetworkTags() bool {
	if o != nil && !IsNil(o.NetworkTags) {
		return true
	}

	return false
}

// SetNetworkTags gets a reference to the given []string and assigns it to the NetworkTags field.
func (o *MerakiNetwork) SetNetworkTags(v []string) {
	o.NetworkTags = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *MerakiNetwork) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerakiNetwork) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *MerakiNetwork) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *MerakiNetwork) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerakiNetwork) GetProductTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerakiNetwork) GetProductTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductTypes) {
		return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *MerakiNetwork) HasProductTypes() bool {
	if o != nil && !IsNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *MerakiNetwork) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerakiNetwork) GetOrganization() MerakiOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret MerakiOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerakiNetwork) GetOrganizationOk() (*MerakiOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *MerakiNetwork) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableMerakiOrganizationRelationship and assigns it to the Organization field.
func (o *MerakiNetwork) SetOrganization(v MerakiOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *MerakiNetwork) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *MerakiNetwork) UnsetOrganization() {
	o.Organization.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerakiNetwork) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerakiNetwork) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MerakiNetwork) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MerakiNetwork) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *MerakiNetwork) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *MerakiNetwork) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o MerakiNetwork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerakiNetwork) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.NetworkId) {
		toSerialize["NetworkId"] = o.NetworkId
	}
	if o.NetworkTags != nil {
		toSerialize["NetworkTags"] = o.NetworkTags
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["OrganizationId"] = o.OrganizationId
	}
	if o.ProductTypes != nil {
		toSerialize["ProductTypes"] = o.ProductTypes
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MerakiNetwork) UnmarshalJSON(data []byte) (err error) {
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
	type MerakiNetworkWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The Meraki network name seamlessly uniting devices for effortless connectivity.
		Name *string `json:"Name,omitempty"`
		// The unique Meraki network id.
		NetworkId   *string  `json:"NetworkId,omitempty"`
		NetworkTags []string `json:"NetworkTags,omitempty"`
		// The unique Meraki organization id.
		OrganizationId   *string                                     `json:"OrganizationId,omitempty"`
		ProductTypes     []string                                    `json:"ProductTypes,omitempty"`
		Organization     NullableMerakiOrganizationRelationship      `json:"Organization,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varMerakiNetworkWithoutEmbeddedStruct := MerakiNetworkWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMerakiNetworkWithoutEmbeddedStruct)
	if err == nil {
		varMerakiNetwork := _MerakiNetwork{}
		varMerakiNetwork.ClassId = varMerakiNetworkWithoutEmbeddedStruct.ClassId
		varMerakiNetwork.ObjectType = varMerakiNetworkWithoutEmbeddedStruct.ObjectType
		varMerakiNetwork.Name = varMerakiNetworkWithoutEmbeddedStruct.Name
		varMerakiNetwork.NetworkId = varMerakiNetworkWithoutEmbeddedStruct.NetworkId
		varMerakiNetwork.NetworkTags = varMerakiNetworkWithoutEmbeddedStruct.NetworkTags
		varMerakiNetwork.OrganizationId = varMerakiNetworkWithoutEmbeddedStruct.OrganizationId
		varMerakiNetwork.ProductTypes = varMerakiNetworkWithoutEmbeddedStruct.ProductTypes
		varMerakiNetwork.Organization = varMerakiNetworkWithoutEmbeddedStruct.Organization
		varMerakiNetwork.RegisteredDevice = varMerakiNetworkWithoutEmbeddedStruct.RegisteredDevice
		*o = MerakiNetwork(varMerakiNetwork)
	} else {
		return err
	}

	varMerakiNetwork := _MerakiNetwork{}

	err = json.Unmarshal(data, &varMerakiNetwork)
	if err == nil {
		o.MoBaseMo = varMerakiNetwork.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NetworkId")
		delete(additionalProperties, "NetworkTags")
		delete(additionalProperties, "OrganizationId")
		delete(additionalProperties, "ProductTypes")
		delete(additionalProperties, "Organization")
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

type NullableMerakiNetwork struct {
	value *MerakiNetwork
	isSet bool
}

func (v NullableMerakiNetwork) Get() *MerakiNetwork {
	return v.value
}

func (v *NullableMerakiNetwork) Set(val *MerakiNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableMerakiNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableMerakiNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerakiNetwork(val *MerakiNetwork) *NullableMerakiNetwork {
	return &NullableMerakiNetwork{value: val, isSet: true}
}

func (v NullableMerakiNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerakiNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
