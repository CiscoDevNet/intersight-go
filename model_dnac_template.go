/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2025062323
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

// checks if the DnacTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnacTemplate{}

// DnacTemplate Collection of information of templates.
type DnacTemplate struct {
	DnacInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Value indicates the template is composite.
	Composite *string `json:"Composite,omitempty"`
	// Identity of the project template.
	ProjectId *string `json:"ProjectId,omitempty"`
	// Name of the project template.
	ProjectName *string `json:"ProjectName,omitempty"`
	// Unique identity of the template.
	TemplateId *string `json:"TemplateId,omitempty"`
	// Name assigned to the template.
	TemplateName         *string `json:"TemplateName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DnacTemplate DnacTemplate

// NewDnacTemplate instantiates a new DnacTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnacTemplate(classId string, objectType string) *DnacTemplate {
	this := DnacTemplate{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewDnacTemplateWithDefaults instantiates a new DnacTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnacTemplateWithDefaults() *DnacTemplate {
	this := DnacTemplate{}
	var classId string = "dnac.Template"
	this.ClassId = classId
	var objectType string = "dnac.Template"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *DnacTemplate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *DnacTemplate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *DnacTemplate) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "dnac.Template" of the ClassId field.
func (o *DnacTemplate) GetDefaultClassId() interface{} {
	return "dnac.Template"
}

// GetObjectType returns the ObjectType field value
func (o *DnacTemplate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *DnacTemplate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *DnacTemplate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "dnac.Template" of the ObjectType field.
func (o *DnacTemplate) GetDefaultObjectType() interface{} {
	return "dnac.Template"
}

// GetComposite returns the Composite field value if set, zero value otherwise.
func (o *DnacTemplate) GetComposite() string {
	if o == nil || IsNil(o.Composite) {
		var ret string
		return ret
	}
	return *o.Composite
}

// GetCompositeOk returns a tuple with the Composite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacTemplate) GetCompositeOk() (*string, bool) {
	if o == nil || IsNil(o.Composite) {
		return nil, false
	}
	return o.Composite, true
}

// HasComposite returns a boolean if a field has been set.
func (o *DnacTemplate) HasComposite() bool {
	if o != nil && !IsNil(o.Composite) {
		return true
	}

	return false
}

// SetComposite gets a reference to the given string and assigns it to the Composite field.
func (o *DnacTemplate) SetComposite(v string) {
	o.Composite = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DnacTemplate) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacTemplate) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DnacTemplate) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *DnacTemplate) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *DnacTemplate) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacTemplate) GetProjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *DnacTemplate) HasProjectName() bool {
	if o != nil && !IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *DnacTemplate) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *DnacTemplate) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacTemplate) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *DnacTemplate) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *DnacTemplate) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *DnacTemplate) GetTemplateName() string {
	if o == nil || IsNil(o.TemplateName) {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacTemplate) GetTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateName) {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *DnacTemplate) HasTemplateName() bool {
	if o != nil && !IsNil(o.TemplateName) {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *DnacTemplate) SetTemplateName(v string) {
	o.TemplateName = &v
}

func (o DnacTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnacTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDnacInventoryEntity, errDnacInventoryEntity := json.Marshal(o.DnacInventoryEntity)
	if errDnacInventoryEntity != nil {
		return map[string]interface{}{}, errDnacInventoryEntity
	}
	errDnacInventoryEntity = json.Unmarshal([]byte(serializedDnacInventoryEntity), &toSerialize)
	if errDnacInventoryEntity != nil {
		return map[string]interface{}{}, errDnacInventoryEntity
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Composite) {
		toSerialize["Composite"] = o.Composite
	}
	if !IsNil(o.ProjectId) {
		toSerialize["ProjectId"] = o.ProjectId
	}
	if !IsNil(o.ProjectName) {
		toSerialize["ProjectName"] = o.ProjectName
	}
	if !IsNil(o.TemplateId) {
		toSerialize["TemplateId"] = o.TemplateId
	}
	if !IsNil(o.TemplateName) {
		toSerialize["TemplateName"] = o.TemplateName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DnacTemplate) UnmarshalJSON(data []byte) (err error) {
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
	type DnacTemplateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Value indicates the template is composite.
		Composite *string `json:"Composite,omitempty"`
		// Identity of the project template.
		ProjectId *string `json:"ProjectId,omitempty"`
		// Name of the project template.
		ProjectName *string `json:"ProjectName,omitempty"`
		// Unique identity of the template.
		TemplateId *string `json:"TemplateId,omitempty"`
		// Name assigned to the template.
		TemplateName *string `json:"TemplateName,omitempty"`
	}

	varDnacTemplateWithoutEmbeddedStruct := DnacTemplateWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varDnacTemplateWithoutEmbeddedStruct)
	if err == nil {
		varDnacTemplate := _DnacTemplate{}
		varDnacTemplate.ClassId = varDnacTemplateWithoutEmbeddedStruct.ClassId
		varDnacTemplate.ObjectType = varDnacTemplateWithoutEmbeddedStruct.ObjectType
		varDnacTemplate.Composite = varDnacTemplateWithoutEmbeddedStruct.Composite
		varDnacTemplate.ProjectId = varDnacTemplateWithoutEmbeddedStruct.ProjectId
		varDnacTemplate.ProjectName = varDnacTemplateWithoutEmbeddedStruct.ProjectName
		varDnacTemplate.TemplateId = varDnacTemplateWithoutEmbeddedStruct.TemplateId
		varDnacTemplate.TemplateName = varDnacTemplateWithoutEmbeddedStruct.TemplateName
		*o = DnacTemplate(varDnacTemplate)
	} else {
		return err
	}

	varDnacTemplate := _DnacTemplate{}

	err = json.Unmarshal(data, &varDnacTemplate)
	if err == nil {
		o.DnacInventoryEntity = varDnacTemplate.DnacInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Composite")
		delete(additionalProperties, "ProjectId")
		delete(additionalProperties, "ProjectName")
		delete(additionalProperties, "TemplateId")
		delete(additionalProperties, "TemplateName")

		// remove fields from embedded structs
		reflectDnacInventoryEntity := reflect.ValueOf(o.DnacInventoryEntity)
		for i := 0; i < reflectDnacInventoryEntity.Type().NumField(); i++ {
			t := reflectDnacInventoryEntity.Type().Field(i)

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

type NullableDnacTemplate struct {
	value *DnacTemplate
	isSet bool
}

func (v NullableDnacTemplate) Get() *DnacTemplate {
	return v.value
}

func (v *NullableDnacTemplate) Set(val *DnacTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableDnacTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableDnacTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnacTemplate(val *DnacTemplate) *NullableDnacTemplate {
	return &NullableDnacTemplate{value: val, isSet: true}
}

func (v NullableDnacTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnacTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
