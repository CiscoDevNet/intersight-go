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

// checks if the VnicEthVethInventoryList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicEthVethInventoryList{}

// VnicEthVethInventoryList This resource list is returned as a response to a HTTP GET request that does not include a specific resource identifier.
type VnicEthVethInventoryList struct {
	MoBaseResponse
	// The total number of 'vnic.EthVethInventory' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty"`
	// The array of 'vnic.EthVethInventory' resources matching the request.
	Results              []VnicEthVethInventory `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicEthVethInventoryList VnicEthVethInventoryList

// NewVnicEthVethInventoryList instantiates a new VnicEthVethInventoryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthVethInventoryList(objectType string) *VnicEthVethInventoryList {
	this := VnicEthVethInventoryList{}
	this.ObjectType = objectType
	return &this
}

// NewVnicEthVethInventoryListWithDefaults instantiates a new VnicEthVethInventoryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthVethInventoryListWithDefaults() *VnicEthVethInventoryList {
	this := VnicEthVethInventoryList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *VnicEthVethInventoryList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthVethInventoryList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *VnicEthVethInventoryList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *VnicEthVethInventoryList) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthVethInventoryList) GetResults() []VnicEthVethInventory {
	if o == nil {
		var ret []VnicEthVethInventory
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthVethInventoryList) GetResultsOk() ([]VnicEthVethInventory, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *VnicEthVethInventoryList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []VnicEthVethInventory and assigns it to the Results field.
func (o *VnicEthVethInventoryList) SetResults(v []VnicEthVethInventory) {
	o.Results = v
}

func (o VnicEthVethInventoryList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicEthVethInventoryList) ToMap() (map[string]interface{}, error) {
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

func (o *VnicEthVethInventoryList) UnmarshalJSON(data []byte) (err error) {
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
	type VnicEthVethInventoryListWithoutEmbeddedStruct struct {
		// The total number of 'vnic.EthVethInventory' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
		Count *int32 `json:"Count,omitempty"`
		// The array of 'vnic.EthVethInventory' resources matching the request.
		Results []VnicEthVethInventory `json:"Results,omitempty"`
	}

	varVnicEthVethInventoryListWithoutEmbeddedStruct := VnicEthVethInventoryListWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicEthVethInventoryListWithoutEmbeddedStruct)
	if err == nil {
		varVnicEthVethInventoryList := _VnicEthVethInventoryList{}
		varVnicEthVethInventoryList.Count = varVnicEthVethInventoryListWithoutEmbeddedStruct.Count
		varVnicEthVethInventoryList.Results = varVnicEthVethInventoryListWithoutEmbeddedStruct.Results
		*o = VnicEthVethInventoryList(varVnicEthVethInventoryList)
	} else {
		return err
	}

	varVnicEthVethInventoryList := _VnicEthVethInventoryList{}

	err = json.Unmarshal(data, &varVnicEthVethInventoryList)
	if err == nil {
		o.MoBaseResponse = varVnicEthVethInventoryList.MoBaseResponse
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

type NullableVnicEthVethInventoryList struct {
	value *VnicEthVethInventoryList
	isSet bool
}

func (v NullableVnicEthVethInventoryList) Get() *VnicEthVethInventoryList {
	return v.value
}

func (v *NullableVnicEthVethInventoryList) Set(val *VnicEthVethInventoryList) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthVethInventoryList) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthVethInventoryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthVethInventoryList(val *VnicEthVethInventoryList) *NullableVnicEthVethInventoryList {
	return &NullableVnicEthVethInventoryList{value: val, isSet: true}
}

func (v NullableVnicEthVethInventoryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthVethInventoryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
