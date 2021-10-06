


package compute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiskAccessAPrivateEndpointConnection(ctx *pulumi.Context, args *LookupDiskAccessAPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupDiskAccessAPrivateEndpointConnectionResult, error) {
	var rv LookupDiskAccessAPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:compute:getDiskAccessAPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskAccessAPrivateEndpointConnectionArgs struct {
	DiskAccessName                string `pulumi:"diskAccessName"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupDiskAccessAPrivateEndpointConnectionResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}
