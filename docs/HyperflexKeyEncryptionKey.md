# HyperflexKeyEncryptionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.KeyEncryptionKey"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.KeyEncryptionKey"]
**IsAccountRecovery** | Pointer to **bool** | Account recovery scenario flag determining parameter population from internal MO, with possibility of ignoring of input parameters when this value is true. | [optional] [default to false]
**IsKekSet** | Pointer to **bool** | Indicates whether the value of the &#39;kek&#39; property has been set. | [optional] [readonly] [default to false]
**IsPassphraseSet** | Pointer to **bool** | Indicates whether the value of the &#39;passphrase&#39; property has been set. | [optional] [readonly] [default to false]
**Iteration** | Pointer to **int64** | Number of hash iterations to run. | [optional] 
**Kek** | Pointer to **string** | Key Encryption Key used to encrypt the DEK&#39;s on the HyperFlex cluster. | [optional] 
**KeyId** | Pointer to **string** | Resource ID and time for Kek retrieval. | [optional] 
**KeyState** | Pointer to **string** | Last known state of the Key Encryption Key. * &#x60;NEW&#x60; - Newly created Key Encryption Key (KEK). * &#x60;ACTIVE&#x60; - Deployed Key Encryption Key on active resources. * &#x60;INACTIVE&#x60; - Inactive and unused Key Encryption Key. * &#x60;INPROGRESS&#x60; - Unconfirmed Key Encryption Key usage on Intersight platform. | [optional] [default to "NEW"]
**Passphrase** | Pointer to **string** | Initial passphrase for encryption policy, requiring a minimum of 12 characters, including 1 lowercase, 1 uppercase, and 1 numeric character. | [optional] 
**ResourceType** | Pointer to **string** | Resource type for key application. * &#x60;CLUSTER&#x60; - Cluster specific encryption per HyperFlex cluster. * &#x60;DATASTORE&#x60; - Data store encryption on the HyperFlex cluster. * &#x60;DRIVE&#x60; - Drive specific encryption on the HyperFlex cluster. | [optional] [default to "CLUSTER"]
**TransitKek** | Pointer to **string** | Temporary copy of KEK used for transfer to remote device endpoint, not persisted anywhere. | [optional] 
**ClusterProfile** | Pointer to [**NullableHyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) |  | [optional] 
**ResourceMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewHyperflexKeyEncryptionKey

`func NewHyperflexKeyEncryptionKey(classId string, objectType string, ) *HyperflexKeyEncryptionKey`

NewHyperflexKeyEncryptionKey instantiates a new HyperflexKeyEncryptionKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexKeyEncryptionKeyWithDefaults

`func NewHyperflexKeyEncryptionKeyWithDefaults() *HyperflexKeyEncryptionKey`

NewHyperflexKeyEncryptionKeyWithDefaults instantiates a new HyperflexKeyEncryptionKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexKeyEncryptionKey) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexKeyEncryptionKey) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexKeyEncryptionKey) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexKeyEncryptionKey) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexKeyEncryptionKey) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexKeyEncryptionKey) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsAccountRecovery

`func (o *HyperflexKeyEncryptionKey) GetIsAccountRecovery() bool`

GetIsAccountRecovery returns the IsAccountRecovery field if non-nil, zero value otherwise.

### GetIsAccountRecoveryOk

`func (o *HyperflexKeyEncryptionKey) GetIsAccountRecoveryOk() (*bool, bool)`

GetIsAccountRecoveryOk returns a tuple with the IsAccountRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountRecovery

`func (o *HyperflexKeyEncryptionKey) SetIsAccountRecovery(v bool)`

SetIsAccountRecovery sets IsAccountRecovery field to given value.

### HasIsAccountRecovery

`func (o *HyperflexKeyEncryptionKey) HasIsAccountRecovery() bool`

HasIsAccountRecovery returns a boolean if a field has been set.

### GetIsKekSet

`func (o *HyperflexKeyEncryptionKey) GetIsKekSet() bool`

GetIsKekSet returns the IsKekSet field if non-nil, zero value otherwise.

### GetIsKekSetOk

`func (o *HyperflexKeyEncryptionKey) GetIsKekSetOk() (*bool, bool)`

GetIsKekSetOk returns a tuple with the IsKekSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKekSet

`func (o *HyperflexKeyEncryptionKey) SetIsKekSet(v bool)`

SetIsKekSet sets IsKekSet field to given value.

### HasIsKekSet

`func (o *HyperflexKeyEncryptionKey) HasIsKekSet() bool`

HasIsKekSet returns a boolean if a field has been set.

### GetIsPassphraseSet

`func (o *HyperflexKeyEncryptionKey) GetIsPassphraseSet() bool`

GetIsPassphraseSet returns the IsPassphraseSet field if non-nil, zero value otherwise.

### GetIsPassphraseSetOk

`func (o *HyperflexKeyEncryptionKey) GetIsPassphraseSetOk() (*bool, bool)`

GetIsPassphraseSetOk returns a tuple with the IsPassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassphraseSet

`func (o *HyperflexKeyEncryptionKey) SetIsPassphraseSet(v bool)`

SetIsPassphraseSet sets IsPassphraseSet field to given value.

### HasIsPassphraseSet

`func (o *HyperflexKeyEncryptionKey) HasIsPassphraseSet() bool`

HasIsPassphraseSet returns a boolean if a field has been set.

### GetIteration

`func (o *HyperflexKeyEncryptionKey) GetIteration() int64`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *HyperflexKeyEncryptionKey) GetIterationOk() (*int64, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *HyperflexKeyEncryptionKey) SetIteration(v int64)`

SetIteration sets Iteration field to given value.

### HasIteration

`func (o *HyperflexKeyEncryptionKey) HasIteration() bool`

HasIteration returns a boolean if a field has been set.

### GetKek

`func (o *HyperflexKeyEncryptionKey) GetKek() string`

GetKek returns the Kek field if non-nil, zero value otherwise.

### GetKekOk

`func (o *HyperflexKeyEncryptionKey) GetKekOk() (*string, bool)`

GetKekOk returns a tuple with the Kek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKek

`func (o *HyperflexKeyEncryptionKey) SetKek(v string)`

SetKek sets Kek field to given value.

### HasKek

`func (o *HyperflexKeyEncryptionKey) HasKek() bool`

HasKek returns a boolean if a field has been set.

### GetKeyId

`func (o *HyperflexKeyEncryptionKey) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *HyperflexKeyEncryptionKey) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *HyperflexKeyEncryptionKey) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *HyperflexKeyEncryptionKey) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetKeyState

`func (o *HyperflexKeyEncryptionKey) GetKeyState() string`

GetKeyState returns the KeyState field if non-nil, zero value otherwise.

### GetKeyStateOk

`func (o *HyperflexKeyEncryptionKey) GetKeyStateOk() (*string, bool)`

GetKeyStateOk returns a tuple with the KeyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyState

`func (o *HyperflexKeyEncryptionKey) SetKeyState(v string)`

SetKeyState sets KeyState field to given value.

### HasKeyState

`func (o *HyperflexKeyEncryptionKey) HasKeyState() bool`

HasKeyState returns a boolean if a field has been set.

### GetPassphrase

`func (o *HyperflexKeyEncryptionKey) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *HyperflexKeyEncryptionKey) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *HyperflexKeyEncryptionKey) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *HyperflexKeyEncryptionKey) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetResourceType

`func (o *HyperflexKeyEncryptionKey) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *HyperflexKeyEncryptionKey) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *HyperflexKeyEncryptionKey) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *HyperflexKeyEncryptionKey) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTransitKek

`func (o *HyperflexKeyEncryptionKey) GetTransitKek() string`

GetTransitKek returns the TransitKek field if non-nil, zero value otherwise.

### GetTransitKekOk

`func (o *HyperflexKeyEncryptionKey) GetTransitKekOk() (*string, bool)`

GetTransitKekOk returns a tuple with the TransitKek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitKek

`func (o *HyperflexKeyEncryptionKey) SetTransitKek(v string)`

SetTransitKek sets TransitKek field to given value.

### HasTransitKek

`func (o *HyperflexKeyEncryptionKey) HasTransitKek() bool`

HasTransitKek returns a boolean if a field has been set.

### GetClusterProfile

`func (o *HyperflexKeyEncryptionKey) GetClusterProfile() HyperflexClusterProfileRelationship`

GetClusterProfile returns the ClusterProfile field if non-nil, zero value otherwise.

### GetClusterProfileOk

`func (o *HyperflexKeyEncryptionKey) GetClusterProfileOk() (*HyperflexClusterProfileRelationship, bool)`

GetClusterProfileOk returns a tuple with the ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfile

`func (o *HyperflexKeyEncryptionKey) SetClusterProfile(v HyperflexClusterProfileRelationship)`

SetClusterProfile sets ClusterProfile field to given value.

### HasClusterProfile

`func (o *HyperflexKeyEncryptionKey) HasClusterProfile() bool`

HasClusterProfile returns a boolean if a field has been set.

### SetClusterProfileNil

`func (o *HyperflexKeyEncryptionKey) SetClusterProfileNil(b bool)`

 SetClusterProfileNil sets the value for ClusterProfile to be an explicit nil

### UnsetClusterProfile
`func (o *HyperflexKeyEncryptionKey) UnsetClusterProfile()`

UnsetClusterProfile ensures that no value is present for ClusterProfile, not even an explicit nil
### GetResourceMo

`func (o *HyperflexKeyEncryptionKey) GetResourceMo() MoBaseMoRelationship`

GetResourceMo returns the ResourceMo field if non-nil, zero value otherwise.

### GetResourceMoOk

`func (o *HyperflexKeyEncryptionKey) GetResourceMoOk() (*MoBaseMoRelationship, bool)`

GetResourceMoOk returns a tuple with the ResourceMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMo

`func (o *HyperflexKeyEncryptionKey) SetResourceMo(v MoBaseMoRelationship)`

SetResourceMo sets ResourceMo field to given value.

### HasResourceMo

`func (o *HyperflexKeyEncryptionKey) HasResourceMo() bool`

HasResourceMo returns a boolean if a field has been set.

### SetResourceMoNil

`func (o *HyperflexKeyEncryptionKey) SetResourceMoNil(b bool)`

 SetResourceMoNil sets the value for ResourceMo to be an explicit nil

### UnsetResourceMo
`func (o *HyperflexKeyEncryptionKey) UnsetResourceMo()`

UnsetResourceMo ensures that no value is present for ResourceMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


