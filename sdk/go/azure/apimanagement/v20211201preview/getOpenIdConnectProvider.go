


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOpenIdConnectProvider(ctx *pulumi.Context, args *LookupOpenIdConnectProviderArgs, opts ...pulumi.InvokeOption) (*LookupOpenIdConnectProviderResult, error) {
	var rv LookupOpenIdConnectProviderResult
	err := ctx.Invoke("azure-native:apimanagement/v20211201preview:getOpenIdConnectProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOpenIdConnectProviderArgs struct {
	Opid              string `pulumi:"opid"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupOpenIdConnectProviderResult struct {
	ClientId         string  `pulumi:"clientId"`
	ClientSecret     *string `pulumi:"clientSecret"`
	Description      *string `pulumi:"description"`
	DisplayName      string  `pulumi:"displayName"`
	Id               string  `pulumi:"id"`
	MetadataEndpoint string  `pulumi:"metadataEndpoint"`
	Name             string  `pulumi:"name"`
	Type             string  `pulumi:"type"`
}

func LookupOpenIdConnectProviderOutput(ctx *pulumi.Context, args LookupOpenIdConnectProviderOutputArgs, opts ...pulumi.InvokeOption) LookupOpenIdConnectProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOpenIdConnectProviderResult, error) {
			args := v.(LookupOpenIdConnectProviderArgs)
			r, err := LookupOpenIdConnectProvider(ctx, &args, opts...)
			var s LookupOpenIdConnectProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOpenIdConnectProviderResultOutput)
}

type LookupOpenIdConnectProviderOutputArgs struct {
	Opid              pulumi.StringInput `pulumi:"opid"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupOpenIdConnectProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOpenIdConnectProviderArgs)(nil)).Elem()
}


type LookupOpenIdConnectProviderResultOutput struct{ *pulumi.OutputState }

func (LookupOpenIdConnectProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOpenIdConnectProviderResult)(nil)).Elem()
}

func (o LookupOpenIdConnectProviderResultOutput) ToLookupOpenIdConnectProviderResultOutput() LookupOpenIdConnectProviderResultOutput {
	return o
}

func (o LookupOpenIdConnectProviderResultOutput) ToLookupOpenIdConnectProviderResultOutputWithContext(ctx context.Context) LookupOpenIdConnectProviderResultOutput {
	return o
}

func (o LookupOpenIdConnectProviderResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o LookupOpenIdConnectProviderResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o LookupOpenIdConnectProviderResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupOpenIdConnectProviderResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupOpenIdConnectProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOpenIdConnectProviderResultOutput) MetadataEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) string { return v.MetadataEndpoint }).(pulumi.StringOutput)
}

func (o LookupOpenIdConnectProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOpenIdConnectProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOpenIdConnectProviderResultOutput{})
}
