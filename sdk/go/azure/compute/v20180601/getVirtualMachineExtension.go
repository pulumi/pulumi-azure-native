


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupVirtualMachineExtension(ctx *pulumi.Context, args *LookupVirtualMachineExtensionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineExtensionResult, error) {
	var rv LookupVirtualMachineExtensionResult
	err := ctx.Invoke("azure-native:compute/v20180601:getVirtualMachineExtension", args, &rv, opts...)
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

func LookupVirtualMachineExtensionOutput(ctx *pulumi.Context, args LookupVirtualMachineExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineExtensionResult, error) {
			args := v.(LookupVirtualMachineExtensionArgs)
			r, err := LookupVirtualMachineExtension(ctx, &args, opts...)
			var s LookupVirtualMachineExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineExtensionResultOutput)
}

type LookupVirtualMachineExtensionOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	VmExtensionName   pulumi.StringInput    `pulumi:"vmExtensionName"`
	VmName            pulumi.StringInput    `pulumi:"vmName"`
}

func (LookupVirtualMachineExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineExtensionArgs)(nil)).Elem()
}


type LookupVirtualMachineExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineExtensionResult)(nil)).Elem()
}

func (o LookupVirtualMachineExtensionResultOutput) ToLookupVirtualMachineExtensionResultOutput() LookupVirtualMachineExtensionResultOutput {
	return o
}

func (o LookupVirtualMachineExtensionResultOutput) ToLookupVirtualMachineExtensionResultOutputWithContext(ctx context.Context) LookupVirtualMachineExtensionResultOutput {
	return o
}

func (o LookupVirtualMachineExtensionResultOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineExtensionResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineExtensionResultOutput) InstanceView() VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) *VirtualMachineExtensionInstanceViewResponse {
		return v.InstanceView
	}).(VirtualMachineExtensionInstanceViewResponsePtrOutput)
}

func (o LookupVirtualMachineExtensionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineExtensionResultOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o LookupVirtualMachineExtensionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineExtensionResultOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineExtensionResultOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o LookupVirtualMachineExtensionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineExtensionResultOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineExtensionResult) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineExtensionResultOutput{})
}
