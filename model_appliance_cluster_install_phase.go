/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9235
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

// ApplianceClusterInstallPhase Represents the Intersight Appliance cluster installation life cycle.
type ApplianceClusterInstallPhase struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Node number install is running on.
	CurrentNode *int64 `json:"CurrentNode,omitempty"`
	// Subphase of the phase currently running.
	CurrentSubphase *int64 `json:"CurrentSubphase,omitempty"`
	// Elapsed time in seconds to complete the current install phase.
	ElapsedTime *int64 `json:"ElapsedTime,omitempty"`
	// End date of the software install phase.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Indicates if the install phase has failed or not.
	Failed *bool `json:"Failed,omitempty"`
	// Status message set during the install phase.
	Message *string `json:"Message,omitempty"`
	// Name of the install phase. * `Backup` - Backup the currently running node. * `EnableBootstrap` - Disable echo and enable nginx on the currently running node. * `UpdateAnsibleHosts` - Update /etc/ansible/hosts on each node. * `SyncImages` - Sync images and manifest to each node. * `ConfigureEtcd` - Configure etcd on each node. * `DeployServices` - Deploy kubernetes services from node 1. * `InstallEquinox` - Configure and install equinox on each node. * `Validate` - Validate equinox is running in primary/secondary mode. * `CheckApplianceClusterState` - Check the appliance cluster status. * `Success` - Upgrade completed successfully. * `Fail` - Indicates that the upgrade process has failed. * `Cancel` - Indicates that the upgrade was canceled by the Intersight Appliance.
	Name         *string `json:"Name,omitempty"`
	PendingNodes []int64 `json:"PendingNodes,omitempty"`
	// Start date of the software install phase.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// Status of the install phase.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceClusterInstallPhase ApplianceClusterInstallPhase

// NewApplianceClusterInstallPhase instantiates a new ApplianceClusterInstallPhase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceClusterInstallPhase(classId string, objectType string) *ApplianceClusterInstallPhase {
	this := ApplianceClusterInstallPhase{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceClusterInstallPhaseWithDefaults instantiates a new ApplianceClusterInstallPhase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceClusterInstallPhaseWithDefaults() *ApplianceClusterInstallPhase {
	this := ApplianceClusterInstallPhase{}
	var classId string = "appliance.ClusterInstallPhase"
	this.ClassId = classId
	var objectType string = "appliance.ClusterInstallPhase"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceClusterInstallPhase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallPhase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceClusterInstallPhase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceClusterInstallPhase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallPhase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceClusterInstallPhase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCurrentNode returns the CurrentNode field value if set, zero value otherwise.
func (o *ApplianceClusterInstallPhase) GetCurrentNode() int64 {
	if o == nil || o.CurrentNode == nil {
		var ret int64
		return ret
	}
	return *o.CurrentNode
}

// GetCurrentNodeOk returns a tuple with the CurrentNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallPhase) GetCurrentNodeOk() (*int64, bool) {
	if o == nil || o.CurrentNode == nil {
		return nil, false
	}
	return o.CurrentNode, true
}

// HasCurrentNode returns a boolean if a field has been set.
func (o *ApplianceClusterInstallPhase) HasCurrentNode() bool {
	if o != nil && o.CurrentNode != nil {
		return true
	}

	return false
}

// SetCurrentNode gets a reference to the given int64 and assigns it to the CurrentNode field.
func (o *ApplianceClusterInstallPhase) SetCurrentNode(v int64) {
	o.CurrentNode = &v
}

// GetCurrentSubphase returns the CurrentSubphase field value if set, zero value otherwise.
func (o *ApplianceClusterInstallPhase) GetCurrentSubphase() int64 {
	if o == nil || o.CurrentSubphase == nil {
		var ret int64
		return ret
	}
	return *o.CurrentSubphase
}

// GetCurrentSubphaseOk returns a tuple with the CurrentSubphase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallPhase) GetCurrentSubphaseOk() (*int64, bool) {
	if o == nil || o.CurrentSubphase == nil {
		return nil, false
	}
	return o.CurrentSubphase, true
}

// HasCurrentSubphase returns a boolean if a field has been set.
func (o *ApplianceClusterInstallPhase) HasCurrentSubphase() bool {
	if o != nil && o.CurrentSubphase != nil {
		return true
	}

	return false
}

// SetCurrentSubphase gets a reference to the given int64 and assigns it to the CurrentSubphase field.
func (o *ApplianceClusterInstallPhase) SetCurrentSubphase(v int64) {
	o.CurrentSubphase = &v
}

// GetElapsedTime returns the ElapsedTime field value if set, zero value otherwise.
func (o *ApplianceClusterInstallPhase) GetElapsedTime() int64 {
	if o == nil || o.ElapsedTime == nil {
		var ret int64
		return ret
	}
	return *o.ElapsedTime
}

// GetElapsedTimeOk returns a tuple with the ElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallPhase) GetElapsedTimeOk() (*int64, bool) {
	if o == nil || o.ElapsedTime == nil {
		return nil, false
	}
	return o.ElapsedTime, true
}

// HasElapsedTime returns a boolean if a field has been set.
func (o *ApplianceClusterInstallPhase) HasElapsedTime() bool {
	if o != nil && o.ElapsedTime != nil {
		return true
	}

	return false
}

// SetElapsedTime gets a reference to the given int64 and assigns it to the ElapsedTime field.
func (o *ApplianceClusterInstallPhase) SetElapsedTime(v int64) {
	o.ElapsedTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ApplianceClusterInstallPhase) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallPhase) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ApplianceClusterInstallPhase) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ApplianceClusterInstallPhase) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *ApplianceClusterInstallPhase) GetFailed() bool {
	if o == nil || o.Failed == nil {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallPhase) GetFailedOk() (*bool, bool) {
	if o == nil || o.Failed == nil {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *ApplianceClusterInstallPhase) HasFailed() bool {
	if o != nil && o.Failed != nil {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *ApplianceClusterInstallPhase) SetFailed(v bool) {
	o.Failed = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApplianceClusterInstallPhase) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallPhase) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApplianceClusterInstallPhase) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApplianceClusterInstallPhase) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplianceClusterInstallPhase) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallPhase) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplianceClusterInstallPhase) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplianceClusterInstallPhase) SetName(v string) {
	o.Name = &v
}

// GetPendingNodes returns the PendingNodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceClusterInstallPhase) GetPendingNodes() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.PendingNodes
}

// GetPendingNodesOk returns a tuple with the PendingNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceClusterInstallPhase) GetPendingNodesOk() ([]int64, bool) {
	if o == nil || o.PendingNodes == nil {
		return nil, false
	}
	return o.PendingNodes, true
}

// HasPendingNodes returns a boolean if a field has been set.
func (o *ApplianceClusterInstallPhase) HasPendingNodes() bool {
	if o != nil && o.PendingNodes != nil {
		return true
	}

	return false
}

// SetPendingNodes gets a reference to the given []int64 and assigns it to the PendingNodes field.
func (o *ApplianceClusterInstallPhase) SetPendingNodes(v []int64) {
	o.PendingNodes = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApplianceClusterInstallPhase) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallPhase) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApplianceClusterInstallPhase) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApplianceClusterInstallPhase) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceClusterInstallPhase) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallPhase) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceClusterInstallPhase) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceClusterInstallPhase) SetStatus(v string) {
	o.Status = &v
}

func (o ApplianceClusterInstallPhase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CurrentNode != nil {
		toSerialize["CurrentNode"] = o.CurrentNode
	}
	if o.CurrentSubphase != nil {
		toSerialize["CurrentSubphase"] = o.CurrentSubphase
	}
	if o.ElapsedTime != nil {
		toSerialize["ElapsedTime"] = o.ElapsedTime
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.Failed != nil {
		toSerialize["Failed"] = o.Failed
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PendingNodes != nil {
		toSerialize["PendingNodes"] = o.PendingNodes
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceClusterInstallPhase) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceClusterInstallPhaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Node number install is running on.
		CurrentNode *int64 `json:"CurrentNode,omitempty"`
		// Subphase of the phase currently running.
		CurrentSubphase *int64 `json:"CurrentSubphase,omitempty"`
		// Elapsed time in seconds to complete the current install phase.
		ElapsedTime *int64 `json:"ElapsedTime,omitempty"`
		// End date of the software install phase.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// Indicates if the install phase has failed or not.
		Failed *bool `json:"Failed,omitempty"`
		// Status message set during the install phase.
		Message *string `json:"Message,omitempty"`
		// Name of the install phase. * `Backup` - Backup the currently running node. * `EnableBootstrap` - Disable echo and enable nginx on the currently running node. * `UpdateAnsibleHosts` - Update /etc/ansible/hosts on each node. * `SyncImages` - Sync images and manifest to each node. * `ConfigureEtcd` - Configure etcd on each node. * `DeployServices` - Deploy kubernetes services from node 1. * `InstallEquinox` - Configure and install equinox on each node. * `Validate` - Validate equinox is running in primary/secondary mode. * `CheckApplianceClusterState` - Check the appliance cluster status. * `Success` - Upgrade completed successfully. * `Fail` - Indicates that the upgrade process has failed. * `Cancel` - Indicates that the upgrade was canceled by the Intersight Appliance.
		Name         *string `json:"Name,omitempty"`
		PendingNodes []int64 `json:"PendingNodes,omitempty"`
		// Start date of the software install phase.
		StartTime *time.Time `json:"StartTime,omitempty"`
		// Status of the install phase.
		Status *string `json:"Status,omitempty"`
	}

	varApplianceClusterInstallPhaseWithoutEmbeddedStruct := ApplianceClusterInstallPhaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceClusterInstallPhaseWithoutEmbeddedStruct)
	if err == nil {
		varApplianceClusterInstallPhase := _ApplianceClusterInstallPhase{}
		varApplianceClusterInstallPhase.ClassId = varApplianceClusterInstallPhaseWithoutEmbeddedStruct.ClassId
		varApplianceClusterInstallPhase.ObjectType = varApplianceClusterInstallPhaseWithoutEmbeddedStruct.ObjectType
		varApplianceClusterInstallPhase.CurrentNode = varApplianceClusterInstallPhaseWithoutEmbeddedStruct.CurrentNode
		varApplianceClusterInstallPhase.CurrentSubphase = varApplianceClusterInstallPhaseWithoutEmbeddedStruct.CurrentSubphase
		varApplianceClusterInstallPhase.ElapsedTime = varApplianceClusterInstallPhaseWithoutEmbeddedStruct.ElapsedTime
		varApplianceClusterInstallPhase.EndTime = varApplianceClusterInstallPhaseWithoutEmbeddedStruct.EndTime
		varApplianceClusterInstallPhase.Failed = varApplianceClusterInstallPhaseWithoutEmbeddedStruct.Failed
		varApplianceClusterInstallPhase.Message = varApplianceClusterInstallPhaseWithoutEmbeddedStruct.Message
		varApplianceClusterInstallPhase.Name = varApplianceClusterInstallPhaseWithoutEmbeddedStruct.Name
		varApplianceClusterInstallPhase.PendingNodes = varApplianceClusterInstallPhaseWithoutEmbeddedStruct.PendingNodes
		varApplianceClusterInstallPhase.StartTime = varApplianceClusterInstallPhaseWithoutEmbeddedStruct.StartTime
		varApplianceClusterInstallPhase.Status = varApplianceClusterInstallPhaseWithoutEmbeddedStruct.Status
		*o = ApplianceClusterInstallPhase(varApplianceClusterInstallPhase)
	} else {
		return err
	}

	varApplianceClusterInstallPhase := _ApplianceClusterInstallPhase{}

	err = json.Unmarshal(bytes, &varApplianceClusterInstallPhase)
	if err == nil {
		o.MoBaseComplexType = varApplianceClusterInstallPhase.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CurrentNode")
		delete(additionalProperties, "CurrentSubphase")
		delete(additionalProperties, "ElapsedTime")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "Failed")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PendingNodes")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")

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

type NullableApplianceClusterInstallPhase struct {
	value *ApplianceClusterInstallPhase
	isSet bool
}

func (v NullableApplianceClusterInstallPhase) Get() *ApplianceClusterInstallPhase {
	return v.value
}

func (v *NullableApplianceClusterInstallPhase) Set(val *ApplianceClusterInstallPhase) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceClusterInstallPhase) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceClusterInstallPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceClusterInstallPhase(val *ApplianceClusterInstallPhase) *NullableApplianceClusterInstallPhase {
	return &NullableApplianceClusterInstallPhase{value: val, isSet: true}
}

func (v NullableApplianceClusterInstallPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceClusterInstallPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}