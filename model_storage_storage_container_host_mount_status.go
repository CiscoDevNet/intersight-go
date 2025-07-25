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

// checks if the StorageStorageContainerHostMountStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageStorageContainerHostMountStatus{}

// StorageStorageContainerHostMountStatus Storage container host mount and accessibility status. Applicable only for NFS type.
type StorageStorageContainerHostMountStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Host specific storage container accessibility status. * `NOT_APPLICABLE` - The HyperFlex storage container accessibility on host is not applicable. * `ACCESSIBLE` - The HyperFlex storage container is accessible on the host. * `NOT_ACCESSIBLE` - The HyperFlex storage container is not accessible on the host. * `PARTIALLY_ACCESSIBLE` - The HyperFlex storage container is partially accessible on the host. * `READONLY` - The HyperFlex storage container accessibility is readonly on the host. * `HOST_NOT_REACHABLE` - The storage container is not accessible via this host because it is not accessible. * `UNKNOWN` - The storage container accessibility via this host is unknown.
	Accessibility *string `json:"Accessibility,omitempty"`
	// The name of the host corresponding to the mount and accessibility status of the storage container.
	HostName *string `json:"HostName,omitempty"`
	// Host specific storage container mount status.
	Mounted *bool `json:"Mounted,omitempty"`
	// Host specific storage container mount status reason.
	Reason               *string `json:"Reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageStorageContainerHostMountStatus StorageStorageContainerHostMountStatus

// NewStorageStorageContainerHostMountStatus instantiates a new StorageStorageContainerHostMountStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageStorageContainerHostMountStatus(classId string, objectType string) *StorageStorageContainerHostMountStatus {
	this := StorageStorageContainerHostMountStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageStorageContainerHostMountStatusWithDefaults instantiates a new StorageStorageContainerHostMountStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageStorageContainerHostMountStatusWithDefaults() *StorageStorageContainerHostMountStatus {
	this := StorageStorageContainerHostMountStatus{}
	var classId string = "storage.StorageContainerHostMountStatus"
	this.ClassId = classId
	var objectType string = "storage.StorageContainerHostMountStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageStorageContainerHostMountStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageStorageContainerHostMountStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageStorageContainerHostMountStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.StorageContainerHostMountStatus" of the ClassId field.
func (o *StorageStorageContainerHostMountStatus) GetDefaultClassId() interface{} {
	return "storage.StorageContainerHostMountStatus"
}

// GetObjectType returns the ObjectType field value
func (o *StorageStorageContainerHostMountStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageStorageContainerHostMountStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageStorageContainerHostMountStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.StorageContainerHostMountStatus" of the ObjectType field.
func (o *StorageStorageContainerHostMountStatus) GetDefaultObjectType() interface{} {
	return "storage.StorageContainerHostMountStatus"
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *StorageStorageContainerHostMountStatus) GetAccessibility() string {
	if o == nil || IsNil(o.Accessibility) {
		var ret string
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageContainerHostMountStatus) GetAccessibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Accessibility) {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *StorageStorageContainerHostMountStatus) HasAccessibility() bool {
	if o != nil && !IsNil(o.Accessibility) {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given string and assigns it to the Accessibility field.
func (o *StorageStorageContainerHostMountStatus) SetAccessibility(v string) {
	o.Accessibility = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *StorageStorageContainerHostMountStatus) GetHostName() string {
	if o == nil || IsNil(o.HostName) {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageContainerHostMountStatus) GetHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.HostName) {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *StorageStorageContainerHostMountStatus) HasHostName() bool {
	if o != nil && !IsNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *StorageStorageContainerHostMountStatus) SetHostName(v string) {
	o.HostName = &v
}

// GetMounted returns the Mounted field value if set, zero value otherwise.
func (o *StorageStorageContainerHostMountStatus) GetMounted() bool {
	if o == nil || IsNil(o.Mounted) {
		var ret bool
		return ret
	}
	return *o.Mounted
}

// GetMountedOk returns a tuple with the Mounted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageContainerHostMountStatus) GetMountedOk() (*bool, bool) {
	if o == nil || IsNil(o.Mounted) {
		return nil, false
	}
	return o.Mounted, true
}

// HasMounted returns a boolean if a field has been set.
func (o *StorageStorageContainerHostMountStatus) HasMounted() bool {
	if o != nil && !IsNil(o.Mounted) {
		return true
	}

	return false
}

// SetMounted gets a reference to the given bool and assigns it to the Mounted field.
func (o *StorageStorageContainerHostMountStatus) SetMounted(v bool) {
	o.Mounted = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *StorageStorageContainerHostMountStatus) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageContainerHostMountStatus) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *StorageStorageContainerHostMountStatus) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *StorageStorageContainerHostMountStatus) SetReason(v string) {
	o.Reason = &v
}

func (o StorageStorageContainerHostMountStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageStorageContainerHostMountStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Accessibility) {
		toSerialize["Accessibility"] = o.Accessibility
	}
	if !IsNil(o.HostName) {
		toSerialize["HostName"] = o.HostName
	}
	if !IsNil(o.Mounted) {
		toSerialize["Mounted"] = o.Mounted
	}
	if !IsNil(o.Reason) {
		toSerialize["Reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageStorageContainerHostMountStatus) UnmarshalJSON(data []byte) (err error) {
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
	type StorageStorageContainerHostMountStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Host specific storage container accessibility status. * `NOT_APPLICABLE` - The HyperFlex storage container accessibility on host is not applicable. * `ACCESSIBLE` - The HyperFlex storage container is accessible on the host. * `NOT_ACCESSIBLE` - The HyperFlex storage container is not accessible on the host. * `PARTIALLY_ACCESSIBLE` - The HyperFlex storage container is partially accessible on the host. * `READONLY` - The HyperFlex storage container accessibility is readonly on the host. * `HOST_NOT_REACHABLE` - The storage container is not accessible via this host because it is not accessible. * `UNKNOWN` - The storage container accessibility via this host is unknown.
		Accessibility *string `json:"Accessibility,omitempty"`
		// The name of the host corresponding to the mount and accessibility status of the storage container.
		HostName *string `json:"HostName,omitempty"`
		// Host specific storage container mount status.
		Mounted *bool `json:"Mounted,omitempty"`
		// Host specific storage container mount status reason.
		Reason *string `json:"Reason,omitempty"`
	}

	varStorageStorageContainerHostMountStatusWithoutEmbeddedStruct := StorageStorageContainerHostMountStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageStorageContainerHostMountStatusWithoutEmbeddedStruct)
	if err == nil {
		varStorageStorageContainerHostMountStatus := _StorageStorageContainerHostMountStatus{}
		varStorageStorageContainerHostMountStatus.ClassId = varStorageStorageContainerHostMountStatusWithoutEmbeddedStruct.ClassId
		varStorageStorageContainerHostMountStatus.ObjectType = varStorageStorageContainerHostMountStatusWithoutEmbeddedStruct.ObjectType
		varStorageStorageContainerHostMountStatus.Accessibility = varStorageStorageContainerHostMountStatusWithoutEmbeddedStruct.Accessibility
		varStorageStorageContainerHostMountStatus.HostName = varStorageStorageContainerHostMountStatusWithoutEmbeddedStruct.HostName
		varStorageStorageContainerHostMountStatus.Mounted = varStorageStorageContainerHostMountStatusWithoutEmbeddedStruct.Mounted
		varStorageStorageContainerHostMountStatus.Reason = varStorageStorageContainerHostMountStatusWithoutEmbeddedStruct.Reason
		*o = StorageStorageContainerHostMountStatus(varStorageStorageContainerHostMountStatus)
	} else {
		return err
	}

	varStorageStorageContainerHostMountStatus := _StorageStorageContainerHostMountStatus{}

	err = json.Unmarshal(data, &varStorageStorageContainerHostMountStatus)
	if err == nil {
		o.MoBaseComplexType = varStorageStorageContainerHostMountStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Accessibility")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "Mounted")
		delete(additionalProperties, "Reason")

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

type NullableStorageStorageContainerHostMountStatus struct {
	value *StorageStorageContainerHostMountStatus
	isSet bool
}

func (v NullableStorageStorageContainerHostMountStatus) Get() *StorageStorageContainerHostMountStatus {
	return v.value
}

func (v *NullableStorageStorageContainerHostMountStatus) Set(val *StorageStorageContainerHostMountStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageStorageContainerHostMountStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageStorageContainerHostMountStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageStorageContainerHostMountStatus(val *StorageStorageContainerHostMountStatus) *NullableStorageStorageContainerHostMountStatus {
	return &NullableStorageStorageContainerHostMountStatus{value: val, isSet: true}
}

func (v NullableStorageStorageContainerHostMountStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageStorageContainerHostMountStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
