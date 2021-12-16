


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineScaleSetVMRunCommand(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetVMRunCommandArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetVMRunCommandResult, error) {
	var rv LookupVirtualMachineScaleSetVMRunCommandResult
	err := ctx.Invoke("azure-native:compute/v20210301:getVirtualMachineScaleSetVMRunCommand", args, &rv, opts...)
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
