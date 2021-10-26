


package v20210827

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionResult, error) {
	var rv LookupPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20210827:getPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionArgs struct {
	ClusterName                   string `pulumi:"clusterName"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointConnectionResult struct {
	GroupId                           string                                            `pulumi:"groupId"`
	Id                                string                                            `pulumi:"id"`
	Name                              string                                            `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                            `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                                `pulumi:"systemData"`
	Type                              string                                            `pulumi:"type"`
}
