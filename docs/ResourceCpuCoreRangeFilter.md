# ResourceCpuCoreRangeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.CpuCoreRangeFilter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.CpuCoreRangeFilter"]
**MaxValue** | Pointer to **int64** | Maximum value for the range limit. It should be greater than minimum value. | [optional] [default to 0]
**MinValue** | Pointer to **int64** | Minimum value for the range limit. | [optional] [default to 0]

## Methods

### NewResourceCpuCoreRangeFilter

`func NewResourceCpuCoreRangeFilter(classId string, objectType string, ) *ResourceCpuCoreRangeFilter`

NewResourceCpuCoreRangeFilter instantiates a new ResourceCpuCoreRangeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceCpuCoreRangeFilterWithDefaults

`func NewResourceCpuCoreRangeFilterWithDefaults() *ResourceCpuCoreRangeFilter`

NewResourceCpuCoreRangeFilterWithDefaults instantiates a new ResourceCpuCoreRangeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceCpuCoreRangeFilter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceCpuCoreRangeFilter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceCpuCoreRangeFilter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceCpuCoreRangeFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceCpuCoreRangeFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceCpuCoreRangeFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMaxValue

`func (o *ResourceCpuCoreRangeFilter) GetMaxValue() int64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ResourceCpuCoreRangeFilter) GetMaxValueOk() (*int64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ResourceCpuCoreRangeFilter) SetMaxValue(v int64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ResourceCpuCoreRangeFilter) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetMinValue

`func (o *ResourceCpuCoreRangeFilter) GetMinValue() int64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ResourceCpuCoreRangeFilter) GetMinValueOk() (*int64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ResourceCpuCoreRangeFilter) SetMinValue(v int64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ResourceCpuCoreRangeFilter) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

