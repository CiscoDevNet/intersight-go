# AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerRedHatOpenStackOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerRedHatOpenStackOptions"]
**Domain** | Pointer to **string** | RedHat OpenStack Identity Service (keystone) domain name. Domain is an additional namespaces you can create in keystone to partition users, groups, and projects. Default domain name value is \&quot;Default\&quot;. | [optional] 
**Tenant** | Pointer to **string** | The name of tenant which has assigned administrator role this RedHat OpenStack target user is in. A tenant or project is referred to as a group of users of RedHat OpenStack. Each tenant can be assigned a role which gives all its member users their rights and privileges to perform a specific set of operations. | [optional] 

## Methods

### NewAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf

`func NewAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf(classId string, objectType string, ) *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf`

NewAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf instantiates a new AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerRedHatOpenStackOptionsAllOfWithDefaults

`func NewAssetWorkloadOptimizerRedHatOpenStackOptionsAllOfWithDefaults() *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf`

NewAssetWorkloadOptimizerRedHatOpenStackOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDomain

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTenant

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


