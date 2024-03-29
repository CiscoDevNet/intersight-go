# TechsupportmanagementDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "techsupportmanagement.Download"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "techsupportmanagement.Download"]
**TechSupportStatus** | Pointer to [**NullableTechsupportmanagementTechSupportStatusRelationship**](TechsupportmanagementTechSupportStatusRelationship.md) |  | [optional] 

## Methods

### NewTechsupportmanagementDownload

`func NewTechsupportmanagementDownload(classId string, objectType string, ) *TechsupportmanagementDownload`

NewTechsupportmanagementDownload instantiates a new TechsupportmanagementDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementDownloadWithDefaults

`func NewTechsupportmanagementDownloadWithDefaults() *TechsupportmanagementDownload`

NewTechsupportmanagementDownloadWithDefaults instantiates a new TechsupportmanagementDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TechsupportmanagementDownload) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TechsupportmanagementDownload) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TechsupportmanagementDownload) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TechsupportmanagementDownload) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TechsupportmanagementDownload) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TechsupportmanagementDownload) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTechSupportStatus

`func (o *TechsupportmanagementDownload) GetTechSupportStatus() TechsupportmanagementTechSupportStatusRelationship`

GetTechSupportStatus returns the TechSupportStatus field if non-nil, zero value otherwise.

### GetTechSupportStatusOk

`func (o *TechsupportmanagementDownload) GetTechSupportStatusOk() (*TechsupportmanagementTechSupportStatusRelationship, bool)`

GetTechSupportStatusOk returns a tuple with the TechSupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSupportStatus

`func (o *TechsupportmanagementDownload) SetTechSupportStatus(v TechsupportmanagementTechSupportStatusRelationship)`

SetTechSupportStatus sets TechSupportStatus field to given value.

### HasTechSupportStatus

`func (o *TechsupportmanagementDownload) HasTechSupportStatus() bool`

HasTechSupportStatus returns a boolean if a field has been set.

### SetTechSupportStatusNil

`func (o *TechsupportmanagementDownload) SetTechSupportStatusNil(b bool)`

 SetTechSupportStatusNil sets the value for TechSupportStatus to be an explicit nil

### UnsetTechSupportStatus
`func (o *TechsupportmanagementDownload) UnsetTechSupportStatus()`

UnsetTechSupportStatus ensures that no value is present for TechSupportStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


