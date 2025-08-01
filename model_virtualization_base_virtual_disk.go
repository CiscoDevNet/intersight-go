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

// checks if the VirtualizationBaseVirtualDisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationBaseVirtualDisk{}

// VirtualizationBaseVirtualDisk The Virtual disk created on a Hypervisor Datastore.
type VirtualizationBaseVirtualDisk struct {
	VirtualizationBaseSourceDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Name of the storage disk. Name must be unique within a Datastore.
	Name *string `json:"Name,omitempty"`
	// Disk size represented in bytes.
	Size                 *int64 `json:"Size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseVirtualDisk VirtualizationBaseVirtualDisk

// NewVirtualizationBaseVirtualDisk instantiates a new VirtualizationBaseVirtualDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseVirtualDisk(classId string, objectType string) *VirtualizationBaseVirtualDisk {
	this := VirtualizationBaseVirtualDisk{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationBaseVirtualDiskWithDefaults instantiates a new VirtualizationBaseVirtualDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseVirtualDiskWithDefaults() *VirtualizationBaseVirtualDisk {
	this := VirtualizationBaseVirtualDisk{}
	var classId string = "virtualization.VmwareVirtualDisk"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVirtualDisk"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseVirtualDisk) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualDisk) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseVirtualDisk) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "virtualization.VmwareVirtualDisk" of the ClassId field.
func (o *VirtualizationBaseVirtualDisk) GetDefaultClassId() interface{} {
	return "virtualization.VmwareVirtualDisk"
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseVirtualDisk) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualDisk) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseVirtualDisk) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "virtualization.VmwareVirtualDisk" of the ObjectType field.
func (o *VirtualizationBaseVirtualDisk) GetDefaultObjectType() interface{} {
	return "virtualization.VmwareVirtualDisk"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualDisk) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualDisk) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualDisk) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseVirtualDisk) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualDisk) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualDisk) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualDisk) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *VirtualizationBaseVirtualDisk) SetSize(v int64) {
	o.Size = &v
}

func (o VirtualizationBaseVirtualDisk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationBaseVirtualDisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseSourceDevice, errVirtualizationBaseSourceDevice := json.Marshal(o.VirtualizationBaseSourceDevice)
	if errVirtualizationBaseSourceDevice != nil {
		return map[string]interface{}{}, errVirtualizationBaseSourceDevice
	}
	errVirtualizationBaseSourceDevice = json.Unmarshal([]byte(serializedVirtualizationBaseSourceDevice), &toSerialize)
	if errVirtualizationBaseSourceDevice != nil {
		return map[string]interface{}{}, errVirtualizationBaseSourceDevice
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
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualizationBaseVirtualDisk) UnmarshalJSON(data []byte) (err error) {
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
	type VirtualizationBaseVirtualDiskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Name of the storage disk. Name must be unique within a Datastore.
		Name *string `json:"Name,omitempty"`
		// Disk size represented in bytes.
		Size *int64 `json:"Size,omitempty"`
	}

	varVirtualizationBaseVirtualDiskWithoutEmbeddedStruct := VirtualizationBaseVirtualDiskWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVirtualizationBaseVirtualDiskWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationBaseVirtualDisk := _VirtualizationBaseVirtualDisk{}
		varVirtualizationBaseVirtualDisk.ClassId = varVirtualizationBaseVirtualDiskWithoutEmbeddedStruct.ClassId
		varVirtualizationBaseVirtualDisk.ObjectType = varVirtualizationBaseVirtualDiskWithoutEmbeddedStruct.ObjectType
		varVirtualizationBaseVirtualDisk.Name = varVirtualizationBaseVirtualDiskWithoutEmbeddedStruct.Name
		varVirtualizationBaseVirtualDisk.Size = varVirtualizationBaseVirtualDiskWithoutEmbeddedStruct.Size
		*o = VirtualizationBaseVirtualDisk(varVirtualizationBaseVirtualDisk)
	} else {
		return err
	}

	varVirtualizationBaseVirtualDisk := _VirtualizationBaseVirtualDisk{}

	err = json.Unmarshal(data, &varVirtualizationBaseVirtualDisk)
	if err == nil {
		o.VirtualizationBaseSourceDevice = varVirtualizationBaseVirtualDisk.VirtualizationBaseSourceDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Size")

		// remove fields from embedded structs
		reflectVirtualizationBaseSourceDevice := reflect.ValueOf(o.VirtualizationBaseSourceDevice)
		for i := 0; i < reflectVirtualizationBaseSourceDevice.Type().NumField(); i++ {
			t := reflectVirtualizationBaseSourceDevice.Type().Field(i)

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

type NullableVirtualizationBaseVirtualDisk struct {
	value *VirtualizationBaseVirtualDisk
	isSet bool
}

func (v NullableVirtualizationBaseVirtualDisk) Get() *VirtualizationBaseVirtualDisk {
	return v.value
}

func (v *NullableVirtualizationBaseVirtualDisk) Set(val *VirtualizationBaseVirtualDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseVirtualDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseVirtualDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseVirtualDisk(val *VirtualizationBaseVirtualDisk) *NullableVirtualizationBaseVirtualDisk {
	return &NullableVirtualizationBaseVirtualDisk{value: val, isSet: true}
}

func (v NullableVirtualizationBaseVirtualDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseVirtualDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
