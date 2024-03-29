# TechsupportmanagementCollectionControlPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "techsupportmanagement.CollectionControlPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "techsupportmanagement.CollectionControlPolicy"]
**DeploymentType** | Pointer to **string** | Deployment type defines whether the policy is associated with a SaaS or Appliance account. * &#x60;None&#x60; - Service deployment type None. * &#x60;SaaS&#x60; - Service deployment type SaaS. * &#x60;Appliance&#x60; - Service deployment type Appliance. | [optional] [readonly] [default to "None"]
**TechSupportCollection** | Pointer to **string** | Enable or Disable techsupport collection for a specific account. * &#x60;Enable&#x60; - Enable techsupport collection. * &#x60;Disable&#x60; - Disable techsupport collection. | [optional] [default to "Enable"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewTechsupportmanagementCollectionControlPolicy

`func NewTechsupportmanagementCollectionControlPolicy(classId string, objectType string, ) *TechsupportmanagementCollectionControlPolicy`

NewTechsupportmanagementCollectionControlPolicy instantiates a new TechsupportmanagementCollectionControlPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementCollectionControlPolicyWithDefaults

`func NewTechsupportmanagementCollectionControlPolicyWithDefaults() *TechsupportmanagementCollectionControlPolicy`

NewTechsupportmanagementCollectionControlPolicyWithDefaults instantiates a new TechsupportmanagementCollectionControlPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TechsupportmanagementCollectionControlPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TechsupportmanagementCollectionControlPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TechsupportmanagementCollectionControlPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TechsupportmanagementCollectionControlPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TechsupportmanagementCollectionControlPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TechsupportmanagementCollectionControlPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeploymentType

`func (o *TechsupportmanagementCollectionControlPolicy) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *TechsupportmanagementCollectionControlPolicy) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *TechsupportmanagementCollectionControlPolicy) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *TechsupportmanagementCollectionControlPolicy) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetTechSupportCollection

`func (o *TechsupportmanagementCollectionControlPolicy) GetTechSupportCollection() string`

GetTechSupportCollection returns the TechSupportCollection field if non-nil, zero value otherwise.

### GetTechSupportCollectionOk

`func (o *TechsupportmanagementCollectionControlPolicy) GetTechSupportCollectionOk() (*string, bool)`

GetTechSupportCollectionOk returns a tuple with the TechSupportCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSupportCollection

`func (o *TechsupportmanagementCollectionControlPolicy) SetTechSupportCollection(v string)`

SetTechSupportCollection sets TechSupportCollection field to given value.

### HasTechSupportCollection

`func (o *TechsupportmanagementCollectionControlPolicy) HasTechSupportCollection() bool`

HasTechSupportCollection returns a boolean if a field has been set.

### GetAccount

`func (o *TechsupportmanagementCollectionControlPolicy) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TechsupportmanagementCollectionControlPolicy) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TechsupportmanagementCollectionControlPolicy) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TechsupportmanagementCollectionControlPolicy) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *TechsupportmanagementCollectionControlPolicy) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *TechsupportmanagementCollectionControlPolicy) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


