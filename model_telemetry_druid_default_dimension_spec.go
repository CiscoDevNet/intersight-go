/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidDefaultDimensionSpec Returns dimension values as is and optionally renames the dimension.
type TelemetryDruidDefaultDimensionSpec struct {
	// the dimension spec type.
	Type                 string `json:"type"`
	Dimension            string `json:"dimension"`
	OutputName           string `json:"outputName"`
	OutputType           string `json:"outputType"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidDefaultDimensionSpec TelemetryDruidDefaultDimensionSpec

// NewTelemetryDruidDefaultDimensionSpec instantiates a new TelemetryDruidDefaultDimensionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidDefaultDimensionSpec(type_ string, dimension string, outputName string, outputType string) *TelemetryDruidDefaultDimensionSpec {
	this := TelemetryDruidDefaultDimensionSpec{}
	this.Type = type_
	this.Dimension = dimension
	this.OutputName = outputName
	this.OutputType = outputType
	return &this
}

// NewTelemetryDruidDefaultDimensionSpecWithDefaults instantiates a new TelemetryDruidDefaultDimensionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidDefaultDimensionSpecWithDefaults() *TelemetryDruidDefaultDimensionSpec {
	this := TelemetryDruidDefaultDimensionSpec{}
	var outputType string = "STRING"
	this.OutputType = outputType
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidDefaultDimensionSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDefaultDimensionSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidDefaultDimensionSpec) SetType(v string) {
	o.Type = v
}

// GetDimension returns the Dimension field value
func (o *TelemetryDruidDefaultDimensionSpec) GetDimension() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDefaultDimensionSpec) GetDimensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *TelemetryDruidDefaultDimensionSpec) SetDimension(v string) {
	o.Dimension = v
}

// GetOutputName returns the OutputName field value
func (o *TelemetryDruidDefaultDimensionSpec) GetOutputName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutputName
}

// GetOutputNameOk returns a tuple with the OutputName field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDefaultDimensionSpec) GetOutputNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputName, true
}

// SetOutputName sets field value
func (o *TelemetryDruidDefaultDimensionSpec) SetOutputName(v string) {
	o.OutputName = v
}

// GetOutputType returns the OutputType field value
func (o *TelemetryDruidDefaultDimensionSpec) GetOutputType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutputType
}

// GetOutputTypeOk returns a tuple with the OutputType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDefaultDimensionSpec) GetOutputTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputType, true
}

// SetOutputType sets field value
func (o *TelemetryDruidDefaultDimensionSpec) SetOutputType(v string) {
	o.OutputType = v
}

func (o TelemetryDruidDefaultDimensionSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["dimension"] = o.Dimension
	}
	if true {
		toSerialize["outputName"] = o.OutputName
	}
	if true {
		toSerialize["outputType"] = o.OutputType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidDefaultDimensionSpec) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidDefaultDimensionSpec := _TelemetryDruidDefaultDimensionSpec{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidDefaultDimensionSpec); err == nil {
		*o = TelemetryDruidDefaultDimensionSpec(varTelemetryDruidDefaultDimensionSpec)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "dimension")
		delete(additionalProperties, "outputName")
		delete(additionalProperties, "outputType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidDefaultDimensionSpec struct {
	value *TelemetryDruidDefaultDimensionSpec
	isSet bool
}

func (v NullableTelemetryDruidDefaultDimensionSpec) Get() *TelemetryDruidDefaultDimensionSpec {
	return v.value
}

func (v *NullableTelemetryDruidDefaultDimensionSpec) Set(val *TelemetryDruidDefaultDimensionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidDefaultDimensionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidDefaultDimensionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidDefaultDimensionSpec(val *TelemetryDruidDefaultDimensionSpec) *NullableTelemetryDruidDefaultDimensionSpec {
	return &NullableTelemetryDruidDefaultDimensionSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidDefaultDimensionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidDefaultDimensionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}