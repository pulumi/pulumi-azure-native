


package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionByName(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionByNameArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionByNameResult, error) {
	var rv LookupPrivateEndpointConnectionByNameResult
	err := ctx.Invoke("azure-native:apimanagement:getPrivateEndpointConnectionByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionByNameArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ServiceName                   string `pulumi:"serviceName"`
}


type LookupPrivateEndpointConnectionByNameResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}
