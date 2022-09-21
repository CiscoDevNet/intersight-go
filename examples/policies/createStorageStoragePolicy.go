####Example to create policy
```
package policy

import (
	"log"
	intersight "github.com/CiscoDevNet/intersight-go"
)

func CreateStorageStoragePolicy(config *Config) string{
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	storageStoragePolicy := intersight.NewStorageStoragePolicyWithDefaults()
	storageStoragePolicy.PolicyAbstractPolicy.SetName("storage_policy_test")
	storageStoragePolicy.PolicyAbstractPolicy.SetDescription( "storage policy test")
	storageStoragePolicy.SetUseJbodForVdCreation(true)
	storageStoragePolicy.SetUnusedDisksState("UnconfiguredGood")
	storageStoragePolicy.SetGlobalHotSpares("3")
	
	m2VirtualDrive := intersight.NewStorageM2VirtualDriveConfigWithDefaults()
	m2VirtualDrive.SetEnable(false)
	m2VirtualDrive1 := intersight.NewNullableStorageM2VirtualDriveConfig(m2VirtualDrive)
	m2Virtual1 := m2VirtualDrive1.Get()
	storageStoragePolicy.SetM2VirtualDrive(*m2Virtual1)
	
	org_moid := getDefaultOrgMoid(config)
	organizationRelationship := getOrganizationRelationship(org_moid)
	storageStoragePolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	resp, _, err := apiClient.StorageApi.CreateStorageStoragePolicy(ctx).StorageStoragePolicy(*storageStoragePolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	
	moid := resp.GetMoid()
	return moid
}
```
