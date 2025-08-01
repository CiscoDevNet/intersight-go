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

// checks if the AssetVmHost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetVmHost{}

// AssetVmHost Type for saving VM host details.
type AssetVmHost struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Reference to virtualization target account ID.
	AccountMoid *string `json:"AccountMoid,omitempty"`
	// Reference to virtualization cluster identity.
	ClusterIdentity *string `json:"ClusterIdentity,omitempty"`
	// Reference to virtualization cluster ID.
	ClusterMoid *string `json:"ClusterMoid,omitempty"`
	// Reference to virtualization cluster name.
	ClusterName          *string `json:"ClusterName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetVmHost AssetVmHost

// NewAssetVmHost instantiates a new AssetVmHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetVmHost(classId string, objectType string) *AssetVmHost {
	this := AssetVmHost{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetVmHostWithDefaults instantiates a new AssetVmHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetVmHostWithDefaults() *AssetVmHost {
	this := AssetVmHost{}
	var classId string = "asset.VmHost"
	this.ClassId = classId
	var objectType string = "asset.VmHost"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetVmHost) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetVmHost) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetVmHost) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "asset.VmHost" of the ClassId field.
func (o *AssetVmHost) GetDefaultClassId() interface{} {
	return "asset.VmHost"
}

// GetObjectType returns the ObjectType field value
func (o *AssetVmHost) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetVmHost) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetVmHost) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "asset.VmHost" of the ObjectType field.
func (o *AssetVmHost) GetDefaultObjectType() interface{} {
	return "asset.VmHost"
}

// GetAccountMoid returns the AccountMoid field value if set, zero value otherwise.
func (o *AssetVmHost) GetAccountMoid() string {
	if o == nil || IsNil(o.AccountMoid) {
		var ret string
		return ret
	}
	return *o.AccountMoid
}

// GetAccountMoidOk returns a tuple with the AccountMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetVmHost) GetAccountMoidOk() (*string, bool) {
	if o == nil || IsNil(o.AccountMoid) {
		return nil, false
	}
	return o.AccountMoid, true
}

// HasAccountMoid returns a boolean if a field has been set.
func (o *AssetVmHost) HasAccountMoid() bool {
	if o != nil && !IsNil(o.AccountMoid) {
		return true
	}

	return false
}

// SetAccountMoid gets a reference to the given string and assigns it to the AccountMoid field.
func (o *AssetVmHost) SetAccountMoid(v string) {
	o.AccountMoid = &v
}

// GetClusterIdentity returns the ClusterIdentity field value if set, zero value otherwise.
func (o *AssetVmHost) GetClusterIdentity() string {
	if o == nil || IsNil(o.ClusterIdentity) {
		var ret string
		return ret
	}
	return *o.ClusterIdentity
}

// GetClusterIdentityOk returns a tuple with the ClusterIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetVmHost) GetClusterIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterIdentity) {
		return nil, false
	}
	return o.ClusterIdentity, true
}

// HasClusterIdentity returns a boolean if a field has been set.
func (o *AssetVmHost) HasClusterIdentity() bool {
	if o != nil && !IsNil(o.ClusterIdentity) {
		return true
	}

	return false
}

// SetClusterIdentity gets a reference to the given string and assigns it to the ClusterIdentity field.
func (o *AssetVmHost) SetClusterIdentity(v string) {
	o.ClusterIdentity = &v
}

// GetClusterMoid returns the ClusterMoid field value if set, zero value otherwise.
func (o *AssetVmHost) GetClusterMoid() string {
	if o == nil || IsNil(o.ClusterMoid) {
		var ret string
		return ret
	}
	return *o.ClusterMoid
}

// GetClusterMoidOk returns a tuple with the ClusterMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetVmHost) GetClusterMoidOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterMoid) {
		return nil, false
	}
	return o.ClusterMoid, true
}

// HasClusterMoid returns a boolean if a field has been set.
func (o *AssetVmHost) HasClusterMoid() bool {
	if o != nil && !IsNil(o.ClusterMoid) {
		return true
	}

	return false
}

// SetClusterMoid gets a reference to the given string and assigns it to the ClusterMoid field.
func (o *AssetVmHost) SetClusterMoid(v string) {
	o.ClusterMoid = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *AssetVmHost) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetVmHost) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *AssetVmHost) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *AssetVmHost) SetClusterName(v string) {
	o.ClusterName = &v
}

func (o AssetVmHost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetVmHost) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AccountMoid) {
		toSerialize["AccountMoid"] = o.AccountMoid
	}
	if !IsNil(o.ClusterIdentity) {
		toSerialize["ClusterIdentity"] = o.ClusterIdentity
	}
	if !IsNil(o.ClusterMoid) {
		toSerialize["ClusterMoid"] = o.ClusterMoid
	}
	if !IsNil(o.ClusterName) {
		toSerialize["ClusterName"] = o.ClusterName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetVmHost) UnmarshalJSON(data []byte) (err error) {
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
	type AssetVmHostWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Reference to virtualization target account ID.
		AccountMoid *string `json:"AccountMoid,omitempty"`
		// Reference to virtualization cluster identity.
		ClusterIdentity *string `json:"ClusterIdentity,omitempty"`
		// Reference to virtualization cluster ID.
		ClusterMoid *string `json:"ClusterMoid,omitempty"`
		// Reference to virtualization cluster name.
		ClusterName *string `json:"ClusterName,omitempty"`
	}

	varAssetVmHostWithoutEmbeddedStruct := AssetVmHostWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetVmHostWithoutEmbeddedStruct)
	if err == nil {
		varAssetVmHost := _AssetVmHost{}
		varAssetVmHost.ClassId = varAssetVmHostWithoutEmbeddedStruct.ClassId
		varAssetVmHost.ObjectType = varAssetVmHostWithoutEmbeddedStruct.ObjectType
		varAssetVmHost.AccountMoid = varAssetVmHostWithoutEmbeddedStruct.AccountMoid
		varAssetVmHost.ClusterIdentity = varAssetVmHostWithoutEmbeddedStruct.ClusterIdentity
		varAssetVmHost.ClusterMoid = varAssetVmHostWithoutEmbeddedStruct.ClusterMoid
		varAssetVmHost.ClusterName = varAssetVmHostWithoutEmbeddedStruct.ClusterName
		*o = AssetVmHost(varAssetVmHost)
	} else {
		return err
	}

	varAssetVmHost := _AssetVmHost{}

	err = json.Unmarshal(data, &varAssetVmHost)
	if err == nil {
		o.MoBaseComplexType = varAssetVmHost.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccountMoid")
		delete(additionalProperties, "ClusterIdentity")
		delete(additionalProperties, "ClusterMoid")
		delete(additionalProperties, "ClusterName")

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

type NullableAssetVmHost struct {
	value *AssetVmHost
	isSet bool
}

func (v NullableAssetVmHost) Get() *AssetVmHost {
	return v.value
}

func (v *NullableAssetVmHost) Set(val *AssetVmHost) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetVmHost) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetVmHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetVmHost(val *AssetVmHost) *NullableAssetVmHost {
	return &NullableAssetVmHost{value: val, isSet: true}
}

func (v NullableAssetVmHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetVmHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
