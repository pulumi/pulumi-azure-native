


package m365securityandcompliance

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionsForMIPPolicySync(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionsForMIPPolicySyncArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionsForMIPPolicySyncResult, error) {
	var rv LookupPrivateEndpointConnectionsForMIPPolicySyncResult
	err := ctx.Invoke("azure-native:m365securityandcompliance:getPrivateEndpointConnectionsForMIPPolicySync", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionsForMIPPolicySyncArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupPrivateEndpointConnectionsForMIPPolicySyncResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}
