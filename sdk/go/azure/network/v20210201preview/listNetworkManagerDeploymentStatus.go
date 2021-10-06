


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNetworkManagerDeploymentStatus(ctx *pulumi.Context, args *ListNetworkManagerDeploymentStatusArgs, opts ...pulumi.InvokeOption) (*ListNetworkManagerDeploymentStatusResult, error) {
	var rv ListNetworkManagerDeploymentStatusResult
	err := ctx.Invoke("azure-native:network/v20210201preview:listNetworkManagerDeploymentStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNetworkManagerDeploymentStatusArgs struct {
	DeploymentTypes    []string `pulumi:"deploymentTypes"`
	NetworkManagerName string   `pulumi:"networkManagerName"`
	Regions            []string `pulumi:"regions"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	SkipToken          *string  `pulumi:"skipToken"`
}


type ListNetworkManagerDeploymentStatusResult struct {
	SkipToken *string                                  `pulumi:"skipToken"`
	Value     []NetworkManagerDeploymentStatusResponse `pulumi:"value"`
}
