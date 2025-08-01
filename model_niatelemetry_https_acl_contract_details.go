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

// checks if the NiatelemetryHttpsAclContractDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryHttpsAclContractDetails{}

// NiatelemetryHttpsAclContractDetails Object to capture the HTTPS ACL contract details in APIC.
type NiatelemetryHttpsAclContractDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Consumer Dn of the HTTPS ACL contract children MOs for APIC.
	ConsumerDn *string `json:"ConsumerDn,omitempty"`
	// Name of HTTPS ACL contract for APIC.
	ContractName *string `json:"ContractName,omitempty"`
	// Mgmt Inst Dn of the HTTPS ACL contract children MOs for APIC.
	MgmtInstpDn *string `json:"MgmtInstpDn,omitempty"`
	// Mgmt subnet address of the HTTPS ACL contract children MOs for APIC.
	MgmtSubnetAddresses *string `json:"MgmtSubnetAddresses,omitempty"`
	// Provider dn of the HTTPS ACL contract children MOs for APIC.
	ProviderDn *string `json:"ProviderDn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                                     `json:"SiteName,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryHttpsAclContractDetails NiatelemetryHttpsAclContractDetails

// NewNiatelemetryHttpsAclContractDetails instantiates a new NiatelemetryHttpsAclContractDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryHttpsAclContractDetails(classId string, objectType string) *NiatelemetryHttpsAclContractDetails {
	this := NiatelemetryHttpsAclContractDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryHttpsAclContractDetailsWithDefaults instantiates a new NiatelemetryHttpsAclContractDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryHttpsAclContractDetailsWithDefaults() *NiatelemetryHttpsAclContractDetails {
	this := NiatelemetryHttpsAclContractDetails{}
	var classId string = "niatelemetry.HttpsAclContractDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.HttpsAclContractDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryHttpsAclContractDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryHttpsAclContractDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.HttpsAclContractDetails" of the ClassId field.
func (o *NiatelemetryHttpsAclContractDetails) GetDefaultClassId() interface{} {
	return "niatelemetry.HttpsAclContractDetails"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryHttpsAclContractDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryHttpsAclContractDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.HttpsAclContractDetails" of the ObjectType field.
func (o *NiatelemetryHttpsAclContractDetails) GetDefaultObjectType() interface{} {
	return "niatelemetry.HttpsAclContractDetails"
}

// GetConsumerDn returns the ConsumerDn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetConsumerDn() string {
	if o == nil || IsNil(o.ConsumerDn) {
		var ret string
		return ret
	}
	return *o.ConsumerDn
}

// GetConsumerDnOk returns a tuple with the ConsumerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetConsumerDnOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerDn) {
		return nil, false
	}
	return o.ConsumerDn, true
}

// HasConsumerDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasConsumerDn() bool {
	if o != nil && !IsNil(o.ConsumerDn) {
		return true
	}

	return false
}

// SetConsumerDn gets a reference to the given string and assigns it to the ConsumerDn field.
func (o *NiatelemetryHttpsAclContractDetails) SetConsumerDn(v string) {
	o.ConsumerDn = &v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetContractName() string {
	if o == nil || IsNil(o.ContractName) {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetContractNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContractName) {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasContractName() bool {
	if o != nil && !IsNil(o.ContractName) {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *NiatelemetryHttpsAclContractDetails) SetContractName(v string) {
	o.ContractName = &v
}

// GetMgmtInstpDn returns the MgmtInstpDn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetMgmtInstpDn() string {
	if o == nil || IsNil(o.MgmtInstpDn) {
		var ret string
		return ret
	}
	return *o.MgmtInstpDn
}

// GetMgmtInstpDnOk returns a tuple with the MgmtInstpDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetMgmtInstpDnOk() (*string, bool) {
	if o == nil || IsNil(o.MgmtInstpDn) {
		return nil, false
	}
	return o.MgmtInstpDn, true
}

// HasMgmtInstpDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasMgmtInstpDn() bool {
	if o != nil && !IsNil(o.MgmtInstpDn) {
		return true
	}

	return false
}

// SetMgmtInstpDn gets a reference to the given string and assigns it to the MgmtInstpDn field.
func (o *NiatelemetryHttpsAclContractDetails) SetMgmtInstpDn(v string) {
	o.MgmtInstpDn = &v
}

// GetMgmtSubnetAddresses returns the MgmtSubnetAddresses field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetMgmtSubnetAddresses() string {
	if o == nil || IsNil(o.MgmtSubnetAddresses) {
		var ret string
		return ret
	}
	return *o.MgmtSubnetAddresses
}

// GetMgmtSubnetAddressesOk returns a tuple with the MgmtSubnetAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetMgmtSubnetAddressesOk() (*string, bool) {
	if o == nil || IsNil(o.MgmtSubnetAddresses) {
		return nil, false
	}
	return o.MgmtSubnetAddresses, true
}

// HasMgmtSubnetAddresses returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasMgmtSubnetAddresses() bool {
	if o != nil && !IsNil(o.MgmtSubnetAddresses) {
		return true
	}

	return false
}

// SetMgmtSubnetAddresses gets a reference to the given string and assigns it to the MgmtSubnetAddresses field.
func (o *NiatelemetryHttpsAclContractDetails) SetMgmtSubnetAddresses(v string) {
	o.MgmtSubnetAddresses = &v
}

// GetProviderDn returns the ProviderDn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetProviderDn() string {
	if o == nil || IsNil(o.ProviderDn) {
		var ret string
		return ret
	}
	return *o.ProviderDn
}

// GetProviderDnOk returns a tuple with the ProviderDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetProviderDnOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderDn) {
		return nil, false
	}
	return o.ProviderDn, true
}

// HasProviderDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasProviderDn() bool {
	if o != nil && !IsNil(o.ProviderDn) {
		return true
	}

	return false
}

// SetProviderDn gets a reference to the given string and assigns it to the ProviderDn field.
func (o *NiatelemetryHttpsAclContractDetails) SetProviderDn(v string) {
	o.ProviderDn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryHttpsAclContractDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryHttpsAclContractDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryHttpsAclContractDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryHttpsAclContractDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryHttpsAclContractDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryHttpsAclContractDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryHttpsAclContractDetails) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryHttpsAclContractDetails) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryHttpsAclContractDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryHttpsAclContractDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ConsumerDn) {
		toSerialize["ConsumerDn"] = o.ConsumerDn
	}
	if !IsNil(o.ContractName) {
		toSerialize["ContractName"] = o.ContractName
	}
	if !IsNil(o.MgmtInstpDn) {
		toSerialize["MgmtInstpDn"] = o.MgmtInstpDn
	}
	if !IsNil(o.MgmtSubnetAddresses) {
		toSerialize["MgmtSubnetAddresses"] = o.MgmtSubnetAddresses
	}
	if !IsNil(o.ProviderDn) {
		toSerialize["ProviderDn"] = o.ProviderDn
	}
	if !IsNil(o.RecordType) {
		toSerialize["RecordType"] = o.RecordType
	}
	if !IsNil(o.RecordVersion) {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if !IsNil(o.SiteName) {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryHttpsAclContractDetails) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Consumer Dn of the HTTPS ACL contract children MOs for APIC.
		ConsumerDn *string `json:"ConsumerDn,omitempty"`
		// Name of HTTPS ACL contract for APIC.
		ContractName *string `json:"ContractName,omitempty"`
		// Mgmt Inst Dn of the HTTPS ACL contract children MOs for APIC.
		MgmtInstpDn *string `json:"MgmtInstpDn,omitempty"`
		// Mgmt subnet address of the HTTPS ACL contract children MOs for APIC.
		MgmtSubnetAddresses *string `json:"MgmtSubnetAddresses,omitempty"`
		// Provider dn of the HTTPS ACL contract children MOs for APIC.
		ProviderDn *string `json:"ProviderDn,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                                     `json:"SiteName,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct := NiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryHttpsAclContractDetails := _NiatelemetryHttpsAclContractDetails{}
		varNiatelemetryHttpsAclContractDetails.ClassId = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryHttpsAclContractDetails.ObjectType = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryHttpsAclContractDetails.ConsumerDn = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.ConsumerDn
		varNiatelemetryHttpsAclContractDetails.ContractName = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.ContractName
		varNiatelemetryHttpsAclContractDetails.MgmtInstpDn = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.MgmtInstpDn
		varNiatelemetryHttpsAclContractDetails.MgmtSubnetAddresses = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.MgmtSubnetAddresses
		varNiatelemetryHttpsAclContractDetails.ProviderDn = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.ProviderDn
		varNiatelemetryHttpsAclContractDetails.RecordType = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryHttpsAclContractDetails.RecordVersion = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryHttpsAclContractDetails.SiteName = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryHttpsAclContractDetails.RegisteredDevice = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryHttpsAclContractDetails(varNiatelemetryHttpsAclContractDetails)
	} else {
		return err
	}

	varNiatelemetryHttpsAclContractDetails := _NiatelemetryHttpsAclContractDetails{}

	err = json.Unmarshal(data, &varNiatelemetryHttpsAclContractDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryHttpsAclContractDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConsumerDn")
		delete(additionalProperties, "ContractName")
		delete(additionalProperties, "MgmtInstpDn")
		delete(additionalProperties, "MgmtSubnetAddresses")
		delete(additionalProperties, "ProviderDn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
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

type NullableNiatelemetryHttpsAclContractDetails struct {
	value *NiatelemetryHttpsAclContractDetails
	isSet bool
}

func (v NullableNiatelemetryHttpsAclContractDetails) Get() *NiatelemetryHttpsAclContractDetails {
	return v.value
}

func (v *NullableNiatelemetryHttpsAclContractDetails) Set(val *NiatelemetryHttpsAclContractDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryHttpsAclContractDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryHttpsAclContractDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryHttpsAclContractDetails(val *NiatelemetryHttpsAclContractDetails) *NullableNiatelemetryHttpsAclContractDetails {
	return &NullableNiatelemetryHttpsAclContractDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryHttpsAclContractDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryHttpsAclContractDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
