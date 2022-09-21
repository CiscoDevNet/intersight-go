####Example to create policy
```
package policy

import (
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func CreateIamEndPointUserPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	userPolicy := intersight.NewIamEndPointUserPolicyWithDefaults()
	userPolicy.PolicyAbstractPolicy.SetName("iam_end_point_user_policy_test")
	userPolicy.PolicyAbstractPolicy.SetDescription("test user policy")
	
	passwordPropertiesVal := intersight.NewIamEndPointPasswordPropertiesWithDefaults()
	passwordPropertiesVal.SetEnforceStrongPassword(true)
	passwordPropertiesVal.SetEnablePasswordExpiry(true)
	passwordPropertiesVal.SetPasswordExpiryDuration(int64(50))
	passwordPropertiesVal.SetPasswordHistory(int64(5))
	passwordPropertiesVal.SetNotificationPeriod(int64(1))
	passwordPropertiesVal.SetGracePeriod(int64(2))
	passproperties := intersight.NewNullableIamEndPointPasswordProperties(passwordPropertiesVal)
	passwordProperties := passproperties.Get()
	userPolicy.SetPasswordProperties(*passwordProperties)
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	userPolicy.SetOrganization(organizationRelationship)
	

	ifMatch := ""
	ifNoneMatch := ""
	resp, _, err := apiClient.IamApi.CreateIamEndPointUserPolicy(ctx).IamEndPointUserPolicy(*userPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	moid := resp.GetMoid()
	return moid
}
```
