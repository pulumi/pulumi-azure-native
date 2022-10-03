


package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetChannelFullUrl(ctx *pulumi.Context, args *GetChannelFullUrlArgs, opts ...pulumi.InvokeOption) (*GetChannelFullUrlResult, error) {
	var rv GetChannelFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid/v20220615:getChannelFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetChannelFullUrlArgs struct {
	ChannelName          string `pulumi:"channelName"`
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type GetChannelFullUrlResult struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
}

func GetChannelFullUrlOutput(ctx *pulumi.Context, args GetChannelFullUrlOutputArgs, opts ...pulumi.InvokeOption) GetChannelFullUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetChannelFullUrlResult, error) {
			args := v.(GetChannelFullUrlArgs)
			r, err := GetChannelFullUrl(ctx, &args, opts...)
			var s GetChannelFullUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetChannelFullUrlResultOutput)
}

type GetChannelFullUrlOutputArgs struct {
	ChannelName          pulumi.StringInput `pulumi:"channelName"`
	PartnerNamespaceName pulumi.StringInput `pulumi:"partnerNamespaceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetChannelFullUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetChannelFullUrlArgs)(nil)).Elem()
}


type GetChannelFullUrlResultOutput struct{ *pulumi.OutputState }

func (GetChannelFullUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetChannelFullUrlResult)(nil)).Elem()
}

func (o GetChannelFullUrlResultOutput) ToGetChannelFullUrlResultOutput() GetChannelFullUrlResultOutput {
	return o
}

func (o GetChannelFullUrlResultOutput) ToGetChannelFullUrlResultOutputWithContext(ctx context.Context) GetChannelFullUrlResultOutput {
	return o
}

func (o GetChannelFullUrlResultOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetChannelFullUrlResult) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetChannelFullUrlResultOutput{})
}
