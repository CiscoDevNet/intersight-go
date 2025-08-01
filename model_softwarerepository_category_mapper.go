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

// checks if the SoftwarerepositoryCategoryMapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwarerepositoryCategoryMapper{}

// SoftwarerepositoryCategoryMapper Maps a Cisco software repository image category identifier to its applicable hardware models.
type SoftwarerepositoryCategoryMapper struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The category of the model series.
	Category *string `json:"Category,omitempty"`
	// The type of distributable image, example huu, scu, driver, os. * `Distributable` - Stores firmware host utility images and fabric images. * `DriverDistributable` - Stores driver distributable images. * `ServerConfigurationUtilityDistributable` - Stores server configuration utility images. * `OperatingSystemFile` - Stores operating system iso images. * `HyperflexDistributable` - It stores HyperFlex images. * `HciDistributable` - It stores HCI images, such as bootstrap iso images.
	FileType *string `json:"FileType,omitempty"`
	// The type of image based on the endpoint it can upgrade. For example, ucs-c420m5-huu-3.2.1a.iso can upgrade standalone servers, so the image type is Standalone Server.
	ImageType *string `json:"ImageType,omitempty"`
	// Defines whether NFS firmware upgrade is supported with this image type.
	IsNfsUpgradeSupported *bool `json:"IsNfsUpgradeSupported,omitempty"`
	// Cisco software repository image category identifier.
	MdfId *string `json:"MdfId,omitempty"`
	// The regex that all images of this category follow.
	RegexPattern *string `json:"RegexPattern,omitempty"`
	// The image can be downloaded from cisco.com or external cloud store. * `Cisco` - External repository hosted on cisco.com. * `IntersightCloud` - Repository hosted by the Intersight Cloud. * `LocalMachine` - The file is available on the local client machine. Used as an upload source type. * `NetworkShare` - External repository in the customer datacenter. This will typically be a file server.
	Source          *string  `json:"Source,omitempty"`
	SupportedModels []string `json:"SupportedModels,omitempty"`
	// The software type id provided by cisco.com.
	SwId     *string  `json:"SwId,omitempty"`
	TagTypes []string `json:"TagTypes,omitempty"`
	// The version from which user can download images from amazon store, if source is external cloud store.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryCategoryMapper SoftwarerepositoryCategoryMapper

// NewSoftwarerepositoryCategoryMapper instantiates a new SoftwarerepositoryCategoryMapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCategoryMapper(classId string, objectType string) *SoftwarerepositoryCategoryMapper {
	this := SoftwarerepositoryCategoryMapper{}
	this.ClassId = classId
	this.ObjectType = objectType
	var fileType string = "Distributable"
	this.FileType = &fileType
	var source string = "Cisco"
	this.Source = &source
	return &this
}

// NewSoftwarerepositoryCategoryMapperWithDefaults instantiates a new SoftwarerepositoryCategoryMapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCategoryMapperWithDefaults() *SoftwarerepositoryCategoryMapper {
	this := SoftwarerepositoryCategoryMapper{}
	var classId string = "softwarerepository.CategoryMapper"
	this.ClassId = classId
	var objectType string = "softwarerepository.CategoryMapper"
	this.ObjectType = objectType
	var fileType string = "Distributable"
	this.FileType = &fileType
	var source string = "Cisco"
	this.Source = &source
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryCategoryMapper) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryCategoryMapper) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "softwarerepository.CategoryMapper" of the ClassId field.
func (o *SoftwarerepositoryCategoryMapper) GetDefaultClassId() interface{} {
	return "softwarerepository.CategoryMapper"
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryCategoryMapper) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryCategoryMapper) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "softwarerepository.CategoryMapper" of the ObjectType field.
func (o *SoftwarerepositoryCategoryMapper) GetDefaultObjectType() interface{} {
	return "softwarerepository.CategoryMapper"
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SoftwarerepositoryCategoryMapper) SetCategory(v string) {
	o.Category = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetFileType() string {
	if o == nil || IsNil(o.FileType) {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetFileTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FileType) {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasFileType() bool {
	if o != nil && !IsNil(o.FileType) {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *SoftwarerepositoryCategoryMapper) SetFileType(v string) {
	o.FileType = &v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetImageType() string {
	if o == nil || IsNil(o.ImageType) {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetImageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ImageType) {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasImageType() bool {
	if o != nil && !IsNil(o.ImageType) {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *SoftwarerepositoryCategoryMapper) SetImageType(v string) {
	o.ImageType = &v
}

// GetIsNfsUpgradeSupported returns the IsNfsUpgradeSupported field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetIsNfsUpgradeSupported() bool {
	if o == nil || IsNil(o.IsNfsUpgradeSupported) {
		var ret bool
		return ret
	}
	return *o.IsNfsUpgradeSupported
}

// GetIsNfsUpgradeSupportedOk returns a tuple with the IsNfsUpgradeSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetIsNfsUpgradeSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNfsUpgradeSupported) {
		return nil, false
	}
	return o.IsNfsUpgradeSupported, true
}

// HasIsNfsUpgradeSupported returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasIsNfsUpgradeSupported() bool {
	if o != nil && !IsNil(o.IsNfsUpgradeSupported) {
		return true
	}

	return false
}

// SetIsNfsUpgradeSupported gets a reference to the given bool and assigns it to the IsNfsUpgradeSupported field.
func (o *SoftwarerepositoryCategoryMapper) SetIsNfsUpgradeSupported(v bool) {
	o.IsNfsUpgradeSupported = &v
}

// GetMdfId returns the MdfId field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetMdfId() string {
	if o == nil || IsNil(o.MdfId) {
		var ret string
		return ret
	}
	return *o.MdfId
}

// GetMdfIdOk returns a tuple with the MdfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetMdfIdOk() (*string, bool) {
	if o == nil || IsNil(o.MdfId) {
		return nil, false
	}
	return o.MdfId, true
}

// HasMdfId returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasMdfId() bool {
	if o != nil && !IsNil(o.MdfId) {
		return true
	}

	return false
}

// SetMdfId gets a reference to the given string and assigns it to the MdfId field.
func (o *SoftwarerepositoryCategoryMapper) SetMdfId(v string) {
	o.MdfId = &v
}

// GetRegexPattern returns the RegexPattern field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetRegexPattern() string {
	if o == nil || IsNil(o.RegexPattern) {
		var ret string
		return ret
	}
	return *o.RegexPattern
}

// GetRegexPatternOk returns a tuple with the RegexPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetRegexPatternOk() (*string, bool) {
	if o == nil || IsNil(o.RegexPattern) {
		return nil, false
	}
	return o.RegexPattern, true
}

// HasRegexPattern returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasRegexPattern() bool {
	if o != nil && !IsNil(o.RegexPattern) {
		return true
	}

	return false
}

// SetRegexPattern gets a reference to the given string and assigns it to the RegexPattern field.
func (o *SoftwarerepositoryCategoryMapper) SetRegexPattern(v string) {
	o.RegexPattern = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *SoftwarerepositoryCategoryMapper) SetSource(v string) {
	o.Source = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryCategoryMapper) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryCategoryMapper) GetSupportedModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedModels) {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasSupportedModels() bool {
	if o != nil && !IsNil(o.SupportedModels) {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *SoftwarerepositoryCategoryMapper) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

// GetSwId returns the SwId field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetSwId() string {
	if o == nil || IsNil(o.SwId) {
		var ret string
		return ret
	}
	return *o.SwId
}

// GetSwIdOk returns a tuple with the SwId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetSwIdOk() (*string, bool) {
	if o == nil || IsNil(o.SwId) {
		return nil, false
	}
	return o.SwId, true
}

// HasSwId returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasSwId() bool {
	if o != nil && !IsNil(o.SwId) {
		return true
	}

	return false
}

// SetSwId gets a reference to the given string and assigns it to the SwId field.
func (o *SoftwarerepositoryCategoryMapper) SetSwId(v string) {
	o.SwId = &v
}

// GetTagTypes returns the TagTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryCategoryMapper) GetTagTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagTypes
}

// GetTagTypesOk returns a tuple with the TagTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryCategoryMapper) GetTagTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.TagTypes) {
		return nil, false
	}
	return o.TagTypes, true
}

// HasTagTypes returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasTagTypes() bool {
	if o != nil && !IsNil(o.TagTypes) {
		return true
	}

	return false
}

// SetTagTypes gets a reference to the given []string and assigns it to the TagTypes field.
func (o *SoftwarerepositoryCategoryMapper) SetTagTypes(v []string) {
	o.TagTypes = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SoftwarerepositoryCategoryMapper) SetVersion(v string) {
	o.Version = &v
}

func (o SoftwarerepositoryCategoryMapper) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwarerepositoryCategoryMapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
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
	if !IsNil(o.FileType) {
		toSerialize["FileType"] = o.FileType
	}
	if !IsNil(o.ImageType) {
		toSerialize["ImageType"] = o.ImageType
	}
	if !IsNil(o.IsNfsUpgradeSupported) {
		toSerialize["IsNfsUpgradeSupported"] = o.IsNfsUpgradeSupported
	}
	if !IsNil(o.MdfId) {
		toSerialize["MdfId"] = o.MdfId
	}
	if !IsNil(o.RegexPattern) {
		toSerialize["RegexPattern"] = o.RegexPattern
	}
	if !IsNil(o.Source) {
		toSerialize["Source"] = o.Source
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}
	if !IsNil(o.SwId) {
		toSerialize["SwId"] = o.SwId
	}
	if o.TagTypes != nil {
		toSerialize["TagTypes"] = o.TagTypes
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SoftwarerepositoryCategoryMapper) UnmarshalJSON(data []byte) (err error) {
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
	type SoftwarerepositoryCategoryMapperWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The category of the model series.
		Category *string `json:"Category,omitempty"`
		// The type of distributable image, example huu, scu, driver, os. * `Distributable` - Stores firmware host utility images and fabric images. * `DriverDistributable` - Stores driver distributable images. * `ServerConfigurationUtilityDistributable` - Stores server configuration utility images. * `OperatingSystemFile` - Stores operating system iso images. * `HyperflexDistributable` - It stores HyperFlex images. * `HciDistributable` - It stores HCI images, such as bootstrap iso images.
		FileType *string `json:"FileType,omitempty"`
		// The type of image based on the endpoint it can upgrade. For example, ucs-c420m5-huu-3.2.1a.iso can upgrade standalone servers, so the image type is Standalone Server.
		ImageType *string `json:"ImageType,omitempty"`
		// Defines whether NFS firmware upgrade is supported with this image type.
		IsNfsUpgradeSupported *bool `json:"IsNfsUpgradeSupported,omitempty"`
		// Cisco software repository image category identifier.
		MdfId *string `json:"MdfId,omitempty"`
		// The regex that all images of this category follow.
		RegexPattern *string `json:"RegexPattern,omitempty"`
		// The image can be downloaded from cisco.com or external cloud store. * `Cisco` - External repository hosted on cisco.com. * `IntersightCloud` - Repository hosted by the Intersight Cloud. * `LocalMachine` - The file is available on the local client machine. Used as an upload source type. * `NetworkShare` - External repository in the customer datacenter. This will typically be a file server.
		Source          *string  `json:"Source,omitempty"`
		SupportedModels []string `json:"SupportedModels,omitempty"`
		// The software type id provided by cisco.com.
		SwId     *string  `json:"SwId,omitempty"`
		TagTypes []string `json:"TagTypes,omitempty"`
		// The version from which user can download images from amazon store, if source is external cloud store.
		Version *string `json:"Version,omitempty"`
	}

	varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct := SoftwarerepositoryCategoryMapperWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryCategoryMapper := _SoftwarerepositoryCategoryMapper{}
		varSoftwarerepositoryCategoryMapper.ClassId = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.ClassId
		varSoftwarerepositoryCategoryMapper.ObjectType = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.ObjectType
		varSoftwarerepositoryCategoryMapper.Category = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.Category
		varSoftwarerepositoryCategoryMapper.FileType = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.FileType
		varSoftwarerepositoryCategoryMapper.ImageType = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.ImageType
		varSoftwarerepositoryCategoryMapper.IsNfsUpgradeSupported = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.IsNfsUpgradeSupported
		varSoftwarerepositoryCategoryMapper.MdfId = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.MdfId
		varSoftwarerepositoryCategoryMapper.RegexPattern = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.RegexPattern
		varSoftwarerepositoryCategoryMapper.Source = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.Source
		varSoftwarerepositoryCategoryMapper.SupportedModels = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.SupportedModels
		varSoftwarerepositoryCategoryMapper.SwId = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.SwId
		varSoftwarerepositoryCategoryMapper.TagTypes = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.TagTypes
		varSoftwarerepositoryCategoryMapper.Version = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.Version
		*o = SoftwarerepositoryCategoryMapper(varSoftwarerepositoryCategoryMapper)
	} else {
		return err
	}

	varSoftwarerepositoryCategoryMapper := _SoftwarerepositoryCategoryMapper{}

	err = json.Unmarshal(data, &varSoftwarerepositoryCategoryMapper)
	if err == nil {
		o.CapabilityCapability = varSoftwarerepositoryCategoryMapper.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "FileType")
		delete(additionalProperties, "ImageType")
		delete(additionalProperties, "IsNfsUpgradeSupported")
		delete(additionalProperties, "MdfId")
		delete(additionalProperties, "RegexPattern")
		delete(additionalProperties, "Source")
		delete(additionalProperties, "SupportedModels")
		delete(additionalProperties, "SwId")
		delete(additionalProperties, "TagTypes")
		delete(additionalProperties, "Version")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableSoftwarerepositoryCategoryMapper struct {
	value *SoftwarerepositoryCategoryMapper
	isSet bool
}

func (v NullableSoftwarerepositoryCategoryMapper) Get() *SoftwarerepositoryCategoryMapper {
	return v.value
}

func (v *NullableSoftwarerepositoryCategoryMapper) Set(val *SoftwarerepositoryCategoryMapper) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCategoryMapper) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCategoryMapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCategoryMapper(val *SoftwarerepositoryCategoryMapper) *NullableSoftwarerepositoryCategoryMapper {
	return &NullableSoftwarerepositoryCategoryMapper{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCategoryMapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCategoryMapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
