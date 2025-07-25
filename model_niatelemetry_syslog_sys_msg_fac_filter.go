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

// checks if the NiatelemetrySyslogSysMsgFacFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetrySyslogSysMsgFacFilter{}

// NiatelemetrySyslogSysMsgFacFilter Object to capture Syslog system Msg details.
type NiatelemetrySyslogSysMsgFacFilter struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Parent common policy for syslog system msg in APIC.
	CommonPolicy *string `json:"CommonPolicy,omitempty"`
	// Dn of the Syslog System msg facility filter in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Facility of Syslog System msg facility filter in APIC.
	Facility *string `json:"Facility,omitempty"`
	// Minimum severity of Syslog System msg facility filter in APIC.
	MinSev *string `json:"MinSev,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// Parent syslog msg for syslog sys msg facility filter in APIC.
	SyslogSysMsg         *string                                     `json:"SyslogSysMsg,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetrySyslogSysMsgFacFilter NiatelemetrySyslogSysMsgFacFilter

// NewNiatelemetrySyslogSysMsgFacFilter instantiates a new NiatelemetrySyslogSysMsgFacFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetrySyslogSysMsgFacFilter(classId string, objectType string) *NiatelemetrySyslogSysMsgFacFilter {
	this := NiatelemetrySyslogSysMsgFacFilter{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetrySyslogSysMsgFacFilterWithDefaults instantiates a new NiatelemetrySyslogSysMsgFacFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetrySyslogSysMsgFacFilterWithDefaults() *NiatelemetrySyslogSysMsgFacFilter {
	this := NiatelemetrySyslogSysMsgFacFilter{}
	var classId string = "niatelemetry.SyslogSysMsgFacFilter"
	this.ClassId = classId
	var objectType string = "niatelemetry.SyslogSysMsgFacFilter"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetrySyslogSysMsgFacFilter) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetrySyslogSysMsgFacFilter) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.SyslogSysMsgFacFilter" of the ClassId field.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetDefaultClassId() interface{} {
	return "niatelemetry.SyslogSysMsgFacFilter"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetrySyslogSysMsgFacFilter) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetrySyslogSysMsgFacFilter) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.SyslogSysMsgFacFilter" of the ObjectType field.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetDefaultObjectType() interface{} {
	return "niatelemetry.SyslogSysMsgFacFilter"
}

// GetCommonPolicy returns the CommonPolicy field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetCommonPolicy() string {
	if o == nil || IsNil(o.CommonPolicy) {
		var ret string
		return ret
	}
	return *o.CommonPolicy
}

// GetCommonPolicyOk returns a tuple with the CommonPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetCommonPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.CommonPolicy) {
		return nil, false
	}
	return o.CommonPolicy, true
}

// HasCommonPolicy returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) HasCommonPolicy() bool {
	if o != nil && !IsNil(o.CommonPolicy) {
		return true
	}

	return false
}

// SetCommonPolicy gets a reference to the given string and assigns it to the CommonPolicy field.
func (o *NiatelemetrySyslogSysMsgFacFilter) SetCommonPolicy(v string) {
	o.CommonPolicy = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetrySyslogSysMsgFacFilter) SetDn(v string) {
	o.Dn = &v
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetFacility() string {
	if o == nil || IsNil(o.Facility) {
		var ret string
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetFacilityOk() (*string, bool) {
	if o == nil || IsNil(o.Facility) {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) HasFacility() bool {
	if o != nil && !IsNil(o.Facility) {
		return true
	}

	return false
}

// SetFacility gets a reference to the given string and assigns it to the Facility field.
func (o *NiatelemetrySyslogSysMsgFacFilter) SetFacility(v string) {
	o.Facility = &v
}

// GetMinSev returns the MinSev field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetMinSev() string {
	if o == nil || IsNil(o.MinSev) {
		var ret string
		return ret
	}
	return *o.MinSev
}

// GetMinSevOk returns a tuple with the MinSev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetMinSevOk() (*string, bool) {
	if o == nil || IsNil(o.MinSev) {
		return nil, false
	}
	return o.MinSev, true
}

// HasMinSev returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) HasMinSev() bool {
	if o != nil && !IsNil(o.MinSev) {
		return true
	}

	return false
}

// SetMinSev gets a reference to the given string and assigns it to the MinSev field.
func (o *NiatelemetrySyslogSysMsgFacFilter) SetMinSev(v string) {
	o.MinSev = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetrySyslogSysMsgFacFilter) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetrySyslogSysMsgFacFilter) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetrySyslogSysMsgFacFilter) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSyslogSysMsg returns the SyslogSysMsg field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetSyslogSysMsg() string {
	if o == nil || IsNil(o.SyslogSysMsg) {
		var ret string
		return ret
	}
	return *o.SyslogSysMsg
}

// GetSyslogSysMsgOk returns a tuple with the SyslogSysMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) GetSyslogSysMsgOk() (*string, bool) {
	if o == nil || IsNil(o.SyslogSysMsg) {
		return nil, false
	}
	return o.SyslogSysMsg, true
}

// HasSyslogSysMsg returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) HasSyslogSysMsg() bool {
	if o != nil && !IsNil(o.SyslogSysMsg) {
		return true
	}

	return false
}

// SetSyslogSysMsg gets a reference to the given string and assigns it to the SyslogSysMsg field.
func (o *NiatelemetrySyslogSysMsgFacFilter) SetSyslogSysMsg(v string) {
	o.SyslogSysMsg = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetrySyslogSysMsgFacFilter) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetrySyslogSysMsgFacFilter) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilter) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetrySyslogSysMsgFacFilter) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetrySyslogSysMsgFacFilter) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetrySyslogSysMsgFacFilter) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetrySyslogSysMsgFacFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetrySyslogSysMsgFacFilter) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CommonPolicy) {
		toSerialize["CommonPolicy"] = o.CommonPolicy
	}
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if !IsNil(o.Facility) {
		toSerialize["Facility"] = o.Facility
	}
	if !IsNil(o.MinSev) {
		toSerialize["MinSev"] = o.MinSev
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
	if !IsNil(o.SyslogSysMsg) {
		toSerialize["SyslogSysMsg"] = o.SyslogSysMsg
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetrySyslogSysMsgFacFilter) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Parent common policy for syslog system msg in APIC.
		CommonPolicy *string `json:"CommonPolicy,omitempty"`
		// Dn of the Syslog System msg facility filter in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Facility of Syslog System msg facility filter in APIC.
		Facility *string `json:"Facility,omitempty"`
		// Minimum severity of Syslog System msg facility filter in APIC.
		MinSev *string `json:"MinSev,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// Parent syslog msg for syslog sys msg facility filter in APIC.
		SyslogSysMsg     *string                                     `json:"SyslogSysMsg,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct := NiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetrySyslogSysMsgFacFilter := _NiatelemetrySyslogSysMsgFacFilter{}
		varNiatelemetrySyslogSysMsgFacFilter.ClassId = varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct.ClassId
		varNiatelemetrySyslogSysMsgFacFilter.ObjectType = varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct.ObjectType
		varNiatelemetrySyslogSysMsgFacFilter.CommonPolicy = varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct.CommonPolicy
		varNiatelemetrySyslogSysMsgFacFilter.Dn = varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct.Dn
		varNiatelemetrySyslogSysMsgFacFilter.Facility = varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct.Facility
		varNiatelemetrySyslogSysMsgFacFilter.MinSev = varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct.MinSev
		varNiatelemetrySyslogSysMsgFacFilter.RecordType = varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct.RecordType
		varNiatelemetrySyslogSysMsgFacFilter.RecordVersion = varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct.RecordVersion
		varNiatelemetrySyslogSysMsgFacFilter.SiteName = varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct.SiteName
		varNiatelemetrySyslogSysMsgFacFilter.SyslogSysMsg = varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct.SyslogSysMsg
		varNiatelemetrySyslogSysMsgFacFilter.RegisteredDevice = varNiatelemetrySyslogSysMsgFacFilterWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetrySyslogSysMsgFacFilter(varNiatelemetrySyslogSysMsgFacFilter)
	} else {
		return err
	}

	varNiatelemetrySyslogSysMsgFacFilter := _NiatelemetrySyslogSysMsgFacFilter{}

	err = json.Unmarshal(data, &varNiatelemetrySyslogSysMsgFacFilter)
	if err == nil {
		o.MoBaseMo = varNiatelemetrySyslogSysMsgFacFilter.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CommonPolicy")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "Facility")
		delete(additionalProperties, "MinSev")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SyslogSysMsg")
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

type NullableNiatelemetrySyslogSysMsgFacFilter struct {
	value *NiatelemetrySyslogSysMsgFacFilter
	isSet bool
}

func (v NullableNiatelemetrySyslogSysMsgFacFilter) Get() *NiatelemetrySyslogSysMsgFacFilter {
	return v.value
}

func (v *NullableNiatelemetrySyslogSysMsgFacFilter) Set(val *NiatelemetrySyslogSysMsgFacFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetrySyslogSysMsgFacFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetrySyslogSysMsgFacFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetrySyslogSysMsgFacFilter(val *NiatelemetrySyslogSysMsgFacFilter) *NullableNiatelemetrySyslogSysMsgFacFilter {
	return &NullableNiatelemetrySyslogSysMsgFacFilter{value: val, isSet: true}
}

func (v NullableNiatelemetrySyslogSysMsgFacFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetrySyslogSysMsgFacFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
