


package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSharedPrivateLinkResource(ctx *pulumi.Context, args *LookupSharedPrivateLinkResourceArgs, opts ...pulumi.InvokeOption) (*LookupSharedPrivateLinkResourceResult, error) {
	var rv LookupSharedPrivateLinkResourceResult
	err := ctx.Invoke("azure-native:search/v20200801:getSharedPrivateLinkResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSharedPrivateLinkResourceArgs struct {
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	SearchServiceName             string `pulumi:"searchServiceName"`
	SharedPrivateLinkResourceName string `pulumi:"sharedPrivateLinkResourceName"`
}


type LookupSharedPrivateLinkResourceResult struct {
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties SharedPrivateLinkResourcePropertiesResponse `pulumi:"properties"`
	Type       string                                      `pulumi:"type"`
}
