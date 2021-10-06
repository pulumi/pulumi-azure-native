


package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIntegrationRuntimeConnectionInfo(ctx *pulumi.Context, args *GetIntegrationRuntimeConnectionInfoArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeConnectionInfoResult, error) {
	var rv GetIntegrationRuntimeConnectionInfoResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getIntegrationRuntimeConnectionInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeConnectionInfoArgs struct {
	FactoryName            string `pulumi:"factoryName"`
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type GetIntegrationRuntimeConnectionInfoResult struct {
	HostServiceUri         string `pulumi:"hostServiceUri"`
	IdentityCertThumbprint string `pulumi:"identityCertThumbprint"`
	IsIdentityCertExprired bool   `pulumi:"isIdentityCertExprired"`
	PublicKey              string `pulumi:"publicKey"`
	ServiceToken           string `pulumi:"serviceToken"`
	Version                string `pulumi:"version"`
}
