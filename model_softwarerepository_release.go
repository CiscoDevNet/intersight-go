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
	"time"
)

// checks if the SoftwarerepositoryRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwarerepositoryRelease{}

// SoftwarerepositoryRelease A Cisco release containing one or more firmware images. Cisco releases images for rack server components or blade server components or for Fabric Interconnect components. The version for the firmware images is the same as specific Cisco release version.
type SoftwarerepositoryRelease struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The date when the file was released or distributed by its vendor.
	ReleaseDate *time.Time `json:"ReleaseDate,omitempty"`
	// The URL for the release notes of this image.
	ReleaseNotesUrl *string  `json:"ReleaseNotesUrl,omitempty"`
	SupportedModels []string `json:"SupportedModels,omitempty"`
	// The platform type for which the images are released. This can be a Fabric Interconnect or compute server hardware. * `FabricSwitch` - The images in a release that correspond to Fabric Interconnect switches. * `ComputeSystem` - The images in a release that correspond to servers.
	Type *string `json:"Type,omitempty"`
	// Cisco provided release version.
	Version              *string                                       `json:"Version,omitempty"`
	Catalog              NullableSoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryRelease SoftwarerepositoryRelease

// NewSoftwarerepositoryRelease instantiates a new SoftwarerepositoryRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryRelease(classId string, objectType string) *SoftwarerepositoryRelease {
	this := SoftwarerepositoryRelease{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "FabricSwitch"
	this.Type = &type_
	return &this
}

// NewSoftwarerepositoryReleaseWithDefaults instantiates a new SoftwarerepositoryRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryReleaseWithDefaults() *SoftwarerepositoryRelease {
	this := SoftwarerepositoryRelease{}
	var classId string = "softwarerepository.Release"
	this.ClassId = classId
	var objectType string = "softwarerepository.Release"
	this.ObjectType = objectType
	var type_ string = "FabricSwitch"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryRelease) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryRelease) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryRelease) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "softwarerepository.Release" of the ClassId field.
func (o *SoftwarerepositoryRelease) GetDefaultClassId() interface{} {
	return "softwarerepository.Release"
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryRelease) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryRelease) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryRelease) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "softwarerepository.Release" of the ObjectType field.
func (o *SoftwarerepositoryRelease) GetDefaultObjectType() interface{} {
	return "softwarerepository.Release"
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *SoftwarerepositoryRelease) GetReleaseDate() time.Time {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryRelease) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *SoftwarerepositoryRelease) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *SoftwarerepositoryRelease) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

// GetReleaseNotesUrl returns the ReleaseNotesUrl field value if set, zero value otherwise.
func (o *SoftwarerepositoryRelease) GetReleaseNotesUrl() string {
	if o == nil || IsNil(o.ReleaseNotesUrl) {
		var ret string
		return ret
	}
	return *o.ReleaseNotesUrl
}

// GetReleaseNotesUrlOk returns a tuple with the ReleaseNotesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryRelease) GetReleaseNotesUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseNotesUrl) {
		return nil, false
	}
	return o.ReleaseNotesUrl, true
}

// HasReleaseNotesUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryRelease) HasReleaseNotesUrl() bool {
	if o != nil && !IsNil(o.ReleaseNotesUrl) {
		return true
	}

	return false
}

// SetReleaseNotesUrl gets a reference to the given string and assigns it to the ReleaseNotesUrl field.
func (o *SoftwarerepositoryRelease) SetReleaseNotesUrl(v string) {
	o.ReleaseNotesUrl = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryRelease) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryRelease) GetSupportedModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedModels) {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryRelease) HasSupportedModels() bool {
	if o != nil && !IsNil(o.SupportedModels) {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *SoftwarerepositoryRelease) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SoftwarerepositoryRelease) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryRelease) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SoftwarerepositoryRelease) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SoftwarerepositoryRelease) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SoftwarerepositoryRelease) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryRelease) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SoftwarerepositoryRelease) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SoftwarerepositoryRelease) SetVersion(v string) {
	o.Version = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryRelease) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || IsNil(o.Catalog.Get()) {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog.Get()
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryRelease) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Catalog.Get(), o.Catalog.IsSet()
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwarerepositoryRelease) HasCatalog() bool {
	if o != nil && o.Catalog.IsSet() {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given NullableSoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwarerepositoryRelease) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog.Set(&v)
}

// SetCatalogNil sets the value for Catalog to be an explicit nil
func (o *SoftwarerepositoryRelease) SetCatalogNil() {
	o.Catalog.Set(nil)
}

// UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
func (o *SoftwarerepositoryRelease) UnsetCatalog() {
	o.Catalog.Unset()
}

func (o SoftwarerepositoryRelease) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwarerepositoryRelease) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ReleaseDate) {
		toSerialize["ReleaseDate"] = o.ReleaseDate
	}
	if !IsNil(o.ReleaseNotesUrl) {
		toSerialize["ReleaseNotesUrl"] = o.ReleaseNotesUrl
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if o.Catalog.IsSet() {
		toSerialize["Catalog"] = o.Catalog.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SoftwarerepositoryRelease) UnmarshalJSON(data []byte) (err error) {
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
	type SoftwarerepositoryReleaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The date when the file was released or distributed by its vendor.
		ReleaseDate *time.Time `json:"ReleaseDate,omitempty"`
		// The URL for the release notes of this image.
		ReleaseNotesUrl *string  `json:"ReleaseNotesUrl,omitempty"`
		SupportedModels []string `json:"SupportedModels,omitempty"`
		// The platform type for which the images are released. This can be a Fabric Interconnect or compute server hardware. * `FabricSwitch` - The images in a release that correspond to Fabric Interconnect switches. * `ComputeSystem` - The images in a release that correspond to servers.
		Type *string `json:"Type,omitempty"`
		// Cisco provided release version.
		Version *string                                       `json:"Version,omitempty"`
		Catalog NullableSoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	}

	varSoftwarerepositoryReleaseWithoutEmbeddedStruct := SoftwarerepositoryReleaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSoftwarerepositoryReleaseWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryRelease := _SoftwarerepositoryRelease{}
		varSoftwarerepositoryRelease.ClassId = varSoftwarerepositoryReleaseWithoutEmbeddedStruct.ClassId
		varSoftwarerepositoryRelease.ObjectType = varSoftwarerepositoryReleaseWithoutEmbeddedStruct.ObjectType
		varSoftwarerepositoryRelease.ReleaseDate = varSoftwarerepositoryReleaseWithoutEmbeddedStruct.ReleaseDate
		varSoftwarerepositoryRelease.ReleaseNotesUrl = varSoftwarerepositoryReleaseWithoutEmbeddedStruct.ReleaseNotesUrl
		varSoftwarerepositoryRelease.SupportedModels = varSoftwarerepositoryReleaseWithoutEmbeddedStruct.SupportedModels
		varSoftwarerepositoryRelease.Type = varSoftwarerepositoryReleaseWithoutEmbeddedStruct.Type
		varSoftwarerepositoryRelease.Version = varSoftwarerepositoryReleaseWithoutEmbeddedStruct.Version
		varSoftwarerepositoryRelease.Catalog = varSoftwarerepositoryReleaseWithoutEmbeddedStruct.Catalog
		*o = SoftwarerepositoryRelease(varSoftwarerepositoryRelease)
	} else {
		return err
	}

	varSoftwarerepositoryRelease := _SoftwarerepositoryRelease{}

	err = json.Unmarshal(data, &varSoftwarerepositoryRelease)
	if err == nil {
		o.MoBaseMo = varSoftwarerepositoryRelease.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ReleaseDate")
		delete(additionalProperties, "ReleaseNotesUrl")
		delete(additionalProperties, "SupportedModels")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Catalog")

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

type NullableSoftwarerepositoryRelease struct {
	value *SoftwarerepositoryRelease
	isSet bool
}

func (v NullableSoftwarerepositoryRelease) Get() *SoftwarerepositoryRelease {
	return v.value
}

func (v *NullableSoftwarerepositoryRelease) Set(val *SoftwarerepositoryRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryRelease(val *SoftwarerepositoryRelease) *NullableSoftwarerepositoryRelease {
	return &NullableSoftwarerepositoryRelease{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
