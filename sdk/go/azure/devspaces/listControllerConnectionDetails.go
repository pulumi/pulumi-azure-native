


package devspaces

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListControllerConnectionDetails(ctx *pulumi.Context, args *ListControllerConnectionDetailsArgs, opts ...pulumi.InvokeOption) (*ListControllerConnectionDetailsResult, error) {
	var rv ListControllerConnectionDetailsResult
	err := ctx.Invoke("azure-native:devspaces:listControllerConnectionDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListControllerConnectionDetailsArgs struct {
	Name                          string `pulumi:"name"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	TargetContainerHostResourceId string `pulumi:"targetContainerHostResourceId"`
}

type ListControllerConnectionDetailsResult struct {
	ConnectionDetailsList []ControllerConnectionDetailsResponse `pulumi:"connectionDetailsList"`
}
