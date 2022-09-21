####Example to create policy
```
package policy

import (
	"context"
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)


type Config struct {
	ApiKey    string
	SecretKey string
	Endpoint  string
	ApiClient *intersight.APIClient
	ctx       context.Context
}


func ExecutePolicies(apiKey string, secret string, host string) {

	config := Config{
		ApiKey:    apiKey,
		SecretKey: secret,
		Endpoint:  host,
	}

	adapterconfigPolicyRelationship := getPolicyRelationship("adapter.ConfigPolicy", config)
	log.Printf("adapterconfigPolicy response: %v\n", adapterconfigPolicyRelationship)
	
	deviceconnectorPolicyRelationship := getPolicyRelationship("deviceconnector.Policy", config)
	log.Printf("deviceconnector response: %v\n", deviceconnectorPolicyRelationship)
	
	ldapPolicyRelationship := getPolicyRelationship("iam.LdapPolicy", config)
	log.Printf("ldapPolicy response: %v\n", ldapPolicyRelationship)
	
	ipmiPolicyRelationship := getPolicyRelationship("ipmioverlan.Policy", config)
	log.Printf("ipmiPolicy response: %v\n", ipmiPolicyRelationship)
	
	kvmPolicyRelationship := getPolicyRelationship("kvm.Policy", config)
	log.Printf("kvmPolicy response: %v\n", kvmPolicyRelationship)
	
	networkConfigPolicyRelationship := getPolicyRelationship("networkconfig.Policy", config)
	log.Printf("networkConfigPolicy response: %v\n", networkConfigPolicyRelationship)
	
	ntpPolicyRelationship := getPolicyRelationship("ntp.Policy", config)
	log.Printf("ntpPolicy response: %v\n", ntpPolicyRelationship)
	
	smtpPolicyRelationship := getPolicyRelationship("smtp.Policy", config)
	log.Printf("smtpPolicy response: %v\n", smtpPolicyRelationship)
	
	snmpPolicyRelationship := getPolicyRelationship("snmp.Policy", config)
	log.Printf("snmpPolicy response: %v\n", snmpPolicyRelationship)
	
	solPolicyRelationship := getPolicyRelationship("sol.Policy", config)
	log.Printf("solPolicy response: %v\n", solPolicyRelationship)
	
	syslogPolicyRelationship := getPolicyRelationship("syslog.Policy", config)
	log.Printf("syslogPolicy response: %v\n", syslogPolicyRelationship)
	
	userPolicyRelationship := getPolicyRelationship("iam.EndPointUserPolicy", config)
	log.Printf("userPolicy response: %v\n", userPolicyRelationship)
	
	storagePolicyRelationship := getPolicyRelationship("storage.StoragePolicy", config)
	log.Printf("storagePolicy response: %v\n", storagePolicyRelationship)
	
	storageDriveGroupRelationship := getPolicyRelationship("storage.DriveGroup", config)
	log.Printf("storageDriveGroup response: %v\n", storageDriveGroupRelationship)
	
	vmediaPolicyRelationship := getPolicyRelationship("vmedia.Policy", config)
	log.Printf("vmedia response: %v\n", vmediaPolicyRelationship)

	ethAdapterPolicyRelationship := getPolicyRelationship("vnic.EthAdapterPolicy", config)
	log.Printf("ethAdapterPolicy response: %v\n", ethAdapterPolicyRelationship)
	
	ethNetworkPolicyRelationship := getPolicyRelationship("vnic.EthNetworkPolicy", config)
	log.Printf("ethNetworkPolicy response: %v\n", ethNetworkPolicyRelationship)

	ethQosPolicyRelationship := getPolicyRelationship("vnic.EthQosPolicy", config)
	log.Printf("ethQosPolicy response: %v\n", ethQosPolicyRelationship)

	lanPolicyRelationship := getPolicyRelationship("vnic.LanConnectivityPolicy", config)
	log.Printf("lanPolicy response: %v\n", lanPolicyRelationship)

	ethIfPolicyRelationship := getPolicyRelationship("vnic.EthIf", config)
	log.Printf("ethIfPolicy response: %v\n", ethIfPolicyRelationship)
	
	fcAdapterPolicyRelationship := getPolicyRelationship("vnic.FcAdapterPolicy", config)
	log.Printf("fcAdapterPolicy response: %v\n", fcAdapterPolicyRelationship)

	fcNetworkPolicyRelationship := getPolicyRelationship("vnic.FcNetworkPolicy", config)
	log.Printf("fcNetworkPolicy response: %v\n", fcNetworkPolicyRelationship)
	
	fcQosPolicyRelationship := getPolicyRelationship("vnic.FcQosPolicy", config)
	log.Printf("fcQosPolicy response: %v\n", fcQosPolicyRelationship)
	
	sanPolicyRelationship := getPolicyRelationship("vnic.SanConnectivityPolicy", config)
	log.Printf("sanPolicy response: %v\n", sanPolicyRelationship)
	
	fcIfPolicyRelationship := getPolicyRelationship("vnic.FcIf", config)
	log.Printf("fcIfPolicy response: %v\n", fcIfPolicyRelationship)
	
	policies := []intersight.PolicyAbstractPolicyRelationship{*adapterconfigPolicyRelationship, *deviceconnectorPolicyRelationship, *ldapPolicyRelationship, *ipmiPolicyRelationship, *kvmPolicyRelationship, *networkConfigPolicyRelationship, *ntpPolicyRelationship, *smtpPolicyRelationship, *snmpPolicyRelationship, *solPolicyRelationship, *syslogPolicyRelationship, *userPolicyRelationship, *vmediaPolicyRelationship, *ethAdapterPolicyRelationship, *ethNetworkPolicyRelationship, *ethQosPolicyRelationship, *lanPolicyRelationship, *fcAdapterPolicyRelationship, *fcNetworkPolicyRelationship, *fcQosPolicyRelationship, *sanPolicyRelationship}
	CreateServerProfile(&config, policies)
}

func getOrganizationRelationship(moid string) intersight.OrganizationOrganizationRelationship {
        organization := new(intersight.OrganizationOrganization)
        organization.ClassId = "mo.MoRef"
        organization.ObjectType = "organization.Organization"
	organization.Moid = &moid
        organizationRelationship := intersight.OrganizationOrganizationAsOrganizationOrganizationRelationship(organization)
        return organizationRelationship
}

func getDefaultOrgMoid(config *Config) string {
        cfg := getApiClient(config)
        apiClient := cfg.ApiClient
        ctx := cfg.ctx

	org_resp, _, org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Fatal("Error: %v\n", org_err)
	}
	org_list := org_resp.OrganizationOrganizationList.GetResults()
	if len(org_list) == 0 {
		log.Fatal("Couldn't find the organization specified")
	}
	org_moid := org_list[0].MoBaseMo.GetMoid()
	return org_moid
}

func getPolicyRelationship(policy string, config Config) *intersight.PolicyAbstractPolicyRelationship {
	var policy_moid string
	switch policy {
		case "adapter.ConfigPolicy":
			policy_moid = CreateAdapterPolicy(&config)
		case "deviceconnector.Policy":
			policy_moid = CreateDeviceConnectorPolicy(&config)
		case "iam.LdapPolicy":
			policy_moid = CreateLdapPolicy(&config)
		case "ipmioverlan.Policy":
			policy_moid = CreateIpmiPolicy(&config)
		case "kvm.Policy":
			policy_moid = CreateKvmPolicy(&config)
		case "networkconfig.Policy":
			policy_moid = CreateNetworkConfigPolicy(&config)
		case "ntp.Policy":
			policy_moid = CreateNtpPolicy(&config)
		case "smtp.Policy":
			policy_moid = CreateSmtpPolicy(&config)
		case "snmp.Policy":
			policy_moid = CreateSnmpPolicy(&config)
		case "sol.Policy":
			policy_moid = CreateSolPolicy(&config)
		case "syslog.Policy":
			policy_moid = CreateSyslogPolicy(&config)
		case "iam.EndPointUserPolicy":
			policy_moid = CreateIamEndPointUserPolicy(&config)
		case "storage.StoragePolicy":
			policy_moid = CreateStorageStoragePolicy(&config)
		case "storage.DriveGroup":
			storageMoid := CreateStorageStoragePolicy(&config)
			policy_moid = CreateSorageDriveGroup(&config,storageMoid)
		case "vmedia.Policy":
			policy_moid = CreateVmediaPolicy(&config)
		case "vnic.EthAdapterPolicy":
			policy_moid = CreateVnicEthAdapterPolicy(&config)
		case "vnic.EthNetworkPolicy":
			policy_moid = CreateVnicEthNetworkPolicy(&config)
		case "vnic.EthQosPolicy":
			policy_moid = CreateVnicEthQosPolicy(&config)
		case "vnic.LanConnectivityPolicy":
			policy_moid = CreateVnicLanConnectivityPolicy(&config)
		case "vnic.EthIf":
			policy_moid = CreateVnicEthIf(&config)
		case "vnic.FcAdapterPolicy":
			policy_moid = CreateVnicFcAdapterPolicy(&config)
		case "vnic.FcNetworkPolicy":
			policy_moid = CreateVnicFcNetworkPolicy(&config)
		case "vnic.FcQosPolicy":
			policy_moid = CreateVnicFcQosPolicy(&config)
		case "vnic.SanConnectivityPolicy":
			policy_moid = CreateVnicSanConnectivityPolicy(&config)
		case "vnic.FcIf":
			policy_moid = CreateVnicFcIf(&config)	
	}
	
	policy_object := new(intersight.PolicyAbstractPolicy)
	policy_object.SetClassId("mo.MoRef")
	policy_object.SetObjectType(policy)
	policy_object.SetMoid(policy_moid)
	policyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(policy_object)
	return &policyRelationship
}
```
