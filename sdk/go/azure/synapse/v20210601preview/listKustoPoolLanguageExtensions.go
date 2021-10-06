


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListKustoPoolLanguageExtensions(ctx *pulumi.Context, args *ListKustoPoolLanguageExtensionsArgs, opts ...pulumi.InvokeOption) (*ListKustoPoolLanguageExtensionsResult, error) {
	var rv ListKustoPoolLanguageExtensionsResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:listKustoPoolLanguageExtensions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListKustoPoolLanguageExtensionsArgs struct {
	KustoPoolName     string `pulumi:"kustoPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListKustoPoolLanguageExtensionsResult struct {
	Value []LanguageExtensionResponse `pulumi:"value"`
}
