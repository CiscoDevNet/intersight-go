/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14828
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NetworkVfc Vfc configured on a Fabric Interconnect.
type NetworkVfc struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description for the vHBA.
	Description *string `json:"Description,omitempty"`
	// Vfc Identifier on a Fabric Interconnect.
	VfcId                  *int64                               `json:"VfcId,omitempty"`
	AdapterHostFcInterface *AdapterHostFcInterfaceRelationship  `json:"AdapterHostFcInterface,omitempty"`
	NetworkElement         *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice       *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _NetworkVfc NetworkVfc

// NewNetworkVfc instantiates a new NetworkVfc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkVfc(classId string, objectType string) *NetworkVfc {
	this := NetworkVfc{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkVfcWithDefaults instantiates a new NetworkVfc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkVfcWithDefaults() *NetworkVfc {
	this := NetworkVfc{}
	var classId string = "network.Vfc"
	this.ClassId = classId
	var objectType string = "network.Vfc"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkVfc) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkVfc) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkVfc) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkVfc) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkVfc) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkVfc) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworkVfc) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVfc) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkVfc) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworkVfc) SetDescription(v string) {
	o.Description = &v
}

// GetVfcId returns the VfcId field value if set, zero value otherwise.
func (o *NetworkVfc) GetVfcId() int64 {
	if o == nil || o.VfcId == nil {
		var ret int64
		return ret
	}
	return *o.VfcId
}

// GetVfcIdOk returns a tuple with the VfcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVfc) GetVfcIdOk() (*int64, bool) {
	if o == nil || o.VfcId == nil {
		return nil, false
	}
	return o.VfcId, true
}

// HasVfcId returns a boolean if a field has been set.
func (o *NetworkVfc) HasVfcId() bool {
	if o != nil && o.VfcId != nil {
		return true
	}

	return false
}

// SetVfcId gets a reference to the given int64 and assigns it to the VfcId field.
func (o *NetworkVfc) SetVfcId(v int64) {
	o.VfcId = &v
}

// GetAdapterHostFcInterface returns the AdapterHostFcInterface field value if set, zero value otherwise.
func (o *NetworkVfc) GetAdapterHostFcInterface() AdapterHostFcInterfaceRelationship {
	if o == nil || o.AdapterHostFcInterface == nil {
		var ret AdapterHostFcInterfaceRelationship
		return ret
	}
	return *o.AdapterHostFcInterface
}

// GetAdapterHostFcInterfaceOk returns a tuple with the AdapterHostFcInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVfc) GetAdapterHostFcInterfaceOk() (*AdapterHostFcInterfaceRelationship, bool) {
	if o == nil || o.AdapterHostFcInterface == nil {
		return nil, false
	}
	return o.AdapterHostFcInterface, true
}

// HasAdapterHostFcInterface returns a boolean if a field has been set.
func (o *NetworkVfc) HasAdapterHostFcInterface() bool {
	if o != nil && o.AdapterHostFcInterface != nil {
		return true
	}

	return false
}

// SetAdapterHostFcInterface gets a reference to the given AdapterHostFcInterfaceRelationship and assigns it to the AdapterHostFcInterface field.
func (o *NetworkVfc) SetAdapterHostFcInterface(v AdapterHostFcInterfaceRelationship) {
	o.AdapterHostFcInterface = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NetworkVfc) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVfc) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkVfc) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkVfc) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkVfc) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVfc) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkVfc) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkVfc) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkVfc) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.VfcId != nil {
		toSerialize["VfcId"] = o.VfcId
	}
	if o.AdapterHostFcInterface != nil {
		toSerialize["AdapterHostFcInterface"] = o.AdapterHostFcInterface
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkVfc) UnmarshalJSON(bytes []byte) (err error) {
	type NetworkVfcWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description for the vHBA.
		Description *string `json:"Description,omitempty"`
		// Vfc Identifier on a Fabric Interconnect.
		VfcId                  *int64                               `json:"VfcId,omitempty"`
		AdapterHostFcInterface *AdapterHostFcInterfaceRelationship  `json:"AdapterHostFcInterface,omitempty"`
		NetworkElement         *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice       *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNetworkVfcWithoutEmbeddedStruct := NetworkVfcWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNetworkVfcWithoutEmbeddedStruct)
	if err == nil {
		varNetworkVfc := _NetworkVfc{}
		varNetworkVfc.ClassId = varNetworkVfcWithoutEmbeddedStruct.ClassId
		varNetworkVfc.ObjectType = varNetworkVfcWithoutEmbeddedStruct.ObjectType
		varNetworkVfc.Description = varNetworkVfcWithoutEmbeddedStruct.Description
		varNetworkVfc.VfcId = varNetworkVfcWithoutEmbeddedStruct.VfcId
		varNetworkVfc.AdapterHostFcInterface = varNetworkVfcWithoutEmbeddedStruct.AdapterHostFcInterface
		varNetworkVfc.NetworkElement = varNetworkVfcWithoutEmbeddedStruct.NetworkElement
		varNetworkVfc.RegisteredDevice = varNetworkVfcWithoutEmbeddedStruct.RegisteredDevice
		*o = NetworkVfc(varNetworkVfc)
	} else {
		return err
	}

	varNetworkVfc := _NetworkVfc{}

	err = json.Unmarshal(bytes, &varNetworkVfc)
	if err == nil {
		o.InventoryBase = varNetworkVfc.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "VfcId")
		delete(additionalProperties, "AdapterHostFcInterface")
		delete(additionalProperties, "NetworkElement")
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

type NullableNetworkVfc struct {
	value *NetworkVfc
	isSet bool
}

func (v NullableNetworkVfc) Get() *NetworkVfc {
	return v.value
}

func (v *NullableNetworkVfc) Set(val *NetworkVfc) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkVfc) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkVfc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkVfc(val *NetworkVfc) *NullableNetworkVfc {
	return &NullableNetworkVfc{value: val, isSet: true}
}

func (v NullableNetworkVfc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkVfc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}