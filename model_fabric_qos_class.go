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

// checks if the FabricQosClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricQosClass{}

// FabricQosClass Type to represent the Best Effort QoS class.
type FabricQosClass struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Administrative state for this QoS class. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
	AdminState *string `json:"AdminState,omitempty"`
	// Percentage of bandwidth received by the traffic tagged with this QoS.
	BandwidthPercent *int64 `json:"BandwidthPercent,omitempty"`
	// Class of service received by the traffic tagged with this QoS.
	Cos *int64 `json:"Cos,omitempty"`
	// Maximum transmission unit (MTU) is the largest size packet or frame, that can be sent in a packet- or frame-based network such as the Internet. User can select from the following: 1. Any value between 1500 and 9216 2. 'Normal' (default) mapping to a value of 1500. 3. 'FC' mapping to a value of 2240.
	Mtu *int64 `json:"Mtu,omitempty"`
	// If enabled, this QoS class will be optimized to send multiple packets.
	MulticastOptimize *bool `json:"MulticastOptimize,omitempty"`
	// The 'name' of this QoS Class. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
	Name *string `json:"Name,omitempty"`
	// If enabled, this QoS class will allow packet drops within an acceptable limit.
	PacketDrop *bool `json:"PacketDrop,omitempty"`
	// The weight of the QoS Class controls the distribution of bandwidth between QoS Classes, with the same priority after the Guarantees for the QoS Classes are reached.
	Weight               *int64 `json:"Weight,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricQosClass FabricQosClass

// NewFabricQosClass instantiates a new FabricQosClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricQosClass(classId string, objectType string) *FabricQosClass {
	this := FabricQosClass{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminState string = "Disabled"
	this.AdminState = &adminState
	var mtu int64 = 1500
	this.Mtu = &mtu
	var multicastOptimize bool = false
	this.MulticastOptimize = &multicastOptimize
	var name string = "Best Effort"
	this.Name = &name
	var packetDrop bool = true
	this.PacketDrop = &packetDrop
	return &this
}

// NewFabricQosClassWithDefaults instantiates a new FabricQosClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricQosClassWithDefaults() *FabricQosClass {
	this := FabricQosClass{}
	var classId string = "fabric.QosClass"
	this.ClassId = classId
	var objectType string = "fabric.QosClass"
	this.ObjectType = objectType
	var adminState string = "Disabled"
	this.AdminState = &adminState
	var mtu int64 = 1500
	this.Mtu = &mtu
	var multicastOptimize bool = false
	this.MulticastOptimize = &multicastOptimize
	var name string = "Best Effort"
	this.Name = &name
	var packetDrop bool = true
	this.PacketDrop = &packetDrop
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricQosClass) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricQosClass) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricQosClass) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.QosClass" of the ClassId field.
func (o *FabricQosClass) GetDefaultClassId() interface{} {
	return "fabric.QosClass"
}

// GetObjectType returns the ObjectType field value
func (o *FabricQosClass) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricQosClass) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricQosClass) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.QosClass" of the ObjectType field.
func (o *FabricQosClass) GetDefaultObjectType() interface{} {
	return "fabric.QosClass"
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *FabricQosClass) GetAdminState() string {
	if o == nil || IsNil(o.AdminState) {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClass) GetAdminStateOk() (*string, bool) {
	if o == nil || IsNil(o.AdminState) {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *FabricQosClass) HasAdminState() bool {
	if o != nil && !IsNil(o.AdminState) {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *FabricQosClass) SetAdminState(v string) {
	o.AdminState = &v
}

// GetBandwidthPercent returns the BandwidthPercent field value if set, zero value otherwise.
func (o *FabricQosClass) GetBandwidthPercent() int64 {
	if o == nil || IsNil(o.BandwidthPercent) {
		var ret int64
		return ret
	}
	return *o.BandwidthPercent
}

// GetBandwidthPercentOk returns a tuple with the BandwidthPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClass) GetBandwidthPercentOk() (*int64, bool) {
	if o == nil || IsNil(o.BandwidthPercent) {
		return nil, false
	}
	return o.BandwidthPercent, true
}

// HasBandwidthPercent returns a boolean if a field has been set.
func (o *FabricQosClass) HasBandwidthPercent() bool {
	if o != nil && !IsNil(o.BandwidthPercent) {
		return true
	}

	return false
}

// SetBandwidthPercent gets a reference to the given int64 and assigns it to the BandwidthPercent field.
func (o *FabricQosClass) SetBandwidthPercent(v int64) {
	o.BandwidthPercent = &v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *FabricQosClass) GetCos() int64 {
	if o == nil || IsNil(o.Cos) {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClass) GetCosOk() (*int64, bool) {
	if o == nil || IsNil(o.Cos) {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *FabricQosClass) HasCos() bool {
	if o != nil && !IsNil(o.Cos) {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *FabricQosClass) SetCos(v int64) {
	o.Cos = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *FabricQosClass) GetMtu() int64 {
	if o == nil || IsNil(o.Mtu) {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClass) GetMtuOk() (*int64, bool) {
	if o == nil || IsNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *FabricQosClass) HasMtu() bool {
	if o != nil && !IsNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *FabricQosClass) SetMtu(v int64) {
	o.Mtu = &v
}

// GetMulticastOptimize returns the MulticastOptimize field value if set, zero value otherwise.
func (o *FabricQosClass) GetMulticastOptimize() bool {
	if o == nil || IsNil(o.MulticastOptimize) {
		var ret bool
		return ret
	}
	return *o.MulticastOptimize
}

// GetMulticastOptimizeOk returns a tuple with the MulticastOptimize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClass) GetMulticastOptimizeOk() (*bool, bool) {
	if o == nil || IsNil(o.MulticastOptimize) {
		return nil, false
	}
	return o.MulticastOptimize, true
}

// HasMulticastOptimize returns a boolean if a field has been set.
func (o *FabricQosClass) HasMulticastOptimize() bool {
	if o != nil && !IsNil(o.MulticastOptimize) {
		return true
	}

	return false
}

// SetMulticastOptimize gets a reference to the given bool and assigns it to the MulticastOptimize field.
func (o *FabricQosClass) SetMulticastOptimize(v bool) {
	o.MulticastOptimize = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricQosClass) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClass) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricQosClass) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricQosClass) SetName(v string) {
	o.Name = &v
}

// GetPacketDrop returns the PacketDrop field value if set, zero value otherwise.
func (o *FabricQosClass) GetPacketDrop() bool {
	if o == nil || IsNil(o.PacketDrop) {
		var ret bool
		return ret
	}
	return *o.PacketDrop
}

// GetPacketDropOk returns a tuple with the PacketDrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClass) GetPacketDropOk() (*bool, bool) {
	if o == nil || IsNil(o.PacketDrop) {
		return nil, false
	}
	return o.PacketDrop, true
}

// HasPacketDrop returns a boolean if a field has been set.
func (o *FabricQosClass) HasPacketDrop() bool {
	if o != nil && !IsNil(o.PacketDrop) {
		return true
	}

	return false
}

// SetPacketDrop gets a reference to the given bool and assigns it to the PacketDrop field.
func (o *FabricQosClass) SetPacketDrop(v bool) {
	o.PacketDrop = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *FabricQosClass) GetWeight() int64 {
	if o == nil || IsNil(o.Weight) {
		var ret int64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricQosClass) GetWeightOk() (*int64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *FabricQosClass) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int64 and assigns it to the Weight field.
func (o *FabricQosClass) SetWeight(v int64) {
	o.Weight = &v
}

func (o FabricQosClass) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricQosClass) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AdminState) {
		toSerialize["AdminState"] = o.AdminState
	}
	if !IsNil(o.BandwidthPercent) {
		toSerialize["BandwidthPercent"] = o.BandwidthPercent
	}
	if !IsNil(o.Cos) {
		toSerialize["Cos"] = o.Cos
	}
	if !IsNil(o.Mtu) {
		toSerialize["Mtu"] = o.Mtu
	}
	if !IsNil(o.MulticastOptimize) {
		toSerialize["MulticastOptimize"] = o.MulticastOptimize
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PacketDrop) {
		toSerialize["PacketDrop"] = o.PacketDrop
	}
	if !IsNil(o.Weight) {
		toSerialize["Weight"] = o.Weight
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricQosClass) UnmarshalJSON(data []byte) (err error) {
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
	type FabricQosClassWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Administrative state for this QoS class. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
		AdminState *string `json:"AdminState,omitempty"`
		// Percentage of bandwidth received by the traffic tagged with this QoS.
		BandwidthPercent *int64 `json:"BandwidthPercent,omitempty"`
		// Class of service received by the traffic tagged with this QoS.
		Cos *int64 `json:"Cos,omitempty"`
		// Maximum transmission unit (MTU) is the largest size packet or frame, that can be sent in a packet- or frame-based network such as the Internet. User can select from the following: 1. Any value between 1500 and 9216 2. 'Normal' (default) mapping to a value of 1500. 3. 'FC' mapping to a value of 2240.
		Mtu *int64 `json:"Mtu,omitempty"`
		// If enabled, this QoS class will be optimized to send multiple packets.
		MulticastOptimize *bool `json:"MulticastOptimize,omitempty"`
		// The 'name' of this QoS Class. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
		Name *string `json:"Name,omitempty"`
		// If enabled, this QoS class will allow packet drops within an acceptable limit.
		PacketDrop *bool `json:"PacketDrop,omitempty"`
		// The weight of the QoS Class controls the distribution of bandwidth between QoS Classes, with the same priority after the Guarantees for the QoS Classes are reached.
		Weight *int64 `json:"Weight,omitempty"`
	}

	varFabricQosClassWithoutEmbeddedStruct := FabricQosClassWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricQosClassWithoutEmbeddedStruct)
	if err == nil {
		varFabricQosClass := _FabricQosClass{}
		varFabricQosClass.ClassId = varFabricQosClassWithoutEmbeddedStruct.ClassId
		varFabricQosClass.ObjectType = varFabricQosClassWithoutEmbeddedStruct.ObjectType
		varFabricQosClass.AdminState = varFabricQosClassWithoutEmbeddedStruct.AdminState
		varFabricQosClass.BandwidthPercent = varFabricQosClassWithoutEmbeddedStruct.BandwidthPercent
		varFabricQosClass.Cos = varFabricQosClassWithoutEmbeddedStruct.Cos
		varFabricQosClass.Mtu = varFabricQosClassWithoutEmbeddedStruct.Mtu
		varFabricQosClass.MulticastOptimize = varFabricQosClassWithoutEmbeddedStruct.MulticastOptimize
		varFabricQosClass.Name = varFabricQosClassWithoutEmbeddedStruct.Name
		varFabricQosClass.PacketDrop = varFabricQosClassWithoutEmbeddedStruct.PacketDrop
		varFabricQosClass.Weight = varFabricQosClassWithoutEmbeddedStruct.Weight
		*o = FabricQosClass(varFabricQosClass)
	} else {
		return err
	}

	varFabricQosClass := _FabricQosClass{}

	err = json.Unmarshal(data, &varFabricQosClass)
	if err == nil {
		o.MoBaseComplexType = varFabricQosClass.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "BandwidthPercent")
		delete(additionalProperties, "Cos")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "MulticastOptimize")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PacketDrop")
		delete(additionalProperties, "Weight")

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

type NullableFabricQosClass struct {
	value *FabricQosClass
	isSet bool
}

func (v NullableFabricQosClass) Get() *FabricQosClass {
	return v.value
}

func (v *NullableFabricQosClass) Set(val *FabricQosClass) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricQosClass) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricQosClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricQosClass(val *FabricQosClass) *NullableFabricQosClass {
	return &NullableFabricQosClass{value: val, isSet: true}
}

func (v NullableFabricQosClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricQosClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
