


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupQueue(ctx *pulumi.Context, args *LookupQueueArgs, opts ...pulumi.InvokeOption) (*LookupQueueResult, error) {
	var rv LookupQueueResult
	err := ctx.Invoke("azure-native:storage/v20210101:getQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueArgs struct {
	AccountName       string `pulumi:"accountName"`
	QueueName         string `pulumi:"queueName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupQueueResult struct {
	ApproximateMessageCount int               `pulumi:"approximateMessageCount"`
	Id                      string            `pulumi:"id"`
	Metadata                map[string]string `pulumi:"metadata"`
	Name                    string            `pulumi:"name"`
	Type                    string            `pulumi:"type"`
}
