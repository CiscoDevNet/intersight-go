#### Example to create policy
```
package policy

import (
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func createSnmpTrap() *intersight.SnmpTrap {
	snmpTrap := intersight.NewSnmpTrap("snmp.Trap", "snmp.Trap")
	snmpTrap.SetDestination("10.10.10.1")
	snmpTrap.SetEnabled(false)
	snmpTrap.SetPort(int64(660))
	snmpTrap.SetType("Trap")
	snmpTrap.SetUser("demouser")
	snmpTrap.SetVersion("V3")
	return snmpTrap
}

func createSnmpUser() *intersight.SnmpUser {
	snmpUser := intersight.NewSnmpUser("snmp.User","snmp.User")
	snmpUser.SetName("demouser")
	snmpUser.SetPrivacyType("AES")
	snmpUser.SetAuthPassword("changeMe")
	snmpUser.SetPrivacyPassword("changeMe")
	snmpUser.SetSecurityLevel("AuthPriv")
	snmpUser.SetAuthType("SHA")
	return snmpUser
}

func CreateSnmpPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	snmpPolicy := intersight.NewSnmpPolicyWithDefaults()
	snmpPolicy.PolicyAbstractPolicy.SetName("snmp_policy_test")
	snmpPolicy.PolicyAbstractPolicy.SetDescription("sample testing snmp policy")
	snmpPolicy.SetEnabled(true)
	snmpPolicy.SetSnmpPort(int64(1983))
	snmpPolicy.SetAccessCommunityString("dummy123")
	snmpPolicy.SetCommunityAccess("Disabled")
	snmpPolicy.SetTrapCommunity("TrapCommunity")
	snmpPolicy.SetSysContact("aanimish")
	snmpPolicy.SetSysLocation("Karnataka")
	snmpPolicy.SetEngineId("vvb")
	snmpTrap1 := createSnmpTrap()
	snmpTraps := []intersight.SnmpTrap{*snmpTrap1}
	snmpPolicy.SetSnmpTraps(snmpTraps)
	snmpUser1 := createSnmpUser()
	snmpUsers := []intersight.SnmpUser{*snmpUser1}
	snmpPolicy.SetSnmpUsers(snmpUsers)

	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	snmpPolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	resp, _, err := apiClient.SnmpApi.CreateSnmpPolicy(ctx).SnmpPolicy(*snmpPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	moid := resp.GetMoid()
	return moid
}
```
