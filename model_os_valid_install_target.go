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

// checks if the OsValidInstallTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsValidInstallTarget{}

// OsValidInstallTarget ValidInstallTarget is used to fetch all the valid Install targets for the servers. The List of Install targets includes Physical Disks and Virtual Drives.
type OsValidInstallTarget struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                   `json:"ObjectType"`
	E3sNvme    []OsPhysicalDiskResponse `json:"E3sNvme,omitempty"`
	// Error message if any errors are encountered while fetching and validating Install targets for the server.
	Error                   *string                   `json:"Error,omitempty"`
	InstallTargets          []OsInstallTargetResponse `json:"InstallTargets,omitempty"`
	M2Jbod                  []OsPhysicalDiskResponse  `json:"M2Jbod,omitempty"`
	M2NvmeRaidJbod          []OsPhysicalDiskResponse  `json:"M2NvmeRaidJbod,omitempty"`
	M2NvmeRaidVirtualDrives []OsVirtualDriveResponse  `json:"M2NvmeRaidVirtualDrives,omitempty"`
	M2VirtualDrives         []OsVirtualDriveResponse  `json:"M2VirtualDrives,omitempty"`
	MraidJbod               []OsPhysicalDiskResponse  `json:"MraidJbod,omitempty"`
	MraidVirtualDrives      []OsVirtualDriveResponse  `json:"MraidVirtualDrives,omitempty"`
	MstorNvme               []OsPhysicalDiskResponse  `json:"MstorNvme,omitempty"`
	// Flag to denote the source of the request. If the call is from Orchestration UI, only the flat list of Install targets can be sent as response.
	Src    *string                  `json:"Src,omitempty"`
	U2Nvme []OsPhysicalDiskResponse `json:"U2Nvme,omitempty"`
	// An array of relationships to computePhysical resources.
	Servers              []ComputePhysicalRelationship `json:"Servers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsValidInstallTarget OsValidInstallTarget

// NewOsValidInstallTarget instantiates a new OsValidInstallTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsValidInstallTarget(classId string, objectType string) *OsValidInstallTarget {
	this := OsValidInstallTarget{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsValidInstallTargetWithDefaults instantiates a new OsValidInstallTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsValidInstallTargetWithDefaults() *OsValidInstallTarget {
	this := OsValidInstallTarget{}
	var classId string = "os.ValidInstallTarget"
	this.ClassId = classId
	var objectType string = "os.ValidInstallTarget"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsValidInstallTarget) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsValidInstallTarget) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsValidInstallTarget) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "os.ValidInstallTarget" of the ClassId field.
func (o *OsValidInstallTarget) GetDefaultClassId() interface{} {
	return "os.ValidInstallTarget"
}

// GetObjectType returns the ObjectType field value
func (o *OsValidInstallTarget) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsValidInstallTarget) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsValidInstallTarget) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "os.ValidInstallTarget" of the ObjectType field.
func (o *OsValidInstallTarget) GetDefaultObjectType() interface{} {
	return "os.ValidInstallTarget"
}

// GetE3sNvme returns the E3sNvme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetE3sNvme() []OsPhysicalDiskResponse {
	if o == nil {
		var ret []OsPhysicalDiskResponse
		return ret
	}
	return o.E3sNvme
}

// GetE3sNvmeOk returns a tuple with the E3sNvme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetE3sNvmeOk() ([]OsPhysicalDiskResponse, bool) {
	if o == nil || IsNil(o.E3sNvme) {
		return nil, false
	}
	return o.E3sNvme, true
}

// HasE3sNvme returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasE3sNvme() bool {
	if o != nil && !IsNil(o.E3sNvme) {
		return true
	}

	return false
}

// SetE3sNvme gets a reference to the given []OsPhysicalDiskResponse and assigns it to the E3sNvme field.
func (o *OsValidInstallTarget) SetE3sNvme(v []OsPhysicalDiskResponse) {
	o.E3sNvme = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *OsValidInstallTarget) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsValidInstallTarget) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *OsValidInstallTarget) SetError(v string) {
	o.Error = &v
}

// GetInstallTargets returns the InstallTargets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetInstallTargets() []OsInstallTargetResponse {
	if o == nil {
		var ret []OsInstallTargetResponse
		return ret
	}
	return o.InstallTargets
}

// GetInstallTargetsOk returns a tuple with the InstallTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetInstallTargetsOk() ([]OsInstallTargetResponse, bool) {
	if o == nil || IsNil(o.InstallTargets) {
		return nil, false
	}
	return o.InstallTargets, true
}

// HasInstallTargets returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasInstallTargets() bool {
	if o != nil && !IsNil(o.InstallTargets) {
		return true
	}

	return false
}

// SetInstallTargets gets a reference to the given []OsInstallTargetResponse and assigns it to the InstallTargets field.
func (o *OsValidInstallTarget) SetInstallTargets(v []OsInstallTargetResponse) {
	o.InstallTargets = v
}

// GetM2Jbod returns the M2Jbod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetM2Jbod() []OsPhysicalDiskResponse {
	if o == nil {
		var ret []OsPhysicalDiskResponse
		return ret
	}
	return o.M2Jbod
}

// GetM2JbodOk returns a tuple with the M2Jbod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetM2JbodOk() ([]OsPhysicalDiskResponse, bool) {
	if o == nil || IsNil(o.M2Jbod) {
		return nil, false
	}
	return o.M2Jbod, true
}

// HasM2Jbod returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasM2Jbod() bool {
	if o != nil && !IsNil(o.M2Jbod) {
		return true
	}

	return false
}

// SetM2Jbod gets a reference to the given []OsPhysicalDiskResponse and assigns it to the M2Jbod field.
func (o *OsValidInstallTarget) SetM2Jbod(v []OsPhysicalDiskResponse) {
	o.M2Jbod = v
}

// GetM2NvmeRaidJbod returns the M2NvmeRaidJbod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetM2NvmeRaidJbod() []OsPhysicalDiskResponse {
	if o == nil {
		var ret []OsPhysicalDiskResponse
		return ret
	}
	return o.M2NvmeRaidJbod
}

// GetM2NvmeRaidJbodOk returns a tuple with the M2NvmeRaidJbod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetM2NvmeRaidJbodOk() ([]OsPhysicalDiskResponse, bool) {
	if o == nil || IsNil(o.M2NvmeRaidJbod) {
		return nil, false
	}
	return o.M2NvmeRaidJbod, true
}

// HasM2NvmeRaidJbod returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasM2NvmeRaidJbod() bool {
	if o != nil && !IsNil(o.M2NvmeRaidJbod) {
		return true
	}

	return false
}

// SetM2NvmeRaidJbod gets a reference to the given []OsPhysicalDiskResponse and assigns it to the M2NvmeRaidJbod field.
func (o *OsValidInstallTarget) SetM2NvmeRaidJbod(v []OsPhysicalDiskResponse) {
	o.M2NvmeRaidJbod = v
}

// GetM2NvmeRaidVirtualDrives returns the M2NvmeRaidVirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetM2NvmeRaidVirtualDrives() []OsVirtualDriveResponse {
	if o == nil {
		var ret []OsVirtualDriveResponse
		return ret
	}
	return o.M2NvmeRaidVirtualDrives
}

// GetM2NvmeRaidVirtualDrivesOk returns a tuple with the M2NvmeRaidVirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetM2NvmeRaidVirtualDrivesOk() ([]OsVirtualDriveResponse, bool) {
	if o == nil || IsNil(o.M2NvmeRaidVirtualDrives) {
		return nil, false
	}
	return o.M2NvmeRaidVirtualDrives, true
}

// HasM2NvmeRaidVirtualDrives returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasM2NvmeRaidVirtualDrives() bool {
	if o != nil && !IsNil(o.M2NvmeRaidVirtualDrives) {
		return true
	}

	return false
}

// SetM2NvmeRaidVirtualDrives gets a reference to the given []OsVirtualDriveResponse and assigns it to the M2NvmeRaidVirtualDrives field.
func (o *OsValidInstallTarget) SetM2NvmeRaidVirtualDrives(v []OsVirtualDriveResponse) {
	o.M2NvmeRaidVirtualDrives = v
}

// GetM2VirtualDrives returns the M2VirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetM2VirtualDrives() []OsVirtualDriveResponse {
	if o == nil {
		var ret []OsVirtualDriveResponse
		return ret
	}
	return o.M2VirtualDrives
}

// GetM2VirtualDrivesOk returns a tuple with the M2VirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetM2VirtualDrivesOk() ([]OsVirtualDriveResponse, bool) {
	if o == nil || IsNil(o.M2VirtualDrives) {
		return nil, false
	}
	return o.M2VirtualDrives, true
}

// HasM2VirtualDrives returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasM2VirtualDrives() bool {
	if o != nil && !IsNil(o.M2VirtualDrives) {
		return true
	}

	return false
}

// SetM2VirtualDrives gets a reference to the given []OsVirtualDriveResponse and assigns it to the M2VirtualDrives field.
func (o *OsValidInstallTarget) SetM2VirtualDrives(v []OsVirtualDriveResponse) {
	o.M2VirtualDrives = v
}

// GetMraidJbod returns the MraidJbod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetMraidJbod() []OsPhysicalDiskResponse {
	if o == nil {
		var ret []OsPhysicalDiskResponse
		return ret
	}
	return o.MraidJbod
}

// GetMraidJbodOk returns a tuple with the MraidJbod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetMraidJbodOk() ([]OsPhysicalDiskResponse, bool) {
	if o == nil || IsNil(o.MraidJbod) {
		return nil, false
	}
	return o.MraidJbod, true
}

// HasMraidJbod returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasMraidJbod() bool {
	if o != nil && !IsNil(o.MraidJbod) {
		return true
	}

	return false
}

// SetMraidJbod gets a reference to the given []OsPhysicalDiskResponse and assigns it to the MraidJbod field.
func (o *OsValidInstallTarget) SetMraidJbod(v []OsPhysicalDiskResponse) {
	o.MraidJbod = v
}

// GetMraidVirtualDrives returns the MraidVirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetMraidVirtualDrives() []OsVirtualDriveResponse {
	if o == nil {
		var ret []OsVirtualDriveResponse
		return ret
	}
	return o.MraidVirtualDrives
}

// GetMraidVirtualDrivesOk returns a tuple with the MraidVirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetMraidVirtualDrivesOk() ([]OsVirtualDriveResponse, bool) {
	if o == nil || IsNil(o.MraidVirtualDrives) {
		return nil, false
	}
	return o.MraidVirtualDrives, true
}

// HasMraidVirtualDrives returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasMraidVirtualDrives() bool {
	if o != nil && !IsNil(o.MraidVirtualDrives) {
		return true
	}

	return false
}

// SetMraidVirtualDrives gets a reference to the given []OsVirtualDriveResponse and assigns it to the MraidVirtualDrives field.
func (o *OsValidInstallTarget) SetMraidVirtualDrives(v []OsVirtualDriveResponse) {
	o.MraidVirtualDrives = v
}

// GetMstorNvme returns the MstorNvme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetMstorNvme() []OsPhysicalDiskResponse {
	if o == nil {
		var ret []OsPhysicalDiskResponse
		return ret
	}
	return o.MstorNvme
}

// GetMstorNvmeOk returns a tuple with the MstorNvme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetMstorNvmeOk() ([]OsPhysicalDiskResponse, bool) {
	if o == nil || IsNil(o.MstorNvme) {
		return nil, false
	}
	return o.MstorNvme, true
}

// HasMstorNvme returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasMstorNvme() bool {
	if o != nil && !IsNil(o.MstorNvme) {
		return true
	}

	return false
}

// SetMstorNvme gets a reference to the given []OsPhysicalDiskResponse and assigns it to the MstorNvme field.
func (o *OsValidInstallTarget) SetMstorNvme(v []OsPhysicalDiskResponse) {
	o.MstorNvme = v
}

// GetSrc returns the Src field value if set, zero value otherwise.
func (o *OsValidInstallTarget) GetSrc() string {
	if o == nil || IsNil(o.Src) {
		var ret string
		return ret
	}
	return *o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsValidInstallTarget) GetSrcOk() (*string, bool) {
	if o == nil || IsNil(o.Src) {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasSrc() bool {
	if o != nil && !IsNil(o.Src) {
		return true
	}

	return false
}

// SetSrc gets a reference to the given string and assigns it to the Src field.
func (o *OsValidInstallTarget) SetSrc(v string) {
	o.Src = &v
}

// GetU2Nvme returns the U2Nvme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetU2Nvme() []OsPhysicalDiskResponse {
	if o == nil {
		var ret []OsPhysicalDiskResponse
		return ret
	}
	return o.U2Nvme
}

// GetU2NvmeOk returns a tuple with the U2Nvme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetU2NvmeOk() ([]OsPhysicalDiskResponse, bool) {
	if o == nil || IsNil(o.U2Nvme) {
		return nil, false
	}
	return o.U2Nvme, true
}

// HasU2Nvme returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasU2Nvme() bool {
	if o != nil && !IsNil(o.U2Nvme) {
		return true
	}

	return false
}

// SetU2Nvme gets a reference to the given []OsPhysicalDiskResponse and assigns it to the U2Nvme field.
func (o *OsValidInstallTarget) SetU2Nvme(v []OsPhysicalDiskResponse) {
	o.U2Nvme = v
}

// GetServers returns the Servers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetServers() []ComputePhysicalRelationship {
	if o == nil {
		var ret []ComputePhysicalRelationship
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetServersOk() ([]ComputePhysicalRelationship, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []ComputePhysicalRelationship and assigns it to the Servers field.
func (o *OsValidInstallTarget) SetServers(v []ComputePhysicalRelationship) {
	o.Servers = v
}

func (o OsValidInstallTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsValidInstallTarget) ToMap() (map[string]interface{}, error) {
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
	if o.E3sNvme != nil {
		toSerialize["E3sNvme"] = o.E3sNvme
	}
	if !IsNil(o.Error) {
		toSerialize["Error"] = o.Error
	}
	if o.InstallTargets != nil {
		toSerialize["InstallTargets"] = o.InstallTargets
	}
	if o.M2Jbod != nil {
		toSerialize["M2Jbod"] = o.M2Jbod
	}
	if o.M2NvmeRaidJbod != nil {
		toSerialize["M2NvmeRaidJbod"] = o.M2NvmeRaidJbod
	}
	if o.M2NvmeRaidVirtualDrives != nil {
		toSerialize["M2NvmeRaidVirtualDrives"] = o.M2NvmeRaidVirtualDrives
	}
	if o.M2VirtualDrives != nil {
		toSerialize["M2VirtualDrives"] = o.M2VirtualDrives
	}
	if o.MraidJbod != nil {
		toSerialize["MraidJbod"] = o.MraidJbod
	}
	if o.MraidVirtualDrives != nil {
		toSerialize["MraidVirtualDrives"] = o.MraidVirtualDrives
	}
	if o.MstorNvme != nil {
		toSerialize["MstorNvme"] = o.MstorNvme
	}
	if !IsNil(o.Src) {
		toSerialize["Src"] = o.Src
	}
	if o.U2Nvme != nil {
		toSerialize["U2Nvme"] = o.U2Nvme
	}
	if o.Servers != nil {
		toSerialize["Servers"] = o.Servers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OsValidInstallTarget) UnmarshalJSON(data []byte) (err error) {
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
	type OsValidInstallTargetWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                   `json:"ObjectType"`
		E3sNvme    []OsPhysicalDiskResponse `json:"E3sNvme,omitempty"`
		// Error message if any errors are encountered while fetching and validating Install targets for the server.
		Error                   *string                   `json:"Error,omitempty"`
		InstallTargets          []OsInstallTargetResponse `json:"InstallTargets,omitempty"`
		M2Jbod                  []OsPhysicalDiskResponse  `json:"M2Jbod,omitempty"`
		M2NvmeRaidJbod          []OsPhysicalDiskResponse  `json:"M2NvmeRaidJbod,omitempty"`
		M2NvmeRaidVirtualDrives []OsVirtualDriveResponse  `json:"M2NvmeRaidVirtualDrives,omitempty"`
		M2VirtualDrives         []OsVirtualDriveResponse  `json:"M2VirtualDrives,omitempty"`
		MraidJbod               []OsPhysicalDiskResponse  `json:"MraidJbod,omitempty"`
		MraidVirtualDrives      []OsVirtualDriveResponse  `json:"MraidVirtualDrives,omitempty"`
		MstorNvme               []OsPhysicalDiskResponse  `json:"MstorNvme,omitempty"`
		// Flag to denote the source of the request. If the call is from Orchestration UI, only the flat list of Install targets can be sent as response.
		Src    *string                  `json:"Src,omitempty"`
		U2Nvme []OsPhysicalDiskResponse `json:"U2Nvme,omitempty"`
		// An array of relationships to computePhysical resources.
		Servers []ComputePhysicalRelationship `json:"Servers,omitempty"`
	}

	varOsValidInstallTargetWithoutEmbeddedStruct := OsValidInstallTargetWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOsValidInstallTargetWithoutEmbeddedStruct)
	if err == nil {
		varOsValidInstallTarget := _OsValidInstallTarget{}
		varOsValidInstallTarget.ClassId = varOsValidInstallTargetWithoutEmbeddedStruct.ClassId
		varOsValidInstallTarget.ObjectType = varOsValidInstallTargetWithoutEmbeddedStruct.ObjectType
		varOsValidInstallTarget.E3sNvme = varOsValidInstallTargetWithoutEmbeddedStruct.E3sNvme
		varOsValidInstallTarget.Error = varOsValidInstallTargetWithoutEmbeddedStruct.Error
		varOsValidInstallTarget.InstallTargets = varOsValidInstallTargetWithoutEmbeddedStruct.InstallTargets
		varOsValidInstallTarget.M2Jbod = varOsValidInstallTargetWithoutEmbeddedStruct.M2Jbod
		varOsValidInstallTarget.M2NvmeRaidJbod = varOsValidInstallTargetWithoutEmbeddedStruct.M2NvmeRaidJbod
		varOsValidInstallTarget.M2NvmeRaidVirtualDrives = varOsValidInstallTargetWithoutEmbeddedStruct.M2NvmeRaidVirtualDrives
		varOsValidInstallTarget.M2VirtualDrives = varOsValidInstallTargetWithoutEmbeddedStruct.M2VirtualDrives
		varOsValidInstallTarget.MraidJbod = varOsValidInstallTargetWithoutEmbeddedStruct.MraidJbod
		varOsValidInstallTarget.MraidVirtualDrives = varOsValidInstallTargetWithoutEmbeddedStruct.MraidVirtualDrives
		varOsValidInstallTarget.MstorNvme = varOsValidInstallTargetWithoutEmbeddedStruct.MstorNvme
		varOsValidInstallTarget.Src = varOsValidInstallTargetWithoutEmbeddedStruct.Src
		varOsValidInstallTarget.U2Nvme = varOsValidInstallTargetWithoutEmbeddedStruct.U2Nvme
		varOsValidInstallTarget.Servers = varOsValidInstallTargetWithoutEmbeddedStruct.Servers
		*o = OsValidInstallTarget(varOsValidInstallTarget)
	} else {
		return err
	}

	varOsValidInstallTarget := _OsValidInstallTarget{}

	err = json.Unmarshal(data, &varOsValidInstallTarget)
	if err == nil {
		o.MoBaseMo = varOsValidInstallTarget.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "E3sNvme")
		delete(additionalProperties, "Error")
		delete(additionalProperties, "InstallTargets")
		delete(additionalProperties, "M2Jbod")
		delete(additionalProperties, "M2NvmeRaidJbod")
		delete(additionalProperties, "M2NvmeRaidVirtualDrives")
		delete(additionalProperties, "M2VirtualDrives")
		delete(additionalProperties, "MraidJbod")
		delete(additionalProperties, "MraidVirtualDrives")
		delete(additionalProperties, "MstorNvme")
		delete(additionalProperties, "Src")
		delete(additionalProperties, "U2Nvme")
		delete(additionalProperties, "Servers")

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

type NullableOsValidInstallTarget struct {
	value *OsValidInstallTarget
	isSet bool
}

func (v NullableOsValidInstallTarget) Get() *OsValidInstallTarget {
	return v.value
}

func (v *NullableOsValidInstallTarget) Set(val *OsValidInstallTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableOsValidInstallTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableOsValidInstallTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsValidInstallTarget(val *OsValidInstallTarget) *NullableOsValidInstallTarget {
	return &NullableOsValidInstallTarget{value: val, isSet: true}
}

func (v NullableOsValidInstallTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsValidInstallTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
