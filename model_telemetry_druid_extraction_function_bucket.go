/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-15473
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// checks if the TelemetryDruidExtractionFunctionBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidExtractionFunctionBucket{}

// TelemetryDruidExtractionFunctionBucket Bucket extraction function is used to bucket numerical values in each range of the given size by converting them to the same base value. Non numeric values are converted to null.
type TelemetryDruidExtractionFunctionBucket struct {
	Type string `json:"type"`
	// The size of the buckets (optional, default 1).
	Size *int32 `json:"size,omitempty"`
	// The offset for the buckets (optional, default 0).
	Offset               *int32 `json:"offset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidExtractionFunctionBucket TelemetryDruidExtractionFunctionBucket

// NewTelemetryDruidExtractionFunctionBucket instantiates a new TelemetryDruidExtractionFunctionBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidExtractionFunctionBucket(type_ string) *TelemetryDruidExtractionFunctionBucket {
	this := TelemetryDruidExtractionFunctionBucket{}
	this.Type = type_
	return &this
}

// NewTelemetryDruidExtractionFunctionBucketWithDefaults instantiates a new TelemetryDruidExtractionFunctionBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidExtractionFunctionBucketWithDefaults() *TelemetryDruidExtractionFunctionBucket {
	this := TelemetryDruidExtractionFunctionBucket{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidExtractionFunctionBucket) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionBucket) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidExtractionFunctionBucket) SetType(v string) {
	o.Type = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionBucket) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionBucket) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionBucket) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *TelemetryDruidExtractionFunctionBucket) SetSize(v int32) {
	o.Size = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionBucket) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionBucket) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionBucket) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *TelemetryDruidExtractionFunctionBucket) SetOffset(v int32) {
	o.Offset = &v
}

func (o TelemetryDruidExtractionFunctionBucket) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidExtractionFunctionBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidExtractionFunctionBucket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varTelemetryDruidExtractionFunctionBucket := _TelemetryDruidExtractionFunctionBucket{}

	err = json.Unmarshal(data, &varTelemetryDruidExtractionFunctionBucket)

	if err != nil {
		return err
	}

	*o = TelemetryDruidExtractionFunctionBucket(varTelemetryDruidExtractionFunctionBucket)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "size")
		delete(additionalProperties, "offset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidExtractionFunctionBucket struct {
	value *TelemetryDruidExtractionFunctionBucket
	isSet bool
}

func (v NullableTelemetryDruidExtractionFunctionBucket) Get() *TelemetryDruidExtractionFunctionBucket {
	return v.value
}

func (v *NullableTelemetryDruidExtractionFunctionBucket) Set(val *TelemetryDruidExtractionFunctionBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidExtractionFunctionBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidExtractionFunctionBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidExtractionFunctionBucket(val *TelemetryDruidExtractionFunctionBucket) *NullableTelemetryDruidExtractionFunctionBucket {
	return &NullableTelemetryDruidExtractionFunctionBucket{value: val, isSet: true}
}

func (v NullableTelemetryDruidExtractionFunctionBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidExtractionFunctionBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}