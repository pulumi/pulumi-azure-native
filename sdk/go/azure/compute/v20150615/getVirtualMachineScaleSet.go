


package v20150615

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineScaleSet(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetResult, error) {
	var rv LookupVirtualMachineScaleSetResult
	err := ctx.Invoke("azure-native:compute/v20150615:getVirtualMachineScaleSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineScaleSetArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VmScaleSetName    string `pulumi:"vmScaleSetName"`
}


type LookupVirtualMachineScaleSetResult struct {
	Id                    string                                   `pulumi:"id"`
	Location              string                                   `pulumi:"location"`
	Name                  string                                   `pulumi:"name"`
	OverProvision         *bool                                    `pulumi:"overProvision"`
	ProvisioningState     *string                                  `pulumi:"provisioningState"`
	Sku                   *SkuResponse                             `pulumi:"sku"`
	Tags                  map[string]string                        `pulumi:"tags"`
	Type                  string                                   `pulumi:"type"`
	UpgradePolicy         *UpgradePolicyResponse                   `pulumi:"upgradePolicy"`
	VirtualMachineProfile *VirtualMachineScaleSetVMProfileResponse `pulumi:"virtualMachineProfile"`
}
