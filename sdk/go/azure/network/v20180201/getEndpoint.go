


package v20180201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	var rv LookupEndpointResult
	err := ctx.Invoke("azure-native:network/v20180201:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	EndpointType      string `pulumi:"endpointType"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEndpointResult struct {
	EndpointLocation      *string  `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string  `pulumi:"endpointMonitorStatus"`
	EndpointStatus        *string  `pulumi:"endpointStatus"`
	GeoMapping            []string `pulumi:"geoMapping"`
	Id                    string   `pulumi:"id"`
	MinChildEndpoints     *float64 `pulumi:"minChildEndpoints"`
	Name                  string   `pulumi:"name"`
	Priority              *float64 `pulumi:"priority"`
	Target                *string  `pulumi:"target"`
	TargetResourceId      *string  `pulumi:"targetResourceId"`
	Type                  string   `pulumi:"type"`
	Weight                *float64 `pulumi:"weight"`
}
