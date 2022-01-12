# CloudCustomAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.CustomAttributes"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.CustomAttributes"]
**AttributeName** | Pointer to **string** | The name of an attribute. If used as a key-value pair then this field represents the key. | [optional] 
**AttributeType** | Pointer to **string** | The data type for attributeValue. For e.g. string, int, float. | [optional] 
**AttributeValue** | Pointer to **string** | The attribute value. If used as a key-value pair then this field represents the value. | [optional] 

## Methods

### NewCloudCustomAttributesAllOf

`func NewCloudCustomAttributesAllOf(classId string, objectType string, ) *CloudCustomAttributesAllOf`

NewCloudCustomAttributesAllOf instantiates a new CloudCustomAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCustomAttributesAllOfWithDefaults

`func NewCloudCustomAttributesAllOfWithDefaults() *CloudCustomAttributesAllOf`

NewCloudCustomAttributesAllOfWithDefaults instantiates a new CloudCustomAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudCustomAttributesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudCustomAttributesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudCustomAttributesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudCustomAttributesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudCustomAttributesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudCustomAttributesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttributeName

`func (o *CloudCustomAttributesAllOf) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *CloudCustomAttributesAllOf) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *CloudCustomAttributesAllOf) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *CloudCustomAttributesAllOf) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetAttributeType

`func (o *CloudCustomAttributesAllOf) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *CloudCustomAttributesAllOf) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *CloudCustomAttributesAllOf) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.

### HasAttributeType

`func (o *CloudCustomAttributesAllOf) HasAttributeType() bool`

HasAttributeType returns a boolean if a field has been set.

### GetAttributeValue

`func (o *CloudCustomAttributesAllOf) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *CloudCustomAttributesAllOf) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *CloudCustomAttributesAllOf) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *CloudCustomAttributesAllOf) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


