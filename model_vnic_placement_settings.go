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

// checks if the VnicPlacementSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicPlacementSettings{}

// VnicPlacementSettings Placement Settings for the virtual interface.
type VnicPlacementSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable or disable automatic assignment of the PCI Link in a dual-link adapter. This option applies only to 13xx series VICs that support dual-link. If enabled, the system determines the placement of the vNIC/vHBA on either of the PCI Links.
	AutoPciLink *bool `json:"AutoPciLink,omitempty"`
	// Enable or disable automatic assignment of the VIC slot ID. If enabled and the server has only one VIC, the same VIC is chosen for all the vNICs. If enabled and the server has multiple VICs, the vNIC/vHBA are deployed on the first VIC. The Slot ID determines the first VIC. MLOM is the first Slot ID and the ID increments to 2, 3, and so on.
	AutoSlotId *bool `json:"AutoSlotId,omitempty"`
	// PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM.
	Id *string `json:"Id,omitempty" validate:"regexp=^$|^([1-9]|1[0-5]|MLOM)$"`
	// The PCI Link used as transport for the virtual interface. PCI Link is only applicable for select Cisco UCS VIC 1300 models (UCSC-PCIE-C40Q-03, UCSB-MLOM-40G-03, UCSB-VIC-M83-8P) that support two PCI links. The value, if specified, for any other VIC model will be ignored.
	PciLink *int64 `json:"PciLink,omitempty"`
	// If the autoPciLink is disabled, the user can either choose to place the vNICs manually or based on a policy.If the autoPciLink is enabled, it will be set to None. * `Custom` - The user needs to specify the PCI Link manually. * `Load-Balanced` - The system will uniformly distribute the interfaces across the PCI Links. * `None` - Assignment is not applicable and will be set when the AutoPciLink is set to true.
	PciLinkAssignmentMode *string `json:"PciLinkAssignmentMode,omitempty"`
	// The fabric port to which the vNICs will be associated. * `None` - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used. * `A` - Fabric A of the FI cluster. * `B` - Fabric B of the FI cluster.
	SwitchId *string `json:"SwitchId,omitempty"`
	// Adapter port on which the virtual interface will be created.
	Uplink               *int64 `json:"Uplink,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicPlacementSettings VnicPlacementSettings

// NewVnicPlacementSettings instantiates a new VnicPlacementSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicPlacementSettings(classId string, objectType string) *VnicPlacementSettings {
	this := VnicPlacementSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var autoPciLink bool = false
	this.AutoPciLink = &autoPciLink
	var autoSlotId bool = false
	this.AutoSlotId = &autoSlotId
	var pciLink int64 = 0
	this.PciLink = &pciLink
	var pciLinkAssignmentMode string = "Custom"
	this.PciLinkAssignmentMode = &pciLinkAssignmentMode
	var switchId string = "None"
	this.SwitchId = &switchId
	return &this
}

// NewVnicPlacementSettingsWithDefaults instantiates a new VnicPlacementSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicPlacementSettingsWithDefaults() *VnicPlacementSettings {
	this := VnicPlacementSettings{}
	var classId string = "vnic.PlacementSettings"
	this.ClassId = classId
	var objectType string = "vnic.PlacementSettings"
	this.ObjectType = objectType
	var autoPciLink bool = false
	this.AutoPciLink = &autoPciLink
	var autoSlotId bool = false
	this.AutoSlotId = &autoSlotId
	var pciLink int64 = 0
	this.PciLink = &pciLink
	var pciLinkAssignmentMode string = "Custom"
	this.PciLinkAssignmentMode = &pciLinkAssignmentMode
	var switchId string = "None"
	this.SwitchId = &switchId
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicPlacementSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicPlacementSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.PlacementSettings" of the ClassId field.
func (o *VnicPlacementSettings) GetDefaultClassId() interface{} {
	return "vnic.PlacementSettings"
}

// GetObjectType returns the ObjectType field value
func (o *VnicPlacementSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicPlacementSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.PlacementSettings" of the ObjectType field.
func (o *VnicPlacementSettings) GetDefaultObjectType() interface{} {
	return "vnic.PlacementSettings"
}

// GetAutoPciLink returns the AutoPciLink field value if set, zero value otherwise.
func (o *VnicPlacementSettings) GetAutoPciLink() bool {
	if o == nil || IsNil(o.AutoPciLink) {
		var ret bool
		return ret
	}
	return *o.AutoPciLink
}

// GetAutoPciLinkOk returns a tuple with the AutoPciLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetAutoPciLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoPciLink) {
		return nil, false
	}
	return o.AutoPciLink, true
}

// HasAutoPciLink returns a boolean if a field has been set.
func (o *VnicPlacementSettings) HasAutoPciLink() bool {
	if o != nil && !IsNil(o.AutoPciLink) {
		return true
	}

	return false
}

// SetAutoPciLink gets a reference to the given bool and assigns it to the AutoPciLink field.
func (o *VnicPlacementSettings) SetAutoPciLink(v bool) {
	o.AutoPciLink = &v
}

// GetAutoSlotId returns the AutoSlotId field value if set, zero value otherwise.
func (o *VnicPlacementSettings) GetAutoSlotId() bool {
	if o == nil || IsNil(o.AutoSlotId) {
		var ret bool
		return ret
	}
	return *o.AutoSlotId
}

// GetAutoSlotIdOk returns a tuple with the AutoSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetAutoSlotIdOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoSlotId) {
		return nil, false
	}
	return o.AutoSlotId, true
}

// HasAutoSlotId returns a boolean if a field has been set.
func (o *VnicPlacementSettings) HasAutoSlotId() bool {
	if o != nil && !IsNil(o.AutoSlotId) {
		return true
	}

	return false
}

// SetAutoSlotId gets a reference to the given bool and assigns it to the AutoSlotId field.
func (o *VnicPlacementSettings) SetAutoSlotId(v bool) {
	o.AutoSlotId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VnicPlacementSettings) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VnicPlacementSettings) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VnicPlacementSettings) SetId(v string) {
	o.Id = &v
}

// GetPciLink returns the PciLink field value if set, zero value otherwise.
func (o *VnicPlacementSettings) GetPciLink() int64 {
	if o == nil || IsNil(o.PciLink) {
		var ret int64
		return ret
	}
	return *o.PciLink
}

// GetPciLinkOk returns a tuple with the PciLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetPciLinkOk() (*int64, bool) {
	if o == nil || IsNil(o.PciLink) {
		return nil, false
	}
	return o.PciLink, true
}

// HasPciLink returns a boolean if a field has been set.
func (o *VnicPlacementSettings) HasPciLink() bool {
	if o != nil && !IsNil(o.PciLink) {
		return true
	}

	return false
}

// SetPciLink gets a reference to the given int64 and assigns it to the PciLink field.
func (o *VnicPlacementSettings) SetPciLink(v int64) {
	o.PciLink = &v
}

// GetPciLinkAssignmentMode returns the PciLinkAssignmentMode field value if set, zero value otherwise.
func (o *VnicPlacementSettings) GetPciLinkAssignmentMode() string {
	if o == nil || IsNil(o.PciLinkAssignmentMode) {
		var ret string
		return ret
	}
	return *o.PciLinkAssignmentMode
}

// GetPciLinkAssignmentModeOk returns a tuple with the PciLinkAssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetPciLinkAssignmentModeOk() (*string, bool) {
	if o == nil || IsNil(o.PciLinkAssignmentMode) {
		return nil, false
	}
	return o.PciLinkAssignmentMode, true
}

// HasPciLinkAssignmentMode returns a boolean if a field has been set.
func (o *VnicPlacementSettings) HasPciLinkAssignmentMode() bool {
	if o != nil && !IsNil(o.PciLinkAssignmentMode) {
		return true
	}

	return false
}

// SetPciLinkAssignmentMode gets a reference to the given string and assigns it to the PciLinkAssignmentMode field.
func (o *VnicPlacementSettings) SetPciLinkAssignmentMode(v string) {
	o.PciLinkAssignmentMode = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *VnicPlacementSettings) GetSwitchId() string {
	if o == nil || IsNil(o.SwitchId) {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetSwitchIdOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchId) {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *VnicPlacementSettings) HasSwitchId() bool {
	if o != nil && !IsNil(o.SwitchId) {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *VnicPlacementSettings) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetUplink returns the Uplink field value if set, zero value otherwise.
func (o *VnicPlacementSettings) GetUplink() int64 {
	if o == nil || IsNil(o.Uplink) {
		var ret int64
		return ret
	}
	return *o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetUplinkOk() (*int64, bool) {
	if o == nil || IsNil(o.Uplink) {
		return nil, false
	}
	return o.Uplink, true
}

// HasUplink returns a boolean if a field has been set.
func (o *VnicPlacementSettings) HasUplink() bool {
	if o != nil && !IsNil(o.Uplink) {
		return true
	}

	return false
}

// SetUplink gets a reference to the given int64 and assigns it to the Uplink field.
func (o *VnicPlacementSettings) SetUplink(v int64) {
	o.Uplink = &v
}

func (o VnicPlacementSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicPlacementSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AutoPciLink) {
		toSerialize["AutoPciLink"] = o.AutoPciLink
	}
	if !IsNil(o.AutoSlotId) {
		toSerialize["AutoSlotId"] = o.AutoSlotId
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.PciLink) {
		toSerialize["PciLink"] = o.PciLink
	}
	if !IsNil(o.PciLinkAssignmentMode) {
		toSerialize["PciLinkAssignmentMode"] = o.PciLinkAssignmentMode
	}
	if !IsNil(o.SwitchId) {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if !IsNil(o.Uplink) {
		toSerialize["Uplink"] = o.Uplink
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicPlacementSettings) UnmarshalJSON(data []byte) (err error) {
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
	type VnicPlacementSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enable or disable automatic assignment of the PCI Link in a dual-link adapter. This option applies only to 13xx series VICs that support dual-link. If enabled, the system determines the placement of the vNIC/vHBA on either of the PCI Links.
		AutoPciLink *bool `json:"AutoPciLink,omitempty"`
		// Enable or disable automatic assignment of the VIC slot ID. If enabled and the server has only one VIC, the same VIC is chosen for all the vNICs. If enabled and the server has multiple VICs, the vNIC/vHBA are deployed on the first VIC. The Slot ID determines the first VIC. MLOM is the first Slot ID and the ID increments to 2, 3, and so on.
		AutoSlotId *bool `json:"AutoSlotId,omitempty"`
		// PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM.
		Id *string `json:"Id,omitempty" validate:"regexp=^$|^([1-9]|1[0-5]|MLOM)$"`
		// The PCI Link used as transport for the virtual interface. PCI Link is only applicable for select Cisco UCS VIC 1300 models (UCSC-PCIE-C40Q-03, UCSB-MLOM-40G-03, UCSB-VIC-M83-8P) that support two PCI links. The value, if specified, for any other VIC model will be ignored.
		PciLink *int64 `json:"PciLink,omitempty"`
		// If the autoPciLink is disabled, the user can either choose to place the vNICs manually or based on a policy.If the autoPciLink is enabled, it will be set to None. * `Custom` - The user needs to specify the PCI Link manually. * `Load-Balanced` - The system will uniformly distribute the interfaces across the PCI Links. * `None` - Assignment is not applicable and will be set when the AutoPciLink is set to true.
		PciLinkAssignmentMode *string `json:"PciLinkAssignmentMode,omitempty"`
		// The fabric port to which the vNICs will be associated. * `None` - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used. * `A` - Fabric A of the FI cluster. * `B` - Fabric B of the FI cluster.
		SwitchId *string `json:"SwitchId,omitempty"`
		// Adapter port on which the virtual interface will be created.
		Uplink *int64 `json:"Uplink,omitempty"`
	}

	varVnicPlacementSettingsWithoutEmbeddedStruct := VnicPlacementSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicPlacementSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicPlacementSettings := _VnicPlacementSettings{}
		varVnicPlacementSettings.ClassId = varVnicPlacementSettingsWithoutEmbeddedStruct.ClassId
		varVnicPlacementSettings.ObjectType = varVnicPlacementSettingsWithoutEmbeddedStruct.ObjectType
		varVnicPlacementSettings.AutoPciLink = varVnicPlacementSettingsWithoutEmbeddedStruct.AutoPciLink
		varVnicPlacementSettings.AutoSlotId = varVnicPlacementSettingsWithoutEmbeddedStruct.AutoSlotId
		varVnicPlacementSettings.Id = varVnicPlacementSettingsWithoutEmbeddedStruct.Id
		varVnicPlacementSettings.PciLink = varVnicPlacementSettingsWithoutEmbeddedStruct.PciLink
		varVnicPlacementSettings.PciLinkAssignmentMode = varVnicPlacementSettingsWithoutEmbeddedStruct.PciLinkAssignmentMode
		varVnicPlacementSettings.SwitchId = varVnicPlacementSettingsWithoutEmbeddedStruct.SwitchId
		varVnicPlacementSettings.Uplink = varVnicPlacementSettingsWithoutEmbeddedStruct.Uplink
		*o = VnicPlacementSettings(varVnicPlacementSettings)
	} else {
		return err
	}

	varVnicPlacementSettings := _VnicPlacementSettings{}

	err = json.Unmarshal(data, &varVnicPlacementSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicPlacementSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutoPciLink")
		delete(additionalProperties, "AutoSlotId")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "PciLink")
		delete(additionalProperties, "PciLinkAssignmentMode")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "Uplink")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableVnicPlacementSettings struct {
	value *VnicPlacementSettings
	isSet bool
}

func (v NullableVnicPlacementSettings) Get() *VnicPlacementSettings {
	return v.value
}

func (v *NullableVnicPlacementSettings) Set(val *VnicPlacementSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicPlacementSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicPlacementSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicPlacementSettings(val *VnicPlacementSettings) *NullableVnicPlacementSettings {
	return &NullableVnicPlacementSettings{value: val, isSet: true}
}

func (v NullableVnicPlacementSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicPlacementSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
