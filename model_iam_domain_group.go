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

// checks if the IamDomainGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamDomainGroup{}

// IamDomainGroup Intersight services are mapped to three different categories of services for scaling purpose. Three categories are defined: Partition1/Partition2/Partition3. Topics for each category are created with a specific number of partitions. For each cloud environment these numbers will be different.
type IamDomainGroup struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the domain-group.
	Name *string `json:"Name,omitempty"`
	// The partition number domain group related messages are produced for 'Partition1' category service topics.
	Partition1 *int64 `json:"Partition1,omitempty"`
	// In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for 'Partition2' category service topics.
	Partition2 *int64 `json:"Partition2,omitempty"`
	// In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for 'Partition3' category service topics.
	Partition3 *int64 `json:"Partition3,omitempty"`
	// Partition key used for producing messages to Kafka partitions. By default Domain-group id will be used as partition key. For Domain-groups belonging to Early adopters domain-group id will be prefixed with 'H' and used as partition key, such partition key will be treated differently and messages will always be produced to partition 0.
	PartitionKey *string `json:"PartitionKey,omitempty"`
	// The number of devices in the domain-group. Device registration notifications are processed to update the usage of the domain-group. The on-boarding account will have multiple domain-groups, and during the device registration least used domain-group will be selected for the device.
	Usage                *int64                         `json:"Usage,omitempty"`
	Account              NullableIamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamDomainGroup IamDomainGroup

// NewIamDomainGroup instantiates a new IamDomainGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamDomainGroup(classId string, objectType string) *IamDomainGroup {
	this := IamDomainGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamDomainGroupWithDefaults instantiates a new IamDomainGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamDomainGroupWithDefaults() *IamDomainGroup {
	this := IamDomainGroup{}
	var classId string = "iam.DomainGroup"
	this.ClassId = classId
	var objectType string = "iam.DomainGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamDomainGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamDomainGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamDomainGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.DomainGroup" of the ClassId field.
func (o *IamDomainGroup) GetDefaultClassId() interface{} {
	return "iam.DomainGroup"
}

// GetObjectType returns the ObjectType field value
func (o *IamDomainGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamDomainGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamDomainGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.DomainGroup" of the ObjectType field.
func (o *IamDomainGroup) GetDefaultObjectType() interface{} {
	return "iam.DomainGroup"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamDomainGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamDomainGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamDomainGroup) SetName(v string) {
	o.Name = &v
}

// GetPartition1 returns the Partition1 field value if set, zero value otherwise.
func (o *IamDomainGroup) GetPartition1() int64 {
	if o == nil || IsNil(o.Partition1) {
		var ret int64
		return ret
	}
	return *o.Partition1
}

// GetPartition1Ok returns a tuple with the Partition1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroup) GetPartition1Ok() (*int64, bool) {
	if o == nil || IsNil(o.Partition1) {
		return nil, false
	}
	return o.Partition1, true
}

// HasPartition1 returns a boolean if a field has been set.
func (o *IamDomainGroup) HasPartition1() bool {
	if o != nil && !IsNil(o.Partition1) {
		return true
	}

	return false
}

// SetPartition1 gets a reference to the given int64 and assigns it to the Partition1 field.
func (o *IamDomainGroup) SetPartition1(v int64) {
	o.Partition1 = &v
}

// GetPartition2 returns the Partition2 field value if set, zero value otherwise.
func (o *IamDomainGroup) GetPartition2() int64 {
	if o == nil || IsNil(o.Partition2) {
		var ret int64
		return ret
	}
	return *o.Partition2
}

// GetPartition2Ok returns a tuple with the Partition2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroup) GetPartition2Ok() (*int64, bool) {
	if o == nil || IsNil(o.Partition2) {
		return nil, false
	}
	return o.Partition2, true
}

// HasPartition2 returns a boolean if a field has been set.
func (o *IamDomainGroup) HasPartition2() bool {
	if o != nil && !IsNil(o.Partition2) {
		return true
	}

	return false
}

// SetPartition2 gets a reference to the given int64 and assigns it to the Partition2 field.
func (o *IamDomainGroup) SetPartition2(v int64) {
	o.Partition2 = &v
}

// GetPartition3 returns the Partition3 field value if set, zero value otherwise.
func (o *IamDomainGroup) GetPartition3() int64 {
	if o == nil || IsNil(o.Partition3) {
		var ret int64
		return ret
	}
	return *o.Partition3
}

// GetPartition3Ok returns a tuple with the Partition3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroup) GetPartition3Ok() (*int64, bool) {
	if o == nil || IsNil(o.Partition3) {
		return nil, false
	}
	return o.Partition3, true
}

// HasPartition3 returns a boolean if a field has been set.
func (o *IamDomainGroup) HasPartition3() bool {
	if o != nil && !IsNil(o.Partition3) {
		return true
	}

	return false
}

// SetPartition3 gets a reference to the given int64 and assigns it to the Partition3 field.
func (o *IamDomainGroup) SetPartition3(v int64) {
	o.Partition3 = &v
}

// GetPartitionKey returns the PartitionKey field value if set, zero value otherwise.
func (o *IamDomainGroup) GetPartitionKey() string {
	if o == nil || IsNil(o.PartitionKey) {
		var ret string
		return ret
	}
	return *o.PartitionKey
}

// GetPartitionKeyOk returns a tuple with the PartitionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroup) GetPartitionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PartitionKey) {
		return nil, false
	}
	return o.PartitionKey, true
}

// HasPartitionKey returns a boolean if a field has been set.
func (o *IamDomainGroup) HasPartitionKey() bool {
	if o != nil && !IsNil(o.PartitionKey) {
		return true
	}

	return false
}

// SetPartitionKey gets a reference to the given string and assigns it to the PartitionKey field.
func (o *IamDomainGroup) SetPartitionKey(v string) {
	o.PartitionKey = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *IamDomainGroup) GetUsage() int64 {
	if o == nil || IsNil(o.Usage) {
		var ret int64
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroup) GetUsageOk() (*int64, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *IamDomainGroup) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given int64 and assigns it to the Usage field.
func (o *IamDomainGroup) SetUsage(v int64) {
	o.Usage = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamDomainGroup) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamDomainGroup) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *IamDomainGroup) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *IamDomainGroup) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *IamDomainGroup) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *IamDomainGroup) UnsetAccount() {
	o.Account.Unset()
}

func (o IamDomainGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamDomainGroup) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Partition1) {
		toSerialize["Partition1"] = o.Partition1
	}
	if !IsNil(o.Partition2) {
		toSerialize["Partition2"] = o.Partition2
	}
	if !IsNil(o.Partition3) {
		toSerialize["Partition3"] = o.Partition3
	}
	if !IsNil(o.PartitionKey) {
		toSerialize["PartitionKey"] = o.PartitionKey
	}
	if !IsNil(o.Usage) {
		toSerialize["Usage"] = o.Usage
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamDomainGroup) UnmarshalJSON(data []byte) (err error) {
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
	type IamDomainGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the domain-group.
		Name *string `json:"Name,omitempty"`
		// The partition number domain group related messages are produced for 'Partition1' category service topics.
		Partition1 *int64 `json:"Partition1,omitempty"`
		// In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for 'Partition2' category service topics.
		Partition2 *int64 `json:"Partition2,omitempty"`
		// In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for 'Partition3' category service topics.
		Partition3 *int64 `json:"Partition3,omitempty"`
		// Partition key used for producing messages to Kafka partitions. By default Domain-group id will be used as partition key. For Domain-groups belonging to Early adopters domain-group id will be prefixed with 'H' and used as partition key, such partition key will be treated differently and messages will always be produced to partition 0.
		PartitionKey *string `json:"PartitionKey,omitempty"`
		// The number of devices in the domain-group. Device registration notifications are processed to update the usage of the domain-group. The on-boarding account will have multiple domain-groups, and during the device registration least used domain-group will be selected for the device.
		Usage   *int64                         `json:"Usage,omitempty"`
		Account NullableIamAccountRelationship `json:"Account,omitempty"`
	}

	varIamDomainGroupWithoutEmbeddedStruct := IamDomainGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamDomainGroupWithoutEmbeddedStruct)
	if err == nil {
		varIamDomainGroup := _IamDomainGroup{}
		varIamDomainGroup.ClassId = varIamDomainGroupWithoutEmbeddedStruct.ClassId
		varIamDomainGroup.ObjectType = varIamDomainGroupWithoutEmbeddedStruct.ObjectType
		varIamDomainGroup.Name = varIamDomainGroupWithoutEmbeddedStruct.Name
		varIamDomainGroup.Partition1 = varIamDomainGroupWithoutEmbeddedStruct.Partition1
		varIamDomainGroup.Partition2 = varIamDomainGroupWithoutEmbeddedStruct.Partition2
		varIamDomainGroup.Partition3 = varIamDomainGroupWithoutEmbeddedStruct.Partition3
		varIamDomainGroup.PartitionKey = varIamDomainGroupWithoutEmbeddedStruct.PartitionKey
		varIamDomainGroup.Usage = varIamDomainGroupWithoutEmbeddedStruct.Usage
		varIamDomainGroup.Account = varIamDomainGroupWithoutEmbeddedStruct.Account
		*o = IamDomainGroup(varIamDomainGroup)
	} else {
		return err
	}

	varIamDomainGroup := _IamDomainGroup{}

	err = json.Unmarshal(data, &varIamDomainGroup)
	if err == nil {
		o.MoBaseMo = varIamDomainGroup.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Partition1")
		delete(additionalProperties, "Partition2")
		delete(additionalProperties, "Partition3")
		delete(additionalProperties, "PartitionKey")
		delete(additionalProperties, "Usage")
		delete(additionalProperties, "Account")

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

type NullableIamDomainGroup struct {
	value *IamDomainGroup
	isSet bool
}

func (v NullableIamDomainGroup) Get() *IamDomainGroup {
	return v.value
}

func (v *NullableIamDomainGroup) Set(val *IamDomainGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableIamDomainGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableIamDomainGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamDomainGroup(val *IamDomainGroup) *NullableIamDomainGroup {
	return &NullableIamDomainGroup{value: val, isSet: true}
}

func (v NullableIamDomainGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamDomainGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
