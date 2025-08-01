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

// checks if the EquipmentIoCardBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentIoCardBase{}

// EquipmentIoCardBase Abstract class for hif and nif ports collection that can be extended by chassis/febric extender.
type EquipmentIoCardBase struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Connectivity Status of FEX/IOM to Switch - A or B or AB.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// This field is to provide description for the iocard module model.
	Description *string `json:"Description,omitempty"`
	// Module Identifier for the IO module.
	ModuleId   *int64   `json:"ModuleId,omitempty"`
	OperReason []string `json:"OperReason,omitempty"`
	// Operational state of IO card or fabric extender.
	OperState *string `json:"OperState,omitempty"`
	// Part Number identifier for the IO module.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the Product ID for the IO module.
	Pid *string `json:"Pid,omitempty"`
	// This field identifies the Product Name for the iocard module model.
	ProductName *string `json:"ProductName,omitempty"`
	// This field identifies the Stock Keeping Unit for the IO card module.
	Sku *string `json:"Sku,omitempty"`
	// This field identifies the version of the IO card module.
	Version *string `json:"Version,omitempty"`
	// This field identifies the Vendor ID for the IO card module.
	Vid *string `json:"Vid,omitempty"`
	// An array of relationships to etherHostPort resources.
	HostPorts      []EtherHostPortRelationship              `json:"HostPorts,omitempty"`
	MgmtController NullableManagementControllerRelationship `json:"MgmtController,omitempty"`
	// An array of relationships to etherNetworkPort resources.
	NetworkPorts []EtherNetworkPortRelationship `json:"NetworkPorts,omitempty"`
	// An array of relationships to portSubGroup resources.
	SubGroup             []PortSubGroupRelationship `json:"SubGroup,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentIoCardBase EquipmentIoCardBase

// NewEquipmentIoCardBase instantiates a new EquipmentIoCardBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIoCardBase(classId string, objectType string) *EquipmentIoCardBase {
	this := EquipmentIoCardBase{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentIoCardBaseWithDefaults instantiates a new EquipmentIoCardBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIoCardBaseWithDefaults() *EquipmentIoCardBase {
	this := EquipmentIoCardBase{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIoCardBase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIoCardBase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIoCardBase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIoCardBase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetConnectionStatus() string {
	if o == nil || IsNil(o.ConnectionStatus) {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetConnectionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionStatus) {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasConnectionStatus() bool {
	if o != nil && !IsNil(o.ConnectionStatus) {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *EquipmentIoCardBase) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentIoCardBase) SetDescription(v string) {
	o.Description = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetModuleId() int64 {
	if o == nil || IsNil(o.ModuleId) {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetModuleIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ModuleId) {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasModuleId() bool {
	if o != nil && !IsNil(o.ModuleId) {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EquipmentIoCardBase) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardBase) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardBase) GetOperReasonOk() ([]string, bool) {
	if o == nil || IsNil(o.OperReason) {
		return nil, false
	}
	return o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasOperReason() bool {
	if o != nil && !IsNil(o.OperReason) {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *EquipmentIoCardBase) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetOperState() string {
	if o == nil || IsNil(o.OperState) {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetOperStateOk() (*string, bool) {
	if o == nil || IsNil(o.OperState) {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasOperState() bool {
	if o != nil && !IsNil(o.OperState) {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentIoCardBase) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentIoCardBase) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *EquipmentIoCardBase) SetPid(v string) {
	o.Pid = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *EquipmentIoCardBase) SetProductName(v string) {
	o.ProductName = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EquipmentIoCardBase) SetSku(v string) {
	o.Sku = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EquipmentIoCardBase) SetVersion(v string) {
	o.Version = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetVid() string {
	if o == nil || IsNil(o.Vid) {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetVidOk() (*string, bool) {
	if o == nil || IsNil(o.Vid) {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasVid() bool {
	if o != nil && !IsNil(o.Vid) {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *EquipmentIoCardBase) SetVid(v string) {
	o.Vid = &v
}

// GetHostPorts returns the HostPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardBase) GetHostPorts() []EtherHostPortRelationship {
	if o == nil {
		var ret []EtherHostPortRelationship
		return ret
	}
	return o.HostPorts
}

// GetHostPortsOk returns a tuple with the HostPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardBase) GetHostPortsOk() ([]EtherHostPortRelationship, bool) {
	if o == nil || IsNil(o.HostPorts) {
		return nil, false
	}
	return o.HostPorts, true
}

// HasHostPorts returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasHostPorts() bool {
	if o != nil && !IsNil(o.HostPorts) {
		return true
	}

	return false
}

// SetHostPorts gets a reference to the given []EtherHostPortRelationship and assigns it to the HostPorts field.
func (o *EquipmentIoCardBase) SetHostPorts(v []EtherHostPortRelationship) {
	o.HostPorts = v
}

// GetMgmtController returns the MgmtController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardBase) GetMgmtController() ManagementControllerRelationship {
	if o == nil || IsNil(o.MgmtController.Get()) {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.MgmtController.Get()
}

// GetMgmtControllerOk returns a tuple with the MgmtController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardBase) GetMgmtControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.MgmtController.Get(), o.MgmtController.IsSet()
}

// HasMgmtController returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasMgmtController() bool {
	if o != nil && o.MgmtController.IsSet() {
		return true
	}

	return false
}

// SetMgmtController gets a reference to the given NullableManagementControllerRelationship and assigns it to the MgmtController field.
func (o *EquipmentIoCardBase) SetMgmtController(v ManagementControllerRelationship) {
	o.MgmtController.Set(&v)
}

// SetMgmtControllerNil sets the value for MgmtController to be an explicit nil
func (o *EquipmentIoCardBase) SetMgmtControllerNil() {
	o.MgmtController.Set(nil)
}

// UnsetMgmtController ensures that no value is present for MgmtController, not even an explicit nil
func (o *EquipmentIoCardBase) UnsetMgmtController() {
	o.MgmtController.Unset()
}

// GetNetworkPorts returns the NetworkPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardBase) GetNetworkPorts() []EtherNetworkPortRelationship {
	if o == nil {
		var ret []EtherNetworkPortRelationship
		return ret
	}
	return o.NetworkPorts
}

// GetNetworkPortsOk returns a tuple with the NetworkPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardBase) GetNetworkPortsOk() ([]EtherNetworkPortRelationship, bool) {
	if o == nil || IsNil(o.NetworkPorts) {
		return nil, false
	}
	return o.NetworkPorts, true
}

// HasNetworkPorts returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasNetworkPorts() bool {
	if o != nil && !IsNil(o.NetworkPorts) {
		return true
	}

	return false
}

// SetNetworkPorts gets a reference to the given []EtherNetworkPortRelationship and assigns it to the NetworkPorts field.
func (o *EquipmentIoCardBase) SetNetworkPorts(v []EtherNetworkPortRelationship) {
	o.NetworkPorts = v
}

// GetSubGroup returns the SubGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardBase) GetSubGroup() []PortSubGroupRelationship {
	if o == nil {
		var ret []PortSubGroupRelationship
		return ret
	}
	return o.SubGroup
}

// GetSubGroupOk returns a tuple with the SubGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardBase) GetSubGroupOk() ([]PortSubGroupRelationship, bool) {
	if o == nil || IsNil(o.SubGroup) {
		return nil, false
	}
	return o.SubGroup, true
}

// HasSubGroup returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasSubGroup() bool {
	if o != nil && !IsNil(o.SubGroup) {
		return true
	}

	return false
}

// SetSubGroup gets a reference to the given []PortSubGroupRelationship and assigns it to the SubGroup field.
func (o *EquipmentIoCardBase) SetSubGroup(v []PortSubGroupRelationship) {
	o.SubGroup = v
}

func (o EquipmentIoCardBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentIoCardBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ConnectionStatus) {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.ModuleId) {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if !IsNil(o.OperState) {
		toSerialize["OperState"] = o.OperState
	}
	if !IsNil(o.PartNumber) {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if !IsNil(o.Pid) {
		toSerialize["Pid"] = o.Pid
	}
	if !IsNil(o.ProductName) {
		toSerialize["ProductName"] = o.ProductName
	}
	if !IsNil(o.Sku) {
		toSerialize["Sku"] = o.Sku
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !IsNil(o.Vid) {
		toSerialize["Vid"] = o.Vid
	}
	if o.HostPorts != nil {
		toSerialize["HostPorts"] = o.HostPorts
	}
	if o.MgmtController.IsSet() {
		toSerialize["MgmtController"] = o.MgmtController.Get()
	}
	if o.NetworkPorts != nil {
		toSerialize["NetworkPorts"] = o.NetworkPorts
	}
	if o.SubGroup != nil {
		toSerialize["SubGroup"] = o.SubGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentIoCardBase) UnmarshalJSON(data []byte) (err error) {
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
	type EquipmentIoCardBaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Connectivity Status of FEX/IOM to Switch - A or B or AB.
		ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
		// This field is to provide description for the iocard module model.
		Description *string `json:"Description,omitempty"`
		// Module Identifier for the IO module.
		ModuleId   *int64   `json:"ModuleId,omitempty"`
		OperReason []string `json:"OperReason,omitempty"`
		// Operational state of IO card or fabric extender.
		OperState *string `json:"OperState,omitempty"`
		// Part Number identifier for the IO module.
		PartNumber *string `json:"PartNumber,omitempty"`
		// This field identifies the Product ID for the IO module.
		Pid *string `json:"Pid,omitempty"`
		// This field identifies the Product Name for the iocard module model.
		ProductName *string `json:"ProductName,omitempty"`
		// This field identifies the Stock Keeping Unit for the IO card module.
		Sku *string `json:"Sku,omitempty"`
		// This field identifies the version of the IO card module.
		Version *string `json:"Version,omitempty"`
		// This field identifies the Vendor ID for the IO card module.
		Vid *string `json:"Vid,omitempty"`
		// An array of relationships to etherHostPort resources.
		HostPorts      []EtherHostPortRelationship              `json:"HostPorts,omitempty"`
		MgmtController NullableManagementControllerRelationship `json:"MgmtController,omitempty"`
		// An array of relationships to etherNetworkPort resources.
		NetworkPorts []EtherNetworkPortRelationship `json:"NetworkPorts,omitempty"`
		// An array of relationships to portSubGroup resources.
		SubGroup []PortSubGroupRelationship `json:"SubGroup,omitempty"`
	}

	varEquipmentIoCardBaseWithoutEmbeddedStruct := EquipmentIoCardBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentIoCardBaseWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentIoCardBase := _EquipmentIoCardBase{}
		varEquipmentIoCardBase.ClassId = varEquipmentIoCardBaseWithoutEmbeddedStruct.ClassId
		varEquipmentIoCardBase.ObjectType = varEquipmentIoCardBaseWithoutEmbeddedStruct.ObjectType
		varEquipmentIoCardBase.ConnectionStatus = varEquipmentIoCardBaseWithoutEmbeddedStruct.ConnectionStatus
		varEquipmentIoCardBase.Description = varEquipmentIoCardBaseWithoutEmbeddedStruct.Description
		varEquipmentIoCardBase.ModuleId = varEquipmentIoCardBaseWithoutEmbeddedStruct.ModuleId
		varEquipmentIoCardBase.OperReason = varEquipmentIoCardBaseWithoutEmbeddedStruct.OperReason
		varEquipmentIoCardBase.OperState = varEquipmentIoCardBaseWithoutEmbeddedStruct.OperState
		varEquipmentIoCardBase.PartNumber = varEquipmentIoCardBaseWithoutEmbeddedStruct.PartNumber
		varEquipmentIoCardBase.Pid = varEquipmentIoCardBaseWithoutEmbeddedStruct.Pid
		varEquipmentIoCardBase.ProductName = varEquipmentIoCardBaseWithoutEmbeddedStruct.ProductName
		varEquipmentIoCardBase.Sku = varEquipmentIoCardBaseWithoutEmbeddedStruct.Sku
		varEquipmentIoCardBase.Version = varEquipmentIoCardBaseWithoutEmbeddedStruct.Version
		varEquipmentIoCardBase.Vid = varEquipmentIoCardBaseWithoutEmbeddedStruct.Vid
		varEquipmentIoCardBase.HostPorts = varEquipmentIoCardBaseWithoutEmbeddedStruct.HostPorts
		varEquipmentIoCardBase.MgmtController = varEquipmentIoCardBaseWithoutEmbeddedStruct.MgmtController
		varEquipmentIoCardBase.NetworkPorts = varEquipmentIoCardBaseWithoutEmbeddedStruct.NetworkPorts
		varEquipmentIoCardBase.SubGroup = varEquipmentIoCardBaseWithoutEmbeddedStruct.SubGroup
		*o = EquipmentIoCardBase(varEquipmentIoCardBase)
	} else {
		return err
	}

	varEquipmentIoCardBase := _EquipmentIoCardBase{}

	err = json.Unmarshal(data, &varEquipmentIoCardBase)
	if err == nil {
		o.EquipmentBase = varEquipmentIoCardBase.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "ModuleId")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "ProductName")
		delete(additionalProperties, "Sku")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "HostPorts")
		delete(additionalProperties, "MgmtController")
		delete(additionalProperties, "NetworkPorts")
		delete(additionalProperties, "SubGroup")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableEquipmentIoCardBase struct {
	value *EquipmentIoCardBase
	isSet bool
}

func (v NullableEquipmentIoCardBase) Get() *EquipmentIoCardBase {
	return v.value
}

func (v *NullableEquipmentIoCardBase) Set(val *EquipmentIoCardBase) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCardBase) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCardBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCardBase(val *EquipmentIoCardBase) *NullableEquipmentIoCardBase {
	return &NullableEquipmentIoCardBase{value: val, isSet: true}
}

func (v NullableEquipmentIoCardBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCardBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
