


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20201001preview:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	VirtualMachineName string `pulumi:"virtualMachineName"`
}


type LookupVirtualMachineResult struct {
	CustomResourceName string                     `pulumi:"customResourceName"`
	ExtendedLocation   *ExtendedLocationResponse  `pulumi:"extendedLocation"`
	FirmwareType       *string                    `pulumi:"firmwareType"`
	FolderPath         string                     `pulumi:"folderPath"`
	GuestAgentProfile  *GuestAgentProfileResponse `pulumi:"guestAgentProfile"`
	HardwareProfile    *HardwareProfileResponse   `pulumi:"hardwareProfile"`
	Id                 string                     `pulumi:"id"`
	Identity           *IdentityResponse          `pulumi:"identity"`
	InstanceUuid       string                     `pulumi:"instanceUuid"`
	InventoryItemId    *string                    `pulumi:"inventoryItemId"`
	Kind               *string                    `pulumi:"kind"`
	Location           string                     `pulumi:"location"`
	MoName             string                     `pulumi:"moName"`
	MoRefId            *string                    `pulumi:"moRefId"`
	Name               string                     `pulumi:"name"`
	NetworkProfile     *NetworkProfileResponse    `pulumi:"networkProfile"`
	OsProfile          *OsProfileResponse         `pulumi:"osProfile"`
	PlacementProfile   *PlacementProfileResponse  `pulumi:"placementProfile"`
	PowerState         string                     `pulumi:"powerState"`
	ProvisioningState  string                     `pulumi:"provisioningState"`
	ResourcePoolId     *string                    `pulumi:"resourcePoolId"`
	SmbiosUuid         *string                    `pulumi:"smbiosUuid"`
	Statuses           []ResourceStatusResponse   `pulumi:"statuses"`
	StorageProfile     *StorageProfileResponse    `pulumi:"storageProfile"`
	SystemData         SystemDataResponse         `pulumi:"systemData"`
	Tags               map[string]string          `pulumi:"tags"`
	TemplateId         *string                    `pulumi:"templateId"`
	Type               string                     `pulumi:"type"`
	Uuid               string                     `pulumi:"uuid"`
	VCenterId          *string                    `pulumi:"vCenterId"`
	VmId               string                     `pulumi:"vmId"`
}
