# SmtpPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "smtp.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "smtp.Policy"]
**AuthPassword** | Pointer to **string** | Authorization password for the process. | [optional] 
**EnableAuth** | Pointer to **bool** | If enabled, lets user input username and password. | [optional] [default to false]
**EnableTls** | Pointer to **bool** | If enabled, lets user input valid CA certificates for authorization. | [optional] [default to false]
**Enabled** | Pointer to **bool** | If enabled, controls the state of the SMTP client service on the managed device. | [optional] [default to true]
**IsAuthPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;authPassword&#39; property has been set. | [optional] [readonly] [default to false]
**MinSeverity** | Pointer to **string** | Minimum fault severity level to receive email notifications. Email notifications are sent for all faults whose severity is equal to or greater than the chosen level. * &#x60;critical&#x60; - Minimum severity to report is critical. * &#x60;condition&#x60; - Minimum severity to report is informational. * &#x60;warning&#x60; - Minimum severity to report is warning. * &#x60;minor&#x60; - Minimum severity to report is minor. * &#x60;major&#x60; - Minimum severity to report is major. | [optional] [default to "critical"]
**SenderEmail** | Pointer to **string** | The email address entered here will be displayed as the from address (mail received from address) of all the SMTP mail alerts that are received. If not configured, the hostname of the server is used in the from address field. | [optional] 
**SmtpPort** | Pointer to **int64** | Port number used by the SMTP server for outgoing SMTP communication. | [optional] [default to 25]
**SmtpRecipients** | Pointer to **[]string** |  | [optional] 
**SmtpServer** | Pointer to **string** | IP address or hostname of the SMTP server. The SMTP server is used by the managed device to send email notifications. | [optional] 
**UserName** | Pointer to **string** | SMTP username from which email notification is sent. | [optional] 
**ApplianceAccount** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Certificate** | Pointer to [**NullableIamTrustPointRelationship**](IamTrustPointRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewSmtpPolicy

`func NewSmtpPolicy(classId string, objectType string, ) *SmtpPolicy`

NewSmtpPolicy instantiates a new SmtpPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpPolicyWithDefaults

`func NewSmtpPolicyWithDefaults() *SmtpPolicy`

NewSmtpPolicyWithDefaults instantiates a new SmtpPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SmtpPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SmtpPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SmtpPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SmtpPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SmtpPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SmtpPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthPassword

`func (o *SmtpPolicy) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *SmtpPolicy) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *SmtpPolicy) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *SmtpPolicy) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetEnableAuth

`func (o *SmtpPolicy) GetEnableAuth() bool`

GetEnableAuth returns the EnableAuth field if non-nil, zero value otherwise.

### GetEnableAuthOk

`func (o *SmtpPolicy) GetEnableAuthOk() (*bool, bool)`

GetEnableAuthOk returns a tuple with the EnableAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuth

`func (o *SmtpPolicy) SetEnableAuth(v bool)`

SetEnableAuth sets EnableAuth field to given value.

### HasEnableAuth

`func (o *SmtpPolicy) HasEnableAuth() bool`

HasEnableAuth returns a boolean if a field has been set.

### GetEnableTls

`func (o *SmtpPolicy) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *SmtpPolicy) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *SmtpPolicy) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *SmtpPolicy) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.

### GetEnabled

`func (o *SmtpPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SmtpPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SmtpPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SmtpPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsAuthPasswordSet

`func (o *SmtpPolicy) GetIsAuthPasswordSet() bool`

GetIsAuthPasswordSet returns the IsAuthPasswordSet field if non-nil, zero value otherwise.

### GetIsAuthPasswordSetOk

`func (o *SmtpPolicy) GetIsAuthPasswordSetOk() (*bool, bool)`

GetIsAuthPasswordSetOk returns a tuple with the IsAuthPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthPasswordSet

`func (o *SmtpPolicy) SetIsAuthPasswordSet(v bool)`

SetIsAuthPasswordSet sets IsAuthPasswordSet field to given value.

### HasIsAuthPasswordSet

`func (o *SmtpPolicy) HasIsAuthPasswordSet() bool`

HasIsAuthPasswordSet returns a boolean if a field has been set.

### GetMinSeverity

`func (o *SmtpPolicy) GetMinSeverity() string`

GetMinSeverity returns the MinSeverity field if non-nil, zero value otherwise.

### GetMinSeverityOk

`func (o *SmtpPolicy) GetMinSeverityOk() (*string, bool)`

GetMinSeverityOk returns a tuple with the MinSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeverity

`func (o *SmtpPolicy) SetMinSeverity(v string)`

SetMinSeverity sets MinSeverity field to given value.

### HasMinSeverity

`func (o *SmtpPolicy) HasMinSeverity() bool`

HasMinSeverity returns a boolean if a field has been set.

### GetSenderEmail

`func (o *SmtpPolicy) GetSenderEmail() string`

GetSenderEmail returns the SenderEmail field if non-nil, zero value otherwise.

### GetSenderEmailOk

`func (o *SmtpPolicy) GetSenderEmailOk() (*string, bool)`

GetSenderEmailOk returns a tuple with the SenderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmail

`func (o *SmtpPolicy) SetSenderEmail(v string)`

SetSenderEmail sets SenderEmail field to given value.

### HasSenderEmail

`func (o *SmtpPolicy) HasSenderEmail() bool`

HasSenderEmail returns a boolean if a field has been set.

### GetSmtpPort

`func (o *SmtpPolicy) GetSmtpPort() int64`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *SmtpPolicy) GetSmtpPortOk() (*int64, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *SmtpPolicy) SetSmtpPort(v int64)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *SmtpPolicy) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpRecipients

`func (o *SmtpPolicy) GetSmtpRecipients() []string`

GetSmtpRecipients returns the SmtpRecipients field if non-nil, zero value otherwise.

### GetSmtpRecipientsOk

`func (o *SmtpPolicy) GetSmtpRecipientsOk() (*[]string, bool)`

GetSmtpRecipientsOk returns a tuple with the SmtpRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpRecipients

`func (o *SmtpPolicy) SetSmtpRecipients(v []string)`

SetSmtpRecipients sets SmtpRecipients field to given value.

### HasSmtpRecipients

`func (o *SmtpPolicy) HasSmtpRecipients() bool`

HasSmtpRecipients returns a boolean if a field has been set.

### SetSmtpRecipientsNil

`func (o *SmtpPolicy) SetSmtpRecipientsNil(b bool)`

 SetSmtpRecipientsNil sets the value for SmtpRecipients to be an explicit nil

### UnsetSmtpRecipients
`func (o *SmtpPolicy) UnsetSmtpRecipients()`

UnsetSmtpRecipients ensures that no value is present for SmtpRecipients, not even an explicit nil
### GetSmtpServer

`func (o *SmtpPolicy) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *SmtpPolicy) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *SmtpPolicy) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *SmtpPolicy) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetUserName

`func (o *SmtpPolicy) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SmtpPolicy) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SmtpPolicy) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SmtpPolicy) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetApplianceAccount

`func (o *SmtpPolicy) GetApplianceAccount() IamAccountRelationship`

GetApplianceAccount returns the ApplianceAccount field if non-nil, zero value otherwise.

### GetApplianceAccountOk

`func (o *SmtpPolicy) GetApplianceAccountOk() (*IamAccountRelationship, bool)`

GetApplianceAccountOk returns a tuple with the ApplianceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceAccount

`func (o *SmtpPolicy) SetApplianceAccount(v IamAccountRelationship)`

SetApplianceAccount sets ApplianceAccount field to given value.

### HasApplianceAccount

`func (o *SmtpPolicy) HasApplianceAccount() bool`

HasApplianceAccount returns a boolean if a field has been set.

### SetApplianceAccountNil

`func (o *SmtpPolicy) SetApplianceAccountNil(b bool)`

 SetApplianceAccountNil sets the value for ApplianceAccount to be an explicit nil

### UnsetApplianceAccount
`func (o *SmtpPolicy) UnsetApplianceAccount()`

UnsetApplianceAccount ensures that no value is present for ApplianceAccount, not even an explicit nil
### GetCertificate

`func (o *SmtpPolicy) GetCertificate() IamTrustPointRelationship`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SmtpPolicy) GetCertificateOk() (*IamTrustPointRelationship, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SmtpPolicy) SetCertificate(v IamTrustPointRelationship)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SmtpPolicy) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *SmtpPolicy) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *SmtpPolicy) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetOrganization

`func (o *SmtpPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SmtpPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SmtpPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SmtpPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *SmtpPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *SmtpPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetProfiles

`func (o *SmtpPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SmtpPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SmtpPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *SmtpPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *SmtpPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *SmtpPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


