


package v20220430preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotHubResourceEventHubConsumerGroup(ctx *pulumi.Context, args *LookupIotHubResourceEventHubConsumerGroupArgs, opts ...pulumi.InvokeOption) (*LookupIotHubResourceEventHubConsumerGroupResult, error) {
	var rv LookupIotHubResourceEventHubConsumerGroupResult
	err := ctx.Invoke("azure-native:devices/v20220430preview:getIotHubResourceEventHubConsumerGroup", args, &rv, opts...)
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
	Etag       string      `pulumi:"etag"`
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

func LookupIotHubResourceEventHubConsumerGroupOutput(ctx *pulumi.Context, args LookupIotHubResourceEventHubConsumerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupIotHubResourceEventHubConsumerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotHubResourceEventHubConsumerGroupResult, error) {
			args := v.(LookupIotHubResourceEventHubConsumerGroupArgs)
			r, err := LookupIotHubResourceEventHubConsumerGroup(ctx, &args, opts...)
			var s LookupIotHubResourceEventHubConsumerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotHubResourceEventHubConsumerGroupResultOutput)
}

type LookupIotHubResourceEventHubConsumerGroupOutputArgs struct {
	EventHubEndpointName pulumi.StringInput `pulumi:"eventHubEndpointName"`
	Name                 pulumi.StringInput `pulumi:"name"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName         pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupIotHubResourceEventHubConsumerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubResourceEventHubConsumerGroupArgs)(nil)).Elem()
}


type LookupIotHubResourceEventHubConsumerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupIotHubResourceEventHubConsumerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubResourceEventHubConsumerGroupResult)(nil)).Elem()
}

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) ToLookupIotHubResourceEventHubConsumerGroupResultOutput() LookupIotHubResourceEventHubConsumerGroupResultOutput {
	return o
}

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) ToLookupIotHubResourceEventHubConsumerGroupResultOutputWithContext(ctx context.Context) LookupIotHubResourceEventHubConsumerGroupResultOutput {
	return o
}

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotHubResourceEventHubConsumerGroupResultOutput{})
}
