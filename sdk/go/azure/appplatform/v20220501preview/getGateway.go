


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGateway(ctx *pulumi.Context, args *LookupGatewayArgs, opts ...pulumi.InvokeOption) (*LookupGatewayResult, error) {
	var rv LookupGatewayResult
	err := ctx.Invoke("azure-native:appplatform/v20220501preview:getGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupGatewayArgs struct {
	GatewayName       string `pulumi:"gatewayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupGatewayResult struct {
	Id         string                    `pulumi:"id"`
	Name       string                    `pulumi:"name"`
	Properties GatewayPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse              `pulumi:"sku"`
	SystemData SystemDataResponse        `pulumi:"systemData"`
	Type       string                    `pulumi:"type"`
}


func (val *LookupGatewayResult) Defaults() *LookupGatewayResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}

func LookupGatewayOutput(ctx *pulumi.Context, args LookupGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayResult, error) {
			args := v.(LookupGatewayArgs)
			r, err := LookupGateway(ctx, &args, opts...)
			var s LookupGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayResultOutput)
}

type LookupGatewayOutputArgs struct {
	GatewayName       pulumi.StringInput `pulumi:"gatewayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayArgs)(nil)).Elem()
}


type LookupGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayResult)(nil)).Elem()
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutput() LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutputWithContext(ctx context.Context) LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) Properties() GatewayPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGatewayResult) GatewayPropertiesResponse { return v.Properties }).(GatewayPropertiesResponseOutput)
}

func (o LookupGatewayResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupGatewayResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupGatewayResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGatewayResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayResultOutput{})
}
