### Example to create boot precision policy

``` go
package example

import (
	"fmt"
	"os"
	"log"
	intersight "github.com/CiscoDevNet/intersight-go"
)


func createBootLocalCdd() *intersight.BootDeviceBase {
	bootLocalCdd := intersight.NewBootDeviceBase("boot.LocalCdd", "boot.LocalCdd")
	bootLocalCdd.SetName("local_cdd1")
	bootLocalCdd.SetEnabled(true)
	return bootLocalCdd
}

func createBootLocalDisk() *intersight.BootDeviceBase {
	bootLocalDisk := intersight.NewBootDeviceBase("boot.LocalDisk", "boot.LocalDisk")
	bootLocalDisk.SetName("local_disk1")
	bootLocalDisk.SetEnabled(true)
	return bootLocalDisk
}

func createBootSdcard() *intersight.BootDeviceBase {
	bootSdcard := intersight.NewBootDeviceBase("boot.SdCard", "boot.SdCard")
	bootSdcard.SetName("boot_sdcard_test")
	bootSdcard.SetEnabled(true)
	return bootSdcard
}

func createBootIscsi() *intersight.BootDeviceBase {
	bootIscsi := intersight.NewBootDeviceBase("boot.Iscsi", "boot.Iscsi")
	bootIscsi.SetName("boot_iscsi_test")
	bootIscsi.SetEnabled(true)
	return bootIscsi
}

func createBootVirtualMedia() *intersight.BootDeviceBase {
	bootVirtualMedia := intersight.NewBootDeviceBase("boot.VirtualMedia", "boot.VirtualMedia")
	bootVirtualMedia.SetName("boot_virtual_media_test")
	bootVirtualMedia.SetEnabled(true)
	return bootVirtualMedia
}

func createOrganization() intersight.OrganizationOrganizationRelationship {
	organization := new(intersight.OrganizationOrganization)
	organization.ClassId = "mo.MoRef"
	organization.ObjectType = "organization.Organization"
	organizationRelationship := intersight.OrganizationOrganizationAsOrganizationOrganizationRelationship(organization)
	return organizationRelationship
}

func createOrganizationWithMoid(moid string) intersight.OrganizationOrganizationRelationship {
	organization := new(intersight.OrganizationOrganization)
	organization.ClassId = "mo.MoRef"
	organization.ObjectType = "organization.Organization"
	organization.Moid = &moid
	organizationRelationship := intersight.OrganizationOrganizationAsOrganizationOrganizationRelationship(organization)
	return organizationRelationship
}

func CreateObject(config *Config) {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	bootLocalCdd := createBootLocalCdd()
	bootLocalDisk := createBootLocalDisk()
	organization := createOrganization()
	bootDevices := []intersight.BootDeviceBase{*bootLocalDisk, *bootLocalCdd}
	bootPrecisionPolicy := intersight.NewBootPrecisionPolicyWithDefaults()
	bootPrecisionPolicy.SetName("sample_boot_policy1")
	bootPrecisionPolicy.SetDescription("sample boot precision policy")
	bootPrecisionPolicy.SetBootDevices(bootDevices)
	bootPrecisionPolicy.SetOrganization(organization)

	ifMatch := ""
	ifNoneMatch := ""
	resp, r, err := apiClient.BootApi.CreateBootPrecisionPolicy(ctx).BootPrecisionPolicy(*bootPrecisionPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)

	//Update
	id := resp.GetMoid()
	getapiResponse, r, err := apiClient.BootApi.GetBootPrecisionPolicyByMoid(ctx, id).Execute()
	if err != nil {
		log.Printf("Error -> GetBootPrecisionPolicyByMoid: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	objMoid := getapiResponse.GetMoid()
	organizationMoid := getapiResponse.GetOrganization().MoMoRef.GetMoid()
	bootSdcard := createBootSdcard()
	bootIscsi := createBootIscsi()
	organization1 := createOrganizationWithMoid(organizationMoid)
	bootDevices1 := []intersight.BootDeviceBase{*bootSdcard, *bootIscsi}
	updatebootPrecisionPolicy := intersight.NewBootPrecisionPolicyWithDefaults()
	updatebootPrecisionPolicy.SetName("updated_boot_precision_policy_for_go_test")
	updatebootPrecisionPolicy.SetDescription("updated description of boot precision policy for testing go example")
	updatebootPrecisionPolicy.SetBootDevices(bootDevices1)
	updatebootPrecisionPolicy.SetOrganization(organization1)
	updateResp, r, err := apiClient.BootApi.UpdateBootPrecisionPolicy(ctx, objMoid).BootPrecisionPolicy(*updatebootPrecisionPolicy).IfMatch(ifMatch).Execute()
	if err != nil {
		log.Printf("Error -> UpdateBootPrecisionPolicy: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response : %v\n", updateResp)

	//Patch
	bootVirtualMedia := createBootVirtualMedia()
	bootDevices2 := []intersight.BootDeviceBase{*bootVirtualMedia}
	patchbootPrecisionPolicy := intersight.NewBootPrecisionPolicyWithDefaults()
	patchbootPrecisionPolicy.SetName("updated_boot_precision_policy_using_patch_go_test")
	patchbootPrecisionPolicy.SetDescription("update the description of boot precision policy with patch for go test")
	patchbootPrecisionPolicy.SetBootDevices(bootDevices2)
	patchbootPrecisionPolicy.SetOrganization(organization1)
	patchResp, r, err := apiClient.BootApi.PatchBootPrecisionPolicy(ctx, objMoid).BootPrecisionPolicy(*patchbootPrecisionPolicy).IfMatch(ifMatch).Execute()
	if err != nil {
		log.Printf("Error -> PatchBootPrecisionPolicy: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response : %v\n", patchResp)

	//Delete
	fullResp, err := apiClient.BootApi.DeleteBootPrecisionPolicy(ctx, objMoid).Execute()
	if err != nil {
		log.Printf("Error -> DeleteBootPrecisionPolicy: %v\n", err)
		log.Printf("HTTP response: %v\n", fullResp)
		return
	}

}
```
