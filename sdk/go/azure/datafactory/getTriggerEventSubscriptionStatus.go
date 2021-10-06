


package datafactory

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTriggerEventSubscriptionStatus(ctx *pulumi.Context, args *GetTriggerEventSubscriptionStatusArgs, opts ...pulumi.InvokeOption) (*GetTriggerEventSubscriptionStatusResult, error) {
	var rv GetTriggerEventSubscriptionStatusResult
	err := ctx.Invoke("azure-native:datafactory:getTriggerEventSubscriptionStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTriggerEventSubscriptionStatusArgs struct {
	FactoryName       string `pulumi:"factoryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TriggerName       string `pulumi:"triggerName"`
}


type GetTriggerEventSubscriptionStatusResult struct {
	Status      string `pulumi:"status"`
	TriggerName string `pulumi:"triggerName"`
}
