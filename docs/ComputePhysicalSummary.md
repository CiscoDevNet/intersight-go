# ComputePhysicalSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.PhysicalSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.PhysicalSummary"]
**AdminPowerState** | Pointer to **string** | The desired power state of the server. | [optional] [readonly] 
**AlarmSummary** | Pointer to [**NullableComputeAlarmSummary**](ComputeAlarmSummary.md) |  | [optional] 
**AssetTag** | Pointer to **string** | The user defined asset tag assigned to the server. | [optional] [readonly] 
**AvailableMemory** | Pointer to **int64** | Total memeory of the server in MB. | [optional] [readonly] 
**BiosPostComplete** | Pointer to **bool** | The BIOS POST completion status of the server. | [optional] [readonly] 
**ChassisId** | Pointer to **string** | The id of the chassis that the blade is discovered in. | [optional] [readonly] 
**ConnectionStatus** | Pointer to **string** | Connectivity Status of RackUnit to Switch - A or B or AB. | [optional] [readonly] 
**CoolingMode** | Pointer to **string** | Cooling mode representation of the server, supported modes include Air and Immersion. * &#x60;Air&#x60; - Cooling mode of the device is set to Air. * &#x60;Immersion&#x60; - Cooling mode of the device is set to Immersion. | [optional] [readonly] [default to "Air"]
**CpuCapacity** | Pointer to **float32** | Total processing capacity of the server. | [optional] [readonly] 
**DeviceMoId** | Pointer to **string** | The MoId of the registered device that coresponds to the server. | [optional] [readonly] 
**Dn** | Pointer to **string** | The Distinguished Name unambiguously identifies an object in the system. | [optional] [readonly] 
**FaultSummary** | Pointer to **int64** | The fault summary for the server. | [optional] [readonly] 
**Firmware** | Pointer to **string** | The firmware version of the Cisco Integrated Management Controller (CIMC) for this server. | [optional] [readonly] 
**FrontPanelLockState** | Pointer to **string** | The actual front panel state of the server. * &#x60;None&#x60; - Front Panel of the server is set to None state. It is required so that the next frontPanelLockState operation can be triggered. * &#x60;Lock&#x60; - Front Panel of the server is set to Locked state. * &#x60;Unlock&#x60; - Front Panel of the server is set to Unlocked state. | [optional] [default to "None"]
**HardwareUuid** | Pointer to **string** | The universally unique hardware identity of the server provided by the manufacturer. | [optional] [readonly] 
**Ipv4Address** | Pointer to **string** | The IPv4 address configured on the management interface of the Integrated Management Controller. | [optional] [readonly] 
**IsUpgraded** | Pointer to **bool** | This field indicates the compute status of the catalog values for the associated component or hardware. | [optional] [readonly] 
**KvmIpAddresses** | Pointer to [**[]ComputeIpAddress**](ComputeIpAddress.md) |  | [optional] 
**KvmServerStateEnabled** | Pointer to **bool** | The KVM server state of the server. | [optional] [readonly] 
**KvmVendor** | Pointer to **string** | The KVM Vendor for the server. | [optional] [readonly] 
**Lifecycle** | Pointer to **string** | The lifecycle of the blade server. * &#x60;None&#x60; - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * &#x60;Active&#x60; - Default Lifecycle State for a physical entity. * &#x60;Decommissioned&#x60; - Decommission Lifecycle state. * &#x60;DiscoveryInProgress&#x60; - DiscoveryInProgress Lifecycle state. * &#x60;DiscoveryFailed&#x60; - DiscoveryFailed Lifecycle state. * &#x60;FirmwareUpgradeInProgress&#x60; - Firmware upgrade is in progress on given physical entity. * &#x60;SecureEraseInProgress&#x60; - Secure Erase is in progress on given physical entity. * &#x60;ScrubInProgress&#x60; - Scrub is in progress on given physical entity. * &#x60;BladeMigrationInProgress&#x60; - Server slot migration is in progress on given physical entity. * &#x60;SlotMismatch&#x60; - The blade server is detected in a different chassis/slot than it was previously. * &#x60;Removed&#x60; - The blade server has been removed from its discovered slot, and not detected anywhere else. Blade inventory can be cleaned up by performing a software remove operation on the physically removed blade. * &#x60;Moved&#x60; - The blade server has been moved from its discovered location to a new location. Blade inventory can be updated by performing a rediscover operation on the moved blade. * &#x60;Replaced&#x60; - The blade server has been removed from its discovered location and another blade has been inserted in that location. Blade inventory can be cleaned up and updated by doing a software remove operation on the physically removed blade. * &#x60;MovedAndReplaced&#x60; - The blade server has been moved from its discovered location to a new location and another blade has been inserted into the old discovered location. Blade inventory can be updated by performing a rediscover operation on the moved blade. | [optional] [readonly] [default to "None"]
**ManagementMode** | Pointer to **string** | The management mode of the server. * &#x60;IntersightStandalone&#x60; - Intersight Standalone mode of operation. * &#x60;UCSM&#x60; - Unified Computing System Manager mode of operation. * &#x60;Intersight&#x60; - Intersight managed mode of operation. | [optional] [readonly] [default to "IntersightStandalone"]
**MemorySpeed** | Pointer to **string** | The maximum memory speed in MHz available on the server. | [optional] [readonly] 
**MgmtIpAddress** | Pointer to **string** | Management address of the server. | [optional] [readonly] 
**Model** | Pointer to **string** | This field identifies the model of the given component. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the UCS Fabric Interconnect cluster or Cisco Integrated Management Controller (CIMC). When this server is attached to a UCS Fabric Interconnect, the value of this property is the name of the UCS Fabric Interconnect along with chassis/server Id. When this server configured in standalone mode, the value of this property is the name of the Cisco Integrated Management Controller. when this server is configired in IMM mode, the value of this property contains model and chassis/server Id. | [optional] [readonly] 
**NumAdaptors** | Pointer to **int64** | The total number of network adapters present on the server. | [optional] [readonly] 
**NumCpuCores** | Pointer to **int64** | The total number of CPU cores enabled on the server. | [optional] [readonly] 
**NumCpuCoresEnabled** | Pointer to **int64** | The total number of CPU cores enabled on the server. | [optional] [readonly] 
**NumCpus** | Pointer to **int64** | The total number of CPUs present on the server. | [optional] [readonly] 
**NumEthHostInterfaces** | Pointer to **int64** | The total number of vNICs which are visible to a host on the server. | [optional] [readonly] 
**NumFcHostInterfaces** | Pointer to **int64** | The total number of vHBAs which are visible to a host on the server. | [optional] [readonly] 
**NumThreads** | Pointer to **int64** | The total number of threads the server is capable of handling. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | The actual power state of the server. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | The operational state of the server. | [optional] [readonly] 
**Operability** | Pointer to **string** | The operability of the server. | [optional] [readonly] 
**PackageVersion** | Pointer to **string** | Bundle version which the firmware belongs to. | [optional] [readonly] 
**Personality** | Pointer to **string** | Unique identity of added software personality. | [optional] [readonly] 
**PlatformType** | Pointer to **string** | The platform type of the registered device - whether managed by UCSM or operating in standalone mode. | [optional] [readonly] 
**Presence** | Pointer to **string** | This field identifies the presence (equipped) or absence of the given component. | [optional] [readonly] 
**Revision** | Pointer to **string** | This field identifies the revision of the given component. | [optional] [readonly] 
**Rn** | Pointer to **string** | The Relative Name uniquely identifies an object within a given context. | [optional] [readonly] 
**ScaledMode** | Pointer to **string** | The mode of the server that determines it is scaled. | [optional] [readonly] 
**Serial** | Pointer to **string** | This field identifies the serial of the given component. | [optional] [readonly] 
**ServerId** | Pointer to **int64** | RackUnit ID that uniquely identifies the server. | [optional] [readonly] 
**ServiceProfile** | Pointer to **string** | The distinguished name of the service profile to which the server is associated to. It is applicable only for servers which are managed via UCSM. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | The slot number in the chassis that the blade is discovered in. | [optional] [readonly] 
**SourceObjectType** | Pointer to **string** | Stores the source object type. This feild will either be RackUnit or Blade. | [optional] [readonly] 
**TopologyScanStatus** | Pointer to **string** | To maintain the Topology workflow run status. | [optional] [readonly] 
**TotalMemory** | Pointer to **int64** | The total memory available on the server. | [optional] [readonly] 
**TunneledKvm** | Pointer to **bool** | The Tunneled vKVM status of the server. | [optional] [readonly] 
**UserLabel** | Pointer to **string** | The user defined label assigned to the server. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The universally unique identity of the server. | [optional] [readonly] 
**Vendor** | Pointer to **string** | This field identifies the vendor of the given component. | [optional] [readonly] 
**CustomPermissionResources** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**EquipmentChassis** | Pointer to [**NullableEquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**InventoryParent** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewComputePhysicalSummary

`func NewComputePhysicalSummary(classId string, objectType string, ) *ComputePhysicalSummary`

NewComputePhysicalSummary instantiates a new ComputePhysicalSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePhysicalSummaryWithDefaults

`func NewComputePhysicalSummaryWithDefaults() *ComputePhysicalSummary`

NewComputePhysicalSummaryWithDefaults instantiates a new ComputePhysicalSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputePhysicalSummary) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputePhysicalSummary) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputePhysicalSummary) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputePhysicalSummary) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputePhysicalSummary) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputePhysicalSummary) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminPowerState

`func (o *ComputePhysicalSummary) GetAdminPowerState() string`

GetAdminPowerState returns the AdminPowerState field if non-nil, zero value otherwise.

### GetAdminPowerStateOk

`func (o *ComputePhysicalSummary) GetAdminPowerStateOk() (*string, bool)`

GetAdminPowerStateOk returns a tuple with the AdminPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPowerState

`func (o *ComputePhysicalSummary) SetAdminPowerState(v string)`

SetAdminPowerState sets AdminPowerState field to given value.

### HasAdminPowerState

`func (o *ComputePhysicalSummary) HasAdminPowerState() bool`

HasAdminPowerState returns a boolean if a field has been set.

### GetAlarmSummary

`func (o *ComputePhysicalSummary) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *ComputePhysicalSummary) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *ComputePhysicalSummary) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *ComputePhysicalSummary) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *ComputePhysicalSummary) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *ComputePhysicalSummary) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetAssetTag

`func (o *ComputePhysicalSummary) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *ComputePhysicalSummary) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *ComputePhysicalSummary) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *ComputePhysicalSummary) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetAvailableMemory

`func (o *ComputePhysicalSummary) GetAvailableMemory() int64`

GetAvailableMemory returns the AvailableMemory field if non-nil, zero value otherwise.

### GetAvailableMemoryOk

`func (o *ComputePhysicalSummary) GetAvailableMemoryOk() (*int64, bool)`

GetAvailableMemoryOk returns a tuple with the AvailableMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMemory

`func (o *ComputePhysicalSummary) SetAvailableMemory(v int64)`

SetAvailableMemory sets AvailableMemory field to given value.

### HasAvailableMemory

`func (o *ComputePhysicalSummary) HasAvailableMemory() bool`

HasAvailableMemory returns a boolean if a field has been set.

### GetBiosPostComplete

`func (o *ComputePhysicalSummary) GetBiosPostComplete() bool`

GetBiosPostComplete returns the BiosPostComplete field if non-nil, zero value otherwise.

### GetBiosPostCompleteOk

`func (o *ComputePhysicalSummary) GetBiosPostCompleteOk() (*bool, bool)`

GetBiosPostCompleteOk returns a tuple with the BiosPostComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosPostComplete

`func (o *ComputePhysicalSummary) SetBiosPostComplete(v bool)`

SetBiosPostComplete sets BiosPostComplete field to given value.

### HasBiosPostComplete

`func (o *ComputePhysicalSummary) HasBiosPostComplete() bool`

HasBiosPostComplete returns a boolean if a field has been set.

### GetChassisId

`func (o *ComputePhysicalSummary) GetChassisId() string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *ComputePhysicalSummary) GetChassisIdOk() (*string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *ComputePhysicalSummary) SetChassisId(v string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *ComputePhysicalSummary) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *ComputePhysicalSummary) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *ComputePhysicalSummary) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *ComputePhysicalSummary) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *ComputePhysicalSummary) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetCoolingMode

`func (o *ComputePhysicalSummary) GetCoolingMode() string`

GetCoolingMode returns the CoolingMode field if non-nil, zero value otherwise.

### GetCoolingModeOk

`func (o *ComputePhysicalSummary) GetCoolingModeOk() (*string, bool)`

GetCoolingModeOk returns a tuple with the CoolingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingMode

`func (o *ComputePhysicalSummary) SetCoolingMode(v string)`

SetCoolingMode sets CoolingMode field to given value.

### HasCoolingMode

`func (o *ComputePhysicalSummary) HasCoolingMode() bool`

HasCoolingMode returns a boolean if a field has been set.

### GetCpuCapacity

`func (o *ComputePhysicalSummary) GetCpuCapacity() float32`

GetCpuCapacity returns the CpuCapacity field if non-nil, zero value otherwise.

### GetCpuCapacityOk

`func (o *ComputePhysicalSummary) GetCpuCapacityOk() (*float32, bool)`

GetCpuCapacityOk returns a tuple with the CpuCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCapacity

`func (o *ComputePhysicalSummary) SetCpuCapacity(v float32)`

SetCpuCapacity sets CpuCapacity field to given value.

### HasCpuCapacity

`func (o *ComputePhysicalSummary) HasCpuCapacity() bool`

HasCpuCapacity returns a boolean if a field has been set.

### GetDeviceMoId

`func (o *ComputePhysicalSummary) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *ComputePhysicalSummary) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *ComputePhysicalSummary) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *ComputePhysicalSummary) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *ComputePhysicalSummary) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ComputePhysicalSummary) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ComputePhysicalSummary) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ComputePhysicalSummary) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFaultSummary

`func (o *ComputePhysicalSummary) GetFaultSummary() int64`

GetFaultSummary returns the FaultSummary field if non-nil, zero value otherwise.

### GetFaultSummaryOk

`func (o *ComputePhysicalSummary) GetFaultSummaryOk() (*int64, bool)`

GetFaultSummaryOk returns a tuple with the FaultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultSummary

`func (o *ComputePhysicalSummary) SetFaultSummary(v int64)`

SetFaultSummary sets FaultSummary field to given value.

### HasFaultSummary

`func (o *ComputePhysicalSummary) HasFaultSummary() bool`

HasFaultSummary returns a boolean if a field has been set.

### GetFirmware

`func (o *ComputePhysicalSummary) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *ComputePhysicalSummary) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *ComputePhysicalSummary) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *ComputePhysicalSummary) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetFrontPanelLockState

`func (o *ComputePhysicalSummary) GetFrontPanelLockState() string`

GetFrontPanelLockState returns the FrontPanelLockState field if non-nil, zero value otherwise.

### GetFrontPanelLockStateOk

`func (o *ComputePhysicalSummary) GetFrontPanelLockStateOk() (*string, bool)`

GetFrontPanelLockStateOk returns a tuple with the FrontPanelLockState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPanelLockState

`func (o *ComputePhysicalSummary) SetFrontPanelLockState(v string)`

SetFrontPanelLockState sets FrontPanelLockState field to given value.

### HasFrontPanelLockState

`func (o *ComputePhysicalSummary) HasFrontPanelLockState() bool`

HasFrontPanelLockState returns a boolean if a field has been set.

### GetHardwareUuid

`func (o *ComputePhysicalSummary) GetHardwareUuid() string`

GetHardwareUuid returns the HardwareUuid field if non-nil, zero value otherwise.

### GetHardwareUuidOk

`func (o *ComputePhysicalSummary) GetHardwareUuidOk() (*string, bool)`

GetHardwareUuidOk returns a tuple with the HardwareUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareUuid

`func (o *ComputePhysicalSummary) SetHardwareUuid(v string)`

SetHardwareUuid sets HardwareUuid field to given value.

### HasHardwareUuid

`func (o *ComputePhysicalSummary) HasHardwareUuid() bool`

HasHardwareUuid returns a boolean if a field has been set.

### GetIpv4Address

`func (o *ComputePhysicalSummary) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *ComputePhysicalSummary) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *ComputePhysicalSummary) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *ComputePhysicalSummary) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIsUpgraded

`func (o *ComputePhysicalSummary) GetIsUpgraded() bool`

GetIsUpgraded returns the IsUpgraded field if non-nil, zero value otherwise.

### GetIsUpgradedOk

`func (o *ComputePhysicalSummary) GetIsUpgradedOk() (*bool, bool)`

GetIsUpgradedOk returns a tuple with the IsUpgraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgraded

`func (o *ComputePhysicalSummary) SetIsUpgraded(v bool)`

SetIsUpgraded sets IsUpgraded field to given value.

### HasIsUpgraded

`func (o *ComputePhysicalSummary) HasIsUpgraded() bool`

HasIsUpgraded returns a boolean if a field has been set.

### GetKvmIpAddresses

`func (o *ComputePhysicalSummary) GetKvmIpAddresses() []ComputeIpAddress`

GetKvmIpAddresses returns the KvmIpAddresses field if non-nil, zero value otherwise.

### GetKvmIpAddressesOk

`func (o *ComputePhysicalSummary) GetKvmIpAddressesOk() (*[]ComputeIpAddress, bool)`

GetKvmIpAddressesOk returns a tuple with the KvmIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmIpAddresses

`func (o *ComputePhysicalSummary) SetKvmIpAddresses(v []ComputeIpAddress)`

SetKvmIpAddresses sets KvmIpAddresses field to given value.

### HasKvmIpAddresses

`func (o *ComputePhysicalSummary) HasKvmIpAddresses() bool`

HasKvmIpAddresses returns a boolean if a field has been set.

### SetKvmIpAddressesNil

`func (o *ComputePhysicalSummary) SetKvmIpAddressesNil(b bool)`

 SetKvmIpAddressesNil sets the value for KvmIpAddresses to be an explicit nil

### UnsetKvmIpAddresses
`func (o *ComputePhysicalSummary) UnsetKvmIpAddresses()`

UnsetKvmIpAddresses ensures that no value is present for KvmIpAddresses, not even an explicit nil
### GetKvmServerStateEnabled

`func (o *ComputePhysicalSummary) GetKvmServerStateEnabled() bool`

GetKvmServerStateEnabled returns the KvmServerStateEnabled field if non-nil, zero value otherwise.

### GetKvmServerStateEnabledOk

`func (o *ComputePhysicalSummary) GetKvmServerStateEnabledOk() (*bool, bool)`

GetKvmServerStateEnabledOk returns a tuple with the KvmServerStateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmServerStateEnabled

`func (o *ComputePhysicalSummary) SetKvmServerStateEnabled(v bool)`

SetKvmServerStateEnabled sets KvmServerStateEnabled field to given value.

### HasKvmServerStateEnabled

`func (o *ComputePhysicalSummary) HasKvmServerStateEnabled() bool`

HasKvmServerStateEnabled returns a boolean if a field has been set.

### GetKvmVendor

`func (o *ComputePhysicalSummary) GetKvmVendor() string`

GetKvmVendor returns the KvmVendor field if non-nil, zero value otherwise.

### GetKvmVendorOk

`func (o *ComputePhysicalSummary) GetKvmVendorOk() (*string, bool)`

GetKvmVendorOk returns a tuple with the KvmVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmVendor

`func (o *ComputePhysicalSummary) SetKvmVendor(v string)`

SetKvmVendor sets KvmVendor field to given value.

### HasKvmVendor

`func (o *ComputePhysicalSummary) HasKvmVendor() bool`

HasKvmVendor returns a boolean if a field has been set.

### GetLifecycle

`func (o *ComputePhysicalSummary) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *ComputePhysicalSummary) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *ComputePhysicalSummary) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *ComputePhysicalSummary) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetManagementMode

`func (o *ComputePhysicalSummary) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *ComputePhysicalSummary) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *ComputePhysicalSummary) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *ComputePhysicalSummary) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetMemorySpeed

`func (o *ComputePhysicalSummary) GetMemorySpeed() string`

GetMemorySpeed returns the MemorySpeed field if non-nil, zero value otherwise.

### GetMemorySpeedOk

`func (o *ComputePhysicalSummary) GetMemorySpeedOk() (*string, bool)`

GetMemorySpeedOk returns a tuple with the MemorySpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySpeed

`func (o *ComputePhysicalSummary) SetMemorySpeed(v string)`

SetMemorySpeed sets MemorySpeed field to given value.

### HasMemorySpeed

`func (o *ComputePhysicalSummary) HasMemorySpeed() bool`

HasMemorySpeed returns a boolean if a field has been set.

### GetMgmtIpAddress

`func (o *ComputePhysicalSummary) GetMgmtIpAddress() string`

GetMgmtIpAddress returns the MgmtIpAddress field if non-nil, zero value otherwise.

### GetMgmtIpAddressOk

`func (o *ComputePhysicalSummary) GetMgmtIpAddressOk() (*string, bool)`

GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpAddress

`func (o *ComputePhysicalSummary) SetMgmtIpAddress(v string)`

SetMgmtIpAddress sets MgmtIpAddress field to given value.

### HasMgmtIpAddress

`func (o *ComputePhysicalSummary) HasMgmtIpAddress() bool`

HasMgmtIpAddress returns a boolean if a field has been set.

### GetModel

`func (o *ComputePhysicalSummary) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComputePhysicalSummary) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComputePhysicalSummary) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComputePhysicalSummary) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *ComputePhysicalSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputePhysicalSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputePhysicalSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputePhysicalSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumAdaptors

`func (o *ComputePhysicalSummary) GetNumAdaptors() int64`

GetNumAdaptors returns the NumAdaptors field if non-nil, zero value otherwise.

### GetNumAdaptorsOk

`func (o *ComputePhysicalSummary) GetNumAdaptorsOk() (*int64, bool)`

GetNumAdaptorsOk returns a tuple with the NumAdaptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAdaptors

`func (o *ComputePhysicalSummary) SetNumAdaptors(v int64)`

SetNumAdaptors sets NumAdaptors field to given value.

### HasNumAdaptors

`func (o *ComputePhysicalSummary) HasNumAdaptors() bool`

HasNumAdaptors returns a boolean if a field has been set.

### GetNumCpuCores

`func (o *ComputePhysicalSummary) GetNumCpuCores() int64`

GetNumCpuCores returns the NumCpuCores field if non-nil, zero value otherwise.

### GetNumCpuCoresOk

`func (o *ComputePhysicalSummary) GetNumCpuCoresOk() (*int64, bool)`

GetNumCpuCoresOk returns a tuple with the NumCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCores

`func (o *ComputePhysicalSummary) SetNumCpuCores(v int64)`

SetNumCpuCores sets NumCpuCores field to given value.

### HasNumCpuCores

`func (o *ComputePhysicalSummary) HasNumCpuCores() bool`

HasNumCpuCores returns a boolean if a field has been set.

### GetNumCpuCoresEnabled

`func (o *ComputePhysicalSummary) GetNumCpuCoresEnabled() int64`

GetNumCpuCoresEnabled returns the NumCpuCoresEnabled field if non-nil, zero value otherwise.

### GetNumCpuCoresEnabledOk

`func (o *ComputePhysicalSummary) GetNumCpuCoresEnabledOk() (*int64, bool)`

GetNumCpuCoresEnabledOk returns a tuple with the NumCpuCoresEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCoresEnabled

`func (o *ComputePhysicalSummary) SetNumCpuCoresEnabled(v int64)`

SetNumCpuCoresEnabled sets NumCpuCoresEnabled field to given value.

### HasNumCpuCoresEnabled

`func (o *ComputePhysicalSummary) HasNumCpuCoresEnabled() bool`

HasNumCpuCoresEnabled returns a boolean if a field has been set.

### GetNumCpus

`func (o *ComputePhysicalSummary) GetNumCpus() int64`

GetNumCpus returns the NumCpus field if non-nil, zero value otherwise.

### GetNumCpusOk

`func (o *ComputePhysicalSummary) GetNumCpusOk() (*int64, bool)`

GetNumCpusOk returns a tuple with the NumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpus

`func (o *ComputePhysicalSummary) SetNumCpus(v int64)`

SetNumCpus sets NumCpus field to given value.

### HasNumCpus

`func (o *ComputePhysicalSummary) HasNumCpus() bool`

HasNumCpus returns a boolean if a field has been set.

### GetNumEthHostInterfaces

`func (o *ComputePhysicalSummary) GetNumEthHostInterfaces() int64`

GetNumEthHostInterfaces returns the NumEthHostInterfaces field if non-nil, zero value otherwise.

### GetNumEthHostInterfacesOk

`func (o *ComputePhysicalSummary) GetNumEthHostInterfacesOk() (*int64, bool)`

GetNumEthHostInterfacesOk returns a tuple with the NumEthHostInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEthHostInterfaces

`func (o *ComputePhysicalSummary) SetNumEthHostInterfaces(v int64)`

SetNumEthHostInterfaces sets NumEthHostInterfaces field to given value.

### HasNumEthHostInterfaces

`func (o *ComputePhysicalSummary) HasNumEthHostInterfaces() bool`

HasNumEthHostInterfaces returns a boolean if a field has been set.

### GetNumFcHostInterfaces

`func (o *ComputePhysicalSummary) GetNumFcHostInterfaces() int64`

GetNumFcHostInterfaces returns the NumFcHostInterfaces field if non-nil, zero value otherwise.

### GetNumFcHostInterfacesOk

`func (o *ComputePhysicalSummary) GetNumFcHostInterfacesOk() (*int64, bool)`

GetNumFcHostInterfacesOk returns a tuple with the NumFcHostInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFcHostInterfaces

`func (o *ComputePhysicalSummary) SetNumFcHostInterfaces(v int64)`

SetNumFcHostInterfaces sets NumFcHostInterfaces field to given value.

### HasNumFcHostInterfaces

`func (o *ComputePhysicalSummary) HasNumFcHostInterfaces() bool`

HasNumFcHostInterfaces returns a boolean if a field has been set.

### GetNumThreads

`func (o *ComputePhysicalSummary) GetNumThreads() int64`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *ComputePhysicalSummary) GetNumThreadsOk() (*int64, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *ComputePhysicalSummary) SetNumThreads(v int64)`

SetNumThreads sets NumThreads field to given value.

### HasNumThreads

`func (o *ComputePhysicalSummary) HasNumThreads() bool`

HasNumThreads returns a boolean if a field has been set.

### GetOperPowerState

`func (o *ComputePhysicalSummary) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *ComputePhysicalSummary) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *ComputePhysicalSummary) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *ComputePhysicalSummary) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperReason

`func (o *ComputePhysicalSummary) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *ComputePhysicalSummary) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *ComputePhysicalSummary) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *ComputePhysicalSummary) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *ComputePhysicalSummary) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *ComputePhysicalSummary) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *ComputePhysicalSummary) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *ComputePhysicalSummary) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *ComputePhysicalSummary) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *ComputePhysicalSummary) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *ComputePhysicalSummary) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *ComputePhysicalSummary) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *ComputePhysicalSummary) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *ComputePhysicalSummary) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPackageVersion

`func (o *ComputePhysicalSummary) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *ComputePhysicalSummary) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *ComputePhysicalSummary) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *ComputePhysicalSummary) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetPersonality

`func (o *ComputePhysicalSummary) GetPersonality() string`

GetPersonality returns the Personality field if non-nil, zero value otherwise.

### GetPersonalityOk

`func (o *ComputePhysicalSummary) GetPersonalityOk() (*string, bool)`

GetPersonalityOk returns a tuple with the Personality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonality

`func (o *ComputePhysicalSummary) SetPersonality(v string)`

SetPersonality sets Personality field to given value.

### HasPersonality

`func (o *ComputePhysicalSummary) HasPersonality() bool`

HasPersonality returns a boolean if a field has been set.

### GetPlatformType

`func (o *ComputePhysicalSummary) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *ComputePhysicalSummary) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *ComputePhysicalSummary) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *ComputePhysicalSummary) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetPresence

`func (o *ComputePhysicalSummary) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *ComputePhysicalSummary) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *ComputePhysicalSummary) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *ComputePhysicalSummary) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetRevision

`func (o *ComputePhysicalSummary) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ComputePhysicalSummary) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ComputePhysicalSummary) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ComputePhysicalSummary) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRn

`func (o *ComputePhysicalSummary) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *ComputePhysicalSummary) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *ComputePhysicalSummary) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *ComputePhysicalSummary) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetScaledMode

`func (o *ComputePhysicalSummary) GetScaledMode() string`

GetScaledMode returns the ScaledMode field if non-nil, zero value otherwise.

### GetScaledModeOk

`func (o *ComputePhysicalSummary) GetScaledModeOk() (*string, bool)`

GetScaledModeOk returns a tuple with the ScaledMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaledMode

`func (o *ComputePhysicalSummary) SetScaledMode(v string)`

SetScaledMode sets ScaledMode field to given value.

### HasScaledMode

`func (o *ComputePhysicalSummary) HasScaledMode() bool`

HasScaledMode returns a boolean if a field has been set.

### GetSerial

`func (o *ComputePhysicalSummary) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ComputePhysicalSummary) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ComputePhysicalSummary) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ComputePhysicalSummary) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetServerId

`func (o *ComputePhysicalSummary) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ComputePhysicalSummary) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ComputePhysicalSummary) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ComputePhysicalSummary) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServiceProfile

`func (o *ComputePhysicalSummary) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *ComputePhysicalSummary) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *ComputePhysicalSummary) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *ComputePhysicalSummary) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.

### GetSlotId

`func (o *ComputePhysicalSummary) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *ComputePhysicalSummary) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *ComputePhysicalSummary) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *ComputePhysicalSummary) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSourceObjectType

`func (o *ComputePhysicalSummary) GetSourceObjectType() string`

GetSourceObjectType returns the SourceObjectType field if non-nil, zero value otherwise.

### GetSourceObjectTypeOk

`func (o *ComputePhysicalSummary) GetSourceObjectTypeOk() (*string, bool)`

GetSourceObjectTypeOk returns a tuple with the SourceObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectType

`func (o *ComputePhysicalSummary) SetSourceObjectType(v string)`

SetSourceObjectType sets SourceObjectType field to given value.

### HasSourceObjectType

`func (o *ComputePhysicalSummary) HasSourceObjectType() bool`

HasSourceObjectType returns a boolean if a field has been set.

### GetTopologyScanStatus

`func (o *ComputePhysicalSummary) GetTopologyScanStatus() string`

GetTopologyScanStatus returns the TopologyScanStatus field if non-nil, zero value otherwise.

### GetTopologyScanStatusOk

`func (o *ComputePhysicalSummary) GetTopologyScanStatusOk() (*string, bool)`

GetTopologyScanStatusOk returns a tuple with the TopologyScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyScanStatus

`func (o *ComputePhysicalSummary) SetTopologyScanStatus(v string)`

SetTopologyScanStatus sets TopologyScanStatus field to given value.

### HasTopologyScanStatus

`func (o *ComputePhysicalSummary) HasTopologyScanStatus() bool`

HasTopologyScanStatus returns a boolean if a field has been set.

### GetTotalMemory

`func (o *ComputePhysicalSummary) GetTotalMemory() int64`

GetTotalMemory returns the TotalMemory field if non-nil, zero value otherwise.

### GetTotalMemoryOk

`func (o *ComputePhysicalSummary) GetTotalMemoryOk() (*int64, bool)`

GetTotalMemoryOk returns a tuple with the TotalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemory

`func (o *ComputePhysicalSummary) SetTotalMemory(v int64)`

SetTotalMemory sets TotalMemory field to given value.

### HasTotalMemory

`func (o *ComputePhysicalSummary) HasTotalMemory() bool`

HasTotalMemory returns a boolean if a field has been set.

### GetTunneledKvm

`func (o *ComputePhysicalSummary) GetTunneledKvm() bool`

GetTunneledKvm returns the TunneledKvm field if non-nil, zero value otherwise.

### GetTunneledKvmOk

`func (o *ComputePhysicalSummary) GetTunneledKvmOk() (*bool, bool)`

GetTunneledKvmOk returns a tuple with the TunneledKvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledKvm

`func (o *ComputePhysicalSummary) SetTunneledKvm(v bool)`

SetTunneledKvm sets TunneledKvm field to given value.

### HasTunneledKvm

`func (o *ComputePhysicalSummary) HasTunneledKvm() bool`

HasTunneledKvm returns a boolean if a field has been set.

### GetUserLabel

`func (o *ComputePhysicalSummary) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ComputePhysicalSummary) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ComputePhysicalSummary) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ComputePhysicalSummary) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetUuid

`func (o *ComputePhysicalSummary) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ComputePhysicalSummary) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ComputePhysicalSummary) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ComputePhysicalSummary) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendor

`func (o *ComputePhysicalSummary) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ComputePhysicalSummary) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ComputePhysicalSummary) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ComputePhysicalSummary) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetCustomPermissionResources

`func (o *ComputePhysicalSummary) GetCustomPermissionResources() []MoBaseMoRelationship`

GetCustomPermissionResources returns the CustomPermissionResources field if non-nil, zero value otherwise.

### GetCustomPermissionResourcesOk

`func (o *ComputePhysicalSummary) GetCustomPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetCustomPermissionResourcesOk returns a tuple with the CustomPermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPermissionResources

`func (o *ComputePhysicalSummary) SetCustomPermissionResources(v []MoBaseMoRelationship)`

SetCustomPermissionResources sets CustomPermissionResources field to given value.

### HasCustomPermissionResources

`func (o *ComputePhysicalSummary) HasCustomPermissionResources() bool`

HasCustomPermissionResources returns a boolean if a field has been set.

### SetCustomPermissionResourcesNil

`func (o *ComputePhysicalSummary) SetCustomPermissionResourcesNil(b bool)`

 SetCustomPermissionResourcesNil sets the value for CustomPermissionResources to be an explicit nil

### UnsetCustomPermissionResources
`func (o *ComputePhysicalSummary) UnsetCustomPermissionResources()`

UnsetCustomPermissionResources ensures that no value is present for CustomPermissionResources, not even an explicit nil
### GetEquipmentChassis

`func (o *ComputePhysicalSummary) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *ComputePhysicalSummary) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *ComputePhysicalSummary) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *ComputePhysicalSummary) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### SetEquipmentChassisNil

`func (o *ComputePhysicalSummary) SetEquipmentChassisNil(b bool)`

 SetEquipmentChassisNil sets the value for EquipmentChassis to be an explicit nil

### UnsetEquipmentChassis
`func (o *ComputePhysicalSummary) UnsetEquipmentChassis()`

UnsetEquipmentChassis ensures that no value is present for EquipmentChassis, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *ComputePhysicalSummary) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ComputePhysicalSummary) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ComputePhysicalSummary) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ComputePhysicalSummary) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *ComputePhysicalSummary) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *ComputePhysicalSummary) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetInventoryParent

`func (o *ComputePhysicalSummary) GetInventoryParent() MoBaseMoRelationship`

GetInventoryParent returns the InventoryParent field if non-nil, zero value otherwise.

### GetInventoryParentOk

`func (o *ComputePhysicalSummary) GetInventoryParentOk() (*MoBaseMoRelationship, bool)`

GetInventoryParentOk returns a tuple with the InventoryParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryParent

`func (o *ComputePhysicalSummary) SetInventoryParent(v MoBaseMoRelationship)`

SetInventoryParent sets InventoryParent field to given value.

### HasInventoryParent

`func (o *ComputePhysicalSummary) HasInventoryParent() bool`

HasInventoryParent returns a boolean if a field has been set.

### SetInventoryParentNil

`func (o *ComputePhysicalSummary) SetInventoryParentNil(b bool)`

 SetInventoryParentNil sets the value for InventoryParent to be an explicit nil

### UnsetInventoryParent
`func (o *ComputePhysicalSummary) UnsetInventoryParent()`

UnsetInventoryParent ensures that no value is present for InventoryParent, not even an explicit nil
### GetRegisteredDevice

`func (o *ComputePhysicalSummary) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputePhysicalSummary) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputePhysicalSummary) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputePhysicalSummary) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *ComputePhysicalSummary) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ComputePhysicalSummary) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


