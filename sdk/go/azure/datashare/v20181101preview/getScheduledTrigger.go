


package v20181101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScheduledTrigger(ctx *pulumi.Context, args *LookupScheduledTriggerArgs, opts ...pulumi.InvokeOption) (*LookupScheduledTriggerResult, error) {
	var rv LookupScheduledTriggerResult
	err := ctx.Invoke("azure-native:datashare/v20181101preview:getScheduledTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledTriggerArgs struct {
	AccountName           string `pulumi:"accountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
	TriggerName           string `pulumi:"triggerName"`
}


type LookupScheduledTriggerResult struct {
	CreatedAt           string  `pulumi:"createdAt"`
	Id                  string  `pulumi:"id"`
	Kind                string  `pulumi:"kind"`
	Name                string  `pulumi:"name"`
	ProvisioningState   string  `pulumi:"provisioningState"`
	RecurrenceInterval  string  `pulumi:"recurrenceInterval"`
	SynchronizationMode *string `pulumi:"synchronizationMode"`
	SynchronizationTime string  `pulumi:"synchronizationTime"`
	TriggerStatus       string  `pulumi:"triggerStatus"`
	Type                string  `pulumi:"type"`
	UserName            string  `pulumi:"userName"`
}
