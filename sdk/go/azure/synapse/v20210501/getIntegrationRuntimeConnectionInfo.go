


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIntegrationRuntimeConnectionInfo(ctx *pulumi.Context, args *GetIntegrationRuntimeConnectionInfoArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeConnectionInfoResult, error) {
	var rv GetIntegrationRuntimeConnectionInfoResult
	err := ctx.Invoke("azure-native:synapse/v20210501:getIntegrationRuntimeConnectionInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeConnectionInfoArgs struct {
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	WorkspaceName          string `pulumi:"workspaceName"`
}


type GetIntegrationRuntimeConnectionInfoResult struct {
	HostServiceUri         string `pulumi:"hostServiceUri"`
	IdentityCertThumbprint string `pulumi:"identityCertThumbprint"`
	IsIdentityCertExprired bool   `pulumi:"isIdentityCertExprired"`
	PublicKey              string `pulumi:"publicKey"`
	ServiceToken           string `pulumi:"serviceToken"`
	Version                string `pulumi:"version"`
}
