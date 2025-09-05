#### Example to create policy
```
package policy

import (
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func createVirtualDrives() *intersight.SdcardVirtualDrive {
	virtualDrives := intersight.NewSdcardVirtualDriveWithDefaults()
	virtualDrives.SetEnable(true)
	virtualDrives.SetObjectType("sdcard.OperatingSystem")
	return virtualDrives
}

func setPartitions() *intersight.SdcardPartition {
	partitions := intersight.NewSdcardPartitionWithDefaults()
	partitions.SetType("OS")
	partitions.SetObjectType("sdcard.Partition")
	virtualDrive := createVirtualDrives()
	virtualDrives := []intersight.SdcardVirtualDrive{*virtualDrive}
	partitions.SetVirtualDrives(virtualDrives)
	return partitions
}

func CreateSdCardPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	org_moid := getDefaultOrgMoid(config)
        organizationRelationship := getOrganizationRelationship(org_moid)
	partition := setPartitions()
	sdCardPolicy := intersight.NewSdcardPolicyWithDefaults()
	sdCardPolicy.SetName("tf_sdcard_sdk")
	sdCardPolicy.SetDescription("test policy")
	sdCardPolicy.SetOrganization(organizationRelationship)
        partitions := []intersight.SdcardPartition{*partition}
	sdCardPolicy.SetPartitions(partitions)
	resp, _, err := apiClient.SdcardApi.CreateSdcardPolicy(ctx).SdcardPolicy(*sdCardPolicy).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	return resp.GetMoid()
}
```
