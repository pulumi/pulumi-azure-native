


package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDomainTopicEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetDomainTopicEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetDomainTopicEventSubscriptionFullUrlResult, error) {
	var rv GetDomainTopicEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid:getDomainTopicEventSubscriptionFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDomainTopicEventSubscriptionFullUrlArgs struct {
	DomainName            string `pulumi:"domainName"`
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TopicName             string `pulumi:"topicName"`
}


type GetDomainTopicEventSubscriptionFullUrlResult struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
}

func GetDomainTopicEventSubscriptionFullUrlOutput(ctx *pulumi.Context, args GetDomainTopicEventSubscriptionFullUrlOutputArgs, opts ...pulumi.InvokeOption) GetDomainTopicEventSubscriptionFullUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainTopicEventSubscriptionFullUrlResult, error) {
			args := v.(GetDomainTopicEventSubscriptionFullUrlArgs)
			r, err := GetDomainTopicEventSubscriptionFullUrl(ctx, &args, opts...)
			var s GetDomainTopicEventSubscriptionFullUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainTopicEventSubscriptionFullUrlResultOutput)
}

type GetDomainTopicEventSubscriptionFullUrlOutputArgs struct {
	DomainName            pulumi.StringInput `pulumi:"domainName"`
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	TopicName             pulumi.StringInput `pulumi:"topicName"`
}

func (GetDomainTopicEventSubscriptionFullUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainTopicEventSubscriptionFullUrlArgs)(nil)).Elem()
}


type GetDomainTopicEventSubscriptionFullUrlResultOutput struct{ *pulumi.OutputState }

func (GetDomainTopicEventSubscriptionFullUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainTopicEventSubscriptionFullUrlResult)(nil)).Elem()
}

func (o GetDomainTopicEventSubscriptionFullUrlResultOutput) ToGetDomainTopicEventSubscriptionFullUrlResultOutput() GetDomainTopicEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetDomainTopicEventSubscriptionFullUrlResultOutput) ToGetDomainTopicEventSubscriptionFullUrlResultOutputWithContext(ctx context.Context) GetDomainTopicEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetDomainTopicEventSubscriptionFullUrlResultOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainTopicEventSubscriptionFullUrlResult) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainTopicEventSubscriptionFullUrlResultOutput{})
}
