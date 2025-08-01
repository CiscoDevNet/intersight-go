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

// checks if the HyperflexStPlatformClusterResiliencyInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexStPlatformClusterResiliencyInfo{}

// HyperflexStPlatformClusterResiliencyInfo The detailed status of the cluster's resiliency health.
type HyperflexStPlatformClusterResiliencyInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of persistent storage device failures tolerable before the storage cluster becomes offline.
	HddFailuresTolerable *int64   `json:"HddFailuresTolerable,omitempty"`
	Messages             []string `json:"Messages,omitempty"`
	// The current message describing the auto-healing process of the cluster.
	MessagesIterator interface{} `json:"MessagesIterator,omitempty"`
	// The number of elements in the messages collection.
	MessagesSize *int64 `json:"MessagesSize,omitempty"`
	// The number of node failures tolerable before the storage cluster becomes offline.
	NodeFailuresTolerable *int64 `json:"NodeFailuresTolerable,omitempty"`
	// The number of caching device failures tolerable before the storage cluster becomes offline.
	SsdFailuresTolerable *int64 `json:"SsdFailuresTolerable,omitempty"`
	// The resiliency state of the cluster. The resiliency status is 'HEALTHY' if there are no failures and the storage cluster is fully operational. The resiliency status is 'WARNING' when the cluster has experienced failures that may adversely affect the cluster. It is 'UNKNOWN' if the cluster is offline or if the state cannot be determined.
	State                *string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexStPlatformClusterResiliencyInfo HyperflexStPlatformClusterResiliencyInfo

// NewHyperflexStPlatformClusterResiliencyInfo instantiates a new HyperflexStPlatformClusterResiliencyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexStPlatformClusterResiliencyInfo(classId string, objectType string) *HyperflexStPlatformClusterResiliencyInfo {
	this := HyperflexStPlatformClusterResiliencyInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexStPlatformClusterResiliencyInfoWithDefaults instantiates a new HyperflexStPlatformClusterResiliencyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexStPlatformClusterResiliencyInfoWithDefaults() *HyperflexStPlatformClusterResiliencyInfo {
	this := HyperflexStPlatformClusterResiliencyInfo{}
	var classId string = "hyperflex.StPlatformClusterResiliencyInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.StPlatformClusterResiliencyInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexStPlatformClusterResiliencyInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexStPlatformClusterResiliencyInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.StPlatformClusterResiliencyInfo" of the ClassId field.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetDefaultClassId() interface{} {
	return "hyperflex.StPlatformClusterResiliencyInfo"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexStPlatformClusterResiliencyInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexStPlatformClusterResiliencyInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.StPlatformClusterResiliencyInfo" of the ObjectType field.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetDefaultObjectType() interface{} {
	return "hyperflex.StPlatformClusterResiliencyInfo"
}

// GetHddFailuresTolerable returns the HddFailuresTolerable field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetHddFailuresTolerable() int64 {
	if o == nil || IsNil(o.HddFailuresTolerable) {
		var ret int64
		return ret
	}
	return *o.HddFailuresTolerable
}

// GetHddFailuresTolerableOk returns a tuple with the HddFailuresTolerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetHddFailuresTolerableOk() (*int64, bool) {
	if o == nil || IsNil(o.HddFailuresTolerable) {
		return nil, false
	}
	return o.HddFailuresTolerable, true
}

// HasHddFailuresTolerable returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) HasHddFailuresTolerable() bool {
	if o != nil && !IsNil(o.HddFailuresTolerable) {
		return true
	}

	return false
}

// SetHddFailuresTolerable gets a reference to the given int64 and assigns it to the HddFailuresTolerable field.
func (o *HyperflexStPlatformClusterResiliencyInfo) SetHddFailuresTolerable(v int64) {
	o.HddFailuresTolerable = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexStPlatformClusterResiliencyInfo) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexStPlatformClusterResiliencyInfo) GetMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *HyperflexStPlatformClusterResiliencyInfo) SetMessages(v []string) {
	o.Messages = v
}

// GetMessagesIterator returns the MessagesIterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexStPlatformClusterResiliencyInfo) GetMessagesIterator() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MessagesIterator
}

// GetMessagesIteratorOk returns a tuple with the MessagesIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexStPlatformClusterResiliencyInfo) GetMessagesIteratorOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MessagesIterator) {
		return nil, false
	}
	return &o.MessagesIterator, true
}

// HasMessagesIterator returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) HasMessagesIterator() bool {
	if o != nil && !IsNil(o.MessagesIterator) {
		return true
	}

	return false
}

// SetMessagesIterator gets a reference to the given interface{} and assigns it to the MessagesIterator field.
func (o *HyperflexStPlatformClusterResiliencyInfo) SetMessagesIterator(v interface{}) {
	o.MessagesIterator = v
}

// GetMessagesSize returns the MessagesSize field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetMessagesSize() int64 {
	if o == nil || IsNil(o.MessagesSize) {
		var ret int64
		return ret
	}
	return *o.MessagesSize
}

// GetMessagesSizeOk returns a tuple with the MessagesSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetMessagesSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MessagesSize) {
		return nil, false
	}
	return o.MessagesSize, true
}

// HasMessagesSize returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) HasMessagesSize() bool {
	if o != nil && !IsNil(o.MessagesSize) {
		return true
	}

	return false
}

// SetMessagesSize gets a reference to the given int64 and assigns it to the MessagesSize field.
func (o *HyperflexStPlatformClusterResiliencyInfo) SetMessagesSize(v int64) {
	o.MessagesSize = &v
}

// GetNodeFailuresTolerable returns the NodeFailuresTolerable field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetNodeFailuresTolerable() int64 {
	if o == nil || IsNil(o.NodeFailuresTolerable) {
		var ret int64
		return ret
	}
	return *o.NodeFailuresTolerable
}

// GetNodeFailuresTolerableOk returns a tuple with the NodeFailuresTolerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetNodeFailuresTolerableOk() (*int64, bool) {
	if o == nil || IsNil(o.NodeFailuresTolerable) {
		return nil, false
	}
	return o.NodeFailuresTolerable, true
}

// HasNodeFailuresTolerable returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) HasNodeFailuresTolerable() bool {
	if o != nil && !IsNil(o.NodeFailuresTolerable) {
		return true
	}

	return false
}

// SetNodeFailuresTolerable gets a reference to the given int64 and assigns it to the NodeFailuresTolerable field.
func (o *HyperflexStPlatformClusterResiliencyInfo) SetNodeFailuresTolerable(v int64) {
	o.NodeFailuresTolerable = &v
}

// GetSsdFailuresTolerable returns the SsdFailuresTolerable field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetSsdFailuresTolerable() int64 {
	if o == nil || IsNil(o.SsdFailuresTolerable) {
		var ret int64
		return ret
	}
	return *o.SsdFailuresTolerable
}

// GetSsdFailuresTolerableOk returns a tuple with the SsdFailuresTolerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetSsdFailuresTolerableOk() (*int64, bool) {
	if o == nil || IsNil(o.SsdFailuresTolerable) {
		return nil, false
	}
	return o.SsdFailuresTolerable, true
}

// HasSsdFailuresTolerable returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) HasSsdFailuresTolerable() bool {
	if o != nil && !IsNil(o.SsdFailuresTolerable) {
		return true
	}

	return false
}

// SetSsdFailuresTolerable gets a reference to the given int64 and assigns it to the SsdFailuresTolerable field.
func (o *HyperflexStPlatformClusterResiliencyInfo) SetSsdFailuresTolerable(v int64) {
	o.SsdFailuresTolerable = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterResiliencyInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *HyperflexStPlatformClusterResiliencyInfo) SetState(v string) {
	o.State = &v
}

func (o HyperflexStPlatformClusterResiliencyInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexStPlatformClusterResiliencyInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.HddFailuresTolerable) {
		toSerialize["HddFailuresTolerable"] = o.HddFailuresTolerable
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.MessagesIterator != nil {
		toSerialize["MessagesIterator"] = o.MessagesIterator
	}
	if !IsNil(o.MessagesSize) {
		toSerialize["MessagesSize"] = o.MessagesSize
	}
	if !IsNil(o.NodeFailuresTolerable) {
		toSerialize["NodeFailuresTolerable"] = o.NodeFailuresTolerable
	}
	if !IsNil(o.SsdFailuresTolerable) {
		toSerialize["SsdFailuresTolerable"] = o.SsdFailuresTolerable
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexStPlatformClusterResiliencyInfo) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of persistent storage device failures tolerable before the storage cluster becomes offline.
		HddFailuresTolerable *int64   `json:"HddFailuresTolerable,omitempty"`
		Messages             []string `json:"Messages,omitempty"`
		// The current message describing the auto-healing process of the cluster.
		MessagesIterator interface{} `json:"MessagesIterator,omitempty"`
		// The number of elements in the messages collection.
		MessagesSize *int64 `json:"MessagesSize,omitempty"`
		// The number of node failures tolerable before the storage cluster becomes offline.
		NodeFailuresTolerable *int64 `json:"NodeFailuresTolerable,omitempty"`
		// The number of caching device failures tolerable before the storage cluster becomes offline.
		SsdFailuresTolerable *int64 `json:"SsdFailuresTolerable,omitempty"`
		// The resiliency state of the cluster. The resiliency status is 'HEALTHY' if there are no failures and the storage cluster is fully operational. The resiliency status is 'WARNING' when the cluster has experienced failures that may adversely affect the cluster. It is 'UNKNOWN' if the cluster is offline or if the state cannot be determined.
		State *string `json:"State,omitempty"`
	}

	varHyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct := HyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexStPlatformClusterResiliencyInfo := _HyperflexStPlatformClusterResiliencyInfo{}
		varHyperflexStPlatformClusterResiliencyInfo.ClassId = varHyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct.ClassId
		varHyperflexStPlatformClusterResiliencyInfo.ObjectType = varHyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct.ObjectType
		varHyperflexStPlatformClusterResiliencyInfo.HddFailuresTolerable = varHyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct.HddFailuresTolerable
		varHyperflexStPlatformClusterResiliencyInfo.Messages = varHyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct.Messages
		varHyperflexStPlatformClusterResiliencyInfo.MessagesIterator = varHyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct.MessagesIterator
		varHyperflexStPlatformClusterResiliencyInfo.MessagesSize = varHyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct.MessagesSize
		varHyperflexStPlatformClusterResiliencyInfo.NodeFailuresTolerable = varHyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct.NodeFailuresTolerable
		varHyperflexStPlatformClusterResiliencyInfo.SsdFailuresTolerable = varHyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct.SsdFailuresTolerable
		varHyperflexStPlatformClusterResiliencyInfo.State = varHyperflexStPlatformClusterResiliencyInfoWithoutEmbeddedStruct.State
		*o = HyperflexStPlatformClusterResiliencyInfo(varHyperflexStPlatformClusterResiliencyInfo)
	} else {
		return err
	}

	varHyperflexStPlatformClusterResiliencyInfo := _HyperflexStPlatformClusterResiliencyInfo{}

	err = json.Unmarshal(data, &varHyperflexStPlatformClusterResiliencyInfo)
	if err == nil {
		o.MoBaseComplexType = varHyperflexStPlatformClusterResiliencyInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HddFailuresTolerable")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "MessagesIterator")
		delete(additionalProperties, "MessagesSize")
		delete(additionalProperties, "NodeFailuresTolerable")
		delete(additionalProperties, "SsdFailuresTolerable")
		delete(additionalProperties, "State")

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

type NullableHyperflexStPlatformClusterResiliencyInfo struct {
	value *HyperflexStPlatformClusterResiliencyInfo
	isSet bool
}

func (v NullableHyperflexStPlatformClusterResiliencyInfo) Get() *HyperflexStPlatformClusterResiliencyInfo {
	return v.value
}

func (v *NullableHyperflexStPlatformClusterResiliencyInfo) Set(val *HyperflexStPlatformClusterResiliencyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexStPlatformClusterResiliencyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexStPlatformClusterResiliencyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexStPlatformClusterResiliencyInfo(val *HyperflexStPlatformClusterResiliencyInfo) *NullableHyperflexStPlatformClusterResiliencyInfo {
	return &NullableHyperflexStPlatformClusterResiliencyInfo{value: val, isSet: true}
}

func (v NullableHyperflexStPlatformClusterResiliencyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexStPlatformClusterResiliencyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
