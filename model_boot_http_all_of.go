/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-13010
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// BootHttpAllOf Definition of the list of properties defined in 'boot.Http', excluding properties defined in parent classes.
type BootHttpAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the underlying virtual ethernet interface used by the HTTP boot device.
	InterfaceName *string `json:"InterfaceName,omitempty"`
	// Lists the supported Interface Source for HTTP device. Supported values are \"name\" and \"mac\". * `name` - Use interface name to select virtual ethernet interface. * `mac` - Use MAC address to select virtual ethernet interface. * `port` - Use port to select virtual ethernet interface.
	InterfaceSource *string `json:"InterfaceSource,omitempty"`
	// The IP config type to use during the HTTP boot process. For DHCP configuration, the IP address, DNS server, netmask and gateway details are obtained from DHCP server. For static configuration, please provide the IP address, DNS server, netmask, and gateway details. * `DHCP` - The type of the IP config is DHCP. * `Static` - The type of the IP config is Static.
	IpConfigType *string `json:"IpConfigType,omitempty"`
	// The IP address family type to use during the HTTP boot process. * `IPv4` - The type of the IP address is IPv4. * `IPv6` - The type of the IP address is IPv6.
	IpType *string `json:"IpType,omitempty"`
	// The MAC Address of the underlying virtual ethernet interface used by the HTTP boot device.
	MacAddress *string `json:"MacAddress,omitempty"`
	// The Port ID of the adapter on which the underlying virtual ethernet interface is present. If no port is specified, the default value is -1. Supported values are 0 to 255.
	Port *int64 `json:"Port,omitempty"`
	// Protocol to be used for HTTP boot. HTTPS require root certificate for authentication. * `HTTPS` - Secure HTTP protocol, certificate required for authentication. * `HTTP` - HTTP protocol without security certificate requirement.
	Protocol *string `json:"Protocol,omitempty"`
	// The slot ID of the adapter on which the underlying virtual ethernet interface is present. Supported values are ( 1 - 255, \"MLOM\", \"L\", \"L1\", \"L2\", \"OCP\").
	Slot               *string                        `json:"Slot,omitempty"`
	StaticIpV4Settings NullableBootStaticIpV4Settings `json:"StaticIpV4Settings,omitempty"`
	StaticIpV6Settings NullableBootStaticIpV6Settings `json:"StaticIpV6Settings,omitempty"`
	// Boot resource location in URI format.
	Uri                  *string `json:"Uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootHttpAllOf BootHttpAllOf

// NewBootHttpAllOf instantiates a new BootHttpAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootHttpAllOf(classId string, objectType string) *BootHttpAllOf {
	this := BootHttpAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var interfaceSource string = "name"
	this.InterfaceSource = &interfaceSource
	var ipConfigType string = "DHCP"
	this.IpConfigType = &ipConfigType
	var ipType string = "IPv4"
	this.IpType = &ipType
	var port int64 = -1
	this.Port = &port
	var protocol string = "HTTPS"
	this.Protocol = &protocol
	return &this
}

// NewBootHttpAllOfWithDefaults instantiates a new BootHttpAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootHttpAllOfWithDefaults() *BootHttpAllOf {
	this := BootHttpAllOf{}
	var classId string = "boot.Http"
	this.ClassId = classId
	var objectType string = "boot.Http"
	this.ObjectType = objectType
	var interfaceSource string = "name"
	this.InterfaceSource = &interfaceSource
	var ipConfigType string = "DHCP"
	this.IpConfigType = &ipConfigType
	var ipType string = "IPv4"
	this.IpType = &ipType
	var port int64 = -1
	this.Port = &port
	var protocol string = "HTTPS"
	this.Protocol = &protocol
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootHttpAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootHttpAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootHttpAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootHttpAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootHttpAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootHttpAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *BootHttpAllOf) GetInterfaceName() string {
	if o == nil || o.InterfaceName == nil {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootHttpAllOf) GetInterfaceNameOk() (*string, bool) {
	if o == nil || o.InterfaceName == nil {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *BootHttpAllOf) HasInterfaceName() bool {
	if o != nil && o.InterfaceName != nil {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *BootHttpAllOf) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetInterfaceSource returns the InterfaceSource field value if set, zero value otherwise.
func (o *BootHttpAllOf) GetInterfaceSource() string {
	if o == nil || o.InterfaceSource == nil {
		var ret string
		return ret
	}
	return *o.InterfaceSource
}

// GetInterfaceSourceOk returns a tuple with the InterfaceSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootHttpAllOf) GetInterfaceSourceOk() (*string, bool) {
	if o == nil || o.InterfaceSource == nil {
		return nil, false
	}
	return o.InterfaceSource, true
}

// HasInterfaceSource returns a boolean if a field has been set.
func (o *BootHttpAllOf) HasInterfaceSource() bool {
	if o != nil && o.InterfaceSource != nil {
		return true
	}

	return false
}

// SetInterfaceSource gets a reference to the given string and assigns it to the InterfaceSource field.
func (o *BootHttpAllOf) SetInterfaceSource(v string) {
	o.InterfaceSource = &v
}

// GetIpConfigType returns the IpConfigType field value if set, zero value otherwise.
func (o *BootHttpAllOf) GetIpConfigType() string {
	if o == nil || o.IpConfigType == nil {
		var ret string
		return ret
	}
	return *o.IpConfigType
}

// GetIpConfigTypeOk returns a tuple with the IpConfigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootHttpAllOf) GetIpConfigTypeOk() (*string, bool) {
	if o == nil || o.IpConfigType == nil {
		return nil, false
	}
	return o.IpConfigType, true
}

// HasIpConfigType returns a boolean if a field has been set.
func (o *BootHttpAllOf) HasIpConfigType() bool {
	if o != nil && o.IpConfigType != nil {
		return true
	}

	return false
}

// SetIpConfigType gets a reference to the given string and assigns it to the IpConfigType field.
func (o *BootHttpAllOf) SetIpConfigType(v string) {
	o.IpConfigType = &v
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *BootHttpAllOf) GetIpType() string {
	if o == nil || o.IpType == nil {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootHttpAllOf) GetIpTypeOk() (*string, bool) {
	if o == nil || o.IpType == nil {
		return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *BootHttpAllOf) HasIpType() bool {
	if o != nil && o.IpType != nil {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *BootHttpAllOf) SetIpType(v string) {
	o.IpType = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *BootHttpAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootHttpAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *BootHttpAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *BootHttpAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *BootHttpAllOf) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootHttpAllOf) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *BootHttpAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *BootHttpAllOf) SetPort(v int64) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *BootHttpAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootHttpAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *BootHttpAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *BootHttpAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *BootHttpAllOf) GetSlot() string {
	if o == nil || o.Slot == nil {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootHttpAllOf) GetSlotOk() (*string, bool) {
	if o == nil || o.Slot == nil {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *BootHttpAllOf) HasSlot() bool {
	if o != nil && o.Slot != nil {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *BootHttpAllOf) SetSlot(v string) {
	o.Slot = &v
}

// GetStaticIpV4Settings returns the StaticIpV4Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootHttpAllOf) GetStaticIpV4Settings() BootStaticIpV4Settings {
	if o == nil || o.StaticIpV4Settings.Get() == nil {
		var ret BootStaticIpV4Settings
		return ret
	}
	return *o.StaticIpV4Settings.Get()
}

// GetStaticIpV4SettingsOk returns a tuple with the StaticIpV4Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootHttpAllOf) GetStaticIpV4SettingsOk() (*BootStaticIpV4Settings, bool) {
	if o == nil {
		return nil, false
	}
	return o.StaticIpV4Settings.Get(), o.StaticIpV4Settings.IsSet()
}

// HasStaticIpV4Settings returns a boolean if a field has been set.
func (o *BootHttpAllOf) HasStaticIpV4Settings() bool {
	if o != nil && o.StaticIpV4Settings.IsSet() {
		return true
	}

	return false
}

// SetStaticIpV4Settings gets a reference to the given NullableBootStaticIpV4Settings and assigns it to the StaticIpV4Settings field.
func (o *BootHttpAllOf) SetStaticIpV4Settings(v BootStaticIpV4Settings) {
	o.StaticIpV4Settings.Set(&v)
}

// SetStaticIpV4SettingsNil sets the value for StaticIpV4Settings to be an explicit nil
func (o *BootHttpAllOf) SetStaticIpV4SettingsNil() {
	o.StaticIpV4Settings.Set(nil)
}

// UnsetStaticIpV4Settings ensures that no value is present for StaticIpV4Settings, not even an explicit nil
func (o *BootHttpAllOf) UnsetStaticIpV4Settings() {
	o.StaticIpV4Settings.Unset()
}

// GetStaticIpV6Settings returns the StaticIpV6Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootHttpAllOf) GetStaticIpV6Settings() BootStaticIpV6Settings {
	if o == nil || o.StaticIpV6Settings.Get() == nil {
		var ret BootStaticIpV6Settings
		return ret
	}
	return *o.StaticIpV6Settings.Get()
}

// GetStaticIpV6SettingsOk returns a tuple with the StaticIpV6Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootHttpAllOf) GetStaticIpV6SettingsOk() (*BootStaticIpV6Settings, bool) {
	if o == nil {
		return nil, false
	}
	return o.StaticIpV6Settings.Get(), o.StaticIpV6Settings.IsSet()
}

// HasStaticIpV6Settings returns a boolean if a field has been set.
func (o *BootHttpAllOf) HasStaticIpV6Settings() bool {
	if o != nil && o.StaticIpV6Settings.IsSet() {
		return true
	}

	return false
}

// SetStaticIpV6Settings gets a reference to the given NullableBootStaticIpV6Settings and assigns it to the StaticIpV6Settings field.
func (o *BootHttpAllOf) SetStaticIpV6Settings(v BootStaticIpV6Settings) {
	o.StaticIpV6Settings.Set(&v)
}

// SetStaticIpV6SettingsNil sets the value for StaticIpV6Settings to be an explicit nil
func (o *BootHttpAllOf) SetStaticIpV6SettingsNil() {
	o.StaticIpV6Settings.Set(nil)
}

// UnsetStaticIpV6Settings ensures that no value is present for StaticIpV6Settings, not even an explicit nil
func (o *BootHttpAllOf) UnsetStaticIpV6Settings() {
	o.StaticIpV6Settings.Unset()
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *BootHttpAllOf) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootHttpAllOf) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BootHttpAllOf) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *BootHttpAllOf) SetUri(v string) {
	o.Uri = &v
}

func (o BootHttpAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InterfaceName != nil {
		toSerialize["InterfaceName"] = o.InterfaceName
	}
	if o.InterfaceSource != nil {
		toSerialize["InterfaceSource"] = o.InterfaceSource
	}
	if o.IpConfigType != nil {
		toSerialize["IpConfigType"] = o.IpConfigType
	}
	if o.IpType != nil {
		toSerialize["IpType"] = o.IpType
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.Slot != nil {
		toSerialize["Slot"] = o.Slot
	}
	if o.StaticIpV4Settings.IsSet() {
		toSerialize["StaticIpV4Settings"] = o.StaticIpV4Settings.Get()
	}
	if o.StaticIpV6Settings.IsSet() {
		toSerialize["StaticIpV6Settings"] = o.StaticIpV6Settings.Get()
	}
	if o.Uri != nil {
		toSerialize["Uri"] = o.Uri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootHttpAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBootHttpAllOf := _BootHttpAllOf{}

	if err = json.Unmarshal(bytes, &varBootHttpAllOf); err == nil {
		*o = BootHttpAllOf(varBootHttpAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InterfaceName")
		delete(additionalProperties, "InterfaceSource")
		delete(additionalProperties, "IpConfigType")
		delete(additionalProperties, "IpType")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "Slot")
		delete(additionalProperties, "StaticIpV4Settings")
		delete(additionalProperties, "StaticIpV6Settings")
		delete(additionalProperties, "Uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBootHttpAllOf struct {
	value *BootHttpAllOf
	isSet bool
}

func (v NullableBootHttpAllOf) Get() *BootHttpAllOf {
	return v.value
}

func (v *NullableBootHttpAllOf) Set(val *BootHttpAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBootHttpAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBootHttpAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootHttpAllOf(val *BootHttpAllOf) *NullableBootHttpAllOf {
	return &NullableBootHttpAllOf{value: val, isSet: true}
}

func (v NullableBootHttpAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootHttpAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}