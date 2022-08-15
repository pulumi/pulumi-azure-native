


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayCustomDomain(ctx *pulumi.Context, args *LookupGatewayCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupGatewayCustomDomainResult, error) {
	var rv LookupGatewayCustomDomainResult
	err := ctx.Invoke("azure-native:appplatform/v20220301preview:getGatewayCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayCustomDomainArgs struct {
	DomainName        string `pulumi:"domainName"`
	GatewayName       string `pulumi:"gatewayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupGatewayCustomDomainResult struct {
	Id         string                                `pulumi:"id"`
	Name       string                                `pulumi:"name"`
	Properties GatewayCustomDomainPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                    `pulumi:"systemData"`
	Type       string                                `pulumi:"type"`
}

func LookupGatewayCustomDomainOutput(ctx *pulumi.Context, args LookupGatewayCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayCustomDomainResult, error) {
			args := v.(LookupGatewayCustomDomainArgs)
			r, err := LookupGatewayCustomDomain(ctx, &args, opts...)
			var s LookupGatewayCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayCustomDomainResultOutput)
}

type LookupGatewayCustomDomainOutputArgs struct {
	DomainName        pulumi.StringInput `pulumi:"domainName"`
	GatewayName       pulumi.StringInput `pulumi:"gatewayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGatewayCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCustomDomainArgs)(nil)).Elem()
}


type LookupGatewayCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCustomDomainResult)(nil)).Elem()
}

func (o LookupGatewayCustomDomainResultOutput) ToLookupGatewayCustomDomainResultOutput() LookupGatewayCustomDomainResultOutput {
	return o
}

func (o LookupGatewayCustomDomainResultOutput) ToLookupGatewayCustomDomainResultOutputWithContext(ctx context.Context) LookupGatewayCustomDomainResultOutput {
	return o
}

func (o LookupGatewayCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGatewayCustomDomainResultOutput) Properties() GatewayCustomDomainPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGatewayCustomDomainResult) GatewayCustomDomainPropertiesResponse { return v.Properties }).(GatewayCustomDomainPropertiesResponseOutput)
}

func (o LookupGatewayCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGatewayCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupGatewayCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayCustomDomainResultOutput{})
}
