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

// checks if the HclHyperflexSoftwareCompatibilityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HclHyperflexSoftwareCompatibilityInfo{}

// HclHyperflexSoftwareCompatibilityInfo Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.
type HclHyperflexSoftwareCompatibilityInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string          `json:"ObjectType"`
	Constraints []HclConstraint `json:"Constraints,omitempty"`
	// HXDP component software version.
	HxdpVersion *string `json:"HxdpVersion,omitempty"`
	// Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// Hypervisor component software version.
	HypervisorVersion *string `json:"HypervisorVersion,omitempty"`
	// Type of the HXDP bundle mgmt or full.
	IsMgmtBuild *string `json:"IsMgmtBuild,omitempty"`
	// Maximum supported HyperFlex Data Platform build version.
	MaxMgmtVersion *string `json:"MaxMgmtVersion,omitempty"`
	// Minimum supported HyperFlex Data Platform build version.
	MinMgmtVersion *string `json:"MinMgmtVersion,omitempty"`
	// UCS Server Firmware component software version.
	ServerFwVersion      *string                                 `json:"ServerFwVersion,omitempty"`
	AppCatalog           NullableHyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclHyperflexSoftwareCompatibilityInfo HclHyperflexSoftwareCompatibilityInfo

// NewHclHyperflexSoftwareCompatibilityInfo instantiates a new HclHyperflexSoftwareCompatibilityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclHyperflexSoftwareCompatibilityInfo(classId string, objectType string) *HclHyperflexSoftwareCompatibilityInfo {
	this := HclHyperflexSoftwareCompatibilityInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// NewHclHyperflexSoftwareCompatibilityInfoWithDefaults instantiates a new HclHyperflexSoftwareCompatibilityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclHyperflexSoftwareCompatibilityInfoWithDefaults() *HclHyperflexSoftwareCompatibilityInfo {
	this := HclHyperflexSoftwareCompatibilityInfo{}
	var classId string = "hcl.HyperflexSoftwareCompatibilityInfo"
	this.ClassId = classId
	var objectType string = "hcl.HyperflexSoftwareCompatibilityInfo"
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclHyperflexSoftwareCompatibilityInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclHyperflexSoftwareCompatibilityInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hcl.HyperflexSoftwareCompatibilityInfo" of the ClassId field.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetDefaultClassId() interface{} {
	return "hcl.HyperflexSoftwareCompatibilityInfo"
}

// GetObjectType returns the ObjectType field value
func (o *HclHyperflexSoftwareCompatibilityInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclHyperflexSoftwareCompatibilityInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hcl.HyperflexSoftwareCompatibilityInfo" of the ObjectType field.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetDefaultObjectType() interface{} {
	return "hcl.HyperflexSoftwareCompatibilityInfo"
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclHyperflexSoftwareCompatibilityInfo) GetConstraints() []HclConstraint {
	if o == nil {
		var ret []HclConstraint
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclHyperflexSoftwareCompatibilityInfo) GetConstraintsOk() ([]HclConstraint, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []HclConstraint and assigns it to the Constraints field.
func (o *HclHyperflexSoftwareCompatibilityInfo) SetConstraints(v []HclConstraint) {
	o.Constraints = v
}

// GetHxdpVersion returns the HxdpVersion field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetHxdpVersion() string {
	if o == nil || IsNil(o.HxdpVersion) {
		var ret string
		return ret
	}
	return *o.HxdpVersion
}

// GetHxdpVersionOk returns a tuple with the HxdpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetHxdpVersionOk() (*string, bool) {
	if o == nil || IsNil(o.HxdpVersion) {
		return nil, false
	}
	return o.HxdpVersion, true
}

// HasHxdpVersion returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) HasHxdpVersion() bool {
	if o != nil && !IsNil(o.HxdpVersion) {
		return true
	}

	return false
}

// SetHxdpVersion gets a reference to the given string and assigns it to the HxdpVersion field.
func (o *HclHyperflexSoftwareCompatibilityInfo) SetHxdpVersion(v string) {
	o.HxdpVersion = &v
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetHypervisorType() string {
	if o == nil || IsNil(o.HypervisorType) {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.HypervisorType) {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) HasHypervisorType() bool {
	if o != nil && !IsNil(o.HypervisorType) {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *HclHyperflexSoftwareCompatibilityInfo) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetHypervisorVersion returns the HypervisorVersion field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetHypervisorVersion() string {
	if o == nil || IsNil(o.HypervisorVersion) {
		var ret string
		return ret
	}
	return *o.HypervisorVersion
}

// GetHypervisorVersionOk returns a tuple with the HypervisorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetHypervisorVersionOk() (*string, bool) {
	if o == nil || IsNil(o.HypervisorVersion) {
		return nil, false
	}
	return o.HypervisorVersion, true
}

// HasHypervisorVersion returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) HasHypervisorVersion() bool {
	if o != nil && !IsNil(o.HypervisorVersion) {
		return true
	}

	return false
}

// SetHypervisorVersion gets a reference to the given string and assigns it to the HypervisorVersion field.
func (o *HclHyperflexSoftwareCompatibilityInfo) SetHypervisorVersion(v string) {
	o.HypervisorVersion = &v
}

// GetIsMgmtBuild returns the IsMgmtBuild field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetIsMgmtBuild() string {
	if o == nil || IsNil(o.IsMgmtBuild) {
		var ret string
		return ret
	}
	return *o.IsMgmtBuild
}

// GetIsMgmtBuildOk returns a tuple with the IsMgmtBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetIsMgmtBuildOk() (*string, bool) {
	if o == nil || IsNil(o.IsMgmtBuild) {
		return nil, false
	}
	return o.IsMgmtBuild, true
}

// HasIsMgmtBuild returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) HasIsMgmtBuild() bool {
	if o != nil && !IsNil(o.IsMgmtBuild) {
		return true
	}

	return false
}

// SetIsMgmtBuild gets a reference to the given string and assigns it to the IsMgmtBuild field.
func (o *HclHyperflexSoftwareCompatibilityInfo) SetIsMgmtBuild(v string) {
	o.IsMgmtBuild = &v
}

// GetMaxMgmtVersion returns the MaxMgmtVersion field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetMaxMgmtVersion() string {
	if o == nil || IsNil(o.MaxMgmtVersion) {
		var ret string
		return ret
	}
	return *o.MaxMgmtVersion
}

// GetMaxMgmtVersionOk returns a tuple with the MaxMgmtVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetMaxMgmtVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MaxMgmtVersion) {
		return nil, false
	}
	return o.MaxMgmtVersion, true
}

// HasMaxMgmtVersion returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) HasMaxMgmtVersion() bool {
	if o != nil && !IsNil(o.MaxMgmtVersion) {
		return true
	}

	return false
}

// SetMaxMgmtVersion gets a reference to the given string and assigns it to the MaxMgmtVersion field.
func (o *HclHyperflexSoftwareCompatibilityInfo) SetMaxMgmtVersion(v string) {
	o.MaxMgmtVersion = &v
}

// GetMinMgmtVersion returns the MinMgmtVersion field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetMinMgmtVersion() string {
	if o == nil || IsNil(o.MinMgmtVersion) {
		var ret string
		return ret
	}
	return *o.MinMgmtVersion
}

// GetMinMgmtVersionOk returns a tuple with the MinMgmtVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetMinMgmtVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinMgmtVersion) {
		return nil, false
	}
	return o.MinMgmtVersion, true
}

// HasMinMgmtVersion returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) HasMinMgmtVersion() bool {
	if o != nil && !IsNil(o.MinMgmtVersion) {
		return true
	}

	return false
}

// SetMinMgmtVersion gets a reference to the given string and assigns it to the MinMgmtVersion field.
func (o *HclHyperflexSoftwareCompatibilityInfo) SetMinMgmtVersion(v string) {
	o.MinMgmtVersion = &v
}

// GetServerFwVersion returns the ServerFwVersion field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetServerFwVersion() string {
	if o == nil || IsNil(o.ServerFwVersion) {
		var ret string
		return ret
	}
	return *o.ServerFwVersion
}

// GetServerFwVersionOk returns a tuple with the ServerFwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) GetServerFwVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFwVersion) {
		return nil, false
	}
	return o.ServerFwVersion, true
}

// HasServerFwVersion returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) HasServerFwVersion() bool {
	if o != nil && !IsNil(o.ServerFwVersion) {
		return true
	}

	return false
}

// SetServerFwVersion gets a reference to the given string and assigns it to the ServerFwVersion field.
func (o *HclHyperflexSoftwareCompatibilityInfo) SetServerFwVersion(v string) {
	o.ServerFwVersion = &v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclHyperflexSoftwareCompatibilityInfo) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || IsNil(o.AppCatalog.Get()) {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog.Get()
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclHyperflexSoftwareCompatibilityInfo) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppCatalog.Get(), o.AppCatalog.IsSet()
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfo) HasAppCatalog() bool {
	if o != nil && o.AppCatalog.IsSet() {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given NullableHyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HclHyperflexSoftwareCompatibilityInfo) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog.Set(&v)
}

// SetAppCatalogNil sets the value for AppCatalog to be an explicit nil
func (o *HclHyperflexSoftwareCompatibilityInfo) SetAppCatalogNil() {
	o.AppCatalog.Set(nil)
}

// UnsetAppCatalog ensures that no value is present for AppCatalog, not even an explicit nil
func (o *HclHyperflexSoftwareCompatibilityInfo) UnsetAppCatalog() {
	o.AppCatalog.Unset()
}

func (o HclHyperflexSoftwareCompatibilityInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HclHyperflexSoftwareCompatibilityInfo) ToMap() (map[string]interface{}, error) {
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
	if o.Constraints != nil {
		toSerialize["Constraints"] = o.Constraints
	}
	if !IsNil(o.HxdpVersion) {
		toSerialize["HxdpVersion"] = o.HxdpVersion
	}
	if !IsNil(o.HypervisorType) {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if !IsNil(o.HypervisorVersion) {
		toSerialize["HypervisorVersion"] = o.HypervisorVersion
	}
	if !IsNil(o.IsMgmtBuild) {
		toSerialize["IsMgmtBuild"] = o.IsMgmtBuild
	}
	if !IsNil(o.MaxMgmtVersion) {
		toSerialize["MaxMgmtVersion"] = o.MaxMgmtVersion
	}
	if !IsNil(o.MinMgmtVersion) {
		toSerialize["MinMgmtVersion"] = o.MinMgmtVersion
	}
	if !IsNil(o.ServerFwVersion) {
		toSerialize["ServerFwVersion"] = o.ServerFwVersion
	}
	if o.AppCatalog.IsSet() {
		toSerialize["AppCatalog"] = o.AppCatalog.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HclHyperflexSoftwareCompatibilityInfo) UnmarshalJSON(data []byte) (err error) {
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
	type HclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string          `json:"ObjectType"`
		Constraints []HclConstraint `json:"Constraints,omitempty"`
		// HXDP component software version.
		HxdpVersion *string `json:"HxdpVersion,omitempty"`
		// Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
		HypervisorType *string `json:"HypervisorType,omitempty"`
		// Hypervisor component software version.
		HypervisorVersion *string `json:"HypervisorVersion,omitempty"`
		// Type of the HXDP bundle mgmt or full.
		IsMgmtBuild *string `json:"IsMgmtBuild,omitempty"`
		// Maximum supported HyperFlex Data Platform build version.
		MaxMgmtVersion *string `json:"MaxMgmtVersion,omitempty"`
		// Minimum supported HyperFlex Data Platform build version.
		MinMgmtVersion *string `json:"MinMgmtVersion,omitempty"`
		// UCS Server Firmware component software version.
		ServerFwVersion *string                                 `json:"ServerFwVersion,omitempty"`
		AppCatalog      NullableHyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	}

	varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct := HclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct)
	if err == nil {
		varHclHyperflexSoftwareCompatibilityInfo := _HclHyperflexSoftwareCompatibilityInfo{}
		varHclHyperflexSoftwareCompatibilityInfo.ClassId = varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct.ClassId
		varHclHyperflexSoftwareCompatibilityInfo.ObjectType = varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct.ObjectType
		varHclHyperflexSoftwareCompatibilityInfo.Constraints = varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct.Constraints
		varHclHyperflexSoftwareCompatibilityInfo.HxdpVersion = varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct.HxdpVersion
		varHclHyperflexSoftwareCompatibilityInfo.HypervisorType = varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct.HypervisorType
		varHclHyperflexSoftwareCompatibilityInfo.HypervisorVersion = varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct.HypervisorVersion
		varHclHyperflexSoftwareCompatibilityInfo.IsMgmtBuild = varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct.IsMgmtBuild
		varHclHyperflexSoftwareCompatibilityInfo.MaxMgmtVersion = varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct.MaxMgmtVersion
		varHclHyperflexSoftwareCompatibilityInfo.MinMgmtVersion = varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct.MinMgmtVersion
		varHclHyperflexSoftwareCompatibilityInfo.ServerFwVersion = varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct.ServerFwVersion
		varHclHyperflexSoftwareCompatibilityInfo.AppCatalog = varHclHyperflexSoftwareCompatibilityInfoWithoutEmbeddedStruct.AppCatalog
		*o = HclHyperflexSoftwareCompatibilityInfo(varHclHyperflexSoftwareCompatibilityInfo)
	} else {
		return err
	}

	varHclHyperflexSoftwareCompatibilityInfo := _HclHyperflexSoftwareCompatibilityInfo{}

	err = json.Unmarshal(data, &varHclHyperflexSoftwareCompatibilityInfo)
	if err == nil {
		o.MoBaseMo = varHclHyperflexSoftwareCompatibilityInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Constraints")
		delete(additionalProperties, "HxdpVersion")
		delete(additionalProperties, "HypervisorType")
		delete(additionalProperties, "HypervisorVersion")
		delete(additionalProperties, "IsMgmtBuild")
		delete(additionalProperties, "MaxMgmtVersion")
		delete(additionalProperties, "MinMgmtVersion")
		delete(additionalProperties, "ServerFwVersion")
		delete(additionalProperties, "AppCatalog")

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

type NullableHclHyperflexSoftwareCompatibilityInfo struct {
	value *HclHyperflexSoftwareCompatibilityInfo
	isSet bool
}

func (v NullableHclHyperflexSoftwareCompatibilityInfo) Get() *HclHyperflexSoftwareCompatibilityInfo {
	return v.value
}

func (v *NullableHclHyperflexSoftwareCompatibilityInfo) Set(val *HclHyperflexSoftwareCompatibilityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHclHyperflexSoftwareCompatibilityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHclHyperflexSoftwareCompatibilityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclHyperflexSoftwareCompatibilityInfo(val *HclHyperflexSoftwareCompatibilityInfo) *NullableHclHyperflexSoftwareCompatibilityInfo {
	return &NullableHclHyperflexSoftwareCompatibilityInfo{value: val, isSet: true}
}

func (v NullableHclHyperflexSoftwareCompatibilityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclHyperflexSoftwareCompatibilityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
