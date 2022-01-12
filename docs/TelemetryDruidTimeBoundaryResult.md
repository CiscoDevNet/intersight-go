# TelemetryDruidTimeBoundaryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The ISO 8601 timestamp. | [optional] 
**Result** | Pointer to **map[string]interface{}** | The corresponding timestamps for query. May contain maxTime, minTime, or both (default). | [optional] 

## Methods

### NewTelemetryDruidTimeBoundaryResult

`func NewTelemetryDruidTimeBoundaryResult() *TelemetryDruidTimeBoundaryResult`

NewTelemetryDruidTimeBoundaryResult instantiates a new TelemetryDruidTimeBoundaryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidTimeBoundaryResultWithDefaults

`func NewTelemetryDruidTimeBoundaryResultWithDefaults() *TelemetryDruidTimeBoundaryResult`

NewTelemetryDruidTimeBoundaryResultWithDefaults instantiates a new TelemetryDruidTimeBoundaryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *TelemetryDruidTimeBoundaryResult) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TelemetryDruidTimeBoundaryResult) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TelemetryDruidTimeBoundaryResult) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TelemetryDruidTimeBoundaryResult) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetResult

`func (o *TelemetryDruidTimeBoundaryResult) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TelemetryDruidTimeBoundaryResult) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TelemetryDruidTimeBoundaryResult) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *TelemetryDruidTimeBoundaryResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


