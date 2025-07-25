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

// checks if the BulkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkRequest{}

// BulkRequest The bulk.Request API allows users to perform API actions (Create, Update or Delete) in bulk, on a given URI. It is possible to operate on multiple subpaths relative to the provided URI (For example, it would be possible to perform a PATCH action on multiple objects of a given REST resource type).
type BulkRequest struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The action to be taken when an error occurs during processing of the request. * `Stop` - Stop the processing of the request after the first error. * `Proceed` - Proceed with the processing of the request even when an error occurs.
	ActionOnError *string  `json:"ActionOnError,omitempty"`
	Actions       []string `json:"Actions,omitempty"`
	// The timestamp when the request processing completed.
	CompletionTime *string          `json:"CompletionTime,omitempty"`
	Headers        []BulkHttpHeader `json:"Headers,omitempty"`
	// The number of sub requests received in this request.
	NumSubRequests *int64 `json:"NumSubRequests,omitempty"`
	// The moid of the organization under which this request was issued.
	OrgMoid *string `json:"OrgMoid,omitempty"`
	// The timestamp when the request was received.
	RequestReceivedTime *string          `json:"RequestReceivedTime,omitempty"`
	Requests            []BulkSubRequest `json:"Requests,omitempty"`
	Results             []BulkApiResult  `json:"Results,omitempty"`
	// Skip the already present objects.
	SkipDuplicates *bool `json:"SkipDuplicates,omitempty"`
	// The processing status of the Request. * `NotStarted` - Indicates that the request processing has not begun yet. * `ObjPresenceCheckInProgress` - Indicates that the object presence check is in progress for this request. * `ObjPresenceCheckComplete` - Indicates that the object presence check is complete. * `ExecutionInProgress` - Indicates that the request processing is in progress. * `Completed` - Indicates that the request processing has been completed successfully. * `CompletedWithErrors` - Indicates that the request processing has one or more failed subrequests. * `Failed` - Indicates that the processing of this request failed. * `TimedOut` - Indicates that the request processing timed out.
	Status *string `json:"Status,omitempty"`
	// The status message corresponding to the status.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// The URI on which this bulk action is to be performed. The value will be used when there is no override in the SubRequest.
	// Deprecated
	Uri *string `json:"Uri,omitempty"`
	// The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). The value will be used when there is no override in the SubRequest. * `POST` - Used to create a REST resource. * `PATCH` - Used to update a REST resource. * `DELETE` - Used to delete a REST resource.
	// Deprecated
	Verb *string `json:"Verb,omitempty"`
	// An array of relationships to bulkSubRequestObj resources.
	AsyncResults []BulkSubRequestObjRelationship `json:"AsyncResults,omitempty"`
	// An array of relationships to bulkSubRequestObj resources.
	AsyncResultsFailed   []BulkSubRequestObjRelationship              `json:"AsyncResultsFailed,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	WorkflowInfo         NullableWorkflowWorkflowInfoRelationship     `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkRequest BulkRequest

// NewBulkRequest instantiates a new BulkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkRequest(classId string, objectType string) *BulkRequest {
	this := BulkRequest{}
	this.ClassId = classId
	this.ObjectType = objectType
	var actionOnError string = "Stop"
	this.ActionOnError = &actionOnError
	var verb string = "POST"
	this.Verb = &verb
	return &this
}

// NewBulkRequestWithDefaults instantiates a new BulkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkRequestWithDefaults() *BulkRequest {
	this := BulkRequest{}
	var classId string = "bulk.Request"
	this.ClassId = classId
	var objectType string = "bulk.Request"
	this.ObjectType = objectType
	var actionOnError string = "Stop"
	this.ActionOnError = &actionOnError
	var verb string = "POST"
	this.Verb = &verb
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkRequest) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkRequest) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "bulk.Request" of the ClassId field.
func (o *BulkRequest) GetDefaultClassId() interface{} {
	return "bulk.Request"
}

// GetObjectType returns the ObjectType field value
func (o *BulkRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "bulk.Request" of the ObjectType field.
func (o *BulkRequest) GetDefaultObjectType() interface{} {
	return "bulk.Request"
}

// GetActionOnError returns the ActionOnError field value if set, zero value otherwise.
func (o *BulkRequest) GetActionOnError() string {
	if o == nil || IsNil(o.ActionOnError) {
		var ret string
		return ret
	}
	return *o.ActionOnError
}

// GetActionOnErrorOk returns a tuple with the ActionOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetActionOnErrorOk() (*string, bool) {
	if o == nil || IsNil(o.ActionOnError) {
		return nil, false
	}
	return o.ActionOnError, true
}

// HasActionOnError returns a boolean if a field has been set.
func (o *BulkRequest) HasActionOnError() bool {
	if o != nil && !IsNil(o.ActionOnError) {
		return true
	}

	return false
}

// SetActionOnError gets a reference to the given string and assigns it to the ActionOnError field.
func (o *BulkRequest) SetActionOnError(v string) {
	o.ActionOnError = &v
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequest) GetActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequest) GetActionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *BulkRequest) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []string and assigns it to the Actions field.
func (o *BulkRequest) SetActions(v []string) {
	o.Actions = v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *BulkRequest) GetCompletionTime() string {
	if o == nil || IsNil(o.CompletionTime) {
		var ret string
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetCompletionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CompletionTime) {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *BulkRequest) HasCompletionTime() bool {
	if o != nil && !IsNil(o.CompletionTime) {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given string and assigns it to the CompletionTime field.
func (o *BulkRequest) SetCompletionTime(v string) {
	o.CompletionTime = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequest) GetHeaders() []BulkHttpHeader {
	if o == nil {
		var ret []BulkHttpHeader
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequest) GetHeadersOk() ([]BulkHttpHeader, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *BulkRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []BulkHttpHeader and assigns it to the Headers field.
func (o *BulkRequest) SetHeaders(v []BulkHttpHeader) {
	o.Headers = v
}

// GetNumSubRequests returns the NumSubRequests field value if set, zero value otherwise.
func (o *BulkRequest) GetNumSubRequests() int64 {
	if o == nil || IsNil(o.NumSubRequests) {
		var ret int64
		return ret
	}
	return *o.NumSubRequests
}

// GetNumSubRequestsOk returns a tuple with the NumSubRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetNumSubRequestsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumSubRequests) {
		return nil, false
	}
	return o.NumSubRequests, true
}

// HasNumSubRequests returns a boolean if a field has been set.
func (o *BulkRequest) HasNumSubRequests() bool {
	if o != nil && !IsNil(o.NumSubRequests) {
		return true
	}

	return false
}

// SetNumSubRequests gets a reference to the given int64 and assigns it to the NumSubRequests field.
func (o *BulkRequest) SetNumSubRequests(v int64) {
	o.NumSubRequests = &v
}

// GetOrgMoid returns the OrgMoid field value if set, zero value otherwise.
func (o *BulkRequest) GetOrgMoid() string {
	if o == nil || IsNil(o.OrgMoid) {
		var ret string
		return ret
	}
	return *o.OrgMoid
}

// GetOrgMoidOk returns a tuple with the OrgMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetOrgMoidOk() (*string, bool) {
	if o == nil || IsNil(o.OrgMoid) {
		return nil, false
	}
	return o.OrgMoid, true
}

// HasOrgMoid returns a boolean if a field has been set.
func (o *BulkRequest) HasOrgMoid() bool {
	if o != nil && !IsNil(o.OrgMoid) {
		return true
	}

	return false
}

// SetOrgMoid gets a reference to the given string and assigns it to the OrgMoid field.
func (o *BulkRequest) SetOrgMoid(v string) {
	o.OrgMoid = &v
}

// GetRequestReceivedTime returns the RequestReceivedTime field value if set, zero value otherwise.
func (o *BulkRequest) GetRequestReceivedTime() string {
	if o == nil || IsNil(o.RequestReceivedTime) {
		var ret string
		return ret
	}
	return *o.RequestReceivedTime
}

// GetRequestReceivedTimeOk returns a tuple with the RequestReceivedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetRequestReceivedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestReceivedTime) {
		return nil, false
	}
	return o.RequestReceivedTime, true
}

// HasRequestReceivedTime returns a boolean if a field has been set.
func (o *BulkRequest) HasRequestReceivedTime() bool {
	if o != nil && !IsNil(o.RequestReceivedTime) {
		return true
	}

	return false
}

// SetRequestReceivedTime gets a reference to the given string and assigns it to the RequestReceivedTime field.
func (o *BulkRequest) SetRequestReceivedTime(v string) {
	o.RequestReceivedTime = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequest) GetRequests() []BulkSubRequest {
	if o == nil {
		var ret []BulkSubRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequest) GetRequestsOk() ([]BulkSubRequest, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *BulkRequest) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given []BulkSubRequest and assigns it to the Requests field.
func (o *BulkRequest) SetRequests(v []BulkSubRequest) {
	o.Requests = v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequest) GetResults() []BulkApiResult {
	if o == nil {
		var ret []BulkApiResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequest) GetResultsOk() ([]BulkApiResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BulkRequest) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BulkApiResult and assigns it to the Results field.
func (o *BulkRequest) SetResults(v []BulkApiResult) {
	o.Results = v
}

// GetSkipDuplicates returns the SkipDuplicates field value if set, zero value otherwise.
func (o *BulkRequest) GetSkipDuplicates() bool {
	if o == nil || IsNil(o.SkipDuplicates) {
		var ret bool
		return ret
	}
	return *o.SkipDuplicates
}

// GetSkipDuplicatesOk returns a tuple with the SkipDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetSkipDuplicatesOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipDuplicates) {
		return nil, false
	}
	return o.SkipDuplicates, true
}

// HasSkipDuplicates returns a boolean if a field has been set.
func (o *BulkRequest) HasSkipDuplicates() bool {
	if o != nil && !IsNil(o.SkipDuplicates) {
		return true
	}

	return false
}

// SetSkipDuplicates gets a reference to the given bool and assigns it to the SkipDuplicates field.
func (o *BulkRequest) SetSkipDuplicates(v bool) {
	o.SkipDuplicates = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BulkRequest) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *BulkRequest) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequest) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *BulkRequest) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *BulkRequest) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
// Deprecated
func (o *BulkRequest) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *BulkRequest) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BulkRequest) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
// Deprecated
func (o *BulkRequest) SetUri(v string) {
	o.Uri = &v
}

// GetVerb returns the Verb field value if set, zero value otherwise.
// Deprecated
func (o *BulkRequest) GetVerb() string {
	if o == nil || IsNil(o.Verb) {
		var ret string
		return ret
	}
	return *o.Verb
}

// GetVerbOk returns a tuple with the Verb field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *BulkRequest) GetVerbOk() (*string, bool) {
	if o == nil || IsNil(o.Verb) {
		return nil, false
	}
	return o.Verb, true
}

// HasVerb returns a boolean if a field has been set.
func (o *BulkRequest) HasVerb() bool {
	if o != nil && !IsNil(o.Verb) {
		return true
	}

	return false
}

// SetVerb gets a reference to the given string and assigns it to the Verb field.
// Deprecated
func (o *BulkRequest) SetVerb(v string) {
	o.Verb = &v
}

// GetAsyncResults returns the AsyncResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequest) GetAsyncResults() []BulkSubRequestObjRelationship {
	if o == nil {
		var ret []BulkSubRequestObjRelationship
		return ret
	}
	return o.AsyncResults
}

// GetAsyncResultsOk returns a tuple with the AsyncResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequest) GetAsyncResultsOk() ([]BulkSubRequestObjRelationship, bool) {
	if o == nil || IsNil(o.AsyncResults) {
		return nil, false
	}
	return o.AsyncResults, true
}

// HasAsyncResults returns a boolean if a field has been set.
func (o *BulkRequest) HasAsyncResults() bool {
	if o != nil && !IsNil(o.AsyncResults) {
		return true
	}

	return false
}

// SetAsyncResults gets a reference to the given []BulkSubRequestObjRelationship and assigns it to the AsyncResults field.
func (o *BulkRequest) SetAsyncResults(v []BulkSubRequestObjRelationship) {
	o.AsyncResults = v
}

// GetAsyncResultsFailed returns the AsyncResultsFailed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequest) GetAsyncResultsFailed() []BulkSubRequestObjRelationship {
	if o == nil {
		var ret []BulkSubRequestObjRelationship
		return ret
	}
	return o.AsyncResultsFailed
}

// GetAsyncResultsFailedOk returns a tuple with the AsyncResultsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequest) GetAsyncResultsFailedOk() ([]BulkSubRequestObjRelationship, bool) {
	if o == nil || IsNil(o.AsyncResultsFailed) {
		return nil, false
	}
	return o.AsyncResultsFailed, true
}

// HasAsyncResultsFailed returns a boolean if a field has been set.
func (o *BulkRequest) HasAsyncResultsFailed() bool {
	if o != nil && !IsNil(o.AsyncResultsFailed) {
		return true
	}

	return false
}

// SetAsyncResultsFailed gets a reference to the given []BulkSubRequestObjRelationship and assigns it to the AsyncResultsFailed field.
func (o *BulkRequest) SetAsyncResultsFailed(v []BulkSubRequestObjRelationship) {
	o.AsyncResultsFailed = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequest) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequest) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkRequest) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkRequest) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *BulkRequest) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *BulkRequest) UnsetOrganization() {
	o.Organization.Unset()
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequest) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || IsNil(o.WorkflowInfo.Get()) {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo.Get()
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequest) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkflowInfo.Get(), o.WorkflowInfo.IsSet()
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *BulkRequest) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo.IsSet() {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given NullableWorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *BulkRequest) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo.Set(&v)
}

// SetWorkflowInfoNil sets the value for WorkflowInfo to be an explicit nil
func (o *BulkRequest) SetWorkflowInfoNil() {
	o.WorkflowInfo.Set(nil)
}

// UnsetWorkflowInfo ensures that no value is present for WorkflowInfo, not even an explicit nil
func (o *BulkRequest) UnsetWorkflowInfo() {
	o.WorkflowInfo.Unset()
}

func (o BulkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ActionOnError) {
		toSerialize["ActionOnError"] = o.ActionOnError
	}
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}
	if !IsNil(o.CompletionTime) {
		toSerialize["CompletionTime"] = o.CompletionTime
	}
	if o.Headers != nil {
		toSerialize["Headers"] = o.Headers
	}
	if !IsNil(o.NumSubRequests) {
		toSerialize["NumSubRequests"] = o.NumSubRequests
	}
	if !IsNil(o.OrgMoid) {
		toSerialize["OrgMoid"] = o.OrgMoid
	}
	if !IsNil(o.RequestReceivedTime) {
		toSerialize["RequestReceivedTime"] = o.RequestReceivedTime
	}
	if o.Requests != nil {
		toSerialize["Requests"] = o.Requests
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}
	if !IsNil(o.SkipDuplicates) {
		toSerialize["SkipDuplicates"] = o.SkipDuplicates
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if !IsNil(o.Uri) {
		toSerialize["Uri"] = o.Uri
	}
	if !IsNil(o.Verb) {
		toSerialize["Verb"] = o.Verb
	}
	if o.AsyncResults != nil {
		toSerialize["AsyncResults"] = o.AsyncResults
	}
	if o.AsyncResultsFailed != nil {
		toSerialize["AsyncResultsFailed"] = o.AsyncResultsFailed
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.WorkflowInfo.IsSet() {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkRequest) UnmarshalJSON(data []byte) (err error) {
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
	type BulkRequestWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The action to be taken when an error occurs during processing of the request. * `Stop` - Stop the processing of the request after the first error. * `Proceed` - Proceed with the processing of the request even when an error occurs.
		ActionOnError *string  `json:"ActionOnError,omitempty"`
		Actions       []string `json:"Actions,omitempty"`
		// The timestamp when the request processing completed.
		CompletionTime *string          `json:"CompletionTime,omitempty"`
		Headers        []BulkHttpHeader `json:"Headers,omitempty"`
		// The number of sub requests received in this request.
		NumSubRequests *int64 `json:"NumSubRequests,omitempty"`
		// The moid of the organization under which this request was issued.
		OrgMoid *string `json:"OrgMoid,omitempty"`
		// The timestamp when the request was received.
		RequestReceivedTime *string          `json:"RequestReceivedTime,omitempty"`
		Requests            []BulkSubRequest `json:"Requests,omitempty"`
		Results             []BulkApiResult  `json:"Results,omitempty"`
		// Skip the already present objects.
		SkipDuplicates *bool `json:"SkipDuplicates,omitempty"`
		// The processing status of the Request. * `NotStarted` - Indicates that the request processing has not begun yet. * `ObjPresenceCheckInProgress` - Indicates that the object presence check is in progress for this request. * `ObjPresenceCheckComplete` - Indicates that the object presence check is complete. * `ExecutionInProgress` - Indicates that the request processing is in progress. * `Completed` - Indicates that the request processing has been completed successfully. * `CompletedWithErrors` - Indicates that the request processing has one or more failed subrequests. * `Failed` - Indicates that the processing of this request failed. * `TimedOut` - Indicates that the request processing timed out.
		Status *string `json:"Status,omitempty"`
		// The status message corresponding to the status.
		StatusMessage *string `json:"StatusMessage,omitempty"`
		// The URI on which this bulk action is to be performed. The value will be used when there is no override in the SubRequest.
		// Deprecated
		Uri *string `json:"Uri,omitempty"`
		// The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). The value will be used when there is no override in the SubRequest. * `POST` - Used to create a REST resource. * `PATCH` - Used to update a REST resource. * `DELETE` - Used to delete a REST resource.
		// Deprecated
		Verb *string `json:"Verb,omitempty"`
		// An array of relationships to bulkSubRequestObj resources.
		AsyncResults []BulkSubRequestObjRelationship `json:"AsyncResults,omitempty"`
		// An array of relationships to bulkSubRequestObj resources.
		AsyncResultsFailed []BulkSubRequestObjRelationship              `json:"AsyncResultsFailed,omitempty"`
		Organization       NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		WorkflowInfo       NullableWorkflowWorkflowInfoRelationship     `json:"WorkflowInfo,omitempty"`
	}

	varBulkRequestWithoutEmbeddedStruct := BulkRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBulkRequestWithoutEmbeddedStruct)
	if err == nil {
		varBulkRequest := _BulkRequest{}
		varBulkRequest.ClassId = varBulkRequestWithoutEmbeddedStruct.ClassId
		varBulkRequest.ObjectType = varBulkRequestWithoutEmbeddedStruct.ObjectType
		varBulkRequest.ActionOnError = varBulkRequestWithoutEmbeddedStruct.ActionOnError
		varBulkRequest.Actions = varBulkRequestWithoutEmbeddedStruct.Actions
		varBulkRequest.CompletionTime = varBulkRequestWithoutEmbeddedStruct.CompletionTime
		varBulkRequest.Headers = varBulkRequestWithoutEmbeddedStruct.Headers
		varBulkRequest.NumSubRequests = varBulkRequestWithoutEmbeddedStruct.NumSubRequests
		varBulkRequest.OrgMoid = varBulkRequestWithoutEmbeddedStruct.OrgMoid
		varBulkRequest.RequestReceivedTime = varBulkRequestWithoutEmbeddedStruct.RequestReceivedTime
		varBulkRequest.Requests = varBulkRequestWithoutEmbeddedStruct.Requests
		varBulkRequest.Results = varBulkRequestWithoutEmbeddedStruct.Results
		varBulkRequest.SkipDuplicates = varBulkRequestWithoutEmbeddedStruct.SkipDuplicates
		varBulkRequest.Status = varBulkRequestWithoutEmbeddedStruct.Status
		varBulkRequest.StatusMessage = varBulkRequestWithoutEmbeddedStruct.StatusMessage
		varBulkRequest.Uri = varBulkRequestWithoutEmbeddedStruct.Uri
		varBulkRequest.Verb = varBulkRequestWithoutEmbeddedStruct.Verb
		varBulkRequest.AsyncResults = varBulkRequestWithoutEmbeddedStruct.AsyncResults
		varBulkRequest.AsyncResultsFailed = varBulkRequestWithoutEmbeddedStruct.AsyncResultsFailed
		varBulkRequest.Organization = varBulkRequestWithoutEmbeddedStruct.Organization
		varBulkRequest.WorkflowInfo = varBulkRequestWithoutEmbeddedStruct.WorkflowInfo
		*o = BulkRequest(varBulkRequest)
	} else {
		return err
	}

	varBulkRequest := _BulkRequest{}

	err = json.Unmarshal(data, &varBulkRequest)
	if err == nil {
		o.MoBaseMo = varBulkRequest.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActionOnError")
		delete(additionalProperties, "Actions")
		delete(additionalProperties, "CompletionTime")
		delete(additionalProperties, "Headers")
		delete(additionalProperties, "NumSubRequests")
		delete(additionalProperties, "OrgMoid")
		delete(additionalProperties, "RequestReceivedTime")
		delete(additionalProperties, "Requests")
		delete(additionalProperties, "Results")
		delete(additionalProperties, "SkipDuplicates")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "Uri")
		delete(additionalProperties, "Verb")
		delete(additionalProperties, "AsyncResults")
		delete(additionalProperties, "AsyncResultsFailed")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "WorkflowInfo")

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

type NullableBulkRequest struct {
	value *BulkRequest
	isSet bool
}

func (v NullableBulkRequest) Get() *BulkRequest {
	return v.value
}

func (v *NullableBulkRequest) Set(val *BulkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkRequest(val *BulkRequest) *NullableBulkRequest {
	return &NullableBulkRequest{value: val, isSet: true}
}

func (v NullableBulkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
