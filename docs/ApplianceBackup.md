# ApplianceBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.Backup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.Backup"]
**ElapsedTime** | Pointer to **int64** | Elapsed time in seconds since the backup process has started. | [optional] [readonly] 
**EndTime** | Pointer to **time.Time** | End date and time of the backup process. | [optional] [readonly] 
**IsManual** | Pointer to **bool** | If true, represents a manual backup. Else represents a scheduled backup. | [optional] [readonly] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Messages** | Pointer to **[]string** |  | [optional] 
**Password** | Pointer to **string** | Password to authenticate the fileserver. | [optional] 
**StartTime** | Pointer to **time.Time** | Start date and time of the backup process. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the backup managed object. * &#x60;Started&#x60; - Backup or restore process has started. * &#x60;Created&#x60; - Backup or restore is in created state. * &#x60;Failed&#x60; - Backup or restore process has failed. * &#x60;Completed&#x60; - Backup or restore process has completed. * &#x60;Copied&#x60; - Backup file has been copied. * &#x60;Cleanup Failed&#x60; - Cleanup of the old backup has failed. | [optional] [readonly] [default to "Started"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceBackup

`func NewApplianceBackup(classId string, objectType string, ) *ApplianceBackup`

NewApplianceBackup instantiates a new ApplianceBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupWithDefaults

`func NewApplianceBackupWithDefaults() *ApplianceBackup`

NewApplianceBackupWithDefaults instantiates a new ApplianceBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceBackup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceBackup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceBackup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceBackup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceBackup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceBackup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetElapsedTime

`func (o *ApplianceBackup) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *ApplianceBackup) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *ApplianceBackup) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *ApplianceBackup) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ApplianceBackup) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceBackup) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceBackup) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceBackup) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetIsManual

`func (o *ApplianceBackup) GetIsManual() bool`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *ApplianceBackup) GetIsManualOk() (*bool, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *ApplianceBackup) SetIsManual(v bool)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *ApplianceBackup) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *ApplianceBackup) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *ApplianceBackup) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *ApplianceBackup) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *ApplianceBackup) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetMessages

`func (o *ApplianceBackup) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ApplianceBackup) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ApplianceBackup) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ApplianceBackup) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *ApplianceBackup) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *ApplianceBackup) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetPassword

`func (o *ApplianceBackup) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApplianceBackup) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApplianceBackup) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApplianceBackup) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetStartTime

`func (o *ApplianceBackup) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceBackup) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceBackup) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceBackup) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceBackup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceBackup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceBackup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceBackup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceBackup) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceBackup) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceBackup) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceBackup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ApplianceBackup) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ApplianceBackup) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


