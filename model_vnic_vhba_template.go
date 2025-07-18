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

// checks if the VnicVhbaTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicVhbaTemplate{}

// VnicVhbaTemplate The vHBA template consists of the common vHBA configuration, which can be reused across multiple vHBAs. vHBAs can be created from the template using the Derive operation. Additionally, an existing vHBA can be attached to a template to use the configuration set in the template. To derive vHBAs from a vHBA template, you must use the asynchronous /v1/bulk/MoCloners bulk API. Deriving vHBAs from a vHBA Template URL: /v1/bulk/MoCloners Method: POST Headers:   - Key: \"prefer\"     Value: \"respond-async\" Body: >  {     \"Sources\": [       {         \"Moid\": \"xxxx\",         \"ObjectType\": \"vnic.VhbaTemplate\"      }     ],     \"Targets\": [       {         \"Name\": \"test-vh2\",         \"SanConnectivityPolicy\": \"aaaaa\",         \"ObjectType\": \"vnic.FcIf\"      }     ],     \"WorkflowNameSuffix\": \"Derive vHBA from a Template\",     \"Organization\": {         \"Moid\": \"oooo\",         \"ObjectType\": \"organization.Organization\"     } } The API response includes the \"AsyncResult\"->bulk.Result MO reference. API clients must poll the bulk.Result MO to monitor the status of the operation. The bulk.Result object also includes a reference to a monitoring workflow, which is accessible from the 'Requests' tab in the UI. Managing Template Updates When the vHBA template is updated, the system initiates an automatic call to the /v1/bulk/MoMergers API to synchronize the template changes to all derived vHBA instances asynchronously. The status of the sync operation can be obtained by performing a query on vnic EthIf MO - $filter=SrcTemplate.Moid eq 'xxx'&$apply=groupby ( (TemplateSyncStatus), aggregate ($count as count)) Override Option for vHBA Templates When enabled, this feature allows the configuration of the derived vHBA to override the configuration available in the template during vHBA create or update. You can query the list of overridable properties for a vHBA template from the Template Catalog. Use the following catalog API query: URL: /v1/capability/TemplateCatalogs Query: $filter= (Name eq ‘VhbaTemplate’)&$select=AllowedOverrideList Once a property is overridden on a derived vHBA, it will be added to the ‘OverriddenList’ property on the vnic FcIf MO. URL: /v1/vnic/FcIfs Query: $select=SrcTemplate, OverriddenList When override is disabled on the template, derived vHBA will have same configuration as the template.
type VnicVhbaTemplate struct {
	VnicBaseFcIf
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the vHBA template.
	Description *string `json:"Description,omitempty"`
	// When enabled, the configuration of the derived instances may override the template configuration.
	EnableOverride *bool `json:"EnableOverride,omitempty"`
	// Name of the vHBA template.
	Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9-._:]+$"`
	// Name of the peer vHBA which belongs to the peer FI.
	PeerVhbaName *string `json:"PeerVhbaName,omitempty"`
	// The count of the San Connectivity Policies using vHBA template.
	ScpUsageCount *int64 `json:"ScpUsageCount,omitempty"`
	// The fabric port to which the vHBAs will be associated. * `None` - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used. * `A` - Fabric A of the FI cluster. * `B` - Fabric B of the FI cluster.
	SwitchId        *string                 `json:"SwitchId,omitempty"`
	TemplateActions []MotemplateActionEntry `json:"TemplateActions,omitempty"`
	// The template sync status with all derived objects. * `None` - The Enum value represents that the object is not attached to any template. * `OK` - The Enum value represents that the object values are in sync with attached template. * `Scheduled` - The Enum value represents that the object sync from attached template is scheduled from template. * `InProgress` - The Enum value represents that the object sync with the attached template is in progress. * `OutOfSync` - The Enum value represents that the object values are not in sync with attached template.
	UpdateStatus *string `json:"UpdateStatus,omitempty"`
	// The number of objects derived from a Template MO instance.
	UsageCount           *int64                                       `json:"UsageCount,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicVhbaTemplate VnicVhbaTemplate

// NewVnicVhbaTemplate instantiates a new VnicVhbaTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicVhbaTemplate(classId string, objectType string) *VnicVhbaTemplate {
	this := VnicVhbaTemplate{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "fc-initiator"
	this.Type = &type_
	var switchId string = "None"
	this.SwitchId = &switchId
	return &this
}

// NewVnicVhbaTemplateWithDefaults instantiates a new VnicVhbaTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicVhbaTemplateWithDefaults() *VnicVhbaTemplate {
	this := VnicVhbaTemplate{}
	var classId string = "vnic.VhbaTemplate"
	this.ClassId = classId
	var objectType string = "vnic.VhbaTemplate"
	this.ObjectType = objectType
	var switchId string = "None"
	this.SwitchId = &switchId
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicVhbaTemplate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicVhbaTemplate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicVhbaTemplate) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.VhbaTemplate" of the ClassId field.
func (o *VnicVhbaTemplate) GetDefaultClassId() interface{} {
	return "vnic.VhbaTemplate"
}

// GetObjectType returns the ObjectType field value
func (o *VnicVhbaTemplate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicVhbaTemplate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicVhbaTemplate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.VhbaTemplate" of the ObjectType field.
func (o *VnicVhbaTemplate) GetDefaultObjectType() interface{} {
	return "vnic.VhbaTemplate"
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VnicVhbaTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVhbaTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VnicVhbaTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VnicVhbaTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetEnableOverride returns the EnableOverride field value if set, zero value otherwise.
func (o *VnicVhbaTemplate) GetEnableOverride() bool {
	if o == nil || IsNil(o.EnableOverride) {
		var ret bool
		return ret
	}
	return *o.EnableOverride
}

// GetEnableOverrideOk returns a tuple with the EnableOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVhbaTemplate) GetEnableOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableOverride) {
		return nil, false
	}
	return o.EnableOverride, true
}

// HasEnableOverride returns a boolean if a field has been set.
func (o *VnicVhbaTemplate) HasEnableOverride() bool {
	if o != nil && !IsNil(o.EnableOverride) {
		return true
	}

	return false
}

// SetEnableOverride gets a reference to the given bool and assigns it to the EnableOverride field.
func (o *VnicVhbaTemplate) SetEnableOverride(v bool) {
	o.EnableOverride = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VnicVhbaTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVhbaTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VnicVhbaTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VnicVhbaTemplate) SetName(v string) {
	o.Name = &v
}

// GetPeerVhbaName returns the PeerVhbaName field value if set, zero value otherwise.
func (o *VnicVhbaTemplate) GetPeerVhbaName() string {
	if o == nil || IsNil(o.PeerVhbaName) {
		var ret string
		return ret
	}
	return *o.PeerVhbaName
}

// GetPeerVhbaNameOk returns a tuple with the PeerVhbaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVhbaTemplate) GetPeerVhbaNameOk() (*string, bool) {
	if o == nil || IsNil(o.PeerVhbaName) {
		return nil, false
	}
	return o.PeerVhbaName, true
}

// HasPeerVhbaName returns a boolean if a field has been set.
func (o *VnicVhbaTemplate) HasPeerVhbaName() bool {
	if o != nil && !IsNil(o.PeerVhbaName) {
		return true
	}

	return false
}

// SetPeerVhbaName gets a reference to the given string and assigns it to the PeerVhbaName field.
func (o *VnicVhbaTemplate) SetPeerVhbaName(v string) {
	o.PeerVhbaName = &v
}

// GetScpUsageCount returns the ScpUsageCount field value if set, zero value otherwise.
func (o *VnicVhbaTemplate) GetScpUsageCount() int64 {
	if o == nil || IsNil(o.ScpUsageCount) {
		var ret int64
		return ret
	}
	return *o.ScpUsageCount
}

// GetScpUsageCountOk returns a tuple with the ScpUsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVhbaTemplate) GetScpUsageCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ScpUsageCount) {
		return nil, false
	}
	return o.ScpUsageCount, true
}

// HasScpUsageCount returns a boolean if a field has been set.
func (o *VnicVhbaTemplate) HasScpUsageCount() bool {
	if o != nil && !IsNil(o.ScpUsageCount) {
		return true
	}

	return false
}

// SetScpUsageCount gets a reference to the given int64 and assigns it to the ScpUsageCount field.
func (o *VnicVhbaTemplate) SetScpUsageCount(v int64) {
	o.ScpUsageCount = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *VnicVhbaTemplate) GetSwitchId() string {
	if o == nil || IsNil(o.SwitchId) {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVhbaTemplate) GetSwitchIdOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchId) {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *VnicVhbaTemplate) HasSwitchId() bool {
	if o != nil && !IsNil(o.SwitchId) {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *VnicVhbaTemplate) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetTemplateActions returns the TemplateActions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicVhbaTemplate) GetTemplateActions() []MotemplateActionEntry {
	if o == nil {
		var ret []MotemplateActionEntry
		return ret
	}
	return o.TemplateActions
}

// GetTemplateActionsOk returns a tuple with the TemplateActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicVhbaTemplate) GetTemplateActionsOk() ([]MotemplateActionEntry, bool) {
	if o == nil || IsNil(o.TemplateActions) {
		return nil, false
	}
	return o.TemplateActions, true
}

// HasTemplateActions returns a boolean if a field has been set.
func (o *VnicVhbaTemplate) HasTemplateActions() bool {
	if o != nil && !IsNil(o.TemplateActions) {
		return true
	}

	return false
}

// SetTemplateActions gets a reference to the given []MotemplateActionEntry and assigns it to the TemplateActions field.
func (o *VnicVhbaTemplate) SetTemplateActions(v []MotemplateActionEntry) {
	o.TemplateActions = v
}

// GetUpdateStatus returns the UpdateStatus field value if set, zero value otherwise.
func (o *VnicVhbaTemplate) GetUpdateStatus() string {
	if o == nil || IsNil(o.UpdateStatus) {
		var ret string
		return ret
	}
	return *o.UpdateStatus
}

// GetUpdateStatusOk returns a tuple with the UpdateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVhbaTemplate) GetUpdateStatusOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateStatus) {
		return nil, false
	}
	return o.UpdateStatus, true
}

// HasUpdateStatus returns a boolean if a field has been set.
func (o *VnicVhbaTemplate) HasUpdateStatus() bool {
	if o != nil && !IsNil(o.UpdateStatus) {
		return true
	}

	return false
}

// SetUpdateStatus gets a reference to the given string and assigns it to the UpdateStatus field.
func (o *VnicVhbaTemplate) SetUpdateStatus(v string) {
	o.UpdateStatus = &v
}

// GetUsageCount returns the UsageCount field value if set, zero value otherwise.
func (o *VnicVhbaTemplate) GetUsageCount() int64 {
	if o == nil || IsNil(o.UsageCount) {
		var ret int64
		return ret
	}
	return *o.UsageCount
}

// GetUsageCountOk returns a tuple with the UsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVhbaTemplate) GetUsageCountOk() (*int64, bool) {
	if o == nil || IsNil(o.UsageCount) {
		return nil, false
	}
	return o.UsageCount, true
}

// HasUsageCount returns a boolean if a field has been set.
func (o *VnicVhbaTemplate) HasUsageCount() bool {
	if o != nil && !IsNil(o.UsageCount) {
		return true
	}

	return false
}

// SetUsageCount gets a reference to the given int64 and assigns it to the UsageCount field.
func (o *VnicVhbaTemplate) SetUsageCount(v int64) {
	o.UsageCount = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicVhbaTemplate) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicVhbaTemplate) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicVhbaTemplate) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicVhbaTemplate) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *VnicVhbaTemplate) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *VnicVhbaTemplate) UnsetOrganization() {
	o.Organization.Unset()
}

func (o VnicVhbaTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicVhbaTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVnicBaseFcIf, errVnicBaseFcIf := json.Marshal(o.VnicBaseFcIf)
	if errVnicBaseFcIf != nil {
		return map[string]interface{}{}, errVnicBaseFcIf
	}
	errVnicBaseFcIf = json.Unmarshal([]byte(serializedVnicBaseFcIf), &toSerialize)
	if errVnicBaseFcIf != nil {
		return map[string]interface{}{}, errVnicBaseFcIf
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.EnableOverride) {
		toSerialize["EnableOverride"] = o.EnableOverride
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PeerVhbaName) {
		toSerialize["PeerVhbaName"] = o.PeerVhbaName
	}
	if !IsNil(o.ScpUsageCount) {
		toSerialize["ScpUsageCount"] = o.ScpUsageCount
	}
	if !IsNil(o.SwitchId) {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.TemplateActions != nil {
		toSerialize["TemplateActions"] = o.TemplateActions
	}
	if !IsNil(o.UpdateStatus) {
		toSerialize["UpdateStatus"] = o.UpdateStatus
	}
	if !IsNil(o.UsageCount) {
		toSerialize["UsageCount"] = o.UsageCount
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicVhbaTemplate) UnmarshalJSON(data []byte) (err error) {
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
	type VnicVhbaTemplateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of the vHBA template.
		Description *string `json:"Description,omitempty"`
		// When enabled, the configuration of the derived instances may override the template configuration.
		EnableOverride *bool `json:"EnableOverride,omitempty"`
		// Name of the vHBA template.
		Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9-._:]+$"`
		// Name of the peer vHBA which belongs to the peer FI.
		PeerVhbaName *string `json:"PeerVhbaName,omitempty"`
		// The count of the San Connectivity Policies using vHBA template.
		ScpUsageCount *int64 `json:"ScpUsageCount,omitempty"`
		// The fabric port to which the vHBAs will be associated. * `None` - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used. * `A` - Fabric A of the FI cluster. * `B` - Fabric B of the FI cluster.
		SwitchId        *string                 `json:"SwitchId,omitempty"`
		TemplateActions []MotemplateActionEntry `json:"TemplateActions,omitempty"`
		// The template sync status with all derived objects. * `None` - The Enum value represents that the object is not attached to any template. * `OK` - The Enum value represents that the object values are in sync with attached template. * `Scheduled` - The Enum value represents that the object sync from attached template is scheduled from template. * `InProgress` - The Enum value represents that the object sync with the attached template is in progress. * `OutOfSync` - The Enum value represents that the object values are not in sync with attached template.
		UpdateStatus *string `json:"UpdateStatus,omitempty"`
		// The number of objects derived from a Template MO instance.
		UsageCount   *int64                                       `json:"UsageCount,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varVnicVhbaTemplateWithoutEmbeddedStruct := VnicVhbaTemplateWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicVhbaTemplateWithoutEmbeddedStruct)
	if err == nil {
		varVnicVhbaTemplate := _VnicVhbaTemplate{}
		varVnicVhbaTemplate.ClassId = varVnicVhbaTemplateWithoutEmbeddedStruct.ClassId
		varVnicVhbaTemplate.ObjectType = varVnicVhbaTemplateWithoutEmbeddedStruct.ObjectType
		varVnicVhbaTemplate.Description = varVnicVhbaTemplateWithoutEmbeddedStruct.Description
		varVnicVhbaTemplate.EnableOverride = varVnicVhbaTemplateWithoutEmbeddedStruct.EnableOverride
		varVnicVhbaTemplate.Name = varVnicVhbaTemplateWithoutEmbeddedStruct.Name
		varVnicVhbaTemplate.PeerVhbaName = varVnicVhbaTemplateWithoutEmbeddedStruct.PeerVhbaName
		varVnicVhbaTemplate.ScpUsageCount = varVnicVhbaTemplateWithoutEmbeddedStruct.ScpUsageCount
		varVnicVhbaTemplate.SwitchId = varVnicVhbaTemplateWithoutEmbeddedStruct.SwitchId
		varVnicVhbaTemplate.TemplateActions = varVnicVhbaTemplateWithoutEmbeddedStruct.TemplateActions
		varVnicVhbaTemplate.UpdateStatus = varVnicVhbaTemplateWithoutEmbeddedStruct.UpdateStatus
		varVnicVhbaTemplate.UsageCount = varVnicVhbaTemplateWithoutEmbeddedStruct.UsageCount
		varVnicVhbaTemplate.Organization = varVnicVhbaTemplateWithoutEmbeddedStruct.Organization
		*o = VnicVhbaTemplate(varVnicVhbaTemplate)
	} else {
		return err
	}

	varVnicVhbaTemplate := _VnicVhbaTemplate{}

	err = json.Unmarshal(data, &varVnicVhbaTemplate)
	if err == nil {
		o.VnicBaseFcIf = varVnicVhbaTemplate.VnicBaseFcIf
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "EnableOverride")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PeerVhbaName")
		delete(additionalProperties, "ScpUsageCount")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "TemplateActions")
		delete(additionalProperties, "UpdateStatus")
		delete(additionalProperties, "UsageCount")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectVnicBaseFcIf := reflect.ValueOf(o.VnicBaseFcIf)
		for i := 0; i < reflectVnicBaseFcIf.Type().NumField(); i++ {
			t := reflectVnicBaseFcIf.Type().Field(i)

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

type NullableVnicVhbaTemplate struct {
	value *VnicVhbaTemplate
	isSet bool
}

func (v NullableVnicVhbaTemplate) Get() *VnicVhbaTemplate {
	return v.value
}

func (v *NullableVnicVhbaTemplate) Set(val *VnicVhbaTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicVhbaTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicVhbaTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicVhbaTemplate(val *VnicVhbaTemplate) *NullableVnicVhbaTemplate {
	return &NullableVnicVhbaTemplate{value: val, isSet: true}
}

func (v NullableVnicVhbaTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicVhbaTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
