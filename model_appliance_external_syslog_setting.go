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

// checks if the ApplianceExternalSyslogSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceExternalSyslogSetting{}

// ApplianceExternalSyslogSetting Configure External Syslog Server.
type ApplianceExternalSyslogSetting struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable or disable External Syslog Server.
	Enabled *bool `json:"Enabled,omitempty"`
	// If the flag is set, the alarms reported in Intersight Appliances are exported to external syslog server based on the alarm severity selection.
	ExportAlarms *bool `json:"ExportAlarms,omitempty"`
	// Enable or disable exporting of Audit logs.
	ExportAudit *bool `json:"ExportAudit,omitempty"`
	// Enable or disable exporting of Web Server access logs.
	ExportNginx *bool `json:"ExportNginx,omitempty"`
	// External Syslog Server Port for connection establishment.
	Port *int64 `json:"Port,omitempty"`
	// Protocol used to connect to external syslog server. * `TCP` - External Syslog messages sent over TCP. * `UDP` - External Syslog messages sent over UDP. * `TLS` - Secure External Syslog messages sent over TLS.
	Protocol *string `json:"Protocol,omitempty"`
	// External Syslog Server Address, can be IP address or hostname.
	Server *string `json:"Server,omitempty"`
	// Minimum severity level to report. * `None` - The Enum value None represents that there is no severity. * `Info` - The Enum value Info represents the Informational level of severity. * `Critical` - The Enum value Critical represents the Critical level of severity. * `Warning` - The Enum value Warning represents the Warning level of severity. * `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.
	Severity             *string                        `json:"Severity,omitempty"`
	Account              NullableIamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceExternalSyslogSetting ApplianceExternalSyslogSetting

// NewApplianceExternalSyslogSetting instantiates a new ApplianceExternalSyslogSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceExternalSyslogSetting(classId string, objectType string) *ApplianceExternalSyslogSetting {
	this := ApplianceExternalSyslogSetting{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	var exportAlarms bool = false
	this.ExportAlarms = &exportAlarms
	var exportAudit bool = false
	this.ExportAudit = &exportAudit
	var exportNginx bool = false
	this.ExportNginx = &exportNginx
	var port int64 = 10514
	this.Port = &port
	var protocol string = "TCP"
	this.Protocol = &protocol
	var severity string = "None"
	this.Severity = &severity
	return &this
}

// NewApplianceExternalSyslogSettingWithDefaults instantiates a new ApplianceExternalSyslogSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceExternalSyslogSettingWithDefaults() *ApplianceExternalSyslogSetting {
	this := ApplianceExternalSyslogSetting{}
	var classId string = "appliance.ExternalSyslogSetting"
	this.ClassId = classId
	var objectType string = "appliance.ExternalSyslogSetting"
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	var exportAlarms bool = false
	this.ExportAlarms = &exportAlarms
	var exportAudit bool = false
	this.ExportAudit = &exportAudit
	var exportNginx bool = false
	this.ExportNginx = &exportNginx
	var port int64 = 10514
	this.Port = &port
	var protocol string = "TCP"
	this.Protocol = &protocol
	var severity string = "None"
	this.Severity = &severity
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceExternalSyslogSetting) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSetting) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceExternalSyslogSetting) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.ExternalSyslogSetting" of the ClassId field.
func (o *ApplianceExternalSyslogSetting) GetDefaultClassId() interface{} {
	return "appliance.ExternalSyslogSetting"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceExternalSyslogSetting) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSetting) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceExternalSyslogSetting) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.ExternalSyslogSetting" of the ObjectType field.
func (o *ApplianceExternalSyslogSetting) GetDefaultObjectType() interface{} {
	return "appliance.ExternalSyslogSetting"
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSetting) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSetting) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSetting) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplianceExternalSyslogSetting) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExportAlarms returns the ExportAlarms field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSetting) GetExportAlarms() bool {
	if o == nil || IsNil(o.ExportAlarms) {
		var ret bool
		return ret
	}
	return *o.ExportAlarms
}

// GetExportAlarmsOk returns a tuple with the ExportAlarms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSetting) GetExportAlarmsOk() (*bool, bool) {
	if o == nil || IsNil(o.ExportAlarms) {
		return nil, false
	}
	return o.ExportAlarms, true
}

// HasExportAlarms returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSetting) HasExportAlarms() bool {
	if o != nil && !IsNil(o.ExportAlarms) {
		return true
	}

	return false
}

// SetExportAlarms gets a reference to the given bool and assigns it to the ExportAlarms field.
func (o *ApplianceExternalSyslogSetting) SetExportAlarms(v bool) {
	o.ExportAlarms = &v
}

// GetExportAudit returns the ExportAudit field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSetting) GetExportAudit() bool {
	if o == nil || IsNil(o.ExportAudit) {
		var ret bool
		return ret
	}
	return *o.ExportAudit
}

// GetExportAuditOk returns a tuple with the ExportAudit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSetting) GetExportAuditOk() (*bool, bool) {
	if o == nil || IsNil(o.ExportAudit) {
		return nil, false
	}
	return o.ExportAudit, true
}

// HasExportAudit returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSetting) HasExportAudit() bool {
	if o != nil && !IsNil(o.ExportAudit) {
		return true
	}

	return false
}

// SetExportAudit gets a reference to the given bool and assigns it to the ExportAudit field.
func (o *ApplianceExternalSyslogSetting) SetExportAudit(v bool) {
	o.ExportAudit = &v
}

// GetExportNginx returns the ExportNginx field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSetting) GetExportNginx() bool {
	if o == nil || IsNil(o.ExportNginx) {
		var ret bool
		return ret
	}
	return *o.ExportNginx
}

// GetExportNginxOk returns a tuple with the ExportNginx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSetting) GetExportNginxOk() (*bool, bool) {
	if o == nil || IsNil(o.ExportNginx) {
		return nil, false
	}
	return o.ExportNginx, true
}

// HasExportNginx returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSetting) HasExportNginx() bool {
	if o != nil && !IsNil(o.ExportNginx) {
		return true
	}

	return false
}

// SetExportNginx gets a reference to the given bool and assigns it to the ExportNginx field.
func (o *ApplianceExternalSyslogSetting) SetExportNginx(v bool) {
	o.ExportNginx = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSetting) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSetting) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSetting) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *ApplianceExternalSyslogSetting) SetPort(v int64) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSetting) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSetting) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSetting) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ApplianceExternalSyslogSetting) SetProtocol(v string) {
	o.Protocol = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSetting) GetServer() string {
	if o == nil || IsNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSetting) GetServerOk() (*string, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSetting) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *ApplianceExternalSyslogSetting) SetServer(v string) {
	o.Server = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSetting) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSetting) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSetting) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ApplianceExternalSyslogSetting) SetSeverity(v string) {
	o.Severity = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceExternalSyslogSetting) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceExternalSyslogSetting) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSetting) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *ApplianceExternalSyslogSetting) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *ApplianceExternalSyslogSetting) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *ApplianceExternalSyslogSetting) UnsetAccount() {
	o.Account.Unset()
}

func (o ApplianceExternalSyslogSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceExternalSyslogSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.ExportAlarms) {
		toSerialize["ExportAlarms"] = o.ExportAlarms
	}
	if !IsNil(o.ExportAudit) {
		toSerialize["ExportAudit"] = o.ExportAudit
	}
	if !IsNil(o.ExportNginx) {
		toSerialize["ExportNginx"] = o.ExportNginx
	}
	if !IsNil(o.Port) {
		toSerialize["Port"] = o.Port
	}
	if !IsNil(o.Protocol) {
		toSerialize["Protocol"] = o.Protocol
	}
	if !IsNil(o.Server) {
		toSerialize["Server"] = o.Server
	}
	if !IsNil(o.Severity) {
		toSerialize["Severity"] = o.Severity
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceExternalSyslogSetting) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceExternalSyslogSettingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enable or disable External Syslog Server.
		Enabled *bool `json:"Enabled,omitempty"`
		// If the flag is set, the alarms reported in Intersight Appliances are exported to external syslog server based on the alarm severity selection.
		ExportAlarms *bool `json:"ExportAlarms,omitempty"`
		// Enable or disable exporting of Audit logs.
		ExportAudit *bool `json:"ExportAudit,omitempty"`
		// Enable or disable exporting of Web Server access logs.
		ExportNginx *bool `json:"ExportNginx,omitempty"`
		// External Syslog Server Port for connection establishment.
		Port *int64 `json:"Port,omitempty"`
		// Protocol used to connect to external syslog server. * `TCP` - External Syslog messages sent over TCP. * `UDP` - External Syslog messages sent over UDP. * `TLS` - Secure External Syslog messages sent over TLS.
		Protocol *string `json:"Protocol,omitempty"`
		// External Syslog Server Address, can be IP address or hostname.
		Server *string `json:"Server,omitempty"`
		// Minimum severity level to report. * `None` - The Enum value None represents that there is no severity. * `Info` - The Enum value Info represents the Informational level of severity. * `Critical` - The Enum value Critical represents the Critical level of severity. * `Warning` - The Enum value Warning represents the Warning level of severity. * `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.
		Severity *string                        `json:"Severity,omitempty"`
		Account  NullableIamAccountRelationship `json:"Account,omitempty"`
	}

	varApplianceExternalSyslogSettingWithoutEmbeddedStruct := ApplianceExternalSyslogSettingWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceExternalSyslogSettingWithoutEmbeddedStruct)
	if err == nil {
		varApplianceExternalSyslogSetting := _ApplianceExternalSyslogSetting{}
		varApplianceExternalSyslogSetting.ClassId = varApplianceExternalSyslogSettingWithoutEmbeddedStruct.ClassId
		varApplianceExternalSyslogSetting.ObjectType = varApplianceExternalSyslogSettingWithoutEmbeddedStruct.ObjectType
		varApplianceExternalSyslogSetting.Enabled = varApplianceExternalSyslogSettingWithoutEmbeddedStruct.Enabled
		varApplianceExternalSyslogSetting.ExportAlarms = varApplianceExternalSyslogSettingWithoutEmbeddedStruct.ExportAlarms
		varApplianceExternalSyslogSetting.ExportAudit = varApplianceExternalSyslogSettingWithoutEmbeddedStruct.ExportAudit
		varApplianceExternalSyslogSetting.ExportNginx = varApplianceExternalSyslogSettingWithoutEmbeddedStruct.ExportNginx
		varApplianceExternalSyslogSetting.Port = varApplianceExternalSyslogSettingWithoutEmbeddedStruct.Port
		varApplianceExternalSyslogSetting.Protocol = varApplianceExternalSyslogSettingWithoutEmbeddedStruct.Protocol
		varApplianceExternalSyslogSetting.Server = varApplianceExternalSyslogSettingWithoutEmbeddedStruct.Server
		varApplianceExternalSyslogSetting.Severity = varApplianceExternalSyslogSettingWithoutEmbeddedStruct.Severity
		varApplianceExternalSyslogSetting.Account = varApplianceExternalSyslogSettingWithoutEmbeddedStruct.Account
		*o = ApplianceExternalSyslogSetting(varApplianceExternalSyslogSetting)
	} else {
		return err
	}

	varApplianceExternalSyslogSetting := _ApplianceExternalSyslogSetting{}

	err = json.Unmarshal(data, &varApplianceExternalSyslogSetting)
	if err == nil {
		o.MoBaseMo = varApplianceExternalSyslogSetting.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "ExportAlarms")
		delete(additionalProperties, "ExportAudit")
		delete(additionalProperties, "ExportNginx")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "Severity")
		delete(additionalProperties, "Account")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableApplianceExternalSyslogSetting struct {
	value *ApplianceExternalSyslogSetting
	isSet bool
}

func (v NullableApplianceExternalSyslogSetting) Get() *ApplianceExternalSyslogSetting {
	return v.value
}

func (v *NullableApplianceExternalSyslogSetting) Set(val *ApplianceExternalSyslogSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceExternalSyslogSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceExternalSyslogSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceExternalSyslogSetting(val *ApplianceExternalSyslogSetting) *NullableApplianceExternalSyslogSetting {
	return &NullableApplianceExternalSyslogSetting{value: val, isSet: true}
}

func (v NullableApplianceExternalSyslogSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceExternalSyslogSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
