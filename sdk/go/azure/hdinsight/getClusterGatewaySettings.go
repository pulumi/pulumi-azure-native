


package hdinsight

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetClusterGatewaySettings(ctx *pulumi.Context, args *GetClusterGatewaySettingsArgs, opts ...pulumi.InvokeOption) (*GetClusterGatewaySettingsResult, error) {
	var rv GetClusterGatewaySettingsResult
	err := ctx.Invoke("azure-native:hdinsight:getClusterGatewaySettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetClusterGatewaySettingsArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetClusterGatewaySettingsResult struct {
	IsCredentialEnabled string `pulumi:"isCredentialEnabled"`
	Password            string `pulumi:"password"`
	UserName            string `pulumi:"userName"`
}
