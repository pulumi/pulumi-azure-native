


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRouteCircuitAuthorization(ctx *pulumi.Context, args *LookupExpressRouteCircuitAuthorizationArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitAuthorizationResult, error) {
	var rv LookupExpressRouteCircuitAuthorizationResult
	err := ctx.Invoke("azure-native:network/v20210201:getExpressRouteCircuitAuthorization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitAuthorizationArgs struct {
	AuthorizationName string `pulumi:"authorizationName"`
	CircuitName       string `pulumi:"circuitName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupExpressRouteCircuitAuthorizationResult struct {
	AuthorizationKey       *string `pulumi:"authorizationKey"`
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	Etag                   string  `pulumi:"etag"`
	Id                     *string `pulumi:"id"`
	Name                   *string `pulumi:"name"`
	ProvisioningState      string  `pulumi:"provisioningState"`
	Type                   string  `pulumi:"type"`
}
