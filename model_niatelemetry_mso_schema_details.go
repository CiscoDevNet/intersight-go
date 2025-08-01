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

// checks if the NiatelemetryMsoSchemaDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryMsoSchemaDetails{}

// NiatelemetryMsoSchemaDetails Details of schema in Multi-Site Orchestrator.
type NiatelemetryMsoSchemaDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Site IDs to which this schema is deployed to.
	DeployedSites *string `json:"DeployedSites,omitempty"`
	// Number of policy objects per schema.
	NumberOfPolicyObjectsPerSchema *int64 `json:"NumberOfPolicyObjectsPerSchema,omitempty"`
	// Number of tenants assigned per schema in Multi-Site Orchestrator.
	NumberOfTemplatesPerSchema *int64 `json:"NumberOfTemplatesPerSchema,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Schema ID in Multi-Site Orchestrator.
	SchemaId *string `json:"SchemaId,omitempty"`
	// Schema name in Multi-Site Orchestrator.
	SchemaName           *string                                     `json:"SchemaName,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryMsoSchemaDetails NiatelemetryMsoSchemaDetails

// NewNiatelemetryMsoSchemaDetails instantiates a new NiatelemetryMsoSchemaDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryMsoSchemaDetails(classId string, objectType string) *NiatelemetryMsoSchemaDetails {
	this := NiatelemetryMsoSchemaDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryMsoSchemaDetailsWithDefaults instantiates a new NiatelemetryMsoSchemaDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryMsoSchemaDetailsWithDefaults() *NiatelemetryMsoSchemaDetails {
	this := NiatelemetryMsoSchemaDetails{}
	var classId string = "niatelemetry.MsoSchemaDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.MsoSchemaDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryMsoSchemaDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryMsoSchemaDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.MsoSchemaDetails" of the ClassId field.
func (o *NiatelemetryMsoSchemaDetails) GetDefaultClassId() interface{} {
	return "niatelemetry.MsoSchemaDetails"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryMsoSchemaDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryMsoSchemaDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.MsoSchemaDetails" of the ObjectType field.
func (o *NiatelemetryMsoSchemaDetails) GetDefaultObjectType() interface{} {
	return "niatelemetry.MsoSchemaDetails"
}

// GetDeployedSites returns the DeployedSites field value if set, zero value otherwise.
func (o *NiatelemetryMsoSchemaDetails) GetDeployedSites() string {
	if o == nil || IsNil(o.DeployedSites) {
		var ret string
		return ret
	}
	return *o.DeployedSites
}

// GetDeployedSitesOk returns a tuple with the DeployedSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetDeployedSitesOk() (*string, bool) {
	if o == nil || IsNil(o.DeployedSites) {
		return nil, false
	}
	return o.DeployedSites, true
}

// HasDeployedSites returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasDeployedSites() bool {
	if o != nil && !IsNil(o.DeployedSites) {
		return true
	}

	return false
}

// SetDeployedSites gets a reference to the given string and assigns it to the DeployedSites field.
func (o *NiatelemetryMsoSchemaDetails) SetDeployedSites(v string) {
	o.DeployedSites = &v
}

// GetNumberOfPolicyObjectsPerSchema returns the NumberOfPolicyObjectsPerSchema field value if set, zero value otherwise.
func (o *NiatelemetryMsoSchemaDetails) GetNumberOfPolicyObjectsPerSchema() int64 {
	if o == nil || IsNil(o.NumberOfPolicyObjectsPerSchema) {
		var ret int64
		return ret
	}
	return *o.NumberOfPolicyObjectsPerSchema
}

// GetNumberOfPolicyObjectsPerSchemaOk returns a tuple with the NumberOfPolicyObjectsPerSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetNumberOfPolicyObjectsPerSchemaOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfPolicyObjectsPerSchema) {
		return nil, false
	}
	return o.NumberOfPolicyObjectsPerSchema, true
}

// HasNumberOfPolicyObjectsPerSchema returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasNumberOfPolicyObjectsPerSchema() bool {
	if o != nil && !IsNil(o.NumberOfPolicyObjectsPerSchema) {
		return true
	}

	return false
}

// SetNumberOfPolicyObjectsPerSchema gets a reference to the given int64 and assigns it to the NumberOfPolicyObjectsPerSchema field.
func (o *NiatelemetryMsoSchemaDetails) SetNumberOfPolicyObjectsPerSchema(v int64) {
	o.NumberOfPolicyObjectsPerSchema = &v
}

// GetNumberOfTemplatesPerSchema returns the NumberOfTemplatesPerSchema field value if set, zero value otherwise.
func (o *NiatelemetryMsoSchemaDetails) GetNumberOfTemplatesPerSchema() int64 {
	if o == nil || IsNil(o.NumberOfTemplatesPerSchema) {
		var ret int64
		return ret
	}
	return *o.NumberOfTemplatesPerSchema
}

// GetNumberOfTemplatesPerSchemaOk returns a tuple with the NumberOfTemplatesPerSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetNumberOfTemplatesPerSchemaOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfTemplatesPerSchema) {
		return nil, false
	}
	return o.NumberOfTemplatesPerSchema, true
}

// HasNumberOfTemplatesPerSchema returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasNumberOfTemplatesPerSchema() bool {
	if o != nil && !IsNil(o.NumberOfTemplatesPerSchema) {
		return true
	}

	return false
}

// SetNumberOfTemplatesPerSchema gets a reference to the given int64 and assigns it to the NumberOfTemplatesPerSchema field.
func (o *NiatelemetryMsoSchemaDetails) SetNumberOfTemplatesPerSchema(v int64) {
	o.NumberOfTemplatesPerSchema = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryMsoSchemaDetails) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryMsoSchemaDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *NiatelemetryMsoSchemaDetails) GetSchemaId() string {
	if o == nil || IsNil(o.SchemaId) {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetSchemaIdOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaId) {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasSchemaId() bool {
	if o != nil && !IsNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *NiatelemetryMsoSchemaDetails) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise.
func (o *NiatelemetryMsoSchemaDetails) GetSchemaName() string {
	if o == nil || IsNil(o.SchemaName) {
		var ret string
		return ret
	}
	return *o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetSchemaNameOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaName) {
		return nil, false
	}
	return o.SchemaName, true
}

// HasSchemaName returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasSchemaName() bool {
	if o != nil && !IsNil(o.SchemaName) {
		return true
	}

	return false
}

// SetSchemaName gets a reference to the given string and assigns it to the SchemaName field.
func (o *NiatelemetryMsoSchemaDetails) SetSchemaName(v string) {
	o.SchemaName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryMsoSchemaDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryMsoSchemaDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryMsoSchemaDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryMsoSchemaDetails) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryMsoSchemaDetails) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryMsoSchemaDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryMsoSchemaDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DeployedSites) {
		toSerialize["DeployedSites"] = o.DeployedSites
	}
	if !IsNil(o.NumberOfPolicyObjectsPerSchema) {
		toSerialize["NumberOfPolicyObjectsPerSchema"] = o.NumberOfPolicyObjectsPerSchema
	}
	if !IsNil(o.NumberOfTemplatesPerSchema) {
		toSerialize["NumberOfTemplatesPerSchema"] = o.NumberOfTemplatesPerSchema
	}
	if !IsNil(o.RecordType) {
		toSerialize["RecordType"] = o.RecordType
	}
	if !IsNil(o.SchemaId) {
		toSerialize["SchemaId"] = o.SchemaId
	}
	if !IsNil(o.SchemaName) {
		toSerialize["SchemaName"] = o.SchemaName
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryMsoSchemaDetails) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Site IDs to which this schema is deployed to.
		DeployedSites *string `json:"DeployedSites,omitempty"`
		// Number of policy objects per schema.
		NumberOfPolicyObjectsPerSchema *int64 `json:"NumberOfPolicyObjectsPerSchema,omitempty"`
		// Number of tenants assigned per schema in Multi-Site Orchestrator.
		NumberOfTemplatesPerSchema *int64 `json:"NumberOfTemplatesPerSchema,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Schema ID in Multi-Site Orchestrator.
		SchemaId *string `json:"SchemaId,omitempty"`
		// Schema name in Multi-Site Orchestrator.
		SchemaName       *string                                     `json:"SchemaName,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct := NiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryMsoSchemaDetails := _NiatelemetryMsoSchemaDetails{}
		varNiatelemetryMsoSchemaDetails.ClassId = varNiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryMsoSchemaDetails.ObjectType = varNiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryMsoSchemaDetails.DeployedSites = varNiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct.DeployedSites
		varNiatelemetryMsoSchemaDetails.NumberOfPolicyObjectsPerSchema = varNiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct.NumberOfPolicyObjectsPerSchema
		varNiatelemetryMsoSchemaDetails.NumberOfTemplatesPerSchema = varNiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct.NumberOfTemplatesPerSchema
		varNiatelemetryMsoSchemaDetails.RecordType = varNiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryMsoSchemaDetails.SchemaId = varNiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct.SchemaId
		varNiatelemetryMsoSchemaDetails.SchemaName = varNiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct.SchemaName
		varNiatelemetryMsoSchemaDetails.RegisteredDevice = varNiatelemetryMsoSchemaDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryMsoSchemaDetails(varNiatelemetryMsoSchemaDetails)
	} else {
		return err
	}

	varNiatelemetryMsoSchemaDetails := _NiatelemetryMsoSchemaDetails{}

	err = json.Unmarshal(data, &varNiatelemetryMsoSchemaDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryMsoSchemaDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeployedSites")
		delete(additionalProperties, "NumberOfPolicyObjectsPerSchema")
		delete(additionalProperties, "NumberOfTemplatesPerSchema")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "SchemaId")
		delete(additionalProperties, "SchemaName")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNiatelemetryMsoSchemaDetails struct {
	value *NiatelemetryMsoSchemaDetails
	isSet bool
}

func (v NullableNiatelemetryMsoSchemaDetails) Get() *NiatelemetryMsoSchemaDetails {
	return v.value
}

func (v *NullableNiatelemetryMsoSchemaDetails) Set(val *NiatelemetryMsoSchemaDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryMsoSchemaDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryMsoSchemaDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryMsoSchemaDetails(val *NiatelemetryMsoSchemaDetails) *NullableNiatelemetryMsoSchemaDetails {
	return &NullableNiatelemetryMsoSchemaDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryMsoSchemaDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryMsoSchemaDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
