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
	"time"
)

// checks if the ApplianceDeviceCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceDeviceCertificate{}

// ApplianceDeviceCertificate DeviceCertificate managed object stores the CA Certificate used by device connector and it allow tracks it renewal.
type ApplianceDeviceCertificate struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The base64 encoded certificate in PEM format.
	CaCertificate *string `json:"CaCertificate,omitempty"`
	// The expiry datetime of new ca certificate which need to be applied on device connector.
	CaCertificateExpiryTime *time.Time `json:"CaCertificateExpiryTime,omitempty"`
	// The date time allocated till cert renewal will be executed. This time used here will be based on cert renewal plan.
	CertificateRenewalExpiryTime *time.Time                  `json:"CertificateRenewalExpiryTime,omitempty"`
	CompletedPhases              []ApplianceCertRenewalPhase `json:"CompletedPhases,omitempty"`
	// The operation configuration MOId.
	ConfigurationMoId *string                           `json:"ConfigurationMoId,omitempty"`
	CurrentPhase      NullableApplianceCertRenewalPhase `json:"CurrentPhase,omitempty"`
	// End date of the certificate renewal.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// The last poll time when data collection was successfull. This time is used to collect data after this time in next cycle.
	LastSuccessPollTime *time.Time `json:"LastSuccessPollTime,omitempty"`
	Messages            []string   `json:"Messages,omitempty"`
	// Start date of the certificate renewal.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// The status of ca certificate renewal.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceDeviceCertificate ApplianceDeviceCertificate

// NewApplianceDeviceCertificate instantiates a new ApplianceDeviceCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceDeviceCertificate(classId string, objectType string) *ApplianceDeviceCertificate {
	this := ApplianceDeviceCertificate{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceDeviceCertificateWithDefaults instantiates a new ApplianceDeviceCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceDeviceCertificateWithDefaults() *ApplianceDeviceCertificate {
	this := ApplianceDeviceCertificate{}
	var classId string = "appliance.DeviceCertificate"
	this.ClassId = classId
	var objectType string = "appliance.DeviceCertificate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceDeviceCertificate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceDeviceCertificate) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.DeviceCertificate" of the ClassId field.
func (o *ApplianceDeviceCertificate) GetDefaultClassId() interface{} {
	return "appliance.DeviceCertificate"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceDeviceCertificate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceDeviceCertificate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.DeviceCertificate" of the ObjectType field.
func (o *ApplianceDeviceCertificate) GetDefaultObjectType() interface{} {
	return "appliance.DeviceCertificate"
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetCaCertificate() string {
	if o == nil || IsNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetCaCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.CaCertificate) {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasCaCertificate() bool {
	if o != nil && !IsNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *ApplianceDeviceCertificate) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

// GetCaCertificateExpiryTime returns the CaCertificateExpiryTime field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetCaCertificateExpiryTime() time.Time {
	if o == nil || IsNil(o.CaCertificateExpiryTime) {
		var ret time.Time
		return ret
	}
	return *o.CaCertificateExpiryTime
}

// GetCaCertificateExpiryTimeOk returns a tuple with the CaCertificateExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetCaCertificateExpiryTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CaCertificateExpiryTime) {
		return nil, false
	}
	return o.CaCertificateExpiryTime, true
}

// HasCaCertificateExpiryTime returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasCaCertificateExpiryTime() bool {
	if o != nil && !IsNil(o.CaCertificateExpiryTime) {
		return true
	}

	return false
}

// SetCaCertificateExpiryTime gets a reference to the given time.Time and assigns it to the CaCertificateExpiryTime field.
func (o *ApplianceDeviceCertificate) SetCaCertificateExpiryTime(v time.Time) {
	o.CaCertificateExpiryTime = &v
}

// GetCertificateRenewalExpiryTime returns the CertificateRenewalExpiryTime field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetCertificateRenewalExpiryTime() time.Time {
	if o == nil || IsNil(o.CertificateRenewalExpiryTime) {
		var ret time.Time
		return ret
	}
	return *o.CertificateRenewalExpiryTime
}

// GetCertificateRenewalExpiryTimeOk returns a tuple with the CertificateRenewalExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetCertificateRenewalExpiryTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CertificateRenewalExpiryTime) {
		return nil, false
	}
	return o.CertificateRenewalExpiryTime, true
}

// HasCertificateRenewalExpiryTime returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasCertificateRenewalExpiryTime() bool {
	if o != nil && !IsNil(o.CertificateRenewalExpiryTime) {
		return true
	}

	return false
}

// SetCertificateRenewalExpiryTime gets a reference to the given time.Time and assigns it to the CertificateRenewalExpiryTime field.
func (o *ApplianceDeviceCertificate) SetCertificateRenewalExpiryTime(v time.Time) {
	o.CertificateRenewalExpiryTime = &v
}

// GetCompletedPhases returns the CompletedPhases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceDeviceCertificate) GetCompletedPhases() []ApplianceCertRenewalPhase {
	if o == nil {
		var ret []ApplianceCertRenewalPhase
		return ret
	}
	return o.CompletedPhases
}

// GetCompletedPhasesOk returns a tuple with the CompletedPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceDeviceCertificate) GetCompletedPhasesOk() ([]ApplianceCertRenewalPhase, bool) {
	if o == nil || IsNil(o.CompletedPhases) {
		return nil, false
	}
	return o.CompletedPhases, true
}

// HasCompletedPhases returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasCompletedPhases() bool {
	if o != nil && !IsNil(o.CompletedPhases) {
		return true
	}

	return false
}

// SetCompletedPhases gets a reference to the given []ApplianceCertRenewalPhase and assigns it to the CompletedPhases field.
func (o *ApplianceDeviceCertificate) SetCompletedPhases(v []ApplianceCertRenewalPhase) {
	o.CompletedPhases = v
}

// GetConfigurationMoId returns the ConfigurationMoId field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetConfigurationMoId() string {
	if o == nil || IsNil(o.ConfigurationMoId) {
		var ret string
		return ret
	}
	return *o.ConfigurationMoId
}

// GetConfigurationMoIdOk returns a tuple with the ConfigurationMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetConfigurationMoIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationMoId) {
		return nil, false
	}
	return o.ConfigurationMoId, true
}

// HasConfigurationMoId returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasConfigurationMoId() bool {
	if o != nil && !IsNil(o.ConfigurationMoId) {
		return true
	}

	return false
}

// SetConfigurationMoId gets a reference to the given string and assigns it to the ConfigurationMoId field.
func (o *ApplianceDeviceCertificate) SetConfigurationMoId(v string) {
	o.ConfigurationMoId = &v
}

// GetCurrentPhase returns the CurrentPhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceDeviceCertificate) GetCurrentPhase() ApplianceCertRenewalPhase {
	if o == nil || IsNil(o.CurrentPhase.Get()) {
		var ret ApplianceCertRenewalPhase
		return ret
	}
	return *o.CurrentPhase.Get()
}

// GetCurrentPhaseOk returns a tuple with the CurrentPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceDeviceCertificate) GetCurrentPhaseOk() (*ApplianceCertRenewalPhase, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPhase.Get(), o.CurrentPhase.IsSet()
}

// HasCurrentPhase returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasCurrentPhase() bool {
	if o != nil && o.CurrentPhase.IsSet() {
		return true
	}

	return false
}

// SetCurrentPhase gets a reference to the given NullableApplianceCertRenewalPhase and assigns it to the CurrentPhase field.
func (o *ApplianceDeviceCertificate) SetCurrentPhase(v ApplianceCertRenewalPhase) {
	o.CurrentPhase.Set(&v)
}

// SetCurrentPhaseNil sets the value for CurrentPhase to be an explicit nil
func (o *ApplianceDeviceCertificate) SetCurrentPhaseNil() {
	o.CurrentPhase.Set(nil)
}

// UnsetCurrentPhase ensures that no value is present for CurrentPhase, not even an explicit nil
func (o *ApplianceDeviceCertificate) UnsetCurrentPhase() {
	o.CurrentPhase.Unset()
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ApplianceDeviceCertificate) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetLastSuccessPollTime returns the LastSuccessPollTime field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetLastSuccessPollTime() time.Time {
	if o == nil || IsNil(o.LastSuccessPollTime) {
		var ret time.Time
		return ret
	}
	return *o.LastSuccessPollTime
}

// GetLastSuccessPollTimeOk returns a tuple with the LastSuccessPollTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetLastSuccessPollTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSuccessPollTime) {
		return nil, false
	}
	return o.LastSuccessPollTime, true
}

// HasLastSuccessPollTime returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasLastSuccessPollTime() bool {
	if o != nil && !IsNil(o.LastSuccessPollTime) {
		return true
	}

	return false
}

// SetLastSuccessPollTime gets a reference to the given time.Time and assigns it to the LastSuccessPollTime field.
func (o *ApplianceDeviceCertificate) SetLastSuccessPollTime(v time.Time) {
	o.LastSuccessPollTime = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceDeviceCertificate) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceDeviceCertificate) GetMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *ApplianceDeviceCertificate) SetMessages(v []string) {
	o.Messages = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApplianceDeviceCertificate) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceDeviceCertificate) SetStatus(v string) {
	o.Status = &v
}

func (o ApplianceDeviceCertificate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceDeviceCertificate) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CaCertificate) {
		toSerialize["CaCertificate"] = o.CaCertificate
	}
	if !IsNil(o.CaCertificateExpiryTime) {
		toSerialize["CaCertificateExpiryTime"] = o.CaCertificateExpiryTime
	}
	if !IsNil(o.CertificateRenewalExpiryTime) {
		toSerialize["CertificateRenewalExpiryTime"] = o.CertificateRenewalExpiryTime
	}
	if o.CompletedPhases != nil {
		toSerialize["CompletedPhases"] = o.CompletedPhases
	}
	if !IsNil(o.ConfigurationMoId) {
		toSerialize["ConfigurationMoId"] = o.ConfigurationMoId
	}
	if o.CurrentPhase.IsSet() {
		toSerialize["CurrentPhase"] = o.CurrentPhase.Get()
	}
	if !IsNil(o.EndTime) {
		toSerialize["EndTime"] = o.EndTime
	}
	if !IsNil(o.LastSuccessPollTime) {
		toSerialize["LastSuccessPollTime"] = o.LastSuccessPollTime
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if !IsNil(o.StartTime) {
		toSerialize["StartTime"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceDeviceCertificate) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceDeviceCertificateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The base64 encoded certificate in PEM format.
		CaCertificate *string `json:"CaCertificate,omitempty"`
		// The expiry datetime of new ca certificate which need to be applied on device connector.
		CaCertificateExpiryTime *time.Time `json:"CaCertificateExpiryTime,omitempty"`
		// The date time allocated till cert renewal will be executed. This time used here will be based on cert renewal plan.
		CertificateRenewalExpiryTime *time.Time                  `json:"CertificateRenewalExpiryTime,omitempty"`
		CompletedPhases              []ApplianceCertRenewalPhase `json:"CompletedPhases,omitempty"`
		// The operation configuration MOId.
		ConfigurationMoId *string                           `json:"ConfigurationMoId,omitempty"`
		CurrentPhase      NullableApplianceCertRenewalPhase `json:"CurrentPhase,omitempty"`
		// End date of the certificate renewal.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// The last poll time when data collection was successfull. This time is used to collect data after this time in next cycle.
		LastSuccessPollTime *time.Time `json:"LastSuccessPollTime,omitempty"`
		Messages            []string   `json:"Messages,omitempty"`
		// Start date of the certificate renewal.
		StartTime *time.Time `json:"StartTime,omitempty"`
		// The status of ca certificate renewal.
		Status *string `json:"Status,omitempty"`
	}

	varApplianceDeviceCertificateWithoutEmbeddedStruct := ApplianceDeviceCertificateWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceDeviceCertificateWithoutEmbeddedStruct)
	if err == nil {
		varApplianceDeviceCertificate := _ApplianceDeviceCertificate{}
		varApplianceDeviceCertificate.ClassId = varApplianceDeviceCertificateWithoutEmbeddedStruct.ClassId
		varApplianceDeviceCertificate.ObjectType = varApplianceDeviceCertificateWithoutEmbeddedStruct.ObjectType
		varApplianceDeviceCertificate.CaCertificate = varApplianceDeviceCertificateWithoutEmbeddedStruct.CaCertificate
		varApplianceDeviceCertificate.CaCertificateExpiryTime = varApplianceDeviceCertificateWithoutEmbeddedStruct.CaCertificateExpiryTime
		varApplianceDeviceCertificate.CertificateRenewalExpiryTime = varApplianceDeviceCertificateWithoutEmbeddedStruct.CertificateRenewalExpiryTime
		varApplianceDeviceCertificate.CompletedPhases = varApplianceDeviceCertificateWithoutEmbeddedStruct.CompletedPhases
		varApplianceDeviceCertificate.ConfigurationMoId = varApplianceDeviceCertificateWithoutEmbeddedStruct.ConfigurationMoId
		varApplianceDeviceCertificate.CurrentPhase = varApplianceDeviceCertificateWithoutEmbeddedStruct.CurrentPhase
		varApplianceDeviceCertificate.EndTime = varApplianceDeviceCertificateWithoutEmbeddedStruct.EndTime
		varApplianceDeviceCertificate.LastSuccessPollTime = varApplianceDeviceCertificateWithoutEmbeddedStruct.LastSuccessPollTime
		varApplianceDeviceCertificate.Messages = varApplianceDeviceCertificateWithoutEmbeddedStruct.Messages
		varApplianceDeviceCertificate.StartTime = varApplianceDeviceCertificateWithoutEmbeddedStruct.StartTime
		varApplianceDeviceCertificate.Status = varApplianceDeviceCertificateWithoutEmbeddedStruct.Status
		*o = ApplianceDeviceCertificate(varApplianceDeviceCertificate)
	} else {
		return err
	}

	varApplianceDeviceCertificate := _ApplianceDeviceCertificate{}

	err = json.Unmarshal(data, &varApplianceDeviceCertificate)
	if err == nil {
		o.MoBaseMo = varApplianceDeviceCertificate.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CaCertificate")
		delete(additionalProperties, "CaCertificateExpiryTime")
		delete(additionalProperties, "CertificateRenewalExpiryTime")
		delete(additionalProperties, "CompletedPhases")
		delete(additionalProperties, "ConfigurationMoId")
		delete(additionalProperties, "CurrentPhase")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "LastSuccessPollTime")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")

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

type NullableApplianceDeviceCertificate struct {
	value *ApplianceDeviceCertificate
	isSet bool
}

func (v NullableApplianceDeviceCertificate) Get() *ApplianceDeviceCertificate {
	return v.value
}

func (v *NullableApplianceDeviceCertificate) Set(val *ApplianceDeviceCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceDeviceCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceDeviceCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceDeviceCertificate(val *ApplianceDeviceCertificate) *NullableApplianceDeviceCertificate {
	return &NullableApplianceDeviceCertificate{value: val, isSet: true}
}

func (v NullableApplianceDeviceCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceDeviceCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
