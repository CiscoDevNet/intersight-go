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

// checks if the PartnerintegrationEtl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnerintegrationEtl{}

// PartnerintegrationEtl ETL definition for the endpoint to translate platform API outputs to Intersight managed objects.
type PartnerintegrationEtl struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Transformation model in yaml format.
	Data interface{} `json:"Data,omitempty"`
	// Placeholder name for the ETL.
	Name                 *string                                         `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.-]{1,64}$"`
	Inventory            NullablePartnerintegrationInventoryRelationship `json:"Inventory,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerintegrationEtl PartnerintegrationEtl

// NewPartnerintegrationEtl instantiates a new PartnerintegrationEtl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerintegrationEtl(classId string, objectType string) *PartnerintegrationEtl {
	this := PartnerintegrationEtl{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPartnerintegrationEtlWithDefaults instantiates a new PartnerintegrationEtl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerintegrationEtlWithDefaults() *PartnerintegrationEtl {
	this := PartnerintegrationEtl{}
	var classId string = "partnerintegration.Etl"
	this.ClassId = classId
	var objectType string = "partnerintegration.Etl"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PartnerintegrationEtl) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationEtl) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PartnerintegrationEtl) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "partnerintegration.Etl" of the ClassId field.
func (o *PartnerintegrationEtl) GetDefaultClassId() interface{} {
	return "partnerintegration.Etl"
}

// GetObjectType returns the ObjectType field value
func (o *PartnerintegrationEtl) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationEtl) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PartnerintegrationEtl) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "partnerintegration.Etl" of the ObjectType field.
func (o *PartnerintegrationEtl) GetDefaultObjectType() interface{} {
	return "partnerintegration.Etl"
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationEtl) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationEtl) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PartnerintegrationEtl) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *PartnerintegrationEtl) SetData(v interface{}) {
	o.Data = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartnerintegrationEtl) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationEtl) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartnerintegrationEtl) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartnerintegrationEtl) SetName(v string) {
	o.Name = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationEtl) GetInventory() PartnerintegrationInventoryRelationship {
	if o == nil || IsNil(o.Inventory.Get()) {
		var ret PartnerintegrationInventoryRelationship
		return ret
	}
	return *o.Inventory.Get()
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationEtl) GetInventoryOk() (*PartnerintegrationInventoryRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inventory.Get(), o.Inventory.IsSet()
}

// HasInventory returns a boolean if a field has been set.
func (o *PartnerintegrationEtl) HasInventory() bool {
	if o != nil && o.Inventory.IsSet() {
		return true
	}

	return false
}

// SetInventory gets a reference to the given NullablePartnerintegrationInventoryRelationship and assigns it to the Inventory field.
func (o *PartnerintegrationEtl) SetInventory(v PartnerintegrationInventoryRelationship) {
	o.Inventory.Set(&v)
}

// SetInventoryNil sets the value for Inventory to be an explicit nil
func (o *PartnerintegrationEtl) SetInventoryNil() {
	o.Inventory.Set(nil)
}

// UnsetInventory ensures that no value is present for Inventory, not even an explicit nil
func (o *PartnerintegrationEtl) UnsetInventory() {
	o.Inventory.Unset()
}

func (o PartnerintegrationEtl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnerintegrationEtl) ToMap() (map[string]interface{}, error) {
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
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.Inventory.IsSet() {
		toSerialize["Inventory"] = o.Inventory.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PartnerintegrationEtl) UnmarshalJSON(data []byte) (err error) {
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
	type PartnerintegrationEtlWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Transformation model in yaml format.
		Data interface{} `json:"Data,omitempty"`
		// Placeholder name for the ETL.
		Name      *string                                         `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.-]{1,64}$"`
		Inventory NullablePartnerintegrationInventoryRelationship `json:"Inventory,omitempty"`
	}

	varPartnerintegrationEtlWithoutEmbeddedStruct := PartnerintegrationEtlWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPartnerintegrationEtlWithoutEmbeddedStruct)
	if err == nil {
		varPartnerintegrationEtl := _PartnerintegrationEtl{}
		varPartnerintegrationEtl.ClassId = varPartnerintegrationEtlWithoutEmbeddedStruct.ClassId
		varPartnerintegrationEtl.ObjectType = varPartnerintegrationEtlWithoutEmbeddedStruct.ObjectType
		varPartnerintegrationEtl.Data = varPartnerintegrationEtlWithoutEmbeddedStruct.Data
		varPartnerintegrationEtl.Name = varPartnerintegrationEtlWithoutEmbeddedStruct.Name
		varPartnerintegrationEtl.Inventory = varPartnerintegrationEtlWithoutEmbeddedStruct.Inventory
		*o = PartnerintegrationEtl(varPartnerintegrationEtl)
	} else {
		return err
	}

	varPartnerintegrationEtl := _PartnerintegrationEtl{}

	err = json.Unmarshal(data, &varPartnerintegrationEtl)
	if err == nil {
		o.MoBaseMo = varPartnerintegrationEtl.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Data")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Inventory")

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

type NullablePartnerintegrationEtl struct {
	value *PartnerintegrationEtl
	isSet bool
}

func (v NullablePartnerintegrationEtl) Get() *PartnerintegrationEtl {
	return v.value
}

func (v *NullablePartnerintegrationEtl) Set(val *PartnerintegrationEtl) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerintegrationEtl) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerintegrationEtl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerintegrationEtl(val *PartnerintegrationEtl) *NullablePartnerintegrationEtl {
	return &NullablePartnerintegrationEtl{value: val, isSet: true}
}

func (v NullablePartnerintegrationEtl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerintegrationEtl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
