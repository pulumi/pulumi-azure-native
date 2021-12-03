


package devtestlab

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualMachineRdpFileContents(ctx *pulumi.Context, args *GetVirtualMachineRdpFileContentsArgs, opts ...pulumi.InvokeOption) (*GetVirtualMachineRdpFileContentsResult, error) {
	var rv GetVirtualMachineRdpFileContentsResult
	err := ctx.Invoke("azure-native:devtestlab:getVirtualMachineRdpFileContents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualMachineRdpFileContentsArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetVirtualMachineRdpFileContentsResult struct {
	Contents *string `pulumi:"contents"`
}
