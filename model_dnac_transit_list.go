/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-17769
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

// checks if the DnacTransitList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnacTransitList{}

// DnacTransitList This resource list is returned as a response to a HTTP GET request that does not include a specific resource identifier.
type DnacTransitList struct {
	MoBaseResponse
	// The total number of 'dnac.Transit' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty"`
	// The array of 'dnac.Transit' resources matching the request.
	Results              []DnacTransit `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DnacTransitList DnacTransitList

// NewDnacTransitList instantiates a new DnacTransitList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnacTransitList(objectType string) *DnacTransitList {
	this := DnacTransitList{}
	this.ObjectType = objectType
	return &this
}

// NewDnacTransitListWithDefaults instantiates a new DnacTransitList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnacTransitListWithDefaults() *DnacTransitList {
	this := DnacTransitList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DnacTransitList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacTransitList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DnacTransitList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *DnacTransitList) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DnacTransitList) GetResults() []DnacTransit {
	if o == nil {
		var ret []DnacTransit
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnacTransitList) GetResultsOk() ([]DnacTransit, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DnacTransitList) HasResults() bool {
	if o != nil && IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DnacTransit and assigns it to the Results field.
func (o *DnacTransitList) SetResults(v []DnacTransit) {
	o.Results = v
}

func (o DnacTransitList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnacTransitList) ToMap() (map[string]interface{}, error) {
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

func (o *DnacTransitList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	type DnacTransitListWithoutEmbeddedStruct struct {
		// The total number of 'dnac.Transit' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
		Count *int32 `json:"Count,omitempty"`
		// The array of 'dnac.Transit' resources matching the request.
		Results []DnacTransit `json:"Results,omitempty"`
	}

	varDnacTransitListWithoutEmbeddedStruct := DnacTransitListWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varDnacTransitListWithoutEmbeddedStruct)
	if err == nil {
		varDnacTransitList := _DnacTransitList{}
		varDnacTransitList.Count = varDnacTransitListWithoutEmbeddedStruct.Count
		varDnacTransitList.Results = varDnacTransitListWithoutEmbeddedStruct.Results
		*o = DnacTransitList(varDnacTransitList)
	} else {
		return err
	}

	varDnacTransitList := _DnacTransitList{}

	err = json.Unmarshal(data, &varDnacTransitList)
	if err == nil {
		o.MoBaseResponse = varDnacTransitList.MoBaseResponse
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

type NullableDnacTransitList struct {
	value *DnacTransitList
	isSet bool
}

func (v NullableDnacTransitList) Get() *DnacTransitList {
	return v.value
}

func (v *NullableDnacTransitList) Set(val *DnacTransitList) {
	v.value = val
	v.isSet = true
}

func (v NullableDnacTransitList) IsSet() bool {
	return v.isSet
}

func (v *NullableDnacTransitList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnacTransitList(val *DnacTransitList) *NullableDnacTransitList {
	return &NullableDnacTransitList{value: val, isSet: true}
}

func (v NullableDnacTransitList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnacTransitList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}