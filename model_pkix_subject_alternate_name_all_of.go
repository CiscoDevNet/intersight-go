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

// PkixSubjectAlternateNameAllOf Definition of the list of properties defined in 'pkix.SubjectAlternateName', excluding properties defined in parent classes.
type PkixSubjectAlternateNameAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string   `json:"ObjectType"`
	DnsName              []string `json:"DnsName,omitempty"`
	EmailAddress         []string `json:"EmailAddress,omitempty"`
	IpAddress            []string `json:"IpAddress,omitempty"`
	Uri                  []string `json:"Uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PkixSubjectAlternateNameAllOf PkixSubjectAlternateNameAllOf

// NewPkixSubjectAlternateNameAllOf instantiates a new PkixSubjectAlternateNameAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixSubjectAlternateNameAllOf(classId string, objectType string) *PkixSubjectAlternateNameAllOf {
	this := PkixSubjectAlternateNameAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPkixSubjectAlternateNameAllOfWithDefaults instantiates a new PkixSubjectAlternateNameAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixSubjectAlternateNameAllOfWithDefaults() *PkixSubjectAlternateNameAllOf {
	this := PkixSubjectAlternateNameAllOf{}
	var classId string = "pkix.SubjectAlternateName"
	this.ClassId = classId
	var objectType string = "pkix.SubjectAlternateName"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PkixSubjectAlternateNameAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PkixSubjectAlternateNameAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PkixSubjectAlternateNameAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PkixSubjectAlternateNameAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PkixSubjectAlternateNameAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PkixSubjectAlternateNameAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixSubjectAlternateNameAllOf) GetDnsName() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixSubjectAlternateNameAllOf) GetDnsNameOk() ([]string, bool) {
	if o == nil || o.DnsName == nil {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *PkixSubjectAlternateNameAllOf) HasDnsName() bool {
	if o != nil && o.DnsName != nil {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given []string and assigns it to the DnsName field.
func (o *PkixSubjectAlternateNameAllOf) SetDnsName(v []string) {
	o.DnsName = v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixSubjectAlternateNameAllOf) GetEmailAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixSubjectAlternateNameAllOf) GetEmailAddressOk() ([]string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *PkixSubjectAlternateNameAllOf) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given []string and assigns it to the EmailAddress field.
func (o *PkixSubjectAlternateNameAllOf) SetEmailAddress(v []string) {
	o.EmailAddress = v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixSubjectAlternateNameAllOf) GetIpAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixSubjectAlternateNameAllOf) GetIpAddressOk() ([]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *PkixSubjectAlternateNameAllOf) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *PkixSubjectAlternateNameAllOf) SetIpAddress(v []string) {
	o.IpAddress = v
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixSubjectAlternateNameAllOf) GetUri() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixSubjectAlternateNameAllOf) GetUriOk() ([]string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *PkixSubjectAlternateNameAllOf) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given []string and assigns it to the Uri field.
func (o *PkixSubjectAlternateNameAllOf) SetUri(v []string) {
	o.Uri = v
}

func (o PkixSubjectAlternateNameAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DnsName != nil {
		toSerialize["DnsName"] = o.DnsName
	}
	if o.EmailAddress != nil {
		toSerialize["EmailAddress"] = o.EmailAddress
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Uri != nil {
		toSerialize["Uri"] = o.Uri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PkixSubjectAlternateNameAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPkixSubjectAlternateNameAllOf := _PkixSubjectAlternateNameAllOf{}

	if err = json.Unmarshal(bytes, &varPkixSubjectAlternateNameAllOf); err == nil {
		*o = PkixSubjectAlternateNameAllOf(varPkixSubjectAlternateNameAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DnsName")
		delete(additionalProperties, "EmailAddress")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "Uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePkixSubjectAlternateNameAllOf struct {
	value *PkixSubjectAlternateNameAllOf
	isSet bool
}

func (v NullablePkixSubjectAlternateNameAllOf) Get() *PkixSubjectAlternateNameAllOf {
	return v.value
}

func (v *NullablePkixSubjectAlternateNameAllOf) Set(val *PkixSubjectAlternateNameAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixSubjectAlternateNameAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixSubjectAlternateNameAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixSubjectAlternateNameAllOf(val *PkixSubjectAlternateNameAllOf) *NullablePkixSubjectAlternateNameAllOf {
	return &NullablePkixSubjectAlternateNameAllOf{value: val, isSet: true}
}

func (v NullablePkixSubjectAlternateNameAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixSubjectAlternateNameAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}