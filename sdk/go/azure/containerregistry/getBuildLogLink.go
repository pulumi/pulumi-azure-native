


package containerregistry

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBuildLogLink(ctx *pulumi.Context, args *GetBuildLogLinkArgs, opts ...pulumi.InvokeOption) (*GetBuildLogLinkResult, error) {
	var rv GetBuildLogLinkResult
	err := ctx.Invoke("azure-native:containerregistry:getBuildLogLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBuildLogLinkArgs struct {
	BuildId           string `pulumi:"buildId"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetBuildLogLinkResult struct {
	LogLink *string `pulumi:"logLink"`
}
