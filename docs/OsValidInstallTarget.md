# OsValidInstallTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.ValidInstallTarget"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.ValidInstallTarget"]
**E3sNvme** | Pointer to [**[]OsPhysicalDiskResponse**](OsPhysicalDiskResponse.md) |  | [optional] 
**Error** | Pointer to **string** | Error message if any errors are encountered while fetching and validating Install targets for the server. | [optional] [readonly] 
**InstallTargets** | Pointer to [**[]OsInstallTargetResponse**](OsInstallTargetResponse.md) |  | [optional] 
**M2Jbod** | Pointer to [**[]OsPhysicalDiskResponse**](OsPhysicalDiskResponse.md) |  | [optional] 
**M2NvmeRaidJbod** | Pointer to [**[]OsPhysicalDiskResponse**](OsPhysicalDiskResponse.md) |  | [optional] 
**M2NvmeRaidVirtualDrives** | Pointer to [**[]OsVirtualDriveResponse**](OsVirtualDriveResponse.md) |  | [optional] 
**M2VirtualDrives** | Pointer to [**[]OsVirtualDriveResponse**](OsVirtualDriveResponse.md) |  | [optional] 
**MraidJbod** | Pointer to [**[]OsPhysicalDiskResponse**](OsPhysicalDiskResponse.md) |  | [optional] 
**MraidVirtualDrives** | Pointer to [**[]OsVirtualDriveResponse**](OsVirtualDriveResponse.md) |  | [optional] 
**MstorNvme** | Pointer to [**[]OsPhysicalDiskResponse**](OsPhysicalDiskResponse.md) |  | [optional] 
**Src** | Pointer to **string** | Flag to denote the source of the request. If the call is from Orchestration UI, only the flat list of Install targets can be sent as response. | [optional] 
**U2Nvme** | Pointer to [**[]OsPhysicalDiskResponse**](OsPhysicalDiskResponse.md) |  | [optional] 
**Servers** | Pointer to [**[]ComputePhysicalRelationship**](ComputePhysicalRelationship.md) | An array of relationships to computePhysical resources. | [optional] 

## Methods

### NewOsValidInstallTarget

`func NewOsValidInstallTarget(classId string, objectType string, ) *OsValidInstallTarget`

NewOsValidInstallTarget instantiates a new OsValidInstallTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsValidInstallTargetWithDefaults

`func NewOsValidInstallTargetWithDefaults() *OsValidInstallTarget`

NewOsValidInstallTargetWithDefaults instantiates a new OsValidInstallTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsValidInstallTarget) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsValidInstallTarget) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsValidInstallTarget) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsValidInstallTarget) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsValidInstallTarget) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsValidInstallTarget) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetE3sNvme

`func (o *OsValidInstallTarget) GetE3sNvme() []OsPhysicalDiskResponse`

GetE3sNvme returns the E3sNvme field if non-nil, zero value otherwise.

### GetE3sNvmeOk

`func (o *OsValidInstallTarget) GetE3sNvmeOk() (*[]OsPhysicalDiskResponse, bool)`

GetE3sNvmeOk returns a tuple with the E3sNvme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE3sNvme

`func (o *OsValidInstallTarget) SetE3sNvme(v []OsPhysicalDiskResponse)`

SetE3sNvme sets E3sNvme field to given value.

### HasE3sNvme

`func (o *OsValidInstallTarget) HasE3sNvme() bool`

HasE3sNvme returns a boolean if a field has been set.

### SetE3sNvmeNil

`func (o *OsValidInstallTarget) SetE3sNvmeNil(b bool)`

 SetE3sNvmeNil sets the value for E3sNvme to be an explicit nil

### UnsetE3sNvme
`func (o *OsValidInstallTarget) UnsetE3sNvme()`

UnsetE3sNvme ensures that no value is present for E3sNvme, not even an explicit nil
### GetError

`func (o *OsValidInstallTarget) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OsValidInstallTarget) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OsValidInstallTarget) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *OsValidInstallTarget) HasError() bool`

HasError returns a boolean if a field has been set.

### GetInstallTargets

`func (o *OsValidInstallTarget) GetInstallTargets() []OsInstallTargetResponse`

GetInstallTargets returns the InstallTargets field if non-nil, zero value otherwise.

### GetInstallTargetsOk

`func (o *OsValidInstallTarget) GetInstallTargetsOk() (*[]OsInstallTargetResponse, bool)`

GetInstallTargetsOk returns a tuple with the InstallTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallTargets

`func (o *OsValidInstallTarget) SetInstallTargets(v []OsInstallTargetResponse)`

SetInstallTargets sets InstallTargets field to given value.

### HasInstallTargets

`func (o *OsValidInstallTarget) HasInstallTargets() bool`

HasInstallTargets returns a boolean if a field has been set.

### SetInstallTargetsNil

`func (o *OsValidInstallTarget) SetInstallTargetsNil(b bool)`

 SetInstallTargetsNil sets the value for InstallTargets to be an explicit nil

### UnsetInstallTargets
`func (o *OsValidInstallTarget) UnsetInstallTargets()`

UnsetInstallTargets ensures that no value is present for InstallTargets, not even an explicit nil
### GetM2Jbod

`func (o *OsValidInstallTarget) GetM2Jbod() []OsPhysicalDiskResponse`

GetM2Jbod returns the M2Jbod field if non-nil, zero value otherwise.

### GetM2JbodOk

`func (o *OsValidInstallTarget) GetM2JbodOk() (*[]OsPhysicalDiskResponse, bool)`

GetM2JbodOk returns a tuple with the M2Jbod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM2Jbod

`func (o *OsValidInstallTarget) SetM2Jbod(v []OsPhysicalDiskResponse)`

SetM2Jbod sets M2Jbod field to given value.

### HasM2Jbod

`func (o *OsValidInstallTarget) HasM2Jbod() bool`

HasM2Jbod returns a boolean if a field has been set.

### SetM2JbodNil

`func (o *OsValidInstallTarget) SetM2JbodNil(b bool)`

 SetM2JbodNil sets the value for M2Jbod to be an explicit nil

### UnsetM2Jbod
`func (o *OsValidInstallTarget) UnsetM2Jbod()`

UnsetM2Jbod ensures that no value is present for M2Jbod, not even an explicit nil
### GetM2NvmeRaidJbod

`func (o *OsValidInstallTarget) GetM2NvmeRaidJbod() []OsPhysicalDiskResponse`

GetM2NvmeRaidJbod returns the M2NvmeRaidJbod field if non-nil, zero value otherwise.

### GetM2NvmeRaidJbodOk

`func (o *OsValidInstallTarget) GetM2NvmeRaidJbodOk() (*[]OsPhysicalDiskResponse, bool)`

GetM2NvmeRaidJbodOk returns a tuple with the M2NvmeRaidJbod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM2NvmeRaidJbod

`func (o *OsValidInstallTarget) SetM2NvmeRaidJbod(v []OsPhysicalDiskResponse)`

SetM2NvmeRaidJbod sets M2NvmeRaidJbod field to given value.

### HasM2NvmeRaidJbod

`func (o *OsValidInstallTarget) HasM2NvmeRaidJbod() bool`

HasM2NvmeRaidJbod returns a boolean if a field has been set.

### SetM2NvmeRaidJbodNil

`func (o *OsValidInstallTarget) SetM2NvmeRaidJbodNil(b bool)`

 SetM2NvmeRaidJbodNil sets the value for M2NvmeRaidJbod to be an explicit nil

### UnsetM2NvmeRaidJbod
`func (o *OsValidInstallTarget) UnsetM2NvmeRaidJbod()`

UnsetM2NvmeRaidJbod ensures that no value is present for M2NvmeRaidJbod, not even an explicit nil
### GetM2NvmeRaidVirtualDrives

`func (o *OsValidInstallTarget) GetM2NvmeRaidVirtualDrives() []OsVirtualDriveResponse`

GetM2NvmeRaidVirtualDrives returns the M2NvmeRaidVirtualDrives field if non-nil, zero value otherwise.

### GetM2NvmeRaidVirtualDrivesOk

`func (o *OsValidInstallTarget) GetM2NvmeRaidVirtualDrivesOk() (*[]OsVirtualDriveResponse, bool)`

GetM2NvmeRaidVirtualDrivesOk returns a tuple with the M2NvmeRaidVirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM2NvmeRaidVirtualDrives

`func (o *OsValidInstallTarget) SetM2NvmeRaidVirtualDrives(v []OsVirtualDriveResponse)`

SetM2NvmeRaidVirtualDrives sets M2NvmeRaidVirtualDrives field to given value.

### HasM2NvmeRaidVirtualDrives

`func (o *OsValidInstallTarget) HasM2NvmeRaidVirtualDrives() bool`

HasM2NvmeRaidVirtualDrives returns a boolean if a field has been set.

### SetM2NvmeRaidVirtualDrivesNil

`func (o *OsValidInstallTarget) SetM2NvmeRaidVirtualDrivesNil(b bool)`

 SetM2NvmeRaidVirtualDrivesNil sets the value for M2NvmeRaidVirtualDrives to be an explicit nil

### UnsetM2NvmeRaidVirtualDrives
`func (o *OsValidInstallTarget) UnsetM2NvmeRaidVirtualDrives()`

UnsetM2NvmeRaidVirtualDrives ensures that no value is present for M2NvmeRaidVirtualDrives, not even an explicit nil
### GetM2VirtualDrives

`func (o *OsValidInstallTarget) GetM2VirtualDrives() []OsVirtualDriveResponse`

GetM2VirtualDrives returns the M2VirtualDrives field if non-nil, zero value otherwise.

### GetM2VirtualDrivesOk

`func (o *OsValidInstallTarget) GetM2VirtualDrivesOk() (*[]OsVirtualDriveResponse, bool)`

GetM2VirtualDrivesOk returns a tuple with the M2VirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM2VirtualDrives

`func (o *OsValidInstallTarget) SetM2VirtualDrives(v []OsVirtualDriveResponse)`

SetM2VirtualDrives sets M2VirtualDrives field to given value.

### HasM2VirtualDrives

`func (o *OsValidInstallTarget) HasM2VirtualDrives() bool`

HasM2VirtualDrives returns a boolean if a field has been set.

### SetM2VirtualDrivesNil

`func (o *OsValidInstallTarget) SetM2VirtualDrivesNil(b bool)`

 SetM2VirtualDrivesNil sets the value for M2VirtualDrives to be an explicit nil

### UnsetM2VirtualDrives
`func (o *OsValidInstallTarget) UnsetM2VirtualDrives()`

UnsetM2VirtualDrives ensures that no value is present for M2VirtualDrives, not even an explicit nil
### GetMraidJbod

`func (o *OsValidInstallTarget) GetMraidJbod() []OsPhysicalDiskResponse`

GetMraidJbod returns the MraidJbod field if non-nil, zero value otherwise.

### GetMraidJbodOk

`func (o *OsValidInstallTarget) GetMraidJbodOk() (*[]OsPhysicalDiskResponse, bool)`

GetMraidJbodOk returns a tuple with the MraidJbod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMraidJbod

`func (o *OsValidInstallTarget) SetMraidJbod(v []OsPhysicalDiskResponse)`

SetMraidJbod sets MraidJbod field to given value.

### HasMraidJbod

`func (o *OsValidInstallTarget) HasMraidJbod() bool`

HasMraidJbod returns a boolean if a field has been set.

### SetMraidJbodNil

`func (o *OsValidInstallTarget) SetMraidJbodNil(b bool)`

 SetMraidJbodNil sets the value for MraidJbod to be an explicit nil

### UnsetMraidJbod
`func (o *OsValidInstallTarget) UnsetMraidJbod()`

UnsetMraidJbod ensures that no value is present for MraidJbod, not even an explicit nil
### GetMraidVirtualDrives

`func (o *OsValidInstallTarget) GetMraidVirtualDrives() []OsVirtualDriveResponse`

GetMraidVirtualDrives returns the MraidVirtualDrives field if non-nil, zero value otherwise.

### GetMraidVirtualDrivesOk

`func (o *OsValidInstallTarget) GetMraidVirtualDrivesOk() (*[]OsVirtualDriveResponse, bool)`

GetMraidVirtualDrivesOk returns a tuple with the MraidVirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMraidVirtualDrives

`func (o *OsValidInstallTarget) SetMraidVirtualDrives(v []OsVirtualDriveResponse)`

SetMraidVirtualDrives sets MraidVirtualDrives field to given value.

### HasMraidVirtualDrives

`func (o *OsValidInstallTarget) HasMraidVirtualDrives() bool`

HasMraidVirtualDrives returns a boolean if a field has been set.

### SetMraidVirtualDrivesNil

`func (o *OsValidInstallTarget) SetMraidVirtualDrivesNil(b bool)`

 SetMraidVirtualDrivesNil sets the value for MraidVirtualDrives to be an explicit nil

### UnsetMraidVirtualDrives
`func (o *OsValidInstallTarget) UnsetMraidVirtualDrives()`

UnsetMraidVirtualDrives ensures that no value is present for MraidVirtualDrives, not even an explicit nil
### GetMstorNvme

`func (o *OsValidInstallTarget) GetMstorNvme() []OsPhysicalDiskResponse`

GetMstorNvme returns the MstorNvme field if non-nil, zero value otherwise.

### GetMstorNvmeOk

`func (o *OsValidInstallTarget) GetMstorNvmeOk() (*[]OsPhysicalDiskResponse, bool)`

GetMstorNvmeOk returns a tuple with the MstorNvme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMstorNvme

`func (o *OsValidInstallTarget) SetMstorNvme(v []OsPhysicalDiskResponse)`

SetMstorNvme sets MstorNvme field to given value.

### HasMstorNvme

`func (o *OsValidInstallTarget) HasMstorNvme() bool`

HasMstorNvme returns a boolean if a field has been set.

### SetMstorNvmeNil

`func (o *OsValidInstallTarget) SetMstorNvmeNil(b bool)`

 SetMstorNvmeNil sets the value for MstorNvme to be an explicit nil

### UnsetMstorNvme
`func (o *OsValidInstallTarget) UnsetMstorNvme()`

UnsetMstorNvme ensures that no value is present for MstorNvme, not even an explicit nil
### GetSrc

`func (o *OsValidInstallTarget) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *OsValidInstallTarget) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *OsValidInstallTarget) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *OsValidInstallTarget) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetU2Nvme

`func (o *OsValidInstallTarget) GetU2Nvme() []OsPhysicalDiskResponse`

GetU2Nvme returns the U2Nvme field if non-nil, zero value otherwise.

### GetU2NvmeOk

`func (o *OsValidInstallTarget) GetU2NvmeOk() (*[]OsPhysicalDiskResponse, bool)`

GetU2NvmeOk returns a tuple with the U2Nvme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU2Nvme

`func (o *OsValidInstallTarget) SetU2Nvme(v []OsPhysicalDiskResponse)`

SetU2Nvme sets U2Nvme field to given value.

### HasU2Nvme

`func (o *OsValidInstallTarget) HasU2Nvme() bool`

HasU2Nvme returns a boolean if a field has been set.

### SetU2NvmeNil

`func (o *OsValidInstallTarget) SetU2NvmeNil(b bool)`

 SetU2NvmeNil sets the value for U2Nvme to be an explicit nil

### UnsetU2Nvme
`func (o *OsValidInstallTarget) UnsetU2Nvme()`

UnsetU2Nvme ensures that no value is present for U2Nvme, not even an explicit nil
### GetServers

`func (o *OsValidInstallTarget) GetServers() []ComputePhysicalRelationship`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *OsValidInstallTarget) GetServersOk() (*[]ComputePhysicalRelationship, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *OsValidInstallTarget) SetServers(v []ComputePhysicalRelationship)`

SetServers sets Servers field to given value.

### HasServers

`func (o *OsValidInstallTarget) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *OsValidInstallTarget) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *OsValidInstallTarget) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


