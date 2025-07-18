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

// checks if the ConnectorpackConnectorPackUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorpackConnectorPackUpdate{}

// ConnectorpackConnectorPackUpdate Connector Pack updates available on UCS Director.
type ConnectorpackConnectorPackUpdate struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Version of connector pack currently running in UCS Director.
	CurrentVersion *string `json:"CurrentVersion,omitempty"`
	// Name of the connector pack.
	Name *string `json:"Name,omitempty"`
	// Version of connector pack to be installed in the next upgrade cycle.
	NewVersion           *string `json:"NewVersion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorpackConnectorPackUpdate ConnectorpackConnectorPackUpdate

// NewConnectorpackConnectorPackUpdate instantiates a new ConnectorpackConnectorPackUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorpackConnectorPackUpdate(classId string, objectType string) *ConnectorpackConnectorPackUpdate {
	this := ConnectorpackConnectorPackUpdate{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorpackConnectorPackUpdateWithDefaults instantiates a new ConnectorpackConnectorPackUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorpackConnectorPackUpdateWithDefaults() *ConnectorpackConnectorPackUpdate {
	this := ConnectorpackConnectorPackUpdate{}
	var classId string = "connectorpack.ConnectorPackUpdate"
	this.ClassId = classId
	var objectType string = "connectorpack.ConnectorPackUpdate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorpackConnectorPackUpdate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorpackConnectorPackUpdate) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "connectorpack.ConnectorPackUpdate" of the ClassId field.
func (o *ConnectorpackConnectorPackUpdate) GetDefaultClassId() interface{} {
	return "connectorpack.ConnectorPackUpdate"
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorpackConnectorPackUpdate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorpackConnectorPackUpdate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "connectorpack.ConnectorPackUpdate" of the ObjectType field.
func (o *ConnectorpackConnectorPackUpdate) GetDefaultObjectType() interface{} {
	return "connectorpack.ConnectorPackUpdate"
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpdate) GetCurrentVersion() string {
	if o == nil || IsNil(o.CurrentVersion) {
		var ret string
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdate) GetCurrentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentVersion) {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpdate) HasCurrentVersion() bool {
	if o != nil && !IsNil(o.CurrentVersion) {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given string and assigns it to the CurrentVersion field.
func (o *ConnectorpackConnectorPackUpdate) SetCurrentVersion(v string) {
	o.CurrentVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectorpackConnectorPackUpdate) SetName(v string) {
	o.Name = &v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpdate) GetNewVersion() string {
	if o == nil || IsNil(o.NewVersion) {
		var ret string
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdate) GetNewVersionOk() (*string, bool) {
	if o == nil || IsNil(o.NewVersion) {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpdate) HasNewVersion() bool {
	if o != nil && !IsNil(o.NewVersion) {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given string and assigns it to the NewVersion field.
func (o *ConnectorpackConnectorPackUpdate) SetNewVersion(v string) {
	o.NewVersion = &v
}

func (o ConnectorpackConnectorPackUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorpackConnectorPackUpdate) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CurrentVersion) {
		toSerialize["CurrentVersion"] = o.CurrentVersion
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.NewVersion) {
		toSerialize["NewVersion"] = o.NewVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectorpackConnectorPackUpdate) UnmarshalJSON(data []byte) (err error) {
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
	type ConnectorpackConnectorPackUpdateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Version of connector pack currently running in UCS Director.
		CurrentVersion *string `json:"CurrentVersion,omitempty"`
		// Name of the connector pack.
		Name *string `json:"Name,omitempty"`
		// Version of connector pack to be installed in the next upgrade cycle.
		NewVersion *string `json:"NewVersion,omitempty"`
	}

	varConnectorpackConnectorPackUpdateWithoutEmbeddedStruct := ConnectorpackConnectorPackUpdateWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varConnectorpackConnectorPackUpdateWithoutEmbeddedStruct)
	if err == nil {
		varConnectorpackConnectorPackUpdate := _ConnectorpackConnectorPackUpdate{}
		varConnectorpackConnectorPackUpdate.ClassId = varConnectorpackConnectorPackUpdateWithoutEmbeddedStruct.ClassId
		varConnectorpackConnectorPackUpdate.ObjectType = varConnectorpackConnectorPackUpdateWithoutEmbeddedStruct.ObjectType
		varConnectorpackConnectorPackUpdate.CurrentVersion = varConnectorpackConnectorPackUpdateWithoutEmbeddedStruct.CurrentVersion
		varConnectorpackConnectorPackUpdate.Name = varConnectorpackConnectorPackUpdateWithoutEmbeddedStruct.Name
		varConnectorpackConnectorPackUpdate.NewVersion = varConnectorpackConnectorPackUpdateWithoutEmbeddedStruct.NewVersion
		*o = ConnectorpackConnectorPackUpdate(varConnectorpackConnectorPackUpdate)
	} else {
		return err
	}

	varConnectorpackConnectorPackUpdate := _ConnectorpackConnectorPackUpdate{}

	err = json.Unmarshal(data, &varConnectorpackConnectorPackUpdate)
	if err == nil {
		o.MoBaseComplexType = varConnectorpackConnectorPackUpdate.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CurrentVersion")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NewVersion")

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

type NullableConnectorpackConnectorPackUpdate struct {
	value *ConnectorpackConnectorPackUpdate
	isSet bool
}

func (v NullableConnectorpackConnectorPackUpdate) Get() *ConnectorpackConnectorPackUpdate {
	return v.value
}

func (v *NullableConnectorpackConnectorPackUpdate) Set(val *ConnectorpackConnectorPackUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorpackConnectorPackUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorpackConnectorPackUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorpackConnectorPackUpdate(val *ConnectorpackConnectorPackUpdate) *NullableConnectorpackConnectorPackUpdate {
	return &NullableConnectorpackConnectorPackUpdate{value: val, isSet: true}
}

func (v NullableConnectorpackConnectorPackUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorpackConnectorPackUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
