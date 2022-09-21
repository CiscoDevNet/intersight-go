####Example to create policy
```
package policy

import (
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)


func setServers() []string {
	servers := []string{"ntp.esl.cisco.com", "time-a-g.nist.gov", "time-b-g.nist.gov"}
	return servers
}

func CreateNtpPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	org_moid := getDefaultOrgMoid(config)
        organizationRelationship := getOrganizationRelationship(org_moid)
	ntpPolicy := intersight.NewNtpPolicyWithDefaults()
	ntpPolicy.SetName("tf_ntp1_sdk")
	ntpPolicy.SetEnabled(true)
	ntpPolicy.SetOrganization(organizationRelationship)
	servers := setServers()
	ntpPolicy.SetNtpServers(servers)
	resp, _, err := apiClient.NtpApi.CreateNtpPolicy(ctx).NtpPolicy(*ntpPolicy).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	return resp.GetMoid()
}
```
