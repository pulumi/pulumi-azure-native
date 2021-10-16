


package v20210308

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionsAdtAPI(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionsAdtAPIArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionsAdtAPIResult, error) {
	var rv LookupPrivateEndpointConnectionsAdtAPIResult
	err := ctx.Invoke("azure-native:securityandcompliance/v20210308:getPrivateEndpointConnectionsAdtAPI", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionsAdtAPIArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupPrivateEndpointConnectionsAdtAPIResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}
