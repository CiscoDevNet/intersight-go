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

// checks if the StorageNetAppHighAvailability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppHighAvailability{}

// StorageNetAppHighAvailability Storage failover and giveback information.
type StorageNetAppHighAvailability struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specifies whether or not giveback is automatically initiated when the node that owns the storage is ready.
	AutoGivebackEnabled *bool `json:"AutoGivebackEnabled,omitempty"`
	// Specifies whether or not storage failover is enabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// The state of the node that is giving storage back to its HA partner. * `unknown` - Default unknown giveback state. * `nothing_to_giveback` - The node has nothing to give back to its HA partner. * `not_attempted` - The node has not attempted to give back storage to its HA partner. * `in_progress` - The node is in progress of giving back storage to its HA partner. * `failed` - The node has failed to give back storage to its HA partner.
	GivebackState *string `json:"GivebackState,omitempty"`
	// The model of the partner in this node's High Availability (HA) group.
	PartnerModel *string `json:"PartnerModel,omitempty"`
	// The partner node name in this node's High Availability (HA) group.
	PartnerName *string `json:"PartnerName,omitempty"`
	// The partner node uuid in this node's High Availability (HA) group.
	PartnerUuid *string `json:"PartnerUuid,omitempty"`
	// The state of the node that is taking over storage from its HA partner. * `unknown` - Default unknown takeover state. * `not_possible` - It is not possible for the node to take over storage from its HA partner. * `not_attempted` - The node has not attempted to take over storage from its HA partner. * `in_takeover` - The node has taken over storage from its HA partner. * `in_progress` - The node is in progress of taking over storage from its HA partner. * `failed` - The node has failed to take over storage from its HA partner.
	TakeoverState        *string `json:"TakeoverState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppHighAvailability StorageNetAppHighAvailability

// NewStorageNetAppHighAvailability instantiates a new StorageNetAppHighAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppHighAvailability(classId string, objectType string) *StorageNetAppHighAvailability {
	this := StorageNetAppHighAvailability{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppHighAvailabilityWithDefaults instantiates a new StorageNetAppHighAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppHighAvailabilityWithDefaults() *StorageNetAppHighAvailability {
	this := StorageNetAppHighAvailability{}
	var classId string = "storage.NetAppHighAvailability"
	this.ClassId = classId
	var objectType string = "storage.NetAppHighAvailability"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppHighAvailability) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppHighAvailability) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppHighAvailability" of the ClassId field.
func (o *StorageNetAppHighAvailability) GetDefaultClassId() interface{} {
	return "storage.NetAppHighAvailability"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppHighAvailability) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppHighAvailability) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppHighAvailability" of the ObjectType field.
func (o *StorageNetAppHighAvailability) GetDefaultObjectType() interface{} {
	return "storage.NetAppHighAvailability"
}

// GetAutoGivebackEnabled returns the AutoGivebackEnabled field value if set, zero value otherwise.
func (o *StorageNetAppHighAvailability) GetAutoGivebackEnabled() bool {
	if o == nil || IsNil(o.AutoGivebackEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoGivebackEnabled
}

// GetAutoGivebackEnabledOk returns a tuple with the AutoGivebackEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetAutoGivebackEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoGivebackEnabled) {
		return nil, false
	}
	return o.AutoGivebackEnabled, true
}

// HasAutoGivebackEnabled returns a boolean if a field has been set.
func (o *StorageNetAppHighAvailability) HasAutoGivebackEnabled() bool {
	if o != nil && !IsNil(o.AutoGivebackEnabled) {
		return true
	}

	return false
}

// SetAutoGivebackEnabled gets a reference to the given bool and assigns it to the AutoGivebackEnabled field.
func (o *StorageNetAppHighAvailability) SetAutoGivebackEnabled(v bool) {
	o.AutoGivebackEnabled = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppHighAvailability) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppHighAvailability) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *StorageNetAppHighAvailability) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetGivebackState returns the GivebackState field value if set, zero value otherwise.
func (o *StorageNetAppHighAvailability) GetGivebackState() string {
	if o == nil || IsNil(o.GivebackState) {
		var ret string
		return ret
	}
	return *o.GivebackState
}

// GetGivebackStateOk returns a tuple with the GivebackState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetGivebackStateOk() (*string, bool) {
	if o == nil || IsNil(o.GivebackState) {
		return nil, false
	}
	return o.GivebackState, true
}

// HasGivebackState returns a boolean if a field has been set.
func (o *StorageNetAppHighAvailability) HasGivebackState() bool {
	if o != nil && !IsNil(o.GivebackState) {
		return true
	}

	return false
}

// SetGivebackState gets a reference to the given string and assigns it to the GivebackState field.
func (o *StorageNetAppHighAvailability) SetGivebackState(v string) {
	o.GivebackState = &v
}

// GetPartnerModel returns the PartnerModel field value if set, zero value otherwise.
func (o *StorageNetAppHighAvailability) GetPartnerModel() string {
	if o == nil || IsNil(o.PartnerModel) {
		var ret string
		return ret
	}
	return *o.PartnerModel
}

// GetPartnerModelOk returns a tuple with the PartnerModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetPartnerModelOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerModel) {
		return nil, false
	}
	return o.PartnerModel, true
}

// HasPartnerModel returns a boolean if a field has been set.
func (o *StorageNetAppHighAvailability) HasPartnerModel() bool {
	if o != nil && !IsNil(o.PartnerModel) {
		return true
	}

	return false
}

// SetPartnerModel gets a reference to the given string and assigns it to the PartnerModel field.
func (o *StorageNetAppHighAvailability) SetPartnerModel(v string) {
	o.PartnerModel = &v
}

// GetPartnerName returns the PartnerName field value if set, zero value otherwise.
func (o *StorageNetAppHighAvailability) GetPartnerName() string {
	if o == nil || IsNil(o.PartnerName) {
		var ret string
		return ret
	}
	return *o.PartnerName
}

// GetPartnerNameOk returns a tuple with the PartnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetPartnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerName) {
		return nil, false
	}
	return o.PartnerName, true
}

// HasPartnerName returns a boolean if a field has been set.
func (o *StorageNetAppHighAvailability) HasPartnerName() bool {
	if o != nil && !IsNil(o.PartnerName) {
		return true
	}

	return false
}

// SetPartnerName gets a reference to the given string and assigns it to the PartnerName field.
func (o *StorageNetAppHighAvailability) SetPartnerName(v string) {
	o.PartnerName = &v
}

// GetPartnerUuid returns the PartnerUuid field value if set, zero value otherwise.
func (o *StorageNetAppHighAvailability) GetPartnerUuid() string {
	if o == nil || IsNil(o.PartnerUuid) {
		var ret string
		return ret
	}
	return *o.PartnerUuid
}

// GetPartnerUuidOk returns a tuple with the PartnerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetPartnerUuidOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerUuid) {
		return nil, false
	}
	return o.PartnerUuid, true
}

// HasPartnerUuid returns a boolean if a field has been set.
func (o *StorageNetAppHighAvailability) HasPartnerUuid() bool {
	if o != nil && !IsNil(o.PartnerUuid) {
		return true
	}

	return false
}

// SetPartnerUuid gets a reference to the given string and assigns it to the PartnerUuid field.
func (o *StorageNetAppHighAvailability) SetPartnerUuid(v string) {
	o.PartnerUuid = &v
}

// GetTakeoverState returns the TakeoverState field value if set, zero value otherwise.
func (o *StorageNetAppHighAvailability) GetTakeoverState() string {
	if o == nil || IsNil(o.TakeoverState) {
		var ret string
		return ret
	}
	return *o.TakeoverState
}

// GetTakeoverStateOk returns a tuple with the TakeoverState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetTakeoverStateOk() (*string, bool) {
	if o == nil || IsNil(o.TakeoverState) {
		return nil, false
	}
	return o.TakeoverState, true
}

// HasTakeoverState returns a boolean if a field has been set.
func (o *StorageNetAppHighAvailability) HasTakeoverState() bool {
	if o != nil && !IsNil(o.TakeoverState) {
		return true
	}

	return false
}

// SetTakeoverState gets a reference to the given string and assigns it to the TakeoverState field.
func (o *StorageNetAppHighAvailability) SetTakeoverState(v string) {
	o.TakeoverState = &v
}

func (o StorageNetAppHighAvailability) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppHighAvailability) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AutoGivebackEnabled) {
		toSerialize["AutoGivebackEnabled"] = o.AutoGivebackEnabled
	}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.GivebackState) {
		toSerialize["GivebackState"] = o.GivebackState
	}
	if !IsNil(o.PartnerModel) {
		toSerialize["PartnerModel"] = o.PartnerModel
	}
	if !IsNil(o.PartnerName) {
		toSerialize["PartnerName"] = o.PartnerName
	}
	if !IsNil(o.PartnerUuid) {
		toSerialize["PartnerUuid"] = o.PartnerUuid
	}
	if !IsNil(o.TakeoverState) {
		toSerialize["TakeoverState"] = o.TakeoverState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppHighAvailability) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppHighAvailabilityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Specifies whether or not giveback is automatically initiated when the node that owns the storage is ready.
		AutoGivebackEnabled *bool `json:"AutoGivebackEnabled,omitempty"`
		// Specifies whether or not storage failover is enabled.
		Enabled *bool `json:"Enabled,omitempty"`
		// The state of the node that is giving storage back to its HA partner. * `unknown` - Default unknown giveback state. * `nothing_to_giveback` - The node has nothing to give back to its HA partner. * `not_attempted` - The node has not attempted to give back storage to its HA partner. * `in_progress` - The node is in progress of giving back storage to its HA partner. * `failed` - The node has failed to give back storage to its HA partner.
		GivebackState *string `json:"GivebackState,omitempty"`
		// The model of the partner in this node's High Availability (HA) group.
		PartnerModel *string `json:"PartnerModel,omitempty"`
		// The partner node name in this node's High Availability (HA) group.
		PartnerName *string `json:"PartnerName,omitempty"`
		// The partner node uuid in this node's High Availability (HA) group.
		PartnerUuid *string `json:"PartnerUuid,omitempty"`
		// The state of the node that is taking over storage from its HA partner. * `unknown` - Default unknown takeover state. * `not_possible` - It is not possible for the node to take over storage from its HA partner. * `not_attempted` - The node has not attempted to take over storage from its HA partner. * `in_takeover` - The node has taken over storage from its HA partner. * `in_progress` - The node is in progress of taking over storage from its HA partner. * `failed` - The node has failed to take over storage from its HA partner.
		TakeoverState *string `json:"TakeoverState,omitempty"`
	}

	varStorageNetAppHighAvailabilityWithoutEmbeddedStruct := StorageNetAppHighAvailabilityWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppHighAvailabilityWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppHighAvailability := _StorageNetAppHighAvailability{}
		varStorageNetAppHighAvailability.ClassId = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.ClassId
		varStorageNetAppHighAvailability.ObjectType = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.ObjectType
		varStorageNetAppHighAvailability.AutoGivebackEnabled = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.AutoGivebackEnabled
		varStorageNetAppHighAvailability.Enabled = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.Enabled
		varStorageNetAppHighAvailability.GivebackState = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.GivebackState
		varStorageNetAppHighAvailability.PartnerModel = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.PartnerModel
		varStorageNetAppHighAvailability.PartnerName = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.PartnerName
		varStorageNetAppHighAvailability.PartnerUuid = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.PartnerUuid
		varStorageNetAppHighAvailability.TakeoverState = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.TakeoverState
		*o = StorageNetAppHighAvailability(varStorageNetAppHighAvailability)
	} else {
		return err
	}

	varStorageNetAppHighAvailability := _StorageNetAppHighAvailability{}

	err = json.Unmarshal(data, &varStorageNetAppHighAvailability)
	if err == nil {
		o.MoBaseComplexType = varStorageNetAppHighAvailability.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutoGivebackEnabled")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "GivebackState")
		delete(additionalProperties, "PartnerModel")
		delete(additionalProperties, "PartnerName")
		delete(additionalProperties, "PartnerUuid")
		delete(additionalProperties, "TakeoverState")

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

type NullableStorageNetAppHighAvailability struct {
	value *StorageNetAppHighAvailability
	isSet bool
}

func (v NullableStorageNetAppHighAvailability) Get() *StorageNetAppHighAvailability {
	return v.value
}

func (v *NullableStorageNetAppHighAvailability) Set(val *StorageNetAppHighAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppHighAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppHighAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppHighAvailability(val *StorageNetAppHighAvailability) *NullableStorageNetAppHighAvailability {
	return &NullableStorageNetAppHighAvailability{value: val, isSet: true}
}

func (v NullableStorageNetAppHighAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppHighAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
