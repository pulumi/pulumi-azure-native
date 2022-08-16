


package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDomainTopicEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetDomainTopicEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetDomainTopicEventSubscriptionDeliveryAttributesResult, error) {
	var rv GetDomainTopicEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid/v20211015preview:getDomainTopicEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDomainTopicEventSubscriptionDeliveryAttributesArgs struct {
	DomainName            string `pulumi:"domainName"`
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TopicName             string `pulumi:"topicName"`
}


type GetDomainTopicEventSubscriptionDeliveryAttributesResult struct {
	Value []interface{} `pulumi:"value"`
}

func GetDomainTopicEventSubscriptionDeliveryAttributesOutput(ctx *pulumi.Context, args GetDomainTopicEventSubscriptionDeliveryAttributesOutputArgs, opts ...pulumi.InvokeOption) GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainTopicEventSubscriptionDeliveryAttributesResult, error) {
			args := v.(GetDomainTopicEventSubscriptionDeliveryAttributesArgs)
			r, err := GetDomainTopicEventSubscriptionDeliveryAttributes(ctx, &args, opts...)
			var s GetDomainTopicEventSubscriptionDeliveryAttributesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput)
}

type GetDomainTopicEventSubscriptionDeliveryAttributesOutputArgs struct {
	DomainName            pulumi.StringInput `pulumi:"domainName"`
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	TopicName             pulumi.StringInput `pulumi:"topicName"`
}

func (GetDomainTopicEventSubscriptionDeliveryAttributesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainTopicEventSubscriptionDeliveryAttributesArgs)(nil)).Elem()
}


type GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput struct{ *pulumi.OutputState }

func (GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainTopicEventSubscriptionDeliveryAttributesResult)(nil)).Elem()
}

func (o GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetDomainTopicEventSubscriptionDeliveryAttributesResultOutput() GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetDomainTopicEventSubscriptionDeliveryAttributesResultOutputWithContext(ctx context.Context) GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetDomainTopicEventSubscriptionDeliveryAttributesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput{})
}
