# VirtualizationStorageCapacityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.StorageCapacity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.StorageCapacity"]
**Capacity** | Pointer to **int64** | The total capacity of the entity (bytes). | [optional] 
**Free** | Pointer to **int64** | Free storage space remaining in the entity (bytes) as a point-in-time snapshot. The available space is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used storage capacity is also reported. | [optional] 
**Used** | Pointer to **int64** | Spaced already used by this entity (bytes), as a point-in-time snapshot. | [optional] 

## Methods

### NewVirtualizationStorageCapacityAllOf

`func NewVirtualizationStorageCapacityAllOf(classId string, objectType string, ) *VirtualizationStorageCapacityAllOf`

NewVirtualizationStorageCapacityAllOf instantiates a new VirtualizationStorageCapacityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationStorageCapacityAllOfWithDefaults

`func NewVirtualizationStorageCapacityAllOfWithDefaults() *VirtualizationStorageCapacityAllOf`

NewVirtualizationStorageCapacityAllOfWithDefaults instantiates a new VirtualizationStorageCapacityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationStorageCapacityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationStorageCapacityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationStorageCapacityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationStorageCapacityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationStorageCapacityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationStorageCapacityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *VirtualizationStorageCapacityAllOf) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationStorageCapacityAllOf) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationStorageCapacityAllOf) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationStorageCapacityAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetFree

`func (o *VirtualizationStorageCapacityAllOf) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *VirtualizationStorageCapacityAllOf) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *VirtualizationStorageCapacityAllOf) SetFree(v int64)`

SetFree sets Free field to given value.

### HasFree

`func (o *VirtualizationStorageCapacityAllOf) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetUsed

`func (o *VirtualizationStorageCapacityAllOf) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *VirtualizationStorageCapacityAllOf) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *VirtualizationStorageCapacityAllOf) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *VirtualizationStorageCapacityAllOf) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


