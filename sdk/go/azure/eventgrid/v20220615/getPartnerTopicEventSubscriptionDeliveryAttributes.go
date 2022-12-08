


package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPartnerTopicEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetPartnerTopicEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetPartnerTopicEventSubscriptionDeliveryAttributesResult, error) {
	var rv GetPartnerTopicEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid/v20220615:getPartnerTopicEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetPartnerTopicEventSubscriptionDeliveryAttributesArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	PartnerTopicName      string `pulumi:"partnerTopicName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type GetPartnerTopicEventSubscriptionDeliveryAttributesResult struct {
	Value []interface{} `pulumi:"value"`
}

func GetPartnerTopicEventSubscriptionDeliveryAttributesOutput(ctx *pulumi.Context, args GetPartnerTopicEventSubscriptionDeliveryAttributesOutputArgs, opts ...pulumi.InvokeOption) GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPartnerTopicEventSubscriptionDeliveryAttributesResult, error) {
			args := v.(GetPartnerTopicEventSubscriptionDeliveryAttributesArgs)
			r, err := GetPartnerTopicEventSubscriptionDeliveryAttributes(ctx, &args, opts...)
			var s GetPartnerTopicEventSubscriptionDeliveryAttributesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput)
}

type GetPartnerTopicEventSubscriptionDeliveryAttributesOutputArgs struct {
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	PartnerTopicName      pulumi.StringInput `pulumi:"partnerTopicName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetPartnerTopicEventSubscriptionDeliveryAttributesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPartnerTopicEventSubscriptionDeliveryAttributesArgs)(nil)).Elem()
}


type GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput struct{ *pulumi.OutputState }

func (GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPartnerTopicEventSubscriptionDeliveryAttributesResult)(nil)).Elem()
}

func (o GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput() GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetPartnerTopicEventSubscriptionDeliveryAttributesResultOutputWithContext(ctx context.Context) GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetPartnerTopicEventSubscriptionDeliveryAttributesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput{})
}
