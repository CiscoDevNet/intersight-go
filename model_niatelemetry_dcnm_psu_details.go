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

// checks if the NiatelemetryDcnmPsuDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryDcnmPsuDetails{}

// NiatelemetryDcnmPsuDetails Inventory Object available for DCNM PSU.
type NiatelemetryDcnmPsuDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the power supply unit.
	Name *string `json:"Name,omitempty"`
	// Product ID of the power supply.
	ProductId *string `json:"ProductId,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Serial number of the power supply unit.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// Vendor Id of the power supply unit.
	VendorId             *string                                     `json:"VendorId,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryDcnmPsuDetails NiatelemetryDcnmPsuDetails

// NewNiatelemetryDcnmPsuDetails instantiates a new NiatelemetryDcnmPsuDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryDcnmPsuDetails(classId string, objectType string) *NiatelemetryDcnmPsuDetails {
	this := NiatelemetryDcnmPsuDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryDcnmPsuDetailsWithDefaults instantiates a new NiatelemetryDcnmPsuDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryDcnmPsuDetailsWithDefaults() *NiatelemetryDcnmPsuDetails {
	this := NiatelemetryDcnmPsuDetails{}
	var classId string = "niatelemetry.DcnmPsuDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.DcnmPsuDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryDcnmPsuDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmPsuDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryDcnmPsuDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.DcnmPsuDetails" of the ClassId field.
func (o *NiatelemetryDcnmPsuDetails) GetDefaultClassId() interface{} {
	return "niatelemetry.DcnmPsuDetails"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryDcnmPsuDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmPsuDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryDcnmPsuDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.DcnmPsuDetails" of the ObjectType field.
func (o *NiatelemetryDcnmPsuDetails) GetDefaultObjectType() interface{} {
	return "niatelemetry.DcnmPsuDetails"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetryDcnmPsuDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmPsuDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetryDcnmPsuDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetryDcnmPsuDetails) SetName(v string) {
	o.Name = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *NiatelemetryDcnmPsuDetails) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmPsuDetails) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *NiatelemetryDcnmPsuDetails) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *NiatelemetryDcnmPsuDetails) SetProductId(v string) {
	o.ProductId = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryDcnmPsuDetails) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmPsuDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryDcnmPsuDetails) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryDcnmPsuDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryDcnmPsuDetails) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmPsuDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryDcnmPsuDetails) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryDcnmPsuDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *NiatelemetryDcnmPsuDetails) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmPsuDetails) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *NiatelemetryDcnmPsuDetails) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *NiatelemetryDcnmPsuDetails) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *NiatelemetryDcnmPsuDetails) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmPsuDetails) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *NiatelemetryDcnmPsuDetails) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *NiatelemetryDcnmPsuDetails) SetVendorId(v string) {
	o.VendorId = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryDcnmPsuDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryDcnmPsuDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryDcnmPsuDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryDcnmPsuDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryDcnmPsuDetails) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryDcnmPsuDetails) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryDcnmPsuDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryDcnmPsuDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ProductId) {
		toSerialize["ProductId"] = o.ProductId
	}
	if !IsNil(o.RecordType) {
		toSerialize["RecordType"] = o.RecordType
	}
	if !IsNil(o.RecordVersion) {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if !IsNil(o.VendorId) {
		toSerialize["VendorId"] = o.VendorId
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryDcnmPsuDetails) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the power supply unit.
		Name *string `json:"Name,omitempty"`
		// Product ID of the power supply.
		ProductId *string `json:"ProductId,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Serial number of the power supply unit.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// Vendor Id of the power supply unit.
		VendorId         *string                                     `json:"VendorId,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct := NiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryDcnmPsuDetails := _NiatelemetryDcnmPsuDetails{}
		varNiatelemetryDcnmPsuDetails.ClassId = varNiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryDcnmPsuDetails.ObjectType = varNiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryDcnmPsuDetails.Name = varNiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct.Name
		varNiatelemetryDcnmPsuDetails.ProductId = varNiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct.ProductId
		varNiatelemetryDcnmPsuDetails.RecordType = varNiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryDcnmPsuDetails.RecordVersion = varNiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryDcnmPsuDetails.SerialNumber = varNiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct.SerialNumber
		varNiatelemetryDcnmPsuDetails.VendorId = varNiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct.VendorId
		varNiatelemetryDcnmPsuDetails.RegisteredDevice = varNiatelemetryDcnmPsuDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryDcnmPsuDetails(varNiatelemetryDcnmPsuDetails)
	} else {
		return err
	}

	varNiatelemetryDcnmPsuDetails := _NiatelemetryDcnmPsuDetails{}

	err = json.Unmarshal(data, &varNiatelemetryDcnmPsuDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryDcnmPsuDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ProductId")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "VendorId")
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

type NullableNiatelemetryDcnmPsuDetails struct {
	value *NiatelemetryDcnmPsuDetails
	isSet bool
}

func (v NullableNiatelemetryDcnmPsuDetails) Get() *NiatelemetryDcnmPsuDetails {
	return v.value
}

func (v *NullableNiatelemetryDcnmPsuDetails) Set(val *NiatelemetryDcnmPsuDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryDcnmPsuDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryDcnmPsuDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryDcnmPsuDetails(val *NiatelemetryDcnmPsuDetails) *NullableNiatelemetryDcnmPsuDetails {
	return &NullableNiatelemetryDcnmPsuDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryDcnmPsuDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryDcnmPsuDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
