


package streamanalytics

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpoint(ctx *pulumi.Context, args *LookupPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointResult, error) {
	var rv LookupPrivateEndpointResult
	err := ctx.Invoke("azure-native:streamanalytics:getPrivateEndpoint", args, &rv, opts...)
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
	Etag       string                            `pulumi:"etag"`
	Id         string                            `pulumi:"id"`
	Name       string                            `pulumi:"name"`
	Properties PrivateEndpointPropertiesResponse `pulumi:"properties"`
	Type       string                            `pulumi:"type"`
}
