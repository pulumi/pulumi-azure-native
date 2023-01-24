


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayRouteConfig(ctx *pulumi.Context, args *LookupGatewayRouteConfigArgs, opts ...pulumi.InvokeOption) (*LookupGatewayRouteConfigResult, error) {
	var rv LookupGatewayRouteConfigResult
	err := ctx.Invoke("azure-native:appplatform/v20221101preview:getGatewayRouteConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupGatewayRouteConfigArgs struct {
	GatewayName       string `pulumi:"gatewayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RouteConfigName   string `pulumi:"routeConfigName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupGatewayRouteConfigResult struct {
	Id         string                               `pulumi:"id"`
	Name       string                               `pulumi:"name"`
	Properties GatewayRouteConfigPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                   `pulumi:"systemData"`
	Type       string                               `pulumi:"type"`
}


func (val *LookupGatewayRouteConfigResult) Defaults() *LookupGatewayRouteConfigResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupGatewayRouteConfigOutput(ctx *pulumi.Context, args LookupGatewayRouteConfigOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayRouteConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayRouteConfigResult, error) {
			args := v.(LookupGatewayRouteConfigArgs)
			r, err := LookupGatewayRouteConfig(ctx, &args, opts...)
			var s LookupGatewayRouteConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayRouteConfigResultOutput)
}

type LookupGatewayRouteConfigOutputArgs struct {
	GatewayName       pulumi.StringInput `pulumi:"gatewayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RouteConfigName   pulumi.StringInput `pulumi:"routeConfigName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGatewayRouteConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayRouteConfigArgs)(nil)).Elem()
}


type LookupGatewayRouteConfigResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayRouteConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayRouteConfigResult)(nil)).Elem()
}

func (o LookupGatewayRouteConfigResultOutput) ToLookupGatewayRouteConfigResultOutput() LookupGatewayRouteConfigResultOutput {
	return o
}

func (o LookupGatewayRouteConfigResultOutput) ToLookupGatewayRouteConfigResultOutputWithContext(ctx context.Context) LookupGatewayRouteConfigResultOutput {
	return o
}

func (o LookupGatewayRouteConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayRouteConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayRouteConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayRouteConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGatewayRouteConfigResultOutput) Properties() GatewayRouteConfigPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGatewayRouteConfigResult) GatewayRouteConfigPropertiesResponse { return v.Properties }).(GatewayRouteConfigPropertiesResponseOutput)
}

func (o LookupGatewayRouteConfigResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGatewayRouteConfigResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupGatewayRouteConfigResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayRouteConfigResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayRouteConfigResultOutput{})
}
