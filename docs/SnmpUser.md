# SnmpUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "snmp.User"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "snmp.User"]
**AuthPassword** | Pointer to **string** | Authorization password for the user. | [optional] 
**AuthType** | Pointer to **string** | Authorization protocol for authenticating the user. * &#x60;NA&#x60; - Authentication protocol is not applicable. * &#x60;MD5&#x60; - MD5 protocol is used to authenticate SNMP user. * &#x60;SHA&#x60; - SHA protocol is used to authenticate SNMP user. * &#x60;SHA-224&#x60; - SHA-224 protocol is used to authenticate SNMP user. * &#x60;SHA-256&#x60; - SHA-256 protocol is used to authenticate SNMP user. * &#x60;SHA-384&#x60; - SHA-384 protocol is used to authenticate SNMP user. * &#x60;SHA-512&#x60; - SHA-512 protocol is used to authenticate SNMP user. | [optional] [default to "NA"]
**IsAuthPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;authPassword&#39; property has been set. | [optional] [readonly] [default to false]
**IsPrivacyPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;privacyPassword&#39; property has been set. | [optional] [readonly] [default to false]
**Name** | Pointer to **string** | SNMP username. Must have a minimum of 1 and and a maximum of 31 characters. | [optional] 
**PrivacyPassword** | Pointer to **string** | Privacy password for the user. | [optional] 
**PrivacyType** | Pointer to **string** | Privacy protocol for the user. * &#x60;NA&#x60; - Privacy protocol is not applicable. * &#x60;DES&#x60; - DES privacy protocol is used for SNMP user. * &#x60;AES&#x60; - AES privacy protocol is used for SNMP user. | [optional] [default to "NA"]
**SecurityLevel** | Pointer to **string** | Security mechanism used for communication between agent and manager. * &#x60;AuthPriv&#x60; - The user requires both an authorization password and a privacy password. * &#x60;NoAuthNoPriv&#x60; - The user does not require an authorization or privacy password. * &#x60;AuthNoPriv&#x60; - The user requires an authorization password but not a privacy password. | [optional] [default to "AuthPriv"]

## Methods

### NewSnmpUser

`func NewSnmpUser(classId string, objectType string, ) *SnmpUser`

NewSnmpUser instantiates a new SnmpUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpUserWithDefaults

`func NewSnmpUserWithDefaults() *SnmpUser`

NewSnmpUserWithDefaults instantiates a new SnmpUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SnmpUser) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SnmpUser) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SnmpUser) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SnmpUser) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SnmpUser) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SnmpUser) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthPassword

`func (o *SnmpUser) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *SnmpUser) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *SnmpUser) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *SnmpUser) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthType

`func (o *SnmpUser) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *SnmpUser) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *SnmpUser) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *SnmpUser) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetIsAuthPasswordSet

`func (o *SnmpUser) GetIsAuthPasswordSet() bool`

GetIsAuthPasswordSet returns the IsAuthPasswordSet field if non-nil, zero value otherwise.

### GetIsAuthPasswordSetOk

`func (o *SnmpUser) GetIsAuthPasswordSetOk() (*bool, bool)`

GetIsAuthPasswordSetOk returns a tuple with the IsAuthPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthPasswordSet

`func (o *SnmpUser) SetIsAuthPasswordSet(v bool)`

SetIsAuthPasswordSet sets IsAuthPasswordSet field to given value.

### HasIsAuthPasswordSet

`func (o *SnmpUser) HasIsAuthPasswordSet() bool`

HasIsAuthPasswordSet returns a boolean if a field has been set.

### GetIsPrivacyPasswordSet

`func (o *SnmpUser) GetIsPrivacyPasswordSet() bool`

GetIsPrivacyPasswordSet returns the IsPrivacyPasswordSet field if non-nil, zero value otherwise.

### GetIsPrivacyPasswordSetOk

`func (o *SnmpUser) GetIsPrivacyPasswordSetOk() (*bool, bool)`

GetIsPrivacyPasswordSetOk returns a tuple with the IsPrivacyPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivacyPasswordSet

`func (o *SnmpUser) SetIsPrivacyPasswordSet(v bool)`

SetIsPrivacyPasswordSet sets IsPrivacyPasswordSet field to given value.

### HasIsPrivacyPasswordSet

`func (o *SnmpUser) HasIsPrivacyPasswordSet() bool`

HasIsPrivacyPasswordSet returns a boolean if a field has been set.

### GetName

`func (o *SnmpUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnmpUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnmpUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnmpUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivacyPassword

`func (o *SnmpUser) GetPrivacyPassword() string`

GetPrivacyPassword returns the PrivacyPassword field if non-nil, zero value otherwise.

### GetPrivacyPasswordOk

`func (o *SnmpUser) GetPrivacyPasswordOk() (*string, bool)`

GetPrivacyPasswordOk returns a tuple with the PrivacyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPassword

`func (o *SnmpUser) SetPrivacyPassword(v string)`

SetPrivacyPassword sets PrivacyPassword field to given value.

### HasPrivacyPassword

`func (o *SnmpUser) HasPrivacyPassword() bool`

HasPrivacyPassword returns a boolean if a field has been set.

### GetPrivacyType

`func (o *SnmpUser) GetPrivacyType() string`

GetPrivacyType returns the PrivacyType field if non-nil, zero value otherwise.

### GetPrivacyTypeOk

`func (o *SnmpUser) GetPrivacyTypeOk() (*string, bool)`

GetPrivacyTypeOk returns a tuple with the PrivacyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyType

`func (o *SnmpUser) SetPrivacyType(v string)`

SetPrivacyType sets PrivacyType field to given value.

### HasPrivacyType

`func (o *SnmpUser) HasPrivacyType() bool`

HasPrivacyType returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *SnmpUser) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *SnmpUser) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *SnmpUser) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *SnmpUser) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


