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

// checks if the RecommendationHardwareExpansionRequestItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommendationHardwareExpansionRequestItem{}

// RecommendationHardwareExpansionRequestItem Entity representing the user request for expansion of each hardware item.
type RecommendationHardwareExpansionRequestItem struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of hardware item for which the capacity increase is requested by the user. For example, CPU. * `None` - The Enum value None represents that no value was set for the hardware type. * `CPU` - The Enum value CPU represents that the hardware type requested for expansion is a physical CPU. * `Memory` - The Enum value Memory represents that the hardware type requested for expansion is a memory unit. * `Storage` - The Enum value Storage represents that the hardware type requested for expansion is a storage unit, ie, storage drives.
	ItemType *string `json:"ItemType,omitempty"`
	// The maximum value allowed for expansion for the hardware type on the referred registered device.
	MaxValue *float32 `json:"MaxValue,omitempty"`
	// Unit type for the maximum value of the resource. For example, TB, GB, MB. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes. * `GB` - The Enum value GB represents that the measurement unit is in gigabytes. * `MHz` - The Enum value MHz represents that the measurement unit is in megahertz. * `GHz` - The Enum value GHz represents that the measurement unit is in gigahertz. * `Percentage` - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity.
	MaxValueUnit *string `json:"MaxValueUnit,omitempty"`
	// Unit type for the expansion request, i.e., if the increase is requested as a raw value in TB, GB, etc., or in percentage increase. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes. * `GB` - The Enum value GB represents that the measurement unit is in gigabytes. * `MHz` - The Enum value MHz represents that the measurement unit is in megahertz. * `GHz` - The Enum value GHz represents that the measurement unit is in gigahertz. * `Percentage` - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity.
	UnitType *string `json:"UnitType,omitempty"`
	// Value of the expansion request which can be absolute value or percentage increase.
	Value                *float32                                                   `json:"Value,omitempty"`
	ExpansionRequest     NullableRecommendationHardwareExpansionRequestRelationship `json:"ExpansionRequest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationHardwareExpansionRequestItem RecommendationHardwareExpansionRequestItem

// NewRecommendationHardwareExpansionRequestItem instantiates a new RecommendationHardwareExpansionRequestItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationHardwareExpansionRequestItem(classId string, objectType string) *RecommendationHardwareExpansionRequestItem {
	this := RecommendationHardwareExpansionRequestItem{}
	this.ClassId = classId
	this.ObjectType = objectType
	var itemType string = "None"
	this.ItemType = &itemType
	var maxValueUnit string = "TB"
	this.MaxValueUnit = &maxValueUnit
	var unitType string = "TB"
	this.UnitType = &unitType
	return &this
}

// NewRecommendationHardwareExpansionRequestItemWithDefaults instantiates a new RecommendationHardwareExpansionRequestItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationHardwareExpansionRequestItemWithDefaults() *RecommendationHardwareExpansionRequestItem {
	this := RecommendationHardwareExpansionRequestItem{}
	var classId string = "recommendation.HardwareExpansionRequestItem"
	this.ClassId = classId
	var objectType string = "recommendation.HardwareExpansionRequestItem"
	this.ObjectType = objectType
	var itemType string = "None"
	this.ItemType = &itemType
	var maxValueUnit string = "TB"
	this.MaxValueUnit = &maxValueUnit
	var unitType string = "TB"
	this.UnitType = &unitType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationHardwareExpansionRequestItem) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationHardwareExpansionRequestItem) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "recommendation.HardwareExpansionRequestItem" of the ClassId field.
func (o *RecommendationHardwareExpansionRequestItem) GetDefaultClassId() interface{} {
	return "recommendation.HardwareExpansionRequestItem"
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationHardwareExpansionRequestItem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationHardwareExpansionRequestItem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "recommendation.HardwareExpansionRequestItem" of the ObjectType field.
func (o *RecommendationHardwareExpansionRequestItem) GetDefaultObjectType() interface{} {
	return "recommendation.HardwareExpansionRequestItem"
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestItem) GetItemType() string {
	if o == nil || IsNil(o.ItemType) {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetItemTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ItemType) {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestItem) HasItemType() bool {
	if o != nil && !IsNil(o.ItemType) {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *RecommendationHardwareExpansionRequestItem) SetItemType(v string) {
	o.ItemType = &v
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestItem) GetMaxValue() float32 {
	if o == nil || IsNil(o.MaxValue) {
		var ret float32
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetMaxValueOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxValue) {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestItem) HasMaxValue() bool {
	if o != nil && !IsNil(o.MaxValue) {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given float32 and assigns it to the MaxValue field.
func (o *RecommendationHardwareExpansionRequestItem) SetMaxValue(v float32) {
	o.MaxValue = &v
}

// GetMaxValueUnit returns the MaxValueUnit field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestItem) GetMaxValueUnit() string {
	if o == nil || IsNil(o.MaxValueUnit) {
		var ret string
		return ret
	}
	return *o.MaxValueUnit
}

// GetMaxValueUnitOk returns a tuple with the MaxValueUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetMaxValueUnitOk() (*string, bool) {
	if o == nil || IsNil(o.MaxValueUnit) {
		return nil, false
	}
	return o.MaxValueUnit, true
}

// HasMaxValueUnit returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestItem) HasMaxValueUnit() bool {
	if o != nil && !IsNil(o.MaxValueUnit) {
		return true
	}

	return false
}

// SetMaxValueUnit gets a reference to the given string and assigns it to the MaxValueUnit field.
func (o *RecommendationHardwareExpansionRequestItem) SetMaxValueUnit(v string) {
	o.MaxValueUnit = &v
}

// GetUnitType returns the UnitType field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestItem) GetUnitType() string {
	if o == nil || IsNil(o.UnitType) {
		var ret string
		return ret
	}
	return *o.UnitType
}

// GetUnitTypeOk returns a tuple with the UnitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetUnitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UnitType) {
		return nil, false
	}
	return o.UnitType, true
}

// HasUnitType returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestItem) HasUnitType() bool {
	if o != nil && !IsNil(o.UnitType) {
		return true
	}

	return false
}

// SetUnitType gets a reference to the given string and assigns it to the UnitType field.
func (o *RecommendationHardwareExpansionRequestItem) SetUnitType(v string) {
	o.UnitType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestItem) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestItem) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *RecommendationHardwareExpansionRequestItem) SetValue(v float32) {
	o.Value = &v
}

// GetExpansionRequest returns the ExpansionRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationHardwareExpansionRequestItem) GetExpansionRequest() RecommendationHardwareExpansionRequestRelationship {
	if o == nil || IsNil(o.ExpansionRequest.Get()) {
		var ret RecommendationHardwareExpansionRequestRelationship
		return ret
	}
	return *o.ExpansionRequest.Get()
}

// GetExpansionRequestOk returns a tuple with the ExpansionRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationHardwareExpansionRequestItem) GetExpansionRequestOk() (*RecommendationHardwareExpansionRequestRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpansionRequest.Get(), o.ExpansionRequest.IsSet()
}

// HasExpansionRequest returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestItem) HasExpansionRequest() bool {
	if o != nil && o.ExpansionRequest.IsSet() {
		return true
	}

	return false
}

// SetExpansionRequest gets a reference to the given NullableRecommendationHardwareExpansionRequestRelationship and assigns it to the ExpansionRequest field.
func (o *RecommendationHardwareExpansionRequestItem) SetExpansionRequest(v RecommendationHardwareExpansionRequestRelationship) {
	o.ExpansionRequest.Set(&v)
}

// SetExpansionRequestNil sets the value for ExpansionRequest to be an explicit nil
func (o *RecommendationHardwareExpansionRequestItem) SetExpansionRequestNil() {
	o.ExpansionRequest.Set(nil)
}

// UnsetExpansionRequest ensures that no value is present for ExpansionRequest, not even an explicit nil
func (o *RecommendationHardwareExpansionRequestItem) UnsetExpansionRequest() {
	o.ExpansionRequest.Unset()
}

func (o RecommendationHardwareExpansionRequestItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationHardwareExpansionRequestItem) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ItemType) {
		toSerialize["ItemType"] = o.ItemType
	}
	if !IsNil(o.MaxValue) {
		toSerialize["MaxValue"] = o.MaxValue
	}
	if !IsNil(o.MaxValueUnit) {
		toSerialize["MaxValueUnit"] = o.MaxValueUnit
	}
	if !IsNil(o.UnitType) {
		toSerialize["UnitType"] = o.UnitType
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	if o.ExpansionRequest.IsSet() {
		toSerialize["ExpansionRequest"] = o.ExpansionRequest.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecommendationHardwareExpansionRequestItem) UnmarshalJSON(data []byte) (err error) {
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
	type RecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of hardware item for which the capacity increase is requested by the user. For example, CPU. * `None` - The Enum value None represents that no value was set for the hardware type. * `CPU` - The Enum value CPU represents that the hardware type requested for expansion is a physical CPU. * `Memory` - The Enum value Memory represents that the hardware type requested for expansion is a memory unit. * `Storage` - The Enum value Storage represents that the hardware type requested for expansion is a storage unit, ie, storage drives.
		ItemType *string `json:"ItemType,omitempty"`
		// The maximum value allowed for expansion for the hardware type on the referred registered device.
		MaxValue *float32 `json:"MaxValue,omitempty"`
		// Unit type for the maximum value of the resource. For example, TB, GB, MB. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes. * `GB` - The Enum value GB represents that the measurement unit is in gigabytes. * `MHz` - The Enum value MHz represents that the measurement unit is in megahertz. * `GHz` - The Enum value GHz represents that the measurement unit is in gigahertz. * `Percentage` - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity.
		MaxValueUnit *string `json:"MaxValueUnit,omitempty"`
		// Unit type for the expansion request, i.e., if the increase is requested as a raw value in TB, GB, etc., or in percentage increase. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes. * `GB` - The Enum value GB represents that the measurement unit is in gigabytes. * `MHz` - The Enum value MHz represents that the measurement unit is in megahertz. * `GHz` - The Enum value GHz represents that the measurement unit is in gigahertz. * `Percentage` - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity.
		UnitType *string `json:"UnitType,omitempty"`
		// Value of the expansion request which can be absolute value or percentage increase.
		Value            *float32                                                   `json:"Value,omitempty"`
		ExpansionRequest NullableRecommendationHardwareExpansionRequestRelationship `json:"ExpansionRequest,omitempty"`
	}

	varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct := RecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct)
	if err == nil {
		varRecommendationHardwareExpansionRequestItem := _RecommendationHardwareExpansionRequestItem{}
		varRecommendationHardwareExpansionRequestItem.ClassId = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.ClassId
		varRecommendationHardwareExpansionRequestItem.ObjectType = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.ObjectType
		varRecommendationHardwareExpansionRequestItem.ItemType = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.ItemType
		varRecommendationHardwareExpansionRequestItem.MaxValue = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.MaxValue
		varRecommendationHardwareExpansionRequestItem.MaxValueUnit = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.MaxValueUnit
		varRecommendationHardwareExpansionRequestItem.UnitType = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.UnitType
		varRecommendationHardwareExpansionRequestItem.Value = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.Value
		varRecommendationHardwareExpansionRequestItem.ExpansionRequest = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.ExpansionRequest
		*o = RecommendationHardwareExpansionRequestItem(varRecommendationHardwareExpansionRequestItem)
	} else {
		return err
	}

	varRecommendationHardwareExpansionRequestItem := _RecommendationHardwareExpansionRequestItem{}

	err = json.Unmarshal(data, &varRecommendationHardwareExpansionRequestItem)
	if err == nil {
		o.MoBaseMo = varRecommendationHardwareExpansionRequestItem.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ItemType")
		delete(additionalProperties, "MaxValue")
		delete(additionalProperties, "MaxValueUnit")
		delete(additionalProperties, "UnitType")
		delete(additionalProperties, "Value")
		delete(additionalProperties, "ExpansionRequest")

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

type NullableRecommendationHardwareExpansionRequestItem struct {
	value *RecommendationHardwareExpansionRequestItem
	isSet bool
}

func (v NullableRecommendationHardwareExpansionRequestItem) Get() *RecommendationHardwareExpansionRequestItem {
	return v.value
}

func (v *NullableRecommendationHardwareExpansionRequestItem) Set(val *RecommendationHardwareExpansionRequestItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationHardwareExpansionRequestItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationHardwareExpansionRequestItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationHardwareExpansionRequestItem(val *RecommendationHardwareExpansionRequestItem) *NullableRecommendationHardwareExpansionRequestItem {
	return &NullableRecommendationHardwareExpansionRequestItem{value: val, isSet: true}
}

func (v NullableRecommendationHardwareExpansionRequestItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationHardwareExpansionRequestItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
