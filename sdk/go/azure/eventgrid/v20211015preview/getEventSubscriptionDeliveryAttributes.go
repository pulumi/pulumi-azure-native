


package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetEventSubscriptionDeliveryAttributesResult, error) {
	var rv GetEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid/v20211015preview:getEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEventSubscriptionDeliveryAttributesArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	Scope                 string `pulumi:"scope"`
}


type GetEventSubscriptionDeliveryAttributesResult struct {
	Value []interface{} `pulumi:"value"`
}

func GetEventSubscriptionDeliveryAttributesOutput(ctx *pulumi.Context, args GetEventSubscriptionDeliveryAttributesOutputArgs, opts ...pulumi.InvokeOption) GetEventSubscriptionDeliveryAttributesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEventSubscriptionDeliveryAttributesResult, error) {
			args := v.(GetEventSubscriptionDeliveryAttributesArgs)
			r, err := GetEventSubscriptionDeliveryAttributes(ctx, &args, opts...)
			var s GetEventSubscriptionDeliveryAttributesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEventSubscriptionDeliveryAttributesResultOutput)
}

type GetEventSubscriptionDeliveryAttributesOutputArgs struct {
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	Scope                 pulumi.StringInput `pulumi:"scope"`
}

func (GetEventSubscriptionDeliveryAttributesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEventSubscriptionDeliveryAttributesArgs)(nil)).Elem()
}


type GetEventSubscriptionDeliveryAttributesResultOutput struct{ *pulumi.OutputState }

func (GetEventSubscriptionDeliveryAttributesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEventSubscriptionDeliveryAttributesResult)(nil)).Elem()
}

func (o GetEventSubscriptionDeliveryAttributesResultOutput) ToGetEventSubscriptionDeliveryAttributesResultOutput() GetEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetEventSubscriptionDeliveryAttributesResultOutput) ToGetEventSubscriptionDeliveryAttributesResultOutputWithContext(ctx context.Context) GetEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetEventSubscriptionDeliveryAttributesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetEventSubscriptionDeliveryAttributesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEventSubscriptionDeliveryAttributesResultOutput{})
}
