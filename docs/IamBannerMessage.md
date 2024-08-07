# IamBannerMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.BannerMessage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.BannerMessage"]
**BannerContents** | Pointer to **string** | Contents of the banner message. | [optional] 
**BannerDisplay** | Pointer to **bool** | Whether or not to display the banner message. | [optional] 
**BannerTitle** | Pointer to **string** | Title of the banner message. | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewIamBannerMessage

`func NewIamBannerMessage(classId string, objectType string, ) *IamBannerMessage`

NewIamBannerMessage instantiates a new IamBannerMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamBannerMessageWithDefaults

`func NewIamBannerMessageWithDefaults() *IamBannerMessage`

NewIamBannerMessageWithDefaults instantiates a new IamBannerMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamBannerMessage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamBannerMessage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamBannerMessage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamBannerMessage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamBannerMessage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamBannerMessage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBannerContents

`func (o *IamBannerMessage) GetBannerContents() string`

GetBannerContents returns the BannerContents field if non-nil, zero value otherwise.

### GetBannerContentsOk

`func (o *IamBannerMessage) GetBannerContentsOk() (*string, bool)`

GetBannerContentsOk returns a tuple with the BannerContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerContents

`func (o *IamBannerMessage) SetBannerContents(v string)`

SetBannerContents sets BannerContents field to given value.

### HasBannerContents

`func (o *IamBannerMessage) HasBannerContents() bool`

HasBannerContents returns a boolean if a field has been set.

### GetBannerDisplay

`func (o *IamBannerMessage) GetBannerDisplay() bool`

GetBannerDisplay returns the BannerDisplay field if non-nil, zero value otherwise.

### GetBannerDisplayOk

`func (o *IamBannerMessage) GetBannerDisplayOk() (*bool, bool)`

GetBannerDisplayOk returns a tuple with the BannerDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerDisplay

`func (o *IamBannerMessage) SetBannerDisplay(v bool)`

SetBannerDisplay sets BannerDisplay field to given value.

### HasBannerDisplay

`func (o *IamBannerMessage) HasBannerDisplay() bool`

HasBannerDisplay returns a boolean if a field has been set.

### GetBannerTitle

`func (o *IamBannerMessage) GetBannerTitle() string`

GetBannerTitle returns the BannerTitle field if non-nil, zero value otherwise.

### GetBannerTitleOk

`func (o *IamBannerMessage) GetBannerTitleOk() (*string, bool)`

GetBannerTitleOk returns a tuple with the BannerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerTitle

`func (o *IamBannerMessage) SetBannerTitle(v string)`

SetBannerTitle sets BannerTitle field to given value.

### HasBannerTitle

`func (o *IamBannerMessage) HasBannerTitle() bool`

HasBannerTitle returns a boolean if a field has been set.

### GetAccount

`func (o *IamBannerMessage) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamBannerMessage) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamBannerMessage) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamBannerMessage) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *IamBannerMessage) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *IamBannerMessage) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


