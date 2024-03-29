# StorageNetAppNonDataIpInterfaceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNonDataIpInterfaceEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNonDataIpInterfaceEvent"]
**IpInterface** | Pointer to [**NullableStorageNetAppNonDataIpInterfaceRelationship**](StorageNetAppNonDataIpInterfaceRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNonDataIpInterfaceEvent

`func NewStorageNetAppNonDataIpInterfaceEvent(classId string, objectType string, ) *StorageNetAppNonDataIpInterfaceEvent`

NewStorageNetAppNonDataIpInterfaceEvent instantiates a new StorageNetAppNonDataIpInterfaceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNonDataIpInterfaceEventWithDefaults

`func NewStorageNetAppNonDataIpInterfaceEventWithDefaults() *StorageNetAppNonDataIpInterfaceEvent`

NewStorageNetAppNonDataIpInterfaceEventWithDefaults instantiates a new StorageNetAppNonDataIpInterfaceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNonDataIpInterfaceEvent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNonDataIpInterfaceEvent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNonDataIpInterfaceEvent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNonDataIpInterfaceEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNonDataIpInterfaceEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNonDataIpInterfaceEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpInterface

`func (o *StorageNetAppNonDataIpInterfaceEvent) GetIpInterface() StorageNetAppNonDataIpInterfaceRelationship`

GetIpInterface returns the IpInterface field if non-nil, zero value otherwise.

### GetIpInterfaceOk

`func (o *StorageNetAppNonDataIpInterfaceEvent) GetIpInterfaceOk() (*StorageNetAppNonDataIpInterfaceRelationship, bool)`

GetIpInterfaceOk returns a tuple with the IpInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInterface

`func (o *StorageNetAppNonDataIpInterfaceEvent) SetIpInterface(v StorageNetAppNonDataIpInterfaceRelationship)`

SetIpInterface sets IpInterface field to given value.

### HasIpInterface

`func (o *StorageNetAppNonDataIpInterfaceEvent) HasIpInterface() bool`

HasIpInterface returns a boolean if a field has been set.

### SetIpInterfaceNil

`func (o *StorageNetAppNonDataIpInterfaceEvent) SetIpInterfaceNil(b bool)`

 SetIpInterfaceNil sets the value for IpInterface to be an explicit nil

### UnsetIpInterface
`func (o *StorageNetAppNonDataIpInterfaceEvent) UnsetIpInterface()`

UnsetIpInterface ensures that no value is present for IpInterface, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


