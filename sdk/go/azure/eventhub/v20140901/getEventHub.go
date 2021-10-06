


package v20140901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventHub(ctx *pulumi.Context, args *LookupEventHubArgs, opts ...pulumi.InvokeOption) (*LookupEventHubResult, error) {
	var rv LookupEventHubResult
	err := ctx.Invoke("azure-native:eventhub/v20140901:getEventHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubArgs struct {
	EventHubName      string `pulumi:"eventHubName"`
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEventHubResult struct {
	CreatedAt              string   `pulumi:"createdAt"`
	Id                     string   `pulumi:"id"`
	Location               *string  `pulumi:"location"`
	MessageRetentionInDays *float64 `pulumi:"messageRetentionInDays"`
	Name                   string   `pulumi:"name"`
	PartitionCount         *float64 `pulumi:"partitionCount"`
	PartitionIds           []string `pulumi:"partitionIds"`
	Status                 *string  `pulumi:"status"`
	Type                   string   `pulumi:"type"`
	UpdatedAt              string   `pulumi:"updatedAt"`
}
