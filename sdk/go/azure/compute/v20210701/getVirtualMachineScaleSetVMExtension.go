


package v20210701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineScaleSetVMExtension(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetVMExtensionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetVMExtensionResult, error) {
	var rv LookupVirtualMachineScaleSetVMExtensionResult
	err := ctx.Invoke("azure-native:compute/v20210701:getVirtualMachineScaleSetVMExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineScaleSetVMExtensionArgs struct {
	Expand            *string `pulumi:"expand"`
	InstanceId        string  `pulumi:"instanceId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VmExtensionName   string  `pulumi:"vmExtensionName"`
	VmScaleSetName    string  `pulumi:"vmScaleSetName"`
}


type LookupVirtualMachineScaleSetVMExtensionResult struct {
	AutoUpgradeMinorVersion *bool                                        `pulumi:"autoUpgradeMinorVersion"`
	EnableAutomaticUpgrade  *bool                                        `pulumi:"enableAutomaticUpgrade"`
	ForceUpdateTag          *string                                      `pulumi:"forceUpdateTag"`
	Id                      string                                       `pulumi:"id"`
	InstanceView            *VirtualMachineExtensionInstanceViewResponse `pulumi:"instanceView"`
	Name                    string                                       `pulumi:"name"`
	ProtectedSettings       interface{}                                  `pulumi:"protectedSettings"`
	ProvisioningState       string                                       `pulumi:"provisioningState"`
	Publisher               *string                                      `pulumi:"publisher"`
	Settings                interface{}                                  `pulumi:"settings"`
	SuppressFailures        *bool                                        `pulumi:"suppressFailures"`
	Type                    string                                       `pulumi:"type"`
	TypeHandlerVersion      *string                                      `pulumi:"typeHandlerVersion"`
}
