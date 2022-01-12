# HyperflexHxZoneInfoDt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxZoneInfoDt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxZoneInfoDt"]
**NumNodes** | Pointer to **int64** | Number of nodes in the Zone. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Zone UUID for this HyperFlex node. | [optional] [readonly] 

## Methods

### NewHyperflexHxZoneInfoDt

`func NewHyperflexHxZoneInfoDt(classId string, objectType string, ) *HyperflexHxZoneInfoDt`

NewHyperflexHxZoneInfoDt instantiates a new HyperflexHxZoneInfoDt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxZoneInfoDtWithDefaults

`func NewHyperflexHxZoneInfoDtWithDefaults() *HyperflexHxZoneInfoDt`

NewHyperflexHxZoneInfoDtWithDefaults instantiates a new HyperflexHxZoneInfoDt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxZoneInfoDt) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxZoneInfoDt) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxZoneInfoDt) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxZoneInfoDt) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxZoneInfoDt) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxZoneInfoDt) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNumNodes

`func (o *HyperflexHxZoneInfoDt) GetNumNodes() int64`

GetNumNodes returns the NumNodes field if non-nil, zero value otherwise.

### GetNumNodesOk

`func (o *HyperflexHxZoneInfoDt) GetNumNodesOk() (*int64, bool)`

GetNumNodesOk returns a tuple with the NumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodes

`func (o *HyperflexHxZoneInfoDt) SetNumNodes(v int64)`

SetNumNodes sets NumNodes field to given value.

### HasNumNodes

`func (o *HyperflexHxZoneInfoDt) HasNumNodes() bool`

HasNumNodes returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexHxZoneInfoDt) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexHxZoneInfoDt) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexHxZoneInfoDt) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexHxZoneInfoDt) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


