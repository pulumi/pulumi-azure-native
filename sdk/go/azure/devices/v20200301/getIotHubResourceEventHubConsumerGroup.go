


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotHubResourceEventHubConsumerGroup(ctx *pulumi.Context, args *LookupIotHubResourceEventHubConsumerGroupArgs, opts ...pulumi.InvokeOption) (*LookupIotHubResourceEventHubConsumerGroupResult, error) {
	var rv LookupIotHubResourceEventHubConsumerGroupResult
	err := ctx.Invoke("azure-native:devices/v20200301:getIotHubResourceEventHubConsumerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotHubResourceEventHubConsumerGroupArgs struct {
	EventHubEndpointName string `pulumi:"eventHubEndpointName"`
	Name                 string `pulumi:"name"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	ResourceName         string `pulumi:"resourceName"`
}


type LookupIotHubResourceEventHubConsumerGroupResult struct {
	Etag       string            `pulumi:"etag"`
	Id         string            `pulumi:"id"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}
