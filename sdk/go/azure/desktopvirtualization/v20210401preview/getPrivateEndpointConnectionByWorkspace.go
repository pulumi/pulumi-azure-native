


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionByWorkspace(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionByWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionByWorkspaceResult, error) {
	var rv LookupPrivateEndpointConnectionByWorkspaceResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20210401preview:getPrivateEndpointConnectionByWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionByWorkspaceArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	WorkspaceName                 string `pulumi:"workspaceName"`
}


type LookupPrivateEndpointConnectionByWorkspaceResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}
