


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetEventSubscriptionFullUrlResult, error) {
	var rv GetEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid/v20210601preview:getEventSubscriptionFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEventSubscriptionFullUrlArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	Scope                 string `pulumi:"scope"`
}


type GetEventSubscriptionFullUrlResult struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
}

func GetEventSubscriptionFullUrlOutput(ctx *pulumi.Context, args GetEventSubscriptionFullUrlOutputArgs, opts ...pulumi.InvokeOption) GetEventSubscriptionFullUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEventSubscriptionFullUrlResult, error) {
			args := v.(GetEventSubscriptionFullUrlArgs)
			r, err := GetEventSubscriptionFullUrl(ctx, &args, opts...)
			var s GetEventSubscriptionFullUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEventSubscriptionFullUrlResultOutput)
}

type GetEventSubscriptionFullUrlOutputArgs struct {
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	Scope                 pulumi.StringInput `pulumi:"scope"`
}

func (GetEventSubscriptionFullUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEventSubscriptionFullUrlArgs)(nil)).Elem()
}


type GetEventSubscriptionFullUrlResultOutput struct{ *pulumi.OutputState }

func (GetEventSubscriptionFullUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEventSubscriptionFullUrlResult)(nil)).Elem()
}

func (o GetEventSubscriptionFullUrlResultOutput) ToGetEventSubscriptionFullUrlResultOutput() GetEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetEventSubscriptionFullUrlResultOutput) ToGetEventSubscriptionFullUrlResultOutputWithContext(ctx context.Context) GetEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetEventSubscriptionFullUrlResultOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEventSubscriptionFullUrlResult) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEventSubscriptionFullUrlResultOutput{})
}
