


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRouteConnection(ctx *pulumi.Context, args *LookupExpressRouteConnectionArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteConnectionResult, error) {
	var rv LookupExpressRouteConnectionResult
	err := ctx.Invoke("azure-native:network/v20210301:getExpressRouteConnection", args, &rv, opts...)
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
	ExpressRouteGatewayBypass  *bool                                `pulumi:"expressRouteGatewayBypass"`
	Id                         *string                              `pulumi:"id"`
	Name                       string                               `pulumi:"name"`
	ProvisioningState          string                               `pulumi:"provisioningState"`
	RoutingConfiguration       *RoutingConfigurationResponse        `pulumi:"routingConfiguration"`
	RoutingWeight              *int                                 `pulumi:"routingWeight"`
}
