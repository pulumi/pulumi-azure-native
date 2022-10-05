


package v20170330

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupVirtualMachineScaleSetExtension(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetExtensionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetExtensionResult, error) {
	var rv LookupVirtualMachineScaleSetExtensionResult
	err := ctx.Invoke("azure-native:compute/v20170330:getVirtualMachineScaleSetExtension", args, &rv, opts...)
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
	AutoUpgradeMinorVersion *bool       `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          *string     `pulumi:"forceUpdateTag"`
	Id                      string      `pulumi:"id"`
	Name                    *string     `pulumi:"name"`
	ProtectedSettings       interface{} `pulumi:"protectedSettings"`
	ProvisioningState       string      `pulumi:"provisioningState"`
	Publisher               *string     `pulumi:"publisher"`
	Settings                interface{} `pulumi:"settings"`
	Type                    *string     `pulumi:"type"`
	TypeHandlerVersion      *string     `pulumi:"typeHandlerVersion"`
}

func LookupVirtualMachineScaleSetExtensionOutput(ctx *pulumi.Context, args LookupVirtualMachineScaleSetExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineScaleSetExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineScaleSetExtensionResult, error) {
			args := v.(LookupVirtualMachineScaleSetExtensionArgs)
			r, err := LookupVirtualMachineScaleSetExtension(ctx, &args, opts...)
			var s LookupVirtualMachineScaleSetExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineScaleSetExtensionResultOutput)
}

type LookupVirtualMachineScaleSetExtensionOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	VmScaleSetName    pulumi.StringInput    `pulumi:"vmScaleSetName"`
	VmssExtensionName pulumi.StringInput    `pulumi:"vmssExtensionName"`
}

func (LookupVirtualMachineScaleSetExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetExtensionArgs)(nil)).Elem()
}


type LookupVirtualMachineScaleSetExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineScaleSetExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetExtensionResult)(nil)).Elem()
}

func (o LookupVirtualMachineScaleSetExtensionResultOutput) ToLookupVirtualMachineScaleSetExtensionResultOutput() LookupVirtualMachineScaleSetExtensionResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetExtensionResultOutput) ToLookupVirtualMachineScaleSetExtensionResultOutputWithContext(ctx context.Context) LookupVirtualMachineScaleSetExtensionResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetExtensionResultOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetExtensionResult) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineScaleSetExtensionResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetExtensionResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetExtensionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetExtensionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetExtensionResultOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetExtensionResult) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o LookupVirtualMachineScaleSetExtensionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetExtensionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetExtensionResultOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetExtensionResult) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetExtensionResultOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetExtensionResult) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o LookupVirtualMachineScaleSetExtensionResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetExtensionResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetExtensionResultOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetExtensionResult) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineScaleSetExtensionResultOutput{})
}
