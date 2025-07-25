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

// checks if the KubernetesClusterCertificateConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesClusterCertificateConfiguration{}

// KubernetesClusterCertificateConfiguration Certifcate and key configuration for Kubernetes cluster creation.
type KubernetesClusterCertificateConfiguration struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Certificate for the Kubernetes API server.
	CaCert *string `json:"CaCert,omitempty"`
	// Private Key for the Kubernetes API server.
	CaKey *string `json:"CaKey,omitempty"`
	// Certificate for the etcd cluster.
	EtcdCert          *string  `json:"EtcdCert,omitempty"`
	EtcdEncryptionKey []string `json:"EtcdEncryptionKey,omitempty"`
	// Private key for the etcd cluster.
	EtcdKey *string `json:"EtcdKey,omitempty"`
	// Certificate for the front proxy to support Kubernetes API aggregators.
	FrontProxyCert *string `json:"FrontProxyCert,omitempty"`
	// Private key for the front proxy to support Kubernetes API aggregators.
	FrontProxyKey *string `json:"FrontProxyKey,omitempty"`
	// Service account private key used by Kubernetes Token Controller to sign generated service account tokens.
	SaPrivateKey *string `json:"SaPrivateKey,omitempty"`
	// Service account public key used by Kubernetes API Server to validate service account tokens generated by the Token Controller.
	SaPublicKey          *string `json:"SaPublicKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesClusterCertificateConfiguration KubernetesClusterCertificateConfiguration

// NewKubernetesClusterCertificateConfiguration instantiates a new KubernetesClusterCertificateConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterCertificateConfiguration(classId string, objectType string) *KubernetesClusterCertificateConfiguration {
	this := KubernetesClusterCertificateConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesClusterCertificateConfigurationWithDefaults instantiates a new KubernetesClusterCertificateConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterCertificateConfigurationWithDefaults() *KubernetesClusterCertificateConfiguration {
	this := KubernetesClusterCertificateConfiguration{}
	var classId string = "kubernetes.ClusterCertificateConfiguration"
	this.ClassId = classId
	var objectType string = "kubernetes.ClusterCertificateConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesClusterCertificateConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesClusterCertificateConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.ClusterCertificateConfiguration" of the ClassId field.
func (o *KubernetesClusterCertificateConfiguration) GetDefaultClassId() interface{} {
	return "kubernetes.ClusterCertificateConfiguration"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesClusterCertificateConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesClusterCertificateConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.ClusterCertificateConfiguration" of the ObjectType field.
func (o *KubernetesClusterCertificateConfiguration) GetDefaultObjectType() interface{} {
	return "kubernetes.ClusterCertificateConfiguration"
}

// GetCaCert returns the CaCert field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetCaCert() string {
	if o == nil || IsNil(o.CaCert) {
		var ret string
		return ret
	}
	return *o.CaCert
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetCaCertOk() (*string, bool) {
	if o == nil || IsNil(o.CaCert) {
		return nil, false
	}
	return o.CaCert, true
}

// HasCaCert returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasCaCert() bool {
	if o != nil && !IsNil(o.CaCert) {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given string and assigns it to the CaCert field.
func (o *KubernetesClusterCertificateConfiguration) SetCaCert(v string) {
	o.CaCert = &v
}

// GetCaKey returns the CaKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetCaKey() string {
	if o == nil || IsNil(o.CaKey) {
		var ret string
		return ret
	}
	return *o.CaKey
}

// GetCaKeyOk returns a tuple with the CaKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetCaKeyOk() (*string, bool) {
	if o == nil || IsNil(o.CaKey) {
		return nil, false
	}
	return o.CaKey, true
}

// HasCaKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasCaKey() bool {
	if o != nil && !IsNil(o.CaKey) {
		return true
	}

	return false
}

// SetCaKey gets a reference to the given string and assigns it to the CaKey field.
func (o *KubernetesClusterCertificateConfiguration) SetCaKey(v string) {
	o.CaKey = &v
}

// GetEtcdCert returns the EtcdCert field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetEtcdCert() string {
	if o == nil || IsNil(o.EtcdCert) {
		var ret string
		return ret
	}
	return *o.EtcdCert
}

// GetEtcdCertOk returns a tuple with the EtcdCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetEtcdCertOk() (*string, bool) {
	if o == nil || IsNil(o.EtcdCert) {
		return nil, false
	}
	return o.EtcdCert, true
}

// HasEtcdCert returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasEtcdCert() bool {
	if o != nil && !IsNil(o.EtcdCert) {
		return true
	}

	return false
}

// SetEtcdCert gets a reference to the given string and assigns it to the EtcdCert field.
func (o *KubernetesClusterCertificateConfiguration) SetEtcdCert(v string) {
	o.EtcdCert = &v
}

// GetEtcdEncryptionKey returns the EtcdEncryptionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterCertificateConfiguration) GetEtcdEncryptionKey() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EtcdEncryptionKey
}

// GetEtcdEncryptionKeyOk returns a tuple with the EtcdEncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterCertificateConfiguration) GetEtcdEncryptionKeyOk() ([]string, bool) {
	if o == nil || IsNil(o.EtcdEncryptionKey) {
		return nil, false
	}
	return o.EtcdEncryptionKey, true
}

// HasEtcdEncryptionKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasEtcdEncryptionKey() bool {
	if o != nil && !IsNil(o.EtcdEncryptionKey) {
		return true
	}

	return false
}

// SetEtcdEncryptionKey gets a reference to the given []string and assigns it to the EtcdEncryptionKey field.
func (o *KubernetesClusterCertificateConfiguration) SetEtcdEncryptionKey(v []string) {
	o.EtcdEncryptionKey = v
}

// GetEtcdKey returns the EtcdKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetEtcdKey() string {
	if o == nil || IsNil(o.EtcdKey) {
		var ret string
		return ret
	}
	return *o.EtcdKey
}

// GetEtcdKeyOk returns a tuple with the EtcdKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetEtcdKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EtcdKey) {
		return nil, false
	}
	return o.EtcdKey, true
}

// HasEtcdKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasEtcdKey() bool {
	if o != nil && !IsNil(o.EtcdKey) {
		return true
	}

	return false
}

// SetEtcdKey gets a reference to the given string and assigns it to the EtcdKey field.
func (o *KubernetesClusterCertificateConfiguration) SetEtcdKey(v string) {
	o.EtcdKey = &v
}

// GetFrontProxyCert returns the FrontProxyCert field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetFrontProxyCert() string {
	if o == nil || IsNil(o.FrontProxyCert) {
		var ret string
		return ret
	}
	return *o.FrontProxyCert
}

// GetFrontProxyCertOk returns a tuple with the FrontProxyCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetFrontProxyCertOk() (*string, bool) {
	if o == nil || IsNil(o.FrontProxyCert) {
		return nil, false
	}
	return o.FrontProxyCert, true
}

// HasFrontProxyCert returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasFrontProxyCert() bool {
	if o != nil && !IsNil(o.FrontProxyCert) {
		return true
	}

	return false
}

// SetFrontProxyCert gets a reference to the given string and assigns it to the FrontProxyCert field.
func (o *KubernetesClusterCertificateConfiguration) SetFrontProxyCert(v string) {
	o.FrontProxyCert = &v
}

// GetFrontProxyKey returns the FrontProxyKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetFrontProxyKey() string {
	if o == nil || IsNil(o.FrontProxyKey) {
		var ret string
		return ret
	}
	return *o.FrontProxyKey
}

// GetFrontProxyKeyOk returns a tuple with the FrontProxyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetFrontProxyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.FrontProxyKey) {
		return nil, false
	}
	return o.FrontProxyKey, true
}

// HasFrontProxyKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasFrontProxyKey() bool {
	if o != nil && !IsNil(o.FrontProxyKey) {
		return true
	}

	return false
}

// SetFrontProxyKey gets a reference to the given string and assigns it to the FrontProxyKey field.
func (o *KubernetesClusterCertificateConfiguration) SetFrontProxyKey(v string) {
	o.FrontProxyKey = &v
}

// GetSaPrivateKey returns the SaPrivateKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetSaPrivateKey() string {
	if o == nil || IsNil(o.SaPrivateKey) {
		var ret string
		return ret
	}
	return *o.SaPrivateKey
}

// GetSaPrivateKeyOk returns a tuple with the SaPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetSaPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SaPrivateKey) {
		return nil, false
	}
	return o.SaPrivateKey, true
}

// HasSaPrivateKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasSaPrivateKey() bool {
	if o != nil && !IsNil(o.SaPrivateKey) {
		return true
	}

	return false
}

// SetSaPrivateKey gets a reference to the given string and assigns it to the SaPrivateKey field.
func (o *KubernetesClusterCertificateConfiguration) SetSaPrivateKey(v string) {
	o.SaPrivateKey = &v
}

// GetSaPublicKey returns the SaPublicKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetSaPublicKey() string {
	if o == nil || IsNil(o.SaPublicKey) {
		var ret string
		return ret
	}
	return *o.SaPublicKey
}

// GetSaPublicKeyOk returns a tuple with the SaPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetSaPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SaPublicKey) {
		return nil, false
	}
	return o.SaPublicKey, true
}

// HasSaPublicKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasSaPublicKey() bool {
	if o != nil && !IsNil(o.SaPublicKey) {
		return true
	}

	return false
}

// SetSaPublicKey gets a reference to the given string and assigns it to the SaPublicKey field.
func (o *KubernetesClusterCertificateConfiguration) SetSaPublicKey(v string) {
	o.SaPublicKey = &v
}

func (o KubernetesClusterCertificateConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesClusterCertificateConfiguration) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CaCert) {
		toSerialize["CaCert"] = o.CaCert
	}
	if !IsNil(o.CaKey) {
		toSerialize["CaKey"] = o.CaKey
	}
	if !IsNil(o.EtcdCert) {
		toSerialize["EtcdCert"] = o.EtcdCert
	}
	if o.EtcdEncryptionKey != nil {
		toSerialize["EtcdEncryptionKey"] = o.EtcdEncryptionKey
	}
	if !IsNil(o.EtcdKey) {
		toSerialize["EtcdKey"] = o.EtcdKey
	}
	if !IsNil(o.FrontProxyCert) {
		toSerialize["FrontProxyCert"] = o.FrontProxyCert
	}
	if !IsNil(o.FrontProxyKey) {
		toSerialize["FrontProxyKey"] = o.FrontProxyKey
	}
	if !IsNil(o.SaPrivateKey) {
		toSerialize["SaPrivateKey"] = o.SaPrivateKey
	}
	if !IsNil(o.SaPublicKey) {
		toSerialize["SaPublicKey"] = o.SaPublicKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesClusterCertificateConfiguration) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesClusterCertificateConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Certificate for the Kubernetes API server.
		CaCert *string `json:"CaCert,omitempty"`
		// Private Key for the Kubernetes API server.
		CaKey *string `json:"CaKey,omitempty"`
		// Certificate for the etcd cluster.
		EtcdCert          *string  `json:"EtcdCert,omitempty"`
		EtcdEncryptionKey []string `json:"EtcdEncryptionKey,omitempty"`
		// Private key for the etcd cluster.
		EtcdKey *string `json:"EtcdKey,omitempty"`
		// Certificate for the front proxy to support Kubernetes API aggregators.
		FrontProxyCert *string `json:"FrontProxyCert,omitempty"`
		// Private key for the front proxy to support Kubernetes API aggregators.
		FrontProxyKey *string `json:"FrontProxyKey,omitempty"`
		// Service account private key used by Kubernetes Token Controller to sign generated service account tokens.
		SaPrivateKey *string `json:"SaPrivateKey,omitempty"`
		// Service account public key used by Kubernetes API Server to validate service account tokens generated by the Token Controller.
		SaPublicKey *string `json:"SaPublicKey,omitempty"`
	}

	varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct := KubernetesClusterCertificateConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesClusterCertificateConfiguration := _KubernetesClusterCertificateConfiguration{}
		varKubernetesClusterCertificateConfiguration.ClassId = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.ClassId
		varKubernetesClusterCertificateConfiguration.ObjectType = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.ObjectType
		varKubernetesClusterCertificateConfiguration.CaCert = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.CaCert
		varKubernetesClusterCertificateConfiguration.CaKey = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.CaKey
		varKubernetesClusterCertificateConfiguration.EtcdCert = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.EtcdCert
		varKubernetesClusterCertificateConfiguration.EtcdEncryptionKey = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.EtcdEncryptionKey
		varKubernetesClusterCertificateConfiguration.EtcdKey = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.EtcdKey
		varKubernetesClusterCertificateConfiguration.FrontProxyCert = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.FrontProxyCert
		varKubernetesClusterCertificateConfiguration.FrontProxyKey = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.FrontProxyKey
		varKubernetesClusterCertificateConfiguration.SaPrivateKey = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.SaPrivateKey
		varKubernetesClusterCertificateConfiguration.SaPublicKey = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.SaPublicKey
		*o = KubernetesClusterCertificateConfiguration(varKubernetesClusterCertificateConfiguration)
	} else {
		return err
	}

	varKubernetesClusterCertificateConfiguration := _KubernetesClusterCertificateConfiguration{}

	err = json.Unmarshal(data, &varKubernetesClusterCertificateConfiguration)
	if err == nil {
		o.MoBaseComplexType = varKubernetesClusterCertificateConfiguration.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CaCert")
		delete(additionalProperties, "CaKey")
		delete(additionalProperties, "EtcdCert")
		delete(additionalProperties, "EtcdEncryptionKey")
		delete(additionalProperties, "EtcdKey")
		delete(additionalProperties, "FrontProxyCert")
		delete(additionalProperties, "FrontProxyKey")
		delete(additionalProperties, "SaPrivateKey")
		delete(additionalProperties, "SaPublicKey")

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

type NullableKubernetesClusterCertificateConfiguration struct {
	value *KubernetesClusterCertificateConfiguration
	isSet bool
}

func (v NullableKubernetesClusterCertificateConfiguration) Get() *KubernetesClusterCertificateConfiguration {
	return v.value
}

func (v *NullableKubernetesClusterCertificateConfiguration) Set(val *KubernetesClusterCertificateConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterCertificateConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterCertificateConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterCertificateConfiguration(val *KubernetesClusterCertificateConfiguration) *NullableKubernetesClusterCertificateConfiguration {
	return &NullableKubernetesClusterCertificateConfiguration{value: val, isSet: true}
}

func (v NullableKubernetesClusterCertificateConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterCertificateConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
