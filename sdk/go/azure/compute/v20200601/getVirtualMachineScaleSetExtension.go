


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineScaleSetExtension(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetExtensionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetExtensionResult, error) {
	var rv LookupVirtualMachineScaleSetExtensionResult
	err := ctx.Invoke("azure-native:compute/v20200601:getVirtualMachineScaleSetExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineScaleSetExtensionArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VmScaleSetName    string  `pulumi:"vmScaleSetName"`
	VmssExtensionName string  `pulumi:"vmssExtensionName"`
}


type LookupVirtualMachineScaleSetExtensionResult struct {
	AutoUpgradeMinorVersion  *bool       `pulumi:"autoUpgradeMinorVersion"`
	EnableAutomaticUpgrade   *bool       `pulumi:"enableAutomaticUpgrade"`
	ForceUpdateTag           *string     `pulumi:"forceUpdateTag"`
	Id                       string      `pulumi:"id"`
	Name                     *string     `pulumi:"name"`
	ProtectedSettings        interface{} `pulumi:"protectedSettings"`
	ProvisionAfterExtensions []string    `pulumi:"provisionAfterExtensions"`
	ProvisioningState        string      `pulumi:"provisioningState"`
	Publisher                *string     `pulumi:"publisher"`
	Settings                 interface{} `pulumi:"settings"`
	Type                     string      `pulumi:"type"`
	TypeHandlerVersion       *string     `pulumi:"typeHandlerVersion"`
}
