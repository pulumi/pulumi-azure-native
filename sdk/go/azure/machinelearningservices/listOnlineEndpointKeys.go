


package machinelearningservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOnlineEndpointKeys(ctx *pulumi.Context, args *ListOnlineEndpointKeysArgs, opts ...pulumi.InvokeOption) (*ListOnlineEndpointKeysResult, error) {
	var rv ListOnlineEndpointKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices:listOnlineEndpointKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOnlineEndpointKeysArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListOnlineEndpointKeysResult struct {
	PrimaryKey   *string `pulumi:"primaryKey"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}
