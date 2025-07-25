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

// checks if the NiatelemetryHttpsAclContractFilterMap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryHttpsAclContractFilterMap{}

// NiatelemetryHttpsAclContractFilterMap Object to capture the HTTPS ACL contract filter map in APIC.
type NiatelemetryHttpsAclContractFilterMap struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of HTTPS ACL contract for APIC.
	ContractName *string `json:"ContractName,omitempty"`
	// Dn of the HTTPS ACL EPGs subject filter relation MO for APIC.
	Dn *string `json:"Dn,omitempty"`
	// Name of HTTPS ACL filter for APIC.
	FilterName *string `json:"FilterName,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                                     `json:"SiteName,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryHttpsAclContractFilterMap NiatelemetryHttpsAclContractFilterMap

// NewNiatelemetryHttpsAclContractFilterMap instantiates a new NiatelemetryHttpsAclContractFilterMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryHttpsAclContractFilterMap(classId string, objectType string) *NiatelemetryHttpsAclContractFilterMap {
	this := NiatelemetryHttpsAclContractFilterMap{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryHttpsAclContractFilterMapWithDefaults instantiates a new NiatelemetryHttpsAclContractFilterMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryHttpsAclContractFilterMapWithDefaults() *NiatelemetryHttpsAclContractFilterMap {
	this := NiatelemetryHttpsAclContractFilterMap{}
	var classId string = "niatelemetry.HttpsAclContractFilterMap"
	this.ClassId = classId
	var objectType string = "niatelemetry.HttpsAclContractFilterMap"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryHttpsAclContractFilterMap) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryHttpsAclContractFilterMap) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.HttpsAclContractFilterMap" of the ClassId field.
func (o *NiatelemetryHttpsAclContractFilterMap) GetDefaultClassId() interface{} {
	return "niatelemetry.HttpsAclContractFilterMap"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryHttpsAclContractFilterMap) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryHttpsAclContractFilterMap) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.HttpsAclContractFilterMap" of the ObjectType field.
func (o *NiatelemetryHttpsAclContractFilterMap) GetDefaultObjectType() interface{} {
	return "niatelemetry.HttpsAclContractFilterMap"
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetContractName() string {
	if o == nil || IsNil(o.ContractName) {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetContractNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContractName) {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasContractName() bool {
	if o != nil && !IsNil(o.ContractName) {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetContractName(v string) {
	o.ContractName = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetDn(v string) {
	o.Dn = &v
}

// GetFilterName returns the FilterName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetFilterName() string {
	if o == nil || IsNil(o.FilterName) {
		var ret string
		return ret
	}
	return *o.FilterName
}

// GetFilterNameOk returns a tuple with the FilterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetFilterNameOk() (*string, bool) {
	if o == nil || IsNil(o.FilterName) {
		return nil, false
	}
	return o.FilterName, true
}

// HasFilterName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasFilterName() bool {
	if o != nil && !IsNil(o.FilterName) {
		return true
	}

	return false
}

// SetFilterName gets a reference to the given string and assigns it to the FilterName field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetFilterName(v string) {
	o.FilterName = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryHttpsAclContractFilterMap) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryHttpsAclContractFilterMap) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryHttpsAclContractFilterMap) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryHttpsAclContractFilterMap) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryHttpsAclContractFilterMap) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryHttpsAclContractFilterMap) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ContractName) {
		toSerialize["ContractName"] = o.ContractName
	}
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if !IsNil(o.FilterName) {
		toSerialize["FilterName"] = o.FilterName
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

func (o *NiatelemetryHttpsAclContractFilterMap) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of HTTPS ACL contract for APIC.
		ContractName *string `json:"ContractName,omitempty"`
		// Dn of the HTTPS ACL EPGs subject filter relation MO for APIC.
		Dn *string `json:"Dn,omitempty"`
		// Name of HTTPS ACL filter for APIC.
		FilterName *string `json:"FilterName,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                                     `json:"SiteName,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct := NiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryHttpsAclContractFilterMap := _NiatelemetryHttpsAclContractFilterMap{}
		varNiatelemetryHttpsAclContractFilterMap.ClassId = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.ClassId
		varNiatelemetryHttpsAclContractFilterMap.ObjectType = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.ObjectType
		varNiatelemetryHttpsAclContractFilterMap.ContractName = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.ContractName
		varNiatelemetryHttpsAclContractFilterMap.Dn = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.Dn
		varNiatelemetryHttpsAclContractFilterMap.FilterName = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.FilterName
		varNiatelemetryHttpsAclContractFilterMap.RecordType = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.RecordType
		varNiatelemetryHttpsAclContractFilterMap.RecordVersion = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryHttpsAclContractFilterMap.SiteName = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.SiteName
		varNiatelemetryHttpsAclContractFilterMap.RegisteredDevice = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryHttpsAclContractFilterMap(varNiatelemetryHttpsAclContractFilterMap)
	} else {
		return err
	}

	varNiatelemetryHttpsAclContractFilterMap := _NiatelemetryHttpsAclContractFilterMap{}

	err = json.Unmarshal(data, &varNiatelemetryHttpsAclContractFilterMap)
	if err == nil {
		o.MoBaseMo = varNiatelemetryHttpsAclContractFilterMap.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ContractName")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "FilterName")
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

type NullableNiatelemetryHttpsAclContractFilterMap struct {
	value *NiatelemetryHttpsAclContractFilterMap
	isSet bool
}

func (v NullableNiatelemetryHttpsAclContractFilterMap) Get() *NiatelemetryHttpsAclContractFilterMap {
	return v.value
}

func (v *NullableNiatelemetryHttpsAclContractFilterMap) Set(val *NiatelemetryHttpsAclContractFilterMap) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryHttpsAclContractFilterMap) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryHttpsAclContractFilterMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryHttpsAclContractFilterMap(val *NiatelemetryHttpsAclContractFilterMap) *NullableNiatelemetryHttpsAclContractFilterMap {
	return &NullableNiatelemetryHttpsAclContractFilterMap{value: val, isSet: true}
}

func (v NullableNiatelemetryHttpsAclContractFilterMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryHttpsAclContractFilterMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
