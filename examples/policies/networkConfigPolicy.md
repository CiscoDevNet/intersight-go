####Example to create policy
```
package policy

import (
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func CreateNetworkConfigPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	org_moid := getDefaultOrgMoid(config)
        organizationRelationship := getOrganizationRelationship(org_moid)
	networkConfigPolicy := intersight.NewNetworkconfigPolicyWithDefaults()
	networkConfigPolicy.SetName("tf_network_config1_sdk")
	networkConfigPolicy.SetDescription("test policy")
	networkConfigPolicy.SetEnableDynamicDns(false)
	networkConfigPolicy.SetPreferredIpv6dnsServer("::")
	networkConfigPolicy.SetEnableIpv6(true)
	networkConfigPolicy.SetEnableIpv6dnsFromDhcp(false)
	networkConfigPolicy.SetPreferredIpv4dnsServer("10.10.10.1")
	networkConfigPolicy.SetAlternateIpv4dnsServer("10.10.10.1")
	networkConfigPolicy.SetAlternateIpv6dnsServer("::")
	networkConfigPolicy.SetEnableIpv4dnsFromDhcp(false)
	networkConfigPolicy.SetOrganization(organizationRelationship)
	resp, _, err := apiClient.NetworkconfigApi.CreateNetworkconfigPolicy(ctx).NetworkconfigPolicy(*networkConfigPolicy).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	return resp.GetMoid()
}
```
