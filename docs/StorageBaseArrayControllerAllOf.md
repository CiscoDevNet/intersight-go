# StorageBaseArrayControllerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Name** | Pointer to **string** | Storage array controller name. | [optional] [readonly] 
**OperationalMode** | Pointer to **string** | Controller running mode, Primary or Secondary. * &#x60;Unknown&#x60; - Component operational state is unknown. * &#x60;Primary&#x60; - Component operates in primary mode and accepts workloads. * &#x60;Secondary&#x60; - Component is running as a secondary or standby mode. * &#x60;Maintenance&#x60; - Component is in maintenance mode for upgrade. During maintenance mode, component does not perform any workload. | [optional] [readonly] [default to "Unknown"]
**Status** | Pointer to **string** | Status of the storage controller. * &#x60;Unknown&#x60; - Component status is not available. * &#x60;Ok&#x60; - Component is healthy and no issues found. * &#x60;Degraded&#x60; - Functioning, but not at full capability due to a non-fatal failure. * &#x60;Critical&#x60; - Not functioning or requiring immediate attention. * &#x60;Offline&#x60; - Component is installed, but powered off. * &#x60;Identifying&#x60; - Component is in initialization process. * &#x60;NotAvailable&#x60; - Component is not installed or not available. * &#x60;Updating&#x60; - Software update is in progress. * &#x60;Unrecognized&#x60; - Component is not recognized. It may be because the component is not installed properly or it is not supported. | [optional] [readonly] [default to "Unknown"]
**Version** | Pointer to **string** | Software version running on a storage controller. | [optional] [readonly] 

## Methods

### NewStorageBaseArrayControllerAllOf

`func NewStorageBaseArrayControllerAllOf(classId string, objectType string, ) *StorageBaseArrayControllerAllOf`

NewStorageBaseArrayControllerAllOf instantiates a new StorageBaseArrayControllerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseArrayControllerAllOfWithDefaults

`func NewStorageBaseArrayControllerAllOfWithDefaults() *StorageBaseArrayControllerAllOf`

NewStorageBaseArrayControllerAllOfWithDefaults instantiates a new StorageBaseArrayControllerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBaseArrayControllerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBaseArrayControllerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBaseArrayControllerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBaseArrayControllerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBaseArrayControllerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBaseArrayControllerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *StorageBaseArrayControllerAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseArrayControllerAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseArrayControllerAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseArrayControllerAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperationalMode

`func (o *StorageBaseArrayControllerAllOf) GetOperationalMode() string`

GetOperationalMode returns the OperationalMode field if non-nil, zero value otherwise.

### GetOperationalModeOk

`func (o *StorageBaseArrayControllerAllOf) GetOperationalModeOk() (*string, bool)`

GetOperationalModeOk returns a tuple with the OperationalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalMode

`func (o *StorageBaseArrayControllerAllOf) SetOperationalMode(v string)`

SetOperationalMode sets OperationalMode field to given value.

### HasOperationalMode

`func (o *StorageBaseArrayControllerAllOf) HasOperationalMode() bool`

HasOperationalMode returns a boolean if a field has been set.

### GetStatus

`func (o *StorageBaseArrayControllerAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageBaseArrayControllerAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageBaseArrayControllerAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageBaseArrayControllerAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *StorageBaseArrayControllerAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageBaseArrayControllerAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageBaseArrayControllerAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageBaseArrayControllerAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


