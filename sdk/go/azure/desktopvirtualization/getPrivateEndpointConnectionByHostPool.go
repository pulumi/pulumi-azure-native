


package desktopvirtualization

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionByHostPool(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionByHostPoolArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionByHostPoolResult, error) {
	var rv LookupPrivateEndpointConnectionByHostPoolResult
	err := ctx.Invoke("azure-native:desktopvirtualization:getPrivateEndpointConnectionByHostPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionByHostPoolArgs struct {
	HostPoolName                  string `pulumi:"hostPoolName"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointConnectionByHostPoolResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}
