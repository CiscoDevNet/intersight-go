# TelemetryDruidAndFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The filter type. | 
**ExtractionFn** | Pointer to **map[string]interface{}** | All filters except the \&quot;spatial\&quot; filter support extraction functions. An extraction function is defined by setting the \&quot;extractionFn\&quot; field on a filter. See Extraction function for more details on extraction functions. If specified, the extraction function will be used to transform input values before the filter is applied. The example below shows a selector filter combined with an extraction function. This filter will transform input values according to the values defined in the lookup map; transformed values will then be matched with the string \&quot;bar_1\&quot;. | [optional] 
**Fields** | [**[]TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | 

## Methods

### NewTelemetryDruidAndFilter

`func NewTelemetryDruidAndFilter(type_ string, fields []TelemetryDruidFilter, ) *TelemetryDruidAndFilter`

NewTelemetryDruidAndFilter instantiates a new TelemetryDruidAndFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidAndFilterWithDefaults

`func NewTelemetryDruidAndFilterWithDefaults() *TelemetryDruidAndFilter`

NewTelemetryDruidAndFilterWithDefaults instantiates a new TelemetryDruidAndFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidAndFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidAndFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidAndFilter) SetType(v string)`

SetType sets Type field to given value.


### GetExtractionFn

`func (o *TelemetryDruidAndFilter) GetExtractionFn() map[string]interface{}`

GetExtractionFn returns the ExtractionFn field if non-nil, zero value otherwise.

### GetExtractionFnOk

`func (o *TelemetryDruidAndFilter) GetExtractionFnOk() (*map[string]interface{}, bool)`

GetExtractionFnOk returns a tuple with the ExtractionFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionFn

`func (o *TelemetryDruidAndFilter) SetExtractionFn(v map[string]interface{})`

SetExtractionFn sets ExtractionFn field to given value.

### HasExtractionFn

`func (o *TelemetryDruidAndFilter) HasExtractionFn() bool`

HasExtractionFn returns a boolean if a field has been set.

### GetFields

`func (o *TelemetryDruidAndFilter) GetFields() []TelemetryDruidFilter`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TelemetryDruidAndFilter) GetFieldsOk() (*[]TelemetryDruidFilter, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TelemetryDruidAndFilter) SetFields(v []TelemetryDruidFilter)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


