# RecommendationHardwareExpansionRequestList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;recommendation.HardwareExpansionRequest&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]RecommendationHardwareExpansionRequest**](RecommendationHardwareExpansionRequest.md) | The array of &#39;recommendation.HardwareExpansionRequest&#39; resources matching the request. | [optional] 

## Methods

### NewRecommendationHardwareExpansionRequestList

`func NewRecommendationHardwareExpansionRequestList() *RecommendationHardwareExpansionRequestList`

NewRecommendationHardwareExpansionRequestList instantiates a new RecommendationHardwareExpansionRequestList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationHardwareExpansionRequestListWithDefaults

`func NewRecommendationHardwareExpansionRequestListWithDefaults() *RecommendationHardwareExpansionRequestList`

NewRecommendationHardwareExpansionRequestListWithDefaults instantiates a new RecommendationHardwareExpansionRequestList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RecommendationHardwareExpansionRequestList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RecommendationHardwareExpansionRequestList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RecommendationHardwareExpansionRequestList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *RecommendationHardwareExpansionRequestList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *RecommendationHardwareExpansionRequestList) GetResults() []RecommendationHardwareExpansionRequest`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RecommendationHardwareExpansionRequestList) GetResultsOk() (*[]RecommendationHardwareExpansionRequest, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RecommendationHardwareExpansionRequestList) SetResults(v []RecommendationHardwareExpansionRequest)`

SetResults sets Results field to given value.

### HasResults

`func (o *RecommendationHardwareExpansionRequestList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *RecommendationHardwareExpansionRequestList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *RecommendationHardwareExpansionRequestList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


