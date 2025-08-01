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

// checks if the StorageExternalPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageExternalPath{}

// StorageExternalPath The following properties are restored for each external path.
type StorageExternalPath struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The time (in seconds) until the external parity group is blocked after all paths to the external parity group are disconnected.
	BlockedPathMonitoring *int64 `json:"BlockedPathMonitoring,omitempty"`
	// WWN of the external storage system.
	ExternalWwn *string `json:"ExternalWwn,omitempty"`
	// The value (in seconds) set for the I/O time over for the external parity group.
	IoTimeOut *int64 `json:"IoTimeOut,omitempty"`
	// Port number of external path on the local storage system.
	PortId *string `json:"PortId,omitempty"`
	// Number of Read-Write commands that can be queued to the external parity group.
	QueueDepth           *int64 `json:"QueueDepth,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageExternalPath StorageExternalPath

// NewStorageExternalPath instantiates a new StorageExternalPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageExternalPath(classId string, objectType string) *StorageExternalPath {
	this := StorageExternalPath{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageExternalPathWithDefaults instantiates a new StorageExternalPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageExternalPathWithDefaults() *StorageExternalPath {
	this := StorageExternalPath{}
	var classId string = "storage.ExternalPath"
	this.ClassId = classId
	var objectType string = "storage.ExternalPath"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageExternalPath) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageExternalPath) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.ExternalPath" of the ClassId field.
func (o *StorageExternalPath) GetDefaultClassId() interface{} {
	return "storage.ExternalPath"
}

// GetObjectType returns the ObjectType field value
func (o *StorageExternalPath) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageExternalPath) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.ExternalPath" of the ObjectType field.
func (o *StorageExternalPath) GetDefaultObjectType() interface{} {
	return "storage.ExternalPath"
}

// GetBlockedPathMonitoring returns the BlockedPathMonitoring field value if set, zero value otherwise.
func (o *StorageExternalPath) GetBlockedPathMonitoring() int64 {
	if o == nil || IsNil(o.BlockedPathMonitoring) {
		var ret int64
		return ret
	}
	return *o.BlockedPathMonitoring
}

// GetBlockedPathMonitoringOk returns a tuple with the BlockedPathMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetBlockedPathMonitoringOk() (*int64, bool) {
	if o == nil || IsNil(o.BlockedPathMonitoring) {
		return nil, false
	}
	return o.BlockedPathMonitoring, true
}

// HasBlockedPathMonitoring returns a boolean if a field has been set.
func (o *StorageExternalPath) HasBlockedPathMonitoring() bool {
	if o != nil && !IsNil(o.BlockedPathMonitoring) {
		return true
	}

	return false
}

// SetBlockedPathMonitoring gets a reference to the given int64 and assigns it to the BlockedPathMonitoring field.
func (o *StorageExternalPath) SetBlockedPathMonitoring(v int64) {
	o.BlockedPathMonitoring = &v
}

// GetExternalWwn returns the ExternalWwn field value if set, zero value otherwise.
func (o *StorageExternalPath) GetExternalWwn() string {
	if o == nil || IsNil(o.ExternalWwn) {
		var ret string
		return ret
	}
	return *o.ExternalWwn
}

// GetExternalWwnOk returns a tuple with the ExternalWwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetExternalWwnOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalWwn) {
		return nil, false
	}
	return o.ExternalWwn, true
}

// HasExternalWwn returns a boolean if a field has been set.
func (o *StorageExternalPath) HasExternalWwn() bool {
	if o != nil && !IsNil(o.ExternalWwn) {
		return true
	}

	return false
}

// SetExternalWwn gets a reference to the given string and assigns it to the ExternalWwn field.
func (o *StorageExternalPath) SetExternalWwn(v string) {
	o.ExternalWwn = &v
}

// GetIoTimeOut returns the IoTimeOut field value if set, zero value otherwise.
func (o *StorageExternalPath) GetIoTimeOut() int64 {
	if o == nil || IsNil(o.IoTimeOut) {
		var ret int64
		return ret
	}
	return *o.IoTimeOut
}

// GetIoTimeOutOk returns a tuple with the IoTimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetIoTimeOutOk() (*int64, bool) {
	if o == nil || IsNil(o.IoTimeOut) {
		return nil, false
	}
	return o.IoTimeOut, true
}

// HasIoTimeOut returns a boolean if a field has been set.
func (o *StorageExternalPath) HasIoTimeOut() bool {
	if o != nil && !IsNil(o.IoTimeOut) {
		return true
	}

	return false
}

// SetIoTimeOut gets a reference to the given int64 and assigns it to the IoTimeOut field.
func (o *StorageExternalPath) SetIoTimeOut(v int64) {
	o.IoTimeOut = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *StorageExternalPath) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *StorageExternalPath) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *StorageExternalPath) SetPortId(v string) {
	o.PortId = &v
}

// GetQueueDepth returns the QueueDepth field value if set, zero value otherwise.
func (o *StorageExternalPath) GetQueueDepth() int64 {
	if o == nil || IsNil(o.QueueDepth) {
		var ret int64
		return ret
	}
	return *o.QueueDepth
}

// GetQueueDepthOk returns a tuple with the QueueDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetQueueDepthOk() (*int64, bool) {
	if o == nil || IsNil(o.QueueDepth) {
		return nil, false
	}
	return o.QueueDepth, true
}

// HasQueueDepth returns a boolean if a field has been set.
func (o *StorageExternalPath) HasQueueDepth() bool {
	if o != nil && !IsNil(o.QueueDepth) {
		return true
	}

	return false
}

// SetQueueDepth gets a reference to the given int64 and assigns it to the QueueDepth field.
func (o *StorageExternalPath) SetQueueDepth(v int64) {
	o.QueueDepth = &v
}

func (o StorageExternalPath) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageExternalPath) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BlockedPathMonitoring) {
		toSerialize["BlockedPathMonitoring"] = o.BlockedPathMonitoring
	}
	if !IsNil(o.ExternalWwn) {
		toSerialize["ExternalWwn"] = o.ExternalWwn
	}
	if !IsNil(o.IoTimeOut) {
		toSerialize["IoTimeOut"] = o.IoTimeOut
	}
	if !IsNil(o.PortId) {
		toSerialize["PortId"] = o.PortId
	}
	if !IsNil(o.QueueDepth) {
		toSerialize["QueueDepth"] = o.QueueDepth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageExternalPath) UnmarshalJSON(data []byte) (err error) {
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
	type StorageExternalPathWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The time (in seconds) until the external parity group is blocked after all paths to the external parity group are disconnected.
		BlockedPathMonitoring *int64 `json:"BlockedPathMonitoring,omitempty"`
		// WWN of the external storage system.
		ExternalWwn *string `json:"ExternalWwn,omitempty"`
		// The value (in seconds) set for the I/O time over for the external parity group.
		IoTimeOut *int64 `json:"IoTimeOut,omitempty"`
		// Port number of external path on the local storage system.
		PortId *string `json:"PortId,omitempty"`
		// Number of Read-Write commands that can be queued to the external parity group.
		QueueDepth *int64 `json:"QueueDepth,omitempty"`
	}

	varStorageExternalPathWithoutEmbeddedStruct := StorageExternalPathWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageExternalPathWithoutEmbeddedStruct)
	if err == nil {
		varStorageExternalPath := _StorageExternalPath{}
		varStorageExternalPath.ClassId = varStorageExternalPathWithoutEmbeddedStruct.ClassId
		varStorageExternalPath.ObjectType = varStorageExternalPathWithoutEmbeddedStruct.ObjectType
		varStorageExternalPath.BlockedPathMonitoring = varStorageExternalPathWithoutEmbeddedStruct.BlockedPathMonitoring
		varStorageExternalPath.ExternalWwn = varStorageExternalPathWithoutEmbeddedStruct.ExternalWwn
		varStorageExternalPath.IoTimeOut = varStorageExternalPathWithoutEmbeddedStruct.IoTimeOut
		varStorageExternalPath.PortId = varStorageExternalPathWithoutEmbeddedStruct.PortId
		varStorageExternalPath.QueueDepth = varStorageExternalPathWithoutEmbeddedStruct.QueueDepth
		*o = StorageExternalPath(varStorageExternalPath)
	} else {
		return err
	}

	varStorageExternalPath := _StorageExternalPath{}

	err = json.Unmarshal(data, &varStorageExternalPath)
	if err == nil {
		o.MoBaseComplexType = varStorageExternalPath.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BlockedPathMonitoring")
		delete(additionalProperties, "ExternalWwn")
		delete(additionalProperties, "IoTimeOut")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "QueueDepth")

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

type NullableStorageExternalPath struct {
	value *StorageExternalPath
	isSet bool
}

func (v NullableStorageExternalPath) Get() *StorageExternalPath {
	return v.value
}

func (v *NullableStorageExternalPath) Set(val *StorageExternalPath) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageExternalPath) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageExternalPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageExternalPath(val *StorageExternalPath) *NullableStorageExternalPath {
	return &NullableStorageExternalPath{value: val, isSet: true}
}

func (v NullableStorageExternalPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageExternalPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
