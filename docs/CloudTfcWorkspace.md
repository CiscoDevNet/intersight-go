# CloudTfcWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.TfcWorkspace"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.TfcWorkspace"]
**AgentPoolId** | Pointer to **string** | The agent pool associated with this workspace. | [optional] [readonly] 
**ApplyMethod** | Pointer to **bool** | Whether or not Terraform Cloud should automatically apply a successful Terraform plan. | [optional] [readonly] 
**ExecutionMode** | Pointer to **string** | Indicates where the Terraform cloud should execute the runs. | [optional] [readonly] 
**Identity** | Pointer to **string** | The unique id for this workspace. | [optional] [readonly] 
**LastRunStatus** | Pointer to **string** | The status of the last executed run in this workspace. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the workspace. | [optional] [readonly] 
**WorkspaceVariables** | Pointer to [**[]CloudTfcWorkspaceVariables**](CloudTfcWorkspaceVariables.md) |  | [optional] 
**Organization** | Pointer to [**NullableCloudTfcOrganizationRelationship**](CloudTfcOrganizationRelationship.md) |  | [optional] 

## Methods

### NewCloudTfcWorkspace

`func NewCloudTfcWorkspace(classId string, objectType string, ) *CloudTfcWorkspace`

NewCloudTfcWorkspace instantiates a new CloudTfcWorkspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTfcWorkspaceWithDefaults

`func NewCloudTfcWorkspaceWithDefaults() *CloudTfcWorkspace`

NewCloudTfcWorkspaceWithDefaults instantiates a new CloudTfcWorkspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudTfcWorkspace) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudTfcWorkspace) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudTfcWorkspace) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudTfcWorkspace) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudTfcWorkspace) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudTfcWorkspace) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAgentPoolId

`func (o *CloudTfcWorkspace) GetAgentPoolId() string`

GetAgentPoolId returns the AgentPoolId field if non-nil, zero value otherwise.

### GetAgentPoolIdOk

`func (o *CloudTfcWorkspace) GetAgentPoolIdOk() (*string, bool)`

GetAgentPoolIdOk returns a tuple with the AgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolId

`func (o *CloudTfcWorkspace) SetAgentPoolId(v string)`

SetAgentPoolId sets AgentPoolId field to given value.

### HasAgentPoolId

`func (o *CloudTfcWorkspace) HasAgentPoolId() bool`

HasAgentPoolId returns a boolean if a field has been set.

### GetApplyMethod

`func (o *CloudTfcWorkspace) GetApplyMethod() bool`

GetApplyMethod returns the ApplyMethod field if non-nil, zero value otherwise.

### GetApplyMethodOk

`func (o *CloudTfcWorkspace) GetApplyMethodOk() (*bool, bool)`

GetApplyMethodOk returns a tuple with the ApplyMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMethod

`func (o *CloudTfcWorkspace) SetApplyMethod(v bool)`

SetApplyMethod sets ApplyMethod field to given value.

### HasApplyMethod

`func (o *CloudTfcWorkspace) HasApplyMethod() bool`

HasApplyMethod returns a boolean if a field has been set.

### GetExecutionMode

`func (o *CloudTfcWorkspace) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *CloudTfcWorkspace) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *CloudTfcWorkspace) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *CloudTfcWorkspace) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetIdentity

`func (o *CloudTfcWorkspace) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudTfcWorkspace) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudTfcWorkspace) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudTfcWorkspace) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetLastRunStatus

`func (o *CloudTfcWorkspace) GetLastRunStatus() string`

GetLastRunStatus returns the LastRunStatus field if non-nil, zero value otherwise.

### GetLastRunStatusOk

`func (o *CloudTfcWorkspace) GetLastRunStatusOk() (*string, bool)`

GetLastRunStatusOk returns a tuple with the LastRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunStatus

`func (o *CloudTfcWorkspace) SetLastRunStatus(v string)`

SetLastRunStatus sets LastRunStatus field to given value.

### HasLastRunStatus

`func (o *CloudTfcWorkspace) HasLastRunStatus() bool`

HasLastRunStatus returns a boolean if a field has been set.

### GetName

`func (o *CloudTfcWorkspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudTfcWorkspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudTfcWorkspace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudTfcWorkspace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWorkspaceVariables

`func (o *CloudTfcWorkspace) GetWorkspaceVariables() []CloudTfcWorkspaceVariables`

GetWorkspaceVariables returns the WorkspaceVariables field if non-nil, zero value otherwise.

### GetWorkspaceVariablesOk

`func (o *CloudTfcWorkspace) GetWorkspaceVariablesOk() (*[]CloudTfcWorkspaceVariables, bool)`

GetWorkspaceVariablesOk returns a tuple with the WorkspaceVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceVariables

`func (o *CloudTfcWorkspace) SetWorkspaceVariables(v []CloudTfcWorkspaceVariables)`

SetWorkspaceVariables sets WorkspaceVariables field to given value.

### HasWorkspaceVariables

`func (o *CloudTfcWorkspace) HasWorkspaceVariables() bool`

HasWorkspaceVariables returns a boolean if a field has been set.

### SetWorkspaceVariablesNil

`func (o *CloudTfcWorkspace) SetWorkspaceVariablesNil(b bool)`

 SetWorkspaceVariablesNil sets the value for WorkspaceVariables to be an explicit nil

### UnsetWorkspaceVariables
`func (o *CloudTfcWorkspace) UnsetWorkspaceVariables()`

UnsetWorkspaceVariables ensures that no value is present for WorkspaceVariables, not even an explicit nil
### GetOrganization

`func (o *CloudTfcWorkspace) GetOrganization() CloudTfcOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CloudTfcWorkspace) GetOrganizationOk() (*CloudTfcOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CloudTfcWorkspace) SetOrganization(v CloudTfcOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CloudTfcWorkspace) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *CloudTfcWorkspace) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *CloudTfcWorkspace) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


