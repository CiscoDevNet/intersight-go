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
)

// checks if the SearchSuggestItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchSuggestItem{}

// SearchSuggestItem This resource is used to search for objects matching a criteria in the endpoint.
type SearchSuggestItem struct {
	AdditionalProperties map[string]interface{}
}

type _SearchSuggestItem SearchSuggestItem

// NewSearchSuggestItem instantiates a new SearchSuggestItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSuggestItem() *SearchSuggestItem {
	this := SearchSuggestItem{}
	return &this
}

// NewSearchSuggestItemWithDefaults instantiates a new SearchSuggestItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSuggestItemWithDefaults() *SearchSuggestItem {
	this := SearchSuggestItem{}
	return &this
}

func (o SearchSuggestItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchSuggestItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchSuggestItem) UnmarshalJSON(data []byte) (err error) {
	varSearchSuggestItem := _SearchSuggestItem{}

	err = json.Unmarshal(data, &varSearchSuggestItem)

	if err != nil {
		return err
	}

	*o = SearchSuggestItem(varSearchSuggestItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchSuggestItem struct {
	value *SearchSuggestItem
	isSet bool
}

func (v NullableSearchSuggestItem) Get() *SearchSuggestItem {
	return v.value
}

func (v *NullableSearchSuggestItem) Set(val *SearchSuggestItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSuggestItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSuggestItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSuggestItem(val *SearchSuggestItem) *NullableSearchSuggestItem {
	return &NullableSearchSuggestItem{value: val, isSet: true}
}

func (v NullableSearchSuggestItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSuggestItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
