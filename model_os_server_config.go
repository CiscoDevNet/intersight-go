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

// checks if the OsServerConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsServerConfig{}

// OsServerConfig The server level OS install parameter for the uploaded CSV file.
type OsServerConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string            `json:"ObjectType"`
	AdditionalParameters []OsPlaceHolder   `json:"AdditionalParameters,omitempty"`
	Answers              NullableOsAnswers `json:"Answers,omitempty"`
	ErrorMsgs            []string          `json:"ErrorMsgs,omitempty"`
	// The WWPN Address of the underlying fibre channel interface at the host side used for SAN accesss. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 20:00:D4:C9:3C:35:02:01.
	InitiatorWwpn *string `json:"InitiatorWwpn,omitempty" validate:"regexp=^$|(^([0-9a-fA-F]{2}:){7}[0-9a-fA-F]{2}$)"`
	// The target in which OS installation triggered, the value represented is StorageControllerID follwed by PhysicalDisk SerialNumber in case of Physcial disk or VirtualDriveId for virtual drive.
	InstallTarget *string `json:"InstallTarget,omitempty"`
	// The Logical Unit Number (LUN) of the install target.
	LunId                     *int64                              `json:"LunId,omitempty"`
	OperatingSystemParameters NullableOsOperatingSystemParameters `json:"OperatingSystemParameters,omitempty"`
	ProcessedInstallTarget    NullableOsInstallTarget             `json:"ProcessedInstallTarget,omitempty"`
	// The Serial Number of the server.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// IQN (iSCSI qualified name) of Storage iSCSI target. Can be up to 255 characters long and has the following format, iqn.yyyy-mm.naming-authority:unique_name.
	TargetIqn *string `json:"TargetIqn,omitempty"`
	// The WWPN Address of the underlying fibre channel interface at the target used by the storage. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 51:4F:0C:50:14:1F:AF:01.
	TargetWwpn *string `json:"TargetWwpn,omitempty" validate:"regexp=^$|(^([0-9a-fA-F]{2}:){7}[0-9a-fA-F]{2}$)"`
	// MAC address of the VNIC used as iSCSI initiator interface.
	VnicMac              *string `json:"VnicMac,omitempty" validate:"regexp=^$|^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$"`
	AdditionalProperties map[string]interface{}
}

type _OsServerConfig OsServerConfig

// NewOsServerConfig instantiates a new OsServerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsServerConfig(classId string, objectType string) *OsServerConfig {
	this := OsServerConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsServerConfigWithDefaults instantiates a new OsServerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsServerConfigWithDefaults() *OsServerConfig {
	this := OsServerConfig{}
	var classId string = "os.ServerConfig"
	this.ClassId = classId
	var objectType string = "os.ServerConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsServerConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsServerConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsServerConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "os.ServerConfig" of the ClassId field.
func (o *OsServerConfig) GetDefaultClassId() interface{} {
	return "os.ServerConfig"
}

// GetObjectType returns the ObjectType field value
func (o *OsServerConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsServerConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsServerConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "os.ServerConfig" of the ObjectType field.
func (o *OsServerConfig) GetDefaultObjectType() interface{} {
	return "os.ServerConfig"
}

// GetAdditionalParameters returns the AdditionalParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsServerConfig) GetAdditionalParameters() []OsPlaceHolder {
	if o == nil {
		var ret []OsPlaceHolder
		return ret
	}
	return o.AdditionalParameters
}

// GetAdditionalParametersOk returns a tuple with the AdditionalParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsServerConfig) GetAdditionalParametersOk() ([]OsPlaceHolder, bool) {
	if o == nil || IsNil(o.AdditionalParameters) {
		return nil, false
	}
	return o.AdditionalParameters, true
}

// HasAdditionalParameters returns a boolean if a field has been set.
func (o *OsServerConfig) HasAdditionalParameters() bool {
	if o != nil && !IsNil(o.AdditionalParameters) {
		return true
	}

	return false
}

// SetAdditionalParameters gets a reference to the given []OsPlaceHolder and assigns it to the AdditionalParameters field.
func (o *OsServerConfig) SetAdditionalParameters(v []OsPlaceHolder) {
	o.AdditionalParameters = v
}

// GetAnswers returns the Answers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsServerConfig) GetAnswers() OsAnswers {
	if o == nil || IsNil(o.Answers.Get()) {
		var ret OsAnswers
		return ret
	}
	return *o.Answers.Get()
}

// GetAnswersOk returns a tuple with the Answers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsServerConfig) GetAnswersOk() (*OsAnswers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Answers.Get(), o.Answers.IsSet()
}

// HasAnswers returns a boolean if a field has been set.
func (o *OsServerConfig) HasAnswers() bool {
	if o != nil && o.Answers.IsSet() {
		return true
	}

	return false
}

// SetAnswers gets a reference to the given NullableOsAnswers and assigns it to the Answers field.
func (o *OsServerConfig) SetAnswers(v OsAnswers) {
	o.Answers.Set(&v)
}

// SetAnswersNil sets the value for Answers to be an explicit nil
func (o *OsServerConfig) SetAnswersNil() {
	o.Answers.Set(nil)
}

// UnsetAnswers ensures that no value is present for Answers, not even an explicit nil
func (o *OsServerConfig) UnsetAnswers() {
	o.Answers.Unset()
}

// GetErrorMsgs returns the ErrorMsgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsServerConfig) GetErrorMsgs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ErrorMsgs
}

// GetErrorMsgsOk returns a tuple with the ErrorMsgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsServerConfig) GetErrorMsgsOk() ([]string, bool) {
	if o == nil || IsNil(o.ErrorMsgs) {
		return nil, false
	}
	return o.ErrorMsgs, true
}

// HasErrorMsgs returns a boolean if a field has been set.
func (o *OsServerConfig) HasErrorMsgs() bool {
	if o != nil && !IsNil(o.ErrorMsgs) {
		return true
	}

	return false
}

// SetErrorMsgs gets a reference to the given []string and assigns it to the ErrorMsgs field.
func (o *OsServerConfig) SetErrorMsgs(v []string) {
	o.ErrorMsgs = v
}

// GetInitiatorWwpn returns the InitiatorWwpn field value if set, zero value otherwise.
func (o *OsServerConfig) GetInitiatorWwpn() string {
	if o == nil || IsNil(o.InitiatorWwpn) {
		var ret string
		return ret
	}
	return *o.InitiatorWwpn
}

// GetInitiatorWwpnOk returns a tuple with the InitiatorWwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsServerConfig) GetInitiatorWwpnOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatorWwpn) {
		return nil, false
	}
	return o.InitiatorWwpn, true
}

// HasInitiatorWwpn returns a boolean if a field has been set.
func (o *OsServerConfig) HasInitiatorWwpn() bool {
	if o != nil && !IsNil(o.InitiatorWwpn) {
		return true
	}

	return false
}

// SetInitiatorWwpn gets a reference to the given string and assigns it to the InitiatorWwpn field.
func (o *OsServerConfig) SetInitiatorWwpn(v string) {
	o.InitiatorWwpn = &v
}

// GetInstallTarget returns the InstallTarget field value if set, zero value otherwise.
func (o *OsServerConfig) GetInstallTarget() string {
	if o == nil || IsNil(o.InstallTarget) {
		var ret string
		return ret
	}
	return *o.InstallTarget
}

// GetInstallTargetOk returns a tuple with the InstallTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsServerConfig) GetInstallTargetOk() (*string, bool) {
	if o == nil || IsNil(o.InstallTarget) {
		return nil, false
	}
	return o.InstallTarget, true
}

// HasInstallTarget returns a boolean if a field has been set.
func (o *OsServerConfig) HasInstallTarget() bool {
	if o != nil && !IsNil(o.InstallTarget) {
		return true
	}

	return false
}

// SetInstallTarget gets a reference to the given string and assigns it to the InstallTarget field.
func (o *OsServerConfig) SetInstallTarget(v string) {
	o.InstallTarget = &v
}

// GetLunId returns the LunId field value if set, zero value otherwise.
func (o *OsServerConfig) GetLunId() int64 {
	if o == nil || IsNil(o.LunId) {
		var ret int64
		return ret
	}
	return *o.LunId
}

// GetLunIdOk returns a tuple with the LunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsServerConfig) GetLunIdOk() (*int64, bool) {
	if o == nil || IsNil(o.LunId) {
		return nil, false
	}
	return o.LunId, true
}

// HasLunId returns a boolean if a field has been set.
func (o *OsServerConfig) HasLunId() bool {
	if o != nil && !IsNil(o.LunId) {
		return true
	}

	return false
}

// SetLunId gets a reference to the given int64 and assigns it to the LunId field.
func (o *OsServerConfig) SetLunId(v int64) {
	o.LunId = &v
}

// GetOperatingSystemParameters returns the OperatingSystemParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsServerConfig) GetOperatingSystemParameters() OsOperatingSystemParameters {
	if o == nil || IsNil(o.OperatingSystemParameters.Get()) {
		var ret OsOperatingSystemParameters
		return ret
	}
	return *o.OperatingSystemParameters.Get()
}

// GetOperatingSystemParametersOk returns a tuple with the OperatingSystemParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsServerConfig) GetOperatingSystemParametersOk() (*OsOperatingSystemParameters, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperatingSystemParameters.Get(), o.OperatingSystemParameters.IsSet()
}

// HasOperatingSystemParameters returns a boolean if a field has been set.
func (o *OsServerConfig) HasOperatingSystemParameters() bool {
	if o != nil && o.OperatingSystemParameters.IsSet() {
		return true
	}

	return false
}

// SetOperatingSystemParameters gets a reference to the given NullableOsOperatingSystemParameters and assigns it to the OperatingSystemParameters field.
func (o *OsServerConfig) SetOperatingSystemParameters(v OsOperatingSystemParameters) {
	o.OperatingSystemParameters.Set(&v)
}

// SetOperatingSystemParametersNil sets the value for OperatingSystemParameters to be an explicit nil
func (o *OsServerConfig) SetOperatingSystemParametersNil() {
	o.OperatingSystemParameters.Set(nil)
}

// UnsetOperatingSystemParameters ensures that no value is present for OperatingSystemParameters, not even an explicit nil
func (o *OsServerConfig) UnsetOperatingSystemParameters() {
	o.OperatingSystemParameters.Unset()
}

// GetProcessedInstallTarget returns the ProcessedInstallTarget field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsServerConfig) GetProcessedInstallTarget() OsInstallTarget {
	if o == nil || IsNil(o.ProcessedInstallTarget.Get()) {
		var ret OsInstallTarget
		return ret
	}
	return *o.ProcessedInstallTarget.Get()
}

// GetProcessedInstallTargetOk returns a tuple with the ProcessedInstallTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsServerConfig) GetProcessedInstallTargetOk() (*OsInstallTarget, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessedInstallTarget.Get(), o.ProcessedInstallTarget.IsSet()
}

// HasProcessedInstallTarget returns a boolean if a field has been set.
func (o *OsServerConfig) HasProcessedInstallTarget() bool {
	if o != nil && o.ProcessedInstallTarget.IsSet() {
		return true
	}

	return false
}

// SetProcessedInstallTarget gets a reference to the given NullableOsInstallTarget and assigns it to the ProcessedInstallTarget field.
func (o *OsServerConfig) SetProcessedInstallTarget(v OsInstallTarget) {
	o.ProcessedInstallTarget.Set(&v)
}

// SetProcessedInstallTargetNil sets the value for ProcessedInstallTarget to be an explicit nil
func (o *OsServerConfig) SetProcessedInstallTargetNil() {
	o.ProcessedInstallTarget.Set(nil)
}

// UnsetProcessedInstallTarget ensures that no value is present for ProcessedInstallTarget, not even an explicit nil
func (o *OsServerConfig) UnsetProcessedInstallTarget() {
	o.ProcessedInstallTarget.Unset()
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *OsServerConfig) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsServerConfig) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *OsServerConfig) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *OsServerConfig) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetTargetIqn returns the TargetIqn field value if set, zero value otherwise.
func (o *OsServerConfig) GetTargetIqn() string {
	if o == nil || IsNil(o.TargetIqn) {
		var ret string
		return ret
	}
	return *o.TargetIqn
}

// GetTargetIqnOk returns a tuple with the TargetIqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsServerConfig) GetTargetIqnOk() (*string, bool) {
	if o == nil || IsNil(o.TargetIqn) {
		return nil, false
	}
	return o.TargetIqn, true
}

// HasTargetIqn returns a boolean if a field has been set.
func (o *OsServerConfig) HasTargetIqn() bool {
	if o != nil && !IsNil(o.TargetIqn) {
		return true
	}

	return false
}

// SetTargetIqn gets a reference to the given string and assigns it to the TargetIqn field.
func (o *OsServerConfig) SetTargetIqn(v string) {
	o.TargetIqn = &v
}

// GetTargetWwpn returns the TargetWwpn field value if set, zero value otherwise.
func (o *OsServerConfig) GetTargetWwpn() string {
	if o == nil || IsNil(o.TargetWwpn) {
		var ret string
		return ret
	}
	return *o.TargetWwpn
}

// GetTargetWwpnOk returns a tuple with the TargetWwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsServerConfig) GetTargetWwpnOk() (*string, bool) {
	if o == nil || IsNil(o.TargetWwpn) {
		return nil, false
	}
	return o.TargetWwpn, true
}

// HasTargetWwpn returns a boolean if a field has been set.
func (o *OsServerConfig) HasTargetWwpn() bool {
	if o != nil && !IsNil(o.TargetWwpn) {
		return true
	}

	return false
}

// SetTargetWwpn gets a reference to the given string and assigns it to the TargetWwpn field.
func (o *OsServerConfig) SetTargetWwpn(v string) {
	o.TargetWwpn = &v
}

// GetVnicMac returns the VnicMac field value if set, zero value otherwise.
func (o *OsServerConfig) GetVnicMac() string {
	if o == nil || IsNil(o.VnicMac) {
		var ret string
		return ret
	}
	return *o.VnicMac
}

// GetVnicMacOk returns a tuple with the VnicMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsServerConfig) GetVnicMacOk() (*string, bool) {
	if o == nil || IsNil(o.VnicMac) {
		return nil, false
	}
	return o.VnicMac, true
}

// HasVnicMac returns a boolean if a field has been set.
func (o *OsServerConfig) HasVnicMac() bool {
	if o != nil && !IsNil(o.VnicMac) {
		return true
	}

	return false
}

// SetVnicMac gets a reference to the given string and assigns it to the VnicMac field.
func (o *OsServerConfig) SetVnicMac(v string) {
	o.VnicMac = &v
}

func (o OsServerConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsServerConfig) ToMap() (map[string]interface{}, error) {
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
	if o.AdditionalParameters != nil {
		toSerialize["AdditionalParameters"] = o.AdditionalParameters
	}
	if o.Answers.IsSet() {
		toSerialize["Answers"] = o.Answers.Get()
	}
	if o.ErrorMsgs != nil {
		toSerialize["ErrorMsgs"] = o.ErrorMsgs
	}
	if !IsNil(o.InitiatorWwpn) {
		toSerialize["InitiatorWwpn"] = o.InitiatorWwpn
	}
	if !IsNil(o.InstallTarget) {
		toSerialize["InstallTarget"] = o.InstallTarget
	}
	if !IsNil(o.LunId) {
		toSerialize["LunId"] = o.LunId
	}
	if o.OperatingSystemParameters.IsSet() {
		toSerialize["OperatingSystemParameters"] = o.OperatingSystemParameters.Get()
	}
	if o.ProcessedInstallTarget.IsSet() {
		toSerialize["ProcessedInstallTarget"] = o.ProcessedInstallTarget.Get()
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if !IsNil(o.TargetIqn) {
		toSerialize["TargetIqn"] = o.TargetIqn
	}
	if !IsNil(o.TargetWwpn) {
		toSerialize["TargetWwpn"] = o.TargetWwpn
	}
	if !IsNil(o.VnicMac) {
		toSerialize["VnicMac"] = o.VnicMac
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OsServerConfig) UnmarshalJSON(data []byte) (err error) {
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
	type OsServerConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string            `json:"ObjectType"`
		AdditionalParameters []OsPlaceHolder   `json:"AdditionalParameters,omitempty"`
		Answers              NullableOsAnswers `json:"Answers,omitempty"`
		ErrorMsgs            []string          `json:"ErrorMsgs,omitempty"`
		// The WWPN Address of the underlying fibre channel interface at the host side used for SAN accesss. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 20:00:D4:C9:3C:35:02:01.
		InitiatorWwpn *string `json:"InitiatorWwpn,omitempty" validate:"regexp=^$|(^([0-9a-fA-F]{2}:){7}[0-9a-fA-F]{2}$)"`
		// The target in which OS installation triggered, the value represented is StorageControllerID follwed by PhysicalDisk SerialNumber in case of Physcial disk or VirtualDriveId for virtual drive.
		InstallTarget *string `json:"InstallTarget,omitempty"`
		// The Logical Unit Number (LUN) of the install target.
		LunId                     *int64                              `json:"LunId,omitempty"`
		OperatingSystemParameters NullableOsOperatingSystemParameters `json:"OperatingSystemParameters,omitempty"`
		ProcessedInstallTarget    NullableOsInstallTarget             `json:"ProcessedInstallTarget,omitempty"`
		// The Serial Number of the server.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// IQN (iSCSI qualified name) of Storage iSCSI target. Can be up to 255 characters long and has the following format, iqn.yyyy-mm.naming-authority:unique_name.
		TargetIqn *string `json:"TargetIqn,omitempty"`
		// The WWPN Address of the underlying fibre channel interface at the target used by the storage. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 51:4F:0C:50:14:1F:AF:01.
		TargetWwpn *string `json:"TargetWwpn,omitempty" validate:"regexp=^$|(^([0-9a-fA-F]{2}:){7}[0-9a-fA-F]{2}$)"`
		// MAC address of the VNIC used as iSCSI initiator interface.
		VnicMac *string `json:"VnicMac,omitempty" validate:"regexp=^$|^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$"`
	}

	varOsServerConfigWithoutEmbeddedStruct := OsServerConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOsServerConfigWithoutEmbeddedStruct)
	if err == nil {
		varOsServerConfig := _OsServerConfig{}
		varOsServerConfig.ClassId = varOsServerConfigWithoutEmbeddedStruct.ClassId
		varOsServerConfig.ObjectType = varOsServerConfigWithoutEmbeddedStruct.ObjectType
		varOsServerConfig.AdditionalParameters = varOsServerConfigWithoutEmbeddedStruct.AdditionalParameters
		varOsServerConfig.Answers = varOsServerConfigWithoutEmbeddedStruct.Answers
		varOsServerConfig.ErrorMsgs = varOsServerConfigWithoutEmbeddedStruct.ErrorMsgs
		varOsServerConfig.InitiatorWwpn = varOsServerConfigWithoutEmbeddedStruct.InitiatorWwpn
		varOsServerConfig.InstallTarget = varOsServerConfigWithoutEmbeddedStruct.InstallTarget
		varOsServerConfig.LunId = varOsServerConfigWithoutEmbeddedStruct.LunId
		varOsServerConfig.OperatingSystemParameters = varOsServerConfigWithoutEmbeddedStruct.OperatingSystemParameters
		varOsServerConfig.ProcessedInstallTarget = varOsServerConfigWithoutEmbeddedStruct.ProcessedInstallTarget
		varOsServerConfig.SerialNumber = varOsServerConfigWithoutEmbeddedStruct.SerialNumber
		varOsServerConfig.TargetIqn = varOsServerConfigWithoutEmbeddedStruct.TargetIqn
		varOsServerConfig.TargetWwpn = varOsServerConfigWithoutEmbeddedStruct.TargetWwpn
		varOsServerConfig.VnicMac = varOsServerConfigWithoutEmbeddedStruct.VnicMac
		*o = OsServerConfig(varOsServerConfig)
	} else {
		return err
	}

	varOsServerConfig := _OsServerConfig{}

	err = json.Unmarshal(data, &varOsServerConfig)
	if err == nil {
		o.MoBaseComplexType = varOsServerConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdditionalParameters")
		delete(additionalProperties, "Answers")
		delete(additionalProperties, "ErrorMsgs")
		delete(additionalProperties, "InitiatorWwpn")
		delete(additionalProperties, "InstallTarget")
		delete(additionalProperties, "LunId")
		delete(additionalProperties, "OperatingSystemParameters")
		delete(additionalProperties, "ProcessedInstallTarget")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "TargetIqn")
		delete(additionalProperties, "TargetWwpn")
		delete(additionalProperties, "VnicMac")

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

type NullableOsServerConfig struct {
	value *OsServerConfig
	isSet bool
}

func (v NullableOsServerConfig) Get() *OsServerConfig {
	return v.value
}

func (v *NullableOsServerConfig) Set(val *OsServerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOsServerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOsServerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsServerConfig(val *OsServerConfig) *NullableOsServerConfig {
	return &NullableOsServerConfig{value: val, isSet: true}
}

func (v NullableOsServerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsServerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
