# HyperflexNodeProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.NodeProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.NodeProfile"]
**HxdpDataIp** | Pointer to **string** | IP address for storage data network (Controller VM interface). | [optional] 
**HxdpMgmtIp** | Pointer to **string** | IP address for HyperFlex management network. | [optional] 
**HxdpStorageClientIp** | Pointer to **string** | IP address for storage client network (Controller VM interface). | [optional] 
**HypervisorControlIp** | Pointer to **string** | IP address for hypervisor control such as VM migration or pod management. | [optional] 
**HypervisorDataIp** | Pointer to **string** | IP address for storage data network (Hypervisor interface). | [optional] 
**HypervisorMgmtIp** | Pointer to **string** | IP address for Hypervisor management network. | [optional] 
**NodeRole** | Pointer to **string** | The role that this node performs in the HyperFlex cluster. * &#x60;Unknown&#x60; - The node role is not available. * &#x60;Storage&#x60; - The node persists data and contributes to the storage capacity of a cluster. * &#x60;Compute&#x60; - The node contributes to the compute capacity of a cluster. | [optional] [default to "Unknown"]
**AssignedServer** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**ClusterProfile** | Pointer to [**NullableHyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) |  | [optional] 
**Node** | Pointer to [**NullableHyperflexNodeRelationship**](HyperflexNodeRelationship.md) |  | [optional] 

## Methods

### NewHyperflexNodeProfile

`func NewHyperflexNodeProfile(classId string, objectType string, ) *HyperflexNodeProfile`

NewHyperflexNodeProfile instantiates a new HyperflexNodeProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNodeProfileWithDefaults

`func NewHyperflexNodeProfileWithDefaults() *HyperflexNodeProfile`

NewHyperflexNodeProfileWithDefaults instantiates a new HyperflexNodeProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexNodeProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexNodeProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexNodeProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexNodeProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexNodeProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexNodeProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHxdpDataIp

`func (o *HyperflexNodeProfile) GetHxdpDataIp() string`

GetHxdpDataIp returns the HxdpDataIp field if non-nil, zero value otherwise.

### GetHxdpDataIpOk

`func (o *HyperflexNodeProfile) GetHxdpDataIpOk() (*string, bool)`

GetHxdpDataIpOk returns a tuple with the HxdpDataIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpDataIp

`func (o *HyperflexNodeProfile) SetHxdpDataIp(v string)`

SetHxdpDataIp sets HxdpDataIp field to given value.

### HasHxdpDataIp

`func (o *HyperflexNodeProfile) HasHxdpDataIp() bool`

HasHxdpDataIp returns a boolean if a field has been set.

### GetHxdpMgmtIp

`func (o *HyperflexNodeProfile) GetHxdpMgmtIp() string`

GetHxdpMgmtIp returns the HxdpMgmtIp field if non-nil, zero value otherwise.

### GetHxdpMgmtIpOk

`func (o *HyperflexNodeProfile) GetHxdpMgmtIpOk() (*string, bool)`

GetHxdpMgmtIpOk returns a tuple with the HxdpMgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpMgmtIp

`func (o *HyperflexNodeProfile) SetHxdpMgmtIp(v string)`

SetHxdpMgmtIp sets HxdpMgmtIp field to given value.

### HasHxdpMgmtIp

`func (o *HyperflexNodeProfile) HasHxdpMgmtIp() bool`

HasHxdpMgmtIp returns a boolean if a field has been set.

### GetHxdpStorageClientIp

`func (o *HyperflexNodeProfile) GetHxdpStorageClientIp() string`

GetHxdpStorageClientIp returns the HxdpStorageClientIp field if non-nil, zero value otherwise.

### GetHxdpStorageClientIpOk

`func (o *HyperflexNodeProfile) GetHxdpStorageClientIpOk() (*string, bool)`

GetHxdpStorageClientIpOk returns a tuple with the HxdpStorageClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpStorageClientIp

`func (o *HyperflexNodeProfile) SetHxdpStorageClientIp(v string)`

SetHxdpStorageClientIp sets HxdpStorageClientIp field to given value.

### HasHxdpStorageClientIp

`func (o *HyperflexNodeProfile) HasHxdpStorageClientIp() bool`

HasHxdpStorageClientIp returns a boolean if a field has been set.

### GetHypervisorControlIp

`func (o *HyperflexNodeProfile) GetHypervisorControlIp() string`

GetHypervisorControlIp returns the HypervisorControlIp field if non-nil, zero value otherwise.

### GetHypervisorControlIpOk

`func (o *HyperflexNodeProfile) GetHypervisorControlIpOk() (*string, bool)`

GetHypervisorControlIpOk returns a tuple with the HypervisorControlIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorControlIp

`func (o *HyperflexNodeProfile) SetHypervisorControlIp(v string)`

SetHypervisorControlIp sets HypervisorControlIp field to given value.

### HasHypervisorControlIp

`func (o *HyperflexNodeProfile) HasHypervisorControlIp() bool`

HasHypervisorControlIp returns a boolean if a field has been set.

### GetHypervisorDataIp

`func (o *HyperflexNodeProfile) GetHypervisorDataIp() string`

GetHypervisorDataIp returns the HypervisorDataIp field if non-nil, zero value otherwise.

### GetHypervisorDataIpOk

`func (o *HyperflexNodeProfile) GetHypervisorDataIpOk() (*string, bool)`

GetHypervisorDataIpOk returns a tuple with the HypervisorDataIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorDataIp

`func (o *HyperflexNodeProfile) SetHypervisorDataIp(v string)`

SetHypervisorDataIp sets HypervisorDataIp field to given value.

### HasHypervisorDataIp

`func (o *HyperflexNodeProfile) HasHypervisorDataIp() bool`

HasHypervisorDataIp returns a boolean if a field has been set.

### GetHypervisorMgmtIp

`func (o *HyperflexNodeProfile) GetHypervisorMgmtIp() string`

GetHypervisorMgmtIp returns the HypervisorMgmtIp field if non-nil, zero value otherwise.

### GetHypervisorMgmtIpOk

`func (o *HyperflexNodeProfile) GetHypervisorMgmtIpOk() (*string, bool)`

GetHypervisorMgmtIpOk returns a tuple with the HypervisorMgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorMgmtIp

`func (o *HyperflexNodeProfile) SetHypervisorMgmtIp(v string)`

SetHypervisorMgmtIp sets HypervisorMgmtIp field to given value.

### HasHypervisorMgmtIp

`func (o *HyperflexNodeProfile) HasHypervisorMgmtIp() bool`

HasHypervisorMgmtIp returns a boolean if a field has been set.

### GetNodeRole

`func (o *HyperflexNodeProfile) GetNodeRole() string`

GetNodeRole returns the NodeRole field if non-nil, zero value otherwise.

### GetNodeRoleOk

`func (o *HyperflexNodeProfile) GetNodeRoleOk() (*string, bool)`

GetNodeRoleOk returns a tuple with the NodeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeRole

`func (o *HyperflexNodeProfile) SetNodeRole(v string)`

SetNodeRole sets NodeRole field to given value.

### HasNodeRole

`func (o *HyperflexNodeProfile) HasNodeRole() bool`

HasNodeRole returns a boolean if a field has been set.

### GetAssignedServer

`func (o *HyperflexNodeProfile) GetAssignedServer() ComputePhysicalRelationship`

GetAssignedServer returns the AssignedServer field if non-nil, zero value otherwise.

### GetAssignedServerOk

`func (o *HyperflexNodeProfile) GetAssignedServerOk() (*ComputePhysicalRelationship, bool)`

GetAssignedServerOk returns a tuple with the AssignedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedServer

`func (o *HyperflexNodeProfile) SetAssignedServer(v ComputePhysicalRelationship)`

SetAssignedServer sets AssignedServer field to given value.

### HasAssignedServer

`func (o *HyperflexNodeProfile) HasAssignedServer() bool`

HasAssignedServer returns a boolean if a field has been set.

### SetAssignedServerNil

`func (o *HyperflexNodeProfile) SetAssignedServerNil(b bool)`

 SetAssignedServerNil sets the value for AssignedServer to be an explicit nil

### UnsetAssignedServer
`func (o *HyperflexNodeProfile) UnsetAssignedServer()`

UnsetAssignedServer ensures that no value is present for AssignedServer, not even an explicit nil
### GetClusterProfile

`func (o *HyperflexNodeProfile) GetClusterProfile() HyperflexClusterProfileRelationship`

GetClusterProfile returns the ClusterProfile field if non-nil, zero value otherwise.

### GetClusterProfileOk

`func (o *HyperflexNodeProfile) GetClusterProfileOk() (*HyperflexClusterProfileRelationship, bool)`

GetClusterProfileOk returns a tuple with the ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfile

`func (o *HyperflexNodeProfile) SetClusterProfile(v HyperflexClusterProfileRelationship)`

SetClusterProfile sets ClusterProfile field to given value.

### HasClusterProfile

`func (o *HyperflexNodeProfile) HasClusterProfile() bool`

HasClusterProfile returns a boolean if a field has been set.

### SetClusterProfileNil

`func (o *HyperflexNodeProfile) SetClusterProfileNil(b bool)`

 SetClusterProfileNil sets the value for ClusterProfile to be an explicit nil

### UnsetClusterProfile
`func (o *HyperflexNodeProfile) UnsetClusterProfile()`

UnsetClusterProfile ensures that no value is present for ClusterProfile, not even an explicit nil
### GetNode

`func (o *HyperflexNodeProfile) GetNode() HyperflexNodeRelationship`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *HyperflexNodeProfile) GetNodeOk() (*HyperflexNodeRelationship, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *HyperflexNodeProfile) SetNode(v HyperflexNodeRelationship)`

SetNode sets Node field to given value.

### HasNode

`func (o *HyperflexNodeProfile) HasNode() bool`

HasNode returns a boolean if a field has been set.

### SetNodeNil

`func (o *HyperflexNodeProfile) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *HyperflexNodeProfile) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


