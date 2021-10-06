


package v20160801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppSyncFunctionTriggers(ctx *pulumi.Context, args *ListWebAppSyncFunctionTriggersArgs, opts ...pulumi.InvokeOption) (*ListWebAppSyncFunctionTriggersResult, error) {
	var rv ListWebAppSyncFunctionTriggersResult
	err := ctx.Invoke("azure-native:web/v20160801:listWebAppSyncFunctionTriggers", args, &rv, opts...)
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
	Id         string  `pulumi:"id"`
	Key        *string `pulumi:"key"`
	Kind       *string `pulumi:"kind"`
	Name       string  `pulumi:"name"`
	TriggerUrl *string `pulumi:"triggerUrl"`
	Type       string  `pulumi:"type"`
}
