


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpoint(ctx *pulumi.Context, args *LookupPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointResult, error) {
	var rv LookupPrivateEndpointResult
	err := ctx.Invoke("azure-native:streamanalytics/v20200301:getPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointArgs struct {
	ClusterName         string `pulumi:"clusterName"`
	PrivateEndpointName string `pulumi:"privateEndpointName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointResult struct {
	CreatedDate                         string                                 `pulumi:"createdDate"`
	Etag                                string                                 `pulumi:"etag"`
	Id                                  string                                 `pulumi:"id"`
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnectionResponse `pulumi:"manualPrivateLinkServiceConnections"`
	Name                                string                                 `pulumi:"name"`
	Type                                string                                 `pulumi:"type"`
}
