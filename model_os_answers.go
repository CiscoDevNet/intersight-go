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

// checks if the OsAnswers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsAnswers{}

// OsAnswers This MO captures the values for the most common set of fields in OS specific answer files. The values provided in this MO are used to construct the OS specific answer files (kickstart, seed, unattended xml) by replacing the fields/placeholders in selected os.ConfigurationFile content with the values of this MO properties.
type OsAnswers struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string   `json:"ObjectType"`
	AlternateNameServers []string `json:"AlternateNameServers,omitempty"`
	// If the source of the answers is a static file, the content of the file is stored as value in this property. The value is mandatory only when the 'Source' property has been set to 'File'.
	AnswerFile *string `json:"AnswerFile,omitempty"`
	// Hostname to be configured for the server in the OS.
	Hostname *string `json:"Hostname,omitempty"`
	// IP configuration type. Values are Static or Dynamic configuration of IP. In case of static IP configuration, IP address, gateway and other details need to be populated. In case of dynamic the IP configuration is obtained dynamically from DHCP. * `static` - In case of static IP configuraton, provide the details such as IP address, netmask, and gateway. * `DHCP` - In case of dynamic IP configuration, the IP address, netmask and gateway detailsare obtained from DHCP.
	IpConfigType    *string                   `json:"IpConfigType,omitempty"`
	IpConfiguration NullableOsIpConfiguration `json:"IpConfiguration,omitempty"`
	// Indicates whether the value of the 'answerFile' property has been set.
	IsAnswerFileSet *bool `json:"IsAnswerFileSet,omitempty"`
	// Enable to indicate Root Password provided is encrypted.
	IsRootPasswordCrypted *bool `json:"IsRootPasswordCrypted,omitempty"`
	// Indicates whether the value of the 'rootPassword' property has been set.
	IsRootPasswordSet *bool `json:"IsRootPasswordSet,omitempty"`
	// IP address of the name server to be configured in the OS.
	Nameserver *string `json:"Nameserver,omitempty"`
	// Network Device where the IP address must be configured. Network Interface names and MAC address are supported. For SUSE Linux Enterprise Server, Network Interface name is a required input and if provided as a MAC address, A persistent interface name is binded to the MAC address and the interface name will be used for network configuration. Refer https://documentation.suse.com/sles/15-SP2/html/SLES-all/cha-configuration-installation-options.html#CreateProfile-Network-names.
	NetworkDevice *string `json:"NetworkDevice,omitempty"`
	// The product key to be used for a specific version of Windows installation.
	ProductKey *string `json:"ProductKey,omitempty"`
	// Password configured for the root / administrator user in the OS. You can enter a plain text or an encrypted password. Intersight encrypts the plaintext password. Enable the Encrypted Password option to provide an encrypted password. For more details on encrypting passwords, see Help Center.
	RootPassword *string `json:"RootPassword,omitempty"`
	// Answer values can be provided from three sources - Embedded in OS image, static file, or as placeholder values for an answer file template. Source of the answers is given as value, Embedded/File/Template. 'Embedded' option indicates that the answer file is embedded within the OS Image. 'File' option indicates that the answers are provided as a file. 'Template' indicates that the placeholders in the selected os.ConfigurationFile MO are replaced with values provided as os.Answers MO. * `None` - Indicates that answers is not sent and values must be populated from Install Template.   * `Embedded` - Indicates that the answer file is embedded within OS image. * `File` - Indicates that the answer file is a static content that has all thevalues populated. * `Template` - Indicates that the given answers are used to populate the answer filetemplate. The template allows the users to refer some server specificanswers as fields/placeholders and replace these placeholders with theactual values for each Server during OS installation using 'Answers' and'AdditionalParameters' properties in os.Install MO.The answer file templates can be created by users as os.ConfigurationFile objects.
	Source               *string `json:"Source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsAnswers OsAnswers

// NewOsAnswers instantiates a new OsAnswers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsAnswers(classId string, objectType string) *OsAnswers {
	this := OsAnswers{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ipConfigType string = "static"
	this.IpConfigType = &ipConfigType
	var source string = "None"
	this.Source = &source
	return &this
}

// NewOsAnswersWithDefaults instantiates a new OsAnswers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsAnswersWithDefaults() *OsAnswers {
	this := OsAnswers{}
	var classId string = "os.Answers"
	this.ClassId = classId
	var objectType string = "os.Answers"
	this.ObjectType = objectType
	var ipConfigType string = "static"
	this.IpConfigType = &ipConfigType
	var source string = "None"
	this.Source = &source
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsAnswers) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsAnswers) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "os.Answers" of the ClassId field.
func (o *OsAnswers) GetDefaultClassId() interface{} {
	return "os.Answers"
}

// GetObjectType returns the ObjectType field value
func (o *OsAnswers) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsAnswers) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "os.Answers" of the ObjectType field.
func (o *OsAnswers) GetDefaultObjectType() interface{} {
	return "os.Answers"
}

// GetAlternateNameServers returns the AlternateNameServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsAnswers) GetAlternateNameServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AlternateNameServers
}

// GetAlternateNameServersOk returns a tuple with the AlternateNameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsAnswers) GetAlternateNameServersOk() ([]string, bool) {
	if o == nil || IsNil(o.AlternateNameServers) {
		return nil, false
	}
	return o.AlternateNameServers, true
}

// HasAlternateNameServers returns a boolean if a field has been set.
func (o *OsAnswers) HasAlternateNameServers() bool {
	if o != nil && !IsNil(o.AlternateNameServers) {
		return true
	}

	return false
}

// SetAlternateNameServers gets a reference to the given []string and assigns it to the AlternateNameServers field.
func (o *OsAnswers) SetAlternateNameServers(v []string) {
	o.AlternateNameServers = v
}

// GetAnswerFile returns the AnswerFile field value if set, zero value otherwise.
func (o *OsAnswers) GetAnswerFile() string {
	if o == nil || IsNil(o.AnswerFile) {
		var ret string
		return ret
	}
	return *o.AnswerFile
}

// GetAnswerFileOk returns a tuple with the AnswerFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetAnswerFileOk() (*string, bool) {
	if o == nil || IsNil(o.AnswerFile) {
		return nil, false
	}
	return o.AnswerFile, true
}

// HasAnswerFile returns a boolean if a field has been set.
func (o *OsAnswers) HasAnswerFile() bool {
	if o != nil && !IsNil(o.AnswerFile) {
		return true
	}

	return false
}

// SetAnswerFile gets a reference to the given string and assigns it to the AnswerFile field.
func (o *OsAnswers) SetAnswerFile(v string) {
	o.AnswerFile = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *OsAnswers) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *OsAnswers) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *OsAnswers) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpConfigType returns the IpConfigType field value if set, zero value otherwise.
func (o *OsAnswers) GetIpConfigType() string {
	if o == nil || IsNil(o.IpConfigType) {
		var ret string
		return ret
	}
	return *o.IpConfigType
}

// GetIpConfigTypeOk returns a tuple with the IpConfigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetIpConfigTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IpConfigType) {
		return nil, false
	}
	return o.IpConfigType, true
}

// HasIpConfigType returns a boolean if a field has been set.
func (o *OsAnswers) HasIpConfigType() bool {
	if o != nil && !IsNil(o.IpConfigType) {
		return true
	}

	return false
}

// SetIpConfigType gets a reference to the given string and assigns it to the IpConfigType field.
func (o *OsAnswers) SetIpConfigType(v string) {
	o.IpConfigType = &v
}

// GetIpConfiguration returns the IpConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsAnswers) GetIpConfiguration() OsIpConfiguration {
	if o == nil || IsNil(o.IpConfiguration.Get()) {
		var ret OsIpConfiguration
		return ret
	}
	return *o.IpConfiguration.Get()
}

// GetIpConfigurationOk returns a tuple with the IpConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsAnswers) GetIpConfigurationOk() (*OsIpConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpConfiguration.Get(), o.IpConfiguration.IsSet()
}

// HasIpConfiguration returns a boolean if a field has been set.
func (o *OsAnswers) HasIpConfiguration() bool {
	if o != nil && o.IpConfiguration.IsSet() {
		return true
	}

	return false
}

// SetIpConfiguration gets a reference to the given NullableOsIpConfiguration and assigns it to the IpConfiguration field.
func (o *OsAnswers) SetIpConfiguration(v OsIpConfiguration) {
	o.IpConfiguration.Set(&v)
}

// SetIpConfigurationNil sets the value for IpConfiguration to be an explicit nil
func (o *OsAnswers) SetIpConfigurationNil() {
	o.IpConfiguration.Set(nil)
}

// UnsetIpConfiguration ensures that no value is present for IpConfiguration, not even an explicit nil
func (o *OsAnswers) UnsetIpConfiguration() {
	o.IpConfiguration.Unset()
}

// GetIsAnswerFileSet returns the IsAnswerFileSet field value if set, zero value otherwise.
func (o *OsAnswers) GetIsAnswerFileSet() bool {
	if o == nil || IsNil(o.IsAnswerFileSet) {
		var ret bool
		return ret
	}
	return *o.IsAnswerFileSet
}

// GetIsAnswerFileSetOk returns a tuple with the IsAnswerFileSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetIsAnswerFileSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAnswerFileSet) {
		return nil, false
	}
	return o.IsAnswerFileSet, true
}

// HasIsAnswerFileSet returns a boolean if a field has been set.
func (o *OsAnswers) HasIsAnswerFileSet() bool {
	if o != nil && !IsNil(o.IsAnswerFileSet) {
		return true
	}

	return false
}

// SetIsAnswerFileSet gets a reference to the given bool and assigns it to the IsAnswerFileSet field.
func (o *OsAnswers) SetIsAnswerFileSet(v bool) {
	o.IsAnswerFileSet = &v
}

// GetIsRootPasswordCrypted returns the IsRootPasswordCrypted field value if set, zero value otherwise.
func (o *OsAnswers) GetIsRootPasswordCrypted() bool {
	if o == nil || IsNil(o.IsRootPasswordCrypted) {
		var ret bool
		return ret
	}
	return *o.IsRootPasswordCrypted
}

// GetIsRootPasswordCryptedOk returns a tuple with the IsRootPasswordCrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetIsRootPasswordCryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRootPasswordCrypted) {
		return nil, false
	}
	return o.IsRootPasswordCrypted, true
}

// HasIsRootPasswordCrypted returns a boolean if a field has been set.
func (o *OsAnswers) HasIsRootPasswordCrypted() bool {
	if o != nil && !IsNil(o.IsRootPasswordCrypted) {
		return true
	}

	return false
}

// SetIsRootPasswordCrypted gets a reference to the given bool and assigns it to the IsRootPasswordCrypted field.
func (o *OsAnswers) SetIsRootPasswordCrypted(v bool) {
	o.IsRootPasswordCrypted = &v
}

// GetIsRootPasswordSet returns the IsRootPasswordSet field value if set, zero value otherwise.
func (o *OsAnswers) GetIsRootPasswordSet() bool {
	if o == nil || IsNil(o.IsRootPasswordSet) {
		var ret bool
		return ret
	}
	return *o.IsRootPasswordSet
}

// GetIsRootPasswordSetOk returns a tuple with the IsRootPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetIsRootPasswordSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRootPasswordSet) {
		return nil, false
	}
	return o.IsRootPasswordSet, true
}

// HasIsRootPasswordSet returns a boolean if a field has been set.
func (o *OsAnswers) HasIsRootPasswordSet() bool {
	if o != nil && !IsNil(o.IsRootPasswordSet) {
		return true
	}

	return false
}

// SetIsRootPasswordSet gets a reference to the given bool and assigns it to the IsRootPasswordSet field.
func (o *OsAnswers) SetIsRootPasswordSet(v bool) {
	o.IsRootPasswordSet = &v
}

// GetNameserver returns the Nameserver field value if set, zero value otherwise.
func (o *OsAnswers) GetNameserver() string {
	if o == nil || IsNil(o.Nameserver) {
		var ret string
		return ret
	}
	return *o.Nameserver
}

// GetNameserverOk returns a tuple with the Nameserver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetNameserverOk() (*string, bool) {
	if o == nil || IsNil(o.Nameserver) {
		return nil, false
	}
	return o.Nameserver, true
}

// HasNameserver returns a boolean if a field has been set.
func (o *OsAnswers) HasNameserver() bool {
	if o != nil && !IsNil(o.Nameserver) {
		return true
	}

	return false
}

// SetNameserver gets a reference to the given string and assigns it to the Nameserver field.
func (o *OsAnswers) SetNameserver(v string) {
	o.Nameserver = &v
}

// GetNetworkDevice returns the NetworkDevice field value if set, zero value otherwise.
func (o *OsAnswers) GetNetworkDevice() string {
	if o == nil || IsNil(o.NetworkDevice) {
		var ret string
		return ret
	}
	return *o.NetworkDevice
}

// GetNetworkDeviceOk returns a tuple with the NetworkDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetNetworkDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkDevice) {
		return nil, false
	}
	return o.NetworkDevice, true
}

// HasNetworkDevice returns a boolean if a field has been set.
func (o *OsAnswers) HasNetworkDevice() bool {
	if o != nil && !IsNil(o.NetworkDevice) {
		return true
	}

	return false
}

// SetNetworkDevice gets a reference to the given string and assigns it to the NetworkDevice field.
func (o *OsAnswers) SetNetworkDevice(v string) {
	o.NetworkDevice = &v
}

// GetProductKey returns the ProductKey field value if set, zero value otherwise.
func (o *OsAnswers) GetProductKey() string {
	if o == nil || IsNil(o.ProductKey) {
		var ret string
		return ret
	}
	return *o.ProductKey
}

// GetProductKeyOk returns a tuple with the ProductKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetProductKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProductKey) {
		return nil, false
	}
	return o.ProductKey, true
}

// HasProductKey returns a boolean if a field has been set.
func (o *OsAnswers) HasProductKey() bool {
	if o != nil && !IsNil(o.ProductKey) {
		return true
	}

	return false
}

// SetProductKey gets a reference to the given string and assigns it to the ProductKey field.
func (o *OsAnswers) SetProductKey(v string) {
	o.ProductKey = &v
}

// GetRootPassword returns the RootPassword field value if set, zero value otherwise.
func (o *OsAnswers) GetRootPassword() string {
	if o == nil || IsNil(o.RootPassword) {
		var ret string
		return ret
	}
	return *o.RootPassword
}

// GetRootPasswordOk returns a tuple with the RootPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetRootPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.RootPassword) {
		return nil, false
	}
	return o.RootPassword, true
}

// HasRootPassword returns a boolean if a field has been set.
func (o *OsAnswers) HasRootPassword() bool {
	if o != nil && !IsNil(o.RootPassword) {
		return true
	}

	return false
}

// SetRootPassword gets a reference to the given string and assigns it to the RootPassword field.
func (o *OsAnswers) SetRootPassword(v string) {
	o.RootPassword = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *OsAnswers) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswers) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *OsAnswers) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *OsAnswers) SetSource(v string) {
	o.Source = &v
}

func (o OsAnswers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsAnswers) ToMap() (map[string]interface{}, error) {
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
	if o.AlternateNameServers != nil {
		toSerialize["AlternateNameServers"] = o.AlternateNameServers
	}
	if !IsNil(o.AnswerFile) {
		toSerialize["AnswerFile"] = o.AnswerFile
	}
	if !IsNil(o.Hostname) {
		toSerialize["Hostname"] = o.Hostname
	}
	if !IsNil(o.IpConfigType) {
		toSerialize["IpConfigType"] = o.IpConfigType
	}
	if o.IpConfiguration.IsSet() {
		toSerialize["IpConfiguration"] = o.IpConfiguration.Get()
	}
	if !IsNil(o.IsAnswerFileSet) {
		toSerialize["IsAnswerFileSet"] = o.IsAnswerFileSet
	}
	if !IsNil(o.IsRootPasswordCrypted) {
		toSerialize["IsRootPasswordCrypted"] = o.IsRootPasswordCrypted
	}
	if !IsNil(o.IsRootPasswordSet) {
		toSerialize["IsRootPasswordSet"] = o.IsRootPasswordSet
	}
	if !IsNil(o.Nameserver) {
		toSerialize["Nameserver"] = o.Nameserver
	}
	if !IsNil(o.NetworkDevice) {
		toSerialize["NetworkDevice"] = o.NetworkDevice
	}
	if !IsNil(o.ProductKey) {
		toSerialize["ProductKey"] = o.ProductKey
	}
	if !IsNil(o.RootPassword) {
		toSerialize["RootPassword"] = o.RootPassword
	}
	if !IsNil(o.Source) {
		toSerialize["Source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OsAnswers) UnmarshalJSON(data []byte) (err error) {
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
	type OsAnswersWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string   `json:"ObjectType"`
		AlternateNameServers []string `json:"AlternateNameServers,omitempty"`
		// If the source of the answers is a static file, the content of the file is stored as value in this property. The value is mandatory only when the 'Source' property has been set to 'File'.
		AnswerFile *string `json:"AnswerFile,omitempty"`
		// Hostname to be configured for the server in the OS.
		Hostname *string `json:"Hostname,omitempty"`
		// IP configuration type. Values are Static or Dynamic configuration of IP. In case of static IP configuration, IP address, gateway and other details need to be populated. In case of dynamic the IP configuration is obtained dynamically from DHCP. * `static` - In case of static IP configuraton, provide the details such as IP address, netmask, and gateway. * `DHCP` - In case of dynamic IP configuration, the IP address, netmask and gateway detailsare obtained from DHCP.
		IpConfigType    *string                   `json:"IpConfigType,omitempty"`
		IpConfiguration NullableOsIpConfiguration `json:"IpConfiguration,omitempty"`
		// Indicates whether the value of the 'answerFile' property has been set.
		IsAnswerFileSet *bool `json:"IsAnswerFileSet,omitempty"`
		// Enable to indicate Root Password provided is encrypted.
		IsRootPasswordCrypted *bool `json:"IsRootPasswordCrypted,omitempty"`
		// Indicates whether the value of the 'rootPassword' property has been set.
		IsRootPasswordSet *bool `json:"IsRootPasswordSet,omitempty"`
		// IP address of the name server to be configured in the OS.
		Nameserver *string `json:"Nameserver,omitempty"`
		// Network Device where the IP address must be configured. Network Interface names and MAC address are supported. For SUSE Linux Enterprise Server, Network Interface name is a required input and if provided as a MAC address, A persistent interface name is binded to the MAC address and the interface name will be used for network configuration. Refer https://documentation.suse.com/sles/15-SP2/html/SLES-all/cha-configuration-installation-options.html#CreateProfile-Network-names.
		NetworkDevice *string `json:"NetworkDevice,omitempty"`
		// The product key to be used for a specific version of Windows installation.
		ProductKey *string `json:"ProductKey,omitempty"`
		// Password configured for the root / administrator user in the OS. You can enter a plain text or an encrypted password. Intersight encrypts the plaintext password. Enable the Encrypted Password option to provide an encrypted password. For more details on encrypting passwords, see Help Center.
		RootPassword *string `json:"RootPassword,omitempty"`
		// Answer values can be provided from three sources - Embedded in OS image, static file, or as placeholder values for an answer file template. Source of the answers is given as value, Embedded/File/Template. 'Embedded' option indicates that the answer file is embedded within the OS Image. 'File' option indicates that the answers are provided as a file. 'Template' indicates that the placeholders in the selected os.ConfigurationFile MO are replaced with values provided as os.Answers MO. * `None` - Indicates that answers is not sent and values must be populated from Install Template.   * `Embedded` - Indicates that the answer file is embedded within OS image. * `File` - Indicates that the answer file is a static content that has all thevalues populated. * `Template` - Indicates that the given answers are used to populate the answer filetemplate. The template allows the users to refer some server specificanswers as fields/placeholders and replace these placeholders with theactual values for each Server during OS installation using 'Answers' and'AdditionalParameters' properties in os.Install MO.The answer file templates can be created by users as os.ConfigurationFile objects.
		Source *string `json:"Source,omitempty"`
	}

	varOsAnswersWithoutEmbeddedStruct := OsAnswersWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOsAnswersWithoutEmbeddedStruct)
	if err == nil {
		varOsAnswers := _OsAnswers{}
		varOsAnswers.ClassId = varOsAnswersWithoutEmbeddedStruct.ClassId
		varOsAnswers.ObjectType = varOsAnswersWithoutEmbeddedStruct.ObjectType
		varOsAnswers.AlternateNameServers = varOsAnswersWithoutEmbeddedStruct.AlternateNameServers
		varOsAnswers.AnswerFile = varOsAnswersWithoutEmbeddedStruct.AnswerFile
		varOsAnswers.Hostname = varOsAnswersWithoutEmbeddedStruct.Hostname
		varOsAnswers.IpConfigType = varOsAnswersWithoutEmbeddedStruct.IpConfigType
		varOsAnswers.IpConfiguration = varOsAnswersWithoutEmbeddedStruct.IpConfiguration
		varOsAnswers.IsAnswerFileSet = varOsAnswersWithoutEmbeddedStruct.IsAnswerFileSet
		varOsAnswers.IsRootPasswordCrypted = varOsAnswersWithoutEmbeddedStruct.IsRootPasswordCrypted
		varOsAnswers.IsRootPasswordSet = varOsAnswersWithoutEmbeddedStruct.IsRootPasswordSet
		varOsAnswers.Nameserver = varOsAnswersWithoutEmbeddedStruct.Nameserver
		varOsAnswers.NetworkDevice = varOsAnswersWithoutEmbeddedStruct.NetworkDevice
		varOsAnswers.ProductKey = varOsAnswersWithoutEmbeddedStruct.ProductKey
		varOsAnswers.RootPassword = varOsAnswersWithoutEmbeddedStruct.RootPassword
		varOsAnswers.Source = varOsAnswersWithoutEmbeddedStruct.Source
		*o = OsAnswers(varOsAnswers)
	} else {
		return err
	}

	varOsAnswers := _OsAnswers{}

	err = json.Unmarshal(data, &varOsAnswers)
	if err == nil {
		o.MoBaseComplexType = varOsAnswers.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlternateNameServers")
		delete(additionalProperties, "AnswerFile")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "IpConfigType")
		delete(additionalProperties, "IpConfiguration")
		delete(additionalProperties, "IsAnswerFileSet")
		delete(additionalProperties, "IsRootPasswordCrypted")
		delete(additionalProperties, "IsRootPasswordSet")
		delete(additionalProperties, "Nameserver")
		delete(additionalProperties, "NetworkDevice")
		delete(additionalProperties, "ProductKey")
		delete(additionalProperties, "RootPassword")
		delete(additionalProperties, "Source")

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

type NullableOsAnswers struct {
	value *OsAnswers
	isSet bool
}

func (v NullableOsAnswers) Get() *OsAnswers {
	return v.value
}

func (v *NullableOsAnswers) Set(val *OsAnswers) {
	v.value = val
	v.isSet = true
}

func (v NullableOsAnswers) IsSet() bool {
	return v.isSet
}

func (v *NullableOsAnswers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsAnswers(val *OsAnswers) *NullableOsAnswers {
	return &NullableOsAnswers{value: val, isSet: true}
}

func (v NullableOsAnswers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsAnswers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
