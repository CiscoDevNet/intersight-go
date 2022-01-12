# FcpoolBlockAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fcpool.Block"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fcpool.Block"]
**From** | Pointer to **string** | Starting WWN identifier of the block must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN&#39;s in the SAN fabric, you are strongly encouraged to use the following WWN prefix; 20:00:00:25:B5:xx:xx:xx. | [optional] 
**To** | Pointer to **string** | Ending WWN identifier of the block must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. | [optional] 

## Methods

### NewFcpoolBlockAllOf

`func NewFcpoolBlockAllOf(classId string, objectType string, ) *FcpoolBlockAllOf`

NewFcpoolBlockAllOf instantiates a new FcpoolBlockAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpoolBlockAllOfWithDefaults

`func NewFcpoolBlockAllOfWithDefaults() *FcpoolBlockAllOf`

NewFcpoolBlockAllOfWithDefaults instantiates a new FcpoolBlockAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcpoolBlockAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcpoolBlockAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcpoolBlockAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcpoolBlockAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcpoolBlockAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcpoolBlockAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFrom

`func (o *FcpoolBlockAllOf) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *FcpoolBlockAllOf) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *FcpoolBlockAllOf) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *FcpoolBlockAllOf) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *FcpoolBlockAllOf) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *FcpoolBlockAllOf) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *FcpoolBlockAllOf) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *FcpoolBlockAllOf) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


