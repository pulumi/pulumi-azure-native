


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJitNetworkAccessPolicy(ctx *pulumi.Context, args *LookupJitNetworkAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupJitNetworkAccessPolicyResult, error) {
	var rv LookupJitNetworkAccessPolicyResult
	err := ctx.Invoke("azure-native:security:getJitNetworkAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJitNetworkAccessPolicyArgs struct {
	AscLocation                string `pulumi:"ascLocation"`
	JitNetworkAccessPolicyName string `pulumi:"jitNetworkAccessPolicyName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
}

type LookupJitNetworkAccessPolicyResult struct {
	Id                string                                         `pulumi:"id"`
	Kind              *string                                        `pulumi:"kind"`
	Location          string                                         `pulumi:"location"`
	Name              string                                         `pulumi:"name"`
	ProvisioningState string                                         `pulumi:"provisioningState"`
	Requests          []JitNetworkAccessRequestResponse              `pulumi:"requests"`
	Type              string                                         `pulumi:"type"`
	VirtualMachines   []JitNetworkAccessPolicyVirtualMachineResponse `pulumi:"virtualMachines"`
}

func LookupJitNetworkAccessPolicyOutput(ctx *pulumi.Context, args LookupJitNetworkAccessPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupJitNetworkAccessPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJitNetworkAccessPolicyResult, error) {
			args := v.(LookupJitNetworkAccessPolicyArgs)
			r, err := LookupJitNetworkAccessPolicy(ctx, &args, opts...)
			var s LookupJitNetworkAccessPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJitNetworkAccessPolicyResultOutput)
}

type LookupJitNetworkAccessPolicyOutputArgs struct {
	AscLocation                pulumi.StringInput `pulumi:"ascLocation"`
	JitNetworkAccessPolicyName pulumi.StringInput `pulumi:"jitNetworkAccessPolicyName"`
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupJitNetworkAccessPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJitNetworkAccessPolicyArgs)(nil)).Elem()
}

type LookupJitNetworkAccessPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupJitNetworkAccessPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJitNetworkAccessPolicyResult)(nil)).Elem()
}

func (o LookupJitNetworkAccessPolicyResultOutput) ToLookupJitNetworkAccessPolicyResultOutput() LookupJitNetworkAccessPolicyResultOutput {
	return o
}

func (o LookupJitNetworkAccessPolicyResultOutput) ToLookupJitNetworkAccessPolicyResultOutputWithContext(ctx context.Context) LookupJitNetworkAccessPolicyResultOutput {
	return o
}

func (o LookupJitNetworkAccessPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJitNetworkAccessPolicyResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupJitNetworkAccessPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupJitNetworkAccessPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJitNetworkAccessPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupJitNetworkAccessPolicyResultOutput) Requests() JitNetworkAccessRequestResponseArrayOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) []JitNetworkAccessRequestResponse { return v.Requests }).(JitNetworkAccessRequestResponseArrayOutput)
}

func (o LookupJitNetworkAccessPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupJitNetworkAccessPolicyResultOutput) VirtualMachines() JitNetworkAccessPolicyVirtualMachineResponseArrayOutput {
	return o.ApplyT(func(v LookupJitNetworkAccessPolicyResult) []JitNetworkAccessPolicyVirtualMachineResponse {
		return v.VirtualMachines
	}).(JitNetworkAccessPolicyVirtualMachineResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJitNetworkAccessPolicyResultOutput{})
}
