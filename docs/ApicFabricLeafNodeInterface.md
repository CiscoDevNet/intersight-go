# ApicFabricLeafNodeInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "apic.FabricLeafNodeInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "apic.FabricLeafNodeInterface"]
**Dn** | Pointer to **string** | Distinguished Name (DN) of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 
**FabricLeafNodeDn** | Pointer to **string** | Fabric leaf node Distinguished Name (DN). | [optional] 
**FabricLeafNodeId** | Pointer to **string** | Fabric leaf node identification number. | [optional] 
**Name** | Pointer to **string** | Object name in Cisco Application Policy Infrastructure Controller (APIC). | [optional] 
**FabricLeafNode** | Pointer to [**NullableApicFabricLeafNodeRelationship**](ApicFabricLeafNodeRelationship.md) |  | [optional] 

## Methods

### NewApicFabricLeafNodeInterface

`func NewApicFabricLeafNodeInterface(classId string, objectType string, ) *ApicFabricLeafNodeInterface`

NewApicFabricLeafNodeInterface instantiates a new ApicFabricLeafNodeInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApicFabricLeafNodeInterfaceWithDefaults

`func NewApicFabricLeafNodeInterfaceWithDefaults() *ApicFabricLeafNodeInterface`

NewApicFabricLeafNodeInterfaceWithDefaults instantiates a new ApicFabricLeafNodeInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApicFabricLeafNodeInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApicFabricLeafNodeInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApicFabricLeafNodeInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApicFabricLeafNodeInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApicFabricLeafNodeInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApicFabricLeafNodeInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *ApicFabricLeafNodeInterface) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ApicFabricLeafNodeInterface) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ApicFabricLeafNodeInterface) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ApicFabricLeafNodeInterface) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFabricLeafNodeDn

`func (o *ApicFabricLeafNodeInterface) GetFabricLeafNodeDn() string`

GetFabricLeafNodeDn returns the FabricLeafNodeDn field if non-nil, zero value otherwise.

### GetFabricLeafNodeDnOk

`func (o *ApicFabricLeafNodeInterface) GetFabricLeafNodeDnOk() (*string, bool)`

GetFabricLeafNodeDnOk returns a tuple with the FabricLeafNodeDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricLeafNodeDn

`func (o *ApicFabricLeafNodeInterface) SetFabricLeafNodeDn(v string)`

SetFabricLeafNodeDn sets FabricLeafNodeDn field to given value.

### HasFabricLeafNodeDn

`func (o *ApicFabricLeafNodeInterface) HasFabricLeafNodeDn() bool`

HasFabricLeafNodeDn returns a boolean if a field has been set.

### GetFabricLeafNodeId

`func (o *ApicFabricLeafNodeInterface) GetFabricLeafNodeId() string`

GetFabricLeafNodeId returns the FabricLeafNodeId field if non-nil, zero value otherwise.

### GetFabricLeafNodeIdOk

`func (o *ApicFabricLeafNodeInterface) GetFabricLeafNodeIdOk() (*string, bool)`

GetFabricLeafNodeIdOk returns a tuple with the FabricLeafNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricLeafNodeId

`func (o *ApicFabricLeafNodeInterface) SetFabricLeafNodeId(v string)`

SetFabricLeafNodeId sets FabricLeafNodeId field to given value.

### HasFabricLeafNodeId

`func (o *ApicFabricLeafNodeInterface) HasFabricLeafNodeId() bool`

HasFabricLeafNodeId returns a boolean if a field has been set.

### GetName

`func (o *ApicFabricLeafNodeInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApicFabricLeafNodeInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApicFabricLeafNodeInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApicFabricLeafNodeInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFabricLeafNode

`func (o *ApicFabricLeafNodeInterface) GetFabricLeafNode() ApicFabricLeafNodeRelationship`

GetFabricLeafNode returns the FabricLeafNode field if non-nil, zero value otherwise.

### GetFabricLeafNodeOk

`func (o *ApicFabricLeafNodeInterface) GetFabricLeafNodeOk() (*ApicFabricLeafNodeRelationship, bool)`

GetFabricLeafNodeOk returns a tuple with the FabricLeafNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricLeafNode

`func (o *ApicFabricLeafNodeInterface) SetFabricLeafNode(v ApicFabricLeafNodeRelationship)`

SetFabricLeafNode sets FabricLeafNode field to given value.

### HasFabricLeafNode

`func (o *ApicFabricLeafNodeInterface) HasFabricLeafNode() bool`

HasFabricLeafNode returns a boolean if a field has been set.

### SetFabricLeafNodeNil

`func (o *ApicFabricLeafNodeInterface) SetFabricLeafNodeNil(b bool)`

 SetFabricLeafNodeNil sets the value for FabricLeafNode to be an explicit nil

### UnsetFabricLeafNode
`func (o *ApicFabricLeafNodeInterface) UnsetFabricLeafNode()`

UnsetFabricLeafNode ensures that no value is present for FabricLeafNode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


