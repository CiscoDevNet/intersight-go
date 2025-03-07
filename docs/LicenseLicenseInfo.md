# LicenseLicenseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.LicenseInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.LicenseInfo"]
**ActiveAdmin** | Pointer to **bool** | The license administrative state. Set this property to &#39;true&#39; to activate the license entitlements. | [optional] [readonly] 
**Balance** | Pointer to **int64** | The total balance we have for licenses. | [optional] [readonly] 
**DaysLeft** | Pointer to **int64** | The number of days left for licenseState to stay in TrialPeriod or OutOfCompliance state. | [optional] [readonly] 
**EndTime** | Pointer to **time.Time** | The date and time when the trial period expires. The value of the &#39;endTime&#39; property is set when the account enters the TrialPeriod or OutOfCompliance state. | [optional] [readonly] 
**EnforceMode** | Pointer to **string** | The entitlement mode reported by Cisco Smart Software Manager. | [optional] [readonly] 
**ErrorDesc** | Pointer to **string** | The detailed error message when there is any error related to this licensing entitlement. | [optional] [readonly] 
**EvaluationPeriod** | Pointer to **int64** | The default Trial or Grace period customer is entitled to. | [optional] 
**ExpireTime** | Pointer to **time.Time** | The date and time when the next expiration time of license subscription. | [optional] [readonly] 
**ExtraEvaluation** | Pointer to **int64** | The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. | [optional] 
**LicenseCount** | Pointer to **int64** | The total number of license consumed in the Intersight account. | [optional] [readonly] 
**LicenseCountPurchased** | Pointer to **int64** | The total number of license purchased from cisco. | [optional] [readonly] 
**LicenseState** | Pointer to **string** | The license state defined by Intersight. The value may be one of NotLicensed, TrialPeriod, OutOfCompliance, Compliance, GraceExpired, or TrialExpired. * &#x60;NotLicensed&#x60; - The license token is neither activated nor registered. * &#x60;GraceExpired&#x60; - The license grace period has expired. * &#x60;TrialPeriod&#x60; - The 90 days of trial period. * &#x60;OutOfCompliance&#x60; - The license is out of compliance. * &#x60;Compliance&#x60; - The license is in compliance. * &#x60;TrialExpired&#x60; - The trial period of 90 days has expired. | [optional] [readonly] [default to "NotLicensed"]
**LicenseType** | Pointer to **string** | The name of the Intersight license entitlement. For example, this property may be set to &#39;Essential&#39;. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. * &#x60;IKS-Advantage&#x60; - IKS-Advantage as a License type. * &#x60;INC-Premier-1GFixed&#x60; - Premier 1G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-10GFixed&#x60; - Premier 10G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-100GFixed&#x60; - Premier 100G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-Mod4Slot&#x60; - Premier Modular 4 slot license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-Mod8Slot&#x60; - Premier Modular 8 slot license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-D2OpsFixed&#x60; - Premier D2Ops fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-D2OpsMod&#x60; - Premier D2Ops modular license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-CentralizedMod8Slot&#x60; - Premier modular license tier of switch type CentralizedMod8Slot for Intersight Nexus Cloud. * &#x60;INC-Premier-DistributedMod8Slot&#x60; - Premier modular license tier of switch type DistributedMod8Slot for Intersight Nexus Cloud. * &#x60;ERP-Advantage&#x60; - Advantage license tier for ERP workflows. * &#x60;IntersightTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers. * &#x60;IWOTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers. * &#x60;IKSTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers. * &#x60;INCTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. | [optional] [readonly] [default to "Base"]
**NetSubstitution** | Pointer to **int64** | The total number of substituted licenses added or removed. | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | The date and time when the licenseState entered the TrialPeriod or OutOfCompliance state. | [optional] [readonly] 
**SubscriptionId** | Pointer to **string** | The id of license subscription. | [optional] [readonly] 
**SubstitutedLicense** | Pointer to [**[]LicenseSubstituteLicense**](LicenseSubstituteLicense.md) |  | [optional] 
**TrialAdmin** | Pointer to **bool** | The administrative state of the trial license. When the LicenseState is set to &#39;NotLicensed&#39;, &#39;trialAdmin&#39; can be set to true to start the trial period, i.e. licenseState is set to be TrialPeriod. | [optional] [readonly] 
**AccountLicenseData** | Pointer to [**NullableLicenseAccountLicenseDataRelationship**](LicenseAccountLicenseDataRelationship.md) |  | [optional] 

## Methods

### NewLicenseLicenseInfo

`func NewLicenseLicenseInfo(classId string, objectType string, ) *LicenseLicenseInfo`

NewLicenseLicenseInfo instantiates a new LicenseLicenseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseLicenseInfoWithDefaults

`func NewLicenseLicenseInfoWithDefaults() *LicenseLicenseInfo`

NewLicenseLicenseInfoWithDefaults instantiates a new LicenseLicenseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseLicenseInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseLicenseInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseLicenseInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseLicenseInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseLicenseInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseLicenseInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveAdmin

`func (o *LicenseLicenseInfo) GetActiveAdmin() bool`

GetActiveAdmin returns the ActiveAdmin field if non-nil, zero value otherwise.

### GetActiveAdminOk

`func (o *LicenseLicenseInfo) GetActiveAdminOk() (*bool, bool)`

GetActiveAdminOk returns a tuple with the ActiveAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAdmin

`func (o *LicenseLicenseInfo) SetActiveAdmin(v bool)`

SetActiveAdmin sets ActiveAdmin field to given value.

### HasActiveAdmin

`func (o *LicenseLicenseInfo) HasActiveAdmin() bool`

HasActiveAdmin returns a boolean if a field has been set.

### GetBalance

`func (o *LicenseLicenseInfo) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *LicenseLicenseInfo) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *LicenseLicenseInfo) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *LicenseLicenseInfo) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetDaysLeft

`func (o *LicenseLicenseInfo) GetDaysLeft() int64`

GetDaysLeft returns the DaysLeft field if non-nil, zero value otherwise.

### GetDaysLeftOk

`func (o *LicenseLicenseInfo) GetDaysLeftOk() (*int64, bool)`

GetDaysLeftOk returns a tuple with the DaysLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLeft

`func (o *LicenseLicenseInfo) SetDaysLeft(v int64)`

SetDaysLeft sets DaysLeft field to given value.

### HasDaysLeft

`func (o *LicenseLicenseInfo) HasDaysLeft() bool`

HasDaysLeft returns a boolean if a field has been set.

### GetEndTime

`func (o *LicenseLicenseInfo) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *LicenseLicenseInfo) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *LicenseLicenseInfo) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *LicenseLicenseInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEnforceMode

`func (o *LicenseLicenseInfo) GetEnforceMode() string`

GetEnforceMode returns the EnforceMode field if non-nil, zero value otherwise.

### GetEnforceModeOk

`func (o *LicenseLicenseInfo) GetEnforceModeOk() (*string, bool)`

GetEnforceModeOk returns a tuple with the EnforceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceMode

`func (o *LicenseLicenseInfo) SetEnforceMode(v string)`

SetEnforceMode sets EnforceMode field to given value.

### HasEnforceMode

`func (o *LicenseLicenseInfo) HasEnforceMode() bool`

HasEnforceMode returns a boolean if a field has been set.

### GetErrorDesc

`func (o *LicenseLicenseInfo) GetErrorDesc() string`

GetErrorDesc returns the ErrorDesc field if non-nil, zero value otherwise.

### GetErrorDescOk

`func (o *LicenseLicenseInfo) GetErrorDescOk() (*string, bool)`

GetErrorDescOk returns a tuple with the ErrorDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDesc

`func (o *LicenseLicenseInfo) SetErrorDesc(v string)`

SetErrorDesc sets ErrorDesc field to given value.

### HasErrorDesc

`func (o *LicenseLicenseInfo) HasErrorDesc() bool`

HasErrorDesc returns a boolean if a field has been set.

### GetEvaluationPeriod

`func (o *LicenseLicenseInfo) GetEvaluationPeriod() int64`

GetEvaluationPeriod returns the EvaluationPeriod field if non-nil, zero value otherwise.

### GetEvaluationPeriodOk

`func (o *LicenseLicenseInfo) GetEvaluationPeriodOk() (*int64, bool)`

GetEvaluationPeriodOk returns a tuple with the EvaluationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPeriod

`func (o *LicenseLicenseInfo) SetEvaluationPeriod(v int64)`

SetEvaluationPeriod sets EvaluationPeriod field to given value.

### HasEvaluationPeriod

`func (o *LicenseLicenseInfo) HasEvaluationPeriod() bool`

HasEvaluationPeriod returns a boolean if a field has been set.

### GetExpireTime

`func (o *LicenseLicenseInfo) GetExpireTime() time.Time`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *LicenseLicenseInfo) GetExpireTimeOk() (*time.Time, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *LicenseLicenseInfo) SetExpireTime(v time.Time)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *LicenseLicenseInfo) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetExtraEvaluation

`func (o *LicenseLicenseInfo) GetExtraEvaluation() int64`

GetExtraEvaluation returns the ExtraEvaluation field if non-nil, zero value otherwise.

### GetExtraEvaluationOk

`func (o *LicenseLicenseInfo) GetExtraEvaluationOk() (*int64, bool)`

GetExtraEvaluationOk returns a tuple with the ExtraEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraEvaluation

`func (o *LicenseLicenseInfo) SetExtraEvaluation(v int64)`

SetExtraEvaluation sets ExtraEvaluation field to given value.

### HasExtraEvaluation

`func (o *LicenseLicenseInfo) HasExtraEvaluation() bool`

HasExtraEvaluation returns a boolean if a field has been set.

### GetLicenseCount

`func (o *LicenseLicenseInfo) GetLicenseCount() int64`

GetLicenseCount returns the LicenseCount field if non-nil, zero value otherwise.

### GetLicenseCountOk

`func (o *LicenseLicenseInfo) GetLicenseCountOk() (*int64, bool)`

GetLicenseCountOk returns a tuple with the LicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCount

`func (o *LicenseLicenseInfo) SetLicenseCount(v int64)`

SetLicenseCount sets LicenseCount field to given value.

### HasLicenseCount

`func (o *LicenseLicenseInfo) HasLicenseCount() bool`

HasLicenseCount returns a boolean if a field has been set.

### GetLicenseCountPurchased

`func (o *LicenseLicenseInfo) GetLicenseCountPurchased() int64`

GetLicenseCountPurchased returns the LicenseCountPurchased field if non-nil, zero value otherwise.

### GetLicenseCountPurchasedOk

`func (o *LicenseLicenseInfo) GetLicenseCountPurchasedOk() (*int64, bool)`

GetLicenseCountPurchasedOk returns a tuple with the LicenseCountPurchased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCountPurchased

`func (o *LicenseLicenseInfo) SetLicenseCountPurchased(v int64)`

SetLicenseCountPurchased sets LicenseCountPurchased field to given value.

### HasLicenseCountPurchased

`func (o *LicenseLicenseInfo) HasLicenseCountPurchased() bool`

HasLicenseCountPurchased returns a boolean if a field has been set.

### GetLicenseState

`func (o *LicenseLicenseInfo) GetLicenseState() string`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *LicenseLicenseInfo) GetLicenseStateOk() (*string, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *LicenseLicenseInfo) SetLicenseState(v string)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *LicenseLicenseInfo) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetLicenseType

`func (o *LicenseLicenseInfo) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *LicenseLicenseInfo) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *LicenseLicenseInfo) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *LicenseLicenseInfo) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetNetSubstitution

`func (o *LicenseLicenseInfo) GetNetSubstitution() int64`

GetNetSubstitution returns the NetSubstitution field if non-nil, zero value otherwise.

### GetNetSubstitutionOk

`func (o *LicenseLicenseInfo) GetNetSubstitutionOk() (*int64, bool)`

GetNetSubstitutionOk returns a tuple with the NetSubstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetSubstitution

`func (o *LicenseLicenseInfo) SetNetSubstitution(v int64)`

SetNetSubstitution sets NetSubstitution field to given value.

### HasNetSubstitution

`func (o *LicenseLicenseInfo) HasNetSubstitution() bool`

HasNetSubstitution returns a boolean if a field has been set.

### GetStartTime

`func (o *LicenseLicenseInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *LicenseLicenseInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *LicenseLicenseInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *LicenseLicenseInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *LicenseLicenseInfo) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *LicenseLicenseInfo) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *LicenseLicenseInfo) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *LicenseLicenseInfo) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubstitutedLicense

`func (o *LicenseLicenseInfo) GetSubstitutedLicense() []LicenseSubstituteLicense`

GetSubstitutedLicense returns the SubstitutedLicense field if non-nil, zero value otherwise.

### GetSubstitutedLicenseOk

`func (o *LicenseLicenseInfo) GetSubstitutedLicenseOk() (*[]LicenseSubstituteLicense, bool)`

GetSubstitutedLicenseOk returns a tuple with the SubstitutedLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutedLicense

`func (o *LicenseLicenseInfo) SetSubstitutedLicense(v []LicenseSubstituteLicense)`

SetSubstitutedLicense sets SubstitutedLicense field to given value.

### HasSubstitutedLicense

`func (o *LicenseLicenseInfo) HasSubstitutedLicense() bool`

HasSubstitutedLicense returns a boolean if a field has been set.

### SetSubstitutedLicenseNil

`func (o *LicenseLicenseInfo) SetSubstitutedLicenseNil(b bool)`

 SetSubstitutedLicenseNil sets the value for SubstitutedLicense to be an explicit nil

### UnsetSubstitutedLicense
`func (o *LicenseLicenseInfo) UnsetSubstitutedLicense()`

UnsetSubstitutedLicense ensures that no value is present for SubstitutedLicense, not even an explicit nil
### GetTrialAdmin

`func (o *LicenseLicenseInfo) GetTrialAdmin() bool`

GetTrialAdmin returns the TrialAdmin field if non-nil, zero value otherwise.

### GetTrialAdminOk

`func (o *LicenseLicenseInfo) GetTrialAdminOk() (*bool, bool)`

GetTrialAdminOk returns a tuple with the TrialAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialAdmin

`func (o *LicenseLicenseInfo) SetTrialAdmin(v bool)`

SetTrialAdmin sets TrialAdmin field to given value.

### HasTrialAdmin

`func (o *LicenseLicenseInfo) HasTrialAdmin() bool`

HasTrialAdmin returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseLicenseInfo) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseLicenseInfo) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseLicenseInfo) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseLicenseInfo) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.

### SetAccountLicenseDataNil

`func (o *LicenseLicenseInfo) SetAccountLicenseDataNil(b bool)`

 SetAccountLicenseDataNil sets the value for AccountLicenseData to be an explicit nil

### UnsetAccountLicenseData
`func (o *LicenseLicenseInfo) UnsetAccountLicenseData()`

UnsetAccountLicenseData ensures that no value is present for AccountLicenseData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


