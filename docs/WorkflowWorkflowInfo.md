# WorkflowWorkflowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.WorkflowInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.WorkflowInfo"]
**Action** | Pointer to **string** | The action of the workflow such as start, cancel, retry, pause. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Create&#x60; - Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow. * &#x60;Start&#x60; - Start a new execution of the workflow. * &#x60;Pause&#x60; - Pause the workflow, this can only be issued on workflows that are in running state. A workflow can be paused for a maximum of 180 days, after 180 days the workflow will be terminated by the system. * &#x60;Resume&#x60; - Resume the workflow which was previously paused through pause action on the workflow. * &#x60;Rerun&#x60; - Rerun the workflow that has previously reached a failed state. The workflow is run from the beginning using inputs from previous execution. Completed and currently running workflows cannot be rerun. Workflows do not have to be marked for retry to use this action. * &#x60;Retry&#x60; - This action has been deprecated. Please use RetryFailed, Rerun or RetryFromTask action. Retry the workflow that has previously reached a final state and has the retryable property set to true. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted from the first task.  The task name in retryFromTaskName must be one of the tasks that completed or failed in the previous run. It is not possible to retry a workflow from a task which wasn&#39;t run in the previous iteration. * &#x60;RetryFailed&#x60; - Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task. * &#x60;RetryFromTask&#x60; - Retry the workflow that has previously reached a failed state and has the retryable property set to true. A running or waiting workflow cannot be retried. RetryFromTaskName must be passed along with this action, and the workflow will be started from that specific task. The task name in RetryFromTaskName must be one of the tasks that was executed in the previous attempt. It is not possible to retry a workflow from a task that wasn&#39;t run in the previous execution attempt. * &#x60;Cancel&#x60; - Cancel the workflow that is in running or waiting state. | [optional] [default to "None"]
**CleanupTime** | Pointer to **time.Time** | The time when the workflow info will be removed from the database. When WorkflowInfo is created, cleanup time will be set to 181 days. As the workflow progresses through different states the cleanup time can be updated. A cleanup time of 0 means the workflow is not scheduled for cleanup. An active workflow that continues to schedule &amp; run tasks can run for any amount of time and there is no upper bound for such workflows. Workflows that are not actively running, say in Paused or Waiting states will be removed after 181 days. | [optional] [readonly] 
**Email** | Pointer to **string** | The email address of the user who started this workflow. In the case of LDAP users, this field can hold either a username or an email. | [optional] [readonly] 
**EndTime** | Pointer to **time.Time** | The time when the workflow reached a final state. | [optional] [readonly] 
**FailedWorkflowCleanupDuration** | Pointer to **int64** | The duration in hours after which the workflow info for failed, terminated or timed out workflow will be removed from database. The minimum is 1 hour, maximum is 365 days and default is 90 days. | [optional] [default to 2160]
**Input** | Pointer to **interface{}** | All the given inputs for the workflow. The schema for the inputs is defined in the InputDefinition section of the WorkflowDefinition. The InputDefinition will provide a list of input fields to be accepted, the associated datatype of the inputs and any additional constraints on the inputs. For more information please refer to InputDefinition property in the the the WorkflowDefinition resource. The inputs for a workflow are provided as a collection of key-value pairs, where key is the name of the input and value is any valid JSON data which conforms to the datatype of the input as specified in the InputDefinition. When the input passed into a workflow does not match the datatype or the constraints specified in the workflow definition, it will not be accepted. For example, if the InputDefinition specified that workflow must accept a string name &#39;key&#39; and the value passed for key must adhere to a regex pattern. If Workflow was started with input where &#39;key&#39; is not a string matching the regex pattern, an error will be generated and workflow will not start execution. During workflow definition design, the input passed into the workflow will be referred using the format &#39;workflow.input.&lt;inputName&gt;&#39;. If the input is referred directly in a mapping it will be in the format &#39;${workflow.input.&lt;inputName&gt;}&#39; or inside a template mapping in the format &#39;{{.global.workflow.input.&lt;inputName&gt;}}&#39;. | [optional] 
**InstId** | Pointer to **string** | A workflow instance Id which is the unique identified for the workflow execution. | [optional] [readonly] 
**Internal** | Pointer to **bool** | Denotes that an Intersight service started this workflow as internal and hence will not be shown in Intersight User Interface. Typically these are internal system maintenance workflows which are triggered by Intersight services. | [optional] [readonly] 
**LastAction** | Pointer to **string** | The last action that was issued on the workflow is saved in this field. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Create&#x60; - Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow. * &#x60;Start&#x60; - Start a new execution of the workflow. * &#x60;Pause&#x60; - Pause the workflow, this can only be issued on workflows that are in running state. A workflow can be paused for a maximum of 180 days, after 180 days the workflow will be terminated by the system. * &#x60;Resume&#x60; - Resume the workflow which was previously paused through pause action on the workflow. * &#x60;Rerun&#x60; - Rerun the workflow that has previously reached a failed state. The workflow is run from the beginning using inputs from previous execution. Completed and currently running workflows cannot be rerun. Workflows do not have to be marked for retry to use this action. * &#x60;Retry&#x60; - This action has been deprecated. Please use RetryFailed, Rerun or RetryFromTask action. Retry the workflow that has previously reached a final state and has the retryable property set to true. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted from the first task.  The task name in retryFromTaskName must be one of the tasks that completed or failed in the previous run. It is not possible to retry a workflow from a task which wasn&#39;t run in the previous iteration. * &#x60;RetryFailed&#x60; - Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task. * &#x60;RetryFromTask&#x60; - Retry the workflow that has previously reached a failed state and has the retryable property set to true. A running or waiting workflow cannot be retried. RetryFromTaskName must be passed along with this action, and the workflow will be started from that specific task. The task name in RetryFromTaskName must be one of the tasks that was executed in the previous attempt. It is not possible to retry a workflow from a task that wasn&#39;t run in the previous execution attempt. * &#x60;Cancel&#x60; - Cancel the workflow that is in running or waiting state. | [optional] [readonly] [default to "None"]
**Message** | Pointer to [**[]WorkflowMessage**](WorkflowMessage.md) |  | [optional] 
**Name** | Pointer to **string** | A name of the workflow execution instance. | [optional] 
**Output** | Pointer to **interface{}** | All the generated outputs for the workflow. The schema for the outputs are defined in the OutputDefinition section of the WorkflowDefinition. The OutputDefinition will provide a list of output fields that could be generated after workflow execution is completed and the associated datatype of the outputs. For more information please refer to OutputDefinition property in WorkflowDefinition resource. The output for the workflow is generated as a collection of key-value pairs, where key is the name of the output and value is any valid JSON data which conforms to the datatype of output as specified in the OutputDefinition. During workflow definition design, if a workflow is included as a sub-workflow inside a parent workflow then the outputs generated by the sub-workflow can be used in the workflow design. For example, if workflow was included into parent workflow as &#39;SubWorkflowSample1&#39;, then that output can be referred as &#39;SubWorkflowSample1.output.&lt;outputName&gt;&#39;. In the output is referred directly in a mapping it will be in the format &#39;${SubWorkflowSample1.output.&lt;outputName&gt;}&#39; or inside a template mapping will be in the format &#39;{{SubWorkflowSample1.output.&lt;outputName&gt;}}&#39;. | [optional] [readonly] 
**PauseReason** | Pointer to **string** | Denotes the reason workflow is in paused status. * &#x60;None&#x60; - Pause reason is none, which indicates there is no reason for the pause state. * &#x60;TaskWithWarning&#x60; - Pause reason indicates the workflow is in this state due to a task that has a status as completed with warnings. * &#x60;SystemMaintenance&#x60; - Pause reason indicates the workflow is in this state based on actions of system admin for maintenance. | [optional] [readonly] [default to "None"]
**Progress** | Pointer to **float32** | This field indicates percentage of workflow task completion based on the total number of tasks in the workflow. The total number of tasks in the workflow is calculated based on the longest path the workflow execution can take. So progress is calculated based on the percentage of tasks that completed out of the total number of tasks that could be executed. Progress is not a representation of the time taken to complete the workflow. A task is considered as completed if the task status is either \&quot;NO_OP\&quot; or \&quot;COMPLETED\&quot;. If the task status is \&quot;SKIP_TO_FAIL\&quot;, the workflow will be terminated and the progress of the workflow will be set to 100. | [optional] [readonly] 
**Properties** | Pointer to [**NullableWorkflowWorkflowInfoProperties**](WorkflowWorkflowInfoProperties.md) |  | [optional] 
**RetryFromTaskName** | Pointer to **string** | This field is required when RetryFromTask action is issued for a workflow that is in a &#39;final&#39; state. The workflow will be retried from the specified task. This field must specify a task name which is the unique name of the task within the workflow. The task name must be one of the tasks that were completed or failed in the previous run. It is not possible to retry a workflow from a task that wasn&#39;t run in the previous execution attempt. | [optional] 
**Src** | Pointer to **string** | The source service that started the workflow execution and hence represents the owning service for this workflow. | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | The time when the workflow was started for execution. | [optional] [readonly] 
**Status** | Pointer to **string** | A status of the workflow (RUNNING, WAITING, COMPLETED, TIME_OUT, FAILED). The \&quot;status\&quot; field has been deprecated and is now replaced with the \&quot;workflowStatus\&quot; field. | [optional] [readonly] 
**SuccessWorkflowCleanupDuration** | Pointer to **int64** | The duration in hours after which the workflow info for successful workflow will be removed from database. The minimum is 1 hour, maximum is 365 days and default is 90 days. | [optional] [default to 2160]
**TaskInfoUpdate** | Pointer to [**NullableWorkflowTaskInfoUpdate**](WorkflowTaskInfoUpdate.md) |  | [optional] 
**TraceId** | Pointer to **string** | The trace id to keep track of workflow execution. | [optional] [readonly] 
**Type** | Pointer to **string** | A type of the workflow (serverconfig, ansible_monitoring). | [optional] [readonly] 
**UserActionRequired** | Pointer to **bool** | Property will be set when a user action is required on the workflow. This can be because the workflow is waiting for a wait task to be updated, workflow is paused or workflow launched by a configuration object has failed and needs to be retried in order to complete successfully. | [optional] [readonly] [default to false]
**UserId** | Pointer to **string** | The user identifier which indicates the user that started this workflow. | [optional] [readonly] 
**Variable** | Pointer to **interface{}** | All the generated variables for the workflow. During workflow execution, the variables will be updated as per the variableParameters specified after each task execution. | [optional] [readonly] 
**WaitReason** | Pointer to **string** | Denotes the reason workflow is in waiting status. * &#x60;None&#x60; - Wait reason is none, which indicates there is no reason for the waiting state. * &#x60;GatherTasks&#x60; - Wait reason is gathering tasks, which indicates the workflow is in this state in order to gather tasks. * &#x60;Duplicate&#x60; - Wait reason is duplicate, which indicates the workflow is a duplicate of current running workflow. * &#x60;RateLimit&#x60; - Wait reason is rate limit, which indicates the workflow is rate limited by account/instance level throttling threshold. * &#x60;WaitTask&#x60; - Wait reason when there are one or more wait tasks in the workflow which are yet to receive a task status update. * &#x60;PendingRetryFailed&#x60; - Wait reason when the workflow is pending a RetryFailed action. * &#x60;WaitingToStart&#x60; - Workflow is waiting to start on workflow engine. | [optional] [readonly] [default to "None"]
**WorkflowCtx** | Pointer to [**NullableWorkflowWorkflowCtx**](WorkflowWorkflowCtx.md) |  | [optional] 
**WorkflowStatus** | Pointer to **string** | The current state of the workflow execution instance. A draft workflow execution will be in NotStarted state and when \&quot;Start\&quot; action is issued then the workflow will move into Waiting state until the first task of the workflow is scheduled at which time it will move into InProgress state. When execution reaches a final state it move to either Completed, Failed or Terminated state. For more details look at the description for each state. * &#x60;NotStarted&#x60; - Initially all the workflow instances are at \&quot;NotStarted\&quot; state. A workflow can be drafted in this state by issuing Create action. When a workflow is in this state the inputs can be updated until the workflow is started. * &#x60;InProgress&#x60; - A workflow execution moves into \&quot;InProgress\&quot; state when the first task of the workflow is scheduled for execution and continues to remain in that state as long as there are tasks executing or yet to be scheduled for execution. * &#x60;Waiting&#x60; - Workflow can go to waiting state due to execution of wait task present in the workflow or the workflow has not started yet either due to duplicate workflow is running or due to workflow throttling. Once Workflow engine picks up the workflow for execution, it will move to in progress state. * &#x60;Completed&#x60; - A workflow execution moves into Completed state when the execution path of the workflow has reached the Success node in the workflow design and there are no more tasks to be executed. Completed is the final state for the workflow execution instance and no further actions are allowed on this workflow instance. * &#x60;Failed&#x60; - A workflow execution moves into a Failed state when the execution path of the workflow has reached the Failed node in the workflow design and there are no more tasks to be scheduled. A Failed node can be reached when the last executed task has failed or timed out and there are no further retries available for the task. Also as per the workflow design, the last executed task did not specify an OnFailure task to be executed and hence by default, the execution will reach the Failed node. Actions like \&quot;Rerun\&quot;, \&quot;RetryFailed\&quot; and \&quot;RetryFromTask\&quot; can be issued on failed workflow instances. Please refer to the \&quot;Action\&quot; description for more details. * &#x60;Terminated&#x60; - A workflow execution moves to Terminated state when user issues a \&quot;Cancel\&quot; action or due to internal errors caused during workflow execution. e.g. - Task input transformation has failed. Terminated is a final state of the workflow, no further action are allowed on this workflow instance. * &#x60;Canceled&#x60; - A workflow execution moves to Canceled state when a user issues a \&quot;Cancel\&quot; action. Cancel is not a final state, the workflow engine will issue cancel to all the running tasks and then move the workflow to the \&quot;Terminated\&quot; state. * &#x60;Paused&#x60; - A workflow execution moves to Paused state when user issues a \&quot;Pause\&quot; action. When in paused state the current running task will complete its execution but no further tasks will be scheduled until the workflow is resumed. A paused workflow is resumed when the user issues a \&quot;Resume\&quot; action. Paused workflows can be canceled by user. | [optional] [readonly] [default to "NotStarted"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**AssociatedObject** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ParentTaskInfo** | Pointer to [**NullableWorkflowTaskInfoRelationship**](WorkflowTaskInfoRelationship.md) |  | [optional] 
**Permission** | Pointer to [**NullableIamPermissionRelationship**](IamPermissionRelationship.md) |  | [optional] 
**TaskInfos** | Pointer to [**[]WorkflowTaskInfoRelationship**](WorkflowTaskInfoRelationship.md) | An array of relationships to workflowTaskInfo resources. | [optional] [readonly] 
**WorkflowDefinition** | Pointer to [**NullableWorkflowWorkflowDefinitionRelationship**](WorkflowWorkflowDefinitionRelationship.md) |  | [optional] 

## Methods

### NewWorkflowWorkflowInfo

`func NewWorkflowWorkflowInfo(classId string, objectType string, ) *WorkflowWorkflowInfo`

NewWorkflowWorkflowInfo instantiates a new WorkflowWorkflowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowInfoWithDefaults

`func NewWorkflowWorkflowInfoWithDefaults() *WorkflowWorkflowInfo`

NewWorkflowWorkflowInfoWithDefaults instantiates a new WorkflowWorkflowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWorkflowInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWorkflowInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWorkflowInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWorkflowInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWorkflowInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWorkflowInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *WorkflowWorkflowInfo) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkflowWorkflowInfo) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkflowWorkflowInfo) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkflowWorkflowInfo) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCleanupTime

`func (o *WorkflowWorkflowInfo) GetCleanupTime() time.Time`

GetCleanupTime returns the CleanupTime field if non-nil, zero value otherwise.

### GetCleanupTimeOk

`func (o *WorkflowWorkflowInfo) GetCleanupTimeOk() (*time.Time, bool)`

GetCleanupTimeOk returns a tuple with the CleanupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupTime

`func (o *WorkflowWorkflowInfo) SetCleanupTime(v time.Time)`

SetCleanupTime sets CleanupTime field to given value.

### HasCleanupTime

`func (o *WorkflowWorkflowInfo) HasCleanupTime() bool`

HasCleanupTime returns a boolean if a field has been set.

### GetEmail

`func (o *WorkflowWorkflowInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WorkflowWorkflowInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WorkflowWorkflowInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *WorkflowWorkflowInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowWorkflowInfo) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowWorkflowInfo) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowWorkflowInfo) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowWorkflowInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFailedWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfo) GetFailedWorkflowCleanupDuration() int64`

GetFailedWorkflowCleanupDuration returns the FailedWorkflowCleanupDuration field if non-nil, zero value otherwise.

### GetFailedWorkflowCleanupDurationOk

`func (o *WorkflowWorkflowInfo) GetFailedWorkflowCleanupDurationOk() (*int64, bool)`

GetFailedWorkflowCleanupDurationOk returns a tuple with the FailedWorkflowCleanupDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfo) SetFailedWorkflowCleanupDuration(v int64)`

SetFailedWorkflowCleanupDuration sets FailedWorkflowCleanupDuration field to given value.

### HasFailedWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfo) HasFailedWorkflowCleanupDuration() bool`

HasFailedWorkflowCleanupDuration returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowWorkflowInfo) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowWorkflowInfo) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowWorkflowInfo) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowWorkflowInfo) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowWorkflowInfo) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowWorkflowInfo) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetInstId

`func (o *WorkflowWorkflowInfo) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *WorkflowWorkflowInfo) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *WorkflowWorkflowInfo) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *WorkflowWorkflowInfo) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInternal

`func (o *WorkflowWorkflowInfo) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *WorkflowWorkflowInfo) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *WorkflowWorkflowInfo) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *WorkflowWorkflowInfo) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetLastAction

`func (o *WorkflowWorkflowInfo) GetLastAction() string`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *WorkflowWorkflowInfo) GetLastActionOk() (*string, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *WorkflowWorkflowInfo) SetLastAction(v string)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *WorkflowWorkflowInfo) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### GetMessage

`func (o *WorkflowWorkflowInfo) GetMessage() []WorkflowMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowWorkflowInfo) GetMessageOk() (*[]WorkflowMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowWorkflowInfo) SetMessage(v []WorkflowMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WorkflowWorkflowInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *WorkflowWorkflowInfo) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *WorkflowWorkflowInfo) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetName

`func (o *WorkflowWorkflowInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowWorkflowInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowWorkflowInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowWorkflowInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowWorkflowInfo) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowWorkflowInfo) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowWorkflowInfo) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowWorkflowInfo) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *WorkflowWorkflowInfo) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *WorkflowWorkflowInfo) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetPauseReason

`func (o *WorkflowWorkflowInfo) GetPauseReason() string`

GetPauseReason returns the PauseReason field if non-nil, zero value otherwise.

### GetPauseReasonOk

`func (o *WorkflowWorkflowInfo) GetPauseReasonOk() (*string, bool)`

GetPauseReasonOk returns a tuple with the PauseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseReason

`func (o *WorkflowWorkflowInfo) SetPauseReason(v string)`

SetPauseReason sets PauseReason field to given value.

### HasPauseReason

`func (o *WorkflowWorkflowInfo) HasPauseReason() bool`

HasPauseReason returns a boolean if a field has been set.

### GetProgress

`func (o *WorkflowWorkflowInfo) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *WorkflowWorkflowInfo) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *WorkflowWorkflowInfo) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *WorkflowWorkflowInfo) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowWorkflowInfo) GetProperties() WorkflowWorkflowInfoProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowWorkflowInfo) GetPropertiesOk() (*WorkflowWorkflowInfoProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowWorkflowInfo) SetProperties(v WorkflowWorkflowInfoProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowWorkflowInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowWorkflowInfo) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowWorkflowInfo) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetRetryFromTaskName

`func (o *WorkflowWorkflowInfo) GetRetryFromTaskName() string`

GetRetryFromTaskName returns the RetryFromTaskName field if non-nil, zero value otherwise.

### GetRetryFromTaskNameOk

`func (o *WorkflowWorkflowInfo) GetRetryFromTaskNameOk() (*string, bool)`

GetRetryFromTaskNameOk returns a tuple with the RetryFromTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFromTaskName

`func (o *WorkflowWorkflowInfo) SetRetryFromTaskName(v string)`

SetRetryFromTaskName sets RetryFromTaskName field to given value.

### HasRetryFromTaskName

`func (o *WorkflowWorkflowInfo) HasRetryFromTaskName() bool`

HasRetryFromTaskName returns a boolean if a field has been set.

### GetSrc

`func (o *WorkflowWorkflowInfo) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *WorkflowWorkflowInfo) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *WorkflowWorkflowInfo) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *WorkflowWorkflowInfo) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkflowWorkflowInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowWorkflowInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowWorkflowInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkflowWorkflowInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowWorkflowInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowWorkflowInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowWorkflowInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowWorkflowInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfo) GetSuccessWorkflowCleanupDuration() int64`

GetSuccessWorkflowCleanupDuration returns the SuccessWorkflowCleanupDuration field if non-nil, zero value otherwise.

### GetSuccessWorkflowCleanupDurationOk

`func (o *WorkflowWorkflowInfo) GetSuccessWorkflowCleanupDurationOk() (*int64, bool)`

GetSuccessWorkflowCleanupDurationOk returns a tuple with the SuccessWorkflowCleanupDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfo) SetSuccessWorkflowCleanupDuration(v int64)`

SetSuccessWorkflowCleanupDuration sets SuccessWorkflowCleanupDuration field to given value.

### HasSuccessWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfo) HasSuccessWorkflowCleanupDuration() bool`

HasSuccessWorkflowCleanupDuration returns a boolean if a field has been set.

### GetTaskInfoUpdate

`func (o *WorkflowWorkflowInfo) GetTaskInfoUpdate() WorkflowTaskInfoUpdate`

GetTaskInfoUpdate returns the TaskInfoUpdate field if non-nil, zero value otherwise.

### GetTaskInfoUpdateOk

`func (o *WorkflowWorkflowInfo) GetTaskInfoUpdateOk() (*WorkflowTaskInfoUpdate, bool)`

GetTaskInfoUpdateOk returns a tuple with the TaskInfoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInfoUpdate

`func (o *WorkflowWorkflowInfo) SetTaskInfoUpdate(v WorkflowTaskInfoUpdate)`

SetTaskInfoUpdate sets TaskInfoUpdate field to given value.

### HasTaskInfoUpdate

`func (o *WorkflowWorkflowInfo) HasTaskInfoUpdate() bool`

HasTaskInfoUpdate returns a boolean if a field has been set.

### SetTaskInfoUpdateNil

`func (o *WorkflowWorkflowInfo) SetTaskInfoUpdateNil(b bool)`

 SetTaskInfoUpdateNil sets the value for TaskInfoUpdate to be an explicit nil

### UnsetTaskInfoUpdate
`func (o *WorkflowWorkflowInfo) UnsetTaskInfoUpdate()`

UnsetTaskInfoUpdate ensures that no value is present for TaskInfoUpdate, not even an explicit nil
### GetTraceId

`func (o *WorkflowWorkflowInfo) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *WorkflowWorkflowInfo) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *WorkflowWorkflowInfo) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *WorkflowWorkflowInfo) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *WorkflowWorkflowInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowWorkflowInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowWorkflowInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowWorkflowInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserActionRequired

`func (o *WorkflowWorkflowInfo) GetUserActionRequired() bool`

GetUserActionRequired returns the UserActionRequired field if non-nil, zero value otherwise.

### GetUserActionRequiredOk

`func (o *WorkflowWorkflowInfo) GetUserActionRequiredOk() (*bool, bool)`

GetUserActionRequiredOk returns a tuple with the UserActionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionRequired

`func (o *WorkflowWorkflowInfo) SetUserActionRequired(v bool)`

SetUserActionRequired sets UserActionRequired field to given value.

### HasUserActionRequired

`func (o *WorkflowWorkflowInfo) HasUserActionRequired() bool`

HasUserActionRequired returns a boolean if a field has been set.

### GetUserId

`func (o *WorkflowWorkflowInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WorkflowWorkflowInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WorkflowWorkflowInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WorkflowWorkflowInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVariable

`func (o *WorkflowWorkflowInfo) GetVariable() interface{}`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *WorkflowWorkflowInfo) GetVariableOk() (*interface{}, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *WorkflowWorkflowInfo) SetVariable(v interface{})`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *WorkflowWorkflowInfo) HasVariable() bool`

HasVariable returns a boolean if a field has been set.

### SetVariableNil

`func (o *WorkflowWorkflowInfo) SetVariableNil(b bool)`

 SetVariableNil sets the value for Variable to be an explicit nil

### UnsetVariable
`func (o *WorkflowWorkflowInfo) UnsetVariable()`

UnsetVariable ensures that no value is present for Variable, not even an explicit nil
### GetWaitReason

`func (o *WorkflowWorkflowInfo) GetWaitReason() string`

GetWaitReason returns the WaitReason field if non-nil, zero value otherwise.

### GetWaitReasonOk

`func (o *WorkflowWorkflowInfo) GetWaitReasonOk() (*string, bool)`

GetWaitReasonOk returns a tuple with the WaitReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitReason

`func (o *WorkflowWorkflowInfo) SetWaitReason(v string)`

SetWaitReason sets WaitReason field to given value.

### HasWaitReason

`func (o *WorkflowWorkflowInfo) HasWaitReason() bool`

HasWaitReason returns a boolean if a field has been set.

### GetWorkflowCtx

`func (o *WorkflowWorkflowInfo) GetWorkflowCtx() WorkflowWorkflowCtx`

GetWorkflowCtx returns the WorkflowCtx field if non-nil, zero value otherwise.

### GetWorkflowCtxOk

`func (o *WorkflowWorkflowInfo) GetWorkflowCtxOk() (*WorkflowWorkflowCtx, bool)`

GetWorkflowCtxOk returns a tuple with the WorkflowCtx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowCtx

`func (o *WorkflowWorkflowInfo) SetWorkflowCtx(v WorkflowWorkflowCtx)`

SetWorkflowCtx sets WorkflowCtx field to given value.

### HasWorkflowCtx

`func (o *WorkflowWorkflowInfo) HasWorkflowCtx() bool`

HasWorkflowCtx returns a boolean if a field has been set.

### SetWorkflowCtxNil

`func (o *WorkflowWorkflowInfo) SetWorkflowCtxNil(b bool)`

 SetWorkflowCtxNil sets the value for WorkflowCtx to be an explicit nil

### UnsetWorkflowCtx
`func (o *WorkflowWorkflowInfo) UnsetWorkflowCtx()`

UnsetWorkflowCtx ensures that no value is present for WorkflowCtx, not even an explicit nil
### GetWorkflowStatus

`func (o *WorkflowWorkflowInfo) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *WorkflowWorkflowInfo) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *WorkflowWorkflowInfo) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *WorkflowWorkflowInfo) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetAccount

`func (o *WorkflowWorkflowInfo) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WorkflowWorkflowInfo) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WorkflowWorkflowInfo) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *WorkflowWorkflowInfo) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *WorkflowWorkflowInfo) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *WorkflowWorkflowInfo) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetAssociatedObject

`func (o *WorkflowWorkflowInfo) GetAssociatedObject() MoBaseMoRelationship`

GetAssociatedObject returns the AssociatedObject field if non-nil, zero value otherwise.

### GetAssociatedObjectOk

`func (o *WorkflowWorkflowInfo) GetAssociatedObjectOk() (*MoBaseMoRelationship, bool)`

GetAssociatedObjectOk returns a tuple with the AssociatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObject

`func (o *WorkflowWorkflowInfo) SetAssociatedObject(v MoBaseMoRelationship)`

SetAssociatedObject sets AssociatedObject field to given value.

### HasAssociatedObject

`func (o *WorkflowWorkflowInfo) HasAssociatedObject() bool`

HasAssociatedObject returns a boolean if a field has been set.

### SetAssociatedObjectNil

`func (o *WorkflowWorkflowInfo) SetAssociatedObjectNil(b bool)`

 SetAssociatedObjectNil sets the value for AssociatedObject to be an explicit nil

### UnsetAssociatedObject
`func (o *WorkflowWorkflowInfo) UnsetAssociatedObject()`

UnsetAssociatedObject ensures that no value is present for AssociatedObject, not even an explicit nil
### GetOrganization

`func (o *WorkflowWorkflowInfo) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowWorkflowInfo) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowWorkflowInfo) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowWorkflowInfo) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *WorkflowWorkflowInfo) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *WorkflowWorkflowInfo) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetParentTaskInfo

`func (o *WorkflowWorkflowInfo) GetParentTaskInfo() WorkflowTaskInfoRelationship`

GetParentTaskInfo returns the ParentTaskInfo field if non-nil, zero value otherwise.

### GetParentTaskInfoOk

`func (o *WorkflowWorkflowInfo) GetParentTaskInfoOk() (*WorkflowTaskInfoRelationship, bool)`

GetParentTaskInfoOk returns a tuple with the ParentTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskInfo

`func (o *WorkflowWorkflowInfo) SetParentTaskInfo(v WorkflowTaskInfoRelationship)`

SetParentTaskInfo sets ParentTaskInfo field to given value.

### HasParentTaskInfo

`func (o *WorkflowWorkflowInfo) HasParentTaskInfo() bool`

HasParentTaskInfo returns a boolean if a field has been set.

### SetParentTaskInfoNil

`func (o *WorkflowWorkflowInfo) SetParentTaskInfoNil(b bool)`

 SetParentTaskInfoNil sets the value for ParentTaskInfo to be an explicit nil

### UnsetParentTaskInfo
`func (o *WorkflowWorkflowInfo) UnsetParentTaskInfo()`

UnsetParentTaskInfo ensures that no value is present for ParentTaskInfo, not even an explicit nil
### GetPermission

`func (o *WorkflowWorkflowInfo) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *WorkflowWorkflowInfo) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *WorkflowWorkflowInfo) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *WorkflowWorkflowInfo) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### SetPermissionNil

`func (o *WorkflowWorkflowInfo) SetPermissionNil(b bool)`

 SetPermissionNil sets the value for Permission to be an explicit nil

### UnsetPermission
`func (o *WorkflowWorkflowInfo) UnsetPermission()`

UnsetPermission ensures that no value is present for Permission, not even an explicit nil
### GetTaskInfos

`func (o *WorkflowWorkflowInfo) GetTaskInfos() []WorkflowTaskInfoRelationship`

GetTaskInfos returns the TaskInfos field if non-nil, zero value otherwise.

### GetTaskInfosOk

`func (o *WorkflowWorkflowInfo) GetTaskInfosOk() (*[]WorkflowTaskInfoRelationship, bool)`

GetTaskInfosOk returns a tuple with the TaskInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInfos

`func (o *WorkflowWorkflowInfo) SetTaskInfos(v []WorkflowTaskInfoRelationship)`

SetTaskInfos sets TaskInfos field to given value.

### HasTaskInfos

`func (o *WorkflowWorkflowInfo) HasTaskInfos() bool`

HasTaskInfos returns a boolean if a field has been set.

### SetTaskInfosNil

`func (o *WorkflowWorkflowInfo) SetTaskInfosNil(b bool)`

 SetTaskInfosNil sets the value for TaskInfos to be an explicit nil

### UnsetTaskInfos
`func (o *WorkflowWorkflowInfo) UnsetTaskInfos()`

UnsetTaskInfos ensures that no value is present for TaskInfos, not even an explicit nil
### GetWorkflowDefinition

`func (o *WorkflowWorkflowInfo) GetWorkflowDefinition() WorkflowWorkflowDefinitionRelationship`

GetWorkflowDefinition returns the WorkflowDefinition field if non-nil, zero value otherwise.

### GetWorkflowDefinitionOk

`func (o *WorkflowWorkflowInfo) GetWorkflowDefinitionOk() (*WorkflowWorkflowDefinitionRelationship, bool)`

GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinition

`func (o *WorkflowWorkflowInfo) SetWorkflowDefinition(v WorkflowWorkflowDefinitionRelationship)`

SetWorkflowDefinition sets WorkflowDefinition field to given value.

### HasWorkflowDefinition

`func (o *WorkflowWorkflowInfo) HasWorkflowDefinition() bool`

HasWorkflowDefinition returns a boolean if a field has been set.

### SetWorkflowDefinitionNil

`func (o *WorkflowWorkflowInfo) SetWorkflowDefinitionNil(b bool)`

 SetWorkflowDefinitionNil sets the value for WorkflowDefinition to be an explicit nil

### UnsetWorkflowDefinition
`func (o *WorkflowWorkflowInfo) UnsetWorkflowDefinition()`

UnsetWorkflowDefinition ensures that no value is present for WorkflowDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


