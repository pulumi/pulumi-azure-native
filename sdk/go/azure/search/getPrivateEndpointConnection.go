


package search

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionResult, error) {
	var rv LookupPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:search:getPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	SearchServiceName             string `pulumi:"searchServiceName"`
}


type LookupPrivateEndpointConnectionResult struct {
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	Type       string                                      `pulumi:"type"`
}


func (val *LookupPrivateEndpointConnectionResult) Defaults() *LookupPrivateEndpointConnectionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
