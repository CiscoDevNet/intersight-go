#### Example to create policy
```
package policy

import (
	"log"
	intersight "github.com/CiscoDevNet/intersight-go"
)

func CreateServerProfile(config *Config,policies []intersight.PolicyAbstractPolicyRelationship) {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	org_moid := getDefaultOrgMoid(config)
    organizationRelationship := getOrganizationRelationship(org_moid)
	serverProfile := intersight.NewServerProfileWithDefaults()
	serverProfile.ServerBaseProfile.PolicyAbstractConfigProfile.PolicyAbstractProfile.SetName("server_test")
	serverProfile.ServerBaseProfile.PolicyAbstractConfigProfile.SetAction("No-op")
	serverProfile.SetOrganization(organizationRelationship)
	serverProfile.ServerBaseProfile.PolicyAbstractConfigProfile.SetPolicyBucket(policies)
	
	resp, _, err := apiClient.ServerApi.CreateServerProfile(ctx).ServerProfile(*serverProfile).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	log.Printf("Response: %v\n", resp)
	log.Printf("Server profile created successfully.")
}

```
