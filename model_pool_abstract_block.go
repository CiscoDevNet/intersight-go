/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// PoolAbstractBlock Abstract base counters for a block of identifiers.
type PoolAbstractBlock struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Free IDs that can be allocated in this block.
	FreeBlockCount *int64 `json:"FreeBlockCount,omitempty"`
	// Moving counter to allocate the next identifier.
	NextIdAllocator      *int64 `json:"NextIdAllocator,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolAbstractBlock PoolAbstractBlock

// NewPoolAbstractBlock instantiates a new PoolAbstractBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractBlock(classId string, objectType string) *PoolAbstractBlock {
	this := PoolAbstractBlock{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPoolAbstractBlockWithDefaults instantiates a new PoolAbstractBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractBlockWithDefaults() *PoolAbstractBlock {
	this := PoolAbstractBlock{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *PoolAbstractBlock) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlock) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PoolAbstractBlock) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PoolAbstractBlock) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlock) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PoolAbstractBlock) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFreeBlockCount returns the FreeBlockCount field value if set, zero value otherwise.
func (o *PoolAbstractBlock) GetFreeBlockCount() int64 {
	if o == nil || o.FreeBlockCount == nil {
		var ret int64
		return ret
	}
	return *o.FreeBlockCount
}

// GetFreeBlockCountOk returns a tuple with the FreeBlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlock) GetFreeBlockCountOk() (*int64, bool) {
	if o == nil || o.FreeBlockCount == nil {
		return nil, false
	}
	return o.FreeBlockCount, true
}

// HasFreeBlockCount returns a boolean if a field has been set.
func (o *PoolAbstractBlock) HasFreeBlockCount() bool {
	if o != nil && o.FreeBlockCount != nil {
		return true
	}

	return false
}

// SetFreeBlockCount gets a reference to the given int64 and assigns it to the FreeBlockCount field.
func (o *PoolAbstractBlock) SetFreeBlockCount(v int64) {
	o.FreeBlockCount = &v
}

// GetNextIdAllocator returns the NextIdAllocator field value if set, zero value otherwise.
func (o *PoolAbstractBlock) GetNextIdAllocator() int64 {
	if o == nil || o.NextIdAllocator == nil {
		var ret int64
		return ret
	}
	return *o.NextIdAllocator
}

// GetNextIdAllocatorOk returns a tuple with the NextIdAllocator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlock) GetNextIdAllocatorOk() (*int64, bool) {
	if o == nil || o.NextIdAllocator == nil {
		return nil, false
	}
	return o.NextIdAllocator, true
}

// HasNextIdAllocator returns a boolean if a field has been set.
func (o *PoolAbstractBlock) HasNextIdAllocator() bool {
	if o != nil && o.NextIdAllocator != nil {
		return true
	}

	return false
}

// SetNextIdAllocator gets a reference to the given int64 and assigns it to the NextIdAllocator field.
func (o *PoolAbstractBlock) SetNextIdAllocator(v int64) {
	o.NextIdAllocator = &v
}

func (o PoolAbstractBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FreeBlockCount != nil {
		toSerialize["FreeBlockCount"] = o.FreeBlockCount
	}
	if o.NextIdAllocator != nil {
		toSerialize["NextIdAllocator"] = o.NextIdAllocator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PoolAbstractBlock) UnmarshalJSON(bytes []byte) (err error) {
	type PoolAbstractBlockWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Free IDs that can be allocated in this block.
		FreeBlockCount *int64 `json:"FreeBlockCount,omitempty"`
		// Moving counter to allocate the next identifier.
		NextIdAllocator *int64 `json:"NextIdAllocator,omitempty"`
	}

	varPoolAbstractBlockWithoutEmbeddedStruct := PoolAbstractBlockWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPoolAbstractBlockWithoutEmbeddedStruct)
	if err == nil {
		varPoolAbstractBlock := _PoolAbstractBlock{}
		varPoolAbstractBlock.ClassId = varPoolAbstractBlockWithoutEmbeddedStruct.ClassId
		varPoolAbstractBlock.ObjectType = varPoolAbstractBlockWithoutEmbeddedStruct.ObjectType
		varPoolAbstractBlock.FreeBlockCount = varPoolAbstractBlockWithoutEmbeddedStruct.FreeBlockCount
		varPoolAbstractBlock.NextIdAllocator = varPoolAbstractBlockWithoutEmbeddedStruct.NextIdAllocator
		*o = PoolAbstractBlock(varPoolAbstractBlock)
	} else {
		return err
	}

	varPoolAbstractBlock := _PoolAbstractBlock{}

	err = json.Unmarshal(bytes, &varPoolAbstractBlock)
	if err == nil {
		o.MoBaseMo = varPoolAbstractBlock.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FreeBlockCount")
		delete(additionalProperties, "NextIdAllocator")

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

type NullablePoolAbstractBlock struct {
	value *PoolAbstractBlock
	isSet bool
}

func (v NullablePoolAbstractBlock) Get() *PoolAbstractBlock {
	return v.value
}

func (v *NullablePoolAbstractBlock) Set(val *PoolAbstractBlock) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractBlock) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractBlock(val *PoolAbstractBlock) *NullablePoolAbstractBlock {
	return &NullablePoolAbstractBlock{value: val, isSet: true}
}

func (v NullablePoolAbstractBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}