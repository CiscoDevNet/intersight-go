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

// checks if the ConnectorBaseMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorBaseMessage{}

// ConnectorBaseMessage A base abstract message that handles the common request properties applicable to all message types.
type ConnectorBaseMessage struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The secure properties that have large text content as value can be encrypted using AES key. In these cases, the AES key needs to be encrypted using the device connector's public key and passed as the value for this property. The secure properties that are encrypted using the AES key are mapped against the property name with prefix 'AES' in SecureProperties dictionary.
	EncryptedAesKey *string `json:"EncryptedAesKey,omitempty"`
	// The public key that was used to encrypt the values present in SecureProperties dictionary. If the given public key is not the same as device connector's public key, an error response with appropriate error message is thrown back.
	EncryptionKey *string `json:"EncryptionKey,omitempty"`
	// A dictionary of encrypted secure values mapped against the secure property name. The values that are encrypted using the AES key must be mapped against the secure property name with an 'AES' prefix Device connector expects the message body to be a golang template and the template can use the secure property names as placeholders.
	SecureProperties     interface{} `json:"SecureProperties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorBaseMessage ConnectorBaseMessage

// NewConnectorBaseMessage instantiates a new ConnectorBaseMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorBaseMessage(classId string, objectType string) *ConnectorBaseMessage {
	this := ConnectorBaseMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorBaseMessageWithDefaults instantiates a new ConnectorBaseMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorBaseMessageWithDefaults() *ConnectorBaseMessage {
	this := ConnectorBaseMessage{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorBaseMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorBaseMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorBaseMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorBaseMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorBaseMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorBaseMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEncryptedAesKey returns the EncryptedAesKey field value if set, zero value otherwise.
func (o *ConnectorBaseMessage) GetEncryptedAesKey() string {
	if o == nil || IsNil(o.EncryptedAesKey) {
		var ret string
		return ret
	}
	return *o.EncryptedAesKey
}

// GetEncryptedAesKeyOk returns a tuple with the EncryptedAesKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorBaseMessage) GetEncryptedAesKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedAesKey) {
		return nil, false
	}
	return o.EncryptedAesKey, true
}

// HasEncryptedAesKey returns a boolean if a field has been set.
func (o *ConnectorBaseMessage) HasEncryptedAesKey() bool {
	if o != nil && !IsNil(o.EncryptedAesKey) {
		return true
	}

	return false
}

// SetEncryptedAesKey gets a reference to the given string and assigns it to the EncryptedAesKey field.
func (o *ConnectorBaseMessage) SetEncryptedAesKey(v string) {
	o.EncryptedAesKey = &v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *ConnectorBaseMessage) GetEncryptionKey() string {
	if o == nil || IsNil(o.EncryptionKey) {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorBaseMessage) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionKey) {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *ConnectorBaseMessage) HasEncryptionKey() bool {
	if o != nil && !IsNil(o.EncryptionKey) {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *ConnectorBaseMessage) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetSecureProperties returns the SecureProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorBaseMessage) GetSecureProperties() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SecureProperties
}

// GetSecurePropertiesOk returns a tuple with the SecureProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorBaseMessage) GetSecurePropertiesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SecureProperties) {
		return nil, false
	}
	return &o.SecureProperties, true
}

// HasSecureProperties returns a boolean if a field has been set.
func (o *ConnectorBaseMessage) HasSecureProperties() bool {
	if o != nil && !IsNil(o.SecureProperties) {
		return true
	}

	return false
}

// SetSecureProperties gets a reference to the given interface{} and assigns it to the SecureProperties field.
func (o *ConnectorBaseMessage) SetSecureProperties(v interface{}) {
	o.SecureProperties = v
}

func (o ConnectorBaseMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorBaseMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.EncryptedAesKey) {
		toSerialize["EncryptedAesKey"] = o.EncryptedAesKey
	}
	if !IsNil(o.EncryptionKey) {
		toSerialize["EncryptionKey"] = o.EncryptionKey
	}
	if o.SecureProperties != nil {
		toSerialize["SecureProperties"] = o.SecureProperties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectorBaseMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type ConnectorBaseMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The secure properties that have large text content as value can be encrypted using AES key. In these cases, the AES key needs to be encrypted using the device connector's public key and passed as the value for this property. The secure properties that are encrypted using the AES key are mapped against the property name with prefix 'AES' in SecureProperties dictionary.
		EncryptedAesKey *string `json:"EncryptedAesKey,omitempty"`
		// The public key that was used to encrypt the values present in SecureProperties dictionary. If the given public key is not the same as device connector's public key, an error response with appropriate error message is thrown back.
		EncryptionKey *string `json:"EncryptionKey,omitempty"`
		// A dictionary of encrypted secure values mapped against the secure property name. The values that are encrypted using the AES key must be mapped against the secure property name with an 'AES' prefix Device connector expects the message body to be a golang template and the template can use the secure property names as placeholders.
		SecureProperties interface{} `json:"SecureProperties,omitempty"`
	}

	varConnectorBaseMessageWithoutEmbeddedStruct := ConnectorBaseMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varConnectorBaseMessageWithoutEmbeddedStruct)
	if err == nil {
		varConnectorBaseMessage := _ConnectorBaseMessage{}
		varConnectorBaseMessage.ClassId = varConnectorBaseMessageWithoutEmbeddedStruct.ClassId
		varConnectorBaseMessage.ObjectType = varConnectorBaseMessageWithoutEmbeddedStruct.ObjectType
		varConnectorBaseMessage.EncryptedAesKey = varConnectorBaseMessageWithoutEmbeddedStruct.EncryptedAesKey
		varConnectorBaseMessage.EncryptionKey = varConnectorBaseMessageWithoutEmbeddedStruct.EncryptionKey
		varConnectorBaseMessage.SecureProperties = varConnectorBaseMessageWithoutEmbeddedStruct.SecureProperties
		*o = ConnectorBaseMessage(varConnectorBaseMessage)
	} else {
		return err
	}

	varConnectorBaseMessage := _ConnectorBaseMessage{}

	err = json.Unmarshal(data, &varConnectorBaseMessage)
	if err == nil {
		o.MoBaseComplexType = varConnectorBaseMessage.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EncryptedAesKey")
		delete(additionalProperties, "EncryptionKey")
		delete(additionalProperties, "SecureProperties")

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

type NullableConnectorBaseMessage struct {
	value *ConnectorBaseMessage
	isSet bool
}

func (v NullableConnectorBaseMessage) Get() *ConnectorBaseMessage {
	return v.value
}

func (v *NullableConnectorBaseMessage) Set(val *ConnectorBaseMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorBaseMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorBaseMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorBaseMessage(val *ConnectorBaseMessage) *NullableConnectorBaseMessage {
	return &NullableConnectorBaseMessage{value: val, isSet: true}
}

func (v NullableConnectorBaseMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorBaseMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
