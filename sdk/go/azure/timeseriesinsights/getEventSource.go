


package timeseriesinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupEventSource(ctx *pulumi.Context, args *LookupEventSourceArgs, opts ...pulumi.InvokeOption) (*LookupEventSourceResult, error) {
	var rv LookupEventSourceResult
	err := ctx.Invoke("azure-native:timeseriesinsights:getEventSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventSourceArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	EventSourceName   string `pulumi:"eventSourceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEventSourceResult struct {
	Id       string            `pulumi:"id"`
	Kind     string            `pulumi:"kind"`
	Location string            `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}
