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

// checks if the FcpoolLease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FcpoolLease{}

// FcpoolLease Lease represents a single WWN ID that is part of the universe, allocated either from a pool or through static assignment.
type FcpoolLease struct {
	PoolAbstractLease
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Purpose of this WWN pool.
	PoolPurpose *string `json:"PoolPurpose,omitempty"`
	// The preferred WWN ID address can be specified only for dynamic lease requests. Intersight will make its best  effort to allocate that WWN ID address if it is available in the pool. If the specified preferred WWN ID address  is not in the range of the pool or if it is already leased or reserved, then the next available WWN ID address  from the pool will be leased. Since this feature is specific to dynamic lease requests only, static lease  request will fail if it specifies the preferred WWN ID address property. When the preferred WWN ID address  property is specified in conjunction with 'migrate' property, existing static or dynamic lease will be  replaced by the new lease. Migration also supported only for dynamic lease requests.
	PreferredWwnId *string                     `json:"PreferredWwnId,omitempty" validate:"regexp=^$|((^20|5[0-9a-fA-F]{1}):([0-9a-fA-F]{2}:){6}([0-9a-fA-F]{2})$)"`
	Reservation    *FcpoolReservationReference `json:"Reservation,omitempty"`
	// WWN ID allocated for pool based allocation.
	WwnId                *string                              `json:"WwnId,omitempty" validate:"regexp=^$|((^20|5[0-9a-fA-F]{1}):([0-9a-fA-F]{2}:){6}([0-9a-fA-F]{2})$)"`
	AssignedToEntity     NullableMoBaseMoRelationship         `json:"AssignedToEntity,omitempty"`
	Pool                 NullableFcpoolPoolRelationship       `json:"Pool,omitempty"`
	PoolMember           NullableFcpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
	Universe             NullableFcpoolUniverseRelationship   `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcpoolLease FcpoolLease

// NewFcpoolLease instantiates a new FcpoolLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpoolLease(classId string, objectType string) *FcpoolLease {
	this := FcpoolLease{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allocationType string = "dynamic"
	this.AllocationType = &allocationType
	var hasDuplicate bool = false
	this.HasDuplicate = &hasDuplicate
	var migrate bool = false
	this.Migrate = &migrate
	return &this
}

// NewFcpoolLeaseWithDefaults instantiates a new FcpoolLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpoolLeaseWithDefaults() *FcpoolLease {
	this := FcpoolLease{}
	var classId string = "fcpool.Lease"
	this.ClassId = classId
	var objectType string = "fcpool.Lease"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FcpoolLease) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FcpoolLease) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FcpoolLease) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fcpool.Lease" of the ClassId field.
func (o *FcpoolLease) GetDefaultClassId() interface{} {
	return "fcpool.Lease"
}

// GetObjectType returns the ObjectType field value
func (o *FcpoolLease) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FcpoolLease) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FcpoolLease) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fcpool.Lease" of the ObjectType field.
func (o *FcpoolLease) GetDefaultObjectType() interface{} {
	return "fcpool.Lease"
}

// GetPoolPurpose returns the PoolPurpose field value if set, zero value otherwise.
func (o *FcpoolLease) GetPoolPurpose() string {
	if o == nil || IsNil(o.PoolPurpose) {
		var ret string
		return ret
	}
	return *o.PoolPurpose
}

// GetPoolPurposeOk returns a tuple with the PoolPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLease) GetPoolPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.PoolPurpose) {
		return nil, false
	}
	return o.PoolPurpose, true
}

// HasPoolPurpose returns a boolean if a field has been set.
func (o *FcpoolLease) HasPoolPurpose() bool {
	if o != nil && !IsNil(o.PoolPurpose) {
		return true
	}

	return false
}

// SetPoolPurpose gets a reference to the given string and assigns it to the PoolPurpose field.
func (o *FcpoolLease) SetPoolPurpose(v string) {
	o.PoolPurpose = &v
}

// GetPreferredWwnId returns the PreferredWwnId field value if set, zero value otherwise.
func (o *FcpoolLease) GetPreferredWwnId() string {
	if o == nil || IsNil(o.PreferredWwnId) {
		var ret string
		return ret
	}
	return *o.PreferredWwnId
}

// GetPreferredWwnIdOk returns a tuple with the PreferredWwnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLease) GetPreferredWwnIdOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredWwnId) {
		return nil, false
	}
	return o.PreferredWwnId, true
}

// HasPreferredWwnId returns a boolean if a field has been set.
func (o *FcpoolLease) HasPreferredWwnId() bool {
	if o != nil && !IsNil(o.PreferredWwnId) {
		return true
	}

	return false
}

// SetPreferredWwnId gets a reference to the given string and assigns it to the PreferredWwnId field.
func (o *FcpoolLease) SetPreferredWwnId(v string) {
	o.PreferredWwnId = &v
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *FcpoolLease) GetReservation() FcpoolReservationReference {
	if o == nil || IsNil(o.Reservation) {
		var ret FcpoolReservationReference
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLease) GetReservationOk() (*FcpoolReservationReference, bool) {
	if o == nil || IsNil(o.Reservation) {
		return nil, false
	}
	return o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *FcpoolLease) HasReservation() bool {
	if o != nil && !IsNil(o.Reservation) {
		return true
	}

	return false
}

// SetReservation gets a reference to the given FcpoolReservationReference and assigns it to the Reservation field.
func (o *FcpoolLease) SetReservation(v FcpoolReservationReference) {
	o.Reservation = &v
}

// GetWwnId returns the WwnId field value if set, zero value otherwise.
func (o *FcpoolLease) GetWwnId() string {
	if o == nil || IsNil(o.WwnId) {
		var ret string
		return ret
	}
	return *o.WwnId
}

// GetWwnIdOk returns a tuple with the WwnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLease) GetWwnIdOk() (*string, bool) {
	if o == nil || IsNil(o.WwnId) {
		return nil, false
	}
	return o.WwnId, true
}

// HasWwnId returns a boolean if a field has been set.
func (o *FcpoolLease) HasWwnId() bool {
	if o != nil && !IsNil(o.WwnId) {
		return true
	}

	return false
}

// SetWwnId gets a reference to the given string and assigns it to the WwnId field.
func (o *FcpoolLease) SetWwnId(v string) {
	o.WwnId = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolLease) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || IsNil(o.AssignedToEntity.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity.Get()
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedToEntity.Get(), o.AssignedToEntity.IsSet()
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *FcpoolLease) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity.IsSet() {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given NullableMoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *FcpoolLease) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity.Set(&v)
}

// SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil
func (o *FcpoolLease) SetAssignedToEntityNil() {
	o.AssignedToEntity.Set(nil)
}

// UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
func (o *FcpoolLease) UnsetAssignedToEntity() {
	o.AssignedToEntity.Unset()
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolLease) GetPool() FcpoolPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret FcpoolPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolLease) GetPoolOk() (*FcpoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *FcpoolLease) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableFcpoolPoolRelationship and assigns it to the Pool field.
func (o *FcpoolLease) SetPool(v FcpoolPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *FcpoolLease) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *FcpoolLease) UnsetPool() {
	o.Pool.Unset()
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolLease) GetPoolMember() FcpoolPoolMemberRelationship {
	if o == nil || IsNil(o.PoolMember.Get()) {
		var ret FcpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember.Get()
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolLease) GetPoolMemberOk() (*FcpoolPoolMemberRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoolMember.Get(), o.PoolMember.IsSet()
}

// HasPoolMember returns a boolean if a field has been set.
func (o *FcpoolLease) HasPoolMember() bool {
	if o != nil && o.PoolMember.IsSet() {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given NullableFcpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *FcpoolLease) SetPoolMember(v FcpoolPoolMemberRelationship) {
	o.PoolMember.Set(&v)
}

// SetPoolMemberNil sets the value for PoolMember to be an explicit nil
func (o *FcpoolLease) SetPoolMemberNil() {
	o.PoolMember.Set(nil)
}

// UnsetPoolMember ensures that no value is present for PoolMember, not even an explicit nil
func (o *FcpoolLease) UnsetPoolMember() {
	o.PoolMember.Unset()
}

// GetUniverse returns the Universe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolLease) GetUniverse() FcpoolUniverseRelationship {
	if o == nil || IsNil(o.Universe.Get()) {
		var ret FcpoolUniverseRelationship
		return ret
	}
	return *o.Universe.Get()
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolLease) GetUniverseOk() (*FcpoolUniverseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Universe.Get(), o.Universe.IsSet()
}

// HasUniverse returns a boolean if a field has been set.
func (o *FcpoolLease) HasUniverse() bool {
	if o != nil && o.Universe.IsSet() {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given NullableFcpoolUniverseRelationship and assigns it to the Universe field.
func (o *FcpoolLease) SetUniverse(v FcpoolUniverseRelationship) {
	o.Universe.Set(&v)
}

// SetUniverseNil sets the value for Universe to be an explicit nil
func (o *FcpoolLease) SetUniverseNil() {
	o.Universe.Set(nil)
}

// UnsetUniverse ensures that no value is present for Universe, not even an explicit nil
func (o *FcpoolLease) UnsetUniverse() {
	o.Universe.Unset()
}

func (o FcpoolLease) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FcpoolLease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractLease, errPoolAbstractLease := json.Marshal(o.PoolAbstractLease)
	if errPoolAbstractLease != nil {
		return map[string]interface{}{}, errPoolAbstractLease
	}
	errPoolAbstractLease = json.Unmarshal([]byte(serializedPoolAbstractLease), &toSerialize)
	if errPoolAbstractLease != nil {
		return map[string]interface{}{}, errPoolAbstractLease
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.PoolPurpose) {
		toSerialize["PoolPurpose"] = o.PoolPurpose
	}
	if !IsNil(o.PreferredWwnId) {
		toSerialize["PreferredWwnId"] = o.PreferredWwnId
	}
	if !IsNil(o.Reservation) {
		toSerialize["Reservation"] = o.Reservation
	}
	if !IsNil(o.WwnId) {
		toSerialize["WwnId"] = o.WwnId
	}
	if o.AssignedToEntity.IsSet() {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity.Get()
	}
	if o.Pool.IsSet() {
		toSerialize["Pool"] = o.Pool.Get()
	}
	if o.PoolMember.IsSet() {
		toSerialize["PoolMember"] = o.PoolMember.Get()
	}
	if o.Universe.IsSet() {
		toSerialize["Universe"] = o.Universe.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FcpoolLease) UnmarshalJSON(data []byte) (err error) {
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
	type FcpoolLeaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Purpose of this WWN pool.
		PoolPurpose *string `json:"PoolPurpose,omitempty"`
		// The preferred WWN ID address can be specified only for dynamic lease requests. Intersight will make its best  effort to allocate that WWN ID address if it is available in the pool. If the specified preferred WWN ID address  is not in the range of the pool or if it is already leased or reserved, then the next available WWN ID address  from the pool will be leased. Since this feature is specific to dynamic lease requests only, static lease  request will fail if it specifies the preferred WWN ID address property. When the preferred WWN ID address  property is specified in conjunction with 'migrate' property, existing static or dynamic lease will be  replaced by the new lease. Migration also supported only for dynamic lease requests.
		PreferredWwnId *string                     `json:"PreferredWwnId,omitempty" validate:"regexp=^$|((^20|5[0-9a-fA-F]{1}):([0-9a-fA-F]{2}:){6}([0-9a-fA-F]{2})$)"`
		Reservation    *FcpoolReservationReference `json:"Reservation,omitempty"`
		// WWN ID allocated for pool based allocation.
		WwnId            *string                              `json:"WwnId,omitempty" validate:"regexp=^$|((^20|5[0-9a-fA-F]{1}):([0-9a-fA-F]{2}:){6}([0-9a-fA-F]{2})$)"`
		AssignedToEntity NullableMoBaseMoRelationship         `json:"AssignedToEntity,omitempty"`
		Pool             NullableFcpoolPoolRelationship       `json:"Pool,omitempty"`
		PoolMember       NullableFcpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
		Universe         NullableFcpoolUniverseRelationship   `json:"Universe,omitempty"`
	}

	varFcpoolLeaseWithoutEmbeddedStruct := FcpoolLeaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFcpoolLeaseWithoutEmbeddedStruct)
	if err == nil {
		varFcpoolLease := _FcpoolLease{}
		varFcpoolLease.ClassId = varFcpoolLeaseWithoutEmbeddedStruct.ClassId
		varFcpoolLease.ObjectType = varFcpoolLeaseWithoutEmbeddedStruct.ObjectType
		varFcpoolLease.PoolPurpose = varFcpoolLeaseWithoutEmbeddedStruct.PoolPurpose
		varFcpoolLease.PreferredWwnId = varFcpoolLeaseWithoutEmbeddedStruct.PreferredWwnId
		varFcpoolLease.Reservation = varFcpoolLeaseWithoutEmbeddedStruct.Reservation
		varFcpoolLease.WwnId = varFcpoolLeaseWithoutEmbeddedStruct.WwnId
		varFcpoolLease.AssignedToEntity = varFcpoolLeaseWithoutEmbeddedStruct.AssignedToEntity
		varFcpoolLease.Pool = varFcpoolLeaseWithoutEmbeddedStruct.Pool
		varFcpoolLease.PoolMember = varFcpoolLeaseWithoutEmbeddedStruct.PoolMember
		varFcpoolLease.Universe = varFcpoolLeaseWithoutEmbeddedStruct.Universe
		*o = FcpoolLease(varFcpoolLease)
	} else {
		return err
	}

	varFcpoolLease := _FcpoolLease{}

	err = json.Unmarshal(data, &varFcpoolLease)
	if err == nil {
		o.PoolAbstractLease = varFcpoolLease.PoolAbstractLease
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PoolPurpose")
		delete(additionalProperties, "PreferredWwnId")
		delete(additionalProperties, "Reservation")
		delete(additionalProperties, "WwnId")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")

		// remove fields from embedded structs
		reflectPoolAbstractLease := reflect.ValueOf(o.PoolAbstractLease)
		for i := 0; i < reflectPoolAbstractLease.Type().NumField(); i++ {
			t := reflectPoolAbstractLease.Type().Field(i)

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

type NullableFcpoolLease struct {
	value *FcpoolLease
	isSet bool
}

func (v NullableFcpoolLease) Get() *FcpoolLease {
	return v.value
}

func (v *NullableFcpoolLease) Set(val *FcpoolLease) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolLease) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolLease(val *FcpoolLease) *NullableFcpoolLease {
	return &NullableFcpoolLease{value: val, isSet: true}
}

func (v NullableFcpoolLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
