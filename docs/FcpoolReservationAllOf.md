# FcpoolReservationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fcpool.Reservation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fcpool.Reservation"]
**IdPurpose** | Pointer to **string** | Purpose of this WWN ID. Purpose can be WWPN or WWNN. | [optional] 
**Identity** | Pointer to **string** | WWN ID that needs to be reserved. | [optional] 
**Block** | Pointer to [**FcpoolFcBlockRelationship**](FcpoolFcBlockRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Pool** | Pointer to [**FcpoolPoolRelationship**](FcpoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**FcpoolPoolMemberRelationship**](FcpoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**FcpoolUniverseRelationship**](FcpoolUniverseRelationship.md) |  | [optional] 

## Methods

### NewFcpoolReservationAllOf

`func NewFcpoolReservationAllOf(classId string, objectType string, ) *FcpoolReservationAllOf`

NewFcpoolReservationAllOf instantiates a new FcpoolReservationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpoolReservationAllOfWithDefaults

`func NewFcpoolReservationAllOfWithDefaults() *FcpoolReservationAllOf`

NewFcpoolReservationAllOfWithDefaults instantiates a new FcpoolReservationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcpoolReservationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcpoolReservationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcpoolReservationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcpoolReservationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcpoolReservationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcpoolReservationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdPurpose

`func (o *FcpoolReservationAllOf) GetIdPurpose() string`

GetIdPurpose returns the IdPurpose field if non-nil, zero value otherwise.

### GetIdPurposeOk

`func (o *FcpoolReservationAllOf) GetIdPurposeOk() (*string, bool)`

GetIdPurposeOk returns a tuple with the IdPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdPurpose

`func (o *FcpoolReservationAllOf) SetIdPurpose(v string)`

SetIdPurpose sets IdPurpose field to given value.

### HasIdPurpose

`func (o *FcpoolReservationAllOf) HasIdPurpose() bool`

HasIdPurpose returns a boolean if a field has been set.

### GetIdentity

`func (o *FcpoolReservationAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *FcpoolReservationAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *FcpoolReservationAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *FcpoolReservationAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetBlock

`func (o *FcpoolReservationAllOf) GetBlock() FcpoolFcBlockRelationship`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *FcpoolReservationAllOf) GetBlockOk() (*FcpoolFcBlockRelationship, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *FcpoolReservationAllOf) SetBlock(v FcpoolFcBlockRelationship)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *FcpoolReservationAllOf) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetOrganization

`func (o *FcpoolReservationAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FcpoolReservationAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FcpoolReservationAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FcpoolReservationAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPool

`func (o *FcpoolReservationAllOf) GetPool() FcpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *FcpoolReservationAllOf) GetPoolOk() (*FcpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *FcpoolReservationAllOf) SetPool(v FcpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *FcpoolReservationAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *FcpoolReservationAllOf) GetPoolMember() FcpoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *FcpoolReservationAllOf) GetPoolMemberOk() (*FcpoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *FcpoolReservationAllOf) SetPoolMember(v FcpoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *FcpoolReservationAllOf) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *FcpoolReservationAllOf) GetUniverse() FcpoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *FcpoolReservationAllOf) GetUniverseOk() (*FcpoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *FcpoolReservationAllOf) SetUniverse(v FcpoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *FcpoolReservationAllOf) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

