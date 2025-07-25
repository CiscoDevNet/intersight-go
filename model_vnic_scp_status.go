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

// checks if the VnicScpStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicScpStatus{}

// VnicScpStatus An internal MO to check if a SCP can be deployed or not on a specific Server Profile.
type VnicScpStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure.
	Reason *string `json:"Reason,omitempty"`
	// Indicates if the LCP is ready for Deploy or not. * `ok` - No issues with the LCP/SCP/VIF. * `error` - The LCP/SCP/VIF cannot be deployed due to error. * `validating` - Validation in progress for the LCP.
	Status               *string                                         `json:"Status,omitempty"`
	VhbaInfo             []VnicVifStatus                                 `json:"VhbaInfo,omitempty"`
	Profile              NullablePolicyAbstractConfigProfileRelationship `json:"Profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicScpStatus VnicScpStatus

// NewVnicScpStatus instantiates a new VnicScpStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicScpStatus(classId string, objectType string) *VnicScpStatus {
	this := VnicScpStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// NewVnicScpStatusWithDefaults instantiates a new VnicScpStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicScpStatusWithDefaults() *VnicScpStatus {
	this := VnicScpStatus{}
	var classId string = "vnic.ScpStatus"
	this.ClassId = classId
	var objectType string = "vnic.ScpStatus"
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicScpStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicScpStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicScpStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.ScpStatus" of the ClassId field.
func (o *VnicScpStatus) GetDefaultClassId() interface{} {
	return "vnic.ScpStatus"
}

// GetObjectType returns the ObjectType field value
func (o *VnicScpStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicScpStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicScpStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.ScpStatus" of the ObjectType field.
func (o *VnicScpStatus) GetDefaultObjectType() interface{} {
	return "vnic.ScpStatus"
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *VnicScpStatus) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicScpStatus) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *VnicScpStatus) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *VnicScpStatus) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VnicScpStatus) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicScpStatus) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VnicScpStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VnicScpStatus) SetStatus(v string) {
	o.Status = &v
}

// GetVhbaInfo returns the VhbaInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicScpStatus) GetVhbaInfo() []VnicVifStatus {
	if o == nil {
		var ret []VnicVifStatus
		return ret
	}
	return o.VhbaInfo
}

// GetVhbaInfoOk returns a tuple with the VhbaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicScpStatus) GetVhbaInfoOk() ([]VnicVifStatus, bool) {
	if o == nil || IsNil(o.VhbaInfo) {
		return nil, false
	}
	return o.VhbaInfo, true
}

// HasVhbaInfo returns a boolean if a field has been set.
func (o *VnicScpStatus) HasVhbaInfo() bool {
	if o != nil && !IsNil(o.VhbaInfo) {
		return true
	}

	return false
}

// SetVhbaInfo gets a reference to the given []VnicVifStatus and assigns it to the VhbaInfo field.
func (o *VnicScpStatus) SetVhbaInfo(v []VnicVifStatus) {
	o.VhbaInfo = v
}

// GetProfile returns the Profile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicScpStatus) GetProfile() PolicyAbstractConfigProfileRelationship {
	if o == nil || IsNil(o.Profile.Get()) {
		var ret PolicyAbstractConfigProfileRelationship
		return ret
	}
	return *o.Profile.Get()
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicScpStatus) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profile.Get(), o.Profile.IsSet()
}

// HasProfile returns a boolean if a field has been set.
func (o *VnicScpStatus) HasProfile() bool {
	if o != nil && o.Profile.IsSet() {
		return true
	}

	return false
}

// SetProfile gets a reference to the given NullablePolicyAbstractConfigProfileRelationship and assigns it to the Profile field.
func (o *VnicScpStatus) SetProfile(v PolicyAbstractConfigProfileRelationship) {
	o.Profile.Set(&v)
}

// SetProfileNil sets the value for Profile to be an explicit nil
func (o *VnicScpStatus) SetProfileNil() {
	o.Profile.Set(nil)
}

// UnsetProfile ensures that no value is present for Profile, not even an explicit nil
func (o *VnicScpStatus) UnsetProfile() {
	o.Profile.Unset()
}

func (o VnicScpStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicScpStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Reason) {
		toSerialize["Reason"] = o.Reason
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if o.VhbaInfo != nil {
		toSerialize["VhbaInfo"] = o.VhbaInfo
	}
	if o.Profile.IsSet() {
		toSerialize["Profile"] = o.Profile.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicScpStatus) UnmarshalJSON(data []byte) (err error) {
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
	type VnicScpStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure.
		Reason *string `json:"Reason,omitempty"`
		// Indicates if the LCP is ready for Deploy or not. * `ok` - No issues with the LCP/SCP/VIF. * `error` - The LCP/SCP/VIF cannot be deployed due to error. * `validating` - Validation in progress for the LCP.
		Status   *string                                         `json:"Status,omitempty"`
		VhbaInfo []VnicVifStatus                                 `json:"VhbaInfo,omitempty"`
		Profile  NullablePolicyAbstractConfigProfileRelationship `json:"Profile,omitempty"`
	}

	varVnicScpStatusWithoutEmbeddedStruct := VnicScpStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicScpStatusWithoutEmbeddedStruct)
	if err == nil {
		varVnicScpStatus := _VnicScpStatus{}
		varVnicScpStatus.ClassId = varVnicScpStatusWithoutEmbeddedStruct.ClassId
		varVnicScpStatus.ObjectType = varVnicScpStatusWithoutEmbeddedStruct.ObjectType
		varVnicScpStatus.Reason = varVnicScpStatusWithoutEmbeddedStruct.Reason
		varVnicScpStatus.Status = varVnicScpStatusWithoutEmbeddedStruct.Status
		varVnicScpStatus.VhbaInfo = varVnicScpStatusWithoutEmbeddedStruct.VhbaInfo
		varVnicScpStatus.Profile = varVnicScpStatusWithoutEmbeddedStruct.Profile
		*o = VnicScpStatus(varVnicScpStatus)
	} else {
		return err
	}

	varVnicScpStatus := _VnicScpStatus{}

	err = json.Unmarshal(data, &varVnicScpStatus)
	if err == nil {
		o.MoBaseMo = varVnicScpStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "VhbaInfo")
		delete(additionalProperties, "Profile")

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

type NullableVnicScpStatus struct {
	value *VnicScpStatus
	isSet bool
}

func (v NullableVnicScpStatus) Get() *VnicScpStatus {
	return v.value
}

func (v *NullableVnicScpStatus) Set(val *VnicScpStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicScpStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicScpStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicScpStatus(val *VnicScpStatus) *NullableVnicScpStatus {
	return &NullableVnicScpStatus{value: val, isSet: true}
}

func (v NullableVnicScpStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicScpStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
