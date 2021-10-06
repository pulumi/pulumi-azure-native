


package v20170901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationRuntimeAuthKeys(ctx *pulumi.Context, args *ListIntegrationRuntimeAuthKeysArgs, opts ...pulumi.InvokeOption) (*ListIntegrationRuntimeAuthKeysResult, error) {
	var rv ListIntegrationRuntimeAuthKeysResult
	err := ctx.Invoke("azure-native:datafactory/v20170901preview:listIntegrationRuntimeAuthKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationRuntimeAuthKeysArgs struct {
	FactoryName            string `pulumi:"factoryName"`
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type ListIntegrationRuntimeAuthKeysResult struct {
	AuthKey1 *string `pulumi:"authKey1"`
	AuthKey2 *string `pulumi:"authKey2"`
}
