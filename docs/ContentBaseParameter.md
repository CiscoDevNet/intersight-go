# ContentBaseParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AcceptSingleValue** | Pointer to **bool** | The flag that allows single values in content to be extracted as a single element collection in case the parameter is of Collection type. This flag is applicable for parameters of type Collection only. | [optional] 
**ComplexType** | Pointer to **string** | The name of the complex type definition in case this is a complex parameter. The content.Grammar object must have a complex type, content.ComplexType, defined with the specified name in types collection property. | [optional] 
**ItemType** | Pointer to **string** | The type of the collection item in case this is a collection parameter. * &#x60;simple&#x60; - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum. * &#x60;string&#x60; - The parameter value to be extracted is of the string type. * &#x60;integer&#x60; - The parameter value to be extracted is of the number type. * &#x60;float&#x60; - The parameter value to be extracted is of the float number type. * &#x60;boolean&#x60; - The parameter value to be extracted is of the boolean type. * &#x60;json&#x60; - The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type. * &#x60;complex&#x60; - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters. * &#x60;collection&#x60; - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. | [optional] [default to "simple"]
**Name** | Pointer to **string** | The name of the parameter. | [optional] 
**Path** | Pointer to **string** | The content specific path information that identifies the parameter value within the content. The value is usually a XPath or JSONPath or a regular expression in case of text content. | [optional] 
**Secure** | Pointer to **bool** | The flag indicates if the extracted value is secure. This flag is applicable for parameters of type String only. | [optional] 
**Type** | Pointer to **string** | The type of the parameter. Accepted values are simple, complex, collection. * &#x60;simple&#x60; - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum. * &#x60;string&#x60; - The parameter value to be extracted is of the string type. * &#x60;integer&#x60; - The parameter value to be extracted is of the number type. * &#x60;float&#x60; - The parameter value to be extracted is of the float number type. * &#x60;boolean&#x60; - The parameter value to be extracted is of the boolean type. * &#x60;json&#x60; - The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type. * &#x60;complex&#x60; - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters. * &#x60;collection&#x60; - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. | [optional] [default to "simple"]

## Methods

### NewContentBaseParameter

`func NewContentBaseParameter(classId string, objectType string, ) *ContentBaseParameter`

NewContentBaseParameter instantiates a new ContentBaseParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentBaseParameterWithDefaults

`func NewContentBaseParameterWithDefaults() *ContentBaseParameter`

NewContentBaseParameterWithDefaults instantiates a new ContentBaseParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ContentBaseParameter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ContentBaseParameter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ContentBaseParameter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ContentBaseParameter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ContentBaseParameter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ContentBaseParameter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAcceptSingleValue

`func (o *ContentBaseParameter) GetAcceptSingleValue() bool`

GetAcceptSingleValue returns the AcceptSingleValue field if non-nil, zero value otherwise.

### GetAcceptSingleValueOk

`func (o *ContentBaseParameter) GetAcceptSingleValueOk() (*bool, bool)`

GetAcceptSingleValueOk returns a tuple with the AcceptSingleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptSingleValue

`func (o *ContentBaseParameter) SetAcceptSingleValue(v bool)`

SetAcceptSingleValue sets AcceptSingleValue field to given value.

### HasAcceptSingleValue

`func (o *ContentBaseParameter) HasAcceptSingleValue() bool`

HasAcceptSingleValue returns a boolean if a field has been set.

### GetComplexType

`func (o *ContentBaseParameter) GetComplexType() string`

GetComplexType returns the ComplexType field if non-nil, zero value otherwise.

### GetComplexTypeOk

`func (o *ContentBaseParameter) GetComplexTypeOk() (*string, bool)`

GetComplexTypeOk returns a tuple with the ComplexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexType

`func (o *ContentBaseParameter) SetComplexType(v string)`

SetComplexType sets ComplexType field to given value.

### HasComplexType

`func (o *ContentBaseParameter) HasComplexType() bool`

HasComplexType returns a boolean if a field has been set.

### GetItemType

`func (o *ContentBaseParameter) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ContentBaseParameter) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ContentBaseParameter) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *ContentBaseParameter) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetName

`func (o *ContentBaseParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentBaseParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentBaseParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentBaseParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *ContentBaseParameter) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ContentBaseParameter) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ContentBaseParameter) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ContentBaseParameter) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSecure

`func (o *ContentBaseParameter) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *ContentBaseParameter) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *ContentBaseParameter) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *ContentBaseParameter) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetType

`func (o *ContentBaseParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentBaseParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentBaseParameter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentBaseParameter) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


