


package v20160129

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspaceCollection(ctx *pulumi.Context, args *LookupWorkspaceCollectionArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceCollectionResult, error) {
	var rv LookupWorkspaceCollectionResult
	err := ctx.Invoke("azure-native:powerbi/v20160129:getWorkspaceCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceCollectionArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	WorkspaceCollectionName string `pulumi:"workspaceCollectionName"`
}

type LookupWorkspaceCollectionResult struct {
	Id         *string           `pulumi:"id"`
	Location   *string           `pulumi:"location"`
	Name       *string           `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *AzureSkuResponse `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       *string           `pulumi:"type"`
}
