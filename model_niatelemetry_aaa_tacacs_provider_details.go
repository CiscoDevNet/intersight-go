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

// checks if the NiatelemetryAaaTacacsProviderDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryAaaTacacsProviderDetails{}

// NiatelemetryAaaTacacsProviderDetails Object to capture AAA Tacacs provider details.
type NiatelemetryAaaTacacsProviderDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the AAA tacacs provider in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Owner Key of the AAA tacacs provider in APIC.
	OwnerKey *string `json:"OwnerKey,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                                     `json:"SiteName,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryAaaTacacsProviderDetails NiatelemetryAaaTacacsProviderDetails

// NewNiatelemetryAaaTacacsProviderDetails instantiates a new NiatelemetryAaaTacacsProviderDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryAaaTacacsProviderDetails(classId string, objectType string) *NiatelemetryAaaTacacsProviderDetails {
	this := NiatelemetryAaaTacacsProviderDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryAaaTacacsProviderDetailsWithDefaults instantiates a new NiatelemetryAaaTacacsProviderDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryAaaTacacsProviderDetailsWithDefaults() *NiatelemetryAaaTacacsProviderDetails {
	this := NiatelemetryAaaTacacsProviderDetails{}
	var classId string = "niatelemetry.AaaTacacsProviderDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.AaaTacacsProviderDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryAaaTacacsProviderDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryAaaTacacsProviderDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.AaaTacacsProviderDetails" of the ClassId field.
func (o *NiatelemetryAaaTacacsProviderDetails) GetDefaultClassId() interface{} {
	return "niatelemetry.AaaTacacsProviderDetails"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryAaaTacacsProviderDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryAaaTacacsProviderDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.AaaTacacsProviderDetails" of the ObjectType field.
func (o *NiatelemetryAaaTacacsProviderDetails) GetDefaultObjectType() interface{} {
	return "niatelemetry.AaaTacacsProviderDetails"
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryAaaTacacsProviderDetails) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryAaaTacacsProviderDetails) SetDn(v string) {
	o.Dn = &v
}

// GetOwnerKey returns the OwnerKey field value if set, zero value otherwise.
func (o *NiatelemetryAaaTacacsProviderDetails) GetOwnerKey() string {
	if o == nil || IsNil(o.OwnerKey) {
		var ret string
		return ret
	}
	return *o.OwnerKey
}

// GetOwnerKeyOk returns a tuple with the OwnerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) GetOwnerKeyOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerKey) {
		return nil, false
	}
	return o.OwnerKey, true
}

// HasOwnerKey returns a boolean if a field has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) HasOwnerKey() bool {
	if o != nil && !IsNil(o.OwnerKey) {
		return true
	}

	return false
}

// SetOwnerKey gets a reference to the given string and assigns it to the OwnerKey field.
func (o *NiatelemetryAaaTacacsProviderDetails) SetOwnerKey(v string) {
	o.OwnerKey = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryAaaTacacsProviderDetails) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryAaaTacacsProviderDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryAaaTacacsProviderDetails) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryAaaTacacsProviderDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryAaaTacacsProviderDetails) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryAaaTacacsProviderDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryAaaTacacsProviderDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryAaaTacacsProviderDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryAaaTacacsProviderDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryAaaTacacsProviderDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryAaaTacacsProviderDetails) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryAaaTacacsProviderDetails) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryAaaTacacsProviderDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryAaaTacacsProviderDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if !IsNil(o.OwnerKey) {
		toSerialize["OwnerKey"] = o.OwnerKey
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

func (o *NiatelemetryAaaTacacsProviderDetails) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryAaaTacacsProviderDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Dn of the AAA tacacs provider in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Owner Key of the AAA tacacs provider in APIC.
		OwnerKey *string `json:"OwnerKey,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                                     `json:"SiteName,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryAaaTacacsProviderDetailsWithoutEmbeddedStruct := NiatelemetryAaaTacacsProviderDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryAaaTacacsProviderDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryAaaTacacsProviderDetails := _NiatelemetryAaaTacacsProviderDetails{}
		varNiatelemetryAaaTacacsProviderDetails.ClassId = varNiatelemetryAaaTacacsProviderDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryAaaTacacsProviderDetails.ObjectType = varNiatelemetryAaaTacacsProviderDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryAaaTacacsProviderDetails.Dn = varNiatelemetryAaaTacacsProviderDetailsWithoutEmbeddedStruct.Dn
		varNiatelemetryAaaTacacsProviderDetails.OwnerKey = varNiatelemetryAaaTacacsProviderDetailsWithoutEmbeddedStruct.OwnerKey
		varNiatelemetryAaaTacacsProviderDetails.RecordType = varNiatelemetryAaaTacacsProviderDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryAaaTacacsProviderDetails.RecordVersion = varNiatelemetryAaaTacacsProviderDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryAaaTacacsProviderDetails.SiteName = varNiatelemetryAaaTacacsProviderDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryAaaTacacsProviderDetails.RegisteredDevice = varNiatelemetryAaaTacacsProviderDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryAaaTacacsProviderDetails(varNiatelemetryAaaTacacsProviderDetails)
	} else {
		return err
	}

	varNiatelemetryAaaTacacsProviderDetails := _NiatelemetryAaaTacacsProviderDetails{}

	err = json.Unmarshal(data, &varNiatelemetryAaaTacacsProviderDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryAaaTacacsProviderDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "OwnerKey")
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

type NullableNiatelemetryAaaTacacsProviderDetails struct {
	value *NiatelemetryAaaTacacsProviderDetails
	isSet bool
}

func (v NullableNiatelemetryAaaTacacsProviderDetails) Get() *NiatelemetryAaaTacacsProviderDetails {
	return v.value
}

func (v *NullableNiatelemetryAaaTacacsProviderDetails) Set(val *NiatelemetryAaaTacacsProviderDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryAaaTacacsProviderDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryAaaTacacsProviderDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryAaaTacacsProviderDetails(val *NiatelemetryAaaTacacsProviderDetails) *NullableNiatelemetryAaaTacacsProviderDetails {
	return &NullableNiatelemetryAaaTacacsProviderDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryAaaTacacsProviderDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryAaaTacacsProviderDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
