# IamLdapGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LdapGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LdapGroup"]
**Domain** | Pointer to **string** | LDAP server domain the Group resides in. | [optional] 
**GroupDn** | Pointer to **string** | LDAP Group DN in the LDAP server database. | [optional] 
**Name** | Pointer to **string** | LDAP Group name in the LDAP server database. | [optional] 
**EndPointRole** | Pointer to [**[]IamEndPointRoleRelationship**](IamEndPointRoleRelationship.md) | An array of relationships to iamEndPointRole resources. | [optional] 
**LdapPolicy** | Pointer to [**NullableIamLdapPolicyRelationship**](IamLdapPolicyRelationship.md) |  | [optional] 

## Methods

### NewIamLdapGroup

`func NewIamLdapGroup(classId string, objectType string, ) *IamLdapGroup`

NewIamLdapGroup instantiates a new IamLdapGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapGroupWithDefaults

`func NewIamLdapGroupWithDefaults() *IamLdapGroup`

NewIamLdapGroupWithDefaults instantiates a new IamLdapGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLdapGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLdapGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLdapGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLdapGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLdapGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLdapGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDomain

`func (o *IamLdapGroup) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IamLdapGroup) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IamLdapGroup) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IamLdapGroup) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetGroupDn

`func (o *IamLdapGroup) GetGroupDn() string`

GetGroupDn returns the GroupDn field if non-nil, zero value otherwise.

### GetGroupDnOk

`func (o *IamLdapGroup) GetGroupDnOk() (*string, bool)`

GetGroupDnOk returns a tuple with the GroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDn

`func (o *IamLdapGroup) SetGroupDn(v string)`

SetGroupDn sets GroupDn field to given value.

### HasGroupDn

`func (o *IamLdapGroup) HasGroupDn() bool`

HasGroupDn returns a boolean if a field has been set.

### GetName

`func (o *IamLdapGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamLdapGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamLdapGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamLdapGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEndPointRole

`func (o *IamLdapGroup) GetEndPointRole() []IamEndPointRoleRelationship`

GetEndPointRole returns the EndPointRole field if non-nil, zero value otherwise.

### GetEndPointRoleOk

`func (o *IamLdapGroup) GetEndPointRoleOk() (*[]IamEndPointRoleRelationship, bool)`

GetEndPointRoleOk returns a tuple with the EndPointRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointRole

`func (o *IamLdapGroup) SetEndPointRole(v []IamEndPointRoleRelationship)`

SetEndPointRole sets EndPointRole field to given value.

### HasEndPointRole

`func (o *IamLdapGroup) HasEndPointRole() bool`

HasEndPointRole returns a boolean if a field has been set.

### SetEndPointRoleNil

`func (o *IamLdapGroup) SetEndPointRoleNil(b bool)`

 SetEndPointRoleNil sets the value for EndPointRole to be an explicit nil

### UnsetEndPointRole
`func (o *IamLdapGroup) UnsetEndPointRole()`

UnsetEndPointRole ensures that no value is present for EndPointRole, not even an explicit nil
### GetLdapPolicy

`func (o *IamLdapGroup) GetLdapPolicy() IamLdapPolicyRelationship`

GetLdapPolicy returns the LdapPolicy field if non-nil, zero value otherwise.

### GetLdapPolicyOk

`func (o *IamLdapGroup) GetLdapPolicyOk() (*IamLdapPolicyRelationship, bool)`

GetLdapPolicyOk returns a tuple with the LdapPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPolicy

`func (o *IamLdapGroup) SetLdapPolicy(v IamLdapPolicyRelationship)`

SetLdapPolicy sets LdapPolicy field to given value.

### HasLdapPolicy

`func (o *IamLdapGroup) HasLdapPolicy() bool`

HasLdapPolicy returns a boolean if a field has been set.

### SetLdapPolicyNil

`func (o *IamLdapGroup) SetLdapPolicyNil(b bool)`

 SetLdapPolicyNil sets the value for LdapPolicy to be an explicit nil

### UnsetLdapPolicy
`func (o *IamLdapGroup) UnsetLdapPolicy()`

UnsetLdapPolicy ensures that no value is present for LdapPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


