#### Example to create policy
```
package policy

import (
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)

var adapterPolicyMoid, fcNetworkMoid, qosMoid, sanConnectivityMoid string

func createFcAdapterPolicyRelationship(moid string) intersight.VnicFcAdapterPolicyRelationship {
	adapterPolicy := new(intersight.VnicFcAdapterPolicy)
	adapterPolicy.ClassId = "mo.MoRef"
	adapterPolicy.ObjectType = "vnic.FcAdapterPolicy"
	adapterPolicy.Moid = &moid
	adapterRelationship := intersight.VnicFcAdapterPolicyAsVnicFcAdapterPolicyRelationship(adapterPolicy)
	return adapterRelationship
}

func createFcNetworkPolicyRelationship(moid string) intersight.VnicFcNetworkPolicyRelationship {
	networkPolicy := new(intersight.VnicFcNetworkPolicy)
	networkPolicy.ClassId = "mo.MoRef"
	networkPolicy.ObjectType = "vnic.FcNetworkPolicy"
	networkPolicy.Moid = &moid
	networkRelationship := intersight.VnicFcNetworkPolicyAsVnicFcNetworkPolicyRelationship(networkPolicy)
	return networkRelationship
}

func createFcQosPolicyRelationship(moid string) intersight.VnicFcQosPolicyRelationship {
	qosPolicy := new(intersight.VnicFcQosPolicy)
	qosPolicy.ClassId = "mo.MoRef"
	qosPolicy.ObjectType = "vnic.FcQosPolicy"
	qosPolicy.Moid = &moid
	qosRelationship := intersight.VnicFcQosPolicyAsVnicFcQosPolicyRelationship(qosPolicy)
	return qosRelationship
}

func createSanConnectivityPolicy(moid string) intersight.VnicSanConnectivityPolicyRelationship {
	sanPolicy := new(intersight.VnicSanConnectivityPolicy)
	sanPolicy.ClassId = "mo.MoRef"
	sanPolicy.ObjectType = "vnic.SanConnectivityPolicy"
	sanPolicy.Moid = &moid
	sanRelationship := intersight.VnicSanConnectivityPolicyAsVnicSanConnectivityPolicyRelationship(sanPolicy)
	return sanRelationship
}

func CreateVnicFcAdapterPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	fcAdapterPolicy := intersight.NewVnicFcAdapterPolicyWithDefaults()
	fcAdapterPolicy.PolicyAbstractPolicy.SetName("fc_adapter_test")
	fcAdapterPolicy.SetErrorDetectionTimeout(int64(100000))
	
	errorRecoverySettingsVal := intersight.NewVnicFcErrorRecoverySettingsWithDefaults()
	errorRecoverySettingsVal.SetEnabled(false)
	errorRecoverySettingsVal.SetIoRetryCount(int64(255))
	errorRecoverySettingsVal.SetIoRetryTimeout(int64(50))
	errorRecoverySettingsVal.SetLinkDownTimeout(int64(240000))
	errorRecoverySettingsVal.SetPortDownTimeout(int64(240000))
	errorRecoverySetting := intersight.NewNullableVnicFcErrorRecoverySettings(errorRecoverySettingsVal)
	errorRecoverySettings := errorRecoverySetting.Get()
	fcAdapterPolicy.SetErrorRecoverySettings(*errorRecoverySettings)
	
	flogiSettingsVal := intersight.NewVnicFlogiSettingsWithDefaults()
	flogiSettingsVal.SetRetries(int64(0))
	flogiSettingsVal.SetTimeout(int64(255000))
	flogiSetting := intersight.NewNullableVnicFlogiSettings(flogiSettingsVal)
	flogiSettings := flogiSetting.Get()
	fcAdapterPolicy.SetFlogiSettings(*flogiSettings)
	
	fcInterruptSettingsVal := intersight.NewVnicFcInterruptSettingsWithDefaults()
	fcInterruptSettingsVal.SetMode("MSIx")
	fcInterruptSetting := intersight.NewNullableVnicFcInterruptSettings(fcInterruptSettingsVal)
	fcInterruptSettings := fcInterruptSetting.Get()
	fcAdapterPolicy.SetInterruptSettings(*fcInterruptSettings)
	
	fcAdapterPolicy.SetIoThrottleCount(int64(1024))
	fcAdapterPolicy.SetLunCount(int64(1024))
	fcAdapterPolicy.SetLunQueueDepth(int64(254))

	plogiSettingsVal := intersight.NewVnicPlogiSettingsWithDefaults()
	plogiSettingsVal.SetRetries(int64(255))
	plogiSettingsVal.SetTimeout(int64(255000))
	plogiSetting := intersight.NewNullableVnicPlogiSettings(plogiSettingsVal)
	plogiSettings := plogiSetting.Get()
	fcAdapterPolicy.SetPlogiSettings(*plogiSettings)
	
	fcAdapterPolicy.SetResourceAllocationTimeout(int64(100000))
	
	rxQueueSettingsVal := intersight.NewVnicFcQueueSettingsWithDefaults()
	rxQueueSettingsVal.SetRingSize(int64(128))
	rxQueueSettingsVal.SetCount(int64(1))
	rxQueueSetting := intersight.NewNullableVnicFcQueueSettings(rxQueueSettingsVal)
	rxQueueSettings := rxQueueSetting.Get()
	fcAdapterPolicy.SetRxQueueSettings(*rxQueueSettings)
	
	txQueueSettingsVal := intersight.NewVnicFcQueueSettingsWithDefaults()
	txQueueSettingsVal.SetRingSize(int64(128))
	txQueueSettingsVal.SetCount(int64(1))
	txQueueSetting := intersight.NewNullableVnicFcQueueSettings(txQueueSettingsVal)
	txQueueSettings := txQueueSetting.Get()
	fcAdapterPolicy.SetTxQueueSettings(*txQueueSettings)
	
	scsiQueueSettingsVal := intersight.NewVnicScsiQueueSettingsWithDefaults()
	scsiQueueSettingsVal.SetRingSize(int64(152))
	scsiQueueSettingsVal.SetCount(int64(8))
	scsiQueueSetting := intersight.NewNullableVnicScsiQueueSettings(scsiQueueSettingsVal)
	scsiQueueSettings := scsiQueueSetting.Get()
	fcAdapterPolicy.SetScsiQueueSettings(*scsiQueueSettings)
	
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	fcAdapterPolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	adapterResp, _, err := apiClient.VnicApi.CreateVnicFcAdapterPolicy(ctx).VnicFcAdapterPolicy(*fcAdapterPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	adapterPolicyMoid = adapterResp.GetMoid()
	return adapterPolicyMoid
}

func CreateVnicFcNetworkPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	fcNetworkPolicy := intersight.NewVnicFcNetworkPolicyWithDefaults()
	fcNetworkPolicy.PolicyAbstractPolicy.SetName("fc_network_test")
	
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	fcNetworkPolicy.SetOrganization(organizationRelationship)
	
	vsanSettingsVal := intersight.NewVnicVsanSettingsWithDefaults()
	vsanSettingsVal.SetId(int64(100))
	vsanSetting := intersight.NewNullableVnicVsanSettings(vsanSettingsVal)
	vsanSettings := vsanSetting.Get()
	fcNetworkPolicy.SetVsanSettings(*vsanSettings)
	
	ifMatch := ""
	ifNoneMatch := ""
	networkResp, _, err := apiClient.VnicApi.CreateVnicFcNetworkPolicy(ctx).VnicFcNetworkPolicy(*fcNetworkPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fcNetworkMoid = networkResp.GetMoid()
	return fcNetworkMoid
}

func CreateVnicFcQosPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	fcQosPolicy := intersight.NewVnicFcQosPolicyWithDefaults()
	fcQosPolicy.PolicyAbstractPolicy.SetName("fc_qos_test")
	fcQosPolicy.SetRateLimit(int64(10000))
	fcQosPolicy.SetCos(int64(6))
	fcQosPolicy.SetMaxDataFieldSize(int64(2112))
	
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	fcQosPolicy.SetOrganization(organizationRelationship)
	
	ifMatch := ""
	ifNoneMatch := ""
	qosResp, _, err := apiClient.VnicApi.CreateVnicFcQosPolicy(ctx).VnicFcQosPolicy(*fcQosPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	qosMoid = qosResp.GetMoid()
	return qosMoid
}

func CreateVnicSanConnectivityPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	sanConnectivityPolicy := intersight.NewVnicSanConnectivityPolicyWithDefaults()
	sanConnectivityPolicy.PolicyAbstractPolicy.SetName("vnic_san_test")
	
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	sanConnectivityPolicy.SetOrganization(organizationRelationship)
	
	ifMatch := ""
	ifNoneMatch := ""
	sanResp, _, err := apiClient.VnicApi.CreateVnicSanConnectivityPolicy(ctx).VnicSanConnectivityPolicy(*sanConnectivityPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	sanConnectivityMoid = sanResp.GetMoid()
	return sanConnectivityMoid
}

func CreateVnicFcIf(config *Config) string{
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	fcIf := intersight.NewVnicFcIfWithDefaults()
	fcIf.SetName("fc0")
	fcIf.SetOrder(int64(1))
	
	placementSettingVal := intersight.NewVnicPlacementSettingsWithDefaults() 
	placementSettingVal.SetId("1")
	placementSettingVal.SetPciLink(int64(0))
	placementSettingVal.SetUplink(int64(0))
	placementSetting := intersight.NewNullableVnicPlacementSettings(placementSettingVal)
	placementSettings := placementSetting.Get()
	fcIf.SetPlacement(*placementSettings)
	
	fcIf.SetPersistentBindings(true)
	
	fcAdapterPolicyRelationship := createFcAdapterPolicyRelationship(adapterPolicyMoid)
	fcIf.SetFcAdapterPolicy(fcAdapterPolicyRelationship)
	
	fcNetworkPolicyRelationship := createFcNetworkPolicyRelationship(fcNetworkMoid)
	fcIf.SetFcNetworkPolicy(fcNetworkPolicyRelationship)
	
	fcQosPolicyRelationship := createFcQosPolicyRelationship(qosMoid)
	fcIf.SetFcQosPolicy(fcQosPolicyRelationship)
	
	sanConnectivityPolicyRelationship := createSanConnectivityPolicy(sanConnectivityMoid)
	fcIf.SetSanConnectivityPolicy(sanConnectivityPolicyRelationship)
	
	ifMatch := ""
	ifNoneMatch := ""
	fcIfResp, _, err := apiClient.VnicApi.CreateVnicFcIf(ctx).VnicFcIf(*fcIf).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	moid := fcIfResp.GetMoid()
	return moid
}
	
```
