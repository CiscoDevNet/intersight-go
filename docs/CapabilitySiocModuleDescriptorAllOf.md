# CapabilitySiocModuleDescriptorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SiocModuleDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SiocModuleDescriptor"]
**Revision** | Pointer to **string** | Revision for the SIOC module. | [optional] 

## Methods

### NewCapabilitySiocModuleDescriptorAllOf

`func NewCapabilitySiocModuleDescriptorAllOf(classId string, objectType string, ) *CapabilitySiocModuleDescriptorAllOf`

NewCapabilitySiocModuleDescriptorAllOf instantiates a new CapabilitySiocModuleDescriptorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySiocModuleDescriptorAllOfWithDefaults

`func NewCapabilitySiocModuleDescriptorAllOfWithDefaults() *CapabilitySiocModuleDescriptorAllOf`

NewCapabilitySiocModuleDescriptorAllOfWithDefaults instantiates a new CapabilitySiocModuleDescriptorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySiocModuleDescriptorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySiocModuleDescriptorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySiocModuleDescriptorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySiocModuleDescriptorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySiocModuleDescriptorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySiocModuleDescriptorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRevision

`func (o *CapabilitySiocModuleDescriptorAllOf) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CapabilitySiocModuleDescriptorAllOf) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CapabilitySiocModuleDescriptorAllOf) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *CapabilitySiocModuleDescriptorAllOf) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


