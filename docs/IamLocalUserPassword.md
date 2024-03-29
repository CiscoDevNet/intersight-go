# IamLocalUserPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LocalUserPassword"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LocalUserPassword"]
**CurrentPassword** | Pointer to **string** | User-entered password to be compared to password for change password function. | [optional] 
**InitialPassword** | Pointer to **string** | Initial password set for the local user for the first time when the local user gets created or when the password gets reset by the Account Administrator. | [optional] 
**IsCurrentPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;currentPassword&#39; property has been set. | [optional] [readonly] [default to false]
**IsInitialPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;initialPassword&#39; property has been set. | [optional] [readonly] [default to false]
**IsNewPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;newPassword&#39; property has been set. | [optional] [readonly] [default to false]
**NeedPasswordReset** | Pointer to **bool** | Indicates whether the user should be prompted to reset their password. | [optional] [readonly] 
**NewPassword** | Pointer to **string** | New password that the user&#39;s password should be changed to. | [optional] 
**Password** | Pointer to **string** | User&#39;s current valid password. | [optional] 
**User** | Pointer to [**NullableIamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewIamLocalUserPassword

`func NewIamLocalUserPassword(classId string, objectType string, ) *IamLocalUserPassword`

NewIamLocalUserPassword instantiates a new IamLocalUserPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLocalUserPasswordWithDefaults

`func NewIamLocalUserPasswordWithDefaults() *IamLocalUserPassword`

NewIamLocalUserPasswordWithDefaults instantiates a new IamLocalUserPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLocalUserPassword) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLocalUserPassword) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLocalUserPassword) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLocalUserPassword) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLocalUserPassword) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLocalUserPassword) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentPassword

`func (o *IamLocalUserPassword) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *IamLocalUserPassword) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *IamLocalUserPassword) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *IamLocalUserPassword) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetInitialPassword

`func (o *IamLocalUserPassword) GetInitialPassword() string`

GetInitialPassword returns the InitialPassword field if non-nil, zero value otherwise.

### GetInitialPasswordOk

`func (o *IamLocalUserPassword) GetInitialPasswordOk() (*string, bool)`

GetInitialPasswordOk returns a tuple with the InitialPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPassword

`func (o *IamLocalUserPassword) SetInitialPassword(v string)`

SetInitialPassword sets InitialPassword field to given value.

### HasInitialPassword

`func (o *IamLocalUserPassword) HasInitialPassword() bool`

HasInitialPassword returns a boolean if a field has been set.

### GetIsCurrentPasswordSet

`func (o *IamLocalUserPassword) GetIsCurrentPasswordSet() bool`

GetIsCurrentPasswordSet returns the IsCurrentPasswordSet field if non-nil, zero value otherwise.

### GetIsCurrentPasswordSetOk

`func (o *IamLocalUserPassword) GetIsCurrentPasswordSetOk() (*bool, bool)`

GetIsCurrentPasswordSetOk returns a tuple with the IsCurrentPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrentPasswordSet

`func (o *IamLocalUserPassword) SetIsCurrentPasswordSet(v bool)`

SetIsCurrentPasswordSet sets IsCurrentPasswordSet field to given value.

### HasIsCurrentPasswordSet

`func (o *IamLocalUserPassword) HasIsCurrentPasswordSet() bool`

HasIsCurrentPasswordSet returns a boolean if a field has been set.

### GetIsInitialPasswordSet

`func (o *IamLocalUserPassword) GetIsInitialPasswordSet() bool`

GetIsInitialPasswordSet returns the IsInitialPasswordSet field if non-nil, zero value otherwise.

### GetIsInitialPasswordSetOk

`func (o *IamLocalUserPassword) GetIsInitialPasswordSetOk() (*bool, bool)`

GetIsInitialPasswordSetOk returns a tuple with the IsInitialPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitialPasswordSet

`func (o *IamLocalUserPassword) SetIsInitialPasswordSet(v bool)`

SetIsInitialPasswordSet sets IsInitialPasswordSet field to given value.

### HasIsInitialPasswordSet

`func (o *IamLocalUserPassword) HasIsInitialPasswordSet() bool`

HasIsInitialPasswordSet returns a boolean if a field has been set.

### GetIsNewPasswordSet

`func (o *IamLocalUserPassword) GetIsNewPasswordSet() bool`

GetIsNewPasswordSet returns the IsNewPasswordSet field if non-nil, zero value otherwise.

### GetIsNewPasswordSetOk

`func (o *IamLocalUserPassword) GetIsNewPasswordSetOk() (*bool, bool)`

GetIsNewPasswordSetOk returns a tuple with the IsNewPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewPasswordSet

`func (o *IamLocalUserPassword) SetIsNewPasswordSet(v bool)`

SetIsNewPasswordSet sets IsNewPasswordSet field to given value.

### HasIsNewPasswordSet

`func (o *IamLocalUserPassword) HasIsNewPasswordSet() bool`

HasIsNewPasswordSet returns a boolean if a field has been set.

### GetNeedPasswordReset

`func (o *IamLocalUserPassword) GetNeedPasswordReset() bool`

GetNeedPasswordReset returns the NeedPasswordReset field if non-nil, zero value otherwise.

### GetNeedPasswordResetOk

`func (o *IamLocalUserPassword) GetNeedPasswordResetOk() (*bool, bool)`

GetNeedPasswordResetOk returns a tuple with the NeedPasswordReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedPasswordReset

`func (o *IamLocalUserPassword) SetNeedPasswordReset(v bool)`

SetNeedPasswordReset sets NeedPasswordReset field to given value.

### HasNeedPasswordReset

`func (o *IamLocalUserPassword) HasNeedPasswordReset() bool`

HasNeedPasswordReset returns a boolean if a field has been set.

### GetNewPassword

`func (o *IamLocalUserPassword) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *IamLocalUserPassword) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *IamLocalUserPassword) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *IamLocalUserPassword) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetPassword

`func (o *IamLocalUserPassword) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamLocalUserPassword) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamLocalUserPassword) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamLocalUserPassword) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUser

`func (o *IamLocalUserPassword) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamLocalUserPassword) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamLocalUserPassword) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamLocalUserPassword) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *IamLocalUserPassword) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *IamLocalUserPassword) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


