


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSystemTopicEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetSystemTopicEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetSystemTopicEventSubscriptionDeliveryAttributesResult, error) {
	var rv GetSystemTopicEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid/v20210601preview:getSystemTopicEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSystemTopicEventSubscriptionDeliveryAttributesArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SystemTopicName       string `pulumi:"systemTopicName"`
}


type GetSystemTopicEventSubscriptionDeliveryAttributesResult struct {
	Value []interface{} `pulumi:"value"`
}

func GetSystemTopicEventSubscriptionDeliveryAttributesOutput(ctx *pulumi.Context, args GetSystemTopicEventSubscriptionDeliveryAttributesOutputArgs, opts ...pulumi.InvokeOption) GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemTopicEventSubscriptionDeliveryAttributesResult, error) {
			args := v.(GetSystemTopicEventSubscriptionDeliveryAttributesArgs)
			r, err := GetSystemTopicEventSubscriptionDeliveryAttributes(ctx, &args, opts...)
			var s GetSystemTopicEventSubscriptionDeliveryAttributesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput)
}

type GetSystemTopicEventSubscriptionDeliveryAttributesOutputArgs struct {
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	SystemTopicName       pulumi.StringInput `pulumi:"systemTopicName"`
}

func (GetSystemTopicEventSubscriptionDeliveryAttributesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemTopicEventSubscriptionDeliveryAttributesArgs)(nil)).Elem()
}


type GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput struct{ *pulumi.OutputState }

func (GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemTopicEventSubscriptionDeliveryAttributesResult)(nil)).Elem()
}

func (o GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetSystemTopicEventSubscriptionDeliveryAttributesResultOutput() GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetSystemTopicEventSubscriptionDeliveryAttributesResultOutputWithContext(ctx context.Context) GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetSystemTopicEventSubscriptionDeliveryAttributesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput{})
}
