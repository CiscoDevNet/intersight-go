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
)

// checks if the TelemetryDruidScanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidScanRequest{}

// TelemetryDruidScanRequest The Scan query returns raw Apache Druid rows in streaming mode. In addition to straightforward usage where a Scan query is issued to the Broker, the Scan query can also be issued directly to Historical processes or streaming ingestion tasks. This can be useful if you want to retrieve large amounts of data in parallel.
type TelemetryDruidScanRequest struct {
	QueryType  string                   `json:"queryType"`
	DataSource TelemetryDruidDataSource `json:"dataSource"`
	// A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over.
	Intervals []string `json:"intervals"`
	// How the results are represented, list or compactedList.
	ResultFormat *string               `json:"resultFormat,omitempty"`
	Filter       *TelemetryDruidFilter `json:"filter,omitempty"`
	// A String array of dimensions and metrics to scan. If left empty, all dimensions and metrics are returned.
	Columns []string `json:"columns,omitempty"`
	// The maximum number of rows buffered before being returned to the client.
	BatchSize *int32 `json:"batchSize,omitempty"`
	// How many rows to return. If not specified, all rows will be returned.
	Limit *int32 `json:"limit,omitempty"`
	// Skip this many rows when returning results. Skipped rows will still need to be generated internally and then discarded, meaning that raising offsets to high values can cause queries to use additional resources. Together, \"limit\" and \"offset\" can be used to implement pagination. However, note that if the underlying datasource is modified in between page fetches in ways that affect overall query results, then the different pages will not necessarily align with each other.
	Offset *int32 `json:"offset,omitempty"`
	// The ordering of returned rows based on timestamp. \"ascending\", \"descending\", and \"none\" (default) are supported. Currently, \"ascending\" and \"descending\" are only supported for queries where the __time column is included in the columns field and the requirements outlined in the time ordering section are met.
	Order *string `json:"order,omitempty"`
	// Return results consistent with the legacy \"scan-query\" contrib extension. Defaults to the value set by druid.query.scan.legacy, which in turn defaults to false.
	Legacy  *bool                       `json:"legacy,omitempty"`
	Context *TelemetryDruidQueryContext `json:"context,omitempty"`
	// Virtual columns are columns that are computed on the fly during query execution. It can potentially draw from multiple underlying columns, although a virtual column always presents itself as a single column. Virtual columns can be referenced by their output names to be used as dimensions or as inputs to filters and aggregators.
	VirtualColumns       []TelemetryDruidVirtualColumn `json:"virtualColumns,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidScanRequest TelemetryDruidScanRequest

// NewTelemetryDruidScanRequest instantiates a new TelemetryDruidScanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidScanRequest(queryType string, dataSource TelemetryDruidDataSource, intervals []string) *TelemetryDruidScanRequest {
	this := TelemetryDruidScanRequest{}
	this.QueryType = queryType
	this.DataSource = dataSource
	this.Intervals = intervals
	var resultFormat string = "list"
	this.ResultFormat = &resultFormat
	var batchSize int32 = 20480
	this.BatchSize = &batchSize
	var order string = "none"
	this.Order = &order
	var legacy bool = false
	this.Legacy = &legacy
	return &this
}

// NewTelemetryDruidScanRequestWithDefaults instantiates a new TelemetryDruidScanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidScanRequestWithDefaults() *TelemetryDruidScanRequest {
	this := TelemetryDruidScanRequest{}
	var resultFormat string = "list"
	this.ResultFormat = &resultFormat
	var batchSize int32 = 20480
	this.BatchSize = &batchSize
	var order string = "none"
	this.Order = &order
	var legacy bool = false
	this.Legacy = &legacy
	return &this
}

// GetQueryType returns the QueryType field value
func (o *TelemetryDruidScanRequest) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *TelemetryDruidScanRequest) SetQueryType(v string) {
	o.QueryType = v
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidScanRequest) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidScanRequest) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetIntervals returns the Intervals field value
func (o *TelemetryDruidScanRequest) GetIntervals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetIntervalsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Intervals, true
}

// SetIntervals sets field value
func (o *TelemetryDruidScanRequest) SetIntervals(v []string) {
	o.Intervals = v
}

// GetResultFormat returns the ResultFormat field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetResultFormat() string {
	if o == nil || IsNil(o.ResultFormat) {
		var ret string
		return ret
	}
	return *o.ResultFormat
}

// GetResultFormatOk returns a tuple with the ResultFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetResultFormatOk() (*string, bool) {
	if o == nil || IsNil(o.ResultFormat) {
		return nil, false
	}
	return o.ResultFormat, true
}

// HasResultFormat returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasResultFormat() bool {
	if o != nil && !IsNil(o.ResultFormat) {
		return true
	}

	return false
}

// SetResultFormat gets a reference to the given string and assigns it to the ResultFormat field.
func (o *TelemetryDruidScanRequest) SetResultFormat(v string) {
	o.ResultFormat = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetFilter() TelemetryDruidFilter {
	if o == nil || IsNil(o.Filter) {
		var ret TelemetryDruidFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given TelemetryDruidFilter and assigns it to the Filter field.
func (o *TelemetryDruidScanRequest) SetFilter(v TelemetryDruidFilter) {
	o.Filter = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetColumns() []string {
	if o == nil || IsNil(o.Columns) {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *TelemetryDruidScanRequest) SetColumns(v []string) {
	o.Columns = v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetBatchSize() int32 {
	if o == nil || IsNil(o.BatchSize) {
		var ret int32
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetBatchSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.BatchSize) {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasBatchSize() bool {
	if o != nil && !IsNil(o.BatchSize) {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given int32 and assigns it to the BatchSize field.
func (o *TelemetryDruidScanRequest) SetBatchSize(v int32) {
	o.BatchSize = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *TelemetryDruidScanRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *TelemetryDruidScanRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *TelemetryDruidScanRequest) SetOrder(v string) {
	o.Order = &v
}

// GetLegacy returns the Legacy field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetLegacy() bool {
	if o == nil || IsNil(o.Legacy) {
		var ret bool
		return ret
	}
	return *o.Legacy
}

// GetLegacyOk returns a tuple with the Legacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetLegacyOk() (*bool, bool) {
	if o == nil || IsNil(o.Legacy) {
		return nil, false
	}
	return o.Legacy, true
}

// HasLegacy returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasLegacy() bool {
	if o != nil && !IsNil(o.Legacy) {
		return true
	}

	return false
}

// SetLegacy gets a reference to the given bool and assigns it to the Legacy field.
func (o *TelemetryDruidScanRequest) SetLegacy(v bool) {
	o.Legacy = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetContext() TelemetryDruidQueryContext {
	if o == nil || IsNil(o.Context) {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidScanRequest) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

// GetVirtualColumns returns the VirtualColumns field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetVirtualColumns() []TelemetryDruidVirtualColumn {
	if o == nil || IsNil(o.VirtualColumns) {
		var ret []TelemetryDruidVirtualColumn
		return ret
	}
	return o.VirtualColumns
}

// GetVirtualColumnsOk returns a tuple with the VirtualColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetVirtualColumnsOk() ([]TelemetryDruidVirtualColumn, bool) {
	if o == nil || IsNil(o.VirtualColumns) {
		return nil, false
	}
	return o.VirtualColumns, true
}

// HasVirtualColumns returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasVirtualColumns() bool {
	if o != nil && !IsNil(o.VirtualColumns) {
		return true
	}

	return false
}

// SetVirtualColumns gets a reference to the given []TelemetryDruidVirtualColumn and assigns it to the VirtualColumns field.
func (o *TelemetryDruidScanRequest) SetVirtualColumns(v []TelemetryDruidVirtualColumn) {
	o.VirtualColumns = v
}

func (o TelemetryDruidScanRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidScanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryType"] = o.QueryType
	toSerialize["dataSource"] = o.DataSource
	toSerialize["intervals"] = o.Intervals
	if !IsNil(o.ResultFormat) {
		toSerialize["resultFormat"] = o.ResultFormat
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	if !IsNil(o.BatchSize) {
		toSerialize["batchSize"] = o.BatchSize
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Legacy) {
		toSerialize["legacy"] = o.Legacy
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.VirtualColumns) {
		toSerialize["virtualColumns"] = o.VirtualColumns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidScanRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryType",
		"dataSource",
		"intervals",
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
	varTelemetryDruidScanRequest := _TelemetryDruidScanRequest{}

	err = json.Unmarshal(data, &varTelemetryDruidScanRequest)

	if err != nil {
		return err
	}

	*o = TelemetryDruidScanRequest(varTelemetryDruidScanRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "queryType")
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "intervals")
		delete(additionalProperties, "resultFormat")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "columns")
		delete(additionalProperties, "batchSize")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "order")
		delete(additionalProperties, "legacy")
		delete(additionalProperties, "context")
		delete(additionalProperties, "virtualColumns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidScanRequest struct {
	value *TelemetryDruidScanRequest
	isSet bool
}

func (v NullableTelemetryDruidScanRequest) Get() *TelemetryDruidScanRequest {
	return v.value
}

func (v *NullableTelemetryDruidScanRequest) Set(val *TelemetryDruidScanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidScanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidScanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidScanRequest(val *TelemetryDruidScanRequest) *NullableTelemetryDruidScanRequest {
	return &NullableTelemetryDruidScanRequest{value: val, isSet: true}
}

func (v NullableTelemetryDruidScanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidScanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
