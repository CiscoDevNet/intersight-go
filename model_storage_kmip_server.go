/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// StorageKmipServer Models the KMIP Server configuration used to fetch the encryption key.
type StorageKmipServer struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable the selected KMIP Server configuration for encryption. This flag just enables the drive security and only after remote key setting configured, the actual encryption will be done.
	EnableDriveSecurity *bool `json:"EnableDriveSecurity,omitempty"`
	// The IP address of the KMIP server. It could be an IPv4 address, an IPv6 address, or a hostname. Hostnames are valid only when Inband is configured for the CIMC address.
	IpAddress *string `json:"IpAddress,omitempty"`
	// The port to which the KMIP client should connect.
	Port *int64 `json:"Port,omitempty"`
	// The timeout before which the KMIP client should connect.
	Timeout              *int64 `json:"Timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageKmipServer StorageKmipServer

// NewStorageKmipServer instantiates a new StorageKmipServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageKmipServer(classId string, objectType string) *StorageKmipServer {
	this := StorageKmipServer{}
	this.ClassId = classId
	this.ObjectType = objectType
	var port int64 = 5696
	this.Port = &port
	var timeout int64 = 60
	this.Timeout = &timeout
	return &this
}

// NewStorageKmipServerWithDefaults instantiates a new StorageKmipServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageKmipServerWithDefaults() *StorageKmipServer {
	this := StorageKmipServer{}
	var classId string = "storage.KmipServer"
	this.ClassId = classId
	var objectType string = "storage.KmipServer"
	this.ObjectType = objectType
	var port int64 = 5696
	this.Port = &port
	var timeout int64 = 60
	this.Timeout = &timeout
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageKmipServer) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageKmipServer) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageKmipServer) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageKmipServer) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageKmipServer) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageKmipServer) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnableDriveSecurity returns the EnableDriveSecurity field value if set, zero value otherwise.
func (o *StorageKmipServer) GetEnableDriveSecurity() bool {
	if o == nil || o.EnableDriveSecurity == nil {
		var ret bool
		return ret
	}
	return *o.EnableDriveSecurity
}

// GetEnableDriveSecurityOk returns a tuple with the EnableDriveSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageKmipServer) GetEnableDriveSecurityOk() (*bool, bool) {
	if o == nil || o.EnableDriveSecurity == nil {
		return nil, false
	}
	return o.EnableDriveSecurity, true
}

// HasEnableDriveSecurity returns a boolean if a field has been set.
func (o *StorageKmipServer) HasEnableDriveSecurity() bool {
	if o != nil && o.EnableDriveSecurity != nil {
		return true
	}

	return false
}

// SetEnableDriveSecurity gets a reference to the given bool and assigns it to the EnableDriveSecurity field.
func (o *StorageKmipServer) SetEnableDriveSecurity(v bool) {
	o.EnableDriveSecurity = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *StorageKmipServer) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageKmipServer) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *StorageKmipServer) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *StorageKmipServer) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *StorageKmipServer) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageKmipServer) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *StorageKmipServer) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *StorageKmipServer) SetPort(v int64) {
	o.Port = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *StorageKmipServer) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageKmipServer) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *StorageKmipServer) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *StorageKmipServer) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o StorageKmipServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnableDriveSecurity != nil {
		toSerialize["EnableDriveSecurity"] = o.EnableDriveSecurity
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageKmipServer) UnmarshalJSON(bytes []byte) (err error) {
	type StorageKmipServerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enable the selected KMIP Server configuration for encryption. This flag just enables the drive security and only after remote key setting configured, the actual encryption will be done.
		EnableDriveSecurity *bool `json:"EnableDriveSecurity,omitempty"`
		// The IP address of the KMIP server. It could be an IPv4 address, an IPv6 address, or a hostname. Hostnames are valid only when Inband is configured for the CIMC address.
		IpAddress *string `json:"IpAddress,omitempty"`
		// The port to which the KMIP client should connect.
		Port *int64 `json:"Port,omitempty"`
		// The timeout before which the KMIP client should connect.
		Timeout *int64 `json:"Timeout,omitempty"`
	}

	varStorageKmipServerWithoutEmbeddedStruct := StorageKmipServerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageKmipServerWithoutEmbeddedStruct)
	if err == nil {
		varStorageKmipServer := _StorageKmipServer{}
		varStorageKmipServer.ClassId = varStorageKmipServerWithoutEmbeddedStruct.ClassId
		varStorageKmipServer.ObjectType = varStorageKmipServerWithoutEmbeddedStruct.ObjectType
		varStorageKmipServer.EnableDriveSecurity = varStorageKmipServerWithoutEmbeddedStruct.EnableDriveSecurity
		varStorageKmipServer.IpAddress = varStorageKmipServerWithoutEmbeddedStruct.IpAddress
		varStorageKmipServer.Port = varStorageKmipServerWithoutEmbeddedStruct.Port
		varStorageKmipServer.Timeout = varStorageKmipServerWithoutEmbeddedStruct.Timeout
		*o = StorageKmipServer(varStorageKmipServer)
	} else {
		return err
	}

	varStorageKmipServer := _StorageKmipServer{}

	err = json.Unmarshal(bytes, &varStorageKmipServer)
	if err == nil {
		o.MoBaseComplexType = varStorageKmipServer.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnableDriveSecurity")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Timeout")

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

type NullableStorageKmipServer struct {
	value *StorageKmipServer
	isSet bool
}

func (v NullableStorageKmipServer) Get() *StorageKmipServer {
	return v.value
}

func (v *NullableStorageKmipServer) Set(val *StorageKmipServer) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageKmipServer) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageKmipServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageKmipServer(val *StorageKmipServer) *NullableStorageKmipServer {
	return &NullableStorageKmipServer{value: val, isSet: true}
}

func (v NullableStorageKmipServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageKmipServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}