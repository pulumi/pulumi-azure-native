


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHubRouteTable(ctx *pulumi.Context, args *LookupHubRouteTableArgs, opts ...pulumi.InvokeOption) (*LookupHubRouteTableResult, error) {
	var rv LookupHubRouteTableResult
	err := ctx.Invoke("azure-native:network/v20200601:getHubRouteTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHubRouteTableArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RouteTableName    string `pulumi:"routeTableName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupHubRouteTableResult struct {
	AssociatedConnections  []string           `pulumi:"associatedConnections"`
	Etag                   string             `pulumi:"etag"`
	Id                     *string            `pulumi:"id"`
	Labels                 []string           `pulumi:"labels"`
	Name                   *string            `pulumi:"name"`
	PropagatingConnections []string           `pulumi:"propagatingConnections"`
	ProvisioningState      string             `pulumi:"provisioningState"`
	Routes                 []HubRouteResponse `pulumi:"routes"`
	Type                   string             `pulumi:"type"`
}
