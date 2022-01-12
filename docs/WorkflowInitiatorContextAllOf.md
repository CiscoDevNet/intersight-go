# WorkflowInitiatorContextAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.InitiatorContext"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.InitiatorContext"]
**InitiatorMoid** | Pointer to **string** | The moid of the Intersight managed object that initiated the workflow. | [optional] 
**InitiatorName** | Pointer to **string** | Name of the initiator who started the workflow. The initiator can be Intersight managed object that triggered the workflow. | [optional] 
**InitiatorType** | Pointer to **string** | Type of Intersight managed object that initiated the workflow. | [optional] 

## Methods

### NewWorkflowInitiatorContextAllOf

`func NewWorkflowInitiatorContextAllOf(classId string, objectType string, ) *WorkflowInitiatorContextAllOf`

NewWorkflowInitiatorContextAllOf instantiates a new WorkflowInitiatorContextAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowInitiatorContextAllOfWithDefaults

`func NewWorkflowInitiatorContextAllOfWithDefaults() *WorkflowInitiatorContextAllOf`

NewWorkflowInitiatorContextAllOfWithDefaults instantiates a new WorkflowInitiatorContextAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowInitiatorContextAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowInitiatorContextAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowInitiatorContextAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowInitiatorContextAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowInitiatorContextAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowInitiatorContextAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInitiatorMoid

`func (o *WorkflowInitiatorContextAllOf) GetInitiatorMoid() string`

GetInitiatorMoid returns the InitiatorMoid field if non-nil, zero value otherwise.

### GetInitiatorMoidOk

`func (o *WorkflowInitiatorContextAllOf) GetInitiatorMoidOk() (*string, bool)`

GetInitiatorMoidOk returns a tuple with the InitiatorMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorMoid

`func (o *WorkflowInitiatorContextAllOf) SetInitiatorMoid(v string)`

SetInitiatorMoid sets InitiatorMoid field to given value.

### HasInitiatorMoid

`func (o *WorkflowInitiatorContextAllOf) HasInitiatorMoid() bool`

HasInitiatorMoid returns a boolean if a field has been set.

### GetInitiatorName

`func (o *WorkflowInitiatorContextAllOf) GetInitiatorName() string`

GetInitiatorName returns the InitiatorName field if non-nil, zero value otherwise.

### GetInitiatorNameOk

`func (o *WorkflowInitiatorContextAllOf) GetInitiatorNameOk() (*string, bool)`

GetInitiatorNameOk returns a tuple with the InitiatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorName

`func (o *WorkflowInitiatorContextAllOf) SetInitiatorName(v string)`

SetInitiatorName sets InitiatorName field to given value.

### HasInitiatorName

`func (o *WorkflowInitiatorContextAllOf) HasInitiatorName() bool`

HasInitiatorName returns a boolean if a field has been set.

### GetInitiatorType

`func (o *WorkflowInitiatorContextAllOf) GetInitiatorType() string`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *WorkflowInitiatorContextAllOf) GetInitiatorTypeOk() (*string, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *WorkflowInitiatorContextAllOf) SetInitiatorType(v string)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *WorkflowInitiatorContextAllOf) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


