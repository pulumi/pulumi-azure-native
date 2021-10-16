


package v20170330

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:compute/v20170330:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VmName            string  `pulumi:"vmName"`
}


type LookupVirtualMachineResult struct {
	AvailabilitySet    *SubResourceResponse               `pulumi:"availabilitySet"`
	DiagnosticsProfile *DiagnosticsProfileResponse        `pulumi:"diagnosticsProfile"`
	HardwareProfile    *HardwareProfileResponse           `pulumi:"hardwareProfile"`
	Id                 string                             `pulumi:"id"`
	Identity           *VirtualMachineIdentityResponse    `pulumi:"identity"`
	InstanceView       VirtualMachineInstanceViewResponse `pulumi:"instanceView"`
	LicenseType        *string                            `pulumi:"licenseType"`
	Location           string                             `pulumi:"location"`
	Name               string                             `pulumi:"name"`
	NetworkProfile     *NetworkProfileResponse            `pulumi:"networkProfile"`
	OsProfile          *OSProfileResponse                 `pulumi:"osProfile"`
	Plan               *PlanResponse                      `pulumi:"plan"`
	ProvisioningState  string                             `pulumi:"provisioningState"`
	Resources          []VirtualMachineExtensionResponse  `pulumi:"resources"`
	StorageProfile     *StorageProfileResponse            `pulumi:"storageProfile"`
	Tags               map[string]string                  `pulumi:"tags"`
	Type               string                             `pulumi:"type"`
	VmId               string                             `pulumi:"vmId"`
	Zones              []string                           `pulumi:"zones"`
}
