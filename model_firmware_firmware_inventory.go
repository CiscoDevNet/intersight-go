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

// checks if the FirmwareFirmwareInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareFirmwareInventory{}

// FirmwareFirmwareInventory Firmware inventory for server component.
type FirmwareFirmwareInventory struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Component category. For example, MRAID is under storage controller, CIMC is under management controller.
	Category *string `json:"Category,omitempty"`
	// The name of the component to reflect on UI.
	Label *string `json:"Label,omitempty"`
	// Model details of component.
	Model *string `json:"Model,omitempty"`
	// The redfish URI to get the firmware inventory of server components.
	UpdateUri *string `json:"UpdateUri,omitempty"`
	// The vendor of the component.
	Vendor *string `json:"Vendor,omitempty"`
	// The firmware running version on the component.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareFirmwareInventory FirmwareFirmwareInventory

// NewFirmwareFirmwareInventory instantiates a new FirmwareFirmwareInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareFirmwareInventory(classId string, objectType string) *FirmwareFirmwareInventory {
	this := FirmwareFirmwareInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareFirmwareInventoryWithDefaults instantiates a new FirmwareFirmwareInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareFirmwareInventoryWithDefaults() *FirmwareFirmwareInventory {
	this := FirmwareFirmwareInventory{}
	var classId string = "firmware.FirmwareInventory"
	this.ClassId = classId
	var objectType string = "firmware.FirmwareInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareFirmwareInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareFirmwareInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "firmware.FirmwareInventory" of the ClassId field.
func (o *FirmwareFirmwareInventory) GetDefaultClassId() interface{} {
	return "firmware.FirmwareInventory"
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareFirmwareInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareFirmwareInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "firmware.FirmwareInventory" of the ObjectType field.
func (o *FirmwareFirmwareInventory) GetDefaultObjectType() interface{} {
	return "firmware.FirmwareInventory"
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *FirmwareFirmwareInventory) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *FirmwareFirmwareInventory) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *FirmwareFirmwareInventory) SetCategory(v string) {
	o.Category = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FirmwareFirmwareInventory) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FirmwareFirmwareInventory) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FirmwareFirmwareInventory) SetLabel(v string) {
	o.Label = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *FirmwareFirmwareInventory) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *FirmwareFirmwareInventory) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *FirmwareFirmwareInventory) SetModel(v string) {
	o.Model = &v
}

// GetUpdateUri returns the UpdateUri field value if set, zero value otherwise.
func (o *FirmwareFirmwareInventory) GetUpdateUri() string {
	if o == nil || IsNil(o.UpdateUri) {
		var ret string
		return ret
	}
	return *o.UpdateUri
}

// GetUpdateUriOk returns a tuple with the UpdateUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetUpdateUriOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateUri) {
		return nil, false
	}
	return o.UpdateUri, true
}

// HasUpdateUri returns a boolean if a field has been set.
func (o *FirmwareFirmwareInventory) HasUpdateUri() bool {
	if o != nil && !IsNil(o.UpdateUri) {
		return true
	}

	return false
}

// SetUpdateUri gets a reference to the given string and assigns it to the UpdateUri field.
func (o *FirmwareFirmwareInventory) SetUpdateUri(v string) {
	o.UpdateUri = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *FirmwareFirmwareInventory) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *FirmwareFirmwareInventory) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *FirmwareFirmwareInventory) SetVendor(v string) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FirmwareFirmwareInventory) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FirmwareFirmwareInventory) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FirmwareFirmwareInventory) SetVersion(v string) {
	o.Version = &v
}

func (o FirmwareFirmwareInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareFirmwareInventory) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Category) {
		toSerialize["Category"] = o.Category
	}
	if !IsNil(o.Label) {
		toSerialize["Label"] = o.Label
	}
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	if !IsNil(o.UpdateUri) {
		toSerialize["UpdateUri"] = o.UpdateUri
	}
	if !IsNil(o.Vendor) {
		toSerialize["Vendor"] = o.Vendor
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirmwareFirmwareInventory) UnmarshalJSON(data []byte) (err error) {
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
	type FirmwareFirmwareInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Component category. For example, MRAID is under storage controller, CIMC is under management controller.
		Category *string `json:"Category,omitempty"`
		// The name of the component to reflect on UI.
		Label *string `json:"Label,omitempty"`
		// Model details of component.
		Model *string `json:"Model,omitempty"`
		// The redfish URI to get the firmware inventory of server components.
		UpdateUri *string `json:"UpdateUri,omitempty"`
		// The vendor of the component.
		Vendor *string `json:"Vendor,omitempty"`
		// The firmware running version on the component.
		Version *string `json:"Version,omitempty"`
	}

	varFirmwareFirmwareInventoryWithoutEmbeddedStruct := FirmwareFirmwareInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFirmwareFirmwareInventoryWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareFirmwareInventory := _FirmwareFirmwareInventory{}
		varFirmwareFirmwareInventory.ClassId = varFirmwareFirmwareInventoryWithoutEmbeddedStruct.ClassId
		varFirmwareFirmwareInventory.ObjectType = varFirmwareFirmwareInventoryWithoutEmbeddedStruct.ObjectType
		varFirmwareFirmwareInventory.Category = varFirmwareFirmwareInventoryWithoutEmbeddedStruct.Category
		varFirmwareFirmwareInventory.Label = varFirmwareFirmwareInventoryWithoutEmbeddedStruct.Label
		varFirmwareFirmwareInventory.Model = varFirmwareFirmwareInventoryWithoutEmbeddedStruct.Model
		varFirmwareFirmwareInventory.UpdateUri = varFirmwareFirmwareInventoryWithoutEmbeddedStruct.UpdateUri
		varFirmwareFirmwareInventory.Vendor = varFirmwareFirmwareInventoryWithoutEmbeddedStruct.Vendor
		varFirmwareFirmwareInventory.Version = varFirmwareFirmwareInventoryWithoutEmbeddedStruct.Version
		*o = FirmwareFirmwareInventory(varFirmwareFirmwareInventory)
	} else {
		return err
	}

	varFirmwareFirmwareInventory := _FirmwareFirmwareInventory{}

	err = json.Unmarshal(data, &varFirmwareFirmwareInventory)
	if err == nil {
		o.MoBaseComplexType = varFirmwareFirmwareInventory.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "UpdateUri")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "Version")

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

type NullableFirmwareFirmwareInventory struct {
	value *FirmwareFirmwareInventory
	isSet bool
}

func (v NullableFirmwareFirmwareInventory) Get() *FirmwareFirmwareInventory {
	return v.value
}

func (v *NullableFirmwareFirmwareInventory) Set(val *FirmwareFirmwareInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareFirmwareInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareFirmwareInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareFirmwareInventory(val *FirmwareFirmwareInventory) *NullableFirmwareFirmwareInventory {
	return &NullableFirmwareFirmwareInventory{value: val, isSet: true}
}

func (v NullableFirmwareFirmwareInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareFirmwareInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
