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

// checks if the HyperflexAlarm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexAlarm{}

// HyperflexAlarm An alarm representing a fault in the HyperFlex cluster configuration or hardware.
type HyperflexAlarm struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The acknowledgement state of the alarm. It is 'true' when the alarm is acknowledged and false otherwise.
	Acknowledged *bool `json:"Acknowledged,omitempty"`
	// The username of the user who acknowledged the alarm.
	AcknowledgedBy *string `json:"AcknowledgedBy,omitempty"`
	// The time when the alarm was acknowledged, represented as a Unix timestamp.
	AcknowledgedTime *int64 `json:"AcknowledgedTime,omitempty"`
	// The time when the alarm was acknowledged in ISO 6801 format.
	AcknowledgedTimeAsUtc *string `json:"AcknowledgedTimeAsUtc,omitempty"`
	// The description of the alarm which includes information about the fault that triggered the alarm.
	Description *string `json:"Description,omitempty"`
	// The data pertaining to this entity.
	EntityData *string `json:"EntityData,omitempty"`
	// The name of entity which triggered the alarm.
	EntityName *string `json:"EntityName,omitempty"`
	// The type of entity which triggered the alarm. For example, this can be the cluster, a node, or VM running on the cluster. * `UNKNOWN` - The type of entity is not known. * `DISK` - The entity is a physical storage device. * `NODE` - The entity is a HyperFlex cluster node. * `CLUSTER` - The entity is the HyperFlex cluster itself. * `DATASTORE` - The entity is a logical datastore configured on the HyperFlex cluster. * `ZONE` - The entity is a logical or physical zone configured on the HyperFlex cluster. * `VIRTUALMACHINE` - The entity is a virtual machine running on the HyperFlex cluster.
	EntityType *string `json:"EntityType,omitempty"`
	// The unique identifier of the entity which triggered the alarm.
	EntityUuId *string `json:"EntityUuId,omitempty"`
	// The localized message displayed to the user which describes the alarm.
	Message *string `json:"Message,omitempty"`
	// The name of the alarm. This name identifies the type of alarm that was triggered.
	Name *string `json:"Name,omitempty"`
	// The severity of the alarm. * `UNKNOWN` - The alarm status is not known. * `CLEARED` - The event that triggered the alarm has been remedied and no longer requires the user's attention. * `INFO` - The alarm represents a message that does not require the user's immediate attention. * `WARNING` - The alarm represents a moderate fault. * `CRITICAL` - The alarm represents a critical fault.
	Status *string `json:"Status,omitempty"`
	// The time when alarm was triggered as a Unix timestamp.
	TriggeredTime *int64 `json:"TriggeredTime,omitempty"`
	// The time when alarm was triggered in ISO 6801 UTC format.
	TriggeredTimeAsUtc *string `json:"TriggeredTimeAsUtc,omitempty"`
	// The unique identifier for this alarm instance.
	Uuid *string `json:"Uuid,omitempty"`
	// An array of relationships to infraBaseCluster resources.
	AncestorMos          []InfraBaseClusterRelationship       `json:"AncestorMos,omitempty"`
	Cluster              NullableHyperflexClusterRelationship `json:"Cluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexAlarm HyperflexAlarm

// NewHyperflexAlarm instantiates a new HyperflexAlarm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAlarm(classId string, objectType string) *HyperflexAlarm {
	this := HyperflexAlarm{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexAlarmWithDefaults instantiates a new HyperflexAlarm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAlarmWithDefaults() *HyperflexAlarm {
	this := HyperflexAlarm{}
	var classId string = "hyperflex.Alarm"
	this.ClassId = classId
	var objectType string = "hyperflex.Alarm"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexAlarm) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexAlarm) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.Alarm" of the ClassId field.
func (o *HyperflexAlarm) GetDefaultClassId() interface{} {
	return "hyperflex.Alarm"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexAlarm) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexAlarm) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.Alarm" of the ObjectType field.
func (o *HyperflexAlarm) GetDefaultObjectType() interface{} {
	return "hyperflex.Alarm"
}

// GetAcknowledged returns the Acknowledged field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetAcknowledged() bool {
	if o == nil || IsNil(o.Acknowledged) {
		var ret bool
		return ret
	}
	return *o.Acknowledged
}

// GetAcknowledgedOk returns a tuple with the Acknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.Acknowledged) {
		return nil, false
	}
	return o.Acknowledged, true
}

// HasAcknowledged returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasAcknowledged() bool {
	if o != nil && !IsNil(o.Acknowledged) {
		return true
	}

	return false
}

// SetAcknowledged gets a reference to the given bool and assigns it to the Acknowledged field.
func (o *HyperflexAlarm) SetAcknowledged(v bool) {
	o.Acknowledged = &v
}

// GetAcknowledgedBy returns the AcknowledgedBy field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetAcknowledgedBy() string {
	if o == nil || IsNil(o.AcknowledgedBy) {
		var ret string
		return ret
	}
	return *o.AcknowledgedBy
}

// GetAcknowledgedByOk returns a tuple with the AcknowledgedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetAcknowledgedByOk() (*string, bool) {
	if o == nil || IsNil(o.AcknowledgedBy) {
		return nil, false
	}
	return o.AcknowledgedBy, true
}

// HasAcknowledgedBy returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasAcknowledgedBy() bool {
	if o != nil && !IsNil(o.AcknowledgedBy) {
		return true
	}

	return false
}

// SetAcknowledgedBy gets a reference to the given string and assigns it to the AcknowledgedBy field.
func (o *HyperflexAlarm) SetAcknowledgedBy(v string) {
	o.AcknowledgedBy = &v
}

// GetAcknowledgedTime returns the AcknowledgedTime field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetAcknowledgedTime() int64 {
	if o == nil || IsNil(o.AcknowledgedTime) {
		var ret int64
		return ret
	}
	return *o.AcknowledgedTime
}

// GetAcknowledgedTimeOk returns a tuple with the AcknowledgedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetAcknowledgedTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.AcknowledgedTime) {
		return nil, false
	}
	return o.AcknowledgedTime, true
}

// HasAcknowledgedTime returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasAcknowledgedTime() bool {
	if o != nil && !IsNil(o.AcknowledgedTime) {
		return true
	}

	return false
}

// SetAcknowledgedTime gets a reference to the given int64 and assigns it to the AcknowledgedTime field.
func (o *HyperflexAlarm) SetAcknowledgedTime(v int64) {
	o.AcknowledgedTime = &v
}

// GetAcknowledgedTimeAsUtc returns the AcknowledgedTimeAsUtc field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetAcknowledgedTimeAsUtc() string {
	if o == nil || IsNil(o.AcknowledgedTimeAsUtc) {
		var ret string
		return ret
	}
	return *o.AcknowledgedTimeAsUtc
}

// GetAcknowledgedTimeAsUtcOk returns a tuple with the AcknowledgedTimeAsUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetAcknowledgedTimeAsUtcOk() (*string, bool) {
	if o == nil || IsNil(o.AcknowledgedTimeAsUtc) {
		return nil, false
	}
	return o.AcknowledgedTimeAsUtc, true
}

// HasAcknowledgedTimeAsUtc returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasAcknowledgedTimeAsUtc() bool {
	if o != nil && !IsNil(o.AcknowledgedTimeAsUtc) {
		return true
	}

	return false
}

// SetAcknowledgedTimeAsUtc gets a reference to the given string and assigns it to the AcknowledgedTimeAsUtc field.
func (o *HyperflexAlarm) SetAcknowledgedTimeAsUtc(v string) {
	o.AcknowledgedTimeAsUtc = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HyperflexAlarm) SetDescription(v string) {
	o.Description = &v
}

// GetEntityData returns the EntityData field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetEntityData() string {
	if o == nil || IsNil(o.EntityData) {
		var ret string
		return ret
	}
	return *o.EntityData
}

// GetEntityDataOk returns a tuple with the EntityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetEntityDataOk() (*string, bool) {
	if o == nil || IsNil(o.EntityData) {
		return nil, false
	}
	return o.EntityData, true
}

// HasEntityData returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasEntityData() bool {
	if o != nil && !IsNil(o.EntityData) {
		return true
	}

	return false
}

// SetEntityData gets a reference to the given string and assigns it to the EntityData field.
func (o *HyperflexAlarm) SetEntityData(v string) {
	o.EntityData = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetEntityName() string {
	if o == nil || IsNil(o.EntityName) {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetEntityNameOk() (*string, bool) {
	if o == nil || IsNil(o.EntityName) {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasEntityName() bool {
	if o != nil && !IsNil(o.EntityName) {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *HyperflexAlarm) SetEntityName(v string) {
	o.EntityName = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *HyperflexAlarm) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityUuId returns the EntityUuId field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetEntityUuId() string {
	if o == nil || IsNil(o.EntityUuId) {
		var ret string
		return ret
	}
	return *o.EntityUuId
}

// GetEntityUuIdOk returns a tuple with the EntityUuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetEntityUuIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityUuId) {
		return nil, false
	}
	return o.EntityUuId, true
}

// HasEntityUuId returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasEntityUuId() bool {
	if o != nil && !IsNil(o.EntityUuId) {
		return true
	}

	return false
}

// SetEntityUuId gets a reference to the given string and assigns it to the EntityUuId field.
func (o *HyperflexAlarm) SetEntityUuId(v string) {
	o.EntityUuId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HyperflexAlarm) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexAlarm) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexAlarm) SetStatus(v string) {
	o.Status = &v
}

// GetTriggeredTime returns the TriggeredTime field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetTriggeredTime() int64 {
	if o == nil || IsNil(o.TriggeredTime) {
		var ret int64
		return ret
	}
	return *o.TriggeredTime
}

// GetTriggeredTimeOk returns a tuple with the TriggeredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetTriggeredTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TriggeredTime) {
		return nil, false
	}
	return o.TriggeredTime, true
}

// HasTriggeredTime returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasTriggeredTime() bool {
	if o != nil && !IsNil(o.TriggeredTime) {
		return true
	}

	return false
}

// SetTriggeredTime gets a reference to the given int64 and assigns it to the TriggeredTime field.
func (o *HyperflexAlarm) SetTriggeredTime(v int64) {
	o.TriggeredTime = &v
}

// GetTriggeredTimeAsUtc returns the TriggeredTimeAsUtc field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetTriggeredTimeAsUtc() string {
	if o == nil || IsNil(o.TriggeredTimeAsUtc) {
		var ret string
		return ret
	}
	return *o.TriggeredTimeAsUtc
}

// GetTriggeredTimeAsUtcOk returns a tuple with the TriggeredTimeAsUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetTriggeredTimeAsUtcOk() (*string, bool) {
	if o == nil || IsNil(o.TriggeredTimeAsUtc) {
		return nil, false
	}
	return o.TriggeredTimeAsUtc, true
}

// HasTriggeredTimeAsUtc returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasTriggeredTimeAsUtc() bool {
	if o != nil && !IsNil(o.TriggeredTimeAsUtc) {
		return true
	}

	return false
}

// SetTriggeredTimeAsUtc gets a reference to the given string and assigns it to the TriggeredTimeAsUtc field.
func (o *HyperflexAlarm) SetTriggeredTimeAsUtc(v string) {
	o.TriggeredTimeAsUtc = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexAlarm) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarm) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexAlarm) SetUuid(v string) {
	o.Uuid = &v
}

// GetAncestorMos returns the AncestorMos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexAlarm) GetAncestorMos() []InfraBaseClusterRelationship {
	if o == nil {
		var ret []InfraBaseClusterRelationship
		return ret
	}
	return o.AncestorMos
}

// GetAncestorMosOk returns a tuple with the AncestorMos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexAlarm) GetAncestorMosOk() ([]InfraBaseClusterRelationship, bool) {
	if o == nil || IsNil(o.AncestorMos) {
		return nil, false
	}
	return o.AncestorMos, true
}

// HasAncestorMos returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasAncestorMos() bool {
	if o != nil && !IsNil(o.AncestorMos) {
		return true
	}

	return false
}

// SetAncestorMos gets a reference to the given []InfraBaseClusterRelationship and assigns it to the AncestorMos field.
func (o *HyperflexAlarm) SetAncestorMos(v []InfraBaseClusterRelationship) {
	o.AncestorMos = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexAlarm) GetCluster() HyperflexClusterRelationship {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexAlarm) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexAlarm) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableHyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexAlarm) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *HyperflexAlarm) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *HyperflexAlarm) UnsetCluster() {
	o.Cluster.Unset()
}

func (o HyperflexAlarm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexAlarm) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Acknowledged) {
		toSerialize["Acknowledged"] = o.Acknowledged
	}
	if !IsNil(o.AcknowledgedBy) {
		toSerialize["AcknowledgedBy"] = o.AcknowledgedBy
	}
	if !IsNil(o.AcknowledgedTime) {
		toSerialize["AcknowledgedTime"] = o.AcknowledgedTime
	}
	if !IsNil(o.AcknowledgedTimeAsUtc) {
		toSerialize["AcknowledgedTimeAsUtc"] = o.AcknowledgedTimeAsUtc
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.EntityData) {
		toSerialize["EntityData"] = o.EntityData
	}
	if !IsNil(o.EntityName) {
		toSerialize["EntityName"] = o.EntityName
	}
	if !IsNil(o.EntityType) {
		toSerialize["EntityType"] = o.EntityType
	}
	if !IsNil(o.EntityUuId) {
		toSerialize["EntityUuId"] = o.EntityUuId
	}
	if !IsNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.TriggeredTime) {
		toSerialize["TriggeredTime"] = o.TriggeredTime
	}
	if !IsNil(o.TriggeredTimeAsUtc) {
		toSerialize["TriggeredTimeAsUtc"] = o.TriggeredTimeAsUtc
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.AncestorMos != nil {
		toSerialize["AncestorMos"] = o.AncestorMos
	}
	if o.Cluster.IsSet() {
		toSerialize["Cluster"] = o.Cluster.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexAlarm) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexAlarmWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The acknowledgement state of the alarm. It is 'true' when the alarm is acknowledged and false otherwise.
		Acknowledged *bool `json:"Acknowledged,omitempty"`
		// The username of the user who acknowledged the alarm.
		AcknowledgedBy *string `json:"AcknowledgedBy,omitempty"`
		// The time when the alarm was acknowledged, represented as a Unix timestamp.
		AcknowledgedTime *int64 `json:"AcknowledgedTime,omitempty"`
		// The time when the alarm was acknowledged in ISO 6801 format.
		AcknowledgedTimeAsUtc *string `json:"AcknowledgedTimeAsUtc,omitempty"`
		// The description of the alarm which includes information about the fault that triggered the alarm.
		Description *string `json:"Description,omitempty"`
		// The data pertaining to this entity.
		EntityData *string `json:"EntityData,omitempty"`
		// The name of entity which triggered the alarm.
		EntityName *string `json:"EntityName,omitempty"`
		// The type of entity which triggered the alarm. For example, this can be the cluster, a node, or VM running on the cluster. * `UNKNOWN` - The type of entity is not known. * `DISK` - The entity is a physical storage device. * `NODE` - The entity is a HyperFlex cluster node. * `CLUSTER` - The entity is the HyperFlex cluster itself. * `DATASTORE` - The entity is a logical datastore configured on the HyperFlex cluster. * `ZONE` - The entity is a logical or physical zone configured on the HyperFlex cluster. * `VIRTUALMACHINE` - The entity is a virtual machine running on the HyperFlex cluster.
		EntityType *string `json:"EntityType,omitempty"`
		// The unique identifier of the entity which triggered the alarm.
		EntityUuId *string `json:"EntityUuId,omitempty"`
		// The localized message displayed to the user which describes the alarm.
		Message *string `json:"Message,omitempty"`
		// The name of the alarm. This name identifies the type of alarm that was triggered.
		Name *string `json:"Name,omitempty"`
		// The severity of the alarm. * `UNKNOWN` - The alarm status is not known. * `CLEARED` - The event that triggered the alarm has been remedied and no longer requires the user's attention. * `INFO` - The alarm represents a message that does not require the user's immediate attention. * `WARNING` - The alarm represents a moderate fault. * `CRITICAL` - The alarm represents a critical fault.
		Status *string `json:"Status,omitempty"`
		// The time when alarm was triggered as a Unix timestamp.
		TriggeredTime *int64 `json:"TriggeredTime,omitempty"`
		// The time when alarm was triggered in ISO 6801 UTC format.
		TriggeredTimeAsUtc *string `json:"TriggeredTimeAsUtc,omitempty"`
		// The unique identifier for this alarm instance.
		Uuid *string `json:"Uuid,omitempty"`
		// An array of relationships to infraBaseCluster resources.
		AncestorMos []InfraBaseClusterRelationship       `json:"AncestorMos,omitempty"`
		Cluster     NullableHyperflexClusterRelationship `json:"Cluster,omitempty"`
	}

	varHyperflexAlarmWithoutEmbeddedStruct := HyperflexAlarmWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexAlarmWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexAlarm := _HyperflexAlarm{}
		varHyperflexAlarm.ClassId = varHyperflexAlarmWithoutEmbeddedStruct.ClassId
		varHyperflexAlarm.ObjectType = varHyperflexAlarmWithoutEmbeddedStruct.ObjectType
		varHyperflexAlarm.Acknowledged = varHyperflexAlarmWithoutEmbeddedStruct.Acknowledged
		varHyperflexAlarm.AcknowledgedBy = varHyperflexAlarmWithoutEmbeddedStruct.AcknowledgedBy
		varHyperflexAlarm.AcknowledgedTime = varHyperflexAlarmWithoutEmbeddedStruct.AcknowledgedTime
		varHyperflexAlarm.AcknowledgedTimeAsUtc = varHyperflexAlarmWithoutEmbeddedStruct.AcknowledgedTimeAsUtc
		varHyperflexAlarm.Description = varHyperflexAlarmWithoutEmbeddedStruct.Description
		varHyperflexAlarm.EntityData = varHyperflexAlarmWithoutEmbeddedStruct.EntityData
		varHyperflexAlarm.EntityName = varHyperflexAlarmWithoutEmbeddedStruct.EntityName
		varHyperflexAlarm.EntityType = varHyperflexAlarmWithoutEmbeddedStruct.EntityType
		varHyperflexAlarm.EntityUuId = varHyperflexAlarmWithoutEmbeddedStruct.EntityUuId
		varHyperflexAlarm.Message = varHyperflexAlarmWithoutEmbeddedStruct.Message
		varHyperflexAlarm.Name = varHyperflexAlarmWithoutEmbeddedStruct.Name
		varHyperflexAlarm.Status = varHyperflexAlarmWithoutEmbeddedStruct.Status
		varHyperflexAlarm.TriggeredTime = varHyperflexAlarmWithoutEmbeddedStruct.TriggeredTime
		varHyperflexAlarm.TriggeredTimeAsUtc = varHyperflexAlarmWithoutEmbeddedStruct.TriggeredTimeAsUtc
		varHyperflexAlarm.Uuid = varHyperflexAlarmWithoutEmbeddedStruct.Uuid
		varHyperflexAlarm.AncestorMos = varHyperflexAlarmWithoutEmbeddedStruct.AncestorMos
		varHyperflexAlarm.Cluster = varHyperflexAlarmWithoutEmbeddedStruct.Cluster
		*o = HyperflexAlarm(varHyperflexAlarm)
	} else {
		return err
	}

	varHyperflexAlarm := _HyperflexAlarm{}

	err = json.Unmarshal(data, &varHyperflexAlarm)
	if err == nil {
		o.MoBaseMo = varHyperflexAlarm.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Acknowledged")
		delete(additionalProperties, "AcknowledgedBy")
		delete(additionalProperties, "AcknowledgedTime")
		delete(additionalProperties, "AcknowledgedTimeAsUtc")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "EntityData")
		delete(additionalProperties, "EntityName")
		delete(additionalProperties, "EntityType")
		delete(additionalProperties, "EntityUuId")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TriggeredTime")
		delete(additionalProperties, "TriggeredTimeAsUtc")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "AncestorMos")
		delete(additionalProperties, "Cluster")

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

type NullableHyperflexAlarm struct {
	value *HyperflexAlarm
	isSet bool
}

func (v NullableHyperflexAlarm) Get() *HyperflexAlarm {
	return v.value
}

func (v *NullableHyperflexAlarm) Set(val *HyperflexAlarm) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAlarm) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAlarm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAlarm(val *HyperflexAlarm) *NullableHyperflexAlarm {
	return &NullableHyperflexAlarm{value: val, isSet: true}
}

func (v NullableHyperflexAlarm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAlarm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
