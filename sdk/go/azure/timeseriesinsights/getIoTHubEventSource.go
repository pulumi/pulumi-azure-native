


package timeseriesinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIoTHubEventSource(ctx *pulumi.Context, args *LookupIoTHubEventSourceArgs, opts ...pulumi.InvokeOption) (*LookupIoTHubEventSourceResult, error) {
	var rv LookupIoTHubEventSourceResult
	err := ctx.Invoke("azure-native:timeseriesinsights:getIoTHubEventSource", args, &rv, opts...)
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
	ConsumerGroupName     string                  `pulumi:"consumerGroupName"`
	CreationTime          string                  `pulumi:"creationTime"`
	EventSourceResourceId string                  `pulumi:"eventSourceResourceId"`
	Id                    string                  `pulumi:"id"`
	IotHubName            string                  `pulumi:"iotHubName"`
	KeyName               string                  `pulumi:"keyName"`
	Kind                  string                  `pulumi:"kind"`
	LocalTimestamp        *LocalTimestampResponse `pulumi:"localTimestamp"`
	Location              string                  `pulumi:"location"`
	Name                  string                  `pulumi:"name"`
	ProvisioningState     string                  `pulumi:"provisioningState"`
	Tags                  map[string]string       `pulumi:"tags"`
	Time                  *string                 `pulumi:"time"`
	TimestampPropertyName *string                 `pulumi:"timestampPropertyName"`
	Type                  string                  `pulumi:"type"`
}

func LookupIoTHubEventSourceOutput(ctx *pulumi.Context, args LookupIoTHubEventSourceOutputArgs, opts ...pulumi.InvokeOption) LookupIoTHubEventSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIoTHubEventSourceResult, error) {
			args := v.(LookupIoTHubEventSourceArgs)
			r, err := LookupIoTHubEventSource(ctx, &args, opts...)
			var s LookupIoTHubEventSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIoTHubEventSourceResultOutput)
}

type LookupIoTHubEventSourceOutputArgs struct {
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	EventSourceName   pulumi.StringInput `pulumi:"eventSourceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIoTHubEventSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoTHubEventSourceArgs)(nil)).Elem()
}


type LookupIoTHubEventSourceResultOutput struct{ *pulumi.OutputState }

func (LookupIoTHubEventSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoTHubEventSourceResult)(nil)).Elem()
}

func (o LookupIoTHubEventSourceResultOutput) ToLookupIoTHubEventSourceResultOutput() LookupIoTHubEventSourceResultOutput {
	return o
}

func (o LookupIoTHubEventSourceResultOutput) ToLookupIoTHubEventSourceResultOutputWithContext(ctx context.Context) LookupIoTHubEventSourceResultOutput {
	return o
}

func (o LookupIoTHubEventSourceResultOutput) ConsumerGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) string { return v.ConsumerGroupName }).(pulumi.StringOutput)
}

func (o LookupIoTHubEventSourceResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupIoTHubEventSourceResultOutput) EventSourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) string { return v.EventSourceResourceId }).(pulumi.StringOutput)
}

func (o LookupIoTHubEventSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIoTHubEventSourceResultOutput) IotHubName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) string { return v.IotHubName }).(pulumi.StringOutput)
}

func (o LookupIoTHubEventSourceResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o LookupIoTHubEventSourceResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupIoTHubEventSourceResultOutput) LocalTimestamp() LocalTimestampResponsePtrOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) *LocalTimestampResponse { return v.LocalTimestamp }).(LocalTimestampResponsePtrOutput)
}

func (o LookupIoTHubEventSourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupIoTHubEventSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIoTHubEventSourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupIoTHubEventSourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIoTHubEventSourceResultOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) *string { return v.Time }).(pulumi.StringPtrOutput)
}

func (o LookupIoTHubEventSourceResultOutput) TimestampPropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) *string { return v.TimestampPropertyName }).(pulumi.StringPtrOutput)
}

func (o LookupIoTHubEventSourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTHubEventSourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIoTHubEventSourceResultOutput{})
}
