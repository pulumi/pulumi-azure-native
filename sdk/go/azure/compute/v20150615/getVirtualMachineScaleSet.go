


package v20150615

import (
	"context"
	"reflect"

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

func LookupVirtualMachineScaleSetOutput(ctx *pulumi.Context, args LookupVirtualMachineScaleSetOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineScaleSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineScaleSetResult, error) {
			args := v.(LookupVirtualMachineScaleSetArgs)
			r, err := LookupVirtualMachineScaleSet(ctx, &args, opts...)
			var s LookupVirtualMachineScaleSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineScaleSetResultOutput)
}

type LookupVirtualMachineScaleSetOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VmScaleSetName    pulumi.StringInput `pulumi:"vmScaleSetName"`
}

func (LookupVirtualMachineScaleSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetArgs)(nil)).Elem()
}


type LookupVirtualMachineScaleSetResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineScaleSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetResult)(nil)).Elem()
}

func (o LookupVirtualMachineScaleSetResultOutput) ToLookupVirtualMachineScaleSetResultOutput() LookupVirtualMachineScaleSetResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetResultOutput) ToLookupVirtualMachineScaleSetResultOutputWithContext(ctx context.Context) LookupVirtualMachineScaleSetResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) OverProvision() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *bool { return v.OverProvision }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) UpgradePolicy() UpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *UpgradePolicyResponse { return v.UpgradePolicy }).(UpgradePolicyResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) VirtualMachineProfile() VirtualMachineScaleSetVMProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *VirtualMachineScaleSetVMProfileResponse {
		return v.VirtualMachineProfile
	}).(VirtualMachineScaleSetVMProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineScaleSetResultOutput{})
}
