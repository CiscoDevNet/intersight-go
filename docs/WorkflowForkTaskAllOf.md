# WorkflowForkTaskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ForkTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ForkTask"]
**ForkedTasks** | Pointer to **[]string** |  | [optional] 
**JoinTask** | Pointer to **string** | Task name for the join control task that must follow a fork control task. | [optional] 

## Methods

### NewWorkflowForkTaskAllOf

`func NewWorkflowForkTaskAllOf(classId string, objectType string, ) *WorkflowForkTaskAllOf`

NewWorkflowForkTaskAllOf instantiates a new WorkflowForkTaskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowForkTaskAllOfWithDefaults

`func NewWorkflowForkTaskAllOfWithDefaults() *WorkflowForkTaskAllOf`

NewWorkflowForkTaskAllOfWithDefaults instantiates a new WorkflowForkTaskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowForkTaskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowForkTaskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowForkTaskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowForkTaskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowForkTaskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowForkTaskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetForkedTasks

`func (o *WorkflowForkTaskAllOf) GetForkedTasks() []string`

GetForkedTasks returns the ForkedTasks field if non-nil, zero value otherwise.

### GetForkedTasksOk

`func (o *WorkflowForkTaskAllOf) GetForkedTasksOk() (*[]string, bool)`

GetForkedTasksOk returns a tuple with the ForkedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkedTasks

`func (o *WorkflowForkTaskAllOf) SetForkedTasks(v []string)`

SetForkedTasks sets ForkedTasks field to given value.

### HasForkedTasks

`func (o *WorkflowForkTaskAllOf) HasForkedTasks() bool`

HasForkedTasks returns a boolean if a field has been set.

### SetForkedTasksNil

`func (o *WorkflowForkTaskAllOf) SetForkedTasksNil(b bool)`

 SetForkedTasksNil sets the value for ForkedTasks to be an explicit nil

### UnsetForkedTasks
`func (o *WorkflowForkTaskAllOf) UnsetForkedTasks()`

UnsetForkedTasks ensures that no value is present for ForkedTasks, not even an explicit nil
### GetJoinTask

`func (o *WorkflowForkTaskAllOf) GetJoinTask() string`

GetJoinTask returns the JoinTask field if non-nil, zero value otherwise.

### GetJoinTaskOk

`func (o *WorkflowForkTaskAllOf) GetJoinTaskOk() (*string, bool)`

GetJoinTaskOk returns a tuple with the JoinTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTask

`func (o *WorkflowForkTaskAllOf) SetJoinTask(v string)`

SetJoinTask sets JoinTask field to given value.

### HasJoinTask

`func (o *WorkflowForkTaskAllOf) HasJoinTask() bool`

HasJoinTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


