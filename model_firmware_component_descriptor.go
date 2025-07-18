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

// checks if the FirmwareComponentDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareComponentDescriptor{}

// FirmwareComponentDescriptor Descriptor to uniquely identify each component.
type FirmwareComponentDescriptor struct {
	CapabilityHardwareDescriptor
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The brand string of the endpoint for which this capability information is applicable.
	BrandString *string `json:"BrandString,omitempty"`
	// The label type for the component.
	Label *string `json:"Label,omitempty"`
	// The revision for the component.
	Revision             *string `json:"Revision,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareComponentDescriptor FirmwareComponentDescriptor

// NewFirmwareComponentDescriptor instantiates a new FirmwareComponentDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareComponentDescriptor(classId string, objectType string) *FirmwareComponentDescriptor {
	this := FirmwareComponentDescriptor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareComponentDescriptorWithDefaults instantiates a new FirmwareComponentDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareComponentDescriptorWithDefaults() *FirmwareComponentDescriptor {
	this := FirmwareComponentDescriptor{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareComponentDescriptor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareComponentDescriptor) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareComponentDescriptor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareComponentDescriptor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBrandString returns the BrandString field value if set, zero value otherwise.
func (o *FirmwareComponentDescriptor) GetBrandString() string {
	if o == nil || IsNil(o.BrandString) {
		var ret string
		return ret
	}
	return *o.BrandString
}

// GetBrandStringOk returns a tuple with the BrandString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptor) GetBrandStringOk() (*string, bool) {
	if o == nil || IsNil(o.BrandString) {
		return nil, false
	}
	return o.BrandString, true
}

// HasBrandString returns a boolean if a field has been set.
func (o *FirmwareComponentDescriptor) HasBrandString() bool {
	if o != nil && !IsNil(o.BrandString) {
		return true
	}

	return false
}

// SetBrandString gets a reference to the given string and assigns it to the BrandString field.
func (o *FirmwareComponentDescriptor) SetBrandString(v string) {
	o.BrandString = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FirmwareComponentDescriptor) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptor) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FirmwareComponentDescriptor) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FirmwareComponentDescriptor) SetLabel(v string) {
	o.Label = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *FirmwareComponentDescriptor) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptor) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *FirmwareComponentDescriptor) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *FirmwareComponentDescriptor) SetRevision(v string) {
	o.Revision = &v
}

func (o FirmwareComponentDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareComponentDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityHardwareDescriptor, errCapabilityHardwareDescriptor := json.Marshal(o.CapabilityHardwareDescriptor)
	if errCapabilityHardwareDescriptor != nil {
		return map[string]interface{}{}, errCapabilityHardwareDescriptor
	}
	errCapabilityHardwareDescriptor = json.Unmarshal([]byte(serializedCapabilityHardwareDescriptor), &toSerialize)
	if errCapabilityHardwareDescriptor != nil {
		return map[string]interface{}{}, errCapabilityHardwareDescriptor
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.BrandString) {
		toSerialize["BrandString"] = o.BrandString
	}
	if !IsNil(o.Label) {
		toSerialize["Label"] = o.Label
	}
	if !IsNil(o.Revision) {
		toSerialize["Revision"] = o.Revision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirmwareComponentDescriptor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type FirmwareComponentDescriptorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The brand string of the endpoint for which this capability information is applicable.
		BrandString *string `json:"BrandString,omitempty"`
		// The label type for the component.
		Label *string `json:"Label,omitempty"`
		// The revision for the component.
		Revision *string `json:"Revision,omitempty"`
	}

	varFirmwareComponentDescriptorWithoutEmbeddedStruct := FirmwareComponentDescriptorWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFirmwareComponentDescriptorWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareComponentDescriptor := _FirmwareComponentDescriptor{}
		varFirmwareComponentDescriptor.ClassId = varFirmwareComponentDescriptorWithoutEmbeddedStruct.ClassId
		varFirmwareComponentDescriptor.ObjectType = varFirmwareComponentDescriptorWithoutEmbeddedStruct.ObjectType
		varFirmwareComponentDescriptor.BrandString = varFirmwareComponentDescriptorWithoutEmbeddedStruct.BrandString
		varFirmwareComponentDescriptor.Label = varFirmwareComponentDescriptorWithoutEmbeddedStruct.Label
		varFirmwareComponentDescriptor.Revision = varFirmwareComponentDescriptorWithoutEmbeddedStruct.Revision
		*o = FirmwareComponentDescriptor(varFirmwareComponentDescriptor)
	} else {
		return err
	}

	varFirmwareComponentDescriptor := _FirmwareComponentDescriptor{}

	err = json.Unmarshal(data, &varFirmwareComponentDescriptor)
	if err == nil {
		o.CapabilityHardwareDescriptor = varFirmwareComponentDescriptor.CapabilityHardwareDescriptor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BrandString")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Revision")

		// remove fields from embedded structs
		reflectCapabilityHardwareDescriptor := reflect.ValueOf(o.CapabilityHardwareDescriptor)
		for i := 0; i < reflectCapabilityHardwareDescriptor.Type().NumField(); i++ {
			t := reflectCapabilityHardwareDescriptor.Type().Field(i)

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

type NullableFirmwareComponentDescriptor struct {
	value *FirmwareComponentDescriptor
	isSet bool
}

func (v NullableFirmwareComponentDescriptor) Get() *FirmwareComponentDescriptor {
	return v.value
}

func (v *NullableFirmwareComponentDescriptor) Set(val *FirmwareComponentDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareComponentDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareComponentDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareComponentDescriptor(val *FirmwareComponentDescriptor) *NullableFirmwareComponentDescriptor {
	return &NullableFirmwareComponentDescriptor{value: val, isSet: true}
}

func (v NullableFirmwareComponentDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareComponentDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
