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

// checks if the NiatelemetryNiaFeatureUsageList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryNiaFeatureUsageList{}

// NiatelemetryNiaFeatureUsageList This resource list is returned as a response to a HTTP GET request that does not include a specific resource identifier.
type NiatelemetryNiaFeatureUsageList struct {
	MoBaseResponse
	// The total number of 'niatelemetry.NiaFeatureUsage' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty"`
	// The array of 'niatelemetry.NiaFeatureUsage' resources matching the request.
	Results              []NiatelemetryNiaFeatureUsage `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNiaFeatureUsageList NiatelemetryNiaFeatureUsageList

// NewNiatelemetryNiaFeatureUsageList instantiates a new NiatelemetryNiaFeatureUsageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNiaFeatureUsageList(objectType string) *NiatelemetryNiaFeatureUsageList {
	this := NiatelemetryNiaFeatureUsageList{}
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNiaFeatureUsageListWithDefaults instantiates a new NiatelemetryNiaFeatureUsageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNiaFeatureUsageListWithDefaults() *NiatelemetryNiaFeatureUsageList {
	this := NiatelemetryNiaFeatureUsageList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *NiatelemetryNiaFeatureUsageList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaFeatureUsageList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaFeatureUsageList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *NiatelemetryNiaFeatureUsageList) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryNiaFeatureUsageList) GetResults() []NiatelemetryNiaFeatureUsage {
	if o == nil {
		var ret []NiatelemetryNiaFeatureUsage
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryNiaFeatureUsageList) GetResultsOk() ([]NiatelemetryNiaFeatureUsage, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *NiatelemetryNiaFeatureUsageList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []NiatelemetryNiaFeatureUsage and assigns it to the Results field.
func (o *NiatelemetryNiaFeatureUsageList) SetResults(v []NiatelemetryNiaFeatureUsage) {
	o.Results = v
}

func (o NiatelemetryNiaFeatureUsageList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryNiaFeatureUsageList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseResponse, errMoBaseResponse := json.Marshal(o.MoBaseResponse)
	if errMoBaseResponse != nil {
		return map[string]interface{}{}, errMoBaseResponse
	}
	errMoBaseResponse = json.Unmarshal([]byte(serializedMoBaseResponse), &toSerialize)
	if errMoBaseResponse != nil {
		return map[string]interface{}{}, errMoBaseResponse
	}
	if !IsNil(o.Count) {
		toSerialize["Count"] = o.Count
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryNiaFeatureUsageList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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
	type NiatelemetryNiaFeatureUsageListWithoutEmbeddedStruct struct {
		// The total number of 'niatelemetry.NiaFeatureUsage' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
		Count *int32 `json:"Count,omitempty"`
		// The array of 'niatelemetry.NiaFeatureUsage' resources matching the request.
		Results []NiatelemetryNiaFeatureUsage `json:"Results,omitempty"`
	}

	varNiatelemetryNiaFeatureUsageListWithoutEmbeddedStruct := NiatelemetryNiaFeatureUsageListWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryNiaFeatureUsageListWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryNiaFeatureUsageList := _NiatelemetryNiaFeatureUsageList{}
		varNiatelemetryNiaFeatureUsageList.Count = varNiatelemetryNiaFeatureUsageListWithoutEmbeddedStruct.Count
		varNiatelemetryNiaFeatureUsageList.Results = varNiatelemetryNiaFeatureUsageListWithoutEmbeddedStruct.Results
		*o = NiatelemetryNiaFeatureUsageList(varNiatelemetryNiaFeatureUsageList)
	} else {
		return err
	}

	varNiatelemetryNiaFeatureUsageList := _NiatelemetryNiaFeatureUsageList{}

	err = json.Unmarshal(data, &varNiatelemetryNiaFeatureUsageList)
	if err == nil {
		o.MoBaseResponse = varNiatelemetryNiaFeatureUsageList.MoBaseResponse
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "Count")
		delete(additionalProperties, "Results")

		// remove fields from embedded structs
		reflectMoBaseResponse := reflect.ValueOf(o.MoBaseResponse)
		for i := 0; i < reflectMoBaseResponse.Type().NumField(); i++ {
			t := reflectMoBaseResponse.Type().Field(i)

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

type NullableNiatelemetryNiaFeatureUsageList struct {
	value *NiatelemetryNiaFeatureUsageList
	isSet bool
}

func (v NullableNiatelemetryNiaFeatureUsageList) Get() *NiatelemetryNiaFeatureUsageList {
	return v.value
}

func (v *NullableNiatelemetryNiaFeatureUsageList) Set(val *NiatelemetryNiaFeatureUsageList) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNiaFeatureUsageList) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNiaFeatureUsageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNiaFeatureUsageList(val *NiatelemetryNiaFeatureUsageList) *NullableNiatelemetryNiaFeatureUsageList {
	return &NullableNiatelemetryNiaFeatureUsageList{value: val, isSet: true}
}

func (v NullableNiatelemetryNiaFeatureUsageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNiaFeatureUsageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
