/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14237
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ResourceSharedResourcesInfoHolderAllOf Definition of the list of properties defined in 'resource.SharedResourcesInfoHolder', excluding properties defined in parent classes.
type ResourceSharedResourcesInfoHolderAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                  `json:"ObjectType"`
	Account    *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to moBaseMo resources.
	PeerObjects          []MoBaseMoRelationship      `json:"PeerObjects,omitempty"`
	SharedResource       *MoBaseMoRelationship       `json:"SharedResource,omitempty"`
	SharedWithResource   *MoBaseMoRelationship       `json:"SharedWithResource,omitempty"`
	SharingRule          *IamSharingRuleRelationship `json:"SharingRule,omitempty"`
	SourceObject         *MoBaseMoRelationship       `json:"SourceObject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSharedResourcesInfoHolderAllOf ResourceSharedResourcesInfoHolderAllOf

// NewResourceSharedResourcesInfoHolderAllOf instantiates a new ResourceSharedResourcesInfoHolderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSharedResourcesInfoHolderAllOf(classId string, objectType string) *ResourceSharedResourcesInfoHolderAllOf {
	this := ResourceSharedResourcesInfoHolderAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourceSharedResourcesInfoHolderAllOfWithDefaults instantiates a new ResourceSharedResourcesInfoHolderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSharedResourcesInfoHolderAllOfWithDefaults() *ResourceSharedResourcesInfoHolderAllOf {
	this := ResourceSharedResourcesInfoHolderAllOf{}
	var classId string = "resource.SharedResourcesInfoHolder"
	this.ClassId = classId
	var objectType string = "resource.SharedResourcesInfoHolder"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceSharedResourcesInfoHolderAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceSharedResourcesInfoHolderAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourceSharedResourcesInfoHolderAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceSharedResourcesInfoHolderAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ResourceSharedResourcesInfoHolderAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ResourceSharedResourcesInfoHolderAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetPeerObjects returns the PeerObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceSharedResourcesInfoHolderAllOf) GetPeerObjects() []MoBaseMoRelationship {
	if o == nil {
		var ret []MoBaseMoRelationship
		return ret
	}
	return o.PeerObjects
}

// GetPeerObjectsOk returns a tuple with the PeerObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceSharedResourcesInfoHolderAllOf) GetPeerObjectsOk() ([]MoBaseMoRelationship, bool) {
	if o == nil || o.PeerObjects == nil {
		return nil, false
	}
	return o.PeerObjects, true
}

// HasPeerObjects returns a boolean if a field has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) HasPeerObjects() bool {
	if o != nil && o.PeerObjects != nil {
		return true
	}

	return false
}

// SetPeerObjects gets a reference to the given []MoBaseMoRelationship and assigns it to the PeerObjects field.
func (o *ResourceSharedResourcesInfoHolderAllOf) SetPeerObjects(v []MoBaseMoRelationship) {
	o.PeerObjects = v
}

// GetSharedResource returns the SharedResource field value if set, zero value otherwise.
func (o *ResourceSharedResourcesInfoHolderAllOf) GetSharedResource() MoBaseMoRelationship {
	if o == nil || o.SharedResource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.SharedResource
}

// GetSharedResourceOk returns a tuple with the SharedResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) GetSharedResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.SharedResource == nil {
		return nil, false
	}
	return o.SharedResource, true
}

// HasSharedResource returns a boolean if a field has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) HasSharedResource() bool {
	if o != nil && o.SharedResource != nil {
		return true
	}

	return false
}

// SetSharedResource gets a reference to the given MoBaseMoRelationship and assigns it to the SharedResource field.
func (o *ResourceSharedResourcesInfoHolderAllOf) SetSharedResource(v MoBaseMoRelationship) {
	o.SharedResource = &v
}

// GetSharedWithResource returns the SharedWithResource field value if set, zero value otherwise.
func (o *ResourceSharedResourcesInfoHolderAllOf) GetSharedWithResource() MoBaseMoRelationship {
	if o == nil || o.SharedWithResource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.SharedWithResource
}

// GetSharedWithResourceOk returns a tuple with the SharedWithResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) GetSharedWithResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.SharedWithResource == nil {
		return nil, false
	}
	return o.SharedWithResource, true
}

// HasSharedWithResource returns a boolean if a field has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) HasSharedWithResource() bool {
	if o != nil && o.SharedWithResource != nil {
		return true
	}

	return false
}

// SetSharedWithResource gets a reference to the given MoBaseMoRelationship and assigns it to the SharedWithResource field.
func (o *ResourceSharedResourcesInfoHolderAllOf) SetSharedWithResource(v MoBaseMoRelationship) {
	o.SharedWithResource = &v
}

// GetSharingRule returns the SharingRule field value if set, zero value otherwise.
func (o *ResourceSharedResourcesInfoHolderAllOf) GetSharingRule() IamSharingRuleRelationship {
	if o == nil || o.SharingRule == nil {
		var ret IamSharingRuleRelationship
		return ret
	}
	return *o.SharingRule
}

// GetSharingRuleOk returns a tuple with the SharingRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) GetSharingRuleOk() (*IamSharingRuleRelationship, bool) {
	if o == nil || o.SharingRule == nil {
		return nil, false
	}
	return o.SharingRule, true
}

// HasSharingRule returns a boolean if a field has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) HasSharingRule() bool {
	if o != nil && o.SharingRule != nil {
		return true
	}

	return false
}

// SetSharingRule gets a reference to the given IamSharingRuleRelationship and assigns it to the SharingRule field.
func (o *ResourceSharedResourcesInfoHolderAllOf) SetSharingRule(v IamSharingRuleRelationship) {
	o.SharingRule = &v
}

// GetSourceObject returns the SourceObject field value if set, zero value otherwise.
func (o *ResourceSharedResourcesInfoHolderAllOf) GetSourceObject() MoBaseMoRelationship {
	if o == nil || o.SourceObject == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.SourceObject
}

// GetSourceObjectOk returns a tuple with the SourceObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) GetSourceObjectOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.SourceObject == nil {
		return nil, false
	}
	return o.SourceObject, true
}

// HasSourceObject returns a boolean if a field has been set.
func (o *ResourceSharedResourcesInfoHolderAllOf) HasSourceObject() bool {
	if o != nil && o.SourceObject != nil {
		return true
	}

	return false
}

// SetSourceObject gets a reference to the given MoBaseMoRelationship and assigns it to the SourceObject field.
func (o *ResourceSharedResourcesInfoHolderAllOf) SetSourceObject(v MoBaseMoRelationship) {
	o.SourceObject = &v
}

func (o ResourceSharedResourcesInfoHolderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.PeerObjects != nil {
		toSerialize["PeerObjects"] = o.PeerObjects
	}
	if o.SharedResource != nil {
		toSerialize["SharedResource"] = o.SharedResource
	}
	if o.SharedWithResource != nil {
		toSerialize["SharedWithResource"] = o.SharedWithResource
	}
	if o.SharingRule != nil {
		toSerialize["SharingRule"] = o.SharingRule
	}
	if o.SourceObject != nil {
		toSerialize["SourceObject"] = o.SourceObject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSharedResourcesInfoHolderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSharedResourcesInfoHolderAllOf := _ResourceSharedResourcesInfoHolderAllOf{}

	if err = json.Unmarshal(bytes, &varResourceSharedResourcesInfoHolderAllOf); err == nil {
		*o = ResourceSharedResourcesInfoHolderAllOf(varResourceSharedResourcesInfoHolderAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "PeerObjects")
		delete(additionalProperties, "SharedResource")
		delete(additionalProperties, "SharedWithResource")
		delete(additionalProperties, "SharingRule")
		delete(additionalProperties, "SourceObject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceSharedResourcesInfoHolderAllOf struct {
	value *ResourceSharedResourcesInfoHolderAllOf
	isSet bool
}

func (v NullableResourceSharedResourcesInfoHolderAllOf) Get() *ResourceSharedResourcesInfoHolderAllOf {
	return v.value
}

func (v *NullableResourceSharedResourcesInfoHolderAllOf) Set(val *ResourceSharedResourcesInfoHolderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSharedResourcesInfoHolderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSharedResourcesInfoHolderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSharedResourcesInfoHolderAllOf(val *ResourceSharedResourcesInfoHolderAllOf) *NullableResourceSharedResourcesInfoHolderAllOf {
	return &NullableResourceSharedResourcesInfoHolderAllOf{value: val, isSet: true}
}

func (v NullableResourceSharedResourcesInfoHolderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSharedResourcesInfoHolderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}