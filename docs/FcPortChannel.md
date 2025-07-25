# FcPortChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fc.PortChannel"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fc.PortChannel"]
**AdminSpeed** | Pointer to **string** | Administrator configured Speed applied on the port channel. | [optional] 
**AdminState** | Pointer to **string** | Administratively configured state (enabled/disabled) for this portchannel. | [optional] [readonly] 
**Mode** | Pointer to **string** | Mode information N_proxy, F or E associated to the Fibre Channel portchannel. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the port channel. | [optional] [readonly] 
**OperSpeed** | Pointer to **string** | Operational speed of this port-channel. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of this port-channel. | [optional] [readonly] 
**OperStateQual** | Pointer to **string** | Reason for this port-channel&#39;s Operational state. | [optional] 
**PortChannelId** | Pointer to **int64** | Unique identifier for this port-channel on the FI. | [optional] [readonly] 
**Role** | Pointer to **string** | This port-channel&#39;s configured role (fcUplink, fcStorage, etc.). | [optional] 
**SwitchId** | Pointer to **string** | Switch Identifier that is local to a cluster. | [optional] [readonly] 
**UserLabel** | Pointer to **string** | The user defined label assigned to the port. | [optional] [readonly] 
**Vsan** | Pointer to **int64** | Virtual San that is associated to the port-channel. | [optional] [readonly] 
**Wwn** | Pointer to **string** | World Wide Name of the port channel. | [optional] [readonly] 
**EquipmentSwitchCard** | Pointer to [**NullableEquipmentSwitchCardRelationship**](EquipmentSwitchCardRelationship.md) |  | [optional] 
**FcNeighbor** | Pointer to [**[]FcNeighborRelationship**](FcNeighborRelationship.md) | An array of relationships to fcNeighbor resources. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewFcPortChannel

`func NewFcPortChannel(classId string, objectType string, ) *FcPortChannel`

NewFcPortChannel instantiates a new FcPortChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcPortChannelWithDefaults

`func NewFcPortChannelWithDefaults() *FcPortChannel`

NewFcPortChannelWithDefaults instantiates a new FcPortChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcPortChannel) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcPortChannel) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcPortChannel) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcPortChannel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcPortChannel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcPortChannel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSpeed

`func (o *FcPortChannel) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FcPortChannel) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FcPortChannel) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FcPortChannel) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetAdminState

`func (o *FcPortChannel) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *FcPortChannel) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *FcPortChannel) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *FcPortChannel) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetMode

`func (o *FcPortChannel) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FcPortChannel) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FcPortChannel) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FcPortChannel) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *FcPortChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FcPortChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FcPortChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FcPortChannel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperSpeed

`func (o *FcPortChannel) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *FcPortChannel) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *FcPortChannel) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *FcPortChannel) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetOperState

`func (o *FcPortChannel) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *FcPortChannel) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *FcPortChannel) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *FcPortChannel) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperStateQual

`func (o *FcPortChannel) GetOperStateQual() string`

GetOperStateQual returns the OperStateQual field if non-nil, zero value otherwise.

### GetOperStateQualOk

`func (o *FcPortChannel) GetOperStateQualOk() (*string, bool)`

GetOperStateQualOk returns a tuple with the OperStateQual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStateQual

`func (o *FcPortChannel) SetOperStateQual(v string)`

SetOperStateQual sets OperStateQual field to given value.

### HasOperStateQual

`func (o *FcPortChannel) HasOperStateQual() bool`

HasOperStateQual returns a boolean if a field has been set.

### GetPortChannelId

`func (o *FcPortChannel) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *FcPortChannel) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *FcPortChannel) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *FcPortChannel) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetRole

`func (o *FcPortChannel) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *FcPortChannel) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *FcPortChannel) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *FcPortChannel) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSwitchId

`func (o *FcPortChannel) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *FcPortChannel) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *FcPortChannel) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *FcPortChannel) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetUserLabel

`func (o *FcPortChannel) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *FcPortChannel) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *FcPortChannel) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *FcPortChannel) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetVsan

`func (o *FcPortChannel) GetVsan() int64`

GetVsan returns the Vsan field if non-nil, zero value otherwise.

### GetVsanOk

`func (o *FcPortChannel) GetVsanOk() (*int64, bool)`

GetVsanOk returns a tuple with the Vsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsan

`func (o *FcPortChannel) SetVsan(v int64)`

SetVsan sets Vsan field to given value.

### HasVsan

`func (o *FcPortChannel) HasVsan() bool`

HasVsan returns a boolean if a field has been set.

### GetWwn

`func (o *FcPortChannel) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *FcPortChannel) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *FcPortChannel) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *FcPortChannel) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetEquipmentSwitchCard

`func (o *FcPortChannel) GetEquipmentSwitchCard() EquipmentSwitchCardRelationship`

GetEquipmentSwitchCard returns the EquipmentSwitchCard field if non-nil, zero value otherwise.

### GetEquipmentSwitchCardOk

`func (o *FcPortChannel) GetEquipmentSwitchCardOk() (*EquipmentSwitchCardRelationship, bool)`

GetEquipmentSwitchCardOk returns a tuple with the EquipmentSwitchCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSwitchCard

`func (o *FcPortChannel) SetEquipmentSwitchCard(v EquipmentSwitchCardRelationship)`

SetEquipmentSwitchCard sets EquipmentSwitchCard field to given value.

### HasEquipmentSwitchCard

`func (o *FcPortChannel) HasEquipmentSwitchCard() bool`

HasEquipmentSwitchCard returns a boolean if a field has been set.

### SetEquipmentSwitchCardNil

`func (o *FcPortChannel) SetEquipmentSwitchCardNil(b bool)`

 SetEquipmentSwitchCardNil sets the value for EquipmentSwitchCard to be an explicit nil

### UnsetEquipmentSwitchCard
`func (o *FcPortChannel) UnsetEquipmentSwitchCard()`

UnsetEquipmentSwitchCard ensures that no value is present for EquipmentSwitchCard, not even an explicit nil
### GetFcNeighbor

`func (o *FcPortChannel) GetFcNeighbor() []FcNeighborRelationship`

GetFcNeighbor returns the FcNeighbor field if non-nil, zero value otherwise.

### GetFcNeighborOk

`func (o *FcPortChannel) GetFcNeighborOk() (*[]FcNeighborRelationship, bool)`

GetFcNeighborOk returns a tuple with the FcNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcNeighbor

`func (o *FcPortChannel) SetFcNeighbor(v []FcNeighborRelationship)`

SetFcNeighbor sets FcNeighbor field to given value.

### HasFcNeighbor

`func (o *FcPortChannel) HasFcNeighbor() bool`

HasFcNeighbor returns a boolean if a field has been set.

### SetFcNeighborNil

`func (o *FcPortChannel) SetFcNeighborNil(b bool)`

 SetFcNeighborNil sets the value for FcNeighbor to be an explicit nil

### UnsetFcNeighbor
`func (o *FcPortChannel) UnsetFcNeighbor()`

UnsetFcNeighbor ensures that no value is present for FcNeighbor, not even an explicit nil
### GetNetworkElement

`func (o *FcPortChannel) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *FcPortChannel) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *FcPortChannel) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *FcPortChannel) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *FcPortChannel) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *FcPortChannel) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
### GetRegisteredDevice

`func (o *FcPortChannel) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *FcPortChannel) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *FcPortChannel) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *FcPortChannel) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *FcPortChannel) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *FcPortChannel) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


