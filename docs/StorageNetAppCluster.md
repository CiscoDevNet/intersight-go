# StorageNetAppCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppCluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppCluster"]
**AutoSupport** | Pointer to [**NullableStorageNetAppAutoSupport**](StorageNetAppAutoSupport.md) |  | [optional] 
**AvgPerformanceMetrics** | Pointer to [**NullableStorageBasePerformanceMetricsAverage**](StorageBasePerformanceMetricsAverage.md) | Average performance metrics data for a NetApp storage resource over a given period of time. | [optional] 
**ClusterEfficiency** | Pointer to [**NullableStorageNetAppStorageClusterEfficiency**](StorageNetAppStorageClusterEfficiency.md) |  | [optional] 
**ClusterHealthStatus** | Pointer to **string** | The health status of the cluster. Possible states are ok, ok-with-suppressed, degraded, and unreachable. * &#x60;Unreachable&#x60; - Cluster status is unreachable. * &#x60;OK&#x60; - Cluster status is either ok or ok-with-suppressed. * &#x60;Degraded&#x60; - Cluster status is degraded. | [optional] [readonly] [default to "Unreachable"]
**DefaultAdminLocked** | Pointer to **bool** | Indicates whether the default admin user is locked out. | [optional] [readonly] 
**DnsDomains** | Pointer to **[]string** |  | [optional] 
**FipsCompliant** | Pointer to **bool** | Indicates whether or not the software FIPS mode is enabled on the cluster. | [optional] [readonly] 
**InsecureCiphers** | Pointer to **int64** | Number of SVMs on the cluster that use insecure ciphers. | [optional] [readonly] 
**Key** | Pointer to **string** | Unique identifier of NetApp Cluster across data center. | [optional] [readonly] 
**Location** | Pointer to **string** | Location of the storage controller. | [optional] [readonly] 
**ManagementAddress** | Pointer to **[]string** |  | [optional] 
**NameServers** | Pointer to **[]string** |  | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**RshEnabled** | Pointer to **bool** | Indicates whether or not rsh is enabled on the cluster. | [optional] [readonly] 
**TelnetEnabled** | Pointer to **bool** | Indicates whether or not telnet is enabled on the cluster. | [optional] [readonly] 
**VersionGeneration** | Pointer to **int64** | The generation portion of the version. | [optional] [readonly] 
**VersionMajor** | Pointer to **int64** | The major portion of the version. | [optional] [readonly] 
**VersionMinor** | Pointer to **int64** | The minor portion of the version. | [optional] [readonly] 
**Events** | Pointer to [**[]StorageNetAppClusterEventRelationship**](StorageNetAppClusterEventRelationship.md) | An array of relationships to storageNetAppClusterEvent resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppCluster

`func NewStorageNetAppCluster(classId string, objectType string, ) *StorageNetAppCluster`

NewStorageNetAppCluster instantiates a new StorageNetAppCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppClusterWithDefaults

`func NewStorageNetAppClusterWithDefaults() *StorageNetAppCluster`

NewStorageNetAppClusterWithDefaults instantiates a new StorageNetAppCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoSupport

`func (o *StorageNetAppCluster) GetAutoSupport() StorageNetAppAutoSupport`

GetAutoSupport returns the AutoSupport field if non-nil, zero value otherwise.

### GetAutoSupportOk

`func (o *StorageNetAppCluster) GetAutoSupportOk() (*StorageNetAppAutoSupport, bool)`

GetAutoSupportOk returns a tuple with the AutoSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSupport

`func (o *StorageNetAppCluster) SetAutoSupport(v StorageNetAppAutoSupport)`

SetAutoSupport sets AutoSupport field to given value.

### HasAutoSupport

`func (o *StorageNetAppCluster) HasAutoSupport() bool`

HasAutoSupport returns a boolean if a field has been set.

### SetAutoSupportNil

`func (o *StorageNetAppCluster) SetAutoSupportNil(b bool)`

 SetAutoSupportNil sets the value for AutoSupport to be an explicit nil

### UnsetAutoSupport
`func (o *StorageNetAppCluster) UnsetAutoSupport()`

UnsetAutoSupport ensures that no value is present for AutoSupport, not even an explicit nil
### GetAvgPerformanceMetrics

`func (o *StorageNetAppCluster) GetAvgPerformanceMetrics() StorageBasePerformanceMetricsAverage`

GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field if non-nil, zero value otherwise.

### GetAvgPerformanceMetricsOk

`func (o *StorageNetAppCluster) GetAvgPerformanceMetricsOk() (*StorageBasePerformanceMetricsAverage, bool)`

GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPerformanceMetrics

`func (o *StorageNetAppCluster) SetAvgPerformanceMetrics(v StorageBasePerformanceMetricsAverage)`

SetAvgPerformanceMetrics sets AvgPerformanceMetrics field to given value.

### HasAvgPerformanceMetrics

`func (o *StorageNetAppCluster) HasAvgPerformanceMetrics() bool`

HasAvgPerformanceMetrics returns a boolean if a field has been set.

### SetAvgPerformanceMetricsNil

`func (o *StorageNetAppCluster) SetAvgPerformanceMetricsNil(b bool)`

 SetAvgPerformanceMetricsNil sets the value for AvgPerformanceMetrics to be an explicit nil

### UnsetAvgPerformanceMetrics
`func (o *StorageNetAppCluster) UnsetAvgPerformanceMetrics()`

UnsetAvgPerformanceMetrics ensures that no value is present for AvgPerformanceMetrics, not even an explicit nil
### GetClusterEfficiency

`func (o *StorageNetAppCluster) GetClusterEfficiency() StorageNetAppStorageClusterEfficiency`

GetClusterEfficiency returns the ClusterEfficiency field if non-nil, zero value otherwise.

### GetClusterEfficiencyOk

`func (o *StorageNetAppCluster) GetClusterEfficiencyOk() (*StorageNetAppStorageClusterEfficiency, bool)`

GetClusterEfficiencyOk returns a tuple with the ClusterEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEfficiency

`func (o *StorageNetAppCluster) SetClusterEfficiency(v StorageNetAppStorageClusterEfficiency)`

SetClusterEfficiency sets ClusterEfficiency field to given value.

### HasClusterEfficiency

`func (o *StorageNetAppCluster) HasClusterEfficiency() bool`

HasClusterEfficiency returns a boolean if a field has been set.

### SetClusterEfficiencyNil

`func (o *StorageNetAppCluster) SetClusterEfficiencyNil(b bool)`

 SetClusterEfficiencyNil sets the value for ClusterEfficiency to be an explicit nil

### UnsetClusterEfficiency
`func (o *StorageNetAppCluster) UnsetClusterEfficiency()`

UnsetClusterEfficiency ensures that no value is present for ClusterEfficiency, not even an explicit nil
### GetClusterHealthStatus

`func (o *StorageNetAppCluster) GetClusterHealthStatus() string`

GetClusterHealthStatus returns the ClusterHealthStatus field if non-nil, zero value otherwise.

### GetClusterHealthStatusOk

`func (o *StorageNetAppCluster) GetClusterHealthStatusOk() (*string, bool)`

GetClusterHealthStatusOk returns a tuple with the ClusterHealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHealthStatus

`func (o *StorageNetAppCluster) SetClusterHealthStatus(v string)`

SetClusterHealthStatus sets ClusterHealthStatus field to given value.

### HasClusterHealthStatus

`func (o *StorageNetAppCluster) HasClusterHealthStatus() bool`

HasClusterHealthStatus returns a boolean if a field has been set.

### GetDefaultAdminLocked

`func (o *StorageNetAppCluster) GetDefaultAdminLocked() bool`

GetDefaultAdminLocked returns the DefaultAdminLocked field if non-nil, zero value otherwise.

### GetDefaultAdminLockedOk

`func (o *StorageNetAppCluster) GetDefaultAdminLockedOk() (*bool, bool)`

GetDefaultAdminLockedOk returns a tuple with the DefaultAdminLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAdminLocked

`func (o *StorageNetAppCluster) SetDefaultAdminLocked(v bool)`

SetDefaultAdminLocked sets DefaultAdminLocked field to given value.

### HasDefaultAdminLocked

`func (o *StorageNetAppCluster) HasDefaultAdminLocked() bool`

HasDefaultAdminLocked returns a boolean if a field has been set.

### GetDnsDomains

`func (o *StorageNetAppCluster) GetDnsDomains() []string`

GetDnsDomains returns the DnsDomains field if non-nil, zero value otherwise.

### GetDnsDomainsOk

`func (o *StorageNetAppCluster) GetDnsDomainsOk() (*[]string, bool)`

GetDnsDomainsOk returns a tuple with the DnsDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomains

`func (o *StorageNetAppCluster) SetDnsDomains(v []string)`

SetDnsDomains sets DnsDomains field to given value.

### HasDnsDomains

`func (o *StorageNetAppCluster) HasDnsDomains() bool`

HasDnsDomains returns a boolean if a field has been set.

### SetDnsDomainsNil

`func (o *StorageNetAppCluster) SetDnsDomainsNil(b bool)`

 SetDnsDomainsNil sets the value for DnsDomains to be an explicit nil

### UnsetDnsDomains
`func (o *StorageNetAppCluster) UnsetDnsDomains()`

UnsetDnsDomains ensures that no value is present for DnsDomains, not even an explicit nil
### GetFipsCompliant

`func (o *StorageNetAppCluster) GetFipsCompliant() bool`

GetFipsCompliant returns the FipsCompliant field if non-nil, zero value otherwise.

### GetFipsCompliantOk

`func (o *StorageNetAppCluster) GetFipsCompliantOk() (*bool, bool)`

GetFipsCompliantOk returns a tuple with the FipsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsCompliant

`func (o *StorageNetAppCluster) SetFipsCompliant(v bool)`

SetFipsCompliant sets FipsCompliant field to given value.

### HasFipsCompliant

`func (o *StorageNetAppCluster) HasFipsCompliant() bool`

HasFipsCompliant returns a boolean if a field has been set.

### GetInsecureCiphers

`func (o *StorageNetAppCluster) GetInsecureCiphers() int64`

GetInsecureCiphers returns the InsecureCiphers field if non-nil, zero value otherwise.

### GetInsecureCiphersOk

`func (o *StorageNetAppCluster) GetInsecureCiphersOk() (*int64, bool)`

GetInsecureCiphersOk returns a tuple with the InsecureCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureCiphers

`func (o *StorageNetAppCluster) SetInsecureCiphers(v int64)`

SetInsecureCiphers sets InsecureCiphers field to given value.

### HasInsecureCiphers

`func (o *StorageNetAppCluster) HasInsecureCiphers() bool`

HasInsecureCiphers returns a boolean if a field has been set.

### GetKey

`func (o *StorageNetAppCluster) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorageNetAppCluster) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorageNetAppCluster) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StorageNetAppCluster) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLocation

`func (o *StorageNetAppCluster) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *StorageNetAppCluster) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *StorageNetAppCluster) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *StorageNetAppCluster) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetManagementAddress

`func (o *StorageNetAppCluster) GetManagementAddress() []string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *StorageNetAppCluster) GetManagementAddressOk() (*[]string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *StorageNetAppCluster) SetManagementAddress(v []string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *StorageNetAppCluster) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### SetManagementAddressNil

`func (o *StorageNetAppCluster) SetManagementAddressNil(b bool)`

 SetManagementAddressNil sets the value for ManagementAddress to be an explicit nil

### UnsetManagementAddress
`func (o *StorageNetAppCluster) UnsetManagementAddress()`

UnsetManagementAddress ensures that no value is present for ManagementAddress, not even an explicit nil
### GetNameServers

`func (o *StorageNetAppCluster) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *StorageNetAppCluster) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *StorageNetAppCluster) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *StorageNetAppCluster) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### SetNameServersNil

`func (o *StorageNetAppCluster) SetNameServersNil(b bool)`

 SetNameServersNil sets the value for NameServers to be an explicit nil

### UnsetNameServers
`func (o *StorageNetAppCluster) UnsetNameServers()`

UnsetNameServers ensures that no value is present for NameServers, not even an explicit nil
### GetNtpServers

`func (o *StorageNetAppCluster) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *StorageNetAppCluster) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *StorageNetAppCluster) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *StorageNetAppCluster) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServersNil

`func (o *StorageNetAppCluster) SetNtpServersNil(b bool)`

 SetNtpServersNil sets the value for NtpServers to be an explicit nil

### UnsetNtpServers
`func (o *StorageNetAppCluster) UnsetNtpServers()`

UnsetNtpServers ensures that no value is present for NtpServers, not even an explicit nil
### GetRshEnabled

`func (o *StorageNetAppCluster) GetRshEnabled() bool`

GetRshEnabled returns the RshEnabled field if non-nil, zero value otherwise.

### GetRshEnabledOk

`func (o *StorageNetAppCluster) GetRshEnabledOk() (*bool, bool)`

GetRshEnabledOk returns a tuple with the RshEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRshEnabled

`func (o *StorageNetAppCluster) SetRshEnabled(v bool)`

SetRshEnabled sets RshEnabled field to given value.

### HasRshEnabled

`func (o *StorageNetAppCluster) HasRshEnabled() bool`

HasRshEnabled returns a boolean if a field has been set.

### GetTelnetEnabled

`func (o *StorageNetAppCluster) GetTelnetEnabled() bool`

GetTelnetEnabled returns the TelnetEnabled field if non-nil, zero value otherwise.

### GetTelnetEnabledOk

`func (o *StorageNetAppCluster) GetTelnetEnabledOk() (*bool, bool)`

GetTelnetEnabledOk returns a tuple with the TelnetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnetEnabled

`func (o *StorageNetAppCluster) SetTelnetEnabled(v bool)`

SetTelnetEnabled sets TelnetEnabled field to given value.

### HasTelnetEnabled

`func (o *StorageNetAppCluster) HasTelnetEnabled() bool`

HasTelnetEnabled returns a boolean if a field has been set.

### GetVersionGeneration

`func (o *StorageNetAppCluster) GetVersionGeneration() int64`

GetVersionGeneration returns the VersionGeneration field if non-nil, zero value otherwise.

### GetVersionGenerationOk

`func (o *StorageNetAppCluster) GetVersionGenerationOk() (*int64, bool)`

GetVersionGenerationOk returns a tuple with the VersionGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGeneration

`func (o *StorageNetAppCluster) SetVersionGeneration(v int64)`

SetVersionGeneration sets VersionGeneration field to given value.

### HasVersionGeneration

`func (o *StorageNetAppCluster) HasVersionGeneration() bool`

HasVersionGeneration returns a boolean if a field has been set.

### GetVersionMajor

`func (o *StorageNetAppCluster) GetVersionMajor() int64`

GetVersionMajor returns the VersionMajor field if non-nil, zero value otherwise.

### GetVersionMajorOk

`func (o *StorageNetAppCluster) GetVersionMajorOk() (*int64, bool)`

GetVersionMajorOk returns a tuple with the VersionMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMajor

`func (o *StorageNetAppCluster) SetVersionMajor(v int64)`

SetVersionMajor sets VersionMajor field to given value.

### HasVersionMajor

`func (o *StorageNetAppCluster) HasVersionMajor() bool`

HasVersionMajor returns a boolean if a field has been set.

### GetVersionMinor

`func (o *StorageNetAppCluster) GetVersionMinor() int64`

GetVersionMinor returns the VersionMinor field if non-nil, zero value otherwise.

### GetVersionMinorOk

`func (o *StorageNetAppCluster) GetVersionMinorOk() (*int64, bool)`

GetVersionMinorOk returns a tuple with the VersionMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMinor

`func (o *StorageNetAppCluster) SetVersionMinor(v int64)`

SetVersionMinor sets VersionMinor field to given value.

### HasVersionMinor

`func (o *StorageNetAppCluster) HasVersionMinor() bool`

HasVersionMinor returns a boolean if a field has been set.

### GetEvents

`func (o *StorageNetAppCluster) GetEvents() []StorageNetAppClusterEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppCluster) GetEventsOk() (*[]StorageNetAppClusterEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppCluster) SetEvents(v []StorageNetAppClusterEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppCluster) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppCluster) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppCluster) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageNetAppCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageNetAppCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageNetAppCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageNetAppCluster) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StorageNetAppCluster) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StorageNetAppCluster) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


