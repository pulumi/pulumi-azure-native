


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListClusterUpgradableVersions(ctx *pulumi.Context, args *ListClusterUpgradableVersionsArgs, opts ...pulumi.InvokeOption) (*ListClusterUpgradableVersionsResult, error) {
	var rv ListClusterUpgradableVersionsResult
	err := ctx.Invoke("azure-native:servicefabric/v20210601:listClusterUpgradableVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListClusterUpgradableVersionsArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TargetVersion     string `pulumi:"targetVersion"`
}


type ListClusterUpgradableVersionsResult struct {
	SupportedPath []string `pulumi:"supportedPath"`
}
