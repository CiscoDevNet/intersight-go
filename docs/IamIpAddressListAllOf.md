# IamIpAddressListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;iam.IpAddress&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]IamIpAddress**](IamIpAddress.md) | The array of &#39;iam.IpAddress&#39; resources matching the request. | [optional] 

## Methods

### NewIamIpAddressListAllOf

`func NewIamIpAddressListAllOf() *IamIpAddressListAllOf`

NewIamIpAddressListAllOf instantiates a new IamIpAddressListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamIpAddressListAllOfWithDefaults

`func NewIamIpAddressListAllOfWithDefaults() *IamIpAddressListAllOf`

NewIamIpAddressListAllOfWithDefaults instantiates a new IamIpAddressListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IamIpAddressListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IamIpAddressListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IamIpAddressListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IamIpAddressListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *IamIpAddressListAllOf) GetResults() []IamIpAddress`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IamIpAddressListAllOf) GetResultsOk() (*[]IamIpAddress, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IamIpAddressListAllOf) SetResults(v []IamIpAddress)`

SetResults sets Results field to given value.

### HasResults

`func (o *IamIpAddressListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *IamIpAddressListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *IamIpAddressListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


