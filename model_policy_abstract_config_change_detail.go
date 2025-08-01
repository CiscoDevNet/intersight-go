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

// checks if the PolicyAbstractConfigChangeDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyAbstractConfigChangeDetail{}

// PolicyAbstractConfigChangeDetail Defines the details of a configuration change including change type, disruption and description of the change.
type PolicyAbstractConfigChangeDetail struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType          string                            `json:"ObjectType"`
	Changes             []string                          `json:"Changes,omitempty"`
	ConfigChangeContext NullablePolicyConfigResultContext `json:"ConfigChangeContext,omitempty"`
	// Config change flag to differentiate Pending-changes and Config-drift. * `Pending-changes` - Config change flag represents changes are due to not deployed changes from Intersight. * `Drift-changes` - Config change flag represents changes are due to endpoint configuration changes.
	ConfigChangeFlag *string  `json:"ConfigChangeFlag,omitempty"`
	Disruptions      []string `json:"Disruptions,omitempty"`
	// Detailed description of the config change.
	Message *string `json:"Message,omitempty"`
	// Modification status of the mo in this config change. * `None` - The 'none' operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration. * `Created` - The 'create' operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet. * `Modified` - The 'update' operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed successfully. * `Deleted` - The 'delete' operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet.
	ModStatus            *string `json:"ModStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAbstractConfigChangeDetail PolicyAbstractConfigChangeDetail

// NewPolicyAbstractConfigChangeDetail instantiates a new PolicyAbstractConfigChangeDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAbstractConfigChangeDetail(classId string, objectType string) *PolicyAbstractConfigChangeDetail {
	this := PolicyAbstractConfigChangeDetail{}
	this.ClassId = classId
	this.ObjectType = objectType
	var configChangeFlag string = "Pending-changes"
	this.ConfigChangeFlag = &configChangeFlag
	var modStatus string = "None"
	this.ModStatus = &modStatus
	return &this
}

// NewPolicyAbstractConfigChangeDetailWithDefaults instantiates a new PolicyAbstractConfigChangeDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAbstractConfigChangeDetailWithDefaults() *PolicyAbstractConfigChangeDetail {
	this := PolicyAbstractConfigChangeDetail{}
	var configChangeFlag string = "Pending-changes"
	this.ConfigChangeFlag = &configChangeFlag
	var modStatus string = "None"
	this.ModStatus = &modStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyAbstractConfigChangeDetail) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigChangeDetail) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyAbstractConfigChangeDetail) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyAbstractConfigChangeDetail) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigChangeDetail) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyAbstractConfigChangeDetail) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChanges returns the Changes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAbstractConfigChangeDetail) GetChanges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAbstractConfigChangeDetail) GetChangesOk() ([]string, bool) {
	if o == nil || IsNil(o.Changes) {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *PolicyAbstractConfigChangeDetail) HasChanges() bool {
	if o != nil && !IsNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []string and assigns it to the Changes field.
func (o *PolicyAbstractConfigChangeDetail) SetChanges(v []string) {
	o.Changes = v
}

// GetConfigChangeContext returns the ConfigChangeContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAbstractConfigChangeDetail) GetConfigChangeContext() PolicyConfigResultContext {
	if o == nil || IsNil(o.ConfigChangeContext.Get()) {
		var ret PolicyConfigResultContext
		return ret
	}
	return *o.ConfigChangeContext.Get()
}

// GetConfigChangeContextOk returns a tuple with the ConfigChangeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAbstractConfigChangeDetail) GetConfigChangeContextOk() (*PolicyConfigResultContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChangeContext.Get(), o.ConfigChangeContext.IsSet()
}

// HasConfigChangeContext returns a boolean if a field has been set.
func (o *PolicyAbstractConfigChangeDetail) HasConfigChangeContext() bool {
	if o != nil && o.ConfigChangeContext.IsSet() {
		return true
	}

	return false
}

// SetConfigChangeContext gets a reference to the given NullablePolicyConfigResultContext and assigns it to the ConfigChangeContext field.
func (o *PolicyAbstractConfigChangeDetail) SetConfigChangeContext(v PolicyConfigResultContext) {
	o.ConfigChangeContext.Set(&v)
}

// SetConfigChangeContextNil sets the value for ConfigChangeContext to be an explicit nil
func (o *PolicyAbstractConfigChangeDetail) SetConfigChangeContextNil() {
	o.ConfigChangeContext.Set(nil)
}

// UnsetConfigChangeContext ensures that no value is present for ConfigChangeContext, not even an explicit nil
func (o *PolicyAbstractConfigChangeDetail) UnsetConfigChangeContext() {
	o.ConfigChangeContext.Unset()
}

// GetConfigChangeFlag returns the ConfigChangeFlag field value if set, zero value otherwise.
func (o *PolicyAbstractConfigChangeDetail) GetConfigChangeFlag() string {
	if o == nil || IsNil(o.ConfigChangeFlag) {
		var ret string
		return ret
	}
	return *o.ConfigChangeFlag
}

// GetConfigChangeFlagOk returns a tuple with the ConfigChangeFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigChangeDetail) GetConfigChangeFlagOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigChangeFlag) {
		return nil, false
	}
	return o.ConfigChangeFlag, true
}

// HasConfigChangeFlag returns a boolean if a field has been set.
func (o *PolicyAbstractConfigChangeDetail) HasConfigChangeFlag() bool {
	if o != nil && !IsNil(o.ConfigChangeFlag) {
		return true
	}

	return false
}

// SetConfigChangeFlag gets a reference to the given string and assigns it to the ConfigChangeFlag field.
func (o *PolicyAbstractConfigChangeDetail) SetConfigChangeFlag(v string) {
	o.ConfigChangeFlag = &v
}

// GetDisruptions returns the Disruptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAbstractConfigChangeDetail) GetDisruptions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Disruptions
}

// GetDisruptionsOk returns a tuple with the Disruptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAbstractConfigChangeDetail) GetDisruptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Disruptions) {
		return nil, false
	}
	return o.Disruptions, true
}

// HasDisruptions returns a boolean if a field has been set.
func (o *PolicyAbstractConfigChangeDetail) HasDisruptions() bool {
	if o != nil && !IsNil(o.Disruptions) {
		return true
	}

	return false
}

// SetDisruptions gets a reference to the given []string and assigns it to the Disruptions field.
func (o *PolicyAbstractConfigChangeDetail) SetDisruptions(v []string) {
	o.Disruptions = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PolicyAbstractConfigChangeDetail) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigChangeDetail) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PolicyAbstractConfigChangeDetail) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PolicyAbstractConfigChangeDetail) SetMessage(v string) {
	o.Message = &v
}

// GetModStatus returns the ModStatus field value if set, zero value otherwise.
func (o *PolicyAbstractConfigChangeDetail) GetModStatus() string {
	if o == nil || IsNil(o.ModStatus) {
		var ret string
		return ret
	}
	return *o.ModStatus
}

// GetModStatusOk returns a tuple with the ModStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigChangeDetail) GetModStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ModStatus) {
		return nil, false
	}
	return o.ModStatus, true
}

// HasModStatus returns a boolean if a field has been set.
func (o *PolicyAbstractConfigChangeDetail) HasModStatus() bool {
	if o != nil && !IsNil(o.ModStatus) {
		return true
	}

	return false
}

// SetModStatus gets a reference to the given string and assigns it to the ModStatus field.
func (o *PolicyAbstractConfigChangeDetail) SetModStatus(v string) {
	o.ModStatus = &v
}

func (o PolicyAbstractConfigChangeDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyAbstractConfigChangeDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if o.Changes != nil {
		toSerialize["Changes"] = o.Changes
	}
	if o.ConfigChangeContext.IsSet() {
		toSerialize["ConfigChangeContext"] = o.ConfigChangeContext.Get()
	}
	if !IsNil(o.ConfigChangeFlag) {
		toSerialize["ConfigChangeFlag"] = o.ConfigChangeFlag
	}
	if o.Disruptions != nil {
		toSerialize["Disruptions"] = o.Disruptions
	}
	if !IsNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if !IsNil(o.ModStatus) {
		toSerialize["ModStatus"] = o.ModStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyAbstractConfigChangeDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type PolicyAbstractConfigChangeDetailWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType          string                            `json:"ObjectType"`
		Changes             []string                          `json:"Changes,omitempty"`
		ConfigChangeContext NullablePolicyConfigResultContext `json:"ConfigChangeContext,omitempty"`
		// Config change flag to differentiate Pending-changes and Config-drift. * `Pending-changes` - Config change flag represents changes are due to not deployed changes from Intersight. * `Drift-changes` - Config change flag represents changes are due to endpoint configuration changes.
		ConfigChangeFlag *string  `json:"ConfigChangeFlag,omitempty"`
		Disruptions      []string `json:"Disruptions,omitempty"`
		// Detailed description of the config change.
		Message *string `json:"Message,omitempty"`
		// Modification status of the mo in this config change. * `None` - The 'none' operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration. * `Created` - The 'create' operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet. * `Modified` - The 'update' operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed successfully. * `Deleted` - The 'delete' operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet.
		ModStatus *string `json:"ModStatus,omitempty"`
	}

	varPolicyAbstractConfigChangeDetailWithoutEmbeddedStruct := PolicyAbstractConfigChangeDetailWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPolicyAbstractConfigChangeDetailWithoutEmbeddedStruct)
	if err == nil {
		varPolicyAbstractConfigChangeDetail := _PolicyAbstractConfigChangeDetail{}
		varPolicyAbstractConfigChangeDetail.ClassId = varPolicyAbstractConfigChangeDetailWithoutEmbeddedStruct.ClassId
		varPolicyAbstractConfigChangeDetail.ObjectType = varPolicyAbstractConfigChangeDetailWithoutEmbeddedStruct.ObjectType
		varPolicyAbstractConfigChangeDetail.Changes = varPolicyAbstractConfigChangeDetailWithoutEmbeddedStruct.Changes
		varPolicyAbstractConfigChangeDetail.ConfigChangeContext = varPolicyAbstractConfigChangeDetailWithoutEmbeddedStruct.ConfigChangeContext
		varPolicyAbstractConfigChangeDetail.ConfigChangeFlag = varPolicyAbstractConfigChangeDetailWithoutEmbeddedStruct.ConfigChangeFlag
		varPolicyAbstractConfigChangeDetail.Disruptions = varPolicyAbstractConfigChangeDetailWithoutEmbeddedStruct.Disruptions
		varPolicyAbstractConfigChangeDetail.Message = varPolicyAbstractConfigChangeDetailWithoutEmbeddedStruct.Message
		varPolicyAbstractConfigChangeDetail.ModStatus = varPolicyAbstractConfigChangeDetailWithoutEmbeddedStruct.ModStatus
		*o = PolicyAbstractConfigChangeDetail(varPolicyAbstractConfigChangeDetail)
	} else {
		return err
	}

	varPolicyAbstractConfigChangeDetail := _PolicyAbstractConfigChangeDetail{}

	err = json.Unmarshal(data, &varPolicyAbstractConfigChangeDetail)
	if err == nil {
		o.MoBaseMo = varPolicyAbstractConfigChangeDetail.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Changes")
		delete(additionalProperties, "ConfigChangeContext")
		delete(additionalProperties, "ConfigChangeFlag")
		delete(additionalProperties, "Disruptions")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "ModStatus")

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

type NullablePolicyAbstractConfigChangeDetail struct {
	value *PolicyAbstractConfigChangeDetail
	isSet bool
}

func (v NullablePolicyAbstractConfigChangeDetail) Get() *PolicyAbstractConfigChangeDetail {
	return v.value
}

func (v *NullablePolicyAbstractConfigChangeDetail) Set(val *PolicyAbstractConfigChangeDetail) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAbstractConfigChangeDetail) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAbstractConfigChangeDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAbstractConfigChangeDetail(val *PolicyAbstractConfigChangeDetail) *NullablePolicyAbstractConfigChangeDetail {
	return &NullablePolicyAbstractConfigChangeDetail{value: val, isSet: true}
}

func (v NullablePolicyAbstractConfigChangeDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAbstractConfigChangeDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
