


package timeseriesinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventHubEventSource(ctx *pulumi.Context, args *LookupEventHubEventSourceArgs, opts ...pulumi.InvokeOption) (*LookupEventHubEventSourceResult, error) {
	var rv LookupEventHubEventSourceResult
	err := ctx.Invoke("azure-native:timeseriesinsights:getEventHubEventSource", args, &rv, opts...)
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

func LookupEventHubEventSourceOutput(ctx *pulumi.Context, args LookupEventHubEventSourceOutputArgs, opts ...pulumi.InvokeOption) LookupEventHubEventSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventHubEventSourceResult, error) {
			args := v.(LookupEventHubEventSourceArgs)
			r, err := LookupEventHubEventSource(ctx, &args, opts...)
			var s LookupEventHubEventSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventHubEventSourceResultOutput)
}

type LookupEventHubEventSourceOutputArgs struct {
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	EventSourceName   pulumi.StringInput `pulumi:"eventSourceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventHubEventSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubEventSourceArgs)(nil)).Elem()
}


type LookupEventHubEventSourceResultOutput struct{ *pulumi.OutputState }

func (LookupEventHubEventSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubEventSourceResult)(nil)).Elem()
}

func (o LookupEventHubEventSourceResultOutput) ToLookupEventHubEventSourceResultOutput() LookupEventHubEventSourceResultOutput {
	return o
}

func (o LookupEventHubEventSourceResultOutput) ToLookupEventHubEventSourceResultOutputWithContext(ctx context.Context) LookupEventHubEventSourceResultOutput {
	return o
}

func (o LookupEventHubEventSourceResultOutput) ConsumerGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.ConsumerGroupName }).(pulumi.StringOutput)
}

func (o LookupEventHubEventSourceResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupEventHubEventSourceResultOutput) EventHubName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.EventHubName }).(pulumi.StringOutput)
}

func (o LookupEventHubEventSourceResultOutput) EventSourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.EventSourceResourceId }).(pulumi.StringOutput)
}

func (o LookupEventHubEventSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEventHubEventSourceResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o LookupEventHubEventSourceResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupEventHubEventSourceResultOutput) LocalTimestamp() LocalTimestampResponsePtrOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) *LocalTimestampResponse { return v.LocalTimestamp }).(LocalTimestampResponsePtrOutput)
}

func (o LookupEventHubEventSourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupEventHubEventSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEventHubEventSourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupEventHubEventSourceResultOutput) ServiceBusNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.ServiceBusNamespace }).(pulumi.StringOutput)
}

func (o LookupEventHubEventSourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEventHubEventSourceResultOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) *string { return v.Time }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubEventSourceResultOutput) TimestampPropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) *string { return v.TimestampPropertyName }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubEventSourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventHubEventSourceResultOutput{})
}
