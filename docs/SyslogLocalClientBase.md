# SyslogLocalClientBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "syslog.LocalFileLoggingClient"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "syslog.LocalFileLoggingClient"]
**MinSeverity** | Pointer to **string** | Lowest level of messages to be included in the local log. * &#x60;warning&#x60; - Use logging level warning for logs classified as warning. * &#x60;emergency&#x60; - Use logging level emergency for logs classified as emergency. * &#x60;alert&#x60; - Use logging level alert for logs classified as alert. * &#x60;critical&#x60; - Use logging level critical for logs classified as critical. * &#x60;error&#x60; - Use logging level error for logs classified as error. * &#x60;notice&#x60; - Use logging level notice for logs classified as notice. * &#x60;informational&#x60; - Use logging level informational for logs classified as informational. * &#x60;debug&#x60; - Use logging level debug for logs classified as debug. | [optional] [default to "warning"]

## Methods

### NewSyslogLocalClientBase

`func NewSyslogLocalClientBase(classId string, objectType string, ) *SyslogLocalClientBase`

NewSyslogLocalClientBase instantiates a new SyslogLocalClientBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogLocalClientBaseWithDefaults

`func NewSyslogLocalClientBaseWithDefaults() *SyslogLocalClientBase`

NewSyslogLocalClientBaseWithDefaults instantiates a new SyslogLocalClientBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SyslogLocalClientBase) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SyslogLocalClientBase) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SyslogLocalClientBase) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SyslogLocalClientBase) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SyslogLocalClientBase) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SyslogLocalClientBase) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMinSeverity

`func (o *SyslogLocalClientBase) GetMinSeverity() string`

GetMinSeverity returns the MinSeverity field if non-nil, zero value otherwise.

### GetMinSeverityOk

`func (o *SyslogLocalClientBase) GetMinSeverityOk() (*string, bool)`

GetMinSeverityOk returns a tuple with the MinSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeverity

`func (o *SyslogLocalClientBase) SetMinSeverity(v string)`

SetMinSeverity sets MinSeverity field to given value.

### HasMinSeverity

`func (o *SyslogLocalClientBase) HasMinSeverity() bool`

HasMinSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


