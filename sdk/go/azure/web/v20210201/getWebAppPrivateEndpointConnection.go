


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppPrivateEndpointConnection(ctx *pulumi.Context, args *LookupWebAppPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPrivateEndpointConnectionResult, error) {
	var rv LookupWebAppPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:web/v20210201:getWebAppPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPrivateEndpointConnectionArgs struct {
	Name                          string `pulumi:"name"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupWebAppPrivateEndpointConnectionResult struct {
	Id                                string                              `pulumi:"id"`
	IpAddresses                       []string                            `pulumi:"ipAddresses"`
	Kind                              *string                             `pulumi:"kind"`
	Name                              string                              `pulumi:"name"`
	PrivateEndpoint                   *ArmIdWrapperResponse               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
	Type                              string                              `pulumi:"type"`
}
