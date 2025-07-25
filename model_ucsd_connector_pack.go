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

// checks if the UcsdConnectorPack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UcsdConnectorPack{}

// UcsdConnectorPack Connector pack details from the backup file.
type UcsdConnectorPack struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// State of connector pack (enabled/ disabled).
	ConnectorFeature *string  `json:"ConnectorFeature,omitempty"`
	DependencyNames  []string `json:"DependencyNames,omitempty"`
	// Last successfully downloaded connector pack version successfully for UCS Director.
	DownloadedVersion *string `json:"DownloadedVersion,omitempty"`
	// UCS Director connector pack name.
	Name     *string  `json:"Name,omitempty"`
	Services []string `json:"Services,omitempty"`
	// Connector pack state (enabled/disabled).
	State *string `json:"State,omitempty"`
	// The connector pack version.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UcsdConnectorPack UcsdConnectorPack

// NewUcsdConnectorPack instantiates a new UcsdConnectorPack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUcsdConnectorPack(classId string, objectType string) *UcsdConnectorPack {
	this := UcsdConnectorPack{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewUcsdConnectorPackWithDefaults instantiates a new UcsdConnectorPack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUcsdConnectorPackWithDefaults() *UcsdConnectorPack {
	this := UcsdConnectorPack{}
	var classId string = "ucsd.ConnectorPack"
	this.ClassId = classId
	var objectType string = "ucsd.ConnectorPack"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UcsdConnectorPack) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UcsdConnectorPack) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "ucsd.ConnectorPack" of the ClassId field.
func (o *UcsdConnectorPack) GetDefaultClassId() interface{} {
	return "ucsd.ConnectorPack"
}

// GetObjectType returns the ObjectType field value
func (o *UcsdConnectorPack) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UcsdConnectorPack) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "ucsd.ConnectorPack" of the ObjectType field.
func (o *UcsdConnectorPack) GetDefaultObjectType() interface{} {
	return "ucsd.ConnectorPack"
}

// GetConnectorFeature returns the ConnectorFeature field value if set, zero value otherwise.
func (o *UcsdConnectorPack) GetConnectorFeature() string {
	if o == nil || IsNil(o.ConnectorFeature) {
		var ret string
		return ret
	}
	return *o.ConnectorFeature
}

// GetConnectorFeatureOk returns a tuple with the ConnectorFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetConnectorFeatureOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectorFeature) {
		return nil, false
	}
	return o.ConnectorFeature, true
}

// HasConnectorFeature returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasConnectorFeature() bool {
	if o != nil && !IsNil(o.ConnectorFeature) {
		return true
	}

	return false
}

// SetConnectorFeature gets a reference to the given string and assigns it to the ConnectorFeature field.
func (o *UcsdConnectorPack) SetConnectorFeature(v string) {
	o.ConnectorFeature = &v
}

// GetDependencyNames returns the DependencyNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UcsdConnectorPack) GetDependencyNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DependencyNames
}

// GetDependencyNamesOk returns a tuple with the DependencyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UcsdConnectorPack) GetDependencyNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DependencyNames) {
		return nil, false
	}
	return o.DependencyNames, true
}

// HasDependencyNames returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasDependencyNames() bool {
	if o != nil && !IsNil(o.DependencyNames) {
		return true
	}

	return false
}

// SetDependencyNames gets a reference to the given []string and assigns it to the DependencyNames field.
func (o *UcsdConnectorPack) SetDependencyNames(v []string) {
	o.DependencyNames = v
}

// GetDownloadedVersion returns the DownloadedVersion field value if set, zero value otherwise.
func (o *UcsdConnectorPack) GetDownloadedVersion() string {
	if o == nil || IsNil(o.DownloadedVersion) {
		var ret string
		return ret
	}
	return *o.DownloadedVersion
}

// GetDownloadedVersionOk returns a tuple with the DownloadedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetDownloadedVersionOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadedVersion) {
		return nil, false
	}
	return o.DownloadedVersion, true
}

// HasDownloadedVersion returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasDownloadedVersion() bool {
	if o != nil && !IsNil(o.DownloadedVersion) {
		return true
	}

	return false
}

// SetDownloadedVersion gets a reference to the given string and assigns it to the DownloadedVersion field.
func (o *UcsdConnectorPack) SetDownloadedVersion(v string) {
	o.DownloadedVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UcsdConnectorPack) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UcsdConnectorPack) SetName(v string) {
	o.Name = &v
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UcsdConnectorPack) GetServices() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UcsdConnectorPack) GetServicesOk() ([]string, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *UcsdConnectorPack) SetServices(v []string) {
	o.Services = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UcsdConnectorPack) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UcsdConnectorPack) SetState(v string) {
	o.State = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UcsdConnectorPack) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdConnectorPack) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UcsdConnectorPack) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UcsdConnectorPack) SetVersion(v string) {
	o.Version = &v
}

func (o UcsdConnectorPack) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UcsdConnectorPack) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ConnectorFeature) {
		toSerialize["ConnectorFeature"] = o.ConnectorFeature
	}
	if o.DependencyNames != nil {
		toSerialize["DependencyNames"] = o.DependencyNames
	}
	if !IsNil(o.DownloadedVersion) {
		toSerialize["DownloadedVersion"] = o.DownloadedVersion
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UcsdConnectorPack) UnmarshalJSON(data []byte) (err error) {
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
	type UcsdConnectorPackWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// State of connector pack (enabled/ disabled).
		ConnectorFeature *string  `json:"ConnectorFeature,omitempty"`
		DependencyNames  []string `json:"DependencyNames,omitempty"`
		// Last successfully downloaded connector pack version successfully for UCS Director.
		DownloadedVersion *string `json:"DownloadedVersion,omitempty"`
		// UCS Director connector pack name.
		Name     *string  `json:"Name,omitempty"`
		Services []string `json:"Services,omitempty"`
		// Connector pack state (enabled/disabled).
		State *string `json:"State,omitempty"`
		// The connector pack version.
		Version *string `json:"Version,omitempty"`
	}

	varUcsdConnectorPackWithoutEmbeddedStruct := UcsdConnectorPackWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUcsdConnectorPackWithoutEmbeddedStruct)
	if err == nil {
		varUcsdConnectorPack := _UcsdConnectorPack{}
		varUcsdConnectorPack.ClassId = varUcsdConnectorPackWithoutEmbeddedStruct.ClassId
		varUcsdConnectorPack.ObjectType = varUcsdConnectorPackWithoutEmbeddedStruct.ObjectType
		varUcsdConnectorPack.ConnectorFeature = varUcsdConnectorPackWithoutEmbeddedStruct.ConnectorFeature
		varUcsdConnectorPack.DependencyNames = varUcsdConnectorPackWithoutEmbeddedStruct.DependencyNames
		varUcsdConnectorPack.DownloadedVersion = varUcsdConnectorPackWithoutEmbeddedStruct.DownloadedVersion
		varUcsdConnectorPack.Name = varUcsdConnectorPackWithoutEmbeddedStruct.Name
		varUcsdConnectorPack.Services = varUcsdConnectorPackWithoutEmbeddedStruct.Services
		varUcsdConnectorPack.State = varUcsdConnectorPackWithoutEmbeddedStruct.State
		varUcsdConnectorPack.Version = varUcsdConnectorPackWithoutEmbeddedStruct.Version
		*o = UcsdConnectorPack(varUcsdConnectorPack)
	} else {
		return err
	}

	varUcsdConnectorPack := _UcsdConnectorPack{}

	err = json.Unmarshal(data, &varUcsdConnectorPack)
	if err == nil {
		o.MoBaseComplexType = varUcsdConnectorPack.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectorFeature")
		delete(additionalProperties, "DependencyNames")
		delete(additionalProperties, "DownloadedVersion")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Services")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Version")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableUcsdConnectorPack struct {
	value *UcsdConnectorPack
	isSet bool
}

func (v NullableUcsdConnectorPack) Get() *UcsdConnectorPack {
	return v.value
}

func (v *NullableUcsdConnectorPack) Set(val *UcsdConnectorPack) {
	v.value = val
	v.isSet = true
}

func (v NullableUcsdConnectorPack) IsSet() bool {
	return v.isSet
}

func (v *NullableUcsdConnectorPack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUcsdConnectorPack(val *UcsdConnectorPack) *NullableUcsdConnectorPack {
	return &NullableUcsdConnectorPack{value: val, isSet: true}
}

func (v NullableUcsdConnectorPack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUcsdConnectorPack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
