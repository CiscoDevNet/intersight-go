/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9235
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NetworkTelemetryCheck Checking configured Telementry data in endpoint.
type NetworkTelemetryCheck struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Failure status for telemetry configured.
	Status *string `json:"Status,omitempty"`
	// The telemetry configuration details from endpoint.
	TelemetryConfig      *string                              `json:"TelemetryConfig,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkTelemetryCheck NetworkTelemetryCheck

// NewNetworkTelemetryCheck instantiates a new NetworkTelemetryCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkTelemetryCheck(classId string, objectType string) *NetworkTelemetryCheck {
	this := NetworkTelemetryCheck{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkTelemetryCheckWithDefaults instantiates a new NetworkTelemetryCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkTelemetryCheckWithDefaults() *NetworkTelemetryCheck {
	this := NetworkTelemetryCheck{}
	var classId string = "network.TelemetryCheck"
	this.ClassId = classId
	var objectType string = "network.TelemetryCheck"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkTelemetryCheck) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkTelemetryCheck) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkTelemetryCheck) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkTelemetryCheck) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkTelemetryCheck) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkTelemetryCheck) SetObjectType(v string) {
	o.ObjectType = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NetworkTelemetryCheck) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTelemetryCheck) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NetworkTelemetryCheck) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NetworkTelemetryCheck) SetStatus(v string) {
	o.Status = &v
}

// GetTelemetryConfig returns the TelemetryConfig field value if set, zero value otherwise.
func (o *NetworkTelemetryCheck) GetTelemetryConfig() string {
	if o == nil || o.TelemetryConfig == nil {
		var ret string
		return ret
	}
	return *o.TelemetryConfig
}

// GetTelemetryConfigOk returns a tuple with the TelemetryConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTelemetryCheck) GetTelemetryConfigOk() (*string, bool) {
	if o == nil || o.TelemetryConfig == nil {
		return nil, false
	}
	return o.TelemetryConfig, true
}

// HasTelemetryConfig returns a boolean if a field has been set.
func (o *NetworkTelemetryCheck) HasTelemetryConfig() bool {
	if o != nil && o.TelemetryConfig != nil {
		return true
	}

	return false
}

// SetTelemetryConfig gets a reference to the given string and assigns it to the TelemetryConfig field.
func (o *NetworkTelemetryCheck) SetTelemetryConfig(v string) {
	o.TelemetryConfig = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkTelemetryCheck) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTelemetryCheck) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkTelemetryCheck) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkTelemetryCheck) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkTelemetryCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TelemetryConfig != nil {
		toSerialize["TelemetryConfig"] = o.TelemetryConfig
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkTelemetryCheck) UnmarshalJSON(bytes []byte) (err error) {
	type NetworkTelemetryCheckWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Failure status for telemetry configured.
		Status *string `json:"Status,omitempty"`
		// The telemetry configuration details from endpoint.
		TelemetryConfig  *string                              `json:"TelemetryConfig,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNetworkTelemetryCheckWithoutEmbeddedStruct := NetworkTelemetryCheckWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNetworkTelemetryCheckWithoutEmbeddedStruct)
	if err == nil {
		varNetworkTelemetryCheck := _NetworkTelemetryCheck{}
		varNetworkTelemetryCheck.ClassId = varNetworkTelemetryCheckWithoutEmbeddedStruct.ClassId
		varNetworkTelemetryCheck.ObjectType = varNetworkTelemetryCheckWithoutEmbeddedStruct.ObjectType
		varNetworkTelemetryCheck.Status = varNetworkTelemetryCheckWithoutEmbeddedStruct.Status
		varNetworkTelemetryCheck.TelemetryConfig = varNetworkTelemetryCheckWithoutEmbeddedStruct.TelemetryConfig
		varNetworkTelemetryCheck.RegisteredDevice = varNetworkTelemetryCheckWithoutEmbeddedStruct.RegisteredDevice
		*o = NetworkTelemetryCheck(varNetworkTelemetryCheck)
	} else {
		return err
	}

	varNetworkTelemetryCheck := _NetworkTelemetryCheck{}

	err = json.Unmarshal(bytes, &varNetworkTelemetryCheck)
	if err == nil {
		o.InventoryBase = varNetworkTelemetryCheck.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TelemetryConfig")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableNetworkTelemetryCheck struct {
	value *NetworkTelemetryCheck
	isSet bool
}

func (v NullableNetworkTelemetryCheck) Get() *NetworkTelemetryCheck {
	return v.value
}

func (v *NullableNetworkTelemetryCheck) Set(val *NetworkTelemetryCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkTelemetryCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkTelemetryCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkTelemetryCheck(val *NetworkTelemetryCheck) *NullableNetworkTelemetryCheck {
	return &NullableNetworkTelemetryCheck{value: val, isSet: true}
}

func (v NullableNetworkTelemetryCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkTelemetryCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}