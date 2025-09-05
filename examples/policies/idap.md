#### Example to create policy
```
package policy

import (
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)


func setBaseProperties() intersight.IamLdapBaseProperties {
	ldapBaseProperties := intersight.NewIamLdapBasePropertiesWithDefaults()
	ldapBaseProperties.SetAttribute("CiscoAvPair")
	ldapBaseProperties.SetBaseDn("DC=QATCSLABTPI02,DC=cisco,DC=com")
	ldapBaseProperties.SetBindDn("CN=administrator,CN=Users,DC=QATCSLABTPI02,DC=cisco,DC=com")
	ldapBaseProperties.SetBindMethod("Anonymous")
	ldapBaseProperties.SetDomain("QATCSLABTPI02.cisco.com")
	ldapBaseProperties.SetEnableEncryption(true)
	ldapBaseProperties.SetEnableGroupAuthorization(true)
	ldapBaseProperties.SetFilter("sAMAccountName")
	ldapBaseProperties.SetGroupAttribute("memberOf")
	ldapBaseProperties.SetNestedGroupSearchDepth(128)
	ldapBaseProperties.SetTimeout(180)
	return *ldapBaseProperties
}

func setDnsProperties() intersight.IamLdapDnsParameters {
	ldapDnsProperties := intersight.NewIamLdapDnsParametersWithDefaults()
	ldapDnsProperties.SetSource("Extracted")
	ldapDnsProperties.SetSearchForest("xyz")
	ldapDnsProperties.SetSearchDomain("abc")
	return *ldapDnsProperties
}

func CreateLdapPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	org_moid := getDefaultOrgMoid(config)
        organizationRelationship := getOrganizationRelationship(org_moid)
	baseProperties := setBaseProperties()
	dnsProperties := setDnsProperties()
	ldapPolicy := intersight.NewIamLdapPolicyWithDefaults()
	ldapPolicy.SetName("tf_ldap1_sdk")
	ldapPolicy.SetDescription("test policy")
	ldapPolicy.SetEnabled(true)
	ldapPolicy.SetEnableDns(true)
	ldapPolicy.SetUserSearchPrecedence("LocalUserDb")
	ldapPolicy.SetOrganization(organizationRelationship)
	ldapPolicy.SetBaseProperties(baseProperties)
	ldapPolicy.SetDnsParameters(dnsProperties)
	resp, _, err := apiClient.IamApi.CreateIamLdapPolicy(ctx).IamLdapPolicy(*ldapPolicy).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	return resp.GetMoid()
}
```
