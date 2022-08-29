


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineScaleSetVMRunCommand(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetVMRunCommandArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetVMRunCommandResult, error) {
	var rv LookupVirtualMachineScaleSetVMRunCommandResult
	err := ctx.Invoke("azure-native:compute/v20210401:getVirtualMachineScaleSetVMRunCommand", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualMachineScaleSetVMRunCommandArgs struct {
	Expand            *string `pulumi:"expand"`
	InstanceId        string  `pulumi:"instanceId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RunCommandName    string  `pulumi:"runCommandName"`
	VmScaleSetName    string  `pulumi:"vmScaleSetName"`
}


type LookupVirtualMachineScaleSetVMRunCommandResult struct {
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


func (val *LookupVirtualMachineScaleSetVMRunCommandResult) Defaults() *LookupVirtualMachineScaleSetVMRunCommandResult {
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

func LookupVirtualMachineScaleSetVMRunCommandOutput(ctx *pulumi.Context, args LookupVirtualMachineScaleSetVMRunCommandOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineScaleSetVMRunCommandResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineScaleSetVMRunCommandResult, error) {
			args := v.(LookupVirtualMachineScaleSetVMRunCommandArgs)
			r, err := LookupVirtualMachineScaleSetVMRunCommand(ctx, &args, opts...)
			var s LookupVirtualMachineScaleSetVMRunCommandResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineScaleSetVMRunCommandResultOutput)
}

type LookupVirtualMachineScaleSetVMRunCommandOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	InstanceId        pulumi.StringInput    `pulumi:"instanceId"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	RunCommandName    pulumi.StringInput    `pulumi:"runCommandName"`
	VmScaleSetName    pulumi.StringInput    `pulumi:"vmScaleSetName"`
}

func (LookupVirtualMachineScaleSetVMRunCommandOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetVMRunCommandArgs)(nil)).Elem()
}


type LookupVirtualMachineScaleSetVMRunCommandResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineScaleSetVMRunCommandResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetVMRunCommandResult)(nil)).Elem()
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) ToLookupVirtualMachineScaleSetVMRunCommandResultOutput() LookupVirtualMachineScaleSetVMRunCommandResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) ToLookupVirtualMachineScaleSetVMRunCommandResultOutputWithContext(ctx context.Context) LookupVirtualMachineScaleSetVMRunCommandResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) AsyncExecution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *bool { return v.AsyncExecution }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) ErrorBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *string { return v.ErrorBlobUri }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) InstanceView() VirtualMachineRunCommandInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) VirtualMachineRunCommandInstanceViewResponse {
		return v.InstanceView
	}).(VirtualMachineRunCommandInstanceViewResponseOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) OutputBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *string { return v.OutputBlobUri }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Parameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) []RunCommandInputParameterResponse {
		return v.Parameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) ProtectedParameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) []RunCommandInputParameterResponse {
		return v.ProtectedParameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) RunAsPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *string { return v.RunAsPassword }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) RunAsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *string { return v.RunAsUser }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Source() VirtualMachineRunCommandScriptSourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *VirtualMachineRunCommandScriptSourceResponse {
		return v.Source
	}).(VirtualMachineRunCommandScriptSourceResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *int { return v.TimeoutInSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineScaleSetVMRunCommandResultOutput{})
}
