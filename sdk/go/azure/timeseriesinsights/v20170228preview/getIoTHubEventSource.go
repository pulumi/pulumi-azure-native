


package v20170228preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIoTHubEventSource(ctx *pulumi.Context, args *LookupIoTHubEventSourceArgs, opts ...pulumi.InvokeOption) (*LookupIoTHubEventSourceResult, error) {
	var rv LookupIoTHubEventSourceResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20170228preview:getIoTHubEventSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIoTHubEventSourceArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	EventSourceName   string `pulumi:"eventSourceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupIoTHubEventSourceResult struct {
	ConsumerGroupName     string            `pulumi:"consumerGroupName"`
	CreationTime          string            `pulumi:"creationTime"`
	EventSourceResourceId string            `pulumi:"eventSourceResourceId"`
	Id                    string            `pulumi:"id"`
	IotHubName            string            `pulumi:"iotHubName"`
	KeyName               string            `pulumi:"keyName"`
	Kind                  string            `pulumi:"kind"`
	Location              string            `pulumi:"location"`
	Name                  string            `pulumi:"name"`
	ProvisioningState     string            `pulumi:"provisioningState"`
	Tags                  map[string]string `pulumi:"tags"`
	TimestampPropertyName *string           `pulumi:"timestampPropertyName"`
	Type                  string            `pulumi:"type"`
}
