


package v20160203

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupIotHubResourceEventHubConsumerGroup(ctx *pulumi.Context, args *LookupIotHubResourceEventHubConsumerGroupArgs, opts ...pulumi.InvokeOption) (*LookupIotHubResourceEventHubConsumerGroupResult, error) {
	var rv LookupIotHubResourceEventHubConsumerGroupResult
	err := ctx.Invoke("azure-native:devices/v20160203:getIotHubResourceEventHubConsumerGroup", args, &rv, opts...)
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

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotHubResourceEventHubConsumerGroupResultOutput{})
}
