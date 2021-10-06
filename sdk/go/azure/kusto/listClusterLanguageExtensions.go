


package kusto

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListClusterLanguageExtensions(ctx *pulumi.Context, args *ListClusterLanguageExtensionsArgs, opts ...pulumi.InvokeOption) (*ListClusterLanguageExtensionsResult, error) {
	var rv ListClusterLanguageExtensionsResult
	err := ctx.Invoke("azure-native:kusto:listClusterLanguageExtensions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListClusterLanguageExtensionsArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListClusterLanguageExtensionsResult struct {
	Value []LanguageExtensionResponse `pulumi:"value"`
}
