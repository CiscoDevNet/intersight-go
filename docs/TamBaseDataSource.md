# TamBaseDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Name** | Pointer to **string** | Name is used to unique identify and refer a given data source in an alert definition. | [optional] 
**Type** | Pointer to **string** | Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.). * &#x60;intersightApi&#x60; - Collector type for this data collection is Intersight APIs. * &#x60;nxos&#x60; - Collector type for this data collection is NXOS. * &#x60;s3File&#x60; - Collector type for this data collection is a file in a cloud hosted object storage bucket. | [optional] [default to "intersightApi"]

## Methods

### NewTamBaseDataSource

`func NewTamBaseDataSource(classId string, objectType string, ) *TamBaseDataSource`

NewTamBaseDataSource instantiates a new TamBaseDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamBaseDataSourceWithDefaults

`func NewTamBaseDataSourceWithDefaults() *TamBaseDataSource`

NewTamBaseDataSourceWithDefaults instantiates a new TamBaseDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamBaseDataSource) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamBaseDataSource) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamBaseDataSource) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamBaseDataSource) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamBaseDataSource) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamBaseDataSource) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *TamBaseDataSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TamBaseDataSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TamBaseDataSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TamBaseDataSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *TamBaseDataSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TamBaseDataSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TamBaseDataSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TamBaseDataSource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


