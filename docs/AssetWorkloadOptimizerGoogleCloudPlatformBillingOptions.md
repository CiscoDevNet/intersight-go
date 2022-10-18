# AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerGoogleCloudPlatformBillingOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerGoogleCloudPlatformBillingOptions"]
**BillingAccountId** | Pointer to **string** | The Google Cloud Platform (GCP) Billing Account ID. | [optional] 
**PricingExportDataSetName** | Pointer to **string** | Name of the BigQuery Pricing Export Data Set which is the dataset for negotiated pricing. | [optional] 
**PricingExportTableName** | Pointer to **string** | The Google BigQuery Pricing Export Table Name field is auto-populated with the table used in BigQuery, cloud_pricing_export. There is no need to modify this name, unless you use a different table for negotiated pricing. The Default name is \&quot;cloud_pricing_export\&quot;. | [optional] 
**ProjectId** | Pointer to **string** | The unique ID assigned to the project containing the cost and pricing exports. If the exports are in separate projects, multiple billing targets will be necessary. | [optional] 
**ResourceLevelCostDataSetName** | Pointer to **string** | This dataset contains the resource level billed cost export table. | [optional] 
**ResourceLevelCostTableName** | Pointer to **string** | This table stores resource level cost export data. | [optional] 
**StandardCostDataSetName** | Pointer to **string** | Name of Google BigQuery Standard Cost Export Data Set which is the dataset for the standard billed cost export. | [optional] 
**StandardCostTableName** | Pointer to **string** | Google BigQuery Standard Cost Export Table Name. This table will store the exported data from the Standard Cost Export. | [optional] 

## Methods

### NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions

`func NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions(classId string, objectType string, ) *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions`

NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptions instantiates a new AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithDefaults

`func NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithDefaults() *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions`

NewAssetWorkloadOptimizerGoogleCloudPlatformBillingOptionsWithDefaults instantiates a new AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBillingAccountId

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetBillingAccountId() string`

GetBillingAccountId returns the BillingAccountId field if non-nil, zero value otherwise.

### GetBillingAccountIdOk

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetBillingAccountIdOk() (*string, bool)`

GetBillingAccountIdOk returns a tuple with the BillingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountId

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetBillingAccountId(v string)`

SetBillingAccountId sets BillingAccountId field to given value.

### HasBillingAccountId

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasBillingAccountId() bool`

HasBillingAccountId returns a boolean if a field has been set.

### GetPricingExportDataSetName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetPricingExportDataSetName() string`

GetPricingExportDataSetName returns the PricingExportDataSetName field if non-nil, zero value otherwise.

### GetPricingExportDataSetNameOk

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetPricingExportDataSetNameOk() (*string, bool)`

GetPricingExportDataSetNameOk returns a tuple with the PricingExportDataSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingExportDataSetName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetPricingExportDataSetName(v string)`

SetPricingExportDataSetName sets PricingExportDataSetName field to given value.

### HasPricingExportDataSetName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasPricingExportDataSetName() bool`

HasPricingExportDataSetName returns a boolean if a field has been set.

### GetPricingExportTableName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetPricingExportTableName() string`

GetPricingExportTableName returns the PricingExportTableName field if non-nil, zero value otherwise.

### GetPricingExportTableNameOk

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetPricingExportTableNameOk() (*string, bool)`

GetPricingExportTableNameOk returns a tuple with the PricingExportTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingExportTableName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetPricingExportTableName(v string)`

SetPricingExportTableName sets PricingExportTableName field to given value.

### HasPricingExportTableName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasPricingExportTableName() bool`

HasPricingExportTableName returns a boolean if a field has been set.

### GetProjectId

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetResourceLevelCostDataSetName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetResourceLevelCostDataSetName() string`

GetResourceLevelCostDataSetName returns the ResourceLevelCostDataSetName field if non-nil, zero value otherwise.

### GetResourceLevelCostDataSetNameOk

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetResourceLevelCostDataSetNameOk() (*string, bool)`

GetResourceLevelCostDataSetNameOk returns a tuple with the ResourceLevelCostDataSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLevelCostDataSetName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetResourceLevelCostDataSetName(v string)`

SetResourceLevelCostDataSetName sets ResourceLevelCostDataSetName field to given value.

### HasResourceLevelCostDataSetName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasResourceLevelCostDataSetName() bool`

HasResourceLevelCostDataSetName returns a boolean if a field has been set.

### GetResourceLevelCostTableName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetResourceLevelCostTableName() string`

GetResourceLevelCostTableName returns the ResourceLevelCostTableName field if non-nil, zero value otherwise.

### GetResourceLevelCostTableNameOk

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetResourceLevelCostTableNameOk() (*string, bool)`

GetResourceLevelCostTableNameOk returns a tuple with the ResourceLevelCostTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLevelCostTableName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetResourceLevelCostTableName(v string)`

SetResourceLevelCostTableName sets ResourceLevelCostTableName field to given value.

### HasResourceLevelCostTableName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasResourceLevelCostTableName() bool`

HasResourceLevelCostTableName returns a boolean if a field has been set.

### GetStandardCostDataSetName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetStandardCostDataSetName() string`

GetStandardCostDataSetName returns the StandardCostDataSetName field if non-nil, zero value otherwise.

### GetStandardCostDataSetNameOk

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetStandardCostDataSetNameOk() (*string, bool)`

GetStandardCostDataSetNameOk returns a tuple with the StandardCostDataSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCostDataSetName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetStandardCostDataSetName(v string)`

SetStandardCostDataSetName sets StandardCostDataSetName field to given value.

### HasStandardCostDataSetName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasStandardCostDataSetName() bool`

HasStandardCostDataSetName returns a boolean if a field has been set.

### GetStandardCostTableName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetStandardCostTableName() string`

GetStandardCostTableName returns the StandardCostTableName field if non-nil, zero value otherwise.

### GetStandardCostTableNameOk

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) GetStandardCostTableNameOk() (*string, bool)`

GetStandardCostTableNameOk returns a tuple with the StandardCostTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCostTableName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) SetStandardCostTableName(v string)`

SetStandardCostTableName sets StandardCostTableName field to given value.

### HasStandardCostTableName

`func (o *AssetWorkloadOptimizerGoogleCloudPlatformBillingOptions) HasStandardCostTableName() bool`

HasStandardCostTableName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


