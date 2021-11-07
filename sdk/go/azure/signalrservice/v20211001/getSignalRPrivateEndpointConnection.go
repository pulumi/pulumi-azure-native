


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSignalRPrivateEndpointConnection(ctx *pulumi.Context, args *LookupSignalRPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupSignalRPrivateEndpointConnectionResult, error) {
	var rv LookupSignalRPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:signalrservice/v20211001:getSignalRPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalRPrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupSignalRPrivateEndpointConnectionResult struct {
	GroupIds                          []string                                   `pulumi:"groupIds"`
	Id                                string                                     `pulumi:"id"`
	Name                              string                                     `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                         `pulumi:"systemData"`
	Type                              string                                     `pulumi:"type"`
}
