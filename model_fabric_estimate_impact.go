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

// checks if the FabricEstimateImpact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricEstimateImpact{}

// FabricEstimateImpact Before submitting switch profile deploy operation, the estimate impact helps to know the list of components be impacted and require switch reboot.
type FabricEstimateImpact struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                  `json:"ObjectType"`
	Profile              NullableFabricSwitchProfileRelationship `json:"Profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricEstimateImpact FabricEstimateImpact

// NewFabricEstimateImpact instantiates a new FabricEstimateImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricEstimateImpact(classId string, objectType string) *FabricEstimateImpact {
	this := FabricEstimateImpact{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricEstimateImpactWithDefaults instantiates a new FabricEstimateImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricEstimateImpactWithDefaults() *FabricEstimateImpact {
	this := FabricEstimateImpact{}
	var classId string = "fabric.EstimateImpact"
	this.ClassId = classId
	var objectType string = "fabric.EstimateImpact"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricEstimateImpact) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricEstimateImpact) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricEstimateImpact) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.EstimateImpact" of the ClassId field.
func (o *FabricEstimateImpact) GetDefaultClassId() interface{} {
	return "fabric.EstimateImpact"
}

// GetObjectType returns the ObjectType field value
func (o *FabricEstimateImpact) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricEstimateImpact) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricEstimateImpact) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.EstimateImpact" of the ObjectType field.
func (o *FabricEstimateImpact) GetDefaultObjectType() interface{} {
	return "fabric.EstimateImpact"
}

// GetProfile returns the Profile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricEstimateImpact) GetProfile() FabricSwitchProfileRelationship {
	if o == nil || IsNil(o.Profile.Get()) {
		var ret FabricSwitchProfileRelationship
		return ret
	}
	return *o.Profile.Get()
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricEstimateImpact) GetProfileOk() (*FabricSwitchProfileRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profile.Get(), o.Profile.IsSet()
}

// HasProfile returns a boolean if a field has been set.
func (o *FabricEstimateImpact) HasProfile() bool {
	if o != nil && o.Profile.IsSet() {
		return true
	}

	return false
}

// SetProfile gets a reference to the given NullableFabricSwitchProfileRelationship and assigns it to the Profile field.
func (o *FabricEstimateImpact) SetProfile(v FabricSwitchProfileRelationship) {
	o.Profile.Set(&v)
}

// SetProfileNil sets the value for Profile to be an explicit nil
func (o *FabricEstimateImpact) SetProfileNil() {
	o.Profile.Set(nil)
}

// UnsetProfile ensures that no value is present for Profile, not even an explicit nil
func (o *FabricEstimateImpact) UnsetProfile() {
	o.Profile.Unset()
}

func (o FabricEstimateImpact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricEstimateImpact) ToMap() (map[string]interface{}, error) {
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
	if o.Profile.IsSet() {
		toSerialize["Profile"] = o.Profile.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricEstimateImpact) UnmarshalJSON(data []byte) (err error) {
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
	type FabricEstimateImpactWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                  `json:"ObjectType"`
		Profile    NullableFabricSwitchProfileRelationship `json:"Profile,omitempty"`
	}

	varFabricEstimateImpactWithoutEmbeddedStruct := FabricEstimateImpactWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricEstimateImpactWithoutEmbeddedStruct)
	if err == nil {
		varFabricEstimateImpact := _FabricEstimateImpact{}
		varFabricEstimateImpact.ClassId = varFabricEstimateImpactWithoutEmbeddedStruct.ClassId
		varFabricEstimateImpact.ObjectType = varFabricEstimateImpactWithoutEmbeddedStruct.ObjectType
		varFabricEstimateImpact.Profile = varFabricEstimateImpactWithoutEmbeddedStruct.Profile
		*o = FabricEstimateImpact(varFabricEstimateImpact)
	} else {
		return err
	}

	varFabricEstimateImpact := _FabricEstimateImpact{}

	err = json.Unmarshal(data, &varFabricEstimateImpact)
	if err == nil {
		o.MoBaseMo = varFabricEstimateImpact.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Profile")

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

type NullableFabricEstimateImpact struct {
	value *FabricEstimateImpact
	isSet bool
}

func (v NullableFabricEstimateImpact) Get() *FabricEstimateImpact {
	return v.value
}

func (v *NullableFabricEstimateImpact) Set(val *FabricEstimateImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricEstimateImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricEstimateImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricEstimateImpact(val *FabricEstimateImpact) *NullableFabricEstimateImpact {
	return &NullableFabricEstimateImpact{value: val, isSet: true}
}

func (v NullableFabricEstimateImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricEstimateImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
