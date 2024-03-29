# IppoolPoolMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.PoolMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.PoolMember"]
**IpType** | Pointer to **string** | Type of the IP address requested. * &#x60;IPv4&#x60; - IP V4 address type requested. * &#x60;IPv6&#x60; - IP V6 address type requested. | [optional] [readonly] [default to "IPv4"]
**IpV4Address** | Pointer to **string** | IPv4 Address of this pool member. | [optional] [readonly] 
**IpV6Address** | Pointer to **string** | IPv6 Address of this pool member. | [optional] [readonly] 
**AssignedToEntity** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**IpBlock** | Pointer to [**NullableIppoolShadowBlockRelationship**](IppoolShadowBlockRelationship.md) |  | [optional] 
**Peer** | Pointer to [**NullableIppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) |  | [optional] 
**Pool** | Pointer to [**NullableIppoolShadowPoolRelationship**](IppoolShadowPoolRelationship.md) |  | [optional] 
**Reservation** | Pointer to [**NullableIppoolReservationRelationship**](IppoolReservationRelationship.md) |  | [optional] 

## Methods

### NewIppoolPoolMember

`func NewIppoolPoolMember(classId string, objectType string, ) *IppoolPoolMember`

NewIppoolPoolMember instantiates a new IppoolPoolMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolPoolMemberWithDefaults

`func NewIppoolPoolMemberWithDefaults() *IppoolPoolMember`

NewIppoolPoolMemberWithDefaults instantiates a new IppoolPoolMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolPoolMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolPoolMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolPoolMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolPoolMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolPoolMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolPoolMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpType

`func (o *IppoolPoolMember) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *IppoolPoolMember) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *IppoolPoolMember) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *IppoolPoolMember) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetIpV4Address

`func (o *IppoolPoolMember) GetIpV4Address() string`

GetIpV4Address returns the IpV4Address field if non-nil, zero value otherwise.

### GetIpV4AddressOk

`func (o *IppoolPoolMember) GetIpV4AddressOk() (*string, bool)`

GetIpV4AddressOk returns a tuple with the IpV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Address

`func (o *IppoolPoolMember) SetIpV4Address(v string)`

SetIpV4Address sets IpV4Address field to given value.

### HasIpV4Address

`func (o *IppoolPoolMember) HasIpV4Address() bool`

HasIpV4Address returns a boolean if a field has been set.

### GetIpV6Address

`func (o *IppoolPoolMember) GetIpV6Address() string`

GetIpV6Address returns the IpV6Address field if non-nil, zero value otherwise.

### GetIpV6AddressOk

`func (o *IppoolPoolMember) GetIpV6AddressOk() (*string, bool)`

GetIpV6AddressOk returns a tuple with the IpV6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Address

`func (o *IppoolPoolMember) SetIpV6Address(v string)`

SetIpV6Address sets IpV6Address field to given value.

### HasIpV6Address

`func (o *IppoolPoolMember) HasIpV6Address() bool`

HasIpV6Address returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *IppoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *IppoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *IppoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *IppoolPoolMember) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### SetAssignedToEntityNil

`func (o *IppoolPoolMember) SetAssignedToEntityNil(b bool)`

 SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil

### UnsetAssignedToEntity
`func (o *IppoolPoolMember) UnsetAssignedToEntity()`

UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
### GetIpBlock

`func (o *IppoolPoolMember) GetIpBlock() IppoolShadowBlockRelationship`

GetIpBlock returns the IpBlock field if non-nil, zero value otherwise.

### GetIpBlockOk

`func (o *IppoolPoolMember) GetIpBlockOk() (*IppoolShadowBlockRelationship, bool)`

GetIpBlockOk returns a tuple with the IpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlock

`func (o *IppoolPoolMember) SetIpBlock(v IppoolShadowBlockRelationship)`

SetIpBlock sets IpBlock field to given value.

### HasIpBlock

`func (o *IppoolPoolMember) HasIpBlock() bool`

HasIpBlock returns a boolean if a field has been set.

### SetIpBlockNil

`func (o *IppoolPoolMember) SetIpBlockNil(b bool)`

 SetIpBlockNil sets the value for IpBlock to be an explicit nil

### UnsetIpBlock
`func (o *IppoolPoolMember) UnsetIpBlock()`

UnsetIpBlock ensures that no value is present for IpBlock, not even an explicit nil
### GetPeer

`func (o *IppoolPoolMember) GetPeer() IppoolIpLeaseRelationship`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *IppoolPoolMember) GetPeerOk() (*IppoolIpLeaseRelationship, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *IppoolPoolMember) SetPeer(v IppoolIpLeaseRelationship)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *IppoolPoolMember) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### SetPeerNil

`func (o *IppoolPoolMember) SetPeerNil(b bool)`

 SetPeerNil sets the value for Peer to be an explicit nil

### UnsetPeer
`func (o *IppoolPoolMember) UnsetPeer()`

UnsetPeer ensures that no value is present for Peer, not even an explicit nil
### GetPool

`func (o *IppoolPoolMember) GetPool() IppoolShadowPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IppoolPoolMember) GetPoolOk() (*IppoolShadowPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IppoolPoolMember) SetPool(v IppoolShadowPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IppoolPoolMember) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *IppoolPoolMember) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *IppoolPoolMember) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetReservation

`func (o *IppoolPoolMember) GetReservation() IppoolReservationRelationship`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *IppoolPoolMember) GetReservationOk() (*IppoolReservationRelationship, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *IppoolPoolMember) SetReservation(v IppoolReservationRelationship)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *IppoolPoolMember) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### SetReservationNil

`func (o *IppoolPoolMember) SetReservationNil(b bool)`

 SetReservationNil sets the value for Reservation to be an explicit nil

### UnsetReservation
`func (o *IppoolPoolMember) UnsetReservation()`

UnsetReservation ensures that no value is present for Reservation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


