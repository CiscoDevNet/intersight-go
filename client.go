/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2025062323
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	JsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	XmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer("%5B", "[", "%5D", "]")
)

// APIClient manages communication with the Cisco Intersight API v1.0.11-2025062323
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AaaApi *AaaApiService

	AccessApi *AccessApiService

	AdapterApi *AdapterApiService

	ApicApi *ApicApiService

	ApplianceApi *ApplianceApiService

	AssetApi *AssetApiService

	BiosApi *BiosApiService

	BootApi *BootApiService

	BulkApi *BulkApiService

	CapabilityApi *CapabilityApiService

	CatalystsdwanApi *CatalystsdwanApiService

	CertificatemanagementApi *CertificatemanagementApiService

	ChangelogApi *ChangelogApiService

	ChassisApi *ChassisApiService

	CloudApi *CloudApiService

	CommApi *CommApiService

	ComputeApi *ComputeApiService

	CondApi *CondApiService

	ConnectorpackApi *ConnectorpackApiService

	ConsoleApi *ConsoleApiService

	ConvergedinfraApi *ConvergedinfraApiService

	CoremanagementApi *CoremanagementApiService

	CrdApi *CrdApiService

	DeviceconnectorApi *DeviceconnectorApiService

	DnacApi *DnacApiService

	EnergyApi *EnergyApiService

	EquipmentApi *EquipmentApiService

	EtherApi *EtherApiService

	ExternalsiteApi *ExternalsiteApiService

	FabricApi *FabricApiService

	FaultApi *FaultApiService

	FcApi *FcApiService

	FcpoolApi *FcpoolApiService

	FeedbackApi *FeedbackApiService

	FirmwareApi *FirmwareApiService

	FmcApi *FmcApiService

	ForecastApi *ForecastApiService

	GraphicsApi *GraphicsApiService

	HciApi *HciApiService

	HclApi *HclApiService

	HyperflexApi *HyperflexApiService

	IaasApi *IaasApiService

	IamApi *IamApiService

	InventoryApi *InventoryApiService

	IpmioverlanApi *IpmioverlanApiService

	IppoolApi *IppoolApiService

	IqnpoolApi *IqnpoolApiService

	IwotenantApi *IwotenantApiService

	KubernetesApi *KubernetesApiService

	KvmApi *KvmApiService

	LicenseApi *LicenseApiService

	LsApi *LsApiService

	MacpoolApi *MacpoolApiService

	ManagementApi *ManagementApiService

	MarketplaceApi *MarketplaceApiService

	MemoryApi *MemoryApiService

	MerakiApi *MerakiApiService

	MetaApi *MetaApiService

	MetricsApi *MetricsApiService

	MonitoringApi *MonitoringApiService

	NetworkApi *NetworkApiService

	NetworkconfigApi *NetworkconfigApiService

	NiaapiApi *NiaapiApiService

	NiatelemetryApi *NiatelemetryApiService

	NotificationApi *NotificationApiService

	NtpApi *NtpApiService

	OauthApi *OauthApiService

	OpenapiApi *OpenapiApiService

	OprsApi *OprsApiService

	OrganizationApi *OrganizationApiService

	OsApi *OsApiService

	PartnerintegrationApi *PartnerintegrationApiService

	PciApi *PciApiService

	PortApi *PortApiService

	PowerApi *PowerApiService

	ProcessorApi *ProcessorApiService

	RackApi *RackApiService

	RecommendationApi *RecommendationApiService

	RecoveryApi *RecoveryApiService

	ResourceApi *ResourceApiService

	ResourcepoolApi *ResourcepoolApiService

	RproxyApi *RproxyApiService

	SchedulerApi *SchedulerApiService

	SdaaciApi *SdaaciApiService

	SdcardApi *SdcardApiService

	SearchApi *SearchApiService

	SecurityApi *SecurityApiService

	ServerApi *ServerApiService

	ServicenowApi *ServicenowApiService

	SmtpApi *SmtpApiService

	SnmpApi *SnmpApiService

	SoftwareApi *SoftwareApiService

	SoftwarerepositoryApi *SoftwarerepositoryApiService

	SolApi *SolApiService

	SshApi *SshApiService

	StorageApi *StorageApiService

	SyslogApi *SyslogApiService

	TamApi *TamApiService

	TaskApi *TaskApiService

	TechsupportmanagementApi *TechsupportmanagementApiService

	TelemetryApi *TelemetryApiService

	TerminalApi *TerminalApiService

	ThermalApi *ThermalApiService

	TopApi *TopApiService

	UcsdApi *UcsdApiService

	UuidpoolApi *UuidpoolApiService

	ViewApi *ViewApiService

	VirtualizationApi *VirtualizationApiService

	VmediaApi *VmediaApiService

	VmrcApi *VmrcApiService

	VnicApi *VnicApiService

	VrfApi *VrfApiService

	WebhookApi *WebhookApiService

	WorkflowApi *WorkflowApiService

	WorkspaceApi *WorkspaceApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AaaApi = (*AaaApiService)(&c.common)
	c.AccessApi = (*AccessApiService)(&c.common)
	c.AdapterApi = (*AdapterApiService)(&c.common)
	c.ApicApi = (*ApicApiService)(&c.common)
	c.ApplianceApi = (*ApplianceApiService)(&c.common)
	c.AssetApi = (*AssetApiService)(&c.common)
	c.BiosApi = (*BiosApiService)(&c.common)
	c.BootApi = (*BootApiService)(&c.common)
	c.BulkApi = (*BulkApiService)(&c.common)
	c.CapabilityApi = (*CapabilityApiService)(&c.common)
	c.CatalystsdwanApi = (*CatalystsdwanApiService)(&c.common)
	c.CertificatemanagementApi = (*CertificatemanagementApiService)(&c.common)
	c.ChangelogApi = (*ChangelogApiService)(&c.common)
	c.ChassisApi = (*ChassisApiService)(&c.common)
	c.CloudApi = (*CloudApiService)(&c.common)
	c.CommApi = (*CommApiService)(&c.common)
	c.ComputeApi = (*ComputeApiService)(&c.common)
	c.CondApi = (*CondApiService)(&c.common)
	c.ConnectorpackApi = (*ConnectorpackApiService)(&c.common)
	c.ConsoleApi = (*ConsoleApiService)(&c.common)
	c.ConvergedinfraApi = (*ConvergedinfraApiService)(&c.common)
	c.CoremanagementApi = (*CoremanagementApiService)(&c.common)
	c.CrdApi = (*CrdApiService)(&c.common)
	c.DeviceconnectorApi = (*DeviceconnectorApiService)(&c.common)
	c.DnacApi = (*DnacApiService)(&c.common)
	c.EnergyApi = (*EnergyApiService)(&c.common)
	c.EquipmentApi = (*EquipmentApiService)(&c.common)
	c.EtherApi = (*EtherApiService)(&c.common)
	c.ExternalsiteApi = (*ExternalsiteApiService)(&c.common)
	c.FabricApi = (*FabricApiService)(&c.common)
	c.FaultApi = (*FaultApiService)(&c.common)
	c.FcApi = (*FcApiService)(&c.common)
	c.FcpoolApi = (*FcpoolApiService)(&c.common)
	c.FeedbackApi = (*FeedbackApiService)(&c.common)
	c.FirmwareApi = (*FirmwareApiService)(&c.common)
	c.FmcApi = (*FmcApiService)(&c.common)
	c.ForecastApi = (*ForecastApiService)(&c.common)
	c.GraphicsApi = (*GraphicsApiService)(&c.common)
	c.HciApi = (*HciApiService)(&c.common)
	c.HclApi = (*HclApiService)(&c.common)
	c.HyperflexApi = (*HyperflexApiService)(&c.common)
	c.IaasApi = (*IaasApiService)(&c.common)
	c.IamApi = (*IamApiService)(&c.common)
	c.InventoryApi = (*InventoryApiService)(&c.common)
	c.IpmioverlanApi = (*IpmioverlanApiService)(&c.common)
	c.IppoolApi = (*IppoolApiService)(&c.common)
	c.IqnpoolApi = (*IqnpoolApiService)(&c.common)
	c.IwotenantApi = (*IwotenantApiService)(&c.common)
	c.KubernetesApi = (*KubernetesApiService)(&c.common)
	c.KvmApi = (*KvmApiService)(&c.common)
	c.LicenseApi = (*LicenseApiService)(&c.common)
	c.LsApi = (*LsApiService)(&c.common)
	c.MacpoolApi = (*MacpoolApiService)(&c.common)
	c.ManagementApi = (*ManagementApiService)(&c.common)
	c.MarketplaceApi = (*MarketplaceApiService)(&c.common)
	c.MemoryApi = (*MemoryApiService)(&c.common)
	c.MerakiApi = (*MerakiApiService)(&c.common)
	c.MetaApi = (*MetaApiService)(&c.common)
	c.MetricsApi = (*MetricsApiService)(&c.common)
	c.MonitoringApi = (*MonitoringApiService)(&c.common)
	c.NetworkApi = (*NetworkApiService)(&c.common)
	c.NetworkconfigApi = (*NetworkconfigApiService)(&c.common)
	c.NiaapiApi = (*NiaapiApiService)(&c.common)
	c.NiatelemetryApi = (*NiatelemetryApiService)(&c.common)
	c.NotificationApi = (*NotificationApiService)(&c.common)
	c.NtpApi = (*NtpApiService)(&c.common)
	c.OauthApi = (*OauthApiService)(&c.common)
	c.OpenapiApi = (*OpenapiApiService)(&c.common)
	c.OprsApi = (*OprsApiService)(&c.common)
	c.OrganizationApi = (*OrganizationApiService)(&c.common)
	c.OsApi = (*OsApiService)(&c.common)
	c.PartnerintegrationApi = (*PartnerintegrationApiService)(&c.common)
	c.PciApi = (*PciApiService)(&c.common)
	c.PortApi = (*PortApiService)(&c.common)
	c.PowerApi = (*PowerApiService)(&c.common)
	c.ProcessorApi = (*ProcessorApiService)(&c.common)
	c.RackApi = (*RackApiService)(&c.common)
	c.RecommendationApi = (*RecommendationApiService)(&c.common)
	c.RecoveryApi = (*RecoveryApiService)(&c.common)
	c.ResourceApi = (*ResourceApiService)(&c.common)
	c.ResourcepoolApi = (*ResourcepoolApiService)(&c.common)
	c.RproxyApi = (*RproxyApiService)(&c.common)
	c.SchedulerApi = (*SchedulerApiService)(&c.common)
	c.SdaaciApi = (*SdaaciApiService)(&c.common)
	c.SdcardApi = (*SdcardApiService)(&c.common)
	c.SearchApi = (*SearchApiService)(&c.common)
	c.SecurityApi = (*SecurityApiService)(&c.common)
	c.ServerApi = (*ServerApiService)(&c.common)
	c.ServicenowApi = (*ServicenowApiService)(&c.common)
	c.SmtpApi = (*SmtpApiService)(&c.common)
	c.SnmpApi = (*SnmpApiService)(&c.common)
	c.SoftwareApi = (*SoftwareApiService)(&c.common)
	c.SoftwarerepositoryApi = (*SoftwarerepositoryApiService)(&c.common)
	c.SolApi = (*SolApiService)(&c.common)
	c.SshApi = (*SshApiService)(&c.common)
	c.StorageApi = (*StorageApiService)(&c.common)
	c.SyslogApi = (*SyslogApiService)(&c.common)
	c.TamApi = (*TamApiService)(&c.common)
	c.TaskApi = (*TaskApiService)(&c.common)
	c.TechsupportmanagementApi = (*TechsupportmanagementApiService)(&c.common)
	c.TelemetryApi = (*TelemetryApiService)(&c.common)
	c.TerminalApi = (*TerminalApiService)(&c.common)
	c.ThermalApi = (*ThermalApiService)(&c.common)
	c.TopApi = (*TopApiService)(&c.common)
	c.UcsdApi = (*UcsdApiService)(&c.common)
	c.UuidpoolApi = (*UuidpoolApiService)(&c.common)
	c.ViewApi = (*ViewApiService)(&c.common)
	c.VirtualizationApi = (*VirtualizationApiService)(&c.common)
	c.VmediaApi = (*VmediaApiService)(&c.common)
	c.VmrcApi = (*VmrcApiService)(&c.common)
	c.VnicApi = (*VnicApiService)(&c.common)
	c.VrfApi = (*VrfApiService)(&c.common)
	c.WebhookApi = (*WebhookApiService)(&c.common)
	c.WorkflowApi = (*WorkflowApiService)(&c.common)
	c.WorkspaceApi = (*WorkspaceApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("expected %s to be of type %s but received %s", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

func parameterValueToString(obj interface{}, key string) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return fmt.Sprintf("%v", obj)
	}
	var param, ok = obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap, err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}, style string, collectionType string) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
		case reflect.Invalid:
			value = "invalid"

		case reflect.Struct:
			if t, ok := obj.(MappedNullable); ok {
				dataMap, err := t.ToMap()
				if err != nil {
					return
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, dataMap, style, collectionType)
				return
			}
			if t, ok := obj.(time.Time); ok {
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, t.Format(time.RFC3339Nano), style, collectionType)
				return
			}
			value = v.Type().String() + " value"
		case reflect.Slice:
			var indValue = reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			var lenIndValue = indValue.Len()
			for i := 0; i < lenIndValue; i++ {
				var arrayValue = indValue.Index(i)
				var keyPrefixForCollectionType = keyPrefix
				if style == "deepObject" {
					keyPrefixForCollectionType = keyPrefix + "[" + strconv.Itoa(i) + "]"
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefixForCollectionType, arrayValue.Interface(), style, collectionType)
			}
			return

		case reflect.Map:
			var indValue = reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			iter := indValue.MapRange()
			for iter.Next() {
				k, v := iter.Key(), iter.Value()
				parameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), style, collectionType)
			}
			return

		case reflect.Interface:
			fallthrough
		case reflect.Ptr:
			parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), style, collectionType)
			return

		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			value = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			value = strconv.FormatUint(v.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
		case reflect.Bool:
			value = strconv.FormatBool(v.Bool())
		case reflect.String:
			value = v.String()
		default:
			value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
	case url.Values:
		if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
			valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix)+","+value)
		} else {
			valuesMap.Add(keyPrefix, value)
		}
		break
	case map[string]string:
		valuesMap[keyPrefix] = value
		break
	}
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

type formFile struct {
	fileBytes    []byte
	fileName     string
	formFileName string
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.fileBytes) > 0 && formFile.fileName != "" {
				w.Boundary()
				part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))
				if err != nil {
					return nil, err
				}
				_, err = part.Write(formFile.fileBytes)
				if err != nil {
					return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		return strings.Join(pieces, "=")
	})

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	if ctx != nil {
		// HTTP Signature Authentication. All request headers must be set (including default headers)
		// because the headers may be included in the signature.
		if auth, ok := ctx.Value(ContextHttpSignatureAuth).(HttpSignatureAuth); ok {
			err = SignRequest(ctx, localVarRequest, auth)
			if err != nil {
				return nil, err
			}
		}
	}
	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(*os.File); ok {
		f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Seek(0, io.SeekStart)
		return
	}
	if f, ok := v.(**os.File); ok {
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(*os.File); ok {
		_, err = bodyBuf.ReadFrom(fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if JsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if XmlCheck.MatchString(contentType) {
		var bs []byte
		bs, err = xml.Marshal(body)
		if err == nil {
			bodyBuf.Write(bs)
		}
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}
