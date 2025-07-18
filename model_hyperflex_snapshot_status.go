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

// checks if the HyperflexSnapshotStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexSnapshotStatus{}

// HyperflexSnapshotStatus The current state of this virtual machine's snapshot.
type HyperflexSnapshotStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of this Snapshot Point.
	Description *string                     `json:"Description,omitempty"`
	Error       NullableHyperflexErrorStack `json:"Error,omitempty"`
	// Completion percentage for this snapshot.
	PctComplete *int64 `json:"PctComplete,omitempty"`
	// Current snapshot state for this snapshot. * `SUCCESS` - This snapshot status code is success. * `FAILED` - This snapshot status code is failed. * `IN_PROGRESS` - This snapshot status code is in progress. * `DELETING` - This snapshot status code is deleting. * `DELETED` - This snapshot status code is deleted. * `NONE` - This snapshot status code is none. * `INIT` - This snapshot status code is initializing.
	Status *string `json:"Status,omitempty"`
	// Timestamp at which the Snapshot is taken.
	Timestamp *int64 `json:"Timestamp,omitempty"`
	// Space Used by this Snapshot Point.
	UsedSpace            *int64 `json:"UsedSpace,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexSnapshotStatus HyperflexSnapshotStatus

// NewHyperflexSnapshotStatus instantiates a new HyperflexSnapshotStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSnapshotStatus(classId string, objectType string) *HyperflexSnapshotStatus {
	this := HyperflexSnapshotStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSnapshotStatusWithDefaults instantiates a new HyperflexSnapshotStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSnapshotStatusWithDefaults() *HyperflexSnapshotStatus {
	this := HyperflexSnapshotStatus{}
	var classId string = "hyperflex.SnapshotStatus"
	this.ClassId = classId
	var objectType string = "hyperflex.SnapshotStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSnapshotStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSnapshotStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.SnapshotStatus" of the ClassId field.
func (o *HyperflexSnapshotStatus) GetDefaultClassId() interface{} {
	return "hyperflex.SnapshotStatus"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSnapshotStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSnapshotStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.SnapshotStatus" of the ObjectType field.
func (o *HyperflexSnapshotStatus) GetDefaultObjectType() interface{} {
	return "hyperflex.SnapshotStatus"
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HyperflexSnapshotStatus) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HyperflexSnapshotStatus) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HyperflexSnapshotStatus) SetDescription(v string) {
	o.Description = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotStatus) GetError() HyperflexErrorStack {
	if o == nil || IsNil(o.Error.Get()) {
		var ret HyperflexErrorStack
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotStatus) GetErrorOk() (*HyperflexErrorStack, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *HyperflexSnapshotStatus) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableHyperflexErrorStack and assigns it to the Error field.
func (o *HyperflexSnapshotStatus) SetError(v HyperflexErrorStack) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *HyperflexSnapshotStatus) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *HyperflexSnapshotStatus) UnsetError() {
	o.Error.Unset()
}

// GetPctComplete returns the PctComplete field value if set, zero value otherwise.
func (o *HyperflexSnapshotStatus) GetPctComplete() int64 {
	if o == nil || IsNil(o.PctComplete) {
		var ret int64
		return ret
	}
	return *o.PctComplete
}

// GetPctCompleteOk returns a tuple with the PctComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetPctCompleteOk() (*int64, bool) {
	if o == nil || IsNil(o.PctComplete) {
		return nil, false
	}
	return o.PctComplete, true
}

// HasPctComplete returns a boolean if a field has been set.
func (o *HyperflexSnapshotStatus) HasPctComplete() bool {
	if o != nil && !IsNil(o.PctComplete) {
		return true
	}

	return false
}

// SetPctComplete gets a reference to the given int64 and assigns it to the PctComplete field.
func (o *HyperflexSnapshotStatus) SetPctComplete(v int64) {
	o.PctComplete = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexSnapshotStatus) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexSnapshotStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexSnapshotStatus) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *HyperflexSnapshotStatus) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *HyperflexSnapshotStatus) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *HyperflexSnapshotStatus) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetUsedSpace returns the UsedSpace field value if set, zero value otherwise.
func (o *HyperflexSnapshotStatus) GetUsedSpace() int64 {
	if o == nil || IsNil(o.UsedSpace) {
		var ret int64
		return ret
	}
	return *o.UsedSpace
}

// GetUsedSpaceOk returns a tuple with the UsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotStatus) GetUsedSpaceOk() (*int64, bool) {
	if o == nil || IsNil(o.UsedSpace) {
		return nil, false
	}
	return o.UsedSpace, true
}

// HasUsedSpace returns a boolean if a field has been set.
func (o *HyperflexSnapshotStatus) HasUsedSpace() bool {
	if o != nil && !IsNil(o.UsedSpace) {
		return true
	}

	return false
}

// SetUsedSpace gets a reference to the given int64 and assigns it to the UsedSpace field.
func (o *HyperflexSnapshotStatus) SetUsedSpace(v int64) {
	o.UsedSpace = &v
}

func (o HyperflexSnapshotStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexSnapshotStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if o.Error.IsSet() {
		toSerialize["Error"] = o.Error.Get()
	}
	if !IsNil(o.PctComplete) {
		toSerialize["PctComplete"] = o.PctComplete
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.Timestamp) {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if !IsNil(o.UsedSpace) {
		toSerialize["UsedSpace"] = o.UsedSpace
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexSnapshotStatus) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexSnapshotStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of this Snapshot Point.
		Description *string                     `json:"Description,omitempty"`
		Error       NullableHyperflexErrorStack `json:"Error,omitempty"`
		// Completion percentage for this snapshot.
		PctComplete *int64 `json:"PctComplete,omitempty"`
		// Current snapshot state for this snapshot. * `SUCCESS` - This snapshot status code is success. * `FAILED` - This snapshot status code is failed. * `IN_PROGRESS` - This snapshot status code is in progress. * `DELETING` - This snapshot status code is deleting. * `DELETED` - This snapshot status code is deleted. * `NONE` - This snapshot status code is none. * `INIT` - This snapshot status code is initializing.
		Status *string `json:"Status,omitempty"`
		// Timestamp at which the Snapshot is taken.
		Timestamp *int64 `json:"Timestamp,omitempty"`
		// Space Used by this Snapshot Point.
		UsedSpace *int64 `json:"UsedSpace,omitempty"`
	}

	varHyperflexSnapshotStatusWithoutEmbeddedStruct := HyperflexSnapshotStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexSnapshotStatusWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexSnapshotStatus := _HyperflexSnapshotStatus{}
		varHyperflexSnapshotStatus.ClassId = varHyperflexSnapshotStatusWithoutEmbeddedStruct.ClassId
		varHyperflexSnapshotStatus.ObjectType = varHyperflexSnapshotStatusWithoutEmbeddedStruct.ObjectType
		varHyperflexSnapshotStatus.Description = varHyperflexSnapshotStatusWithoutEmbeddedStruct.Description
		varHyperflexSnapshotStatus.Error = varHyperflexSnapshotStatusWithoutEmbeddedStruct.Error
		varHyperflexSnapshotStatus.PctComplete = varHyperflexSnapshotStatusWithoutEmbeddedStruct.PctComplete
		varHyperflexSnapshotStatus.Status = varHyperflexSnapshotStatusWithoutEmbeddedStruct.Status
		varHyperflexSnapshotStatus.Timestamp = varHyperflexSnapshotStatusWithoutEmbeddedStruct.Timestamp
		varHyperflexSnapshotStatus.UsedSpace = varHyperflexSnapshotStatusWithoutEmbeddedStruct.UsedSpace
		*o = HyperflexSnapshotStatus(varHyperflexSnapshotStatus)
	} else {
		return err
	}

	varHyperflexSnapshotStatus := _HyperflexSnapshotStatus{}

	err = json.Unmarshal(data, &varHyperflexSnapshotStatus)
	if err == nil {
		o.MoBaseComplexType = varHyperflexSnapshotStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Error")
		delete(additionalProperties, "PctComplete")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "UsedSpace")

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

type NullableHyperflexSnapshotStatus struct {
	value *HyperflexSnapshotStatus
	isSet bool
}

func (v NullableHyperflexSnapshotStatus) Get() *HyperflexSnapshotStatus {
	return v.value
}

func (v *NullableHyperflexSnapshotStatus) Set(val *HyperflexSnapshotStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSnapshotStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSnapshotStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSnapshotStatus(val *HyperflexSnapshotStatus) *NullableHyperflexSnapshotStatus {
	return &NullableHyperflexSnapshotStatus{value: val, isSet: true}
}

func (v NullableHyperflexSnapshotStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSnapshotStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
