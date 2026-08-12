#### Example to create policy
```
package policy

import (
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func CreateSolPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	solPolicy := intersight.NewSolPolicyWithDefaults()
	solPolicy.PolicyAbstractPolicy.SetName("sol_test")
	solPolicy.SetEnabled(false)
	solPolicy.SetBaudRate(int32(9600))
	solPolicy.SetComPort("com1")
	solPolicy.SetSshPort(int64(1096))
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	solPolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	resp, _, err := apiClient.SolApi.CreateSolPolicy(ctx).SolPolicy(*solPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	moid := resp.GetMoid()
	return moid
	}
```
