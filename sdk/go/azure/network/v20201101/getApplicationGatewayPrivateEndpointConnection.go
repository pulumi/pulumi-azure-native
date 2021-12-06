


package v20201101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationGatewayPrivateEndpointConnection(ctx *pulumi.Context, args *LookupApplicationGatewayPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGatewayPrivateEndpointConnectionResult, error) {
	var rv LookupApplicationGatewayPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:network/v20201101:getApplicationGatewayPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationGatewayPrivateEndpointConnectionArgs struct {
	ApplicationGatewayName string `pulumi:"applicationGatewayName"`
	ConnectionName         string `pulumi:"connectionName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupApplicationGatewayPrivateEndpointConnectionResult struct {
	Etag                              string                                     `pulumi:"etag"`
	Id                                *string                                    `pulumi:"id"`
	LinkIdentifier                    string                                     `pulumi:"linkIdentifier"`
	Name                              *string                                    `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponse                    `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	Type                              string                                     `pulumi:"type"`
}
