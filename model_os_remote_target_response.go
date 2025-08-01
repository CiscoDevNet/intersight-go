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

// checks if the OsRemoteTargetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsRemoteTargetResponse{}

// OsRemoteTargetResponse Remote Install target type for the flat array for Orchestration UI.
type OsRemoteTargetResponse struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Interface name of the iSCSI/FC used as initiator.
	InterfaceName        *string `json:"InterfaceName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsRemoteTargetResponse OsRemoteTargetResponse

// NewOsRemoteTargetResponse instantiates a new OsRemoteTargetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsRemoteTargetResponse(classId string, objectType string) *OsRemoteTargetResponse {
	this := OsRemoteTargetResponse{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsRemoteTargetResponseWithDefaults instantiates a new OsRemoteTargetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsRemoteTargetResponseWithDefaults() *OsRemoteTargetResponse {
	this := OsRemoteTargetResponse{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsRemoteTargetResponse) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsRemoteTargetResponse) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsRemoteTargetResponse) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsRemoteTargetResponse) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsRemoteTargetResponse) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsRemoteTargetResponse) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *OsRemoteTargetResponse) GetInterfaceName() string {
	if o == nil || IsNil(o.InterfaceName) {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsRemoteTargetResponse) GetInterfaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceName) {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *OsRemoteTargetResponse) HasInterfaceName() bool {
	if o != nil && !IsNil(o.InterfaceName) {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *OsRemoteTargetResponse) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

func (o OsRemoteTargetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsRemoteTargetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.InterfaceName) {
		toSerialize["InterfaceName"] = o.InterfaceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OsRemoteTargetResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type OsRemoteTargetResponseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Interface name of the iSCSI/FC used as initiator.
		InterfaceName *string `json:"InterfaceName,omitempty"`
	}

	varOsRemoteTargetResponseWithoutEmbeddedStruct := OsRemoteTargetResponseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOsRemoteTargetResponseWithoutEmbeddedStruct)
	if err == nil {
		varOsRemoteTargetResponse := _OsRemoteTargetResponse{}
		varOsRemoteTargetResponse.ClassId = varOsRemoteTargetResponseWithoutEmbeddedStruct.ClassId
		varOsRemoteTargetResponse.ObjectType = varOsRemoteTargetResponseWithoutEmbeddedStruct.ObjectType
		varOsRemoteTargetResponse.InterfaceName = varOsRemoteTargetResponseWithoutEmbeddedStruct.InterfaceName
		*o = OsRemoteTargetResponse(varOsRemoteTargetResponse)
	} else {
		return err
	}

	varOsRemoteTargetResponse := _OsRemoteTargetResponse{}

	err = json.Unmarshal(data, &varOsRemoteTargetResponse)
	if err == nil {
		o.MoBaseComplexType = varOsRemoteTargetResponse.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InterfaceName")

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

type NullableOsRemoteTargetResponse struct {
	value *OsRemoteTargetResponse
	isSet bool
}

func (v NullableOsRemoteTargetResponse) Get() *OsRemoteTargetResponse {
	return v.value
}

func (v *NullableOsRemoteTargetResponse) Set(val *OsRemoteTargetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOsRemoteTargetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOsRemoteTargetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsRemoteTargetResponse(val *OsRemoteTargetResponse) *NullableOsRemoteTargetResponse {
	return &NullableOsRemoteTargetResponse{value: val, isSet: true}
}

func (v NullableOsRemoteTargetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsRemoteTargetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
