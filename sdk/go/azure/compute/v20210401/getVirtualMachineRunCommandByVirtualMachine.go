


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineRunCommandByVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineRunCommandByVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineRunCommandByVirtualMachineResult, error) {
	var rv LookupVirtualMachineRunCommandByVirtualMachineResult
	err := ctx.Invoke("azure-native:compute/v20210401:getVirtualMachineRunCommandByVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualMachineRunCommandByVirtualMachineArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RunCommandName    string  `pulumi:"runCommandName"`
	VmName            string  `pulumi:"vmName"`
}


type LookupVirtualMachineRunCommandByVirtualMachineResult struct {
	AsyncExecution      *bool                                         `pulumi:"asyncExecution"`
	ErrorBlobUri        *string                                       `pulumi:"errorBlobUri"`
	Id                  string                                        `pulumi:"id"`
	InstanceView        VirtualMachineRunCommandInstanceViewResponse  `pulumi:"instanceView"`
	Location            string                                        `pulumi:"location"`
	Name                string                                        `pulumi:"name"`
	OutputBlobUri       *string                                       `pulumi:"outputBlobUri"`
	Parameters          []RunCommandInputParameterResponse            `pulumi:"parameters"`
	ProtectedParameters []RunCommandInputParameterResponse            `pulumi:"protectedParameters"`
	ProvisioningState   string                                        `pulumi:"provisioningState"`
	RunAsPassword       *string                                       `pulumi:"runAsPassword"`
	RunAsUser           *string                                       `pulumi:"runAsUser"`
	Source              *VirtualMachineRunCommandScriptSourceResponse `pulumi:"source"`
	Tags                map[string]string                             `pulumi:"tags"`
	TimeoutInSeconds    *int                                          `pulumi:"timeoutInSeconds"`
	Type                string                                        `pulumi:"type"`
}


func (val *LookupVirtualMachineRunCommandByVirtualMachineResult) Defaults() *LookupVirtualMachineRunCommandByVirtualMachineResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AsyncExecution) {
		asyncExecution_ := false
		tmp.AsyncExecution = &asyncExecution_
	}
	return &tmp
}

func LookupVirtualMachineRunCommandByVirtualMachineOutput(ctx *pulumi.Context, args LookupVirtualMachineRunCommandByVirtualMachineOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineRunCommandByVirtualMachineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineRunCommandByVirtualMachineResult, error) {
			args := v.(LookupVirtualMachineRunCommandByVirtualMachineArgs)
			r, err := LookupVirtualMachineRunCommandByVirtualMachine(ctx, &args, opts...)
			var s LookupVirtualMachineRunCommandByVirtualMachineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineRunCommandByVirtualMachineResultOutput)
}

type LookupVirtualMachineRunCommandByVirtualMachineOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	RunCommandName    pulumi.StringInput    `pulumi:"runCommandName"`
	VmName            pulumi.StringInput    `pulumi:"vmName"`
}

func (LookupVirtualMachineRunCommandByVirtualMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineRunCommandByVirtualMachineArgs)(nil)).Elem()
}


type LookupVirtualMachineRunCommandByVirtualMachineResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineRunCommandByVirtualMachineResult)(nil)).Elem()
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ToLookupVirtualMachineRunCommandByVirtualMachineResultOutput() LookupVirtualMachineRunCommandByVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ToLookupVirtualMachineRunCommandByVirtualMachineResultOutputWithContext(ctx context.Context) LookupVirtualMachineRunCommandByVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) AsyncExecution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *bool { return v.AsyncExecution }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ErrorBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *string { return v.ErrorBlobUri }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) InstanceView() VirtualMachineRunCommandInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) VirtualMachineRunCommandInstanceViewResponse {
		return v.InstanceView
	}).(VirtualMachineRunCommandInstanceViewResponseOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) OutputBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *string { return v.OutputBlobUri }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Parameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) []RunCommandInputParameterResponse {
		return v.Parameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ProtectedParameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) []RunCommandInputParameterResponse {
		return v.ProtectedParameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) RunAsPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *string { return v.RunAsPassword }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) RunAsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *string { return v.RunAsUser }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Source() VirtualMachineRunCommandScriptSourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *VirtualMachineRunCommandScriptSourceResponse {
		return v.Source
	}).(VirtualMachineRunCommandScriptSourceResponsePtrOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *int { return v.TimeoutInSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineRunCommandByVirtualMachineResultOutput{})
}
