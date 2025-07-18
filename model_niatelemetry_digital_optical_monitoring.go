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

// checks if the NiatelemetryDigitalOpticalMonitoring type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryDigitalOpticalMonitoring{}

// NiatelemetryDigitalOpticalMonitoring Digital optical monitoring details for aci nodes.
type NiatelemetryDigitalOpticalMonitoring struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Alerts count for the interface in the node.
	Alerts *string `json:"Alerts,omitempty"`
	// Dn with interface name for the aci nodes.
	Dn *string `json:"Dn,omitempty"`
	// RxLos count for the interface in the node.
	RxLos *string `json:"RxLos,omitempty"`
	// TxfaultCount for the interface in the node.
	TxFaultCount         *string `json:"TxFaultCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryDigitalOpticalMonitoring NiatelemetryDigitalOpticalMonitoring

// NewNiatelemetryDigitalOpticalMonitoring instantiates a new NiatelemetryDigitalOpticalMonitoring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryDigitalOpticalMonitoring(classId string, objectType string) *NiatelemetryDigitalOpticalMonitoring {
	this := NiatelemetryDigitalOpticalMonitoring{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryDigitalOpticalMonitoringWithDefaults instantiates a new NiatelemetryDigitalOpticalMonitoring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryDigitalOpticalMonitoringWithDefaults() *NiatelemetryDigitalOpticalMonitoring {
	this := NiatelemetryDigitalOpticalMonitoring{}
	var classId string = "niatelemetry.DigitalOpticalMonitoring"
	this.ClassId = classId
	var objectType string = "niatelemetry.DigitalOpticalMonitoring"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryDigitalOpticalMonitoring) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryDigitalOpticalMonitoring) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.DigitalOpticalMonitoring" of the ClassId field.
func (o *NiatelemetryDigitalOpticalMonitoring) GetDefaultClassId() interface{} {
	return "niatelemetry.DigitalOpticalMonitoring"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryDigitalOpticalMonitoring) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryDigitalOpticalMonitoring) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.DigitalOpticalMonitoring" of the ObjectType field.
func (o *NiatelemetryDigitalOpticalMonitoring) GetDefaultObjectType() interface{} {
	return "niatelemetry.DigitalOpticalMonitoring"
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *NiatelemetryDigitalOpticalMonitoring) GetAlerts() string {
	if o == nil || IsNil(o.Alerts) {
		var ret string
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) GetAlertsOk() (*string, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given string and assigns it to the Alerts field.
func (o *NiatelemetryDigitalOpticalMonitoring) SetAlerts(v string) {
	o.Alerts = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryDigitalOpticalMonitoring) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryDigitalOpticalMonitoring) SetDn(v string) {
	o.Dn = &v
}

// GetRxLos returns the RxLos field value if set, zero value otherwise.
func (o *NiatelemetryDigitalOpticalMonitoring) GetRxLos() string {
	if o == nil || IsNil(o.RxLos) {
		var ret string
		return ret
	}
	return *o.RxLos
}

// GetRxLosOk returns a tuple with the RxLos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) GetRxLosOk() (*string, bool) {
	if o == nil || IsNil(o.RxLos) {
		return nil, false
	}
	return o.RxLos, true
}

// HasRxLos returns a boolean if a field has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) HasRxLos() bool {
	if o != nil && !IsNil(o.RxLos) {
		return true
	}

	return false
}

// SetRxLos gets a reference to the given string and assigns it to the RxLos field.
func (o *NiatelemetryDigitalOpticalMonitoring) SetRxLos(v string) {
	o.RxLos = &v
}

// GetTxFaultCount returns the TxFaultCount field value if set, zero value otherwise.
func (o *NiatelemetryDigitalOpticalMonitoring) GetTxFaultCount() string {
	if o == nil || IsNil(o.TxFaultCount) {
		var ret string
		return ret
	}
	return *o.TxFaultCount
}

// GetTxFaultCountOk returns a tuple with the TxFaultCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) GetTxFaultCountOk() (*string, bool) {
	if o == nil || IsNil(o.TxFaultCount) {
		return nil, false
	}
	return o.TxFaultCount, true
}

// HasTxFaultCount returns a boolean if a field has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) HasTxFaultCount() bool {
	if o != nil && !IsNil(o.TxFaultCount) {
		return true
	}

	return false
}

// SetTxFaultCount gets a reference to the given string and assigns it to the TxFaultCount field.
func (o *NiatelemetryDigitalOpticalMonitoring) SetTxFaultCount(v string) {
	o.TxFaultCount = &v
}

func (o NiatelemetryDigitalOpticalMonitoring) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryDigitalOpticalMonitoring) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Alerts) {
		toSerialize["Alerts"] = o.Alerts
	}
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if !IsNil(o.RxLos) {
		toSerialize["RxLos"] = o.RxLos
	}
	if !IsNil(o.TxFaultCount) {
		toSerialize["TxFaultCount"] = o.TxFaultCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryDigitalOpticalMonitoring) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Alerts count for the interface in the node.
		Alerts *string `json:"Alerts,omitempty"`
		// Dn with interface name for the aci nodes.
		Dn *string `json:"Dn,omitempty"`
		// RxLos count for the interface in the node.
		RxLos *string `json:"RxLos,omitempty"`
		// TxfaultCount for the interface in the node.
		TxFaultCount *string `json:"TxFaultCount,omitempty"`
	}

	varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct := NiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryDigitalOpticalMonitoring := _NiatelemetryDigitalOpticalMonitoring{}
		varNiatelemetryDigitalOpticalMonitoring.ClassId = varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct.ClassId
		varNiatelemetryDigitalOpticalMonitoring.ObjectType = varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct.ObjectType
		varNiatelemetryDigitalOpticalMonitoring.Alerts = varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct.Alerts
		varNiatelemetryDigitalOpticalMonitoring.Dn = varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct.Dn
		varNiatelemetryDigitalOpticalMonitoring.RxLos = varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct.RxLos
		varNiatelemetryDigitalOpticalMonitoring.TxFaultCount = varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct.TxFaultCount
		*o = NiatelemetryDigitalOpticalMonitoring(varNiatelemetryDigitalOpticalMonitoring)
	} else {
		return err
	}

	varNiatelemetryDigitalOpticalMonitoring := _NiatelemetryDigitalOpticalMonitoring{}

	err = json.Unmarshal(data, &varNiatelemetryDigitalOpticalMonitoring)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryDigitalOpticalMonitoring.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Alerts")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "RxLos")
		delete(additionalProperties, "TxFaultCount")

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

type NullableNiatelemetryDigitalOpticalMonitoring struct {
	value *NiatelemetryDigitalOpticalMonitoring
	isSet bool
}

func (v NullableNiatelemetryDigitalOpticalMonitoring) Get() *NiatelemetryDigitalOpticalMonitoring {
	return v.value
}

func (v *NullableNiatelemetryDigitalOpticalMonitoring) Set(val *NiatelemetryDigitalOpticalMonitoring) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryDigitalOpticalMonitoring) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryDigitalOpticalMonitoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryDigitalOpticalMonitoring(val *NiatelemetryDigitalOpticalMonitoring) *NullableNiatelemetryDigitalOpticalMonitoring {
	return &NullableNiatelemetryDigitalOpticalMonitoring{value: val, isSet: true}
}

func (v NullableNiatelemetryDigitalOpticalMonitoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryDigitalOpticalMonitoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
