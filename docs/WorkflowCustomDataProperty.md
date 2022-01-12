# WorkflowCustomDataProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.CustomDataProperty"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.CustomDataProperty"]
**CatalogMoid** | Pointer to **string** | Specify the catalog moid that this custom data type belongs. | [optional] 
**CustomDataTypeId** | Pointer to **string** | The resolved custom data type definition managed object. | [optional] [readonly] 
**CustomDataTypeName** | Pointer to **string** | Name of the custom data type for this input. | [optional] 

## Methods

### NewWorkflowCustomDataProperty

`func NewWorkflowCustomDataProperty(classId string, objectType string, ) *WorkflowCustomDataProperty`

NewWorkflowCustomDataProperty instantiates a new WorkflowCustomDataProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCustomDataPropertyWithDefaults

`func NewWorkflowCustomDataPropertyWithDefaults() *WorkflowCustomDataProperty`

NewWorkflowCustomDataPropertyWithDefaults instantiates a new WorkflowCustomDataProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowCustomDataProperty) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowCustomDataProperty) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowCustomDataProperty) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowCustomDataProperty) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowCustomDataProperty) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowCustomDataProperty) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCatalogMoid

`func (o *WorkflowCustomDataProperty) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *WorkflowCustomDataProperty) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *WorkflowCustomDataProperty) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *WorkflowCustomDataProperty) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetCustomDataTypeId

`func (o *WorkflowCustomDataProperty) GetCustomDataTypeId() string`

GetCustomDataTypeId returns the CustomDataTypeId field if non-nil, zero value otherwise.

### GetCustomDataTypeIdOk

`func (o *WorkflowCustomDataProperty) GetCustomDataTypeIdOk() (*string, bool)`

GetCustomDataTypeIdOk returns a tuple with the CustomDataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDataTypeId

`func (o *WorkflowCustomDataProperty) SetCustomDataTypeId(v string)`

SetCustomDataTypeId sets CustomDataTypeId field to given value.

### HasCustomDataTypeId

`func (o *WorkflowCustomDataProperty) HasCustomDataTypeId() bool`

HasCustomDataTypeId returns a boolean if a field has been set.

### GetCustomDataTypeName

`func (o *WorkflowCustomDataProperty) GetCustomDataTypeName() string`

GetCustomDataTypeName returns the CustomDataTypeName field if non-nil, zero value otherwise.

### GetCustomDataTypeNameOk

`func (o *WorkflowCustomDataProperty) GetCustomDataTypeNameOk() (*string, bool)`

GetCustomDataTypeNameOk returns a tuple with the CustomDataTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDataTypeName

`func (o *WorkflowCustomDataProperty) SetCustomDataTypeName(v string)`

SetCustomDataTypeName sets CustomDataTypeName field to given value.

### HasCustomDataTypeName

`func (o *WorkflowCustomDataProperty) HasCustomDataTypeName() bool`

HasCustomDataTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


