


package v20191017preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateLinkScope(ctx *pulumi.Context, args *LookupPrivateLinkScopeArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkScopeResult, error) {
	var rv LookupPrivateLinkScopeResult
	err := ctx.Invoke("azure-native:insights/v20191017preview:getPrivateLinkScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkScopeArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScopeName         string `pulumi:"scopeName"`
}


type LookupPrivateLinkScopeResult struct {
	Id                         string                              `pulumi:"id"`
	Location                   string                              `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
}
