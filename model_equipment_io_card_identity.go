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

// checks if the EquipmentIoCardIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentIoCardIdentity{}

// EquipmentIoCardIdentity IoCardIdentity Complex type referenced in ChassisIdentity concrete MO.
type EquipmentIoCardIdentity struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// MO Reference to equipmentIoCard MO in inventory service.
	IoCardMoid *string `json:"IoCardMoid,omitempty"`
	// IO Card inventory lifecycle status. * `Unknown` - Default lifecycle state of a iocard. This should be an initial state when no state is defined for a specific iocard. * `Decommissioned` - Lifecycle state is set to this value after the chassis is successfully decommissioned. * `DiscoveryInProgress` - Lifecycle state is set to this value after the successful start of the iocard connection discovery process. * `DiscoveryFailed` - Lifecycle state is set to this value after the iocard connection discovery has failed. * `DiscoveryCompleted` - Lifecycle state is set to this value after the connection discovery of the iocard is completed successfully. * `None` - Lifecycle state is set to this value before the start of connection discovery and inventory collection process for a iocard. * `InventoryCompleted` - Lifecycle state is set to this value after the chassis inventory collection process is completed for a specific iocard. * `InventoryInProgress` - Lifecycle state is set to this value after successful  start of the chassis inventory collection process for a specific iocard. * `InventoryFailed` - Lifecycle state is set to this value after the chassis inventory collection process failed for a iocard.
	Lifecycle *string `json:"Lifecycle,omitempty"`
	// IO Card or intelligent fabric module model.
	Model *string `json:"Model,omitempty"`
	// IOM/MUX Module ID connected to the FI.
	ModuleId *int64 `json:"ModuleId,omitempty"`
	// MO Reference to networkElement MO in inventory service.
	NetworkElementMoid *string `json:"NetworkElementMoid,omitempty"`
	// IO Card or intelligent fabric module serial number.
	Serial *string `json:"Serial,omitempty"`
	// Identifier of the Switch where the IOM is connected. ID can be either 1 or 2, equivalent to A or B.
	SwitchId *int64 `json:"SwitchId,omitempty"`
	// IO Card or intelligent fabric module vendor.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentIoCardIdentity EquipmentIoCardIdentity

// NewEquipmentIoCardIdentity instantiates a new EquipmentIoCardIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIoCardIdentity(classId string, objectType string) *EquipmentIoCardIdentity {
	this := EquipmentIoCardIdentity{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentIoCardIdentityWithDefaults instantiates a new EquipmentIoCardIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIoCardIdentityWithDefaults() *EquipmentIoCardIdentity {
	this := EquipmentIoCardIdentity{}
	var classId string = "equipment.IoCardIdentity"
	this.ClassId = classId
	var objectType string = "equipment.IoCardIdentity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIoCardIdentity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIoCardIdentity) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "equipment.IoCardIdentity" of the ClassId field.
func (o *EquipmentIoCardIdentity) GetDefaultClassId() interface{} {
	return "equipment.IoCardIdentity"
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIoCardIdentity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIoCardIdentity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "equipment.IoCardIdentity" of the ObjectType field.
func (o *EquipmentIoCardIdentity) GetDefaultObjectType() interface{} {
	return "equipment.IoCardIdentity"
}

// GetIoCardMoid returns the IoCardMoid field value if set, zero value otherwise.
func (o *EquipmentIoCardIdentity) GetIoCardMoid() string {
	if o == nil || IsNil(o.IoCardMoid) {
		var ret string
		return ret
	}
	return *o.IoCardMoid
}

// GetIoCardMoidOk returns a tuple with the IoCardMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentity) GetIoCardMoidOk() (*string, bool) {
	if o == nil || IsNil(o.IoCardMoid) {
		return nil, false
	}
	return o.IoCardMoid, true
}

// HasIoCardMoid returns a boolean if a field has been set.
func (o *EquipmentIoCardIdentity) HasIoCardMoid() bool {
	if o != nil && !IsNil(o.IoCardMoid) {
		return true
	}

	return false
}

// SetIoCardMoid gets a reference to the given string and assigns it to the IoCardMoid field.
func (o *EquipmentIoCardIdentity) SetIoCardMoid(v string) {
	o.IoCardMoid = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *EquipmentIoCardIdentity) GetLifecycle() string {
	if o == nil || IsNil(o.Lifecycle) {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentity) GetLifecycleOk() (*string, bool) {
	if o == nil || IsNil(o.Lifecycle) {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *EquipmentIoCardIdentity) HasLifecycle() bool {
	if o != nil && !IsNil(o.Lifecycle) {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *EquipmentIoCardIdentity) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *EquipmentIoCardIdentity) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentity) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *EquipmentIoCardIdentity) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *EquipmentIoCardIdentity) SetModel(v string) {
	o.Model = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EquipmentIoCardIdentity) GetModuleId() int64 {
	if o == nil || IsNil(o.ModuleId) {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentity) GetModuleIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ModuleId) {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EquipmentIoCardIdentity) HasModuleId() bool {
	if o != nil && !IsNil(o.ModuleId) {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EquipmentIoCardIdentity) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetNetworkElementMoid returns the NetworkElementMoid field value if set, zero value otherwise.
func (o *EquipmentIoCardIdentity) GetNetworkElementMoid() string {
	if o == nil || IsNil(o.NetworkElementMoid) {
		var ret string
		return ret
	}
	return *o.NetworkElementMoid
}

// GetNetworkElementMoidOk returns a tuple with the NetworkElementMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentity) GetNetworkElementMoidOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkElementMoid) {
		return nil, false
	}
	return o.NetworkElementMoid, true
}

// HasNetworkElementMoid returns a boolean if a field has been set.
func (o *EquipmentIoCardIdentity) HasNetworkElementMoid() bool {
	if o != nil && !IsNil(o.NetworkElementMoid) {
		return true
	}

	return false
}

// SetNetworkElementMoid gets a reference to the given string and assigns it to the NetworkElementMoid field.
func (o *EquipmentIoCardIdentity) SetNetworkElementMoid(v string) {
	o.NetworkElementMoid = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *EquipmentIoCardIdentity) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentity) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *EquipmentIoCardIdentity) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *EquipmentIoCardIdentity) SetSerial(v string) {
	o.Serial = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *EquipmentIoCardIdentity) GetSwitchId() int64 {
	if o == nil || IsNil(o.SwitchId) {
		var ret int64
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentity) GetSwitchIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SwitchId) {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *EquipmentIoCardIdentity) HasSwitchId() bool {
	if o != nil && !IsNil(o.SwitchId) {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given int64 and assigns it to the SwitchId field.
func (o *EquipmentIoCardIdentity) SetSwitchId(v int64) {
	o.SwitchId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *EquipmentIoCardIdentity) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentity) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *EquipmentIoCardIdentity) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *EquipmentIoCardIdentity) SetVendor(v string) {
	o.Vendor = &v
}

func (o EquipmentIoCardIdentity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentIoCardIdentity) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IoCardMoid) {
		toSerialize["IoCardMoid"] = o.IoCardMoid
	}
	if !IsNil(o.Lifecycle) {
		toSerialize["Lifecycle"] = o.Lifecycle
	}
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	if !IsNil(o.ModuleId) {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if !IsNil(o.NetworkElementMoid) {
		toSerialize["NetworkElementMoid"] = o.NetworkElementMoid
	}
	if !IsNil(o.Serial) {
		toSerialize["Serial"] = o.Serial
	}
	if !IsNil(o.SwitchId) {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if !IsNil(o.Vendor) {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentIoCardIdentity) UnmarshalJSON(data []byte) (err error) {
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
	type EquipmentIoCardIdentityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// MO Reference to equipmentIoCard MO in inventory service.
		IoCardMoid *string `json:"IoCardMoid,omitempty"`
		// IO Card inventory lifecycle status. * `Unknown` - Default lifecycle state of a iocard. This should be an initial state when no state is defined for a specific iocard. * `Decommissioned` - Lifecycle state is set to this value after the chassis is successfully decommissioned. * `DiscoveryInProgress` - Lifecycle state is set to this value after the successful start of the iocard connection discovery process. * `DiscoveryFailed` - Lifecycle state is set to this value after the iocard connection discovery has failed. * `DiscoveryCompleted` - Lifecycle state is set to this value after the connection discovery of the iocard is completed successfully. * `None` - Lifecycle state is set to this value before the start of connection discovery and inventory collection process for a iocard. * `InventoryCompleted` - Lifecycle state is set to this value after the chassis inventory collection process is completed for a specific iocard. * `InventoryInProgress` - Lifecycle state is set to this value after successful  start of the chassis inventory collection process for a specific iocard. * `InventoryFailed` - Lifecycle state is set to this value after the chassis inventory collection process failed for a iocard.
		Lifecycle *string `json:"Lifecycle,omitempty"`
		// IO Card or intelligent fabric module model.
		Model *string `json:"Model,omitempty"`
		// IOM/MUX Module ID connected to the FI.
		ModuleId *int64 `json:"ModuleId,omitempty"`
		// MO Reference to networkElement MO in inventory service.
		NetworkElementMoid *string `json:"NetworkElementMoid,omitempty"`
		// IO Card or intelligent fabric module serial number.
		Serial *string `json:"Serial,omitempty"`
		// Identifier of the Switch where the IOM is connected. ID can be either 1 or 2, equivalent to A or B.
		SwitchId *int64 `json:"SwitchId,omitempty"`
		// IO Card or intelligent fabric module vendor.
		Vendor *string `json:"Vendor,omitempty"`
	}

	varEquipmentIoCardIdentityWithoutEmbeddedStruct := EquipmentIoCardIdentityWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentIoCardIdentityWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentIoCardIdentity := _EquipmentIoCardIdentity{}
		varEquipmentIoCardIdentity.ClassId = varEquipmentIoCardIdentityWithoutEmbeddedStruct.ClassId
		varEquipmentIoCardIdentity.ObjectType = varEquipmentIoCardIdentityWithoutEmbeddedStruct.ObjectType
		varEquipmentIoCardIdentity.IoCardMoid = varEquipmentIoCardIdentityWithoutEmbeddedStruct.IoCardMoid
		varEquipmentIoCardIdentity.Lifecycle = varEquipmentIoCardIdentityWithoutEmbeddedStruct.Lifecycle
		varEquipmentIoCardIdentity.Model = varEquipmentIoCardIdentityWithoutEmbeddedStruct.Model
		varEquipmentIoCardIdentity.ModuleId = varEquipmentIoCardIdentityWithoutEmbeddedStruct.ModuleId
		varEquipmentIoCardIdentity.NetworkElementMoid = varEquipmentIoCardIdentityWithoutEmbeddedStruct.NetworkElementMoid
		varEquipmentIoCardIdentity.Serial = varEquipmentIoCardIdentityWithoutEmbeddedStruct.Serial
		varEquipmentIoCardIdentity.SwitchId = varEquipmentIoCardIdentityWithoutEmbeddedStruct.SwitchId
		varEquipmentIoCardIdentity.Vendor = varEquipmentIoCardIdentityWithoutEmbeddedStruct.Vendor
		*o = EquipmentIoCardIdentity(varEquipmentIoCardIdentity)
	} else {
		return err
	}

	varEquipmentIoCardIdentity := _EquipmentIoCardIdentity{}

	err = json.Unmarshal(data, &varEquipmentIoCardIdentity)
	if err == nil {
		o.MoBaseComplexType = varEquipmentIoCardIdentity.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IoCardMoid")
		delete(additionalProperties, "Lifecycle")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "ModuleId")
		delete(additionalProperties, "NetworkElementMoid")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "Vendor")

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

type NullableEquipmentIoCardIdentity struct {
	value *EquipmentIoCardIdentity
	isSet bool
}

func (v NullableEquipmentIoCardIdentity) Get() *EquipmentIoCardIdentity {
	return v.value
}

func (v *NullableEquipmentIoCardIdentity) Set(val *EquipmentIoCardIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCardIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCardIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCardIdentity(val *EquipmentIoCardIdentity) *NullableEquipmentIoCardIdentity {
	return &NullableEquipmentIoCardIdentity{value: val, isSet: true}
}

func (v NullableEquipmentIoCardIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCardIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
