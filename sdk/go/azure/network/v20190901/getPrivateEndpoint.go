


package v20190901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpoint(ctx *pulumi.Context, args *LookupPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointResult, error) {
	var rv LookupPrivateEndpointResult
	err := ctx.Invoke("azure-native:network/v20190901:getPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointArgs struct {
	Expand              *string `pulumi:"expand"`
	PrivateEndpointName string  `pulumi:"privateEndpointName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointResult struct {
	Etag                                string                                 `pulumi:"etag"`
	Id                                  *string                                `pulumi:"id"`
	Location                            *string                                `pulumi:"location"`
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnectionResponse `pulumi:"manualPrivateLinkServiceConnections"`
	Name                                string                                 `pulumi:"name"`
	NetworkInterfaces                   []NetworkInterfaceResponse             `pulumi:"networkInterfaces"`
	PrivateLinkServiceConnections       []PrivateLinkServiceConnectionResponse `pulumi:"privateLinkServiceConnections"`
	ProvisioningState                   string                                 `pulumi:"provisioningState"`
	Subnet                              *SubnetResponse                        `pulumi:"subnet"`
	Tags                                map[string]string                      `pulumi:"tags"`
	Type                                string                                 `pulumi:"type"`
}
