/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-17769
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

// checks if the SchedulerTaskResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulerTaskResult{}

// SchedulerTaskResult An instance of a task execution result. Each Recurring schedule will have an instance of this object in the database. A one-time schedule will have just one entry.
type SchedulerTaskResult struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// End time for the scheduled task.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Start time for the scheduled task.
	StartTime    *time.Time                        `json:"StartTime,omitempty"`
	StatusDetail NullableSchedulerTaskResultStatus `json:"StatusDetail,omitempty"`
	// The timezone for the startTime specified. * `Pacific/Niue` -  * `Pacific/Pago_Pago` -  * `Pacific/Honolulu` -  * `Pacific/Rarotonga` -  * `Pacific/Tahiti` -  * `Pacific/Marquesas` -  * `America/Anchorage` -  * `Pacific/Gambier` -  * `America/Los_Angeles` -  * `America/Tijuana` -  * `America/Vancouver` -  * `America/Whitehorse` -  * `Pacific/Pitcairn` -  * `America/Dawson_Creek` -  * `America/Denver` -  * `America/Edmonton` -  * `America/Hermosillo` -  * `America/Mazatlan` -  * `America/Phoenix` -  * `America/Yellowknife` -  * `America/Belize` -  * `America/Chicago` -  * `America/Costa_Rica` -  * `America/El_Salvador` -  * `America/Guatemala` -  * `America/Managua` -  * `America/Mexico_City` -  * `America/Regina` -  * `America/Tegucigalpa` -  * `America/Winnipeg` -  * `Pacific/Galapagos` -  * `America/Bogota` -  * `America/Cancun` -  * `America/Cayman` -  * `America/Guayaquil` -  * `America/Havana` -  * `America/Iqaluit` -  * `America/Jamaica` -  * `America/Lima` -  * `America/Nassau` -  * `America/New_York` -  * `America/Nuuk` -  * `America/Panama` -  * `America/Port-au-Prince` -  * `America/Rio_Branco` -  * `America/Toronto` -  * `Pacific/Easter` -  * `America/Caracas` -  * `America/Asuncion` -  * `America/Barbados` -  * `America/Boa_Vista` -  * `America/Campo_Grande` -  * `America/Cuiaba` -  * `America/Curacao` -  * `America/Grand_Turk` -  * `America/Guyana` -  * `America/Halifax` -  * `America/La_Paz` -  * `America/Manaus` -  * `America/Martinique` -  * `America/Port_of_Spain` -  * `America/Porto_Velho` -  * `America/Puerto_Rico` -  * `America/Santo_Domingo` -  * `America/Thule` -  * `Atlantic/Bermuda` -  * `America/St_Johns` -  * `America/Araguaina` -  * `America/Argentina/Buenos_Aires` -  * `America/Bahia` -  * `America/Belem` -  * `America/Cayenne` -  * `America/Fortaleza` -  * `America/Godthab` -  * `America/Maceio` -  * `America/Miquelon` -  * `America/Montevideo` -  * `America/Paramaribo` -  * `America/Recife` -  * `America/Santiago` -  * `America/Sao_Paulo` -  * `Antarctica/Palmer` -  * `Antarctica/Rothera` -  * `Atlantic/Stanley` -  * `America/Noronha` -  * `Atlantic/South_Georgia` -  * `America/Scoresbysund` -  * `Atlantic/Azores` -  * `Atlantic/Cape_Verde` -  * `Africa/Abidjan` -  * `Africa/Accra` -  * `Africa/Bissau` -  * `Africa/Casablanca` -  * `Africa/El_Aaiun` -  * `Africa/Monrovia` -  * `America/Danmarkshavn` -  * `Atlantic/Canary` -  * `Atlantic/Faroe` -  * `Atlantic/Reykjavik` -  * `Etc/GMT` -  * `Europe/Dublin` -  * `Europe/Lisbon` -  * `Europe/London` -  * `Africa/Algiers` -  * `Africa/Ceuta` -  * `Africa/Lagos` -  * `Africa/Ndjamena` -  * `Africa/Tunis` -  * `Africa/Windhoek` -  * `Europe/Amsterdam` -  * `Europe/Andorra` -  * `Europe/Belgrade` -  * `Europe/Berlin` -  * `Europe/Brussels` -  * `Europe/Budapest` -  * `Europe/Copenhagen` -  * `Europe/Gibraltar` -  * `Europe/Luxembourg` -  * `Europe/Madrid` -  * `Europe/Malta` -  * `Europe/Monaco` -  * `Europe/Oslo` -  * `Europe/Paris` -  * `Europe/Prague` -  * `Europe/Rome` -  * `Europe/Stockholm` -  * `Europe/Tirane` -  * `Europe/Vienna` -  * `Europe/Warsaw` -  * `Europe/Zurich` -  * `Africa/Cairo` -  * `Africa/Johannesburg` -  * `Africa/Maputo` -  * `Africa/Tripoli` -  * `Asia/Amman` -  * `Asia/Beirut` -  * `Asia/Damascus` -  * `Asia/Gaza` -  * `Asia/Jerusalem` -  * `Asia/Nicosia` -  * `Europe/Athens` -  * `Europe/Bucharest` -  * `Europe/Chisinau` -  * `Europe/Helsinki` -  * `Europe/Istanbul` -  * `Europe/Kaliningrad` -  * `Europe/Kiev` -  * `Europe/Riga` -  * `Europe/Sofia` -  * `Europe/Tallinn` -  * `Europe/Vilnius` -  * `Africa/Khartoum` -  * `Africa/Nairobi` -  * `Antarctica/Syowa` -  * `Asia/Baghdad` -  * `Asia/Qatar` -  * `Asia/Riyadh` -  * `Europe/Minsk` -  * `Europe/Moscow` -  * `Asia/Tehran` -  * `Asia/Baku` -  * `Asia/Dubai` -  * `Asia/Tbilisi` -  * `Asia/Yerevan` -  * `Europe/Samara` -  * `Indian/Mahe` -  * `Indian/Mauritius` -  * `Indian/Reunion` -  * `Asia/Kabul` -  * `Antarctica/Mawson` -  * `Asia/Aqtau` -  * `Asia/Aqtobe` -  * `Asia/Ashgabat` -  * `Asia/Dushanbe` -  * `Asia/Karachi` -  * `Asia/Tashkent` -  * `Asia/Yekaterinburg` -  * `Indian/Kerguelen` -  * `Indian/Maldives` -  * `Asia/Calcutta` -  * `Asia/Kolkata` -  * `Asia/Colombo` -  * `Asia/Kathmandu` -  * `Asia/Katmandu` -  * `Antarctica/Vostok` -  * `Asia/Almaty` -  * `Asia/Bishkek` -  * `Asia/Dhaka` -  * `Asia/Omsk` -  * `Asia/Thimphu` -  * `Indian/Chagos` -  * `Asia/Rangoon` -  * `Indian/Cocos` -  * `Antarctica/Davis` -  * `Asia/Bangkok` -  * `Asia/Ho_Chi_Minh` -  * `Asia/Hovd` -  * `Asia/Jakarta` -  * `Asia/Krasnoyarsk` -  * `Asia/Saigon` -  * `Indian/Christmas` -  * `Antarctica/Casey` -  * `Asia/Brunei` -  * `Asia/Choibalsan` -  * `Asia/Hong_Kong` -  * `Asia/Irkutsk` -  * `Asia/Kuala_Lumpur` -  * `Asia/Macau` -  * `Asia/Makassar` -  * `Asia/Manila` -  * `Asia/Shanghai` -  * `Asia/Singapore` -  * `Asia/Taipei` -  * `Asia/Ulaanbaatar` -  * `Australia/Perth` -  * `Asia/Pyongyang` -  * `Asia/Dili` -  * `Asia/Jayapura` -  * `Asia/Seoul` -  * `Asia/Tokyo` -  * `Asia/Yakutsk` -  * `Asia/Yangon` -  * `Pacific/Palau` -  * `Australia/Adelaide` -  * `Australia/Darwin` -  * `Antarctica/DumontDUrville` -  * `Asia/Magadan` -  * `Asia/Vladivostok` -  * `Australia/Brisbane` -  * `Australia/Hobart` -  * `Australia/Sydney` -  * `Pacific/Chuuk` -  * `Pacific/Guam` -  * `Pacific/Port_Moresby` -  * `Pacific/Efate` -  * `Pacific/Guadalcanal` -  * `Pacific/Kosrae` -  * `Pacific/Norfolk` -  * `Pacific/Noumea` -  * `Pacific/Pohnpei` -  * `Asia/Kamchatka` -  * `Pacific/Auckland` -  * `Pacific/Fiji` -  * `Pacific/Funafuti` -  * `Pacific/Kwajalein` -  * `Pacific/Majuro` -  * `Pacific/Nauru` -  * `Pacific/Tarawa` -  * `Pacific/Wake` -  * `Pacific/Wallis` -  * `Pacific/Apia` -  * `Pacific/Enderbury` -  * `Pacific/Fakaofo` -  * `Pacific/Tongatapu` -  * `Pacific/Kiritimati` -  * `UTC` -
	TimeZone             *string                                   `json:"TimeZone,omitempty"`
	AssociatedObject     NullableMoBaseMoRelationship              `json:"AssociatedObject,omitempty"`
	TaskSchedule         NullableSchedulerTaskScheduleRelationship `json:"TaskSchedule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchedulerTaskResult SchedulerTaskResult

// NewSchedulerTaskResult instantiates a new SchedulerTaskResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulerTaskResult(classId string, objectType string) *SchedulerTaskResult {
	this := SchedulerTaskResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSchedulerTaskResultWithDefaults instantiates a new SchedulerTaskResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulerTaskResultWithDefaults() *SchedulerTaskResult {
	this := SchedulerTaskResult{}
	var classId string = "scheduler.TaskResult"
	this.ClassId = classId
	var objectType string = "scheduler.TaskResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SchedulerTaskResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SchedulerTaskResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SchedulerTaskResult) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SchedulerTaskResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SchedulerTaskResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SchedulerTaskResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *SchedulerTaskResult) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerTaskResult) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *SchedulerTaskResult) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *SchedulerTaskResult) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *SchedulerTaskResult) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerTaskResult) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *SchedulerTaskResult) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *SchedulerTaskResult) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatusDetail returns the StatusDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulerTaskResult) GetStatusDetail() SchedulerTaskResultStatus {
	if o == nil || IsNil(o.StatusDetail.Get()) {
		var ret SchedulerTaskResultStatus
		return ret
	}
	return *o.StatusDetail.Get()
}

// GetStatusDetailOk returns a tuple with the StatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulerTaskResult) GetStatusDetailOk() (*SchedulerTaskResultStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusDetail.Get(), o.StatusDetail.IsSet()
}

// HasStatusDetail returns a boolean if a field has been set.
func (o *SchedulerTaskResult) HasStatusDetail() bool {
	if o != nil && o.StatusDetail.IsSet() {
		return true
	}

	return false
}

// SetStatusDetail gets a reference to the given NullableSchedulerTaskResultStatus and assigns it to the StatusDetail field.
func (o *SchedulerTaskResult) SetStatusDetail(v SchedulerTaskResultStatus) {
	o.StatusDetail.Set(&v)
}

// SetStatusDetailNil sets the value for StatusDetail to be an explicit nil
func (o *SchedulerTaskResult) SetStatusDetailNil() {
	o.StatusDetail.Set(nil)
}

// UnsetStatusDetail ensures that no value is present for StatusDetail, not even an explicit nil
func (o *SchedulerTaskResult) UnsetStatusDetail() {
	o.StatusDetail.Unset()
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *SchedulerTaskResult) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerTaskResult) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *SchedulerTaskResult) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *SchedulerTaskResult) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetAssociatedObject returns the AssociatedObject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulerTaskResult) GetAssociatedObject() MoBaseMoRelationship {
	if o == nil || IsNil(o.AssociatedObject.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssociatedObject.Get()
}

// GetAssociatedObjectOk returns a tuple with the AssociatedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulerTaskResult) GetAssociatedObjectOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssociatedObject.Get(), o.AssociatedObject.IsSet()
}

// HasAssociatedObject returns a boolean if a field has been set.
func (o *SchedulerTaskResult) HasAssociatedObject() bool {
	if o != nil && o.AssociatedObject.IsSet() {
		return true
	}

	return false
}

// SetAssociatedObject gets a reference to the given NullableMoBaseMoRelationship and assigns it to the AssociatedObject field.
func (o *SchedulerTaskResult) SetAssociatedObject(v MoBaseMoRelationship) {
	o.AssociatedObject.Set(&v)
}

// SetAssociatedObjectNil sets the value for AssociatedObject to be an explicit nil
func (o *SchedulerTaskResult) SetAssociatedObjectNil() {
	o.AssociatedObject.Set(nil)
}

// UnsetAssociatedObject ensures that no value is present for AssociatedObject, not even an explicit nil
func (o *SchedulerTaskResult) UnsetAssociatedObject() {
	o.AssociatedObject.Unset()
}

// GetTaskSchedule returns the TaskSchedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulerTaskResult) GetTaskSchedule() SchedulerTaskScheduleRelationship {
	if o == nil || IsNil(o.TaskSchedule.Get()) {
		var ret SchedulerTaskScheduleRelationship
		return ret
	}
	return *o.TaskSchedule.Get()
}

// GetTaskScheduleOk returns a tuple with the TaskSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulerTaskResult) GetTaskScheduleOk() (*SchedulerTaskScheduleRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskSchedule.Get(), o.TaskSchedule.IsSet()
}

// HasTaskSchedule returns a boolean if a field has been set.
func (o *SchedulerTaskResult) HasTaskSchedule() bool {
	if o != nil && o.TaskSchedule.IsSet() {
		return true
	}

	return false
}

// SetTaskSchedule gets a reference to the given NullableSchedulerTaskScheduleRelationship and assigns it to the TaskSchedule field.
func (o *SchedulerTaskResult) SetTaskSchedule(v SchedulerTaskScheduleRelationship) {
	o.TaskSchedule.Set(&v)
}

// SetTaskScheduleNil sets the value for TaskSchedule to be an explicit nil
func (o *SchedulerTaskResult) SetTaskScheduleNil() {
	o.TaskSchedule.Set(nil)
}

// UnsetTaskSchedule ensures that no value is present for TaskSchedule, not even an explicit nil
func (o *SchedulerTaskResult) UnsetTaskSchedule() {
	o.TaskSchedule.Unset()
}

func (o SchedulerTaskResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulerTaskResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.EndTime) {
		toSerialize["EndTime"] = o.EndTime
	}
	if !IsNil(o.StartTime) {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.StatusDetail.IsSet() {
		toSerialize["StatusDetail"] = o.StatusDetail.Get()
	}
	if !IsNil(o.TimeZone) {
		toSerialize["TimeZone"] = o.TimeZone
	}
	if o.AssociatedObject.IsSet() {
		toSerialize["AssociatedObject"] = o.AssociatedObject.Get()
	}
	if o.TaskSchedule.IsSet() {
		toSerialize["TaskSchedule"] = o.TaskSchedule.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchedulerTaskResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type SchedulerTaskResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// End time for the scheduled task.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// Start time for the scheduled task.
		StartTime    *time.Time                        `json:"StartTime,omitempty"`
		StatusDetail NullableSchedulerTaskResultStatus `json:"StatusDetail,omitempty"`
		// The timezone for the startTime specified. * `Pacific/Niue` -  * `Pacific/Pago_Pago` -  * `Pacific/Honolulu` -  * `Pacific/Rarotonga` -  * `Pacific/Tahiti` -  * `Pacific/Marquesas` -  * `America/Anchorage` -  * `Pacific/Gambier` -  * `America/Los_Angeles` -  * `America/Tijuana` -  * `America/Vancouver` -  * `America/Whitehorse` -  * `Pacific/Pitcairn` -  * `America/Dawson_Creek` -  * `America/Denver` -  * `America/Edmonton` -  * `America/Hermosillo` -  * `America/Mazatlan` -  * `America/Phoenix` -  * `America/Yellowknife` -  * `America/Belize` -  * `America/Chicago` -  * `America/Costa_Rica` -  * `America/El_Salvador` -  * `America/Guatemala` -  * `America/Managua` -  * `America/Mexico_City` -  * `America/Regina` -  * `America/Tegucigalpa` -  * `America/Winnipeg` -  * `Pacific/Galapagos` -  * `America/Bogota` -  * `America/Cancun` -  * `America/Cayman` -  * `America/Guayaquil` -  * `America/Havana` -  * `America/Iqaluit` -  * `America/Jamaica` -  * `America/Lima` -  * `America/Nassau` -  * `America/New_York` -  * `America/Nuuk` -  * `America/Panama` -  * `America/Port-au-Prince` -  * `America/Rio_Branco` -  * `America/Toronto` -  * `Pacific/Easter` -  * `America/Caracas` -  * `America/Asuncion` -  * `America/Barbados` -  * `America/Boa_Vista` -  * `America/Campo_Grande` -  * `America/Cuiaba` -  * `America/Curacao` -  * `America/Grand_Turk` -  * `America/Guyana` -  * `America/Halifax` -  * `America/La_Paz` -  * `America/Manaus` -  * `America/Martinique` -  * `America/Port_of_Spain` -  * `America/Porto_Velho` -  * `America/Puerto_Rico` -  * `America/Santo_Domingo` -  * `America/Thule` -  * `Atlantic/Bermuda` -  * `America/St_Johns` -  * `America/Araguaina` -  * `America/Argentina/Buenos_Aires` -  * `America/Bahia` -  * `America/Belem` -  * `America/Cayenne` -  * `America/Fortaleza` -  * `America/Godthab` -  * `America/Maceio` -  * `America/Miquelon` -  * `America/Montevideo` -  * `America/Paramaribo` -  * `America/Recife` -  * `America/Santiago` -  * `America/Sao_Paulo` -  * `Antarctica/Palmer` -  * `Antarctica/Rothera` -  * `Atlantic/Stanley` -  * `America/Noronha` -  * `Atlantic/South_Georgia` -  * `America/Scoresbysund` -  * `Atlantic/Azores` -  * `Atlantic/Cape_Verde` -  * `Africa/Abidjan` -  * `Africa/Accra` -  * `Africa/Bissau` -  * `Africa/Casablanca` -  * `Africa/El_Aaiun` -  * `Africa/Monrovia` -  * `America/Danmarkshavn` -  * `Atlantic/Canary` -  * `Atlantic/Faroe` -  * `Atlantic/Reykjavik` -  * `Etc/GMT` -  * `Europe/Dublin` -  * `Europe/Lisbon` -  * `Europe/London` -  * `Africa/Algiers` -  * `Africa/Ceuta` -  * `Africa/Lagos` -  * `Africa/Ndjamena` -  * `Africa/Tunis` -  * `Africa/Windhoek` -  * `Europe/Amsterdam` -  * `Europe/Andorra` -  * `Europe/Belgrade` -  * `Europe/Berlin` -  * `Europe/Brussels` -  * `Europe/Budapest` -  * `Europe/Copenhagen` -  * `Europe/Gibraltar` -  * `Europe/Luxembourg` -  * `Europe/Madrid` -  * `Europe/Malta` -  * `Europe/Monaco` -  * `Europe/Oslo` -  * `Europe/Paris` -  * `Europe/Prague` -  * `Europe/Rome` -  * `Europe/Stockholm` -  * `Europe/Tirane` -  * `Europe/Vienna` -  * `Europe/Warsaw` -  * `Europe/Zurich` -  * `Africa/Cairo` -  * `Africa/Johannesburg` -  * `Africa/Maputo` -  * `Africa/Tripoli` -  * `Asia/Amman` -  * `Asia/Beirut` -  * `Asia/Damascus` -  * `Asia/Gaza` -  * `Asia/Jerusalem` -  * `Asia/Nicosia` -  * `Europe/Athens` -  * `Europe/Bucharest` -  * `Europe/Chisinau` -  * `Europe/Helsinki` -  * `Europe/Istanbul` -  * `Europe/Kaliningrad` -  * `Europe/Kiev` -  * `Europe/Riga` -  * `Europe/Sofia` -  * `Europe/Tallinn` -  * `Europe/Vilnius` -  * `Africa/Khartoum` -  * `Africa/Nairobi` -  * `Antarctica/Syowa` -  * `Asia/Baghdad` -  * `Asia/Qatar` -  * `Asia/Riyadh` -  * `Europe/Minsk` -  * `Europe/Moscow` -  * `Asia/Tehran` -  * `Asia/Baku` -  * `Asia/Dubai` -  * `Asia/Tbilisi` -  * `Asia/Yerevan` -  * `Europe/Samara` -  * `Indian/Mahe` -  * `Indian/Mauritius` -  * `Indian/Reunion` -  * `Asia/Kabul` -  * `Antarctica/Mawson` -  * `Asia/Aqtau` -  * `Asia/Aqtobe` -  * `Asia/Ashgabat` -  * `Asia/Dushanbe` -  * `Asia/Karachi` -  * `Asia/Tashkent` -  * `Asia/Yekaterinburg` -  * `Indian/Kerguelen` -  * `Indian/Maldives` -  * `Asia/Calcutta` -  * `Asia/Kolkata` -  * `Asia/Colombo` -  * `Asia/Kathmandu` -  * `Asia/Katmandu` -  * `Antarctica/Vostok` -  * `Asia/Almaty` -  * `Asia/Bishkek` -  * `Asia/Dhaka` -  * `Asia/Omsk` -  * `Asia/Thimphu` -  * `Indian/Chagos` -  * `Asia/Rangoon` -  * `Indian/Cocos` -  * `Antarctica/Davis` -  * `Asia/Bangkok` -  * `Asia/Ho_Chi_Minh` -  * `Asia/Hovd` -  * `Asia/Jakarta` -  * `Asia/Krasnoyarsk` -  * `Asia/Saigon` -  * `Indian/Christmas` -  * `Antarctica/Casey` -  * `Asia/Brunei` -  * `Asia/Choibalsan` -  * `Asia/Hong_Kong` -  * `Asia/Irkutsk` -  * `Asia/Kuala_Lumpur` -  * `Asia/Macau` -  * `Asia/Makassar` -  * `Asia/Manila` -  * `Asia/Shanghai` -  * `Asia/Singapore` -  * `Asia/Taipei` -  * `Asia/Ulaanbaatar` -  * `Australia/Perth` -  * `Asia/Pyongyang` -  * `Asia/Dili` -  * `Asia/Jayapura` -  * `Asia/Seoul` -  * `Asia/Tokyo` -  * `Asia/Yakutsk` -  * `Asia/Yangon` -  * `Pacific/Palau` -  * `Australia/Adelaide` -  * `Australia/Darwin` -  * `Antarctica/DumontDUrville` -  * `Asia/Magadan` -  * `Asia/Vladivostok` -  * `Australia/Brisbane` -  * `Australia/Hobart` -  * `Australia/Sydney` -  * `Pacific/Chuuk` -  * `Pacific/Guam` -  * `Pacific/Port_Moresby` -  * `Pacific/Efate` -  * `Pacific/Guadalcanal` -  * `Pacific/Kosrae` -  * `Pacific/Norfolk` -  * `Pacific/Noumea` -  * `Pacific/Pohnpei` -  * `Asia/Kamchatka` -  * `Pacific/Auckland` -  * `Pacific/Fiji` -  * `Pacific/Funafuti` -  * `Pacific/Kwajalein` -  * `Pacific/Majuro` -  * `Pacific/Nauru` -  * `Pacific/Tarawa` -  * `Pacific/Wake` -  * `Pacific/Wallis` -  * `Pacific/Apia` -  * `Pacific/Enderbury` -  * `Pacific/Fakaofo` -  * `Pacific/Tongatapu` -  * `Pacific/Kiritimati` -  * `UTC` -
		TimeZone         *string                                   `json:"TimeZone,omitempty"`
		AssociatedObject NullableMoBaseMoRelationship              `json:"AssociatedObject,omitempty"`
		TaskSchedule     NullableSchedulerTaskScheduleRelationship `json:"TaskSchedule,omitempty"`
	}

	varSchedulerTaskResultWithoutEmbeddedStruct := SchedulerTaskResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSchedulerTaskResultWithoutEmbeddedStruct)
	if err == nil {
		varSchedulerTaskResult := _SchedulerTaskResult{}
		varSchedulerTaskResult.ClassId = varSchedulerTaskResultWithoutEmbeddedStruct.ClassId
		varSchedulerTaskResult.ObjectType = varSchedulerTaskResultWithoutEmbeddedStruct.ObjectType
		varSchedulerTaskResult.EndTime = varSchedulerTaskResultWithoutEmbeddedStruct.EndTime
		varSchedulerTaskResult.StartTime = varSchedulerTaskResultWithoutEmbeddedStruct.StartTime
		varSchedulerTaskResult.StatusDetail = varSchedulerTaskResultWithoutEmbeddedStruct.StatusDetail
		varSchedulerTaskResult.TimeZone = varSchedulerTaskResultWithoutEmbeddedStruct.TimeZone
		varSchedulerTaskResult.AssociatedObject = varSchedulerTaskResultWithoutEmbeddedStruct.AssociatedObject
		varSchedulerTaskResult.TaskSchedule = varSchedulerTaskResultWithoutEmbeddedStruct.TaskSchedule
		*o = SchedulerTaskResult(varSchedulerTaskResult)
	} else {
		return err
	}

	varSchedulerTaskResult := _SchedulerTaskResult{}

	err = json.Unmarshal(data, &varSchedulerTaskResult)
	if err == nil {
		o.MoBaseMo = varSchedulerTaskResult.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "StatusDetail")
		delete(additionalProperties, "TimeZone")
		delete(additionalProperties, "AssociatedObject")
		delete(additionalProperties, "TaskSchedule")

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

type NullableSchedulerTaskResult struct {
	value *SchedulerTaskResult
	isSet bool
}

func (v NullableSchedulerTaskResult) Get() *SchedulerTaskResult {
	return v.value
}

func (v *NullableSchedulerTaskResult) Set(val *SchedulerTaskResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulerTaskResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulerTaskResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulerTaskResult(val *SchedulerTaskResult) *NullableSchedulerTaskResult {
	return &NullableSchedulerTaskResult{value: val, isSet: true}
}

func (v NullableSchedulerTaskResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulerTaskResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}