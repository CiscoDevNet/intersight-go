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

// checks if the FabricSwitchClusterProfileTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricSwitchClusterProfileTemplate{}

// FabricSwitchClusterProfileTemplate The Switch Cluster Profile Template consists of common switch profile configurations, which can be reused across multiple profiles. Switch Cluster Profiles can be created from the template using the Derive operation. Additionally, an existing profile can be attached to a template to use the configuration set in the template. To derive switch cluster profiles from a switch cluster profile template, you must use the synchronous /v1/bulk/MoCloners bulk API. Deriving profiles from a Switch Profile Template URL: /v1/bulk/MoCloners Method: POST Body: >  {     \"Sources\":[       {         \"Moid\":\"64fb5d17656e6f301e43045b\",         \"ObjectType\":\"fabric.SwitchClusterProfileTemplate\"       }],     \"Targets\":[       {         \"Name\":\"template1_DERIVED-1\",         \"ObjectType\":\"fabric.SwitchClusterProfile\",         \"ClusterAssignments\": [           {             \"SourceSwitchProfileOrTemplateName\":\"template1-A\",             \"NetworkElement\":               {                 \"ObjectType\": \"network.Element\",                 \"Moid\": \"64fe8802617675301eb3bdaf\"               }           },           {             \"SourceSwitchProfileOrTemplateName\":\"template1-B\",             \"NetworkElement\":               {                 \"ObjectType\": \"network.Element\",                 \"Moid\": \"64fe8802617675301eb3bdc0\"               }           }],         \"Organization\":           {             \"ObjectType\":\"organization.Organization\",             \"Moid\":\"64b0b9ef697265301e52ea0c\"           },         \"Description\":\"\",         \"Tags\":[]       }]  } The API response includes the derived SwitchClusterProfile object details. Template Updates When the profile template is updated, a call to the /v1/bulk/MoMergers API is to be made by the client to synchronize the template changes to all derived profile instances. Updating profiles from a Switch Profile Template URL: /v1/bulk/MoMergers Method: POST Body: >  {     \"Sources\":[       {         \"Moid\":\"64fb5d17656e6f301e43045b\",         \"ObjectType\":\"fabric.SwitchClusterProfileTemplate\"       }],     \"Targets\":[       {         \"Moid\":\"6502ffc8656e6f301e5e9f6b\",         \"ObjectType\":\"fabric.SwitchClusterProfile\"       }],     \"MergeAction\":\"Replace\"  } The response of the MoMerger API call would contain the changed profiles.
type FabricSwitchClusterProfileTemplate struct {
	FabricBaseClusterProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The count of switch cluster profiles derived from the template.
	Usage        *int64                                       `json:"Usage,omitempty"`
	Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to fabricSwitchProfileTemplate resources.
	SwitchProfileTemplates []FabricSwitchProfileTemplateRelationship `json:"SwitchProfileTemplates,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _FabricSwitchClusterProfileTemplate FabricSwitchClusterProfileTemplate

// NewFabricSwitchClusterProfileTemplate instantiates a new FabricSwitchClusterProfileTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSwitchClusterProfileTemplate(classId string, objectType string) *FabricSwitchClusterProfileTemplate {
	this := FabricSwitchClusterProfileTemplate{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	return &this
}

// NewFabricSwitchClusterProfileTemplateWithDefaults instantiates a new FabricSwitchClusterProfileTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSwitchClusterProfileTemplateWithDefaults() *FabricSwitchClusterProfileTemplate {
	this := FabricSwitchClusterProfileTemplate{}
	var classId string = "fabric.SwitchClusterProfileTemplate"
	this.ClassId = classId
	var objectType string = "fabric.SwitchClusterProfileTemplate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSwitchClusterProfileTemplate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchClusterProfileTemplate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSwitchClusterProfileTemplate) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.SwitchClusterProfileTemplate" of the ClassId field.
func (o *FabricSwitchClusterProfileTemplate) GetDefaultClassId() interface{} {
	return "fabric.SwitchClusterProfileTemplate"
}

// GetObjectType returns the ObjectType field value
func (o *FabricSwitchClusterProfileTemplate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchClusterProfileTemplate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSwitchClusterProfileTemplate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.SwitchClusterProfileTemplate" of the ObjectType field.
func (o *FabricSwitchClusterProfileTemplate) GetDefaultObjectType() interface{} {
	return "fabric.SwitchClusterProfileTemplate"
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *FabricSwitchClusterProfileTemplate) GetUsage() int64 {
	if o == nil || IsNil(o.Usage) {
		var ret int64
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchClusterProfileTemplate) GetUsageOk() (*int64, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *FabricSwitchClusterProfileTemplate) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given int64 and assigns it to the Usage field.
func (o *FabricSwitchClusterProfileTemplate) SetUsage(v int64) {
	o.Usage = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchClusterProfileTemplate) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchClusterProfileTemplate) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricSwitchClusterProfileTemplate) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricSwitchClusterProfileTemplate) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *FabricSwitchClusterProfileTemplate) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *FabricSwitchClusterProfileTemplate) UnsetOrganization() {
	o.Organization.Unset()
}

// GetSwitchProfileTemplates returns the SwitchProfileTemplates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchClusterProfileTemplate) GetSwitchProfileTemplates() []FabricSwitchProfileTemplateRelationship {
	if o == nil {
		var ret []FabricSwitchProfileTemplateRelationship
		return ret
	}
	return o.SwitchProfileTemplates
}

// GetSwitchProfileTemplatesOk returns a tuple with the SwitchProfileTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchClusterProfileTemplate) GetSwitchProfileTemplatesOk() ([]FabricSwitchProfileTemplateRelationship, bool) {
	if o == nil || IsNil(o.SwitchProfileTemplates) {
		return nil, false
	}
	return o.SwitchProfileTemplates, true
}

// HasSwitchProfileTemplates returns a boolean if a field has been set.
func (o *FabricSwitchClusterProfileTemplate) HasSwitchProfileTemplates() bool {
	if o != nil && !IsNil(o.SwitchProfileTemplates) {
		return true
	}

	return false
}

// SetSwitchProfileTemplates gets a reference to the given []FabricSwitchProfileTemplateRelationship and assigns it to the SwitchProfileTemplates field.
func (o *FabricSwitchClusterProfileTemplate) SetSwitchProfileTemplates(v []FabricSwitchProfileTemplateRelationship) {
	o.SwitchProfileTemplates = v
}

func (o FabricSwitchClusterProfileTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricSwitchClusterProfileTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricBaseClusterProfile, errFabricBaseClusterProfile := json.Marshal(o.FabricBaseClusterProfile)
	if errFabricBaseClusterProfile != nil {
		return map[string]interface{}{}, errFabricBaseClusterProfile
	}
	errFabricBaseClusterProfile = json.Unmarshal([]byte(serializedFabricBaseClusterProfile), &toSerialize)
	if errFabricBaseClusterProfile != nil {
		return map[string]interface{}{}, errFabricBaseClusterProfile
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Usage) {
		toSerialize["Usage"] = o.Usage
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.SwitchProfileTemplates != nil {
		toSerialize["SwitchProfileTemplates"] = o.SwitchProfileTemplates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricSwitchClusterProfileTemplate) UnmarshalJSON(data []byte) (err error) {
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
	type FabricSwitchClusterProfileTemplateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The count of switch cluster profiles derived from the template.
		Usage        *int64                                       `json:"Usage,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to fabricSwitchProfileTemplate resources.
		SwitchProfileTemplates []FabricSwitchProfileTemplateRelationship `json:"SwitchProfileTemplates,omitempty"`
	}

	varFabricSwitchClusterProfileTemplateWithoutEmbeddedStruct := FabricSwitchClusterProfileTemplateWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricSwitchClusterProfileTemplateWithoutEmbeddedStruct)
	if err == nil {
		varFabricSwitchClusterProfileTemplate := _FabricSwitchClusterProfileTemplate{}
		varFabricSwitchClusterProfileTemplate.ClassId = varFabricSwitchClusterProfileTemplateWithoutEmbeddedStruct.ClassId
		varFabricSwitchClusterProfileTemplate.ObjectType = varFabricSwitchClusterProfileTemplateWithoutEmbeddedStruct.ObjectType
		varFabricSwitchClusterProfileTemplate.Usage = varFabricSwitchClusterProfileTemplateWithoutEmbeddedStruct.Usage
		varFabricSwitchClusterProfileTemplate.Organization = varFabricSwitchClusterProfileTemplateWithoutEmbeddedStruct.Organization
		varFabricSwitchClusterProfileTemplate.SwitchProfileTemplates = varFabricSwitchClusterProfileTemplateWithoutEmbeddedStruct.SwitchProfileTemplates
		*o = FabricSwitchClusterProfileTemplate(varFabricSwitchClusterProfileTemplate)
	} else {
		return err
	}

	varFabricSwitchClusterProfileTemplate := _FabricSwitchClusterProfileTemplate{}

	err = json.Unmarshal(data, &varFabricSwitchClusterProfileTemplate)
	if err == nil {
		o.FabricBaseClusterProfile = varFabricSwitchClusterProfileTemplate.FabricBaseClusterProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Usage")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "SwitchProfileTemplates")

		// remove fields from embedded structs
		reflectFabricBaseClusterProfile := reflect.ValueOf(o.FabricBaseClusterProfile)
		for i := 0; i < reflectFabricBaseClusterProfile.Type().NumField(); i++ {
			t := reflectFabricBaseClusterProfile.Type().Field(i)

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

type NullableFabricSwitchClusterProfileTemplate struct {
	value *FabricSwitchClusterProfileTemplate
	isSet bool
}

func (v NullableFabricSwitchClusterProfileTemplate) Get() *FabricSwitchClusterProfileTemplate {
	return v.value
}

func (v *NullableFabricSwitchClusterProfileTemplate) Set(val *FabricSwitchClusterProfileTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSwitchClusterProfileTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSwitchClusterProfileTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSwitchClusterProfileTemplate(val *FabricSwitchClusterProfileTemplate) *NullableFabricSwitchClusterProfileTemplate {
	return &NullableFabricSwitchClusterProfileTemplate{value: val, isSet: true}
}

func (v NullableFabricSwitchClusterProfileTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSwitchClusterProfileTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
