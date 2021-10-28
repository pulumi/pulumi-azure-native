


package v20170901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIntegrationRuntimeStatus(ctx *pulumi.Context, args *GetIntegrationRuntimeStatusArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeStatusResult, error) {
	var rv GetIntegrationRuntimeStatusResult
	err := ctx.Invoke("azure-native:datafactory/v20170901preview:getIntegrationRuntimeStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeStatusArgs struct {
	FactoryName            string `pulumi:"factoryName"`
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type GetIntegrationRuntimeStatusResult struct {
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
}
