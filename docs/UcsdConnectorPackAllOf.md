# UcsdConnectorPackAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ucsd.ConnectorPack"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ucsd.ConnectorPack"]
**ConnectorFeature** | Pointer to **string** | State of the connector pack whether it is enabled or disabled. | [optional] [readonly] 
**DependencyNames** | Pointer to **[]string** |  | [optional] 
**DownloadedVersion** | Pointer to **string** | Version of the connector pack that is last downloaded successfully to UCS Director. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the connector pack running on the UCS Director. | [optional] [readonly] 
**Services** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** | State of the connector pack whether it is enabled or disabled. | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the connector pack. | [optional] [readonly] 

## Methods

### NewUcsdConnectorPackAllOf

`func NewUcsdConnectorPackAllOf(classId string, objectType string, ) *UcsdConnectorPackAllOf`

NewUcsdConnectorPackAllOf instantiates a new UcsdConnectorPackAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUcsdConnectorPackAllOfWithDefaults

`func NewUcsdConnectorPackAllOfWithDefaults() *UcsdConnectorPackAllOf`

NewUcsdConnectorPackAllOfWithDefaults instantiates a new UcsdConnectorPackAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *UcsdConnectorPackAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *UcsdConnectorPackAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *UcsdConnectorPackAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *UcsdConnectorPackAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UcsdConnectorPackAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UcsdConnectorPackAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectorFeature

`func (o *UcsdConnectorPackAllOf) GetConnectorFeature() string`

GetConnectorFeature returns the ConnectorFeature field if non-nil, zero value otherwise.

### GetConnectorFeatureOk

`func (o *UcsdConnectorPackAllOf) GetConnectorFeatureOk() (*string, bool)`

GetConnectorFeatureOk returns a tuple with the ConnectorFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorFeature

`func (o *UcsdConnectorPackAllOf) SetConnectorFeature(v string)`

SetConnectorFeature sets ConnectorFeature field to given value.

### HasConnectorFeature

`func (o *UcsdConnectorPackAllOf) HasConnectorFeature() bool`

HasConnectorFeature returns a boolean if a field has been set.

### GetDependencyNames

`func (o *UcsdConnectorPackAllOf) GetDependencyNames() []string`

GetDependencyNames returns the DependencyNames field if non-nil, zero value otherwise.

### GetDependencyNamesOk

`func (o *UcsdConnectorPackAllOf) GetDependencyNamesOk() (*[]string, bool)`

GetDependencyNamesOk returns a tuple with the DependencyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyNames

`func (o *UcsdConnectorPackAllOf) SetDependencyNames(v []string)`

SetDependencyNames sets DependencyNames field to given value.

### HasDependencyNames

`func (o *UcsdConnectorPackAllOf) HasDependencyNames() bool`

HasDependencyNames returns a boolean if a field has been set.

### SetDependencyNamesNil

`func (o *UcsdConnectorPackAllOf) SetDependencyNamesNil(b bool)`

 SetDependencyNamesNil sets the value for DependencyNames to be an explicit nil

### UnsetDependencyNames
`func (o *UcsdConnectorPackAllOf) UnsetDependencyNames()`

UnsetDependencyNames ensures that no value is present for DependencyNames, not even an explicit nil
### GetDownloadedVersion

`func (o *UcsdConnectorPackAllOf) GetDownloadedVersion() string`

GetDownloadedVersion returns the DownloadedVersion field if non-nil, zero value otherwise.

### GetDownloadedVersionOk

`func (o *UcsdConnectorPackAllOf) GetDownloadedVersionOk() (*string, bool)`

GetDownloadedVersionOk returns a tuple with the DownloadedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedVersion

`func (o *UcsdConnectorPackAllOf) SetDownloadedVersion(v string)`

SetDownloadedVersion sets DownloadedVersion field to given value.

### HasDownloadedVersion

`func (o *UcsdConnectorPackAllOf) HasDownloadedVersion() bool`

HasDownloadedVersion returns a boolean if a field has been set.

### GetName

`func (o *UcsdConnectorPackAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UcsdConnectorPackAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UcsdConnectorPackAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UcsdConnectorPackAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServices

`func (o *UcsdConnectorPackAllOf) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *UcsdConnectorPackAllOf) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *UcsdConnectorPackAllOf) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *UcsdConnectorPackAllOf) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *UcsdConnectorPackAllOf) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *UcsdConnectorPackAllOf) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetState

`func (o *UcsdConnectorPackAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UcsdConnectorPackAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UcsdConnectorPackAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UcsdConnectorPackAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVersion

`func (o *UcsdConnectorPackAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UcsdConnectorPackAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UcsdConnectorPackAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UcsdConnectorPackAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


