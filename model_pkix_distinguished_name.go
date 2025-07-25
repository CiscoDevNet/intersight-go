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

// checks if the PkixDistinguishedName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PkixDistinguishedName{}

// PkixDistinguishedName The identifier for the owner of an X.509 certificate and the authority that issued the certificate.
type PkixDistinguishedName struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A required component that identifies a person or an object.
	CommonName           *string  `json:"CommonName,omitempty"`
	Country              []string `json:"Country,omitempty"`
	Locality             []string `json:"Locality,omitempty"`
	Organization         []string `json:"Organization,omitempty"`
	OrganizationalUnit   []string `json:"OrganizationalUnit,omitempty"`
	State                []string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PkixDistinguishedName PkixDistinguishedName

// NewPkixDistinguishedName instantiates a new PkixDistinguishedName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixDistinguishedName(classId string, objectType string) *PkixDistinguishedName {
	this := PkixDistinguishedName{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPkixDistinguishedNameWithDefaults instantiates a new PkixDistinguishedName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixDistinguishedNameWithDefaults() *PkixDistinguishedName {
	this := PkixDistinguishedName{}
	var classId string = "pkix.DistinguishedName"
	this.ClassId = classId
	var objectType string = "pkix.DistinguishedName"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PkixDistinguishedName) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PkixDistinguishedName) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PkixDistinguishedName) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "pkix.DistinguishedName" of the ClassId field.
func (o *PkixDistinguishedName) GetDefaultClassId() interface{} {
	return "pkix.DistinguishedName"
}

// GetObjectType returns the ObjectType field value
func (o *PkixDistinguishedName) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PkixDistinguishedName) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PkixDistinguishedName) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "pkix.DistinguishedName" of the ObjectType field.
func (o *PkixDistinguishedName) GetDefaultObjectType() interface{} {
	return "pkix.DistinguishedName"
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *PkixDistinguishedName) GetCommonName() string {
	if o == nil || IsNil(o.CommonName) {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixDistinguishedName) GetCommonNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommonName) {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *PkixDistinguishedName) HasCommonName() bool {
	if o != nil && !IsNil(o.CommonName) {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *PkixDistinguishedName) SetCommonName(v string) {
	o.CommonName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedName) GetCountry() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedName) GetCountryOk() ([]string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PkixDistinguishedName) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given []string and assigns it to the Country field.
func (o *PkixDistinguishedName) SetCountry(v []string) {
	o.Country = v
}

// GetLocality returns the Locality field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedName) GetLocality() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedName) GetLocalityOk() ([]string, bool) {
	if o == nil || IsNil(o.Locality) {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *PkixDistinguishedName) HasLocality() bool {
	if o != nil && !IsNil(o.Locality) {
		return true
	}

	return false
}

// SetLocality gets a reference to the given []string and assigns it to the Locality field.
func (o *PkixDistinguishedName) SetLocality(v []string) {
	o.Locality = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedName) GetOrganization() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedName) GetOrganizationOk() ([]string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PkixDistinguishedName) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given []string and assigns it to the Organization field.
func (o *PkixDistinguishedName) SetOrganization(v []string) {
	o.Organization = v
}

// GetOrganizationalUnit returns the OrganizationalUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedName) GetOrganizationalUnit() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OrganizationalUnit
}

// GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedName) GetOrganizationalUnitOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationalUnit) {
		return nil, false
	}
	return o.OrganizationalUnit, true
}

// HasOrganizationalUnit returns a boolean if a field has been set.
func (o *PkixDistinguishedName) HasOrganizationalUnit() bool {
	if o != nil && !IsNil(o.OrganizationalUnit) {
		return true
	}

	return false
}

// SetOrganizationalUnit gets a reference to the given []string and assigns it to the OrganizationalUnit field.
func (o *PkixDistinguishedName) SetOrganizationalUnit(v []string) {
	o.OrganizationalUnit = v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedName) GetState() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedName) GetStateOk() ([]string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PkixDistinguishedName) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given []string and assigns it to the State field.
func (o *PkixDistinguishedName) SetState(v []string) {
	o.State = v
}

func (o PkixDistinguishedName) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PkixDistinguishedName) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CommonName) {
		toSerialize["CommonName"] = o.CommonName
	}
	if o.Country != nil {
		toSerialize["Country"] = o.Country
	}
	if o.Locality != nil {
		toSerialize["Locality"] = o.Locality
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.OrganizationalUnit != nil {
		toSerialize["OrganizationalUnit"] = o.OrganizationalUnit
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PkixDistinguishedName) UnmarshalJSON(data []byte) (err error) {
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
	type PkixDistinguishedNameWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// A required component that identifies a person or an object.
		CommonName         *string  `json:"CommonName,omitempty"`
		Country            []string `json:"Country,omitempty"`
		Locality           []string `json:"Locality,omitempty"`
		Organization       []string `json:"Organization,omitempty"`
		OrganizationalUnit []string `json:"OrganizationalUnit,omitempty"`
		State              []string `json:"State,omitempty"`
	}

	varPkixDistinguishedNameWithoutEmbeddedStruct := PkixDistinguishedNameWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPkixDistinguishedNameWithoutEmbeddedStruct)
	if err == nil {
		varPkixDistinguishedName := _PkixDistinguishedName{}
		varPkixDistinguishedName.ClassId = varPkixDistinguishedNameWithoutEmbeddedStruct.ClassId
		varPkixDistinguishedName.ObjectType = varPkixDistinguishedNameWithoutEmbeddedStruct.ObjectType
		varPkixDistinguishedName.CommonName = varPkixDistinguishedNameWithoutEmbeddedStruct.CommonName
		varPkixDistinguishedName.Country = varPkixDistinguishedNameWithoutEmbeddedStruct.Country
		varPkixDistinguishedName.Locality = varPkixDistinguishedNameWithoutEmbeddedStruct.Locality
		varPkixDistinguishedName.Organization = varPkixDistinguishedNameWithoutEmbeddedStruct.Organization
		varPkixDistinguishedName.OrganizationalUnit = varPkixDistinguishedNameWithoutEmbeddedStruct.OrganizationalUnit
		varPkixDistinguishedName.State = varPkixDistinguishedNameWithoutEmbeddedStruct.State
		*o = PkixDistinguishedName(varPkixDistinguishedName)
	} else {
		return err
	}

	varPkixDistinguishedName := _PkixDistinguishedName{}

	err = json.Unmarshal(data, &varPkixDistinguishedName)
	if err == nil {
		o.MoBaseComplexType = varPkixDistinguishedName.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CommonName")
		delete(additionalProperties, "Country")
		delete(additionalProperties, "Locality")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "OrganizationalUnit")
		delete(additionalProperties, "State")

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

type NullablePkixDistinguishedName struct {
	value *PkixDistinguishedName
	isSet bool
}

func (v NullablePkixDistinguishedName) Get() *PkixDistinguishedName {
	return v.value
}

func (v *NullablePkixDistinguishedName) Set(val *PkixDistinguishedName) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixDistinguishedName) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixDistinguishedName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixDistinguishedName(val *PkixDistinguishedName) *NullablePkixDistinguishedName {
	return &NullablePkixDistinguishedName{value: val, isSet: true}
}

func (v NullablePkixDistinguishedName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixDistinguishedName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
