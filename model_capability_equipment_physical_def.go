/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// CapabilityEquipmentPhysicalDef Type to represent additional switch specific capabilities.
type CapabilityEquipmentPhysicalDef struct {
	CapabilitySwitchCapabilityDef
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Depth information for a Switch/Fabric-Interconnect.
	Depth *float32 `json:"Depth,omitempty"`
	// Height information for a Switch/Fabric-Interconnect.
	Height *float32 `json:"Height,omitempty"`
	// Max Power information for a Switch/Fabric-Interconnect.
	MaxPower *int64 `json:"MaxPower,omitempty"`
	// Min Power information for a Switch/Fabric-Interconnect.
	MinPower *int64 `json:"MinPower,omitempty"`
	// Nominal Power information for a Switch/Fabric-Interconnect.
	NominalPower *int64 `json:"NominalPower,omitempty"`
	// Weight information for a Switch/Fabric-Interconnect.
	Weight *float32 `json:"Weight,omitempty"`
	// Width information for a Switch/Fabric-Interconnect.
	Width                *float32 `json:"Width,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityEquipmentPhysicalDef CapabilityEquipmentPhysicalDef

// NewCapabilityEquipmentPhysicalDef instantiates a new CapabilityEquipmentPhysicalDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityEquipmentPhysicalDef(classId string, objectType string) *CapabilityEquipmentPhysicalDef {
	this := CapabilityEquipmentPhysicalDef{}
	this.ClassId = classId
	this.ObjectType = objectType
	var pid string = "UCS-FI-6454"
	this.Pid = &pid
	return &this
}

// NewCapabilityEquipmentPhysicalDefWithDefaults instantiates a new CapabilityEquipmentPhysicalDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityEquipmentPhysicalDefWithDefaults() *CapabilityEquipmentPhysicalDef {
	this := CapabilityEquipmentPhysicalDef{}
	var classId string = "capability.EquipmentPhysicalDef"
	this.ClassId = classId
	var objectType string = "capability.EquipmentPhysicalDef"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityEquipmentPhysicalDef) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDef) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityEquipmentPhysicalDef) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityEquipmentPhysicalDef) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDef) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityEquipmentPhysicalDef) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDef) GetDepth() float32 {
	if o == nil || o.Depth == nil {
		var ret float32
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDef) GetDepthOk() (*float32, bool) {
	if o == nil || o.Depth == nil {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDef) HasDepth() bool {
	if o != nil && o.Depth != nil {
		return true
	}

	return false
}

// SetDepth gets a reference to the given float32 and assigns it to the Depth field.
func (o *CapabilityEquipmentPhysicalDef) SetDepth(v float32) {
	o.Depth = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDef) GetHeight() float32 {
	if o == nil || o.Height == nil {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDef) GetHeightOk() (*float32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDef) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *CapabilityEquipmentPhysicalDef) SetHeight(v float32) {
	o.Height = &v
}

// GetMaxPower returns the MaxPower field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDef) GetMaxPower() int64 {
	if o == nil || o.MaxPower == nil {
		var ret int64
		return ret
	}
	return *o.MaxPower
}

// GetMaxPowerOk returns a tuple with the MaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDef) GetMaxPowerOk() (*int64, bool) {
	if o == nil || o.MaxPower == nil {
		return nil, false
	}
	return o.MaxPower, true
}

// HasMaxPower returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDef) HasMaxPower() bool {
	if o != nil && o.MaxPower != nil {
		return true
	}

	return false
}

// SetMaxPower gets a reference to the given int64 and assigns it to the MaxPower field.
func (o *CapabilityEquipmentPhysicalDef) SetMaxPower(v int64) {
	o.MaxPower = &v
}

// GetMinPower returns the MinPower field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDef) GetMinPower() int64 {
	if o == nil || o.MinPower == nil {
		var ret int64
		return ret
	}
	return *o.MinPower
}

// GetMinPowerOk returns a tuple with the MinPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDef) GetMinPowerOk() (*int64, bool) {
	if o == nil || o.MinPower == nil {
		return nil, false
	}
	return o.MinPower, true
}

// HasMinPower returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDef) HasMinPower() bool {
	if o != nil && o.MinPower != nil {
		return true
	}

	return false
}

// SetMinPower gets a reference to the given int64 and assigns it to the MinPower field.
func (o *CapabilityEquipmentPhysicalDef) SetMinPower(v int64) {
	o.MinPower = &v
}

// GetNominalPower returns the NominalPower field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDef) GetNominalPower() int64 {
	if o == nil || o.NominalPower == nil {
		var ret int64
		return ret
	}
	return *o.NominalPower
}

// GetNominalPowerOk returns a tuple with the NominalPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDef) GetNominalPowerOk() (*int64, bool) {
	if o == nil || o.NominalPower == nil {
		return nil, false
	}
	return o.NominalPower, true
}

// HasNominalPower returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDef) HasNominalPower() bool {
	if o != nil && o.NominalPower != nil {
		return true
	}

	return false
}

// SetNominalPower gets a reference to the given int64 and assigns it to the NominalPower field.
func (o *CapabilityEquipmentPhysicalDef) SetNominalPower(v int64) {
	o.NominalPower = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDef) GetWeight() float32 {
	if o == nil || o.Weight == nil {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDef) GetWeightOk() (*float32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDef) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *CapabilityEquipmentPhysicalDef) SetWeight(v float32) {
	o.Weight = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDef) GetWidth() float32 {
	if o == nil || o.Width == nil {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDef) GetWidthOk() (*float32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDef) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *CapabilityEquipmentPhysicalDef) SetWidth(v float32) {
	o.Width = &v
}

func (o CapabilityEquipmentPhysicalDef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilitySwitchCapabilityDef, errCapabilitySwitchCapabilityDef := json.Marshal(o.CapabilitySwitchCapabilityDef)
	if errCapabilitySwitchCapabilityDef != nil {
		return []byte{}, errCapabilitySwitchCapabilityDef
	}
	errCapabilitySwitchCapabilityDef = json.Unmarshal([]byte(serializedCapabilitySwitchCapabilityDef), &toSerialize)
	if errCapabilitySwitchCapabilityDef != nil {
		return []byte{}, errCapabilitySwitchCapabilityDef
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Depth != nil {
		toSerialize["Depth"] = o.Depth
	}
	if o.Height != nil {
		toSerialize["Height"] = o.Height
	}
	if o.MaxPower != nil {
		toSerialize["MaxPower"] = o.MaxPower
	}
	if o.MinPower != nil {
		toSerialize["MinPower"] = o.MinPower
	}
	if o.NominalPower != nil {
		toSerialize["NominalPower"] = o.NominalPower
	}
	if o.Weight != nil {
		toSerialize["Weight"] = o.Weight
	}
	if o.Width != nil {
		toSerialize["Width"] = o.Width
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityEquipmentPhysicalDef) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilityEquipmentPhysicalDefWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Depth information for a Switch/Fabric-Interconnect.
		Depth *float32 `json:"Depth,omitempty"`
		// Height information for a Switch/Fabric-Interconnect.
		Height *float32 `json:"Height,omitempty"`
		// Max Power information for a Switch/Fabric-Interconnect.
		MaxPower *int64 `json:"MaxPower,omitempty"`
		// Min Power information for a Switch/Fabric-Interconnect.
		MinPower *int64 `json:"MinPower,omitempty"`
		// Nominal Power information for a Switch/Fabric-Interconnect.
		NominalPower *int64 `json:"NominalPower,omitempty"`
		// Weight information for a Switch/Fabric-Interconnect.
		Weight *float32 `json:"Weight,omitempty"`
		// Width information for a Switch/Fabric-Interconnect.
		Width *float32 `json:"Width,omitempty"`
	}

	varCapabilityEquipmentPhysicalDefWithoutEmbeddedStruct := CapabilityEquipmentPhysicalDefWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilityEquipmentPhysicalDefWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityEquipmentPhysicalDef := _CapabilityEquipmentPhysicalDef{}
		varCapabilityEquipmentPhysicalDef.ClassId = varCapabilityEquipmentPhysicalDefWithoutEmbeddedStruct.ClassId
		varCapabilityEquipmentPhysicalDef.ObjectType = varCapabilityEquipmentPhysicalDefWithoutEmbeddedStruct.ObjectType
		varCapabilityEquipmentPhysicalDef.Depth = varCapabilityEquipmentPhysicalDefWithoutEmbeddedStruct.Depth
		varCapabilityEquipmentPhysicalDef.Height = varCapabilityEquipmentPhysicalDefWithoutEmbeddedStruct.Height
		varCapabilityEquipmentPhysicalDef.MaxPower = varCapabilityEquipmentPhysicalDefWithoutEmbeddedStruct.MaxPower
		varCapabilityEquipmentPhysicalDef.MinPower = varCapabilityEquipmentPhysicalDefWithoutEmbeddedStruct.MinPower
		varCapabilityEquipmentPhysicalDef.NominalPower = varCapabilityEquipmentPhysicalDefWithoutEmbeddedStruct.NominalPower
		varCapabilityEquipmentPhysicalDef.Weight = varCapabilityEquipmentPhysicalDefWithoutEmbeddedStruct.Weight
		varCapabilityEquipmentPhysicalDef.Width = varCapabilityEquipmentPhysicalDefWithoutEmbeddedStruct.Width
		*o = CapabilityEquipmentPhysicalDef(varCapabilityEquipmentPhysicalDef)
	} else {
		return err
	}

	varCapabilityEquipmentPhysicalDef := _CapabilityEquipmentPhysicalDef{}

	err = json.Unmarshal(bytes, &varCapabilityEquipmentPhysicalDef)
	if err == nil {
		o.CapabilitySwitchCapabilityDef = varCapabilityEquipmentPhysicalDef.CapabilitySwitchCapabilityDef
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Depth")
		delete(additionalProperties, "Height")
		delete(additionalProperties, "MaxPower")
		delete(additionalProperties, "MinPower")
		delete(additionalProperties, "NominalPower")
		delete(additionalProperties, "Weight")
		delete(additionalProperties, "Width")

		// remove fields from embedded structs
		reflectCapabilitySwitchCapabilityDef := reflect.ValueOf(o.CapabilitySwitchCapabilityDef)
		for i := 0; i < reflectCapabilitySwitchCapabilityDef.Type().NumField(); i++ {
			t := reflectCapabilitySwitchCapabilityDef.Type().Field(i)

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

type NullableCapabilityEquipmentPhysicalDef struct {
	value *CapabilityEquipmentPhysicalDef
	isSet bool
}

func (v NullableCapabilityEquipmentPhysicalDef) Get() *CapabilityEquipmentPhysicalDef {
	return v.value
}

func (v *NullableCapabilityEquipmentPhysicalDef) Set(val *CapabilityEquipmentPhysicalDef) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityEquipmentPhysicalDef) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityEquipmentPhysicalDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityEquipmentPhysicalDef(val *CapabilityEquipmentPhysicalDef) *NullableCapabilityEquipmentPhysicalDef {
	return &NullableCapabilityEquipmentPhysicalDef{value: val, isSet: true}
}

func (v NullableCapabilityEquipmentPhysicalDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityEquipmentPhysicalDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}