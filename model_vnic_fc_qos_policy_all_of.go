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

// VnicFcQosPolicyAllOf Definition of the list of properties defined in 'vnic.FcQosPolicy', excluding properties defined in parent classes.
type VnicFcQosPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The burst traffic, in bytes, allowed on the vHBA.
	Burst *int64 `json:"Burst,omitempty"`
	// Class of Service to be associated to the traffic on the virtual interface.
	Cos *int64 `json:"Cos,omitempty"`
	// The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports.
	MaxDataFieldSize *int64 `json:"MaxDataFieldSize,omitempty"`
	// The priortity matching the System QoS specified in the fabric profile. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
	Priority *string `json:"Priority,omitempty"`
	// The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
	RateLimit            *int64                                `json:"RateLimit,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcQosPolicyAllOf VnicFcQosPolicyAllOf

// NewVnicFcQosPolicyAllOf instantiates a new VnicFcQosPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcQosPolicyAllOf(classId string, objectType string) *VnicFcQosPolicyAllOf {
	this := VnicFcQosPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var burst int64 = 10240
	this.Burst = &burst
	var cos int64 = 3
	this.Cos = &cos
	var maxDataFieldSize int64 = 2112
	this.MaxDataFieldSize = &maxDataFieldSize
	var rateLimit int64 = 0
	this.RateLimit = &rateLimit
	return &this
}

// NewVnicFcQosPolicyAllOfWithDefaults instantiates a new VnicFcQosPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcQosPolicyAllOfWithDefaults() *VnicFcQosPolicyAllOf {
	this := VnicFcQosPolicyAllOf{}
	var classId string = "vnic.FcQosPolicy"
	this.ClassId = classId
	var objectType string = "vnic.FcQosPolicy"
	this.ObjectType = objectType
	var burst int64 = 10240
	this.Burst = &burst
	var cos int64 = 3
	this.Cos = &cos
	var maxDataFieldSize int64 = 2112
	this.MaxDataFieldSize = &maxDataFieldSize
	var rateLimit int64 = 0
	this.RateLimit = &rateLimit
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcQosPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcQosPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcQosPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcQosPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBurst returns the Burst field value if set, zero value otherwise.
func (o *VnicFcQosPolicyAllOf) GetBurst() int64 {
	if o == nil || o.Burst == nil {
		var ret int64
		return ret
	}
	return *o.Burst
}

// GetBurstOk returns a tuple with the Burst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyAllOf) GetBurstOk() (*int64, bool) {
	if o == nil || o.Burst == nil {
		return nil, false
	}
	return o.Burst, true
}

// HasBurst returns a boolean if a field has been set.
func (o *VnicFcQosPolicyAllOf) HasBurst() bool {
	if o != nil && o.Burst != nil {
		return true
	}

	return false
}

// SetBurst gets a reference to the given int64 and assigns it to the Burst field.
func (o *VnicFcQosPolicyAllOf) SetBurst(v int64) {
	o.Burst = &v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *VnicFcQosPolicyAllOf) GetCos() int64 {
	if o == nil || o.Cos == nil {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyAllOf) GetCosOk() (*int64, bool) {
	if o == nil || o.Cos == nil {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *VnicFcQosPolicyAllOf) HasCos() bool {
	if o != nil && o.Cos != nil {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *VnicFcQosPolicyAllOf) SetCos(v int64) {
	o.Cos = &v
}

// GetMaxDataFieldSize returns the MaxDataFieldSize field value if set, zero value otherwise.
func (o *VnicFcQosPolicyAllOf) GetMaxDataFieldSize() int64 {
	if o == nil || o.MaxDataFieldSize == nil {
		var ret int64
		return ret
	}
	return *o.MaxDataFieldSize
}

// GetMaxDataFieldSizeOk returns a tuple with the MaxDataFieldSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyAllOf) GetMaxDataFieldSizeOk() (*int64, bool) {
	if o == nil || o.MaxDataFieldSize == nil {
		return nil, false
	}
	return o.MaxDataFieldSize, true
}

// HasMaxDataFieldSize returns a boolean if a field has been set.
func (o *VnicFcQosPolicyAllOf) HasMaxDataFieldSize() bool {
	if o != nil && o.MaxDataFieldSize != nil {
		return true
	}

	return false
}

// SetMaxDataFieldSize gets a reference to the given int64 and assigns it to the MaxDataFieldSize field.
func (o *VnicFcQosPolicyAllOf) SetMaxDataFieldSize(v int64) {
	o.MaxDataFieldSize = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *VnicFcQosPolicyAllOf) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyAllOf) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *VnicFcQosPolicyAllOf) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *VnicFcQosPolicyAllOf) SetPriority(v string) {
	o.Priority = &v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *VnicFcQosPolicyAllOf) GetRateLimit() int64 {
	if o == nil || o.RateLimit == nil {
		var ret int64
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyAllOf) GetRateLimitOk() (*int64, bool) {
	if o == nil || o.RateLimit == nil {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *VnicFcQosPolicyAllOf) HasRateLimit() bool {
	if o != nil && o.RateLimit != nil {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given int64 and assigns it to the RateLimit field.
func (o *VnicFcQosPolicyAllOf) SetRateLimit(v int64) {
	o.RateLimit = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicFcQosPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicFcQosPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicFcQosPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o VnicFcQosPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Burst != nil {
		toSerialize["Burst"] = o.Burst
	}
	if o.Cos != nil {
		toSerialize["Cos"] = o.Cos
	}
	if o.MaxDataFieldSize != nil {
		toSerialize["MaxDataFieldSize"] = o.MaxDataFieldSize
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.RateLimit != nil {
		toSerialize["RateLimit"] = o.RateLimit
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicFcQosPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicFcQosPolicyAllOf := _VnicFcQosPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varVnicFcQosPolicyAllOf); err == nil {
		*o = VnicFcQosPolicyAllOf(varVnicFcQosPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Burst")
		delete(additionalProperties, "Cos")
		delete(additionalProperties, "MaxDataFieldSize")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "RateLimit")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicFcQosPolicyAllOf struct {
	value *VnicFcQosPolicyAllOf
	isSet bool
}

func (v NullableVnicFcQosPolicyAllOf) Get() *VnicFcQosPolicyAllOf {
	return v.value
}

func (v *NullableVnicFcQosPolicyAllOf) Set(val *VnicFcQosPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcQosPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcQosPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcQosPolicyAllOf(val *VnicFcQosPolicyAllOf) *NullableVnicFcQosPolicyAllOf {
	return &NullableVnicFcQosPolicyAllOf{value: val, isSet: true}
}

func (v NullableVnicFcQosPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcQosPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}