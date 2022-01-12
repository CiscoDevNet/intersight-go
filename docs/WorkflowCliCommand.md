# WorkflowCliCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.CliCommand"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.CliCommand"]
**Command** | Pointer to **string** | The command to run on the device connector. | [optional] 
**EndPrompt** | Pointer to **string** | The regex string that identifies the end of the command response. | [optional] 
**ExpectPrompts** | Pointer to [**[]WorkflowExpectPrompt**](WorkflowExpectPrompt.md) |  | [optional] 
**ExpectedExitCodes** | Pointer to **[]int64** |  | [optional] 
**SkipStatusCheck** | Pointer to **bool** | Skips the execution status check of the terminal command. One use case for this is while exiting the terminal session from esxi host. | [optional] 
**TerminalEnd** | Pointer to **bool** | If this flag is set, it marks the end of the terminal session where the previous commands were executed. | [optional] 
**TerminalStart** | Pointer to **bool** | If this flag is set, the execution of this command initiates a terminal session in which the subsequent CLI commands are executed until a command with terminalEnd flag is encountered or the end of the batch. | [optional] 
**Type** | Pointer to **string** | The type of the command - can be interactive or non-interactive. * &#x60;NonInteractive&#x60; - The CLI command is not an interactive command. * &#x60;Interactive&#x60; - The CLI command is executed in interactive mode and the command must provide the expects andanswers. | [optional] [default to "NonInteractive"]

## Methods

### NewWorkflowCliCommand

`func NewWorkflowCliCommand(classId string, objectType string, ) *WorkflowCliCommand`

NewWorkflowCliCommand instantiates a new WorkflowCliCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCliCommandWithDefaults

`func NewWorkflowCliCommandWithDefaults() *WorkflowCliCommand`

NewWorkflowCliCommandWithDefaults instantiates a new WorkflowCliCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowCliCommand) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowCliCommand) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowCliCommand) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowCliCommand) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowCliCommand) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowCliCommand) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommand

`func (o *WorkflowCliCommand) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WorkflowCliCommand) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WorkflowCliCommand) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *WorkflowCliCommand) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEndPrompt

`func (o *WorkflowCliCommand) GetEndPrompt() string`

GetEndPrompt returns the EndPrompt field if non-nil, zero value otherwise.

### GetEndPromptOk

`func (o *WorkflowCliCommand) GetEndPromptOk() (*string, bool)`

GetEndPromptOk returns a tuple with the EndPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPrompt

`func (o *WorkflowCliCommand) SetEndPrompt(v string)`

SetEndPrompt sets EndPrompt field to given value.

### HasEndPrompt

`func (o *WorkflowCliCommand) HasEndPrompt() bool`

HasEndPrompt returns a boolean if a field has been set.

### GetExpectPrompts

`func (o *WorkflowCliCommand) GetExpectPrompts() []WorkflowExpectPrompt`

GetExpectPrompts returns the ExpectPrompts field if non-nil, zero value otherwise.

### GetExpectPromptsOk

`func (o *WorkflowCliCommand) GetExpectPromptsOk() (*[]WorkflowExpectPrompt, bool)`

GetExpectPromptsOk returns a tuple with the ExpectPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectPrompts

`func (o *WorkflowCliCommand) SetExpectPrompts(v []WorkflowExpectPrompt)`

SetExpectPrompts sets ExpectPrompts field to given value.

### HasExpectPrompts

`func (o *WorkflowCliCommand) HasExpectPrompts() bool`

HasExpectPrompts returns a boolean if a field has been set.

### SetExpectPromptsNil

`func (o *WorkflowCliCommand) SetExpectPromptsNil(b bool)`

 SetExpectPromptsNil sets the value for ExpectPrompts to be an explicit nil

### UnsetExpectPrompts
`func (o *WorkflowCliCommand) UnsetExpectPrompts()`

UnsetExpectPrompts ensures that no value is present for ExpectPrompts, not even an explicit nil
### GetExpectedExitCodes

`func (o *WorkflowCliCommand) GetExpectedExitCodes() []int64`

GetExpectedExitCodes returns the ExpectedExitCodes field if non-nil, zero value otherwise.

### GetExpectedExitCodesOk

`func (o *WorkflowCliCommand) GetExpectedExitCodesOk() (*[]int64, bool)`

GetExpectedExitCodesOk returns a tuple with the ExpectedExitCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedExitCodes

`func (o *WorkflowCliCommand) SetExpectedExitCodes(v []int64)`

SetExpectedExitCodes sets ExpectedExitCodes field to given value.

### HasExpectedExitCodes

`func (o *WorkflowCliCommand) HasExpectedExitCodes() bool`

HasExpectedExitCodes returns a boolean if a field has been set.

### SetExpectedExitCodesNil

`func (o *WorkflowCliCommand) SetExpectedExitCodesNil(b bool)`

 SetExpectedExitCodesNil sets the value for ExpectedExitCodes to be an explicit nil

### UnsetExpectedExitCodes
`func (o *WorkflowCliCommand) UnsetExpectedExitCodes()`

UnsetExpectedExitCodes ensures that no value is present for ExpectedExitCodes, not even an explicit nil
### GetSkipStatusCheck

`func (o *WorkflowCliCommand) GetSkipStatusCheck() bool`

GetSkipStatusCheck returns the SkipStatusCheck field if non-nil, zero value otherwise.

### GetSkipStatusCheckOk

`func (o *WorkflowCliCommand) GetSkipStatusCheckOk() (*bool, bool)`

GetSkipStatusCheckOk returns a tuple with the SkipStatusCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipStatusCheck

`func (o *WorkflowCliCommand) SetSkipStatusCheck(v bool)`

SetSkipStatusCheck sets SkipStatusCheck field to given value.

### HasSkipStatusCheck

`func (o *WorkflowCliCommand) HasSkipStatusCheck() bool`

HasSkipStatusCheck returns a boolean if a field has been set.

### GetTerminalEnd

`func (o *WorkflowCliCommand) GetTerminalEnd() bool`

GetTerminalEnd returns the TerminalEnd field if non-nil, zero value otherwise.

### GetTerminalEndOk

`func (o *WorkflowCliCommand) GetTerminalEndOk() (*bool, bool)`

GetTerminalEndOk returns a tuple with the TerminalEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalEnd

`func (o *WorkflowCliCommand) SetTerminalEnd(v bool)`

SetTerminalEnd sets TerminalEnd field to given value.

### HasTerminalEnd

`func (o *WorkflowCliCommand) HasTerminalEnd() bool`

HasTerminalEnd returns a boolean if a field has been set.

### GetTerminalStart

`func (o *WorkflowCliCommand) GetTerminalStart() bool`

GetTerminalStart returns the TerminalStart field if non-nil, zero value otherwise.

### GetTerminalStartOk

`func (o *WorkflowCliCommand) GetTerminalStartOk() (*bool, bool)`

GetTerminalStartOk returns a tuple with the TerminalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalStart

`func (o *WorkflowCliCommand) SetTerminalStart(v bool)`

SetTerminalStart sets TerminalStart field to given value.

### HasTerminalStart

`func (o *WorkflowCliCommand) HasTerminalStart() bool`

HasTerminalStart returns a boolean if a field has been set.

### GetType

`func (o *WorkflowCliCommand) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowCliCommand) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowCliCommand) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowCliCommand) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


