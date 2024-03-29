# KubernetesAddonPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.AddonPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.AddonPolicy"]
**AddonConfiguration** | Pointer to [**NullableKubernetesAddonConfiguration**](KubernetesAddonConfiguration.md) |  | [optional] 
**AddonDefinition** | Pointer to [**NullableKubernetesAddonDefinitionRelationship**](KubernetesAddonDefinitionRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesAddonPolicy

`func NewKubernetesAddonPolicy(classId string, objectType string, ) *KubernetesAddonPolicy`

NewKubernetesAddonPolicy instantiates a new KubernetesAddonPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAddonPolicyWithDefaults

`func NewKubernetesAddonPolicyWithDefaults() *KubernetesAddonPolicy`

NewKubernetesAddonPolicyWithDefaults instantiates a new KubernetesAddonPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAddonPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAddonPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAddonPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAddonPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAddonPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAddonPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddonConfiguration

`func (o *KubernetesAddonPolicy) GetAddonConfiguration() KubernetesAddonConfiguration`

GetAddonConfiguration returns the AddonConfiguration field if non-nil, zero value otherwise.

### GetAddonConfigurationOk

`func (o *KubernetesAddonPolicy) GetAddonConfigurationOk() (*KubernetesAddonConfiguration, bool)`

GetAddonConfigurationOk returns a tuple with the AddonConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonConfiguration

`func (o *KubernetesAddonPolicy) SetAddonConfiguration(v KubernetesAddonConfiguration)`

SetAddonConfiguration sets AddonConfiguration field to given value.

### HasAddonConfiguration

`func (o *KubernetesAddonPolicy) HasAddonConfiguration() bool`

HasAddonConfiguration returns a boolean if a field has been set.

### SetAddonConfigurationNil

`func (o *KubernetesAddonPolicy) SetAddonConfigurationNil(b bool)`

 SetAddonConfigurationNil sets the value for AddonConfiguration to be an explicit nil

### UnsetAddonConfiguration
`func (o *KubernetesAddonPolicy) UnsetAddonConfiguration()`

UnsetAddonConfiguration ensures that no value is present for AddonConfiguration, not even an explicit nil
### GetAddonDefinition

`func (o *KubernetesAddonPolicy) GetAddonDefinition() KubernetesAddonDefinitionRelationship`

GetAddonDefinition returns the AddonDefinition field if non-nil, zero value otherwise.

### GetAddonDefinitionOk

`func (o *KubernetesAddonPolicy) GetAddonDefinitionOk() (*KubernetesAddonDefinitionRelationship, bool)`

GetAddonDefinitionOk returns a tuple with the AddonDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonDefinition

`func (o *KubernetesAddonPolicy) SetAddonDefinition(v KubernetesAddonDefinitionRelationship)`

SetAddonDefinition sets AddonDefinition field to given value.

### HasAddonDefinition

`func (o *KubernetesAddonPolicy) HasAddonDefinition() bool`

HasAddonDefinition returns a boolean if a field has been set.

### SetAddonDefinitionNil

`func (o *KubernetesAddonPolicy) SetAddonDefinitionNil(b bool)`

 SetAddonDefinitionNil sets the value for AddonDefinition to be an explicit nil

### UnsetAddonDefinition
`func (o *KubernetesAddonPolicy) UnsetAddonDefinition()`

UnsetAddonDefinition ensures that no value is present for AddonDefinition, not even an explicit nil
### GetOrganization

`func (o *KubernetesAddonPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesAddonPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesAddonPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesAddonPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *KubernetesAddonPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *KubernetesAddonPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


