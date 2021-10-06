


package v20210422preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateLinkScope(ctx *pulumi.Context, args *LookupPrivateLinkScopeArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkScopeResult, error) {
	var rv LookupPrivateLinkScopeResult
	err := ctx.Invoke("azure-native:hybridcompute/v20210422preview:getPrivateLinkScope", args, &rv, opts...)
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
	Id         string                                          `pulumi:"id"`
	Location   string                                          `pulumi:"location"`
	Name       string                                          `pulumi:"name"`
	Properties HybridComputePrivateLinkScopePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                              `pulumi:"systemData"`
	Tags       map[string]string                               `pulumi:"tags"`
	Type       string                                          `pulumi:"type"`
}
