


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateLinkHub(ctx *pulumi.Context, args *LookupPrivateLinkHubArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkHubResult, error) {
	var rv LookupPrivateLinkHubResult
	err := ctx.Invoke("azure-native:synapse/v20210601:getPrivateLinkHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkHubArgs struct {
	PrivateLinkHubName string `pulumi:"privateLinkHubName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupPrivateLinkHubResult struct {
	Id                         string                                                    `pulumi:"id"`
	Location                   string                                                    `pulumi:"location"`
	Name                       string                                                    `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionForPrivateLinkHubBasicResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          *string                                                   `pulumi:"provisioningState"`
	Tags                       map[string]string                                         `pulumi:"tags"`
	Type                       string                                                    `pulumi:"type"`
}
