


package v20191017preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateLinkScopedResource(ctx *pulumi.Context, args *LookupPrivateLinkScopedResourceArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkScopedResourceResult, error) {
	var rv LookupPrivateLinkScopedResourceResult
	err := ctx.Invoke("azure-native:insights/v20191017preview:getPrivateLinkScopedResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkScopedResourceArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScopeName         string `pulumi:"scopeName"`
}


type LookupPrivateLinkScopedResourceResult struct {
	Id                string  `pulumi:"id"`
	LinkedResourceId  *string `pulumi:"linkedResourceId"`
	Name              string  `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Type              string  `pulumi:"type"`
}
