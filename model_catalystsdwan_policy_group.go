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

// checks if the CatalystsdwanPolicyGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalystsdwanPolicyGroup{}

// CatalystsdwanPolicyGroup Details for the Catalyst SDWAN policy groups.
type CatalystsdwanPolicyGroup struct {
	CatalystsdwanInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string   `json:"ObjectType"`
	Devices    []string `json:"Devices,omitempty"`
	// The Catalyst SDWAN policy group name.
	Name *string `json:"Name,omitempty"`
	// The Catalyst SDWAN policy group number of devices.
	NumberOfDevices *int64 `json:"NumberOfDevices,omitempty"`
	// The Catalyst SDWAN policy group number of devices up to date.
	NumberOfDevicesUpToDate *int64 `json:"NumberOfDevicesUpToDate,omitempty"`
	// UUID for the Catalyst SDWAN policy group.
	PolicyGroupId *string `json:"PolicyGroupId,omitempty"`
	// The Catalyst SDWAN policy group solution.
	Solution *string `json:"Solution,omitempty"`
	// The Catalyst SDWAN policy group version.
	Version              *int64 `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatalystsdwanPolicyGroup CatalystsdwanPolicyGroup

// NewCatalystsdwanPolicyGroup instantiates a new CatalystsdwanPolicyGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalystsdwanPolicyGroup(classId string, objectType string) *CatalystsdwanPolicyGroup {
	this := CatalystsdwanPolicyGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCatalystsdwanPolicyGroupWithDefaults instantiates a new CatalystsdwanPolicyGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalystsdwanPolicyGroupWithDefaults() *CatalystsdwanPolicyGroup {
	this := CatalystsdwanPolicyGroup{}
	var classId string = "catalystsdwan.PolicyGroup"
	this.ClassId = classId
	var objectType string = "catalystsdwan.PolicyGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CatalystsdwanPolicyGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CatalystsdwanPolicyGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CatalystsdwanPolicyGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "catalystsdwan.PolicyGroup" of the ClassId field.
func (o *CatalystsdwanPolicyGroup) GetDefaultClassId() interface{} {
	return "catalystsdwan.PolicyGroup"
}

// GetObjectType returns the ObjectType field value
func (o *CatalystsdwanPolicyGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CatalystsdwanPolicyGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CatalystsdwanPolicyGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "catalystsdwan.PolicyGroup" of the ObjectType field.
func (o *CatalystsdwanPolicyGroup) GetDefaultObjectType() interface{} {
	return "catalystsdwan.PolicyGroup"
}

// GetDevices returns the Devices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalystsdwanPolicyGroup) GetDevices() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalystsdwanPolicyGroup) GetDevicesOk() ([]string, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *CatalystsdwanPolicyGroup) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []string and assigns it to the Devices field.
func (o *CatalystsdwanPolicyGroup) SetDevices(v []string) {
	o.Devices = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalystsdwanPolicyGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanPolicyGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalystsdwanPolicyGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalystsdwanPolicyGroup) SetName(v string) {
	o.Name = &v
}

// GetNumberOfDevices returns the NumberOfDevices field value if set, zero value otherwise.
func (o *CatalystsdwanPolicyGroup) GetNumberOfDevices() int64 {
	if o == nil || IsNil(o.NumberOfDevices) {
		var ret int64
		return ret
	}
	return *o.NumberOfDevices
}

// GetNumberOfDevicesOk returns a tuple with the NumberOfDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanPolicyGroup) GetNumberOfDevicesOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfDevices) {
		return nil, false
	}
	return o.NumberOfDevices, true
}

// HasNumberOfDevices returns a boolean if a field has been set.
func (o *CatalystsdwanPolicyGroup) HasNumberOfDevices() bool {
	if o != nil && !IsNil(o.NumberOfDevices) {
		return true
	}

	return false
}

// SetNumberOfDevices gets a reference to the given int64 and assigns it to the NumberOfDevices field.
func (o *CatalystsdwanPolicyGroup) SetNumberOfDevices(v int64) {
	o.NumberOfDevices = &v
}

// GetNumberOfDevicesUpToDate returns the NumberOfDevicesUpToDate field value if set, zero value otherwise.
func (o *CatalystsdwanPolicyGroup) GetNumberOfDevicesUpToDate() int64 {
	if o == nil || IsNil(o.NumberOfDevicesUpToDate) {
		var ret int64
		return ret
	}
	return *o.NumberOfDevicesUpToDate
}

// GetNumberOfDevicesUpToDateOk returns a tuple with the NumberOfDevicesUpToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanPolicyGroup) GetNumberOfDevicesUpToDateOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfDevicesUpToDate) {
		return nil, false
	}
	return o.NumberOfDevicesUpToDate, true
}

// HasNumberOfDevicesUpToDate returns a boolean if a field has been set.
func (o *CatalystsdwanPolicyGroup) HasNumberOfDevicesUpToDate() bool {
	if o != nil && !IsNil(o.NumberOfDevicesUpToDate) {
		return true
	}

	return false
}

// SetNumberOfDevicesUpToDate gets a reference to the given int64 and assigns it to the NumberOfDevicesUpToDate field.
func (o *CatalystsdwanPolicyGroup) SetNumberOfDevicesUpToDate(v int64) {
	o.NumberOfDevicesUpToDate = &v
}

// GetPolicyGroupId returns the PolicyGroupId field value if set, zero value otherwise.
func (o *CatalystsdwanPolicyGroup) GetPolicyGroupId() string {
	if o == nil || IsNil(o.PolicyGroupId) {
		var ret string
		return ret
	}
	return *o.PolicyGroupId
}

// GetPolicyGroupIdOk returns a tuple with the PolicyGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanPolicyGroup) GetPolicyGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyGroupId) {
		return nil, false
	}
	return o.PolicyGroupId, true
}

// HasPolicyGroupId returns a boolean if a field has been set.
func (o *CatalystsdwanPolicyGroup) HasPolicyGroupId() bool {
	if o != nil && !IsNil(o.PolicyGroupId) {
		return true
	}

	return false
}

// SetPolicyGroupId gets a reference to the given string and assigns it to the PolicyGroupId field.
func (o *CatalystsdwanPolicyGroup) SetPolicyGroupId(v string) {
	o.PolicyGroupId = &v
}

// GetSolution returns the Solution field value if set, zero value otherwise.
func (o *CatalystsdwanPolicyGroup) GetSolution() string {
	if o == nil || IsNil(o.Solution) {
		var ret string
		return ret
	}
	return *o.Solution
}

// GetSolutionOk returns a tuple with the Solution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanPolicyGroup) GetSolutionOk() (*string, bool) {
	if o == nil || IsNil(o.Solution) {
		return nil, false
	}
	return o.Solution, true
}

// HasSolution returns a boolean if a field has been set.
func (o *CatalystsdwanPolicyGroup) HasSolution() bool {
	if o != nil && !IsNil(o.Solution) {
		return true
	}

	return false
}

// SetSolution gets a reference to the given string and assigns it to the Solution field.
func (o *CatalystsdwanPolicyGroup) SetSolution(v string) {
	o.Solution = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CatalystsdwanPolicyGroup) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalystsdwanPolicyGroup) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CatalystsdwanPolicyGroup) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *CatalystsdwanPolicyGroup) SetVersion(v int64) {
	o.Version = &v
}

func (o CatalystsdwanPolicyGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalystsdwanPolicyGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCatalystsdwanInventoryEntity, errCatalystsdwanInventoryEntity := json.Marshal(o.CatalystsdwanInventoryEntity)
	if errCatalystsdwanInventoryEntity != nil {
		return map[string]interface{}{}, errCatalystsdwanInventoryEntity
	}
	errCatalystsdwanInventoryEntity = json.Unmarshal([]byte(serializedCatalystsdwanInventoryEntity), &toSerialize)
	if errCatalystsdwanInventoryEntity != nil {
		return map[string]interface{}{}, errCatalystsdwanInventoryEntity
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Devices != nil {
		toSerialize["Devices"] = o.Devices
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.NumberOfDevices) {
		toSerialize["NumberOfDevices"] = o.NumberOfDevices
	}
	if !IsNil(o.NumberOfDevicesUpToDate) {
		toSerialize["NumberOfDevicesUpToDate"] = o.NumberOfDevicesUpToDate
	}
	if !IsNil(o.PolicyGroupId) {
		toSerialize["PolicyGroupId"] = o.PolicyGroupId
	}
	if !IsNil(o.Solution) {
		toSerialize["Solution"] = o.Solution
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalystsdwanPolicyGroup) UnmarshalJSON(data []byte) (err error) {
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
	type CatalystsdwanPolicyGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string   `json:"ObjectType"`
		Devices    []string `json:"Devices,omitempty"`
		// The Catalyst SDWAN policy group name.
		Name *string `json:"Name,omitempty"`
		// The Catalyst SDWAN policy group number of devices.
		NumberOfDevices *int64 `json:"NumberOfDevices,omitempty"`
		// The Catalyst SDWAN policy group number of devices up to date.
		NumberOfDevicesUpToDate *int64 `json:"NumberOfDevicesUpToDate,omitempty"`
		// UUID for the Catalyst SDWAN policy group.
		PolicyGroupId *string `json:"PolicyGroupId,omitempty"`
		// The Catalyst SDWAN policy group solution.
		Solution *string `json:"Solution,omitempty"`
		// The Catalyst SDWAN policy group version.
		Version *int64 `json:"Version,omitempty"`
	}

	varCatalystsdwanPolicyGroupWithoutEmbeddedStruct := CatalystsdwanPolicyGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCatalystsdwanPolicyGroupWithoutEmbeddedStruct)
	if err == nil {
		varCatalystsdwanPolicyGroup := _CatalystsdwanPolicyGroup{}
		varCatalystsdwanPolicyGroup.ClassId = varCatalystsdwanPolicyGroupWithoutEmbeddedStruct.ClassId
		varCatalystsdwanPolicyGroup.ObjectType = varCatalystsdwanPolicyGroupWithoutEmbeddedStruct.ObjectType
		varCatalystsdwanPolicyGroup.Devices = varCatalystsdwanPolicyGroupWithoutEmbeddedStruct.Devices
		varCatalystsdwanPolicyGroup.Name = varCatalystsdwanPolicyGroupWithoutEmbeddedStruct.Name
		varCatalystsdwanPolicyGroup.NumberOfDevices = varCatalystsdwanPolicyGroupWithoutEmbeddedStruct.NumberOfDevices
		varCatalystsdwanPolicyGroup.NumberOfDevicesUpToDate = varCatalystsdwanPolicyGroupWithoutEmbeddedStruct.NumberOfDevicesUpToDate
		varCatalystsdwanPolicyGroup.PolicyGroupId = varCatalystsdwanPolicyGroupWithoutEmbeddedStruct.PolicyGroupId
		varCatalystsdwanPolicyGroup.Solution = varCatalystsdwanPolicyGroupWithoutEmbeddedStruct.Solution
		varCatalystsdwanPolicyGroup.Version = varCatalystsdwanPolicyGroupWithoutEmbeddedStruct.Version
		*o = CatalystsdwanPolicyGroup(varCatalystsdwanPolicyGroup)
	} else {
		return err
	}

	varCatalystsdwanPolicyGroup := _CatalystsdwanPolicyGroup{}

	err = json.Unmarshal(data, &varCatalystsdwanPolicyGroup)
	if err == nil {
		o.CatalystsdwanInventoryEntity = varCatalystsdwanPolicyGroup.CatalystsdwanInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Devices")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NumberOfDevices")
		delete(additionalProperties, "NumberOfDevicesUpToDate")
		delete(additionalProperties, "PolicyGroupId")
		delete(additionalProperties, "Solution")
		delete(additionalProperties, "Version")

		// remove fields from embedded structs
		reflectCatalystsdwanInventoryEntity := reflect.ValueOf(o.CatalystsdwanInventoryEntity)
		for i := 0; i < reflectCatalystsdwanInventoryEntity.Type().NumField(); i++ {
			t := reflectCatalystsdwanInventoryEntity.Type().Field(i)

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

type NullableCatalystsdwanPolicyGroup struct {
	value *CatalystsdwanPolicyGroup
	isSet bool
}

func (v NullableCatalystsdwanPolicyGroup) Get() *CatalystsdwanPolicyGroup {
	return v.value
}

func (v *NullableCatalystsdwanPolicyGroup) Set(val *CatalystsdwanPolicyGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalystsdwanPolicyGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalystsdwanPolicyGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalystsdwanPolicyGroup(val *CatalystsdwanPolicyGroup) *NullableCatalystsdwanPolicyGroup {
	return &NullableCatalystsdwanPolicyGroup{value: val, isSet: true}
}

func (v NullableCatalystsdwanPolicyGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalystsdwanPolicyGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
