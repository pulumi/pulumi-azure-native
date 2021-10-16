


package v20160801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppSyncFunctionTriggersSlot(ctx *pulumi.Context, args *ListWebAppSyncFunctionTriggersSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppSyncFunctionTriggersSlotResult, error) {
	var rv ListWebAppSyncFunctionTriggersSlotResult
	err := ctx.Invoke("azure-native:web/v20160801:listWebAppSyncFunctionTriggersSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSyncFunctionTriggersSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppSyncFunctionTriggersSlotResult struct {
	Id         string  `pulumi:"id"`
	Key        *string `pulumi:"key"`
	Kind       *string `pulumi:"kind"`
	Name       string  `pulumi:"name"`
	TriggerUrl *string `pulumi:"triggerUrl"`
	Type       string  `pulumi:"type"`
}
