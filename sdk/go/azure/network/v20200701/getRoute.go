


package v20200701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoute(ctx *pulumi.Context, args *LookupRouteArgs, opts ...pulumi.InvokeOption) (*LookupRouteResult, error) {
	var rv LookupRouteResult
	err := ctx.Invoke("azure-native:network/v20200701:getRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RouteName         string `pulumi:"routeName"`
	RouteTableName    string `pulumi:"routeTableName"`
}


type LookupRouteResult struct {
	AddressPrefix     *string `pulumi:"addressPrefix"`
	Etag              string  `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	NextHopIpAddress  *string `pulumi:"nextHopIpAddress"`
	NextHopType       string  `pulumi:"nextHopType"`
	ProvisioningState string  `pulumi:"provisioningState"`
}
