


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppSyncFunctionTriggers(ctx *pulumi.Context, args *ListWebAppSyncFunctionTriggersArgs, opts ...pulumi.InvokeOption) (*ListWebAppSyncFunctionTriggersResult, error) {
	var rv ListWebAppSyncFunctionTriggersResult
	err := ctx.Invoke("azure-native:web/v20201201:listWebAppSyncFunctionTriggers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSyncFunctionTriggersArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppSyncFunctionTriggersResult struct {
	Key        *string `pulumi:"key"`
	TriggerUrl *string `pulumi:"triggerUrl"`
}
