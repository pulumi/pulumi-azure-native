


package v20211201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAuthorization(ctx *pulumi.Context, args *LookupAuthorizationArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationResult, error) {
	var rv LookupAuthorizationResult
	err := ctx.Invoke("azure-native:avs/v20211201:getAuthorization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizationArgs struct {
	AuthorizationName string `pulumi:"authorizationName"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAuthorizationResult struct {
	ExpressRouteAuthorizationId  string  `pulumi:"expressRouteAuthorizationId"`
	ExpressRouteAuthorizationKey string  `pulumi:"expressRouteAuthorizationKey"`
	ExpressRouteId               *string `pulumi:"expressRouteId"`
	Id                           string  `pulumi:"id"`
	Name                         string  `pulumi:"name"`
	ProvisioningState            string  `pulumi:"provisioningState"`
	Type                         string  `pulumi:"type"`
}
