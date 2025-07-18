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
	"time"
)

// checks if the TechsupportmanagementTechSupportStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TechsupportmanagementTechSupportStatus{}

// TechsupportmanagementTechSupportStatus The techsupport collection status.
type TechsupportmanagementTechSupportStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the Techsupport bundle file.
	FileName *string `json:"FileName,omitempty"`
	// Techsupport file size in bytes.
	FileSize *int64 `json:"FileSize,omitempty"`
	// Reason for techsupport failure, if any.
	Reason *string `json:"Reason,omitempty"`
	// Reason for status relay failure, if any.
	RelayReason *string `json:"RelayReason,omitempty"`
	// Status of techsupport status relay. Valid values are NoRelay, Pending, Completed, and Failed.
	RelayStatus *string `json:"RelayStatus,omitempty"`
	// The time at which the techsupport request was initiated.
	RequestTs *time.Time `json:"RequestTs,omitempty"`
	// Status of the techsupport collection. Valid values are Scheduled, Pending, CollectionInProgress, CollectionFailed, CollectionComplete, UploadPending, UploadInProgress, UploadPartsComplete, UploadPreparingNextFile, UploadFailed, TechsupportDownloadUrlCreationFailed, PartiallyCompleted, and Completed. The final status will be one of CollectionFailed, UploadFailed, or TechsupportDownloadUrlCreationFailed if there is a failure, Completed if the request completed successfully and the file (or files) were uploaded to Intersight Storage Service, or PartiallyCompleted if at least one file in a multiple file collection uploaded successfully. All the remaining status values indicates the progress of techsupport collection.
	Status *string `json:"Status,omitempty"`
	// The Url to download the techsupport file.
	TechsupportDownloadUrl *string                                    `json:"TechsupportDownloadUrl,omitempty"`
	TechsupportFiles       []TechsupportmanagementTechSupportFileInfo `json:"TechsupportFiles,omitempty"`
	// The name of the role granted to the user that issued the techsupport request.
	UserRole             *string                                                    `json:"UserRole,omitempty"`
	ClusterMember        NullableAssetClusterMemberRelationship                     `json:"ClusterMember,omitempty"`
	DeviceRegistration   NullableAssetDeviceRegistrationRelationship                `json:"DeviceRegistration,omitempty"`
	OriginResource       NullableMoBaseMoRelationship                               `json:"OriginResource,omitempty"`
	TechSupportRequest   NullableTechsupportmanagementTechSupportBundleRelationship `json:"TechSupportRequest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TechsupportmanagementTechSupportStatus TechsupportmanagementTechSupportStatus

// NewTechsupportmanagementTechSupportStatus instantiates a new TechsupportmanagementTechSupportStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementTechSupportStatus(classId string, objectType string) *TechsupportmanagementTechSupportStatus {
	this := TechsupportmanagementTechSupportStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTechsupportmanagementTechSupportStatusWithDefaults instantiates a new TechsupportmanagementTechSupportStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementTechSupportStatusWithDefaults() *TechsupportmanagementTechSupportStatus {
	this := TechsupportmanagementTechSupportStatus{}
	var classId string = "techsupportmanagement.TechSupportStatus"
	this.ClassId = classId
	var objectType string = "techsupportmanagement.TechSupportStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TechsupportmanagementTechSupportStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TechsupportmanagementTechSupportStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "techsupportmanagement.TechSupportStatus" of the ClassId field.
func (o *TechsupportmanagementTechSupportStatus) GetDefaultClassId() interface{} {
	return "techsupportmanagement.TechSupportStatus"
}

// GetObjectType returns the ObjectType field value
func (o *TechsupportmanagementTechSupportStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TechsupportmanagementTechSupportStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "techsupportmanagement.TechSupportStatus" of the ObjectType field.
func (o *TechsupportmanagementTechSupportStatus) GetDefaultObjectType() interface{} {
	return "techsupportmanagement.TechSupportStatus"
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *TechsupportmanagementTechSupportStatus) SetFileName(v string) {
	o.FileName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetFileSize() int64 {
	if o == nil || IsNil(o.FileSize) {
		var ret int64
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetFileSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int64 and assigns it to the FileSize field.
func (o *TechsupportmanagementTechSupportStatus) SetFileSize(v int64) {
	o.FileSize = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *TechsupportmanagementTechSupportStatus) SetReason(v string) {
	o.Reason = &v
}

// GetRelayReason returns the RelayReason field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetRelayReason() string {
	if o == nil || IsNil(o.RelayReason) {
		var ret string
		return ret
	}
	return *o.RelayReason
}

// GetRelayReasonOk returns a tuple with the RelayReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetRelayReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RelayReason) {
		return nil, false
	}
	return o.RelayReason, true
}

// HasRelayReason returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasRelayReason() bool {
	if o != nil && !IsNil(o.RelayReason) {
		return true
	}

	return false
}

// SetRelayReason gets a reference to the given string and assigns it to the RelayReason field.
func (o *TechsupportmanagementTechSupportStatus) SetRelayReason(v string) {
	o.RelayReason = &v
}

// GetRelayStatus returns the RelayStatus field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetRelayStatus() string {
	if o == nil || IsNil(o.RelayStatus) {
		var ret string
		return ret
	}
	return *o.RelayStatus
}

// GetRelayStatusOk returns a tuple with the RelayStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetRelayStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RelayStatus) {
		return nil, false
	}
	return o.RelayStatus, true
}

// HasRelayStatus returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasRelayStatus() bool {
	if o != nil && !IsNil(o.RelayStatus) {
		return true
	}

	return false
}

// SetRelayStatus gets a reference to the given string and assigns it to the RelayStatus field.
func (o *TechsupportmanagementTechSupportStatus) SetRelayStatus(v string) {
	o.RelayStatus = &v
}

// GetRequestTs returns the RequestTs field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetRequestTs() time.Time {
	if o == nil || IsNil(o.RequestTs) {
		var ret time.Time
		return ret
	}
	return *o.RequestTs
}

// GetRequestTsOk returns a tuple with the RequestTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetRequestTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestTs) {
		return nil, false
	}
	return o.RequestTs, true
}

// HasRequestTs returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasRequestTs() bool {
	if o != nil && !IsNil(o.RequestTs) {
		return true
	}

	return false
}

// SetRequestTs gets a reference to the given time.Time and assigns it to the RequestTs field.
func (o *TechsupportmanagementTechSupportStatus) SetRequestTs(v time.Time) {
	o.RequestTs = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TechsupportmanagementTechSupportStatus) SetStatus(v string) {
	o.Status = &v
}

// GetTechsupportDownloadUrl returns the TechsupportDownloadUrl field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetTechsupportDownloadUrl() string {
	if o == nil || IsNil(o.TechsupportDownloadUrl) {
		var ret string
		return ret
	}
	return *o.TechsupportDownloadUrl
}

// GetTechsupportDownloadUrlOk returns a tuple with the TechsupportDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetTechsupportDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TechsupportDownloadUrl) {
		return nil, false
	}
	return o.TechsupportDownloadUrl, true
}

// HasTechsupportDownloadUrl returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasTechsupportDownloadUrl() bool {
	if o != nil && !IsNil(o.TechsupportDownloadUrl) {
		return true
	}

	return false
}

// SetTechsupportDownloadUrl gets a reference to the given string and assigns it to the TechsupportDownloadUrl field.
func (o *TechsupportmanagementTechSupportStatus) SetTechsupportDownloadUrl(v string) {
	o.TechsupportDownloadUrl = &v
}

// GetTechsupportFiles returns the TechsupportFiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TechsupportmanagementTechSupportStatus) GetTechsupportFiles() []TechsupportmanagementTechSupportFileInfo {
	if o == nil {
		var ret []TechsupportmanagementTechSupportFileInfo
		return ret
	}
	return o.TechsupportFiles
}

// GetTechsupportFilesOk returns a tuple with the TechsupportFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TechsupportmanagementTechSupportStatus) GetTechsupportFilesOk() ([]TechsupportmanagementTechSupportFileInfo, bool) {
	if o == nil || IsNil(o.TechsupportFiles) {
		return nil, false
	}
	return o.TechsupportFiles, true
}

// HasTechsupportFiles returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasTechsupportFiles() bool {
	if o != nil && !IsNil(o.TechsupportFiles) {
		return true
	}

	return false
}

// SetTechsupportFiles gets a reference to the given []TechsupportmanagementTechSupportFileInfo and assigns it to the TechsupportFiles field.
func (o *TechsupportmanagementTechSupportStatus) SetTechsupportFiles(v []TechsupportmanagementTechSupportFileInfo) {
	o.TechsupportFiles = v
}

// GetUserRole returns the UserRole field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetUserRole() string {
	if o == nil || IsNil(o.UserRole) {
		var ret string
		return ret
	}
	return *o.UserRole
}

// GetUserRoleOk returns a tuple with the UserRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetUserRoleOk() (*string, bool) {
	if o == nil || IsNil(o.UserRole) {
		return nil, false
	}
	return o.UserRole, true
}

// HasUserRole returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasUserRole() bool {
	if o != nil && !IsNil(o.UserRole) {
		return true
	}

	return false
}

// SetUserRole gets a reference to the given string and assigns it to the UserRole field.
func (o *TechsupportmanagementTechSupportStatus) SetUserRole(v string) {
	o.UserRole = &v
}

// GetClusterMember returns the ClusterMember field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TechsupportmanagementTechSupportStatus) GetClusterMember() AssetClusterMemberRelationship {
	if o == nil || IsNil(o.ClusterMember.Get()) {
		var ret AssetClusterMemberRelationship
		return ret
	}
	return *o.ClusterMember.Get()
}

// GetClusterMemberOk returns a tuple with the ClusterMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TechsupportmanagementTechSupportStatus) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterMember.Get(), o.ClusterMember.IsSet()
}

// HasClusterMember returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasClusterMember() bool {
	if o != nil && o.ClusterMember.IsSet() {
		return true
	}

	return false
}

// SetClusterMember gets a reference to the given NullableAssetClusterMemberRelationship and assigns it to the ClusterMember field.
func (o *TechsupportmanagementTechSupportStatus) SetClusterMember(v AssetClusterMemberRelationship) {
	o.ClusterMember.Set(&v)
}

// SetClusterMemberNil sets the value for ClusterMember to be an explicit nil
func (o *TechsupportmanagementTechSupportStatus) SetClusterMemberNil() {
	o.ClusterMember.Set(nil)
}

// UnsetClusterMember ensures that no value is present for ClusterMember, not even an explicit nil
func (o *TechsupportmanagementTechSupportStatus) UnsetClusterMember() {
	o.ClusterMember.Unset()
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TechsupportmanagementTechSupportStatus) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.DeviceRegistration.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration.Get()
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TechsupportmanagementTechSupportStatus) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceRegistration.Get(), o.DeviceRegistration.IsSet()
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration.IsSet() {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *TechsupportmanagementTechSupportStatus) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration.Set(&v)
}

// SetDeviceRegistrationNil sets the value for DeviceRegistration to be an explicit nil
func (o *TechsupportmanagementTechSupportStatus) SetDeviceRegistrationNil() {
	o.DeviceRegistration.Set(nil)
}

// UnsetDeviceRegistration ensures that no value is present for DeviceRegistration, not even an explicit nil
func (o *TechsupportmanagementTechSupportStatus) UnsetDeviceRegistration() {
	o.DeviceRegistration.Unset()
}

// GetOriginResource returns the OriginResource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TechsupportmanagementTechSupportStatus) GetOriginResource() MoBaseMoRelationship {
	if o == nil || IsNil(o.OriginResource.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.OriginResource.Get()
}

// GetOriginResourceOk returns a tuple with the OriginResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TechsupportmanagementTechSupportStatus) GetOriginResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginResource.Get(), o.OriginResource.IsSet()
}

// HasOriginResource returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasOriginResource() bool {
	if o != nil && o.OriginResource.IsSet() {
		return true
	}

	return false
}

// SetOriginResource gets a reference to the given NullableMoBaseMoRelationship and assigns it to the OriginResource field.
func (o *TechsupportmanagementTechSupportStatus) SetOriginResource(v MoBaseMoRelationship) {
	o.OriginResource.Set(&v)
}

// SetOriginResourceNil sets the value for OriginResource to be an explicit nil
func (o *TechsupportmanagementTechSupportStatus) SetOriginResourceNil() {
	o.OriginResource.Set(nil)
}

// UnsetOriginResource ensures that no value is present for OriginResource, not even an explicit nil
func (o *TechsupportmanagementTechSupportStatus) UnsetOriginResource() {
	o.OriginResource.Unset()
}

// GetTechSupportRequest returns the TechSupportRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TechsupportmanagementTechSupportStatus) GetTechSupportRequest() TechsupportmanagementTechSupportBundleRelationship {
	if o == nil || IsNil(o.TechSupportRequest.Get()) {
		var ret TechsupportmanagementTechSupportBundleRelationship
		return ret
	}
	return *o.TechSupportRequest.Get()
}

// GetTechSupportRequestOk returns a tuple with the TechSupportRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TechsupportmanagementTechSupportStatus) GetTechSupportRequestOk() (*TechsupportmanagementTechSupportBundleRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TechSupportRequest.Get(), o.TechSupportRequest.IsSet()
}

// HasTechSupportRequest returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasTechSupportRequest() bool {
	if o != nil && o.TechSupportRequest.IsSet() {
		return true
	}

	return false
}

// SetTechSupportRequest gets a reference to the given NullableTechsupportmanagementTechSupportBundleRelationship and assigns it to the TechSupportRequest field.
func (o *TechsupportmanagementTechSupportStatus) SetTechSupportRequest(v TechsupportmanagementTechSupportBundleRelationship) {
	o.TechSupportRequest.Set(&v)
}

// SetTechSupportRequestNil sets the value for TechSupportRequest to be an explicit nil
func (o *TechsupportmanagementTechSupportStatus) SetTechSupportRequestNil() {
	o.TechSupportRequest.Set(nil)
}

// UnsetTechSupportRequest ensures that no value is present for TechSupportRequest, not even an explicit nil
func (o *TechsupportmanagementTechSupportStatus) UnsetTechSupportRequest() {
	o.TechSupportRequest.Unset()
}

func (o TechsupportmanagementTechSupportStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TechsupportmanagementTechSupportStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.FileName) {
		toSerialize["FileName"] = o.FileName
	}
	if !IsNil(o.FileSize) {
		toSerialize["FileSize"] = o.FileSize
	}
	if !IsNil(o.Reason) {
		toSerialize["Reason"] = o.Reason
	}
	if !IsNil(o.RelayReason) {
		toSerialize["RelayReason"] = o.RelayReason
	}
	if !IsNil(o.RelayStatus) {
		toSerialize["RelayStatus"] = o.RelayStatus
	}
	if !IsNil(o.RequestTs) {
		toSerialize["RequestTs"] = o.RequestTs
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.TechsupportDownloadUrl) {
		toSerialize["TechsupportDownloadUrl"] = o.TechsupportDownloadUrl
	}
	if o.TechsupportFiles != nil {
		toSerialize["TechsupportFiles"] = o.TechsupportFiles
	}
	if !IsNil(o.UserRole) {
		toSerialize["UserRole"] = o.UserRole
	}
	if o.ClusterMember.IsSet() {
		toSerialize["ClusterMember"] = o.ClusterMember.Get()
	}
	if o.DeviceRegistration.IsSet() {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration.Get()
	}
	if o.OriginResource.IsSet() {
		toSerialize["OriginResource"] = o.OriginResource.Get()
	}
	if o.TechSupportRequest.IsSet() {
		toSerialize["TechSupportRequest"] = o.TechSupportRequest.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TechsupportmanagementTechSupportStatus) UnmarshalJSON(data []byte) (err error) {
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
	type TechsupportmanagementTechSupportStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the Techsupport bundle file.
		FileName *string `json:"FileName,omitempty"`
		// Techsupport file size in bytes.
		FileSize *int64 `json:"FileSize,omitempty"`
		// Reason for techsupport failure, if any.
		Reason *string `json:"Reason,omitempty"`
		// Reason for status relay failure, if any.
		RelayReason *string `json:"RelayReason,omitempty"`
		// Status of techsupport status relay. Valid values are NoRelay, Pending, Completed, and Failed.
		RelayStatus *string `json:"RelayStatus,omitempty"`
		// The time at which the techsupport request was initiated.
		RequestTs *time.Time `json:"RequestTs,omitempty"`
		// Status of the techsupport collection. Valid values are Scheduled, Pending, CollectionInProgress, CollectionFailed, CollectionComplete, UploadPending, UploadInProgress, UploadPartsComplete, UploadPreparingNextFile, UploadFailed, TechsupportDownloadUrlCreationFailed, PartiallyCompleted, and Completed. The final status will be one of CollectionFailed, UploadFailed, or TechsupportDownloadUrlCreationFailed if there is a failure, Completed if the request completed successfully and the file (or files) were uploaded to Intersight Storage Service, or PartiallyCompleted if at least one file in a multiple file collection uploaded successfully. All the remaining status values indicates the progress of techsupport collection.
		Status *string `json:"Status,omitempty"`
		// The Url to download the techsupport file.
		TechsupportDownloadUrl *string                                    `json:"TechsupportDownloadUrl,omitempty"`
		TechsupportFiles       []TechsupportmanagementTechSupportFileInfo `json:"TechsupportFiles,omitempty"`
		// The name of the role granted to the user that issued the techsupport request.
		UserRole           *string                                                    `json:"UserRole,omitempty"`
		ClusterMember      NullableAssetClusterMemberRelationship                     `json:"ClusterMember,omitempty"`
		DeviceRegistration NullableAssetDeviceRegistrationRelationship                `json:"DeviceRegistration,omitempty"`
		OriginResource     NullableMoBaseMoRelationship                               `json:"OriginResource,omitempty"`
		TechSupportRequest NullableTechsupportmanagementTechSupportBundleRelationship `json:"TechSupportRequest,omitempty"`
	}

	varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct := TechsupportmanagementTechSupportStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct)
	if err == nil {
		varTechsupportmanagementTechSupportStatus := _TechsupportmanagementTechSupportStatus{}
		varTechsupportmanagementTechSupportStatus.ClassId = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.ClassId
		varTechsupportmanagementTechSupportStatus.ObjectType = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.ObjectType
		varTechsupportmanagementTechSupportStatus.FileName = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.FileName
		varTechsupportmanagementTechSupportStatus.FileSize = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.FileSize
		varTechsupportmanagementTechSupportStatus.Reason = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.Reason
		varTechsupportmanagementTechSupportStatus.RelayReason = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.RelayReason
		varTechsupportmanagementTechSupportStatus.RelayStatus = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.RelayStatus
		varTechsupportmanagementTechSupportStatus.RequestTs = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.RequestTs
		varTechsupportmanagementTechSupportStatus.Status = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.Status
		varTechsupportmanagementTechSupportStatus.TechsupportDownloadUrl = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.TechsupportDownloadUrl
		varTechsupportmanagementTechSupportStatus.TechsupportFiles = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.TechsupportFiles
		varTechsupportmanagementTechSupportStatus.UserRole = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.UserRole
		varTechsupportmanagementTechSupportStatus.ClusterMember = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.ClusterMember
		varTechsupportmanagementTechSupportStatus.DeviceRegistration = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.DeviceRegistration
		varTechsupportmanagementTechSupportStatus.OriginResource = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.OriginResource
		varTechsupportmanagementTechSupportStatus.TechSupportRequest = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.TechSupportRequest
		*o = TechsupportmanagementTechSupportStatus(varTechsupportmanagementTechSupportStatus)
	} else {
		return err
	}

	varTechsupportmanagementTechSupportStatus := _TechsupportmanagementTechSupportStatus{}

	err = json.Unmarshal(data, &varTechsupportmanagementTechSupportStatus)
	if err == nil {
		o.MoBaseMo = varTechsupportmanagementTechSupportStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileName")
		delete(additionalProperties, "FileSize")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "RelayReason")
		delete(additionalProperties, "RelayStatus")
		delete(additionalProperties, "RequestTs")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TechsupportDownloadUrl")
		delete(additionalProperties, "TechsupportFiles")
		delete(additionalProperties, "UserRole")
		delete(additionalProperties, "ClusterMember")
		delete(additionalProperties, "DeviceRegistration")
		delete(additionalProperties, "OriginResource")
		delete(additionalProperties, "TechSupportRequest")

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

type NullableTechsupportmanagementTechSupportStatus struct {
	value *TechsupportmanagementTechSupportStatus
	isSet bool
}

func (v NullableTechsupportmanagementTechSupportStatus) Get() *TechsupportmanagementTechSupportStatus {
	return v.value
}

func (v *NullableTechsupportmanagementTechSupportStatus) Set(val *TechsupportmanagementTechSupportStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementTechSupportStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementTechSupportStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementTechSupportStatus(val *TechsupportmanagementTechSupportStatus) *NullableTechsupportmanagementTechSupportStatus {
	return &NullableTechsupportmanagementTechSupportStatus{value: val, isSet: true}
}

func (v NullableTechsupportmanagementTechSupportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementTechSupportStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
