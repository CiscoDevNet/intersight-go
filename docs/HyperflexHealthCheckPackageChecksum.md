# HyperflexHealthCheckPackageChecksum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HealthCheckPackageChecksum"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HealthCheckPackageChecksum"]
**Checksum** | Pointer to **string** | SHA512 checksum of the health check package. | [optional] 
**Name** | Pointer to **string** | Name of the health check Debian package. | [optional] 
**PackageName** | Pointer to **string** | HyperFlex health check Debian package file name. | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp of last update of HyperFlex health check package checksum. | [optional] 
**Version** | Pointer to **string** | HyperFlex health check Debian Package Version. | [optional] 

## Methods

### NewHyperflexHealthCheckPackageChecksum

`func NewHyperflexHealthCheckPackageChecksum(classId string, objectType string, ) *HyperflexHealthCheckPackageChecksum`

NewHyperflexHealthCheckPackageChecksum instantiates a new HyperflexHealthCheckPackageChecksum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHealthCheckPackageChecksumWithDefaults

`func NewHyperflexHealthCheckPackageChecksumWithDefaults() *HyperflexHealthCheckPackageChecksum`

NewHyperflexHealthCheckPackageChecksumWithDefaults instantiates a new HyperflexHealthCheckPackageChecksum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHealthCheckPackageChecksum) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHealthCheckPackageChecksum) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHealthCheckPackageChecksum) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHealthCheckPackageChecksum) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHealthCheckPackageChecksum) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHealthCheckPackageChecksum) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChecksum

`func (o *HyperflexHealthCheckPackageChecksum) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *HyperflexHealthCheckPackageChecksum) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *HyperflexHealthCheckPackageChecksum) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *HyperflexHealthCheckPackageChecksum) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetName

`func (o *HyperflexHealthCheckPackageChecksum) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexHealthCheckPackageChecksum) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexHealthCheckPackageChecksum) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexHealthCheckPackageChecksum) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageName

`func (o *HyperflexHealthCheckPackageChecksum) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *HyperflexHealthCheckPackageChecksum) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *HyperflexHealthCheckPackageChecksum) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *HyperflexHealthCheckPackageChecksum) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetTimestamp

`func (o *HyperflexHealthCheckPackageChecksum) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HyperflexHealthCheckPackageChecksum) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HyperflexHealthCheckPackageChecksum) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *HyperflexHealthCheckPackageChecksum) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexHealthCheckPackageChecksum) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexHealthCheckPackageChecksum) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexHealthCheckPackageChecksum) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexHealthCheckPackageChecksum) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


