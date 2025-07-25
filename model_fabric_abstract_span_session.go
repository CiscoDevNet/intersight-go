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

// checks if the FabricAbstractSpanSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricAbstractSpanSession{}

// FabricAbstractSpanSession Abstract class for configuring a SPAN session per Fabric Interconnect.
type FabricAbstractSpanSession struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Admin state to enable or disable the SPAN session. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
	AdminState *string `json:"AdminState,omitempty"`
	// The configured state of SPAN configuration. If the configuration fails to deploy to the Fabric Interconnect, it can be redeployed by toggling the admin state. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	ConfigState *string `json:"ConfigState,omitempty"`
	// Name of the SPAN session. The name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_).
	Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.:-]{1,64}$"`
	// Operational state of the SPAN session. * `Unknown` - SPAN session operational state is Unknown. * `Up` - SPAN session operational state is Up. * `Down` - SPAN session operational state is Down. * `Error` - SPAN session operational state is Error.
	OperState *string `json:"OperState,omitempty"`
	// Operational state reason of the SPAN session. * `Unknown` - Unknown operational state. * `Active` - Active and operational SPAN session. * `NoHardwareResource` - No hardware resources available. * `NoOperationalSrcDst` - No operational SPAN source or destination. * `GenericError` - Generic operational error. * `NoSourcesConfigured` - No source interfaces configured. * `NoDestinationConfigured` - No destination port configured. * `NoSourceDestinationConfigured` - No source or destination interface configured. * `SessionAdminShut` - Session is administratively disabled. * `WrongDestinationMode` - Wrong Destination mode configured. * `WrongSourceMode` - Wrong Source mode configured. * `TunnelMisconfDown` - Tunnel Misconfigured or Down. * `NoFlowIdSpecified` - No Flow ID specified for ERSPAN.
	OperStateReason *string `json:"OperStateReason,omitempty"`
	// Session ID identifies the SPAN session on the Fabric Interconnect.
	SessionId *int64 `json:"SessionId,omitempty"`
	// Count of total number of sources in the SPAN session.
	SourceCount *int64 `json:"SourceCount,omitempty"`
	// Admin state to enable or disable spanning control packets. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
	SpanControlPackets *string `json:"SpanControlPackets,omitempty"`
	// Switch ID (A or B) corresponding to network element. * `A` - Switch Identifier of Fabric Interconnect A. * `B` - Switch Identifier of Fabric Interconnect B.
	SwitchId             *string                            `json:"SwitchId,omitempty"`
	NetworkElement       NullableNetworkElementRelationship `json:"NetworkElement,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricAbstractSpanSession FabricAbstractSpanSession

// NewFabricAbstractSpanSession instantiates a new FabricAbstractSpanSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricAbstractSpanSession(classId string, objectType string) *FabricAbstractSpanSession {
	this := FabricAbstractSpanSession{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminState string = "Disabled"
	this.AdminState = &adminState
	var spanControlPackets string = "Disabled"
	this.SpanControlPackets = &spanControlPackets
	return &this
}

// NewFabricAbstractSpanSessionWithDefaults instantiates a new FabricAbstractSpanSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricAbstractSpanSessionWithDefaults() *FabricAbstractSpanSession {
	this := FabricAbstractSpanSession{}
	var classId string = "fabric.SpanSession"
	this.ClassId = classId
	var objectType string = "fabric.SpanSession"
	this.ObjectType = objectType
	var adminState string = "Disabled"
	this.AdminState = &adminState
	var spanControlPackets string = "Disabled"
	this.SpanControlPackets = &spanControlPackets
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricAbstractSpanSession) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSession) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricAbstractSpanSession) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.SpanSession" of the ClassId field.
func (o *FabricAbstractSpanSession) GetDefaultClassId() interface{} {
	return "fabric.SpanSession"
}

// GetObjectType returns the ObjectType field value
func (o *FabricAbstractSpanSession) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSession) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricAbstractSpanSession) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.SpanSession" of the ObjectType field.
func (o *FabricAbstractSpanSession) GetDefaultObjectType() interface{} {
	return "fabric.SpanSession"
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *FabricAbstractSpanSession) GetAdminState() string {
	if o == nil || IsNil(o.AdminState) {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSession) GetAdminStateOk() (*string, bool) {
	if o == nil || IsNil(o.AdminState) {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *FabricAbstractSpanSession) HasAdminState() bool {
	if o != nil && !IsNil(o.AdminState) {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *FabricAbstractSpanSession) SetAdminState(v string) {
	o.AdminState = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *FabricAbstractSpanSession) GetConfigState() string {
	if o == nil || IsNil(o.ConfigState) {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSession) GetConfigStateOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigState) {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *FabricAbstractSpanSession) HasConfigState() bool {
	if o != nil && !IsNil(o.ConfigState) {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *FabricAbstractSpanSession) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricAbstractSpanSession) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSession) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricAbstractSpanSession) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricAbstractSpanSession) SetName(v string) {
	o.Name = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *FabricAbstractSpanSession) GetOperState() string {
	if o == nil || IsNil(o.OperState) {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSession) GetOperStateOk() (*string, bool) {
	if o == nil || IsNil(o.OperState) {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *FabricAbstractSpanSession) HasOperState() bool {
	if o != nil && !IsNil(o.OperState) {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *FabricAbstractSpanSession) SetOperState(v string) {
	o.OperState = &v
}

// GetOperStateReason returns the OperStateReason field value if set, zero value otherwise.
func (o *FabricAbstractSpanSession) GetOperStateReason() string {
	if o == nil || IsNil(o.OperStateReason) {
		var ret string
		return ret
	}
	return *o.OperStateReason
}

// GetOperStateReasonOk returns a tuple with the OperStateReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSession) GetOperStateReasonOk() (*string, bool) {
	if o == nil || IsNil(o.OperStateReason) {
		return nil, false
	}
	return o.OperStateReason, true
}

// HasOperStateReason returns a boolean if a field has been set.
func (o *FabricAbstractSpanSession) HasOperStateReason() bool {
	if o != nil && !IsNil(o.OperStateReason) {
		return true
	}

	return false
}

// SetOperStateReason gets a reference to the given string and assigns it to the OperStateReason field.
func (o *FabricAbstractSpanSession) SetOperStateReason(v string) {
	o.OperStateReason = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *FabricAbstractSpanSession) GetSessionId() int64 {
	if o == nil || IsNil(o.SessionId) {
		var ret int64
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSession) GetSessionIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *FabricAbstractSpanSession) HasSessionId() bool {
	if o != nil && !IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given int64 and assigns it to the SessionId field.
func (o *FabricAbstractSpanSession) SetSessionId(v int64) {
	o.SessionId = &v
}

// GetSourceCount returns the SourceCount field value if set, zero value otherwise.
func (o *FabricAbstractSpanSession) GetSourceCount() int64 {
	if o == nil || IsNil(o.SourceCount) {
		var ret int64
		return ret
	}
	return *o.SourceCount
}

// GetSourceCountOk returns a tuple with the SourceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSession) GetSourceCountOk() (*int64, bool) {
	if o == nil || IsNil(o.SourceCount) {
		return nil, false
	}
	return o.SourceCount, true
}

// HasSourceCount returns a boolean if a field has been set.
func (o *FabricAbstractSpanSession) HasSourceCount() bool {
	if o != nil && !IsNil(o.SourceCount) {
		return true
	}

	return false
}

// SetSourceCount gets a reference to the given int64 and assigns it to the SourceCount field.
func (o *FabricAbstractSpanSession) SetSourceCount(v int64) {
	o.SourceCount = &v
}

// GetSpanControlPackets returns the SpanControlPackets field value if set, zero value otherwise.
func (o *FabricAbstractSpanSession) GetSpanControlPackets() string {
	if o == nil || IsNil(o.SpanControlPackets) {
		var ret string
		return ret
	}
	return *o.SpanControlPackets
}

// GetSpanControlPacketsOk returns a tuple with the SpanControlPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSession) GetSpanControlPacketsOk() (*string, bool) {
	if o == nil || IsNil(o.SpanControlPackets) {
		return nil, false
	}
	return o.SpanControlPackets, true
}

// HasSpanControlPackets returns a boolean if a field has been set.
func (o *FabricAbstractSpanSession) HasSpanControlPackets() bool {
	if o != nil && !IsNil(o.SpanControlPackets) {
		return true
	}

	return false
}

// SetSpanControlPackets gets a reference to the given string and assigns it to the SpanControlPackets field.
func (o *FabricAbstractSpanSession) SetSpanControlPackets(v string) {
	o.SpanControlPackets = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *FabricAbstractSpanSession) GetSwitchId() string {
	if o == nil || IsNil(o.SwitchId) {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSession) GetSwitchIdOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchId) {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *FabricAbstractSpanSession) HasSwitchId() bool {
	if o != nil && !IsNil(o.SwitchId) {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *FabricAbstractSpanSession) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricAbstractSpanSession) GetNetworkElement() NetworkElementRelationship {
	if o == nil || IsNil(o.NetworkElement.Get()) {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement.Get()
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricAbstractSpanSession) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkElement.Get(), o.NetworkElement.IsSet()
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *FabricAbstractSpanSession) HasNetworkElement() bool {
	if o != nil && o.NetworkElement.IsSet() {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NullableNetworkElementRelationship and assigns it to the NetworkElement field.
func (o *FabricAbstractSpanSession) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement.Set(&v)
}

// SetNetworkElementNil sets the value for NetworkElement to be an explicit nil
func (o *FabricAbstractSpanSession) SetNetworkElementNil() {
	o.NetworkElement.Set(nil)
}

// UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
func (o *FabricAbstractSpanSession) UnsetNetworkElement() {
	o.NetworkElement.Unset()
}

func (o FabricAbstractSpanSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricAbstractSpanSession) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AdminState) {
		toSerialize["AdminState"] = o.AdminState
	}
	if !IsNil(o.ConfigState) {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.OperState) {
		toSerialize["OperState"] = o.OperState
	}
	if !IsNil(o.OperStateReason) {
		toSerialize["OperStateReason"] = o.OperStateReason
	}
	if !IsNil(o.SessionId) {
		toSerialize["SessionId"] = o.SessionId
	}
	if !IsNil(o.SourceCount) {
		toSerialize["SourceCount"] = o.SourceCount
	}
	if !IsNil(o.SpanControlPackets) {
		toSerialize["SpanControlPackets"] = o.SpanControlPackets
	}
	if !IsNil(o.SwitchId) {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.NetworkElement.IsSet() {
		toSerialize["NetworkElement"] = o.NetworkElement.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricAbstractSpanSession) UnmarshalJSON(data []byte) (err error) {
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
	type FabricAbstractSpanSessionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Admin state to enable or disable the SPAN session. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
		AdminState *string `json:"AdminState,omitempty"`
		// The configured state of SPAN configuration. If the configuration fails to deploy to the Fabric Interconnect, it can be redeployed by toggling the admin state. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
		ConfigState *string `json:"ConfigState,omitempty"`
		// Name of the SPAN session. The name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_).
		Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.:-]{1,64}$"`
		// Operational state of the SPAN session. * `Unknown` - SPAN session operational state is Unknown. * `Up` - SPAN session operational state is Up. * `Down` - SPAN session operational state is Down. * `Error` - SPAN session operational state is Error.
		OperState *string `json:"OperState,omitempty"`
		// Operational state reason of the SPAN session. * `Unknown` - Unknown operational state. * `Active` - Active and operational SPAN session. * `NoHardwareResource` - No hardware resources available. * `NoOperationalSrcDst` - No operational SPAN source or destination. * `GenericError` - Generic operational error. * `NoSourcesConfigured` - No source interfaces configured. * `NoDestinationConfigured` - No destination port configured. * `NoSourceDestinationConfigured` - No source or destination interface configured. * `SessionAdminShut` - Session is administratively disabled. * `WrongDestinationMode` - Wrong Destination mode configured. * `WrongSourceMode` - Wrong Source mode configured. * `TunnelMisconfDown` - Tunnel Misconfigured or Down. * `NoFlowIdSpecified` - No Flow ID specified for ERSPAN.
		OperStateReason *string `json:"OperStateReason,omitempty"`
		// Session ID identifies the SPAN session on the Fabric Interconnect.
		SessionId *int64 `json:"SessionId,omitempty"`
		// Count of total number of sources in the SPAN session.
		SourceCount *int64 `json:"SourceCount,omitempty"`
		// Admin state to enable or disable spanning control packets. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
		SpanControlPackets *string `json:"SpanControlPackets,omitempty"`
		// Switch ID (A or B) corresponding to network element. * `A` - Switch Identifier of Fabric Interconnect A. * `B` - Switch Identifier of Fabric Interconnect B.
		SwitchId       *string                            `json:"SwitchId,omitempty"`
		NetworkElement NullableNetworkElementRelationship `json:"NetworkElement,omitempty"`
	}

	varFabricAbstractSpanSessionWithoutEmbeddedStruct := FabricAbstractSpanSessionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricAbstractSpanSessionWithoutEmbeddedStruct)
	if err == nil {
		varFabricAbstractSpanSession := _FabricAbstractSpanSession{}
		varFabricAbstractSpanSession.ClassId = varFabricAbstractSpanSessionWithoutEmbeddedStruct.ClassId
		varFabricAbstractSpanSession.ObjectType = varFabricAbstractSpanSessionWithoutEmbeddedStruct.ObjectType
		varFabricAbstractSpanSession.AdminState = varFabricAbstractSpanSessionWithoutEmbeddedStruct.AdminState
		varFabricAbstractSpanSession.ConfigState = varFabricAbstractSpanSessionWithoutEmbeddedStruct.ConfigState
		varFabricAbstractSpanSession.Name = varFabricAbstractSpanSessionWithoutEmbeddedStruct.Name
		varFabricAbstractSpanSession.OperState = varFabricAbstractSpanSessionWithoutEmbeddedStruct.OperState
		varFabricAbstractSpanSession.OperStateReason = varFabricAbstractSpanSessionWithoutEmbeddedStruct.OperStateReason
		varFabricAbstractSpanSession.SessionId = varFabricAbstractSpanSessionWithoutEmbeddedStruct.SessionId
		varFabricAbstractSpanSession.SourceCount = varFabricAbstractSpanSessionWithoutEmbeddedStruct.SourceCount
		varFabricAbstractSpanSession.SpanControlPackets = varFabricAbstractSpanSessionWithoutEmbeddedStruct.SpanControlPackets
		varFabricAbstractSpanSession.SwitchId = varFabricAbstractSpanSessionWithoutEmbeddedStruct.SwitchId
		varFabricAbstractSpanSession.NetworkElement = varFabricAbstractSpanSessionWithoutEmbeddedStruct.NetworkElement
		*o = FabricAbstractSpanSession(varFabricAbstractSpanSession)
	} else {
		return err
	}

	varFabricAbstractSpanSession := _FabricAbstractSpanSession{}

	err = json.Unmarshal(data, &varFabricAbstractSpanSession)
	if err == nil {
		o.MoBaseMo = varFabricAbstractSpanSession.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "OperStateReason")
		delete(additionalProperties, "SessionId")
		delete(additionalProperties, "SourceCount")
		delete(additionalProperties, "SpanControlPackets")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "NetworkElement")

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

type NullableFabricAbstractSpanSession struct {
	value *FabricAbstractSpanSession
	isSet bool
}

func (v NullableFabricAbstractSpanSession) Get() *FabricAbstractSpanSession {
	return v.value
}

func (v *NullableFabricAbstractSpanSession) Set(val *FabricAbstractSpanSession) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricAbstractSpanSession) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricAbstractSpanSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricAbstractSpanSession(val *FabricAbstractSpanSession) *NullableFabricAbstractSpanSession {
	return &NullableFabricAbstractSpanSession{value: val, isSet: true}
}

func (v NullableFabricAbstractSpanSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricAbstractSpanSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
