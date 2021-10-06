


package v20171201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineExtension(ctx *pulumi.Context, args *LookupVirtualMachineExtensionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineExtensionResult, error) {
	var rv LookupVirtualMachineExtensionResult
	err := ctx.Invoke("azure-native:compute/v20171201:getVirtualMachineExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineExtensionArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VmExtensionName   string  `pulumi:"vmExtensionName"`
	VmName            string  `pulumi:"vmName"`
}


type LookupVirtualMachineExtensionResult struct {
	AutoUpgradeMinorVersion *bool                                        `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          *string                                      `pulumi:"forceUpdateTag"`
	Id                      string                                       `pulumi:"id"`
	InstanceView            *VirtualMachineExtensionInstanceViewResponse `pulumi:"instanceView"`
	Location                string                                       `pulumi:"location"`
	Name                    string                                       `pulumi:"name"`
	ProtectedSettings       interface{}                                  `pulumi:"protectedSettings"`
	ProvisioningState       string                                       `pulumi:"provisioningState"`
	Publisher               *string                                      `pulumi:"publisher"`
	Settings                interface{}                                  `pulumi:"settings"`
	Tags                    map[string]string                            `pulumi:"tags"`
	Type                    string                                       `pulumi:"type"`
	TypeHandlerVersion      *string                                      `pulumi:"typeHandlerVersion"`
}
