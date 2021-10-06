


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationRuntimeAuthKey(ctx *pulumi.Context, args *ListIntegrationRuntimeAuthKeyArgs, opts ...pulumi.InvokeOption) (*ListIntegrationRuntimeAuthKeyResult, error) {
	var rv ListIntegrationRuntimeAuthKeyResult
	err := ctx.Invoke("azure-native:synapse/v20201201:listIntegrationRuntimeAuthKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationRuntimeAuthKeyArgs struct {
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	WorkspaceName          string `pulumi:"workspaceName"`
}


type ListIntegrationRuntimeAuthKeyResult struct {
	AuthKey1 *string `pulumi:"authKey1"`
	AuthKey2 *string `pulumi:"authKey2"`
}
