# ChassisConfigChangeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "chassis.ConfigChangeDetail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "chassis.ConfigChangeDetail"]
**IomProfile** | Pointer to [**NullableChassisIomProfileRelationship**](ChassisIomProfileRelationship.md) |  | [optional] 
**Profile** | Pointer to [**NullableChassisProfileRelationship**](ChassisProfileRelationship.md) |  | [optional] 

## Methods

### NewChassisConfigChangeDetail

`func NewChassisConfigChangeDetail(classId string, objectType string, ) *ChassisConfigChangeDetail`

NewChassisConfigChangeDetail instantiates a new ChassisConfigChangeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisConfigChangeDetailWithDefaults

`func NewChassisConfigChangeDetailWithDefaults() *ChassisConfigChangeDetail`

NewChassisConfigChangeDetailWithDefaults instantiates a new ChassisConfigChangeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChassisConfigChangeDetail) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChassisConfigChangeDetail) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChassisConfigChangeDetail) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChassisConfigChangeDetail) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChassisConfigChangeDetail) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChassisConfigChangeDetail) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIomProfile

`func (o *ChassisConfigChangeDetail) GetIomProfile() ChassisIomProfileRelationship`

GetIomProfile returns the IomProfile field if non-nil, zero value otherwise.

### GetIomProfileOk

`func (o *ChassisConfigChangeDetail) GetIomProfileOk() (*ChassisIomProfileRelationship, bool)`

GetIomProfileOk returns a tuple with the IomProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIomProfile

`func (o *ChassisConfigChangeDetail) SetIomProfile(v ChassisIomProfileRelationship)`

SetIomProfile sets IomProfile field to given value.

### HasIomProfile

`func (o *ChassisConfigChangeDetail) HasIomProfile() bool`

HasIomProfile returns a boolean if a field has been set.

### SetIomProfileNil

`func (o *ChassisConfigChangeDetail) SetIomProfileNil(b bool)`

 SetIomProfileNil sets the value for IomProfile to be an explicit nil

### UnsetIomProfile
`func (o *ChassisConfigChangeDetail) UnsetIomProfile()`

UnsetIomProfile ensures that no value is present for IomProfile, not even an explicit nil
### GetProfile

`func (o *ChassisConfigChangeDetail) GetProfile() ChassisProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ChassisConfigChangeDetail) GetProfileOk() (*ChassisProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ChassisConfigChangeDetail) SetProfile(v ChassisProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ChassisConfigChangeDetail) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *ChassisConfigChangeDetail) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *ChassisConfigChangeDetail) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


