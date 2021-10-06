


package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateLinkServicePrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateLinkServicePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkServicePrivateEndpointConnectionResult, error) {
	var rv LookupPrivateLinkServicePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:network/v20200801:getPrivateLinkServicePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkServicePrivateEndpointConnectionArgs struct {
	Expand            *string `pulumi:"expand"`
	PeConnectionName  string  `pulumi:"peConnectionName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type LookupPrivateLinkServicePrivateEndpointConnectionResult struct {
	Etag                              string                                     `pulumi:"etag"`
	Id                                *string                                    `pulumi:"id"`
	LinkIdentifier                    string                                     `pulumi:"linkIdentifier"`
	Name                              *string                                    `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponse                    `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	Type                              string                                     `pulumi:"type"`
}
