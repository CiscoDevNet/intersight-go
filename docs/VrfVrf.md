# VrfVrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vrf.Vrf"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vrf.Vrf"]
**Description** | Pointer to **string** | Description to help identify or describe this VRF. | [optional] 
**Name** | Pointer to **string** | Name of the Virtual Routing and Forwarding Instance. | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewVrfVrf

`func NewVrfVrf(classId string, objectType string, ) *VrfVrf`

NewVrfVrf instantiates a new VrfVrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfVrfWithDefaults

`func NewVrfVrfWithDefaults() *VrfVrf`

NewVrfVrfWithDefaults instantiates a new VrfVrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VrfVrf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VrfVrf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VrfVrf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VrfVrf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VrfVrf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VrfVrf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *VrfVrf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VrfVrf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VrfVrf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VrfVrf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VrfVrf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VrfVrf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VrfVrf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VrfVrf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *VrfVrf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *VrfVrf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *VrfVrf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *VrfVrf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *VrfVrf) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *VrfVrf) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


