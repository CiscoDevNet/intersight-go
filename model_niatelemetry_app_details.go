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

// checks if the NiatelemetryAppDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryAppDetails{}

// NiatelemetryAppDetails Details of apps installed on Nexus Dashboard.
type NiatelemetryAppDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Names of apps running on ND.
	AppName *string `json:"AppName,omitempty"`
	// Status of apps running on ND.
	AppStatus *string `json:"AppStatus,omitempty"`
	// Versions of apps running on ND.
	AppVersion *string `json:"AppVersion,omitempty"`
	// Clustername on which apps are running on ND.
	NexusDashboard *string `json:"NexusDashboard,omitempty"`
	// Number of sites on which particular app installed on ND.
	NumberOfSitesOnboarded *int64 `json:"NumberOfSitesOnboarded,omitempty"`
	// Type of apps running on ND.
	Type                 *string                                     `json:"Type,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryAppDetails NiatelemetryAppDetails

// NewNiatelemetryAppDetails instantiates a new NiatelemetryAppDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryAppDetails(classId string, objectType string) *NiatelemetryAppDetails {
	this := NiatelemetryAppDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryAppDetailsWithDefaults instantiates a new NiatelemetryAppDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryAppDetailsWithDefaults() *NiatelemetryAppDetails {
	this := NiatelemetryAppDetails{}
	var classId string = "niatelemetry.AppDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.AppDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryAppDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryAppDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.AppDetails" of the ClassId field.
func (o *NiatelemetryAppDetails) GetDefaultClassId() interface{} {
	return "niatelemetry.AppDetails"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryAppDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryAppDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.AppDetails" of the ObjectType field.
func (o *NiatelemetryAppDetails) GetDefaultObjectType() interface{} {
	return "niatelemetry.AppDetails"
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *NiatelemetryAppDetails) SetAppName(v string) {
	o.AppName = &v
}

// GetAppStatus returns the AppStatus field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetAppStatus() string {
	if o == nil || IsNil(o.AppStatus) {
		var ret string
		return ret
	}
	return *o.AppStatus
}

// GetAppStatusOk returns a tuple with the AppStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetAppStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AppStatus) {
		return nil, false
	}
	return o.AppStatus, true
}

// HasAppStatus returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasAppStatus() bool {
	if o != nil && !IsNil(o.AppStatus) {
		return true
	}

	return false
}

// SetAppStatus gets a reference to the given string and assigns it to the AppStatus field.
func (o *NiatelemetryAppDetails) SetAppStatus(v string) {
	o.AppStatus = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *NiatelemetryAppDetails) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetNexusDashboard returns the NexusDashboard field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetNexusDashboard() string {
	if o == nil || IsNil(o.NexusDashboard) {
		var ret string
		return ret
	}
	return *o.NexusDashboard
}

// GetNexusDashboardOk returns a tuple with the NexusDashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetNexusDashboardOk() (*string, bool) {
	if o == nil || IsNil(o.NexusDashboard) {
		return nil, false
	}
	return o.NexusDashboard, true
}

// HasNexusDashboard returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasNexusDashboard() bool {
	if o != nil && !IsNil(o.NexusDashboard) {
		return true
	}

	return false
}

// SetNexusDashboard gets a reference to the given string and assigns it to the NexusDashboard field.
func (o *NiatelemetryAppDetails) SetNexusDashboard(v string) {
	o.NexusDashboard = &v
}

// GetNumberOfSitesOnboarded returns the NumberOfSitesOnboarded field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetNumberOfSitesOnboarded() int64 {
	if o == nil || IsNil(o.NumberOfSitesOnboarded) {
		var ret int64
		return ret
	}
	return *o.NumberOfSitesOnboarded
}

// GetNumberOfSitesOnboardedOk returns a tuple with the NumberOfSitesOnboarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetNumberOfSitesOnboardedOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfSitesOnboarded) {
		return nil, false
	}
	return o.NumberOfSitesOnboarded, true
}

// HasNumberOfSitesOnboarded returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasNumberOfSitesOnboarded() bool {
	if o != nil && !IsNil(o.NumberOfSitesOnboarded) {
		return true
	}

	return false
}

// SetNumberOfSitesOnboarded gets a reference to the given int64 and assigns it to the NumberOfSitesOnboarded field.
func (o *NiatelemetryAppDetails) SetNumberOfSitesOnboarded(v int64) {
	o.NumberOfSitesOnboarded = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NiatelemetryAppDetails) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetails) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NiatelemetryAppDetails) SetType(v string) {
	o.Type = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryAppDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryAppDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryAppDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryAppDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryAppDetails) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryAppDetails) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryAppDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryAppDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AppName) {
		toSerialize["AppName"] = o.AppName
	}
	if !IsNil(o.AppStatus) {
		toSerialize["AppStatus"] = o.AppStatus
	}
	if !IsNil(o.AppVersion) {
		toSerialize["AppVersion"] = o.AppVersion
	}
	if !IsNil(o.NexusDashboard) {
		toSerialize["NexusDashboard"] = o.NexusDashboard
	}
	if !IsNil(o.NumberOfSitesOnboarded) {
		toSerialize["NumberOfSitesOnboarded"] = o.NumberOfSitesOnboarded
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryAppDetails) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryAppDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Names of apps running on ND.
		AppName *string `json:"AppName,omitempty"`
		// Status of apps running on ND.
		AppStatus *string `json:"AppStatus,omitempty"`
		// Versions of apps running on ND.
		AppVersion *string `json:"AppVersion,omitempty"`
		// Clustername on which apps are running on ND.
		NexusDashboard *string `json:"NexusDashboard,omitempty"`
		// Number of sites on which particular app installed on ND.
		NumberOfSitesOnboarded *int64 `json:"NumberOfSitesOnboarded,omitempty"`
		// Type of apps running on ND.
		Type             *string                                     `json:"Type,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryAppDetailsWithoutEmbeddedStruct := NiatelemetryAppDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryAppDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryAppDetails := _NiatelemetryAppDetails{}
		varNiatelemetryAppDetails.ClassId = varNiatelemetryAppDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryAppDetails.ObjectType = varNiatelemetryAppDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryAppDetails.AppName = varNiatelemetryAppDetailsWithoutEmbeddedStruct.AppName
		varNiatelemetryAppDetails.AppStatus = varNiatelemetryAppDetailsWithoutEmbeddedStruct.AppStatus
		varNiatelemetryAppDetails.AppVersion = varNiatelemetryAppDetailsWithoutEmbeddedStruct.AppVersion
		varNiatelemetryAppDetails.NexusDashboard = varNiatelemetryAppDetailsWithoutEmbeddedStruct.NexusDashboard
		varNiatelemetryAppDetails.NumberOfSitesOnboarded = varNiatelemetryAppDetailsWithoutEmbeddedStruct.NumberOfSitesOnboarded
		varNiatelemetryAppDetails.Type = varNiatelemetryAppDetailsWithoutEmbeddedStruct.Type
		varNiatelemetryAppDetails.RegisteredDevice = varNiatelemetryAppDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryAppDetails(varNiatelemetryAppDetails)
	} else {
		return err
	}

	varNiatelemetryAppDetails := _NiatelemetryAppDetails{}

	err = json.Unmarshal(data, &varNiatelemetryAppDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryAppDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AppName")
		delete(additionalProperties, "AppStatus")
		delete(additionalProperties, "AppVersion")
		delete(additionalProperties, "NexusDashboard")
		delete(additionalProperties, "NumberOfSitesOnboarded")
		delete(additionalProperties, "Type")
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

type NullableNiatelemetryAppDetails struct {
	value *NiatelemetryAppDetails
	isSet bool
}

func (v NullableNiatelemetryAppDetails) Get() *NiatelemetryAppDetails {
	return v.value
}

func (v *NullableNiatelemetryAppDetails) Set(val *NiatelemetryAppDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryAppDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryAppDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryAppDetails(val *NiatelemetryAppDetails) *NullableNiatelemetryAppDetails {
	return &NullableNiatelemetryAppDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryAppDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryAppDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
