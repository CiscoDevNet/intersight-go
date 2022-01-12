# AssetDeviceTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.DeviceTransaction"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.DeviceTransaction"]
**Action** | Pointer to **string** | The action taken by Cisco Install Base on the device. | [optional] [readonly] 
**StatusDescription** | Pointer to **string** | Description of status of Cisco device reported by Cisco Install Base. | [optional] [readonly] 
**StatusId** | Pointer to **int32** | Status ID of device reported by Cisco Install Base. * &#x60;0&#x60; - A default value to catch cases where device status is not correctly detected. * &#x60;10000&#x60; - Device is installed successfully. * &#x60;1010041&#x60; - Device is currently in Return Material Authorization process. * &#x60;10005&#x60; - Device is replaced successfully with another device. * &#x60;10007&#x60; - Device is returned succcessfuly. * &#x60;10009&#x60; - Device is terminated at customer&#39;s end. | [optional] [readonly] [default to 0]
**Timestamp** | Pointer to **string** | Timestamp field reported by Cisco Install Base. | [optional] [readonly] 
**TransactionBatchId** | Pointer to **int64** | Transaction batch ID reported by Cisco Install Base. | [optional] [readonly] 
**TransactionDate** | Pointer to **string** | Transaction Date reported by Cisco Install Base. | [optional] [readonly] 
**TransactionSequence** | Pointer to **int64** | Transaction sequence reported by Cisco Install Base. | [optional] [readonly] 

## Methods

### NewAssetDeviceTransaction

`func NewAssetDeviceTransaction(classId string, objectType string, ) *AssetDeviceTransaction`

NewAssetDeviceTransaction instantiates a new AssetDeviceTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceTransactionWithDefaults

`func NewAssetDeviceTransactionWithDefaults() *AssetDeviceTransaction`

NewAssetDeviceTransactionWithDefaults instantiates a new AssetDeviceTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeviceTransaction) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeviceTransaction) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeviceTransaction) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeviceTransaction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeviceTransaction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeviceTransaction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *AssetDeviceTransaction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AssetDeviceTransaction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AssetDeviceTransaction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AssetDeviceTransaction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetStatusDescription

`func (o *AssetDeviceTransaction) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *AssetDeviceTransaction) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *AssetDeviceTransaction) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *AssetDeviceTransaction) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetStatusId

`func (o *AssetDeviceTransaction) GetStatusId() int32`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *AssetDeviceTransaction) GetStatusIdOk() (*int32, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *AssetDeviceTransaction) SetStatusId(v int32)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *AssetDeviceTransaction) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AssetDeviceTransaction) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AssetDeviceTransaction) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AssetDeviceTransaction) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AssetDeviceTransaction) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTransactionBatchId

`func (o *AssetDeviceTransaction) GetTransactionBatchId() int64`

GetTransactionBatchId returns the TransactionBatchId field if non-nil, zero value otherwise.

### GetTransactionBatchIdOk

`func (o *AssetDeviceTransaction) GetTransactionBatchIdOk() (*int64, bool)`

GetTransactionBatchIdOk returns a tuple with the TransactionBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionBatchId

`func (o *AssetDeviceTransaction) SetTransactionBatchId(v int64)`

SetTransactionBatchId sets TransactionBatchId field to given value.

### HasTransactionBatchId

`func (o *AssetDeviceTransaction) HasTransactionBatchId() bool`

HasTransactionBatchId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *AssetDeviceTransaction) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *AssetDeviceTransaction) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *AssetDeviceTransaction) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *AssetDeviceTransaction) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetTransactionSequence

`func (o *AssetDeviceTransaction) GetTransactionSequence() int64`

GetTransactionSequence returns the TransactionSequence field if non-nil, zero value otherwise.

### GetTransactionSequenceOk

`func (o *AssetDeviceTransaction) GetTransactionSequenceOk() (*int64, bool)`

GetTransactionSequenceOk returns a tuple with the TransactionSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSequence

`func (o *AssetDeviceTransaction) SetTransactionSequence(v int64)`

SetTransactionSequence sets TransactionSequence field to given value.

### HasTransactionSequence

`func (o *AssetDeviceTransaction) HasTransactionSequence() bool`

HasTransactionSequence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


