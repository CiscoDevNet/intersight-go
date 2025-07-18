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

// checks if the TamAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TamAction{}

// TamAction An action is used to react when an object satifies the condition specified in alert query. For e.g. an action in case of an object matching a fieldNotice query would be to create an alert instance of type 'fieldNotice' for the affected object.
type TamAction struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove).
	AffectedObjectType *string `json:"AffectedObjectType,omitempty"`
	// Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.). * `psirt` - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x). * `fieldNotice` - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html). * `eolAdvisory` - Represents product End of Life (EOL) type (https://www.cisco.com/c/en/us/products/eos-eol-policy.html).
	AlertType   *string          `json:"AlertType,omitempty"`
	Identifiers []TamIdentifiers `json:"Identifiers,omitempty"`
	// Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries.
	Name *string `json:"Name,omitempty"`
	// Operation type for the alert action. An action is used to carry out the process of \"reacting\" to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied. * `create` - Create an instance of AdvisoryInstance. * `remove` - Remove an instance of AdvisoryInstance.
	OperationType *string         `json:"OperationType,omitempty"`
	Queries       []TamQueryEntry `json:"Queries,omitempty"`
	// Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type. * `restApi` - Repesents the use of REST API for carrying out alert actions.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamAction TamAction

// NewTamAction instantiates a new TamAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamAction(classId string, objectType string) *TamAction {
	this := TamAction{}
	this.ClassId = classId
	this.ObjectType = objectType
	var alertType string = "psirt"
	this.AlertType = &alertType
	var operationType string = "create"
	this.OperationType = &operationType
	var type_ string = "restApi"
	this.Type = &type_
	return &this
}

// NewTamActionWithDefaults instantiates a new TamAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamActionWithDefaults() *TamAction {
	this := TamAction{}
	var classId string = "tam.Action"
	this.ClassId = classId
	var objectType string = "tam.Action"
	this.ObjectType = objectType
	var alertType string = "psirt"
	this.AlertType = &alertType
	var operationType string = "create"
	this.OperationType = &operationType
	var type_ string = "restApi"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamAction) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamAction) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamAction) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "tam.Action" of the ClassId field.
func (o *TamAction) GetDefaultClassId() interface{} {
	return "tam.Action"
}

// GetObjectType returns the ObjectType field value
func (o *TamAction) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamAction) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamAction) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "tam.Action" of the ObjectType field.
func (o *TamAction) GetDefaultObjectType() interface{} {
	return "tam.Action"
}

// GetAffectedObjectType returns the AffectedObjectType field value if set, zero value otherwise.
func (o *TamAction) GetAffectedObjectType() string {
	if o == nil || IsNil(o.AffectedObjectType) {
		var ret string
		return ret
	}
	return *o.AffectedObjectType
}

// GetAffectedObjectTypeOk returns a tuple with the AffectedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAction) GetAffectedObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AffectedObjectType) {
		return nil, false
	}
	return o.AffectedObjectType, true
}

// HasAffectedObjectType returns a boolean if a field has been set.
func (o *TamAction) HasAffectedObjectType() bool {
	if o != nil && !IsNil(o.AffectedObjectType) {
		return true
	}

	return false
}

// SetAffectedObjectType gets a reference to the given string and assigns it to the AffectedObjectType field.
func (o *TamAction) SetAffectedObjectType(v string) {
	o.AffectedObjectType = &v
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *TamAction) GetAlertType() string {
	if o == nil || IsNil(o.AlertType) {
		var ret string
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAction) GetAlertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AlertType) {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *TamAction) HasAlertType() bool {
	if o != nil && !IsNil(o.AlertType) {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given string and assigns it to the AlertType field.
func (o *TamAction) SetAlertType(v string) {
	o.AlertType = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamAction) GetIdentifiers() []TamIdentifiers {
	if o == nil {
		var ret []TamIdentifiers
		return ret
	}
	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamAction) GetIdentifiersOk() ([]TamIdentifiers, bool) {
	if o == nil || IsNil(o.Identifiers) {
		return nil, false
	}
	return o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *TamAction) HasIdentifiers() bool {
	if o != nil && !IsNil(o.Identifiers) {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []TamIdentifiers and assigns it to the Identifiers field.
func (o *TamAction) SetIdentifiers(v []TamIdentifiers) {
	o.Identifiers = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TamAction) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAction) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TamAction) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TamAction) SetName(v string) {
	o.Name = &v
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *TamAction) GetOperationType() string {
	if o == nil || IsNil(o.OperationType) {
		var ret string
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAction) GetOperationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OperationType) {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *TamAction) HasOperationType() bool {
	if o != nil && !IsNil(o.OperationType) {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given string and assigns it to the OperationType field.
func (o *TamAction) SetOperationType(v string) {
	o.OperationType = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamAction) GetQueries() []TamQueryEntry {
	if o == nil {
		var ret []TamQueryEntry
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamAction) GetQueriesOk() ([]TamQueryEntry, bool) {
	if o == nil || IsNil(o.Queries) {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *TamAction) HasQueries() bool {
	if o != nil && !IsNil(o.Queries) {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []TamQueryEntry and assigns it to the Queries field.
func (o *TamAction) SetQueries(v []TamQueryEntry) {
	o.Queries = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TamAction) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAction) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TamAction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TamAction) SetType(v string) {
	o.Type = &v
}

func (o TamAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TamAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AffectedObjectType) {
		toSerialize["AffectedObjectType"] = o.AffectedObjectType
	}
	if !IsNil(o.AlertType) {
		toSerialize["AlertType"] = o.AlertType
	}
	if o.Identifiers != nil {
		toSerialize["Identifiers"] = o.Identifiers
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.OperationType) {
		toSerialize["OperationType"] = o.OperationType
	}
	if o.Queries != nil {
		toSerialize["Queries"] = o.Queries
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TamAction) UnmarshalJSON(data []byte) (err error) {
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
	type TamActionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove).
		AffectedObjectType *string `json:"AffectedObjectType,omitempty"`
		// Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.). * `psirt` - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x). * `fieldNotice` - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html). * `eolAdvisory` - Represents product End of Life (EOL) type (https://www.cisco.com/c/en/us/products/eos-eol-policy.html).
		AlertType   *string          `json:"AlertType,omitempty"`
		Identifiers []TamIdentifiers `json:"Identifiers,omitempty"`
		// Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries.
		Name *string `json:"Name,omitempty"`
		// Operation type for the alert action. An action is used to carry out the process of \"reacting\" to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied. * `create` - Create an instance of AdvisoryInstance. * `remove` - Remove an instance of AdvisoryInstance.
		OperationType *string         `json:"OperationType,omitempty"`
		Queries       []TamQueryEntry `json:"Queries,omitempty"`
		// Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type. * `restApi` - Repesents the use of REST API for carrying out alert actions.
		Type *string `json:"Type,omitempty"`
	}

	varTamActionWithoutEmbeddedStruct := TamActionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varTamActionWithoutEmbeddedStruct)
	if err == nil {
		varTamAction := _TamAction{}
		varTamAction.ClassId = varTamActionWithoutEmbeddedStruct.ClassId
		varTamAction.ObjectType = varTamActionWithoutEmbeddedStruct.ObjectType
		varTamAction.AffectedObjectType = varTamActionWithoutEmbeddedStruct.AffectedObjectType
		varTamAction.AlertType = varTamActionWithoutEmbeddedStruct.AlertType
		varTamAction.Identifiers = varTamActionWithoutEmbeddedStruct.Identifiers
		varTamAction.Name = varTamActionWithoutEmbeddedStruct.Name
		varTamAction.OperationType = varTamActionWithoutEmbeddedStruct.OperationType
		varTamAction.Queries = varTamActionWithoutEmbeddedStruct.Queries
		varTamAction.Type = varTamActionWithoutEmbeddedStruct.Type
		*o = TamAction(varTamAction)
	} else {
		return err
	}

	varTamAction := _TamAction{}

	err = json.Unmarshal(data, &varTamAction)
	if err == nil {
		o.MoBaseComplexType = varTamAction.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AffectedObjectType")
		delete(additionalProperties, "AlertType")
		delete(additionalProperties, "Identifiers")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperationType")
		delete(additionalProperties, "Queries")
		delete(additionalProperties, "Type")

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

type NullableTamAction struct {
	value *TamAction
	isSet bool
}

func (v NullableTamAction) Get() *TamAction {
	return v.value
}

func (v *NullableTamAction) Set(val *TamAction) {
	v.value = val
	v.isSet = true
}

func (v NullableTamAction) IsSet() bool {
	return v.isSet
}

func (v *NullableTamAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamAction(val *TamAction) *NullableTamAction {
	return &NullableTamAction{value: val, isSet: true}
}

func (v NullableTamAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
