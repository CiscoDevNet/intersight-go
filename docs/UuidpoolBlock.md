# UuidpoolBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "uuidpool.Block"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "uuidpool.Block"]
**Prefix** | Pointer to **string** | Prefix of the UUID pool. UUID is constructed as &lt;prefix&gt;-&lt;suffix&gt;. | [optional] [readonly] 
**UuidSuffixBlock** | Pointer to [**UuidpoolUuidBlock**](UuidpoolUuidBlock.md) |  | [optional] 
**Pool** | Pointer to [**NullableUuidpoolPoolRelationship**](UuidpoolPoolRelationship.md) |  | [optional] 
**Reservations** | Pointer to [**[]UuidpoolReservationRelationship**](UuidpoolReservationRelationship.md) | An array of relationships to uuidpoolReservation resources. | [optional] [readonly] 

## Methods

### NewUuidpoolBlock

`func NewUuidpoolBlock(classId string, objectType string, ) *UuidpoolBlock`

NewUuidpoolBlock instantiates a new UuidpoolBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUuidpoolBlockWithDefaults

`func NewUuidpoolBlockWithDefaults() *UuidpoolBlock`

NewUuidpoolBlockWithDefaults instantiates a new UuidpoolBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *UuidpoolBlock) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *UuidpoolBlock) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *UuidpoolBlock) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *UuidpoolBlock) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UuidpoolBlock) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UuidpoolBlock) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPrefix

`func (o *UuidpoolBlock) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *UuidpoolBlock) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *UuidpoolBlock) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *UuidpoolBlock) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetUuidSuffixBlock

`func (o *UuidpoolBlock) GetUuidSuffixBlock() UuidpoolUuidBlock`

GetUuidSuffixBlock returns the UuidSuffixBlock field if non-nil, zero value otherwise.

### GetUuidSuffixBlockOk

`func (o *UuidpoolBlock) GetUuidSuffixBlockOk() (*UuidpoolUuidBlock, bool)`

GetUuidSuffixBlockOk returns a tuple with the UuidSuffixBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidSuffixBlock

`func (o *UuidpoolBlock) SetUuidSuffixBlock(v UuidpoolUuidBlock)`

SetUuidSuffixBlock sets UuidSuffixBlock field to given value.

### HasUuidSuffixBlock

`func (o *UuidpoolBlock) HasUuidSuffixBlock() bool`

HasUuidSuffixBlock returns a boolean if a field has been set.

### GetPool

`func (o *UuidpoolBlock) GetPool() UuidpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UuidpoolBlock) GetPoolOk() (*UuidpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UuidpoolBlock) SetPool(v UuidpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UuidpoolBlock) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *UuidpoolBlock) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *UuidpoolBlock) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetReservations

`func (o *UuidpoolBlock) GetReservations() []UuidpoolReservationRelationship`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *UuidpoolBlock) GetReservationsOk() (*[]UuidpoolReservationRelationship, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *UuidpoolBlock) SetReservations(v []UuidpoolReservationRelationship)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *UuidpoolBlock) HasReservations() bool`

HasReservations returns a boolean if a field has been set.

### SetReservationsNil

`func (o *UuidpoolBlock) SetReservationsNil(b bool)`

 SetReservationsNil sets the value for Reservations to be an explicit nil

### UnsetReservations
`func (o *UuidpoolBlock) UnsetReservations()`

UnsetReservations ensures that no value is present for Reservations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


