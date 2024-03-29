# IaasLicenseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.LicenseInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.LicenseInfo"]
**LicenseExpirationDate** | Pointer to **string** | UCS Director license expiration date. | [optional] [readonly] 
**LicenseKeysInfo** | Pointer to [**[]IaasLicenseKeysInfo**](IaasLicenseKeysInfo.md) |  | [optional] 
**LicenseType** | Pointer to **string** | License type of UCSD whether it is EVAL/Permanent/Subscription.. | [optional] [readonly] 
**LicenseUtilizationInfo** | Pointer to [**[]IaasLicenseUtilizationInfo**](IaasLicenseUtilizationInfo.md) |  | [optional] 
**Guid** | Pointer to [**NullableIaasUcsdInfoRelationship**](IaasUcsdInfoRelationship.md) |  | [optional] 

## Methods

### NewIaasLicenseInfo

`func NewIaasLicenseInfo(classId string, objectType string, ) *IaasLicenseInfo`

NewIaasLicenseInfo instantiates a new IaasLicenseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasLicenseInfoWithDefaults

`func NewIaasLicenseInfoWithDefaults() *IaasLicenseInfo`

NewIaasLicenseInfoWithDefaults instantiates a new IaasLicenseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasLicenseInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasLicenseInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasLicenseInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasLicenseInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasLicenseInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasLicenseInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLicenseExpirationDate

`func (o *IaasLicenseInfo) GetLicenseExpirationDate() string`

GetLicenseExpirationDate returns the LicenseExpirationDate field if non-nil, zero value otherwise.

### GetLicenseExpirationDateOk

`func (o *IaasLicenseInfo) GetLicenseExpirationDateOk() (*string, bool)`

GetLicenseExpirationDateOk returns a tuple with the LicenseExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpirationDate

`func (o *IaasLicenseInfo) SetLicenseExpirationDate(v string)`

SetLicenseExpirationDate sets LicenseExpirationDate field to given value.

### HasLicenseExpirationDate

`func (o *IaasLicenseInfo) HasLicenseExpirationDate() bool`

HasLicenseExpirationDate returns a boolean if a field has been set.

### GetLicenseKeysInfo

`func (o *IaasLicenseInfo) GetLicenseKeysInfo() []IaasLicenseKeysInfo`

GetLicenseKeysInfo returns the LicenseKeysInfo field if non-nil, zero value otherwise.

### GetLicenseKeysInfoOk

`func (o *IaasLicenseInfo) GetLicenseKeysInfoOk() (*[]IaasLicenseKeysInfo, bool)`

GetLicenseKeysInfoOk returns a tuple with the LicenseKeysInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKeysInfo

`func (o *IaasLicenseInfo) SetLicenseKeysInfo(v []IaasLicenseKeysInfo)`

SetLicenseKeysInfo sets LicenseKeysInfo field to given value.

### HasLicenseKeysInfo

`func (o *IaasLicenseInfo) HasLicenseKeysInfo() bool`

HasLicenseKeysInfo returns a boolean if a field has been set.

### SetLicenseKeysInfoNil

`func (o *IaasLicenseInfo) SetLicenseKeysInfoNil(b bool)`

 SetLicenseKeysInfoNil sets the value for LicenseKeysInfo to be an explicit nil

### UnsetLicenseKeysInfo
`func (o *IaasLicenseInfo) UnsetLicenseKeysInfo()`

UnsetLicenseKeysInfo ensures that no value is present for LicenseKeysInfo, not even an explicit nil
### GetLicenseType

`func (o *IaasLicenseInfo) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *IaasLicenseInfo) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *IaasLicenseInfo) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *IaasLicenseInfo) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetLicenseUtilizationInfo

`func (o *IaasLicenseInfo) GetLicenseUtilizationInfo() []IaasLicenseUtilizationInfo`

GetLicenseUtilizationInfo returns the LicenseUtilizationInfo field if non-nil, zero value otherwise.

### GetLicenseUtilizationInfoOk

`func (o *IaasLicenseInfo) GetLicenseUtilizationInfoOk() (*[]IaasLicenseUtilizationInfo, bool)`

GetLicenseUtilizationInfoOk returns a tuple with the LicenseUtilizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUtilizationInfo

`func (o *IaasLicenseInfo) SetLicenseUtilizationInfo(v []IaasLicenseUtilizationInfo)`

SetLicenseUtilizationInfo sets LicenseUtilizationInfo field to given value.

### HasLicenseUtilizationInfo

`func (o *IaasLicenseInfo) HasLicenseUtilizationInfo() bool`

HasLicenseUtilizationInfo returns a boolean if a field has been set.

### SetLicenseUtilizationInfoNil

`func (o *IaasLicenseInfo) SetLicenseUtilizationInfoNil(b bool)`

 SetLicenseUtilizationInfoNil sets the value for LicenseUtilizationInfo to be an explicit nil

### UnsetLicenseUtilizationInfo
`func (o *IaasLicenseInfo) UnsetLicenseUtilizationInfo()`

UnsetLicenseUtilizationInfo ensures that no value is present for LicenseUtilizationInfo, not even an explicit nil
### GetGuid

`func (o *IaasLicenseInfo) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasLicenseInfo) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasLicenseInfo) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasLicenseInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *IaasLicenseInfo) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *IaasLicenseInfo) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


