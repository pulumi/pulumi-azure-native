


package v20210325preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionsForSCCPowershell(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionsForSCCPowershellArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionsForSCCPowershellResult, error) {
	var rv LookupPrivateEndpointConnectionsForSCCPowershellResult
	err := ctx.Invoke("azure-native:m365securityandcompliance/v20210325preview:getPrivateEndpointConnectionsForSCCPowershell", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionsForSCCPowershellArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupPrivateEndpointConnectionsForSCCPowershellResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}
