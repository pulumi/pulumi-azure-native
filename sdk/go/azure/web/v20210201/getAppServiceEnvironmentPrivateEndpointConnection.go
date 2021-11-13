


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServiceEnvironmentPrivateEndpointConnection(ctx *pulumi.Context, args *LookupAppServiceEnvironmentPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceEnvironmentPrivateEndpointConnectionResult, error) {
	var rv LookupAppServiceEnvironmentPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:web/v20210201:getAppServiceEnvironmentPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServiceEnvironmentPrivateEndpointConnectionArgs struct {
	Name                          string `pulumi:"name"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupAppServiceEnvironmentPrivateEndpointConnectionResult struct {
	Id                                string                              `pulumi:"id"`
	IpAddresses                       []string                            `pulumi:"ipAddresses"`
	Kind                              *string                             `pulumi:"kind"`
	Name                              string                              `pulumi:"name"`
	PrivateEndpoint                   *ArmIdWrapperResponse               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
	Type                              string                              `pulumi:"type"`
}
