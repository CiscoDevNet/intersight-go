####Example to create policy
```
package policy

import (
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)

var ethAdapterPolicyMoid, ethNetworkMoid, ethQosMoid, lanConnectivityMoid string

func createEthAdapterPolicyRelationship(moid string) intersight.VnicEthAdapterPolicyRelationship {
	adapterPolicy := new(intersight.VnicEthAdapterPolicy)
	adapterPolicy.ClassId = "mo.MoRef"
	adapterPolicy.ObjectType = "vnic.EthAdapterPolicy"
	adapterPolicy.Moid = &moid
	adapterRelationship := intersight.VnicEthAdapterPolicyAsVnicEthAdapterPolicyRelationship(adapterPolicy)
	return adapterRelationship
}

func createEthNetworkPolicyRelationship(moid string) intersight.VnicEthNetworkPolicyRelationship {
	networkPolicy := new(intersight.VnicEthNetworkPolicy)
	networkPolicy.ClassId = "mo.MoRef"
	networkPolicy.ObjectType = "vnic.EthNetworkPolicy"
	networkPolicy.Moid = &moid
	networkRelationship := intersight.VnicEthNetworkPolicyAsVnicEthNetworkPolicyRelationship(networkPolicy)
	return networkRelationship
}

func createEthQosPolicyRelationship(moid string) intersight.VnicEthQosPolicyRelationship {
	qosPolicy := new(intersight.VnicEthQosPolicy)
	qosPolicy.ClassId = "mo.MoRef"
	qosPolicy.ObjectType = "vnic.EthQosPolicy"
	qosPolicy.Moid = &moid
	qosRelationship := intersight.VnicEthQosPolicyAsVnicEthQosPolicyRelationship(qosPolicy)
	return qosRelationship
}

func createLanConnectivityPolicy(moid string) intersight.VnicLanConnectivityPolicyRelationship {
	lanPolicy := new(intersight.VnicLanConnectivityPolicy)
	lanPolicy.ClassId = "mo.MoRef"
	lanPolicy.ObjectType = "vnic.LanConnectivityPolicy"
	lanPolicy.Moid = &moid
	lanRelationship := intersight.VnicLanConnectivityPolicyAsVnicLanConnectivityPolicyRelationship(lanPolicy)
	return lanRelationship
}

func CreateVnicEthAdapterPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	ethadapterPolicy := intersight.NewVnicEthAdapterPolicyWithDefaults()
	ethadapterPolicy.PolicyAbstractPolicy.SetName("v_eth_adapter_test")
	ethadapterPolicy.SetRssSettings(true)
	ethadapterPolicy.SetUplinkFailbackTimeout(int64(5))
	
	vxlanSettingsVal := intersight.NewVnicVxlanSettingsWithDefaults()
	vxlanSettingsVal.SetEnabled(false)
	vxlanSetting := intersight.NewNullableVnicVxlanSettings(vxlanSettingsVal)
	vxlanSettings := vxlanSetting.Get()
	ethadapterPolicy.SetVxlanSettings(*vxlanSettings)
	
	nvgreSettingsVal := intersight.NewVnicNvgreSettingsWithDefaults()
	nvgreSettingsVal.SetEnabled(true)
	nvgreSetting := intersight.NewNullableVnicNvgreSettings(nvgreSettingsVal)
	nvgreSettings := nvgreSetting.Get()
	ethadapterPolicy.SetNvgreSettings(*nvgreSettings)
	
	arfsSettingsVal := intersight.NewVnicArfsSettingsWithDefaults()
	arfsSettingsVal.SetEnabled(true)
	arfsSetting := intersight.NewNullableVnicArfsSettings(arfsSettingsVal)
	arfsSettings := arfsSetting.Get()
	ethadapterPolicy.SetArfsSettings(*arfsSettings)
	
	interruptSettingsVal := intersight.NewVnicEthInterruptSettingsWithDefaults()
	interruptSettingsVal.SetCoalescingTime(int64(125))
	interruptSettingsVal.SetCoalescingType("MIN")
	interruptSettingsVal.SetCount(int64(4))
	interruptSettingsVal.SetMode("MSI")
	interruptSetting := intersight.NewNullableVnicEthInterruptSettings(interruptSettingsVal)
	interruptSettings := interruptSetting.Get()
	ethadapterPolicy.SetInterruptSettings(*interruptSettings)
	
	completionQueueSettingsVal := intersight.NewVnicCompletionQueueSettingsWithDefaults()
	completionQueueSettingsVal.SetRingSize(int64(1))
	interruptSettingsVal.SetCount(int64(4))
	completionQueueSetting := intersight.NewNullableVnicCompletionQueueSettings(completionQueueSettingsVal)
	completionQueueSettings := completionQueueSetting.Get()
	ethadapterPolicy.SetCompletionQueueSettings(*completionQueueSettings)
	
	rxQueueSettingsVal := intersight.NewVnicEthRxQueueSettingsWithDefaults()
	rxQueueSettingsVal.SetRingSize(int64(512))
	rxQueueSettingsVal.SetCount(int64(4))
	rxQueueSetting := intersight.NewNullableVnicEthRxQueueSettings(rxQueueSettingsVal)
	rxQueueSettings := rxQueueSetting.Get()
	ethadapterPolicy.SetRxQueueSettings(*rxQueueSettings)
	
	txQueueSettingsVal := intersight.NewVnicEthTxQueueSettingsWithDefaults()
	txQueueSettingsVal.SetRingSize(int64(512))
	txQueueSettingsVal.SetCount(int64(4))
	txQueueSetting := intersight.NewNullableVnicEthTxQueueSettings(txQueueSettingsVal)
	txQueueSettings := txQueueSetting.Get()
	ethadapterPolicy.SetTxQueueSettings(*txQueueSettings)
	
	tcpOffloadSettingsVal := intersight.NewVnicTcpOffloadSettingsWithDefaults()
	tcpOffloadSettingsVal.SetLargeReceive(true)
	tcpOffloadSettingsVal.SetLargeSend(true)
	tcpOffloadSettingsVal.SetRxChecksum(true)
	tcpOffloadSettingsVal.SetTxChecksum(true)
	tcpOffloadSetting := intersight.NewNullableVnicTcpOffloadSettings(tcpOffloadSettingsVal)
	tcpOffloadSettings := tcpOffloadSetting.Get()
	ethadapterPolicy.SetTcpOffloadSettings(*tcpOffloadSettings)
	
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	ethadapterPolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	respAdapter, _, err := apiClient.VnicApi.CreateVnicEthAdapterPolicy(ctx).VnicEthAdapterPolicy(*ethadapterPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	ethAdapterPolicyMoid = respAdapter.GetMoid()
	return ethAdapterPolicyMoid
}

func CreateVnicEthNetworkPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	ethNetworkPolicy := intersight.NewVnicEthNetworkPolicyWithDefaults()
	ethNetworkPolicy.PolicyAbstractPolicy.SetName("v_eth_network_test")
	
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	ethNetworkPolicy.SetOrganization(organizationRelationship)
	
	vlanSettingsVal := intersight.NewVnicVlanSettingsWithDefaults()
	vlanSettingsVal.SetDefaultVlan(int64(1))
	vlanSettingsVal.SetMode("ACCESS")
	vlanSetting := intersight.NewNullableVnicVlanSettings(vlanSettingsVal)
	vlanSettings := vlanSetting.Get()
	ethNetworkPolicy.SetVlanSettings(*vlanSettings)
	
	ifMatch := ""
	ifNoneMatch := ""
	respNetwork, _, err := apiClient.VnicApi.CreateVnicEthNetworkPolicy(ctx).VnicEthNetworkPolicy(*ethNetworkPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	ethNetworkMoid = respNetwork.GetMoid()
	return ethNetworkMoid
}

func CreateVnicEthQosPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	ethQosPolicy := intersight.NewVnicEthQosPolicyWithDefaults()
	ethQosPolicy.PolicyAbstractPolicy.SetName("v_eth_qos_test")
	ethQosPolicy.SetMtu(int64(1500))
	ethQosPolicy.SetRateLimit(int64(0))
	ethQosPolicy.SetCos(int64(0))
	ethQosPolicy.SetTrustHostCos(false)
	
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	ethQosPolicy.SetOrganization(organizationRelationship)
	
	ifMatch := ""
	ifNoneMatch := ""
	respQos, _, err := apiClient.VnicApi.CreateVnicEthQosPolicy(ctx).VnicEthQosPolicy(*ethQosPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	ethQosMoid = respQos.GetMoid()
	return ethQosMoid
}

func CreateVnicLanConnectivityPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	vniclanConnectivityPolicy := intersight.NewVnicLanConnectivityPolicyWithDefaults()
	vniclanConnectivityPolicy.PolicyAbstractPolicy.SetName("vnic_lan_test")
	
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	vniclanConnectivityPolicy.SetOrganization(organizationRelationship)
	
	ifMatch := ""
	ifNoneMatch := ""
	respLan, _, err := apiClient.VnicApi.CreateVnicLanConnectivityPolicy(ctx).VnicLanConnectivityPolicy(*vniclanConnectivityPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	lanConnectivityMoid = respLan.GetMoid()
	return lanConnectivityMoid
}

func CreateVnicEthIf(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	ethIf := intersight.NewVnicEthIfWithDefaults()
	ethIf.SetName("vnic_eth_if_test")
	ethIf.SetOrder(int64(0))
	
	placementSettingVal := intersight.NewVnicPlacementSettingsWithDefaults()
	placementSettingVal.SetId("1")
	placementSettingVal.SetPciLink(int64(0))
	placementSettingVal.SetUplink(int64(0))
	placementSetting := intersight.NewNullableVnicPlacementSettings(placementSettingVal)
	placementSettings := placementSetting.Get()
	ethIf.SetPlacement(*placementSettings)
	
	cdnVal := intersight.NewVnicCdnWithDefaults()
	cdnVal.SetValue("VIC-1-eth00")
	cdnVal.SetSource("user")
	cdn := intersight.NewNullableVnicCdn(cdnVal)
	cdn1 := cdn.Get()
	ethIf.SetCdn(*cdn1)
	
	usnicSettingsVal := intersight.NewVnicUsnicSettingsWithDefaults()
	usnicSettingsVal.SetCos(int64(5))
	usnicSettingsVal.SetCount(int64(0))
	usnicSetting := intersight.NewNullableVnicUsnicSettings(usnicSettingsVal)
	usnicSettings := usnicSetting.Get()
	ethIf.SetUsnicSettings(*usnicSettings)
	
	vmqSettingsVal := intersight.NewVnicVmqSettingsWithDefaults()
	vmqSettingsVal.SetEnabled(true)
	vmqSettingsVal.SetMultiQueueSupport(false)
	vmqSettingsVal.SetNumInterrupts(int64(1))
	vmqSettingsVal.SetNumVmqs(int64(1))
	vmqSetting := intersight.NewNullableVnicVmqSettings(vmqSettingsVal)
	vmqSettings := vmqSetting.Get()
	ethIf.SetVmqSettings(*vmqSettings)
	
	ethAdapterPolicyRelationship := createEthAdapterPolicyRelationship(ethAdapterPolicyMoid)
	ethIf.SetEthAdapterPolicy(ethAdapterPolicyRelationship)
	
	ethNetworkPolicyRelationship := createEthNetworkPolicyRelationship(ethNetworkMoid)
	ethIf.SetEthNetworkPolicy(ethNetworkPolicyRelationship)
	
	ethQosPolicyRelationship := createEthQosPolicyRelationship(ethQosMoid)
	ethIf.SetEthQosPolicy(ethQosPolicyRelationship)
	
	lanConnectivityPolicyRelationship := createLanConnectivityPolicy(lanConnectivityMoid)
	ethIf.SetLanConnectivityPolicy(lanConnectivityPolicyRelationship)
	
	ifMatch := ""
	ifNoneMatch := ""
	ethIfResp, _, err := apiClient.VnicApi.CreateVnicEthIf(ctx).VnicEthIf(*ethIf).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	moid := ethIfResp.GetMoid()
	return moid
}
	
```
