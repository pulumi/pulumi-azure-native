


package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPartnerTopicEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetPartnerTopicEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetPartnerTopicEventSubscriptionFullUrlResult, error) {
	var rv GetPartnerTopicEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid/v20211015preview:getPartnerTopicEventSubscriptionFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetPartnerTopicEventSubscriptionFullUrlArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	PartnerTopicName      string `pulumi:"partnerTopicName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type GetPartnerTopicEventSubscriptionFullUrlResult struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
}

func GetPartnerTopicEventSubscriptionFullUrlOutput(ctx *pulumi.Context, args GetPartnerTopicEventSubscriptionFullUrlOutputArgs, opts ...pulumi.InvokeOption) GetPartnerTopicEventSubscriptionFullUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPartnerTopicEventSubscriptionFullUrlResult, error) {
			args := v.(GetPartnerTopicEventSubscriptionFullUrlArgs)
			r, err := GetPartnerTopicEventSubscriptionFullUrl(ctx, &args, opts...)
			var s GetPartnerTopicEventSubscriptionFullUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPartnerTopicEventSubscriptionFullUrlResultOutput)
}

type GetPartnerTopicEventSubscriptionFullUrlOutputArgs struct {
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	PartnerTopicName      pulumi.StringInput `pulumi:"partnerTopicName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetPartnerTopicEventSubscriptionFullUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPartnerTopicEventSubscriptionFullUrlArgs)(nil)).Elem()
}


type GetPartnerTopicEventSubscriptionFullUrlResultOutput struct{ *pulumi.OutputState }

func (GetPartnerTopicEventSubscriptionFullUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPartnerTopicEventSubscriptionFullUrlResult)(nil)).Elem()
}

func (o GetPartnerTopicEventSubscriptionFullUrlResultOutput) ToGetPartnerTopicEventSubscriptionFullUrlResultOutput() GetPartnerTopicEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetPartnerTopicEventSubscriptionFullUrlResultOutput) ToGetPartnerTopicEventSubscriptionFullUrlResultOutputWithContext(ctx context.Context) GetPartnerTopicEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetPartnerTopicEventSubscriptionFullUrlResultOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPartnerTopicEventSubscriptionFullUrlResult) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPartnerTopicEventSubscriptionFullUrlResultOutput{})
}
