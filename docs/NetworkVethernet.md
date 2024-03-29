# NetworkVethernet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.Vethernet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.Vethernet"]
**Description** | Pointer to **string** | Description for the vNIC. | [optional] [readonly] 
**VethId** | Pointer to **int64** | Vethernet Identifier on a Fabric Interconnect. | [optional] [readonly] 
**AdapterHostEthInterface** | Pointer to [**NullableAdapterHostEthInterfaceRelationship**](AdapterHostEthInterfaceRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkVethernet

`func NewNetworkVethernet(classId string, objectType string, ) *NetworkVethernet`

NewNetworkVethernet instantiates a new NetworkVethernet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVethernetWithDefaults

`func NewNetworkVethernetWithDefaults() *NetworkVethernet`

NewNetworkVethernetWithDefaults instantiates a new NetworkVethernet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkVethernet) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVethernet) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVethernet) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkVethernet) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVethernet) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVethernet) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *NetworkVethernet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkVethernet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkVethernet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkVethernet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVethId

`func (o *NetworkVethernet) GetVethId() int64`

GetVethId returns the VethId field if non-nil, zero value otherwise.

### GetVethIdOk

`func (o *NetworkVethernet) GetVethIdOk() (*int64, bool)`

GetVethIdOk returns a tuple with the VethId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVethId

`func (o *NetworkVethernet) SetVethId(v int64)`

SetVethId sets VethId field to given value.

### HasVethId

`func (o *NetworkVethernet) HasVethId() bool`

HasVethId returns a boolean if a field has been set.

### GetAdapterHostEthInterface

`func (o *NetworkVethernet) GetAdapterHostEthInterface() AdapterHostEthInterfaceRelationship`

GetAdapterHostEthInterface returns the AdapterHostEthInterface field if non-nil, zero value otherwise.

### GetAdapterHostEthInterfaceOk

`func (o *NetworkVethernet) GetAdapterHostEthInterfaceOk() (*AdapterHostEthInterfaceRelationship, bool)`

GetAdapterHostEthInterfaceOk returns a tuple with the AdapterHostEthInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterHostEthInterface

`func (o *NetworkVethernet) SetAdapterHostEthInterface(v AdapterHostEthInterfaceRelationship)`

SetAdapterHostEthInterface sets AdapterHostEthInterface field to given value.

### HasAdapterHostEthInterface

`func (o *NetworkVethernet) HasAdapterHostEthInterface() bool`

HasAdapterHostEthInterface returns a boolean if a field has been set.

### SetAdapterHostEthInterfaceNil

`func (o *NetworkVethernet) SetAdapterHostEthInterfaceNil(b bool)`

 SetAdapterHostEthInterfaceNil sets the value for AdapterHostEthInterface to be an explicit nil

### UnsetAdapterHostEthInterface
`func (o *NetworkVethernet) UnsetAdapterHostEthInterface()`

UnsetAdapterHostEthInterface ensures that no value is present for AdapterHostEthInterface, not even an explicit nil
### GetNetworkElement

`func (o *NetworkVethernet) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVethernet) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVethernet) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVethernet) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *NetworkVethernet) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *NetworkVethernet) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
### GetRegisteredDevice

`func (o *NetworkVethernet) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVethernet) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVethernet) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVethernet) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NetworkVethernet) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NetworkVethernet) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


