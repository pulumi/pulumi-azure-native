


package v20170801beta

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListServerGatewayStatus(ctx *pulumi.Context, args *ListServerGatewayStatusArgs, opts ...pulumi.InvokeOption) (*ListServerGatewayStatusResult, error) {
	var rv ListServerGatewayStatusResult
	err := ctx.Invoke("azure-native:analysisservices/v20170801beta:listServerGatewayStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServerGatewayStatusArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type ListServerGatewayStatusResult struct {
	Status *int `pulumi:"status"`
}
