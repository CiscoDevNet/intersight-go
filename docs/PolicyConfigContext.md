# PolicyConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.ConfigContext"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.ConfigContext"]
**ConfigState** | Pointer to **string** | Indicates a profile&#39;s configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Out-of-sync, Validating, Configuring, Failed. | [optional] [readonly] 
**ConfigStateSummary** | Pointer to **string** | Indicates a profile&#39;s configuration deploying state. Values -- Assigned, Not-assigned, Associated, InConsistent, Validating, Configuring, Failed, Activating, UnConfiguring. * &#x60;None&#x60; - The default state is none. * &#x60;Not-assigned&#x60; - Server is not assigned to the profile. * &#x60;Assigned&#x60; - Server is assigned to the profile and the configurations are not yet deployed. * &#x60;Preparing&#x60; - Preparing to deploy the configuration. * &#x60;Validating&#x60; - Profile validation in progress. * &#x60;Configuring&#x60; - Profile deploy operation is in progress. * &#x60;UnConfiguring&#x60; - Server is unassigned and config cleanup is in progress. * &#x60;Analyzing&#x60; - Profile changes are being analyzed. * &#x60;Activating&#x60; - Configuration is being activated at the endpoint. * &#x60;Inconsistent&#x60; - Profile is inconsistent with the endpoint configuration. * &#x60;Associated&#x60; - The profile configuration has been applied to the endpoint and no inconsistencies have been detected. * &#x60;Failed&#x60; - The last action on the profile has failed. * &#x60;Not-complete&#x60; - Config import operation on the profile is not complete. * &#x60;Waiting-for-resource&#x60; - Waiting for the resource to be allocated for the profile. * &#x60;Partially-deployed&#x60; - The profile configuration has been applied on a subset of endpoints. | [optional] [readonly] [default to "None"]
**ConfigType** | Pointer to **string** | The type of configuration running on the profile. Since profile deployments can configure multiple different settings, configType indicates which type of configuration is currently in progress. | [optional] [readonly] 
**ControlAction** | Pointer to **string** | System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind. | [optional] 
**ErrorState** | Pointer to **string** | Indicates a profile&#39;s error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error). | [optional] 
**InconsistencyReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed. | [optional] [readonly] 

## Methods

### NewPolicyConfigContext

`func NewPolicyConfigContext(classId string, objectType string, ) *PolicyConfigContext`

NewPolicyConfigContext instantiates a new PolicyConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyConfigContextWithDefaults

`func NewPolicyConfigContextWithDefaults() *PolicyConfigContext`

NewPolicyConfigContextWithDefaults instantiates a new PolicyConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyConfigContext) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyConfigContext) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyConfigContext) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyConfigContext) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyConfigContext) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyConfigContext) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigState

`func (o *PolicyConfigContext) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *PolicyConfigContext) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *PolicyConfigContext) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *PolicyConfigContext) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetConfigStateSummary

`func (o *PolicyConfigContext) GetConfigStateSummary() string`

GetConfigStateSummary returns the ConfigStateSummary field if non-nil, zero value otherwise.

### GetConfigStateSummaryOk

`func (o *PolicyConfigContext) GetConfigStateSummaryOk() (*string, bool)`

GetConfigStateSummaryOk returns a tuple with the ConfigStateSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStateSummary

`func (o *PolicyConfigContext) SetConfigStateSummary(v string)`

SetConfigStateSummary sets ConfigStateSummary field to given value.

### HasConfigStateSummary

`func (o *PolicyConfigContext) HasConfigStateSummary() bool`

HasConfigStateSummary returns a boolean if a field has been set.

### GetConfigType

`func (o *PolicyConfigContext) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *PolicyConfigContext) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *PolicyConfigContext) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.

### HasConfigType

`func (o *PolicyConfigContext) HasConfigType() bool`

HasConfigType returns a boolean if a field has been set.

### GetControlAction

`func (o *PolicyConfigContext) GetControlAction() string`

GetControlAction returns the ControlAction field if non-nil, zero value otherwise.

### GetControlActionOk

`func (o *PolicyConfigContext) GetControlActionOk() (*string, bool)`

GetControlActionOk returns a tuple with the ControlAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlAction

`func (o *PolicyConfigContext) SetControlAction(v string)`

SetControlAction sets ControlAction field to given value.

### HasControlAction

`func (o *PolicyConfigContext) HasControlAction() bool`

HasControlAction returns a boolean if a field has been set.

### GetErrorState

`func (o *PolicyConfigContext) GetErrorState() string`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *PolicyConfigContext) GetErrorStateOk() (*string, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *PolicyConfigContext) SetErrorState(v string)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *PolicyConfigContext) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.

### GetInconsistencyReason

`func (o *PolicyConfigContext) GetInconsistencyReason() []string`

GetInconsistencyReason returns the InconsistencyReason field if non-nil, zero value otherwise.

### GetInconsistencyReasonOk

`func (o *PolicyConfigContext) GetInconsistencyReasonOk() (*[]string, bool)`

GetInconsistencyReasonOk returns a tuple with the InconsistencyReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInconsistencyReason

`func (o *PolicyConfigContext) SetInconsistencyReason(v []string)`

SetInconsistencyReason sets InconsistencyReason field to given value.

### HasInconsistencyReason

`func (o *PolicyConfigContext) HasInconsistencyReason() bool`

HasInconsistencyReason returns a boolean if a field has been set.

### SetInconsistencyReasonNil

`func (o *PolicyConfigContext) SetInconsistencyReasonNil(b bool)`

 SetInconsistencyReasonNil sets the value for InconsistencyReason to be an explicit nil

### UnsetInconsistencyReason
`func (o *PolicyConfigContext) UnsetInconsistencyReason()`

UnsetInconsistencyReason ensures that no value is present for InconsistencyReason, not even an explicit nil
### GetOperState

`func (o *PolicyConfigContext) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PolicyConfigContext) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PolicyConfigContext) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PolicyConfigContext) HasOperState() bool`

HasOperState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


