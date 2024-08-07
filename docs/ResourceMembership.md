# ResourceMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.Membership"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.Membership"]
**GroupPermissionRoles** | Pointer to [**[]IamGroupPermissionToRoles**](IamGroupPermissionToRoles.md) |  | [optional] 
**Reevaluate** | Pointer to **bool** | Set Reevaluate to true to reevaluate the membership of a resource. | [optional] 
**TargetApp** | Pointer to **string** | Name of the Service owning the resource. | [optional] [readonly] 
**Holder** | Pointer to [**NullableResourceMembershipHolderRelationship**](ResourceMembershipHolderRelationship.md) |  | [optional] 
**Resource** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**ResourceAncestors** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] 

## Methods

### NewResourceMembership

`func NewResourceMembership(classId string, objectType string, ) *ResourceMembership`

NewResourceMembership instantiates a new ResourceMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMembershipWithDefaults

`func NewResourceMembershipWithDefaults() *ResourceMembership`

NewResourceMembershipWithDefaults instantiates a new ResourceMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceMembership) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceMembership) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceMembership) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceMembership) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceMembership) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceMembership) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGroupPermissionRoles

`func (o *ResourceMembership) GetGroupPermissionRoles() []IamGroupPermissionToRoles`

GetGroupPermissionRoles returns the GroupPermissionRoles field if non-nil, zero value otherwise.

### GetGroupPermissionRolesOk

`func (o *ResourceMembership) GetGroupPermissionRolesOk() (*[]IamGroupPermissionToRoles, bool)`

GetGroupPermissionRolesOk returns a tuple with the GroupPermissionRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPermissionRoles

`func (o *ResourceMembership) SetGroupPermissionRoles(v []IamGroupPermissionToRoles)`

SetGroupPermissionRoles sets GroupPermissionRoles field to given value.

### HasGroupPermissionRoles

`func (o *ResourceMembership) HasGroupPermissionRoles() bool`

HasGroupPermissionRoles returns a boolean if a field has been set.

### SetGroupPermissionRolesNil

`func (o *ResourceMembership) SetGroupPermissionRolesNil(b bool)`

 SetGroupPermissionRolesNil sets the value for GroupPermissionRoles to be an explicit nil

### UnsetGroupPermissionRoles
`func (o *ResourceMembership) UnsetGroupPermissionRoles()`

UnsetGroupPermissionRoles ensures that no value is present for GroupPermissionRoles, not even an explicit nil
### GetReevaluate

`func (o *ResourceMembership) GetReevaluate() bool`

GetReevaluate returns the Reevaluate field if non-nil, zero value otherwise.

### GetReevaluateOk

`func (o *ResourceMembership) GetReevaluateOk() (*bool, bool)`

GetReevaluateOk returns a tuple with the Reevaluate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReevaluate

`func (o *ResourceMembership) SetReevaluate(v bool)`

SetReevaluate sets Reevaluate field to given value.

### HasReevaluate

`func (o *ResourceMembership) HasReevaluate() bool`

HasReevaluate returns a boolean if a field has been set.

### GetTargetApp

`func (o *ResourceMembership) GetTargetApp() string`

GetTargetApp returns the TargetApp field if non-nil, zero value otherwise.

### GetTargetAppOk

`func (o *ResourceMembership) GetTargetAppOk() (*string, bool)`

GetTargetAppOk returns a tuple with the TargetApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApp

`func (o *ResourceMembership) SetTargetApp(v string)`

SetTargetApp sets TargetApp field to given value.

### HasTargetApp

`func (o *ResourceMembership) HasTargetApp() bool`

HasTargetApp returns a boolean if a field has been set.

### GetHolder

`func (o *ResourceMembership) GetHolder() ResourceMembershipHolderRelationship`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *ResourceMembership) GetHolderOk() (*ResourceMembershipHolderRelationship, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *ResourceMembership) SetHolder(v ResourceMembershipHolderRelationship)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *ResourceMembership) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### SetHolderNil

`func (o *ResourceMembership) SetHolderNil(b bool)`

 SetHolderNil sets the value for Holder to be an explicit nil

### UnsetHolder
`func (o *ResourceMembership) UnsetHolder()`

UnsetHolder ensures that no value is present for Holder, not even an explicit nil
### GetResource

`func (o *ResourceMembership) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceMembership) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceMembership) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourceMembership) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *ResourceMembership) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *ResourceMembership) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetResourceAncestors

`func (o *ResourceMembership) GetResourceAncestors() []MoBaseMoRelationship`

GetResourceAncestors returns the ResourceAncestors field if non-nil, zero value otherwise.

### GetResourceAncestorsOk

`func (o *ResourceMembership) GetResourceAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetResourceAncestorsOk returns a tuple with the ResourceAncestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAncestors

`func (o *ResourceMembership) SetResourceAncestors(v []MoBaseMoRelationship)`

SetResourceAncestors sets ResourceAncestors field to given value.

### HasResourceAncestors

`func (o *ResourceMembership) HasResourceAncestors() bool`

HasResourceAncestors returns a boolean if a field has been set.

### SetResourceAncestorsNil

`func (o *ResourceMembership) SetResourceAncestorsNil(b bool)`

 SetResourceAncestorsNil sets the value for ResourceAncestors to be an explicit nil

### UnsetResourceAncestors
`func (o *ResourceMembership) UnsetResourceAncestors()`

UnsetResourceAncestors ensures that no value is present for ResourceAncestors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


