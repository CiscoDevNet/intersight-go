# VnicFcVhbaPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcVhbaPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcVhbaPolicyInventory"]
**Cos** | Pointer to **int64** | Class of Service to be associated to the traffic on the virtual interface. | [optional] [default to 3]
**MaxDataFieldSize** | Pointer to **int64** | The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports. | [optional] [default to 2112]
**Name** | Pointer to **string** | Name of the virtual HBA interface. | [optional] 
**Priority** | Pointer to **string** | The priortity matching the System QoS specified in the fabric profile. * &#x60;Best Effort&#x60; - QoS Priority for Best-effort traffic. * &#x60;FC&#x60; - QoS Priority for FC traffic. * &#x60;Platinum&#x60; - QoS Priority for Platinum traffic. * &#x60;Gold&#x60; - QoS Priority for Gold traffic. * &#x60;Silver&#x60; - QoS Priority for Silver traffic. * &#x60;Bronze&#x60; - QoS Priority for Bronze traffic. | [optional] [readonly] [default to "Best Effort"]
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVnicFcVhbaPolicyInventory

`func NewVnicFcVhbaPolicyInventory(classId string, objectType string, ) *VnicFcVhbaPolicyInventory`

NewVnicFcVhbaPolicyInventory instantiates a new VnicFcVhbaPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcVhbaPolicyInventoryWithDefaults

`func NewVnicFcVhbaPolicyInventoryWithDefaults() *VnicFcVhbaPolicyInventory`

NewVnicFcVhbaPolicyInventoryWithDefaults instantiates a new VnicFcVhbaPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcVhbaPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcVhbaPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcVhbaPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcVhbaPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcVhbaPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcVhbaPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCos

`func (o *VnicFcVhbaPolicyInventory) GetCos() int64`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *VnicFcVhbaPolicyInventory) GetCosOk() (*int64, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *VnicFcVhbaPolicyInventory) SetCos(v int64)`

SetCos sets Cos field to given value.

### HasCos

`func (o *VnicFcVhbaPolicyInventory) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetMaxDataFieldSize

`func (o *VnicFcVhbaPolicyInventory) GetMaxDataFieldSize() int64`

GetMaxDataFieldSize returns the MaxDataFieldSize field if non-nil, zero value otherwise.

### GetMaxDataFieldSizeOk

`func (o *VnicFcVhbaPolicyInventory) GetMaxDataFieldSizeOk() (*int64, bool)`

GetMaxDataFieldSizeOk returns a tuple with the MaxDataFieldSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataFieldSize

`func (o *VnicFcVhbaPolicyInventory) SetMaxDataFieldSize(v int64)`

SetMaxDataFieldSize sets MaxDataFieldSize field to given value.

### HasMaxDataFieldSize

`func (o *VnicFcVhbaPolicyInventory) HasMaxDataFieldSize() bool`

HasMaxDataFieldSize returns a boolean if a field has been set.

### GetName

`func (o *VnicFcVhbaPolicyInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicFcVhbaPolicyInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicFcVhbaPolicyInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicFcVhbaPolicyInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *VnicFcVhbaPolicyInventory) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *VnicFcVhbaPolicyInventory) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *VnicFcVhbaPolicyInventory) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *VnicFcVhbaPolicyInventory) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTargetMo

`func (o *VnicFcVhbaPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicFcVhbaPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicFcVhbaPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicFcVhbaPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *VnicFcVhbaPolicyInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *VnicFcVhbaPolicyInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


