


package v20210630preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventHubEventSource(ctx *pulumi.Context, args *LookupEventHubEventSourceArgs, opts ...pulumi.InvokeOption) (*LookupEventHubEventSourceResult, error) {
	var rv LookupEventHubEventSourceResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20210630preview:getEventHubEventSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubEventSourceArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	EventSourceName   string `pulumi:"eventSourceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEventHubEventSourceResult struct {
	ConsumerGroupName     string                  `pulumi:"consumerGroupName"`
	CreationTime          string                  `pulumi:"creationTime"`
	EventHubName          string                  `pulumi:"eventHubName"`
	EventSourceResourceId string                  `pulumi:"eventSourceResourceId"`
	Id                    string                  `pulumi:"id"`
	KeyName               string                  `pulumi:"keyName"`
	Kind                  string                  `pulumi:"kind"`
	LocalTimestamp        *LocalTimestampResponse `pulumi:"localTimestamp"`
	Location              string                  `pulumi:"location"`
	Name                  string                  `pulumi:"name"`
	ProvisioningState     string                  `pulumi:"provisioningState"`
	ServiceBusNamespace   string                  `pulumi:"serviceBusNamespace"`
	Tags                  map[string]string       `pulumi:"tags"`
	Time                  *string                 `pulumi:"time"`
	TimestampPropertyName *string                 `pulumi:"timestampPropertyName"`
	Type                  string                  `pulumi:"type"`
}
