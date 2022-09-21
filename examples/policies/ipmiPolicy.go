####Example to create policy
```
package policy

import (
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)


func CreateIpmiPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	org_moid := getDefaultOrgMoid(config)
        organizationRelationship := getOrganizationRelationship(org_moid)
	ipmiPolicy := intersight.NewIpmioverlanPolicyWithDefaults()
	ipmiPolicy.SetName("tf_ipmi_sdk")
	ipmiPolicy.SetDescription("demo ipmi policy")
	ipmiPolicy.SetEnabled(true)
	ipmiPolicy.SetPrivilege("admin")
	ipmiPolicy.SetOrganization(organizationRelationship)
	resp, _, err := apiClient.IpmioverlanApi.CreateIpmioverlanPolicy(ctx).IpmioverlanPolicy(*ipmiPolicy).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	return resp.GetMoid()
}
```
