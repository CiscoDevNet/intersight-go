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

// checks if the TamSecurityAdvisoryDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TamSecurityAdvisoryDetails{}

// TamSecurityAdvisoryDetails Details pertaining to a security advisory defined by a given advisory.
type TamSecurityAdvisoryDetails struct {
	TamBaseAdvisoryDetails
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// CVSS version 3 base score for the security Advisory.
	BaseScore *float32 `json:"BaseScore,omitempty"`
	CveIds    []string `json:"CveIds,omitempty"`
	// CVSS version 3 environmental score for the security Advisory.
	EnvironmentalScore *float32 `json:"EnvironmentalScore,omitempty"`
	// Cisco assigned status of the published advisory. Depends on whether the investigation is complete or on-going. * `interim` - The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available. * `final` - Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability.
	Status *string `json:"Status,omitempty"`
	// CVSS version 3 temporal score for the security Advisory.
	TemporalScore        *float32 `json:"TemporalScore,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamSecurityAdvisoryDetails TamSecurityAdvisoryDetails

// NewTamSecurityAdvisoryDetails instantiates a new TamSecurityAdvisoryDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamSecurityAdvisoryDetails(classId string, objectType string) *TamSecurityAdvisoryDetails {
	this := TamSecurityAdvisoryDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "interim"
	this.Status = &status
	return &this
}

// NewTamSecurityAdvisoryDetailsWithDefaults instantiates a new TamSecurityAdvisoryDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamSecurityAdvisoryDetailsWithDefaults() *TamSecurityAdvisoryDetails {
	this := TamSecurityAdvisoryDetails{}
	var classId string = "tam.SecurityAdvisoryDetails"
	this.ClassId = classId
	var objectType string = "tam.SecurityAdvisoryDetails"
	this.ObjectType = objectType
	var status string = "interim"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamSecurityAdvisoryDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamSecurityAdvisoryDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "tam.SecurityAdvisoryDetails" of the ClassId field.
func (o *TamSecurityAdvisoryDetails) GetDefaultClassId() interface{} {
	return "tam.SecurityAdvisoryDetails"
}

// GetObjectType returns the ObjectType field value
func (o *TamSecurityAdvisoryDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamSecurityAdvisoryDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "tam.SecurityAdvisoryDetails" of the ObjectType field.
func (o *TamSecurityAdvisoryDetails) GetDefaultObjectType() interface{} {
	return "tam.SecurityAdvisoryDetails"
}

// GetBaseScore returns the BaseScore field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryDetails) GetBaseScore() float32 {
	if o == nil || IsNil(o.BaseScore) {
		var ret float32
		return ret
	}
	return *o.BaseScore
}

// GetBaseScoreOk returns a tuple with the BaseScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryDetails) GetBaseScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseScore) {
		return nil, false
	}
	return o.BaseScore, true
}

// HasBaseScore returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryDetails) HasBaseScore() bool {
	if o != nil && !IsNil(o.BaseScore) {
		return true
	}

	return false
}

// SetBaseScore gets a reference to the given float32 and assigns it to the BaseScore field.
func (o *TamSecurityAdvisoryDetails) SetBaseScore(v float32) {
	o.BaseScore = &v
}

// GetCveIds returns the CveIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamSecurityAdvisoryDetails) GetCveIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CveIds
}

// GetCveIdsOk returns a tuple with the CveIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamSecurityAdvisoryDetails) GetCveIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CveIds) {
		return nil, false
	}
	return o.CveIds, true
}

// HasCveIds returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryDetails) HasCveIds() bool {
	if o != nil && !IsNil(o.CveIds) {
		return true
	}

	return false
}

// SetCveIds gets a reference to the given []string and assigns it to the CveIds field.
func (o *TamSecurityAdvisoryDetails) SetCveIds(v []string) {
	o.CveIds = v
}

// GetEnvironmentalScore returns the EnvironmentalScore field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryDetails) GetEnvironmentalScore() float32 {
	if o == nil || IsNil(o.EnvironmentalScore) {
		var ret float32
		return ret
	}
	return *o.EnvironmentalScore
}

// GetEnvironmentalScoreOk returns a tuple with the EnvironmentalScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryDetails) GetEnvironmentalScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.EnvironmentalScore) {
		return nil, false
	}
	return o.EnvironmentalScore, true
}

// HasEnvironmentalScore returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryDetails) HasEnvironmentalScore() bool {
	if o != nil && !IsNil(o.EnvironmentalScore) {
		return true
	}

	return false
}

// SetEnvironmentalScore gets a reference to the given float32 and assigns it to the EnvironmentalScore field.
func (o *TamSecurityAdvisoryDetails) SetEnvironmentalScore(v float32) {
	o.EnvironmentalScore = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryDetails) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryDetails) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryDetails) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TamSecurityAdvisoryDetails) SetStatus(v string) {
	o.Status = &v
}

// GetTemporalScore returns the TemporalScore field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryDetails) GetTemporalScore() float32 {
	if o == nil || IsNil(o.TemporalScore) {
		var ret float32
		return ret
	}
	return *o.TemporalScore
}

// GetTemporalScoreOk returns a tuple with the TemporalScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryDetails) GetTemporalScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.TemporalScore) {
		return nil, false
	}
	return o.TemporalScore, true
}

// HasTemporalScore returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryDetails) HasTemporalScore() bool {
	if o != nil && !IsNil(o.TemporalScore) {
		return true
	}

	return false
}

// SetTemporalScore gets a reference to the given float32 and assigns it to the TemporalScore field.
func (o *TamSecurityAdvisoryDetails) SetTemporalScore(v float32) {
	o.TemporalScore = &v
}

func (o TamSecurityAdvisoryDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TamSecurityAdvisoryDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTamBaseAdvisoryDetails, errTamBaseAdvisoryDetails := json.Marshal(o.TamBaseAdvisoryDetails)
	if errTamBaseAdvisoryDetails != nil {
		return map[string]interface{}{}, errTamBaseAdvisoryDetails
	}
	errTamBaseAdvisoryDetails = json.Unmarshal([]byte(serializedTamBaseAdvisoryDetails), &toSerialize)
	if errTamBaseAdvisoryDetails != nil {
		return map[string]interface{}{}, errTamBaseAdvisoryDetails
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.BaseScore) {
		toSerialize["BaseScore"] = o.BaseScore
	}
	if o.CveIds != nil {
		toSerialize["CveIds"] = o.CveIds
	}
	if !IsNil(o.EnvironmentalScore) {
		toSerialize["EnvironmentalScore"] = o.EnvironmentalScore
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.TemporalScore) {
		toSerialize["TemporalScore"] = o.TemporalScore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TamSecurityAdvisoryDetails) UnmarshalJSON(data []byte) (err error) {
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
	type TamSecurityAdvisoryDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// CVSS version 3 base score for the security Advisory.
		BaseScore *float32 `json:"BaseScore,omitempty"`
		CveIds    []string `json:"CveIds,omitempty"`
		// CVSS version 3 environmental score for the security Advisory.
		EnvironmentalScore *float32 `json:"EnvironmentalScore,omitempty"`
		// Cisco assigned status of the published advisory. Depends on whether the investigation is complete or on-going. * `interim` - The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available. * `final` - Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability.
		Status *string `json:"Status,omitempty"`
		// CVSS version 3 temporal score for the security Advisory.
		TemporalScore *float32 `json:"TemporalScore,omitempty"`
	}

	varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct := TamSecurityAdvisoryDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct)
	if err == nil {
		varTamSecurityAdvisoryDetails := _TamSecurityAdvisoryDetails{}
		varTamSecurityAdvisoryDetails.ClassId = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.ClassId
		varTamSecurityAdvisoryDetails.ObjectType = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.ObjectType
		varTamSecurityAdvisoryDetails.BaseScore = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.BaseScore
		varTamSecurityAdvisoryDetails.CveIds = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.CveIds
		varTamSecurityAdvisoryDetails.EnvironmentalScore = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.EnvironmentalScore
		varTamSecurityAdvisoryDetails.Status = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.Status
		varTamSecurityAdvisoryDetails.TemporalScore = varTamSecurityAdvisoryDetailsWithoutEmbeddedStruct.TemporalScore
		*o = TamSecurityAdvisoryDetails(varTamSecurityAdvisoryDetails)
	} else {
		return err
	}

	varTamSecurityAdvisoryDetails := _TamSecurityAdvisoryDetails{}

	err = json.Unmarshal(data, &varTamSecurityAdvisoryDetails)
	if err == nil {
		o.TamBaseAdvisoryDetails = varTamSecurityAdvisoryDetails.TamBaseAdvisoryDetails
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BaseScore")
		delete(additionalProperties, "CveIds")
		delete(additionalProperties, "EnvironmentalScore")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TemporalScore")

		// remove fields from embedded structs
		reflectTamBaseAdvisoryDetails := reflect.ValueOf(o.TamBaseAdvisoryDetails)
		for i := 0; i < reflectTamBaseAdvisoryDetails.Type().NumField(); i++ {
			t := reflectTamBaseAdvisoryDetails.Type().Field(i)

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

type NullableTamSecurityAdvisoryDetails struct {
	value *TamSecurityAdvisoryDetails
	isSet bool
}

func (v NullableTamSecurityAdvisoryDetails) Get() *TamSecurityAdvisoryDetails {
	return v.value
}

func (v *NullableTamSecurityAdvisoryDetails) Set(val *TamSecurityAdvisoryDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTamSecurityAdvisoryDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTamSecurityAdvisoryDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamSecurityAdvisoryDetails(val *TamSecurityAdvisoryDetails) *NullableTamSecurityAdvisoryDetails {
	return &NullableTamSecurityAdvisoryDetails{value: val, isSet: true}
}

func (v NullableTamSecurityAdvisoryDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamSecurityAdvisoryDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
