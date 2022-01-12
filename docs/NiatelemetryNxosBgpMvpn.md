# NiatelemetryNxosBgpMvpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NxosBgpMvpn"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NxosBgpMvpn"]
**CapablePeers** | Pointer to **int64** | Return count of BGP MVPN table capable peers. | [optional] 
**ConfiguredPeers** | Pointer to **int64** | Return count of BGP MVPN configured peers. | [optional] 
**MemoryUsed** | Pointer to **int64** | Return value of BGP MVPN memory used. | [optional] 
**NumberOfClusterLists** | Pointer to **int64** | Return value of BGP MVPN cluster list. | [optional] 
**NumberOfCommunities** | Pointer to **int64** | Return count of BGP MVPN communities. | [optional] 
**TableVersion** | Pointer to **int64** | Return value of BGP MVPN table version. | [optional] 
**TotalNetworks** | Pointer to **int64** | Return count of BGP MVPN networks. | [optional] 
**TotalPaths** | Pointer to **int64** | Return count of BGP MVPN paths. | [optional] 

## Methods

### NewNiatelemetryNxosBgpMvpn

`func NewNiatelemetryNxosBgpMvpn(classId string, objectType string, ) *NiatelemetryNxosBgpMvpn`

NewNiatelemetryNxosBgpMvpn instantiates a new NiatelemetryNxosBgpMvpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNxosBgpMvpnWithDefaults

`func NewNiatelemetryNxosBgpMvpnWithDefaults() *NiatelemetryNxosBgpMvpn`

NewNiatelemetryNxosBgpMvpnWithDefaults instantiates a new NiatelemetryNxosBgpMvpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNxosBgpMvpn) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNxosBgpMvpn) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNxosBgpMvpn) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNxosBgpMvpn) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNxosBgpMvpn) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNxosBgpMvpn) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapablePeers

`func (o *NiatelemetryNxosBgpMvpn) GetCapablePeers() int64`

GetCapablePeers returns the CapablePeers field if non-nil, zero value otherwise.

### GetCapablePeersOk

`func (o *NiatelemetryNxosBgpMvpn) GetCapablePeersOk() (*int64, bool)`

GetCapablePeersOk returns a tuple with the CapablePeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapablePeers

`func (o *NiatelemetryNxosBgpMvpn) SetCapablePeers(v int64)`

SetCapablePeers sets CapablePeers field to given value.

### HasCapablePeers

`func (o *NiatelemetryNxosBgpMvpn) HasCapablePeers() bool`

HasCapablePeers returns a boolean if a field has been set.

### GetConfiguredPeers

`func (o *NiatelemetryNxosBgpMvpn) GetConfiguredPeers() int64`

GetConfiguredPeers returns the ConfiguredPeers field if non-nil, zero value otherwise.

### GetConfiguredPeersOk

`func (o *NiatelemetryNxosBgpMvpn) GetConfiguredPeersOk() (*int64, bool)`

GetConfiguredPeersOk returns a tuple with the ConfiguredPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredPeers

`func (o *NiatelemetryNxosBgpMvpn) SetConfiguredPeers(v int64)`

SetConfiguredPeers sets ConfiguredPeers field to given value.

### HasConfiguredPeers

`func (o *NiatelemetryNxosBgpMvpn) HasConfiguredPeers() bool`

HasConfiguredPeers returns a boolean if a field has been set.

### GetMemoryUsed

`func (o *NiatelemetryNxosBgpMvpn) GetMemoryUsed() int64`

GetMemoryUsed returns the MemoryUsed field if non-nil, zero value otherwise.

### GetMemoryUsedOk

`func (o *NiatelemetryNxosBgpMvpn) GetMemoryUsedOk() (*int64, bool)`

GetMemoryUsedOk returns a tuple with the MemoryUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsed

`func (o *NiatelemetryNxosBgpMvpn) SetMemoryUsed(v int64)`

SetMemoryUsed sets MemoryUsed field to given value.

### HasMemoryUsed

`func (o *NiatelemetryNxosBgpMvpn) HasMemoryUsed() bool`

HasMemoryUsed returns a boolean if a field has been set.

### GetNumberOfClusterLists

`func (o *NiatelemetryNxosBgpMvpn) GetNumberOfClusterLists() int64`

GetNumberOfClusterLists returns the NumberOfClusterLists field if non-nil, zero value otherwise.

### GetNumberOfClusterListsOk

`func (o *NiatelemetryNxosBgpMvpn) GetNumberOfClusterListsOk() (*int64, bool)`

GetNumberOfClusterListsOk returns a tuple with the NumberOfClusterLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfClusterLists

`func (o *NiatelemetryNxosBgpMvpn) SetNumberOfClusterLists(v int64)`

SetNumberOfClusterLists sets NumberOfClusterLists field to given value.

### HasNumberOfClusterLists

`func (o *NiatelemetryNxosBgpMvpn) HasNumberOfClusterLists() bool`

HasNumberOfClusterLists returns a boolean if a field has been set.

### GetNumberOfCommunities

`func (o *NiatelemetryNxosBgpMvpn) GetNumberOfCommunities() int64`

GetNumberOfCommunities returns the NumberOfCommunities field if non-nil, zero value otherwise.

### GetNumberOfCommunitiesOk

`func (o *NiatelemetryNxosBgpMvpn) GetNumberOfCommunitiesOk() (*int64, bool)`

GetNumberOfCommunitiesOk returns a tuple with the NumberOfCommunities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCommunities

`func (o *NiatelemetryNxosBgpMvpn) SetNumberOfCommunities(v int64)`

SetNumberOfCommunities sets NumberOfCommunities field to given value.

### HasNumberOfCommunities

`func (o *NiatelemetryNxosBgpMvpn) HasNumberOfCommunities() bool`

HasNumberOfCommunities returns a boolean if a field has been set.

### GetTableVersion

`func (o *NiatelemetryNxosBgpMvpn) GetTableVersion() int64`

GetTableVersion returns the TableVersion field if non-nil, zero value otherwise.

### GetTableVersionOk

`func (o *NiatelemetryNxosBgpMvpn) GetTableVersionOk() (*int64, bool)`

GetTableVersionOk returns a tuple with the TableVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableVersion

`func (o *NiatelemetryNxosBgpMvpn) SetTableVersion(v int64)`

SetTableVersion sets TableVersion field to given value.

### HasTableVersion

`func (o *NiatelemetryNxosBgpMvpn) HasTableVersion() bool`

HasTableVersion returns a boolean if a field has been set.

### GetTotalNetworks

`func (o *NiatelemetryNxosBgpMvpn) GetTotalNetworks() int64`

GetTotalNetworks returns the TotalNetworks field if non-nil, zero value otherwise.

### GetTotalNetworksOk

`func (o *NiatelemetryNxosBgpMvpn) GetTotalNetworksOk() (*int64, bool)`

GetTotalNetworksOk returns a tuple with the TotalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetworks

`func (o *NiatelemetryNxosBgpMvpn) SetTotalNetworks(v int64)`

SetTotalNetworks sets TotalNetworks field to given value.

### HasTotalNetworks

`func (o *NiatelemetryNxosBgpMvpn) HasTotalNetworks() bool`

HasTotalNetworks returns a boolean if a field has been set.

### GetTotalPaths

`func (o *NiatelemetryNxosBgpMvpn) GetTotalPaths() int64`

GetTotalPaths returns the TotalPaths field if non-nil, zero value otherwise.

### GetTotalPathsOk

`func (o *NiatelemetryNxosBgpMvpn) GetTotalPathsOk() (*int64, bool)`

GetTotalPathsOk returns a tuple with the TotalPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaths

`func (o *NiatelemetryNxosBgpMvpn) SetTotalPaths(v int64)`

SetTotalPaths sets TotalPaths field to given value.

### HasTotalPaths

`func (o *NiatelemetryNxosBgpMvpn) HasTotalPaths() bool`

HasTotalPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


