


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSourceControlRepositories(ctx *pulumi.Context, args *ListSourceControlRepositoriesArgs, opts ...pulumi.InvokeOption) (*ListSourceControlRepositoriesResult, error) {
	var rv ListSourceControlRepositoriesResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:listSourceControlRepositories", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSourceControlRepositoriesArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type ListSourceControlRepositoriesResult struct {
	NextLink string         `pulumi:"nextLink"`
	Value    []RepoResponse `pulumi:"value"`
}
