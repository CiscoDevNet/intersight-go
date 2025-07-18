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
	"time"
)

// checks if the HciEntitlementDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HciEntitlementDetail{}

// HciEntitlementDetail The detail of a license entitlement detail.
type HciEntitlementDetail struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The category of the entitlement.
	Category *string `json:"Category,omitempty"`
	// The earliest expiry date of the entitlement.
	EarliestExpiryDate *time.Time `json:"EarliestExpiryDate,omitempty"`
	// The meter of the entitlement.
	Meter *string `json:"Meter,omitempty"`
	// The name of the entitlement.
	Name *string `json:"Name,omitempty"`
	// The quantity of the entitlement.
	Quantity *float64 `json:"Quantity,omitempty"`
	// The scope of the entitlement.
	Scope *string `json:"Scope,omitempty"`
	// The subCategory of the entitlement.
	SubCategory *string `json:"SubCategory,omitempty"`
	// The type of the license entitlement.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HciEntitlementDetail HciEntitlementDetail

// NewHciEntitlementDetail instantiates a new HciEntitlementDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHciEntitlementDetail(classId string, objectType string) *HciEntitlementDetail {
	this := HciEntitlementDetail{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHciEntitlementDetailWithDefaults instantiates a new HciEntitlementDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHciEntitlementDetailWithDefaults() *HciEntitlementDetail {
	this := HciEntitlementDetail{}
	var classId string = "hci.EntitlementDetail"
	this.ClassId = classId
	var objectType string = "hci.EntitlementDetail"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HciEntitlementDetail) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HciEntitlementDetail) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HciEntitlementDetail) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hci.EntitlementDetail" of the ClassId field.
func (o *HciEntitlementDetail) GetDefaultClassId() interface{} {
	return "hci.EntitlementDetail"
}

// GetObjectType returns the ObjectType field value
func (o *HciEntitlementDetail) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HciEntitlementDetail) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HciEntitlementDetail) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hci.EntitlementDetail" of the ObjectType field.
func (o *HciEntitlementDetail) GetDefaultObjectType() interface{} {
	return "hci.EntitlementDetail"
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *HciEntitlementDetail) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEntitlementDetail) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *HciEntitlementDetail) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *HciEntitlementDetail) SetCategory(v string) {
	o.Category = &v
}

// GetEarliestExpiryDate returns the EarliestExpiryDate field value if set, zero value otherwise.
func (o *HciEntitlementDetail) GetEarliestExpiryDate() time.Time {
	if o == nil || IsNil(o.EarliestExpiryDate) {
		var ret time.Time
		return ret
	}
	return *o.EarliestExpiryDate
}

// GetEarliestExpiryDateOk returns a tuple with the EarliestExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEntitlementDetail) GetEarliestExpiryDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EarliestExpiryDate) {
		return nil, false
	}
	return o.EarliestExpiryDate, true
}

// HasEarliestExpiryDate returns a boolean if a field has been set.
func (o *HciEntitlementDetail) HasEarliestExpiryDate() bool {
	if o != nil && !IsNil(o.EarliestExpiryDate) {
		return true
	}

	return false
}

// SetEarliestExpiryDate gets a reference to the given time.Time and assigns it to the EarliestExpiryDate field.
func (o *HciEntitlementDetail) SetEarliestExpiryDate(v time.Time) {
	o.EarliestExpiryDate = &v
}

// GetMeter returns the Meter field value if set, zero value otherwise.
func (o *HciEntitlementDetail) GetMeter() string {
	if o == nil || IsNil(o.Meter) {
		var ret string
		return ret
	}
	return *o.Meter
}

// GetMeterOk returns a tuple with the Meter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEntitlementDetail) GetMeterOk() (*string, bool) {
	if o == nil || IsNil(o.Meter) {
		return nil, false
	}
	return o.Meter, true
}

// HasMeter returns a boolean if a field has been set.
func (o *HciEntitlementDetail) HasMeter() bool {
	if o != nil && !IsNil(o.Meter) {
		return true
	}

	return false
}

// SetMeter gets a reference to the given string and assigns it to the Meter field.
func (o *HciEntitlementDetail) SetMeter(v string) {
	o.Meter = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HciEntitlementDetail) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEntitlementDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HciEntitlementDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HciEntitlementDetail) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *HciEntitlementDetail) GetQuantity() float64 {
	if o == nil || IsNil(o.Quantity) {
		var ret float64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEntitlementDetail) GetQuantityOk() (*float64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *HciEntitlementDetail) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float64 and assigns it to the Quantity field.
func (o *HciEntitlementDetail) SetQuantity(v float64) {
	o.Quantity = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *HciEntitlementDetail) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEntitlementDetail) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *HciEntitlementDetail) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *HciEntitlementDetail) SetScope(v string) {
	o.Scope = &v
}

// GetSubCategory returns the SubCategory field value if set, zero value otherwise.
func (o *HciEntitlementDetail) GetSubCategory() string {
	if o == nil || IsNil(o.SubCategory) {
		var ret string
		return ret
	}
	return *o.SubCategory
}

// GetSubCategoryOk returns a tuple with the SubCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEntitlementDetail) GetSubCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.SubCategory) {
		return nil, false
	}
	return o.SubCategory, true
}

// HasSubCategory returns a boolean if a field has been set.
func (o *HciEntitlementDetail) HasSubCategory() bool {
	if o != nil && !IsNil(o.SubCategory) {
		return true
	}

	return false
}

// SetSubCategory gets a reference to the given string and assigns it to the SubCategory field.
func (o *HciEntitlementDetail) SetSubCategory(v string) {
	o.SubCategory = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HciEntitlementDetail) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HciEntitlementDetail) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HciEntitlementDetail) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HciEntitlementDetail) SetType(v string) {
	o.Type = &v
}

func (o HciEntitlementDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HciEntitlementDetail) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Category) {
		toSerialize["Category"] = o.Category
	}
	if !IsNil(o.EarliestExpiryDate) {
		toSerialize["EarliestExpiryDate"] = o.EarliestExpiryDate
	}
	if !IsNil(o.Meter) {
		toSerialize["Meter"] = o.Meter
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Quantity) {
		toSerialize["Quantity"] = o.Quantity
	}
	if !IsNil(o.Scope) {
		toSerialize["Scope"] = o.Scope
	}
	if !IsNil(o.SubCategory) {
		toSerialize["SubCategory"] = o.SubCategory
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HciEntitlementDetail) UnmarshalJSON(data []byte) (err error) {
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
	type HciEntitlementDetailWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The category of the entitlement.
		Category *string `json:"Category,omitempty"`
		// The earliest expiry date of the entitlement.
		EarliestExpiryDate *time.Time `json:"EarliestExpiryDate,omitempty"`
		// The meter of the entitlement.
		Meter *string `json:"Meter,omitempty"`
		// The name of the entitlement.
		Name *string `json:"Name,omitempty"`
		// The quantity of the entitlement.
		Quantity *float64 `json:"Quantity,omitempty"`
		// The scope of the entitlement.
		Scope *string `json:"Scope,omitempty"`
		// The subCategory of the entitlement.
		SubCategory *string `json:"SubCategory,omitempty"`
		// The type of the license entitlement.
		Type *string `json:"Type,omitempty"`
	}

	varHciEntitlementDetailWithoutEmbeddedStruct := HciEntitlementDetailWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHciEntitlementDetailWithoutEmbeddedStruct)
	if err == nil {
		varHciEntitlementDetail := _HciEntitlementDetail{}
		varHciEntitlementDetail.ClassId = varHciEntitlementDetailWithoutEmbeddedStruct.ClassId
		varHciEntitlementDetail.ObjectType = varHciEntitlementDetailWithoutEmbeddedStruct.ObjectType
		varHciEntitlementDetail.Category = varHciEntitlementDetailWithoutEmbeddedStruct.Category
		varHciEntitlementDetail.EarliestExpiryDate = varHciEntitlementDetailWithoutEmbeddedStruct.EarliestExpiryDate
		varHciEntitlementDetail.Meter = varHciEntitlementDetailWithoutEmbeddedStruct.Meter
		varHciEntitlementDetail.Name = varHciEntitlementDetailWithoutEmbeddedStruct.Name
		varHciEntitlementDetail.Quantity = varHciEntitlementDetailWithoutEmbeddedStruct.Quantity
		varHciEntitlementDetail.Scope = varHciEntitlementDetailWithoutEmbeddedStruct.Scope
		varHciEntitlementDetail.SubCategory = varHciEntitlementDetailWithoutEmbeddedStruct.SubCategory
		varHciEntitlementDetail.Type = varHciEntitlementDetailWithoutEmbeddedStruct.Type
		*o = HciEntitlementDetail(varHciEntitlementDetail)
	} else {
		return err
	}

	varHciEntitlementDetail := _HciEntitlementDetail{}

	err = json.Unmarshal(data, &varHciEntitlementDetail)
	if err == nil {
		o.MoBaseComplexType = varHciEntitlementDetail.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "EarliestExpiryDate")
		delete(additionalProperties, "Meter")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Quantity")
		delete(additionalProperties, "Scope")
		delete(additionalProperties, "SubCategory")
		delete(additionalProperties, "Type")

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

type NullableHciEntitlementDetail struct {
	value *HciEntitlementDetail
	isSet bool
}

func (v NullableHciEntitlementDetail) Get() *HciEntitlementDetail {
	return v.value
}

func (v *NullableHciEntitlementDetail) Set(val *HciEntitlementDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableHciEntitlementDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableHciEntitlementDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHciEntitlementDetail(val *HciEntitlementDetail) *NullableHciEntitlementDetail {
	return &NullableHciEntitlementDetail{value: val, isSet: true}
}

func (v NullableHciEntitlementDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHciEntitlementDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
