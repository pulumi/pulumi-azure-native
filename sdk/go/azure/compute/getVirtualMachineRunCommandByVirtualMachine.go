


package compute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineRunCommandByVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineRunCommandByVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineRunCommandByVirtualMachineResult, error) {
	var rv LookupVirtualMachineRunCommandByVirtualMachineResult
	err := ctx.Invoke("azure-native:compute:getVirtualMachineRunCommandByVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
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
