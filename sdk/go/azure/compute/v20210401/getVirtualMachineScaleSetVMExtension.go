


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineScaleSetVMExtension(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetVMExtensionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetVMExtensionResult, error) {
	var rv LookupVirtualMachineScaleSetVMExtensionResult
	err := ctx.Invoke("azure-native:compute/v20210401:getVirtualMachineScaleSetVMExtension", args, &rv, opts...)
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
	Type                    string                                       `pulumi:"type"`
	TypeHandlerVersion      *string                                      `pulumi:"typeHandlerVersion"`
}

func LookupVirtualMachineScaleSetVMExtensionOutput(ctx *pulumi.Context, args LookupVirtualMachineScaleSetVMExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineScaleSetVMExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineScaleSetVMExtensionResult, error) {
			args := v.(LookupVirtualMachineScaleSetVMExtensionArgs)
			r, err := LookupVirtualMachineScaleSetVMExtension(ctx, &args, opts...)
			var s LookupVirtualMachineScaleSetVMExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineScaleSetVMExtensionResultOutput)
}

type LookupVirtualMachineScaleSetVMExtensionOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	InstanceId        pulumi.StringInput    `pulumi:"instanceId"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	VmExtensionName   pulumi.StringInput    `pulumi:"vmExtensionName"`
	VmScaleSetName    pulumi.StringInput    `pulumi:"vmScaleSetName"`
}

func (LookupVirtualMachineScaleSetVMExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetVMExtensionArgs)(nil)).Elem()
}


type LookupVirtualMachineScaleSetVMExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineScaleSetVMExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetVMExtensionResult)(nil)).Elem()
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) ToLookupVirtualMachineScaleSetVMExtensionResultOutput() LookupVirtualMachineScaleSetVMExtensionResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) ToLookupVirtualMachineScaleSetVMExtensionResultOutputWithContext(ctx context.Context) LookupVirtualMachineScaleSetVMExtensionResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMExtensionResult) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMExtensionResult) *bool { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMExtensionResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) InstanceView() VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMExtensionResult) *VirtualMachineExtensionInstanceViewResponse {
		return v.InstanceView
	}).(VirtualMachineExtensionInstanceViewResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMExtensionResult) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMExtensionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMExtensionResult) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMExtensionResult) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMExtensionResultOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMExtensionResult) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineScaleSetVMExtensionResultOutput{})
}
