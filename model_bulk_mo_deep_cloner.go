/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-13010
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// BulkMoDeepCloner The MODeepCloner interface facilitates making n number of deep copies of any resource instance which supports the CREATE operation. The copy is limited to creating copies of only the child references of the resource. The MO to be cloned should be specified as an MoRef object in the \"Source\". The \"Targets\" array should contain n JSON documents each compliant to RFC 7386.  For each target mo to be created, the user can specify the following - - new values for the identity properties, if applicable - new values for specific properties or references of the source MO which need to be overridden in the cloned object.
type BulkMoDeepCloner struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string     `json:"ObjectType"`
	ExcludeProperties []string   `json:"ExcludeProperties,omitempty"`
	Source            *MoMoRef   `json:"Source,omitempty"`
	Targets           []MoBaseMo `json:"Targets,omitempty"`
	// A user friendly short name to identify the workflow, optionally. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), comma or an underscore (_).
	WorkflowNameSuffix   *string                               `json:"WorkflowNameSuffix,omitempty"`
	AsyncResult          *BulkResultRelationship               `json:"AsyncResult,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkMoDeepCloner BulkMoDeepCloner

// NewBulkMoDeepCloner instantiates a new BulkMoDeepCloner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkMoDeepCloner(classId string, objectType string) *BulkMoDeepCloner {
	this := BulkMoDeepCloner{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBulkMoDeepClonerWithDefaults instantiates a new BulkMoDeepCloner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkMoDeepClonerWithDefaults() *BulkMoDeepCloner {
	this := BulkMoDeepCloner{}
	var classId string = "bulk.MoDeepCloner"
	this.ClassId = classId
	var objectType string = "bulk.MoDeepCloner"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkMoDeepCloner) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkMoDeepCloner) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkMoDeepCloner) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkMoDeepCloner) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkMoDeepCloner) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkMoDeepCloner) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExcludeProperties returns the ExcludeProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoDeepCloner) GetExcludeProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeProperties
}

// GetExcludePropertiesOk returns a tuple with the ExcludeProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoDeepCloner) GetExcludePropertiesOk() ([]string, bool) {
	if o == nil || o.ExcludeProperties == nil {
		return nil, false
	}
	return o.ExcludeProperties, true
}

// HasExcludeProperties returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasExcludeProperties() bool {
	if o != nil && o.ExcludeProperties != nil {
		return true
	}

	return false
}

// SetExcludeProperties gets a reference to the given []string and assigns it to the ExcludeProperties field.
func (o *BulkMoDeepCloner) SetExcludeProperties(v []string) {
	o.ExcludeProperties = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BulkMoDeepCloner) GetSource() MoMoRef {
	if o == nil || o.Source == nil {
		var ret MoMoRef
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoDeepCloner) GetSourceOk() (*MoMoRef, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given MoMoRef and assigns it to the Source field.
func (o *BulkMoDeepCloner) SetSource(v MoMoRef) {
	o.Source = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoDeepCloner) GetTargets() []MoBaseMo {
	if o == nil {
		var ret []MoBaseMo
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoDeepCloner) GetTargetsOk() ([]MoBaseMo, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []MoBaseMo and assigns it to the Targets field.
func (o *BulkMoDeepCloner) SetTargets(v []MoBaseMo) {
	o.Targets = v
}

// GetWorkflowNameSuffix returns the WorkflowNameSuffix field value if set, zero value otherwise.
func (o *BulkMoDeepCloner) GetWorkflowNameSuffix() string {
	if o == nil || o.WorkflowNameSuffix == nil {
		var ret string
		return ret
	}
	return *o.WorkflowNameSuffix
}

// GetWorkflowNameSuffixOk returns a tuple with the WorkflowNameSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoDeepCloner) GetWorkflowNameSuffixOk() (*string, bool) {
	if o == nil || o.WorkflowNameSuffix == nil {
		return nil, false
	}
	return o.WorkflowNameSuffix, true
}

// HasWorkflowNameSuffix returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasWorkflowNameSuffix() bool {
	if o != nil && o.WorkflowNameSuffix != nil {
		return true
	}

	return false
}

// SetWorkflowNameSuffix gets a reference to the given string and assigns it to the WorkflowNameSuffix field.
func (o *BulkMoDeepCloner) SetWorkflowNameSuffix(v string) {
	o.WorkflowNameSuffix = &v
}

// GetAsyncResult returns the AsyncResult field value if set, zero value otherwise.
func (o *BulkMoDeepCloner) GetAsyncResult() BulkResultRelationship {
	if o == nil || o.AsyncResult == nil {
		var ret BulkResultRelationship
		return ret
	}
	return *o.AsyncResult
}

// GetAsyncResultOk returns a tuple with the AsyncResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoDeepCloner) GetAsyncResultOk() (*BulkResultRelationship, bool) {
	if o == nil || o.AsyncResult == nil {
		return nil, false
	}
	return o.AsyncResult, true
}

// HasAsyncResult returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasAsyncResult() bool {
	if o != nil && o.AsyncResult != nil {
		return true
	}

	return false
}

// SetAsyncResult gets a reference to the given BulkResultRelationship and assigns it to the AsyncResult field.
func (o *BulkMoDeepCloner) SetAsyncResult(v BulkResultRelationship) {
	o.AsyncResult = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *BulkMoDeepCloner) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoDeepCloner) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkMoDeepCloner) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkMoDeepCloner) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o BulkMoDeepCloner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExcludeProperties != nil {
		toSerialize["ExcludeProperties"] = o.ExcludeProperties
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.Targets != nil {
		toSerialize["Targets"] = o.Targets
	}
	if o.WorkflowNameSuffix != nil {
		toSerialize["WorkflowNameSuffix"] = o.WorkflowNameSuffix
	}
	if o.AsyncResult != nil {
		toSerialize["AsyncResult"] = o.AsyncResult
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkMoDeepCloner) UnmarshalJSON(bytes []byte) (err error) {
	type BulkMoDeepClonerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string     `json:"ObjectType"`
		ExcludeProperties []string   `json:"ExcludeProperties,omitempty"`
		Source            *MoMoRef   `json:"Source,omitempty"`
		Targets           []MoBaseMo `json:"Targets,omitempty"`
		// A user friendly short name to identify the workflow, optionally. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), comma or an underscore (_).
		WorkflowNameSuffix *string                               `json:"WorkflowNameSuffix,omitempty"`
		AsyncResult        *BulkResultRelationship               `json:"AsyncResult,omitempty"`
		Organization       *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varBulkMoDeepClonerWithoutEmbeddedStruct := BulkMoDeepClonerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBulkMoDeepClonerWithoutEmbeddedStruct)
	if err == nil {
		varBulkMoDeepCloner := _BulkMoDeepCloner{}
		varBulkMoDeepCloner.ClassId = varBulkMoDeepClonerWithoutEmbeddedStruct.ClassId
		varBulkMoDeepCloner.ObjectType = varBulkMoDeepClonerWithoutEmbeddedStruct.ObjectType
		varBulkMoDeepCloner.ExcludeProperties = varBulkMoDeepClonerWithoutEmbeddedStruct.ExcludeProperties
		varBulkMoDeepCloner.Source = varBulkMoDeepClonerWithoutEmbeddedStruct.Source
		varBulkMoDeepCloner.Targets = varBulkMoDeepClonerWithoutEmbeddedStruct.Targets
		varBulkMoDeepCloner.WorkflowNameSuffix = varBulkMoDeepClonerWithoutEmbeddedStruct.WorkflowNameSuffix
		varBulkMoDeepCloner.AsyncResult = varBulkMoDeepClonerWithoutEmbeddedStruct.AsyncResult
		varBulkMoDeepCloner.Organization = varBulkMoDeepClonerWithoutEmbeddedStruct.Organization
		*o = BulkMoDeepCloner(varBulkMoDeepCloner)
	} else {
		return err
	}

	varBulkMoDeepCloner := _BulkMoDeepCloner{}

	err = json.Unmarshal(bytes, &varBulkMoDeepCloner)
	if err == nil {
		o.MoBaseMo = varBulkMoDeepCloner.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExcludeProperties")
		delete(additionalProperties, "Source")
		delete(additionalProperties, "Targets")
		delete(additionalProperties, "WorkflowNameSuffix")
		delete(additionalProperties, "AsyncResult")
		delete(additionalProperties, "Organization")

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

type NullableBulkMoDeepCloner struct {
	value *BulkMoDeepCloner
	isSet bool
}

func (v NullableBulkMoDeepCloner) Get() *BulkMoDeepCloner {
	return v.value
}

func (v *NullableBulkMoDeepCloner) Set(val *BulkMoDeepCloner) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkMoDeepCloner) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkMoDeepCloner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkMoDeepCloner(val *BulkMoDeepCloner) *NullableBulkMoDeepCloner {
	return &NullableBulkMoDeepCloner{value: val, isSet: true}
}

func (v NullableBulkMoDeepCloner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkMoDeepCloner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}