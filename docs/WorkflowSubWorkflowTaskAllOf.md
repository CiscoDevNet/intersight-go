# WorkflowSubWorkflowTaskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SubWorkflowTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SubWorkflowTask"]
**CatalogMoid** | Pointer to **string** | Specify the catalog moid that this task belongs. | [optional] 
**Version** | Pointer to **int64** | The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used. | [optional] 
**WorkflowDefinitionId** | Pointer to **string** | The resolved referenced workflow definition managed object. | [optional] [readonly] 
**WorkflowDefinitionName** | Pointer to **string** | The qualified name of workflow that should be executed as a task. | [optional] 

## Methods

### NewWorkflowSubWorkflowTaskAllOf

`func NewWorkflowSubWorkflowTaskAllOf(classId string, objectType string, ) *WorkflowSubWorkflowTaskAllOf`

NewWorkflowSubWorkflowTaskAllOf instantiates a new WorkflowSubWorkflowTaskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSubWorkflowTaskAllOfWithDefaults

`func NewWorkflowSubWorkflowTaskAllOfWithDefaults() *WorkflowSubWorkflowTaskAllOf`

NewWorkflowSubWorkflowTaskAllOfWithDefaults instantiates a new WorkflowSubWorkflowTaskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSubWorkflowTaskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSubWorkflowTaskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSubWorkflowTaskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSubWorkflowTaskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSubWorkflowTaskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSubWorkflowTaskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCatalogMoid

`func (o *WorkflowSubWorkflowTaskAllOf) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *WorkflowSubWorkflowTaskAllOf) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *WorkflowSubWorkflowTaskAllOf) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *WorkflowSubWorkflowTaskAllOf) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowSubWorkflowTaskAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowSubWorkflowTaskAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowSubWorkflowTaskAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowSubWorkflowTaskAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowDefinitionId

`func (o *WorkflowSubWorkflowTaskAllOf) GetWorkflowDefinitionId() string`

GetWorkflowDefinitionId returns the WorkflowDefinitionId field if non-nil, zero value otherwise.

### GetWorkflowDefinitionIdOk

`func (o *WorkflowSubWorkflowTaskAllOf) GetWorkflowDefinitionIdOk() (*string, bool)`

GetWorkflowDefinitionIdOk returns a tuple with the WorkflowDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinitionId

`func (o *WorkflowSubWorkflowTaskAllOf) SetWorkflowDefinitionId(v string)`

SetWorkflowDefinitionId sets WorkflowDefinitionId field to given value.

### HasWorkflowDefinitionId

`func (o *WorkflowSubWorkflowTaskAllOf) HasWorkflowDefinitionId() bool`

HasWorkflowDefinitionId returns a boolean if a field has been set.

### GetWorkflowDefinitionName

`func (o *WorkflowSubWorkflowTaskAllOf) GetWorkflowDefinitionName() string`

GetWorkflowDefinitionName returns the WorkflowDefinitionName field if non-nil, zero value otherwise.

### GetWorkflowDefinitionNameOk

`func (o *WorkflowSubWorkflowTaskAllOf) GetWorkflowDefinitionNameOk() (*string, bool)`

GetWorkflowDefinitionNameOk returns a tuple with the WorkflowDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinitionName

`func (o *WorkflowSubWorkflowTaskAllOf) SetWorkflowDefinitionName(v string)`

SetWorkflowDefinitionName sets WorkflowDefinitionName field to given value.

### HasWorkflowDefinitionName

`func (o *WorkflowSubWorkflowTaskAllOf) HasWorkflowDefinitionName() bool`

HasWorkflowDefinitionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


