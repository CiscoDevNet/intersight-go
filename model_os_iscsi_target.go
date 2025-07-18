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

// checks if the OsIscsiTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsIscsiTarget{}

// OsIscsiTarget ISCSI LUN to install Operating System.
type OsIscsiTarget struct {
	OsSanTarget
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Logical Unit Number (LUN) of the install target.
	LunId *int64 `json:"LunId,omitempty"`
	// IQN (iSCSI qualified name) of Storage iSCSI target.Can be up to 255 characters long and has the following format, iqn.yyyy-mm.naming-authority:unique_name.
	TargetIqn *string `json:"TargetIqn,omitempty" validate:"regexp=^(iqn\\\\.(\\\\d{4}-(0[1-9]|1[0-2]))\\\\.([a-zA-Z0-9]+)([a-zA-Z0-9-.]*)(:)+([^,\\\\s']+))$"`
	// MAC address of the VNIC used as iSCSI initiator interface.
	VnicMac              *string `json:"VnicMac,omitempty" validate:"regexp=^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$"`
	AdditionalProperties map[string]interface{}
}

type _OsIscsiTarget OsIscsiTarget

// NewOsIscsiTarget instantiates a new OsIscsiTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsIscsiTarget(classId string, objectType string) *OsIscsiTarget {
	this := OsIscsiTarget{}
	this.ClassId = classId
	this.ObjectType = objectType
	var lunId int64 = 0
	this.LunId = &lunId
	return &this
}

// NewOsIscsiTargetWithDefaults instantiates a new OsIscsiTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsIscsiTargetWithDefaults() *OsIscsiTarget {
	this := OsIscsiTarget{}
	var classId string = "os.IscsiTarget"
	this.ClassId = classId
	var objectType string = "os.IscsiTarget"
	this.ObjectType = objectType
	var lunId int64 = 0
	this.LunId = &lunId
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsIscsiTarget) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsIscsiTarget) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsIscsiTarget) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "os.IscsiTarget" of the ClassId field.
func (o *OsIscsiTarget) GetDefaultClassId() interface{} {
	return "os.IscsiTarget"
}

// GetObjectType returns the ObjectType field value
func (o *OsIscsiTarget) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsIscsiTarget) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsIscsiTarget) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "os.IscsiTarget" of the ObjectType field.
func (o *OsIscsiTarget) GetDefaultObjectType() interface{} {
	return "os.IscsiTarget"
}

// GetLunId returns the LunId field value if set, zero value otherwise.
func (o *OsIscsiTarget) GetLunId() int64 {
	if o == nil || IsNil(o.LunId) {
		var ret int64
		return ret
	}
	return *o.LunId
}

// GetLunIdOk returns a tuple with the LunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsIscsiTarget) GetLunIdOk() (*int64, bool) {
	if o == nil || IsNil(o.LunId) {
		return nil, false
	}
	return o.LunId, true
}

// HasLunId returns a boolean if a field has been set.
func (o *OsIscsiTarget) HasLunId() bool {
	if o != nil && !IsNil(o.LunId) {
		return true
	}

	return false
}

// SetLunId gets a reference to the given int64 and assigns it to the LunId field.
func (o *OsIscsiTarget) SetLunId(v int64) {
	o.LunId = &v
}

// GetTargetIqn returns the TargetIqn field value if set, zero value otherwise.
func (o *OsIscsiTarget) GetTargetIqn() string {
	if o == nil || IsNil(o.TargetIqn) {
		var ret string
		return ret
	}
	return *o.TargetIqn
}

// GetTargetIqnOk returns a tuple with the TargetIqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsIscsiTarget) GetTargetIqnOk() (*string, bool) {
	if o == nil || IsNil(o.TargetIqn) {
		return nil, false
	}
	return o.TargetIqn, true
}

// HasTargetIqn returns a boolean if a field has been set.
func (o *OsIscsiTarget) HasTargetIqn() bool {
	if o != nil && !IsNil(o.TargetIqn) {
		return true
	}

	return false
}

// SetTargetIqn gets a reference to the given string and assigns it to the TargetIqn field.
func (o *OsIscsiTarget) SetTargetIqn(v string) {
	o.TargetIqn = &v
}

// GetVnicMac returns the VnicMac field value if set, zero value otherwise.
func (o *OsIscsiTarget) GetVnicMac() string {
	if o == nil || IsNil(o.VnicMac) {
		var ret string
		return ret
	}
	return *o.VnicMac
}

// GetVnicMacOk returns a tuple with the VnicMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsIscsiTarget) GetVnicMacOk() (*string, bool) {
	if o == nil || IsNil(o.VnicMac) {
		return nil, false
	}
	return o.VnicMac, true
}

// HasVnicMac returns a boolean if a field has been set.
func (o *OsIscsiTarget) HasVnicMac() bool {
	if o != nil && !IsNil(o.VnicMac) {
		return true
	}

	return false
}

// SetVnicMac gets a reference to the given string and assigns it to the VnicMac field.
func (o *OsIscsiTarget) SetVnicMac(v string) {
	o.VnicMac = &v
}

func (o OsIscsiTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsIscsiTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedOsSanTarget, errOsSanTarget := json.Marshal(o.OsSanTarget)
	if errOsSanTarget != nil {
		return map[string]interface{}{}, errOsSanTarget
	}
	errOsSanTarget = json.Unmarshal([]byte(serializedOsSanTarget), &toSerialize)
	if errOsSanTarget != nil {
		return map[string]interface{}{}, errOsSanTarget
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.LunId) {
		toSerialize["LunId"] = o.LunId
	}
	if !IsNil(o.TargetIqn) {
		toSerialize["TargetIqn"] = o.TargetIqn
	}
	if !IsNil(o.VnicMac) {
		toSerialize["VnicMac"] = o.VnicMac
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OsIscsiTarget) UnmarshalJSON(data []byte) (err error) {
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
	type OsIscsiTargetWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The Logical Unit Number (LUN) of the install target.
		LunId *int64 `json:"LunId,omitempty"`
		// IQN (iSCSI qualified name) of Storage iSCSI target.Can be up to 255 characters long and has the following format, iqn.yyyy-mm.naming-authority:unique_name.
		TargetIqn *string `json:"TargetIqn,omitempty" validate:"regexp=^(iqn\\\\.(\\\\d{4}-(0[1-9]|1[0-2]))\\\\.([a-zA-Z0-9]+)([a-zA-Z0-9-.]*)(:)+([^,\\\\s']+))$"`
		// MAC address of the VNIC used as iSCSI initiator interface.
		VnicMac *string `json:"VnicMac,omitempty" validate:"regexp=^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$"`
	}

	varOsIscsiTargetWithoutEmbeddedStruct := OsIscsiTargetWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOsIscsiTargetWithoutEmbeddedStruct)
	if err == nil {
		varOsIscsiTarget := _OsIscsiTarget{}
		varOsIscsiTarget.ClassId = varOsIscsiTargetWithoutEmbeddedStruct.ClassId
		varOsIscsiTarget.ObjectType = varOsIscsiTargetWithoutEmbeddedStruct.ObjectType
		varOsIscsiTarget.LunId = varOsIscsiTargetWithoutEmbeddedStruct.LunId
		varOsIscsiTarget.TargetIqn = varOsIscsiTargetWithoutEmbeddedStruct.TargetIqn
		varOsIscsiTarget.VnicMac = varOsIscsiTargetWithoutEmbeddedStruct.VnicMac
		*o = OsIscsiTarget(varOsIscsiTarget)
	} else {
		return err
	}

	varOsIscsiTarget := _OsIscsiTarget{}

	err = json.Unmarshal(data, &varOsIscsiTarget)
	if err == nil {
		o.OsSanTarget = varOsIscsiTarget.OsSanTarget
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LunId")
		delete(additionalProperties, "TargetIqn")
		delete(additionalProperties, "VnicMac")

		// remove fields from embedded structs
		reflectOsSanTarget := reflect.ValueOf(o.OsSanTarget)
		for i := 0; i < reflectOsSanTarget.Type().NumField(); i++ {
			t := reflectOsSanTarget.Type().Field(i)

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

type NullableOsIscsiTarget struct {
	value *OsIscsiTarget
	isSet bool
}

func (v NullableOsIscsiTarget) Get() *OsIscsiTarget {
	return v.value
}

func (v *NullableOsIscsiTarget) Set(val *OsIscsiTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableOsIscsiTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableOsIscsiTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsIscsiTarget(val *OsIscsiTarget) *NullableOsIscsiTarget {
	return &NullableOsIscsiTarget{value: val, isSet: true}
}

func (v NullableOsIscsiTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsIscsiTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
