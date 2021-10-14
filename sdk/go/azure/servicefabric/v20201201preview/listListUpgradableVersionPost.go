


package v20201201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListListUpgradableVersionPost(ctx *pulumi.Context, args *ListListUpgradableVersionPostArgs, opts ...pulumi.InvokeOption) (*ListListUpgradableVersionPostResult, error) {
	var rv ListListUpgradableVersionPostResult
	err := ctx.Invoke("azure-native:servicefabric/v20201201preview:listListUpgradableVersionPost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListListUpgradableVersionPostArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TargetVersion     string `pulumi:"targetVersion"`
}


type ListListUpgradableVersionPostResult struct {
	SupportedPath []string `pulumi:"supportedPath"`
}
