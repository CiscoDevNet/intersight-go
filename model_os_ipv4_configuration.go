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

// OsIpv4Configuration In case of static IPv4 configuration, IP address, netmask and gateway details are provided.
type OsIpv4Configuration struct {
	OsIpConfiguration
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                    `json:"ObjectType"`
	IpV4Config           NullableCommIpV4Interface `json:"IpV4Config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsIpv4Configuration OsIpv4Configuration

// NewOsIpv4Configuration instantiates a new OsIpv4Configuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsIpv4Configuration(classId string, objectType string) *OsIpv4Configuration {
	this := OsIpv4Configuration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsIpv4ConfigurationWithDefaults instantiates a new OsIpv4Configuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsIpv4ConfigurationWithDefaults() *OsIpv4Configuration {
	this := OsIpv4Configuration{}
	var classId string = "os.Ipv4Configuration"
	this.ClassId = classId
	var objectType string = "os.Ipv4Configuration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsIpv4Configuration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsIpv4Configuration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsIpv4Configuration) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsIpv4Configuration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsIpv4Configuration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsIpv4Configuration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpV4Config returns the IpV4Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsIpv4Configuration) GetIpV4Config() CommIpV4Interface {
	if o == nil || o.IpV4Config.Get() == nil {
		var ret CommIpV4Interface
		return ret
	}
	return *o.IpV4Config.Get()
}

// GetIpV4ConfigOk returns a tuple with the IpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsIpv4Configuration) GetIpV4ConfigOk() (*CommIpV4Interface, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpV4Config.Get(), o.IpV4Config.IsSet()
}

// HasIpV4Config returns a boolean if a field has been set.
func (o *OsIpv4Configuration) HasIpV4Config() bool {
	if o != nil && o.IpV4Config.IsSet() {
		return true
	}

	return false
}

// SetIpV4Config gets a reference to the given NullableCommIpV4Interface and assigns it to the IpV4Config field.
func (o *OsIpv4Configuration) SetIpV4Config(v CommIpV4Interface) {
	o.IpV4Config.Set(&v)
}

// SetIpV4ConfigNil sets the value for IpV4Config to be an explicit nil
func (o *OsIpv4Configuration) SetIpV4ConfigNil() {
	o.IpV4Config.Set(nil)
}

// UnsetIpV4Config ensures that no value is present for IpV4Config, not even an explicit nil
func (o *OsIpv4Configuration) UnsetIpV4Config() {
	o.IpV4Config.Unset()
}

func (o OsIpv4Configuration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedOsIpConfiguration, errOsIpConfiguration := json.Marshal(o.OsIpConfiguration)
	if errOsIpConfiguration != nil {
		return []byte{}, errOsIpConfiguration
	}
	errOsIpConfiguration = json.Unmarshal([]byte(serializedOsIpConfiguration), &toSerialize)
	if errOsIpConfiguration != nil {
		return []byte{}, errOsIpConfiguration
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IpV4Config.IsSet() {
		toSerialize["IpV4Config"] = o.IpV4Config.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsIpv4Configuration) UnmarshalJSON(bytes []byte) (err error) {
	type OsIpv4ConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                    `json:"ObjectType"`
		IpV4Config NullableCommIpV4Interface `json:"IpV4Config,omitempty"`
	}

	varOsIpv4ConfigurationWithoutEmbeddedStruct := OsIpv4ConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOsIpv4ConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varOsIpv4Configuration := _OsIpv4Configuration{}
		varOsIpv4Configuration.ClassId = varOsIpv4ConfigurationWithoutEmbeddedStruct.ClassId
		varOsIpv4Configuration.ObjectType = varOsIpv4ConfigurationWithoutEmbeddedStruct.ObjectType
		varOsIpv4Configuration.IpV4Config = varOsIpv4ConfigurationWithoutEmbeddedStruct.IpV4Config
		*o = OsIpv4Configuration(varOsIpv4Configuration)
	} else {
		return err
	}

	varOsIpv4Configuration := _OsIpv4Configuration{}

	err = json.Unmarshal(bytes, &varOsIpv4Configuration)
	if err == nil {
		o.OsIpConfiguration = varOsIpv4Configuration.OsIpConfiguration
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpV4Config")

		// remove fields from embedded structs
		reflectOsIpConfiguration := reflect.ValueOf(o.OsIpConfiguration)
		for i := 0; i < reflectOsIpConfiguration.Type().NumField(); i++ {
			t := reflectOsIpConfiguration.Type().Field(i)

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

type NullableOsIpv4Configuration struct {
	value *OsIpv4Configuration
	isSet bool
}

func (v NullableOsIpv4Configuration) Get() *OsIpv4Configuration {
	return v.value
}

func (v *NullableOsIpv4Configuration) Set(val *OsIpv4Configuration) {
	v.value = val
	v.isSet = true
}

func (v NullableOsIpv4Configuration) IsSet() bool {
	return v.isSet
}

func (v *NullableOsIpv4Configuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsIpv4Configuration(val *OsIpv4Configuration) *NullableOsIpv4Configuration {
	return &NullableOsIpv4Configuration{value: val, isSet: true}
}

func (v NullableOsIpv4Configuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsIpv4Configuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}