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

// checks if the ApplianceCertRenewalPhase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceCertRenewalPhase{}

// ApplianceCertRenewalPhase CertRenewalPhase represents a phase of the Intersight Appliance certificate renewal process.
type ApplianceCertRenewalPhase struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// End date of the certificate renewal phase.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Indicates if the certificate renewal phase has failed or not.
	Failed *bool `json:"Failed,omitempty"`
	// Status message set during the certificate renewal phase.
	Message *string `json:"Message,omitempty"`
	// Name of the certificate renewal phase. * `Init` - New certificate detected, cleanup the old process if any running. * `ScheduleCertificateAddOperation` - Certificate Add Operation Schedulled. * `WaitForCertCollectionByEndpoint` - Monitor cert collection by endpoint. * `Success` - Certificate Renewal Task Success. * `Fail` - Certificate Renewal Task Fail. * `Cancel` - Certificate Renewal Task Cancel.
	Name *string `json:"Name,omitempty"`
	// Start date of the certificate renewal phase.
	StartTime            *time.Time `json:"StartTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceCertRenewalPhase ApplianceCertRenewalPhase

// NewApplianceCertRenewalPhase instantiates a new ApplianceCertRenewalPhase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceCertRenewalPhase(classId string, objectType string) *ApplianceCertRenewalPhase {
	this := ApplianceCertRenewalPhase{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceCertRenewalPhaseWithDefaults instantiates a new ApplianceCertRenewalPhase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceCertRenewalPhaseWithDefaults() *ApplianceCertRenewalPhase {
	this := ApplianceCertRenewalPhase{}
	var classId string = "appliance.CertRenewalPhase"
	this.ClassId = classId
	var objectType string = "appliance.CertRenewalPhase"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceCertRenewalPhase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceCertRenewalPhase) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.CertRenewalPhase" of the ClassId field.
func (o *ApplianceCertRenewalPhase) GetDefaultClassId() interface{} {
	return "appliance.CertRenewalPhase"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceCertRenewalPhase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceCertRenewalPhase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.CertRenewalPhase" of the ObjectType field.
func (o *ApplianceCertRenewalPhase) GetDefaultObjectType() interface{} {
	return "appliance.CertRenewalPhase"
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ApplianceCertRenewalPhase) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhase) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ApplianceCertRenewalPhase) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ApplianceCertRenewalPhase) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *ApplianceCertRenewalPhase) GetFailed() bool {
	if o == nil || IsNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhase) GetFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *ApplianceCertRenewalPhase) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *ApplianceCertRenewalPhase) SetFailed(v bool) {
	o.Failed = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApplianceCertRenewalPhase) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhase) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApplianceCertRenewalPhase) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApplianceCertRenewalPhase) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplianceCertRenewalPhase) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhase) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplianceCertRenewalPhase) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplianceCertRenewalPhase) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApplianceCertRenewalPhase) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhase) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApplianceCertRenewalPhase) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApplianceCertRenewalPhase) SetStartTime(v time.Time) {
	o.StartTime = &v
}

func (o ApplianceCertRenewalPhase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceCertRenewalPhase) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EndTime) {
		toSerialize["EndTime"] = o.EndTime
	}
	if !IsNil(o.Failed) {
		toSerialize["Failed"] = o.Failed
	}
	if !IsNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.StartTime) {
		toSerialize["StartTime"] = o.StartTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceCertRenewalPhase) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceCertRenewalPhaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// End date of the certificate renewal phase.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// Indicates if the certificate renewal phase has failed or not.
		Failed *bool `json:"Failed,omitempty"`
		// Status message set during the certificate renewal phase.
		Message *string `json:"Message,omitempty"`
		// Name of the certificate renewal phase. * `Init` - New certificate detected, cleanup the old process if any running. * `ScheduleCertificateAddOperation` - Certificate Add Operation Schedulled. * `WaitForCertCollectionByEndpoint` - Monitor cert collection by endpoint. * `Success` - Certificate Renewal Task Success. * `Fail` - Certificate Renewal Task Fail. * `Cancel` - Certificate Renewal Task Cancel.
		Name *string `json:"Name,omitempty"`
		// Start date of the certificate renewal phase.
		StartTime *time.Time `json:"StartTime,omitempty"`
	}

	varApplianceCertRenewalPhaseWithoutEmbeddedStruct := ApplianceCertRenewalPhaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceCertRenewalPhaseWithoutEmbeddedStruct)
	if err == nil {
		varApplianceCertRenewalPhase := _ApplianceCertRenewalPhase{}
		varApplianceCertRenewalPhase.ClassId = varApplianceCertRenewalPhaseWithoutEmbeddedStruct.ClassId
		varApplianceCertRenewalPhase.ObjectType = varApplianceCertRenewalPhaseWithoutEmbeddedStruct.ObjectType
		varApplianceCertRenewalPhase.EndTime = varApplianceCertRenewalPhaseWithoutEmbeddedStruct.EndTime
		varApplianceCertRenewalPhase.Failed = varApplianceCertRenewalPhaseWithoutEmbeddedStruct.Failed
		varApplianceCertRenewalPhase.Message = varApplianceCertRenewalPhaseWithoutEmbeddedStruct.Message
		varApplianceCertRenewalPhase.Name = varApplianceCertRenewalPhaseWithoutEmbeddedStruct.Name
		varApplianceCertRenewalPhase.StartTime = varApplianceCertRenewalPhaseWithoutEmbeddedStruct.StartTime
		*o = ApplianceCertRenewalPhase(varApplianceCertRenewalPhase)
	} else {
		return err
	}

	varApplianceCertRenewalPhase := _ApplianceCertRenewalPhase{}

	err = json.Unmarshal(data, &varApplianceCertRenewalPhase)
	if err == nil {
		o.MoBaseComplexType = varApplianceCertRenewalPhase.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "Failed")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "StartTime")

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

type NullableApplianceCertRenewalPhase struct {
	value *ApplianceCertRenewalPhase
	isSet bool
}

func (v NullableApplianceCertRenewalPhase) Get() *ApplianceCertRenewalPhase {
	return v.value
}

func (v *NullableApplianceCertRenewalPhase) Set(val *ApplianceCertRenewalPhase) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceCertRenewalPhase) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceCertRenewalPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceCertRenewalPhase(val *ApplianceCertRenewalPhase) *NullableApplianceCertRenewalPhase {
	return &NullableApplianceCertRenewalPhase{value: val, isSet: true}
}

func (v NullableApplianceCertRenewalPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceCertRenewalPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
