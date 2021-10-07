


package v20191201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualHubRouteTableV2(ctx *pulumi.Context, args *LookupVirtualHubRouteTableV2Args, opts ...pulumi.InvokeOption) (*LookupVirtualHubRouteTableV2Result, error) {
	var rv LookupVirtualHubRouteTableV2Result
	err := ctx.Invoke("azure-native:network/v20191201:getVirtualHubRouteTableV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubRouteTableV2Args struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RouteTableName    string `pulumi:"routeTableName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupVirtualHubRouteTableV2Result struct {
	AttachedConnections []string                    `pulumi:"attachedConnections"`
	Etag                string                      `pulumi:"etag"`
	Id                  *string                     `pulumi:"id"`
	Name                *string                     `pulumi:"name"`
	ProvisioningState   string                      `pulumi:"provisioningState"`
	Routes              []VirtualHubRouteV2Response `pulumi:"routes"`
}
