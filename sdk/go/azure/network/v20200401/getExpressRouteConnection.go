


package v20200401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRouteConnection(ctx *pulumi.Context, args *LookupExpressRouteConnectionArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteConnectionResult, error) {
	var rv LookupExpressRouteConnectionResult
	err := ctx.Invoke("azure-native:network/v20200401:getExpressRouteConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteConnectionArgs struct {
	ConnectionName          string `pulumi:"connectionName"`
	ExpressRouteGatewayName string `pulumi:"expressRouteGatewayName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupExpressRouteConnectionResult struct {
	AuthorizationKey           *string                              `pulumi:"authorizationKey"`
	EnableInternetSecurity     *bool                                `pulumi:"enableInternetSecurity"`
	ExpressRouteCircuitPeering ExpressRouteCircuitPeeringIdResponse `pulumi:"expressRouteCircuitPeering"`
	Id                         *string                              `pulumi:"id"`
	Name                       string                               `pulumi:"name"`
	ProvisioningState          string                               `pulumi:"provisioningState"`
	RoutingConfiguration       *RoutingConfigurationResponse        `pulumi:"routingConfiguration"`
	RoutingWeight              *int                                 `pulumi:"routingWeight"`
}

func LookupExpressRouteConnectionOutput(ctx *pulumi.Context, args LookupExpressRouteConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteConnectionResult, error) {
			args := v.(LookupExpressRouteConnectionArgs)
			r, err := LookupExpressRouteConnection(ctx, &args, opts...)
			var s LookupExpressRouteConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRouteConnectionResultOutput)
}

type LookupExpressRouteConnectionOutputArgs struct {
	ConnectionName          pulumi.StringInput `pulumi:"connectionName"`
	ExpressRouteGatewayName pulumi.StringInput `pulumi:"expressRouteGatewayName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteConnectionArgs)(nil)).Elem()
}


type LookupExpressRouteConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteConnectionResult)(nil)).Elem()
}

func (o LookupExpressRouteConnectionResultOutput) ToLookupExpressRouteConnectionResultOutput() LookupExpressRouteConnectionResultOutput {
	return o
}

func (o LookupExpressRouteConnectionResultOutput) ToLookupExpressRouteConnectionResultOutputWithContext(ctx context.Context) LookupExpressRouteConnectionResultOutput {
	return o
}

func (o LookupExpressRouteConnectionResultOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteConnectionResult) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteConnectionResultOutput) EnableInternetSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteConnectionResult) *bool { return v.EnableInternetSecurity }).(pulumi.BoolPtrOutput)
}

func (o LookupExpressRouteConnectionResultOutput) ExpressRouteCircuitPeering() ExpressRouteCircuitPeeringIdResponseOutput {
	return o.ApplyT(func(v LookupExpressRouteConnectionResult) ExpressRouteCircuitPeeringIdResponse {
		return v.ExpressRouteCircuitPeering
	}).(ExpressRouteCircuitPeeringIdResponseOutput)
}

func (o LookupExpressRouteConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExpressRouteConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupExpressRouteConnectionResultOutput) RoutingConfiguration() RoutingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteConnectionResult) *RoutingConfigurationResponse {
		return v.RoutingConfiguration
	}).(RoutingConfigurationResponsePtrOutput)
}

func (o LookupExpressRouteConnectionResultOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteConnectionResult) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteConnectionResultOutput{})
}
