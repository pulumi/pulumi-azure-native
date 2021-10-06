


package v20170701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotHubResourceEventHubConsumerGroup(ctx *pulumi.Context, args *LookupIotHubResourceEventHubConsumerGroupArgs, opts ...pulumi.InvokeOption) (*LookupIotHubResourceEventHubConsumerGroupResult, error) {
	var rv LookupIotHubResourceEventHubConsumerGroupResult
	err := ctx.Invoke("azure-native:devices/v20170701:getIotHubResourceEventHubConsumerGroup", args, &rv, opts...)
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
	Id   *string           `pulumi:"id"`
	Name *string           `pulumi:"name"`
	Tags map[string]string `pulumi:"tags"`
}
