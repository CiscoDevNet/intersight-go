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

// checks if the StorageVirtualDrivePolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageVirtualDrivePolicy{}

// StorageVirtualDrivePolicy This models the manual drive selection configuration.
type StorageVirtualDrivePolicy struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Access policy that host has on this virtual drive. * `Default` - Use platform default access mode. * `ReadWrite` - Enables host to perform read-write on the VD. * `ReadOnly` - Host can only read from the VD. * `Blocked` - Host can neither read nor write to the VD.
	AccessPolicy *string `json:"AccessPolicy,omitempty"`
	// Disk cache policy for the virtual drive. * `Default` - Use platform default drive cache mode. * `NoChange` - Drive cache policy is unchanged. * `Enable` - Enables IO caching on the drive. * `Disable` - Disables IO caching on the drive.
	DriveCache *string `json:"DriveCache,omitempty"`
	// Read ahead mode to be used to read data from this virtual drive. * `Default` - Use platform default read ahead mode. * `ReadAhead` - Use read ahead mode for the policy. * `NoReadAhead` - Do not use read ahead mode for the policy.
	ReadPolicy *string `json:"ReadPolicy,omitempty"`
	// Desired strip size - Allowed values are 64KiB, 128KiB, 256KiB, 512KiB, 1024KiB. * `64` - Number of bytes in a strip is 64 Kibibytes. * `128` - Number of bytes in a strip is 128 Kibibytes. * `256` - Number of bytes in a strip is 256 Kibibytes. * `512` - Number of bytes in a strip is 512 Kibibytes. * `1024` - Number of bytes in a strip is 1024 Kibibytes or 1 Mebibyte.
	StripSize *int32 `json:"StripSize,omitempty"`
	// Write mode to be used to write data to this virtual drive. * `Default` - Use platform default write mode. * `WriteThrough` - Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache. * `WriteBackGoodBbu` - Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure. * `AlwaysWriteBack` - With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged.
	WritePolicy          *string `json:"WritePolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageVirtualDrivePolicy StorageVirtualDrivePolicy

// NewStorageVirtualDrivePolicy instantiates a new StorageVirtualDrivePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDrivePolicy(classId string, objectType string) *StorageVirtualDrivePolicy {
	this := StorageVirtualDrivePolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var accessPolicy string = "Default"
	this.AccessPolicy = &accessPolicy
	var driveCache string = "Default"
	this.DriveCache = &driveCache
	var readPolicy string = "Default"
	this.ReadPolicy = &readPolicy
	var stripSize int32 = 64
	this.StripSize = &stripSize
	var writePolicy string = "Default"
	this.WritePolicy = &writePolicy
	return &this
}

// NewStorageVirtualDrivePolicyWithDefaults instantiates a new StorageVirtualDrivePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDrivePolicyWithDefaults() *StorageVirtualDrivePolicy {
	this := StorageVirtualDrivePolicy{}
	var classId string = "storage.VirtualDrivePolicy"
	this.ClassId = classId
	var objectType string = "storage.VirtualDrivePolicy"
	this.ObjectType = objectType
	var accessPolicy string = "Default"
	this.AccessPolicy = &accessPolicy
	var driveCache string = "Default"
	this.DriveCache = &driveCache
	var readPolicy string = "Default"
	this.ReadPolicy = &readPolicy
	var stripSize int32 = 64
	this.StripSize = &stripSize
	var writePolicy string = "Default"
	this.WritePolicy = &writePolicy
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVirtualDrivePolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVirtualDrivePolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.VirtualDrivePolicy" of the ClassId field.
func (o *StorageVirtualDrivePolicy) GetDefaultClassId() interface{} {
	return "storage.VirtualDrivePolicy"
}

// GetObjectType returns the ObjectType field value
func (o *StorageVirtualDrivePolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVirtualDrivePolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.VirtualDrivePolicy" of the ObjectType field.
func (o *StorageVirtualDrivePolicy) GetDefaultObjectType() interface{} {
	return "storage.VirtualDrivePolicy"
}

// GetAccessPolicy returns the AccessPolicy field value if set, zero value otherwise.
func (o *StorageVirtualDrivePolicy) GetAccessPolicy() string {
	if o == nil || IsNil(o.AccessPolicy) {
		var ret string
		return ret
	}
	return *o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetAccessPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessPolicy) {
		return nil, false
	}
	return o.AccessPolicy, true
}

// HasAccessPolicy returns a boolean if a field has been set.
func (o *StorageVirtualDrivePolicy) HasAccessPolicy() bool {
	if o != nil && !IsNil(o.AccessPolicy) {
		return true
	}

	return false
}

// SetAccessPolicy gets a reference to the given string and assigns it to the AccessPolicy field.
func (o *StorageVirtualDrivePolicy) SetAccessPolicy(v string) {
	o.AccessPolicy = &v
}

// GetDriveCache returns the DriveCache field value if set, zero value otherwise.
func (o *StorageVirtualDrivePolicy) GetDriveCache() string {
	if o == nil || IsNil(o.DriveCache) {
		var ret string
		return ret
	}
	return *o.DriveCache
}

// GetDriveCacheOk returns a tuple with the DriveCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetDriveCacheOk() (*string, bool) {
	if o == nil || IsNil(o.DriveCache) {
		return nil, false
	}
	return o.DriveCache, true
}

// HasDriveCache returns a boolean if a field has been set.
func (o *StorageVirtualDrivePolicy) HasDriveCache() bool {
	if o != nil && !IsNil(o.DriveCache) {
		return true
	}

	return false
}

// SetDriveCache gets a reference to the given string and assigns it to the DriveCache field.
func (o *StorageVirtualDrivePolicy) SetDriveCache(v string) {
	o.DriveCache = &v
}

// GetReadPolicy returns the ReadPolicy field value if set, zero value otherwise.
func (o *StorageVirtualDrivePolicy) GetReadPolicy() string {
	if o == nil || IsNil(o.ReadPolicy) {
		var ret string
		return ret
	}
	return *o.ReadPolicy
}

// GetReadPolicyOk returns a tuple with the ReadPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetReadPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.ReadPolicy) {
		return nil, false
	}
	return o.ReadPolicy, true
}

// HasReadPolicy returns a boolean if a field has been set.
func (o *StorageVirtualDrivePolicy) HasReadPolicy() bool {
	if o != nil && !IsNil(o.ReadPolicy) {
		return true
	}

	return false
}

// SetReadPolicy gets a reference to the given string and assigns it to the ReadPolicy field.
func (o *StorageVirtualDrivePolicy) SetReadPolicy(v string) {
	o.ReadPolicy = &v
}

// GetStripSize returns the StripSize field value if set, zero value otherwise.
func (o *StorageVirtualDrivePolicy) GetStripSize() int32 {
	if o == nil || IsNil(o.StripSize) {
		var ret int32
		return ret
	}
	return *o.StripSize
}

// GetStripSizeOk returns a tuple with the StripSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetStripSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.StripSize) {
		return nil, false
	}
	return o.StripSize, true
}

// HasStripSize returns a boolean if a field has been set.
func (o *StorageVirtualDrivePolicy) HasStripSize() bool {
	if o != nil && !IsNil(o.StripSize) {
		return true
	}

	return false
}

// SetStripSize gets a reference to the given int32 and assigns it to the StripSize field.
func (o *StorageVirtualDrivePolicy) SetStripSize(v int32) {
	o.StripSize = &v
}

// GetWritePolicy returns the WritePolicy field value if set, zero value otherwise.
func (o *StorageVirtualDrivePolicy) GetWritePolicy() string {
	if o == nil || IsNil(o.WritePolicy) {
		var ret string
		return ret
	}
	return *o.WritePolicy
}

// GetWritePolicyOk returns a tuple with the WritePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetWritePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.WritePolicy) {
		return nil, false
	}
	return o.WritePolicy, true
}

// HasWritePolicy returns a boolean if a field has been set.
func (o *StorageVirtualDrivePolicy) HasWritePolicy() bool {
	if o != nil && !IsNil(o.WritePolicy) {
		return true
	}

	return false
}

// SetWritePolicy gets a reference to the given string and assigns it to the WritePolicy field.
func (o *StorageVirtualDrivePolicy) SetWritePolicy(v string) {
	o.WritePolicy = &v
}

func (o StorageVirtualDrivePolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageVirtualDrivePolicy) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AccessPolicy) {
		toSerialize["AccessPolicy"] = o.AccessPolicy
	}
	if !IsNil(o.DriveCache) {
		toSerialize["DriveCache"] = o.DriveCache
	}
	if !IsNil(o.ReadPolicy) {
		toSerialize["ReadPolicy"] = o.ReadPolicy
	}
	if !IsNil(o.StripSize) {
		toSerialize["StripSize"] = o.StripSize
	}
	if !IsNil(o.WritePolicy) {
		toSerialize["WritePolicy"] = o.WritePolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageVirtualDrivePolicy) UnmarshalJSON(data []byte) (err error) {
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
	type StorageVirtualDrivePolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Access policy that host has on this virtual drive. * `Default` - Use platform default access mode. * `ReadWrite` - Enables host to perform read-write on the VD. * `ReadOnly` - Host can only read from the VD. * `Blocked` - Host can neither read nor write to the VD.
		AccessPolicy *string `json:"AccessPolicy,omitempty"`
		// Disk cache policy for the virtual drive. * `Default` - Use platform default drive cache mode. * `NoChange` - Drive cache policy is unchanged. * `Enable` - Enables IO caching on the drive. * `Disable` - Disables IO caching on the drive.
		DriveCache *string `json:"DriveCache,omitempty"`
		// Read ahead mode to be used to read data from this virtual drive. * `Default` - Use platform default read ahead mode. * `ReadAhead` - Use read ahead mode for the policy. * `NoReadAhead` - Do not use read ahead mode for the policy.
		ReadPolicy *string `json:"ReadPolicy,omitempty"`
		// Desired strip size - Allowed values are 64KiB, 128KiB, 256KiB, 512KiB, 1024KiB. * `64` - Number of bytes in a strip is 64 Kibibytes. * `128` - Number of bytes in a strip is 128 Kibibytes. * `256` - Number of bytes in a strip is 256 Kibibytes. * `512` - Number of bytes in a strip is 512 Kibibytes. * `1024` - Number of bytes in a strip is 1024 Kibibytes or 1 Mebibyte.
		StripSize *int32 `json:"StripSize,omitempty"`
		// Write mode to be used to write data to this virtual drive. * `Default` - Use platform default write mode. * `WriteThrough` - Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache. * `WriteBackGoodBbu` - Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure. * `AlwaysWriteBack` - With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged.
		WritePolicy *string `json:"WritePolicy,omitempty"`
	}

	varStorageVirtualDrivePolicyWithoutEmbeddedStruct := StorageVirtualDrivePolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageVirtualDrivePolicyWithoutEmbeddedStruct)
	if err == nil {
		varStorageVirtualDrivePolicy := _StorageVirtualDrivePolicy{}
		varStorageVirtualDrivePolicy.ClassId = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.ClassId
		varStorageVirtualDrivePolicy.ObjectType = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.ObjectType
		varStorageVirtualDrivePolicy.AccessPolicy = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.AccessPolicy
		varStorageVirtualDrivePolicy.DriveCache = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.DriveCache
		varStorageVirtualDrivePolicy.ReadPolicy = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.ReadPolicy
		varStorageVirtualDrivePolicy.StripSize = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.StripSize
		varStorageVirtualDrivePolicy.WritePolicy = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.WritePolicy
		*o = StorageVirtualDrivePolicy(varStorageVirtualDrivePolicy)
	} else {
		return err
	}

	varStorageVirtualDrivePolicy := _StorageVirtualDrivePolicy{}

	err = json.Unmarshal(data, &varStorageVirtualDrivePolicy)
	if err == nil {
		o.MoBaseComplexType = varStorageVirtualDrivePolicy.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessPolicy")
		delete(additionalProperties, "DriveCache")
		delete(additionalProperties, "ReadPolicy")
		delete(additionalProperties, "StripSize")
		delete(additionalProperties, "WritePolicy")

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

type NullableStorageVirtualDrivePolicy struct {
	value *StorageVirtualDrivePolicy
	isSet bool
}

func (v NullableStorageVirtualDrivePolicy) Get() *StorageVirtualDrivePolicy {
	return v.value
}

func (v *NullableStorageVirtualDrivePolicy) Set(val *StorageVirtualDrivePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDrivePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDrivePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDrivePolicy(val *StorageVirtualDrivePolicy) *NullableStorageVirtualDrivePolicy {
	return &NullableStorageVirtualDrivePolicy{value: val, isSet: true}
}

func (v NullableStorageVirtualDrivePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDrivePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
