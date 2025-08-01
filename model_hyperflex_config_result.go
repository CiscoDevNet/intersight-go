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

// checks if the HyperflexConfigResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexConfigResult{}

// HyperflexConfigResult Profile configuration (deploy, validation) results with the overall state and detailed result messages.
type HyperflexConfigResult struct {
	PolicyAbstractConfigResult
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The progress percentage of the running configuration or workflow.
	ConfigProgress *string `json:"ConfigProgress,omitempty"`
	// The duration of the running configuration or workflow.
	Duration *string `json:"Duration,omitempty"`
	// The start time of the configuration or workflow.
	StartTime      *string                                     `json:"StartTime,omitempty"`
	ClusterProfile NullableHyperflexClusterProfileRelationship `json:"ClusterProfile,omitempty"`
	// An array of relationships to hyperflexConfigResultEntry resources.
	ResultEntries        []HyperflexConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexConfigResult HyperflexConfigResult

// NewHyperflexConfigResult instantiates a new HyperflexConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexConfigResult(classId string, objectType string) *HyperflexConfigResult {
	this := HyperflexConfigResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexConfigResultWithDefaults instantiates a new HyperflexConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexConfigResultWithDefaults() *HyperflexConfigResult {
	this := HyperflexConfigResult{}
	var classId string = "hyperflex.ConfigResult"
	this.ClassId = classId
	var objectType string = "hyperflex.ConfigResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexConfigResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexConfigResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexConfigResult) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ConfigResult" of the ClassId field.
func (o *HyperflexConfigResult) GetDefaultClassId() interface{} {
	return "hyperflex.ConfigResult"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexConfigResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexConfigResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexConfigResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ConfigResult" of the ObjectType field.
func (o *HyperflexConfigResult) GetDefaultObjectType() interface{} {
	return "hyperflex.ConfigResult"
}

// GetConfigProgress returns the ConfigProgress field value if set, zero value otherwise.
func (o *HyperflexConfigResult) GetConfigProgress() string {
	if o == nil || IsNil(o.ConfigProgress) {
		var ret string
		return ret
	}
	return *o.ConfigProgress
}

// GetConfigProgressOk returns a tuple with the ConfigProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexConfigResult) GetConfigProgressOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigProgress) {
		return nil, false
	}
	return o.ConfigProgress, true
}

// HasConfigProgress returns a boolean if a field has been set.
func (o *HyperflexConfigResult) HasConfigProgress() bool {
	if o != nil && !IsNil(o.ConfigProgress) {
		return true
	}

	return false
}

// SetConfigProgress gets a reference to the given string and assigns it to the ConfigProgress field.
func (o *HyperflexConfigResult) SetConfigProgress(v string) {
	o.ConfigProgress = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *HyperflexConfigResult) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexConfigResult) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *HyperflexConfigResult) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *HyperflexConfigResult) SetDuration(v string) {
	o.Duration = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *HyperflexConfigResult) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexConfigResult) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *HyperflexConfigResult) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *HyperflexConfigResult) SetStartTime(v string) {
	o.StartTime = &v
}

// GetClusterProfile returns the ClusterProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexConfigResult) GetClusterProfile() HyperflexClusterProfileRelationship {
	if o == nil || IsNil(o.ClusterProfile.Get()) {
		var ret HyperflexClusterProfileRelationship
		return ret
	}
	return *o.ClusterProfile.Get()
}

// GetClusterProfileOk returns a tuple with the ClusterProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexConfigResult) GetClusterProfileOk() (*HyperflexClusterProfileRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterProfile.Get(), o.ClusterProfile.IsSet()
}

// HasClusterProfile returns a boolean if a field has been set.
func (o *HyperflexConfigResult) HasClusterProfile() bool {
	if o != nil && o.ClusterProfile.IsSet() {
		return true
	}

	return false
}

// SetClusterProfile gets a reference to the given NullableHyperflexClusterProfileRelationship and assigns it to the ClusterProfile field.
func (o *HyperflexConfigResult) SetClusterProfile(v HyperflexClusterProfileRelationship) {
	o.ClusterProfile.Set(&v)
}

// SetClusterProfileNil sets the value for ClusterProfile to be an explicit nil
func (o *HyperflexConfigResult) SetClusterProfileNil() {
	o.ClusterProfile.Set(nil)
}

// UnsetClusterProfile ensures that no value is present for ClusterProfile, not even an explicit nil
func (o *HyperflexConfigResult) UnsetClusterProfile() {
	o.ClusterProfile.Unset()
}

// GetResultEntries returns the ResultEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexConfigResult) GetResultEntries() []HyperflexConfigResultEntryRelationship {
	if o == nil {
		var ret []HyperflexConfigResultEntryRelationship
		return ret
	}
	return o.ResultEntries
}

// GetResultEntriesOk returns a tuple with the ResultEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexConfigResult) GetResultEntriesOk() ([]HyperflexConfigResultEntryRelationship, bool) {
	if o == nil || IsNil(o.ResultEntries) {
		return nil, false
	}
	return o.ResultEntries, true
}

// HasResultEntries returns a boolean if a field has been set.
func (o *HyperflexConfigResult) HasResultEntries() bool {
	if o != nil && !IsNil(o.ResultEntries) {
		return true
	}

	return false
}

// SetResultEntries gets a reference to the given []HyperflexConfigResultEntryRelationship and assigns it to the ResultEntries field.
func (o *HyperflexConfigResult) SetResultEntries(v []HyperflexConfigResultEntryRelationship) {
	o.ResultEntries = v
}

func (o HyperflexConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexConfigResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigResult, errPolicyAbstractConfigResult := json.Marshal(o.PolicyAbstractConfigResult)
	if errPolicyAbstractConfigResult != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigResult
	}
	errPolicyAbstractConfigResult = json.Unmarshal([]byte(serializedPolicyAbstractConfigResult), &toSerialize)
	if errPolicyAbstractConfigResult != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigResult
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ConfigProgress) {
		toSerialize["ConfigProgress"] = o.ConfigProgress
	}
	if !IsNil(o.Duration) {
		toSerialize["Duration"] = o.Duration
	}
	if !IsNil(o.StartTime) {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.ClusterProfile.IsSet() {
		toSerialize["ClusterProfile"] = o.ClusterProfile.Get()
	}
	if o.ResultEntries != nil {
		toSerialize["ResultEntries"] = o.ResultEntries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexConfigResult) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexConfigResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The progress percentage of the running configuration or workflow.
		ConfigProgress *string `json:"ConfigProgress,omitempty"`
		// The duration of the running configuration or workflow.
		Duration *string `json:"Duration,omitempty"`
		// The start time of the configuration or workflow.
		StartTime      *string                                     `json:"StartTime,omitempty"`
		ClusterProfile NullableHyperflexClusterProfileRelationship `json:"ClusterProfile,omitempty"`
		// An array of relationships to hyperflexConfigResultEntry resources.
		ResultEntries []HyperflexConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	}

	varHyperflexConfigResultWithoutEmbeddedStruct := HyperflexConfigResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexConfigResultWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexConfigResult := _HyperflexConfigResult{}
		varHyperflexConfigResult.ClassId = varHyperflexConfigResultWithoutEmbeddedStruct.ClassId
		varHyperflexConfigResult.ObjectType = varHyperflexConfigResultWithoutEmbeddedStruct.ObjectType
		varHyperflexConfigResult.ConfigProgress = varHyperflexConfigResultWithoutEmbeddedStruct.ConfigProgress
		varHyperflexConfigResult.Duration = varHyperflexConfigResultWithoutEmbeddedStruct.Duration
		varHyperflexConfigResult.StartTime = varHyperflexConfigResultWithoutEmbeddedStruct.StartTime
		varHyperflexConfigResult.ClusterProfile = varHyperflexConfigResultWithoutEmbeddedStruct.ClusterProfile
		varHyperflexConfigResult.ResultEntries = varHyperflexConfigResultWithoutEmbeddedStruct.ResultEntries
		*o = HyperflexConfigResult(varHyperflexConfigResult)
	} else {
		return err
	}

	varHyperflexConfigResult := _HyperflexConfigResult{}

	err = json.Unmarshal(data, &varHyperflexConfigResult)
	if err == nil {
		o.PolicyAbstractConfigResult = varHyperflexConfigResult.PolicyAbstractConfigResult
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigProgress")
		delete(additionalProperties, "Duration")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "ClusterProfile")
		delete(additionalProperties, "ResultEntries")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigResult := reflect.ValueOf(o.PolicyAbstractConfigResult)
		for i := 0; i < reflectPolicyAbstractConfigResult.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigResult.Type().Field(i)

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

type NullableHyperflexConfigResult struct {
	value *HyperflexConfigResult
	isSet bool
}

func (v NullableHyperflexConfigResult) Get() *HyperflexConfigResult {
	return v.value
}

func (v *NullableHyperflexConfigResult) Set(val *HyperflexConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexConfigResult(val *HyperflexConfigResult) *NullableHyperflexConfigResult {
	return &NullableHyperflexConfigResult{value: val, isSet: true}
}

func (v NullableHyperflexConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
