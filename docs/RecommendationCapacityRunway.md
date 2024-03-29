# RecommendationCapacityRunway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recommendation.CapacityRunway"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recommendation.CapacityRunway"]
**AdditionalCapacity** | Pointer to **int64** | Additional capacity is the capacity which is needed more after exhausting all hardware on current cluster. | [optional] [readonly] 
**Period** | Pointer to **int64** | Number of months in future for which recommendation is provided for. | [optional] [readonly] 
**Runway** | Pointer to **int64** | This represents the new runway, that is the number of days remaining before the cluster&#39;s storage utilization reaches the recommended capacity limit after the recommended hardware is added. | [optional] [readonly] 
**TotalCapacity** | Pointer to **int64** | Total capacity of the cluster after the recommended hardware is added. | [optional] [readonly] 
**Unit** | Pointer to **string** | Unit for the new capacity. * &#x60;TB&#x60; - The Enum value TB represents that the measurement unit is in terabytes. * &#x60;MB&#x60; - The Enum value MB represents that the measurement unit is in megabytes. * &#x60;GB&#x60; - The Enum value GB represents that the measurement unit is in gigabytes. * &#x60;MHz&#x60; - The Enum value MHz represents that the measurement unit is in megahertz. * &#x60;GHz&#x60; - The Enum value GHz represents that the measurement unit is in gigahertz. * &#x60;Percentage&#x60; - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity. | [optional] [readonly] [default to "TB"]
**ForecastInstance** | Pointer to [**NullableForecastInstanceRelationship**](ForecastInstanceRelationship.md) |  | [optional] 
**PhysicalItem** | Pointer to [**[]RecommendationPhysicalItemRelationship**](RecommendationPhysicalItemRelationship.md) | An array of relationships to recommendationPhysicalItem resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewRecommendationCapacityRunway

`func NewRecommendationCapacityRunway(classId string, objectType string, ) *RecommendationCapacityRunway`

NewRecommendationCapacityRunway instantiates a new RecommendationCapacityRunway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationCapacityRunwayWithDefaults

`func NewRecommendationCapacityRunwayWithDefaults() *RecommendationCapacityRunway`

NewRecommendationCapacityRunwayWithDefaults instantiates a new RecommendationCapacityRunway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationCapacityRunway) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationCapacityRunway) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationCapacityRunway) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationCapacityRunway) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationCapacityRunway) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationCapacityRunway) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdditionalCapacity

`func (o *RecommendationCapacityRunway) GetAdditionalCapacity() int64`

GetAdditionalCapacity returns the AdditionalCapacity field if non-nil, zero value otherwise.

### GetAdditionalCapacityOk

`func (o *RecommendationCapacityRunway) GetAdditionalCapacityOk() (*int64, bool)`

GetAdditionalCapacityOk returns a tuple with the AdditionalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCapacity

`func (o *RecommendationCapacityRunway) SetAdditionalCapacity(v int64)`

SetAdditionalCapacity sets AdditionalCapacity field to given value.

### HasAdditionalCapacity

`func (o *RecommendationCapacityRunway) HasAdditionalCapacity() bool`

HasAdditionalCapacity returns a boolean if a field has been set.

### GetPeriod

`func (o *RecommendationCapacityRunway) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *RecommendationCapacityRunway) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *RecommendationCapacityRunway) SetPeriod(v int64)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *RecommendationCapacityRunway) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRunway

`func (o *RecommendationCapacityRunway) GetRunway() int64`

GetRunway returns the Runway field if non-nil, zero value otherwise.

### GetRunwayOk

`func (o *RecommendationCapacityRunway) GetRunwayOk() (*int64, bool)`

GetRunwayOk returns a tuple with the Runway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunway

`func (o *RecommendationCapacityRunway) SetRunway(v int64)`

SetRunway sets Runway field to given value.

### HasRunway

`func (o *RecommendationCapacityRunway) HasRunway() bool`

HasRunway returns a boolean if a field has been set.

### GetTotalCapacity

`func (o *RecommendationCapacityRunway) GetTotalCapacity() int64`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *RecommendationCapacityRunway) GetTotalCapacityOk() (*int64, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *RecommendationCapacityRunway) SetTotalCapacity(v int64)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *RecommendationCapacityRunway) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetUnit

`func (o *RecommendationCapacityRunway) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RecommendationCapacityRunway) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RecommendationCapacityRunway) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RecommendationCapacityRunway) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetForecastInstance

`func (o *RecommendationCapacityRunway) GetForecastInstance() ForecastInstanceRelationship`

GetForecastInstance returns the ForecastInstance field if non-nil, zero value otherwise.

### GetForecastInstanceOk

`func (o *RecommendationCapacityRunway) GetForecastInstanceOk() (*ForecastInstanceRelationship, bool)`

GetForecastInstanceOk returns a tuple with the ForecastInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastInstance

`func (o *RecommendationCapacityRunway) SetForecastInstance(v ForecastInstanceRelationship)`

SetForecastInstance sets ForecastInstance field to given value.

### HasForecastInstance

`func (o *RecommendationCapacityRunway) HasForecastInstance() bool`

HasForecastInstance returns a boolean if a field has been set.

### SetForecastInstanceNil

`func (o *RecommendationCapacityRunway) SetForecastInstanceNil(b bool)`

 SetForecastInstanceNil sets the value for ForecastInstance to be an explicit nil

### UnsetForecastInstance
`func (o *RecommendationCapacityRunway) UnsetForecastInstance()`

UnsetForecastInstance ensures that no value is present for ForecastInstance, not even an explicit nil
### GetPhysicalItem

`func (o *RecommendationCapacityRunway) GetPhysicalItem() []RecommendationPhysicalItemRelationship`

GetPhysicalItem returns the PhysicalItem field if non-nil, zero value otherwise.

### GetPhysicalItemOk

`func (o *RecommendationCapacityRunway) GetPhysicalItemOk() (*[]RecommendationPhysicalItemRelationship, bool)`

GetPhysicalItemOk returns a tuple with the PhysicalItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalItem

`func (o *RecommendationCapacityRunway) SetPhysicalItem(v []RecommendationPhysicalItemRelationship)`

SetPhysicalItem sets PhysicalItem field to given value.

### HasPhysicalItem

`func (o *RecommendationCapacityRunway) HasPhysicalItem() bool`

HasPhysicalItem returns a boolean if a field has been set.

### SetPhysicalItemNil

`func (o *RecommendationCapacityRunway) SetPhysicalItemNil(b bool)`

 SetPhysicalItemNil sets the value for PhysicalItem to be an explicit nil

### UnsetPhysicalItem
`func (o *RecommendationCapacityRunway) UnsetPhysicalItem()`

UnsetPhysicalItem ensures that no value is present for PhysicalItem, not even an explicit nil
### GetRegisteredDevice

`func (o *RecommendationCapacityRunway) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *RecommendationCapacityRunway) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *RecommendationCapacityRunway) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *RecommendationCapacityRunway) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *RecommendationCapacityRunway) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *RecommendationCapacityRunway) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


