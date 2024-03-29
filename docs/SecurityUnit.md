# SecurityUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "security.Unit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "security.Unit"]
**OperState** | Pointer to **string** | Operational state of the security unit. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability state of the security unit. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | The part number of the security unit. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | PCIe slot of the security unit in the server. | [optional] [readonly] 
**Power** | Pointer to **string** | Power state of the security unit. | [optional] [readonly] 
**Thermal** | Pointer to **string** | Thermal state of the security unit. | [optional] [readonly] 
**UnitId** | Pointer to **int64** | The unique identifier assigned to the security unit within the server. | [optional] [readonly] 
**Vid** | Pointer to **string** | The vendor identifier of the security unit. | [optional] [readonly] 
**Voltage** | Pointer to **string** | The voltage state of the security unit. | [optional] [readonly] 
**ComputeBoard** | Pointer to [**NullableComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewSecurityUnit

`func NewSecurityUnit(classId string, objectType string, ) *SecurityUnit`

NewSecurityUnit instantiates a new SecurityUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityUnitWithDefaults

`func NewSecurityUnitWithDefaults() *SecurityUnit`

NewSecurityUnitWithDefaults instantiates a new SecurityUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SecurityUnit) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SecurityUnit) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SecurityUnit) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SecurityUnit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SecurityUnit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SecurityUnit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperState

`func (o *SecurityUnit) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *SecurityUnit) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *SecurityUnit) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *SecurityUnit) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *SecurityUnit) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *SecurityUnit) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *SecurityUnit) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *SecurityUnit) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPartNumber

`func (o *SecurityUnit) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *SecurityUnit) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *SecurityUnit) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *SecurityUnit) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPciSlot

`func (o *SecurityUnit) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *SecurityUnit) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *SecurityUnit) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *SecurityUnit) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetPower

`func (o *SecurityUnit) GetPower() string`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *SecurityUnit) GetPowerOk() (*string, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *SecurityUnit) SetPower(v string)`

SetPower sets Power field to given value.

### HasPower

`func (o *SecurityUnit) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetThermal

`func (o *SecurityUnit) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *SecurityUnit) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *SecurityUnit) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *SecurityUnit) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetUnitId

`func (o *SecurityUnit) GetUnitId() int64`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *SecurityUnit) GetUnitIdOk() (*int64, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *SecurityUnit) SetUnitId(v int64)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *SecurityUnit) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### GetVid

`func (o *SecurityUnit) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *SecurityUnit) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *SecurityUnit) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *SecurityUnit) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetVoltage

`func (o *SecurityUnit) GetVoltage() string`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *SecurityUnit) GetVoltageOk() (*string, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *SecurityUnit) SetVoltage(v string)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *SecurityUnit) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetComputeBoard

`func (o *SecurityUnit) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *SecurityUnit) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *SecurityUnit) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *SecurityUnit) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### SetComputeBoardNil

`func (o *SecurityUnit) SetComputeBoardNil(b bool)`

 SetComputeBoardNil sets the value for ComputeBoard to be an explicit nil

### UnsetComputeBoard
`func (o *SecurityUnit) UnsetComputeBoard()`

UnsetComputeBoard ensures that no value is present for ComputeBoard, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *SecurityUnit) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *SecurityUnit) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *SecurityUnit) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *SecurityUnit) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *SecurityUnit) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *SecurityUnit) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetRegisteredDevice

`func (o *SecurityUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *SecurityUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *SecurityUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *SecurityUnit) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *SecurityUnit) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *SecurityUnit) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


