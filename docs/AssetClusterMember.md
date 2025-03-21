# AssetClusterMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.ClusterMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.ClusterMember"]
**ConnectedHost** | Pointer to **string** | The DNS hostname the device connector has used to connect to Intersight. Devices may be configured to connect to a set of DNS hostnames which all resolve to the same Intersight instance, the connected host is the latest hostname the device used to connect successfully to Intersight. | [optional] [readonly] 
**ConnectionFlapStatus** | Pointer to [**NullableAssetConnectionFlapStatus**](AssetConnectionFlapStatus.md) |  | [optional] 
**Leadership** | Pointer to **string** | The current leadershipstate of this member. Updated by the device connector on failover or leadership change. If a member is elected as Primary within the cluster this connection will be the same as the DeviceRegistration connection. E.g a message addressed to the DeviceRegistration will be forwarded to the ClusterMember connection. * &#x60;Unknown&#x60; - The node is unable to complete election or determine the current state. If the device has been registered before and the node has access to the current credentials, it will establish a connection to Intersight with limited capabilities that can be used to debug the HA failure from Intersight. * &#x60;Primary&#x60; - The node has been elected as the primary and will establish a connection to the Intersight service and accept all message types enabled for a primary node. There can only be one primary node in a given cluster, while the underlying platform may be active. If it is active, only one connector will assume the primary role. * &#x60;Secondary&#x60; - The node has been elected as a secondary node in the cluster. The device connector will establish a connection to the Intersight service with limited capabilities. E.g. file upload will be enabled, but requests to the underlying platform management will be disabled. | [optional] [readonly] [default to "Unknown"]
**LockedLeader** | Pointer to **bool** | Devices lock their leadership on failure to heartbeat with peers. Value acts as a third party tie breaker in election between nodes. Intersight enforces that only one cluster member exists with this value set to true. | [optional] 
**MemberIdentity** | Pointer to **string** | The unique identity of the member within the cluster. The identity is retrieved from the platform and reported by the device connector at connection time. | [optional] [readonly] 
**ParentClusterMemberIdentity** | Pointer to **string** | The member idenity of the cluster member through which this device is connected if applicable. | [optional] [readonly] 
**Sudi** | Pointer to [**NullableAssetSudiInfo**](AssetSudiInfo.md) |  | [optional] 
**Device** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewAssetClusterMember

`func NewAssetClusterMember(classId string, objectType string, ) *AssetClusterMember`

NewAssetClusterMember instantiates a new AssetClusterMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetClusterMemberWithDefaults

`func NewAssetClusterMemberWithDefaults() *AssetClusterMember`

NewAssetClusterMemberWithDefaults instantiates a new AssetClusterMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetClusterMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetClusterMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetClusterMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetClusterMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetClusterMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetClusterMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectedHost

`func (o *AssetClusterMember) GetConnectedHost() string`

GetConnectedHost returns the ConnectedHost field if non-nil, zero value otherwise.

### GetConnectedHostOk

`func (o *AssetClusterMember) GetConnectedHostOk() (*string, bool)`

GetConnectedHostOk returns a tuple with the ConnectedHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedHost

`func (o *AssetClusterMember) SetConnectedHost(v string)`

SetConnectedHost sets ConnectedHost field to given value.

### HasConnectedHost

`func (o *AssetClusterMember) HasConnectedHost() bool`

HasConnectedHost returns a boolean if a field has been set.

### GetConnectionFlapStatus

`func (o *AssetClusterMember) GetConnectionFlapStatus() AssetConnectionFlapStatus`

GetConnectionFlapStatus returns the ConnectionFlapStatus field if non-nil, zero value otherwise.

### GetConnectionFlapStatusOk

`func (o *AssetClusterMember) GetConnectionFlapStatusOk() (*AssetConnectionFlapStatus, bool)`

GetConnectionFlapStatusOk returns a tuple with the ConnectionFlapStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionFlapStatus

`func (o *AssetClusterMember) SetConnectionFlapStatus(v AssetConnectionFlapStatus)`

SetConnectionFlapStatus sets ConnectionFlapStatus field to given value.

### HasConnectionFlapStatus

`func (o *AssetClusterMember) HasConnectionFlapStatus() bool`

HasConnectionFlapStatus returns a boolean if a field has been set.

### SetConnectionFlapStatusNil

`func (o *AssetClusterMember) SetConnectionFlapStatusNil(b bool)`

 SetConnectionFlapStatusNil sets the value for ConnectionFlapStatus to be an explicit nil

### UnsetConnectionFlapStatus
`func (o *AssetClusterMember) UnsetConnectionFlapStatus()`

UnsetConnectionFlapStatus ensures that no value is present for ConnectionFlapStatus, not even an explicit nil
### GetLeadership

`func (o *AssetClusterMember) GetLeadership() string`

GetLeadership returns the Leadership field if non-nil, zero value otherwise.

### GetLeadershipOk

`func (o *AssetClusterMember) GetLeadershipOk() (*string, bool)`

GetLeadershipOk returns a tuple with the Leadership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadership

`func (o *AssetClusterMember) SetLeadership(v string)`

SetLeadership sets Leadership field to given value.

### HasLeadership

`func (o *AssetClusterMember) HasLeadership() bool`

HasLeadership returns a boolean if a field has been set.

### GetLockedLeader

`func (o *AssetClusterMember) GetLockedLeader() bool`

GetLockedLeader returns the LockedLeader field if non-nil, zero value otherwise.

### GetLockedLeaderOk

`func (o *AssetClusterMember) GetLockedLeaderOk() (*bool, bool)`

GetLockedLeaderOk returns a tuple with the LockedLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedLeader

`func (o *AssetClusterMember) SetLockedLeader(v bool)`

SetLockedLeader sets LockedLeader field to given value.

### HasLockedLeader

`func (o *AssetClusterMember) HasLockedLeader() bool`

HasLockedLeader returns a boolean if a field has been set.

### GetMemberIdentity

`func (o *AssetClusterMember) GetMemberIdentity() string`

GetMemberIdentity returns the MemberIdentity field if non-nil, zero value otherwise.

### GetMemberIdentityOk

`func (o *AssetClusterMember) GetMemberIdentityOk() (*string, bool)`

GetMemberIdentityOk returns a tuple with the MemberIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIdentity

`func (o *AssetClusterMember) SetMemberIdentity(v string)`

SetMemberIdentity sets MemberIdentity field to given value.

### HasMemberIdentity

`func (o *AssetClusterMember) HasMemberIdentity() bool`

HasMemberIdentity returns a boolean if a field has been set.

### GetParentClusterMemberIdentity

`func (o *AssetClusterMember) GetParentClusterMemberIdentity() string`

GetParentClusterMemberIdentity returns the ParentClusterMemberIdentity field if non-nil, zero value otherwise.

### GetParentClusterMemberIdentityOk

`func (o *AssetClusterMember) GetParentClusterMemberIdentityOk() (*string, bool)`

GetParentClusterMemberIdentityOk returns a tuple with the ParentClusterMemberIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentClusterMemberIdentity

`func (o *AssetClusterMember) SetParentClusterMemberIdentity(v string)`

SetParentClusterMemberIdentity sets ParentClusterMemberIdentity field to given value.

### HasParentClusterMemberIdentity

`func (o *AssetClusterMember) HasParentClusterMemberIdentity() bool`

HasParentClusterMemberIdentity returns a boolean if a field has been set.

### GetSudi

`func (o *AssetClusterMember) GetSudi() AssetSudiInfo`

GetSudi returns the Sudi field if non-nil, zero value otherwise.

### GetSudiOk

`func (o *AssetClusterMember) GetSudiOk() (*AssetSudiInfo, bool)`

GetSudiOk returns a tuple with the Sudi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudi

`func (o *AssetClusterMember) SetSudi(v AssetSudiInfo)`

SetSudi sets Sudi field to given value.

### HasSudi

`func (o *AssetClusterMember) HasSudi() bool`

HasSudi returns a boolean if a field has been set.

### SetSudiNil

`func (o *AssetClusterMember) SetSudiNil(b bool)`

 SetSudiNil sets the value for Sudi to be an explicit nil

### UnsetSudi
`func (o *AssetClusterMember) UnsetSudi()`

UnsetSudi ensures that no value is present for Sudi, not even an explicit nil
### GetDevice

`func (o *AssetClusterMember) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *AssetClusterMember) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *AssetClusterMember) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *AssetClusterMember) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *AssetClusterMember) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *AssetClusterMember) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


