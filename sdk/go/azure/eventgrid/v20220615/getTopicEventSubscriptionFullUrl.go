


package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTopicEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetTopicEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetTopicEventSubscriptionFullUrlResult, error) {
	var rv GetTopicEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid/v20220615:getTopicEventSubscriptionFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTopicEventSubscriptionFullUrlArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TopicName             string `pulumi:"topicName"`
}


type GetTopicEventSubscriptionFullUrlResult struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
}

func GetTopicEventSubscriptionFullUrlOutput(ctx *pulumi.Context, args GetTopicEventSubscriptionFullUrlOutputArgs, opts ...pulumi.InvokeOption) GetTopicEventSubscriptionFullUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTopicEventSubscriptionFullUrlResult, error) {
			args := v.(GetTopicEventSubscriptionFullUrlArgs)
			r, err := GetTopicEventSubscriptionFullUrl(ctx, &args, opts...)
			var s GetTopicEventSubscriptionFullUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTopicEventSubscriptionFullUrlResultOutput)
}

type GetTopicEventSubscriptionFullUrlOutputArgs struct {
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	TopicName             pulumi.StringInput `pulumi:"topicName"`
}

func (GetTopicEventSubscriptionFullUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTopicEventSubscriptionFullUrlArgs)(nil)).Elem()
}


type GetTopicEventSubscriptionFullUrlResultOutput struct{ *pulumi.OutputState }

func (GetTopicEventSubscriptionFullUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTopicEventSubscriptionFullUrlResult)(nil)).Elem()
}

func (o GetTopicEventSubscriptionFullUrlResultOutput) ToGetTopicEventSubscriptionFullUrlResultOutput() GetTopicEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetTopicEventSubscriptionFullUrlResultOutput) ToGetTopicEventSubscriptionFullUrlResultOutputWithContext(ctx context.Context) GetTopicEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetTopicEventSubscriptionFullUrlResultOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopicEventSubscriptionFullUrlResult) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTopicEventSubscriptionFullUrlResultOutput{})
}
