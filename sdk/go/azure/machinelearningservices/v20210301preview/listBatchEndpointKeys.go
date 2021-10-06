


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBatchEndpointKeys(ctx *pulumi.Context, args *ListBatchEndpointKeysArgs, opts ...pulumi.InvokeOption) (*ListBatchEndpointKeysResult, error) {
	var rv ListBatchEndpointKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:listBatchEndpointKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBatchEndpointKeysArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListBatchEndpointKeysResult struct {
	PrimaryKey   *string `pulumi:"primaryKey"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}
