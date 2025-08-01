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

// checks if the BulkSubRequestObj type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkSubRequestObj{}

// BulkSubRequestObj The sub request object is created for every subrequest in the incoming request.
type BulkSubRequestObj struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string    `json:"ObjectType"`
	Body       *MoBaseMo `json:"Body,omitempty"`
	// The body of the sub-request in string format.
	BodyString *string `json:"BodyString,omitempty"`
	// The time at which processing of this request completed.
	ExecutionCompletionTime *string `json:"ExecutionCompletionTime,omitempty"`
	// The time at which processing of this request started.
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty"`
	// For Async Bulk Mo Operations this flag will be set to true.
	IsBulkMoOp *bool `json:"IsBulkMoOp,omitempty"`
	// This flag indicates if an already existing object was found or not after execution of the action CheckObjectPresence.
	IsObjectPresent *bool                 `json:"IsObjectPresent,omitempty"`
	Result          NullableBulkApiResult `json:"Result,omitempty"`
	// Skip the already present objects. The value from the Request.
	SkipDuplicates *bool `json:"SkipDuplicates,omitempty"`
	// The status of the request. * `Pending` - Indicates that the request is yet to be processed. * `ObjPresenceCheckInProgress` - Indicates that the checking for object presence is in progress. * `ObjPresenceCheckInComplete` - Indicates that the request is being processed. * `ObjPresenceCheckFailed` - Indicates that the checking for object presence failed. * `Processing` - Indicates that the request is being processed. * `TimedOut` - Indicates that the request processing timed out. * `Failed` - Indicates that the request processing failed. * `Completed` - Indicates that the request processing is complete. * `Skipped` - Indicates that the request was skipped.
	Status *string `json:"Status,omitempty"`
	// This flag indicates if the a system defined object was detected after execution of the action CheckObjectPresence.
	SystemDefinedObjectDetected *bool `json:"SystemDefinedObjectDetected,omitempty"`
	// Used with PATCH & DELETE actions. The moid of an existing object instance.
	TargetMoid *string `json:"TargetMoid,omitempty"`
	// The URI on which this bulk action is to be performed.
	Uri *string `json:"Uri,omitempty"`
	// The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). * `POST` - Used to create a REST resource. * `PATCH` - Used to update a REST resource. * `DELETE` - Used to delete a REST resource.
	Verb                 *string                         `json:"Verb,omitempty"`
	AsyncRequest         NullableBulkResultRelationship  `json:"AsyncRequest,omitempty"`
	Request              NullableBulkRequestRelationship `json:"Request,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkSubRequestObj BulkSubRequestObj

// NewBulkSubRequestObj instantiates a new BulkSubRequestObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkSubRequestObj(classId string, objectType string) *BulkSubRequestObj {
	this := BulkSubRequestObj{}
	this.ClassId = classId
	this.ObjectType = objectType
	var verb string = "POST"
	this.Verb = &verb
	return &this
}

// NewBulkSubRequestObjWithDefaults instantiates a new BulkSubRequestObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkSubRequestObjWithDefaults() *BulkSubRequestObj {
	this := BulkSubRequestObj{}
	var classId string = "bulk.SubRequestObj"
	this.ClassId = classId
	var objectType string = "bulk.SubRequestObj"
	this.ObjectType = objectType
	var verb string = "POST"
	this.Verb = &verb
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkSubRequestObj) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkSubRequestObj) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "bulk.SubRequestObj" of the ClassId field.
func (o *BulkSubRequestObj) GetDefaultClassId() interface{} {
	return "bulk.SubRequestObj"
}

// GetObjectType returns the ObjectType field value
func (o *BulkSubRequestObj) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkSubRequestObj) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "bulk.SubRequestObj" of the ObjectType field.
func (o *BulkSubRequestObj) GetDefaultObjectType() interface{} {
	return "bulk.SubRequestObj"
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetBody() MoBaseMo {
	if o == nil || IsNil(o.Body) {
		var ret MoBaseMo
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetBodyOk() (*MoBaseMo, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given MoBaseMo and assigns it to the Body field.
func (o *BulkSubRequestObj) SetBody(v MoBaseMo) {
	o.Body = &v
}

// GetBodyString returns the BodyString field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetBodyString() string {
	if o == nil || IsNil(o.BodyString) {
		var ret string
		return ret
	}
	return *o.BodyString
}

// GetBodyStringOk returns a tuple with the BodyString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetBodyStringOk() (*string, bool) {
	if o == nil || IsNil(o.BodyString) {
		return nil, false
	}
	return o.BodyString, true
}

// HasBodyString returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasBodyString() bool {
	if o != nil && !IsNil(o.BodyString) {
		return true
	}

	return false
}

// SetBodyString gets a reference to the given string and assigns it to the BodyString field.
func (o *BulkSubRequestObj) SetBodyString(v string) {
	o.BodyString = &v
}

// GetExecutionCompletionTime returns the ExecutionCompletionTime field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetExecutionCompletionTime() string {
	if o == nil || IsNil(o.ExecutionCompletionTime) {
		var ret string
		return ret
	}
	return *o.ExecutionCompletionTime
}

// GetExecutionCompletionTimeOk returns a tuple with the ExecutionCompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetExecutionCompletionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionCompletionTime) {
		return nil, false
	}
	return o.ExecutionCompletionTime, true
}

// HasExecutionCompletionTime returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasExecutionCompletionTime() bool {
	if o != nil && !IsNil(o.ExecutionCompletionTime) {
		return true
	}

	return false
}

// SetExecutionCompletionTime gets a reference to the given string and assigns it to the ExecutionCompletionTime field.
func (o *BulkSubRequestObj) SetExecutionCompletionTime(v string) {
	o.ExecutionCompletionTime = &v
}

// GetExecutionStartTime returns the ExecutionStartTime field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetExecutionStartTime() string {
	if o == nil || IsNil(o.ExecutionStartTime) {
		var ret string
		return ret
	}
	return *o.ExecutionStartTime
}

// GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetExecutionStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionStartTime) {
		return nil, false
	}
	return o.ExecutionStartTime, true
}

// HasExecutionStartTime returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasExecutionStartTime() bool {
	if o != nil && !IsNil(o.ExecutionStartTime) {
		return true
	}

	return false
}

// SetExecutionStartTime gets a reference to the given string and assigns it to the ExecutionStartTime field.
func (o *BulkSubRequestObj) SetExecutionStartTime(v string) {
	o.ExecutionStartTime = &v
}

// GetIsBulkMoOp returns the IsBulkMoOp field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetIsBulkMoOp() bool {
	if o == nil || IsNil(o.IsBulkMoOp) {
		var ret bool
		return ret
	}
	return *o.IsBulkMoOp
}

// GetIsBulkMoOpOk returns a tuple with the IsBulkMoOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetIsBulkMoOpOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBulkMoOp) {
		return nil, false
	}
	return o.IsBulkMoOp, true
}

// HasIsBulkMoOp returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasIsBulkMoOp() bool {
	if o != nil && !IsNil(o.IsBulkMoOp) {
		return true
	}

	return false
}

// SetIsBulkMoOp gets a reference to the given bool and assigns it to the IsBulkMoOp field.
func (o *BulkSubRequestObj) SetIsBulkMoOp(v bool) {
	o.IsBulkMoOp = &v
}

// GetIsObjectPresent returns the IsObjectPresent field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetIsObjectPresent() bool {
	if o == nil || IsNil(o.IsObjectPresent) {
		var ret bool
		return ret
	}
	return *o.IsObjectPresent
}

// GetIsObjectPresentOk returns a tuple with the IsObjectPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetIsObjectPresentOk() (*bool, bool) {
	if o == nil || IsNil(o.IsObjectPresent) {
		return nil, false
	}
	return o.IsObjectPresent, true
}

// HasIsObjectPresent returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasIsObjectPresent() bool {
	if o != nil && !IsNil(o.IsObjectPresent) {
		return true
	}

	return false
}

// SetIsObjectPresent gets a reference to the given bool and assigns it to the IsObjectPresent field.
func (o *BulkSubRequestObj) SetIsObjectPresent(v bool) {
	o.IsObjectPresent = &v
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkSubRequestObj) GetResult() BulkApiResult {
	if o == nil || IsNil(o.Result.Get()) {
		var ret BulkApiResult
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkSubRequestObj) GetResultOk() (*BulkApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableBulkApiResult and assigns it to the Result field.
func (o *BulkSubRequestObj) SetResult(v BulkApiResult) {
	o.Result.Set(&v)
}

// SetResultNil sets the value for Result to be an explicit nil
func (o *BulkSubRequestObj) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *BulkSubRequestObj) UnsetResult() {
	o.Result.Unset()
}

// GetSkipDuplicates returns the SkipDuplicates field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetSkipDuplicates() bool {
	if o == nil || IsNil(o.SkipDuplicates) {
		var ret bool
		return ret
	}
	return *o.SkipDuplicates
}

// GetSkipDuplicatesOk returns a tuple with the SkipDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetSkipDuplicatesOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipDuplicates) {
		return nil, false
	}
	return o.SkipDuplicates, true
}

// HasSkipDuplicates returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasSkipDuplicates() bool {
	if o != nil && !IsNil(o.SkipDuplicates) {
		return true
	}

	return false
}

// SetSkipDuplicates gets a reference to the given bool and assigns it to the SkipDuplicates field.
func (o *BulkSubRequestObj) SetSkipDuplicates(v bool) {
	o.SkipDuplicates = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BulkSubRequestObj) SetStatus(v string) {
	o.Status = &v
}

// GetSystemDefinedObjectDetected returns the SystemDefinedObjectDetected field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetSystemDefinedObjectDetected() bool {
	if o == nil || IsNil(o.SystemDefinedObjectDetected) {
		var ret bool
		return ret
	}
	return *o.SystemDefinedObjectDetected
}

// GetSystemDefinedObjectDetectedOk returns a tuple with the SystemDefinedObjectDetected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetSystemDefinedObjectDetectedOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemDefinedObjectDetected) {
		return nil, false
	}
	return o.SystemDefinedObjectDetected, true
}

// HasSystemDefinedObjectDetected returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasSystemDefinedObjectDetected() bool {
	if o != nil && !IsNil(o.SystemDefinedObjectDetected) {
		return true
	}

	return false
}

// SetSystemDefinedObjectDetected gets a reference to the given bool and assigns it to the SystemDefinedObjectDetected field.
func (o *BulkSubRequestObj) SetSystemDefinedObjectDetected(v bool) {
	o.SystemDefinedObjectDetected = &v
}

// GetTargetMoid returns the TargetMoid field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetTargetMoid() string {
	if o == nil || IsNil(o.TargetMoid) {
		var ret string
		return ret
	}
	return *o.TargetMoid
}

// GetTargetMoidOk returns a tuple with the TargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetTargetMoidOk() (*string, bool) {
	if o == nil || IsNil(o.TargetMoid) {
		return nil, false
	}
	return o.TargetMoid, true
}

// HasTargetMoid returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasTargetMoid() bool {
	if o != nil && !IsNil(o.TargetMoid) {
		return true
	}

	return false
}

// SetTargetMoid gets a reference to the given string and assigns it to the TargetMoid field.
func (o *BulkSubRequestObj) SetTargetMoid(v string) {
	o.TargetMoid = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *BulkSubRequestObj) SetUri(v string) {
	o.Uri = &v
}

// GetVerb returns the Verb field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetVerb() string {
	if o == nil || IsNil(o.Verb) {
		var ret string
		return ret
	}
	return *o.Verb
}

// GetVerbOk returns a tuple with the Verb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetVerbOk() (*string, bool) {
	if o == nil || IsNil(o.Verb) {
		return nil, false
	}
	return o.Verb, true
}

// HasVerb returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasVerb() bool {
	if o != nil && !IsNil(o.Verb) {
		return true
	}

	return false
}

// SetVerb gets a reference to the given string and assigns it to the Verb field.
func (o *BulkSubRequestObj) SetVerb(v string) {
	o.Verb = &v
}

// GetAsyncRequest returns the AsyncRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkSubRequestObj) GetAsyncRequest() BulkResultRelationship {
	if o == nil || IsNil(o.AsyncRequest.Get()) {
		var ret BulkResultRelationship
		return ret
	}
	return *o.AsyncRequest.Get()
}

// GetAsyncRequestOk returns a tuple with the AsyncRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkSubRequestObj) GetAsyncRequestOk() (*BulkResultRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AsyncRequest.Get(), o.AsyncRequest.IsSet()
}

// HasAsyncRequest returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasAsyncRequest() bool {
	if o != nil && o.AsyncRequest.IsSet() {
		return true
	}

	return false
}

// SetAsyncRequest gets a reference to the given NullableBulkResultRelationship and assigns it to the AsyncRequest field.
func (o *BulkSubRequestObj) SetAsyncRequest(v BulkResultRelationship) {
	o.AsyncRequest.Set(&v)
}

// SetAsyncRequestNil sets the value for AsyncRequest to be an explicit nil
func (o *BulkSubRequestObj) SetAsyncRequestNil() {
	o.AsyncRequest.Set(nil)
}

// UnsetAsyncRequest ensures that no value is present for AsyncRequest, not even an explicit nil
func (o *BulkSubRequestObj) UnsetAsyncRequest() {
	o.AsyncRequest.Unset()
}

// GetRequest returns the Request field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkSubRequestObj) GetRequest() BulkRequestRelationship {
	if o == nil || IsNil(o.Request.Get()) {
		var ret BulkRequestRelationship
		return ret
	}
	return *o.Request.Get()
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkSubRequestObj) GetRequestOk() (*BulkRequestRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Request.Get(), o.Request.IsSet()
}

// HasRequest returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasRequest() bool {
	if o != nil && o.Request.IsSet() {
		return true
	}

	return false
}

// SetRequest gets a reference to the given NullableBulkRequestRelationship and assigns it to the Request field.
func (o *BulkSubRequestObj) SetRequest(v BulkRequestRelationship) {
	o.Request.Set(&v)
}

// SetRequestNil sets the value for Request to be an explicit nil
func (o *BulkSubRequestObj) SetRequestNil() {
	o.Request.Set(nil)
}

// UnsetRequest ensures that no value is present for Request, not even an explicit nil
func (o *BulkSubRequestObj) UnsetRequest() {
	o.Request.Unset()
}

func (o BulkSubRequestObj) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkSubRequestObj) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Body) {
		toSerialize["Body"] = o.Body
	}
	if !IsNil(o.BodyString) {
		toSerialize["BodyString"] = o.BodyString
	}
	if !IsNil(o.ExecutionCompletionTime) {
		toSerialize["ExecutionCompletionTime"] = o.ExecutionCompletionTime
	}
	if !IsNil(o.ExecutionStartTime) {
		toSerialize["ExecutionStartTime"] = o.ExecutionStartTime
	}
	if !IsNil(o.IsBulkMoOp) {
		toSerialize["IsBulkMoOp"] = o.IsBulkMoOp
	}
	if !IsNil(o.IsObjectPresent) {
		toSerialize["IsObjectPresent"] = o.IsObjectPresent
	}
	if o.Result.IsSet() {
		toSerialize["Result"] = o.Result.Get()
	}
	if !IsNil(o.SkipDuplicates) {
		toSerialize["SkipDuplicates"] = o.SkipDuplicates
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.SystemDefinedObjectDetected) {
		toSerialize["SystemDefinedObjectDetected"] = o.SystemDefinedObjectDetected
	}
	if !IsNil(o.TargetMoid) {
		toSerialize["TargetMoid"] = o.TargetMoid
	}
	if !IsNil(o.Uri) {
		toSerialize["Uri"] = o.Uri
	}
	if !IsNil(o.Verb) {
		toSerialize["Verb"] = o.Verb
	}
	if o.AsyncRequest.IsSet() {
		toSerialize["AsyncRequest"] = o.AsyncRequest.Get()
	}
	if o.Request.IsSet() {
		toSerialize["Request"] = o.Request.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkSubRequestObj) UnmarshalJSON(data []byte) (err error) {
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
	type BulkSubRequestObjWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string    `json:"ObjectType"`
		Body       *MoBaseMo `json:"Body,omitempty"`
		// The body of the sub-request in string format.
		BodyString *string `json:"BodyString,omitempty"`
		// The time at which processing of this request completed.
		ExecutionCompletionTime *string `json:"ExecutionCompletionTime,omitempty"`
		// The time at which processing of this request started.
		ExecutionStartTime *string `json:"ExecutionStartTime,omitempty"`
		// For Async Bulk Mo Operations this flag will be set to true.
		IsBulkMoOp *bool `json:"IsBulkMoOp,omitempty"`
		// This flag indicates if an already existing object was found or not after execution of the action CheckObjectPresence.
		IsObjectPresent *bool                 `json:"IsObjectPresent,omitempty"`
		Result          NullableBulkApiResult `json:"Result,omitempty"`
		// Skip the already present objects. The value from the Request.
		SkipDuplicates *bool `json:"SkipDuplicates,omitempty"`
		// The status of the request. * `Pending` - Indicates that the request is yet to be processed. * `ObjPresenceCheckInProgress` - Indicates that the checking for object presence is in progress. * `ObjPresenceCheckInComplete` - Indicates that the request is being processed. * `ObjPresenceCheckFailed` - Indicates that the checking for object presence failed. * `Processing` - Indicates that the request is being processed. * `TimedOut` - Indicates that the request processing timed out. * `Failed` - Indicates that the request processing failed. * `Completed` - Indicates that the request processing is complete. * `Skipped` - Indicates that the request was skipped.
		Status *string `json:"Status,omitempty"`
		// This flag indicates if the a system defined object was detected after execution of the action CheckObjectPresence.
		SystemDefinedObjectDetected *bool `json:"SystemDefinedObjectDetected,omitempty"`
		// Used with PATCH & DELETE actions. The moid of an existing object instance.
		TargetMoid *string `json:"TargetMoid,omitempty"`
		// The URI on which this bulk action is to be performed.
		Uri *string `json:"Uri,omitempty"`
		// The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). * `POST` - Used to create a REST resource. * `PATCH` - Used to update a REST resource. * `DELETE` - Used to delete a REST resource.
		Verb         *string                         `json:"Verb,omitempty"`
		AsyncRequest NullableBulkResultRelationship  `json:"AsyncRequest,omitempty"`
		Request      NullableBulkRequestRelationship `json:"Request,omitempty"`
	}

	varBulkSubRequestObjWithoutEmbeddedStruct := BulkSubRequestObjWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBulkSubRequestObjWithoutEmbeddedStruct)
	if err == nil {
		varBulkSubRequestObj := _BulkSubRequestObj{}
		varBulkSubRequestObj.ClassId = varBulkSubRequestObjWithoutEmbeddedStruct.ClassId
		varBulkSubRequestObj.ObjectType = varBulkSubRequestObjWithoutEmbeddedStruct.ObjectType
		varBulkSubRequestObj.Body = varBulkSubRequestObjWithoutEmbeddedStruct.Body
		varBulkSubRequestObj.BodyString = varBulkSubRequestObjWithoutEmbeddedStruct.BodyString
		varBulkSubRequestObj.ExecutionCompletionTime = varBulkSubRequestObjWithoutEmbeddedStruct.ExecutionCompletionTime
		varBulkSubRequestObj.ExecutionStartTime = varBulkSubRequestObjWithoutEmbeddedStruct.ExecutionStartTime
		varBulkSubRequestObj.IsBulkMoOp = varBulkSubRequestObjWithoutEmbeddedStruct.IsBulkMoOp
		varBulkSubRequestObj.IsObjectPresent = varBulkSubRequestObjWithoutEmbeddedStruct.IsObjectPresent
		varBulkSubRequestObj.Result = varBulkSubRequestObjWithoutEmbeddedStruct.Result
		varBulkSubRequestObj.SkipDuplicates = varBulkSubRequestObjWithoutEmbeddedStruct.SkipDuplicates
		varBulkSubRequestObj.Status = varBulkSubRequestObjWithoutEmbeddedStruct.Status
		varBulkSubRequestObj.SystemDefinedObjectDetected = varBulkSubRequestObjWithoutEmbeddedStruct.SystemDefinedObjectDetected
		varBulkSubRequestObj.TargetMoid = varBulkSubRequestObjWithoutEmbeddedStruct.TargetMoid
		varBulkSubRequestObj.Uri = varBulkSubRequestObjWithoutEmbeddedStruct.Uri
		varBulkSubRequestObj.Verb = varBulkSubRequestObjWithoutEmbeddedStruct.Verb
		varBulkSubRequestObj.AsyncRequest = varBulkSubRequestObjWithoutEmbeddedStruct.AsyncRequest
		varBulkSubRequestObj.Request = varBulkSubRequestObjWithoutEmbeddedStruct.Request
		*o = BulkSubRequestObj(varBulkSubRequestObj)
	} else {
		return err
	}

	varBulkSubRequestObj := _BulkSubRequestObj{}

	err = json.Unmarshal(data, &varBulkSubRequestObj)
	if err == nil {
		o.MoBaseMo = varBulkSubRequestObj.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Body")
		delete(additionalProperties, "BodyString")
		delete(additionalProperties, "ExecutionCompletionTime")
		delete(additionalProperties, "ExecutionStartTime")
		delete(additionalProperties, "IsBulkMoOp")
		delete(additionalProperties, "IsObjectPresent")
		delete(additionalProperties, "Result")
		delete(additionalProperties, "SkipDuplicates")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "SystemDefinedObjectDetected")
		delete(additionalProperties, "TargetMoid")
		delete(additionalProperties, "Uri")
		delete(additionalProperties, "Verb")
		delete(additionalProperties, "AsyncRequest")
		delete(additionalProperties, "Request")

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

type NullableBulkSubRequestObj struct {
	value *BulkSubRequestObj
	isSet bool
}

func (v NullableBulkSubRequestObj) Get() *BulkSubRequestObj {
	return v.value
}

func (v *NullableBulkSubRequestObj) Set(val *BulkSubRequestObj) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkSubRequestObj) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkSubRequestObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkSubRequestObj(val *BulkSubRequestObj) *NullableBulkSubRequestObj {
	return &NullableBulkSubRequestObj{value: val, isSet: true}
}

func (v NullableBulkSubRequestObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkSubRequestObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
