# NiatelemetryProcSysMem15min

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ProcSysMem15min"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ProcSysMem15min"]
**FreeAvg** | Pointer to **string** | Returns the free average value. | [optional] 
**TotalAvg** | Pointer to **string** | Returns the total average value. | [optional] 

## Methods

### NewNiatelemetryProcSysMem15min

`func NewNiatelemetryProcSysMem15min(classId string, objectType string, ) *NiatelemetryProcSysMem15min`

NewNiatelemetryProcSysMem15min instantiates a new NiatelemetryProcSysMem15min object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryProcSysMem15minWithDefaults

`func NewNiatelemetryProcSysMem15minWithDefaults() *NiatelemetryProcSysMem15min`

NewNiatelemetryProcSysMem15minWithDefaults instantiates a new NiatelemetryProcSysMem15min object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryProcSysMem15min) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryProcSysMem15min) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryProcSysMem15min) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryProcSysMem15min) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryProcSysMem15min) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryProcSysMem15min) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFreeAvg

`func (o *NiatelemetryProcSysMem15min) GetFreeAvg() string`

GetFreeAvg returns the FreeAvg field if non-nil, zero value otherwise.

### GetFreeAvgOk

`func (o *NiatelemetryProcSysMem15min) GetFreeAvgOk() (*string, bool)`

GetFreeAvgOk returns a tuple with the FreeAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeAvg

`func (o *NiatelemetryProcSysMem15min) SetFreeAvg(v string)`

SetFreeAvg sets FreeAvg field to given value.

### HasFreeAvg

`func (o *NiatelemetryProcSysMem15min) HasFreeAvg() bool`

HasFreeAvg returns a boolean if a field has been set.

### GetTotalAvg

`func (o *NiatelemetryProcSysMem15min) GetTotalAvg() string`

GetTotalAvg returns the TotalAvg field if non-nil, zero value otherwise.

### GetTotalAvgOk

`func (o *NiatelemetryProcSysMem15min) GetTotalAvgOk() (*string, bool)`

GetTotalAvgOk returns a tuple with the TotalAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAvg

`func (o *NiatelemetryProcSysMem15min) SetTotalAvg(v string)`

SetTotalAvg sets TotalAvg field to given value.

### HasTotalAvg

`func (o *NiatelemetryProcSysMem15min) HasTotalAvg() bool`

HasTotalAvg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


