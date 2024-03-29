# MemoryPersistentMemoryUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "memory.PersistentMemoryUnit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "memory.PersistentMemoryUnit"]
**AppDirectCapacity** | Pointer to **string** | AppDirect capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**CountStatus** | Pointer to **string** | Count status of the Persistent Memory Module on a server. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | Firmware version of the firware running on the Persistent Memory Module on a server. | [optional] [readonly] 
**FrozenStatus** | Pointer to **string** | Frozen status of the Persistent Memory Module on a server. | [optional] [readonly] 
**HealthState** | Pointer to **string** | Health state of the Persistent Memory Module on a server. | [optional] [readonly] 
**LockStatus** | Pointer to **string** | Lock status of the Persistent Memory Module on a server. | [optional] [readonly] 
**MemoryCapacity** | Pointer to **string** | Memory capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**MemoryId** | Pointer to **int64** | ID of the Persistent Memory Module on a server. | [optional] [readonly] 
**PersistentMemoryCapacity** | Pointer to **string** | Persistent Memory capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**ReservedCapacity** | Pointer to **string** | Reserved capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**SecurityStatus** | Pointer to **string** | Security status of the Persistent Memory Module on a server. | [optional] [readonly] 
**SocketId** | Pointer to **string** | Socket ID of the Persistent Memory Module on a server. | [optional] [readonly] 
**SocketMemoryId** | Pointer to **string** | Socket Memory ID of the Persistent Memory Module on a server. | [optional] [readonly] 
**TotalCapacity** | Pointer to **string** | Total capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**Uid** | Pointer to **string** | UID of the Persistent Memory Module on a server. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**MemoryArray** | Pointer to [**NullableMemoryArrayRelationship**](MemoryArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewMemoryPersistentMemoryUnit

`func NewMemoryPersistentMemoryUnit(classId string, objectType string, ) *MemoryPersistentMemoryUnit`

NewMemoryPersistentMemoryUnit instantiates a new MemoryPersistentMemoryUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryUnitWithDefaults

`func NewMemoryPersistentMemoryUnitWithDefaults() *MemoryPersistentMemoryUnit`

NewMemoryPersistentMemoryUnitWithDefaults instantiates a new MemoryPersistentMemoryUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MemoryPersistentMemoryUnit) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryPersistentMemoryUnit) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryPersistentMemoryUnit) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MemoryPersistentMemoryUnit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryPersistentMemoryUnit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryPersistentMemoryUnit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAppDirectCapacity

`func (o *MemoryPersistentMemoryUnit) GetAppDirectCapacity() string`

GetAppDirectCapacity returns the AppDirectCapacity field if non-nil, zero value otherwise.

### GetAppDirectCapacityOk

`func (o *MemoryPersistentMemoryUnit) GetAppDirectCapacityOk() (*string, bool)`

GetAppDirectCapacityOk returns a tuple with the AppDirectCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDirectCapacity

`func (o *MemoryPersistentMemoryUnit) SetAppDirectCapacity(v string)`

SetAppDirectCapacity sets AppDirectCapacity field to given value.

### HasAppDirectCapacity

`func (o *MemoryPersistentMemoryUnit) HasAppDirectCapacity() bool`

HasAppDirectCapacity returns a boolean if a field has been set.

### GetCountStatus

`func (o *MemoryPersistentMemoryUnit) GetCountStatus() string`

GetCountStatus returns the CountStatus field if non-nil, zero value otherwise.

### GetCountStatusOk

`func (o *MemoryPersistentMemoryUnit) GetCountStatusOk() (*string, bool)`

GetCountStatusOk returns a tuple with the CountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountStatus

`func (o *MemoryPersistentMemoryUnit) SetCountStatus(v string)`

SetCountStatus sets CountStatus field to given value.

### HasCountStatus

`func (o *MemoryPersistentMemoryUnit) HasCountStatus() bool`

HasCountStatus returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *MemoryPersistentMemoryUnit) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *MemoryPersistentMemoryUnit) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *MemoryPersistentMemoryUnit) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *MemoryPersistentMemoryUnit) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetFrozenStatus

`func (o *MemoryPersistentMemoryUnit) GetFrozenStatus() string`

GetFrozenStatus returns the FrozenStatus field if non-nil, zero value otherwise.

### GetFrozenStatusOk

`func (o *MemoryPersistentMemoryUnit) GetFrozenStatusOk() (*string, bool)`

GetFrozenStatusOk returns a tuple with the FrozenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenStatus

`func (o *MemoryPersistentMemoryUnit) SetFrozenStatus(v string)`

SetFrozenStatus sets FrozenStatus field to given value.

### HasFrozenStatus

`func (o *MemoryPersistentMemoryUnit) HasFrozenStatus() bool`

HasFrozenStatus returns a boolean if a field has been set.

### GetHealthState

`func (o *MemoryPersistentMemoryUnit) GetHealthState() string`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *MemoryPersistentMemoryUnit) GetHealthStateOk() (*string, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *MemoryPersistentMemoryUnit) SetHealthState(v string)`

SetHealthState sets HealthState field to given value.

### HasHealthState

`func (o *MemoryPersistentMemoryUnit) HasHealthState() bool`

HasHealthState returns a boolean if a field has been set.

### GetLockStatus

`func (o *MemoryPersistentMemoryUnit) GetLockStatus() string`

GetLockStatus returns the LockStatus field if non-nil, zero value otherwise.

### GetLockStatusOk

`func (o *MemoryPersistentMemoryUnit) GetLockStatusOk() (*string, bool)`

GetLockStatusOk returns a tuple with the LockStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockStatus

`func (o *MemoryPersistentMemoryUnit) SetLockStatus(v string)`

SetLockStatus sets LockStatus field to given value.

### HasLockStatus

`func (o *MemoryPersistentMemoryUnit) HasLockStatus() bool`

HasLockStatus returns a boolean if a field has been set.

### GetMemoryCapacity

`func (o *MemoryPersistentMemoryUnit) GetMemoryCapacity() string`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *MemoryPersistentMemoryUnit) GetMemoryCapacityOk() (*string, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *MemoryPersistentMemoryUnit) SetMemoryCapacity(v string)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *MemoryPersistentMemoryUnit) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### GetMemoryId

`func (o *MemoryPersistentMemoryUnit) GetMemoryId() int64`

GetMemoryId returns the MemoryId field if non-nil, zero value otherwise.

### GetMemoryIdOk

`func (o *MemoryPersistentMemoryUnit) GetMemoryIdOk() (*int64, bool)`

GetMemoryIdOk returns a tuple with the MemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryId

`func (o *MemoryPersistentMemoryUnit) SetMemoryId(v int64)`

SetMemoryId sets MemoryId field to given value.

### HasMemoryId

`func (o *MemoryPersistentMemoryUnit) HasMemoryId() bool`

HasMemoryId returns a boolean if a field has been set.

### GetPersistentMemoryCapacity

`func (o *MemoryPersistentMemoryUnit) GetPersistentMemoryCapacity() string`

GetPersistentMemoryCapacity returns the PersistentMemoryCapacity field if non-nil, zero value otherwise.

### GetPersistentMemoryCapacityOk

`func (o *MemoryPersistentMemoryUnit) GetPersistentMemoryCapacityOk() (*string, bool)`

GetPersistentMemoryCapacityOk returns a tuple with the PersistentMemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryCapacity

`func (o *MemoryPersistentMemoryUnit) SetPersistentMemoryCapacity(v string)`

SetPersistentMemoryCapacity sets PersistentMemoryCapacity field to given value.

### HasPersistentMemoryCapacity

`func (o *MemoryPersistentMemoryUnit) HasPersistentMemoryCapacity() bool`

HasPersistentMemoryCapacity returns a boolean if a field has been set.

### GetReservedCapacity

`func (o *MemoryPersistentMemoryUnit) GetReservedCapacity() string`

GetReservedCapacity returns the ReservedCapacity field if non-nil, zero value otherwise.

### GetReservedCapacityOk

`func (o *MemoryPersistentMemoryUnit) GetReservedCapacityOk() (*string, bool)`

GetReservedCapacityOk returns a tuple with the ReservedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCapacity

`func (o *MemoryPersistentMemoryUnit) SetReservedCapacity(v string)`

SetReservedCapacity sets ReservedCapacity field to given value.

### HasReservedCapacity

`func (o *MemoryPersistentMemoryUnit) HasReservedCapacity() bool`

HasReservedCapacity returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *MemoryPersistentMemoryUnit) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *MemoryPersistentMemoryUnit) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *MemoryPersistentMemoryUnit) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *MemoryPersistentMemoryUnit) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetSocketId

`func (o *MemoryPersistentMemoryUnit) GetSocketId() string`

GetSocketId returns the SocketId field if non-nil, zero value otherwise.

### GetSocketIdOk

`func (o *MemoryPersistentMemoryUnit) GetSocketIdOk() (*string, bool)`

GetSocketIdOk returns a tuple with the SocketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketId

`func (o *MemoryPersistentMemoryUnit) SetSocketId(v string)`

SetSocketId sets SocketId field to given value.

### HasSocketId

`func (o *MemoryPersistentMemoryUnit) HasSocketId() bool`

HasSocketId returns a boolean if a field has been set.

### GetSocketMemoryId

`func (o *MemoryPersistentMemoryUnit) GetSocketMemoryId() string`

GetSocketMemoryId returns the SocketMemoryId field if non-nil, zero value otherwise.

### GetSocketMemoryIdOk

`func (o *MemoryPersistentMemoryUnit) GetSocketMemoryIdOk() (*string, bool)`

GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketMemoryId

`func (o *MemoryPersistentMemoryUnit) SetSocketMemoryId(v string)`

SetSocketMemoryId sets SocketMemoryId field to given value.

### HasSocketMemoryId

`func (o *MemoryPersistentMemoryUnit) HasSocketMemoryId() bool`

HasSocketMemoryId returns a boolean if a field has been set.

### GetTotalCapacity

`func (o *MemoryPersistentMemoryUnit) GetTotalCapacity() string`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *MemoryPersistentMemoryUnit) GetTotalCapacityOk() (*string, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *MemoryPersistentMemoryUnit) SetTotalCapacity(v string)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *MemoryPersistentMemoryUnit) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetUid

`func (o *MemoryPersistentMemoryUnit) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MemoryPersistentMemoryUnit) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MemoryPersistentMemoryUnit) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MemoryPersistentMemoryUnit) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryUnit) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryUnit) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryUnit) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryUnit) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *MemoryPersistentMemoryUnit) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *MemoryPersistentMemoryUnit) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetMemoryArray

`func (o *MemoryPersistentMemoryUnit) GetMemoryArray() MemoryArrayRelationship`

GetMemoryArray returns the MemoryArray field if non-nil, zero value otherwise.

### GetMemoryArrayOk

`func (o *MemoryPersistentMemoryUnit) GetMemoryArrayOk() (*MemoryArrayRelationship, bool)`

GetMemoryArrayOk returns a tuple with the MemoryArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArray

`func (o *MemoryPersistentMemoryUnit) SetMemoryArray(v MemoryArrayRelationship)`

SetMemoryArray sets MemoryArray field to given value.

### HasMemoryArray

`func (o *MemoryPersistentMemoryUnit) HasMemoryArray() bool`

HasMemoryArray returns a boolean if a field has been set.

### SetMemoryArrayNil

`func (o *MemoryPersistentMemoryUnit) SetMemoryArrayNil(b bool)`

 SetMemoryArrayNil sets the value for MemoryArray to be an explicit nil

### UnsetMemoryArray
`func (o *MemoryPersistentMemoryUnit) UnsetMemoryArray()`

UnsetMemoryArray ensures that no value is present for MemoryArray, not even an explicit nil
### GetRegisteredDevice

`func (o *MemoryPersistentMemoryUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryUnit) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *MemoryPersistentMemoryUnit) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *MemoryPersistentMemoryUnit) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


