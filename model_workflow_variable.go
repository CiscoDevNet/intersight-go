/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2024101709
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the WorkflowVariable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowVariable{}

// WorkflowVariable Variables are user-defined entities that can be shared across workflows. They allow users to set a value once and then reference it from different workflows within the same scope. The variables can be of any type that is supported by the workflow system.
type WorkflowVariable struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The user identifier who created the environment variable.
	CreateUser *string                      `json:"CreateUser,omitempty"`
	Definition NullableWorkflowBaseDataType `json:"Definition,omitempty"`
	// The user identifier who last updated the environment variable.
	ModUser *string `json:"ModUser,omitempty"`
	// This defines the name of the variable and it is set by the system. The name used inside the datatype definition will be set as the name of the variable.
	Name *string `json:"Name,omitempty"`
	// This defines the value of the variable and it will be validated against the datatype defined in the definition.
	Value                interface{}                         `json:"Value,omitempty"`
	Catalog              NullableWorkflowCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowVariable WorkflowVariable

// NewWorkflowVariable instantiates a new WorkflowVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowVariable(classId string, objectType string) *WorkflowVariable {
	this := WorkflowVariable{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowVariableWithDefaults instantiates a new WorkflowVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowVariableWithDefaults() *WorkflowVariable {
	this := WorkflowVariable{}
	var classId string = "workflow.Variable"
	this.ClassId = classId
	var objectType string = "workflow.Variable"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowVariable) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowVariable) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowVariable) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.Variable" of the ClassId field.
func (o *WorkflowVariable) GetDefaultClassId() interface{} {
	return "workflow.Variable"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowVariable) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowVariable) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowVariable) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.Variable" of the ObjectType field.
func (o *WorkflowVariable) GetDefaultObjectType() interface{} {
	return "workflow.Variable"
}

// GetCreateUser returns the CreateUser field value if set, zero value otherwise.
func (o *WorkflowVariable) GetCreateUser() string {
	if o == nil || IsNil(o.CreateUser) {
		var ret string
		return ret
	}
	return *o.CreateUser
}

// GetCreateUserOk returns a tuple with the CreateUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariable) GetCreateUserOk() (*string, bool) {
	if o == nil || IsNil(o.CreateUser) {
		return nil, false
	}
	return o.CreateUser, true
}

// HasCreateUser returns a boolean if a field has been set.
func (o *WorkflowVariable) HasCreateUser() bool {
	if o != nil && !IsNil(o.CreateUser) {
		return true
	}

	return false
}

// SetCreateUser gets a reference to the given string and assigns it to the CreateUser field.
func (o *WorkflowVariable) SetCreateUser(v string) {
	o.CreateUser = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariable) GetDefinition() WorkflowBaseDataType {
	if o == nil || IsNil(o.Definition.Get()) {
		var ret WorkflowBaseDataType
		return ret
	}
	return *o.Definition.Get()
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariable) GetDefinitionOk() (*WorkflowBaseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Definition.Get(), o.Definition.IsSet()
}

// HasDefinition returns a boolean if a field has been set.
func (o *WorkflowVariable) HasDefinition() bool {
	if o != nil && o.Definition.IsSet() {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given NullableWorkflowBaseDataType and assigns it to the Definition field.
func (o *WorkflowVariable) SetDefinition(v WorkflowBaseDataType) {
	o.Definition.Set(&v)
}

// SetDefinitionNil sets the value for Definition to be an explicit nil
func (o *WorkflowVariable) SetDefinitionNil() {
	o.Definition.Set(nil)
}

// UnsetDefinition ensures that no value is present for Definition, not even an explicit nil
func (o *WorkflowVariable) UnsetDefinition() {
	o.Definition.Unset()
}

// GetModUser returns the ModUser field value if set, zero value otherwise.
func (o *WorkflowVariable) GetModUser() string {
	if o == nil || IsNil(o.ModUser) {
		var ret string
		return ret
	}
	return *o.ModUser
}

// GetModUserOk returns a tuple with the ModUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariable) GetModUserOk() (*string, bool) {
	if o == nil || IsNil(o.ModUser) {
		return nil, false
	}
	return o.ModUser, true
}

// HasModUser returns a boolean if a field has been set.
func (o *WorkflowVariable) HasModUser() bool {
	if o != nil && !IsNil(o.ModUser) {
		return true
	}

	return false
}

// SetModUser gets a reference to the given string and assigns it to the ModUser field.
func (o *WorkflowVariable) SetModUser(v string) {
	o.ModUser = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowVariable) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariable) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowVariable) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowVariable) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariable) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariable) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkflowVariable) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *WorkflowVariable) SetValue(v interface{}) {
	o.Value = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariable) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || IsNil(o.Catalog.Get()) {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog.Get()
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariable) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Catalog.Get(), o.Catalog.IsSet()
}

// HasCatalog returns a boolean if a field has been set.
func (o *WorkflowVariable) HasCatalog() bool {
	if o != nil && o.Catalog.IsSet() {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given NullableWorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *WorkflowVariable) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog.Set(&v)
}

// SetCatalogNil sets the value for Catalog to be an explicit nil
func (o *WorkflowVariable) SetCatalogNil() {
	o.Catalog.Set(nil)
}

// UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
func (o *WorkflowVariable) UnsetCatalog() {
	o.Catalog.Unset()
}

func (o WorkflowVariable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowVariable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.CreateUser) {
		toSerialize["CreateUser"] = o.CreateUser
	}
	if o.Definition.IsSet() {
		toSerialize["Definition"] = o.Definition.Get()
	}
	if !IsNil(o.ModUser) {
		toSerialize["ModUser"] = o.ModUser
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	if o.Catalog.IsSet() {
		toSerialize["Catalog"] = o.Catalog.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowVariable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type WorkflowVariableWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The user identifier who created the environment variable.
		CreateUser *string                      `json:"CreateUser,omitempty"`
		Definition NullableWorkflowBaseDataType `json:"Definition,omitempty"`
		// The user identifier who last updated the environment variable.
		ModUser *string `json:"ModUser,omitempty"`
		// This defines the name of the variable and it is set by the system. The name used inside the datatype definition will be set as the name of the variable.
		Name *string `json:"Name,omitempty"`
		// This defines the value of the variable and it will be validated against the datatype defined in the definition.
		Value   interface{}                         `json:"Value,omitempty"`
		Catalog NullableWorkflowCatalogRelationship `json:"Catalog,omitempty"`
	}

	varWorkflowVariableWithoutEmbeddedStruct := WorkflowVariableWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowVariableWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowVariable := _WorkflowVariable{}
		varWorkflowVariable.ClassId = varWorkflowVariableWithoutEmbeddedStruct.ClassId
		varWorkflowVariable.ObjectType = varWorkflowVariableWithoutEmbeddedStruct.ObjectType
		varWorkflowVariable.CreateUser = varWorkflowVariableWithoutEmbeddedStruct.CreateUser
		varWorkflowVariable.Definition = varWorkflowVariableWithoutEmbeddedStruct.Definition
		varWorkflowVariable.ModUser = varWorkflowVariableWithoutEmbeddedStruct.ModUser
		varWorkflowVariable.Name = varWorkflowVariableWithoutEmbeddedStruct.Name
		varWorkflowVariable.Value = varWorkflowVariableWithoutEmbeddedStruct.Value
		varWorkflowVariable.Catalog = varWorkflowVariableWithoutEmbeddedStruct.Catalog
		*o = WorkflowVariable(varWorkflowVariable)
	} else {
		return err
	}

	varWorkflowVariable := _WorkflowVariable{}

	err = json.Unmarshal(data, &varWorkflowVariable)
	if err == nil {
		o.MoBaseMo = varWorkflowVariable.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CreateUser")
		delete(additionalProperties, "Definition")
		delete(additionalProperties, "ModUser")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Value")
		delete(additionalProperties, "Catalog")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowVariable struct {
	value *WorkflowVariable
	isSet bool
}

func (v NullableWorkflowVariable) Get() *WorkflowVariable {
	return v.value
}

func (v *NullableWorkflowVariable) Set(val *WorkflowVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowVariable(val *WorkflowVariable) *NullableWorkflowVariable {
	return &NullableWorkflowVariable{value: val, isSet: true}
}

func (v NullableWorkflowVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}