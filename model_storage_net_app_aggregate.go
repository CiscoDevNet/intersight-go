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

// checks if the StorageNetAppAggregate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppAggregate{}

// StorageNetAppAggregate NetApp aggregate is a collection of disks arranged into one or more RAID groups.
type StorageNetAppAggregate struct {
	StorageBaseDiskPool
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Storage disk type for NetApp aggregate. * `HDD` - Hard Disk Drive disk type. * `Hybrid` - Solid State Hard Disk Drive. * `Hybrid (Flash Pool)` - SSHD in a flash pool disk type. * `SSD` - Solid State Disk disk type. * `SSD (FabricPool)` - SSD in a flash pool disk type. * `VMDisk (SDS)` - Storage disk with Hard disk drive. * `VMDisk (FabricPool)` - Storage disk with Non-volatile random-access memory drives. * `LUN (FlexArray)` - LUN (FlexArray) disk type. * `Not Mapped` - Storage disk is not mapped.
	AggregateType         *string                                 `json:"AggregateType,omitempty"`
	AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage `json:"AvgPerformanceMetrics,omitempty"`
	CloudStorage          []string                                `json:"CloudStorage,omitempty"`
	// Data reduction ratio (logical_used / used).
	EfficiencyRatio *float32 `json:"EfficiencyRatio,omitempty"`
	// Unique identifier of NetApp Aggregate across data center.
	Key *string `json:"Key,omitempty"`
	// The node name for the aggregate.
	NodeName *string `json:"NodeName,omitempty"`
	// Size of the RAID group represented by number of disks in the group.
	RaidSize *int64 `json:"RaidSize,omitempty"`
	// The RAID configuration type. * `Unknown` - Default unknown RAID type. * `RAID0` - RAID0 splits (\"stripes\") data evenly across two or more disks, without parity information. * `RAID1` - RAID1 requires a minimum of two disks to work, and provides data redundancy and failover. * `RAID4` - RAID4 stripes block level data and dedicates a disk to parity. * `RAID5` - RAID5  distributes striping and parity at a block level. * `RAID6` - RAID6 level operates like RAID5 with distributed parity and striping. * `RAID10` - RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy. * `RAIDDP` - RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group. * `RAIDTEC` - With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group.
	RaidType *string `json:"RaidType,omitempty"`
	// Current state of the NetApp aggregate. * `Unknown` - Specifies that the aggregate is discovered, but the aggregate information is not yet retrieved by the Unified Manager server. * `Online` - Aggregate is ready and available. * `Onlining` - The state is transitioning online. * `Offline` - Aggregate is unavailable. * `Offlining` - The state is transitioning offline. * `Relocating` - The aggregate is being relocated. * `Restricted` - Limited operations (e.g., parity reconstruction) are allowed, but data access is not allowed. * `Failed` - The aggregate cannot be brought online. * `Inconsistent` - The aggregate has been marked corrupted; contact technical support. * `Unmounted` - The aggregate is not mounted.
	State *string `json:"State,omitempty"`
	// Uuid of NetApp Aggregate.
	Uuid            *string                               `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	ArrayController NullableStorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppAggregateEvent resources.
	Events               []StorageNetAppAggregateEventRelationship `json:"Events,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppAggregate StorageNetAppAggregate

// NewStorageNetAppAggregate instantiates a new StorageNetAppAggregate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppAggregate(classId string, objectType string) *StorageNetAppAggregate {
	this := StorageNetAppAggregate{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppAggregateWithDefaults instantiates a new StorageNetAppAggregate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppAggregateWithDefaults() *StorageNetAppAggregate {
	this := StorageNetAppAggregate{}
	var classId string = "storage.NetAppAggregate"
	this.ClassId = classId
	var objectType string = "storage.NetAppAggregate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppAggregate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppAggregate) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppAggregate" of the ClassId field.
func (o *StorageNetAppAggregate) GetDefaultClassId() interface{} {
	return "storage.NetAppAggregate"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppAggregate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppAggregate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppAggregate" of the ObjectType field.
func (o *StorageNetAppAggregate) GetDefaultObjectType() interface{} {
	return "storage.NetAppAggregate"
}

// GetAggregateType returns the AggregateType field value if set, zero value otherwise.
func (o *StorageNetAppAggregate) GetAggregateType() string {
	if o == nil || IsNil(o.AggregateType) {
		var ret string
		return ret
	}
	return *o.AggregateType
}

// GetAggregateTypeOk returns a tuple with the AggregateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregate) GetAggregateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AggregateType) {
		return nil, false
	}
	return o.AggregateType, true
}

// HasAggregateType returns a boolean if a field has been set.
func (o *StorageNetAppAggregate) HasAggregateType() bool {
	if o != nil && !IsNil(o.AggregateType) {
		return true
	}

	return false
}

// SetAggregateType gets a reference to the given string and assigns it to the AggregateType field.
func (o *StorageNetAppAggregate) SetAggregateType(v string) {
	o.AggregateType = &v
}

// GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field value if set, zero value otherwise.
func (o *StorageNetAppAggregate) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage {
	if o == nil || IsNil(o.AvgPerformanceMetrics) {
		var ret StorageNetAppPerformanceMetricsAverage
		return ret
	}
	return *o.AvgPerformanceMetrics
}

// GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregate) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool) {
	if o == nil || IsNil(o.AvgPerformanceMetrics) {
		return nil, false
	}
	return o.AvgPerformanceMetrics, true
}

// HasAvgPerformanceMetrics returns a boolean if a field has been set.
func (o *StorageNetAppAggregate) HasAvgPerformanceMetrics() bool {
	if o != nil && !IsNil(o.AvgPerformanceMetrics) {
		return true
	}

	return false
}

// SetAvgPerformanceMetrics gets a reference to the given StorageNetAppPerformanceMetricsAverage and assigns it to the AvgPerformanceMetrics field.
func (o *StorageNetAppAggregate) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage) {
	o.AvgPerformanceMetrics = &v
}

// GetCloudStorage returns the CloudStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppAggregate) GetCloudStorage() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CloudStorage
}

// GetCloudStorageOk returns a tuple with the CloudStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppAggregate) GetCloudStorageOk() ([]string, bool) {
	if o == nil || IsNil(o.CloudStorage) {
		return nil, false
	}
	return o.CloudStorage, true
}

// HasCloudStorage returns a boolean if a field has been set.
func (o *StorageNetAppAggregate) HasCloudStorage() bool {
	if o != nil && !IsNil(o.CloudStorage) {
		return true
	}

	return false
}

// SetCloudStorage gets a reference to the given []string and assigns it to the CloudStorage field.
func (o *StorageNetAppAggregate) SetCloudStorage(v []string) {
	o.CloudStorage = v
}

// GetEfficiencyRatio returns the EfficiencyRatio field value if set, zero value otherwise.
func (o *StorageNetAppAggregate) GetEfficiencyRatio() float32 {
	if o == nil || IsNil(o.EfficiencyRatio) {
		var ret float32
		return ret
	}
	return *o.EfficiencyRatio
}

// GetEfficiencyRatioOk returns a tuple with the EfficiencyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregate) GetEfficiencyRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.EfficiencyRatio) {
		return nil, false
	}
	return o.EfficiencyRatio, true
}

// HasEfficiencyRatio returns a boolean if a field has been set.
func (o *StorageNetAppAggregate) HasEfficiencyRatio() bool {
	if o != nil && !IsNil(o.EfficiencyRatio) {
		return true
	}

	return false
}

// SetEfficiencyRatio gets a reference to the given float32 and assigns it to the EfficiencyRatio field.
func (o *StorageNetAppAggregate) SetEfficiencyRatio(v float32) {
	o.EfficiencyRatio = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StorageNetAppAggregate) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregate) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StorageNetAppAggregate) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StorageNetAppAggregate) SetKey(v string) {
	o.Key = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *StorageNetAppAggregate) GetNodeName() string {
	if o == nil || IsNil(o.NodeName) {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregate) GetNodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.NodeName) {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *StorageNetAppAggregate) HasNodeName() bool {
	if o != nil && !IsNil(o.NodeName) {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *StorageNetAppAggregate) SetNodeName(v string) {
	o.NodeName = &v
}

// GetRaidSize returns the RaidSize field value if set, zero value otherwise.
func (o *StorageNetAppAggregate) GetRaidSize() int64 {
	if o == nil || IsNil(o.RaidSize) {
		var ret int64
		return ret
	}
	return *o.RaidSize
}

// GetRaidSizeOk returns a tuple with the RaidSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregate) GetRaidSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.RaidSize) {
		return nil, false
	}
	return o.RaidSize, true
}

// HasRaidSize returns a boolean if a field has been set.
func (o *StorageNetAppAggregate) HasRaidSize() bool {
	if o != nil && !IsNil(o.RaidSize) {
		return true
	}

	return false
}

// SetRaidSize gets a reference to the given int64 and assigns it to the RaidSize field.
func (o *StorageNetAppAggregate) SetRaidSize(v int64) {
	o.RaidSize = &v
}

// GetRaidType returns the RaidType field value if set, zero value otherwise.
func (o *StorageNetAppAggregate) GetRaidType() string {
	if o == nil || IsNil(o.RaidType) {
		var ret string
		return ret
	}
	return *o.RaidType
}

// GetRaidTypeOk returns a tuple with the RaidType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregate) GetRaidTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RaidType) {
		return nil, false
	}
	return o.RaidType, true
}

// HasRaidType returns a boolean if a field has been set.
func (o *StorageNetAppAggregate) HasRaidType() bool {
	if o != nil && !IsNil(o.RaidType) {
		return true
	}

	return false
}

// SetRaidType gets a reference to the given string and assigns it to the RaidType field.
func (o *StorageNetAppAggregate) SetRaidType(v string) {
	o.RaidType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppAggregate) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregate) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppAggregate) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppAggregate) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppAggregate) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregate) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppAggregate) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppAggregate) SetUuid(v string) {
	o.Uuid = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppAggregate) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || IsNil(o.ArrayController.Get()) {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController.Get()
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppAggregate) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArrayController.Get(), o.ArrayController.IsSet()
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppAggregate) HasArrayController() bool {
	if o != nil && o.ArrayController.IsSet() {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given NullableStorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppAggregate) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController.Set(&v)
}

// SetArrayControllerNil sets the value for ArrayController to be an explicit nil
func (o *StorageNetAppAggregate) SetArrayControllerNil() {
	o.ArrayController.Set(nil)
}

// UnsetArrayController ensures that no value is present for ArrayController, not even an explicit nil
func (o *StorageNetAppAggregate) UnsetArrayController() {
	o.ArrayController.Unset()
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppAggregate) GetEvents() []StorageNetAppAggregateEventRelationship {
	if o == nil {
		var ret []StorageNetAppAggregateEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppAggregate) GetEventsOk() ([]StorageNetAppAggregateEventRelationship, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppAggregate) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppAggregateEventRelationship and assigns it to the Events field.
func (o *StorageNetAppAggregate) SetEvents(v []StorageNetAppAggregateEventRelationship) {
	o.Events = v
}

func (o StorageNetAppAggregate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppAggregate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseDiskPool, errStorageBaseDiskPool := json.Marshal(o.StorageBaseDiskPool)
	if errStorageBaseDiskPool != nil {
		return map[string]interface{}{}, errStorageBaseDiskPool
	}
	errStorageBaseDiskPool = json.Unmarshal([]byte(serializedStorageBaseDiskPool), &toSerialize)
	if errStorageBaseDiskPool != nil {
		return map[string]interface{}{}, errStorageBaseDiskPool
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AggregateType) {
		toSerialize["AggregateType"] = o.AggregateType
	}
	if !IsNil(o.AvgPerformanceMetrics) {
		toSerialize["AvgPerformanceMetrics"] = o.AvgPerformanceMetrics
	}
	if o.CloudStorage != nil {
		toSerialize["CloudStorage"] = o.CloudStorage
	}
	if !IsNil(o.EfficiencyRatio) {
		toSerialize["EfficiencyRatio"] = o.EfficiencyRatio
	}
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.NodeName) {
		toSerialize["NodeName"] = o.NodeName
	}
	if !IsNil(o.RaidSize) {
		toSerialize["RaidSize"] = o.RaidSize
	}
	if !IsNil(o.RaidType) {
		toSerialize["RaidType"] = o.RaidType
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.ArrayController.IsSet() {
		toSerialize["ArrayController"] = o.ArrayController.Get()
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppAggregate) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppAggregateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Storage disk type for NetApp aggregate. * `HDD` - Hard Disk Drive disk type. * `Hybrid` - Solid State Hard Disk Drive. * `Hybrid (Flash Pool)` - SSHD in a flash pool disk type. * `SSD` - Solid State Disk disk type. * `SSD (FabricPool)` - SSD in a flash pool disk type. * `VMDisk (SDS)` - Storage disk with Hard disk drive. * `VMDisk (FabricPool)` - Storage disk with Non-volatile random-access memory drives. * `LUN (FlexArray)` - LUN (FlexArray) disk type. * `Not Mapped` - Storage disk is not mapped.
		AggregateType         *string                                 `json:"AggregateType,omitempty"`
		AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage `json:"AvgPerformanceMetrics,omitempty"`
		CloudStorage          []string                                `json:"CloudStorage,omitempty"`
		// Data reduction ratio (logical_used / used).
		EfficiencyRatio *float32 `json:"EfficiencyRatio,omitempty"`
		// Unique identifier of NetApp Aggregate across data center.
		Key *string `json:"Key,omitempty"`
		// The node name for the aggregate.
		NodeName *string `json:"NodeName,omitempty"`
		// Size of the RAID group represented by number of disks in the group.
		RaidSize *int64 `json:"RaidSize,omitempty"`
		// The RAID configuration type. * `Unknown` - Default unknown RAID type. * `RAID0` - RAID0 splits (\"stripes\") data evenly across two or more disks, without parity information. * `RAID1` - RAID1 requires a minimum of two disks to work, and provides data redundancy and failover. * `RAID4` - RAID4 stripes block level data and dedicates a disk to parity. * `RAID5` - RAID5  distributes striping and parity at a block level. * `RAID6` - RAID6 level operates like RAID5 with distributed parity and striping. * `RAID10` - RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy. * `RAIDDP` - RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group. * `RAIDTEC` - With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group.
		RaidType *string `json:"RaidType,omitempty"`
		// Current state of the NetApp aggregate. * `Unknown` - Specifies that the aggregate is discovered, but the aggregate information is not yet retrieved by the Unified Manager server. * `Online` - Aggregate is ready and available. * `Onlining` - The state is transitioning online. * `Offline` - Aggregate is unavailable. * `Offlining` - The state is transitioning offline. * `Relocating` - The aggregate is being relocated. * `Restricted` - Limited operations (e.g., parity reconstruction) are allowed, but data access is not allowed. * `Failed` - The aggregate cannot be brought online. * `Inconsistent` - The aggregate has been marked corrupted; contact technical support. * `Unmounted` - The aggregate is not mounted.
		State *string `json:"State,omitempty"`
		// Uuid of NetApp Aggregate.
		Uuid            *string                               `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
		ArrayController NullableStorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
		// An array of relationships to storageNetAppAggregateEvent resources.
		Events []StorageNetAppAggregateEventRelationship `json:"Events,omitempty"`
	}

	varStorageNetAppAggregateWithoutEmbeddedStruct := StorageNetAppAggregateWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppAggregateWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppAggregate := _StorageNetAppAggregate{}
		varStorageNetAppAggregate.ClassId = varStorageNetAppAggregateWithoutEmbeddedStruct.ClassId
		varStorageNetAppAggregate.ObjectType = varStorageNetAppAggregateWithoutEmbeddedStruct.ObjectType
		varStorageNetAppAggregate.AggregateType = varStorageNetAppAggregateWithoutEmbeddedStruct.AggregateType
		varStorageNetAppAggregate.AvgPerformanceMetrics = varStorageNetAppAggregateWithoutEmbeddedStruct.AvgPerformanceMetrics
		varStorageNetAppAggregate.CloudStorage = varStorageNetAppAggregateWithoutEmbeddedStruct.CloudStorage
		varStorageNetAppAggregate.EfficiencyRatio = varStorageNetAppAggregateWithoutEmbeddedStruct.EfficiencyRatio
		varStorageNetAppAggregate.Key = varStorageNetAppAggregateWithoutEmbeddedStruct.Key
		varStorageNetAppAggregate.NodeName = varStorageNetAppAggregateWithoutEmbeddedStruct.NodeName
		varStorageNetAppAggregate.RaidSize = varStorageNetAppAggregateWithoutEmbeddedStruct.RaidSize
		varStorageNetAppAggregate.RaidType = varStorageNetAppAggregateWithoutEmbeddedStruct.RaidType
		varStorageNetAppAggregate.State = varStorageNetAppAggregateWithoutEmbeddedStruct.State
		varStorageNetAppAggregate.Uuid = varStorageNetAppAggregateWithoutEmbeddedStruct.Uuid
		varStorageNetAppAggregate.ArrayController = varStorageNetAppAggregateWithoutEmbeddedStruct.ArrayController
		varStorageNetAppAggregate.Events = varStorageNetAppAggregateWithoutEmbeddedStruct.Events
		*o = StorageNetAppAggregate(varStorageNetAppAggregate)
	} else {
		return err
	}

	varStorageNetAppAggregate := _StorageNetAppAggregate{}

	err = json.Unmarshal(data, &varStorageNetAppAggregate)
	if err == nil {
		o.StorageBaseDiskPool = varStorageNetAppAggregate.StorageBaseDiskPool
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AggregateType")
		delete(additionalProperties, "AvgPerformanceMetrics")
		delete(additionalProperties, "CloudStorage")
		delete(additionalProperties, "EfficiencyRatio")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "NodeName")
		delete(additionalProperties, "RaidSize")
		delete(additionalProperties, "RaidType")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "Events")

		// remove fields from embedded structs
		reflectStorageBaseDiskPool := reflect.ValueOf(o.StorageBaseDiskPool)
		for i := 0; i < reflectStorageBaseDiskPool.Type().NumField(); i++ {
			t := reflectStorageBaseDiskPool.Type().Field(i)

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

type NullableStorageNetAppAggregate struct {
	value *StorageNetAppAggregate
	isSet bool
}

func (v NullableStorageNetAppAggregate) Get() *StorageNetAppAggregate {
	return v.value
}

func (v *NullableStorageNetAppAggregate) Set(val *StorageNetAppAggregate) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppAggregate) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppAggregate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppAggregate(val *StorageNetAppAggregate) *NullableStorageNetAppAggregate {
	return &NullableStorageNetAppAggregate{value: val, isSet: true}
}

func (v NullableStorageNetAppAggregate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppAggregate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
