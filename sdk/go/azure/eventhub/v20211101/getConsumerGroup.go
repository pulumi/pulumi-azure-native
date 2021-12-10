


package v20211101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConsumerGroup(ctx *pulumi.Context, args *LookupConsumerGroupArgs, opts ...pulumi.InvokeOption) (*LookupConsumerGroupResult, error) {
	var rv LookupConsumerGroupResult
	err := ctx.Invoke("azure-native:eventhub/v20211101:getConsumerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConsumerGroupArgs struct {
	ConsumerGroupName string `pulumi:"consumerGroupName"`
	EventHubName      string `pulumi:"eventHubName"`
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConsumerGroupResult struct {
	CreatedAt    string             `pulumi:"createdAt"`
	Id           string             `pulumi:"id"`
	Location     string             `pulumi:"location"`
	Name         string             `pulumi:"name"`
	SystemData   SystemDataResponse `pulumi:"systemData"`
	Type         string             `pulumi:"type"`
	UpdatedAt    string             `pulumi:"updatedAt"`
	UserMetadata *string            `pulumi:"userMetadata"`
}
