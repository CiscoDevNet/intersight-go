# ThermalPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "thermal.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "thermal.Policy"]
**FanControlMode** | Pointer to **string** | Sets the Fan Control Mode of the System. High Power, Maximum Power and Acoustic modes are only supported for Cisco UCS X series Chassis. * &#x60;Balanced&#x60; - The fans run faster when needed based on the heat generated by the server. When possible, the fans returns to the minimum required speed. * &#x60;LowPower&#x60; - The Fans run at the minimum speed required to keep the server cool. * &#x60;HighPower&#x60; - The fans are kept at higher speed to emphasizes performance over power consumption. This Mode is only supported for UCS X series Chassis. * &#x60;MaximumPower&#x60; - The fans are always kept at maximum speed. This option provides the most cooling and consumes the most power. This Mode is only supported for UCS X series Chassis. * &#x60;Acoustic&#x60; - The fan speed is reduced to reduce noise levels in acoustic-sensitive environments. This Mode is only supported for UCS X series Chassis. | [optional] [default to "Balanced"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewThermalPolicyAllOf

`func NewThermalPolicyAllOf(classId string, objectType string, ) *ThermalPolicyAllOf`

NewThermalPolicyAllOf instantiates a new ThermalPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThermalPolicyAllOfWithDefaults

`func NewThermalPolicyAllOfWithDefaults() *ThermalPolicyAllOf`

NewThermalPolicyAllOfWithDefaults instantiates a new ThermalPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ThermalPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ThermalPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ThermalPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ThermalPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ThermalPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ThermalPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFanControlMode

`func (o *ThermalPolicyAllOf) GetFanControlMode() string`

GetFanControlMode returns the FanControlMode field if non-nil, zero value otherwise.

### GetFanControlModeOk

`func (o *ThermalPolicyAllOf) GetFanControlModeOk() (*string, bool)`

GetFanControlModeOk returns a tuple with the FanControlMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanControlMode

`func (o *ThermalPolicyAllOf) SetFanControlMode(v string)`

SetFanControlMode sets FanControlMode field to given value.

### HasFanControlMode

`func (o *ThermalPolicyAllOf) HasFanControlMode() bool`

HasFanControlMode returns a boolean if a field has been set.

### GetOrganization

`func (o *ThermalPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ThermalPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ThermalPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ThermalPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *ThermalPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *ThermalPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *ThermalPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *ThermalPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *ThermalPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *ThermalPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

