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

// checks if the HciDomainManager type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HciDomainManager{}

// HciDomainManager The current HCI Prism Central instance reported by Prism Central.
type HciDomainManager struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string        `json:"ObjectType"`
	ApiLimits  []HciApiLimit `json:"ApiLimits,omitempty"`
	// The string representation of the API limits as a string. It can be used by Alarm.
	ApiLimitsString *string `json:"ApiLimitsString,omitempty"`
	// The name of the domain manager.
	Name *string `json:"Name,omitempty"`
	// The unique identifier of the domain manager (Prism Central) instance.
	PcExtId *string `json:"PcExtId,omitempty"`
	// The size of the domain manager such as STARTER, SMALL, LARGE, EXTRALARGE. It determines the resources used by the domain manager.
	Size *string `json:"Size,omitempty"`
	// An array of relationships to hciCluster resources.
	Clusters []HciClusterRelationship `json:"Clusters,omitempty"`
	// An array of relationships to hciLicense resources.
	Licenses             []HciLicenseRelationship                    `json:"Licenses,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HciDomainManager HciDomainManager

// NewHciDomainManager instantiates a new HciDomainManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHciDomainManager(classId string, objectType string) *HciDomainManager {
	this := HciDomainManager{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHciDomainManagerWithDefaults instantiates a new HciDomainManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHciDomainManagerWithDefaults() *HciDomainManager {
	this := HciDomainManager{}
	var classId string = "hci.DomainManager"
	this.ClassId = classId
	var objectType string = "hci.DomainManager"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HciDomainManager) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HciDomainManager) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HciDomainManager) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hci.DomainManager" of the ClassId field.
func (o *HciDomainManager) GetDefaultClassId() interface{} {
	return "hci.DomainManager"
}

// GetObjectType returns the ObjectType field value
func (o *HciDomainManager) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HciDomainManager) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HciDomainManager) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hci.DomainManager" of the ObjectType field.
func (o *HciDomainManager) GetDefaultObjectType() interface{} {
	return "hci.DomainManager"
}

// GetApiLimits returns the ApiLimits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HciDomainManager) GetApiLimits() []HciApiLimit {
	if o == nil {
		var ret []HciApiLimit
		return ret
	}
	return o.ApiLimits
}

// GetApiLimitsOk returns a tuple with the ApiLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HciDomainManager) GetApiLimitsOk() ([]HciApiLimit, bool) {
	if o == nil || IsNil(o.ApiLimits) {
		return nil, false
	}
	return o.ApiLimits, true
}

// HasApiLimits returns a boolean if a field has been set.
func (o *HciDomainManager) HasApiLimits() bool {
	if o != nil && !IsNil(o.ApiLimits) {
		return true
	}

	return false
}

// SetApiLimits gets a reference to the given []HciApiLimit and assigns it to the ApiLimits field.
func (o *HciDomainManager) SetApiLimits(v []HciApiLimit) {
	o.ApiLimits = v
}

// GetApiLimitsString returns the ApiLimitsString field value if set, zero value otherwise.
func (o *HciDomainManager) GetApiLimitsString() string {
	if o == nil || IsNil(o.ApiLimitsString) {
		var ret string
		return ret
	}
	return *o.ApiLimitsString
}

// GetApiLimitsStringOk returns a tuple with the ApiLimitsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciDomainManager) GetApiLimitsStringOk() (*string, bool) {
	if o == nil || IsNil(o.ApiLimitsString) {
		return nil, false
	}
	return o.ApiLimitsString, true
}

// HasApiLimitsString returns a boolean if a field has been set.
func (o *HciDomainManager) HasApiLimitsString() bool {
	if o != nil && !IsNil(o.ApiLimitsString) {
		return true
	}

	return false
}

// SetApiLimitsString gets a reference to the given string and assigns it to the ApiLimitsString field.
func (o *HciDomainManager) SetApiLimitsString(v string) {
	o.ApiLimitsString = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HciDomainManager) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciDomainManager) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HciDomainManager) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HciDomainManager) SetName(v string) {
	o.Name = &v
}

// GetPcExtId returns the PcExtId field value if set, zero value otherwise.
func (o *HciDomainManager) GetPcExtId() string {
	if o == nil || IsNil(o.PcExtId) {
		var ret string
		return ret
	}
	return *o.PcExtId
}

// GetPcExtIdOk returns a tuple with the PcExtId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciDomainManager) GetPcExtIdOk() (*string, bool) {
	if o == nil || IsNil(o.PcExtId) {
		return nil, false
	}
	return o.PcExtId, true
}

// HasPcExtId returns a boolean if a field has been set.
func (o *HciDomainManager) HasPcExtId() bool {
	if o != nil && !IsNil(o.PcExtId) {
		return true
	}

	return false
}

// SetPcExtId gets a reference to the given string and assigns it to the PcExtId field.
func (o *HciDomainManager) SetPcExtId(v string) {
	o.PcExtId = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *HciDomainManager) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciDomainManager) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *HciDomainManager) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *HciDomainManager) SetSize(v string) {
	o.Size = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HciDomainManager) GetClusters() []HciClusterRelationship {
	if o == nil {
		var ret []HciClusterRelationship
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HciDomainManager) GetClustersOk() ([]HciClusterRelationship, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *HciDomainManager) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []HciClusterRelationship and assigns it to the Clusters field.
func (o *HciDomainManager) SetClusters(v []HciClusterRelationship) {
	o.Clusters = v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HciDomainManager) GetLicenses() []HciLicenseRelationship {
	if o == nil {
		var ret []HciLicenseRelationship
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HciDomainManager) GetLicensesOk() ([]HciLicenseRelationship, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *HciDomainManager) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []HciLicenseRelationship and assigns it to the Licenses field.
func (o *HciDomainManager) SetLicenses(v []HciLicenseRelationship) {
	o.Licenses = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HciDomainManager) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HciDomainManager) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HciDomainManager) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HciDomainManager) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *HciDomainManager) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *HciDomainManager) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o HciDomainManager) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HciDomainManager) ToMap() (map[string]interface{}, error) {
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
	if o.ApiLimits != nil {
		toSerialize["ApiLimits"] = o.ApiLimits
	}
	if !IsNil(o.ApiLimitsString) {
		toSerialize["ApiLimitsString"] = o.ApiLimitsString
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PcExtId) {
		toSerialize["PcExtId"] = o.PcExtId
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	if o.Clusters != nil {
		toSerialize["Clusters"] = o.Clusters
	}
	if o.Licenses != nil {
		toSerialize["Licenses"] = o.Licenses
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HciDomainManager) UnmarshalJSON(data []byte) (err error) {
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
	type HciDomainManagerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string        `json:"ObjectType"`
		ApiLimits  []HciApiLimit `json:"ApiLimits,omitempty"`
		// The string representation of the API limits as a string. It can be used by Alarm.
		ApiLimitsString *string `json:"ApiLimitsString,omitempty"`
		// The name of the domain manager.
		Name *string `json:"Name,omitempty"`
		// The unique identifier of the domain manager (Prism Central) instance.
		PcExtId *string `json:"PcExtId,omitempty"`
		// The size of the domain manager such as STARTER, SMALL, LARGE, EXTRALARGE. It determines the resources used by the domain manager.
		Size *string `json:"Size,omitempty"`
		// An array of relationships to hciCluster resources.
		Clusters []HciClusterRelationship `json:"Clusters,omitempty"`
		// An array of relationships to hciLicense resources.
		Licenses         []HciLicenseRelationship                    `json:"Licenses,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varHciDomainManagerWithoutEmbeddedStruct := HciDomainManagerWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHciDomainManagerWithoutEmbeddedStruct)
	if err == nil {
		varHciDomainManager := _HciDomainManager{}
		varHciDomainManager.ClassId = varHciDomainManagerWithoutEmbeddedStruct.ClassId
		varHciDomainManager.ObjectType = varHciDomainManagerWithoutEmbeddedStruct.ObjectType
		varHciDomainManager.ApiLimits = varHciDomainManagerWithoutEmbeddedStruct.ApiLimits
		varHciDomainManager.ApiLimitsString = varHciDomainManagerWithoutEmbeddedStruct.ApiLimitsString
		varHciDomainManager.Name = varHciDomainManagerWithoutEmbeddedStruct.Name
		varHciDomainManager.PcExtId = varHciDomainManagerWithoutEmbeddedStruct.PcExtId
		varHciDomainManager.Size = varHciDomainManagerWithoutEmbeddedStruct.Size
		varHciDomainManager.Clusters = varHciDomainManagerWithoutEmbeddedStruct.Clusters
		varHciDomainManager.Licenses = varHciDomainManagerWithoutEmbeddedStruct.Licenses
		varHciDomainManager.RegisteredDevice = varHciDomainManagerWithoutEmbeddedStruct.RegisteredDevice
		*o = HciDomainManager(varHciDomainManager)
	} else {
		return err
	}

	varHciDomainManager := _HciDomainManager{}

	err = json.Unmarshal(data, &varHciDomainManager)
	if err == nil {
		o.MoBaseMo = varHciDomainManager.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApiLimits")
		delete(additionalProperties, "ApiLimitsString")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PcExtId")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "Clusters")
		delete(additionalProperties, "Licenses")
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

type NullableHciDomainManager struct {
	value *HciDomainManager
	isSet bool
}

func (v NullableHciDomainManager) Get() *HciDomainManager {
	return v.value
}

func (v *NullableHciDomainManager) Set(val *HciDomainManager) {
	v.value = val
	v.isSet = true
}

func (v NullableHciDomainManager) IsSet() bool {
	return v.isSet
}

func (v *NullableHciDomainManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHciDomainManager(val *HciDomainManager) *NullableHciDomainManager {
	return &NullableHciDomainManager{value: val, isSet: true}
}

func (v NullableHciDomainManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHciDomainManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
