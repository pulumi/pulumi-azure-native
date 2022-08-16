


package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDomainEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetDomainEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetDomainEventSubscriptionFullUrlResult, error) {
	var rv GetDomainEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid/v20220615:getDomainEventSubscriptionFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDomainEventSubscriptionFullUrlArgs struct {
	DomainName            string `pulumi:"domainName"`
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type GetDomainEventSubscriptionFullUrlResult struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
}

func GetDomainEventSubscriptionFullUrlOutput(ctx *pulumi.Context, args GetDomainEventSubscriptionFullUrlOutputArgs, opts ...pulumi.InvokeOption) GetDomainEventSubscriptionFullUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainEventSubscriptionFullUrlResult, error) {
			args := v.(GetDomainEventSubscriptionFullUrlArgs)
			r, err := GetDomainEventSubscriptionFullUrl(ctx, &args, opts...)
			var s GetDomainEventSubscriptionFullUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainEventSubscriptionFullUrlResultOutput)
}

type GetDomainEventSubscriptionFullUrlOutputArgs struct {
	DomainName            pulumi.StringInput `pulumi:"domainName"`
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetDomainEventSubscriptionFullUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainEventSubscriptionFullUrlArgs)(nil)).Elem()
}


type GetDomainEventSubscriptionFullUrlResultOutput struct{ *pulumi.OutputState }

func (GetDomainEventSubscriptionFullUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainEventSubscriptionFullUrlResult)(nil)).Elem()
}

func (o GetDomainEventSubscriptionFullUrlResultOutput) ToGetDomainEventSubscriptionFullUrlResultOutput() GetDomainEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetDomainEventSubscriptionFullUrlResultOutput) ToGetDomainEventSubscriptionFullUrlResultOutputWithContext(ctx context.Context) GetDomainEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetDomainEventSubscriptionFullUrlResultOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainEventSubscriptionFullUrlResult) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainEventSubscriptionFullUrlResultOutput{})
}
