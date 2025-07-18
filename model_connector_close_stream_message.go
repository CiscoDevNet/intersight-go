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

// checks if the ConnectorCloseStreamMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorCloseStreamMessage{}

// ConnectorCloseStreamMessage Close a running stream. After close the stream will no longer emit events or be available to write to.
type ConnectorCloseStreamMessage struct {
	ConnectorStreamMessage
	AdditionalProperties map[string]interface{}
}

type _ConnectorCloseStreamMessage ConnectorCloseStreamMessage

// NewConnectorCloseStreamMessage instantiates a new ConnectorCloseStreamMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorCloseStreamMessage(classId string, objectType string) *ConnectorCloseStreamMessage {
	this := ConnectorCloseStreamMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorCloseStreamMessageWithDefaults instantiates a new ConnectorCloseStreamMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorCloseStreamMessageWithDefaults() *ConnectorCloseStreamMessage {
	this := ConnectorCloseStreamMessage{}
	return &this
}

func (o ConnectorCloseStreamMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorCloseStreamMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorStreamMessage, errConnectorStreamMessage := json.Marshal(o.ConnectorStreamMessage)
	if errConnectorStreamMessage != nil {
		return map[string]interface{}{}, errConnectorStreamMessage
	}
	errConnectorStreamMessage = json.Unmarshal([]byte(serializedConnectorStreamMessage), &toSerialize)
	if errConnectorStreamMessage != nil {
		return map[string]interface{}{}, errConnectorStreamMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectorCloseStreamMessage) UnmarshalJSON(data []byte) (err error) {
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
	type ConnectorCloseStreamMessageWithoutEmbeddedStruct struct {
	}

	varConnectorCloseStreamMessageWithoutEmbeddedStruct := ConnectorCloseStreamMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varConnectorCloseStreamMessageWithoutEmbeddedStruct)
	if err == nil {
		varConnectorCloseStreamMessage := _ConnectorCloseStreamMessage{}
		*o = ConnectorCloseStreamMessage(varConnectorCloseStreamMessage)
	} else {
		return err
	}

	varConnectorCloseStreamMessage := _ConnectorCloseStreamMessage{}

	err = json.Unmarshal(data, &varConnectorCloseStreamMessage)
	if err == nil {
		o.ConnectorStreamMessage = varConnectorCloseStreamMessage.ConnectorStreamMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectConnectorStreamMessage := reflect.ValueOf(o.ConnectorStreamMessage)
		for i := 0; i < reflectConnectorStreamMessage.Type().NumField(); i++ {
			t := reflectConnectorStreamMessage.Type().Field(i)

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

type NullableConnectorCloseStreamMessage struct {
	value *ConnectorCloseStreamMessage
	isSet bool
}

func (v NullableConnectorCloseStreamMessage) Get() *ConnectorCloseStreamMessage {
	return v.value
}

func (v *NullableConnectorCloseStreamMessage) Set(val *ConnectorCloseStreamMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorCloseStreamMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorCloseStreamMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorCloseStreamMessage(val *ConnectorCloseStreamMessage) *NullableConnectorCloseStreamMessage {
	return &NullableConnectorCloseStreamMessage{value: val, isSet: true}
}

func (v NullableConnectorCloseStreamMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorCloseStreamMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
