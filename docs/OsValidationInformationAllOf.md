# OsValidationInformationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.ValidationInformation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.ValidationInformation"]
**ErrorMsg** | Pointer to **string** | Validation error message. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the validation step. * &#x60;NotValidated&#x60; - The validation not started. * &#x60;Valid&#x60; - The step status marked as valid when respective step validation condition is successful. * &#x60;Invalid&#x60; - The step status marked as invalid when respective step validation condition is failed. * &#x60;InProgress&#x60; - The validation is in progress. | [optional] [readonly] [default to "NotValidated"]
**StepName** | Pointer to **string** | The validation step name. * &#x60;OS Install Schema Validation&#x60; - The step to validate the CSV file schema. * &#x60;OS Image Validation&#x60; - The Operating System Image parameter validation step. * &#x60;SCU Image Validation&#x60; - The SCU Image parameter validation step. * &#x60;Configuration source and file validation&#x60; - The Configuration Source and Configuration file validation step. * &#x60;Server level data validation&#x60; - The server level parameters validation. | [optional] [readonly] [default to "OS Install Schema Validation"]

## Methods

### NewOsValidationInformationAllOf

`func NewOsValidationInformationAllOf(classId string, objectType string, ) *OsValidationInformationAllOf`

NewOsValidationInformationAllOf instantiates a new OsValidationInformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsValidationInformationAllOfWithDefaults

`func NewOsValidationInformationAllOfWithDefaults() *OsValidationInformationAllOf`

NewOsValidationInformationAllOfWithDefaults instantiates a new OsValidationInformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsValidationInformationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsValidationInformationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsValidationInformationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsValidationInformationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsValidationInformationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsValidationInformationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrorMsg

`func (o *OsValidationInformationAllOf) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *OsValidationInformationAllOf) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *OsValidationInformationAllOf) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *OsValidationInformationAllOf) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### GetStatus

`func (o *OsValidationInformationAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OsValidationInformationAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OsValidationInformationAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OsValidationInformationAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStepName

`func (o *OsValidationInformationAllOf) GetStepName() string`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *OsValidationInformationAllOf) GetStepNameOk() (*string, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *OsValidationInformationAllOf) SetStepName(v string)`

SetStepName sets StepName field to given value.

### HasStepName

`func (o *OsValidationInformationAllOf) HasStepName() bool`

HasStepName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


