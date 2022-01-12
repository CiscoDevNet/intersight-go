# TamPsirtSeverityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.PsirtSeverity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.PsirtSeverity"]
**Level** | Pointer to **string** | Severity level associated with the security advisory. * &#x60;critical&#x60; - &lt; If applicable, it may expose users to critical failures and it needs to be addressed immediately. For e.g. a PSIRT advisory with a corresponding CVSS score of above 9.0. * &#x60;high&#x60; - &lt; If applicable, it may expose the users to critical failure and it needs to be addressed immediately. For e.g. a PSIRT advisory with a corresponding CVSS score between 7.0-8.9. These may be the vulnerabilities that are more difficult to exploit than the ones deemed critical but once exploited, the will cause critical failures. * &#x60;medium&#x60; - &lt; If applicable, it may expose the users to failure of certain functions. for e.g. a PSIRT advisory with a corresponding CVSS score between 4.0-6.9. These may be the vulnerabilities that are mitigated to a large extent but that may still be exploited in certain restricted cases. * &#x60;info&#x60; - &lt; If applicable, it may have some minimal impact for certain functions in the user environment. For e.g. a PSIRT advisory with a corresponding CVSS score below 4.0. | [optional] [default to "critical"]

## Methods

### NewTamPsirtSeverityAllOf

`func NewTamPsirtSeverityAllOf(classId string, objectType string, ) *TamPsirtSeverityAllOf`

NewTamPsirtSeverityAllOf instantiates a new TamPsirtSeverityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamPsirtSeverityAllOfWithDefaults

`func NewTamPsirtSeverityAllOfWithDefaults() *TamPsirtSeverityAllOf`

NewTamPsirtSeverityAllOfWithDefaults instantiates a new TamPsirtSeverityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamPsirtSeverityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamPsirtSeverityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamPsirtSeverityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamPsirtSeverityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamPsirtSeverityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamPsirtSeverityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLevel

`func (o *TamPsirtSeverityAllOf) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *TamPsirtSeverityAllOf) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *TamPsirtSeverityAllOf) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *TamPsirtSeverityAllOf) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


