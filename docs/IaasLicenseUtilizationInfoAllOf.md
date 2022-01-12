# IaasLicenseUtilizationInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.LicenseUtilizationInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.LicenseUtilizationInfo"]
**ActualUsed** | Pointer to **int64** | Number of licenses actually used for this feature. | [optional] [readonly] 
**Label** | Pointer to **string** | License label of UCSD license. | [optional] [readonly] 
**LicensedLimit** | Pointer to **string** | License limit for this license feature. | [optional] [readonly] 
**Sku** | Pointer to **string** | SKU for the UCSD license. | [optional] [readonly] 

## Methods

### NewIaasLicenseUtilizationInfoAllOf

`func NewIaasLicenseUtilizationInfoAllOf(classId string, objectType string, ) *IaasLicenseUtilizationInfoAllOf`

NewIaasLicenseUtilizationInfoAllOf instantiates a new IaasLicenseUtilizationInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasLicenseUtilizationInfoAllOfWithDefaults

`func NewIaasLicenseUtilizationInfoAllOfWithDefaults() *IaasLicenseUtilizationInfoAllOf`

NewIaasLicenseUtilizationInfoAllOfWithDefaults instantiates a new IaasLicenseUtilizationInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasLicenseUtilizationInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasLicenseUtilizationInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasLicenseUtilizationInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasLicenseUtilizationInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasLicenseUtilizationInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasLicenseUtilizationInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActualUsed

`func (o *IaasLicenseUtilizationInfoAllOf) GetActualUsed() int64`

GetActualUsed returns the ActualUsed field if non-nil, zero value otherwise.

### GetActualUsedOk

`func (o *IaasLicenseUtilizationInfoAllOf) GetActualUsedOk() (*int64, bool)`

GetActualUsedOk returns a tuple with the ActualUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualUsed

`func (o *IaasLicenseUtilizationInfoAllOf) SetActualUsed(v int64)`

SetActualUsed sets ActualUsed field to given value.

### HasActualUsed

`func (o *IaasLicenseUtilizationInfoAllOf) HasActualUsed() bool`

HasActualUsed returns a boolean if a field has been set.

### GetLabel

`func (o *IaasLicenseUtilizationInfoAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IaasLicenseUtilizationInfoAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IaasLicenseUtilizationInfoAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *IaasLicenseUtilizationInfoAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLicensedLimit

`func (o *IaasLicenseUtilizationInfoAllOf) GetLicensedLimit() string`

GetLicensedLimit returns the LicensedLimit field if non-nil, zero value otherwise.

### GetLicensedLimitOk

`func (o *IaasLicenseUtilizationInfoAllOf) GetLicensedLimitOk() (*string, bool)`

GetLicensedLimitOk returns a tuple with the LicensedLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedLimit

`func (o *IaasLicenseUtilizationInfoAllOf) SetLicensedLimit(v string)`

SetLicensedLimit sets LicensedLimit field to given value.

### HasLicensedLimit

`func (o *IaasLicenseUtilizationInfoAllOf) HasLicensedLimit() bool`

HasLicensedLimit returns a boolean if a field has been set.

### GetSku

`func (o *IaasLicenseUtilizationInfoAllOf) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *IaasLicenseUtilizationInfoAllOf) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *IaasLicenseUtilizationInfoAllOf) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *IaasLicenseUtilizationInfoAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


