/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-15830
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

// checks if the SoftwareHciDistributable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwareHciDistributable{}

// SoftwareHciDistributable An HCI image distributed by Cisco.
type SoftwareHciDistributable struct {
	FirmwareBaseDistributable
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Additional information associated with the distributable object. This data is provided as a JSON blob by Nutanix, a partner vendor, and the format is not fixed.
	MetaInfo             interface{}                                   `json:"MetaInfo,omitempty"`
	Catalog              NullableSoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareHciDistributable SoftwareHciDistributable

// NewSoftwareHciDistributable instantiates a new SoftwareHciDistributable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareHciDistributable(classId string, objectType string) *SoftwareHciDistributable {
	this := SoftwareHciDistributable{}
	this.ClassId = classId
	this.ObjectType = objectType
	var importAction string = "None"
	this.ImportAction = &importAction
	var vendor string = "Cisco"
	this.Vendor = &vendor
	return &this
}

// NewSoftwareHciDistributableWithDefaults instantiates a new SoftwareHciDistributable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareHciDistributableWithDefaults() *SoftwareHciDistributable {
	this := SoftwareHciDistributable{}
	var classId string = "software.HciDistributable"
	this.ClassId = classId
	var objectType string = "software.HciDistributable"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareHciDistributable) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareHciDistributable) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareHciDistributable) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareHciDistributable) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareHciDistributable) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareHciDistributable) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMetaInfo returns the MetaInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwareHciDistributable) GetMetaInfo() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MetaInfo
}

// GetMetaInfoOk returns a tuple with the MetaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwareHciDistributable) GetMetaInfoOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MetaInfo) {
		return nil, false
	}
	return &o.MetaInfo, true
}

// HasMetaInfo returns a boolean if a field has been set.
func (o *SoftwareHciDistributable) HasMetaInfo() bool {
	if o != nil && IsNil(o.MetaInfo) {
		return true
	}

	return false
}

// SetMetaInfo gets a reference to the given interface{} and assigns it to the MetaInfo field.
func (o *SoftwareHciDistributable) SetMetaInfo(v interface{}) {
	o.MetaInfo = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwareHciDistributable) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || IsNil(o.Catalog.Get()) {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog.Get()
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwareHciDistributable) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Catalog.Get(), o.Catalog.IsSet()
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwareHciDistributable) HasCatalog() bool {
	if o != nil && o.Catalog.IsSet() {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given NullableSoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwareHciDistributable) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog.Set(&v)
}

// SetCatalogNil sets the value for Catalog to be an explicit nil
func (o *SoftwareHciDistributable) SetCatalogNil() {
	o.Catalog.Set(nil)
}

// UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
func (o *SoftwareHciDistributable) UnsetCatalog() {
	o.Catalog.Unset()
}

func (o SoftwareHciDistributable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwareHciDistributable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareBaseDistributable, errFirmwareBaseDistributable := json.Marshal(o.FirmwareBaseDistributable)
	if errFirmwareBaseDistributable != nil {
		return map[string]interface{}{}, errFirmwareBaseDistributable
	}
	errFirmwareBaseDistributable = json.Unmarshal([]byte(serializedFirmwareBaseDistributable), &toSerialize)
	if errFirmwareBaseDistributable != nil {
		return map[string]interface{}{}, errFirmwareBaseDistributable
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if o.MetaInfo != nil {
		toSerialize["MetaInfo"] = o.MetaInfo
	}
	if o.Catalog.IsSet() {
		toSerialize["Catalog"] = o.Catalog.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SoftwareHciDistributable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type SoftwareHciDistributableWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Additional information associated with the distributable object. This data is provided as a JSON blob by Nutanix, a partner vendor, and the format is not fixed.
		MetaInfo interface{}                                   `json:"MetaInfo,omitempty"`
		Catalog  NullableSoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	}

	varSoftwareHciDistributableWithoutEmbeddedStruct := SoftwareHciDistributableWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSoftwareHciDistributableWithoutEmbeddedStruct)
	if err == nil {
		varSoftwareHciDistributable := _SoftwareHciDistributable{}
		varSoftwareHciDistributable.ClassId = varSoftwareHciDistributableWithoutEmbeddedStruct.ClassId
		varSoftwareHciDistributable.ObjectType = varSoftwareHciDistributableWithoutEmbeddedStruct.ObjectType
		varSoftwareHciDistributable.MetaInfo = varSoftwareHciDistributableWithoutEmbeddedStruct.MetaInfo
		varSoftwareHciDistributable.Catalog = varSoftwareHciDistributableWithoutEmbeddedStruct.Catalog
		*o = SoftwareHciDistributable(varSoftwareHciDistributable)
	} else {
		return err
	}

	varSoftwareHciDistributable := _SoftwareHciDistributable{}

	err = json.Unmarshal(data, &varSoftwareHciDistributable)
	if err == nil {
		o.FirmwareBaseDistributable = varSoftwareHciDistributable.FirmwareBaseDistributable
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MetaInfo")
		delete(additionalProperties, "Catalog")

		// remove fields from embedded structs
		reflectFirmwareBaseDistributable := reflect.ValueOf(o.FirmwareBaseDistributable)
		for i := 0; i < reflectFirmwareBaseDistributable.Type().NumField(); i++ {
			t := reflectFirmwareBaseDistributable.Type().Field(i)

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

type NullableSoftwareHciDistributable struct {
	value *SoftwareHciDistributable
	isSet bool
}

func (v NullableSoftwareHciDistributable) Get() *SoftwareHciDistributable {
	return v.value
}

func (v *NullableSoftwareHciDistributable) Set(val *SoftwareHciDistributable) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareHciDistributable) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareHciDistributable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareHciDistributable(val *SoftwareHciDistributable) *NullableSoftwareHciDistributable {
	return &NullableSoftwareHciDistributable{value: val, isSet: true}
}

func (v NullableSoftwareHciDistributable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareHciDistributable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}