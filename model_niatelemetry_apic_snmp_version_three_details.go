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

// checks if the NiatelemetryApicSnmpVersionThreeDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryApicSnmpVersionThreeDetails{}

// NiatelemetryApicSnmpVersionThreeDetails Object to capture Snmp V3 details in APIC.
type NiatelemetryApicSnmpVersionThreeDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// AuthType of SNMP V3 in APIC.
	AuthType *string `json:"AuthType,omitempty"`
	// Dn of SNMP V3 attribute in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Name of SNMP V3 attribute in APIC.
	Name *string `json:"Name,omitempty"`
	// PrivType of SNMP V3 in APIC.
	PrivType *string `json:"PrivType,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                                     `json:"SiteName,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicSnmpVersionThreeDetails NiatelemetryApicSnmpVersionThreeDetails

// NewNiatelemetryApicSnmpVersionThreeDetails instantiates a new NiatelemetryApicSnmpVersionThreeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicSnmpVersionThreeDetails(classId string, objectType string) *NiatelemetryApicSnmpVersionThreeDetails {
	this := NiatelemetryApicSnmpVersionThreeDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicSnmpVersionThreeDetailsWithDefaults instantiates a new NiatelemetryApicSnmpVersionThreeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicSnmpVersionThreeDetailsWithDefaults() *NiatelemetryApicSnmpVersionThreeDetails {
	this := NiatelemetryApicSnmpVersionThreeDetails{}
	var classId string = "niatelemetry.ApicSnmpVersionThreeDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicSnmpVersionThreeDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.ApicSnmpVersionThreeDetails" of the ClassId field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetDefaultClassId() interface{} {
	return "niatelemetry.ApicSnmpVersionThreeDetails"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.ApicSnmpVersionThreeDetails" of the ObjectType field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetDefaultObjectType() interface{} {
	return "niatelemetry.ApicSnmpVersionThreeDetails"
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetAuthType() string {
	if o == nil || IsNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetAuthType(v string) {
	o.AuthType = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetDn(v string) {
	o.Dn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetName(v string) {
	o.Name = &v
}

// GetPrivType returns the PrivType field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetPrivType() string {
	if o == nil || IsNil(o.PrivType) {
		var ret string
		return ret
	}
	return *o.PrivType
}

// GetPrivTypeOk returns a tuple with the PrivType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetPrivTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PrivType) {
		return nil, false
	}
	return o.PrivType, true
}

// HasPrivType returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasPrivType() bool {
	if o != nil && !IsNil(o.PrivType) {
		return true
	}

	return false
}

// SetPrivType gets a reference to the given string and assigns it to the PrivType field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetPrivType(v string) {
	o.PrivType = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpVersionThreeDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryApicSnmpVersionThreeDetails) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryApicSnmpVersionThreeDetails) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryApicSnmpVersionThreeDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryApicSnmpVersionThreeDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AuthType) {
		toSerialize["AuthType"] = o.AuthType
	}
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PrivType) {
		toSerialize["PrivType"] = o.PrivType
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

func (o *NiatelemetryApicSnmpVersionThreeDetails) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// AuthType of SNMP V3 in APIC.
		AuthType *string `json:"AuthType,omitempty"`
		// Dn of SNMP V3 attribute in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Name of SNMP V3 attribute in APIC.
		Name *string `json:"Name,omitempty"`
		// PrivType of SNMP V3 in APIC.
		PrivType *string `json:"PrivType,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                                     `json:"SiteName,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct := NiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryApicSnmpVersionThreeDetails := _NiatelemetryApicSnmpVersionThreeDetails{}
		varNiatelemetryApicSnmpVersionThreeDetails.ClassId = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryApicSnmpVersionThreeDetails.ObjectType = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryApicSnmpVersionThreeDetails.AuthType = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.AuthType
		varNiatelemetryApicSnmpVersionThreeDetails.Dn = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.Dn
		varNiatelemetryApicSnmpVersionThreeDetails.Name = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.Name
		varNiatelemetryApicSnmpVersionThreeDetails.PrivType = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.PrivType
		varNiatelemetryApicSnmpVersionThreeDetails.RecordType = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryApicSnmpVersionThreeDetails.RecordVersion = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryApicSnmpVersionThreeDetails.SiteName = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryApicSnmpVersionThreeDetails.RegisteredDevice = varNiatelemetryApicSnmpVersionThreeDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryApicSnmpVersionThreeDetails(varNiatelemetryApicSnmpVersionThreeDetails)
	} else {
		return err
	}

	varNiatelemetryApicSnmpVersionThreeDetails := _NiatelemetryApicSnmpVersionThreeDetails{}

	err = json.Unmarshal(data, &varNiatelemetryApicSnmpVersionThreeDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryApicSnmpVersionThreeDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PrivType")
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

type NullableNiatelemetryApicSnmpVersionThreeDetails struct {
	value *NiatelemetryApicSnmpVersionThreeDetails
	isSet bool
}

func (v NullableNiatelemetryApicSnmpVersionThreeDetails) Get() *NiatelemetryApicSnmpVersionThreeDetails {
	return v.value
}

func (v *NullableNiatelemetryApicSnmpVersionThreeDetails) Set(val *NiatelemetryApicSnmpVersionThreeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicSnmpVersionThreeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicSnmpVersionThreeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicSnmpVersionThreeDetails(val *NiatelemetryApicSnmpVersionThreeDetails) *NullableNiatelemetryApicSnmpVersionThreeDetails {
	return &NullableNiatelemetryApicSnmpVersionThreeDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryApicSnmpVersionThreeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicSnmpVersionThreeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
