


package v20161201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRouteTable(ctx *pulumi.Context, args *LookupRouteTableArgs, opts ...pulumi.InvokeOption) (*LookupRouteTableResult, error) {
	var rv LookupRouteTableResult
	err := ctx.Invoke("azure-native:network/v20161201:getRouteTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteTableArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RouteTableName    string  `pulumi:"routeTableName"`
}


type LookupRouteTableResult struct {
	Etag              *string           `pulumi:"etag"`
	Id                *string           `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	Routes            []RouteResponse   `pulumi:"routes"`
	Subnets           []SubnetResponse  `pulumi:"subnets"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}
