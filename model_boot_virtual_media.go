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

// checks if the BootVirtualMedia type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BootVirtualMedia{}

// BootVirtualMedia Device type used when booting from Virtual Media device.
type BootVirtualMedia struct {
	BootDeviceBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The subtype for the selected device type. * `None` - No sub type for virtual media. * `cimc-mapped-dvd` - The virtual media device is mapped to a virtual DVD device. * `cimc-mapped-hdd` - The virtual media device is mapped to a virtual HDD device. * `kvm-mapped-dvd` - A KVM mapped DVD virtual media device. * `kvm-mapped-hdd` - A KVM mapped HDD virtual media device. * `kvm-mapped-fdd` - A KVM mapped FDD virtual media device.
	Subtype              *string `json:"Subtype,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootVirtualMedia BootVirtualMedia

// NewBootVirtualMedia instantiates a new BootVirtualMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootVirtualMedia(classId string, objectType string) *BootVirtualMedia {
	this := BootVirtualMedia{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	var subtype string = "None"
	this.Subtype = &subtype
	return &this
}

// NewBootVirtualMediaWithDefaults instantiates a new BootVirtualMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootVirtualMediaWithDefaults() *BootVirtualMedia {
	this := BootVirtualMedia{}
	var classId string = "boot.VirtualMedia"
	this.ClassId = classId
	var objectType string = "boot.VirtualMedia"
	this.ObjectType = objectType
	var subtype string = "None"
	this.Subtype = &subtype
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootVirtualMedia) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootVirtualMedia) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootVirtualMedia) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "boot.VirtualMedia" of the ClassId field.
func (o *BootVirtualMedia) GetDefaultClassId() interface{} {
	return "boot.VirtualMedia"
}

// GetObjectType returns the ObjectType field value
func (o *BootVirtualMedia) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootVirtualMedia) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootVirtualMedia) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "boot.VirtualMedia" of the ObjectType field.
func (o *BootVirtualMedia) GetDefaultObjectType() interface{} {
	return "boot.VirtualMedia"
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *BootVirtualMedia) GetSubtype() string {
	if o == nil || IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVirtualMedia) GetSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *BootVirtualMedia) HasSubtype() bool {
	if o != nil && !IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *BootVirtualMedia) SetSubtype(v string) {
	o.Subtype = &v
}

func (o BootVirtualMedia) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BootVirtualMedia) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBootDeviceBase, errBootDeviceBase := json.Marshal(o.BootDeviceBase)
	if errBootDeviceBase != nil {
		return map[string]interface{}{}, errBootDeviceBase
	}
	errBootDeviceBase = json.Unmarshal([]byte(serializedBootDeviceBase), &toSerialize)
	if errBootDeviceBase != nil {
		return map[string]interface{}{}, errBootDeviceBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Subtype) {
		toSerialize["Subtype"] = o.Subtype
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BootVirtualMedia) UnmarshalJSON(data []byte) (err error) {
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
	type BootVirtualMediaWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The subtype for the selected device type. * `None` - No sub type for virtual media. * `cimc-mapped-dvd` - The virtual media device is mapped to a virtual DVD device. * `cimc-mapped-hdd` - The virtual media device is mapped to a virtual HDD device. * `kvm-mapped-dvd` - A KVM mapped DVD virtual media device. * `kvm-mapped-hdd` - A KVM mapped HDD virtual media device. * `kvm-mapped-fdd` - A KVM mapped FDD virtual media device.
		Subtype *string `json:"Subtype,omitempty"`
	}

	varBootVirtualMediaWithoutEmbeddedStruct := BootVirtualMediaWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBootVirtualMediaWithoutEmbeddedStruct)
	if err == nil {
		varBootVirtualMedia := _BootVirtualMedia{}
		varBootVirtualMedia.ClassId = varBootVirtualMediaWithoutEmbeddedStruct.ClassId
		varBootVirtualMedia.ObjectType = varBootVirtualMediaWithoutEmbeddedStruct.ObjectType
		varBootVirtualMedia.Subtype = varBootVirtualMediaWithoutEmbeddedStruct.Subtype
		*o = BootVirtualMedia(varBootVirtualMedia)
	} else {
		return err
	}

	varBootVirtualMedia := _BootVirtualMedia{}

	err = json.Unmarshal(data, &varBootVirtualMedia)
	if err == nil {
		o.BootDeviceBase = varBootVirtualMedia.BootDeviceBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Subtype")

		// remove fields from embedded structs
		reflectBootDeviceBase := reflect.ValueOf(o.BootDeviceBase)
		for i := 0; i < reflectBootDeviceBase.Type().NumField(); i++ {
			t := reflectBootDeviceBase.Type().Field(i)

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

type NullableBootVirtualMedia struct {
	value *BootVirtualMedia
	isSet bool
}

func (v NullableBootVirtualMedia) Get() *BootVirtualMedia {
	return v.value
}

func (v *NullableBootVirtualMedia) Set(val *BootVirtualMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableBootVirtualMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableBootVirtualMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootVirtualMedia(val *BootVirtualMedia) *NullableBootVirtualMedia {
	return &NullableBootVirtualMedia{value: val, isSet: true}
}

func (v NullableBootVirtualMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootVirtualMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
