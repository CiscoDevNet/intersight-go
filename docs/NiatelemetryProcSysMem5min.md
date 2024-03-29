# NiatelemetryProcSysMem5min

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ProcSysMem5min"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ProcSysMem5min"]
**FreeAvg** | Pointer to **string** | Returns the free average value. | [optional] 
**TotalAvg** | Pointer to **string** | Returns the total average value. | [optional] 

## Methods

### NewNiatelemetryProcSysMem5min

`func NewNiatelemetryProcSysMem5min(classId string, objectType string, ) *NiatelemetryProcSysMem5min`

NewNiatelemetryProcSysMem5min instantiates a new NiatelemetryProcSysMem5min object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryProcSysMem5minWithDefaults

`func NewNiatelemetryProcSysMem5minWithDefaults() *NiatelemetryProcSysMem5min`

NewNiatelemetryProcSysMem5minWithDefaults instantiates a new NiatelemetryProcSysMem5min object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryProcSysMem5min) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryProcSysMem5min) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryProcSysMem5min) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryProcSysMem5min) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryProcSysMem5min) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryProcSysMem5min) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFreeAvg

`func (o *NiatelemetryProcSysMem5min) GetFreeAvg() string`

GetFreeAvg returns the FreeAvg field if non-nil, zero value otherwise.

### GetFreeAvgOk

`func (o *NiatelemetryProcSysMem5min) GetFreeAvgOk() (*string, bool)`

GetFreeAvgOk returns a tuple with the FreeAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeAvg

`func (o *NiatelemetryProcSysMem5min) SetFreeAvg(v string)`

SetFreeAvg sets FreeAvg field to given value.

### HasFreeAvg

`func (o *NiatelemetryProcSysMem5min) HasFreeAvg() bool`

HasFreeAvg returns a boolean if a field has been set.

### GetTotalAvg

`func (o *NiatelemetryProcSysMem5min) GetTotalAvg() string`

GetTotalAvg returns the TotalAvg field if non-nil, zero value otherwise.

### GetTotalAvgOk

`func (o *NiatelemetryProcSysMem5min) GetTotalAvgOk() (*string, bool)`

GetTotalAvgOk returns a tuple with the TotalAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAvg

`func (o *NiatelemetryProcSysMem5min) SetTotalAvg(v string)`

SetTotalAvg sets TotalAvg field to given value.

### HasTotalAvg

`func (o *NiatelemetryProcSysMem5min) HasTotalAvg() bool`

HasTotalAvg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


