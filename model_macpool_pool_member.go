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

// checks if the MacpoolPoolMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MacpoolPoolMember{}

// MacpoolPoolMember PoolMember represents a single MAC address that is part of a pool.
type MacpoolPoolMember struct {
	PoolAbstractIdPoolMember
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// MAC Address of this pool member.
	MacAddress           *string                                `json:"MacAddress,omitempty"`
	AssignedToEntity     NullableMoBaseMoRelationship           `json:"AssignedToEntity,omitempty"`
	BlockHead            NullableMacpoolIdBlockRelationship     `json:"BlockHead,omitempty"`
	Peer                 NullableMacpoolLeaseRelationship       `json:"Peer,omitempty"`
	Pool                 NullableMacpoolPoolRelationship        `json:"Pool,omitempty"`
	Reservation          NullableMacpoolReservationRelationship `json:"Reservation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MacpoolPoolMember MacpoolPoolMember

// NewMacpoolPoolMember instantiates a new MacpoolPoolMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacpoolPoolMember(classId string, objectType string) *MacpoolPoolMember {
	this := MacpoolPoolMember{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assigned bool = false
	this.Assigned = &assigned
	var assignedByAnother bool = false
	this.AssignedByAnother = &assignedByAnother
	var reserved bool = false
	this.Reserved = &reserved
	return &this
}

// NewMacpoolPoolMemberWithDefaults instantiates a new MacpoolPoolMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacpoolPoolMemberWithDefaults() *MacpoolPoolMember {
	this := MacpoolPoolMember{}
	var classId string = "macpool.PoolMember"
	this.ClassId = classId
	var objectType string = "macpool.PoolMember"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MacpoolPoolMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MacpoolPoolMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MacpoolPoolMember) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "macpool.PoolMember" of the ClassId field.
func (o *MacpoolPoolMember) GetDefaultClassId() interface{} {
	return "macpool.PoolMember"
}

// GetObjectType returns the ObjectType field value
func (o *MacpoolPoolMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MacpoolPoolMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MacpoolPoolMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "macpool.PoolMember" of the ObjectType field.
func (o *MacpoolPoolMember) GetDefaultObjectType() interface{} {
	return "macpool.PoolMember"
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *MacpoolPoolMember) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolPoolMember) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *MacpoolPoolMember) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *MacpoolPoolMember) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MacpoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || IsNil(o.AssignedToEntity.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity.Get()
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MacpoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedToEntity.Get(), o.AssignedToEntity.IsSet()
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *MacpoolPoolMember) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity.IsSet() {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given NullableMoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *MacpoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity.Set(&v)
}

// SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil
func (o *MacpoolPoolMember) SetAssignedToEntityNil() {
	o.AssignedToEntity.Set(nil)
}

// UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
func (o *MacpoolPoolMember) UnsetAssignedToEntity() {
	o.AssignedToEntity.Unset()
}

// GetBlockHead returns the BlockHead field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MacpoolPoolMember) GetBlockHead() MacpoolIdBlockRelationship {
	if o == nil || IsNil(o.BlockHead.Get()) {
		var ret MacpoolIdBlockRelationship
		return ret
	}
	return *o.BlockHead.Get()
}

// GetBlockHeadOk returns a tuple with the BlockHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MacpoolPoolMember) GetBlockHeadOk() (*MacpoolIdBlockRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockHead.Get(), o.BlockHead.IsSet()
}

// HasBlockHead returns a boolean if a field has been set.
func (o *MacpoolPoolMember) HasBlockHead() bool {
	if o != nil && o.BlockHead.IsSet() {
		return true
	}

	return false
}

// SetBlockHead gets a reference to the given NullableMacpoolIdBlockRelationship and assigns it to the BlockHead field.
func (o *MacpoolPoolMember) SetBlockHead(v MacpoolIdBlockRelationship) {
	o.BlockHead.Set(&v)
}

// SetBlockHeadNil sets the value for BlockHead to be an explicit nil
func (o *MacpoolPoolMember) SetBlockHeadNil() {
	o.BlockHead.Set(nil)
}

// UnsetBlockHead ensures that no value is present for BlockHead, not even an explicit nil
func (o *MacpoolPoolMember) UnsetBlockHead() {
	o.BlockHead.Unset()
}

// GetPeer returns the Peer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MacpoolPoolMember) GetPeer() MacpoolLeaseRelationship {
	if o == nil || IsNil(o.Peer.Get()) {
		var ret MacpoolLeaseRelationship
		return ret
	}
	return *o.Peer.Get()
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MacpoolPoolMember) GetPeerOk() (*MacpoolLeaseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peer.Get(), o.Peer.IsSet()
}

// HasPeer returns a boolean if a field has been set.
func (o *MacpoolPoolMember) HasPeer() bool {
	if o != nil && o.Peer.IsSet() {
		return true
	}

	return false
}

// SetPeer gets a reference to the given NullableMacpoolLeaseRelationship and assigns it to the Peer field.
func (o *MacpoolPoolMember) SetPeer(v MacpoolLeaseRelationship) {
	o.Peer.Set(&v)
}

// SetPeerNil sets the value for Peer to be an explicit nil
func (o *MacpoolPoolMember) SetPeerNil() {
	o.Peer.Set(nil)
}

// UnsetPeer ensures that no value is present for Peer, not even an explicit nil
func (o *MacpoolPoolMember) UnsetPeer() {
	o.Peer.Unset()
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MacpoolPoolMember) GetPool() MacpoolPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret MacpoolPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MacpoolPoolMember) GetPoolOk() (*MacpoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *MacpoolPoolMember) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableMacpoolPoolRelationship and assigns it to the Pool field.
func (o *MacpoolPoolMember) SetPool(v MacpoolPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *MacpoolPoolMember) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *MacpoolPoolMember) UnsetPool() {
	o.Pool.Unset()
}

// GetReservation returns the Reservation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MacpoolPoolMember) GetReservation() MacpoolReservationRelationship {
	if o == nil || IsNil(o.Reservation.Get()) {
		var ret MacpoolReservationRelationship
		return ret
	}
	return *o.Reservation.Get()
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MacpoolPoolMember) GetReservationOk() (*MacpoolReservationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reservation.Get(), o.Reservation.IsSet()
}

// HasReservation returns a boolean if a field has been set.
func (o *MacpoolPoolMember) HasReservation() bool {
	if o != nil && o.Reservation.IsSet() {
		return true
	}

	return false
}

// SetReservation gets a reference to the given NullableMacpoolReservationRelationship and assigns it to the Reservation field.
func (o *MacpoolPoolMember) SetReservation(v MacpoolReservationRelationship) {
	o.Reservation.Set(&v)
}

// SetReservationNil sets the value for Reservation to be an explicit nil
func (o *MacpoolPoolMember) SetReservationNil() {
	o.Reservation.Set(nil)
}

// UnsetReservation ensures that no value is present for Reservation, not even an explicit nil
func (o *MacpoolPoolMember) UnsetReservation() {
	o.Reservation.Unset()
}

func (o MacpoolPoolMember) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MacpoolPoolMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractIdPoolMember, errPoolAbstractIdPoolMember := json.Marshal(o.PoolAbstractIdPoolMember)
	if errPoolAbstractIdPoolMember != nil {
		return map[string]interface{}{}, errPoolAbstractIdPoolMember
	}
	errPoolAbstractIdPoolMember = json.Unmarshal([]byte(serializedPoolAbstractIdPoolMember), &toSerialize)
	if errPoolAbstractIdPoolMember != nil {
		return map[string]interface{}{}, errPoolAbstractIdPoolMember
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.MacAddress) {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.AssignedToEntity.IsSet() {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity.Get()
	}
	if o.BlockHead.IsSet() {
		toSerialize["BlockHead"] = o.BlockHead.Get()
	}
	if o.Peer.IsSet() {
		toSerialize["Peer"] = o.Peer.Get()
	}
	if o.Pool.IsSet() {
		toSerialize["Pool"] = o.Pool.Get()
	}
	if o.Reservation.IsSet() {
		toSerialize["Reservation"] = o.Reservation.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MacpoolPoolMember) UnmarshalJSON(data []byte) (err error) {
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
	type MacpoolPoolMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// MAC Address of this pool member.
		MacAddress       *string                                `json:"MacAddress,omitempty"`
		AssignedToEntity NullableMoBaseMoRelationship           `json:"AssignedToEntity,omitempty"`
		BlockHead        NullableMacpoolIdBlockRelationship     `json:"BlockHead,omitempty"`
		Peer             NullableMacpoolLeaseRelationship       `json:"Peer,omitempty"`
		Pool             NullableMacpoolPoolRelationship        `json:"Pool,omitempty"`
		Reservation      NullableMacpoolReservationRelationship `json:"Reservation,omitempty"`
	}

	varMacpoolPoolMemberWithoutEmbeddedStruct := MacpoolPoolMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMacpoolPoolMemberWithoutEmbeddedStruct)
	if err == nil {
		varMacpoolPoolMember := _MacpoolPoolMember{}
		varMacpoolPoolMember.ClassId = varMacpoolPoolMemberWithoutEmbeddedStruct.ClassId
		varMacpoolPoolMember.ObjectType = varMacpoolPoolMemberWithoutEmbeddedStruct.ObjectType
		varMacpoolPoolMember.MacAddress = varMacpoolPoolMemberWithoutEmbeddedStruct.MacAddress
		varMacpoolPoolMember.AssignedToEntity = varMacpoolPoolMemberWithoutEmbeddedStruct.AssignedToEntity
		varMacpoolPoolMember.BlockHead = varMacpoolPoolMemberWithoutEmbeddedStruct.BlockHead
		varMacpoolPoolMember.Peer = varMacpoolPoolMemberWithoutEmbeddedStruct.Peer
		varMacpoolPoolMember.Pool = varMacpoolPoolMemberWithoutEmbeddedStruct.Pool
		varMacpoolPoolMember.Reservation = varMacpoolPoolMemberWithoutEmbeddedStruct.Reservation
		*o = MacpoolPoolMember(varMacpoolPoolMember)
	} else {
		return err
	}

	varMacpoolPoolMember := _MacpoolPoolMember{}

	err = json.Unmarshal(data, &varMacpoolPoolMember)
	if err == nil {
		o.PoolAbstractIdPoolMember = varMacpoolPoolMember.PoolAbstractIdPoolMember
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "BlockHead")
		delete(additionalProperties, "Peer")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "Reservation")

		// remove fields from embedded structs
		reflectPoolAbstractIdPoolMember := reflect.ValueOf(o.PoolAbstractIdPoolMember)
		for i := 0; i < reflectPoolAbstractIdPoolMember.Type().NumField(); i++ {
			t := reflectPoolAbstractIdPoolMember.Type().Field(i)

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

type NullableMacpoolPoolMember struct {
	value *MacpoolPoolMember
	isSet bool
}

func (v NullableMacpoolPoolMember) Get() *MacpoolPoolMember {
	return v.value
}

func (v *NullableMacpoolPoolMember) Set(val *MacpoolPoolMember) {
	v.value = val
	v.isSet = true
}

func (v NullableMacpoolPoolMember) IsSet() bool {
	return v.isSet
}

func (v *NullableMacpoolPoolMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacpoolPoolMember(val *MacpoolPoolMember) *NullableMacpoolPoolMember {
	return &NullableMacpoolPoolMember{value: val, isSet: true}
}

func (v NullableMacpoolPoolMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacpoolPoolMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
