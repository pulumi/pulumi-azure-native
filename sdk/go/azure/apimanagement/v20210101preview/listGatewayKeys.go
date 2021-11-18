


package v20210101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListGatewayKeys(ctx *pulumi.Context, args *ListGatewayKeysArgs, opts ...pulumi.InvokeOption) (*ListGatewayKeysResult, error) {
	var rv ListGatewayKeysResult
	err := ctx.Invoke("azure-native:apimanagement/v20210101preview:listGatewayKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListGatewayKeysArgs struct {
	GatewayId         string `pulumi:"gatewayId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListGatewayKeysResult struct {
	Primary   *string `pulumi:"primary"`
	Secondary *string `pulumi:"secondary"`
}
