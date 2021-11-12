


package v20190401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRouteGateway(ctx *pulumi.Context, args *LookupExpressRouteGatewayArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteGatewayResult, error) {
	var rv LookupExpressRouteGatewayResult
	err := ctx.Invoke("azure-native:network/v20190401:getExpressRouteGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteGatewayArgs struct {
	ExpressRouteGatewayName string `pulumi:"expressRouteGatewayName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupExpressRouteGatewayResult struct {
	AutoScaleConfiguration  *ExpressRouteGatewayPropertiesResponseAutoScaleConfiguration `pulumi:"autoScaleConfiguration"`
	Etag                    string                                                       `pulumi:"etag"`
	ExpressRouteConnections []ExpressRouteConnectionResponse                             `pulumi:"expressRouteConnections"`
	Id                      *string                                                      `pulumi:"id"`
	Location                *string                                                      `pulumi:"location"`
	Name                    string                                                       `pulumi:"name"`
	ProvisioningState       string                                                       `pulumi:"provisioningState"`
	Tags                    map[string]string                                            `pulumi:"tags"`
	Type                    string                                                       `pulumi:"type"`
	VirtualHub              VirtualHubIdResponse                                         `pulumi:"virtualHub"`
}
