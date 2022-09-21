####Example to create policy
```
package policy

import (
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func CreateSyslogPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	syslogPolicy := intersight.NewSyslogPolicyWithDefaults()
	syslogPolicy.PolicyAbstractPolicy.SetName("syslog_test")
	syslogPolicy.PolicyAbstractPolicy.SetDescription("demo syslog policy")
	
	remoteClient1 := intersight.NewSyslogRemoteClientBaseWithDefaults()
	remoteClient1.SetEnabled(true)
	remoteClient1.SetHostname("10.10.10.10")
	remoteClient1.SetPort(int64(514))
	remoteClient1.SetProtocol("tcp")
	remoteClient1.SetMinSeverity("emergency")
	
	
	remoteClient2 := intersight.NewSyslogRemoteClientBaseWithDefaults()
	remoteClient2.SetEnabled(true)
	remoteClient2.SetHostname("2001:0db8:0a0b:12f0:0000:0000:0000:0004")
	remoteClient2.SetPort(int64(64000))
	remoteClient2.SetProtocol("udp")
	remoteClient2.SetMinSeverity("emergency")
	remoteClients := []intersight.SyslogRemoteClientBase{*remoteClient1, *remoteClient2}
	syslogPolicy.SetRemoteClients(remoteClients)
	
	localClient1 := intersight.NewSyslogLocalClientBaseWithDefaults()
	localClient1.SetMinSeverity("emergency")
	localClients := []intersight.SyslogLocalClientBase{*localClient1}
	syslogPolicy.SetLocalClients(localClients)
	
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	syslogPolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	resp, _, err := apiClient.SyslogApi.CreateSyslogPolicy(ctx).SyslogPolicy(*syslogPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	moid := resp.GetMoid()
	return moid
}
```
