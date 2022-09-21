####Example to create policy
```
package policy

import (
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func setEth() intersight.AdapterEthSettings {
	eth := intersight.NewAdapterEthSettings("adapter.EthSettings", "adapter.EthSettings")
	eth.SetLldpEnabled(true)
	return *eth
}

func setFc() intersight.AdapterFcSettings {
	fc := intersight.NewAdapterFcSettings("adapter.FcSettings", "adapter.FcSettings")
	fc.SetFipEnabled(true)
	return *fc
}

func setSettings() *intersight.AdapterAdapterConfig {
	adapterConfig := intersight.NewAdapterAdapterConfigWithDefaults()
	ethSettings := setEth()
	fcSettings := setFc()
	adapterConfig.SetSlotId("1")
	adapterConfig.SetEthSettings(ethSettings)
	adapterConfig.SetFcSettings(fcSettings)
	return adapterConfig
}

func setSettingsMlom() *intersight.AdapterAdapterConfig {
	adapterConfig := intersight.NewAdapterAdapterConfigWithDefaults()
	ethSettings := setEth()
	fcSettings := setFc()
	adapterConfig.SetSlotId("MLOM")
	adapterConfig.SetEthSettings(ethSettings)
	adapterConfig.SetFcSettings(fcSettings)
	return adapterConfig
}


func CreateAdapterPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	settings := setSettings()
	settingsMlom := setSettingsMlom()
	settingsArr := []intersight.AdapterAdapterConfig{*settings, *settingsMlom}
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	adapterConfigPolicy := intersight.NewAdapterConfigPolicyWithDefaults()
	adapterConfigPolicy.SetSettings(settingsArr)
	adapterConfigPolicy.SetName("tf_adapter_config_sdk")
	adapterConfigPolicy.SetDescription("test policy")
	adapterConfigPolicy.SetOrganization(organizationRelationship)
	resp, _, err := apiClient.AdapterApi.CreateAdapterConfigPolicy(ctx).AdapterConfigPolicy(*adapterConfigPolicy).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	return resp.GetMoid()
}
```
