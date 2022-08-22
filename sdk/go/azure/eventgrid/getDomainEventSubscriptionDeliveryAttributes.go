


package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDomainEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetDomainEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetDomainEventSubscriptionDeliveryAttributesResult, error) {
	var rv GetDomainEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid:getDomainEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDomainEventSubscriptionDeliveryAttributesArgs struct {
	DomainName            string `pulumi:"domainName"`
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type GetDomainEventSubscriptionDeliveryAttributesResult struct {
	Value []interface{} `pulumi:"value"`
}

func GetDomainEventSubscriptionDeliveryAttributesOutput(ctx *pulumi.Context, args GetDomainEventSubscriptionDeliveryAttributesOutputArgs, opts ...pulumi.InvokeOption) GetDomainEventSubscriptionDeliveryAttributesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainEventSubscriptionDeliveryAttributesResult, error) {
			args := v.(GetDomainEventSubscriptionDeliveryAttributesArgs)
			r, err := GetDomainEventSubscriptionDeliveryAttributes(ctx, &args, opts...)
			var s GetDomainEventSubscriptionDeliveryAttributesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainEventSubscriptionDeliveryAttributesResultOutput)
}

type GetDomainEventSubscriptionDeliveryAttributesOutputArgs struct {
	DomainName            pulumi.StringInput `pulumi:"domainName"`
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetDomainEventSubscriptionDeliveryAttributesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainEventSubscriptionDeliveryAttributesArgs)(nil)).Elem()
}


type GetDomainEventSubscriptionDeliveryAttributesResultOutput struct{ *pulumi.OutputState }

func (GetDomainEventSubscriptionDeliveryAttributesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainEventSubscriptionDeliveryAttributesResult)(nil)).Elem()
}

func (o GetDomainEventSubscriptionDeliveryAttributesResultOutput) ToGetDomainEventSubscriptionDeliveryAttributesResultOutput() GetDomainEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetDomainEventSubscriptionDeliveryAttributesResultOutput) ToGetDomainEventSubscriptionDeliveryAttributesResultOutputWithContext(ctx context.Context) GetDomainEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetDomainEventSubscriptionDeliveryAttributesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetDomainEventSubscriptionDeliveryAttributesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainEventSubscriptionDeliveryAttributesResultOutput{})
}
