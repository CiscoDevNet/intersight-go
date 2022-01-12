# CapabilityPortGroupAggregationDefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.PortGroupAggregationDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.PortGroupAggregationDef"]
**AggregationCap** | Pointer to **string** | Aggregation capability for port group. | [optional] 
**Hw40GPortGroupCap** | Pointer to **bool** | Indicates support for 40G port group capability. | [optional] 
**Pgtype** | Pointer to **string** | The type of port group for which this capability is defined. | [optional] 

## Methods

### NewCapabilityPortGroupAggregationDefAllOf

`func NewCapabilityPortGroupAggregationDefAllOf(classId string, objectType string, ) *CapabilityPortGroupAggregationDefAllOf`

NewCapabilityPortGroupAggregationDefAllOf instantiates a new CapabilityPortGroupAggregationDefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityPortGroupAggregationDefAllOfWithDefaults

`func NewCapabilityPortGroupAggregationDefAllOfWithDefaults() *CapabilityPortGroupAggregationDefAllOf`

NewCapabilityPortGroupAggregationDefAllOfWithDefaults instantiates a new CapabilityPortGroupAggregationDefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityPortGroupAggregationDefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityPortGroupAggregationDefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityPortGroupAggregationDefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityPortGroupAggregationDefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityPortGroupAggregationDefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityPortGroupAggregationDefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregationCap

`func (o *CapabilityPortGroupAggregationDefAllOf) GetAggregationCap() string`

GetAggregationCap returns the AggregationCap field if non-nil, zero value otherwise.

### GetAggregationCapOk

`func (o *CapabilityPortGroupAggregationDefAllOf) GetAggregationCapOk() (*string, bool)`

GetAggregationCapOk returns a tuple with the AggregationCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationCap

`func (o *CapabilityPortGroupAggregationDefAllOf) SetAggregationCap(v string)`

SetAggregationCap sets AggregationCap field to given value.

### HasAggregationCap

`func (o *CapabilityPortGroupAggregationDefAllOf) HasAggregationCap() bool`

HasAggregationCap returns a boolean if a field has been set.

### GetHw40GPortGroupCap

`func (o *CapabilityPortGroupAggregationDefAllOf) GetHw40GPortGroupCap() bool`

GetHw40GPortGroupCap returns the Hw40GPortGroupCap field if non-nil, zero value otherwise.

### GetHw40GPortGroupCapOk

`func (o *CapabilityPortGroupAggregationDefAllOf) GetHw40GPortGroupCapOk() (*bool, bool)`

GetHw40GPortGroupCapOk returns a tuple with the Hw40GPortGroupCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHw40GPortGroupCap

`func (o *CapabilityPortGroupAggregationDefAllOf) SetHw40GPortGroupCap(v bool)`

SetHw40GPortGroupCap sets Hw40GPortGroupCap field to given value.

### HasHw40GPortGroupCap

`func (o *CapabilityPortGroupAggregationDefAllOf) HasHw40GPortGroupCap() bool`

HasHw40GPortGroupCap returns a boolean if a field has been set.

### GetPgtype

`func (o *CapabilityPortGroupAggregationDefAllOf) GetPgtype() string`

GetPgtype returns the Pgtype field if non-nil, zero value otherwise.

### GetPgtypeOk

`func (o *CapabilityPortGroupAggregationDefAllOf) GetPgtypeOk() (*string, bool)`

GetPgtypeOk returns a tuple with the Pgtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgtype

`func (o *CapabilityPortGroupAggregationDefAllOf) SetPgtype(v string)`

SetPgtype sets Pgtype field to given value.

### HasPgtype

`func (o *CapabilityPortGroupAggregationDefAllOf) HasPgtype() bool`

HasPgtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


